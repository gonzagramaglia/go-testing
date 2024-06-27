// To do stuff repeatedly in Go, you'll need `for`.
// In Go there are no `while`, `do`, `until` keywords, you can only use for.
// Which is a good thing!
package iterations

// This function chains X times the characters entered as an argument
// and returns said chain as a return
func Repeat(characters string, repeatCount int) string {
	var chainOfCharacters string
	for i := 0; i < repeatCount; i++ {
		chainOfCharacters += characters
	}
	return chainOfCharacters
}
