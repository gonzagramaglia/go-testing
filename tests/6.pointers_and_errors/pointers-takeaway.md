### Pointers let us point to some values and then let us change them. 

So rather than taking a copy of the whole Wallet, we instead take a pointer to that wallet so that we can change the original values within it.
```
func (w *Wallet) Deposit(amount int) {
	w.balance += amount
}

func (w *Wallet) Balance() int {
	return w.balance
}
```
The difference is the receiver type is *Wallet rather than Wallet which you can read as "a pointer to a wallet".

Try and re-run the tests and they should pass.

Now you might wonder, why did they pass?
We didn't dereference the pointer in the function, like so:
```
func (w *Wallet) Balance() int {
	return (*w).balance
}
```
and seemingly addressed the object directly.
In fact, the code above using (*w) is absolutely valid.
However, the makers of Go deemed this notation cumbersome,
so the language permits us to write w.balance, without
an explicit dereference. These pointers to structs even have
their own name: struct pointers and they are automatically
dereferenced.

Technically you do not need to change Balance to use a pointer
receiver as taking a copy of the balance is fine.
However, by convention you should keep your method receiver
types the same for consistency.
