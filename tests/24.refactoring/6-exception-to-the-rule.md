### Exceptions to the rule

There are improvements you can make to your code that require a change in your tests, which I would still be happy to put into the "refactoring" bucket, even though it breaks the rule.

A simple example would be **renaming a public symbol** (e.g., a method, type, or function) **with `shift+F6`**. This will, of course, change the production and test codes.

However, as **it is an automated and safe change**, **the risk** of going into a spiral of breaking tests and production code that so many fall into with other kinds of design changes **is minimal**.

For that reason, any changes you can safely perform with your IDE/editor, I would still happily call refactoring.