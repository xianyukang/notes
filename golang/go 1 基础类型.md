## Table of Contents
  - [数值类型](#%E6%95%B0%E5%80%BC%E7%B1%BB%E5%9E%8B)
    - [全部数值类型](#%E5%85%A8%E9%83%A8%E6%95%B0%E5%80%BC%E7%B1%BB%E5%9E%8B)
    - [16 进制整数字面量](#16-%E8%BF%9B%E5%88%B6%E6%95%B4%E6%95%B0%E5%AD%97%E9%9D%A2%E9%87%8F)
    - [数值类型的别名](#%E6%95%B0%E5%80%BC%E7%B1%BB%E5%9E%8B%E7%9A%84%E5%88%AB%E5%90%8D)
    - [怎么选择数值类型?](#%E6%80%8E%E4%B9%88%E9%80%89%E6%8B%A9%E6%95%B0%E5%80%BC%E7%B1%BB%E5%9E%8B)
    - [浮点数](#%E6%B5%AE%E7%82%B9%E6%95%B0)
    - [数学上的复数](#%E6%95%B0%E5%AD%A6%E4%B8%8A%E7%9A%84%E5%A4%8D%E6%95%B0)
  - [切片](#%E5%88%87%E7%89%87)
    - [数组](#%E6%95%B0%E7%BB%84)
    - [切片结构](#%E5%88%87%E7%89%87%E7%BB%93%E6%9E%84)
    - [make 函数](#make-%E5%87%BD%E6%95%B0)
    - [空切片和 nil 切片](#%E7%A9%BA%E5%88%87%E7%89%87%E5%92%8C-nil-%E5%88%87%E7%89%87)
    - [子切片和源切片共享同一底层数组](#%E5%AD%90%E5%88%87%E7%89%87%E5%92%8C%E6%BA%90%E5%88%87%E7%89%87%E5%85%B1%E4%BA%AB%E5%90%8C%E4%B8%80%E5%BA%95%E5%B1%82%E6%95%B0%E7%BB%84)
    - [小心对子切片调用 append](#%E5%B0%8F%E5%BF%83%E5%AF%B9%E5%AD%90%E5%88%87%E7%89%87%E8%B0%83%E7%94%A8-append)
    - [指向切片的指针有什么用?](#%E6%8C%87%E5%90%91%E5%88%87%E7%89%87%E7%9A%84%E6%8C%87%E9%92%88%E6%9C%89%E4%BB%80%E4%B9%88%E7%94%A8)
    - [copy 函数](#copy-%E5%87%BD%E6%95%B0)
    - [拼接两个切片](#%E6%8B%BC%E6%8E%A5%E4%B8%A4%E4%B8%AA%E5%88%87%E7%89%87)
  - [字符串](#%E5%AD%97%E7%AC%A6%E4%B8%B2)
    - [字符串如何工作](#%E5%AD%97%E7%AC%A6%E4%B8%B2%E5%A6%82%E4%BD%95%E5%B7%A5%E4%BD%9C)
    - [字符串的 for-range 循环](#%E5%AD%97%E7%AC%A6%E4%B8%B2%E7%9A%84-forrange-%E5%BE%AA%E7%8E%AF)
    - [rune 字面量](#rune-%E5%AD%97%E9%9D%A2%E9%87%8F)
    - [其他细节](#%E5%85%B6%E4%BB%96%E7%BB%86%E8%8A%82)
  - [Map](#Map)
    - [创建 Map](#%E5%88%9B%E5%BB%BA-Map)
    - [区分零值与未设置过值](#%E5%8C%BA%E5%88%86%E9%9B%B6%E5%80%BC%E4%B8%8E%E6%9C%AA%E8%AE%BE%E7%BD%AE%E8%BF%87%E5%80%BC)
    - [Go 没有 Set 类型但可以用 Map 实现](#Go-%E6%B2%A1%E6%9C%89-Set-%E7%B1%BB%E5%9E%8B%E4%BD%86%E5%8F%AF%E4%BB%A5%E7%94%A8-Map-%E5%AE%9E%E7%8E%B0)
  - [Struct](#Struct)
    - [结构体字面量](#%E7%BB%93%E6%9E%84%E4%BD%93%E5%AD%97%E9%9D%A2%E9%87%8F)
    - [匿名结构体](#%E5%8C%BF%E5%90%8D%E7%BB%93%E6%9E%84%E4%BD%93)
  - [指针](#%E6%8C%87%E9%92%88)
    - [理解 nil](#%E7%90%86%E8%A7%A3-nil)
    - [理解 new 函数](#%E7%90%86%E8%A7%A3-new-%E5%87%BD%E6%95%B0)
    - [无法把字面量赋值给指针](#%E6%97%A0%E6%B3%95%E6%8A%8A%E5%AD%97%E9%9D%A2%E9%87%8F%E8%B5%8B%E5%80%BC%E7%BB%99%E6%8C%87%E9%92%88)
  - [指针使用建议](#%E6%8C%87%E9%92%88%E4%BD%BF%E7%94%A8%E5%BB%BA%E8%AE%AE)
    - [使用 Pointer 还是 Value?](#%E4%BD%BF%E7%94%A8-Pointer-%E8%BF%98%E6%98%AF-Value)
    - [The Zero Value Versus No Value](#The-Zero-Value-Versus-No-Value)
    - [不要返回 nil 表示数据不存在](#%E4%B8%8D%E8%A6%81%E8%BF%94%E5%9B%9E-nil-%E8%A1%A8%E7%A4%BA%E6%95%B0%E6%8D%AE%E4%B8%8D%E5%AD%98%E5%9C%A8)
    - [栈/堆上分配](#%E6%A0%88%E5%A0%86%E4%B8%8A%E5%88%86%E9%85%8D)

## 数值类型

### 全部数值类型

![image-20220427000932962](https://static.xianyukang.com/img/image-20220427000932962.png) 

### 16 进制整数字面量

Integer literals are normally base ten, but different prefixes are used to indicate other bases:  

- 0b for binary (base two), 
- 0o for octal (base eight), 
- 0x for hexadecimal (base sixteen). 
- 数字中可以用下划线, 比如 1_000 表示 1000,  另外 `0644` 也是八进制数但更推荐写成 `0o644`

### 数值类型的别名

Go does have some special names for integer types. A `byte` is an alias for `uint8`. A `rune` is an alias for `int32` and is equivalent to int32 in all ways. The second special name is `int`. On a 32-bit CPU, int is a 32-bit signed integer like an int32. On most 64-bit CPUs, int is a 64-bit signed integer, just like an int64.

Because int isn’t consistent from platform to platform, it is a compile-time error to assign, compare, or perform mathematical operations between an int and an int32 or int64 without a type conversion.

### 怎么选择数值类型?

Given all of these choices, you might wonder when you should use each of them. There are three simple rules to follow:

1. If you are working with a binary file format or network protocol that has an integer of a specific size or sign, use the corresponding integer type.
2. If you are writing a library function that should work with any integer type, write a pair of functions, one with `int64` for the parameters and variables and the other with `uint64`. 
3. <font color='#D05'>In all other cases, just use int</font>. Unless you need to be explicit about the size or sign of an integer for performance or integration purposes, use the int type. Consider any other type to be a premature optimization until proven otherwise.

其中第二条是指用 int64 和 uint64 来处理所有整数类型,  这在 Go 1.18 之前是惯用写法 (因为之前没泛型)

> Using int64 and uint64 means that you can write the code once and let your callers use type conversions to pass values in and convert data that’s returned. You can see this pattern in the Go standard library with the functions `FormatInt` and `ParseInt` in the strconv package.

### 浮点数

一般推荐用 float64 而不是 float32. It also helps mitigate floating point accuracy issues since a float32 only has six- or seven-decimal digits of precision. Don’t worry about the difference in memory size unless you have used the profiler to determine that it is a significant source of problems. 

<font color='#D05'>A floating point number cannot represent a decimal value exactly</font>.  
<font color='#D05'>Do not use them to represent money or any other value that must have an exact decimal representation!</font>

#### ➤ 整数除以 0 会 panic, 但浮点数除以 0 不会 panic、且有三种返回值

Floating point division has a couple of interesting properties. Dividing a nonzero floating point variable by 0 returns `+Inf` or `-Inf` (positive or negative infinity), depending on the sign of the number. Dividing a floating point variable set to 0 by 0 returns `NaN` (Not a Number).

#### ➤ 无论用什么编程语言,  比较浮点数都不能用 `==`、`!=`

While Go lets you use == and != to compare floats, don’t do it. Due to the inexact nature of floats, two floating point values might not be equal when you think they should be. Instead, define a maximum allowed variance and see if the difference between two floats is less than that.

### 数学上的复数

There is one more numeric type and it is pretty unusual. Go has first-class support for `complex` numbers.  If you don’t know what complex numbers are, you are not the target audience for this feature; feel free to skip ahead. 所以可以先跳过.

## 切片

### 数组

golang 数组是一种值类型，数组本身的赋值和函数传参都是以整体复制的方式处理的.  
若在函数中修改数组,  记得传数组指针,  否则改的是拷贝,  外面看不到

<font color='#D05'>另外数组的长度是数组类型的一部分，两个不同长度的数组属于不同的类型</font>，这意味着

1. `[2]int` 无法赋值给 `[3]int`
2. 一个处理 [2]int 类型的函数只能处理 [2]int,  无法写出一个函数来处理任意长度的数组
3. 无法用一个变量来指定数组长度, 因为数组的长度必须在编译时确定

### 切片结构

Golang 用 `切片结构体` 来表示一个切片,  它包含 length、capacity 和一个 array 指针

![image-20220511110657088](https://static.xianyukang.com/img/image-20220511110657088.png) 

Every slice has a `capacity`, which is the number of consecutive memory locations reserved.  
如果 len < cap 那么调用 `append(slice, 123)` 往切片中添加元素时,  会把元素放到空余位置,  并返回新的切片结构体.  

If you try to add additional values when the length equals the capacity, the `append` function uses the Go runtime to allocate a new slice with a larger capacity. The values in the original slice are copied to the new slice, the new values are added to the end, and the new slice is returned.

#### ➤ 为什么 append 后要赋值给原变量

![image-20220511133850765](https://static.xianyukang.com/img/image-20220511133850765.png) 

#### ➤ 切片容量增长策略

The rules as of Go 1.14 are to double the size of the slice when the capacity is less than 1,024 and then grow by at least 25% afterward. If you know how many things you plan to put into a slice, create the slice with the correct initial capacity. We do that with the `make` function.

### make 函数

`make` allows us to specify the type, length, and, optionally, the capacity. The built-in function `make(T,args)` serves a purpose different from `new(T)`. It creates slices, maps, and channels only.

```go
x := make([]int, 5)    // len:5, cap:5
x := make([]int, 0, 5) // len:0, cap:5, if len > cap, your program will panic at runtime.

// 新手可能会犯下面的错误,  元素 10 会被添加到索引 5,  而不是索引 0:
x := make([]int, 5)  // 这样创建的 slice 并不是空的,  它的长度和容量都为 5,  5 个元素都是零值
x = append(x, 10)    // append always increases the length of a slice!  
```

### 空切片和 nil 切片

You can create a slice using an empty slice literal: `var x = []int{}`  
This creates a zero-length slice, which is non-nil (comparing it to `nil` returns false).  
Otherwise, a nil slice works identically to a zero-length slice.  
The only situation where a zero-length slice is useful is when converting a slice to JSON.

### 子切片和源切片共享同一底层数组

When you take a slice from a slice, you are not making a copy of the data. Instead, you now have two variables that are sharing memory. This means that *changes to an element in a slice affect all slices that share that element*. 

```go
// 从一个数组或 slice 创建另一个 slice 时,  并没有复制底层数组,  只是产生一个新的切片结构体
// 虽然每个子切片有独立的长度和容量信息,  但和源切片使用同一个底层数组
x := []int{1, 2, 3, 4}
y := x[:2]                   // y 反映切片的前两个元素
z := x[1:]                   // z 反映切片的后三个元素
z[0] = 333                   // 修改 z[0] 会同时影响 x,y,z 三个切片
```

### 小心对子切片调用 append

![image-20220511163518149](https://static.xianyukang.com/img/image-20220511163518149.png) 

#### ➤ 为什么 y 的容量是 4 而不是 2

Whenever you take a slice from another slice, the subslice’s capacity is set to the capacity of the original slice, minus the offset of the subslice within the original slice.  
对于 y 切片来说,  源切片的容量是 4,  y 在源切片中的偏移量为 0,  所以容量为 4 - 0 = 4

#### ➤ 为什么往 append(y, 30) 会导致 x 变化

y 的长度为 2 容量为 4,  所以往 y 中添加元素时无需创建更大的数组,  把 30 放在索引 2 就行了  
因为 y 和 x 共享同一个底层数组,  所以 x 中索引为 2 的元素也会变成 30

#### ➤ full slice expression

*To avoid complicated slice situations, you should either never use append with a sub-slice or make sure that append doesn’t cause an overwrite by using a full slice expression*. 

The full slice expression includes a third part, which indicates the last position in the parent slice’s capacity that’s available for the subslice. Subtract the starting offset from this number to get the subslice’s capacity.

比如 `y := x[0:2:2]` 中第三部分的 2 表示,  y 与源切片共享 [0, 2) 这块内存  
比如 `z := x[2:4:4]` 中第三部分的 4 表示,  z 与源切片共享 [2, 4) 这块内存

### 指向切片的指针有什么用?

[Why pointers to slices are useful and how ignoring them can lead to tricky bugs](https://link.medium.com/tOemYJhIAub).  
如果只需修改切片中的元素,  那么函数参数用 `[]string` 类型  
如果需要增/删切片元素,  那么函数参数用 `*[]string` 类型,  否则外面会看不到修改,  
因为切片变量的实质是一个 `reflect.SliceHeader`,  切片按值传递也是传一个结构体

### copy 函数

If you need to create a slice that’s independent of the original, use the built-in `copy` function. The copy function takes two parameters. The first is the destination slice and the second is the source slice. *It copies as many values as it can from source to destination*, limited by whichever slice is smaller, and returns the number of elements copied. The capacity of `x` and `y` doesn’t matter; it’s the length that’s important.

![image-20220511170137047](https://static.xianyukang.com/img/image-20220511170137047.png) 

### 拼接两个切片

#### ➤ [参考回答](https://stackoverflow.com/a/58726780)

可以用 `c := append(a, b...)` 但这并不像 `a := append(a, b...)` 那样安全  
因为 a 和 c 可能使用同一底层数组,  替换 a 中元素、或往 a 添加元素、都会影响到 c

```go
func main() {
    a := make([]int, 3, 6)
    b := []int{1, 1, 1}
    c := append(a, b...) // c 与 a 共用同一底层数组
    a[0] = 666           // 替换 a 中元素会影响到 c
    fmt.Println(c)       // c[0] 也变成了 666
}
```

#### ➤ 可以自己写个 `Append` 函数

```go
func Append(a, b []int) []int {
    newLen := len(a) + len(b)
    newSlice := make([]int, newLen, newLen*2) // Allocate double what's needed, for future growth.
    copy(newSlice, a)
    copy(newSlice[len(a):], b)
    return newSlice
}

func main() {
    a := make([]int, 3, 6)
    b := []int{1, 1, 1}
    c := Append(a, b)              // c 是重新分配的
    a[0] = 666                     //
    fmt.Println(c, len(c), cap(c)) // c[0] 依旧是 0, 不会被 a 影响
}
```



## 字符串

### 字符串如何工作

```go
func how_strings_work_in_go() {
    // You might think that a string in Go is made out of runes, but that’s not the case. 
    // Under the covers, Go uses a sequence of bytes to represent a string. 
    // string 就是一个 utf-8 编码的 byte array, 中文字符占三个字节、狗头占 4 个字节
    str := "你好🐶"
    for i := 0; i < len(str); i++ {
        fmt.Printf("%x ", str[i])
    }
    fmt.Println(str[:3])                     // 要前三个字节才能取出 '你' 这个字
    fmt.Println(utf8.RuneCountInString(str)) // len 返回的是字节数, 计算字符 (unicode码点) 个数要这样写

    // rune 类型是 int32 的别名,  所以 rune 存储的是 unicode 码点编号，而不是字符本身，
    // 如果你把 rune 传递给 fmt.Println，你会在输出中看到一个数字，而不是原始字符。
    fmt.Printf("狗头的 unicode 码点: 0x%x \n", '🐶')
}
```

### 字符串的 for-range 循环

For strings, the `range` does more work for you, breaking out individual Unicode code points by parsing the UTF-8. Erroneous encodings consume one byte and produce the replacement rune U+FFFD.

```go
for pos, char := range "日本\x80語" { // \x80 is an illegal UTF-8 encoding
    fmt.Printf("character %#U starts at byte position %d\n", char, pos)
}
```

prints

```bash
character U+65E5 '日' starts at byte position 0
character U+672C '本' starts at byte position 3
character U+FFFD '�' starts at byte position 6
character U+8A9E '語' starts at byte position 7
```

### rune 字面量

`Rune literals` represent characters and are surrounded by single quotes. Unlike many other languages, in Go single quotes and double quotes are not interchangeable. Rune literals can be written as 

- single Unicode characters ('a'), 
- 8-bit octal numbers ('\141'), 8-bit hexadecimal numbers ('\x61'), 
- 16-bit hexadecimal numbers ('\u0061'), 32-bit Unicode numbers ('\U00000061'). 
- There are also several backslash escaped rune literals, with the most useful ones being 
  newline ('\n'), tab('\t'), single quote ('\''), double quote ('\"'), and backslash ('\\').

### 其他细节

1. 与切片类似、传递字符串参数时，不会复制底层数组, 只需复制一下 `reflect.StringHeader`

2. A single rune or byte can be converted to a string,  `string('x')`、`string(byte('y'))`
   

## Map

### 创建 Map

```go
m := map[string]int{} // 这创建了一个空的 map
var m map[string]int  // 这是一个 nil map, 比较危险
```

A `nil` map behaves like an empty map when reading,  
but attempts to write to a `nil` map will cause a runtime panic; don't do that.

如果事先知道 map 中会有多少键值对, 为了减少底层数组的扩容开销  
You can use `make` to create a map with a default size: `ages := make(map[int]string, 10)`  
Maps created with `make` still have a length of 0, and they can grow past the initially specified size.  

#### ➤ Map Key 必须是可比较的类型

The key for a map can be any comparable type. Go doesn’t require (or even allow) you to define your own hash algorithm or equality definition. Instead, the Go runtime that’s compiled into every Go program has code that implements hash algorithms for all types that are allowed to be keys.  

### 区分零值与未设置过值

When we try to read the value assigned to a map key that was never set,  
the map returns the zero value for the map’s value type.

有时候这很方便,  直接 `m["xxx"]++` 就行了, Java 的 `HashMap<String, Integer>` 则需要显式初始化  
为了判断 map 中是否设置过 xxx 这个键,  可以用第二个返回值 ok ,  `v, ok = m["xxx"]`

### Go 没有 Set 类型但可以用 Map 实现

例如 `intSet := map[int]bool{}`  
Go doesn’t include a set, but you can use a map to simulate some of its features.  
Use the key of the map for the type that you want to put into the set and use a bool for the value.  

If you need sets that provide operations like union, intersection, and subtraction, you can either write one yourself or use one of the many third-party libraries that provide the functionality.

Some people prefer to use `struct{}` for the value when a map is being used to implement a set. The advantage is that an empty struct uses zero bytes, while a boolean uses one byte. 缺点是看起来丑一点, and you need to use the comma ok idiom to check if a value is in the set.

## Struct

### 结构体字面量

A `struct literal` can be specified as a comma-separated list of values for the fields inside of braces:

```go
// When using this struct literal format, a value for every field in the struct must be specified, 
// and the values are assigned to the fields in the order they were declared in the struct definition.
julia := person{ "Julia", 40, "cat"}

// 一般这种形式用的更多
var p = person{
    age: 17,
    name: "ichigo",
}
```

### 匿名结构体

```go
var person struct{ name string }             // person 变量的类型为匿名结构体
person := struct{ name string }{"ichigo"}    // 把匿名结构体的 name 字段初始化为 ichigo
```

There are two common situations where anonymous structs are handy. 

1. The first is when you translate external data into a struct or a struct into external data (like JSON or protocol buffers). This is called unmarshaling and marshaling data. 
2. Writing tests is another place where anonymous structs pop up.

## 指针

### 理解 nil

`nil` is slightly different from the `null` that’s found in other languages. In Go, nil is an identifier that represents the lack of a value for some types. Like the untyped numeric constants we saw in the previous chapter, *nil has no type*, so it can be assigned or compared against values of different types.

`pointer`、`slice`、`map`、`function`、`channel`、`interface` 变量的零值都是 `nil`.  `nil` is an untyped identifier that represents the lack of a value for certain types.  Unlike NULL in C, nil is not another name for 0; you can’t convert it back and forth with a number.

<font color='#D05'>Before dereferencing a pointer, you must make sure that the pointer is non-nil</font>. Your program will panic if you attempt to dereference a nil pointer.

### 理解 new 函数

![image-20220510205008434](https://static.xianyukang.com/img/image-20220510205008434.png) 

#### ➤ `new(Point)` 和 `&Point{}` 都能返回指针, 他们有什么区别?

[参考: Is there a difference between new() and "regular" allocation?](https://stackoverflow.com/questions/13244947/is-there-a-difference-between-new-and-regular-allocation)  
`new()` is the only way to get a pointer to an unnamed integer or other basic type.   
You can write `p := new(int)` but you can't write `p := &int{0}`. Other than that, it's a matter of preference.

#### ➤ `new` 并不负责初始化

`new` is a built-in function that allocates memory, but unlike its namesakes in some other languages it does not *initialize* the memory, it only *zeros* it. That is, `new(T)` allocates zeroed storage for a new item of type `T` and returns its address, a value of type `*T`.

#### ➤ 尽量让类型的零值直接可用

Since the memory returned by `new` is zeroed, it's helpful to arrange when designing your data structures that the zero value of each type can be used without further initialization. This means a user of the data structure can create one with `new` and get right to work. For example, the documentation for `bytes.Buffer` states that "the zero value for `Buffer` is an empty buffer ready to use." Similarly, `sync.Mutex` does not have an explicit constructor or `Init` method. Instead, the zero value for a `sync.Mutex` is defined to be an unlocked mutex. Sometimes the zero value isn't good enough and an initializing constructor is necessary.

#### ➤ new 与 make 的区别

`make` applies only to maps, slices and channels and does not return a pointer.

```go
var p *[]int = new([]int)       // allocates slice structure; *p == nil; rarely useful
var v  []int = make([]int, 100) // the slice v now refers to a new array of 100 ints

var p *[]int = new([]int)       // Unnecessarily complex:
*p = make([]int, 100, 100)
```



### 无法把字面量赋值给指针

![image-20220510211155165](https://static.xianyukang.com/img/image-20220510211155165.png) 

解决办法:

1. 写一个 stringp("ichigo") 函数返回指针类型
2. 定义 name 变量存储 "ichigo",  然后传 &name

## 指针使用建议

### 使用 Pointer 还是 Value?

1. 一般情况都推荐用 Value,  有理由用指针时才会用指针
2. 如果要在函数中修改传入的参数,  那么用 Pointer
3. 如果函数的返回值在后续代码中会被修改,  那么用 Pointer
4. 定义类型的的方法时,  经常用 pointer receiver,  因为经常会定义 `obj.add` 这样的方法来修改对象本身

**Pointers indicate mutable parameters**. That said, you should be careful when using pointers in Go. As discussed earlier, they make it harder to understand data flow and can create extra work for the garbage collector.

#### ➤ 总而言之指针有三个内涵:

1. 数据会被修改
2. 是同一个对象、不是拷贝后的对象 
3. 减少巨型对象拷贝、提升性能 (偶尔会遇到大对象)

#### ➤ Pointer Passing Performance

If a struct is large enough, there are performance improvements from using a pointer to the struct as either an input parameter or a return value. For the vast majority of cases, the difference between using a pointer and a value won’t affect your program’s performance. But if you are passing megabytes of data between functions, consider using a pointer even if the data is meant to be immutable.

### The Zero Value Versus No Value

The other common usage of pointers in Go is to indicate the difference between a variable or field that’s been assigned the zero value and a variable or field that hasn’t been assigned a value at all. If this distinction matters in your program, use a nil pointer to represent an unassigned variable or struct field.

在解析 JSON 请求时,  如果要区分 `age: 0` 和 `根本没传 age 字段` 这两种情况  
需要把 Person 结构体的 Age 字段设为 *int 类型,  因为指针类型的零值是 nil,  Age 等于 nil 说明没传  
如果要区分 `age: null` 和 `不传 age 字段`,  参考 [JSON field set to null vs field not there](https://stackoverflow.com/questions/36601367/json-field-set-to-null-vs-field-not-there)

### 不要返回 nil 表示数据不存在

其他语言返回 null 表示数据不存在,  但 golang 使用第二个返回值 (类型为 bool/error) 来表示数据不存在  
Rather than return a pointer set to nil from a function, use the comma ok idiom that we saw for maps and return a value type and a boolean.

Resist the temptation to use a pointer field to indicate no value. While a pointer does provide a handy way to indicate no value, if you are not going to modify the value, you should use a value type instead, paired with a boolean. Because pointers also indicate mutability.

### 栈/堆上分配

#### ➤ 为什么栈上分配效率高

Allocating memory on the stack is fast and simple. A stack pointer tracks the last location where memory was allocated; allocating additional memory is done by moving the stack pointer. When a function is invoked, a new stack frame is created for the function’s data. *Local variables are stored on the stack, along with parameters passed into a function*. When a function exits, its return values are copied back to the calling function via the stack and the stack pointer is moved back to the beginning of the stack frame for the exited function, deallocating all of the stack memory that was used by that function’s local variables and parameters.  

①因为栈是一块连续的内存,  所以内存访问效率高.  
②因为函数退出后栈中所有数据都变成垃圾,  不涉及复杂的垃圾收集算法,  所以垃圾回收效率高.

#### ➤ 什么东西在栈上分配

To store something on the stack, you have to know exactly how big it is at compile time. When you look at the *value types in Go (primitive values, arrays, and structs)*, they all have one thing in common: we know exactly how much memory they take at compile time. Because their sizes are known, they can be allocated on the stack instead of the heap.

#### ➤ `new(obj)` 创建的对象可能在栈上分配

In order to allocate the data the pointer points to on the stack, several conditions must be true.

1. It must be a local variable whose data size is known at compile time. 
2. The pointer cannot be returned from the function. 
3. If the pointer is passed into a function, the compiler must be able to ensure that these conditions still hold. 

If the size isn’t known, you can’t make space for it by simply moving the stack pointer. If the pointer variable is returned, the memory that the pointer points to will no longer be valid when the function exits.

```go
func count指针所指对象在栈上分配() *int {
    count := new(int)
    if count != nil {
        fmt.Println("haha")
    }    
    return nil // 如果 return count 那么 new(int) 就会在堆上分配
}
```

#### ➤ `new(obj)` 也可能在堆上分配

A common source of bugs in C programs is returning a pointer to a local variable. In C, this results in a pointer pointing to invalid memory. The Go compiler is smarter. When it sees that a pointer to a local variable is returned, the local variable’s value is stored on the heap.

When the compiler determines that the data can’t be stored on the stack, we say that the data the pointer points to *escapes the stack* and the compiler stores the data on the heap. The heap is the memory that’s managed by the garbage collector (or by hand in languages like C and C++).   

#### ➤ 其他细节

Any data that’s stored on the heap is valid as long as it can be tracked back to a pointer type variable on a stack. Once there are no more pointers pointing to that data (or to data that points to that data), the data becomes garbage and it’s the job of the garbage collector to clear it out.

What’s so bad about storing things on the heap? There are two problems related to performance. 

1. First is that the garbage collector takes time to do its work. It isn’t trivial to keep track of all of the available chunks of free memory on the heap or tracking which used blocks of memory still have valid pointers. This is time that’s taken away from doing the processing that your program is written to do.  
2. The second problem deals with the nature of computer hardware. RAM might mean “random access memory,” but the fastest way to read from memory is to read it sequentially. A slice of structs in Go has all of the data laid out sequentially in memory. This makes it fast to load and fast to process. A slice of pointers to structs (or structs whose fields are pointers) has its data scattered across RAM, making it far slower to read and process.  
