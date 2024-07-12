package internal

import (
	"os"

	"github.com/gookit/color"
)

func Validate() ([]byte, string) {
	var error string

	//	if len(os.Args) != 2 {
	//		error = color.Red.Text("Error: please select a file to parse")
	//		return nil, error
	//	}
	//
	//	args := os.Args[1]

	path, err := os.Getwd()
	if err != nil {
		error = color.Red.Text("Error: unable to find current working directory")
		return nil, error
	}

	filename := ListFiles()

	f, err := os.ReadFile(path + "/" + filename)
	if err != nil {
		error = color.Red.Text("Error: unable to find current working directory")
		return nil, error
	}

	return f, error
}
