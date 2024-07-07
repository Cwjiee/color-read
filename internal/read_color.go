package internal

import (
	"regexp"
	"strings"
)

func ReadColor(f []byte) []string {
	var hexs []string
	var hex []string
	var start bool

	for _, byt := range f {
		char := string(byt)

		if char == "#" {
			start = !start
		}

		if start {

			hex = append(hex, char)
			if len(hex) == 7 {
				pattern := `^#([A-Fa-f0-9]{6}|[A-Fa-f0-9]{3})$`

				match, _ := regexp.MatchString(pattern, strings.Join(hex, ""))
				if !match {
					hex = hex[:0]
					start = !start
					continue
				}

				hexs = append(hexs, strings.Join(hex, ""))
				hex = hex[:0]
				start = !start
			}
		}
	}

	return hexs
}
