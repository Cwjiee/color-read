package internal

import (
	"fmt"
	"os"

	"github.com/gookit/color"
)

func RenderColor() {

	f, err := Validate()
	if err != "" {
		fmt.Println(err)
		os.Exit(1)
	}

	hexColors := ReadColor(f)

	if len(hexColors) > 0 {
		for _, c := range hexColors {
			color.Hex(c, true).Print("#" + c)
			color.Hex(c, false).Println(" #" + c)
		}
	} else {
		fmt.Println("There are no colors found in this file")
	}
}
