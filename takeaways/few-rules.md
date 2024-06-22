Writing a test is just like writing a function, with a few rules
1. It needs to be in a file with a name like xxx_test.go
2. The test function must start with the word Test
3. The test function takes one argument only t *testing.T
4. To use the *testing.T type, you need to import "testing"