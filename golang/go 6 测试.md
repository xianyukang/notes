## Table of Contents
  - [测试](#%E6%B5%8B%E8%AF%95)
    - [The Basics of Testing](#The-Basics-of-Testing)
    - [Reporting Test Failures](#Reporting-Test-Failures)
    - [Setting Up and Tearing Down](#Setting-Up-and-Tearing-Down)
    - [Storing Sample Test Data  ](#Storing-Sample-Test-Data)
    - [Caching Test Results](#Caching-Test-Results)
    - [Testing Your Public API](#Testing-Your-Public-API)
    - [Use go-cmp to Compare Test Results](#Use-gocmp-to-Compare-Test-Results)
    - [Table Tests](#Table-Tests)
    - [Checking Your Code Coverage](#Checking-Your-Code-Coverage)
    - [Benchmarks](#Benchmarks)
    - [Mocks and Stubs](#Mocks-and-Stubs)
    - [httptest](#httptest)

## 测试

### The Basics of Testing

Over the past two decades, the widespread adoption of automated testing has probably done more to improve code quality than any other software engineering technique. As a language and ecosystem focused on improving software quality, it’s not surprising that *Go includes testing support as part of its standard library*.

Go’s testing support has two parts: libraries and tooling. The `testing` package in the standard library provides the types and functions to write tests, while the `go test` tool that’s bundled with Go runs your tests and generates reports.

![image-20220716191517337](https://static.xianyukang.com/img/image-20220716191517337.png) 

(1) Every test is written in a file whose name ends with `_test.go`. If you are writing tests against foo.go, place your tests in a file named foo_test.go.  

(2) Test functions start with the word `Test` and take in a single parameter of type `*testing.T`. By convention, this parameter is named `t`. Test functions do not return any values.   

(3) Also note that we use standard Go code to call the code being tested and to validate if the responses are as expected. When there’s an incorrect result, we report the error with the `t.Error` method, which works like the `fmt.Print` function.

(4) Just as `go build` builds a binary and `go run` runs a file, the command `go test` runs the tests in the current directory. The `go test` command allows you to specify which packages to test. Using `./...` for the package name specifies that you want to run tests in the current directory and all of the subdirectories of the current directory. Include a `-v` flag to get verbose testing output.

### Reporting Test Failures

While Error and Errorf mark a test as failed, the test function continues running. If you think a test function should stop processing as soon as a failure is found, use the Fatal and Fatalf methods.  The difference is that the test function exits immediately after the test failure message is generated. Note that this doesn’t exit all tests; any remaining test functions will execute after the current test function exits.  

When should you use Fatal/Fatalf and when should you use Error/Errorf? If the failure of a check in a test means that further checks in the same test function will always fail or cause the test to panic, use Fatal or Fatalf. If you are testing several independent items (such as validating fields in a struct), then use Error or Errorf so you can report as many problems at once. This makes it easier to fix multiple problems without rerunning your tests over and over.

### Setting Up and Tearing Down

Sometimes you have some common state that you want to set up before any tests run and remove when testing is complete. Use a TestMain function to manage this state and run your tests:  

![image-20220717130654161](https://static.xianyukang.com/img/image-20220717130654161.png) 

Both TestFirst and TestSecond refer to the package-level variable testTime. Once the state is configured, call the Run method on *testing.M to run the test functions. The Run method returns the exit code; 0 indicates that all tests passed. Finally, you must call os.Exit with the exit code returned from Run.  

Be aware that TestMain is invoked once, not before and after each individual test. Also be aware that you can have only one TestMain per package. There are two common situations where TestMain is useful:  

- When you need to set up data in an external repository, such as a database.
- When the code being tested depends on package-level variables that need to be initialized.

*As mentioned before (and will be mentioned again!) you should avoid package-level variables in your programs*. They make it hard to understand how data flows through your program. If you are using TestMain for this reason, consider refactoring your code.  

The `Cleanup` method on *testing.T. is used to clean up temporary resources created for a single test. This method has a single parameter, a function with no input parameters or return values. The function runs when the test completes. For simple tests, you can achieve the same result by using a defer statement, but Cleanup is useful when tests rely on helper functions to set up sample data. It’s fine to call Cleanup multiple times. Just like defer, the functions are invoked in last added, first called order.  

### Storing Sample Test Data  

As `go test` walks your source code tree, *it uses the current package directory as the current working directory*. If you want to use sample data to test functions in a package, create a subdirectory named `testdata` to hold your files. Go reserves this directory name as a place to hold test files. When reading from testdata, always use a relative file reference. Since go test changes the current working directory to the current package, each package accesses its own testdata via a relative file path.  

### Caching Test Results

Go caches compiled packages if they haven’t changed, Go also caches test results when running tests across multiple packages if they have passed and their code hasn’t changed. *The tests are recompiled and rerun if you change any file in the package or in the testdata directory*. You can also force tests to always run if you pass the flag `-count=1` to go test.  

### Testing Your Public API

The tests that we’ve written are in the same package as the production code. This allows us to test both exported and unexported functions.

If you want to test just the public API of your package, Go has a convention for specifying this. You still keep your test source code in the same directory as the production source code, but you use `packagename_test` for the package name. (一个目录中只能有一个包, 但 aaa 和 aaa_test 可以共存)

![image-20220717135343270](https://static.xianyukang.com/img/image-20220717135343270.png) 

Notice that the package name for our test file is adder_test. We have to import test_examples/adder even though the files are in the same directory. To follow the convention for naming tests, the test function name matches the name of the AddNumbers function. Also note that we use adder.AddNumbers, since we are calling an exported function in a different package.  

The advantage of using the _test package suffix is that it lets you treat your package as a “black box”; you are forced to interact with it only via its exported functions, methods, types, constants, and variables.

### Use go-cmp to Compare Test Results

It can be verbose to write a thorough comparison between two instances of a compound type. While you can use reflect.DeepEqual to compare structs, maps, and slices, there’s a better way. Google released a third-party module called `go-cmp` that does the comparison for you and returns a detailed description of what does not match.

![image-20220718131426057](https://static.xianyukang.com/img/image-20220718131426057.png) 

The cmp.Diff function takes in the expected output and the output that was returned by the function that we’re testing. It returns a string that describes any mismatches between the two inputs. If the inputs match, it returns an empty string.

![image-20220718131921204](https://static.xianyukang.com/img/image-20220718131921204.png) 

### Table Tests

Most of the time, it takes more than a single test case to validate that a function is working correctly. You could write multiple test functions to validate your function or multiple tests within the same function, but you’ll find that a great deal of the testing logic is repetitive. You set up supporting data and functions, specify inputs, check the outputs, and compare to see if they match your expectations. Rather than writing this over and over, you can take advantage of a pattern called table tests.

![image-20220718135000092](https://static.xianyukang.com/img/image-20220718135000092.png) 

![image-20220718134919140](https://static.xianyukang.com/img/image-20220718134919140.png) 

![image-20220718135257635](https://static.xianyukang.com/img/image-20220718135257635.png) 

### Checking Your Code Coverage

Code coverage is a very useful tool for knowing if you’ve missed any obvious cases. However, reaching 100% code coverage doesn’t guarantee that there aren’t bugs in your code for some inputs.   

Adding the -cover flag to the go test command calculates coverage information and includes a summary in the test output. If you include a second flag -coverprofile, you can save the coverage information to a file:  `go test -v -cover -coverprofile=c.out`

It’d be more useful if we could see what we missed. The cover tool included with Go generates an HTML representation of your source code with that information: `go tool cover -html=c.out`

Every file that’s tested appears in the combo box in the upper left. The source code is in one of three colors. Gray is used for lines of code that aren’t testable, green is used for code that’s been covered by a test, and red is used for code that hasn’t been tested. From looking at this, we can see that we didn’t write a test to cover the default case, when a bad operator is passed to our function. Let’s add that case to our slice of test cases: {"bad_op", 2, 2, "?", 0, "unknown operator ?"}

Code coverage is a great thing, but it’s not enough. *You can have 100% code coverage and still have bugs in your code!* There’s actually a bug in our code, even though we have 100% coverage.

### Benchmarks

建议看书上 13 章的对应小节

Our goal is to find out what size buffer we should use to read from the file.  

Before you spend time going down an optimization rabbit hole, be sure that you need to optimize. If your program is already fast enough to meet your responsiveness requirements and is using an acceptable amount of memory, then your time is better spent on adding features and fixing bugs. Your business requirements determine what “fast enough” and “acceptable amount of memory” mean  

In Go, benchmarks are functions in your test files that start with the word `Benchmark` and take in a single parameter of type `*testing.B`. This type includes all of the functionality of a *testing.T as well as additional support for benchmarking.

Every Go benchmark must have a loop that iterates from 0 to `b.N`. The testing framework calls our benchmark functions over and over with larger and larger values for N until it is sure that the timing results are accurate. 

![image-20220719154716241](https://static.xianyukang.com/img/image-20220719154716241.png) 

### Mocks and Stubs

So far, we’ve written tests for functions that didn’t depend on other code. This is not typical as most code is filled with dependencies. There are two ways that Go allows us to abstract function calls: defining a function type and defining an interface.

### httptest

Now let’s see how to use the httptest library to test this code without standing up a server.  
