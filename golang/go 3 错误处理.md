## Table of Contents
  - [错误处理](#%E9%94%99%E8%AF%AF%E5%A4%84%E7%90%86)
    - [错误处理概述](#%E9%94%99%E8%AF%AF%E5%A4%84%E7%90%86%E6%A6%82%E8%BF%B0)
    - [字符串与简单错误](#%E5%AD%97%E7%AC%A6%E4%B8%B2%E4%B8%8E%E7%AE%80%E5%8D%95%E9%94%99%E8%AF%AF)
    - [Sentinel Errors](#Sentinel-Errors)
    - [自定义错误类型](#%E8%87%AA%E5%AE%9A%E4%B9%89%E9%94%99%E8%AF%AF%E7%B1%BB%E5%9E%8B)
    - [显式返回 nil 表示没有错误](#%E6%98%BE%E5%BC%8F%E8%BF%94%E5%9B%9E-nil-%E8%A1%A8%E7%A4%BA%E6%B2%A1%E6%9C%89%E9%94%99%E8%AF%AF)
    - [返回错误的最佳实践](#%E8%BF%94%E5%9B%9E%E9%94%99%E8%AF%AF%E7%9A%84%E6%9C%80%E4%BD%B3%E5%AE%9E%E8%B7%B5)
  - [Wrapping Errors](#Wrapping-Errors)
    - [Why Wrapping Errors](#Why-Wrapping-Errors)
    - [The Unwrap method](#The-Unwrap-method)
    - [Wrapping errors with %w](#Wrapping-errors-with-w)
    - [Whether to Wrap](#Whether-to-Wrap)
    - [Examining errors with Is and As](#Examining-errors-with-Is-and-As)
    - [注意 errors.Is 和 errors.As](#%E6%B3%A8%E6%84%8F-errorsIs-%E5%92%8C-errorsAs)
  - [处理 Panic](#%E5%A4%84%E7%90%86-Panic)
    - [什么时候适合 panic](#%E4%BB%80%E4%B9%88%E6%97%B6%E5%80%99%E9%80%82%E5%90%88-panic)
    - [使用 recover 恢复 panic](#%E4%BD%BF%E7%94%A8-recover-%E6%81%A2%E5%A4%8D-panic)
    - [什么时候适合 recover](#%E4%BB%80%E4%B9%88%E6%97%B6%E5%80%99%E9%80%82%E5%90%88-recover)
    - [用 recover 简化错误处理](#%E7%94%A8-recover-%E7%AE%80%E5%8C%96%E9%94%99%E8%AF%AF%E5%A4%84%E7%90%86)

## 错误处理

### 错误处理概述

➤ 返回错误值、而不是抛出错误

Go handles errors by returning a value of type `error` as the last return value for a function. The calling function then checks the error return value by comparing it to `nil`, handling the error, or returning an error of its own.  A new error is created from a string by calling the `errors.New` function. Error messages should not be capitalized nor should they end with punctuation or a newline.  

**Go errors are values**. Unlike languages with exceptions, Go doesn’t have special constructs to detect if an error was returned.  `error` is a built-in interface that defines a single method: `Error() string`. Anything that implements this interface is considered an error.   

➤ Golang 这种错误处理风格の背后的考虑

There are two very good reasons why Go uses a returned error instead of thrown exceptions. First, exceptions add at least one new code path through the code. These paths are sometimes unclear, especially in languages whose functions don’t include a declaration that an exception is possible. This produces code that crashes in surprising ways when exceptions aren’t properly handled, or, even worse, code that doesn’t crash but whose data is not properly initialized, modified, or stored.

The second reason is more subtle, but demonstrates how Go’s features work together. The Go compiler requires that all variables must be read. Making errors returned values forces developers to either check and handle error conditions or make it explicit that they are ignoring errors by using an underscore `_` for the returned error value.  

Exception handling may produce shorter code, but having fewer lines doesn’t necessarily make code easier to understand or maintain. As we’ve seen, idiomatic Go favors clear code, even if it takes more lines.  

### 字符串与简单错误

Go’s standard library provides two ways to create an error from a string. The first is the `errors.New` function. It takes in a string and returns an `error`. This string is returned when you call the `Error` method on the returned error instance. If you pass an error to `fmt.Println`, it calls the `Error` method automatically. The second way is to use the `fmt.Errorf` function. This function allows you to use all of the formatting verbs for `fmt.Printf` to create an error.

### Sentinel Errors

➤ 例子: `var ErrNotExist = errors.New("file does not exist")`

Sentinel errors are one of the few variables that are declared at the package level. By convention, their names start with `Err` (with the notable exception of `io.EOF`).  Be sure you need a sentinel error before you define one. Once you define one, it is part of your public API and you have committed to it being available in all future backward-compatible releases. It’s far better to reuse one of the existing ones in the standard library.

### 自定义错误类型

Since `error` is an interface, you can define your own errors that include additional information for logging or error handling. For example, you might want to include a status code as part of the error to indicate the kind of error that should be reported back to the user. 

```go
type NotFoundError struct {
    File string
}

// (1) 实现 error 接口时使用 pointer receiver
func (e *NotFoundError) Error() string {
    return fmt.Sprintf("file %q not found", e.File)
}

// (2) 返回错误时,  返回指针
func Open(file string) error {
    return &NotFoundError{File: file}
}

func main() {
    if err := Open("abc.txt"); err != nil {
        // (3) 接收错误时,  也使用指针,  不要用值类型 var notFound NotFoundError
        var notFound *NotFoundError
        if errors.As(err, &notFound) {
            fmt.Println("not found error:", notFound.File)
        } else {
            panic("unknown error")
        }
    }

    // 这里有个坑,  (2) 和 (3) 的类型必须相同,  都是指针或者都是值
    // 推荐 (1) 处用指针接收器,  这样 IDE 能在 (2) 和 (3) 不匹配时发出警告
}
```

### 显式返回 nil 表示没有错误

You shouldn’t declare a variable to be the type of your custom error and then return that variable. The reason why err is non-nil is that `error` is an interface. For an interface to be considered `nil`, both the underlying type and the underlying value must be `nil`.  

```go
type SomeErr []string
func (e SomeErr) Error() string { return e[0] }

func HandleSomething() error {
    var err SomeErr
    fmt.Println(err == nil) // true
    return err              // 返回的 error 接口为 {type: SomeErr, value: nil}
    // return nil           // 如果想表示没有错误,  应该显式返回 nil,  若返回变量可能导致 err != nil
}

func main() {
    err := HandleSomething()
    fmt.Println(err == nil) // false,  因为 type 不为空
}
```

There are two ways to fix this. The most common approach is to explicitly return `nil`. Another approach is to make sure that any local variable that holds an error is of type `error`.

### 返回错误的最佳实践

➤ [具体例子](https://github.com/uber-go/guide/blob/master/style.md#errors)

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

## Wrapping Errors

[➤ 参考文档](https://go.dev/blog/go1.13-errors)

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

When adding additional context to an error, either with `fmt.Errorf` or by implementing a custom type, you need to decide whether the new error should wrap the original. There is no single answer to this question. Wrap an error to expose it to callers. Do not wrap an error when doing so would expose implementation details. 

[点此查看两者的适用情况.](https://go.dev/blog/go1.13-errors#:~:text=Is(err%2C%20ErrPermission)%20...-,Whether%20to%20Wrap,-When%20adding%20additional)  
总而言之,  如果调用方需要基于底层错误做逻辑,  那么就 wrap 错误,  让被 wrap 的错误成为 API 的一部分

### Examining errors with Is and As

Wrapping errors is a useful way to get additional information about an error, but it introduces problems. If a sentinel error is wrapped, you cannot use `==` to check for it, nor can you use a type assertion or type switch to match a wrapped custom error. Go solves this problem with two functions in the `errors` package, Is and As.

- The `errors.Is` function compares an error to a value.  
- The `errors.As` function tests whether an error is a specific type.  
- The `errors.Unwrap(e)` function returns `e.Unwrap()`, or `nil` when the error has no `Unwrap` method.

观察 `errors.Is` 的源码可知,  他会递归的调用 `e.Unwrap` 方法,  
对 error chain 上每一层 error 执行 == 和 Is 判断,  只要有一层匹配就返回 true

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
// errors.Is 如何判定两个对象的相等性?  先尝试 == 后尝试 Is,  其中一个是 true 则返回 true
```

### 注意 errors.Is 和 errors.As

➤ 用 Is 检查错误链

```go
errors.Is(err, sql.ErrNoRows)   // good
err == sql.ErrNoRows            // bad, 因为如果哪天 wrap 了错误, 这行代码就会失效
```

➤ 用 As 获取错误类型

```go
// errors.As 和类型断言的作用一样 (但能处理嵌套的 error):
//   if e, ok := err.(*QueryError); ok { … }
var err = RunQuery()
var e *QueryError           // 因为 RunQuery() 返回指针而不是结构体,  所以 e 也要用v
if errors.As(err, &e) {
    fmt.Println(e.Query)
}
```

## 处理 Panic

### 什么时候适合 panic

The first class of errors are expected errors that may occur during normal operation.  
比如数据库查询超时、网络不通、错误的用户输入...,  对于这种可以预见的错误要返回错误值而不是 panic

The other class of errors are unexpected errors. These are errors which *should not happen during normal operation*, and if they do it is probably the result of a developer *mistake or a logical error* in your codebase. These errors are truly exceptional, and using panic in these circumstances is more widely accepted. In fact, the Go standard library frequently does this when you make a logical error or try to use the language features in an unintended way — such as when trying to access an out-of-bounds index in a slice, or trying to close an already-closed channel.

One possible counterexample is during initialization: if the library truly cannot set itself up, it might be reasonable to panic.

### 使用 recover 恢复 panic

The built-in function `panic` takes one parameter, which can be of any type. Usually, it is a string.  Go provides a way to capture a panic to provide a more graceful shutdown or to prevent shutdown at all. The built-in `recover` function is called from within a `defer` to check if a panic happened. If there was a panic, the value assigned to the panic is returned.

![image-20220516203534170](https://static.xianyukang.com/img/image-20220516203534170.png) 

➤ 为什么不要滥用 panic

While panic and recover look a lot like exception handling in other languages, they are not intended to be used that way. The reason we don’t rely on panic and recover is that recover *doesn’t make clear what could fail*. It just ensures that if something fails, we can print out a message and continue. Idiomatic Go favors code that explicitly outlines the possible failure conditions over shorter code that handles anything while saying nothing.  

Reserve panics for fatal situations and use `recover` as a way to gracefully handle these situations. If your program panics, be very careful about trying to continue executing after the panic. It’s very rare that you want to keep your program running after a panic occurs.   

➤ 一般返回 error 而不是 panic

In the preceding sample program, it would be idiomatic to check for division by zero and return an error if one was passed in.  既然能预料到除零会导致 panic,  就应该校验参数是否为 0 并返回错误值.

### 什么时候适合 recover

➤ 用于避免 panic 泄漏给调用者

There is one situation where `recover` is recommended. If you are creating a library for third parties, do not let panics escape the boundaries of your public API. If a panic is possible, a public function should use a recover to convert the panic into an error, return it, and let the calling code decide what to do with them.  

➤ 用于避免 goroutine 中的 panic 导致整个应用挂掉

One application of `recover` is to shut down a failing goroutine inside a server without killing the other executing goroutines. In this example, if `do(work)` panics, the result will be logged and the goroutine will exit cleanly without disturbing the others.

```go
func server(workChan <-chan *Work) {
    for work := range workChan {
        go safelyDo(work)
    }
}

func safelyDo(work *Work) {
    defer func() {
        if err := recover(); err != nil {
            log.Println("work failed:", err)
        }
    }()
    do(work)
}
```

### 用 recover 简化错误处理

➤ 有时候逐层返回错误会让人无法忍受

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
    // doParse() 中可以用 regexp.error("xxx error") 抛出错误,  而不是逐层返回错误值
    return regexp.doParse(str), nil
}
```

➤ 短短几行代码,  却并不简单 (非常巧妙):

If `doParse` panics, the recovery block will set the return value to `nil`—deferred functions can modify named return values. It will then check, in the assignment to `err`, that the problem was a parse error by asserting that it has the local type `Error`. If it does not, the type assertion will fail, causing a run-time error that continues the stack unwinding as though nothing had interrupted it. This check means that if something unexpected happens, such as an index out of bounds, the code will fail even though we are using `panic` and `recover` to handle parse errors.

➤ re-panic 修改了 panic value,  但也没什么大问题:

By the way, this re-panic idiom changes the panic value if an actual error occurs. However, both the original and new failures will be presented in the crash report, so the root cause of the problem will still be visible (崩溃报告会打印整个 panic 链). Thus this simple re-panic approach is usually sufficient—it's a crash after all—but if you want to display only the original value, you can write a little more code to filter unexpected problems and re-panic with the original error. 

➤ 若不想修改 panic value 可以这样做:

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

