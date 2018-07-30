package debugger

import (
	"fmt"
	"io"
	"os"
	"os/user"
	"path/filepath"

	"github.com/fatih/color"
	"github.com/gen0cide/gscript/logger/standard"
	"github.com/gohxs/readline"
	isatty "github.com/mattn/go-isatty"
	"github.com/robertkrimen/otto"
)

func (d *Debugger) runDebugger() error {
	prompt := fmt.Sprintf("%s%s", color.HiRedString("gscript"), color.HiWhiteString("> "))

	var err error

	interactive := isatty.IsTerminal(os.Stdout.Fd()) && isatty.IsTerminal(os.Stdin.Fd())
	cygwin := isatty.IsCygwinTerminal(os.Stdout.Fd()) && isatty.IsCygwinTerminal(os.Stdin.Fd())

	var closers []func() error

	var stdin io.ReadCloser
	if cygwin {
		stdin = os.Stdin
	} else {
		stdin = readline.Stdin
	}

	var histfile string
	cu, err := user.Current()
	if err != nil {
		return err
	}
	histfile = filepath.Join(cu.HomeDir, ".gscript_history")

	// configure stdout
	var stdout io.Writer

	if cygwin {
		stdout = color.Output
	} else {
		stdout = readline.Stdout
	}

	// configure stderr
	var stderr io.Writer = os.Stderr
	if !cygwin {
		stderr = readline.Stderr
	}

	if interactive {
		stdin = readline.NewCancelableStdin(stdin)
	}

	// create readline instance
	rl, err := readline.NewEx(&readline.Config{
		HistoryFile:            histfile,
		DisableAutoSaveHistory: true,
		InterruptPrompt:        "^C",
		HistorySearchFold:      true,
		Stdin:                  stdin,
		Stdout:                 stdout,
		Stderr:                 stderr,
		Prompt:                 prompt,
		FuncIsTerminal: func() bool {
			return interactive || cygwin
		},
		FuncFilterInputRune: func(r rune) (rune, bool) {
			if r == readline.CharCtrlZ {
				return r, false
			}
			return r, true
		},
	})
	if err != nil {
		return err
	}

	closers = append(closers, rl.Close)

	standard.PrintLogo()
	title := fmt.Sprintf(
		"%s %s %s %s",
		color.HiWhiteString("***"),
		color.HiRedString("GSCRIPT"),
		color.YellowString("INTERACTIVE SHELL"),
		color.HiWhiteString("***"),
	)
	fmt.Fprintf(color.Output, "%s\n", title)
	rl.Refresh()

	for {
		l, err := rl.Readline()
		if err != nil {
			if err == readline.ErrInterrupt {
				if d != nil {
					d = nil
					rl.SetPrompt(prompt)
					rl.Refresh()
					continue
				}
				break
			}
			return err
		}
		if l == "" {
			continue
		}
		if l == "exit" {
			break
		}
		if l == "halt" {
			fmt.Println("")
			os.Exit(0)
		}
		s, err := d.VM.VM.Compile("debugger", l)
		if err != nil {
			d.Logger.Errorf("%v", err)
			rl.SetPrompt(prompt)
			rl.Refresh()
			continue
		}
		v, err := d.VM.VM.Eval(s)
		if err != nil {
			if oerr, ok := err.(*otto.Error); ok {
				d.Logger.Error(oerr.Error())
			} else {
				d.Logger.Error(err.Error())
			}
		} else {
			rl.Write([]byte(fmt.Sprintf(">>> %s\n", v.String())))
		}
		rl.Refresh()
	}

	return rl.Close()
}
