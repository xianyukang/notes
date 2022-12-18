## Table of Contents
  - [Methods](#Methods)
    - [指针接收器和值接收器](#%E6%8C%87%E9%92%88%E6%8E%A5%E6%94%B6%E5%99%A8%E5%92%8C%E5%80%BC%E6%8E%A5%E6%94%B6%E5%99%A8)
    - [指针的 method set 包含更多方法](#%E6%8C%87%E9%92%88%E7%9A%84-method-set-%E5%8C%85%E5%90%AB%E6%9B%B4%E5%A4%9A%E6%96%B9%E6%B3%95)
    - [能在值变量上调用指针方法](#%E8%83%BD%E5%9C%A8%E5%80%BC%E5%8F%98%E9%87%8F%E4%B8%8A%E8%B0%83%E7%94%A8%E6%8C%87%E9%92%88%E6%96%B9%E6%B3%95)
    - [在 nil 指针上调用方法?](#%E5%9C%A8-nil-%E6%8C%87%E9%92%88%E4%B8%8A%E8%B0%83%E7%94%A8%E6%96%B9%E6%B3%95)
    - [能传函数的地方也能传方法](#%E8%83%BD%E4%BC%A0%E5%87%BD%E6%95%B0%E7%9A%84%E5%9C%B0%E6%96%B9%E4%B9%9F%E8%83%BD%E4%BC%A0%E6%96%B9%E6%B3%95)
    - [两个限制](#%E4%B8%A4%E4%B8%AA%E9%99%90%E5%88%B6)
  - [Types](#Types)
    - [类型声明不是继承](#%E7%B1%BB%E5%9E%8B%E5%A3%B0%E6%98%8E%E4%B8%8D%E6%98%AF%E7%BB%A7%E6%89%BF)
    - [底层类型相同时允许类型转换](#%E5%BA%95%E5%B1%82%E7%B1%BB%E5%9E%8B%E7%9B%B8%E5%90%8C%E6%97%B6%E5%85%81%E8%AE%B8%E7%B1%BB%E5%9E%8B%E8%BD%AC%E6%8D%A2)
    - [使用 iota 做枚举](#%E4%BD%BF%E7%94%A8-iota-%E5%81%9A%E6%9E%9A%E4%B8%BE)
    - [使用 Embedding 来组合类型](#%E4%BD%BF%E7%94%A8-Embedding-%E6%9D%A5%E7%BB%84%E5%90%88%E7%B1%BB%E5%9E%8B)
    - [嵌入一个类型、并重定义某些方法](#%E5%B5%8C%E5%85%A5%E4%B8%80%E4%B8%AA%E7%B1%BB%E5%9E%8B%E5%B9%B6%E9%87%8D%E5%AE%9A%E4%B9%89%E6%9F%90%E4%BA%9B%E6%96%B9%E6%B3%95)
  - [Interfaces](#Interfaces)
    - [定义接口](#%E5%AE%9A%E4%B9%89%E6%8E%A5%E5%8F%A3)
    - [隐式接口与鸭子类型](#%E9%9A%90%E5%BC%8F%E6%8E%A5%E5%8F%A3%E4%B8%8E%E9%B8%AD%E5%AD%90%E7%B1%BB%E5%9E%8B)
    - [也能显式声明要实现的接口](#%E4%B9%9F%E8%83%BD%E6%98%BE%E5%BC%8F%E5%A3%B0%E6%98%8E%E8%A6%81%E5%AE%9E%E7%8E%B0%E7%9A%84%E6%8E%A5%E5%8F%A3)
    - [嵌入接口](#%E5%B5%8C%E5%85%A5%E6%8E%A5%E5%8F%A3)
    - [入参用接口、反参用具体类型](#%E5%85%A5%E5%8F%82%E7%94%A8%E6%8E%A5%E5%8F%A3%E5%8F%8D%E5%8F%82%E7%94%A8%E5%85%B7%E4%BD%93%E7%B1%BB%E5%9E%8B)
    - [接口包含 type 和 value 两个字段](#%E6%8E%A5%E5%8F%A3%E5%8C%85%E5%90%AB-type-%E5%92%8C-value-%E4%B8%A4%E4%B8%AA%E5%AD%97%E6%AE%B5)
    - [关于空接口](#%E5%85%B3%E4%BA%8E%E7%A9%BA%E6%8E%A5%E5%8F%A3)
    - [Type Assertions and Type Switches](#Type-Assertions-and-Type-Switches)
    - [Function Types Are a Bridge to Interfaces](#Function-Types-Are-a-Bridge-to-Interfaces)

## Methods

### 指针接收器和值接收器

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

- If your method modifies the receiver, you must use a pointer receiver.
- If your method needs to handle `nil` instances, then it must use a pointer receiver.
- If your method doesn’t modify the receiver, you can use a value receiver.

When a type has any pointer receiver methods, a common practice is to be consistent and use pointer receivers for all methods, even the ones that don’t modify the receiver.

### 指针的 method set 包含更多方法

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

One thing you might notice is that we were able to call the pointer receiver method even though `c` is a value type. When you use a pointer method with a local variable that’s a value type, Go automatically converts it to a pointer type. In this case, `c.Add()` is converted to `(&c).Add()`. But `c().Add()` can't be converted to `(&c()).Add()`.

It's OK to call a pointer receiver method on a value as long as the value is addressable. Not every variable is addressable though. Map elements are not addressable (you can't `&someMap["key"]`).

```go
func (c *Character) SayHi() {
    fmt.Println("Hi, I am", c.Name)
}

func main() {
    var c = Character{Name: "Cloud", From: "FF7", Age: 21}
    c.SayHi()            // c 是值类型、并且 addressable
    (&c).SayHi()         // 所以编译器会自动把上一行重写成这一行

    var foo SayHier = &c // 这里必须用 &c 取指针, 如果用 c 会编译错误
    foo.SayHi()          // 因为 c 是值类型,  c 的方法集不包括 SayHi()
}

type SayHier interface {
    SayHi()
}
```

### 在 nil 指针上调用方法?

What happens when you call a method on a `nil` instance? In most languages, this produces some sort of error.  Go does something a little different. It actually tries to invoke the method. If it’s a method with a value receiver, you’ll get a panic, as there is no value being pointed to by the pointer. If it’s a method with a pointer receiver, it can work if the method is written to handle the possibility of a nil instance. 

如果方法使用 pointer receiver,  且在方法中处理了指针为 `nil` 的情况,  那么在 `nil` 上调用方法也不会出错.

However, most of the time `nil` receiver is not very useful. If your method has a pointer receiver and won’t work for a `nil` receiver, check for `nil` and return an error.

### 能传函数的地方也能传方法

Methods in Go are so much like functions that you can use a method *as a replacement for a function any time there’s a variable or parameter of a function type*. We can assign the method to a variable or pass it to a parameter. This is called a *method value*. A method value is a bit like a closure, since it can access the values in the fields of the instance from which it was created.  

![image-20220512105855885](https://static.xianyukang.com/img/image-20220512105855885.png) 

#### ➤ `ichigo.greet` 与 `Boy.greet` 的区别?

You can also create a function from the type itself. This is called a *method expression*. In the case of a method expression, the first parameter is the receiver for the method.

![image-20220512111201885](https://static.xianyukang.com/img/image-20220512111201885.png) 

### 两个限制

#### ➤ [参考回答](https://stackoverflow.com/a/65441367)

1. 如果 base type 是指针或接口,  则不能添加任何方法,  例如 `type Counter *int`
2. 无法为另一个包中的类型添加方法,  例如 `func (t *time.Time) Hello() {}`

## Types

### 类型声明不是继承

In addition to declaring types based on built-in Go types and struct literals, you can also declare a user-defined type based on another user-defined type:  `type HighScore Score`.

Declaring a type based on another type looks a bit like inheritance, but it isn’t. The two types have the same underlying type, but that’s all. There is no hierarchy between these types. You can’t assign an instance of type `HighScore` to a variable of type `Score` or vice versa without a type conversion. Furthermore, any methods defined on `Score` aren’t defined on `HighScore`. A type conversion between types that share an underlying type keeps the same underlying data but associates different methods.  

### 底层类型相同时允许类型转换

`h := Hour(1)`   // 如果 Hour 和 Second 的底层类型都是 int, 那么 Hour、Second、int 之间能相互转换  
`s := Second(h)` // 虽然允许这么转,  但有时候逻辑上是错的,  比如 1 小时应该等于 3600 秒

### 使用 iota 做枚举

Go doesn’t have an enumeration type. Instead, it has `iota`, which lets you assign an increasing value to a set of constants. When using `iota`, the best practice is:

![image-20220512155100349](https://static.xianyukang.com/img/image-20220512155100349.png) 

#### ➤ `iota` 的值是它在 constant block 中的索引

When the Go compiler sees this constant block, it repeats the type and the assignment to all of the subsequent constants in the block, and increments the value of `iota` on each line. This means that it assigns 0 to the first constant (Uncategorized), 1 to the second constant (Personal), and so on. When a new const block is created, `iota` is set back to 0.

#### ➤ 在 Personal 上面加一行 XXX 会让后续所有常量的值都加一

If you insert a new identifier in the middle of your list of literals, all of the subsequent ones will be renumbered. This will break your application in a subtle way if those constants represented values in another system or in a database. 

#### ➤ 什么时候别用 `iota`,  应该显式指定常量值?

Use iota for “internal” purposes only. iota-based enumerations only make sense when you care about being able to differentiate between a set of values, and don’t particularly care what the value is behind the scenes. If the actual value matters, specify it explicitly.  

#### ➤ 把首个 `iota` 命名为 Invalid

Be aware that iota starts numbering from 0.  When mail first arrives, it is uncategorized, so the zero value makes sense. If there isn’t a sensical default value for your constants, a common pattern is to assign the first `iota` value in the constant block to a constant that indicates the value is invalid.

### 使用 Embedding 来组合类型

The software engineering advice “Favor object composition over class inheritance”. While Go doesn’t have inheritance, it encourages code reuse via built-in support for composition and promotion.

#### ➤ 嵌入另一个类型的所有方法和字段

Note that `Manager` contains a field of type `Employee`, but no name is assigned to that field. This makes `Employee` an `embedded field`. Any fields or methods declared on an embedded field are promoted to the containing struct and can be invoked directly on it.

![image-20220512170745813](https://static.xianyukang.com/img/image-20220512170745813.png) 

#### ➤ 嵌入不是继承

You cannot assign a variable of type `Manager` to a variable of type `Employee`.  
If you want to access the Employee field in Manager, you must do so explicitly: `var e = m.Employee`.

#### ➤ 调用被提升的方法时,  receiver 不是外层类型

When we embed a type, the methods of that type become methods of the outer type, but when they are invoked the receiver of the method is the inner type, not the outer one. 

```go
type Manager struct { company.Employee }          // 所谓嵌入 company.Employee 类型
type Manager struct { Employee company.Employee } // 相当于添加一个 Employee 字段,  但是方法被提升到外层
m.Name   <=> m.Employee.Name
m.Work() <=> m.Employee.Work()                    // m.Work() 的 receiver 是 m.Employee 而不是 m
```

### 嵌入一个类型、并重定义某些方法

![image-20221015194510999](https://static.xianyukang.com/img/image-20221015194510999.png) 

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

### 也能显式声明要实现的接口

If `json.RawMessage` needs a custom JSON representation, it should implement `json.Marshaler`, but there are no static conversions that would cause the compiler to verify this automatically. To guarantee that the implementation is correct, a global declaration using the blank identifier can be used in the package:

```go
var _ json.Marshaler = (*RawMessage)(nil) // 让编译器检查 RawMessage 是否实现了 json.Marshaler 接口
```

### 嵌入接口

Just like you can embed a type in a struct, you can also embed an interface in an interface. For example, the `io.ReadCloser` interface is built out of an `io.Reader` and an `io.Closer`. Just like you can embed a concrete type in a struct, you can also embed an interface in a struct. 

#### ➤ [在结构体中嵌入一个接口有何好处?](https://stackoverflow.com/a/24537547) 相当于添加一个字段,  但类型为接口,  并且接口方法提升到结构体

```go
type Weapon interface{ WeaponAbility() string }
type WeaponFunc func() string
func (f WeaponFunc) WeaponAbility() string { return f() }

type Character struct {
    Name   string
    From   string
    Age    int
    Weapon // 嵌入一个接口
}

func main() {
    var dante = Character{Name: "但丁", From: "DMC", Age: 40}
    var 魔剑但丁 = WeaponFunc(func() string { return "咿呀剑法" })
    var 黑檀木与白象牙 = WeaponFunc(func() string { return "旋转跳跃突突突" })

    dante.Weapon = 魔剑但丁
    fmt.Println(dante.WeaponAbility())
    dante.Weapon = 黑檀木与白象牙
    fmt.Println(dante.WeaponAbility())
}
```

### 入参用接口、反参用具体类型

#### ➤ 入参应该用「 接口 」

如果 `CreateUserService` 用 `DB` 接口作为入参,  而不是依赖某个具体类型,  
那么能方便地切换实现,  比如从 mysql 切换到 postgres 数据库.  (官方推荐让 consumer 定义它用到的接口)

#### ➤ 返回值应该用「 具体类型 」

The implementing package should return concrete (usually pointer or struct) types: that way, new methods can be added to implementations without requiring extensive refactoring. Because adding a new method to an interface means that you need to update all existing implementations of the interface, or your code breaks. 

Do not define interfaces before they are used: without a realistic example of usage, it is too difficult to see whether an interface is even necessary, let alone what methods it ought to contain.

### 接口包含 type 和 value 两个字段

We also use `nil` to represent the zero value for an interface instance, but it’s not as simple as it is for concrete types. In order for an interface to be considered `nil` **both the type and the value must be nil**. In the Go runtime, interfaces are implemented as a pair of pointers, one to the underlying type and one to the underlying value. As long as the type is non-nil, the interface is non-nil.  

![image-20220512211552639](https://static.xianyukang.com/img/image-20220512211552639.png) 

What `nil` indicates for an interface is whether or not you can invoke methods on it. As we covered earlier, you can invoke methods on `nil` concrete instances, so it makes sense that you can invoke methods on an interface variable that was assigned a nil concrete instance. 

#### ➤ 用 nil 接口调用方法会报错,  当接口的 value 为 nil 时也可能报错

If an interface is `nil`, invoking any methods on it triggers a panic. If an interface is non-nil, you can invoke methods on it. (But note that if the value is nil and the methods of the assigned type don’t properly handle nil, you could still trigger a panic.)

Since an interface instance with a non-nil type is not equal to `nil`, it is not straight-forward to tell whether or not the value associated with the interface is nil when the type is non-nil. You must use reflection to find out.  

### 关于空接口

Sometimes in a statically typed language, you need a way to say that a variable could store a value of any type. Go uses `interface{}` to represent this. Because an empty interface doesn’t tell you anything about the value it represents, there isn’t a lot you can do with it. 

Avoid using `interface{}`. If you see a function that takes in an empty interface, it’s likely that it is using reflection to either populate or read the value. These situations should be relatively rare.

### Type Assertions and Type Switches

Go provides two ways to see if a variable of an interface type has a specific concrete type or if the concrete type implements another interface.

![image-20220512233721375](https://static.xianyukang.com/img/image-20220512233721375.png) 



### Function Types Are a Bridge to Interfaces

Go allows methods on any user-defined type, including user-defined function types.  
They allow functions to implement interfaces. The most common usage is for HTTP handlers.

![image-20220513133303489](https://static.xianyukang.com/img/image-20220513133303489.png) 

By using a type conversion to `http.HandlerFunc`, any function that has the signature `func(http.ResponseWriter,*http.Request)` can be used as an `http.Handler`. This lets you implement HTTP handlers using functions, methods, or closures that meet the `http.Handler` interface. 
