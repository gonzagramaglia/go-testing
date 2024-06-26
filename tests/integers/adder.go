// This way you can generate an overview in the documentation
// of this file that is generated with `godoc -http=:XXXX`
package integers

// Add takes two integers and returns the sum of them.
func Add(x, y int) int {
	return x + y
}

/*

If you have noticed we learnt about named return value
in the hello_test but aren't using the same here.
It should generally be used when the meaning of the result
isn't clear from context, in our case it's pretty much clear
that Add function will add the parameters.

*/
