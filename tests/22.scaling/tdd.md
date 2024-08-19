
### Test-Driven Development

TDD is focused on letting you design for the behaviour you precisely need, iteratively. When starting a new area, you must identify a key, necessary behaviour and aggressively cut scope.
**Follow a "top-down" approach, starting with an acceptance test (AT) that exercises the behaviour from the outside. This will act as a north-star for your efforts. All you should be focused on is making that test pass.** This test will likely be failing for a while whilst you develop enough code to make it pass.
Once your AT is set up, you can break into the TDD process to drive out enough units to make the AT pass. The trick is to not worry too much about design at this point; get enough code to make the AT pass because you're still learning and exploring the problem.
Taking this first step is often more extensive than you think, setting up web servers, routing, configuration, etc., which is why keeping the scope of the work small is essential. We want to make that first positive step on our blank canvas and have it backed by a passing AT so we can continue to iterate quickly and safely.
As you develop, listen to your tests, and they should give you signals to help you push your design in a better direction but, again, anchored to the behaviour rather than our imagination.
**Typically, your first "unit" that does the hard work to make the AT pass will grow too big to be comfortable, even for this small amount of behaviour. This is when you can start thinking about how to break the problem down and introduce new collaborators**.
This is where test doubles (e.g. fakes, mocks) are handy because most of the complexity that lives internally within software doesn't usually reside in implementation detail but "between" the units and how they interact.