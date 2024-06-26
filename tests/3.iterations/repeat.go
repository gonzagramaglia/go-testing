package iterations

func Repeat(character string, repeatCount int) string {
	var chainOfCharacters string
	for i := 0; i < repeatCount; i++ {
		chainOfCharacters += character
	}
	return chainOfCharacters
}
