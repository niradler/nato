package nato

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

type LoopArgs struct {
	Input     string
	Pattern   string
	Separator string
	Command   string
	DryRun    bool
}

type Variables struct {
	I     int
	Index int
	V     string
	Value string
}

func Loop(args LoopArgs) (re bool) {

	var inputSlices []string
	switch args.Pattern {
	case "split":
		inputSlices = strings.Split(args.Input, args.Separator)
	default:
		inputSlices = strings.Fields(args.Input)
	}
	for i, item := range inputSlices {
		value := strings.TrimSuffix(item, "\r\n")
		variables := Variables{
			i, i, value, value,
		}

		t, err := template.New("raw").Parse(args.Command)
		if err != nil {
			panic(err)
		}
		var tpl bytes.Buffer
		if err := t.Execute(&tpl, variables); err != nil {
			panic(err)
		}

		compiled := tpl.String()
		if args.DryRun {
			fmt.Println(compiled)
		} else {
			commandArgs := strings.Fields(compiled)
			var shell string
			var shellArgs []string
			if runtime.GOOS == "windows" {
				shell = "cmd"
				shellArgs = []string{"/C"}
			} else {
				shell = "bash"
				shellArgs = []string{"-c"}
			}

			cmd := exec.Command(shell, append(shellArgs, commandArgs...)...)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			err := cmd.Run()
			if err != nil {
				panic(err)
			}
		}

	}

	return true
}

func GetStdin() string {
	data := ""
	fi, _ := os.Stdin.Stat() // get the FileInfo struct describing the standard input.
	if (fi.Mode() & os.ModeCharDevice) == 0 {
		bytes, _ := io.ReadAll(os.Stdin)
		str := string(bytes)
		data = str
	}

	return data
}
