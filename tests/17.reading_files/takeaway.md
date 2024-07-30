### Thinking about the kind of test we want to see

Let's remind ourselves of our mindset and goals when starting:
**Write the test we want to see**. Think about how we'd like to use the code we're going to write from a consumer's point of view.
Focus on *what* and *why*, but don't get distracted by *how*.


##### Notice that the package of our test is blogposts_test. 
Remember, when TDD is practiced well we take a consumer-driven approach: we don't want to test internal details because consumers don't care about them. By appending _test to our intended package name, we only access exported members from our package - just like a real user of our package.