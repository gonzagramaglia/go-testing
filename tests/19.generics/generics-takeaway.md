### Generics are simpler than using interface{} in most cases

Using interface{} makes your code:
* Less safe (mix apples and oranges), requires more error handling
* Less expressive, interface{} tells you nothing about the data
* More likely to rely on reflection, type-assertions etc which makes your code more difficult to work with and more error prone as it pushes checks from compile-time to runtime


Using statically typed languages is an act of describing constraints. If you do it well, you create code that is not only safe and simple to use but also simpler to write because the possible solution space is smaller.

Generics gives us a new way to express constraints in our code, which as demonstrated will allow us to consolidate and simplify code that was not possible until Go 1.18.
