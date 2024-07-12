package internal

import (
	"os"

	"github.com/gookit/color"
)

func Validate() ([]byte, string) {
	var error string
	var filename string

	path, err := os.Getwd()
	if err != nil {
		error = color.Red.Text("Error: unable to find current working directory")
		return nil, error
	}

	if len(os.Args) == 2 {
		filename = os.Args[1]
	} else if len(os.Args) == 1 {
		filename = ListFiles()
	} else {
		error = color.Red.Text("Error: please select a file to parse")
		return nil, error
	}

	f, err := os.ReadFile(path + "/" + filename)
	if err != nil {
		error = color.Red.Text("Error: unable to find current working directory")
		return nil, error
	}

	return f, error
}
