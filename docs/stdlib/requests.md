
# Package: requests

## Functions

### GetURLAsBytes

GetURLAsBytes will fetch a url, headers, and a bool for ignoring ssl errors. this returns an http response object, the body as a string, and an error

``````
func GetURLAsBytes(url string, headers map[string]interface{}, ignoresslerrors bool) (*http.Response, []byte, error)
``````
#### Example
``````

``````

### GetURLAsString

GetURLAsString will fetch a url and return an http response object, the body as a string, and an error

``````
func GetURLAsString(url string, headers map[string]interface{}, ignoresslerrors bool) (*http.Response, string, error)
``````
#### Example
``````

``````

### PostBinary

PostBinary posts the specified data to a url endpoint as application/octet-stream data

``````
func PostBinary(url string, readPath string, headers map[string]interface{}, ignoresslerrors bool) (*http.Response, string, error)
``````
#### Example
``````

``````

### PostJSON

PostJSON takes a url, json data, a map of headers, and a bool to ignore ssl errors, posts json data to url

``````
func PostJSON(url string, data map[string]interface{}, headers map[string]interface{}, ignoresslerrors bool) (*http.Response, string, error)
``````
#### Example
``````

``````

### PostURL

PostURL posts the specified data to a url endpoint as text/plain data

``````
func PostURL(url string, data string, headers map[string]interface{}, ignoresslerrors bool) (*http.Response, string, error)
``````
#### Example
``````

``````
