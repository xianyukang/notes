## Table of Contents
  - [Gin](#Gin)
    - [参考](#%E5%8F%82%E8%80%83)
    - [概述](#%E6%A6%82%E8%BF%B0)
  - [基础概念](#%E5%9F%BA%E7%A1%80%E6%A6%82%E5%BF%B5)
    - [路由](#%E8%B7%AF%E7%94%B1)
    - [查询参数、表单](#%E6%9F%A5%E8%AF%A2%E5%8F%82%E6%95%B0%E8%A1%A8%E5%8D%95)
    - [中间件](#%E4%B8%AD%E9%97%B4%E4%BB%B6)
    - [模型绑定](#%E6%A8%A1%E5%9E%8B%E7%BB%91%E5%AE%9A)
    - [各种类型的响应](#%E5%90%84%E7%A7%8D%E7%B1%BB%E5%9E%8B%E7%9A%84%E5%93%8D%E5%BA%94)
    - [配置 HTTP Server](#%E9%85%8D%E7%BD%AE-HTTP-Server)
    - [编写测试](#%E7%BC%96%E5%86%99%E6%B5%8B%E8%AF%95)
    - [优雅地关闭服务](#%E4%BC%98%E9%9B%85%E5%9C%B0%E5%85%B3%E9%97%AD%E6%9C%8D%E5%8A%A1)
  - [模型校验](#%E6%A8%A1%E5%9E%8B%E6%A0%A1%E9%AA%8C)
    - [概述](#%E6%A6%82%E8%BF%B0)
    - [简单例子](#%E7%AE%80%E5%8D%95%E4%BE%8B%E5%AD%90)
    - [各种自定义](#%E5%90%84%E7%A7%8D%E8%87%AA%E5%AE%9A%E4%B9%89)
    - [各种校验 tag](#%E5%90%84%E7%A7%8D%E6%A0%A1%E9%AA%8C-tag)

## Gin

### 参考

- [官方文档](https://github.com/gin-gonic/gin)
- [官方提供了众多的学习例子](https://github.com/gin-gonic/examples)

### 概述

Why gin? ①标准库功能不够强大 ②用框架更方便、提高生产力 ③流行度最高的是 gin

While the `net/http` package allows you to craft a web application easily, the routing mechanism is not so powerful, especially for complex applications. That's where a web framework comes into play. 

Gin is a minimalistic web framework suitable for building web applications, microservices, and RESTful APIs. It reduces boilerplate code by creating reusable and extensible pieces of code: you can write a piece of middleware that can be plugged into one or more request handlers. Moreover, it comes with the following key features:  

- *Well documented*: The documentation for Gin is broad and comprehensive. Most tasks that you will need to do relating to the router can be found easily in the docs.
- *Simplicity*: Gin is a pretty minimalistic framework. Only the most essential features and libraries are included, with little to no boilerplate to bootstrap applications, making Gin a great framework for developing highly available REST APIs.
- *Extensible*: The Gin community has created numerous pieces of well-tested middleware that make developing for Gin a charm. Features include compression with GZip, authentication with an authorization middleware, and logging with external solutions such as Sentry.
- *Performance*: Gin runs 40x faster than Martini and runs comparatively well compared to other Golang frameworks.

## 基础概念

### 路由

#### ➤ `/user/:name` 会匹配什么?  `/user/:name/*action` 又会匹配什么?  [答案](https://github.com/gin-gonic/gin#parameters-in-path)  
`:name` 只能匹配非空字符串,  所以 /user/:name 不会匹配 /user 和 /user/   
`*action` 相比 :name 则多了空字符串匹配, 比如 /user/:name/*action 可以匹配 /user/john/  
另外没有其他路由能匹配 /user/john 时,  /user/john 会被重定向到 /user/john/

#### ➤ 相关函数  
①用 c.Param("name") 读取路径参数 Path Parameter 的值  
②用 c.FullPath() 获取路由定义, 比如  c.FullPath() == "/user/:name/*action"

#### ➤ 精确匹配的优先级更高  
`router.GET("/user/groups", ...)` will add a new router for /user/groups.  
*Exact routes are resolved before param routes, regardless of the order they were defined*.

#### ➤ 路由分组

```go
v1 := router.Group("/v1")                // Simple group: v1
{
    v1.POST("/login", loginEndpoint)     // 用大括号包起来是为了美观、可读性
    v1.POST("/submit", submitEndpoint)
    v1.POST("/read", readEndpoint)
}
```

### 查询参数、表单

#### ➤ 获取查询参数、读取表单

```go
firstname := c.DefaultQuery("firstname", "Guest")   // 读取 Query Parameter 时提供默认值
lastname := c.Query("lastname")                     // shortcut for c.Request.URL.Query().Get("lastname")
nick := c.DefaultPostForm("nick", "anonymous")      // 读取表单时提供默认值
message := c.PostForm("message")                    // 如果 message 表单字段不存在则返回空字符串
/post?ids[a]=1234&ids[b]=hello                      // Query Parameter 和表单都可以使用 Map 语法
var ids map[string]string = c.QueryMap("ids")       // 然后用 c.QueryMap 或 c.PostFormMap 解析出 Map
```

#### ➤ 上传单个或多个文件  

- [参考这里](https://github.com/gin-gonic/gin#upload-files)

### 中间件

#### ➤ New() 与 Default() 的区别 ?

```go
r := gin.New()     // Blank Gin without middleware by default
r := gin.Default() // Default With the Logger and Recovery middleware already attached
```

#### ➤ 使用中间件

```go
func 使用中间件() {
    // Creates a router without any middleware by default
    r := gin.New()

    // Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
    // By default, gin.DefaultWriter = os.Stdout
    r.Use(gin.Logger())

    // Recovery middleware recovers from any panics and writes a 500 if there was one.
    r.Use(gin.Recovery())

    // Per route middleware, you can add as many as you desire.
    r.GET("/benchmark", MyBenchLogger(), benchEndpoint)

    // Authorization group, and per group middleware!
    authorized := r.Group("/")
    authorized.Use(AuthRequired())
    {
        authorized.POST("/submit", submitEndpoint)
        authorized.POST("/read", readEndpoint)

        testing := authorized.Group("testing")       // nested group
        testing.GET("/analytics", analyticsEndpoint) // a route for /testing/analytics
    }
}
```

#### [➤ 如何编写自定义中间件、以及 c.Set() 的作用](https://github.com/gin-gonic/gin#custom-middleware)

#### [➤ 自定义 Panic Recovery 中间件](https://github.com/gin-gonic/gin#custom-recovery-behavior)

#### ➤ 注意在 Middleware/Handler 中启动新的 Goroutine 时,  必须 c.Copy() 复制一下 Context, [例子](https://github.com/gin-gonic/gin#goroutines-inside-a-middleware)

#### ➤ 日志相关

- [如何修改 DefaultWriter 把日志保存到文件](https://github.com/gin-gonic/gin#how-to-write-log-file)
- [如何自定义日志格式: 比如请求时间、IP、用什么方法访问了什么端点、正常返回还是有错误](https://github.com/gin-gonic/gin#custom-log-format)
- [如果把日志保存到文件,  那么应该关掉 coloring](https://github.com/gin-gonic/gin#controlling-log-output-coloring)
- [define format for the log of routes](https://github.com/gin-gonic/gin#define-format-for-the-log-of-routes)

### 模型绑定

To bind a request body into a type, use model binding. We currently support binding of JSON, XML, YAML, TOML and standard form values (foo=bar&boo=baz). Note that you need to set the corresponding binding tag on all fields you want to bind. For example, when binding from JSON, set `json:"fieldname"`.

#### ➤ 两类 Binding 方法

- Must bind
  - **Methods** - `Bind`, `BindJSON`, `BindXML`, `BindQuery`, `BindYAML`, `BindHeader`, `BindTOML`
  - **Behavior** - These methods use `MustBindWith` under the hood. If there is a binding error, the request is aborted with `c.AbortWithError(400, err).SetType(ErrorTypeBind)`. This sets the response status code to 400 and the `Content-Type` header is set to `text/plain; charset=utf-8`. Note that if you try to set the response code after this, it will result in a warning `[GIN-debug] [WARNING] Headers were already written. Wanted to override status code 400 with 422`. If you wish to have greater control over the behavior, consider using the `ShouldBind` equivalent method.
- Should bind
  - **Methods** - `ShouldBind`, `ShouldBindJSON`, `ShouldBindXML`, `ShouldBindQuery`, `ShouldBindYAML`, ...
  - **Behavior** - These methods use `ShouldBindWith` under the hood. If there is a binding error, the error is returned and it is the developer's responsibility to handle the request and error appropriately.

When using the `Bind` method, Gin tries to infer the binder depending on the Content-Type header. If you are sure what you are binding, you can use `MustBindWith` or `ShouldBindWith`.

#### ➤ 理解 binding:"required" 与 empty value

1. You can also specify that specific fields are required.  
    If a field is decorated with `binding:"required"` and has a empty value when binding, an error will be returned.
2. 类型的零值、null值、和不提供某字段, 这些都算 empty value, 所以 {id:0} 中的 id 也会被当做 empty value
3. 如果想区分 {id:null} 和 {id:0}, 试试指针类型 `*int` 并检查字段值是否为 `nil`
4. 另外还可以用 `binding:"-"` 表示不校验某字段

```go
type Login struct {
    User     string `form:"user" json:"user" xml:"user"  binding:"required"`            // 同时设置 form 和 json 绑定
    Password string `form:"password" json:"password" xml:"password" binding:"required"` // binding:"required" 表示必须
}

func main() {
    router := gin.Default()
    // Example for binding JSON ({"user": "manu", "password": "123"})
    router.POST("/loginJSON", func(c *gin.Context) {
        var json Login
        if err := c.ShouldBindJSON(&json); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }
    })
    // Example for binding a HTML form (user=manu&password=123)
    router.POST("/loginForm", func(c *gin.Context) {
        var form Login
        // This will infer what binder to use depending on the content-type header.
        if err := c.ShouldBind(&form); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }
})}
```

#### ➤ 更多例子

- [#only-bind-query-string](https://github.com/gin-gonic/gin#only-bind-query-string)
- [#bind-query-string-or-post-data](https://github.com/gin-gonic/gin#bind-query-string-or-post-data)
- [#bind-header](https://github.com/gin-gonic/gin#bind-header)、[#bind-uri](https://github.com/gin-gonic/gin#bind-uri) 通过绑定 header、url parameter,  就能校验 header、url 中的参数是否有误
- [#bind-html-checkboxes](https://github.com/gin-gonic/gin#bind-html-checkboxes)、[#upload-files](https://github.com/gin-gonic/gin#upload-files)

### 各种类型的响应

#### ➤ JSON

[JSON 响应例子](https://github.com/gin-gonic/gin#xml-json-yaml-and-protobuf-rendering): 除了 JSON 还支持输出 XML/YAML/ProtoBuf 等格式  
[c.PureJSON 和 c.JSON 的区别?](https://github.com/gin-gonic/gin#purejson)

#### ➤ 文件

返回文件: [①serving static files](https://github.com/gin-gonic/gin#serving-static-files)、[②serving data from file](https://github.com/gin-gonic/gin#serving-data-from-file)  
[返回 Reader 中的数据](https://github.com/gin-gonic/gin#serving-data-from-reader),  比如用 http.Get 请求某个网址得到一个 Reader,  然后返回 Reader 中的数据

#### [➤ 渲染 HTML 模板](https://github.com/gin-gonic/gin#html-rendering)  
#### [➤ 返回重定向](https://github.com/gin-gonic/gin#redirects)  
#### [➤ 读写 Cookie](https://github.com/gin-gonic/gin#set-and-get-a-cookie)

### 配置 HTTP Server

- [修改 HTTP 服务器设置](https://github.com/gin-gonic/gin#custom-http-configuration)
- [配置 HTTPS](https://github.com/gin-gonic/gin#support-lets-encrypt),  但如果在 gin 应用前面套个 nginx 就得去 nginx 中配置 https
- Gin 应用的前面有代理服务器,  如何获取用户真实 IP?  [参考1](https://github.com/gin-gonic/gin/issues/2697)、[参考2](https://github.com/gin-gonic/gin#dont-trust-all-proxies)

### 编写测试

- [例子](https://github.com/gin-gonic/gin#testing): The `net/http/httptest` package is preferable way for HTTP testing.

### 优雅地关闭服务

#### ➤ 如何平滑关闭 Gin 应用?

[参考此处的文档](https://github.com/gin-gonic/gin#manually)

#### ➤ 为什么 Web 服务需要平滑退出?

At the moment, when we stop our API application (usually by pressing Ctrl+C) it is terminated immediately with no opportunity for in-flight HTTP requests to complete. This isn’t ideal for two reasons:

1. It means that clients won’t receive responses to their in-flight requests — all they will experience is a hard closure of the HTTP connection.
2. Any work being carried out by our handlers may be left in an incomplete state. 

#### ➤ Sending Shutdown Signals

When our application is running, we can terminate it at any time by sending it a specific signal. A common way to do this, which you’ve probably been using, is by pressing Ctrl+C on your keyboard to send an interrupt signal — also known as a `SIGINT`.

| Signal  | Description                          | Keyboard Shortcut | Catchable |
| ------- | ------------------------------------ | ----------------- | --------- |
| SIGINT  | Interrupt from keyboard              | Ctrl + C          | Yes       |
| SIGQUIT | Quit from keyboard                   | Ctrl + \          | Yes       |
| SIGKILL | Kill process (terminate immediately) | -                 | No        |
| SIGTERM | Terminate proces                     | -                 | Yes       |

Catachable signals can be intercepted by our application and either ignored, or used to trigger a certain action (such as a graceful shutdown). Other signals, like SIGKILL, are not catchable and cannot be intercepted. Go provides tools in the `os/signals` package that we can use to intercept catchable signals and trigger a graceful shutdown of our application.   

```bash
pgrep -l api          # 为了确认进程是否存在、可以搜索进程名中包含 api 的进程、用 -l 显示搜到的进程名
pkill -SIGKILL api    # 向指定进程名发送 SIGKILL 信号
```

#### ➤ Intercepting Shutdown Signals

To catch the signals, we’ll need to spin up a background goroutine which runs for the lifetime of our application. In this background goroutine, we can use the `signal.Notify()` function to listen for specific signals and relay them to a channel for further processing.  

```go
func main() {
    router := http.NewServeMux()
    router.HandleFunc("/", index)

    go func() {
        quit := make(chan os.Signal, 1)                      // 注意大小为 1, 因为发送方是非阻塞的
        signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM) // 监听这两个 signal
        s := <-quit                                          // 在此处等待,  直到收到了某个 signal
        fmt.Println("caught signal", s.String())             // 打印收到的信号
        os.Exit(0)                                           // 等待后台任务、清理资源、关闭服务...
    }()

    _ = http.ListenAndServe("127.0.0.1:3000", router)
}

func index(writer http.ResponseWriter, request *http.Request) {
    _, _ = writer.Write([]byte("hello world!"))
}
```

#### ➤ Executing the Shutdown

Intercepting the signals is all well and good, but it’s not very useful until we do something with them! Specifically, after receiving one of these signals we will call the `Shutdown()` method on our HTTP server. The official documentation describes this as follows:

> Shutdown gracefully shuts down the server without interrupting any active connections. Shutdown works by first closing all open listeners, then closing all idle connections, and then waiting indefinitely for connections to return to idle and then shut down. The `Shutdown()` method does not wait for any background tasks to complete. You will need to implement your own logic to coordinate a graceful shutdown of these things. 

```go
func main() {
    router := http.NewServeMux()
    router.HandleFunc("/index", index)
    srv := &http.Server{
        Addr:    "127.0.0.1:3000",
        Handler: router,
    }
    shutdownError := make(chan error)
    go func() {
        quit := make(chan os.Signal, 1)
        signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
        s := <-quit
        fmt.Println("caught signal", s.String())
        ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
        defer cancel()
        // Shutdown 可能会返回错误, 比如无法在 5 秒内关闭服务
        shutdownError <- srv.Shutdown(ctx)
    }()
    // 调用 Shutdown() 会让 ListenAndServe() 立即返回 http.ErrServerClosed
    err := srv.ListenAndServe()
    if !errors.Is(err, http.ErrServerClosed) {
        panic(err)
    }
    // 判断调用 Shutdown() 时有没有发生错误
    err = <-shutdownError
    if err != nil {
        panic(err)
    }
    // 执行到这说明平滑关闭了服务器
    fmt.Println("stopped server")
}

func index(writer http.ResponseWriter, request *http.Request) {
    fmt.Println("处理请求中...")
    time.Sleep(4 * time.Second) // 在处理请求时用 Ctrl+C 关闭应用,  应该能看到平滑关闭
    _, _ = writer.Write([]byte("hello world!"))
    fmt.Println("处理完成")
}
```

## 模型校验

### 概述

(1) Gin uses [go-playground/validator/v10](https://github.com/go-playground/validator) for validation  
(2) 通过 v := binding.Validator.Engine().(*validator.Validate) 获取 gin 的校验器、然后能修改配置  
(3) 如果已经学会了 validator/v10 包的使用,  [参考这里的例子](https://github.com/gin-gonic/gin#custom-validators),  为 gin 添加自定义的 validator

### 简单例子

#### ➤ [参考这个 Simple 例子](https://github.com/go-playground/validator/blob/master/_examples/simple/main.go)

(1) 使用 validator 包的流程如下

1. `validate = validator.New()`, 创建 validate 单例, 它线程安全, 且带有缓存能提高性能
2. 给结构体加上各种校验 tag
3. 用 `err := validate.Struct(user)` 看有没有校验错误,  如果有则用 err.Field() 之类的方法创建错误消息
4. 也可以对非结构体进行校验, 比如 errs := validate.Var(myEmail, "required,email") 中 myEamil 是个字符串

(2) 可用逗号连接多个校验器  
逗号表示 AND 关系: `validate:"gte=0,lte=130"`、`validate:"required,email"`  
也可以用 | 表示 OR 关系,  比如 `omitempty,rgb|rgba` 表示 omitempty && (rgb || rgba)  
多个校验器会依次执行,  比如 `validate:"max=10,min=1"` 会先执行 max 然后执行 min  
如果需要在 tag 中使用 `,` 或 `|` 那么应该用 utf-8 形式, 比如 0x2C 和 0x7C  
可以创建 alias, 比如 `validate:"iscolor"` 是 validate:"hexcolor|rgb|rgba|hsl|hsla" 的别名  

### 各种自定义

#### ➤ 自定义 Validator

```go
func customFunc(fl validator.FieldLevel) bool {
    if fl.Field().String() == "invalid" {
        return false
    }
    return true
}
// NOTES: using the same tag name as an existing function will overwrite the existing one
validate.RegisterValidation("custom tag name", customFunc)
```

注意 validator 包不提供正则匹配,  如果需要正则匹配,  官方建议添加自定义的 validator  
A regex can be used within the validator function and even be precompiled for better efficiency.  
另外有一个非标准的 validator 也需要手动注册: validate.RegisterValidation("notblank", validators.NotBlank)

#### ➤ [自定义值提取器](https://github.com/go-playground/validator/blob/master/_examples/custom/main.go)

```go
type DbBackedUser struct {
    Name sql.NullString `validate:"required"` // NullString 是个结构体,  但我们不关心它的具体结构
    Age  sql.NullInt64  `validate:"required"` // 定义一个提取器从 NullString 中提取字符串,  就能应用各种字符串校验
}
```

#### ➤ [注册结构体级别的校验器](https://github.com/go-playground/validator/blob/master/_examples/struct-level/main.go#L48)

有时候不只是对单个字段做校验,  需要联动其他字段,  比如要求 FirstName 或 LastName 至少一个不为空

#### ➤ [注册字段名提取器](https://github.com/go-playground/validator/blob/master/_examples/struct-level/main.go#L37)

err.Field() 能返回错误字段名,  但返回的是结构体的字段名(对前端无意义), 所以希望它能返回 json 字段名

#### ➤ [使用 Translator 返回错误消息](https://github.com/go-playground/validator/blob/master/_examples/translations/main.go)

自己使用 err 的 Field()、Tag()、Param()、... 创建错误消息,  也太麻烦了,  所以出现了 Translator  
可以覆盖 translator 中的错误消息, 也能为自定义的 validator 添加错误消息,  参考例子中的 translateOverride  
有时候错误消息需要直接展示给用户,  这就需要中文错误消息了, 读完链接中的例子、再读读这个 [issue](https://github.com/go-playground/validator/issues/524#issuecomment-542654236)

### 各种校验 tag

#### ➤ 列举一些校验 tag, 参考这里的 [完整 tag 列表](https://pkg.go.dev/github.com/go-playground/validator/v10#hdr-Baked_In_Validators_and_Tags)

```bash
-                      # 表示不校验某字段; 也可用于不校验 embeded struct
omitempty              # 设置了字段才校验,  判断「 设置与否 」的标准与 required 一致  
required               # 要求字段值不是零值, 比如 0(int)、""(string)、nil 都不满足 required  
isdefault              # 要求字段的值是零值, 刚好和 required 相反

unique                 # 要求 slice 没有重复的元素、map 没有重复的 value
unique=name            # 要求 struct slice 中各结构体的 name 字段各不相同

ascii                  # 要求字符串只包含 ASCII 字符 (允许空字符串)
printascii             # 只包含可打印的 ASCII 字符   (允许空字符串)
alpha                  # 只包含字母
number                 # 只包含数字
alphanum               # 只包含字母、和数字
alphaunicode           # 任意语言中的 letter, 比如 a(英文)、帅(中文)、の(日文) 都是 letter
alphanumunicode        # 同上但加上数字, 注意标点符号不算 letter 也不算 number

email                  # 要求字符串符合 email 格式
json                   # 要求字符串符合 json 格式
url                    # 要求字符串符合 url 格式
datetime=2006-01-02    # 要求字符串符合指定的日期格式,  参考 https://golang.org/pkg/time/
boolean                # 字符串值能解析成 boolean, 比如 "true"
numeric                # 字符串值能解析成十进制数, 比如 "1.23"

len=10                 # 要求字符串、列表、哈希表的长度为 10,  对于数值类型则要求值为 10  
max=10                 # 要求字符串、列表、哈希表的最大长度为 10,  或限制数值类型的最大值为 10
min=10                 # 同上但「 最小值 」
eq=10                  # 要求字符串、数值类型的值是等于 10
ne=10                  # 同上但「 不等于 」

oneof=red green        # 适用于字符串、整数类型,  要求值为给定列表中的某一个  
oneof='nice try' red   # 若某个字符串含有空格应该用单引号括起来

contains=abc           # 字符串包含 abc
containsany=abc        # 字符串包含 a 或 b 或 c
containsrune=@         # 字符串包含 @

excludes=abc           # 字符串不包含 abc
excludesall=abc        # 字符串不包含 a、不包含 b、不包含 c
excludesrune=@         # 字符串不包含 @

startswith=hello       # 以 hello 开头
endswith=goodbye       # 以 goodbye 结尾
startsnotwith=hello    # 不以 hello 开头
endsnotwith=goodbye    # 不以 goodbye 结尾


gt=10                  # 数值类型大于 10,  string/slice/map 的长度的大于 10  
gt=1h30m               # time.Duration 要长于一个小时三十分钟  
gte=10                 # 同上但「 大于或等于 」
lt=10                  # 同上但「 小于 」
lte=10                 # 同上但「 小于或等于 」
```

`dive` dive,required 是什么意思? 它表示进入一层并应用 required. 另外怎么校验 map key? [总之参考这里](https://pkg.go.dev/github.com/go-playground/validator/v10#hdr-Dive)

#### ➤ conditional required、conditional excluded 的 [例子在这](https://pkg.go.dev/github.com/go-playground/validator/v10#hdr-Required_If)

```bash
required_if=Field1 xxx        # 此字段是 required、如果 Field1 值为 xxx
required_unlesss=Field1 xxx   # 此字段是 required、除非 Field1 值为 xxx
excluded_if=Field1 xxx        # 不校验此字段, 如果 Field1 值为 xxx
excluded_unless=Field1 xxx    # 不校验此字段, 除非 Field1 值为 xxx
required_with                 # 如果 Field1 或 Field2 存在  
required_without              # 如果 Field1 或 Field2 不存在  
required_with_all             # 如果 Field1 且 Field2 都存在
required_without_all          # 如果 Field1 且 Field2 都不存在
```

#### ➤ 跨字段校验

什么是跨字段校验? 怎么用? [参考这里](https://pkg.go.dev/github.com/go-playground/validator/v10#hdr-Cross_Field_Validation)  
tag 命名规律为: 在后面加一个 field,  比如 contains=... -> containsfield=InnerStructField.Field

*eqfield=ConfirmPassword* 要求一个字段等于另一个字段,  [更多例子](https://pkg.go.dev/github.com/go-playground/validator/v10#hdr-Field_Equals_Another_Field),  另外 `nefield` 表示不等于某字段  
*gtfield=Start* 要求一个字段大于另一个字段, 另外 `gtefield` 表示大于等于, `ltfield` 表示小于  
*gtcsfield=InnerStructField.Field* 同上但能使用嵌套字段



