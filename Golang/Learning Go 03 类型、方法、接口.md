## Table of Contents
  - [Methods](#Methods)
    - [接收器是方法的首个参数](#%E6%8E%A5%E6%94%B6%E5%99%A8%E6%98%AF%E6%96%B9%E6%B3%95%E7%9A%84%E9%A6%96%E4%B8%AA%E5%8F%82%E6%95%B0)
    - [指针接收器 vs 值接收器](#%E6%8C%87%E9%92%88%E6%8E%A5%E6%94%B6%E5%99%A8-vs-%E5%80%BC%E6%8E%A5%E6%94%B6%E5%99%A8)
    - [指针的 method set 更大](#%E6%8C%87%E9%92%88%E7%9A%84-method-set-%E6%9B%B4%E5%A4%A7)
    - [能在值变量上调用指针方法](#%E8%83%BD%E5%9C%A8%E5%80%BC%E5%8F%98%E9%87%8F%E4%B8%8A%E8%B0%83%E7%94%A8%E6%8C%87%E9%92%88%E6%96%B9%E6%B3%95)
    - [允许在 nil 指针上调用方法](#%E5%85%81%E8%AE%B8%E5%9C%A8-nil-%E6%8C%87%E9%92%88%E4%B8%8A%E8%B0%83%E7%94%A8%E6%96%B9%E6%B3%95)
    - [能传函数的地方也能传方法](#%E8%83%BD%E4%BC%A0%E5%87%BD%E6%95%B0%E7%9A%84%E5%9C%B0%E6%96%B9%E4%B9%9F%E8%83%BD%E4%BC%A0%E6%96%B9%E6%B3%95)
    - [为类型添加方法的两个限制](#%E4%B8%BA%E7%B1%BB%E5%9E%8B%E6%B7%BB%E5%8A%A0%E6%96%B9%E6%B3%95%E7%9A%84%E4%B8%A4%E4%B8%AA%E9%99%90%E5%88%B6)
  - [Types](#Types)
    - [类型定义不是继承](#%E7%B1%BB%E5%9E%8B%E5%AE%9A%E4%B9%89%E4%B8%8D%E6%98%AF%E7%BB%A7%E6%89%BF)
    - [底层类型相同时允许转换](#%E5%BA%95%E5%B1%82%E7%B1%BB%E5%9E%8B%E7%9B%B8%E5%90%8C%E6%97%B6%E5%85%81%E8%AE%B8%E8%BD%AC%E6%8D%A2)
    - [那么底层类型是什么意思](#%E9%82%A3%E4%B9%88%E5%BA%95%E5%B1%82%E7%B1%BB%E5%9E%8B%E6%98%AF%E4%BB%80%E4%B9%88%E6%84%8F%E6%80%9D)
    - [命名类型 vs 非命名类型](#%E5%91%BD%E5%90%8D%E7%B1%BB%E5%9E%8B-vs-%E9%9D%9E%E5%91%BD%E5%90%8D%E7%B1%BB%E5%9E%8B)
  - [Type Embedding](#Type-Embedding)
    - [组合 vs 继承](#%E7%BB%84%E5%90%88-vs-%E7%BB%A7%E6%89%BF)
    - [把方法和字段提升到外层](#%E6%8A%8A%E6%96%B9%E6%B3%95%E5%92%8C%E5%AD%97%E6%AE%B5%E6%8F%90%E5%8D%87%E5%88%B0%E5%A4%96%E5%B1%82)
    - [Embedding 与 Forwarding](#Embedding-%E4%B8%8E-Forwarding)
    - [type A B 包含哪些方法](#type-A-B-%E5%8C%85%E5%90%AB%E5%93%AA%E4%BA%9B%E6%96%B9%E6%B3%95)
    - [选择深度最浅的方法](#%E9%80%89%E6%8B%A9%E6%B7%B1%E5%BA%A6%E6%9C%80%E6%B5%85%E7%9A%84%E6%96%B9%E6%B3%95)
    - [嵌入值还是指针呢](#%E5%B5%8C%E5%85%A5%E5%80%BC%E8%BF%98%E6%98%AF%E6%8C%87%E9%92%88%E5%91%A2)
    - [嵌入接口有何好处](#%E5%B5%8C%E5%85%A5%E6%8E%A5%E5%8F%A3%E6%9C%89%E4%BD%95%E5%A5%BD%E5%A4%84)
    - [查看类型的 method set](#%E6%9F%A5%E7%9C%8B%E7%B1%BB%E5%9E%8B%E7%9A%84-method-set)
    - [判断类型的 method set](#%E5%88%A4%E6%96%AD%E7%B1%BB%E5%9E%8B%E7%9A%84-method-set)
  - [Interfaces](#Interfaces)
    - [定义接口](#%E5%AE%9A%E4%B9%89%E6%8E%A5%E5%8F%A3)
    - [隐式接口与鸭子类型](#%E9%9A%90%E5%BC%8F%E6%8E%A5%E5%8F%A3%E4%B8%8E%E9%B8%AD%E5%AD%90%E7%B1%BB%E5%9E%8B)
    - [显式 Interface Check](#%E6%98%BE%E5%BC%8F-Interface-Check)
    - [返回接口还是具体类型?](#%E8%BF%94%E5%9B%9E%E6%8E%A5%E5%8F%A3%E8%BF%98%E6%98%AF%E5%85%B7%E4%BD%93%E7%B1%BB%E5%9E%8B)
    - [接口包含 type 和 value 两个指针](#%E6%8E%A5%E5%8F%A3%E5%8C%85%E5%90%AB-type-%E5%92%8C-value-%E4%B8%A4%E4%B8%AA%E6%8C%87%E9%92%88)
    - [空接口一般和反射或序列化相关](#%E7%A9%BA%E6%8E%A5%E5%8F%A3%E4%B8%80%E8%88%AC%E5%92%8C%E5%8F%8D%E5%B0%84%E6%88%96%E5%BA%8F%E5%88%97%E5%8C%96%E7%9B%B8%E5%85%B3)
    - [经常能看到用函数实现接口的套路](#%E7%BB%8F%E5%B8%B8%E8%83%BD%E7%9C%8B%E5%88%B0%E7%94%A8%E5%87%BD%E6%95%B0%E5%AE%9E%E7%8E%B0%E6%8E%A5%E5%8F%A3%E7%9A%84%E5%A5%97%E8%B7%AF)
    - [接口被设置了类型信息，就不等于 nil](#%E6%8E%A5%E5%8F%A3%E8%A2%AB%E8%AE%BE%E7%BD%AE%E4%BA%86%E7%B1%BB%E5%9E%8B%E4%BF%A1%E6%81%AF%E5%B0%B1%E4%B8%8D%E7%AD%89%E4%BA%8E-nil)

## Methods

### 接收器是方法的首个参数

比如 `p.ChangeName("nia")` 等价于 `ChangeName(p, "nia")`，这样想方便理解值和指针接收器的区别

```go
type Person struct{ name string }
func (p Person)  ChangeName1(name string) { p.name = name }
func (p *Person) ChangeName2(name string) { p.name = name }

func TestName(t *testing.T) {
    p := Person{name: "nia"}

    p.ChangeName1("homura   ") // 无法修改, 因为 ChangeName1(p, "homura") 会把 p 结构体复制
    _ = ""                     // 所以改掉了副本, 但原来的 p 不受影响

    (&p).ChangeName2("hikari") // 能够修改, 因为 ChangeName2(&p, "hikari") 会把 &p 指针复制
    _ = ""                     // 所以通过指针修改了原来的 p 结构体
    p.ChangeName2("hikari")    // 调用指针方法时, 不必像上面一样用 &p, 编译器会自动把 p 改写成 &p
}
```

### 指针接收器 vs 值接收器

```go
type Counter struct {
    count int
}
// (c *Counter) 是指针接收器,  通过指针能修改 count 字段
func (c *Counter) Add() {
    c.count++
}
// (c Counter)  是值接收器,  无法修改 count 字段,  因为 c 是结构体的拷贝
func (c Counter) Add2() {
    c.count++
}
```

Go uses parameters of pointer type to indicate that a parameter might be modified by the function. The same rules apply for method receivers, too. They can be pointer receivers (the type is a pointer) or value receivers (the type is a value type). The following rules help you determine when to use each kind of receiver:

- If your method modifies the receiver, you `must` use a pointer receiver.
- If your method needs to handle `nil` instances, then it `must` use a pointer receiver.
- If your method doesn’t modify the receiver, you `can` use a value receiver.
- When a type has any pointer receiver methods, a common practice is to `be consistent` and use pointer receivers for all methods, even the ones that don’t modify the receiver.

**Correctness wins over speed or simplicity.** There are cases where you must use a pointer value. In other cases, pick pointers for large types or as future-proofing if you don’t have a good sense of how the code will grow, and use values for simple [plain old data](https://en.wikipedia.org/wiki/Passive_data_structure). The list below spells out each case in further detail:

```go
// (1) 不希望发生复制，或不能被安全复制，此时用 pointer receiver，典型例子为包含 sync.Mutex 字段
type Counter struct {
    mu    sync.Mutex
    total int
}
func (c *Counter) Inc() { // 必须用 pointer receiver
    c.mu.Lock()           // 因为如果复制了结构体中的锁, 会让 Inc 的并发同步iu
    defer c.mu.Unlock()
    c.total++
}

// (2) 如果调用方法会产生任何修改行为，比如会修改深层的字段，那么用 pointer receiver 暗示会发生修改
type Counter struct {
    m *Metric
}
func (c *Counter) Inc() {
    c.m.Add(1)
}

// (3) 如果 receiver 是 map/function/channel 类型或不涉及 reallocate 操作的切片，那么用 value receiver
type Header map[string][]string
func (h Header) Add(key, value string) { }

// (4) 不可变对象应该用 value receiver，典型例子比如 time.Time
type Time struct { /* omitted */ }
func (t Time) Add(d Duration) Time { }
```

### 指针的 method set 更大

```go
func main() {
    type HasTwoAdd interface {
        Add()
        Add2()
    }
    var c Counter
    var i HasTwoAdd
    i = c  // 错误,  因为 c  是 value instance, 方法集只包含 Add2
    i = &c // 正确,  因为 &c 是 pointer instance,  方法集包含两个 Add 方法
}
```

Go considers both pointer and value receiver methods to be in the method set for a pointer instance. For a value instance, only the value receiver methods are in the method set.

### 能在值变量上调用指针方法

The rule about pointers vs. values for receivers is that value methods can be invoked on pointers and values, but pointer methods can only be invoked on pointers. This rule arises because pointer methods can modify the receiver; invoking them on a value would cause the method to receive a copy of the value, so any modifications would be discarded. The language therefore disallows this mistake. There is a handy exception, though. When the value is `addressable`, the language takes care of the common case of invoking a pointer method on a value *by inserting the address operator automatically*.

One thing you might notice is that we were able to call the pointer receiver method even though `c` is a value type. When you use a pointer method with a local variable that’s a value type, Go automatically converts it to a pointer type. In this case, `c.Add()` is converted to `(&c).Add()`. But `c().Add()` can't be converted to `(&c()).Add()`.

It's OK to call a pointer receiver method on a value as long as the value is addressable. Not every variable is addressable though. Map elements are not addressable (you can't `&someMap["key"]`).

```go
func (c *Character) SayHi() {
    fmt.Println("Hi, I am", c.Name)
}

func main() {
    var c = Character{Name: "Cloud", From: "FF7", Age: 21}
    c.SayHi()            // c 是值类型、并且 addressable、且 SayHi 是指针方法
    (&c).SayHi()         // 所以编译器会自动把上一行重写成这一行

    var foo SayHier = &c // 这里必须用 &c 取指针, 如果用 c 会编译错误
    foo.SayHi()          // 因为 c 是一个值而不是指针,  c 的方法集不包括指针方法
}

type SayHier interface {
    SayHi()
}
```

### 允许在 nil 指针上调用方法

What happens when you call a method on a `nil` instance? In most languages, this produces some sort of error.  Go does something a little different. It actually tries to invoke the method. If it’s a method with a value receiver, you’ll get a panic, as there is no value being pointed to by the pointer. If it’s a method with a pointer receiver, it can work if the method is written to handle the possibility of a nil instance. 

- 通常而言，在 `nil` 上调用方法会空指针异常

- 但如果方法使用 pointer receiver，且在方法中处理了指针为 `nil` 的情况，那么在 `nil` 上调用方法也不会出错

However, most of the time `nil` receiver is not very useful. If your method has a pointer receiver and won’t work for a `nil` receiver, check for `nil` and return an error.

### 能传函数的地方也能传方法

#### ➤ 传方法时就像传了一个闭包

Methods in Go are so much like functions that you can use a method *as a replacement for a function any time there’s a variable or parameter of a function type*. We can assign the method to a variable or pass it to a parameter. This is called a *method value*. A method value is a bit like a closure, since it can access the values in the fields of the instance from which it was created. 

```go
type Boy struct{ name string }

func (b Boy) greet() { fmt.Printf("Hi, my name is %s.", b.name) }

func 方法是一个闭包() {
    ichigo := Boy{name: "ichigo"}
    sayHi := ichigo.greet // 即使把方法提取出来保存到变量,  也能继续访问 receiver
    sayHi()               // Hi, my name is ichigo.
}
```

#### ➤ 甚至能用 `Boy.greet` 把方法提取成函数，函数的第一个参数是 receiver

You can also create a function from the type itself. This is called a *method expression*. In the case of a method expression, the first parameter is the receiver for the method.

```go
func method_expression() {
    cloud := Boy{name: "cloud"}
    dante := Boy{name: "dante"}
    sayHi := Boy.greet
    sayHi(cloud)
    sayHi(dante)
}
```

### 为类型添加方法的两个限制

#### ➤ [参考回答](https://stackoverflow.com/a/65441367)

1. 无法为「 接口类型 」和「 指针类型 」添加任何方法，例如 `type Counter *int` 就不能为 Counter 添加方法
2. 无法为另一个包中的类型添加方法,  例如 `func (t *time.Time) Hello() {}`

## Types

### 类型定义不是继承

You can declare a user-defined type based on another user-defined type:  `type HighScore Score`. Declaring a type based on another type looks a bit like inheritance, but it isn’t. The two types have the same underlying type `int`, but that’s all. There is no hierarchy between these types. You can’t assign an instance of type `HighScore` to a variable of type `Score` or vice versa without a type conversion. Furthermore, any methods defined on `Score` aren’t defined on `HighScore`. A type conversion between types that share an underlying type keeps the same underlying data but associates different methods.  

#### ➤ type A B 中的 A B 有什么关系?

```markdown
- A 和 B 是两个完全不同的类型，各有各的方法
- A 和 B 的字段相同，A 和 B 的类型允许互相转换
```

### 底层类型相同时允许转换

```go
h := Hour(1)   // 如果 Hour 和 Second 的底层类型都是 int, 那么 Hour、Second、int 之间能相互转换  
s := Second(h) // 虽然允许这么转,  但有时候逻辑上是错的,  比如 1 小时应该等于 3600 秒
```

#### ➤ 匿名结构体相等性

Two struct types are identical if they have the same sequence of fields, and if corresponding fields have the same names, and identical types, and identical tags. [Non-exported](https://go.dev/ref/spec#Exported_identifiers) field names from different packages are always different.

```go
func TestAnonymous(t *testing.T) {
    var homura any = struct{ Age int }{}
    var nia any = struct{ Age int }{}
    var mio any = struct {
        Age int `json:"age"`
    }{}
    t.Log(homura == nia) // 类型都是 struct{Age int}, 所以类型相等, 然后值也相等
    t.Log(homura == mio) // 不相等, 因为 struct tag 不一样, 所以类型不一样
}
```

#### ➤ 两个结构体什么时候允许类型转换?

When explicitly converting a value from one struct type to another, as of Go 1.8 the tags are ignored. Thus two structs that differ only in their tags may be converted from one to the other:

```go
func example() {
    type T1 struct {
        X int `json:"foo"`
    }
    type T2 struct {
        X int `json:"bar"`
    }
    var v1 T1
    var v2 T2
    v1 = T1(v2) // now legal
}
```

### 那么底层类型是什么意思

#### ➤ 参考 [Golang underlying types](https://stackoverflow.com/questions/29332879/golang-underlying-types)

```go
type T1 string // T1 的底层类型为 string
type T2 T1     // T2 的底层类型为 T1 的底层类型, 即 string
type T3 []T1   // T3 的底层类型为 []T1, 因为 []T1 是 type literal
type T4 T3     // T4 的底层类型为 T3 的底层类型即 []T1
```

The [spec mentions](https://golang.org/ref/spec#Types):

> Each type `T` has an underlying type: If `T` is one of the predeclared boolean, numeric, or string types, **or a type literal**, the corresponding underlying type is `T` itself. **Otherwise, `T`'s underlying type is the underlying type of the type to which `T` refers in its type declaration**.

#### ➤ 什么是 type literal

```go
// A type may also be specified using a type literal, which composes a type from existing types.
var a []T1
var b map[T1]T2
var c func(T1) T2
```

#### ➤ 题外话

```markdown
- T2 的底层类型不是 T1，那么 T1 叫啥好呢?  
- 首先不能叫 base type，因为指针 *T 的 base type 是 T，这个概念已被占用  
- 然后网上有人把 T1 叫做 T2 的 source type，但这不是官方命名，官方对 T1 的形容就一句 given type:  
- A type definition creates a new, distinct type with the same underlying type and operations as the given type.
```

### 命名类型 vs 非命名类型

- A [named type](https://go.dev/ref/spec#Types) is always different from any other type.
- A value `x` of type `V` is *assignable* to a variable of type `T`
  - `V` and `T` have identical [underlying types](https://go.dev/ref/spec#Underlying_types) and 
  - at least one of `V` or `T` is not a [named type](https://go.dev/ref/spec#Types).

```go
// Predeclared types, defined types, and type parameters are called named types.
string, int, bool
type MyString string

// 不属于上述情况的类型, 那就不是 named type, 例如各种未命名的 type literal
var _ []string
var _ map[int]int

// 变量 f1 和 f2 的底层类型相等都是 func()int 并且变量 f2 的类型没有名字, 所以可以相互赋值
func TestAssignability(t *testing.T) {
    type MyFunc func() int
    var f1 MyFunc
    var f2 = func() int { return 0 }
    f1 = f2
    f2 = f1
}
```



## Type Embedding

### 组合 vs 继承

The software engineering advice “Favor object composition over class inheritance”. While Go doesn’t have inheritance, it encourages code reuse via built-in support for composition and promotion. There is a big difference between the two:

- If a type `T` inherits another type, then type `T` obtains the abilities of the other type. At the same time, each value of type `T` can also be viewed as a value of the other type.
- If a type `T` embeds another type, then type other type becomes a part of type `T`, and type `T` obtains the abilities of the other type, but none values of type `T` can be viewed as values of the other type.
- 两种方式都能调用 A 的能力，如果 T 继承了 A 那么 T 可以看作一种 A，但在组合中这种关系不成立

### 把方法和字段提升到外层

- type A { B } 相当于在 `A` 类型中加一个名为 `B` 的字段
- 但可以像 `a.MethodFromB()` 这样快捷地访问 B 的方法
- `a.MethodFromB()` 的完整形式为 `a.B.MethodFromB()`

Note that `Manager` contains a field of type `Employee`, but no name is assigned to that field. This makes `Employee` an `embedded field`. Any fields or methods declared on an embedded field are promoted to the containing struct and can be invoked directly on it.

```go
type Employee struct {
    ID     string
    Name   string
    Salary int
}

func (e Employee) Description() {
    fmt.Printf("%s %s \n", e.Name, e.ID)
}

type Manager struct {
    Employee     // 嵌入 Employee 类型
    Salary   int // 经理也有一个 Salary 字段
}

func TestEmbedding(t *testing.T) {
    m := Manager{
        Employee: Employee{
            ID:     "111",
            Name:   "ichigo",
            Salary: 233,
        },
        Salary: 8888888,
    }

    fmt.Println(m.ID)              // Manager 嵌入了 Employee,  Employee 的字段和方法会被提升到 Manager
    m.Description()                // 所以能在 Manager 上直接使用 Employee 的 Description 方法
    fmt.Println(m.Salary)          // m.Salary 是经理的薪水, 因为最近的 Salary 字段在 Manager 类型上
    fmt.Println(m.Employee.Salary) // 可用完整形式访问深层的同名字段
}
```

#### ➤ 嵌入不是继承

- You cannot assign a variable of type `Manager` to a variable of type `Employee`.  
- If you want to access the Employee field in Manager, you must do so explicitly: `var e = m.Employee`.

### Embedding 与 Forwarding

- Embedding 实际上是 Forwarding
- 注意 `a.MethodFromB()` 的 receiver 不是 `a` 而是 `a.B`

#### ➤ 嵌入类型就是加了一个字段，并把方法和字段提升到外层，方便调用

```go
type Manager struct {
    Employee     // 虽然没写字段名, 其实字段名是 Employee
    Salary   int // 经理的 Salary 字段会遮盖员工的, 可用完整形式 m.Employee.Salary 访问被遮盖的字段
}

m := Manager{
    Employee: Employee{Name:"ichigo"},  // 可证明 Manager 加了个 Employee 字段
    Salary: 8888888,                    // 用法和普通字段一模一样: Manager{ Employee{Name:"ichigo"}, 8888888 }
}
```

#### ➤ 调用被提升的方法时，receiver 不是外层类型

When we embed a type, the methods of that type become methods of the outer type, but when they are invoked the receiver of the method is the inner type, not the outer one. 

```go
type Manager struct { company.Employee }          // 嵌入 company.Employee 类型
type Manager struct { Employee company.Employee } // 就像这样添加一个 Employee 字段，但是方法和字段可直接访问
m.Name   <=> m.Employee.Name                      // 等价关系
m.Work() <=> m.Employee.Work()                    // m.Work() 的 receiver 是 m.Employee 而不是 m
```

### type A B 包含哪些方法

The method set of a type is composed of the methods declared directly for the type and the method set of the type's underlying type. 

```go
// A 和 B 是同一个底层类型, 即 struct{ C }
// 所以 A 隐式地嵌入了 C, 所以 A 包含 C 的方法
type A B
type B struct{ C }
type C struct{}

func (c B) MethodFromB() {}
func (c C) MethodFromC() {}

func TestMethod(t *testing.T) {
    var a A
    a.MethodFromC()
    a.MethodFromB() // 编译错误, type A B 不会获得 B 的方法
}
```

### 选择深度最浅的方法

#### ➤ 规则

```go
type Job struct {
    Command string
    *log.Logger
}
```

Embedding types introduces the problem of name conflicts but the rules to resolve them are simple. First, a field or method `X` hides any other item `X` in a more deeply nested part of the type. If `log.Logger` contained a field or method called `Command`, the `Command` field of `Job` would dominate it.

Second, if the same name appears at the same nesting level, it is usually an error; it would be erroneous to embed `log.Logger` if the `Job` struct contained another field or method called `Logger`. However, if the duplicate name is never mentioned in the program outside the type definition, it is OK. 

#### ➤ 例子

- `a.b` 会寻找名为 `b` 的字段或方法，并选择深度最浅的那一个
- 如果有两个最浅的 `b` 字段，则编译报错 `ambiguous selector` ( 仅当代码中用到了 `a.b` 才会报错 )

```go
type A struct{ *B } // A 嵌入 B
type B struct{ *C } // B 嵌入 C
type C struct{}

// B 和 C 都有 Hello 方法, 那么 a.Hello 选择哪一个呢?
func (c *C) Hello() { fmt.Println("i am ccc") }
func (b *B) Hello() { fmt.Println("i am bbb") }

func TestShortestPath(t *testing.T) {
    c := &C{}
    b := &B{c}
    a := &A{b}
    a.Hello()   // 深度最浅的 Hello 方法来自 B, 所以执行 B.Hello
    a.B.Hello() // 等价于这一行

    a.C.Hello()   // 先找到 C ，然后再从 C 找深度最浅的 Hello 方法
    a.B.C.Hello() // 等价于这一行
}
```

有点像 OOP 中的 Mehotd Overriding，继承全部方法并重写部分方法:

```go
type Employee struct{}
type Manager struct{ *Employee }

func (e *Employee) Salary() int { return 233 }
func (m *Manager) Salary() int  { return 888 }

func TestOverride(t *testing.T) {
    e := Employee{}
    m := Manager{&e}
    fmt.Println(e.Salary()) // 233
    fmt.Println(m.Salary()) // 重写了 Salary 方法所以是 888
}

// 只要方法名相同就会产生遮盖, 不管入参和返回值类型 ( Golang 没有函数重载 )
func (m *Manager) Salary() string { return "八八八" }
```

### 嵌入值还是指针呢

#### ➤ 参考: [It depends. There are several possibilities here.](https://stackoverflow.com/a/28505394)

```markdown
嵌入值 `T` 能获得更好的内存性能 ( 指内存局部性更好，内存分配次数更少，垃圾回收负担更小 )
如果外层对象按值传递，那么复制外层对象时，内层的值 `T` 也会被复制，可能会产生逻辑问题，比如内层的锁被复制
如果外层对象按指针传递，那么内层对象使用 `T` 或 `*T` 都可以，代码逻辑上没有任何区别，仅仅是内存结构的区别

然后是 method set 有区别，例如 `type Outer struct { T }` 嵌入了值 `T` 
那么 `Outer{}` 的方法集只包含 `T` 的值方法，另外 `&Outer{}` 包含 `T` 的全部方法

最后如果 T 的零值不可用或无意义，需要通过构造函数创建可用实例，那么应该嵌入 `*T`，暗示这个字段需要初始化
```

#### ➤ 外层对象被复制时，内层对象想共享同一个，那就应该嵌入 `*T`，例如:

```go
type File struct{ data int }
func (c *File) Increase() { c.data++ }

type App struct {
    // 外层的 App 可能会被复制, 例如 app2 := app
    // 我们希望 app2.Increase() 能修改同一个文件, 所以嵌入 *File 指针而不是 File
    *File
}
```

### 嵌入接口有何好处

Just like you can embed a type in a struct, you can also embed an interface in an interface. For example, the `io.ReadCloser` interface is built out of an `io.Reader` and an `io.Closer`. Just like you can embed a concrete type in a struct, you can also embed an interface in a struct. 

#### ➤ [在结构体中嵌入一个接口有何好处?](https://stackoverflow.com/a/24537547) 其实就是添加一个接口字段，并且接口方法提升到结构体

```go
type Weapon interface{ WeaponAbility() string }
                                                          // 函数类型 WeaponFunc 实现了 Weapon 接口
type WeaponFunc func() string                             // 方便用函数实现 Weapon 接口, 不必新建类型
func (f WeaponFunc) WeaponAbility() string { return f() } // 这种做法类似于标准库中的 http.HandlerFunc

type Character struct {
    Name   string
    From   string
    Age    int
    Weapon // 嵌入一个接口
}

func TestEmbedInterface(t *testing.T) {
    var dante = Character{Name: "但丁", From: "DMC", Age: 40}
    var 魔剑但丁 = WeaponFunc(func() string { return "咿呀剑法" })             // 直接用函数实现 Weapon 接口
    var 黑檀木与白象牙 = WeaponFunc(func() string { return "旋转跳跃突突突" }) //

    dante.Weapon = 魔剑但丁
    fmt.Println(dante.WeaponAbility()) // 接口的方法被提升到外层, 能方便调用
    dante.Weapon = 黑檀木与白象牙      // 让接口字段引用另一个对象, 能切换实现
    fmt.Println(dante.WeaponAbility())
}
```

### 查看类型的 method set

```go
type A struct{}
type B struct{ A }

func (a A) ValueMethod()    {}
func (a *A) PointerMethod() {}

func TestMethodSet(t *testing.T) {
    var a A
    PrintMethods(B{a})  // B{A} 只包含 A 的 value method
    PrintMethods(&B{a}) // *B{A} 包含 A 的 value method 和 pointer method
}

func PrintMethods(i any) {
    t := reflect.TypeOf(i)
    fmt.Println(t, "has", t.NumMethod(), "methods:")
    for i := 0; i < t.NumMethod(); i++ {
        fmt.Print(" method#", i, ": ", t.Method(i).Name, "\n")
    }
}
```

### 判断类型的 method set

[参考此处:](https://go101.org/article/type-embedding.html#:~:text=Implicit%20Methods%20for%20Embedding%20Types)

- type `struct{T}` and type `*struct{T}` both obtain all the methods of the type denoted by `T`.
- type `*struct{T}`, type `struct{*T}`, and type `*struct{*T}` all obtain all the methods of type `*T`.
- `*T`'s method set is a super set of `T`'s method set. ( 两者都包含 value method, `*T` 多了 pointer method )

示例:

```go
type A struct{}
type B struct{ A }

func (a A) ValueMethod()    {}
func (a *A) PointerMethod() {}

type HasValueMethod interface{ ValueMethod() }
type HasPointerMethod interface{ PointerMethod() }

func TestEmbedPointerOrValue(t *testing.T) {
    var a A
    var _ HasValueMethod = B{a}   // 类型 B{A}  可以访问 A 的 value method
    var _ HasPointerMethod = B{a} // 类型 B{A}  无法访问 A 的 pointer method ( 编译错误 )

    var _ HasPointerMethod = &B{a} // 类型 *B{A} 可以访问 A 的 pointer method
    var _ HasValueMethod = &B{a}   // 总之用指针就能访问 A 的全部方法, 不管是值方法还是指针方法
}
```





## Interfaces

### 定义接口

```go
type Stringer interface {
    String() string
}
```

Interfaces are usually named with “er” endings. We’ve already seen `fmt.Stringer`, but there are many more, including io.Reader, io.Closer, io.ReadCloser, json.Marshaler, and http.Handler.

### 隐式接口与鸭子类型

What makes Go’s interfaces special is that they are implemented implicitly. A concrete type does not declare that it implements an interface. If a concrete type contains all the methods of an interface, the concrete type implements the interface.   

Dynamically typed languages like Python, Ruby, and JavaScript don’t have interfaces. Instead, those developers use “duck typing,” which is based on the expression “If it walks like a duck and quacks like a duck, it’s a duck.” The concept is that you can pass an instance of a type as a parameter to a function as long as the function can find a method to invoke that it expects.

```js
function Vergil(thatGuy) {
    if (thatGuy.hasPower()) {         // 在这段 javascript 中,  不管 thatGuy 是什么类型
        console.log("这是什么抛瓦?")  // 只要它有 hasPower() 方法,  就能传入 Vergil 方法
    }
}
```

### 显式 Interface Check

As we saw in the discussion of [interfaces](https://go.dev/doc/effective_go#interfaces_and_types) above, a type need not declare explicitly that it implements an interface. Instead, a type implements the interface just by implementing the interface's methods. In practice, most interface conversions are static and therefore checked at compile time. For example, passing an `*os.File` to a function expecting an `io.Reader` will not compile unless `*os.File` implements the `io.Reader` interface.

Some interface checks do happen at run-time, though. One instance is in the `encoding/json` package, which defines a `Marshaler` interface. When the JSON encoder receives a value that implements that interface, the encoder invokes the value's marshaling method to convert it to JSON instead of doing the standard conversion. The encoder checks this property at run time with a [type assertion](https://go.dev/doc/effective_go#interface_conversions) like:

```go
m, ok := val.(json.Marshaler)
```

If `json.RawMessage` needs a custom JSON representation, it should implement `json.Marshaler`, but there are no static conversions that would cause the compiler to verify this automatically. To guarantee that the implementation is correct, a global declaration using the blank identifier can be used in the package:

```go
var _ json.Marshaler = (*RawMessage)(nil) // 让编译器检查 RawMessage 是否实现了 json.Marshaler 接口
```

Don't do this for every type that satisfies an interface, though. By convention, such declarations are only used when there are no static conversions already present in the code, which is a rare event.

### 返回接口还是具体类型?

#### ➤ 入参应该用「 接口 」

如果 `CreateUserService` 用 `DB` 接口作为入参,  而不是依赖某个具体类型,  
那么能方便地切换实现,  比如从 mysql 切换到 postgres 数据库.

#### ➤ 返回值一般用「 具体类型 」，推荐让调用方定义接口

Go interfaces generally belong in the package that uses values of the interface type, not the package that implements those values. The implementing package should return concrete (usually pointer or struct) types: that way, new methods can be added to implementations without requiring extensive refactoring. Because adding a new method to an interface means that you need to update all existing implementations of the interface, or your code breaks.

#### ➤ 如果类型仅用来实现特定接口，没有任何其它用途，可以返回接口

If a type exists only to implement an interface and will never have exported methods beyond that interface, there is no need to export the type itself. Exporting just the interface makes it clear the value has no interesting behavior beyond what is described in the interface. It also avoids the need to repeat the documentation on every instance of a common method.

In such cases, the constructor should return an interface value rather than the implementing type. As an example, in the hash libraries both `crc32.NewIEEE` and `adler32.New` return the interface type `hash.Hash32`. Substituting the CRC-32 algorithm for Adler-32 in a Go program requires only changing the constructor call; the rest of the code is unaffected by the change of algorithm.

### 接口包含 type 和 value 两个指针

#### ➤ 当两个指针都为空时接口才等于 nil

We also use `nil` to represent the zero value for an interface instance, but it’s not as simple as it is for concrete types. In order for an interface to be considered `nil` **both the type and the value must be nil**. In the Go runtime, interfaces are implemented as a pair of pointers, one to the underlying type and one to the underlying value. As long as the type is non-nil, the interface is non-nil.  

```go
func TestInterface(t *testing.T) {
    var s *string
    var i any
    println(s == nil) // true
    println(i == nil) // true
    i = s
    println(i == nil) // false, 因为赋值后接口 i 中的类型字段变成了 *string
}
```

What `nil` indicates for an interface is whether or not you can invoke methods on it. As we covered earlier, you can invoke methods on `nil` concrete instances, so it makes sense that you can invoke methods on an interface variable that was assigned a `nil` concrete instance. 

#### ➤ 用 nil 接口调用方法一定会 panic，但就算用 non-nil 接口也不安全 ( 例如类型非 nil 但值为 nil )

If an interface is `nil`, invoking any methods on it triggers a panic. If an interface is non-nil, you can invoke methods on it. (But note that if the value is `nil` and the methods of the assigned type don’t properly handle `nil`, you could still trigger a panic.)

Since an interface instance with a non-nil type is not equal to `nil`, it is not straight-forward to tell whether or not the value associated with the interface is `nil` when the type is non-nil. You must use reflection to find out.  

### 空接口一般和反射或序列化相关

Sometimes in a statically typed language, you need a way to say that a variable could store a value of any type. Go uses `interface{}` to represent this. Because an empty interface doesn’t tell you anything about the value it represents, there isn’t a lot you can do with it. 

Avoid using `interface{}`. If you see a function that takes in an empty interface, it’s likely that it is using reflection to either populate or read the value. These situations should be relatively rare.

### 经常能看到用函数实现接口的套路

Go allows methods on any user-defined type, including user-defined function types.  
They allow functions to implement interfaces. The most common usage is for HTTP handlers.

```go
func Test函数适配器(t *testing.T) {
    f := func(w http.ResponseWriter, r *http.Request) {} // f 是一个普通函数、http.HandlerFunc 是一个函数类型
    handler := http.HandlerFunc(f)                       // f 与 http.HandlerFunc 底层类型相同,  允许类型转换
    handler.ServeHTTP(nil, nil)                          // 转换后多了个 ServeHTTP 方法! 实现了 http.Handler 接口

    router := http.NewServeMux()
    router.Handle("/xxx", http.HandlerFunc(f)) // 所以有了这种惯用写法, 直接用函数实现接口, 不必定义 type
}
```

By using a type conversion to `http.HandlerFunc`, any function that has the signature `func(http.ResponseWriter,*http.Request)` can be used as an `http.Handler`. This lets you implement HTTP handlers using functions, methods, or closures that meet the `http.Handler` interface. 

### 接口被设置了类型信息，就不等于 nil

#### ➤ 什么是 nil interface ?

An interface is two fields:

1. A pointer to some type-specific information. You can think of this as "type."
2. Data pointer. If the data stored is a pointer, it’s stored directly.  
   If the data stored is a value, then a pointer to the value is stored.
3. Interface variables will be "nil" only when their `type` and `value` fields are "nil".

```go
func main() {
    var m map[string]int
    var i111 interface{} = m // type 字段不为 nil,  value 字段为 nil
    var i222 interface{}     // 接口尚未被赋值,  所以两个字段都是 nil

    fmt.Println("m    == nil", m == nil)    // true
    fmt.Println("i111 == nil", i111 == nil) // false
    fmt.Println("i222 == nil", i222 == nil) // true

    if i111 != nil {
        m := i111.(map[string]int)
        m["key"]++ // 虽然检测过接口不为 nil，但还是空指针异常了，所以这是个坑
    }
}
```

#### ➤ 函数返回接口时，小心接口被具体类型污染 ( 总之想返回 nil interface 那么可以直接显式 `return nil` )

```go
func DeleteFile(path string) error {
    var err *os.PathError
    if path == "/root" {
        err = &os.PathError{}
    }
    return err // err 的类型会让返回值永远不可能是 nil interface
}

func main() {
    fmt.Println(DeleteFile("abc.txt") == nil) // false
}
```
