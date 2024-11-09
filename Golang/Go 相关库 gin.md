## Table of Contents
  - [Gin](#Gin)
    - [参考](#%E5%8F%82%E8%80%83)
    - [为什么用 Gin](#%E4%B8%BA%E4%BB%80%E4%B9%88%E7%94%A8-Gin)
    - [相关配置](#%E7%9B%B8%E5%85%B3%E9%85%8D%E7%BD%AE)
  - [路由](#%E8%B7%AF%E7%94%B1)
    - [入门](#%E5%85%A5%E9%97%A8)
    - [路径参数](#%E8%B7%AF%E5%BE%84%E5%8F%82%E6%95%B0)
    - [查询参数](#%E6%9F%A5%E8%AF%A2%E5%8F%82%E6%95%B0)
    - [路由冲突](#%E8%B7%AF%E7%94%B1%E5%86%B2%E7%AA%81)
    - [路由分组](#%E8%B7%AF%E7%94%B1%E5%88%86%E7%BB%84)
  - [中间件](#%E4%B8%AD%E9%97%B4%E4%BB%B6)
    - [概述](#%E6%A6%82%E8%BF%B0)
    - [使用入门](#%E4%BD%BF%E7%94%A8%E5%85%A5%E9%97%A8)
    - [常用中间件](#%E5%B8%B8%E7%94%A8%E4%B8%AD%E9%97%B4%E4%BB%B6)
    - [并发注意事项](#%E5%B9%B6%E5%8F%91%E6%B3%A8%E6%84%8F%E4%BA%8B%E9%A1%B9)
  - [模型绑定](#%E6%A8%A1%E5%9E%8B%E7%BB%91%E5%AE%9A)
    - [入门示例](#%E5%85%A5%E9%97%A8%E7%A4%BA%E4%BE%8B)
    - [各种绑定方法](#%E5%90%84%E7%A7%8D%E7%BB%91%E5%AE%9A%E6%96%B9%E6%B3%95)
    - [表单注意事项](#%E8%A1%A8%E5%8D%95%E6%B3%A8%E6%84%8F%E4%BA%8B%E9%A1%B9)
  - [返回响应](#%E8%BF%94%E5%9B%9E%E5%93%8D%E5%BA%94)
    - [渲染 JSON](#%E6%B8%B2%E6%9F%93-JSON)
    - [静态文件](#%E9%9D%99%E6%80%81%E6%96%87%E4%BB%B6)
    - [渲染模板](#%E6%B8%B2%E6%9F%93%E6%A8%A1%E6%9D%BF)
    - [其他响应](#%E5%85%B6%E4%BB%96%E5%93%8D%E5%BA%94)
  - [数据校验](#%E6%95%B0%E6%8D%AE%E6%A0%A1%E9%AA%8C)
    - [入门例子](#%E5%85%A5%E9%97%A8%E4%BE%8B%E5%AD%90)
    - [自定义校验器](#%E8%87%AA%E5%AE%9A%E4%B9%89%E6%A0%A1%E9%AA%8C%E5%99%A8)
    - [各种自定义](#%E5%90%84%E7%A7%8D%E8%87%AA%E5%AE%9A%E4%B9%89)
    - [包内置的校验 tag](#%E5%8C%85%E5%86%85%E7%BD%AE%E7%9A%84%E6%A0%A1%E9%AA%8C-tag)
  - [编写测试](#%E7%BC%96%E5%86%99%E6%B5%8B%E8%AF%95)
    - [单元测试](#%E5%8D%95%E5%85%83%E6%B5%8B%E8%AF%95)
    - [集成测试](#%E9%9B%86%E6%88%90%E6%B5%8B%E8%AF%95)
  - [Server 相关](#Server-%E7%9B%B8%E5%85%B3)
    - [HTTP Server](#HTTP-Server)
    - [支持 HTTPS](#%E6%94%AF%E6%8C%81-HTTPS)
    - [获取客户端 IP](#%E8%8E%B7%E5%8F%96%E5%AE%A2%E6%88%B7%E7%AB%AF-IP)
    - [优雅地关闭服务](#%E4%BC%98%E9%9B%85%E5%9C%B0%E5%85%B3%E9%97%AD%E6%9C%8D%E5%8A%A1)
  - [各种示例](#%E5%90%84%E7%A7%8D%E7%A4%BA%E4%BE%8B)
    - [上传文件](#%E4%B8%8A%E4%BC%A0%E6%96%87%E4%BB%B6)
    - [日志相关](#%E6%97%A5%E5%BF%97%E7%9B%B8%E5%85%B3)
    - [HTTP2](#HTTP2)

## Gin

### 参考

- [官方文档](https://github.com/gin-gonic/gin/blob/master/docs/doc.md)

### 为什么用 Gin

Why gin? ①标准库功能不够强大 ②用框架更方便、提高生产力 ③目前流行度最高的是 gin，[star 数超第二名一倍](https://github.com/mingrammer/go-web-framework-stars)

While the `net/http` package allows you to craft a web application easily, the routing mechanism is not so powerful, especially for complex applications. That's where a web framework comes into play. 

Gin is a minimalistic web framework suitable for building web applications, microservices, and RESTful APIs. It reduces boilerplate code by creating reusable and extensible pieces of code: you can write a piece of middleware that can be plugged into one or more request handlers. Moreover, it comes with the following key features:  

- *Well documented*: The documentation for Gin is broad and comprehensive. Most tasks that you will need to do relating to the router can be found easily in the docs.
- *Simplicity*: Gin is a pretty minimalistic framework. Only the most essential features and libraries are included, with little to no boilerplate to bootstrap applications, making Gin a great framework for developing highly available REST APIs.
- *Extensible*: The Gin community has created numerous pieces of well-tested middleware that make developing for Gin a charm. Features include compression with GZip, authentication with an authorization middleware, and logging with external solutions such as Sentry.
- *Performance*: Gin runs 40x faster than Martini and runs comparatively well compared to other Golang frameworks.

### 相关配置

#### ➤ 可以切换 json 库，比如 [jsoniter](https://github.com/json-iterator/go) 和 [go-json](https://github.com/goccy/go-json)，然后 [sonic](https://github.com/bytedance/sonic/blob/main/README_ZH_CN.md) 说自己是最快的

```bash
# 没有啥要求就用标准库，最好根据自己的场景做测试
# 按照社区活跃度依次推荐 sonic、go-json、jsoniter
go build -tags="sonic avx" . # you have to ensure that your cpu support avx instruction.
go build -tags=go_json .
go build -tags=jsoniter .
```

#### ➤ 可以关掉 [MessagePack](https://msgpack.org/) 格式，能减少可执行文件的体积

Gin enables `MsgPack` rendering feature by default. But you can disable this feature by specifying `nomsgpack` build tag. This is useful to reduce the binary size of executable files. See the [detail information](https://github.com/gin-gonic/gin/pull/1852).

```bash
go build -tags=nomsgpack .
```





## 路由
### 入门

#### ➤ Basic Route Definition

In the world of web development, routing is a fundamental concept that allows us to define how our application responds to different HTTP requests. Routing in Gin is straightforward and intuitive, making it easy for developers to map URLs to specific handler functions. Let's look at a simple example:

```go
func TestGin(t *testing.T) {
    router := gin.Default()

    router.GET("/hello", func(c *gin.Context) {
        c.String(http.StatusOK, "Hello, World!")
    })

    router.Run(":8080")
}
```

#### ➤ Handler Functions

Handler functions in Gin receive a `*gin.Context` parameter, which provides access to the request details and allows you to send responses. These functions can be defined inline or as separate named functions. Here's an example of a named handler function:

```go
func TestGin(t *testing.T) {
    router := gin.Default()
    router.GET("/users", getUsers)
    router.Run(":8080")
}

func getUsers(c *gin.Context) {
    users := []string{"Alice", "Bob", "Charlie"}
    c.JSON(http.StatusOK, gin.H{"users": users})
}
```

#### ➤ Multiple Handlers

Gin allows you to assign multiple handler functions to a single route. This is useful for implementing middleware or breaking down complex logic into smaller, reusable functions. In this example, the "/protected" route will first run the `authenticate` function, then the `authorize` function, and finally the `handleRequest` function.

```go
func TestGin(t *testing.T) {
    gin.SetMode(gin.ReleaseMode) // 不要 Debug 模式的各种提示, 也可以设为 gin.TestMode
    router := gin.New()          // 不要 Logger, Recovery 等中间件, 方便看打印的消息

    router.GET("/protected", authenticate, authorize, handleRequest) // 执行顺序为从左到右

    r, _ := httptest.NewRequest("GET", "/protected", nil)
    w := httptest.NewRecorder()
    router.ServeHTTP(w, r)
}

func getUsers(c *gin.Context) {
    users := []string{"Alice", "Bob", "Charlie"}
    c.JSON(http.StatusOK, gin.H{"users": users})
}

func authenticate(c *gin.Context) { fmt.Println("Authentication logic") }

func authorize(c *gin.Context) { fmt.Println("Authorization logic") }

func handleRequest(c *gin.Context) { fmt.Println("Main request handling logic") }
```

#### ➤ NoRoute() and NoMethod() Handlers

Gin allows you to define custom handlers for requests that don't match any defined routes or use unsupported HTTP methods:

```go
func TestGin(t *testing.T) {
    // 如果 url 存在但请求方法不对, 默认会返回 404 响应
    // 除非把 router.HandleMethodNotAllowed 设为 true
    router := gin.Default()
    router.HandleMethodNotAllowed = true // 请求方法不对时返回 405 和响应头 Allow: POST, PUT

    router.NoRoute(func(c *gin.Context) {
        // 默认返回 404 page not found 网页, 这里改成返回 JSON
        c.JSON(http.StatusNotFound, gin.H{"message": "Page not found"})
    })

    router.NoMethod(func(c *gin.Context) {
        // 默认返回 405 method not allowed 网页, 这里改成返回 JSON
        c.JSON(http.StatusMethodNotAllowed, gin.H{"message": "Method not allowed"})
    })
}
```



### 路径参数

#### ➤ Colon-prefixed parameters

Path parameters, also known as URL parameters or route parameters, are variables embedded in the URL path itself. They allow you to capture specific parts of the URL and use them as input in your application logic. This feature is particularly useful when you need to handle requests for resources with varying identifiers, such as user profiles, product details, or blog posts.

In Gin, path parameters are defined using a colon (:) followed by the parameter name in the route definition. When a request matches the route pattern, Gin extracts the values from the URL and makes them available to your handler function. In this example, `:id` is a path parameter. When a request is made to a URL like `/users/123`, Gin will extract the string value `"123"` and make it available through the `c.Param("id")` method.

```go
router.GET("/users/:id", func (c *gin.Context) {
    id := c.Param("id")
    c.String(http.StatusOK, "User ID: %s", id)
})
```

#### ➤ Wildcards in routes

Gin supports wildcard parameters in routes, which can match any number of path segments:

```go
func TestGin(t *testing.T) {
    gin.SetMode(gin.ReleaseMode)
    router := gin.New()
    router.GET("/files/*path", func(c *gin.Context) {
        t.Log(c.Param("path"))
    })

    w := httptest.NewRecorder()
    router.ServeHTTP(w, newRequest("GET", "/files/"))          // 打印 /
    router.ServeHTTP(w, newRequest("GET", "/files/img/x.png")) // 打印 /img/x.png

    // 通配符只能出现在末尾, 放中间会报错:
    // catch-all routes are only allowed at the end of the path
    router.GET("/a/*b/c", func(c *gin.Context) {})
}

func newRequest(method, url string) *http.Request {
    r, _ := httptest.NewRequest(method, url, nil)
    return r
}
```

#### ➤ 注意事项

- `:param` 对应的部分必须存在，并且斜杠敏感，例如 `/users/:id` 不会匹配 */users/* 和 */users/123/*
- `*param` 对应的部分可有可无，返回值以斜杠开头，例如 `/files/*path` 能匹配 */files/* 且返回 `"/"`

- 当没有其他路由能匹配 /home 时，/home 会被 301 重定向到 /home/，反之亦然  
  所以 `/users/:id` 在触发重定向机制时，能匹配 */users/123/*

```go
func TestGin(t *testing.T) {
    router := gin.Default()
    router.GET("/home", func(c *gin.Context) {
        c.String(http.StatusOK, "url: "+c.Request.URL.Path)
    })
    // 注意 /home 和 /home/ 是两个路由
    // 但如果没有其他路由能成功匹配, 那么对 /home/ 的访问会被重定向到 /home
    router.Run(":8080")
}
```

### 查询参数

#### ➤ 入门

Query parameters are an essential part of web applications, allowing clients to send additional data to the server through the URL. Query parameters are key-value pairs appended to the end of a URL. They follow the question mark (?) in the URL and are separated by ampersands (&). For example:

```bash
https://example.com/api/users?page=1&limit=10
```

Query parameters can be accessed using the `Query()` method of the `gin.Context`:

```go
func TestGin(t *testing.T) {
    router := gin.Default()
    router.GET("/search", func(c *gin.Context) {

        t.Log(c.Query("q"))                 // 对于 /search?q=golang 会返回 golang
        t.Log(c.DefaultQuery("q", "value")) // 当 url 中不存在 q 时 DefaultQuery() 返回默认值
        t.Log(c.GetQuery("q"))              // 第二个返回值表示 q 是否存在, 例如 /search?q 会返回 "", true

        // 可能是 /search 或 /search?q 或 /search?q= 或 search?id=123
        if c.Query("q") == "" {
        }

        // 多次使用同一参数例如 /search?tag=golang&tag=web 会形成数组 tag = ["golang", "web"]
        t.Log(c.Query("tag"), c.QueryArray("tag")) // 注意 Query() 只返回首个元素
        
        // 还可以使用 map 语法例如:
        //     /search?user[name]=homura&user[age]=18
        // 注意 array 和 map 语法由后端框架定义, 不同框架的语法不同
        c.JSON(http.StatusOK, c.QueryMap("user"))
    })
    router.Run(":8080")
}
```

#### ➤ 表单也是这 7 个函数

```go
c.PostForm("msg")                      // 例如 msg=love&msg=you 返回首个即 love
c.PostFormArray("tags")                // 返回 slice
c.PostFormMap("user")                  // 返回 map
c.DefaultPostForm("name", "anonymous") // name 不存在时使用默认值, 注意不存在和空字符串是两个概念

c.GetPostForm("msg") // 第二个返回值表示是否存在, 注意 name=tifa&msg 会返回 "", true
c.GetPostFormArray()
c.GetPostFormMap()
```

#### ➤ 把 query string 解析到结构体

```go
func TestQuery(t *testing.T) {
    router := gin.Default()
    router.GET("/search", func(c *gin.Context) {
        query := struct {
            Keyword string   `form:"q"`   // 使用 form 标签把参数 q 解析到 Keyword 字段
            Tag     []string `form:"tag"` // 把参数 tag 解析为字符串切片
        }{}
        // 使用 c.ShouldBindQuery 并处理错误
        if err := c.ShouldBindQuery(&query); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }
        c.JSON(http.StatusOK, query)
    })
    router.Run(":8080")
}
```

#### ➤ 加上数据校验

```go
func TestQuery(t *testing.T) {
    router := gin.Default()
    router.GET("/search", func(c *gin.Context) {
        query := struct {
            Keyword string   `form:"q" binding:"required"` // 必须提供参数 q, 且不能为空字符串 ""
            Tag     []string `form:"tag" binding:"min=2"`  // 要求切片长度 >=2, 所以必须提供两个 tag 参数
            ID      int      `form:"id" binding:"max=3"`   // 限制最大值为 3, 所以字段的零值 0 也满足
        }{}
        if err := c.ShouldBindQuery(&query); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }
        c.JSON(http.StatusOK, query)
    })
    router.Run(":8080")
}
```



### 路由冲突

- Exact routes are resolved before param routes, regardless of the order they were defined.
- Routes starting with /user/groups are never interpreted as /user/:name/... routes

#### ➤ 写死的前缀越长优先级越高

```go
func TestGin(t *testing.T) {
    router := gin.Default()
    printRoute := func(c *gin.Context) {
        c.String(http.StatusOK, c.FullPath()) // 使用 c.FullPath() 获取当前是哪条路由
    }

    // 写死的前缀越长优先级越高, 此处 /users/123 比 /users/ 更长
    router.GET("/users/123", printRoute)
    router.GET("/users/:id", printRoute)

    // 如果第一部分相等, 那就继续比下一步部分, 此处 123 比 :target 更优先
    router.GET("/:action/123", printRoute)
    router.GET("/:action/:target", printRoute)

    // 通配符是前缀匹配, 所有以 /files/ 开头的路由都与这个通配符路由冲突
    router.GET("/files/*path")
    // router.GET("/files/:path") // 冲突
    // router.GET("/files/x.png") // 冲突

    router.Run(":8080")
}
```

#### ➤ 如果 path paramter 要包含斜杠该怎么办? ( 当然最好是避开这种奇怪的设计，可以放 query string 里 )

```go
func TestGin(t *testing.T) {
    gin.SetMode(gin.ReleaseMode)
    router := gin.New()
    // RawPath 可以用来区分 / 和 %2F, 参考 https://youtu.be/mzJCXk6H8I8?t=289
    // The RawPath field is an optional field which is only set when the
    // default encoding of Path is different from the escaped path. ( 两者相同则不设置 RawPath )
    router.UseRawPath = true
    router.UnescapePathValues = false

    router.GET("/users/:id/haha", func(c *gin.Context) {
        t.Log("Route  ", c.FullPath())
        t.Log("Path   ", c.Request.URL.Path)
        t.Log("RawPath", c.Request.URL.RawPath)
        t.Log("id =   ", c.Param("id"))
    })

    w := httptest.NewRecorder()
    r, _ := httptest.NewRequest("GET", "/users/21%2Fadmin%2F527/haha", nil)
    router.ServeHTTP(w, r) // 打印  21%2Fadmin%2F527
}
```

### 路由分组

Gin provides a powerful feature called route groups, which allow you to organize related routes and apply common middleware or prefixes to a set of routes.

```go
    v1 := router.Group("/v1")
    {
        v1.GET("/users", getUsers)       // 等价于 GET /v1/users
        v1.POST("/users", createUser)    // 以后也可以建个 /v2 组
    }

    // 这样就设置了一组应用了 auth 中间件、需要认证的路由, 挺方便的啊
    authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))
    authRoutes.POST("/accounts", server.createAccount)
    authRoutes.GET("/accounts/:id", server.getAccount)

    // Group 可以嵌套, 加大括号只是为了格式更美观
    api := router.Group("/api")
    {
        users := api.Group("/users")
        {
            users.GET("/:id", nil) // 等价 GET /api/users/:id
            users.POST("", nil)    // 等价 POST /api/users
        }

    }
```

#### ➤ 路径拼接使用了 path.Join 所以双斜杠也没问题，详见 gin.joinPaths

```go
func TestGin(t *testing.T) {
    router := gin.Default()
    slash := router.Group("/")
    slash.GET("/")                 // joinPaths("/", "/")      == "/"
    slash.GET("/home")             // joinPaths("/", "/home")  == "/home"
    users := slash.Group("/users") // joinPaths("/", "/users") == "/users"
    {
        users.POST("add")       // joinPaths("/users", "add") == "/users/add"
        users.POST("/")         // joinPaths("/users", "/")   == "/users/"
        users.POST("")          // joinPaths("/users", "")    == "/users"
        users.POST("/activate") // 一般推荐用 / 开头可以看上去比较统一
    }
    router.Run(":8080")
}
```





## 中间件

### 概述

#### ➤ 中间件就是在处理请求前，或生成响应后，做点什么

Middleware is a piece of code that sits between the incoming request and the final handler in a web application. It intercepts the request, performs some operations, and then passes control to the next middleware or the final handler. Middleware can also modify the response before it's sent back to the client.

Think of middleware as a series of layers that a request must pass through before reaching its final destination. Each layer can perform various tasks, such as: logging, authentication, error handling, request parsing, response compression, CORS, rate limiting.

#### ➤ 从左往右执行下一个，然后从右往左返回上一个

Let's break down the process:

- **Request Arrival**
- **Middleware Execution**: Each middleware function in the chain is executed in order. Each function can:
  - Perform operations before passing control to the next middleware
  - Modify the request or context
  - Abort the request processing
  - Call the next middleware in the chain
- **Handler Execution**: The final handler processes the request and generates a response.
- **Response Processing**: The control flow then travels back through the middleware chain in reverse order, allowing middleware to perform additional operations or modifications.
- **Response Sent**

### 使用入门

#### ➤ 这个 c.Next() 和 c.Abort() 的作用

```go
func MyMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        // Do something before the request is handled
        fmt.Println("Before request")

        // 条件不满足, 停止往后执行, 提前返回响应
        if !isAuthorized(c) {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
            return
        }

        // 执行下一个中间件或 handler
        // 如果仅需要在处理请求前做点什么, 不需要在之后做什么, 那么不调 c.Next() 也行
        c.Next()

        // Do something after the request is handled
        fmt.Println("After request")
    }
}
```

#### ➤ 可以往 gin.Context 添加数据，方便后续使用

```go
func UserIDMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        userID := extractUserID(c)
        c.Set("userID", userID) // 后面的中间件或 handler 可用 c.Get("userID") 取出数据
        c.Next()
    }
}
```

#### ➤ 按照应用范围，中间件可以分成三类

```go
func TestGin(t *testing.T) {
    // 全局中间件
    router := gin.New()
    router.Use(gin.Logger())

    // 应用于某个组的中间件
    v1 := router.Group("/v1")
    v1.Use(AuthMiddleware())

    // 应用于特定路由的中间件
    router.GET("/protected", AuthMiddleware(), protectedHandler)
}
```
#### ➤ New() 与 Default() 的区别是 Logger 和 Recovery 中间件

```go
func TestGin(t *testing.T) {
    // 这一行等价于后面三行
    // Default With the Logger and Recovery middleware already attached
    r := gin.Default()

    // Blank Gin without middleware by default
    r := gin.New()

    // Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
    // By default, gin.DefaultWriter = os.Stdout
    r.Use(gin.Logger())

    // Recovery middleware recovers from any panics and writes a 500 if there was one.
    r.Use(gin.Recovery())
}
```

### 常用中间件

#### ➤ Custom Recovery behavior

```go
func TestGin(t *testing.T) {
    router := gin.New()
    router.Use(gin.Logger())

    router.Use(gin.CustomRecovery(func(c *gin.Context, recovered any) {
        // 自定义 Recovery 中间件返回 JSON 格式的错误消息
        // 这里 recovered 一定不是 nil, 因为 gin.CustomRecoveryWithWriter() 检查过了
        c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("%v", recovered)})
    }))

    router.GET("/", func(c *gin.Context) {
        // 遇到 panic 时 Recovery 中间件做了这些事:
        // (1)把请求头、错误消息、堆栈信息写入日志 (2)调用上面注册的 handler
        panic(errors.New("foo"))
    })

    router.Run(":8080")
}
```



- 一些常用的中间件已经有人写好了，参考
  - [gin-contrib repositories](https://github.com/orgs/gin-contrib/repositories?type=all&q=sort%3Astars)
  - [Collection of middlewares created by the community](https://github.com/gin-gonic/contrib)

- [gin-contrib/cors](https://github.com/gin-contrib/cors) ( 比如前端项目运行在 3000 端口，后端服务运行在 8000 端口，就需要 CORS )
- [gin-contrib/sessions](https://github.com/gin-contrib/sessions) ( 适合 Gin 的 session 中间件，有 1.4k 颗星 )

### 并发注意事项

When starting new Goroutines inside a middleware or handler,  
you **SHOULD NOT** use the original context inside it, you have to use a read-only copy.

```go
func TestGin(t *testing.T) {
    router := gin.Default()

    router.GET("/long_async", func(c *gin.Context) {
        // create copy to be used inside the goroutine
        cCp := c.Copy()
        go func() {
            time.Sleep(5 * time.Second)
            // note that you are using the copied context "cCp", IMPORTANT!
            log.Println("Done! in path " + cCp.Request.URL.Path)
        }()
    })

    router.GET("/long_sync", func(c *gin.Context) {
        time.Sleep(5 * time.Second)
        // since we are NOT using a goroutine, we do not have to copy the context
        log.Println("Done! in path " + c.Request.URL.Path)
    })

    router.Run(":8080")
}
```

## 模型绑定

### 入门示例

```go
func TestGin(t *testing.T) {
    type Req struct {
        ID        int    `uri:"id"`
        Name      string `form:"name"`
        Age       int    `form:"age"`
        Sort      string `form:"sort"`
        UserAgent string `header:"User-Agent"`
    }

    router := gin.Default()
    router.POST("/question/:id", func(c *gin.Context) {
        var r Req
        _ = c.ShouldBindQuery(&r)                  // 只绑定 query string
        _ = c.ShouldBindWith(&r, binding.FormPost) // 只绑定 body 中的表单
        _ = c.ShouldBindWith(&r, binding.Form)     // 同时绑定 body 表单和 query string
        _ = c.ShouldBindUri(&r)                    // 只绑定 path parameter
        _ = c.ShouldBindHeader(&r)                 // 只绑定 header
        c.JSON(http.StatusOK, r)
    })
    router.Run(":8080")
}
```

可执行如下命令做测试:

```bash
curl -X POST 'http://192.168.2.11:8080/question/123?sort=like' \
 -H "Content-Type: application/x-www-form-urlencoded" \
 -H "User-Agent: Chrome" \
 -d "name=Tifa&age=20"
```

### 各种绑定方法

#### ➤ 不是 Should 开头的就是 Must Bind

例如: Bind, BindJSON, BindQuery, BindHeader, BindUri

These methods use `MustBindWith` under the hood. If there is a binding error, the request is aborted with `c.AbortWithError(400, err).SetType(ErrorTypeBind)`. This sets the response status code to 400 and the `Content-Type` header is set to `text/plain; charset=utf-8`. Note that if you try to set the response code after this, it will result in a warning `[GIN-debug] [WARNING] Headers were already written. Wanted to override status code 400 with 422`. If you wish to have greater control over the behavior, consider using the `ShouldBind` equivalent method.

#### ➤ 以 Should 开头的是 Should Bind，自己决定怎么处理错误

These methods use `ShouldBindWith` under the hood. If there is a binding error, the error is returned and it is the developer's responsibility to handle the request and error appropriately.

#### ➤ 比起 ShouldBind 的推断，用 ShouldBindJSON 更清晰

When using the Bind-method, Gin tries to infer the binder depending on the Content-Type header. If you are sure what you are binding, you can use `MustBindWith` or `ShouldBindWith`.

```go
// ShouldBindJSON is a shortcut for c.ShouldBindWith(obj, binding.JSON).
func (c *Context) ShouldBindJSON(obj any) error {
    return c.ShouldBindWith(obj, binding.JSON)
}

// Bind 方法的推断逻辑如下
func Default(method, contentType string) Binding {
    if method == http.MethodGet {
        return Form
    }
    // 根据请求头 Content-Type 返回:
    // JSON / XML / ProtoBuf / MsgPack / YAML / TOML / FormMultipart
    // 如果不是以上类型，那么默认返回 Form
}
```

#### ➤ c.ShouldBindBodyWith() 用来做什么?

```go
// ShouldBindBodyWith is similar with ShouldBindWith, but it stores the request
// body into the context, and reuse when it is called again.
// NOTE: You should use ShouldBindWith for better performance if you need to call only once.
```

注意 r.Body 只能读一次:

```go
func handler(c *gin.Context) {
    var data struct{}
    c.ShouldBindJSON(&data)         // 第一次调用会消耗 r.Body
    c.ShouldBindJSON(&data)         // 所以第二次调用会失败
    c.ShouldBindBodyWithJSON(&data) // 混用 Bind 和 BindBody 也会失败 ( 它们都需要读一次 r.Body )

    c.ShouldBindBodyWithJSON(&data) // 只用 BindBody 系列方法
    c.ShouldBindBodyWithJSON(&data) // 那么可以多次调用
}
```

### 表单注意事项

#### ➤ 不存在 c.ShouldBindForm 方法，对于表单可以用 c.ShouldBind 的推断

```go
func TestGinHandler(t *testing.T) {
    gin.SetMode(gin.TestMode)
    form := url.Values{
        "age":  {"21"},
        "name": {"cloud"},
    }
    
    // 注意 url 中设置了 name, 然后表单中设置了 age 和 name
    r := httptest.NewRequest("POST", "/?name=tifa", strings.NewReader(form.Encode()))
    r.Header.Add("Content-Type", "application/x-www-form-urlencoded")

    w := httptest.NewRecorder()
    c, _ := gin.CreateTestContext(w)
    c.Request = r

    handler(t, c)
}

func handler(t *testing.T, c *gin.Context) {
    data := struct {
        Names []string `form:"name"`
        Age   int      `form:"age"`
    }{}

    // binding.Form 内部用了 r.Form, 所以会同时绑定 query string 和 body 中的数据
    err := c.ShouldBindWith(&data, binding.Form)
    t.Log(err, data)

    // binding.FormPost 内部用了 r.PostForm, 所以只绑定 body 中的数据
    err = c.ShouldBindWith(&data, binding.FormPost)
    t.Log(err, data)

    // 如果利用 ShouldBind 的推断, 那么只会推断出 bindind.Form, 没有 binding.FormPost
    t.Log(binding.Form == binding.Default(c.Request.Method, c.ContentType()))
}
```

#### ➤ r.Form 和 r.PostForm 的区别

- `r.PostForm` contains the parsed form data from PATCH, POST or PUT body parameters. 

- `r.Form` contains both the URL field's query parameters and the PATCH, POST, or PUT form data.
- 观察 `r.ParseForm()` 源码可知，`r.Form` 是在 `r.PostForm` 的基础上再添加 url 中的数据，所以 url 的数据排在后面

```go
func TestParseForm(t *testing.T) {
    // body 和 url 都设置了 name 字段
    form := url.Values{
        "age":  {"21"},
        "name": {"cloud"},
    }
    r := httptest.NewRequest("POST", "/?name=tifa", strings.NewReader(form.Encode()))
    r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
    _ = r.ParseForm()
    t.Log("r.Form    ", r.Form)     // 注意顺序是 cloud,tifa 因为 url 数据排后面
    t.Log("r.PostForm", r.PostForm) // PostForm 只包含 request body 的数据
}
```

#### ➤ time_format 标签可以用于表单，不能用于 JSON

- `time.Time.MarshalJSON()` 限制了时间格式必须是 `time.RFC3339Nano`，所以反序列化时如果格式不符合会报错
- 注意 Gin 里面有个 [time_format](https://github.com/gin-gonic/gin/blob/master/docs/doc.md#custom-validators) 标签，这个可以用于表单，[但不能用于 JSON](https://github.com/gin-gonic/gin/issues/1193)

```go
type Person struct {
    Birthday   time.Time `form:"birthday" time_format:"2006-01-02" time_utc:"1"`
    CreateTime time.Time `form:"createTime" time_format:"unixNano"`
    UnixTime   time.Time `form:"unixTime" time_format:"unix"`
}
// 请求示例:
// curl -X GET "localhost:8085/testing?birthday=1992-03-15&createTime=1562400033000000123&unixTime=1562400033"
```

#### ➤ 绑定表单时可以提供默认值

```go
func TestDefault(t *testing.T) {
    type User struct {
        Age        int      `form:"age,default=20"`                      // 这个 default=20 竟然没有文档
        Name       string   `form:"name,default=Tifa"`                   // 然后他们在 Gin 1.11 补上了文档
        Friends    []string `form:"friends,default=Cloud;Alice"`         // 给数组设置默认值需要 Gin 1.11
        FriendsCSV []string `form:"friends_csv" collection_format:"csv"` // 让数组支持 csv 格式需要 Gin 1.11
    }

    router := gin.Default()
    router.GET("/", func(c *gin.Context) {
        var user User
        _ = c.ShouldBindQuery(&user)
        c.JSON(http.StatusOK, user)
    })
    router.Run(":8080")
}
```

#### ➤ 自定义表单绑定

```go
type Birthday string

// 实现 binding.BindUnmarshaler 接口
func (b *Birthday) UnmarshalParam(param string) error {
    *b = Birthday(strings.ReplaceAll(param, "-", "/"))
    return nil
}

func TestDefault(t *testing.T) {
    type User struct {
        Birthday Birthday `form:"birthday" binding:"required"`
    }

    router := gin.Default()
    router.GET("/", func(c *gin.Context) {
        var user User
        _ = c.ShouldBindQuery(&user)
        c.JSON(http.StatusOK, user)
    })
    router.Run(":8080") // 访问 /?birthday=2000-01-01 做测试
}
```

#### ➤ 表单也能绑定到结构体切片 ( 在 url 里传数组，每个元素还是个对象! )

```go
func TestGin(t *testing.T) {
    // 嵌套一个结构体很正常, 更厉害的是可以嵌套一个结构体切片
    type Girl struct {
        Name string `form:"name"`
    }
    type Boy struct {
        Name        string `form:"name"`
        Girlfriends []Girl `form:"girlfriends"`
    }

    router := gin.Default()
    router.GET("/", func(c *gin.Context) {
        var b Boy
        if err := c.ShouldBindQuery(&b); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }
        c.JSON(http.StatusOK, b)
    })
    // 列表中的结构体要用 JSON 表示:
    // http://localhost:8080/?name=Cloud&girlfriends={%22name%22:%22Tifa%22}&girlfriends={%22name%22:%22Aerith%22}
    router.Run(":8080")
}
```

#### ➤ 自定义表单绑定

参考 [#bind-form-data-request-with-custom-struct-and-custom-tag](https://github.com/gin-gonic/gin/blob/master/docs/doc.md#bind-form-data-request-with-custom-struct-and-custom-tag)

## 返回响应

### 渲染 JSON

#### ➤ c.JSON()、c.AsciiJSON()、c.PureJSON() 的作用

```go
func handler(c *gin.Context) {
    var user struct {
        Name string `json:"name"`
        Age  int    `json:"age"`
    }
    c.JSON(http.StatusOK, user) // 使用结构体
    c.JSON(http.StatusOK, gin.H{ // 或者手写 map
        "name": "tifa",
        "age":  20,
    })

    // Using AsciiJSON to Generates ASCII-only JSON with escaped non-ASCII characters.
    // This will output : {"lang":"GO\u8bed\u8a00","tag":"\u003cbr\u003e"}
    c.AsciiJSON(http.StatusOK, gin.H{
        "lang": "GO语言",
        "tag":  "<br>",
    })

    // Normally, JSON replaces special HTML characters with their Unicode entities, e.g. < becomes \u003c.
    // If you want to encode such characters literally, you can use PureJSON instead.
    c.PureJSON(http.StatusOK, gin.H{
        "html": "<b>Hello, world!</b>",
    })
}
```

### 静态文件

- 可能会遇到路由冲突，参考 [Wildcard route conflicts with static files · Issue #360](https://github.com/gin-gonic/gin/issues/360)
- 可以用中间件减少路由冲突: [gin-contrib/static: Static middleware](https://github.com/gin-contrib/static)
- 另外还可以让 Nginx / Caddy 处理静态文件，然后 Golang 只处理动态内容

#### ➤ router.Static() 系列方法

```go
func TestGin(t *testing.T) {
    router := gin.Default()
    router.Static("/assets", "./assets")                         // 返回此目录中的文件
    router.StaticFS("/more_static", http.Dir("my_file_system"))  // FS 版本接受 文件系统 作为入参
    router.StaticFile("/favicon.ico", "./resources/favicon.ico") // 让特定 URL 返回特定文件
    router.StaticFileFS("/more_favicon.ico", "more_favicon.ico", http.Dir("my_file_system"))

    // 标准库的 http.Dir 默认开启文件列表
    // 出于安全考量 Gin 默认不开文件列表, 但我恰好想看看磁盘上的文件列表
    router.StaticFS("/files", gin.Dir("D:\\Music", true))

    router.Run(":8080")
}
```

#### ➤ 返回文件数据

```go
func TestGin(t *testing.T) {
    // 只能返回文件, 不能干别的
    router := gin.Default()
    router.StaticFile("/img0", "D:\\Pictures\\Screenshots\\诺艾尔.jpg")

    // 返回文件作为响应, handler 中还能干点别的
    router.GET("/img1", func(c *gin.Context) {
        c.File("D:\\Pictures\\Screenshots\\特务支援科.png")
    })

    // 返回文件作为响应, 使用 fs 作为入参
    fs := http.Dir("D:\\Pictures\\Screenshots")
    router.GET("/img2", func(c *gin.Context) {
        c.FileFromFS("全家福.png", fs)
    })
    router.Run(":8080")
}
```

#### ➤ 返回 io.Reader 中的数据

```go
func TestGin(t *testing.T) {
    router := gin.Default()
    // 某种意义上的代理, 请求其他网址得到 io.Reader 然后返回其中的数据
    router.GET("/img", func(c *gin.Context) {
        resp, err := http.Get("https://i0.hdslb.com/bfs/new_dyn/6531e88855abd02e59f6731a06f2d8a434674679.png")
        if err != nil {
            c.JSON(http.StatusServiceUnavailable, gin.H{"error": err.Error()})
            return
        }
        defer resp.Body.Close()
        if resp.StatusCode != http.StatusOK {
            c.JSON(http.StatusServiceUnavailable, gin.H{"code": resp.StatusCode})
            return
        }

        reader := resp.Body
        contentLength := resp.ContentLength
        contentType := resp.Header.Get("Content-Type")
        extraHeaders := map[string]string{
            // 让文件作为附件下载, 而不是直接在网页中展示
            "Content-Disposition": `attachment; filename="downloaded image.png"`,
        }
        c.DataFromReader(http.StatusOK, contentLength, contentType, reader, extraHeaders)
    })

    router.Run(":8080")
}
```

### 渲染模板

模板内容:

```html
<!doctype html>
<html lang=en>
<head>
    <meta charset=utf-8>
    <title>{{ .title }}</title>
</head>
<body>
<p>I'm the content</p>
</body>
</html>
```

渲染模板:

```go
func TestGin(t *testing.T) {
    // 可以用 LoadHTMLGlob() / LoadHTMLFiles() / SetHTMLTemplate() 设置模板
    gin.SetMode(gin.ReleaseMode)
    router := gin.Default()
    router.LoadHTMLGlob("templates/*")
    t.Log(router.HTMLRender.(render.HTMLProduction).Template.DefinedTemplates())

    router.GET("/", func(c *gin.Context) {
        c.HTML(http.StatusOK, "index.tmpl", gin.H{
            "title": "Hello, World!",
        })
    })

    router.Run(":8080")
}
```

模板相关配置:

```go
func TestGin(t *testing.T) {
    router := gin.Default()
    ts := template.Must(template.New("").Funcs(template.FuncMap{}).ParseFiles("templates/index.tmpl"))
    router.SetHTMLTemplate(ts) // 自己提供 Template 对象

    router.Delims("{{", "}}")             // 在模板中使用 {{ }} 作为分隔符
    router.SetFuncMap(template.FuncMap{}) // 设置 FuncMap 供模板调用
    router.LoadHTMLGlob("templates/*")    // 需要一次 LoadHTML 让前两行的配置起作用

    router.GET("/", func(c *gin.Context) {
        c.HTML(http.StatusOK, "index.tmpl", gin.H{
            "title": "Hello, World!",
        })
    })

    router.Run(":8080")
}
```

如果 Gin 需要模板继承可以[看这里](https://github.com/gin-contrib/multitemplate):

```go
// 大致原理就是用 map 存储模板
// 然后一个页面的模板由 layout 和 content 构成, 分别定义在不同的文件
files := []string{"./layout.tmpl", "./about.tmpl"}

// 传给 template.New() 的参数是想要执行的命名模板, 比如 layout.tmpl
tname := filepath.Base(files[0])

templates := map[string]*template.Template{
    "about.tmpl": template.Must(template.New(tname).Funcs(template.FuncMap{}).ParseFiles(files...)),
}

templates["about.tmpl"].Execute(writer, data)
```

把模板嵌入可执行文件:

```go
//go:embed templates/* assets/tifa.jpg assets/favicon.ico
var f embed.FS

func TestGin(t *testing.T) {
    // 取 embed.FS 中的子文件夹
    assets, err := fs.Sub(f, "assets")
    if err != nil {
        log.Fatal(err)
    }

    // 从 embed.FS 加载模板
    router := gin.Default()
    ts := template.Must(template.New("").ParseFS(f, "templates/*"))
    router.SetHTMLTemplate(ts)

    router.StaticFS("/images", http.FS(assets))                            // 用 /images/tifa.jpg 访问 assets 中的文件
    router.StaticFileFS("/img/tifa.jpg", "/assets/tifa.jpg", http.FS(f))   // 从 FS 返回单个文件
    router.StaticFileFS("/favicon.ico", "/assets/favicon.ico", http.FS(f)) // 设置网站图标

    router.GET("/", func(c *gin.Context) {
        c.HTML(http.StatusOK, "index.tmpl", gin.H{"title": "Hi~"})
    })
    router.Run(":8080")
}
```

### 其他响应

#### ➤ 返回重定向

```go
func TestGin(t *testing.T) {
    router := gin.Default()
    router.GET("/tifa", func(c *gin.Context) {
        c.Redirect(http.StatusFound, "/aerith") // 302 重定向
    })
    router.GET("/aerith", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{"name": "Aerith"})
    })
    router.GET("/cloud", func(c *gin.Context) {
        c.Request.URL.Path = "/aerith"
        router.HandleContext(c) // 应用内重定向 ( 内容有变化但 url 不变 )
    })
    router.Run(":8080")
}
```

#### ➤ 读写 Cookie

```go
func TestGin(t *testing.T) {
    router := gin.Default()
    router.GET("/", func(c *gin.Context) {
        // 尝试取 cookie, 取不到就设置一下, 同一用户第二次访问能取到 cookie
        cookie, err := c.Cookie("gin_cookie")
        if err != nil {
            cookie = "NotSet"
            c.SetCookie("gin_cookie", "cookie_value", 3600, "/", "localhost", false, true)
        }
        c.JSON(http.StatusOK, gin.H{
            "gin_cookie": cookie,
        })
    })
    router.Run(":8080")
}
```

## 数据校验

### 入门例子

(1) Gin uses [go-playground/validator/v10](https://github.com/go-playground/validator) for validation

(2) validator 包默认标签名为 `validate:"required"`，但在 Gin 里面要用 `binding:"required"`

#### ➤ 单独使用 validator 包的流程如下

1. `validate = validator.New()`, 创建 validate 单例, 它线程安全
2. 给结构体加上校验 tag，比如 validate:"required"
3. 调用 `err := validate.Struct(user)` 看有没有校验错误,  如果有则用 err.Field() 之类的方法创建错误消息
4. 也可以对非结构体进行校验, 比如 errs := validate.Var(myEmail, "required,email") 中 myEamil 是个字符串

```go
var validate *validator.Validate

type User struct {
    Name           string     `validate:"required"`
    Age            uint8      `validate:"gte=0,lte=130"`
    Email          string     `validate:"required,email"`
    Gender         string     `validate:"oneof=male female prefer_not_to"`
    FavouriteColor string     `validate:"iscolor"`                // alias for 'hexcolor|rgb|rgba|hsl|hsla'
    Addresses      []*Address `validate:"required,dive,required"` // a person can have a home and cottage...
}

type Address struct {
    Street string `validate:"required"`
    City   string `validate:"required"`
    Phone  string `validate:"required"`
}

func TestValidator(t *testing.T) {
    validate = validator.New(validator.WithRequiredStructEnabled())
    validateStruct()
    validateVariable()
}

func validateStruct() {
    address := &Address{
        Street: "Eavesdown Docks",
        Phone:  "none",
    }

    user := &User{
        Name:           "Badger Smith",
        Age:            135,
        Gender:         "male",
        Email:          "Badger.Smith@gmail.com",
        FavouriteColor: "#000-",
        Addresses:      []*Address{address},
    }

    // returns nil or ValidationErrors ( []FieldError )
    err := validate.Struct(user)
    if err != nil {
        for _, err := range err.(validator.ValidationErrors) {
            fmt.Printf("%-20s %s\n", "Namespace", err.Namespace())
            fmt.Printf("%-20s %s\n", "Field", err.Field())
            fmt.Printf("%-20s %s\n", "StructNamespace", err.StructNamespace())
            fmt.Printf("%-20s %s\n", "StructField", err.StructField())
            fmt.Printf("%-20s %s\n", "Tag", err.Tag())
            fmt.Printf("%-20s %s\n", "ActualTag", err.ActualTag())
            fmt.Printf("%-20s %s\n", "Kind", err.Kind())
            fmt.Printf("%-20s %s\n", "Type", err.Type())
            fmt.Printf("%-20s %v\n", "Value", err.Value())
            fmt.Printf("%-20s %s\n", "Param", err.Param())
            fmt.Println()
        }
        // from here you can create your own error messages in whatever language you wish
        return
    }
    // save user to database
}

//goland:noinspection GoTypeAssertionOnErrors
func validateVariable() {
    myEmail := "joeybloggs.gmail.com"
    err := validate.Var(myEmail, "required,email")

    if err != nil {
        // this check is only needed when your code could produce
        // an invalid value for validation such as interface with nil
        // value most including myself do not usually have code like this.
        if _, ok := err.(*validator.InvalidValidationError); ok {
            fmt.Println(err)
            return
        }
        fmt.Println(err) // output: Key: '' Error:Field validation for '' failed on the 'email' tag
        return
    }
    // email ok, move on
}
```

#### ➤ 可用逗号连接多个校验器

逗号表示 AND 关系: `validate:"gte=0,lte=130"`、`validate:"required,email"` (中间不要加空格)  
也可以用 | 表示 OR 关系,  比如 `omitempty,rgb|rgba` 表示 omitempty && (rgb || rgba)  
多个校验器会依次执行,  比如 `validate:"max=10,min=1"` 会先执行 max 然后执行 min  
如果需要在 tag 中使用 `,` 或 `|` 那么应该用它们的 utf-8 形式, 比如 excludesall=0x2C 排除逗号  
可以创建 alias, 比如 `validate:"iscolor"` 是 validate:"hexcolor|rgb|rgba|hsl|hsla" 的别名  

#### ➤ required 怎么判定 empty value ?

> This validates that the value is not the data types default zero value. For numbers ensures value is not zero. For strings ensures value is not "". For booleans ensures value is not false. For slices, maps, pointers, interfaces, channels and functions ensures the value is not nil. For structs ensures value is not the zero value when using WithRequiredStructEnabled.
>
> WithRequiredStructEnabled enables required tag on non-pointer structs to be applied instead of ignored.

1. JSON 中提供类型的零值、或值为null、和不提供某字段, 这些都不满足 required, 注意 {id:0} 也不满足
2. 如果想区分 {id:null} 和 {id:0}, 试试指针类型 `*int` 并检查字段值是否为 `nil`

### 自定义校验器

#### ➤ 例如注册 adult 标签验证年龄 >= 18

```go
func TestCustomValidator(t *testing.T) {
    gin.SetMode(gin.TestMode)
    registerAdultValidator()

    body := strings.NewReader(`{"name": "tifa", "age": 17}`)
    c, _, _ := createTestContext(http.MethodPost, "/", body)

    func(c *gin.Context) {
        user := struct {
            Name string `binding:"required"`
            Age  int    `binding:"adult"`
        }{}
        err := c.ShouldBindJSON(&user)
        t.Log(err, user)
    }(c)
}

func createTestContext(method, url string, body io.Reader) (*gin.Context, *httptest.ResponseRecorder, *http.Request) {
    w := httptest.NewRecorder()
    r := httptest.NewRequest(method, url, body)
    c, _ := gin.CreateTestContext(w)
    c.Request = r
    return c, w, r
}

func registerAdultValidator() {
    var adult validator.Func = func(fl validator.FieldLevel) bool {
        if age, ok := fl.Field().Interface().(int); ok {
            return age >= 18
        }
        return false
    }
    // NOTES: using the same tag name as an existing function will overwrite the existing one
    if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
        err := v.RegisterValidation("adult", adult)
        if err != nil {
            panic(err) // 程序初始化阶段可以 panic
        }
    }
}
```

### 各种自定义

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

### 包内置的校验 tag

#### ➤ 正则表达式

注意 validator 包不提供正则匹配，如果需要正则匹配，官方建议添加自定义的 validator  
A regex can be used within the validator function and even be precompiled for better efficiency.  
另外有一个非标准的 validator 也需要手动注册: validate.RegisterValidation("notblank", validators.NotBlank)

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
tag 命名规律为: 在后面加一个 field,  比如 contains= -> containsfield=InnerStructField.Field

*eqfield=ConfirmPassword* 要求一个字段等于另一个字段,  [更多例子](https://pkg.go.dev/github.com/go-playground/validator/v10#hdr-Field_Equals_Another_Field),  另外 `nefield` 表示不等于某字段  
*gtfield=Start* 要求一个字段大于另一个字段, 另外 `gtefield` 表示大于等于, `ltfield` 表示小于  
*gtcsfield=InnerStructField.Field* 同上但能使用嵌套字段



## 编写测试

### 单元测试

#### ➤ 单独测某个 handler 或 middleware

```go
func TestGinHandler(t *testing.T) {
    // 避免调用 gin.New 时打印一段警告, 这一行也可以放 init() 或 TestMain() 里
    gin.SetMode(gin.TestMode)

    // 设置 ResponseWriter 和 Request
    w := httptest.NewRecorder()
    c, _ := gin.CreateTestContext(w)
    c.Request, _ = httptest.NewRequest("GET", "/", nil)

    // 想测中间件可以改成 authMiddleware()(c)
    // 当然后面检查结果的代码也得相应地改改
    indexHandler(c)

    // 检查响应/结果是否符合预期
    body, _ := io.ReadAll(w.Body)
    assert.Equal(t, "Welcome!", string(body))
    assert.Equal(t, http.StatusOK, w.Code)
}
```

### 集成测试

#### ➤ 把两个或多个组件组合在一起，看看它们之间能否配合工作

```go
func TestGinHandler(t *testing.T) {
    gin.SetMode(gin.TestMode)
    gin.DefaultErrorWriter = io.Discard

    // 把 router, middleware, handler 全组装上了
    router := gin.New()
    router.Use(gin.Recovery())
    router.GET("/", func(c *gin.Context) {
        panic("123")
    })

    // 设置 ResponseWriter 和 Request
    w := httptest.NewRecorder()
    r, _ := httptest.NewRequest("GET", "/", nil)

    router.ServeHTTP(w, r)

    // 检查响应/结果是否符合预期
    body, _ := io.ReadAll(w.Body)
    assert.Equal(t, "", string(body))
    assert.Equal(t, http.StatusInternalServerError, w.Code)
}
```

#### ➤ [集成测试和端到端测试的区别](https://www.reddit.com/r/devops/comments/xztk9a/comment/iro616y/?utm_source=share&utm_medium=web3x&utm_name=web3xcss&utm_term=1&utm_content=share_button)

> There is some overlap between the two teams, but there is a slight difference.
>
> Integration test just means you're combining two or more systems in your test. You might be testing front-end with the backend and a mocked database, or two backend services working together, etc.
>
> End-to-end tests are made from the user's point of view, and tend to be done against all components of the system, in a staging environment that is very similar to production.

## Server 相关

### HTTP Server

#### ➤ 修改 http.Server{} 结构体

```go
func TestGin(t *testing.T) {
    router := gin.Default()
    s := &http.Server{
        Addr:           ":8080",
        Handler:        router,
        ReadTimeout:    10 * time.Second,
        WriteTimeout:   10 * time.Second,
        MaxHeaderBytes: 1 << 20,
    }
    s.ListenAndServe()
}
```

#### ➤ 一个进程监听两个端口，分别干不同的事

```go
func createRouter(name string) *gin.Engine {
    router := gin.New()
    router.Use(gin.Recovery())
    router.GET("/", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{"server_name": name})
    })
    return router
}

func TestGin(t *testing.T) {
    fn := func(router *gin.Engine, addr string) func() error {
        return func() error {
            err := router.Run(addr)
            if err != nil && !errors.Is(err, http.ErrServerClosed) {
                log.Fatal(err)
            }
            return nil
        }
    }

    var group errgroup.Group
    group.Go(fn(createRouter("server cloud"), ":8080"))
    group.Go(fn(createRouter("server tifa"), ":8081"))
    if err := group.Wait(); err != nil {
        log.Fatal(err)
    }
}
```

### 支持 HTTPS

先了解一下 [ACME 是什么](https://www.youtube.com/watch?v=rIwszDULXvc)

然后 Nginx / Caddy 可以利用 ACME Client 从 Let's Encrypt 获取免费 TLS 证书

Golang 也有相关包实现了 ACME Client (比如 [autocert](https://pkg.go.dev/golang.org/x/crypto/acme/autocert))，Go 服务也可以用做 HTTPS 服务器 ([Gin配置HTTPS](https://github.com/gin-gonic/gin/blob/master/docs/doc.md#support-lets-encrypt))

但经常会在 Go 服务前放一个反向代理 (如 Caddy) 作为对外公开的 HTTPS Server，然后只在 Caddy 中配置 HTTPS


### 获取客户端 IP

#### ➤ 请求头 X-Forwarded-For

它的格式为 `X-Forwarded-For: <client>, <proxy1>, <proxy2>`

代理服务器会把它的 client ( 可能是用户也可能是另一个代理 ) 的 ip 追加到这个请求头的右侧，再转发给目标  
注意用户能随意设置 X-Forwarded-For，若直接把最左侧的地址当做用户的 ip，那么用户能[随意伪造 ip 轻松绕过限流](https://serverfault.com/a/414166)

假设 `user->proxy1->proxy2->server`，然后 proxy1 和 proxy2 都完全可信，那么用户就无法伪造 ip  
因为不管用户怎么设置，服务器只相信 proxy1 记录的 ip ( 规则是从右往左读，把首个不信任的 ip 当做用户 ip )

#### ➤ 使用 c.ClientIP() 和 router.SetTrustedProxies() 获取用户 IP

```go
func TestGin(t *testing.T) {
    router := gin.Default()
    router.SetTrustedProxies([]string{"::1"}) // 127.0.0.1 的 IPv6 写法, 适合 Caddy 和应用在同一机器上
    router.GET("/", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "1": c.Request.RemoteAddr, // 可能是用户或代理的地址, 格式为 IP:port
            "2": c.RemoteIP(),         // 从 c.Request.RemoteAddr 中分割出 IP 信息
            "3": c.ClientIP(),         // 利用 trusted proxy 和请求头 X-Forwarded-For 确定用户 IP
        })
    })
    router.Run(":8080")
}
```

### 优雅地关闭服务

#### ➤ 参考: [Graceful shutdown or restart](https://github.com/gin-gonic/gin/blob/master/docs/doc.md#graceful-shutdown-or-restart)

```go
func main() {
    router := gin.Default()
    router.GET("/", func(c *gin.Context) {
        time.Sleep(5 * time.Second)
        c.String(http.StatusOK, "Hello, World!")
    })

    srv := &http.Server{
        Addr:    ":8080",
        Handler: router,
    }

    go func() {
        if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
            log.Println(err)
        }
    }()

    // 等待退出信号
    quit := make(chan os.Signal, 1)
    signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
    <-quit
    log.Println("Shutting down server...")

    // 使用 5 秒超时调用 Shutdown
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    if err := srv.Shutdown(ctx); err != nil {
        // 未能平滑关闭服务, 可能等了 5 秒还是有连接不在空闲状态
        log.Fatal("Server forced to shutdown: ", err)
    }
    log.Println("Server exiting")
}

// 还有一种交互是第一次按 Ctrl+C 触发平滑关闭, 不想等了, 再按一次直接退出进程
// https://github.com/gin-gonic/examples/blob/master/graceful-shutdown/graceful-shutdown/notify-with-context/server.go
```

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
        // Shutdown 可能会返回错误, 比如 5 秒后还是有连接不空闲会返回超时
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

## 各种示例

### 上传文件

- 参考 [#upload-files](https://github.com/gin-gonic/gin/blob/master/docs/doc.md#upload-files)

### 日志相关

- [把日志写入文件](https://github.com/gin-gonic/gin/blob/master/docs/doc.md#how-to-write-log-file)
- [自定义日志格式](https://github.com/gin-gonic/gin/blob/master/docs/doc.md#custom-log-format)

- [某些 url 或情况下，不需要打日志](https://github.com/gin-gonic/gin/blob/master/docs/doc.md#skip-logging)
- 以 gin.Debug 模式启动时，会打印所有路由，这个格式也[可以自定义](https://github.com/gin-gonic/gin/blob/master/docs/doc.md#define-format-for-the-log-of-routes)

### HTTP2

- [http2 server push](https://github.com/gin-gonic/gin/blob/master/docs/doc.md#http2-server-push)