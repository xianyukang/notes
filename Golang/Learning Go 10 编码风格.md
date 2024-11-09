## Table of Contents
  - [Naming in Go](#Naming-in-Go)
    - [概述](#%E6%A6%82%E8%BF%B0)
    - [局部变量要短](#%E5%B1%80%E9%83%A8%E5%8F%98%E9%87%8F%E8%A6%81%E7%9F%AD)
    - [变量起名规则](#%E5%8F%98%E9%87%8F%E8%B5%B7%E5%90%8D%E8%A7%84%E5%88%99)
    - [拒绝命名重复](#%E6%8B%92%E7%BB%9D%E5%91%BD%E5%90%8D%E9%87%8D%E5%A4%8D)
    - [一些规则](#%E4%B8%80%E4%BA%9B%E8%A7%84%E5%88%99)
    - [构造函数](#%E6%9E%84%E9%80%A0%E5%87%BD%E6%95%B0)
    - [关于下划线](#%E5%85%B3%E4%BA%8E%E4%B8%8B%E5%88%92%E7%BA%BF)
    - [Interface](#Interface)
    - [Package Name](#Package-Name)
  - [SOLID Principle](#SOLID-Principle)
    - [Single Responsibility](#Single-Responsibility)
    - [Open Closed Principle](#Open-Closed-Principle)
    - [Interface Segregation Principle](#Interface-Segregation-Principle)
    - [Dependency Inversion Principle](#Dependency-Inversion-Principle)
  - [编码风格](#%E7%BC%96%E7%A0%81%E9%A3%8E%E6%A0%BC)
    - [Formatting](#Formatting)
    - [Commentary](#Commentary)
    - [var vs :=](#var-vs-)
    - [Literal Formatting](#Literal-Formatting)
    - [Nil Slice](#Nil-Slice)
    - [Copying](#Copying)
    - [Doc Comment](#Doc-Comment)
  - [Coding Style](#Coding-Style)
    - [Interfaces](#Interfaces)
    - [性能小提示](#%E6%80%A7%E8%83%BD%E5%B0%8F%E6%8F%90%E7%A4%BA)
    - [记得拷贝 slice/map](#%E8%AE%B0%E5%BE%97%E6%8B%B7%E8%B4%9D-slicemap)
    - [让枚举值从 1 开始](#%E8%AE%A9%E6%9E%9A%E4%B8%BE%E5%80%BC%E4%BB%8E-1-%E5%BC%80%E5%A7%8B)
    - [推荐加上 json tag](#%E6%8E%A8%E8%8D%90%E5%8A%A0%E4%B8%8A-json-tag)
    - [避免使用全局"变量"](#%E9%81%BF%E5%85%8D%E4%BD%BF%E7%94%A8%E5%85%A8%E5%B1%80%E5%8F%98%E9%87%8F)
    - [用依赖注入避免全局变量](#%E7%94%A8%E4%BE%9D%E8%B5%96%E6%B3%A8%E5%85%A5%E9%81%BF%E5%85%8D%E5%85%A8%E5%B1%80%E5%8F%98%E9%87%8F)
    - [不要滥用 Embedding](#%E4%B8%8D%E8%A6%81%E6%BB%A5%E7%94%A8-Embedding)
    - [少用 else, 多用 if-return](#%E5%B0%91%E7%94%A8-else-%E5%A4%9A%E7%94%A8-ifreturn)
    - [命名返回值的两个作用](#%E5%91%BD%E5%90%8D%E8%BF%94%E5%9B%9E%E5%80%BC%E7%9A%84%E4%B8%A4%E4%B8%AA%E4%BD%9C%E7%94%A8)
    - [方法推荐用 pointer receiver](#%E6%96%B9%E6%B3%95%E6%8E%A8%E8%8D%90%E7%94%A8-pointer-receiver)
    - [Context](#Context)
    - [Crypto Rand](#Crypto-Rand)
    - [关于编写测试的建议](#%E5%85%B3%E4%BA%8E%E7%BC%96%E5%86%99%E6%B5%8B%E8%AF%95%E7%9A%84%E5%BB%BA%E8%AE%AE)
  - [Go Doc Comments](#Go-Doc-Comments)
    - [Introduction](#Introduction)
    - [Packages](#Packages)
    - [Commands](#Commands)
    - [Types](#Types)
    - [Funcs](#Funcs)
    - [Consts and Vars](#Consts-and-Vars)
    - [Syntax](#Syntax)

#### ➤ 参考资料

- [Effective Go](https://go.dev/doc/effective_go)

- [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)

- [Go Styleguide](https://github.com/bahlo/go-styleguide)

- [Uber Go Style Guide](https://github.com/uber-go/guide/blob/master/style.md#uber-go-style-guide)

#### ➤ 为什么要学习 Go 的编码风格

A straightforward translation of a C++ or Java program into Go is unlikely to produce a satisfactory result—Java programs are written in Java, not Go. In other words, to write Go well, it's important to understand its properties and idioms. It's also important to know the established conventions for programming in Go, such as naming, formatting, program construction, and so on, so that programs you write will be easy for other Go programmers to understand.

## Naming in Go

参考 [Meetup Golang Paris - What's in a name?](https://www.youtube.com/watch?v=sFUSP8Au_PE)

参考 [Google Go Style Guide - Naming](https://google.github.io/styleguide/go/decisions#naming)

参考 [Package Names - The Go Blog](https://go.dev/blog/package-names)

### 概述

Good names are critical to redability. A good name is:

- Consistent ( easy to guess )
- Short ( easy to type )
- Accurate ( easy to understand )
- 一致性是指遵守同一个模式/套路，比如 `json.Marshal` 和 `yaml.Marshal`，以及 `atomic.StoreInt32` 和 `atomic.StoreInt64`

Some rules: 

- Use short names. 
- Think about context. 
- Avoid repetitive names.

### 局部变量要短

#### ➤ 作用域越小，使用次数越多，那么变量名越短

The general rule of thumb is that the length of a name should be proportional to the size of its scope and inversely proportional to the number of times that it is used within that scope. A variable created at file scope may require multiple words, whereas a variable scoped to a single inner block may be a single word or even just a character or two, to keep the code clear and avoid extraneous information. The greater the distance between a name's declaration and its uses, the longer the name should be. If you declare a name on one line and you only use it on the next line then it should be really short. 

#### ➤ Common variable may use really short names. Using familiar variable names for common types is often helpful

- `r` for an `io.Reader` or `*http.Request`
- `w` for an `io.Writer` or `http.ResponseWriter`

```go
// Prefer i to index. 
// Prefer r to reader. 
// Prefer b to buffer. 
// A name might be perfectly clear within a small scope.
for _, n := range nodes {
    fmt.Println(n.Name)
}
```

#### ➤ 用 cur 还是 curr 作为 current 的缩写? [代码常用单词的缩写](https://github.com/kisvegabor/abbreviations-in-code)

#### ➤ Local variables should be short. Long names obscure what the code does. 

```go
// Avoid redundant names, given their context: 
// Prefer count to runeCount inside a function named RuneCount.

// Bad: 长的名字让代码看起来很复杂的样子
func RuneCount(buffer []byte) int {
    runeCount := 0
    for index := 0; index < len(buffer); {
        if buffer[index] < RuneSelf {
            index++
        } else {
            _, size := DecodeRune(buffer[index:])
            index += size
        }
        runeCount++
    }
    return runeCount
}

// Good: 短的名字让代码看起来更简洁
func RuneCount(b []byte) int {
    count := 0
    for i := 0; i < len(b); {
        if b[i] < RuneSelf {
            i++
        } else {
            _, n := DecodeRune(b[i:])
            i += n
        }
        count++
    }
    return count
}
```

#### ➤ Function parameters are like local variables, but they also serve as documentation. 

```go
// Where the types are descriptive, they should be short:
func AfterFunc(d Duration, f func()) *Timer

// Where the types are more ambiguous, the names may provide documentation:
func HasPrefix(s, prefix []byte) bool
func Unix(sec, nsec int64) Time
```

#### ➤ Receivers are a special kind of argument. By convention, they are one or two characters that reflect the receiver type, because they typically appear on almost every line:

```go
// Receivers use one or two characters
func (b *Reader) Read(p []byte) (n int, err error) {
}
```

#### ➤ 一般不用命名返回值，但它有两个作用: ①偶尔能提高可读性 ②在 defer 函数中修改返回值

```go
// The name of the function itself and the type of the result parameters are often sufficiently clear.
func (n *Node) Parent1() *Node
func (n *Node) Parent2() (*Node, error)

// If a function returns two or more parameters of the same type, adding names can be useful.
func (n *Node) Children() (left, right *Node, err error)
```

If the caller must take action on particular result parameters, naming them can help suggest what the action is:

```go
// WithTimeout returns a context that will be canceled no later than d duration from now.
//
// The caller must arrange for the returned cancel function to be called when
// the context is no longer needed to prevent a resource leak.
func WithTimeout(parent Context, d time.Duration) (ctx Context, cancel func())
```

In the code above, cancellation is a particular action a caller must take. However, were the result parameters written as `(Context, func())` alone, it would be unclear what is meant by “cancel function”.

### 变量起名规则

In general:

- Single-word names like `count` or `options` are a good starting point.
- Additional words can be added to disambiguate similar names, for example `userCount` and `projectCount`.
- Omit types and type-like words from most variable names.
  - For a number, `userCount` is a better name than `numUsers` or `usersInt`.
  - For a slice, `users` is a better name than `userSlice`.
  - It is acceptable to include a type-like qualifier if there are two versions of a value in scope, for example you might have an input stored in `ageString` and use `age` for the parsed value.
- Omit words that are clear from the surrounding context. For example, in the implementation of a `UserCount` method, a local variable called `userCount` is probably redundant; `count`, `users`, or even `c` are just as readable.

### 拒绝命名重复

A piece of Go source code should avoid unnecessary repetition. One common source of this is repetitive names, which often include unnecessary words or repeat their context or type. Repetitive naming can come in many forms, including:

#### ➤ Package vs. exported symbol name

Exported names are qualified by their package names. Remember this when naming exported variables, functions, and types. That's why we have `bytes.Buffer / strings.Reader`, not `bytes.ByteBuffer / strings.StringReader`. **Examples:** Repetitive Name -> Better Name:

- widget.NewWidget -> widget.New
- widget.NewWidgetWithName -> widget.NewWithName
- db.LoadFromDatabase -> db.Load

#### ➤ Variable name vs. type

In most cases it is clear to the reader what type a variable is by how it is used. It is only necessary to clarify the type of a variable if its value appears twice in the same scope.

```go
// Repetitive Name                  Better Name
var numUsers int                  var users int
var nameString string              var name string
var primaryProject *Project        var primary *Project
```

If the value appears in multiple forms, this can be clarified either with an extra word like `raw` and `parsed` or with the underlying representation:

```go
// Good:
limitStr := r.FormValue("limit")
limit, err := strconv.Atoi(limitStr)
// Good:
limitRaw := r.FormValue("limit")
limit, err := strconv.Atoi(limitRaw)
```

#### ➤ External context vs. local names

Names that include information from their surrounding context often create extra noise without benefit. The package name, method name, type name, function name, import path, and even filename can all provide context that automatically qualifies all names within.

```go
// In package "ads/targeting/revenue/reporting"
type AdsTargetingRevenueReport struct{} // Bad
type Report struct{}                    // Good

func (p *Project) ProjectName() string  // Bad
func (p *Project) Name() string         // Good

// In package "sqldb"
type DBConnection struct{} // Bad
type Connection struct{}   // Good
```

Names that are clear from context or usage can often be omitted:

```go
// Bad:
func (db *DB) UserCount() (userCount int, err error) {
    var userCountInt64 int64
    if dbLoadError := db.LoadFromDatabase("count(distinct users)", &userCountInt64); dbLoadError != nil {
        return 0, fmt.Errorf("failed to load user count: %s", dbLoadError)
    }
    userCount = int(userCountInt64)
    return userCount, nil
}

// Good:
func (db *DB) UserCount() (int, error) {
    var count int64
    if err := db.Load("count(distinct users)", &count); err != nil {
        return 0, fmt.Errorf("failed to load user count: %s", err)
    }
    return int(count), nil
}
```

### 一些规则

(1) Words in names that are initialisms or acronyms should have the same case. URL should appear as `URL` or `url` (as in urlPony, or URLPony), never as `Url`. This also applies to `ID` when it is short for “identifier”; write `appID` instead of `appId`.

(3) Errors:

```go
// Error types should be of the form FooError:
type ExitError struct {
    ...
}
// Error values should be of the form ErrFoo:
var ErrFormat = errors.New("image: unknown format")
```

(3) If you have a field called `owner` (lower case, unexported), the getter method should be called `Owner` (upper case, exported), not `GetOwner`. A setter function, if needed, will likely be called `SetOwner`. 

If the function involves performing a complex computation or executing a remote call, a different word like `Compute` or `Fetch` can be used in place of `Get`, to make it clear to a reader that the function call may take time and could block or fail.

### 构造函数

Golang 中的构造函数是普通函数，用 `New` 作为前缀，例如 `list.NewList()`，但此处和包名产生重复，所以叫 `list.New` 就行

```go
// A function named New in package pkg returns a value of type pkg.Pkg
q := list.New()  // q is a *list.List
```

**Simplify function names.** When a function in package pkg returns a value of type `pkg.Pkg` (or `*pkg.Pkg`), the function name can often omit the type name without confusion:  

```go
start := time.Now()                                  // 不一定要用 New, 显然 Now 的内涵/可读性比 New 更好
t, err := time.Parse(time.Kitchen, "6:06PM")         // Parse 算经典名称了, 根据字符串创建实例
ctx = context.WithTimeout(ctx, 10*time.Millisecond)  // WithXXX 常用于不可变对象, 返回一个副本, 但修改了 XXX 字段
ip, ok := userip.FromContext(ctx)                    // ip is a net.IP
```

When a function returns a value of type `pkg.T`, where `T` is not `Pkg`, the function name may include `T` to make client code easier to understand.

```go
d, err := time.ParseDuration("10s")  // d is a time.Duration
elapsed := time.Since(start)         // elapsed is a time.Duration
ticker := time.NewTicker(d)          // ticker is a *time.Ticker
timer := time.NewTimer(d)            // timer is a *time.Timer
```

### 关于下划线

#### ➤ 一般不用下划线

1. Go source code uses `MixedCaps` (camel case) rather than underscores (snake case) when writing multi-word names.
2. Names in Go should in general not contain underscores. There are some exceptions:
   - Test, Benchmark and Example function names within `*_test.go` files may include underscores.
   - Package names that are only imported by generated code may contain underscores.

#### ➤ 文件和文件夹用小写字母和下划线，用子文件夹分隔包名，例如 `foo/bar` 而不是 `foo_bar`

1. All filenames should be lowercase.
2. Go source files and directories use underscores, not dashes.
3. Package directories should generally avoid using separators as much as possible. When package names are multiple words, they usually should be in nested subdirectories.

### Interface

Interfaces that specify just one method are usually just that function name with 'er' appended to it.

```go
type Reader interface {
    Read(p []byte) (n int, err error)
}
// Sometimes the result isn't correct English, but we do it anyway:
type Execer interface {
    Exec(query string, args []Value) (Result, error)
}
// Sometimes we use English to make it nicer:
type ByteReader interface {
    ReadByte() (c byte, err error)
}
// When an interface includes multiple methods, choose a name that accurately describes its purpose
// ( examples: net.Conn, http.ResponseWriter, io.ReadWriter ).

```

- 若不想实现 Reader/Writer/Stringer 接口，那么别用 Read/Write/String 之类的方法名，会让人误以为你想实现 Reader
- 另一方面，如果方法的作用和意义有相应的接口，应该积极实现，例如转成字符串的方法叫 `String` 而不是 `ToString`

### Package Name

Package names are central to good naming in Go programs. Take the time to choose good package names and organize your code well. This helps clients understand and use your packages and helps maintainers to grow them gracefully. The name of the package is a critical piece of its design. Work to eliminate meaningless package names from your projects.

#### ➤ 注意事项

1. 包名应该短小简洁、尽量使用单个单词、用单数形式 (strings/bytes算例外)，全用小写 (例如 `usercount`)
2. 避免常用变量名，比如 `count` 是常用的变量名，当包名和变量名冲突时会让人不爽，推荐用 `usercount` 替代
3. 有时候可以用缩写。Widely-used packages often have compressed names:
   - `strconv` (string conversion)
   - `syscall` (system call)
   - `fmt` (formatted I/O)

#### ➤ 拆成具体的小包

- 避免无意义的包名，例如 util, common, helper, ... 这类包名相当于啥也没说，其中的内容也是乱七八糟，没有分类
- 推荐把泛用包拆成具体的小包
  - 例如 `once.Do()` 比 `util.DoOrWaitUntilDone()` 更好
  - 例如 `stringset.New()` 比 `util.NewStringSet()` 更好

**Avoid meaningless package names.** Packages named `util`, `common`, or `misc` provide clients with no sense of what the package contains. This makes it harder for clients to use the package and makes it harder for maintainers to keep the package focused. Over time, they accumulate dependencies that can make compilation significantly and unnecessarily slower, especially in large programs. And since such package names are generic, they are more likely to collide with other packages imported by client code, forcing clients to invent names to distinguish them.

**Break up generic packages.** To fix such packages, look for types and functions with common name elements and pull them into their own package. For example, if you have:

```go
package util
func NewStringSet(...string) map[string]bool {...}
func SortStringSet(map[string]bool) []string {...}

set := util.NewStringSet("c", "a", "b")
fmt.Println(util.SortStringSet(set))
```

Pull these functions out of `util` into a new package, choosing a name that fits the contents:

```go
package stringset
func New(...string) map[string]bool {...}
func Sort(map[string]bool) []string {...}

set := stringset.New("c", "a", "b")
fmt.Println(stringset.Sort(set))
```

**Don’t use a single package for all your APIs.** Many well-intentioned programmers put all the interfaces exposed by their program into a single package named `api`, `types`, or `interfaces`, thinking it makes it easier to find the entry points to their code base. This is a mistake. Such packages suffer from the same problems as those named `util` or `common`, growing without bound, providing no guidance to users, accumulating dependencies, and colliding with other imports. Break them up, perhaps using directories to separate public packages from implementation.

**Avoid unnecessary package name collisions.** While packages in different directories may have the same name, packages that are frequently used together should have distinct names. This reduces confusion and the need for local renaming in client code. For the same reason, avoid using the same name as popular standard packages like `io` or `http`.

#### ➤ 利用包名信息，拒绝重复

- 使用包中内容时总会带上包名，命名时应该利用包名信息，比如叫 `bufio.Reader` 就好、而 `bufio.BufReader` 是重复的
- **Avoid repetition.** Since client code uses the package name as a prefix when referring to the package contents, the names for those contents need not repeat the package name. The HTTP server provided by the `http` package is called `Server`, not `HTTPServer`.
- Types in different packages can have the same name, because from the client’s point of view such names are discriminated by the package name. For example, the standard library includes several types named `Reader`, including `jpeg.Reader`, `bufio.Reader`, and `csv.Reader`. Each package name fits with `Reader` to yield a good type name.

## SOLID Principle

参考 [Golang UK Conference 2016 - Dave Cheney - SOLID Go Design](https://www.youtube.com/watch?v=zzAdEt3xZ1M)

### Single Responsibility

- A class or module should have only one reason to change.
  - 意思是单一职责，让类更加简单和稳定，易于理解和修改
  - 单一职责的东西更易于重用，通过组装简单的零件，可构件复杂的应用

- When your code does have to change, it should do so for it's own reason.
  - 意思是代码要稳定、要抽象，不要因为依赖的东西发生了变化，自己也要跟着变
  - ~~程序员，你也不想自己的代码经常被迫修改吧 🐶~~

### Open Closed Principle

Software entities should be open for extension but closed for modification.  
This means that the behaviour of a module can be extended without modifying its source code. 

- Open 是指代码要易于扩展，想加点东西或功能，能轻松实现
- Closed 意思是源码要稳定，允许在不改变它的源代码的前提下变更它的行为

### Interface Segregation Principle

Clients should not be forced to depend on interfaces they do not use.  
The dependency of one class to another one should depend on the smallest possible interface.

- 意思是接口要尽量小，两个东西之间的协作接口越简洁越好，这样能降低耦合
- 比如 `Save(f *os.File, doc *Document)` 使用 `*os.File` 作为入参，不如改成 `io.Writer`

>  A great rule of thumb for Go is accept interfaces, return structs. -- Jack Lindamood

### Dependency Inversion Principle

在传统的应用架构中，高层次的组件直接依赖于低层次的组件去实现一些任务。  
这会导致上层和下层产生耦合，限制了高层次组件被重用的可行性。

依赖反转原则规定:

- Abstractions should not depend on details. 
- High-level modules should not import anything from low-level modules.
- 上层模块不应该依赖下层模块，上层模块应该定义一些抽象接口，然后依赖这些抽象接口
- 让下层组件实现上层组件定义的抽象接口，也就是让下层组件依赖上层接口，从而形成依赖反转

举个例子:

- 图 1 中，高层对象 A 依赖于底层对象 B 的实现；
- 图 2 中，把高层对象 A 对底层对象的需求抽象为一个接口 A，底层对象 B 实现了接口 A，这就是依赖反转。
- 总之就是 A 不直接调用 B，而是提供一个用于注入依赖的函数，方便以后切换各种各样的 B 的实现

![img](https://upload.wikimedia.org/wikipedia/commons/9/96/Dependency_inversion.png) 


## 编码风格

### Formatting

- We use tabs for indentation and `gofmt` emits them by default. Use spaces only if you must.
- Go has no line length limit. If a line feels too long, it should be refactored instead of broken.
- Go needs fewer parentheses than C and Java: control structures (`if`, `for`, `switch`) do not have parentheses in their syntax. Also, the operator precedence hierarchy is shorter and clearer, so `x<<8 + y<<16` means what the spacing implies, unlike in the other languages.

### Commentary

Go provides C-style `/* */` block comments and C++-style `//` line comments. Line comments are the norm; block comments appear mostly as package comments, but are useful within an expression or to disable large swaths of code.

Comments that appear before top-level declarations, with no intervening newlines, are considered to document the declaration itself. These “doc comments” are the primary documentation for a given Go package or command. For more about doc comments, see “[Go Doc Comments](https://go.dev/doc/comment)”.

### var vs :=

Short variable declarations (`:=`) should be used if a variable is being set to some value explicitly. However, there are cases where the default value is clearer when the `var` keyword is used.

```go
func main() {
    s2 := []int{1, 2} // 若需要显式设置值, 用 :=
    var s1 []int      // 若使用变量的零值, 用 var

    var s []int       // 推荐用 nil 切片, 一般不用空切片
    if len(s) == 0 {  // 用 len 判断空切片,  别用 s == nil
    }

    // 能写成一行, 则推荐写成一行,  拆成两行有点啰嗦
    if err := os.WriteFile(name, data, 0644); err != nil {
        return err
    }

    // 通常用 &T{} 而不是 new(T) 获取指针,  因为这样一致性更好,  看起来更整齐
    c1 := &Character{}
    c2 := Character{Name: "Homura"}

    // 注意 nil map 很危险,  记得用 make 进行初始化,  可选提供 capacity hint
    m := make(map[string]int, 666)
}
```

### Literal Formatting

Go has an exceptionally powerful [composite literal syntax](https://golang.org/ref/spec#Composite_literals), with which it is possible to express deeply-nested, complicated values in a single expression. Where possible, this literal syntax should be used instead of building values field-by-field. Struct literals should usually specify **field names**. The position of fields in a struct and the full set of fields are not usually considered to be part of a struct’s public API; specifying the field name is needed to avoid unnecessary coupling.

```go
good := otherpkg.Type{A: 42}                             // Good
r := csv.Reader{',', '#', 4, false, false, false, false} // Bad, 改了字段顺序或增删字段都会让当前写法失效
```

### Nil Slice

For most purposes, there is no functional difference between `nil` and the empty slice. Built-in functions like `len` and `cap` behave as expected on `nil` slices. If you declare an empty slice as a local variable (especially if it can be the source of a return value), prefer the nil initialization to reduce the risk of bugs by callers. 

```go
func Foo() ([]int, []int) {
    var s1 []int  // Good
    s2 := []int{} // Bad: 万一 caller 没用 len(s) 判断空切片, 可能出错
    return s1, s2
}
```

Do not create APIs that force their clients to make distinctions between nil and the empty slice. When designing interfaces, avoid making a distinction between a `nil` slice and a non-nil, zero-length slice, as this can lead to subtle programming errors. This is typically accomplished by using `len` to check for emptiness, rather than `== nil`.

### Copying

To avoid unexpected aliasing and similar bugs, be careful when copying a struct. For example, synchronization objects such as `sync.Mutex` must not be copied. If you copy a `bytes.Buffer`, the slice in the copy may alias the array in the original, causing subsequent method calls to have surprising effects. In general, do not copy a value of type `T` if its methods are associated with the pointer type, `*T`.

```go
// Bad:
var b1 bytes.Buffer
b2 := b1
```

Invoking a method that takes a value receiver can hide the copy. When you author an API, you should generally take and return pointer types if your structs contain fields that should not be copied.

```go
type Record struct {
  buf bytes.Buffer
}

// Good:
func New() *Record {...}
func (r *Record) Process(...) {...}
func Consumer(r *Record) {...}

// Bad:
func (r Record) Process(...) {...} // Makes a copy of r.buf
func Consumer(r Record) {...}      // Makes a copy of r.buf

// 既然 Record 中包含不想被 copy 的 bytes.Buffer
// 那就应该避免 copy, 例如方法使用 *Record 作为接收器, 函数使用 *Record 作为参数
```

### Doc Comment

1. Comments should begin with the name of the thing being described and end in a period. 

2. All top-level, exported names should have doc comments, as should non-trivial unexported type or function declarations. 
3. When adding a new package, include examples of intended usage: a runnable Example, or a simple test demonstrating a complete call sequence. Read more about [testable Example() functions](https://go.dev/blog/examples).

```go
// Request represents a request to run a command.
type Request struct { ...
// Encode writes the JSON encoding of req to w.
func Encode(w io.Writer, req *Request) { ...
```


## Coding Style

### Interfaces

#### ➤ 显示声明要实现的接口

```go
var _ http.Handler = LogHandler{}
var _ http.Handler = (*Handler)(nil)
// The statement will fail to compile if *Handler ever stops matching the http.Handler interface.
```

#### ➤ 接口实现方返回具体类型

Go interfaces generally belong in the package that uses values of the interface type, not the package that implements those values. The implementing package should return concrete (usually pointer or struct) types: that way, new methods can be added to implementations without requiring extensive refactoring.

```go
// 接口属于消费者而不是实现者，在消费端定义接口，
// 实现端返回具体类型而不是接口类型，这样在 implemention 中添加更多方法，也不会影响消费者
package consumer
type Thinger interface { Thing() bool }
func Foo(t Thinger) string { ... }
```

#### ➤ 接口要有意义，被用到

1. Do not define interfaces before they are used. It is too difficult to see whether an interface is even necessary.
2. Do not use interface-typed parameters if the users of the package do not need to pass different types for them.
3. Do not export interfaces that the users of the package do not need.

### 性能小提示

1. When converting primitives to strings, `strconv` is faster than `fmt`: `strconv.Itoa(123)`
2. Where possible, provide capacity hints when initializing maps with make: `make(map[T1]T2, hint)`
3. Where possible, provide capacity when initializing slices with make: `make([]T, length, capacity)`

### 记得拷贝 slice/map

#### ➤ [示例](https://github.com/uber-go/guide/blob/master/style.md#copy-slices-and-maps-at-boundaries)

Slices and maps contain pointers to the underlying data so be wary of scenarios when they need to be copied. Keep in mind that users can modify a map or slice you received as an argument if you store a reference to it.

### 让枚举值从 1 开始

#### ➤ [示例](https://github.com/uber-go/guide/blob/master/style.md#start-enums-at-one)

Since variables have a 0 default value, you should usually start your enums on a non-zero value. There are cases where using the zero value makes sense, for example when the zero value case is the desirable default behavior.

### 推荐加上 json tag

Any struct field that is marshaled into JSON, YAML, or other formats that support tag-based field naming should be annotated with the relevant tag. The serialized form of the structure is a contract between different systems. Specifying field names inside tags makes the contract explicit, and it guards against accidentally breaking the contract by refactoring or renaming fields.

```go
type Stock struct {
  Price int    `json:"price"`
  Name  string `json:"name"`     // Safe to rename Name to Symbol.
}
```

### 避免使用全局"变量"

You should rarely declare variables outside of functions, in what’s called the `package block`. *Package-level variables whose values change are a bad idea*. When you have a variable outside of a function:

1. It can be difficult to track the changes made to it, 
2. which makes it hard to understand how data is flowing through your program. 
3. This can lead to subtle bugs.

As a general rule, you should only declare variables in the package block that are *effectively immutable*. Avoid declaring variables outside of functions because they *complicate data flow analysis*.

If you have package-level variables that need to be modified while your program is running, see if you can refactor your code to put that state into a struct that’s initialized and returned by a function in the package.  

### 用依赖注入避免全局变量

```go
func main() {
    db := //
    handlers := Handlers{DB: db}
    http.HandleFunc("/drop", handlers.DropHandler)
}
```

### 不要滥用 Embedding

#### ➤ Type embedding in Go

Go has the ability to “borrow” pieces of an implementation by embedding types. The outer type gets implicit copies of the embedded type's methods. These methods, by default, delegate to the same method of the embedded instance. Embedding `company.Employee` will generate a field named as `Employee` in the outer struct. So, if the embedded type is public, the field is public. To maintain backward compatibility, every future version of the outer type must keep the embedded type.

#### ➤ The embedded type places limits on the evolution of the type.

- Adding methods to an embedded interface is a breaking change.
- Removing methods from an embedded struct is a breaking change.
- Removing the embedded type is a breaking change.

#### ➤ Avoid Embedding Types in Public Structs

These embedded types leak implementation details, inhibit type evolution. Avoid embedding the `AbstractList` in your concrete list implementations. Instead, hand-write only the methods to your concrete list that will delegate to the abstract list. Although writing these delegate methods is tedious, the additional effort hides an implementation detail, leaves more opportunities for change.

#### ➤ 总结、总而言之

Simply put, embed consciously and intentionally. A good litmus test is, "would all of these exported inner methods/fields be added directly to the outer type"; if the answer is "some" or "no", don't embed the inner type - use a field instead.

### 少用 else, 多用 if-return

For instance, don't write:

```go
if err != nil {
    // error handling
} else {
    // normal code
}
```

Instead, write:

```go
if err != nil {
    // error handling
    return
}
// normal code
// 好处有两个，①减少嵌套层数 ②让错误情况被优先处理，后续的代码无需顾虑错误情况

// 另外可以用默认值减少 else 的使用
xxxOption := ...
if c {
    xxxOpton = ...
}
```

### 命名返回值的两个作用

(1) 偶尔能增加可读性

```go
func (f *Foo) Location() (float64, float64, error)      // 很模糊
// Location returns f's latitude and longitude.
// Negative values mean south and west, respectively.
func (f *Foo) Location() (lat, long float64, err error) // 加上了注释和返回值名字、更清晰
```

(2) 另一个作用是在 deferred closure 中修改函数返回值

### 方法推荐用 pointer receiver

Choosing whether to use a value or pointer receiver on methods can be difficult, especially to new Go programmers. If in doubt, use a pointer, but there are times when a value receiver makes sense, usually for reasons of efficiency, such as for small unchanging structs or values of basic type. 

#### ➤ First:

- When in doubt, use a pointer receiver.
- Don't mix receiver types. Choose either pointers or struct types for all available methods.

#### ➤ Some useful guidelines:

- If the receiver is a map, func or chan, don't use a pointer to them. If the receiver is a slice and the method doesn't reslice or reallocate the slice, don't use a pointer to it.
- If the method needs to mutate the receiver, the receiver must be a pointer.
- If the receiver is a struct that contains a `sync.Mutex` or similar synchronizing field, the receiver must be a pointer to avoid copying. (因为拷贝一个锁,  锁就失去了意义)
- If the receiver is a large struct or array, a pointer receiver is more efficient. How large is large? Assume it's equivalent to passing all its elements as arguments to the method.
- Can methods be mutating the receiver? A value type creates a copy of the receiver when the method is invoked, so outside updates will not be applied to this receiver. If changes must be visible in the original receiver, the receiver must be a pointer.
- If the receiver is a small array or struct that is naturally a value type (for instance, something like the time.Time type), with no mutable fields and no pointers, or is just a simple basic type such as int or string, a value receiver makes sense. A value receiver can reduce the amount of garbage that can be generated; if a value is passed to a value method, an on-stack copy can be used instead of allocating on the heap.

### Context

Go programs pass Contexts explicitly along the entire function call chain. Most functions that use a `context.Context` should accept it as their first parameter: `func F(ctx context.Context, ...) {}`. Values of the [`context.Context`](https://pkg.go.dev/context) type carry security credentials, tracing information, deadlines, and cancellation signals across API and process boundaries. Unlike C++ and Java, which in the Google codebase use thread-local storage, Go programs pass contexts explicitly along the entire function call chain. 

1. Don't add a Context member to a struct type; instead add a `ctx` parameter to each method on that type that needs to pass it along.
2. Contexts are immutable, so it's fine to pass the same ctx to multiple calls that share the same deadline, cancellation signal, credentials, parent trace, etc.
3. It is very rare for code in the middle of a callchain to require creating a base context of its own using `context.Background()`. Always prefer taking a context from your caller, unless it’s the wrong context.

### Crypto Rand

Do not use package `math/rand` to generate keys, even throwaway ones. Unseeded, the generator is completely predictable. Seeded with `time.Nanoseconds()`, there are just a few bits of entropy. Instead, use `crypto/rand`'s Reader, and if you need text, print to hexadecimal or base64:

```go
import (
    "crypto/rand"
    "fmt"
)

func Key() string {
    buf := make([]byte, 16)
    _, err := rand.Read(buf)
    if err != nil {
        panic(err)  // out of randomness, should never happen
    }
    return fmt.Sprintf("%x", buf)
    // or hex.EncodeToString(buf)
    // or base64.StdEncoding.EncodeToString(buf)
}
```

### 关于编写测试的建议

#### ➤ 错误报告要详细

Tests should fail with helpful messages saying what was wrong, with what inputs, what was actually got, and what was expected. It may be tempting to write a bunch of assertFoo helpers, but be sure your helpers produce useful error messages. Assume that the person debugging your failing test is not you, and is not your team. A typical Go test fails like:

```go
// Note that the order here is actual != expected, and the message uses that order too.
if got != tt.want {
    t.Errorf("Foo(%q) = %d; want %d", tt.in, got, tt.want)
}
```

#### ➤ 请掌握 [table-driven test](https://github.com/uber-go/guide/blob/master/style.md#test-tables)

```go
import "github.com/stretchr/testify/assert"
func Add(a, b int) int { return a + b }
func TestAdd(t *testing.T) {
   cases := []struct {
      Name           string
      A, B, Expected int
   }{
      {"一加一", 1, 1, 2},
      {"零加零", 0, 0, 0},
      {"一加负一", 1, -1, 0},
   }
   for _, tc := range cases {
      tc := tc
      t.Run(tc.Name, func(t *testing.T) {
         // t.Parallel()
         assert.Equal(t, tc.Expected, Add(tc.A, tc.B))
      })
   }
}
```

#### ➤ Use an assert library

Using [assert libraries](https://github.com/stretchr/testify) makes your tests more readable, requires less code and provides consistent error output.

```go
import "github.com/stretchr/testify/assert"

func TestAdd(t *testing.T) {
    actual := 2 + 2
    expected := 4
    assert.Equal(t, expected, actual)
}
```

#### ➤ 可以写个 [testString()](https://github.com/bahlo/go-styleguide#avoid-deepequal) 方法比较结构体, 或者用 go-cmp

```go
import "github.com/google/go-cmp/cmp"
func TestGoCmp(t *testing.T) {
   actual := Character{Name: "Ruby", From: "RWBY"}
   expected := Character{Name: "Blake", From: "RWBY"}
   if diff := cmp.Diff(expected, actual); diff != "" {
      t.Error(diff)
   }
}
```

## Go Doc Comments

### Introduction

“Doc comments” are comments that appear immediately before top-level package, const, func, type, and var declarations with no intervening newlines. Every exported (capitalized) name should have a doc comment.

The [go/doc](https://go.dev/pkg/go/doc) and [go/doc/comment](https://go.dev/pkg/go/doc/comment) packages provide the ability to extract documentation from Go source code, and a variety of tools make use of this functionality. The [`go doc io.ReadAll`](https://go.dev/cmd/go#hdr-Show_documentation_for_package_or_symbol) command looks up and prints the doc comment for a given package or symbol. (A symbol is a top-level const, func, type, or var.) The web server [pkg.go.dev](https://pkg.go.dev/) shows the documentation for public Go packages. The program serving that site is [golang.org/x/pkgsite/cmd/pkgsite](https://pkg.go.dev/golang.org/x/pkgsite/cmd/pkgsite), which can also be run locally to view documentation for private modules or without an internet connection.

### Packages

Every package should have a package comment introducing the package. It provides information relevant to the package as a whole and generally sets expectations for the package. Especially in large packages, it can be helpful for the package comment to give a brief overview of the most important parts of the API, linking to other doc comments as needed. For example:

```go
// Package path implements utility routines for manipulating slash-separated
// paths.
//
// The path package should only be used for paths separated by forward
// slashes, such as the paths in URLs. This package does not deal with
// Windows paths with drive letters or backslashes; to manipulate
// operating system paths, use the [path/filepath] package.
package path
```

- The square brackets in `[path/filepath]` create a [documentation link](https://go.dev/doc/comment#links).
- For a package comment, the [first sentence](https://go.dev/pkg/go/doc/#Package.Synopsis) begins with “Package ”.
- For multi-file packages, the package comment should only be in one source file ( e.g., `doc.go` or `main.go` ).

### Commands

A package comment for a command is similar, but it describes the behavior of the program rather than the Go symbols in the package. The first sentence conventionally begins with the name of the program itself, capitalized because it is at the start of a sentence. For example, here is an abridged version of the package comment for [gofmt](https://go.dev/cmd/gofmt):

```go
/*
// 可以随意换行, go doc 和 pkgsite 会去掉换行符, 把这些行连成一个段落.
Gofmt formats Go programs.
It uses tabs for indentation and blanks for alignment.
Alignment assumes that an editor is using a fixed-width font.

// 像下面一样用 tab 缩进, 被缩进的行会在网页中显示为一个 code block
Usage:

    gofmt [flags] [path ...]

The flags are:

    -d
        Do not print reformatted sources to standard output.
        If a file's formatting is different than gofmt's, print diffs
        to standard output.
    -w
        Do not print reformatted sources to standard output.
        If a file's formatting is different from gofmt's, overwrite it
        with gofmt's version. If an error occurred during overwriting,
        the original file is restored from an automatic backup.

When gofmt reads from standard input, it accepts either a full Go program
or a program fragment. A program fragment must be a syntactically
valid declaration list, statement list, or expression.
*/
package main
```

### Types

A type’s doc comment should explain what each instance of that type represents or provides. For example:

```go
package zip

// A Reader serves content from a ZIP archive.
type Reader struct {
    ...
}
```

By default, programmers should expect that a type is safe for use only by a single goroutine at a time. If a type provides stronger guarantees, the doc comment should state them. For example:

```go
package regexp

// Regexp is the representation of a compiled regular expression.
// A Regexp is safe for concurrent use by multiple goroutines,
// except for configuration methods, such as Longest.
type Regexp struct {
    ...
}
```

Go types should also aim to make the zero value have a useful meaning. If it isn’t obvious, that meaning should be documented. For example:

```go
package bytes

// A Buffer is a variable-sized buffer of bytes with Read and Write methods.
// The zero value for Buffer is an empty buffer ready to use.
type Buffer struct {
    ...
}
```

For a struct with exported fields, either the doc comment or per-field comments should explain the meaning of each exported field. For example, this type’s doc comment explains the fields:

```go
package io

// A LimitedReader reads from R but limits the amount of
// data returned to just N bytes. Each call to Read
// updates N to reflect the new amount remaining.
// Read returns EOF when N <= 0.
type LimitedReader struct {
    R   Reader // underlying reader
    N   int64  // max bytes remaining
}
```

In contrast, this type’s doc comment leaves the explanations to per-field comments:

```go
package comment

// A Printer is a doc comment printer.
// The fields in the struct can be filled in before calling
// any of the printing methods
// in order to customize the details of the printing process.
type Printer struct {
    // HeadingLevel is the nesting level used for
    // HTML and Markdown headings.
    // If HeadingLevel is zero, it defaults to level 3,
    // meaning to use <h3> and ###.
    HeadingLevel int
    ...
}
```

### Funcs

A func’s doc comment should explain what the function returns or, for functions called for side effects, what it does. Named arguments or results can be referred to directly in the comment, without any special syntax like backquotes.

```go
package strconv

// Quote returns a double-quoted Go string literal representing s.
// The returned string uses Go escape sequences (\t, \n, \xFF, \u0100)
// for control characters and non-printable characters as defined by IsPrint.
func Quote(s string) string {
    ...
}
```

And:

```go
package os

// Exit causes the current program to exit with the given status code.
// Conventionally, code zero indicates success, non-zero an error.
// The program terminates immediately; deferred functions are not run.
//
// For portability, the status code should be in the range [0, 125].
func Exit(code int) {
    ...
}
```

If a doc comment needs to explain multiple results, naming the results can make the doc comment more understandable, even if the names are not used in the body of the function. For example:

```go
package io

// Copy copies from src to dst until either EOF is reached
// on src or an error occurs. It returns the total number of bytes
// written and the first error encountered while copying, if any.
//
// A successful Copy returns err == nil, not err == EOF.
// Because Copy is defined to read from src until EOF, it does
// not treat an EOF from Read as an error to be reported.
func Copy(dst Writer, src Reader) (n int64, err error) {
    ...
}
```

In the output of `go doc T` command, top-level functions returning a type `T` or pointer `*T`, perhaps with an additional error result, are shown alongside the type `T` and its methods, under the assumption that they are `T`’s constructors.

By default, programmers can assume that a top-level func is safe to call from multiple goroutines; this fact need not be stated explicitly.

On the other hand, as noted in the previous section, using an instance of a type in any way, including calling a method, is typically assumed to be restricted to a single goroutine at a time. If the methods that are safe for concurrent use are not documented in the type’s doc comment, they should be documented in per-method comments. For example:

```go
package sql

// Close returns the connection to the connection pool.
// All operations after a Close will return with ErrConnDone.
// Close is safe to call concurrently with other operations and will
// block until all other operations finish. It may be useful to first
// cancel any used context and then call Close directly after.
func (c *Conn) Close() error {
    ...
}
```

Note that func and method doc comments focus on what the operation returns or does, detailing what the caller needs to know. Special cases can be particularly important to document. For example:

```go
package math

// Sqrt returns the square root of x.
//
// Special cases are:
//
//  Sqrt(+Inf) = +Inf
//  Sqrt(±0) = ±0
//  Sqrt(x < 0) = NaN
//  Sqrt(NaN) = NaN
func Sqrt(x float64) float64 {
    ...
}
```

Doc comments should not explain internal details such as the algorithm used in the current implementation. Those are best left to comments inside the function body. It may be appropriate to give asymptotic time or space bounds when that detail is particularly important to callers. For example:

```go
package sort

// Sort sorts data in ascending order as determined by the Less method.
// It makes one call to data.Len to determine n and O(n*log(n)) calls to
// data.Less and data.Swap. The sort is not guaranteed to be stable.
func Sort(data Interface) {
    ...
}
```

Because this doc comment makes no mention of which sorting algorithm is used, it is easier to change the implementation to use a different algorithm in the future.

### Consts and Vars

Go’s declaration syntax allows grouping of declarations, in which case a single doc comment can introduce a group of related constants, with individual constants only documented by short end-of-line comments. For example:

```go
package scanner

// The result of Scan is one of these tokens or a Unicode character.
const (
    EOF = -(iota + 1)
    Ident
    Int
    Float
    Char
    ...
)
```

Sometimes the group needs no doc comment at all. For example:

```go
package unicode

const (
    MaxRune         = '\U0010FFFF' // maximum valid Unicode code point.
    ReplacementChar = '\uFFFD'     // represents invalid code points.
    MaxASCII        = '\u007F'     // maximum ASCII value.
    MaxLatin1       = '\u00FF'     // maximum Latin-1 value.
)
```

On the other hand, ungrouped constants typically warrant a full doc comment starting with a complete sentence. For example:

```go
package unicode

// Version is the Unicode edition from which the tables are derived.
const Version = "13.0.0"
```

Typed constants are displayed next to the declaration of their type and as a result often omit a const group doc comment in favor of the type’s doc comment. For example:

```go
package syntax

// An Op is a single regular expression operator.
type Op uint8

const (
    OpNoMatch        Op = 1 + iota // matches no strings
    OpEmptyMatch                   // matches empty string
    OpLiteral                      // matches Runes sequence
    OpCharClass                    // matches Runes interpreted as range pair list
    OpAnyCharNotNL                 // matches any character except newline
    ...
)
```

The conventions for variables are the same as those for constants.

### Syntax

详情参考 [Go Doc Comments #Syntax](https://go.dev/doc/comment#:~:text=Arabic%2C%0A%20%20%20%20%22Armenian%22%3A%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20Armenian%2C%0A%20%20%20%20...%0A%7D-,Syntax,-Go%20doc%20comments)

Go doc comments are written in a simple syntax that supports paragraphs, headings, links, lists, and preformatted code blocks. To keep comments lightweight and readable in source files, there is no support for complex features like font changes or raw HTML. Markdown aficionados can view the syntax as a simplified subset of Markdown.

1. `# Numeric Conversions` 就像 Markdown 中的标题，但上下两行需要为空白行 (`#` 语法在 Go 1.19 才加入)
2. `[RFC 7159]` 是一个链接，地址是末尾的 `[RFC 7159]: https://tools.ietf.org/html/rfc7159`
3. `[Atoi]` 也是链接，引用本包的 Atoi 函数，另外可用 `[encoding/json.Decoder]`，引用其他包的东西
4. 用 tab 或两个空格缩进，然后用 `-` 或 `1.` ，随后再加一个空格，就能形成列表
5. 添加一个空行，然后用 tab 缩进接下来的行，就能形成 code block

```go
// Package strconv implements conversions to and from string representations
// of basic data types.
//
// # Numeric Conversions
//
// The most common numeric conversions are [Atoi] (string to int) and [Itoa] (int to string).
...
package strconv

// Package json implements encoding and decoding of JSON as defined in
// [RFC 7159]. The mapping between JSON and Go values is described
// in the documentation for the Marshal and Unmarshal functions.
//
// For an introduction to this package, see the article
// “[JSON and Go].”
//
// [RFC 7159]: https://tools.ietf.org/html/rfc7159
// [JSON and Go]: https://golang.org/doc/articles/json_and_go.html
package json


// PublicSuffixList provides the public suffix of a domain. For example:
//   - the public suffix of "example.com" is "com",
//   - the public suffix of "foo1.foo2.foo3.co.uk" is "co.uk", and
//   - the public suffix of "bar.pvt.k12.ma.us" is "pvt.k12.ma.us".

package sort
// As a more whimsical example, this program guesses your number:
//
//    func GuessingGame() {
//        var s string
//        fmt.Printf("Pick an integer from 0 to 100.\n")
//        answer := sort.Search(100, func(i int) bool {
//            fmt.Printf("Is your number <= %d? ", i)
//            fmt.Scanf("%s", &s)
//            return s != "" && s[0] == 'y'
//        })
//        fmt.Printf("Your number is %d.\n", answer)
//    }
func Search(n int, f func(int) bool) int {
```

