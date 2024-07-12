package internal

import (
	"fmt"
	"os"

	"github.com/ktr0731/go-fuzzyfinder"
)

func ListFiles() string {
	f, err := os.ReadDir(".")
	if err != nil {
		fmt.Println(err)
	}

	var files []string

	for _, file := range f {
		files = append(files, file.Name())
	}

	idx, err := fuzzyfinder.Find(
		files,
		func(i int) string {
			return files[i]
		})

	if err != nil {
		fmt.Println(err)
	}

	return files[idx]
}
