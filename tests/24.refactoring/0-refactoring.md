### Refactoring

Refactoring is just improving existing code and not changing behaviour; therefore, tests shouldn't have to change. This is why it's the 3rd step of the TDD cycle.

You're doing something else if you are "refactoring" some code and having to change tests at the same time.

Many very helpful refactorings are simple to learn and easy to do (your IDE almost entirely automates many) but, over time, become hugely impactful to the quality of our system.



What if you want to change the signature of a method?
```
func (b BirthdayGreeter) WishHappyBirthday(age int, firstname, lastname string, email Email) {
	// some fascinating emailing code
}
```
You may feel its argument list is too long and want to bring more cohesion and meaning to the code.
```
func (b BirthdayGreeter) WishHappyBirthday(person Person)
```

You should still be able to drive this change with a test first. You can split hairs over whether this is a "behaviour" change, but you want your method to behave differently.

As this is a behaviour change, apply the TDD process here too. One benefit of TDD is that it gives you a simple, safe, repeatable way of driving behaviour change in your system; why abandon it in these situations just because it feels different?

In this case, you'll change your existing tests to use the new type. The iterative, small steps you usually do with TDD to reduce risk and bring discipline & clarity will help you in these situations, too.

Chances are you'll have several tests that call WishHappyBirthday; in these scenarios, I'd suggest commenting out all but one of the tests, driving out the change, and then working through the rest of the tests as you see fit.