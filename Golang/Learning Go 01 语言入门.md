## Table of Contents
  - [变量](#%E5%8F%98%E9%87%8F)
    - [变量定义](#%E5%8F%98%E9%87%8F%E5%AE%9A%E4%B9%89)
    - [变量遮蔽](#%E5%8F%98%E9%87%8F%E9%81%AE%E8%94%BD)
    - [显式类型转换](#%E6%98%BE%E5%BC%8F%E7%B1%BB%E5%9E%8B%E8%BD%AC%E6%8D%A2)
    - [Package 变量和 init()](#Package-%E5%8F%98%E9%87%8F%E5%92%8C-init)
    - [Package 的初始化顺序](#Package-%E7%9A%84%E5%88%9D%E5%A7%8B%E5%8C%96%E9%A1%BA%E5%BA%8F)
    - [Package Initialization](#Package-Initialization)
  - [常量](#%E5%B8%B8%E9%87%8F)
    - [字面量没有类型](#%E5%AD%97%E9%9D%A2%E9%87%8F%E6%B2%A1%E6%9C%89%E7%B1%BB%E5%9E%8B)
    - [通过 const 给字面量命名](#%E9%80%9A%E8%BF%87-const-%E7%BB%99%E5%AD%97%E9%9D%A2%E9%87%8F%E5%91%BD%E5%90%8D)
    - [使用 iota 做枚举](#%E4%BD%BF%E7%94%A8-iota-%E5%81%9A%E6%9E%9A%E4%B8%BE)
  - [理解 Comparable](#%E7%90%86%E8%A7%A3-Comparable)
    - [可比较 vs 比大小](#%E5%8F%AF%E6%AF%94%E8%BE%83-vs-%E6%AF%94%E5%A4%A7%E5%B0%8F)
    - [可比较类型](#%E5%8F%AF%E6%AF%94%E8%BE%83%E7%B1%BB%E5%9E%8B)
    - [不可比较类型](#%E4%B8%8D%E5%8F%AF%E6%AF%94%E8%BE%83%E7%B1%BB%E5%9E%8B)
    - [比较 interface 时要注意](#%E6%AF%94%E8%BE%83-interface-%E6%97%B6%E8%A6%81%E6%B3%A8%E6%84%8F)
    - [比较 struct types](#%E6%AF%94%E8%BE%83-struct-types)
  - [各种控制结构](#%E5%90%84%E7%A7%8D%E6%8E%A7%E5%88%B6%E7%BB%93%E6%9E%84)
    - [对比其他语言](#%E5%AF%B9%E6%AF%94%E5%85%B6%E4%BB%96%E8%AF%AD%E8%A8%80)
    - [独特的 if 语句](#%E7%8B%AC%E7%89%B9%E7%9A%84-if-%E8%AF%AD%E5%8F%A5)
    - [For 循环的四种写法](#For-%E5%BE%AA%E7%8E%AF%E7%9A%84%E5%9B%9B%E7%A7%8D%E5%86%99%E6%B3%95)
    - [For 变量可能是同一个](#For-%E5%8F%98%E9%87%8F%E5%8F%AF%E8%83%BD%E6%98%AF%E5%90%8C%E4%B8%80%E4%B8%AA)
    - [独特的 switch 语句](#%E7%8B%AC%E7%89%B9%E7%9A%84-switch-%E8%AF%AD%E5%8F%A5)
    - [Type Assertion](#Type-Assertion)
    - [Type Switch](#Type-Switch)
  - [函数](#%E5%87%BD%E6%95%B0)
    - [参数按值传递](#%E5%8F%82%E6%95%B0%E6%8C%89%E5%80%BC%E4%BC%A0%E9%80%92)
    - [命名参数、可选参数](#%E5%91%BD%E5%90%8D%E5%8F%82%E6%95%B0%E5%8F%AF%E9%80%89%E5%8F%82%E6%95%B0)
    - [变长参数](#%E5%8F%98%E9%95%BF%E5%8F%82%E6%95%B0)
    - [命名返回值](#%E5%91%BD%E5%90%8D%E8%BF%94%E5%9B%9E%E5%80%BC)
    - [函数变量/类型](#%E5%87%BD%E6%95%B0%E5%8F%98%E9%87%8F%E7%B1%BB%E5%9E%8B)
    - [匿名函数](#%E5%8C%BF%E5%90%8D%E5%87%BD%E6%95%B0)
    - [返回清理函数](#%E8%BF%94%E5%9B%9E%E6%B8%85%E7%90%86%E5%87%BD%E6%95%B0)
  - [闭包](#%E9%97%AD%E5%8C%85)
    - [什么是 Closure](#%E4%BB%80%E4%B9%88%E6%98%AF-Closure)
    - [什么是 Method Value](#%E4%BB%80%E4%B9%88%E6%98%AF-Method-Value)
    - [使用 Method Value 的例子](#%E4%BD%BF%E7%94%A8-Method-Value-%E7%9A%84%E4%BE%8B%E5%AD%90)
  - [defer](#defer)
    - [defer 的好处](#defer-%E7%9A%84%E5%A5%BD%E5%A4%84)
    - [defer 前要检查错误](#defer-%E5%89%8D%E8%A6%81%E6%A3%80%E6%9F%A5%E9%94%99%E8%AF%AF)
    - [defer 关闭文件](#defer-%E5%85%B3%E9%97%AD%E6%96%87%E4%BB%B6)
    - [defer 在什么时候执行](#defer-%E5%9C%A8%E4%BB%80%E4%B9%88%E6%97%B6%E5%80%99%E6%89%A7%E8%A1%8C)
    - [defer / go 函数的参数在何时求值](#defer--go-%E5%87%BD%E6%95%B0%E7%9A%84%E5%8F%82%E6%95%B0%E5%9C%A8%E4%BD%95%E6%97%B6%E6%B1%82%E5%80%BC)
    - [defer 与函数相关，与块作用域无关](#defer-%E4%B8%8E%E5%87%BD%E6%95%B0%E7%9B%B8%E5%85%B3%E4%B8%8E%E5%9D%97%E4%BD%9C%E7%94%A8%E5%9F%9F%E6%97%A0%E5%85%B3)
    - [recover 需要在 defer 中直接调用](#recover-%E9%9C%80%E8%A6%81%E5%9C%A8-defer-%E4%B8%AD%E7%9B%B4%E6%8E%A5%E8%B0%83%E7%94%A8)

## 变量

### 变量定义

```go
func 变量定义() {
    var a int = 10        // 最完整的写法
    var b = 10            // 自动推断类型 or 使用默认类型
    c := 10               // 最简单的写法，省略 var关键字
    var d int             // 使用类型的零值
    g, h := 10, "hello"   // 可以一次性创建多个变量，但一般建议分成两行，可读性更好

    // 可以用 := 对旧变量赋值, 但有两个条件
    m := 10             // ①这个 m 和下一行的 m 位于同一层作用域
    m, n := 30, "hello" // ②操作符 := 的左边存在新创建的变量 ( 比如 n )

    // 所以呢 := 有三种行为,  ①创建新变量  ②对同层现有变量赋值
    x := 10
    if x > 5 {
        x := 2233      // ③在局部作用域中，创建同名的新变量，遮蔽外层变量
        fmt.Println(x) // 注意两个 x 变量只是名字相同，但各有各的内容
    }
}
// 包级别的变量只能用 var 创建, 因为 := 语法仅限于函数中使用
var packageLevelVariable int
```

### 变量遮蔽

- `x := 5` 会遮蔽外部的 x 变量,  这个「 外部 」可能是当前大括号外的 x 变量、或包级别的 x 变量
- 可以用 `math := "oops"` 遮蔽 math 包，甚至允许用 `true := 10` 遮蔽 universe block 中的 true 变量 (但不推荐)

```go
func TestShadowing(t *testing.T) {
    x := 10
    {
        // 这里的 := 会创建新的 x 变量, 而不是对外部的 x 赋值, 注意 := 的赋值行为仅发生在同层作用域
        x, y := 5, 20
        fmt.Println(x, y)
    }
    fmt.Println(x) // 值是 10 而不是 5
}
```

### 显式类型转换

Most languages that have multiple numeric types automatically convert from one to another when needed. This is called `automatic type promotion`, and while it seems very convenient, it turns out that the rules to properly convert one type to another can get complicated and produce unexpected results.

As a language that values clarity of intent and readability, Go doesn’t allow automatic type promotion between variables. You must use a type conversion when variable types do not match. Even different-sized integers and floats must be converted to the same type to interact.

```go
func TestExplicitTypeConversion(t *testing.T) {
    var a int32
    var b int64
    fmt.Println(a + b)        // 在其他语言会自动把 int32 提升为 int64
    fmt.Println(int64(a) + b) // 在 Golang 需要显式类型转换, 不允许 int32 和 int64 做计算
}
```

*Since all type conversions in Go are explicit, you cannot treat another Go type as a boolean*. In many languages, a nonzero number or a nonempty string can be interpreted as a boolean true. Just like automatic type promotion, the rules for “truthy” values vary from language to language and can be confusing. Unsurprisingly, Go doesn’t allow truthiness. In fact, no other type can be converted to a bool, implicitly or explicitly. If you want to convert from another data type to boolean, you must use one of the comparison operators (==, !=, >, <, <=, or >=)  

### Package 变量和 init()

#### ➤ `init()` 可用来初始化包级别的变量/状态

```go
// 先执行 package variable 的初始化, 比如 var age = GetAge()
var age = GetAge()
var name = "nia"

// 然后执行 init 函数, 每个 .go 文件可以有多个 init 函数, 它们按顺序执行
func init() {
    fmt.Println(1, name)    // 可以读取或修改 name 变量
    fmt.Println("2 homura")
}

func init() {
    fmt.Println("3 hikari")
}

func GetAge() int {
    fmt.Println("0 GetAge")
    return 18
}
```

#### ➤ 一般推荐显式注册/初始化，可读性比 `init()` 的隐式注册要好

Some packages, like database drivers, use `init` functions to register the database driver. However, you don’t use any of the identifiers in the package ( 比如用 `import _ "github.com/lib/pq"` 注册数据库驱动 ). This pattern is considered obsolete because it’s unclear that a registration operation is being performed. If you have a registry pattern in your own code, register your plug-ins explicitly. The primary use of `init` functions today is to initialize package-level variables that can’t be configured in a single assignment.

### Package 的初始化顺序

#### ➤ 大致顺序: `main -> import -> package-level variables -> init` ( [参考此处](https://yourbasic.org/golang/package-init-function-main-execution-order/) )

> - First the `main` package is initialized.
>   - Imported packages are initialized before the package itself.
>   - Packages are initialized one at a time:
>     - first package-level variables are initialized in declaration order,
>     - then the `init` functions are run.
> - Finally the `main` function is called.

#### ➤ 同一包中 `init()` 的执行顺序

> 1. You can have multiple `init()` functions per package; they will be executed in the order they show up in the file (after all variables are initialized of course).
>
> 2. If `init()` functions span multiple files, they will be called in the order in which they are presented to the compiler: 
>
>    - In most cases they called in order of filenames, A.go, a.go, d.go, ... 
>
>    - The [Go spec](https://golang.org/ref/spec#Package_initialization) says "build systems **are encouraged to** present multiple files belonging to the same package in lexical file name order to a compiler" ( 仅仅是建议、不强制，所以无法依赖 )
>
> 3. 如果需要以特定顺序执行 `init()` 函数，可以只写一个 `init()` 函数，在其中显式调用 `initA()`、`initB()`、`initC()`、...
>
> 4. Keep in mind that `init()` is always called, regardless if there's `main()` or not, so if you import a package that has an `init` function, it will be executed. 

#### ➤ 其他注意事项

- Initialization cycles are not allowed. There can be **no cyclic dependencies**.
- Each package is initialized **once**, regardless if it’s imported by multiple other packages.
- Package initialization happens in a single goroutine, sequentially, one package at a time.

### Package Initialization

#### ➤ [参考 Language Specification](https://go.dev/ref/spec#Package_initialization)

- A package with no imports is initialized
  - by assigning initial values to all its package-level variables,
  - followed by calling all `init` functions in the order they appear in the source.
- Package-level variables are initialized in **declaration order**, but after any of the variables they **depend** on.
- The **declaration order** of variables declared in multiple files is determined by the order in which the files are presented to the compiler.

> A package-level variable is considered *ready for initialization* if it has no [initialization expression](https://go.dev/ref/spec#Variable_declarations) or its initialization expression has no *dependencies* on uninitialized variables. Initialization proceeds by repeatedly initializing the next package-level variable that is earliest in declaration order and ready for initialization, until there are no variables ready for initialization. ( 简而言之按声明顺序初始化，如果依赖的变量尚未初始化，那么要等等 )
>
> The declaration order of variables declared in multiple files is determined by the order in which the files are presented to the compiler: Variables declared in the first file are declared before any of the variables declared in the second file, and so on. 
>
> To ensure reproducible initialization behavior, build systems are encouraged to present multiple files belonging to the same package in lexical file name order to a compiler.
>
> - 比如 `go run main.go a.go b.go` 会依次执行 main.go -> a.go -> b.go 中的 `init()` 函数
> - 比如 `go run a.go b.go main.go` 会依次执行 a.go -> b.go -> main.go 中的 `init()` 函数

#### ➤ 实际例子

```go
var (
    a = c + b // == 9
    b = f()   // == 4
    c = f()   // == 5
    x = y()   // == 0
    d = 3     // == 5 after initialization has finished
)

func f() int {
    d++
    return d
}

func y() int {
    return 0
}

func main() {
    // (1) 首先 a 依赖 c 和 b 所以要在它们初始化之后初始化

    // (2) 下一行是 b，所以尝试初始化 b，但是 f() 依赖 d，所以 b 要在 d 之后初始化
    //     下一行是 c，逻辑同上, 也要在 d 之后初始化

    // (3) 下一行是 x，表达式 y() 不依赖任何未初始化的变量，所以直接执行 y() 并初始化 x
    //     下一行是 d，表达式 3 不依赖任何未初始化的变量，所以初始化 d

    // (4) 初始化完 d，会让 b 和 c 都 ready for initialization，按照声明顺序，先处理 b 后处理 c
    // (5) 所以总顺序是: x d b c a
}
```

## 常量

### 字面量没有类型

You can’t even add two integer variables together if they are declared to be of different sizes. However, Go lets you use an integer literal in floating point expressions or even assign an integer literal to a floating point variable. *This is because literals in Go are untyped; they can interact with any variable that’s compatible with the literal.* We’ll see that we can even use literals with user-defined types based on primitive types.

```go
func TestLiteralHasNoType(t *testing.T) {
    // 字面量 1 既能赋值给 int 又能赋值给 float64, 它是什么类型?
    // 其实字面量没有类型, 只要字面量与变量的类型兼容, 就可以进行赋值
    var _ int = 1
    var _ float64 = 1
}
```

Being untyped only goes so far; you can’t assign a literal string to a variable with a numeric type or a literal number to a string variable, nor can you assign a float literal to an int.

### 通过 const 给字面量命名

Constants in Go are just that—constant. They are created at compile time. However, `const` in Go is very limited. Constants in Go are a way to give names to literals. They can only hold values that the compiler can figure out `at compile time`. This means that they can be assigned:

- Numeric literals、true and false、Strings、Runes
- The built-in functions complex, real, imag, len, and cap ( `const Length = len("abc")` )
- Expressions that consist of operators and the preceding values

For instance, `1<<3` is a constant expression, while `math.Sin(math.Pi/4)` is not because the function call to `math.Sin` needs to happen at run time.

Go doesn’t provide a way to specify that a value calculated at runtime is immutable. As we’ll see in the next chapter, there are no immutable arrays, slices , maps, or structs, and there’s no way to declare that a field in a struct is immutable. *Constants in Go are a way to give names to literals. There is no way in Go to declare that a variable is immutable.*

#### ➤ Typed and Untyped Constants

Constants can be typed or untyped. An untyped constant works exactly like a literal; it has no type of its own, but does have a default type that is used when no other type can be inferred. A typed constant can only be directly assigned to a variable of that type.

```go
const x = 10              // untyped
const typedX int = 10     // typed
var y byte = x            // 合法,  因为 x 常量没有类型
var z byte = typedX       // 非法,  因为不能把 int 类型赋值给 byte 类型
```

### 使用 iota 做枚举

Go doesn’t have an enumeration type. In Go, enumerated constants are created using the `iota` enumerator. Since `iota` can be part of an expression and expressions can be implicitly repeated, it is easy to build intricate sets of values.

```go
type ByteSize float64

const (
    _           = iota             // ignore first value by assigning to blank identifier
    KB ByteSize = 1 << (10 * iota) // iota 的值是它在 constant block 中的索引
    MB                             // 隐式重复上一行表达式
    GB
    TB
)
```

#### ➤ 使用 iota 的最佳实践

```go
// First, define a type based on int that will represent all the valid values.
type MailCategory int

// Next, use a const block to define a set of values for your type.
const (
    Uncategorized MailCategory = iota
    Personal
    Spam
    Social
    Advertisements
)
```

#### ➤ `iota` 的值是它在 constant block 中的索引

When the Go compiler sees this constant block, it repeats the type and the assignment to all of the subsequent constants in the block, and increments the value of `iota` on each line. This means that it assigns 0 to the first constant (Uncategorized), 1 to the second constant (Personal), and so on. When a new const block is created, `iota` is set back to 0.

#### ➤ 插入新枚举、交换两个枚举的位置都会导致枚举值发生变动

If you insert a new identifier in the middle of your list of literals, all of the subsequent ones will be renumbered. This will break your application in a subtle way if those constants represented values in another system or in a database. 

#### ➤ 如果关心枚举的具体值，应该显式指定常量值，别用 `iota`

Use iota for “internal” purposes only. iota-based enumerations only make sense when you care about being able to differentiate between a set of values, and don’t particularly care what the value is behind the scenes. If the actual value matters, specify it explicitly.

#### ➤ 把首个 `iota` 命名为 Invalid / Uncategorized / ... 用来表示默认状态

Be aware that iota starts numbering from 0. When mail first arrives, it is uncategorized, so the zero value makes sense. If there isn’t a sensical default value for your constants, a common pattern is to assign the first `iota` value in the constant block to a constant that indicates the value is invalid. 

## 理解 Comparable

### 可比较 vs 比大小

- The equality operators `==` and `!=` apply to operands of *comparable* types. 

- The ordering operators `<`, `<=`, `>`, and `>=` apply to operands of *ordered* types.

```go
func Test字典序(t *testing.T) {
    s1 := "21"
    s2 := "23"
    s3 := "3"
    fmt.Println(s1 < s2) // 两个字符串逐字节比较, 如果第一个字节判断不出来, 就看下一个字节
    fmt.Println(s2 < s3) // 即通过首个不相等的字节, 判断两个字符串的先后顺序
}
```

### 可比较类型

- `Channel` types are comparable. Two channel values are equal if they were created by the same call to [`make`](https://go.dev/ref/spec#Making_slices_maps_and_channels) or if both have value `nil`.
- `Array` types are comparable if their array element types are comparable. Two array values are equal if their corresponding element values are equal. 
- `String` types are comparable and ordered. Two string values are compared lexically byte-wise.
- `Struct` types are comparable if all their field types are comparable.
- `Interface` types that are not type parameters are comparable. 

#### ➤ `Pointer` types are comparable.

- Two pointer values are equal if they point to the same variable or if both have value `nil`. 
- Pointers to distinct [zero-size](https://go.dev/ref/spec#Size_and_alignment_guarantees) variables may or may not be equal. ( [这有例子](https://www.reddit.com/r/golang/comments/orfjdr/why_do_zerosize_variables_seem_to_have_the_same/) )
- A struct or array type has size zero if it contains no fields (or elements, respectively) that have a size greater than zero. 

### 不可比较类型

- slice、map、function 是不可比较的
- The only thing you can compare a slice/map/function with is `nil`. 
- It is a compile-time error to use `==` to see if two slices are identical or `!=` to see if they are different. 

#### ➤ 另外反射包有个 `reflect.DeepEqual` 方法几乎能比较任意类型，包括 slice、map

```go
func TestDeepEqual(t *testing.T) {
    m1 := map[string]int{"a": 1}
    m2 := map[string]int{"a": 1}
    fmt.Println(reflect.DeepEqual(m1, m2)) // true

    var s1 []byte
    var s2 = []byte{}
    fmt.Println(reflect.DeepEqual(s1, s2)) // false,  DeepEqual 认为 nil 切片和空切片不相等
    fmt.Println(bytes.Equal(s1, s2))       // true,   bytes.Equal 认为 nil 切片和空切片相等

    s3 := []string{"one", "two"}
    s4 := []interface{}{"one", "two"}
    fmt.Println(reflect.DeepEqual(s3, s4)) // false,  因为两个切片的类型不一样
}
```

### 比较 interface 时要注意

- Interface types that are not type parameters are comparable. 
- Two interface values are equal if they have [identical](https://go.dev/ref/spec#Type_identity) dynamic types and equal dynamic values 
- or if both have value `nil`.

#### ➤ 两个 interface 作比较可能会 panic

A comparison of two interface values with identical dynamic types causes a run-time `panic` if that type is not comparable. This behavior applies not only to direct interface value comparisons but also when comparing arrays of interface values or structs with interface-valued fields.

```go
func TestInterfaceComparePanic(t *testing.T) {
    var a any = []int{1, 2, 3}
    var b any = []float64{1, 2, 3}
    var c any = []float64{1, 2, 3}
    fmt.Println(a == b) // 两个 interface 类型不相等, 所以是 false
    fmt.Println(b == c) // 两个 interface 类型相等, 但却是不可比较类型, 所以 panic !!!
}
```

#### ➤ 有时候 interface 能和具体类型作比较

A value `x` of non-interface type `X` and a value `t` of interface type `T` can be compared if type `X` is comparable and `X` [implements](https://go.dev/ref/spec#Implementing_an_interface) `T`. They are equal if `t`'s dynamic type is identical to `X` and `t`'s dynamic value is equal to `x`.

```go
type X struct{ data int } // X 是可比较类型
type T interface{}        // T 不包含任何方法, 所有类型包括 X 都实现了 T 接口

func TestCompareTypeAndInterface(_ *testing.T) {
    var x X
    var t T
    // 如果具体类型 X 实现了 T 接口, 并且 X 是可比较类型, 那么 X 变量和 T 变量能进行比较
    // 因为 x 和 t 都是 T, 它们有相等的可能性, 所以应该允许 x 和 t 进行比较
    fmt.Println(t == x)

    // 总而言之, 如果接口变量 t 的动态类型和值都与 x 变量相等, 那么 t == x
    t = X{data: 0}
    fmt.Println(t == x)
}
```

### 比较 struct types

- Struct types are comparable if all their field types are comparable.
- Whether or not a struct is comparable depends on the struct’s fields. Structs that are entirely composed of comparable types are comparable; those with slice or map fields are not. 

Just like Go doesn’t allow comparisons between variables of different primitive types, Go doesn’t allow comparisons between variables that represent structs of different types. But Go does allow you to perform a type conversion from one struct type to another if the fields of both structs have the same names, order, and types. 

```go
type Hour struct {
    data int
}
type Minute struct {
    data int
}

func TestCompareStruct(t *testing.T) {
    var h Hour
    var m Minute
    fmt.Println(h == m)       // 两个结构体不能比较, 因为他们是不同的类型
    fmt.Println(h == Hour(m)) // 两个结构体的字段顺序、字段名、字段类型完全相等, 所以允许类型转换
}
```

Unlike in Python or Ruby, in Go there’s no magic method that can be overridden to redefine equality and make == and != work for incomparable structs. You can, of course, write your own function that you use to compare structs.

一个 anonymous struct 变量可以与另一个 struct 变量，直接进行比较操作、赋值操作，无需类型转换,  
但前提是 If the fields of both structs have the same names, order, and types. 

## 各种控制结构

### 对比其他语言

- There is no `do` or `while` loop, only a slightly generalized `for`; 
- `switch` is more flexible; 
- `if` and `switch` accept an optional initialization statement like that of `for`; 
- `break` and `continue` statements take an optional label to identify what to break or continue; 
- there are new control structures including a `type switch` and a multiway communications multiplexer, `select`. 

### 独特的 if 语句

The most visible difference between if statements in Go and other languages is that you don’t put parenthesis around the condition. What Go adds is the ability to declare variables that are scoped to the condition and to both the if and else blocks.

```go
func if_statement() {
    if n := rand.Intn(10); n == 0 {
        // 这里可以使用 n
    } else if n > 5 {
        // 这里可以使用 n
    } else {
        // 这里可以使用 n
    }
    // 这里不能使用 n
}
```

### For 循环的四种写法

#### ➤ 这不是茴香豆的四种写法，他们都有对应的作用

```go
// #### ➤ A complete, C-style for
for i := 0; i < 10; i++ {
    fmt.Println(i)
}

// #### ➤ A condition-only for, like the while statement found in C
for i < 100 {
    fmt.Println(i)
    i = i * 2
}

// #### ➤ An infinite for
for {
    fmt.Println("Hello")
}
```

#### ➤ for-range

If you're looping over an array, slice, string, or map, or reading from a channel, a `range` clause can manage the loop. You can only use a for-range loop to iterate over the built-in compound types and user-defined types that are based on them.

```go
// 第一个返回值是 key 或 index,  第二个返回值是 value,  可以忽略其中一个返回值
for i, v := range ARRAY {fmt.Println(i, v)}
for k, v := range MAP {fmt.Println(k, v)}
for k := range MAP {fmt.Println(k)}            // 忽略 value
for _, v := range MAP {fmt.Println(v)}         // 忽略 key

// 字符串的 for-range 循环比较新鲜,  可以观察到 byte_offset 不是每次都加一,  因为中文字符占 3 个字节
// 第一个返回值是 unicode 码点在 utf-8 字节数组中的偏移量,  第二个返回值是 rune 类型的 unicode 码点值
for byte_offset, rune_value := range "abc字符串🐶!" {
    fmt.Println(byte_offset, rune_value, string(rune_value))
}
```

For strings, the `range` does more work for you, breaking out individual Unicode code points by parsing the UTF-8. Erroneous encodings consume one byte and produce the replacement rune U+FFFD.

```go
func TestForRangeString(t *testing.T) {
    // "焰" 的 utf-8 编码为 [e7 84 b0], 把最后一个字节篡改成 0x21
    // 这三个字节无法构成一个汉字, for-range 把前两个字节解释成 U+FFFD 并把 0x21 解码为感叹号
    for pos, char := range "我爱光\xe7\x84\x21" {
        fmt.Printf("character %#U starts at byte position %d\n", char, pos)
    }
}
character U+6211 '我' starts at byte position 0
character U+7231 '爱' starts at byte position 3
character U+5149 '光' starts at byte position 6
character U+FFFD '�' starts at byte position 9
character U+FFFD '�' starts at byte position 10
character U+0021 '!' starts at byte position 11
```

#### ➤ 使用多个变量

Finally, Go has no comma operator and `++` and `--` are statements not expressions. Thus if you want to run multiple variables in a `for` you should use parallel assignment:

```go
// Reverse a
for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
    a[i], a[j] = a[j], a[i]
}
```

#### ➤ The for-range value is a copy

You should be aware that each time the for-range loop iterates over your compound type, it copies the value from the compound type to the value variable. Modifying the value variable will not modify the value in the compound type.  
不要用 `for _, v := range items { v.id = 123 }` 修改取出来的结构体，因为 `v` 是一个拷贝

#### ➤ Labeling Your for Statements

By default, the break and continue keywords apply to the for loop that directly contains them. What if you have nested for loops and you want to exit or skip over an iterator of an outer loop?  

```go
func break_continue_嵌套循环() {
    samples := []string{"hello", "apple_π!"}
outer:
    for _, sample := range samples {
        for i, r := range sample {
            fmt.Println(i, r, string(r))
            if r == 'l' {
                continue outer // 我们想 continue 外层的循环
            }
        }
        fmt.Println()
    }
}
```

### For 变量可能是同一个

#### ➤ for 循环的迭代变量是同一个 ( Go <= 1.21 )

```go
var userList = []User{
    {Name: "Cloud"},
    {Name: "Tifa"},
}

func TestLoopVariable(t *testing.T) {
    // 很常见的需求, 把一个 list 转成以 Name 字段为键的 map
    um := make(map[string]*User, len(userList))

    for i, u := range userList {
        um[u.Name] = &u           // 这么写在旧版本是 bug
        um[u.Name] = &userList[i] // 正确写法
    }

    // 因为 for 循环的迭代变量是同一个 ( Go <= 1.21 )
    // i 和 u 变量每次循环被赋予新值, 但依旧是同一个变量, 所以 &u 会得到完全相同的地址
    t.Log(um["Tifa"] == um["Cloud"], um)
}
```

#### ➤ 在值变量上调用指针方法，会隐式地取地址

如果 `Print` 是指针方法，下述调用会变成 `defer (&n).Print()`，隐式地用到了 `&n`，所以也会遇到前文的问题

```go
func TestPointerReceiver(t *testing.T) {
    names := []Name{"Cloud", "Alice", "Tifa"}
    for _, n := range names {
        defer n.Print() // Go <= 1.21 会全都打印 Tifa
    }
}

type Name string
func (n *Name) Print() { fmt.Println("I am", *n) }
```

#### ➤ for 循环的迭代变量如今不再是同一个 ( Go >= 1.22 )

```go
func TestGo122(t *testing.T) {
    // 对于 go.mod 版本 >= Go 1.22 的模块, 其中的包会使用新语义: 每次循环都分配新变量
    for i, u := range userList {
        // 相当于在每次循环的开头, 执行这两行代码
        i := i
        u := u
    }
}
```

### 独特的 switch 语句

Go's `switch` is more general than C's. The expressions need not be constants or even integers, the cases are evaluated top to bottom until a match is found, and if the `switch` has no expression it switches on `true`. It's therefore possible—and idiomatic—to write an `if`-`else`-`if`-`else` chain as a `switch`.

```go
func unhex(c byte) byte {
    switch {
    case '0' <= c && c <= '9':
        return c - '0'
    case 'a' <= c && c <= 'f':
        return c - 'a' + 10
    case 'A' <= c && c <= 'F':
        return c - 'A' + 10
    }
    return 0
}
```

There is no automatic fall through, but cases can be presented in comma-separated lists.

```go
func switch_statement() {
    words := []string{"a", "cow", "smile", "gopher", "octopus", "anthropologist"}
    for _, word := range words {
        // 每个分支都自动包裹在一个大括号里面,  默认不会 fallthrough
        // 所以不用像 C++ / Java 一样写 {} 把分支包起来,  不用像其他语言一样写 break 语句
        // Also like an if statement, you can declare a variable that’s scoped to all the branches.
        switch size := len(word); size {
        case 1, 2, 3, 4:
            fmt.Println(word, "is a short word!")
        case 5:
            wordLen := len(word)
            fmt.Println(word, "is exactly the right length:", wordLen)
        case 6, 7, 8, 9:
            // 这一行加上 fallthrough 语句才会往下掉
            // 注意 case 中可以用逗号分隔多种情况 (case 1,2,3,4:),  所以 fallthrough 一般用不到
        default:
            fmt.Println(word, "is a long word!")
        }
    }
}
```

In our sample program we are switching on the value of an integer, but that’s not all you can do. You can switch on any type that can be compared with `==`, which includes all of the built-in types except slices, maps, channels, functions, and structs that contain fields of these types.  

#### ➤ Blank Switch

You can write a switch statement that doesn’t specify the value that you’re comparing against. This is called a blank switch. A regular switch only allows you to check a value for equality. A blank switch allows you to use any boolean comparison for each case.

```go
func blank_switch_statement(err error) {
    switch {
    case errors.Is(err, io.ErrClosedPipe):
        fmt.Println("aaa")
    case errors.Is(err, io.ErrShortBuffer):
        fmt.Println("bbb")
    default:
        fmt.Println("ccc")
    }
}
```

#### ➤ switch/select 中的 break 只是跳出当前 switch/select

在 switch 语句中一般不会写 break，但如果写了 break 那么会提前退出当前 switch 语句  
如果 for 循环里面嵌套了一个 switch 语句，可以用上一节讲的 labeled for statement 退出外层循环

```go
func TestBreakOuterLoop(t *testing.T) {
    var x string
loop:
    for {
        switch x {
        case "A":
            break      // X 仅退出当前 switch, 会无限循环
        case "B":
            break loop // √ 用标签跳出外层循环
        }
    }
}
```

### Type Assertion

Go provides two ways to see if a variable of an interface type has a specific concrete type or if the concrete type implements another interface. If the type assertion ( `str, ok := value.(string)` ) fails, `str` will still exist and be of type string, but it will have the zero value, an empty string.

```go
func TestTypeAssertion(t *testing.T) {
    var num Number = 123
    var i interface{}
    i = num

    // 断言 i 的具体类型为 int,  如果 i 不是 int 会导致 panic
    // 注意 i 的具体类型是 Number 而不是 int,  所以下面这行代码会 panic!!
    // 虽然 Number 使用了 int 作为底层类型,  但 Number 和 int 是两个独立的类型
    _ = i.(int)

    // 可以用第二个返回值 ok 来避免 panic,  如果类型断言不 ok 那么 n 为 Number 的零值
    if n, ok := i.(Number); ok {
        n.PrintDouble()
    }
}
```

### Type Switch

A switch can also be used to discover the dynamic type of an interface variable. Such a *type switch* uses the syntax of a type assertion with the keyword `type` inside the parentheses. If the switch declares a variable in the expression, the variable will have the corresponding type in each clause. 

```go
func 当接口变量可能是多种类型时_使用_type_switch(i interface{}) {
    switch j := i.(type) {
    case nil:
        // i is nil, type of j is interface{} ( 这里接口 i 的类型和值都是 nil
    case Number:
        // j is of type Number
    case io.Reader:
        // j is of type io.Reader (这里判断 i 是否实现了 io.Reader 接口
    case bool, rune:
        // i is either a bool or rune, so j is of type interface{}
    default:
        // no idea what i is, so j is of type interface{}
        fmt.Println(j)
    }
    // 惯用的写法是把结果赋值给同名变量, 例如: switch i := i.(type) {...}
}
```



## 函数

### 参数按值传递

- It means that when you supply a variable for a parameter to a function,  
  Go always makes a copy of the value of the variable.  
  这种复制是浅复制，但对于不含指针的对象，浅复制等价于深复制
- 如果用 √ X 表示是否影响外部:
  - 在函数内修改传入的 `int`、`string`、`struct` (X)
  - 在函数内修改传入的 `map`,  增删 key (√)、赋予 key 新的 value (√)
  - 在函数内修改传入的 `slice`,  修改索引对应的值 (√),  用 append 修改切片长度 (X)
- 把 string/slice 传给函数的复制开销很小，只需要复制 SliceHeader 就好，其中包含 cap、len 字段和一个 data 指针

### 命名参数、可选参数

Go doesn’t have: named and optional input parameters.  You must supply all of the parameters for a function.  If you want to emulate named and optional parameters, define a struct that has fields that match the desired parameters, and pass the struct to your function.

### 变长参数

Like many languages, Go supports variadic functions. The variadic parameter must be the last (or only) parameter in the input parameter list. You indicate it with three dots `...` before the type. The variable that’s created within the function is a slice of the specified type. You use it just like any other slice.  

```go
func sum(nums ...int) int {
    total := 0
    for _, num := range nums {
        total += num
    }
    return total
}
```

### 命名返回值

#### ➤ 命名返回值就是用变量存储返回值

Go also allows you to specify names for your return values. When you supply names to your return values, what you are doing is pre-declaring variables that you use within the function to hold the return values. When named, they are initialized to the zero values for their types when the function begins;

```go
func div(a int, b int) (result int, err error) {
    // result 和 err 变量的值就是函数的返回值
    if b == 0 {
        err = errors.New("dividing by zero")
        return
    }
    result = a / b
    return
}
// return 后面啥也不跟就叫 blank return.
// However, most experienced Go developers consider blank returns 
// a bad idea because they make it harder to understand data flow.
```

#### ➤ 一些注意点:

1. 返回值变量其实就是普通变量啦，相当于在函数开头用 `var result int` 和 `var err error` 创建变量
2. `return 0, nil` 就是把 0、nil 分别赋值给 result、err 变量

#### [➤ 为什么 defer 配合 named return vlaue 能修改函数返回值?](https://stackoverflow.com/questions/37248898/how-does-defer-and-named-return-value-work)

1. defer 函数在 return 语句之后，在函数返回之前执行
2. 在函数返回之前，在 defer 函数中修改 `result` 变量，就能改掉函数的最终返回值

### 函数变量/类型

#### ➤ 声明函数变量: `var f func(a, b int) int`.

Just like in many other languages, functions in Go are values. The type of a function is built out of the keyword `func` and the types of the parameters and return values. This combination is called the signature of the function. Any function that has the exact same number and types of parameters and return values meets the type signature.

#### ➤ 创建函数类型: `type opFuncType func(int,int) int`

Just like you can use the type keyword to define a struct, you can use it to define a function type, too. Any function that has two input parameters of type int and a single return value of type int automatically meets this type.

### 匿名函数

However, there are two situations where declaring anonymous functions without assigning them to variables is useful: defer statements and launching goroutines. 

```go
// 用变量存储匿名函数
hello := func(name string) {
    fmt.Println("hello", name)
}
hello("ichigo")

// 常用场景: 用于启动 goroutine
go func() {
    fmt.Println("async task")
}()

// 常用场景: 用于 defer 清理资源
defer func() {
    fmt.Println("close resource")
}()
```

#### ➤ 使用匿名函数对切片进行排序

```go
func Passing_Functions_as_Parameters() {

    type Person struct {
        FirstName string
        LastName  string
        Age       int
    }

    people := []Person{
        {"Rukia", "Kuchiki", 150},
        {"Ichigo", "Kurosaki", 17},
        {"Tifa", "Lockhart", 20},
        {"Cloud", "Strife", 21},
    }

    // 让切片中的数据根据 Age 排序,  注意这里 i,j 是索引
    // 这里使用了闭包,  匿名函数中访问了外部的 people 变量
    sort.Slice(people, func(i int, j int) bool {
        return people[i].Age < people[j].Age
    })

    fmt.Println(people)
}
```

### 返回清理函数

A common pattern in Go is for a function that allocates a resource to also return a closure that cleans up the resource.
```go
func 打开文件并返回cleanup函数() {
    f, cleanup, err := getFile(`D:\temp\x.txt`)
    if err != nil {
        log.Fatalln(err)
    }
    defer cleanup()

    bytes, err := ioutil.ReadAll(f)
    if err != nil {
        log.Fatalln(err)
    }
    fmt.Println(string(bytes))
}
```

## 闭包

### 什么是 Closure

#### ➤ 闭包就是函数读写它 body 外的变量

Functions declared inside of functions are special; they are closures. This is a computer science word that means that functions declared inside of functions are able to access and modify variables declared in the outer function.  

Go functions may be closures. A closure is a function value that references variables from outside its body. The function may access and assign to the referenced variables; in this sense the function is "bound" to the variables. For example, the `Counter` function returns a closure. Each closure is bound to its own `value` variable.

```go
func TestClosure(t *testing.T) {
    count := Counter()
    t.Log(count())
    t.Log(count())
    t.Log(count())
}

func Counter() func() int {
    var value int
    return func() int {
        value++
        return value
    }
}
```

#### ➤ 如果闭包捕获了一个会变的变量，那么要小心注意

```go
func TestCaptureVariable(t *testing.T) {
    var id int
    for i := 0; i < 4; i++ {
        id = i
        // 最后全都打印 3, 因为等闭包执行时, 变量的值变成了 3
        defer func() {
            t.Log(id)
        }()
    }

    // 解决办法有两个:
    // (1) 循环中加一行 id := id 遮蔽外层作用域的同名变量
    // (2) 利用 defer 函数的求值机制: defer func(id int){...}(i)
    // 这两个方法的思路一致: 让闭包捕获一个不会发生改变的变量
}
```

### 什么是 Method Value

从实例身上提取方法，比如 `m := a.Method`，这叫 method value，调用 `m()` 时的 receiver 是 `a` 或 `&a` 的副本

```go
type A int
func (a A) ValueMethod()    { fmt.Println(a) }
func (a *A) PointerMethod() { fmt.Println(*a) }

func TestValueMethod(t *testing.T) {
    a := A(1)
    m := a.ValueMethod // m 是一个闭包, 它会记住 a=1 这个值
    a = A(2)           // 修改 a 变量
    m()                // 依旧打印 1, 因为 m 保存了 a=1 这个值, 不随 a 变量发生变化

}

func TestPointerMethod(t *testing.T) {
    a := A(1)
    m := a.PointerMethod // m 是一个闭包, 它会记住 &a 这个指针
    a = A(2)             // 修改 a 变量
    m()                  // 这次打印 2, 因为 m 保存了 &a 这个指针, 通过指针能获取 a 的当前值

}

func TestExtractMethodFromType(t *testing.T) {
    m := A.ValueMethod   // 从类型提取方法, 得到一个函数, 这种做法叫做 method expression
    m(A(1))
    m(A(2))
}
```

### 使用 Method Value 的例子

```go
func TestLazyInit(t *testing.T) {
    var d lazyData
    fmt.Println(d.Data()) // 首次使用时初始化 data
    fmt.Println(d.Data()) // 并发环境下也只初始化一次
}

type lazyData struct {
    once sync.Once // 确保 data 在并发环境下也只初始化一次
    data []int
}

func (l *lazyData) init() {
    fmt.Println("init data")
    l.data = []int{1, 2, 3}
}

func (l *lazyData) Data() []int {
    l.once.Do(l.init) // 这里使用了 method value
    return l.data
}
```



## defer

### defer 的好处

Programs often create temporary resources, like files or network connections, that need to be cleaned up. This cleanup has to happen, no matter how many exit points a function has, or whether a function completed successfully or not. In Go, the cleanup code is attached to the function with the `defer` keyword. Deferring a call to a function such as `Close` has two advantages. First, it guarantees that you will never forget to close the file, a mistake that's easy to make if you later edit the function to add a new return path. Second, it means that the close sits near the open, which is much clearer than placing it at the end of the function.

### defer 前要检查错误

#### ➤ 别忘了关 HTTP Response Body

When you make requests using the standard http library you get a http response variable. If you don't read the response body you still need to close it. Note that you must do it for empty responses too. It's very easy to forget especially for new Go developers. [go - What could happen if I don't close response.Body?](https://stackoverflow.com/questions/33238518/what-could-happen-if-i-dont-close-response-body)

- 在 http client 中 `http.Response` 的 `Body` 字段无论用没用到都别忘了关，否则会泄露资源
- `http.Request` 的 `Body` 字段不必自行关闭，因为 server/client 都会帮你关 ( 多次关闭 `io.Closer` 是未定义行为 )

#### ➤ 要先检查错误再设置 defer

```go
func main() {
    resp, err := http.Get("https://127.0.0.1")
    defer resp.Body.Close() // 不要这样写! 如果有错那么 resp 是 nil,  这一行会空指针异常
    if err != nil {
        fmt.Println(err)
        return
    }
    // 应该挪到这一行的位置 (即检查错误之后) 设置 defer

    // 打印 Body
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(string(body))
}
```

### defer 关闭文件

- 注意 `defer f.Close()` 并非最正确的写法，因为关闭文件可能出错，这么写忽略了 `f.Close()` 返回的错误
- 可以像下面的 `cat` 函数一样，如果 `f.Close()` 出错了且之前没有发生错误，则修改函数返回的错误值

```go
func main() {
    // 使用 golang 写一个 cat 工具
    if len(os.Args) < 2 {
        log.Fatal("no file specified")
    }
    err := cat(os.Args[1])
    if err != nil {
        log.Fatal(err)
    }
}

func cat(file string) (err error) {
    f, err := os.Open(file)
    if err != nil {
        return err
    }
    // 使用 defer 关闭文件,  并且检查关闭文件是否有错误,  这里要使用 named return value 特性
    defer func() {
        cErr := f.Close()
        if nil == err { // 不要让 Close() 错误覆盖先前存在的错误
            err = cErr
        }
    }()

    // 每次最多读 2048 个字节然后写到 stdout
    data := make([]byte, 2048)
    for {
        count, err := f.Read(data)
        if err != nil { // 如果不是 io.EOF 则返回错误
            if err != io.EOF {
                return err
            }
            break
        }
        _, err = os.Stdout.Write(data[:count])
        if err != nil {
            return err
        }
    }
    return nil
}
```

### defer 在什么时候执行

defer 的执行时机为: 在 retrun/panic 语句之后、在函数返回之前  
(1) [Is golang defer statement execute before or after return statement?](https://stackoverflow.com/questions/52718143/is-golang-defer-statement-execute-before-or-after-return-statement)  
(2) [go - How does defer and named return value work?](https://stackoverflow.com/questions/37248898/how-does-defer-and-named-return-value-work)

defer 的执行顺序为 LIFO，最后一个注册的 defer 最先执行，位于下面的 defer 语句先执行  
defer 有时候不执行, 比如 `log.Fatal`、`os.Exit`、终端中按下 `Ctrl+C`,  这些都会导致进程立即退出

#### ➤ 用 defer 修改函数返回值

```go
func 使用_defer_修改_named_返回值() (guy string) {
    // 在使用了 named return values 特性的函数中,  return "rukia" 相当于把 "rukia" 赋值给 guy 变量
    // defer 函数在 return 语句之后执行,  之所以能修改函数返回值,  因为我们拿到了真正存储返回值的 guy 变量
    defer func() {
        fmt.Println(guy) // "rukia"
        guy = "ichigo"
    }()
    return "rukia"
}
```

### defer / go 函数的参数在何时求值

The arguments to the deferred function (which include the receiver if the function is a method) are evaluated when the *defer* executes, not when the *call* executes. Besides avoiding worries about variables changing values as the function executes, this means that a single deferred call site can defer multiple function executions. Here's a silly example.

```go
for i := 0; i < 5; i++ {
    go fmt.Printf("go %d\n", i)       // 此时便会对函数参数 i 进行求值,  相当于存了 i 的快照
    defer fmt.Printf("defer %d\n", i) // 此时便会对函数参数 i 进行求值,  相当于存了 i 的快照
}
```

- 总之 `defer someFunc(theArg)` 中的 `theArg` 部分在执行 `defer` 这一行时求值
- 另外 `defer i.Love(homura)` 这样的方法等价于 `defer Love(i, homura)`，两个参数在执行 `defer` 这一行时求值

### defer 与函数相关，与块作用域无关

`defer` is not block-based but function-based

```go
func Test2(t *testing.T) {
    {
        // defer 和块作用域无关, 并非在退出块作用域时执行 defer 函数
        // 所以打印顺序是 2 1
        defer fmt.Println(1)
    }
    defer fmt.Println(2)
}
```

### recover 需要在 defer 中直接调用

```go
func TestRecover(t *testing.T) {
    defer func() {
        doRecover() // panic is not recovered
    }()

    defer doRecover() // 这么写就没问题

    panic("error")
}

func doRecover() {
    // 若 recover was not called directly by a deferred function.
    // 则 recover 不会起作用,  语言就是这么规定的
    fmt.Println("recovered =>", recover())
}
```

