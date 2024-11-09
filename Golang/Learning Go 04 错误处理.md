## Table of Contents
  - [入门](#%E5%85%A5%E9%97%A8)
    - [错误处理概述](#%E9%94%99%E8%AF%AF%E5%A4%84%E7%90%86%E6%A6%82%E8%BF%B0)
    - [返回 error/ok 表示数据无效](#%E8%BF%94%E5%9B%9E-errorok-%E8%A1%A8%E7%A4%BA%E6%95%B0%E6%8D%AE%E6%97%A0%E6%95%88)
  - [创建错误](#%E5%88%9B%E5%BB%BA%E9%94%99%E8%AF%AF)
    - [errors.New()](#errorsNew)
    - [Sentinel Errors](#Sentinel-Errors)
    - [Custom Errors](#Custom-Errors)
    - [Error Strings](#Error-Strings)
  - [使用哪种方式返回错误](#%E4%BD%BF%E7%94%A8%E5%93%AA%E7%A7%8D%E6%96%B9%E5%BC%8F%E8%BF%94%E5%9B%9E%E9%94%99%E8%AF%AF)
    - [如何返回 nil error](#%E5%A6%82%E4%BD%95%E8%BF%94%E5%9B%9E-nil-error)
    - [New/Sentinel/Custom/Wrap](#NewSentinelCustomWrap)
  - [处理错误](#%E5%A4%84%E7%90%86%E9%94%99%E8%AF%AF)
    - [有四种方法](#%E6%9C%89%E5%9B%9B%E7%A7%8D%E6%96%B9%E6%B3%95)
    - [一般不忽略错误](#%E4%B8%80%E8%88%AC%E4%B8%8D%E5%BF%BD%E7%95%A5%E9%94%99%E8%AF%AF)
    - [通常要避免 panic](#%E9%80%9A%E5%B8%B8%E8%A6%81%E9%81%BF%E5%85%8D-panic)
  - [Wrapping Errors](#Wrapping-Errors)
    - [Why Wrapping Errors](#Why-Wrapping-Errors)
    - [The Unwrap method](#The-Unwrap-method)
    - [Wrapping errors with %w](#Wrapping-errors-with-w)
    - [Whether to Wrap](#Whether-to-Wrap)
    - [Examining errors with Is and As](#Examining-errors-with-Is-and-As)
    - [使用 errors.Is 和 errors.As](#%E4%BD%BF%E7%94%A8-errorsIs-%E5%92%8C-errorsAs)
  - [处理 Panic](#%E5%A4%84%E7%90%86-Panic)
    - [一般推荐返回 error](#%E4%B8%80%E8%88%AC%E6%8E%A8%E8%8D%90%E8%BF%94%E5%9B%9E-error)
    - [什么时候适合 panic](#%E4%BB%80%E4%B9%88%E6%97%B6%E5%80%99%E9%80%82%E5%90%88-panic)
    - [使用 recover 恢复 panic](#%E4%BD%BF%E7%94%A8-recover-%E6%81%A2%E5%A4%8D-panic)
    - [什么时候适合 recover](#%E4%BB%80%E4%B9%88%E6%97%B6%E5%80%99%E9%80%82%E5%90%88-recover)
    - [用 panic + recover 简化错误处理](#%E7%94%A8-panic--recover-%E7%AE%80%E5%8C%96%E9%94%99%E8%AF%AF%E5%A4%84%E7%90%86)
    - [在 defer 函数里 panic 会怎样?](#%E5%9C%A8-defer-%E5%87%BD%E6%95%B0%E9%87%8C-panic-%E4%BC%9A%E6%80%8E%E6%A0%B7)

## 入门

### 错误处理概述

#### ➤ 返回 error 而不是抛出异常

Go handles errors by returning a value of type `error` as the last return value for a function. The calling function then checks the error return value by comparing it to `nil`, handling the error, or returning an error of its own. A new error is created from a string by calling the `errors.New` function.

**Go errors are values**. Unlike languages with exceptions, Go doesn’t have special constructs to detect if an error was returned.  `error` is a built-in interface that defines a single method: `Error() string`. Anything that implements this interface is considered an error.   

#### ➤ 在返回值中添加 error 明示函数可能遇到错误

Use `error` to signal that a function can fail. By convention, `error` is the last result parameter.

```go
func Lookup() (*Result, error) {}
```

Returning a `nil` error is the idiomatic way to signal a successful operation that could otherwise fail. If a function returns an error, callers must treat all non-error return values as `unspecified` unless explicitly documented otherwise. Commonly, the non-error return values are their zero values, but this cannot be assumed.

#### ➤ 为什么不像大多数语言那样抛出异常呢?

There are two very good reasons why Go uses a returned error instead of thrown exceptions. First, exceptions add at least one new code path through the code. These paths are sometimes unclear, especially in languages whose functions don’t include a declaration that an exception is possible. This produces code that crashes in surprising ways when exceptions aren’t properly handled, or, even worse, code that doesn’t crash but whose data is not properly initialized, modified, or stored.

The second reason is more subtle, but demonstrates how Go’s features work together. The Go compiler requires that all variables must be read. Making errors returned values forces developers to either check and handle error conditions or make it explicit that they are ignoring errors by using an underscore `_` for the returned error value.  

Exception handling may produce shorter code, but having fewer lines doesn’t necessarily make code easier to understand or maintain. As we’ve seen, idiomatic Go favors clear code, even if it takes more lines.  

总而言之 Golang 把错误当做函数的另一个返回值，让你不会忘记/忽略错误处理，更容易写出健壮的程序

### 返回 error/ok 表示数据无效

In C and similar languages, it is common for functions to return values like -1, null, or the empty string to signal errors or missing results. This is known as in-band error handling. Failing to check for an in-band error value can lead to bugs and can attribute errors to the wrong function.

```go
// Bad:
// Lookup returns the value for key or -1 if there is no mapping for key.
func Lookup(key string) int
// 这种做法容易出错，容易忘记处理异常返回值
```

Go's support for multiple return values provides a better solution. A function should return an additional value to indicate whether its other return values are valid. This return value may be an error, or a boolean when no explanation is needed.

如果只返回一个 null 表示失败，那么很容易忘了检查 null，然后 NullPointerException.

```go
GetUser().PrintName()  // 在 Java 中容易 NullPointerException
GetUser().PrintName()  // 在 Go 中不存在这种问题,  因为 GetUser() 返回两个值
```

Returning errors in this way encourages more robust and explicit error handling:

```go
value, ok := Lookup(key)
if !ok {
    return fmt.Errorf("no value for %q", key)
}
return Parse(value)
```

Some standard library functions, like those in package `strings`, return in-band error values. This greatly simplifies string-manipulation code at the cost of requiring more diligence from the programmer. In general, Go code in the Google codebase should return additional values for errors.

## 创建错误

### errors.New()

Go’s standard library provides two ways to create an error from a string. The first is the `errors.New` function. It takes in a string and returns an `error`. This string is returned when you call the `Error` method on the returned error instance. If you pass an error to `fmt.Println`, it calls the `Error` method automatically. The second way is to use the `fmt.Errorf` function. This function allows you to use all of the formatting verbs for `fmt.Printf` to create an error.

### Sentinel Errors

#### ➤ 例子: `var ErrNotExist = errors.New("file does not exist")`

Sentinel errors are one of the few variables that are declared at the package level. By convention, their names start with `Err` (with the notable exception of `io.EOF`).  Be sure you need a sentinel error before you define one. Once you define one, it is part of your public API and you have committed to it being available in all future backward-compatible releases. It’s far better to reuse one of the existing ones in the standard library.

#### ➤ 为什么 errors.New("hey") != errors.New("hey")

```markdown
- 每次 errors.New("hey") 都会创建一个新实例并返回指针
- 每次调用都返回不同的指针，所以不相等
```

#### ➤ 为什么要让 errors.New("hey") != errors.New("hey")，这么做的原因是什么?

```go
// 如果两次调用 errors.New("hey") 的返回值相等, 就会导致 windows.ErrSystem == linux.ErrSystem
var ErrSystem = errors.New("System Error") // 在 windows 包中定义
var ErrSystem = errors.New("System Error") // 在 linux 包中定义

// 总而言之, 我们希望 errors.New() 创建一个唯一的错误
// 而不是只要两个 error 的错误消息相同, 就认为它们相等
// 然后确保了 errors.New() 的唯一性, 就能放心使用 if err == windows.ErrSystem {...}
```

### Custom Errors

Since `error` is an interface, you can define your own errors that include additional information for logging or error handling. For example, you might want to include a status code as part of the error to indicate the kind of error that should be reported back to the user. 

```go
type NotFoundError struct {
    File string
}

func (e *NotFoundError) Error() string {
    return fmt.Sprintf("file %q not found", e.File)
}

func Open(file string) error {
    return &NotFoundError{File: file}
}

func TestError(t *testing.T) {
    if err := Open("abc.txt"); err != nil {
        var notFound *NotFoundError    // 此处类型必须和 Open 的返回值类型完全一致
        if errors.As(err, &notFound) {
            fmt.Println("not found error:", notFound.File)
        } else {
            panic("unknown error")
        }
    }
}
```

#### ➤ 为什么实现 error 接口要用 pointer receiver ?

[Receiver vs pointer reciever on custom error : r/golang](https://www.reddit.com/r/golang/comments/16eoxsv/receiver_vs_pointer_reciever_on_custom_error/)

```go
type MyError struct{}

// 这里使用 value receiver 做演示, 实现 error 接口一般推荐用 pointer receiver
func (e MyError) Error() string {
    return "my error"
}

func Open() error {
    // 因为实现接口时使用了 value receiver 所以如下两种返回形式都合法
    // 那就选择困难了呗, 然后接收方也要纠结你返回的是指针还是值, 所以我们约定一个惯例
    return MyError{}
    return &MyError{}
}
```

### Error Strings

#### ➤ 无需首字母大写，不要加句点

Error strings should not be capitalized (unless beginning with an exported name, a proper noun or an acronym) and should not end with punctuation. This is because error strings usually appear within other context before being printed to the user.

```go
err := fmt.Errorf("something bad happened")
```

#### ➤ 错误消息要详细，比如包含操作类型，操作对象，底层错误，当前的 package name

```go
type PathError struct {
    Op   string // "open", "unlink", etc.
    Path string // The associated file.
    Err  error  // Returned by the system call.
}

func (e *PathError) Error() string {
    return e.Op + " " + e.Path + ": " + e.Err.Error() // 其实就是把函数的输入/输出记下来，能方便排错
}
```

PathError's `Error()` generates a string like this: `open /etc/passwx: no such file or directory`. Such an error, which includes the problematic file name, the operation, and the operating system error it triggered, is useful even if printed far from the call that caused it; **it is much more informative than the plain "no such file or directory"**.

When feasible, error strings should identify their origin, such as by having a prefix naming the operation or package that generated the error. For example, in package `image`, the string representation for a decoding error due to an unknown format is `image: unknown format`.

#### ➤ 添加一些上下文信息，方便查错

Can you suggest some problems with the following piece of code? An obvious suggestion is that the five lines of the function could be replaced with `return authenticate(r.User)`. But this is the simple stuff that everyone should be catching in code review. More fundamentally the problem with this code is I cannot tell where the original error came from. Don’t just check errors, handle them gracefully.

```go
func AuthenticateRequest(r *Request) error {
    err := authenticate(r.User)
    if err != nil {
        return err
    }
    return nil
}
```

If `authenticate` returns an error, then `AuthenticateRequest` will return the error to its caller, who will probably do the same, and so on. At the top of the program the main body of the program will print the error to the screen or a log file, and all that will be printed is: `No such file or directory`.

There is no information of file and line where the error was generated. There is no stack trace of the call stack leading up to the error. The author of this code will be forced to a long session of bisecting their code to discover which code path trigged the file not found error. 

#### ➤ 所以推荐用 %w 动词补充上下文, 让错误消息有迹可循

```go
file, err := os.Open("foo.txt")
if err != nil {
    return fmt.Errorf("open foo.txt failed: %w", err) // %w 表示 wrap
}
```

## 使用哪种方式返回错误

### 如何返回 nil error

#### ➤ 两个办法，返回 nil 或 nil interface

You shouldn’t declare a variable to be the type of your custom error and then return that variable. The reason why err is non-nil is that `error` is an interface. For an interface to be considered `nil`, both the underlying type and the underlying value must be `nil`. Just keep in mind that if any concrete value has been stored in the interface, the interface will not be `nil`.

```go
type CustomError1 interface{ error }
type CustomError2 struct{}
func (e *CustomError2) Error() string { return "" }

func ReturnError() (error, error) {
    var e1 CustomError1  // 返回的 error 接口为 {type: nil, value: nil}
    var e2 *CustomError2 // 返回的 error 接口为 {type: *CustomError2, value: nil}
    return e1, e2
}

func TestReturnError(t *testing.T) {
    e1, e2 := ReturnError()
    fmt.Println("returned e1 == nil:", e1 == nil)
    fmt.Println("returned e2 == nil:", e2 == nil)
}
```

There are two ways to fix this. The most common approach is to explicitly return `nil`. Another approach is to make sure that any local variable that holds an error is of type `error`.

#### ➤ 函数的返回值用 error 而不是具体类型

Exported functions that return errors should return them using the `error` type. Concrete error types are susceptible to subtle bugs: a concrete `nil` pointer can get wrapped into an interface and thus become a non-nil value. 

```go
// Bad:
func ReturnConcreteError() *os.PathError {
    return nil
}

func TestReturnConcreteError(t *testing.T) {
    var err error
    // 如果把具体类型赋值给接口类型, 会让接口不等于 nil
    err = ReturnConcreteError()
    if err != nil {
        fmt.Println("error is:", err)
    }
}
```

### New/Sentinel/Custom/Wrap

#### ➤ [选哪种方案? 点此查看示例](https://github.com/uber-go/guide/blob/master/style.md#errors)

There are few options for declaring errors. Consider the following before picking the option best suited for your use case.

- Does the caller need to match the error so that they can handle it? If yes, we must support the [`errors.Is`](https://golang.org/pkg/errors/#Is) or [`errors.As`](https://golang.org/pkg/errors/#As) functions by declaring a top-level error variable or a custom type.
- Is the error message a static string, or is it a dynamic string that requires contextual information? For the former, we can use [`errors.New`](https://golang.org/pkg/errors/#New), but for the latter we must use [`fmt.Errorf`](https://golang.org/pkg/fmt/#Errorf) or a custom error type.
- Are we propagating a new error returned by a downstream function? If so, see the [section on error wrapping](https://github.com/uber-go/guide/blob/master/style.md#error-wrapping).

| Error matching? | Error Message | Guidance                                                     |
| --------------- | ------------- | ------------------------------------------------------------ |
| No              | static        | [`errors.New`](https://golang.org/pkg/errors/#New)           |
| No              | dynamic       | [`fmt.Errorf`](https://golang.org/pkg/fmt/#Errorf)           |
| Yes             | static        | top-level `var` with [`errors.New`](https://golang.org/pkg/errors/#New) |
| Yes             | dynamic       | custom `error` type                                          |

#### ➤ 要不要在 error string 中添加上下文 ?

There are three main options for propagating errors if a call fails:

- return the original error as-is
- add context with `fmt.Errorf` and the `%w` verb
- add context with `fmt.Errorf` and the `%v` verb

Return the original error as-is if there is no additional context to add. This maintains the original error type and message. This is well suited for cases when the underlying error message has sufficient information to track down where it came from.

Otherwise, add context to the error message where possible so that **instead of a vague error such as "connection refused", you get more useful errors such as "call service foo: connection refused"**.

#### ➤ 使用 %w 还是 %v ?

Use `fmt.Errorf` to add context to your errors, picking between the `%w` or `%v` verbs based on whether the caller should be able to match and extract the underlying cause.

- Use `%w` if the caller should have access to the underlying error. This is a good default for most wrapped errors, but be aware that callers may begin to rely on this behavior. So for cases where the wrapped error is a known `var` or type, document and test it as part of your function's contract.
- Use `%v` to obfuscate the underlying error. Callers will be unable to match it, but you can switch to `%w` in the future if needed.

When adding context to returned errors, keep the context succinct by avoiding phrases like "failed to", which state the obvious and pile up as the error percolates up through the stack:

```go
return fmt.Errorf("failed to create new store: %w", err) // bad,  嵌套时产生一堆 fail to 很难看
return fmt.Errorf("new store: %w", err)                  // good, 简洁一点, 嵌套时也能保证可读性
```

## 处理错误

### 有四种方法

Code that encounters an error should make a deliberate choice about how to handle it. It is not usually appropriate to discard errors using `_` variables. If a function returns an error, do one of the following:

- Handle and address the error immediately.
- Return the error to the caller.
- In exceptional situations, call [`log.Fatal`](https://pkg.go.dev/github.com/golang/glog#Fatal) or (if absolutely necessary) `panic`.
- If it is appropriate to ignore or discard an error, an accompanying comment should explain why this is safe.

```go
var b *bytes.Buffer
n, _ := b.Write(p) // never returns a non-nil error
```

### 一般不忽略错误

示例:

```go
// Bad! This code will crash if path does not exist.
fi, _ := os.Stat(path)
if fi.IsDir() {
    fmt.Printf("%s is a directory\n", path)
}
// 应该始终检查函数返回的错误值
// 如果 err 不为 nil, 那么和错误一起返回的返回值就是无效的,  不应该继续使用
```

#### ➤ **Don’t write fragile programs**.

The core logic for this example is relatively short. Of the 22 lines inside the for loop, 6 of them implement the actual algorithm and the other 16 are error checking and data validation. You might be tempted to not validate incoming data or check errors, but doing so produces unstable, unmaintainable code. 

#### ➤ **Error handling is what separates the professionals from the amateurs**.  

### 通常要避免 panic

#### ➤ Don't Panic

Code running in production must avoid panics. If an error occurs, the function must return an error and allow the caller to decide how to handle it.

```go
func run(args []string) error {
    if len(args) == 0 {
        panic("an argument is required")             // bad, 调用方没得选
        return errors.New("an argument is required") // good, 调用方想怎样处理错误都行
    }
    return nil
}
```

#### ➤ Handle Type Assertion Failures

- This is bad: `t := i.(string)`
- The single return value form of a type assertion will panic on an incorrect type.
- Therefore, always use the "comma ok" idiom.

#### ➤ Panic/recover is not an error handling strategy. 

A program must panic only when something irrecoverable happens such as a `nil` dereference. Even in tests, prefer `t.Fatal` or `t.FailNow` over panics to ensure that the test is marked as failed. An exception to this is program initialization: bad things at program startup that should abort the program may cause panic.

```go
var _statusTemplate = template.Must(template.New("name").Parse("_statusHTML")) // must 意思是不成功就 panic 终止程序启动
```

#### ➤ Exit in Main

Call one of `os.Exit` or `log.Fatal` only in `main()`. All other functions should return errors to signal failure. Programs with multiple functions that exit present a few issues: 

- Non-obvious control flow: Any function can exit the program so it becomes difficult to reason about the control flow.

- Skipped cleanup: When a function exits the program, it skips function calls enqueued with `defer` statements.

- This adds risk of skipping important cleanup tasks. ( 跳过必要的清理会导致资源泄露、数据不一致 )


## Wrapping Errors

#### [➤ 参考文档](https://go.dev/blog/go1.13-errors)

### Why Wrapping Errors

We may sometimes want to define a new error type that contains the underlying error, adding information to it, preserving it for inspection by code. 

```go
// 如何 Wrap 错误并添加额外信息:
// (1) 定义一个 QueryError 类型,  保留底层 error 的同时添加了额外的 query 信息
type QueryError struct {
    Query string
    Err   error
}
// (2) 定义 Unrap 和 Error 方法
func (e *QueryError) Unwrap() error { return e.Err }

func (e *QueryError) Error() string { return e.Query + ": " + e.Err.Error() }

// 如果是错误 QueryError 类型并且底层错误是 ErrPermission,  针对这种情况做特殊处理
if e, ok := err.(*QueryError); ok && e.Err == ErrPermission {
    // 因为 Wrap 了错误,  这里既能知道底层错误是 ErrPermission,  还能知道引发错误的 Query
}
```

When you preserve an error while adding additional information, it is called wrapping the error. When you have a series of wrapped errors, it is called an error chain.

### The Unwrap method

An error which contains another may implement an `Unwrap` method returning the underlying error. If `e1.Unwrap()` retlurns `e2`, then we say that `e1` wraps `e2`, and that you can unwrap `e1` to get `e2`. The result of unwrapping an error may itself have an `Unwrap` method.

### Wrapping errors with %w

In Go 1.13, the `fmt.Errorf` function supports a new `%w` verb. When this verb is present, the error returned by `fmt.Errorf` will have an `Unwrap` method returning the argument of `%w`, which must be an error. Use `%w` to create an error whose formatted string includes the formatted string of another error and which contains the original error as well. The convention is to write `: %w` at the end of the error format string and make the error to be wrapped the last parameter passed to `fmt.Errorf`.  

```go
// %w 和 %v 产生的错误消息相同,  区别在于有没有 wrap 错误
if err != nil {
    // Return an error which unwraps to err.
    return fmt.Errorf("decompress %v: %w", name, err)
    // Creating a new error with `fmt.Errorf` discards everything from the original error except the text. 
    return fmt.Errorf("decompress %v: %v", name, err)
}
```

### Whether to Wrap

#### ➤ [参考此处](https://go.dev/blog/go1.13-errors#:~:text=Is(err%2C%20ErrPermission)%20...-,Whether%20to%20Wrap,-%C2%B6)

When adding additional context to an error, either with `fmt.Errorf` or by implementing a custom type, you need to decide whether the new error should wrap the original. There is no single answer to this question. Wrap an error to expose it to callers. Do not wrap an error when doing so would expose implementation details. 

总而言之，如果调用方需要判断底层错误是什么，并且这个底层错误不可替代，那么就 wrap 错误

一般而言，调用方只需要知道数据库查不到数据，我们返回自定义的 `data.ErrRecordNotFound` 就好，别返回 postgresql 的底层错误

如果底层错误被 wrap 或直接返回，那它就成了 API 的一部分，调用方可以依赖它，我们为了兼容性也必须始终返回它，这不利于切换实现

### Examining errors with Is and As

Wrapping errors is a useful way to get additional information about an error, but it introduces problems. If a sentinel error is wrapped, you cannot use `==` to check for it, nor can you use a type assertion or type switch to match a wrapped custom error. Go solves this problem with two functions in the `errors` package, Is and As.

- The `errors.Is` function compares an error to a value.  
- The `errors.As` function tests whether an error is a specific type.  
- The `errors.Unwrap(e)` function returns `e.Unwrap()`, or `nil` when the error has no `Unwrap` method.

> 观察 `errors.Is(err, target)` 的源码可知，它会递归地调用 `e.Unwrap` 方法，
>
> 对 error chain 上的每一层执行 `err == target || err.Is(target)`，只要有一层匹配就返回 true

By default, `errors.Is` uses `==` to compare each wrapped error with the specified error. If this does not work for an error type that you define (for example, if your error is a noncomparable type), implement the `Is` method on your error.

```go
func (m MyErr) Is(target error) bool {
    if me2, ok := target.(MyErr); ok {
        return reflect.DeepEqual(me, me2)
    }
    return false
}
```

Another use for defining your own `Is` method is to allow comparisons against errors that aren’t identical instances. You might want to pattern match your errors:  

```go
if errors.Is(err, ResourceErr{Resource: "Database"}) {
    fmt.Println("The database is broken:", err)
}
```

### 使用 errors.Is 和 errors.As

#### ➤ 用 Is 检查错误链

```go
errors.Is(err, sql.ErrNoRows)   // good
err == sql.ErrNoRows            // bad, 因为如果哪天 wrap 了错误, 这行代码就会失效
```

#### ➤ 用 As 获取错误类型

```go
// errors.As 和类型断言的作用一样 (但能处理嵌套的 error):
//   if e, ok := err.(*QueryError); ok { … }
var err = RunQuery()
var e *QueryError           // 因为 RunQuery() 返回了指针,  所以 e 要保持一致使用 *QueryError 类型
if errors.As(err, &e) {     // 实现 error 接口一般用指针接收器，所以函数返回的自定义错误通常也是 *CustomError 指针
    fmt.Println(e.Query)
}

// 注意第二个参数必须是非空指针, 并且指向的东西:
// - 要么实现了 error 接口
// - 要么是接口类型
if errors.As(err, &myErr) {
    // 编译期会做上述类型检查，总之我们确保第二个参数是 non-nil pointer 那就不会 panic
}
```

## 处理 Panic

### 一般推荐返回 error

#### ➤ 启动时遇到初始化错误，可以 panic

The usual way to report an error to a caller is to return an `error` as an extra return value. Real library functions should avoid `panic`. If the problem can be masked or worked around, it's always better to let things continue to run rather than taking down the whole program.

But what if the error is unrecoverable? Sometimes the program simply cannot continue. One possible example is during initialization: if the library truly cannot set itself up, it might be reasonable to panic.

```go
var user = os.Getenv("USER")

func init() {
    if user == "" {
        panic("no value for $USER")
    }
}
```

### 什么时候适合 panic

#### ➤ 错误有两类，一类可以预见，另一类则不符合预期，可以用 panic 表示程序有逻辑错误

The first class of errors are expected errors that may occur during normal operation.  
比如数据库查询超时、暂时网络不通、错误的用户输入...，以下这种情况要返回错误值而不是 panic:

- `可以预见、且应该预见的错误`
- `这些错误正常流程可能会发生、且允许发生、不属于代码 bug`

The other class of errors are unexpected errors. These are errors which *should not happen during normal operation*, and if they do it is probably the result of a developer *mistake or a logical error* in your codebase. These errors are truly exceptional, and using panic in these circumstances is more widely accepted. In fact, the Go standard library frequently does this when you make a logical error or try to use the language features in an unintended way — such as when trying to access an out-of-bounds index in a slice, or trying to close an already-closed channel.

One possible counterexample is during initialization: if the library truly cannot set itself up, it might be reasonable to panic. Another exception to this is when returning the error adds an unacceptable amount of error handling to the rest of your codebase.  

### 使用 recover 恢复 panic

When `panic` is called, including implicitly for run-time errors such as indexing a slice out of bounds or failing a type assertion, it immediately stops execution of the current function and begins unwinding the stack of the goroutine, running any deferred functions along the way. *If that unwinding reaches the top of the goroutine's stack, the program dies.* However, it is possible to use the built-in function `recover` to regain control of the goroutine and resume normal execution. One application of `recover` is to shut down a failing goroutine inside a server without killing the other executing goroutines.

The built-in function `panic` takes one parameter, which can be of any type. Usually, it is a string.  Go provides a way to capture a panic to provide a more graceful shutdown or to prevent shutdown at all. The built-in `recover` function is called from within a `defer` to check if a panic happened. If there was a panic, the value assigned to the panic is returned.

```go
func div(i int) {
    // 在 defer 中调用 recover 就能恢复 panic
    // The value returned by the builtin recover() function has the type interface{}, and its
    // underlying type could be string, error, or something else — whatever the parameter passed to panic() was.
    defer func() {
        if v := recover(); v != nil {
            fmt.Println(v)
        }
    }()
    fmt.Println(60 / i)
}

func 恢复panic() {
    for _, val := range []int{1, 2, 0, 6} {
        div(val)
    }
}
```

#### ➤ 为什么不要滥用 panic

While panic and recover look a lot like exception handling in other languages, they are not intended to be used that way. The reason we don’t rely on panic and recover is that recover *doesn’t make clear what could fail*. It just ensures that if something fails, we can print out a message and continue. Idiomatic Go favors code that explicitly outlines the possible failure conditions over shorter code that handles anything while saying nothing.  

Reserve panics for fatal situations and use `recover` as a way to gracefully handle these situations. If your program panics, be very careful about trying to continue executing after the panic. It’s very rare that you want to keep your program running after a panic occurs.   

#### ➤ 一般返回 error 而不是 panic

In the preceding sample program, it would be idiomatic to check for division by zero and return an error if one was passed in.



### 什么时候适合 recover

#### ➤ 用于避免 panic 传播到调用者，用于把 panic 转成 error 再返回

There is one situation where `recover` is recommended. If you are creating a library for third parties, do not let panics escape the boundaries of your public API. If a panic is possible, a public function should use a recover to convert the panic into an error, return it, and let the calling code decide what to do with them.  

#### ➤ 某个 goroutine panic 会导致整个应用挂掉，可用 recover 避免这种情况

One application of `recover` is to shut down a failing goroutine inside a server without killing the other executing goroutines. In this example, if `do(work)` panics, the result will be logged and the goroutine will exit cleanly without disturbing the others.

```go
func server(workChan <-chan *Work) {
    for work := range workChan {
        go safelyDo(work)
    }
}

func safelyDo(work *Work) {
    // recover 一下, 避免一个 goroutine panic 让整个应用挂掉
    defer func() {
        if err := recover(); err != nil {
            log.Println("work failed:", err)
        }
    }()
    do(work)
}
```

### 用 panic + recover 简化错误处理

#### ➤ 有时候逐层返回错误会让人无法忍受

We can use that idea to simplify error handling in complex software. Let's look at an idealized version of a `regexp` package, which reports parsing errors by calling `panic` with a local error type.

```go
// 这个 Error 类型用于区分不同的 panic
type Error string
func (e Error) Error() string { return string(e) }

// error is a method of *Regexp that reports parsing errors by panicking with an Error.

func (regexp *Regexp) error(err string) {
    panic(Error(err))
}

func Compile(str string) (regexp *Regexp, err error) {
    regexp = new(Regexp)
    defer func() {
        if e := recover(); e != nil {
            regexp = nil    // Clear return value.
            err = e.(Error) // Will re-panic and change the panic value if not a parse error.
        }
    }()
    // 重点: doParse() 中会用 regexp.error("xxx error") 抛出错误, 而不是逐层返回错误值 ( 逐层返回很麻烦 )
    return regexp.doParse(str), nil
}
```

#### ➤ 短短几行代码，却并不简单 (非常巧妙):

If `doParse` panics, the recovery block will set the return value to `nil`—deferred functions can modify named return values. It will then check, in the assignment to `err`, that the problem was a parse error by asserting that it has the local type `Error`. If it does not, the type assertion will fail, causing a run-time error that continues the stack unwinding as though nothing had interrupted it. This check means that if something unexpected happens, such as an index out of bounds, the code will fail even though we are using `panic` and `recover` to handle parse errors.

#### ➤ re-panic 修改了 panic value,  但也没什么大问题:

By the way, this re-panic idiom changes the panic value if an actual error occurs. However, both the original and new failures will be presented in the crash report, so the root cause of the problem will still be visible (崩溃报告会打印整个 panic 链). Thus this simple re-panic approach is usually sufficient—it's a crash after all—but if you want to display only the original value, you can write a little more code to filter unexpected problems and re-panic with the original error. 

#### ➤ 若不想修改 panic value 可以这样做:

```go
func main() {
    defer func() {
        err := recover().(int)
        fmt.Println(err + 2)
    }()
    defer func() {
        err := recover()
        // var _ = err.(string)              // 这会修改 panic value
        if _, ok := err.(string); !ok {
            panic(err)                       // 重新抛出原来的错误
        }
    }()
    panic(123)
}
```

### 在 defer 函数里 panic 会怎样?

[go - Is it okay to panic inside defer function, especially when it's already panicking?](https://stackoverflow.com/questions/41139447/is-it-okay-to-panic-inside-defer-function-especially-when-its-already-panickin)

- 可以同时存在多个 panic，它们都会在错误堆栈中显示，但 recover 只返回最近一个 panic，再次 recover 会返回 nil

- 程序因 panic 异常结束时会打印所有的 panic，包括那些已经被 recover 过的 ( 例如 `panic: 123 [recovered]` )

```go
func TestMultiPanic(t *testing.T) {
    defer panic(3) // will replace panic 2
    defer panic(2) // will replace panic 1
    defer panic(1) // will replace panic 0
    panic(0)
}
```

- defer 函数中 panic 没啥特殊的，其他 defer 函数会继续执行，这个 panic 也能被上层函数恢复

```go
func TestPanicDefer(t *testing.T) {
    defer func() {
        fmt.Println("defer 1:", recover()) // 注意 recover 只返回最近的一个 panic
        fmt.Println("defer 1:", recover()) // 再次调用 recover 会返回 nil
    }()
    defer func() {
        fmt.Println("defer 2:           ") // 再次 panic, 加上 defer 3 里的 panic
        panic("re-panic one more time")    // 注意现在同时有两个 panic
    }()
    defer func() {
        fmt.Println("defer 3:", recover()) // 这里返回 nil 因为已经被处理了
        panic("re-panic")                  // 在 defer 中 panic, 可以被上面的 recover 捕获
    }()
    defer func() {
        fmt.Println("defer 4:", recover()) // 恢复下面的 panic(111)
    }()
    panic(111)                                // 先执行这行, 然后从下往上执行 defer 函数
}
```

