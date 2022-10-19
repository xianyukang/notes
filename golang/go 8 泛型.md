## Table of Contents
  - [Why people want generics?](#Why-people-want-generics)
    - [概述](#%E6%A6%82%E8%BF%B0)
    - [没有泛型的诸多限制](#%E6%B2%A1%E6%9C%89%E6%B3%9B%E5%9E%8B%E7%9A%84%E8%AF%B8%E5%A4%9A%E9%99%90%E5%88%B6)
    - [Idiomatic Go and Generics](#Idiomatic-Go-and-Generics)
    - [什么时候适合用泛型?](#%E4%BB%80%E4%B9%88%E6%97%B6%E5%80%99%E9%80%82%E5%90%88%E7%94%A8%E6%B3%9B%E5%9E%8B)
  - [使用泛型](#%E4%BD%BF%E7%94%A8%E6%B3%9B%E5%9E%8B)
    - [基础使用](#%E5%9F%BA%E7%A1%80%E4%BD%BF%E7%94%A8)
    - [类型约束](#%E7%B1%BB%E5%9E%8B%E7%BA%A6%E6%9D%9F)
    - [语法糖](#%E8%AF%AD%E6%B3%95%E7%B3%96)
    - [泛型接口](#%E6%B3%9B%E5%9E%8B%E6%8E%A5%E5%8F%A3)
    - [泛型函数](#%E6%B3%9B%E5%9E%8B%E5%87%BD%E6%95%B0)
    - [显式指定类型](#%E6%98%BE%E5%BC%8F%E6%8C%87%E5%AE%9A%E7%B1%BB%E5%9E%8B)
  - [Type Set](#Type-Set)
    - [操作符取交集](#%E6%93%8D%E4%BD%9C%E7%AC%A6%E5%8F%96%E4%BA%A4%E9%9B%86)
    - [声明过的方法才能用](#%E5%A3%B0%E6%98%8E%E8%BF%87%E7%9A%84%E6%96%B9%E6%B3%95%E6%89%8D%E8%83%BD%E7%94%A8)
  - [问题](#%E9%97%AE%E9%A2%98)
    - [试解释图中的问题](#%E8%AF%95%E8%A7%A3%E9%87%8A%E5%9B%BE%E4%B8%AD%E7%9A%84%E9%97%AE%E9%A2%98)

## Why people want generics?

### 概述

<font color='#D05'>Generics Reduce Repetitive Code and Increase Type Safety.</font>

If we wanted a binary tree for strings or float64s and we wanted type safety, we would have to write a custom tree for each type. *That’s verbose and error-prone*. It would be nice to write a single data structure that could handle any type that can be compared with a `<`, but Go doesn’t let you do that today. 

While data structures without generics are inconvenient, the real limitation is in writing functions. Rather than write multiple functions to handle different numeric types, Go implements functions like `math.Max`, `math.Min`, and `math.Mod` using `float64` parameters, which have a range big enough to represent nearly every other numeric type exactly. (The exceptions are an int, int64, or uint with a value greater than 2^53 – 1 or less than –2^53 – 1)

### 没有泛型的诸多限制

You also cannot create a new instance of a variable that’s specified by interface, nor can you specify that two parameters that are of the same interface type are also of the same concrete type. Go also doesn’t provide a way to process a slice of any type; you cannot assign a `[]string` or `[]int` to a variable of `[]interface{}`. This means functions that operate on slices have to be repeated for each type of slice, unless you resort to reflection and give up some performance along with compile-time type safety (this is how `sort.Slice` works).

The result is that many common algorithms, such as `map`, `reduce`, and `filter`, end up being reimplemented for different types. While simple algorithms are easy enough to copy, many (if not most) software engineers find it grating to duplicate code simply because the compiler isn’t smart enough to do it automatically.

### Idiomatic Go and Generics

Adding generics clearly changes some of the advice for how to use Go idiomatically. The use of `float64` to represent any numeric type will end. We will no longer use `interface{}` to represent any possible value in a data structure or function parameter. You can handle different slice types with a single function. *But don’t feel the need to switch all of your code over to using type parameters immediately*. Your old code will still work as new design patterns are invented and refined.  

### 什么时候适合用泛型?

- Functions that work on slices, maps, and channels of any element type.
- General purpose data structures. For exmaple, a linked list or binary tree.
- When a method looks the same for all types. 比如排序算法.

## 使用泛型

### 基础使用

![image-20220731192911891](https://static.xianyukang.com/img/image-20220731192911891.png) 

### 类型约束

按下图添加一个 Contains 方法会失败,  因为 T 的类型约束是 any,  不能对 any 使用 `==` 操作符:

![image-20220731193213863](https://static.xianyukang.com/img/image-20220731193213863.png) 

Just as `interface{}` doesn’t say anything, neither does `any`. We can only store values of any type and retrieve them. To use `==` and `!=`, we need a different type: `type Stack[T comparable] struct {...}`.  

➤ 更多类型约束可以下载 golang.org/x/exp/constraints 包,  比如 constraints.Ordered 表示可比较大小:

![image-20220801193358580](https://static.xianyukang.com/img/image-20220801193358580.png) 

➤ Golang 用接口表示类型约束,  如下是 Ordered 的实现, 其中 `~string` 表示任意底层类型为 `string` 的类型:

![image-20220801194059456](https://static.xianyukang.com/img/image-20220801194059456.png) 

### 语法糖

(1) 用 `interface{}` 表示类型约束: [S <font color='#D05'>interface{~[]E}</font>, E <font color='#2E66FF'>interface{}</font>]  
    这里对 S 的约束: 它是一个 E 列表  
    这里对 E 的约束: 它可以是任意类型

(2) 可以省略 `interface{}`: [S ~[]E, E interface{}]

(3) 可以用 `any` 表示 interface{}: [S ~[]E, E any]

### 泛型接口

![image-20220731214239340](https://static.xianyukang.com/img/image-20220731214239340.png) 

### 泛型函数

Earlier we mentioned that not having generics made it difficult to write map, reduce, and filter implementations that work for all types. Generics make it easy.  

![image-20220801150619604](https://static.xianyukang.com/img/image-20220801150619604.png) 

### 显式指定类型

In some situations, type inference isn’t possible (for example, when a type parameter is only used as a return value). When that happens, all of the type arguments must be specified:  

![image-20220801153537958](https://static.xianyukang.com/img/image-20220801153537958.png) 





## Type Set

### 操作符取交集

While we can use our Tree with built-in types when there’s an associated type that meets Orderable, it might be nice to use a Tree with built-in types that didn’t require the wrappers. To do that, we need a way to specify that we can use the `<` operator. Go generics do that with a type set, which is simply a list of types specified within an interface:  

![image-20220801144739900](https://static.xianyukang.com/img/image-20220801144739900.png) 

When a type constraint is specified with an interface containing a type list, *any listed type can match. However, the allowed operators are the ones that apply for all of the listed types*. In this case, those are the operators ==, !=, >, <, >=, <=, and +.   

Type lists also specify which constants can be assigned to variables of the generic type. There are no constants that can be assigned to all of the listed types in BuiltInOrdered, so you cannot assign a constant to a variable of that generic type.  

### 声明过的方法才能用

Specifying a user-defined type does not give you access to the methods on the type. Any methods on user-defined types that you want to access must be defined in the interface, but all types specified in the type list must implement those methods, or they will not match the interface.

在下面的例子中,  `int` 类型虽然能匹配 `~int`,  但没有实现 String 方法,  所以 DoubleString 函数不能传 int

![image-20220801163402061](https://static.xianyukang.com/img/image-20220801163402061.png) 

## 问题

### 试解释图中的问题

➤ 为什么 `Scale(p, 10)` 返回的是 `[]int` 类型?

![image-20220801201717300](https://static.xianyukang.com/img/image-20220801201717300.png) 

调用 Scale 时,  E 被实例化成 int,  所以返回值是 `[]int`,  而不是 `Point`.  
解决办法如下,  S 可以是任意整数切片,  调用 Scale2 时 S 被实例化成 Point,  所以返回值也是 Point 类型:

![image-20220801202312424](https://static.xianyukang.com/img/image-20220801202312424.png) 

