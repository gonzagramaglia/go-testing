### Unit Tests

As a project grows you often may be in situations where the acceptance tests can take a few minutes to run. To offer a friendly developer experience for people checking out your project, you can enable developers to run the different kinds of tests separately.
It's preferable that running `go test ./...` should be runnable with no further set up from an engineer, beyond say a few key dependencies such as the Go compiler (obviously) and perhaps Docker.

##### go test -short
Go provides a mechanism for engineers to run only "short" tests with the short flag
`go test -short ./...`
We can add to our acceptance tests to see if the user wants to run our acceptance tests by inspecting the value of the flag

```
if testing.Short() {
	t.Skip()
}
I made a Makefile to show this usage
```

```
build:
	golangci-lint run
	go test ./...

unit-tests:
	go test -short ./...
```