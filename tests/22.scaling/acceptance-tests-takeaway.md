### Acceptance Tests

**Acceptance tests are essential as they directly impact your ability to verify your system and confidently evolve it over time in a more deliberate and methodical way, reducing wasted effort.**

They're also a fantastic tool to help you work with legacy code. When faced with a poor codebase without any tests, please resist the temptation to start refactoring. Instead, write some acceptance tests to give you a safety net to freely change the system's internals without affecting its functional external behaviour. ATs need not be concerned with internal quality, so they're a great fit in these situations.

Think about what prompts acceptance tests to change:
* An external behaviour change. If you want to change what the system does, changing the acceptance test suite seems reasonable, if not desirable.
* An implementation detail change / refactoring. Ideally, this shouldn't prompt a change, or if it does, a minor one.
Too often, though, the latter is the reason acceptance tests have to change. To the point where engineers even become reluctant to change their system because of the perceived effort of updating tests!

When done right, this approach gives us flexibility in our implementation detail and stability in our specifications. Importantly, **it provides a simple and obvious structure for managing change**, which becomes essential as a system and its team grows.


_'I test in prod'. A must read: https://increment.com/testing/i-test-in-production/_


### When should I write acceptance tests?
The best practice is to favour having lots of fast running unit tests and a few acceptance tests, but how do you decide when you should write an acceptance test, vs unit tests?
It's difficult to give a concrete rule, but the questions I typically ask myself are:
* Is this an edge case (or a permutation of business rules)? I'd prefer to unit test those
* Is this something that the non-computer people talk about a lot? I would prefer to have a lot of confidence the key thing "really" works, so I'd add an acceptance test
* Am I describing a user journey, rather than a specific function? Acceptance test
* Would unit tests give me enough confidence? Sometimes you're taking an existing journey that already has an acceptance test, but you're adding other functionality to deal with different scenarios due to different inputs. In this case, adding another acceptance test adds a cost but brings little value, so I'd prefer some unit tests.
