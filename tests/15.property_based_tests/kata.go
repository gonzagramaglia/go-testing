package kata

import "strings"

func ConvertToRoman(arabic int) string {

	var result strings.Builder

	for i := arabic; i > 0; i-- {
		if result.String() == "IV" {
			writeRomanNumbers(&result, "V")
		} else if result.String() == "III" {
			writeRomanNumbers(&result, "IV")
		} else {
			result.WriteString("I")
		}
	}

	return result.String()
}

func writeRomanNumbers(r *strings.Builder, stringToWrite string) {
	r.Reset()
	r.WriteString(stringToWrite)
}
