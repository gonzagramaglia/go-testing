### Mocking

* **Without mocking important areas of your code will be untested**. In our case we would not be able to test that our code paused between each print but there are countless other examples. Calling a service that can fail? Wanting to test your system in a particular state? It is very hard to test these scenarios without mocking.
* Without mocks you may have to set up databases and other third parties things just to test simple business rules. You're likely to have slow tests, resulting in **slow feedback loops**.
* By having to spin up a database or a webservice to test something you're likely to have fragile tests due to the unreliability of such services.
Once a developer learns about mocking it becomes very easy to over-test every single facet of a system in terms of the way it works rather than what it does. Always be mindful about **the value of your tests** and what impact they would have in future refactoring.