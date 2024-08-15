### Acceptance Tests

**Acceptance tests are essential as they directly impact your ability to verify your system and confidently evolve it over time in a more deliberate and methodical way, reducing wasted effort.**

They're also a fantastic tool to help you work with legacy code. When faced with a poor codebase without any tests, please resist the temptation to start refactoring. Instead, write some acceptance tests to give you a safety net to freely change the system's internals without affecting its functional external behaviour. ATs need not be concerned with internal quality, so they're a great fit in these situations.

Think about what prompts acceptance tests to change:
* An external behaviour change. If you want to change what the system does, changing the acceptance test suite seems reasonable, if not desirable.
* An implementation detail change / refactoring. Ideally, this shouldn't prompt a change, or if it does, a minor one.
Too often, though, the latter is the reason acceptance tests have to change. To the point where engineers even become reluctant to change their system because of the perceived effort of updating tests!

When done right, this approach gives us flexibility in our implementation detail and stability in our specifications. Importantly, **it provides a simple and obvious structure for managing change**, which becomes essential as a system and its team grows.


_'I test in prod'. A must read: https://increment.com/testing/i-test-in-production/_


### Test-Driven Development

TDD is focused on letting you design for the behaviour you precisely need, iteratively. When starting a new area, you must identify a key, necessary behaviour and aggressively cut scope.
**Follow a "top-down" approach, starting with an acceptance test (AT) that exercises the behaviour from the outside. This will act as a north-star for your efforts. All you should be focused on is making that test pass.** This test will likely be failing for a while whilst you develop enough code to make it pass.
Once your AT is set up, you can break into the TDD process to drive out enough units to make the AT pass. The trick is to not worry too much about design at this point; get enough code to make the AT pass because you're still learning and exploring the problem.
Taking this first step is often more extensive than you think, setting up web servers, routing, configuration, etc., which is why keeping the scope of the work small is essential. We want to make that first positive step on our blank canvas and have it backed by a passing AT so we can continue to iterate quickly and safely.
As you develop, listen to your tests, and they should give you signals to help you push your design in a better direction but, again, anchored to the behaviour rather than our imagination.
**Typically, your first "unit" that does the hard work to make the AT pass will grow too big to be comfortable, even for this small amount of behaviour. This is when you can start thinking about how to break the problem down and introduce new collaborators**.
This is where test doubles (e.g. fakes, mocks) are handy because most of the complexity that lives internally within software doesn't usually reside in implementation detail but "between" the units and how they interact.
