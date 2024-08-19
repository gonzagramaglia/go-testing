### DRY up values with extract variables

Using the same value multiple times in a function? Consider extracting and capturing a variable in a meaningful variable name (command+option+v --> shortcut for Intellij/Goland).

This helps with readability and makes changing the value easier in future, as you won't have to remember to update multiple occurrences of the same value.


### Dont Repeat Yourself (DRY)

Reducing the number of lines of code is often a side-effect of DRY, but it is not the actual goal.

Rather than being extremist on either side of "must DRY everything" or "DRY is bad", engage your brain and think about the code you see in front of you. What is repeated? Does it need to be? Does the parameter list look sensible if you encapsulate some repeated code into a method? Does it feel self-documenting and encapsulate the "idea" clearly?

Nine times out of 10, you can look at the argument list of a function, and if it looks messy and confusing, then it is likely to be a poor application of DRY.
If making some code DRY feels hard, you're probably making things more complex; consider stopping.

DRY with care, **but practising this frequently will improve your judgement**. I encourage my colleagues to "just try it" and use source control to get back to safety if it is wrong.

**Trying these things will teach you more than discussing it**, and source control coupled with good automated tests gives you the perfect setup to experiment and learn.