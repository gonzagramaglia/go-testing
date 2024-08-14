### Refactoring

Using generics as a tool for simplifying code via the refactoring step is far more likely to guide you to useful improvements, rather than premature abstractions.

### Good source control habits

Practice combining TDD with good source control habits. Commit your work when your test is passing, before trying to refactor. This way if you make a mess, you can easily get yourself back to your working state.

### Names matter

Make an effort to do some research outside of Go, so you don't re-invent patterns that already exist with an already established name.
Writing a function takes a collection of A and converts them to B? Don't call it Convert, that's Map. Using the "proper" name for these items will reduce the cognitive burden for others and make it more search engine friendly to learn more.

### Fold

Fold is a real fundamental in computer science. Here's an interesting resource if you wish to dig more into it: http://www.cs.nott.ac.uk/~pszgmh/fold.pdf 