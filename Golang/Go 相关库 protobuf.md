## Table of Contents
  - [Protobuf](#Protobuf)
    - [什么是 Protobuf](#%E4%BB%80%E4%B9%88%E6%98%AF-Protobuf)
  - [Protobuf 语法入门](#Protobuf-%E8%AF%AD%E6%B3%95%E5%85%A5%E9%97%A8)
    - [protoc 命令行选项](#protoc-%E5%91%BD%E4%BB%A4%E8%A1%8C%E9%80%89%E9%A1%B9)
    - [简单了解 Protobuf 编码](#%E7%AE%80%E5%8D%95%E4%BA%86%E8%A7%A3-Protobuf-%E7%BC%96%E7%A0%81)
  - [在 Golang 中使用 Protobuf](#%E5%9C%A8-Golang-%E4%B8%AD%E4%BD%BF%E7%94%A8-Protobuf)
    - [入门例子](#%E5%85%A5%E9%97%A8%E4%BE%8B%E5%AD%90)
    - [go_package 的含义](#gopackage-%E7%9A%84%E5%90%AB%E4%B9%89)
    - [可选使用 JSON 序列化](#%E5%8F%AF%E9%80%89%E4%BD%BF%E7%94%A8-JSON-%E5%BA%8F%E5%88%97%E5%8C%96)
    - [一些生成 Go 代码的规则](#%E4%B8%80%E4%BA%9B%E7%94%9F%E6%88%90-Go-%E4%BB%A3%E7%A0%81%E7%9A%84%E8%A7%84%E5%88%99)
  - [标准库的 RPC](#%E6%A0%87%E5%87%86%E5%BA%93%E7%9A%84-RPC)
    - [RPC 是什么](#RPC-%E6%98%AF%E4%BB%80%E4%B9%88)
    - [RPC 入门](#RPC-%E5%85%A5%E9%97%A8)
    - [重构上述例子](#%E9%87%8D%E6%9E%84%E4%B8%8A%E8%BF%B0%E4%BE%8B%E5%AD%90)

## Protobuf

### 什么是 Protobuf

#### 概述

Protocol buffers are Google's language-neutral mechanism for serializing structured data – think XML/JSON, but smaller, faster, and simpler(?). Protocol buffers are the most commonly-used data format at Google. They are used extensively in inter-server communications as well as for archival storage of data on disk.

Protocol buffers are ideal for any situation in which you need to serialize structured, record-like, typed data in a language-neutral, platform-neutral, extensible manner. They are most often used for defining communications protocols (together with gRPC) and for data storage.

它包含四个部分:

- A definition language (`.proto` files)
- The code that the proto compiler generates to interface with data
- Language-specific runtime libraries
- The serialization format for data that is written to a file (or sent across a network connection)

#### 使用流程如下

(1) 创建 .proto 文件,  在其中定义一个 Message,  用来描述数据中有哪些字段  
(2) 用编译工具从 .proto 文件生成 Go 结构体或 Java 类, 我们用生成的类型来读写数据字段  
(3) 最后在 Go/Java 中安装 Protobuf 库,  用来序列化或反序列化上面生成的类型

![img](https://minkj1992.github.io/images/protocol-buffers-concepts.png) 

#### JSON 也行为何 Protobuf

JSON 是文本格式、即便数字 123 用一个字节就能表示,  但序列化成 JSON 要占用 [0x31, 0x32, 0x33] 共三个字节  
JSON 的优点是简单、通用性好、易于查看和修改,  但如果用于服务间通讯,  JSON 有两个缺点  
①编解码速度不是最快 ②编码结果更臃肿、消耗更多网络带宽  

Protobuf 是一种二进制格式,  它表示数字 123 只需一个字节  
它不会像 JSON 一样存储字段名,  所以编/解码速度更快、编码结果更小

```javascript
// JSON 字符串, 占用 23 个字节
{"name":"ivy","age":24}

// 这是上述 JSON 对应的 Protobuf,  只占 10 个字节,  网络传输大小减少了一倍
// 缺点是不够直观,  除非你有这串字节的 .proto 定义，否则无法解析 Protobuf 的内容
[10 6 69 108 108 105 111 116 16 24]
```

## Protobuf 语法入门

#### (0) [详细的语法和示例参考官方文档](https://developers.google.com/protocol-buffers/docs/proto3)

```protobuf
syntax = "proto3";

message SearchRequest {
  string query = 1;
  int32 page_number = 2;
  int32 result_per_page = 3;
}
```

#### (1) Specifying Field Types

Protocol buffers support the usual primitive data types, such as integers, booleans, and floats. For the full list, see [Scalar Value Types](https://developers.google.com/protocol-buffers/docs/proto#scalar). However, you can also specify composite types for your fields, including [enumerations](https://developers.google.com/protocol-buffers/docs/proto3#enum) and other message types.

A field can also be of:

- A `message` type, so that you can nest parts of the definition, such as for repeating sets of data.
- An `enum` type, so you can specify a set of values to choose from.
- A `oneof` type, which you can use when a message has many optional fields and at most one field will be set at the same time.
- A `map` type, to add key-value pairs to your definition.

In addition to the simple and composite value types, several common types are published.

- [`Duration`](https://github.com/protocolbuffers/protobuf/blob/master/src/google/protobuf/duration.proto) is a signed, fixed-length span of time, such as 42s.
- [`Timestamp`](https://github.com/protocolbuffers/protobuf/blob/master/src/google/protobuf/timestamp.proto) is a point in time independent of any time zone or calendar, such as 2017-01-15T01:30:15.01Z.
- [`Date`](https://github.com/googleapis/googleapis/blob/master/google/type/date.proto) is a whole calendar date, such as 2025-09-19.

#### (2) Assigning Field Numbers

As you can see, each field in the message definition has a **unique number**. These field numbers are used to identify your fields in the [message binary format](https://developers.google.com/protocol-buffers/docs/encoding), and should not be changed once your message type is in use. 

Field numbers cannot be repurposed or reused. If you delete a field, you should reserve its field number to prevent someone from accidentally reusing the number.

Note that field numbers in the range 1 through 15 take one byte to encode, including the field number and the field's type. Field numbers in the range 16 through 2047 take two bytes. So you should reserve the numbers 1 through 15 for very frequently occurring message elements.

#### (3) Specifying Field Rules

Message fields can be one of the following:

- `singular`: a well-formed message can have zero or one of this field (but not more than one). When using proto3 syntax, this is the default field rule when no other field rules are specified for a given field. You cannot determine whether it was parsed from the wire bytes. It will be serialized to the wire unless it is the default value.
- `optional`: the same as `singular`, except that you can check to see if the value was explicitly set.
- `repeated`: If a field is `repeated`, the field may be repeated any number of times (including zero). The order of the repeated values will be preserved in the protocol buffer. Think of repeated fields as dynamically sized arrays.

#### Reserved Fields

If you [update](https://developers.google.com/protocol-buffers/docs/proto3#updating) a message type by entirely removing a field, or commenting it out, future users can reuse the field number when making their own updates to the type. This can cause severe issues if they later load old versions of the same `.proto`, including data corruption, privacy bugs, and so on. One way to make sure this doesn't happen is to specify that the field numbers (or names, which can also cause issues for JSON serialization) of your deleted fields are `reserved`. The protocol buffer compiler will complain if any future users try to use these field identifiers.

```protobuf
message Foo {
  reserved 2, 15, 9 to 11;
  reserved "foo", "bar";
}
```

#### Default Values

If a field value isn't set, a [default value](https://developers.google.com/protocol-buffers/docs/proto3#default) is used: zero for numeric types, the empty string for strings, false for bools. For embedded messages, the default value is always the "default instance" of the message, which has none of its fields set. Calling the accessor to get the value of a field which has not been explicitly set always returns that field's default value.

- For strings, the default value is the empty string.
- For bytes, the default value is empty bytes.
- For bools, the default value is false.
- For numeric types, the default value is zero.
- For [enums](https://developers.google.com/protocol-buffers/docs/proto3#enum), the default value is the **first defined enum value**, which must be 0.
- For message fields, the field is not set. Its exact value is language-dependent.
- The default value for repeated fields is empty (generally an empty list in the appropriate language).

#### ➤ 如果想默认启用 `XXX` 那么字段名应该叫 `bool disableXXX;`

Note that for scalar message fields, once a message is parsed there's no way of telling whether a field was explicitly set to the default value (for example whether a boolean was set to `false`) or just not set at all: you should bear this in mind when defining your message types.

#### Enumerations

When you're defining a message type, you might want one of its fields to only have one of a pre-defined list of values. For example, let's say you want to add a `corpus` field for each `SearchRequest`, where the corpus can be `UNIVERSAL`, `WEB`, `IMAGES`, `VIDEO`. You can do this very simply by adding an `enum` to your message definition with a constant for each possible value.

```protobuf
enum Corpus {
  CORPUS_UNSPECIFIED = 0;
  CORPUS_UNIVERSAL = 1;
  CORPUS_WEB = 2;
  CORPUS_IMAGES = 3;
  CORPUS_VIDEO = 4;
}
message SearchRequest {
  string query = 1;
  int32 page_number = 2;
  int32 result_per_page = 3;
  Corpus corpus = 4;
}
```

As you can see, the `Corpus` enum's first constant maps to zero:  
every enum definition **must** contain a constant that maps to zero as its first element.

If you [update](https://developers.google.com/protocol-buffers/docs/proto3#updating) an enum type by entirely removing an enum entry, or commenting it out, future users can reuse the numeric value when making their own updates to the type. This can cause severe issues if they later load old versions of the same `.proto`, including data corruption, privacy bugs, and so on.

```protobuf
enum Foo {
  reserved 2, 15, 9 to 11, 40 to max;
  reserved "FOO", "BAR";
}
```

#### Using Other Message Types

You can use other message types as field types. For example, let's say you wanted to include `Result` messages in each `SearchResponse` message – to do this, you can define a `Result` message type in the same `.proto` and then specify a field of type `Result` in `SearchResponse`:

```protobuf
message SearchResponse {
  repeated Result results = 1;
}

message Result {
  string url = 1;
  string title = 2;
  repeated string snippets = 3;
}
```

You can use definitions from other `.proto` files by *importing* them. To import another `.proto`'s definitions, you add an import statement to the top of your file:

```protobuf
import "myproject/other_protos.proto";
```

By default, you can use definitions only from directly imported `.proto` files. However, sometimes you may need to move a `.proto` file to a new location. Instead of moving the `.proto` file directly and updating all the call sites in a single change, you can put a placeholder `.proto` file in the old location to forward all the imports to the new location using the `import public` notion. ( `import public` is not available in Java )

```protobuf
// new.proto
// All definitions are moved here

// old.proto
import public "new.proto";
import "other.proto";

// client.proto
import "old.proto";
// You use definitions from old.proto and new.proto, but not other.proto
```

The protocol compiler searches for imported files in a set of directories specified on the protocol compiler command line using the `-I`/`--proto_path` flag. If no flag was given, it looks in the directory in which the compiler was invoked. In general you should set the `--proto_path` flag to the root of your project and use fully qualified names for all imports.

#### Nested Types

You can define and use message types inside other message types, as in the following example – here the `Result` message is defined inside the `SearchResponse` message:

```protobuf
message SearchResponse {
  message Result {
    string url = 1;
    string title = 2;
    repeated string snippets = 3;
  }
  repeated Result results = 1;
}
```

If you want to reuse this message type outside its parent message type, you refer to it as `Parent.Type`:

```protobuf
message SomeOtherMessage {
  SearchResponse.Result result = 1;
}
```

#### Updating A Message Type

Don't worry! It's very simple to update message types without breaking any of your existing code. `Just` remember the following (亿点点) rules:

- Don't change the field numbers for any existing fields.
- If you add new fields, any messages serialized by code using your "old" message format can still be parsed by your new generated code. Similarly, messages created by your new code can be parsed by your old code: old binaries simply ignore the new field when parsing. 
- Fields can be removed, as long as the field number is not used again in your updated message type. You may want to rename the field instead, perhaps adding the prefix `OBSOLETE_`, or make the field number [reserved](https://developers.google.com/protocol-buffers/docs/proto3#reserved), so that future users of your `.proto` can't accidentally reuse the number.
- `int32`, `uint32`, `int64`, `uint64`, and `bool` are all compatible – this means you can change a field from one of these types to another without breaking forwards- or backwards-compatibility. You will get the same effect as if you had cast the number to that type in C++ (for example, if a 64-bit number is read as an int32, it will be truncated to 32 bits).
- `string` and `bytes` are compatible as long as the bytes are valid UTF-8.
- ......

For example, when an old binary parses data sent by a new binary with new fields, those new fields become unknown fields in the old binary. Originally, proto3 messages always discarded unknown fields during parsing, but in version 3.5 we reintroduced the preservation of unknown fields to match the proto2 behavior. In versions 3.5 and later, unknown fields are retained during parsing and included in the serialized output.

#### Packages

You can add an optional `package` specifier to a `.proto` file to prevent name clashes between protocol message types.

```protobuf
package foo.bar;
message Open { ... }
```

You can then use the package specifier when defining fields of your message type:

```protobuf
message Foo {
  foo.bar.Open open = 1;
}
```

#### Defining Services

If you want to use your message types with an RPC (Remote Procedure Call) system, you can define an RPC service interface in a `.proto` file and the protocol buffer compiler will generate service interface code and stubs in your chosen language. 

```protobuf
service SearchService {
  rpc Search(SearchRequest) returns (SearchResponse);
}
```

The most straightforward RPC system to use with protocol buffers is [gRPC](https://grpc.io/): a language-neutral open source RPC system developed at Google. gRPC works particularly well with protocol buffers and lets you generate the relevant RPC code directly from your `.proto` files.

有两种不同的结尾: `;`、`{}`,  [它们的区别?](https://stackoverflow.com/questions/30106667/grpc-protobuf-3-syntax-what-is-the-difference-between-rpc-lines-that-end-with-s)

#### Style Guide

- Use an indent of 2 spaces. Prefer the use of double quotes for strings.
- Files should be named `lower_snake_case.proto`.
- All files should be ordered in the following manner:  
  Version (`syntax = "proto3";`) -> Package -> Imports -> File options -> Everything else
- Package names should be in lowercase. Package names should have unique names based on the project name, and possibly based on the path of the `.proto` file.
- Use `CamelCase` (with an initial capital) for message names – for example, `SongServerRequest`. Use `underscore_separated_names` for field names – for example, `song_name`. 
- If your field name contains a number, the number should appear after the letter instead of after the underscore. For example, use `song_name1` instead of `song_name_1`
- Use pluralized names for repeated fields. `  repeated string keys = 1;`
- The zero value enum should have the suffix _UNSPECIFIED. `FOO_BAR_UNSPECIFIED = 0;`
- You should use CamelCase (with an initial capital) for both the service name and any RPC method names.
- Consider creating a subpackage `proto` for `.proto` files, under the root package for your project.

### protoc 命令行选项

To generate the Java, C++, Go code you need to work with the message types defined in a `.proto` file, you need to run the protocol buffer compiler `protoc` on the `.proto`. If you haven't installed the compiler, [download the package](https://developers.google.com/protocol-buffers/docs/downloads) and follow the instructions in the README. For Go, you also need to install a special code generator plugin for the compiler: you can find this and installation instructions in the [golang/protobuf](https://github.com/golang/protobuf/) repository on GitHub.

```bash
protoc --proto_path=IMPORT_PATH --java_out=DST_DIR --go_out=DST_DIR path/to/file.proto
```

- `IMPORT_PATH` specifies a directory in which to look for `.proto` files when resolving `import` directives. If omitted, the current directory is used. `-I` can be used as a short form of `--proto_path`.
- You can provide one or more output directives: `--go_out` generates Go code in `DST_DIR`.
- You must provide one or more `.proto` files as input. Each file must reside in one of the `IMPORT_PATH`s so that the compiler can determine its canonical name.

### [简单了解 Protobuf 编码](https://developers.google.com/protocol-buffers/docs/encoding)

A protocol buffer message is a series of key-value pairs. The binary version of a message just uses the field's number as the key -- the name and declared type for each field can only be determined on the decoding end by referencing the message type's definition (that is, the `.proto` file). When a message is encoded, each key-value pair is turned into a *record* consisting of the field number, a wire type and a payload. This type of scheme is sometimes called Tag-Length-Value, or TLV.

Protobuf 中用 varint 表示整数,  这是一种变长编码,  在编码 `较小的正整数` 时能有效节省字节  
varint 中每个字节只有 7 个有效位,  如果最高位是 1 则表示下一个字节也是当前 varint 的一部分

```go
// 150 的 varint 编码是 [0x96 0x01],  按照下述步骤可从 [0x96 0x01] 解析出 150
10010110 00000001        // Original inputs.
 0010110  0000001        // Drop continuation bits.
 0000001  0010110        // Put into little-endian order.
 10010110                // Concatenate.
 128 + 16 + 4 + 2 = 150  // Interpret as integer.
```

若创建如下的 Test1 消息并把 a 字段设为 `150`,  那么在序列化后占用三个字节: `[08 96 01]`  
其中 [96 01] 是 150 的 varint 编码,  那么 08 是怎么来的呢? 公式为 `varint_encode((field_number << 3) | wire_type)`  
它的含义是先把 field_number 左移三位,  然后用低三位表示 wire_type,  最后用 varint 编码这个整数  
因为 a 字段的编号是 1,  int32 对应的 wire_type 是 0,  代入上述公式可得 `varint_encode(0x08) = 0x08`

```protobuf
message Test1 {
  optional int32 a = 1;
}
```

Tag numbers 1-15 require one less byte to encode than higher numbers, so as an optimization you can decide to use those tags for the commonly used or repeated elements, leaving tags 16 and higher for less-commonly used optional elements.

## 在 Golang 中使用 Protobuf

### 入门例子

(1) 安装 Protobuf 编译器和 Go 语言插件

```bash
# 去 https://github.com/protocolbuffers/protobuf/releases 下载 Protobuf 编译器
# 解压到某个位置, 然后把其中的 bin 目录添加到 PATH
# 还需要执行下述命令安装 Protobuf 的  Go 语言插件
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
```

(2) 在 proto 文件夹中创建 addressbook.proto 文件

```protobuf
syntax = "proto3";
package tutorial;
import "google/protobuf/timestamp.proto";

// 假设 --go_out=$DIR 那么会把生成的代码放到 $DIR/./tutorialpb
option go_package = "./tutorialpb";

message Person {
  int32 id = 1;
  string name = 2;
  string email = 3;
  repeated PhoneNumber phones = 4;
  google.protobuf.Timestamp last_updated = 5;

  // 定义两个嵌套类型
  enum PhoneType {
    MOBILE = 0;
    HOME = 1;
    WORK = 2;
  }
  message PhoneNumber {
    string number = 1;
    PhoneType type = 2;
  }
}

// AddressBook 是一个 Person 列表
message AddressBook {
  repeated Person people = 1;
}
```

(3) 编译 .proto 文件

```bash
# 进入 proto 文件夹,  然后执行
protoc --go_out=. --go_opt=module=xianyukang.com/learn_go/proto addressbook.proto
```

(4) 通过生成的 Person 结构读写数据字段,  通过 proto 包进行序列化和反序列化

```go
import (
    "google.golang.org/protobuf/proto"
    "google.golang.org/protobuf/types/known/timestamppb"
    "xianyukang.com/learn_go/proto/tutorialpb"
)

func main() {
    filename := "proto/person.dat"
    person := &tutorialpb.Person{
        Id:          1,
        Name:        "Homura",
        Email:       "Homura@xenoblade2.com",
        LastUpdated: timestamppb.New(time.Now()),
        Phones:      []*tutorialpb.Person_PhoneNumber{{Number: "123", Type: tutorialpb.Person_MOBILE}},
    }

    // 序列化、然后写入文件
    out, err := proto.Marshal(person)
    if err != nil {
        log.Fatalln("Failed to encode person:", err)
    }
    if err := ioutil.WriteFile(filename, out, 0644); err != nil {
        log.Fatalln("Failed to write person:", err)
    }

    // 读取文件、然后反序列化
    in, err := ioutil.ReadFile(filename)
    if err != nil {
        log.Fatalln("Error reading file:", err)
    }
    person2 := &tutorialpb.Person{}
    if err := proto.Unmarshal(in, person2); err != nil {
        log.Fatalln("Failed to parse person:", err)
    }

    fmt.Println(person)
    fmt.Println(person2)
}
```

### go_package 的含义

- `--go_out=.` 表示生成的代码放到当前目录
- `option go_package` 表示当前文件的「 import path 」,  
  在其他 Go 代码中通过这个 import path 来导入 .proto 定义的结构

```protobuf
// 假设当前文件为 bar.proto, 设置如下的 go_package
option go_package = "xianyukang.com/learn_go/proto/bar";

// 那么 main.go 可以用下面这行 import 来使用 bar.proto 定义的结构
// 另外如果 foo.proto 导入了 bar.proto,  在生成的 foo.pb.go 也会自动添加这样一行 import:
import "xianyukang.com/learn_go/proto/bar"
```

- `--go_opt=module=example.com` 的用法如下:

```protobuf
syntax = "proto3";
package bar;

option go_package = "example.com/bar";

message Guy {
  int32 id = 1;
  string name = 2;
}

// 若执行 protoc --go_out=. bar.go 生成的代码会放在 ./example.com/bar
// 但我们不想要 example.com,  希望直接放在 ./bar

// 可以这样去掉目录结构中的 example.com
// protoc --go_out=. --go_opt=module=example.com bar.proto
```

### 可选使用 JSON 序列化

- Package "google.golang.org/protobuf/encoding/protojson" converts messages to and from JSON.
- This package produces a different output than the standard `encoding/json` package, which does not operate correctly on protocol buffer messages.

### 一些生成 Go 代码的规则

#### (1) Given a simple message declaration:

```protobuf
message Foo {}
```

1. The protocol buffer compiler generates a struct called `Foo`. A `*Foo` implements the [`proto.Message`](https://pkg.go.dev/google.golang.org/protobuf/proto?tab=doc#Message) interface.
2. The [`proto` package](https://pkg.go.dev/google.golang.org/protobuf/proto?tab=doc) provides functions which operate on messages, including conversion to and from binary format.
3. The `proto.Message` interface defines a `ProtoReflect` method. This method returns a [`protoreflect.Message`](https://pkg.go.dev/google.golang.org/protobuf/reflect/protoreflect?tab=doc#Message) which provides a reflection-based view of the message.

#### (2) Nested Types

A message can be declared inside another message. In this case, the compiler generates two structs: `Foo` and `Foo_Bar`.

```protobuf
message Foo {
  message Bar {
  }
}
```

#### (3) Singular Scalar Fields

Note that the generated Go field names always use camel-case naming, even if the field name in the .proto file uses lower-case with underscores. The first letter is capitalized for export. Thus, the proto field `foo_bar_baz` becomes `FooBarBaz` in Go, and `_my_field_name_2` becomes `XMyFieldName_2`.

For this field definition:

```protobuf
int32 foo = 1;
```

The compiler will generate a struct with an `int32` field named `Foo` and an accessor method `GetFoo()` which returns the `int32` value in `Foo` or the zero value of that type if the field is unset. For other field types `int32` is replaced with the corresponding Go type according to the [scalar value types table](https://developers.google.com/protocol-buffers/docs/proto3#scalar).

#### (4) Singular Message Fields

For a message with a `Bar` field:

```protobuf
// proto2
message Baz {
  optional Bar foo = 1;
}

// proto3
message Baz {
  Bar foo = 1;
}
```

The compiler will generate a Go struct:

```go
type Baz struct {
    Foo *Bar
}
```

Message fields can be set to `nil`, which means that the field is unset, effectively clearing the field. This is not equivalent to setting the value to an "empty" instance of the message struct.

The compiler also generates a `func (m *Baz) GetFoo() *Bar` helper function. This function returns a `nil` `*Bar` if `m` is nil or `foo` is unset. This makes it possible to chain get calls without intermediate `nil` checks.

```go
f := (*foo.Foo)(nil)               // f 是 nil
f.GetBar()                         // 但调用 f.GetBar() 也不会爆空指针,  会返回 nil *Bar
f.GetBar().GetValue()              // 所以方便了链式调用、无需逐层检查 if ... != nil
fmt.Println(f.GetBar().GetValue()) // 妹有设置则返回零值, 比如这里会打印空字符串
```

#### (5) Repeated Fields

Each repeated field generates a slice field. For this message with a repeated field:

```protobuf
message Baz {
  repeated Bar foo = 1;
}
```

the compiler generates the Go struct:

```go
type Baz struct {
    Foo  []*Bar
}
```

#### (6) Map Fields

For this message with a map field:

```protobuf
message Bar {}

message Baz {
  map<string, Bar> foo = 1;
}
```

the compiler generates the Go struct:

```go
type Baz struct {
    Foo map[string]*Bar
}
```

#### (7) Oneof Fields

用到再看吧,  参考[这个文档](https://developers.google.com/protocol-buffers/docs/reference/go-generated#oneof),  什么时候使用 oneof:  
When a message has many optional fields and at most one field will be set at the same time.

#### (8) Enumerations

(1) Given an enumeration like this, the compiler generates a type and a series of constants with that type.

```protobuf
message SearchRequest {
  enum Corpus {
    UNIVERSAL = 0;
    WEB = 1;
    IMAGES = 2;
    VIDEO = 3;
  }
  Corpus corpus = 1;
}
```

For enums within a message (like the one above), the type name begins with the message name:

```go
type SearchRequest_Corpus int32
```

(2) For a package-level enum:

```protobuf
enum Foo {
  DEFAULT_BAR = 0;
  BAR_BELLS = 1;
  BAR_B_CUE = 2;
}
```

the Go type name is unmodified from the proto enum name:

```go
type Foo int32
```

(3) This type has a `String()` method that returns the name of a given value.

(4) The protocol buffer compiler generates a constant for each value in the enum. For enums within a message, the constants begin with the enclosing message's name: `const SearchRequest_UNIVERSAL SearchRequest_Corpus = 0`. For a package-level enum, the constants begin with the enum name instead: `const Corpus_UNIVERSAL Corpus = 0`

(5) The protobuf compiler also generates a map from integer values to the string names and a map from the names to the values:

```go
var (
    Corpus_name = map[int32]string{
        0: "UNIVERSAL",
        1: "WEB",
        2: "IMAGES",
        3: "VIDEO",
    }
    Corpus_value = map[string]int32{
        "UNIVERSAL": 0,
        "WEB":       1,
        "IMAGES":    2,
        "VIDEO":     3,
    }
)
```

#### (9) Services

The Go code generator does not produce output for services by default. If you enable the [gRPC](https://www.grpc.io/) plugin (see the [gRPC Go Quickstart guide](https://github.com/grpc/grpc-go/tree/master/examples)) then code will be generated to support gRPC.

## 标准库的 RPC

### RPC 是什么

RPC 是远程过程调用（Remote Procedure Call）的缩写，通俗地说就是调用远处的一个函数  
远处可以是「 同一个机器的另一个进程的函数 」, 也可以是 「 远程机器中的某个函数 」  
RPC 是为了调用远程机器的功能,  比如浏览器调用 server、微服务间的互相调用,  都是一种 RPC

### RPC 入门

RPC 是远程过程调用的简称，是分布式系统中不同节点间流行的通信方式。在互联网时代，RPC 已经和 IPC 一样成为一个不可或缺的基础构件。因此 Go 语言的标准库也提供了一个简单的 RPC 实现。由这个例子可以看出 RPC 的使用其实非常简单。

```go
//////////////////////////////////////////// server.go
func main() {
    // RegisterName 会注册所有满足 RPC 规则的方法
    // 客户端可通过 HelloService.Hello 调用服务端的 Hello 方法
    _ = rpc.RegisterName("HelloService", &HelloService{})
    listener, err := net.Listen("tcp", ":1234")
    if err != nil {
        log.Fatal("listen tcp error:", err)
    }
    conn, err := listener.Accept()
    if err != nil {
        log.Fatal("accept error:", err)
    }
    rpc.ServeConn(conn)
}

type HelloService struct{}

// Hello 方法必须满足 Go 语言的 RPC 规则:
// 方法只能有两个可序列化的参数，其中第二个参数是指针类型，
// 并且返回一个 error 类型，同时必须是公开的方法。
func (h *HelloService) Hello(req string, resp *string) error {
    *resp = "Hello: " + req
    return nil
}

//////////////////////////////////////////// client.go
func main() {
    client, err := rpc.Dial("tcp", "localhost:1234")
    if err != nil {
        log.Fatal("dial error:", err)
    }
    var resp string
    err = client.Call("HelloService.Hello", "Homura", &resp)
    if err != nil {
        log.Fatal("call error:", err)
    }
    fmt.Println(resp)
}
```

### 重构上述例子

Hello 服务的接口:

```go
package hello
import "net/rpc"

// (1) 服务名
const ServiceName = "com.xxx.HelloService"

// (2) 服务中的方法列表,  这个接口供两端共同使用
type ServiceInterface interface {
    Hello(req string, resp *string) error
}

// (3) 用于注册服务的函数
func RegisterService(s ServiceInterface) error {
    return rpc.RegisterName(ServiceName, s)
}
```

客户端实现:

```go
type HelloServiceClient struct {
    *rpc.Client
}

// (1) 让编译器确保客户端实现了 HelloServiceInterface 接口
var _ hello.ServiceInterface = (*HelloServiceClient)(nil)

// (2) 客户端构造函数
func DialHelloService(network, address string) (*HelloServiceClient, error) {
    c, err := rpc.Dial(network, address)
    if err != nil {
        return nil, err
    }
    return &HelloServiceClient{Client: c}, nil
}

// (3) 实现接口中定义的函数,  注意 Hello 具有类型安全而 Client.Call 的参数都是 any
//     现在客户端用户不用再担心 RPC 方法名字或参数类型不匹配等低级错误的发生。
func (h *HelloServiceClient) Hello(req string, resp *string) error {
    return h.Client.Call(hello.ServiceName+".Hello", "Homura", resp)
}

func main() {
    client, err := DialHelloService("tcp", "localhost:1234")
    if err != nil {
        log.Fatal("dial error: ", err)
    }
    var resp string
    err = client.Hello("Homura", &resp)
    if err != nil {
        log.Fatal("call error: ", err)
    }
    fmt.Println(resp)
}
```

服务端实现:

```go
func main() {
    // 服务端只管具体实现,  像服务名、服务的方法列表则在接口中定义
    _ = hello.RegisterService(&HelloService{})
    listener, err := net.Listen("tcp", ":1234")
    if err != nil {
        log.Fatal("listen tcp error:", err)
    }
    for {
        // Accept 会返回一个连接,  然后用一个 goroutine 来处理这个连接
        conn, err := listener.Accept()
        if err != nil {
            log.Fatal("accept error:", err)
        }
        go rpc.ServeConn(conn)
    }
}

type HelloService struct{}

func (h *HelloService) Hello(req string, resp *string) error {
    *resp = "Hello: " + req
    return nil
}
```

#### JSON 序列化方式

- 标准库的 RPC 默认用 Golang 自家的 Gob 序列化,  如果改成 JSON 则方便其他语言调用
- 只要知道了 RPC 请求的格式,  在其他语言中构造并发送相同格式的请求,  就能完成跨语言调用

```go
// 服务端改动
// go rpc.ServeConn(conn)
go rpc.ServeCodec(jsonrpc.NewServerCodec(conn)) // 自定义编解码方式

// 客户端改动
// client, err := rpc.Dial(network, address)
conn, err := net.Dial(network, address)
client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn)) // 自定义编解码方式
```

查看客户端的 JSON RPC 请求:

```bash
nc -l 1234                                                           # 开一个 TCP 服务打印收到的请求
go run ...                                                           # 启动客户端
{"method":"com.xxx.HelloService.Hello","params":["Homura"],"id":0}   # 这就是 json rpc 请求
```

查看服务端返回的 JSON RPC 响应:

```bash
go run ...                                                                                       # 启动服务端
echo -e '{"method":"com.xxx.HelloService.Hello","params":["Homura"],"id":0}' | nc localhost 1234 # 发送请求
{"id":0,"result":"Hello: Homura","error":null}                                                   # 这就是 json rpc 响应
```

#### HTTP 传输协议

- 使用 HTTP 而不是 TCP 传输 RPC 请求,  能方便从不同语言中访问 RPC 服务。

```go
type HelloService struct{}
func (h *HelloService) Hello(req string, resp *string) error {
    *resp = "Hello: " + req
    return nil
}
func main() {
    _ = rpc.RegisterName("HelloService", new(HelloService))
    http.HandleFunc("/jsonrpc", func(w http.ResponseWriter, r *http.Request) {
        var conn io.ReadWriteCloser = struct {
            io.Writer
            io.ReadCloser
        }{
            ReadCloser: r.Body,
            Writer:     w,
        }
        // http 请求的 body 是一个 json rpc 请求,  总之把 http 请求交给 rpc 服务进行处理
        _ = rpc.ServeRequest(jsonrpc.NewServerCodec(conn))
    })
    _ = http.ListenAndServe(":1234", nil)
}

// 验证方法:
// curl localhost:1234/jsonrpc -X POST --data '{"method":"HelloService.Hello","params":["hello"],"id":0}'
```

