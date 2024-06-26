# Common Go Test Flags: Detailed Examples and Outputs

When working with Go, the go test command is essential for running tests. Here, we will delve into some of the most commonly used flags, providing detailed examples and expected outputs.

### 1. Verbose (-v)
```
go test -v
```
Description: This flag enables detailed output for each test that is executed.

Example:
```
go test -v
```
Output:
```
=== RUN   TestFunction1
--- PASS: TestFunction1 (0.00s)
=== RUN   TestFunction2
--- PASS: TestFunction2 (0.00s)
PASS
ok      example.com/mypackage   0.002s
```

### 2. Run (-run)
```
go test -run TestNamePattern
```
Description: Runs tests matching the specified pattern. Useful for running specific tests without executing the entire suite.

Example:
```
go test -run ^TestFunction1$
```
Output:
```
=== RUN   TestFunction1
--- PASS: TestFunction1 (0.00s)
PASS
ok      example.com/mypackage   0.001s
```

### 3. Short (-short)
```
go test -short
```
Description: Signals tests to run in "short mode," which is often used to skip long-running tests.

Example:
```
go test -short
```
Output:
```
PASS
ok      example.com/mypackage   0.002s
```

### 4. Count (-count)
```
go test -count N
```
Description: Specifies the number of times to run each test. This is useful for catching flaky tests that may fail intermittently.

Example:
```
go test -count=3
```
Output:
```
=== RUN   TestFunction1
--- PASS: TestFunction1 (0.00s)
=== RUN   TestFunction1
--- PASS: TestFunction1 (0.00s)
=== RUN   TestFunction1
--- PASS: TestFunction1 (0.00s)
PASS
ok      example.com/mypackage   0.003s
```

### 5. Cover (-cover)
```
go test -cover
```
Description: Enables code coverage analysis, showing the percentage of code that is covered by tests.

Example:
```
go test -cover
```

Output:
```
PASS
coverage: 75.0% of statements
ok      example.com/mypackage   0.002s
```

### 6. Coverprofile (-coverprofile)
```
go test -coverprofile=cover.out
```
Description: Writes the coverage profile to the specified file, which can be used with other tools for detailed analysis.

Example:
```
go test -coverprofile=cover.out
```

Output:
```
PASS
coverage: 75.0% of statements
ok      example.com/mypackage   0.002s
```
Contents of cover.out:
```
mode: set
example.com/mypackage/mypackage.go:12.12,14.6 2 1
example.com/mypackage/mypackage.go:17.6,20.12 3 1
...
```

### 7. CPU (-cpu)

Usage: go test -cpu 1,2,4

Description: Specifies the number of CPU cores to use for the test run, useful for performance testing.

Example:
```
go test -cpu 1,2,4
```
Output:
```
PASS
ok      example.com/mypackage   0.004s
```

### 8. Parallel (-parallel)
```
go test -parallel N
```
Description: Specifies the maximum number of tests to run in parallel. This can speed up testing for packages with many tests.

Example:
```
go test -parallel=4
```

Output:
```
PASS
ok      example.com/mypackage   0.003s
```

### 9. Timeout (-timeout)
```
go test -timeout 30s
```
Description: Sets a timeout for the test run. This prevents tests from running indefinitely.

Example:
```
go test -timeout 30s
```
Output:
```
PASS
ok      example.com/mypackage   0.002s
```


By understanding and utilizing these flags, you can significantly enhance your testing workflow in Go. These examples illustrate how each flag can be applied and what kind of output you can expect, providing a clearer picture of how to leverage go test for effective testing.