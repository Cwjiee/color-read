package internal

import (
	"regexp"
)

func ReadColor(f []byte) []string {
	hexPattern := `#([A-Fa-f0-9]{6}|[A-Fa-f0-9]{3})\b`
	//	rgbPattern := `rgb\(\s*\d{1,3}\s*,\s*\d{1,3}\s*,\s*\d{1,3}\s*\)`

	hexRegex := regexp.MustCompile(hexPattern)
	//	rgbRegex := regexp.MustCompile(rgbPattern)

	hexColors := hexRegex.FindAllString(string(f), -1)
	//	rgbColors := rgbRegex.FindAllString(string(f), -1)

	return hexColors
}

// -------- old method --------

// func ReadColor(f []byte) []string {
// 	var hexs []string
// 	var hex []string
// 	var start bool
//
// 	for _, byt := range f {
// 		char := string(byt)
//
// 		if char == "#" {
// 			start = !start
// 		}
//
// 		if start {
//
// 			hex = append(hex, char)
// 			if len(hex) == 7 {
// 				pattern := `^#([A-Fa-f0-9]{6}|[A-Fa-f0-9]{3})$`
//
// 				match, _ := regexp.MatchString(pattern, strings.Join(hex, ""))
// 				if !match {
// 					hex = hex[:0]
// 					start = !start
// 					continue
// 				}
//
// 				hexs = append(hexs, strings.Join(hex, ""))
// 				hex = hex[:0]
// 				start = !start
// 			}
// 		}
// 	}
//
// 	return hexs
// }
