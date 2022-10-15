## Table of Contents
  - [语言基础](#%E8%AF%AD%E8%A8%80%E5%9F%BA%E7%A1%80)
    - [创建变量的各种方式](#%E5%88%9B%E5%BB%BA%E5%8F%98%E9%87%8F%E7%9A%84%E5%90%84%E7%A7%8D%E6%96%B9%E5%BC%8F)
    - [全都需要显式类型转换](#%E5%85%A8%E9%83%BD%E9%9C%80%E8%A6%81%E6%98%BE%E5%BC%8F%E7%B1%BB%E5%9E%8B%E8%BD%AC%E6%8D%A2)
    - [理解 Comparable](#%E7%90%86%E8%A7%A3-Comparable)
    - [字面量没有类型](#%E5%AD%97%E9%9D%A2%E9%87%8F%E6%B2%A1%E6%9C%89%E7%B1%BB%E5%9E%8B)
    - [通过 Const 给字面量命名](#%E9%80%9A%E8%BF%87-Const-%E7%BB%99%E5%AD%97%E9%9D%A2%E9%87%8F%E5%91%BD%E5%90%8D)
    - [变量遮蔽](#%E5%8F%98%E9%87%8F%E9%81%AE%E8%94%BD)
    - [独特的 if 语句](#%E7%8B%AC%E7%89%B9%E7%9A%84-if-%E8%AF%AD%E5%8F%A5)
    - [For 循环的四种写法](#For-%E5%BE%AA%E7%8E%AF%E7%9A%84%E5%9B%9B%E7%A7%8D%E5%86%99%E6%B3%95)
    - [独特的 switch 语句](#%E7%8B%AC%E7%89%B9%E7%9A%84-switch-%E8%AF%AD%E5%8F%A5)
  - [函数](#%E5%87%BD%E6%95%B0)
    - [参数按值传递](#%E5%8F%82%E6%95%B0%E6%8C%89%E5%80%BC%E4%BC%A0%E9%80%92)
    - [命名参数、可选参数](#%E5%91%BD%E5%90%8D%E5%8F%82%E6%95%B0%E5%8F%AF%E9%80%89%E5%8F%82%E6%95%B0)
    - [变长参数](#%E5%8F%98%E9%95%BF%E5%8F%82%E6%95%B0)
    - [命名返回值就是用变量存储返回值](#%E5%91%BD%E5%90%8D%E8%BF%94%E5%9B%9E%E5%80%BC%E5%B0%B1%E6%98%AF%E7%94%A8%E5%8F%98%E9%87%8F%E5%AD%98%E5%82%A8%E8%BF%94%E5%9B%9E%E5%80%BC)
  - [函数类型](#%E5%87%BD%E6%95%B0%E7%B1%BB%E5%9E%8B)
    - [函数变量/类型](#%E5%87%BD%E6%95%B0%E5%8F%98%E9%87%8F%E7%B1%BB%E5%9E%8B)
    - [匿名函数](#%E5%8C%BF%E5%90%8D%E5%87%BD%E6%95%B0)
    - [返回清理函数](#%E8%BF%94%E5%9B%9E%E6%B8%85%E7%90%86%E5%87%BD%E6%95%B0)
  - [defer](#defer)
    - [defer 通常用来关闭资源](#defer-%E9%80%9A%E5%B8%B8%E7%94%A8%E6%9D%A5%E5%85%B3%E9%97%AD%E8%B5%84%E6%BA%90)
    - [defer 在什么时候执行](#defer-%E5%9C%A8%E4%BB%80%E4%B9%88%E6%97%B6%E5%80%99%E6%89%A7%E8%A1%8C)

## 语言基础

### 创建变量的各种方式

```go
func 变量声明() {
	var a int = 10        // 最完整的写法
	var b = 10            // 自动推断类型 or 使用默认类型
	c := 10               // 省略 var关键字
	var d int             // 使用零值,  不初始化
	g, h := 10, "hello"   // 可以一次性创建多个变量,  但更建议分成两行,  可读性更好
	res, err := xxxFunc() // only use this style when assigning multiple values returned from a function

	// 可以用 := 进行赋值. As long as there is one new variable on the left-hand side of the :=
	m := 10
	m, n := 30, "hello"

	// 所以呢 := 有三种行为,  ①创建新变量  ②对现有变量赋值
	x := 10
	if x > 5 {
		x := 2233      // ③在局部作用域中,  创建同名的新变量
		fmt.Println(x) // 注意两个变量 x 只是名字相同,  但各有各的内容 (尽量避免这样的代码,  易混淆)
	}

    // If you are declaring a variable at package level, 
    // you must use var because := is not legal outside of functions
	var packageLevelVariable int
}
```

### 全都需要显式类型转换

Most languages that have multiple numeric types automatically convert from one to another when needed. This is called `automatic type promotion`, and while it seems very convenient, it turns out that the rules to properly convert one type to another can get complicated and produce unexpected results.

As a language that values clarity of intent and readability, Go doesn’t allow automatic type promotion between variables. You must use a type conversion when variable types do not match. Even different-sized integers and floats must be converted to the same type to interact.

*Since all type conversions in Go are explicit, you cannot treat another Go type as a boolean*. In many languages, a nonzero number or a nonempty string can be interpreted as a boolean true. Just like automatic type promotion, the rules for “truthy” values vary from language to language and can be confusing. Unsurprisingly, Go doesn’t allow truthiness. In fact, no other type can be converted to a bool, implicitly or explicitly. If you want to convert from another data type to boolean, you must use one of the comparison operators (==, !=, >, <, <=, or >=)  

### 理解 Comparable

➤ slice、map、function、channel 是不可比较的

The only thing you can compare a slice with is `nil`. It is a compile-time error to use `==` to see if two slices are identical or `!=` to see if they are different. 注意反射包有个 DeepEqual 方法几乎能比较任意类型,  包括 slice.

Whether or not a struct is comparable depends on the struct’s fields. Structs that are entirely composed of comparable types are comparable; those with slice or map fields are not. 

Unlike in Python or Ruby, in Go there’s no magic method that can be overridden to redefine equality and make == and != work for incomparable structs. You can, of course, write your own function that you use to compare structs.

Just like Go doesn’t allow comparisons between variables of different primitive types, Go doesn’t allow comparisons between variables that represent structs of different types. But Go does allow you to perform a type conversion from one struct type to another if the fields of both structs have the same names, order, and types. 

一个 anonymous struct 变量可以与另一个 struct 变量,  直接进行比较操作、赋值操作, 无需类型转换,  
但前提是 If the fields of both structs have the same names, order, and types. 

### 字面量没有类型

You can’t even add two integer variables together if they are declared to be of different sizes. However, Go lets you use an integer literal in floating point expressions or even assign an integer literal to a floating point variable. <font color='#D05'>This is because literals in Go are untyped; they can interact with any variable that’s compatible with the literal</font>. We’ll see that we can even use literals with user-defined types based on primitive types.

Being untyped only goes so far; you can’t assign a literal string to a variable with a numeric type or a literal number to a string variable, nor can you assign a float literal to an int.

### 通过 Const 给字面量命名

However, `const` in Go is very limited. Constants in Go are a way to give names to literals. They can only hold values that the compiler can figure out at compile time. This means that they can be assigned:

- Numeric literals、true and false、Strings、Runes
- The built-in functions complex, real, imag, len, and cap
- Expressions that consist of operators and the preceding values

Go doesn’t provide a way to specify that a value calculated at runtime is immutable. As we’ll see in the next chapter, there are no immutable arrays, slices , maps, or structs, and there’s no way to declare that a field in a struct is immutable. <font color='#D05'>Constants in Go are a way to give names to literals. There is no way in Go to declare that a variable is immutable</font>.  

➤ Typed and Untyped Constants

Constants can be typed or untyped. An untyped constant works exactly like a literal; it has no type of its own, but does have a default type that is used when no other type can be inferred. A typed constant can only be directly assigned to a variable of that type.

```go
const x = 10              // untyped
const typedX int = 10     // typed
var y byte = x            // 合法,  因为 x 常量没有类型
var z byte = typedX       // 非法,  因为不能把 int 类型赋值给 byte 类型
```

### 变量遮蔽

详情参考 Learning Go 第四章,  要点如下:

- `x := 5` 会遮蔽外部的 x 变量,  这个「 外部 」可能是当前大括号外的 x 变量、或包级别的 x 变量
- `math := "oops"` 会遮蔽 math 包,  甚至允许用 `true := 10` 遮蔽 universe block 中的 true 变量 (但不推荐)

```go
x := 10
{
    // 无法对外部的 x 变量进行赋值,  除非 x 变量也定义在当前块作用域
    // 所以此处创建了 x 和 y 两个新变量
    x, y := 5, 20
}
fmt.Println(x) // 注意值是 10 而不是 5
```

### 独特的 if 语句

The most visible difference between if statements in Go and other languages is that you don’t put parenthesis around the condition. What Go adds is the ability to declare variables that are scoped to the condition and to both the if and else blocks.

![image-20220508161113546](https://static.xianyukang.com/img/image-20220508161113546.png) 

### For 循环的四种写法



```go
// ➤ A complete, C-style for
for i := 0; i < 10; i++ {
	fmt.Println(i)
}

// ➤ A condition-only for, like the while statement found in C
for i < 100 {
	fmt.Println(i)
	i = i * 2
}

// ➤ An infinite for
for {
	fmt.Println("Hello")
}
```

➤ for-range

You can only use a for-range loop to iterate over the built-in compound types and user-defined types that are based on them.

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

➤ The for-range value is a copy

You should be aware that each time the for-range loop iterates over your compound type, it copies the value from the compound type to the value variable. Modifying the value variable will not modify the value in the compound type. 注意不要用 `for _, v := range items { v.id = 123 }` 修改数组中的结构体,  因为 `v` 是一个拷贝.

➤ Labeling Your for Statements

By default, the break and continue keywords apply to the for loop that directly contains them. What if you have nested for loops and you want to exit or skip over an iterator of an outer loop?  

![image-20220509110608855](https://static.xianyukang.com/img/image-20220509110608855.png) 

### 独特的 switch 语句

![image-20220509115714882](https://static.xianyukang.com/img/image-20220509115714882.png) 

In our sample program we are switching on the value of an integer, but that’s not all you can do. You can switch on any type that can be compared with `==`, which includes all of the built-in types except slices, maps, channels, functions, and structs that contain fields of these types.  

在 switch 语句中一般不会写 break,  但是也允许写,  break 的行为是提前退出当前 switch 语句  
如果 for 循环里面嵌套了一个 switch 语句,  可以用上一节讲的 labeled for statement 退出外层循环

➤ Blank Switch

You can write a switch statement that doesn’t specify the value that you’re comparing against. This is called a blank switch. A regular switch only allows you to check a value for equality. A blank switch allows you to use any boolean comparison for each case.

![image-20220509121137866](https://static.xianyukang.com/img/image-20220509121137866.png) 

## 函数

### 参数按值传递

- It means that when you supply a variable for a parameter to a function,  
  Go always makes a copy of the value of the variable.  
  这种复制是浅复制,  但对于不含指针的对象,  浅复制等价于深复制
- 如果用 √ X 表示是否影响外部:
  - 在函数内修改传入的 `int`、`string`、`struct` (X)
  - 在函数内修改传入的 `map`,  增删 key (√)、修改 key 对应的 value (√)
  - 在函数内修改传入的 `slice`,  修改索引对应的值 (√),  用 append 修改切片长度 (X)
  - 另外在函数内修改传入的结构体的 map/slice 字段,  也会影响外部

- 把 string/slice 传给函数的复制开销很小,  只需要复制 SliceHeader 就好,  其中包含 cap、len 字段和一个 data 指针

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

### 命名返回值就是用变量存储返回值

Go also allows you to specify names for your return values. When you supply names to your return values, what you are doing is pre-declaring variables that you use within the function to hold the return values. 

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
```

 [➤ 为什么 defer 配合 named return vlaue 能修改函数返回值?](https://stackoverflow.com/questions/37248898/how-does-defer-and-named-return-value-work)

一些注意点:

1. 返回值变量像普通变量一样,  可能发生遮蔽
2. `return 0, nil` 相当于把 0、nil 赋值给 result、err 变量
3. `return` 后面啥也不跟就叫 blank return.  
   However, most experienced Go developers consider blank returns a bad idea because they make it harder to understand data flow.

## 函数类型

### 函数变量/类型

➤ 声明函数变量: `var f func(a, b int) int`.

Just like in many other languages, functions in Go are values. The type of a function is built out of the keyword `func` and the types of the parameters and return values. This combination is called the signature of the function. Any function that has the exact same number and types of parameters and return values meets the type signature.

➤ 创建函数类型: `type opFuncType func(int,int) int`

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

➤ 闭包就是带状态的函数

Functions declared inside of functions are special; they are closures. This is a computer science word that means that functions declared inside of functions are able to access and modify variables declared in the outer function.  

```go
func main() {
	count := counter()
	count() // 打印 1
	count() // 打印 2
}

func counter() func() {
	val := 0
	return func() {
		val++
		fmt.Println(val)
	}
}
```

➤ 使用匿名函数对切片进行排序

![image-20220509191650192](https://static.xianyukang.com/img/image-20220509191650192.png) 

### 返回清理函数

A common pattern in Go is for a function that allocates a resource to also return a closure that cleans up the resource.
![image-20220510002413763](https://static.xianyukang.com/img/image-20220510002413763.png)

## defer

### defer 通常用来关闭资源

Programs often create temporary resources, like files or network connections, that need to be cleaned up. This cleanup has to happen, no matter how many exit points a function has, or whether a function completed successfully or not. In Go, the cleanup code is attached to the function with the defer keyword.

![image-20220509233838326](https://static.xianyukang.com/img/image-20220509233838326.png) 

### defer 在什么时候执行

defer 的执行时机为: 在 retrun/panic 语句之后、在函数返回之前  
(1) [Is golang defer statement execute before or after return statement?](https://stackoverflow.com/questions/52718143/is-golang-defer-statement-execute-before-or-after-return-statement)  
(2) [go - How does defer and named return value work?](https://stackoverflow.com/questions/37248898/how-does-defer-and-named-return-value-work)

defer 的执行顺序为 LIFO,  最后一个注册的 defer 最先执行  
defer 有时候不执行, 比如 `log.Fatal`、`os.Exit`、终端中按下 `Ctrl+C`,  这些都会导致进程立即退出

➤ 用 defer 修改函数返回值

![image-20220509234632112](https://static.xianyukang.com/img/image-20220509234632112.png)  

