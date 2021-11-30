package input

import (
	"errors"
	"io/ioutil"
	"os"
	"path"
	"runtime"
	"strings"
)

func GetInputChars() []string {
	return cleanEmpty(strings.Split(inputChain(), ""))
}

func GetInputString() string {
	return inputChain()
}

func GetInputStrings() []string {
	return cleanEmpty(strings.Split(inputChain(), "\n"))
}

// We usually want the input from a file called "input", but its useful
// to load alternative data from the CLI.
// Input chain:
// - stdin
// - "input" file in the same directory as the caller
func inputChain() string {
	stdin, err := inputFromStdIn()
	if err == nil {
		return stdin
	}

	_, callerFilename, _, _ := runtime.Caller(2)
	file, err := inputFileIn(callerFilename)
	if err == nil {
		return file
	}

	panic("No input at all")
}

func inputFromStdIn() (string, error) {
	stat, _ := os.Stdin.Stat()
	dataPiped := (stat.Mode() & os.ModeCharDevice) == 0

	if dataPiped {
		bytes, err := ioutil.ReadAll(os.Stdin)
		if len(bytes) > 0 {
			return string(bytes), err
		}
	}

	return "", errors.New("No STDIN")
}

func inputFileIn(callerFilename string) (string, error) {
	inputFilename := path.Join(path.Dir(callerFilename), "input")
	bytes, err := ioutil.ReadFile(inputFilename)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

func cleanEmpty(s []string) []string {
	var r []string
	for _, str := range s {
		if str != "" {
			r = append(r, str)
		}
	}
	return r
}
