//go:generate fileb0x assets.toml

package main

import (
	"bytes"
	"errors"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

const (
	stdLibRoot = "https://godoc.org/github.com/gen0cide/gscript/stdlib"
)

type functionDef struct {
	Name    string
	Sig     string
	Doc     string
	Example string
}

type gDocPage struct {
	PackageName  string
	FunctionDefs []functionDef
}

func checkFatalErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func getParsedHTMLDoc(url string) (*html.Node, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	return html.Parse(resp.Body)
}

func makeGDocsFromURL(url string) {

	// init
	innerHTMLBuffer := bytes.NewBuffer([]byte{})
	gDocPageObj := gDocPage{}

	// retrive godoc page
	doc, err := getParsedHTMLDoc(url)
	if err != nil {
		log.Println(err)
		return
	}

	// make extract data for making a GDoc file
	var parseGoDoc func(n *html.Node)
	parseGoDoc = func(n *html.Node) {
		if n.Type == html.ElementNode {
			switch n.Data {
			case "h2":
				for _, attr := range n.Attr {
					if attr.Key == "id" {
						if attr.Val == "pkg-overview" {
							if err := html.Render(innerHTMLBuffer, n); err != nil {
								log.Println("error parsing element with id: pkg-overview!")
								checkFatalErr(err)
							}
							gDocPageObj.PackageName = innerHTMLBuffer.String()[30:]
							innerHTMLBuffer.Reset()
							gDocPageObj.PackageName = gDocPageObj.PackageName[:len(gDocPageObj.PackageName)-5]
						}
					}
				}
			case "div":
				for _, attr := range n.Attr {
					if attr.Key == "class" {
						if attr.Val == "funcdecl decl" {
							funcDeclPre := n.FirstChild.NextSibling
							if funcDeclPre == nil {
								checkFatalErr(errors.New("error, could not locate <pre> tag containing function decleration after finding 'funcdecl decl' element"))
							}
							if err := html.Render(innerHTMLBuffer, funcDeclPre); err != nil {
								log.Println("error parsing function decleration <pre> element")
								checkFatalErr(err)
							}
							funcSig := innerHTMLBuffer.String()
							innerHTMLBuffer.Reset()

							// get GoDoc comment
							funcDesc := n.NextSibling
							var funcComment string
							if funcDesc != nil {
								if funcDesc.Type == html.ElementNode {
									if funcDesc.Data == "p" {
										if err := html.Render(innerHTMLBuffer, funcDesc); err != nil {
											log.Println("error parsing function comment")
											checkFatalErr(err)
										}
										funcComment = innerHTMLBuffer.String()
										innerHTMLBuffer.Reset()
										funcComment = html.UnescapeString(funcComment[3 : len(funcComment)-4])
									}
								}
							}

							// signature  HTML cleanup
							funcSig = strings.Replace(funcSig, "<pre>", "", -1)
							funcSig = strings.Replace(funcSig, "</pre>", "", -1)
							funcSig = strings.Replace(funcSig, "</a>", "", -1)
							for {
								aTagStart := strings.Index(funcSig, `<a href="`)
								if aTagStart == -1 {
									break
								}
								funcSig = funcSig[:aTagStart] + funcSig[strings.Index(funcSig, `">`)+2:]
							}

							// get func name from signature
							var name string
							if strings.HasPrefix(funcSig, "func (") {
								name = funcSig[strings.Index(funcSig, ") "):]
							} else {
								name = funcSig[5:strings.Index(funcSig, "(")]
							}

							// output
							gDocPageObj.FunctionDefs = append(gDocPageObj.FunctionDefs, functionDef{
								Name: name,
								Sig:  funcSig,
								Doc:  funcComment,
							})
						}
					}
				}
				// TODO: Examples
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			parseGoDoc(c)
		}
	}
	parseGoDoc(doc)

	// make GDoc page
	gdocTemplate, err := ReadFile("docs.gotmpl")
	checkFatalErr(err)
	gDocTemplateInstance := template.Must(template.New("doc").Parse(string(gdocTemplate)))
	gDoc := bytes.NewBuffer([]byte{})
	gDocTemplateInstance.Execute(gDoc, gDocPageObj)
	ioutil.WriteFile("stdlib/"+gDocPageObj.PackageName+".md", gDoc.Bytes(), 0644)
}

func main() {
	// get the root godoc page for GScript
	doc, err := getParsedHTMLDoc(stdLibRoot[:len(stdLibRoot)-7])
	checkFatalErr(err)

	// loop through all the HTML nodes
	var findPkgList func(n *html.Node)
	findPkgList = func(n *html.Node) {
		if n.Type == html.ElementNode {

			// for each <a> tag
			if n.Data == "a" {

				// find "href" attribute
				for _, attr := range n.Attr {
					if attr.Key == "href" {

						// if this a tag's href attribute points to a stdlib package...
						if strings.Contains(attr.Val, strings.Split(stdLibRoot, "/")[strings.Count(stdLibRoot, "/")]) {
							makeGDocsFromURL(stdLibRoot + attr.Val[strings.LastIndex(attr.Val, "/"):])
						}
					}
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			findPkgList(c)
		}
	}
	findPkgList(doc)
}
