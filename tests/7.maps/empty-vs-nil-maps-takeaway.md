### Pointers, copies and others

An interesting property of maps is that you can modify them without passing as an address to it (e.g &myMap)
This may make them feel like a "reference type", but as Dave Cheney describes they are not.
>"A map value is a pointer to a runtime.hmap structure."

So **when you pass a map to a function/method, you are indeed copying it, but just the pointer part, not the underlying data structure that contains the data**.

**A gotcha with maps is that they can be a nil value. A nil map behaves like an empty map when reading, but attempts to write to a nil map will cause a runtime panic**. You can read more about maps [here](https://go.dev/blog/maps).
**Therefore, you should never initialize a nil map variable**:
```
var m map[string]string
```
Instead, you can initialize an empty map or use the make keyword to create a map for you:
```
var dictionary = map[string]string{}

// OR

var dictionary = make(map[string]string)
```
Both approaches create an empty hash map and point dictionary at it. Which ensures that you will never get a runtime panic.