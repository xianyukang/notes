## Table of Contents
  - [错误处理](#%E9%94%99%E8%AF%AF%E5%A4%84%E7%90%86)
    - [错误处理概述](#%E9%94%99%E8%AF%AF%E5%A4%84%E7%90%86%E6%A6%82%E8%BF%B0)
    - [使用字符串表示简单错误](#%E4%BD%BF%E7%94%A8%E5%AD%97%E7%AC%A6%E4%B8%B2%E8%A1%A8%E7%A4%BA%E7%AE%80%E5%8D%95%E9%94%99%E8%AF%AF)
    - [Sentinel Errors](#Sentinel-Errors)
    - [自定义错误类型](#%E8%87%AA%E5%AE%9A%E4%B9%89%E9%94%99%E8%AF%AF%E7%B1%BB%E5%9E%8B)
    - [显式返回 nil 表示没有错误](#%E6%98%BE%E5%BC%8F%E8%BF%94%E5%9B%9E-nil-%E8%A1%A8%E7%A4%BA%E6%B2%A1%E6%9C%89%E9%94%99%E8%AF%AF)
  - [Wrapping Errors](#Wrapping-Errors)
    - [Why Wrapping Errors](#Why-Wrapping-Errors)
    - [The Unwrap method](#The-Unwrap-method)
    - [Wrapping errors with %w](#Wrapping-errors-with-w)
    - [Whether to Wrap](#Whether-to-Wrap)
    - [Examining errors with Is and As](#Examining-errors-with-Is-and-As)
    - [注意 errors.Is 和 errors.As](#%E6%B3%A8%E6%84%8F-errorsIs-%E5%92%8C-errorsAs)
  - [处理 Panic](#%E5%A4%84%E7%90%86-Panic)
    - [什么时候适合 panic](#%E4%BB%80%E4%B9%88%E6%97%B6%E5%80%99%E9%80%82%E5%90%88-panic)
    - [实现 Panic Recovery](#%E5%AE%9E%E7%8E%B0-Panic-Recovery)

## 错误处理

### 错误处理概述

➤ 返回错误值、而不是抛出错误

Go handles errors by returning a value of type `error` as the last return value for a function. The calling function then checks the error return value by comparing it to `nil`, handling the error, or returning an error of its own.  A new error is created from a string by calling the `errors.New` function. Error messages should not be capitalized nor should they end with punctuation or a newline.  

**Go errors are values**. Unlike languages with exceptions, Go doesn’t have special constructs to detect if an error was returned.  `error` is a built-in interface that defines a single method: `Error() string`. Anything that implements this interface is considered an error.   

➤ Golang 这种错误处理风格、背后的考虑

There are two very good reasons why Go uses a returned error instead of thrown exceptions. First, exceptions add at least one new code path through the code. These paths are sometimes unclear, especially in languages whose functions don’t include a declaration that an exception is possible. This produces code that crashes in surprising ways when exceptions aren’t properly handled, or, even worse, code that doesn’t crash but whose data is not properly initialized, modified, or stored.

The second reason is more subtle, but demonstrates how Go’s features work together. The Go compiler requires that all variables must be read. Making errors returned values forces developers to either check and handle error conditions or make it explicit that they are ignoring errors by using an underscore `_` for the returned error value.  

Exception handling may produce shorter code, but having fewer lines doesn’t necessarily make code easier to understand or maintain. As we’ve seen, idiomatic Go favors clear code, even if it takes more lines.  

### 使用字符串表示简单错误

Go’s standard library provides two ways to create an error from a string. The first is the `errors.New` function. It takes in a string and returns an `error`. This string is returned when you call the `Error` method on the returned error instance. If you pass an error to `fmt.Println`, it calls the `Error` method automatically.  The second way is to use the `fmt.Errorf` function. This function allows you to use all of the formatting verbs for `fmt.Printf` to create an error.

### Sentinel Errors

➤ 例子: `var ErrNotExist = errors.New("file does not exist")`

Sentinel errors are one of the few variables that are declared at the package level. By convention, their names start with `Err` (with the notable exception of `io.EOF`).  Be sure you need a sentinel error before you define one. Once you define one, it is part of your public API and you have committed to it being available in all future backward-compatible releases. It’s far better to reuse one of the existing ones in the standard library.

### 自定义错误类型

Since error is an interface, you can define your own errors that include additional information for logging or error handling. For example, you might want to include a status code as part of the error to indicate the kind of error that should be reported back to the user. This lets you avoid string comparisons (whose text might change) to determine error causes.

![image-20220515113743677](https://static.xianyukang.com/img/image-20220515113743677.png) 

### 显式返回 nil 表示没有错误

You shouldn’t declare a variable to be the type of your custom error and then return that variable. The reason why err is non-nil is that `error` is an interface. For an interface to be considered `nil`, both the underlying type and the underlying value must be `nil`.  

```go
type SomeErr []string
func (e SomeErr) Error() string { return e[0] }

func HandleSomething() error {
	var err SomeErr
	fmt.Println(err == nil) // true
	return err              // 返回的 error 接口为 {type: SomeErr, value: nil}
	// return nil           // 如果想表示没有错误,  只能显式返回 nil,  若返回变量会导致 err != nil
}

func main() {
	err := HandleSomething()
	fmt.Println(err == nil) // false,  因为 type 不为空
}
```

There are two ways to fix this. The most common approach is to explicitly return `nil`. Another approach is to make sure that any local variable that holds an error is of type `error`.

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
var myErr MyErr
// Similar to:  if e, ok := err.(*QueryError); ok { … }
// The second is a pointer to a variable of the type that you are looking for
if errors.As(err, &myErr) {
	fmt.Println(myErr.Code)
}
var coder interface {      // 除了用 MyErr 类型的变量,  也能用匿名接口
	Code() int
}
if errors.As(err, &coder) {
	fmt.Println(coder.Code())
}
```

## 处理 Panic

### 什么时候适合 panic

The first class of errors are expected errors that may occur during normal operation.  
比如数据库查询超时、网络不通、错误的用户输入...,  对于这种可以预见的错误要返回错误值而不是 panic

The other class of errors are unexpected errors. These are errors which *should not happen during normal operation*, and if they do it is probably the result of a developer *mistake or a logical error* in your codebase. These errors are truly exceptional, and using panic in these circumstances is more widely accepted. In fact, the Go standard library frequently does this when you make a logical error or try to use the language features in an unintended way — such as when trying to access an out-of-bounds index in a slice, or trying to close an already-closed channel.

### 实现 Panic Recovery

The built-in function `panic` takes one parameter, which can be of any type. Usually, it is a string.  Go provides a way to capture a panic to provide a more graceful shutdown or to prevent shutdown at all. The built-in `recover` function is called from within a `defer` to check if a panic happened. If there was a panic, the value assigned to the panic is returned.

![image-20220516203534170](https://static.xianyukang.com/img/image-20220516203534170.png) 

While panic and recover look a lot like exception handling in other languages, they are not intended to be used that way. The reason we don’t rely on panic and recover is that recover *doesn’t make clear what could fail*. It just ensures that if something fails, we can print out a message and continue. Idiomatic Go favors code that explicitly outlines the possible failure conditions over shorter code that handles anything while saying nothing.  

Reserve panics for fatal situations and use `recover` as a way to gracefully handle these situations. If your program panics, be very careful about trying to continue executing after the panic. It’s very rare that you want to keep your program running after a panic occurs.   

In the preceding sample program, it would be idiomatic to check for division by zero and return an error if one was passed in.  既然能预料到除零会导致 panic,  就应该校验参数是否为 0 并返回错误值.

There is one situation where `recover` is recommended. If you are creating a library for third parties, do not let panics escape the boundaries of your public API. If a panic is possible, a public function should use a recover to convert the panic into an error, return it, and let the calling code decide what to do with them.  
