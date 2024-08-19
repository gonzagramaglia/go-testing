### Try to remove comments

`A heuristic we follow is that whenever we feel the need to comment something, we write a method instead. -- Martin Fowler`

This quote suggests that instead of adding a comment to explain a complex piece of code, you should extract that code into a separate method with a descriptive name. This approach has several benefits:

1. *Self-explanatory code*: By moving the code into a method, you can give it a name that describes its purpose, making the code more self-explanatory.
2. *Reduced comments*: Comments can become outdated or misleading over time. By using method names to convey meaning, you reduce the need for comments.
3. *Improved readability*: Short, focused methods are easier to understand than large blocks of code with comments.
4. *Reusability*: Extracted methods can be reused in other parts of the codebase, reducing duplication.
5. *Easier maintenance*: If the logic needs to change, it's easier to modify a single method than to update comments and code in multiple places.

For example, instead of:
```
// Calculate the total cost with tax and discount
total = price * (1 + taxRate) * (1 - discountRate);
```

You would extract a method:

```
calculateTotalWithTaxAndDiscount(price, taxRate, discountRate) {
return price * (1 + taxRate) * (1 - discountRate);
}
```

This approach promotes cleaner, more maintainable code and reduces the need for comments.