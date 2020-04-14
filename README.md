# go-best-test
# Good testing practice in building web applications in Go
# Unit, integration, end-to-end and benchmark tests (performance) 

# Running

1. Run `go mod init github.com/monkrus/go-best-test`

2. Interesting testing-related packages : `testing`, `testing/quick`, `testing/iotest` and `net/http/httptest `

3. Other projects 
  (assertions)[github.com/stretchcom/testif],  
   github.com/onsi/ginkgo (BDD), 
   goconvey.com (visual results in the browser),
   github.com/gavv/httpexpect (end-to-end testing, e.g. RestAPI), 
   code.google.com/p/gomock (injecting mocks into your tests), 
   github.com/DATA-DOG/go-sqlmock (DB testing)

# Test creating rules

1. Add `_test` to filenames
2. Prefix tests with `Test`
3. Accept one parameter `*testing.T` 
4. Remove `_test suffix` for whitebox tests

# Writing tests

1. Receive `*testing.T` object

2. No special assertions API, use normal Go code

3. Report errors using methods on `*testing.T object`

4. Immediate failure 
   FailNow(), 
   Fatal(args ..interface{}),
   Fatalf(format strings, args ..interface{})

5. Non-immediate failure
   Fail(),
   Error(args ...interface{}),
   Errorf(format string, args ..interface{})

# Testing

1. `go test` or `go test {pkg1}`

2. `go test./...` Will fund tests on current directory and subtree

3. `go test -v` Generates verbose output

4. `go test -run {regexp}` To concentrate on the spesific tests (matching)

5. `go help testflags` 

6. `go test run {function name}`- to run specific function

# Test coverage report

1. `go test -cover`

2. Reports `go test -coverprofile cover.out` or `go tool cover -func cover.out`

3. Graphical coverage `go test -coverprofile count.out -covermode count` and then `go tool cover -html cover.out`

# Other methods

1. Log and logf-(write a message testing output)

2. Helper (for helper method)

3. Skip, Skipf, SkipNow

4. Run (create subtest)

5. Parallel









