## Table of Contents
  - [Golang 中使用 gRPC](#Golang-%E4%B8%AD%E4%BD%BF%E7%94%A8-gRPC)
    - [基础教程](#%E5%9F%BA%E7%A1%80%E6%95%99%E7%A8%8B)
    - [入门例子](#%E5%85%A5%E9%97%A8%E4%BE%8B%E5%AD%90)
    - [错误处理](#%E9%94%99%E8%AF%AF%E5%A4%84%E7%90%86)
    - [命令行工具](#%E5%91%BD%E4%BB%A4%E8%A1%8C%E5%B7%A5%E5%85%B7)
    - [Interceptor](#Interceptor)
    - [生成 RESTful API](#%E7%94%9F%E6%88%90-RESTful-API)
  - [gRPC 的相关概念](#gRPC-%E7%9A%84%E7%9B%B8%E5%85%B3%E6%A6%82%E5%BF%B5)
    - [概述](#%E6%A6%82%E8%BF%B0)
    - [核心概念](#%E6%A0%B8%E5%BF%83%E6%A6%82%E5%BF%B5)
    - [RPC Life Cycle](#RPC-Life-Cycle)

## Golang 中使用 gRPC

### 基础教程

- [The complete gRPC course (强烈推荐:代码风格很好)](https://www.youtube.com/playlist?list=PLy_6D98if3UJd5hxWNfAqKMr15HZqFnqf)
- 有四种类型的 RPC: [服务端示例](https://grpc.io/docs/languages/go/basics/#simple-rpc)、[客户端示例](https://grpc.io/docs/languages/go/basics/#calling-service-methods)
- [grpc-go 官方例子](https://github.com/grpc/grpc-go/tree/master/examples)

### 入门例子

(1) 在安装 protoc 编译器之后,  要安装 Go 语言插件,  gRPC 插件

```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

(2) 编写 .proto 文件,  然后生成代码

```protobuf
syntax = "proto3";
package helloworld;

option go_package = "grpc-example/helloworld";

service Greeter {
  rpc SayHello (HelloRequest) returns (HelloReply) {}
}

message HelloRequest {
  string name = 1;
}

message HelloReply {
  string message = 1;
}
// 在项目根目录执行这个生成代码:
// protoc.exe -I=helloworld --go_out=.. --go-grpc_out=.. helloworld/*.proto
```

(3) 使用生成的代码编写服务端

```go
import (
    "context"
    "google.golang.org/grpc"
    "google.golang.org/grpc/reflection"
    pb "grpc-example/helloworld"
)

type server struct {
    // 如果 Service 接口添加了新的方法,  那么 ServiceImpl 也得实现这些方法,  否则无法编译
    // 所以 ServiceImpl 没有 Forward Capability,  通过嵌入 UnimplementedServer 就能让旧实现也满足新接口
    pb.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
    // 实现 RPC 方法
    return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}

func main() {
    // 启动 RPC 服务
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    s := grpc.NewServer()
    pb.RegisterGreeterServer(s, &server{})
    // Register reflection service on gRPC server.
    // 用于反射服务端有哪些 RPC 方法,  有哪些 message 结构
    // 所以用 grpcurl 来测试 RPC 接口时无需提供 .proto 文件
    reflection.Register(s)
    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}
```

(4) 使用生成的代码编写客户端

```go
import (
    "context"
    "google.golang.org/grpc"
    "google.golang.org/grpc/credentials/insecure"
    pb "grpc-example/helloworld"
)

func main() {
    // Set up a connection to the server.
    conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
    if err != nil {
        log.Fatalf("did not connect: %v", err)
    }
    defer conn.Close()
    client := pb.NewGreeterClient(conn)

    // Contact the server and print out its response.
    name := "Hikari"
    if len(os.Args) > 1 {
        name = os.Args[1]
    }
    ctx, cancel := context.WithTimeout(context.Background(), time.Second)
    defer cancel()
    // 使用生成的 RPC 客户端以及生成的 HelloRequest 来调用服务端方法
    r, err := client.SayHello(ctx, &pb.HelloRequest{Name: name})
    if err != nil {
        log.Fatalf("could not greet: %v", err)
    }
    log.Printf("Greeting: %s", r.Message)
}
```

### 错误处理

- [返回错误码和消息 (status.Errorf)](https://github.com/avinassh/grpc-errors/tree/master/go)

- [返回额外的错误详情 (st.WithDetails)](https://medium.com/utility-warehouse-technology/advanced-grpc-error-usage-1b37398f0ff4)

``` go
// 如果服务端返回这个, 客户端会打印 rpc error: code = Unknown desc = xxx error
return errors.New("xxx error")

// 如果服务端返回这个, 客户端会打印 rpc error: code = InvalidArgument desc = xxx error
return status.Errorf(codes.InvalidArgument, "%v error", "xxx")

// 客户端用 status.Convert 获取调用状态, 另外 status.FromError 也是同样的效果但它的第二个返回值没用
resp, err = client.SayHello(context.Background(), &api.HelloReq{Name: "Leonhard Euler"})
st := status.Convert(err)

// 客户端这样解析服务器发过来的错误详情:
for _, detail := range st.Details() {
    switch t := detail.(type) {
    case *errdetails.BadRequest:
        // 参考上面第 2 个链接
    }
}
```

#### Richer error model

The simple error model (code and message only, no details) is the official gRPC error model, is supported by all gRPC client/server libraries, and is independent of the gRPC data format (whether protocol buffers or something else). You may have noticed that it’s quite limited and doesn’t include the ability to communicate error details.

If you’re using protocol buffers as your data format, however, you may wish to consider using the richer error model developed and used by Google as described [here](https://cloud.google.com/apis/design/errors#error_model). This model enables servers to return and clients to consume additional error details expressed as one or more protobuf messages. It further specifies a [standard set of error message types](https://github.com/googleapis/googleapis/blob/master/google/rpc/error_details.proto) to cover the most common needs (such as invalid parameters, quota violations, and stack traces). The protobuf binary encoding of this extra error information is provided as trailing metadata in the response.

### 命令行工具

- [ktr0731/evans: Evans: more expressive universal gRPC client](https://github.com/ktr0731/evans)
- [fullstorydev/grpcurl: Like cURL, but for gRPC](https://github.com/fullstorydev/grpcurl)

```bash
# 一些 evans 命令
evans -r -p 8080 repl    # 以反射模式启动 evans
show package             # 列出所有包
package api              # 选择一个包
show service             # 列出所有服务和方法
show message             # 列出所有 message
desc RateLaptopRequest   # 展示某个 message 的定义
show header              # 列出请求头
header foo=bar           # 设置请求头
service Example          # 选择某个服务
call Unary               # 调用某个方法, 会提示输入每个字段, 按 Ctrl+D 退出 repeated field
```

### Interceptor

- [A Guide to gRPC and Interceptors](https://edgehog.blog/a-guide-to-grpc-and-interceptors-265c306d3773)
- 总之相当于中间件，可以用来做日志追踪、限流、用户验证(Authentication)、权限验证(Authorization)

Interceptors are neat little components of a gRPC application that allow us to interact with a proto message or context either before it is sent by the client or after it is received by the server. 

For example, we may want to modify each request before it is sent, perhaps by adding some information or metadata about the environment the client is running in — and we can do that! On the server side — we can also intercept that message before the actual function call is executed — perhaps running some validation or a check before our main business logic is run.

有四种 Interceptor 类型:

1. `Client Unary Interceptor` These are executed on the client before a singular request is made to a function on the server side. This is where we might enrich a message with client side metadata.
2. `Client Stream Interceptor` For example, if we were streaming a list of 100 objects to the server, such as chunks of a file, we could intercept before sending each chunk.
3. `Server Unary Interceptor` These are executed server side when a singular request is received from a client. It’s at this point we may want to perform some checks on the authenticity of the request or checking that certain fields are present.
4. `Server Stream Interceptor` 这个和上一个的作用一样，如果客户端请求是 stream 而不是 unary 则需要用这个

### 生成 RESTful API

This helps you provide your APIs in both gRPC and RESTful style at the same time. 

The [gRPC-Gateway](https://github.com/grpc-ecosystem/grpc-gateway) is a plugin of the Google protocol buffers compiler [protoc](https://github.com/protocolbuffers/protobuf). It reads protobuf service definitions and generates a reverse-proxy server which translates a RESTful HTTP API into gRPC. This server is generated according to the [`google.api.http`](https://github.com/googleapis/googleapis/blob/master/google/api/http.proto#L46) annotations in your service definitions.

## gRPC 的相关概念

### 概述

#### 是什么?

gRPC is a modern, open source remote procedure call (RPC) framework that can run anywhere. It enables client and server applications to communicate transparently, and makes it easier to build connected systems.

#### 可以干什么?

It can efficiently connect services in and across data centers with pluggable support for load balancing, tracing, health checking and authentication. It is also applicable to connect devices, mobile applications and browsers to backend services.

The main usage scenarios:

- Low latency, highly scalable, distributed systems.
- Developing mobile clients which are communicating to a cloud server.
- Designing a new protocol that needs to be accurate, efficient and language independent.

#### gRPC vs REST

- [Comparison](https://youtu.be/3ViS8DAqLME?t=100)
- [Where to use gRPC?](https://youtu.be/3ViS8DAqLME?t=208)

#### Overview

In gRPC, a client application can directly call a method on a server application on a different machine as if it were a local object, making it easier for you to create distributed applications and services. As in many RPC systems, gRPC is based around the idea of defining a service, specifying the methods that can be called remotely with their parameters and return types. On the server side, the server implements this interface and runs a gRPC server to handle client calls. On the client side, the client has a stub (placeholder) that provides the same methods as the server.

![Concept Diagram](https://grpc.io/img/landing-2.svg) 

- 可以跨语言调用,  比如 Android-Java 调用后端 C++ 服务
- 上至集群中的服务器, 下至家里的电脑都能用 gRPC 通讯
- 一些 Google API 甚至提供 gRPC 版本来方便你调用

gRPC clients and servers can run and talk to each other in a variety of environments - from servers inside Google to your own desktop - and can be written in any of gRPC’s supported languages. So, for example, you can easily create a gRPC server in Java with clients in Go, Python, or Ruby. In addition, the latest Google APIs will have gRPC versions of their interfaces, letting you easily build Google functionality into your applications.

#### 和 Protobuf 的关系

By default, gRPC uses [Protocol Buffers](https://developers.google.com/protocol-buffers/docs/overview), Google’s mature open source mechanism for serializing structured data (although it can be used with other data formats such as JSON). gRPC can use protocol buffers as both its Interface Definition Language (**IDL**) and as its underlying message interchange format. 

You define gRPC services in ordinary proto files, with RPC method parameters and return types specified as protocol buffer messages. gRPC uses `protoc` with a special gRPC plugin to generate code from your proto file: you get generated gRPC client and server code, as well as the regular protocol buffer code for populating, serializing, and retrieving your message types.

```protobuf
// 定义 Greeter 服务的方法列表
service Greeter {
  rpc SayHello (HelloRequest) returns (HelloReply) {}
}

message HelloRequest {
  string name = 1;
}
message HelloReply {
  string message = 1;
}
```

### 核心概念

#### Service definition 

Like many RPC systems, gRPC is based around the idea of defining a service, specifying the methods that can be called remotely with their parameters and return types. By default, gRPC uses protocol buffers as the Interface Definition Language (IDL) for describing both the service interface and the structure of the payload messages.

gRPC lets you define four kinds of service method:

- `Unary RPCs` where the client sends a single request to the server and gets a single response back, just like a normal function call.
- `Server streaming RPCs` where the client sends a request to the server and gets a stream to read a sequence of messages back. The client reads from the returned stream until there are no more messages. gRPC guarantees message ordering within an individual RPC call.
- `Client streaming RPCs` where the client writes a sequence of messages and sends them to the server, again using a provided stream. Once the client has finished writing the messages, it waits for the server to read them and return its response. Again gRPC guarantees message ordering within an individual RPC call.
- `Bidirectional streaming RPCs` where both sides send a sequence of messages using a read-write stream. The two streams operate independently, so clients and servers can read and write in whatever order they like: for example, the server could wait to receive all the client messages before writing its responses, or it could alternately read a message then write a message, or some other combination of reads and writes. The order of messages in each stream is preserved.

```protobuf
rpc SayHello(HelloRequest) returns (HelloResponse);                  // unary rpc
rpc LotsOfReplies(HelloRequest) returns (stream HelloResponse);      // server streaming rpc
rpc LotsOfGreetings(stream HelloRequest) returns (HelloResponse);    // client streaming rpc
rpc BidiHello(stream HelloRequest) returns (stream HelloResponse);   // Bidirectional streaming RPCs
```

#### Using the API 

Starting from a service definition in a `.proto` file, gRPC provides protocol buffer compiler plugins that generate client- and server-side code. gRPC users typically call these APIs on the client side and implement the corresponding API on the server side.

- On the server side, the server implements the methods declared by the service and runs a gRPC server to handle client calls. The gRPC infrastructure decodes incoming requests, executes service methods, and encodes service responses.
- On the client side, the client has a local object known as stub that implements the same methods as the service. The client can then just call those methods on the local object, and the methods wrap the parameters for the call in the appropriate protocol buffer message type, send the requests to the server, and return the server’s protocol buffer responses.

#### Synchronous vs. asynchronous

Synchronous RPC calls that block until a response arrives from the server (在响应到达前会一直阻塞的同步 RPC 调用) are the closest approximation to (最接近于) the abstraction of a procedure call that RPC aspires to. (RPC 所追求的对「过程调用」的抽象)

On the other hand, networks are inherently asynchronous and in many scenarios it’s useful to be able to start RPCs without blocking the current thread. The gRPC programming API in most languages comes in both synchronous and asynchronous flavors.

### RPC Life Cycle

#### (1) Unary RPC

总之先发送 metadata 然后再发送 request message,  另外返回 response 时还会附带 status 信息  
First consider the simplest type of RPC where the client sends a single request and gets back a single response.

1. Once the client calls a stub method, the server is notified that the RPC has been invoked with the client’s [metadata](https://grpc.io/docs/what-is-grpc/core-concepts/#metadata) for this call, the method name, and the specified deadline if applicable.
2. The server can then either send back its own initial metadata (which must be sent before any response) straight away, or wait for the client’s request message. Which happens first, is application-specific.
3. Once the server has the client’s request message, it does whatever work is necessary to create and populate a response. The response is then returned (if successful) to the client together with status details (status code and optional status message) and optional trailing metadata.
4. If the response status is OK, then the client gets the response, which completes the call on the client side.

#### (2) Server streaming RPC

A server-streaming RPC is similar to a unary RPC, except that the server returns a stream of messages in response to a client’s request. *After sending all its messages, the server’s status details* (status code and optional status message) and optional trailing metadata are sent to the client. *This completes processing on the server side*. The client completes once it has all the server’s messages.

#### (3) Client streaming RPC

A client-streaming RPC is similar to a unary RPC, except that the client sends a stream of messages to the server instead of a single message. The server responds with a single message (along with its status details and optional trailing metadata), typically but not necessarily after it has received all the client’s messages.

#### (4) Bidirectional streaming RPC

In a bidirectional streaming RPC, the call is initiated by the client invoking the method and the server receiving the client metadata, method name, and deadline. The server can choose to send back its initial metadata or wait for the client to start streaming messages.

Client- and server-side stream processing is application specific. Since the two streams are independent, the client and server can read and write messages in any order. For example, a server can wait until it has received all of a client’s messages before writing its messages, or the server and client can play “ping-pong” – the server gets a request, then sends back a response, then the client sends another request based on the response, and so on.

#### (5) Deadlines/Timeouts

gRPC allows clients to specify how long they are willing to wait for an RPC to complete before the RPC is terminated with a `DEADLINE_EXCEEDED` error. On the server side, the server can query to see if a particular RPC has timed out, or how much time is left to complete the RPC.

#### (6) RPC termination

In gRPC, both the client and server make independent and local determinations of the success of the call, and their conclusions may not match. This means that, for example, you could have an RPC that finishes successfully on the server side (“I have sent all my responses!”) but fails on the client side (“The responses arrived after my deadline!”). It’s also possible for a server to decide to complete before a client has sent all its requests.

#### (7) Cancelling an RPC

Either the client or the server can cancel an RPC at any time. A cancellation terminates the RPC immediately so that no further work is done. **Warning: Changes made before a cancellation are not rolled back.**

#### (8) Metadata

Metadata is information about a particular RPC call (such as [authentication details](https://grpc.io/docs/guides/auth/)) in the form of a list of key-value pairs, where the keys are strings and the values are typically strings, but can be binary data.

Keys are case insensitive and consist of ASCII letters, digits, and special characters `-`, `_`, `.` and must not start with `grpc-` (which is reserved for gRPC itself). Binary-valued keys end in `-bin` while ASCII-valued keys do not.

User-defined metadata is not used by gRPC, which allows the client to provide information associated with the call to the server and vice versa. Access to metadata is language dependent.

#### (9) Channels

A gRPC channel provides a connection to a gRPC server on a specified host and port. It is used when creating a client stub. Clients can specify channel arguments to modify gRPC’s default behavior, such as switching message compression on or off. A channel has state, including `connected` and `idle`.

