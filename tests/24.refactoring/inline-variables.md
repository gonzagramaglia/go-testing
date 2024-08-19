### Inline variables

If you create a variable, only for it to be passed on to another method/function:
```
url := baseURL + "/user/" + id
res, err := client.Get(url)
```
Consider inlining it (command+option+n --> shortcut for Intellij/Goland) unless the variable name adds significant meaning.
```
res, err := client.Get(baseURL + "/user/" + id)
```
Don't be too clever about inlining; the goal is not to have zero variables and instead have ridiculous one-liners that no one can read. If you can add significant naming to a value, it might be best to leave it be.