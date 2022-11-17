## Table of Contents
  - [从官方文档开始](#%E4%BB%8E%E5%AE%98%E6%96%B9%E6%96%87%E6%A1%A3%E5%BC%80%E5%A7%8B)
  - [fmt](#fmt)
    - [方便的 %v](#%E6%96%B9%E4%BE%BF%E7%9A%84-v)
    - [Stringer 接口](#Stringer-%E6%8E%A5%E5%8F%A3)
  - [io](#io)
    - [Reader/Writer 概述](#ReaderWriter-%E6%A6%82%E8%BF%B0)
    - [如何使用 Reader 接口](#%E5%A6%82%E4%BD%95%E4%BD%BF%E7%94%A8-Reader-%E6%8E%A5%E5%8F%A3)
    - [从字符串或Gzip文件创建 Reader](#%E4%BB%8E%E5%AD%97%E7%AC%A6%E4%B8%B2%E6%88%96Gzip%E6%96%87%E4%BB%B6%E5%88%9B%E5%BB%BA-Reader)
    - [其他 io 函数和工具](#%E5%85%B6%E4%BB%96-io-%E5%87%BD%E6%95%B0%E5%92%8C%E5%B7%A5%E5%85%B7)
    - [为现有类型添加方法、实现接口](#%E4%B8%BA%E7%8E%B0%E6%9C%89%E7%B1%BB%E5%9E%8B%E6%B7%BB%E5%8A%A0%E6%96%B9%E6%B3%95%E5%AE%9E%E7%8E%B0%E6%8E%A5%E5%8F%A3)
  - [time](#time)
    - [time.Duration](#timeDuration)
    - [time.Time](#timeTime)
    - [Monotonic Time](#Monotonic-Time)
    - [Timers and Timeouts](#Timers-and-Timeouts)
  - [json](#json)
    - [Use Struct Tags to Add Metadata](#Use-Struct-Tags-to-Add-Metadata)
    - [Unmarshaling and Marshaling](#Unmarshaling-and-Marshaling)
    - [JSON, Readers, and Writers](#JSON-Readers-and-Writers)
    - [Encoding and Decoding JSON Streams](#Encoding-and-Decoding-JSON-Streams)
    - [Custom JSON Parsing](#Custom-JSON-Parsing)
    - [json.RawMessage 有什么用](#jsonRawMessage-%E6%9C%89%E4%BB%80%E4%B9%88%E7%94%A8)
  - [net/http](#nethttp)
    - [The Client](#The-Client)
    - [The Server](#The-Server)
  - [os](#os)
    - [os/exec](#osexec)
  - [sort](#sort)
    - [例子](#%E4%BE%8B%E5%AD%90)

## 从官方文档开始

We can’t cover all of the standard library packages, and luckily, we don’t have to, as there are many excellent sources of information on the standard library, [starting with the documentation](https://pkg.go.dev/std). Instead, we’ll focus on several of the most important packages and how their design and use demonstrate the principles of idiomatic Go. Some packages (errors, sync, context, testing, reflect, and unsafe) are covered in their own chapters. In this chapter, we’ll look at Go’s built-in support for I/O, time, JSON, and HTTP.

Moreover, many of the packages contain working, self-contained executable examples you can run directly from the [golang.org](https://golang.org/) web site, such as [this one](https://go.dev/pkg/strings/#example_Map) (if necessary, click on the word "Example" to open it up). If you have a question about how to approach a problem or how something might be implemented, the documentation, code and examples in the library can provide answers, ideas and background.

## fmt

### 方便的 %v

If you just want the default conversion, such as decimal for integers, you can use the catchall format `%v` (for “value”); the result is exactly what `Print` and `Println` would produce. Moreover, that format can print *any* value, even arrays, slices, structs, and maps.

```go
fmt.Printf("%v\n", timeZone)   // 等价于 fmt.Println(timeZone)
var c = Character{Name: "Cloud", From: "FF7", Age: 21}
fmt.Printf("%+v\n", c)         // 打印结构体时带上字段名
fmt.Printf("%#v\n", c)         // prints the value in full Go syntax.
```

### Stringer 接口

如果 `User` 类型有个 `time.Time` 字段,  那么直接用 `fmt` 打印会很丑,  
可以实现 `Stringer` 接口,  然后 `fmt.Println` 就会调用这个接口把 `User` 转成 `string`

```go
type Character struct {
	Name string
	From string
	Age  int
}
func (c Character) String() string {
	return fmt.Sprintf("(%q, %q, %v)", c.Name, c.From, c.Age)
    // return fmt.Sprintf("%v", c)  // 小心无限递归,  fmt 处理 %s/%v/%q 时又会调用 String()
}
func main() {
	var c = Character{Name: "Cloud", From: "FF7", Age: 21}
	fmt.Println(c)
}
```

## io

### Reader/Writer 概述

For a program to be useful, it needs to read in and write out data. The heart of Go’s input/output philosophy can be found in the io package. In particular, two interfaces defined in this package are probably the second and third most-used interfaces in Go: `io.Reader` and `io.Writer`. (What’s number one? That’d be `error`). Both io.Reader and io.Writer define a single method:

![image-20220627113707965](https://static.xianyukang.com/img/image-20220627113707965.png) 

The `Write` method on the `io.Writer` interface takes in a slice of bytes, which are written to the interface implementation. It returns the number of bytes written and an error if something went wrong. The `Read` method on `io.Reader` is more interesting. Rather than return data through a return parameter, a slice input parameter is passed into the implementation and modified. Up to `len(p)` bytes will be written into the slice. The method returns the number of bytes written.

### 如何使用 Reader 接口

This might seem a little strange. You might expect this:  `Read() ([]byte, error)`. There’s a very good reason why `io.Reader` is defined the way it is. Let’s write a function that’s representative of how to work with an io.Reader to understand:  

![image-20220628012603433](https://static.xianyukang.com/img/image-20220628012603433.png) 

(1) First, we create our buffer once and reuse it on every call to `r.Read`. This allows us to use a single memory allocation to read from a potentially large data source. If the `Read` method were written to return a `[]byte`, it would require a new allocation on every single call. Each allocation would end up on the heap, which would make quite a lot of work for the garbage collector. By passing in a slice to io.Reader, memory allocation is under the control of the developer.

(2) Second, we use the `n` value returned from `r.Read` to know how many bytes were written to the buffer and iterate over a subslice of our buf slice, processing the data that was read.  

(3) Finally, we know that we’re done reading from `r` when the error returned from `r.Read` is `io.EOF`. This error is a bit odd, in that it isn’t really an error. It indicates that there’s nothing left to read from the `io.Reader`. When `io.EOF` is returned, we are finished processing and return our result.  

There is one unusual thing about the `Read` method in `io.Reader`. In most cases when a function or method has an error return value, we check the error before we try to process the nonerror return values. We do the opposite for `Read` *(指没有检查 err 就使用了返回值)* because there might have been bytes returned before an error was triggered by the end of the data stream or by an unexpected condition.

### 从字符串或Gzip文件创建 Reader

Because `io.Reader` and `io.Writer` are such simple interfaces, they can be implemented many different ways. We can create an `io.Reader` from a string using the `strings.NewReader` function:

![image-20220628014225941](https://static.xianyukang.com/img/image-20220628014225941.png) 

➤ 装饰器模式指 decorator 拦截了函数调用、不改变对外接口、也不改变函数本身,  但多了一层逻辑.

Implementations of `io.Reader` and `io.Writer` are often chained together in a decorator pattern. Because countLetters depends on an `io.Reader`, we can use the exact same `countLetters` function to count English letters in a gzip-compressed file. First we write a function that, when given a filename, returns a `*gzip.Reader`:  

![image-20220628145210796](https://static.xianyukang.com/img/image-20220628145210796.png) 

*This function demonstrates the way to properly wrap types that implement io.Reader*. We create an `*os.File` (which meets the `io.Reader` interface), and after making sure it’s valid, we pass it to the `gzip.NewReader` function, which returns a `*gzip.Reader` instance. If it is valid, we return the `*gzip.Reader` and a closer closure that properly cleans up our resources when it is invoked. Since `*gzip.Reader` implements `io.Reader`, we can use it with `countLetters` just like we used the `*strings.Reader` previously:  

![image-20220628145855156](https://static.xianyukang.com/img/image-20220628145855156.png) 

### 其他 io 函数和工具

There’s a standard function in the io package for copying from an io.Reader to an io.Writer, `io.Copy`. There are other standard functions for adding new functionality to existing io.Reader and io.Writer instances. These include:  

- `io.MultiReader` returns an io.Reader that reads from multiple io.Reader instances, one after another.  

- `io.LimitReader` returns an io.Reader that only reads up to a specified number of bytes from the supplied io.Reader.  
- `io.MultiWriter` returns an io.Writer that writes to multiple io.Writer instances at the same time.  

The `io.Seeker` interface is used for random access to a resource. The `io.Closer` interface is implemented by types like os.File that need to do cleanup when reading or writing is complete. Usually, Close is called via a defer:

![image-20220629030746327](https://static.xianyukang.com/img/image-20220629030746327.png) 

➤ 不要在循环里用 defer、推荐手动关闭、因为 defer 与块作用域无关、只在函数结束后运行

*If you are opening the resource in a loop, do not use defer*, as it will not run until the function exits. Instead, you should call Close before the end of the loop iteration. If there are errors that can lead to an exit, you must call `Close` there, too.  

➤ 推荐用 io.ReadCloser 这样的整合接口而不是用 os.File 作为参数,  因为前者更通用、并且意图更清晰

The io package defines interfaces that combine these four interfaces in various ways. They include io.ReadCloser, io.ReadSeeker, io.ReadWriteCloser, io.ReadWrite Seeker, io.ReadWriter, io.WriteCloser, and io.WriteSeeker. Use these interfaces to specify what your functions expect to do with the data. For example, rather than just using an `os.File` as a parameter, use the interfaces to specify exactly what your function will do with the parameter. Not only does it make your functions more general purpose, it also makes your intent clearer.

➤ 小文件可以用 ioutil、大文件应该用 bufio

The `ioutil` package provides some simple utilities for things like reading entire `io.Reader` implementations into byte slices at once, reading and writing files, and  working with temporary files. The `ioutil.ReadAll`, `ioutil.ReadFile`, and `ioutil.WriteFile` functions are fine for small data sources, *but it’s better to use the Reader, Writer, and Scanner in the bufio package to work with larger data sources*.

### 为现有类型添加方法、实现接口

One of the more clever functions in ioutil demonstrates a pattern for adding a method to a Go type. If you have a type that implements io.Reader but not io.Closer (such as strings.Reader) and need to pass it to a function that expects an io.ReadCloser, pass your `io.Reader` into `ioutil.NopCloser` and get back a type that implements `io.ReadCloser`. If you look at the implementation, it’s very simple:  

![image-20220629032621291](https://static.xianyukang.com/img/image-20220629032621291.png) 

Any time you need to add additional methods to a type so that it can meet an interface, use this embedded type pattern.

## time

### time.Duration

There are two main types used to represent time, `time.Duration` and `time.Time`. A period of time is represented with a time.Duration, a type based on an `int64`. The smallest amount of time that Go can represent is one nanosecond, but the time package defines constants of type time.Duration to represent a nanosecond, microsecond, millisecond, second, minute, and hour. 

For example, you represent a duration of 2 hours and 30 minutes with: `d := 2 * time.Hour + 30 * time.Minute`. `d` is of type time.Duration. These constants make the use of a time.Duration both readable and type-safe. They demonstrate a good use of a typed constant.  

Go defines a sensible string format, a series of numbers, that can be parsed into a time.Duration with the `time.ParseDuration` function. As described in the standard library documentation:  

> A duration string is a possibly signed sequence of decimal numbers, each with optional fraction and a unit suffix, such as “300ms”, “-1.5h” or “2h45m”. Valid time units are “ns”, “us” (or “µs”), “ms”, “s”, “m”, “h”.  

There are several methods defined on time.Duration. It meets the `fmt.Stringer` interface and returns a formatted duration string via the `String` method. It also has methods to get the value as a number of hours, minutes, seconds, milliseconds, microseconds, or nanoseconds. The Truncate and Round methods truncate or round a time.Duration to the units of the specified time.Duration.  

![image-20220629125746379](https://static.xianyukang.com/img/image-20220629125746379.png) 

### time.Time

An moment of time is represented with the `time.Time` type, complete with a time zone. You acquire a reference to the current time with the function `time.Now`. This returns a `time.Time` instance set to the current local time.  

> The fact that a `time.Time` instance contains a time zone means that you should not use `==` to check if two time.Time instances refer to the same moment in time. Instead, use the `Equal` method, which corrects for time zone.  

The `time.Parse` function converts from a string to a time.Time, while the `Format` method converts a time.Time to a string. `time.Parse` parses a formatted string and returns the time value it represents. 

![image-20220630034152962](https://static.xianyukang.com/img/image-20220630034152962.png) 

➤ 提取时间分量、比较时间

Just as there are methods on time.Duration to extract portions of it, there are methods defined on time.Time to do the same, including Day, Month, Year, Hour, Minute, Second, Weekday, `Clock` (which returns the time portion of a time.Time as separate hour, minute, and second int values), and `Date` (which returns the year, month, and day as separate int values). You can compare one time.Time instance against another with the `After`, `Before`, and `Equal` methods.  

➤ 时间的加减计算

The `Sub` method returns a `time.Duration` that represents the elapsed time between two `time.Time` instances, while the `Add` method returns a `time.Time` that is `time.Duration` later, and the `AddDate` method returns a new `time.Time` instance that’s incremented by the specified number of years, months, and days. As with time.Duration, there are `Truncate` and `Round` methods defined as well. All of these methods are defined on a value receiver, so they do not modify the time.Time instance.  

![image-20220701131746995](https://static.xianyukang.com/img/image-20220701131746995.png) 

### Monotonic Time

Most operating systems keep track of two different sorts of time: the wall clock, which corresponds to the current time, and the *monotonic clock which simply counts up from the time the computer was booted*. The reason for tracking two different clocks is that the wall clock doesn’t uniformly increase. NTP (Network Time Protocol) updates can make the wall clock move unexpectedly forward or backward. This can cause problems when setting a timer or finding the amount of time that’s elapsed.  

To address this potential problem, Go uses monotonic time to track elapsed time whenever a timer is set or a time.Time instance is created with `time.Now`. This support is invisible; timers use it automatically. The `Sub` method uses the montonic clock to calculate the `time.Duration` if both of the `time.Time` instances have it set. If they don’t (because one or both of the instances was not created with `time.Now`), the `Sub` method uses the time specified in the instances to calculate the `time.Duration` instead.  

### Timers and Timeouts

The time package includes functions that return channels that output values after a specified time. The `time.After` function returns a channel that outputs once, while the channel returned by `time.Tick` returns a new value every time the specified time.Duration elapses. You can also trigger a single function to run after a specified time.Duration with the `time.AfterFunc` function. Don’t use `time.Tick` outside of trivial programs, because the underlying `time.Ticker` cannot be shut down (and therefore cannot be garbage collected). Use the `time.NewTicker` function instead, which returns a `*time.Ticker` that has the channel to listen to, as well as methods to reset and stop the ticker.

## json

Go’s standard library includes support for converting Go data types to and from JSON. The word marshaling means converting from a Go data type to an encoding, and unmarshaling means converting to a Go data type.

### Use Struct Tags to Add Metadata

We specify the rules for processing our JSON with struct tags, strings that are written after the fields in a struct. Even though struct tags are strings marked with backticks, they cannot extend past a single line. Struct tags are composed of one or more tag-value pairs, written as `tagName:"tagValue"` and separated by spaces. Also, note that all of these fields are exported. Like any other package, the code in the `encoding/json` package cannot access an unexported field on a struct in another package.  

➤ 推荐用 json 标签显式指定字段名

For JSON processing, we use the tag name `json` to specify the name of the JSON field that should be associated with the struct field. If no json tag is provided, the default behavior is to assume that the name of the JSON object field matches the name of the Go struct field. Despite this default behavior, it’s best to use the struct tag to specify the name of the field explicitly, even if the field names are identical.  

When unmarshaling from JSON into a struct field with no json tag, the name match is case-insensitive. When marshaling a struct field with no json tag back to JSON , the JSON field will always have an uppercase first letter, because the field is exported.

➤ 如何忽略某字段、以及 empty 的定义

If a field should be ignored when marshaling or unmarshaling, use a dash `-` for the name. If the field should be left out of the output when it is empty, add `,omitempty` after the name. Unfortunately, the definition of “empty” doesn’t exactly align with the zero value, as you might expect. The zero value of a struct doesn’t count as empty, but a zero-length slice or map does.

![image-20220702141843132](https://static.xianyukang.com/img/image-20220702141843132.png) 

Struct tags allow you to use metadata to control how your program behaves. Other languages, most notably Java, encourage developers to place annotations on various program elements to describe how they should be processed, without explicitly specifying what is going to do the processing. While declarative programming allows for more concise programs, automatic processing of metadata makes it difficult to understand how a program behaves. Anyone who has worked on a large Java project with annotations has had a moment of panic when something goes wrong and they don’t understand which code is processing a particular annotation and what changes it made. *Go favors explicit code over short code*.

### Unmarshaling and Marshaling

The `Unmarshal` function in the encoding/json package is used to convert a slice of bytes into a struct. If we have a string named data, this is the code to convert data to a struct of type Order:

![image-20220702143744196](https://static.xianyukang.com/img/image-20220702143744196.png) 

The `json.Unmarshal` function populates data into an input parameter, just like the implementations of the `io.Reader` interface. There are two reasons for this. First, just like `io.Reader` implementations, this allows for efficient reuse of the same struct over and over, giving you control over memory usage. Second, there’s simply no other way to do it. Because Go doesn’t currently have generics, there’s no way to specify what type should be instantiated to store the bytes being read. Even when Go adopts generics, the memory usage advantages will remain.

We use the `Marshal` function in the encoding/json package to write an Order instance back as JSON, stored in a slice of bytes: `out, err := json.Marshal(o)`.

This leads to the question: how are you able to evaluate struct tags? You might also be wondering how `json.Marshal` and `json.Unmarshal` are able to read and write a struct of any type. After all, every other method that we’ve written has only worked with types that were known when the program was compiled (even the types listed in a type switch are enumerated ahead of time). The answer to both questions is reflection.

### JSON, Readers, and Writers

The `json.Marshal` and `json.Unmarshal` functions work on slices of bytes. As we just saw, most data sources and sinks in Go implement the `io.Reader` and `io.Writer` interfaces. While you could use `ioutil.ReadAll` to copy the entire contents of an `io.Reader` into a byte slice so it can be read by `json.Unmarshal`, *this is inefficient*. Similarly, we could write to an in-memory byte slice buffer using `json.Marshal` and then write that byte slice to the network or disk, but it’d be better if we could write to an `io.Writer` directly.

The encoding/json package includes two types that allow us to handle these situations. The `json.Decoder` and `json.Encoder` types read from and write to anything that meets the `io.Reader` and `io.Writer` interfaces, respectively.  

The `os.File` type implements both the `io.Reader` and `io.Writer` interfaces, so we can use it to demonstrate json.Decoder and json.Encoder. First, we write toFile to a temp file by passing the temp file to json.NewEncoder, which returns a json.Encoder for the temp file. We then pass toFile to the Encode method:  

![image-20220702194202138](https://static.xianyukang.com/img/image-20220702194202138.png) 

Once toFile is written, we can read the JSON back in by passing a reference to the temp file to json.NewDecoder and then calling the Decode method on the returned json.Decoder with a variable of type Person:  

![image-20220702195121179](https://static.xianyukang.com/img/image-20220702195121179.png) 

### Encoding and Decoding JSON Streams

![image-20220702195641773](https://static.xianyukang.com/img/image-20220702195641773.png) 

### Custom JSON Parsing

While the default functionality is often sufficient, there are times you need to override it. While time.Time supports JSON fields in RFC 339 format out of the box, you might have to deal with other time formats. We can handle this by creating a new type that implements two interfaces, `json.Marshaler` and `json.Unmarshaler`:  

![image-20220702202319546](https://static.xianyukang.com/img/image-20220702202319546.png) 

We embedded a time.Time instance into a new struct called `RFC822ZTime` so that we still have access to the other methods on `time.Time`. As we discussed in “Pointer Receivers and Value Receivers” on page 131, the method that reads the time value is declared on a value receiver, *while the method that modifies the time value is declared on a pointer receiver*.

➤ 分离序列化逻辑与业务逻辑

为了让自定义的时间格式生效, Order 结构体中需要用 RFC822ZTime 类型,  
导致格式化逻辑影响了数据的存储结构, 所以这种方式有一点耦合 (Order 中用 time.Time 存储时间更合适). 

To limit the amount of code that cares about what your JSON looks like, define two different structs. Use one for converting to and from JSON and the other for data processing. Read in JSON to your JSON-aware type, and then copy it to the other. When you want to write out JSON, do the reverse. This does create some duplication, but it keeps your business logic from depending on wire protocols.  

While JSON is probably the most commonly used encoder in the standard library, Go ships with others, including XML and Base64. If you have a data format that you want to encode and you can’t find support for it in the standard library or a thirdparty module, you can write one yourself. We’ll learn how to implement our own encoder in “Use Reflection to Write a Data Marshaler” on page 307.  

### json.RawMessage 有什么用

➤ 若 status 字段可能为 string 或 number,  可以用 json.RawMessage 延迟解析

```go
func main() {
	records := [][]byte{
		[]byte(`{"status": 200, "tag":"one"}`),
		[]byte(`{"status":"ok", "tag":"two"}`),
	}

	for idx, record := range records {
		var result struct {
			StatusCode uint64
			StatusName string
			Status     json.RawMessage `json:"status"`
			Tag        string          `json:"tag"`
		}

		// 注意 Status 字段的类型为 json.RawMessage
		// 所谓 json.RawMessage 是个 byte slice,  表示先把数据存下来,  等下再进一步解析
		if err := json.NewDecoder(bytes.NewReader(record)).Decode(&result); err != nil {
			fmt.Println("error:", err)
			return
		}

		// records 中的 status 字段可能是字符串,  也可能是数字
		var stringStatus string
		if err := json.Unmarshal(result.Status, &stringStatus); err == nil {
			result.StatusName = stringStatus
		}

		var numberStatus uint64
		if err := json.Unmarshal(result.Status, &numberStatus); err == nil {
			result.StatusCode = numberStatus
		}

		fmt.Printf("[%v] result => %+v\n", idx, result)
	}
}
```

## net/http

Every language ships with a standard library, but the expectations of what a standard library should include have changed over time. As a language launched in the 2010s, Go’s standard library includes something that *other language distributions had considered the responsibility of a third party*: a production quality HTTP/2 client and server.

### The Client

(1) The net/http package defines a Client type to make HTTP requests and receive HTTP responses. A default client instance (cleverly named DefaultClient) is found in the net/http package, but you should avoid using it in production applications, because it defaults to having no timeout. Instead, instantiate your own. You only need to create a single http.Client for your entire program, as it properly handles multiple simultaneous requests across goroutines:  `client := &http.Client{Timeout: 30 * time.Second}`

(2) When you want to make a request, you create a new `*http.Request` instance with the `http.NewRequestWithContext` function, passing it a context, the method, and URL that you are connecting to. If you are making a PUT, POST, or PATCH request, specify the body of the request with the last parameter as an `io.Reader`. If there is no body, use `nil`.

(3) Once you have an `*http.Request` instance, you can set any headers via the `Headers` field of the instance. Call the `Do` method on the `http.Client` with your http.Request and the result is returned in an http.Response.

(4) The response has several fields with information on the request. The numeric code of the response status is in the `StatusCode` field, the text of the response code is in the `Status` field, the response headers are in the `Header` field, and any returned content is in a `Body` field of type `io.ReadCloser`. This allows us to use it with `json.Decoder` to process REST API responses.

![image-20220702214421424](https://static.xianyukang.com/img/image-20220702214421424.png) 

There are functions in the net/http package to make GET, HEAD, and POST calls. Avoid using these functions because they use the default client, which means they don’t set a request timeout.  

### The Server

The http.Server is responsible for listening for HTTP requests. It is a performant HTTP/2 server that supports TLS. A request to a server is handled by an implementation of the http.Handler interface that’s assigned to the Handler field (用 server 里面的 Handler 字段处理请求).

![image-20220705134700753](https://static.xianyukang.com/img/image-20220705134700753.png) 

*These methods must be called in a specific order*. 

1. First, call `Header` to get an instance of http.Header and set any response headers you need.  
   If you don’t need to set any headers, you don’t need to call it. 

2. Next, call `WriteHeader` with the HTTP status code for your response.  
   If you are sending a response that has a 200 status code, you can skip WriteHeader. 

3. Finally, call the `Write` method to set the body for the response.  

![image-20220705144053741](https://static.xianyukang.com/img/image-20220705144053741.png) 

The `Addr` field specifies the host and port the server listens on. If you don’t specify them, your server defaults to listening on all hosts on the standard HTTP port, 80. You specify timeouts for the server’s reads, writes, and idles using time.Duration values. Be sure to set these to properly handle malicious or broken HTTP clients, *as the default behavior is to not time out at all*. Finally, you specify the http.Handler for your server with the `Handler` field.  

A server that only handles a single request isn’t terribly useful, so the Go standard library includes a `request router`, *http.ServeMux. You create an instance with the http.NewServeMux function. It meets the http.Handler interface, so it can be assigned to the Handler field in http.Server. It also includes two methods that allow it to dispatch requests. The first method is simply called `Handle` and takes in two parameters, a path and an http.Handler. 

![image-20220705165122482](https://static.xianyukang.com/img/image-20220705165122482.png) 

> There are package-level functions, http.Handle, http.HandleFunc, http.ListenAndServe, and http.ListenAndServeTLS that work with a package-level instance of the *http.ServeMux called http.DefaultServeMux. Don’t use them outside of trivial test programs. Keep your application under control by avoiding shared state.

## os

### os/exec

[➤ 官方文档](https://pkg.go.dev/os/exec)

Package exec runs external commands. It wraps os.StartProcess to make it easier to remap stdin and stdout, connect I/O with pipes, and do other adjustments.

Unlike the "system" library call from C and other languages, <font color='#D05'>the os/exec package intentionally does not invoke the system shell</font> and does not expand any glob patterns or handle other expansions, pipelines, or redirections typically done by shells. The package behaves more like C's "exec" family of functions. To expand glob patterns, either call the shell directly, taking care to escape any dangerous input, or use the path/filepath package's Glob function. To expand environment variables, use package os's ExpandEnv.

## sort

### 例子

➤ [参考例子](https://go.dev/doc/effective_go#:~:text=bytes.Buffer.-,Interfaces%20and%20other%20types,-Interfaces)

1. 想让自定义类型支持排序,  可以实现 `sort.Interface`
2. `sort.Sort` 和 `sort.Stable` 分别是不稳定排序、和稳定排序
3. 如果实现 `Less` 方法时需要比较两个浮点数,  要考虑 `NaN`,  参考 `sort.Float64Slice` 的 `Less` 实现

➤ 有现成函数对 `[]int`、`[]string`、`[]float64` 排序

```go
type Sequence []int

func main() {
	var s = Sequence{3, 2, 1}
	fmt.Println(s)
	sort.Ints(s)            // It's an idiom in Go programs to convert the type of
	sort.IntSlice(s).Sort() // an expression to access a different set of methods.
	fmt.Println(s)
}

func (s Sequence) Copy() Sequence {
    copy := make(Sequence, 0, len(s))
    return append(copy, s...)
}

func (s Sequence) String() string {
    // Now, instead of having Sequence implement multiple interfaces (sorting and printing), 
    // we're using the ability of a data item to be converted to multiple types (Sequence, sort.IntSlice and []int), 
    // each of which does some part of the job.
    s = s.Copy()                 // 转成 Sequence 利用它的 Copy f
    sort.IntSlice(s).Sort()      // 转成 sort.IntSlice 利用它的排序方法
    return fmt.Sprint([]int(s))  // 转成 []int 利用 fmt.Sprint 对 []int 的支持
}
```

➤ The other way is to use `sort.Slice` with a custom `Less` function.

```go
sort.Slice(people, func(i, j int) bool {
	return people[i].Age < people[j].Age
})
```

➤ 若需要支持多种排序方式, [参考 sort 包的三个例子](https://pkg.go.dev/sort#example-package-SortKeys)

```go
// Sort the planets by the various criteria.
By(name).Sort(planets)
By(mass).Sort(planets)

OrderedBy(user).Sort(changes)
OrderedBy(user, increasingLines).Sort(changes)

sort.Sort(ByWeight{s})
sort.Sort(ByName{s})
```

