package internal

import (
	"os"

	"github.com/gookit/color"
)

func RenderColor() {

	f, err := Validate()
	if err != "" {
		color.Red.Println(err)
		os.Exit(1)
	}

	hexColors := ReadColor(f)

	if len(hexColors) > 0 {
		for _, c := range hexColors {
			color.Hex(c, true).Print("#" + c)
			color.Hex(c, false).Println(" #" + c)
		}
	} else {
		color.Red.Println("There are no colours found in this file")
	}
}
