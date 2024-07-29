### Wrapping up

##### Table based tests
When you're testing something which feels like it's a matter of "given input X, we expect Y" you should probably use **table based tests**.

##### Intro to property based tests
There have been a few rules in the domain of Roman Numerals that we have worked with in this chapter:
* Can't have more than 3 consecutive symbols.
* Only I (1), X (10) and C (100) can be "subtractors".
* Taking the result of `ConvertToRoman(N)` and passing it to `ConvertToArabic` should return us `N`.
The tests we have written so far can be described as "example" based tests where we provide examples for the tooling to verify.
What if we could take these rules that we know about our domain and somehow exercise them against our code?
Property based tests help you do this by throwing random data at your code and verifying the rules you describe always hold true. A lot of people think property based tests are mainly about random data but they would be mistaken. The real challenge about property based tests is having a good understanding of your domain so you can write these properties.

##### More TDD practice with iterative development
Did the thought of writing code that converts 1984 into MCMLXXXIV feel intimidating to you at first? It did to me and I've been writing software for quite a long time.
The trick, as always, is to **get started with something simple and take small steps**.
The skill is knowing how to split work up, and that comes with practice and with some lovely TDD to help you on your way.