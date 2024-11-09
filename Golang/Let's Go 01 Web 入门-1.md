## Table of Contents
  - [Web 应用入门](#Web-%E5%BA%94%E7%94%A8%E5%85%A5%E9%97%A8)
    - [Hello World](#Hello-World)
    - [监听地址](#%E7%9B%91%E5%90%AC%E5%9C%B0%E5%9D%80)
    - [目录结构](#%E7%9B%AE%E5%BD%95%E7%BB%93%E6%9E%84)
  - [标准库的 Router](#%E6%A0%87%E5%87%86%E5%BA%93%E7%9A%84-Router)
    - [Router 的作用](#Router-%E7%9A%84%E4%BD%9C%E7%94%A8)
    - [Pattern 越长优先级越高](#Pattern-%E8%B6%8A%E9%95%BF%E4%BC%98%E5%85%88%E7%BA%A7%E8%B6%8A%E9%AB%98)
    - [Subtree Pattern 是什么](#Subtree-Pattern-%E6%98%AF%E4%BB%80%E4%B9%88)
    - [RESTful Routing 怎么搞](#RESTful-Routing-%E6%80%8E%E4%B9%88%E6%90%9E)
    - [还有其他的一些注意事项](#%E8%BF%98%E6%9C%89%E5%85%B6%E4%BB%96%E7%9A%84%E4%B8%80%E4%BA%9B%E6%B3%A8%E6%84%8F%E4%BA%8B%E9%A1%B9)
  - [第三方的 Router](#%E7%AC%AC%E4%B8%89%E6%96%B9%E7%9A%84-Router)
    - [为啥不用标准库](#%E4%B8%BA%E5%95%A5%E4%B8%8D%E7%94%A8%E6%A0%87%E5%87%86%E5%BA%93)
    - [选择一个 Router](#%E9%80%89%E6%8B%A9%E4%B8%80%E4%B8%AA-Router)
    - [httprouter 入门](#httprouter-%E5%85%A5%E9%97%A8)
    - [自定义 404 Handler](#%E8%87%AA%E5%AE%9A%E4%B9%89-404-Handler)
    - [不允许路由冲突](#%E4%B8%8D%E5%85%81%E8%AE%B8%E8%B7%AF%E7%94%B1%E5%86%B2%E7%AA%81)
  - [一点点杂谈](#%E4%B8%80%E7%82%B9%E7%82%B9%E6%9D%82%E8%B0%88)
    - [Handler 套娃](#Handler-%E5%A5%97%E5%A8%83)
    - [请求会被并发处理](#%E8%AF%B7%E6%B1%82%E4%BC%9A%E8%A2%AB%E5%B9%B6%E5%8F%91%E5%A4%84%E7%90%86)
    - [封装 helper 函数](#%E5%B0%81%E8%A3%85-helper-%E5%87%BD%E6%95%B0)
    - [什么是依赖注入](#%E4%BB%80%E4%B9%88%E6%98%AF%E4%BE%9D%E8%B5%96%E6%B3%A8%E5%85%A5)
    - [为 Handler 注入依赖](#%E4%B8%BA-Handler-%E6%B3%A8%E5%85%A5%E4%BE%9D%E8%B5%96)
    - [也可以用闭包存储依赖](#%E4%B9%9F%E5%8F%AF%E4%BB%A5%E7%94%A8%E9%97%AD%E5%8C%85%E5%AD%98%E5%82%A8%E4%BE%9D%E8%B5%96)
  - [HTML 模板入门](#HTML-%E6%A8%A1%E6%9D%BF%E5%85%A5%E9%97%A8)
    - [Template 的命名结构](#Template-%E7%9A%84%E5%91%BD%E5%90%8D%E7%BB%93%E6%9E%84)
    - [首个 HTML 模板](#%E9%A6%96%E4%B8%AA-HTML-%E6%A8%A1%E6%9D%BF)
    - [HTML 模板继承](#HTML-%E6%A8%A1%E6%9D%BF%E7%BB%A7%E6%89%BF)
    - [可重用 HTML 模板](#%E5%8F%AF%E9%87%8D%E7%94%A8-HTML-%E6%A8%A1%E6%9D%BF)
    - [定义和引用模板](#%E5%AE%9A%E4%B9%89%E5%92%8C%E5%BC%95%E7%94%A8%E6%A8%A1%E6%9D%BF)
    - [打包 HTML 模板](#%E6%89%93%E5%8C%85-HTML-%E6%A8%A1%E6%9D%BF)
  - [HTML 模板进阶](#HTML-%E6%A8%A1%E6%9D%BF%E8%BF%9B%E9%98%B6)
    - [渲染动态内容](#%E6%B8%B2%E6%9F%93%E5%8A%A8%E6%80%81%E5%86%85%E5%AE%B9)
    - [渲染多个数据](#%E6%B8%B2%E6%9F%93%E5%A4%9A%E4%B8%AA%E6%95%B0%E6%8D%AE)
    - [自动转义 HTML](#%E8%87%AA%E5%8A%A8%E8%BD%AC%E4%B9%89-HTML)
    - [建立模板缓存](#%E5%BB%BA%E7%AB%8B%E6%A8%A1%E6%9D%BF%E7%BC%93%E5%AD%98)
    - [自动注册 partials](#%E8%87%AA%E5%8A%A8%E6%B3%A8%E5%86%8C-partials)
    - [模板中的运行时异常](#%E6%A8%A1%E6%9D%BF%E4%B8%AD%E7%9A%84%E8%BF%90%E8%A1%8C%E6%97%B6%E5%BC%82%E5%B8%B8)
    - [添加通用的模板数据](#%E6%B7%BB%E5%8A%A0%E9%80%9A%E7%94%A8%E7%9A%84%E6%A8%A1%E6%9D%BF%E6%95%B0%E6%8D%AE)
    - [添加自定义模板函数](#%E6%B7%BB%E5%8A%A0%E8%87%AA%E5%AE%9A%E4%B9%89%E6%A8%A1%E6%9D%BF%E5%87%BD%E6%95%B0)
    - [模板语法 - 条件和循环](#%E6%A8%A1%E6%9D%BF%E8%AF%AD%E6%B3%95--%E6%9D%A1%E4%BB%B6%E5%92%8C%E5%BE%AA%E7%8E%AF)
    - [模板语法 - 内置的函数](#%E6%A8%A1%E6%9D%BF%E8%AF%AD%E6%B3%95--%E5%86%85%E7%BD%AE%E7%9A%84%E5%87%BD%E6%95%B0)

## Web 应用入门

### Hello World

```go
func main() {
    router := http.NewServeMux()
    router.HandleFunc("/", home)

    server := &http.Server{
        Addr:         "127.0.0.1:4000",
        Handler:      router,
        ReadTimeout:  30 * time.Second,
        WriteTimeout: 90 * time.Second,
        IdleTimeout:  120 * time.Second,
    }

    log.Print("Starting server on ", server.Addr)
    err := server.ListenAndServe()
    if !errors.Is(err, http.ErrServerClosed) {
        log.Fatal(err)
    }
}

func home(w http.ResponseWriter, r *http.Request) {
    // 访问不存在的页面时返回 404, 而不是返回主页
    if r.URL.Path != "/" {
        http.NotFound(w, r)
        return
    }
    _, _ = w.Write([]byte("Hello world!"))
}
```

### 监听地址

`server.ListenAndServe()` 实际上会调用 `net.Listen("tcp", addr)`，其中 `addr` 的格式为 `host:port`，一些注意事项:

1. `127.0.0.1:4000` ( 等价于 `localhost:4000` ) 表示仅本机可访问，不向外暴露服务
2. `:4000` ( 省略了 host, 等价于 `0.0.0.0:4000` ) 表示在所有网卡上监听 4000 端口，可通过局域网/公网访问服务
3. 省略端口号 ( `127.0.0.1:` ) 或用端口号 0 表示自动选择一个可用端口，可通过返回值的 `Addr()` 方法获取端口号

### 目录结构

It’s important to explain upfront that there’s no single right — or even recommended — way to structure web applications in Go. As a starting point, the best advice I can give you is don’t over-complicate things. Try hard to add structure and complexity only when it’s demonstrably needed.

```bash
mkdir -p cmd/web internal ui/html ui/static
touch cmd/web/main.go
touch cmd/web/handlers.go
```

- `cmd` directory will contain the application-specific code for the executable applications in the project. For now we’ll have just one executable application — the web application — which will live under the cmd/web directory.  

- `internal` directory will contain the ancillary non-application-specific code used in the project. We’ll use it to hold potentially reusable code like validation helpers and the SQL database models for the project.
- `ui` directory will contain the user-interface assets used by the web application. Specifically, the ui/html directory will contain HTML templates, and the ui/static directory will contain static files (like CSS and images).

It scales really nicely if you want to add another executable application to your project. For example, you might want to add a CLI (Command Line Interface) to automate some administrative tasks in the future. With this structure, you could create this CLI application under cmd/cli and it will be able to import and reuse all the code you’ve written under the internal directory.

It’s important to point out that the directory name `internal` carries a special meaning and behavior in Go: any packages which live under this directory can only be imported by code inside the parent of the `internal` directory. In our case, this means that any packages which live in `internal` can only be imported by code inside our snippetbox project directory. 

Or, looking at it the other way, this means that any packages under `internal` cannot be imported by code outside of our project. This is useful because it prevents other codebases from importing and relying on the (potentially unversioned and unsupported) packages in our `internal` directory — even if the project code is publicly available somewhere like GitHub. 别人不能依赖这些东西，所以我能随意改。

## 标准库的 Router

### Router 的作用

When our server receives a new HTTP request it calls the servemux’s `ServeHTTP()` method. This looks up the relevant handler based on the request URL path, and in turn calls that handler’s `ServeHTTP()` method. You can think of a Go web application as a chain of `ServeHTTP()` methods being called one after another.

总之 Server 收到的所有请求都会交给 ServeMux (aka router)，然后 router 根据 URL 再把请求交给对应的 Handler

### Pattern 越长优先级越高

- 虽然 `"/"` 能匹配所有 URL,  比如 "/"、"/home"、"/other"，但因为它最短所以它的优先级也最低
- Go matches patterns based on length, with longer patterns taking precedence over shorter ones. So you can register patterns in any order and it won’t change how the servemux behaves.
- 访问不存在的路径也返回 home 页面会比较奇怪，所以要在 handler 中

```go
func home(w http.ResponseWriter, r *http.Request) {
    // 访问不存在的页面时返回 404, 而不是返回主页
    if r.URL.Path != "/" {
        http.NotFound(w, r)
        return
    }
    _, _ = w.Write([]byte("Hello world!"))
}
```

### Subtree Pattern 是什么

Go’s servemux supports two different types of URL patterns: fixed paths and subtree paths. Fixed paths don’t end with a trailing slash, whereas subtree paths do end with a trailing slash. `"/snippet"` and `"/snippet/create"` are both examples of fixed paths. Fixed path patterns are only matched when the request URL path exactly matches the fixed path.

In contrast, our pattern `"/"` is an example of a subtree path (because it ends in a trailing slash). Another example would be something like `"/static/"`. Subtree path patterns are matched whenever the start of a request URL path matches the subtree path. If it helps your understanding, you can think of subtree paths as acting a bit like they have a wildcard at the end, like `/**` or `/static/**`.

### RESTful Routing 怎么搞

It’s important to acknowledge that the routing functionality provided by Go’s servemux is pretty lightweight. 

- It doesn’t support routing based on the request method, 
- it doesn’t support semantic URLs with variables in them, 
- and it doesn’t support regexp-based patterns. 

If you have a background in using frameworks like Rails, Django or Laravel you might find this a bit restrictive. But don’t let that put you off. The reality is that Go’s servemux can still get you quite far, and for many applications is perfectly sufficient. For the times that you need more, there’s a huge choice of third-party routers that you can use instead of Go’s servemux.

### 还有其他的一些注意事项

#### ➤ 别用 DefaultServeMux

`http.Handle()` and `http.HandleFunc()` functions. These allow you to register routes without declaring a servemux. Behind the scenes, these functions register their routes with something called the DefaultServeMux. There’s nothing special about this — it’s just regular servemux like `DefaultServeMux := NewServeMux()`

Although this approach can make your code slightly shorter, I don’t recommend it for production applications. Because DefaultServeMux is a global variable, any package can access it and register a route — 容易引起冲突

#### ➤ 自动 URL 清洗和重定向

- If the request path contains any `.` or `..` elements or repeated slashes, the user will automatically be redirected to an equivalent clean URL. For example, if a user makes a request to `/foo/bar/..//baz` they will automatically be sent a 301 Permanent Redirect to `/foo/baz` instead.  

- 如果注册了 `/foo/` 但没有注册 `/foo`，那么当用户访问 `/foo` 时，会被 301 重定向到 `/foo/`

#### ➤ Host Name Matching

It’s possible to include host names in your URL patterns. This can be useful if your application is acting as the back end for multiple sites or services. For example:

```go
// 注意 HTTP 请求中有个 Host 请求头
// 同一服务器可以收到两个 URL 相同但 Host 不同的请求, 例如 foo.com/home 和 bar.com/home
mux := http.NewServeMux()
mux.HandleFunc("foo.example.org/", fooHandler)
mux.HandleFunc("bar.example.org/", barHandler)
mux.HandleFunc("/baz", bazHandler)
```

When it comes to pattern matching, any host-specific patterns will be checked first and if there is a match the request will be dispatched to the corresponding handler. Only when there isn’t a host-specific match found will the non-host specific patterns also be checked.

## 第三方的 Router

### 为啥不用标准库

Go’s servemux doesn’t support method based routing or clean URLs with variables in them. [There are some tricks](https://youtu.be/yi5A3cK1LNA?t=778) you can use to get around this, but most people tend to decide that it’s easier to reach for a third-party package to help with routing. ( **标准库也很好，一些项目如果用标准库就能满足需求，那就没必要加入三方库** )

### 选择一个 Router

- `julienschmidt/httprouter` is the most focused, lightweight and fastest of the three packages, and is about as close to ‘perfect’ as any third-party router gets ( 最近乎完美的第三方路由器, as as any/ever 表示最高级 ) in terms of its compliance with the HTTP specs ( 在兼容 HTTP 规范方面 ). It automatically handles `OPTIONS` requests and sends 405 responses correctly, and allows you to set custom handlers for 404 and 405 responses too.
- `go-chi/chi` is generally similar to `httprouter` in terms of its features, with the main differences being that it also supports regexp route patterns and ‘grouping’ of routes which use specific middleware. This route grouping feature is really valuable in larger applications where you have lots routes and middleware to manage. ( `Gin` 使用自定义的 `httprouter` 所以也支持路由分组 )
- `gorilla/mux` is the most full-featured of the three routers. It supports regexp route patterns, and allows you to route requests based on scheme, host and headers. It’s also the only one to support custom routing rules and route ‘reversing’ (like you get in Django, Rails or Laravel).
- In our case, our application is fairly small and we don’t need support for anything beyond basic method-based routing and clean URLs. So, for the sake of performance and correctness, we’ll opt to use `julienschmidt/httprouter` in this project.

### httprouter 入门

- Install the latest version of httprouter like so: `go get github.com/julienschmidt/httprouter@v1`. 

- The second argument of `router.HandlerFunc()` is the pattern that the request URL path must match. A request with a URL path like `/snippet/view/123` or `/snippet/view/foo` would match our example pattern `/snippet/view/:id`, but a request for `/snippet/view/` or `/snippet/view/foo/baz` wouldn’t.
- Patterns can also include a single catch-all parameter in the form `*name`. These match everything and should be used at the end of a pattern, like as `/static/*filepath`.
- The pattern `"/"` will only match requests where the URL path is exactly `"/"`.

```go
func routes(app *application) http.Handler {
    router := httprouter.New()
    router.HandlerFunc(http.MethodGet, "/", app.homePage)
    router.HandlerFunc(http.MethodGet, "/snippet/view/:id", app.snippetView)
    router.HandlerFunc(http.MethodGet, "/snippet/create", app.snippetCreate)
    router.HandlerFunc(http.MethodPost, "/snippet/create", app.snippetCreatePost)

    // 创建 File Server, 所有以 /static/ 开头的 URL 都交给它处理
    // 把请求交给 File Server 处理前, 先去掉请求 URL 中的 /static 前缀
    fileServer := http.FileServer(http.Dir("./web/ui/static/"))
    router.Handler(http.MethodGet, "/static/*filepath", http.StripPrefix("/static", fileServer))

    // 创建 standard middleware chain, 并把它用于每一个请求
    standard := alice.New(app.recoverPanic, app.logRequest, secureHeaders)
    return standard.Then(router)
}

func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {
    // 这样取出名为 id 的 URL parameter, 并校验 id 是数字且为正数
    params := httprouter.ParamsFromContext(r.Context())
    id, err := strconv.Atoi(params.ByName("id"))
    if err != nil || id < 1 {
        app.notFound(w)
        return
    }
}
```

### 自定义 404 Handler

```go
func routes(app *application) http.Handler {
    router := httprouter.New()
    // 自定义 404 Handler, 统一用 app.NotFound() 处理
    router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        app.notFound(w)
    })
}
```

### 不允许路由冲突

It’s important to be aware that `httprouter` doesn’t allow conflicting route patterns which potentially match the same request. So, for example, you cannot register a route like `GET /foo/new` and another route with a named parameter segment or catch-all parameter that conflicts with it — like `GET /foo/:name` or GET `/foo/*name`.

In most cases this is a positive thing. Because conflicting routes aren’t allowed, there are no routing-priority rules that you need to worry about, and it reduces the risk of bugs and unintended behavior in your application.  

But if you do need to support conflicting routes (for example, you might need to replicate the endpoints of an existing application exactly for backwards-compatibility), then I would recommend using `chi` or `gorilla/mux` instead — both of which do permit conflicting routes.

[关于路由冲突, 参考 httprouter 作者给出的答疑](https://github.com/julienschmidt/httprouter/issues/175)

```go
func TestHttpRouter(t *testing.T) {
    router := httprouter.New()
    // 这两个路由会冲突, 因为 /i/love/homura 可以同时匹配两个路由
    router.HandlerFunc(http.MethodGet, "/i/love/:name", nil)
    router.HandlerFunc(http.MethodGet, "/i/love/homura", nil)

    // 想注册如下路由, 理论上它应该不和 /i/love/:name 产生冲突
    // 但 httprouter 使用前缀树实现路由匹配, 这个路由的 /i/love/homura 前缀会和 /i/love/:name 冲突
    router.HandlerFunc(http.MethodGet, "/i/love/homura/and/hikari", nil)

    // 解决办法是这样注册路由, 然后在 Handler 中检查 :name 是否为 homura
    router.HandlerFunc(http.MethodGet, "/i/love/:name/and/hikari", func(w http.ResponseWriter, r *http.Request) {
        params := httprouter.ParamsFromContext(r.Context())
        if params.ByName("name") != "homura" {
            http.NotFound(w, r)
            return
        }
    })
}
```

## 一点点杂谈

### Handler 套娃

The servemux also has a `ServeHTTP()` method, meaning that it too satisfies the `http.Handler` interface. For me it simplifies things to think of the servemux as just being a special kind of handler, which instead of providing a response itself passes the request on to a second handler. 

In fact, what exactly is happening is this: When our server receives a new HTTP request, it calls the servemux’s `ServeHTTP()` method. This looks up the relevant handler based on the request URL path, and in turn calls that handler’s `ServeHTTP()` method. You can think of a Go web application as a chain of `ServeHTTP()` methods being called one after another.

### 请求会被并发处理

There is one more thing that’s really important to point out: all incoming HTTP requests are served in their own goroutine. For busy servers, this means it’s very likely that the code in or called by your handlers will be running concurrently. While this helps make Go blazingly fast, the downside is that you need to be aware of (and protect against) race conditions when accessing shared resources from your handlers.

### 封装 helper 函数

Let’s neaten up our application by moving some of the error handling code into helper methods. This will help separate our concerns and stop us repeating code as we progress through the build.

```go
// serverError 打印错误日志和堆栈，并返回 500 响应
func (app *application) serverError(w http.ResponseWriter, err error) {
    trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
    _ = app.errorLog.Output(2, trace) // 在日志中记录这一行没意义，应该记录调用 app.serverError 那一行
    http.Error(w, http.StatusText(500), 500)
}

// clientError 向用户返回错误消息
func (app *application) clientError(w http.ResponseWriter, status int) {
    http.Error(w, http.StatusText(status), status)
}

func (app *application) notFound(w http.ResponseWriter) {
    app.clientError(w, http.StatusNotFound)
}
```

 效果如下，能减少一些重复:

```go
    // before
    if err != nil {
        app.errorLog.Print(err.Error())
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        return
    }
    // after
    if err != nil {
        app.serverError(w, err)
        return
    }
```

### 什么是依赖注入

- 如果 A 组件依赖 B 组件，那么可以在 A 组件中创建 B 组件，但一般不这么做，
- 而是 A 组件预留一个用于注入依赖的方法，靠依赖注入器把 B 组件传给 A 组件。
- 比如通过构造函数、Setter 方法把一个组件传给另一个组件，就是一种简单的依赖注入。

如何为 http.Handler 注入依赖?

```go
// 如果用结构体实现 http.Handler，那么把依赖存到结构体字段就行
// 如果利用闭包存储依赖，函数也能支持依赖注入，就像这样:
mux.HandleFunc("/", Home(app))
func Home(app *application) http.HandlerFunc {
   return func(w http.ResponseWriter, r *http.Request) {
      // 因为闭包，即使被返回，这个匿名函数也能访问 app 变量
      app.infoLog.Println("i need more power !!!")
   }
}
```

### 为 Handler 注入依赖

This raises a good question: how can we make our new `errorLog` logger available to our `home` function? And this question generalizes further. Most web applications will have multiple dependencies that their handlers need to access, such as a database connection pool, centralized error handlers, and template caches. What we really want to answer is: how can we make any dependency available to our handlers?

There are [a few different ways](https://www.alexedwards.net/blog/organising-database-access) to do this, the simplest being to just put the dependencies in global variables. But in general, it is good practice to inject dependencies into your handlers. It makes your code more explicit, less error-prone and easier to unit test than if you use global variables.

For applications where all your handlers are in the same package, like ours, a neat way to inject dependencies is to put them into a custom `application` struct, and then define your handler functions as methods against `application`.

```go
// 让 Handler 作为 application 的方法，就能访问所有依赖
type application struct {
    errorLog *log.Logger
    infoLog  *log.Logger
}
func main() {
    app := &application{
        errorLog: errorLog,
        infoLog:  infoLog,
    }
    router.HandleFunc("/", app.home)
}
func (app *application) home(w http.ResponseWriter, r *http.Request) {
    if err != nil {
        app.errorLog.Print(err.Error()) // 因为是方法所以能访问 app.errorLog
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        return
    }
}
```

### 也可以用闭包存储依赖

The pattern that we’re using to inject dependencies won’t work if your handlers are spread across multiple packages. In that case, an alternative approach is to create a `config` package exporting an `Application` struct and have your handler functions close over this to form a closure. Very roughly:

```go
func main() {
    app := &config.Application{
        ErrorLog: log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile),
    }
    mux.Handle("/", examplePackage.ExampleHandler(app))
}

func ExampleHandler(app *config.Application) http.HandlerFunc {
    // 用闭包存储 app 这个依赖
    return func(w http.ResponseWriter, r *http.Request) {
        ts, err := template.ParseFiles(files...)
        if err != nil {
            app.ErrorLog.Print(err.Error())
            http.Error(w, "Internal Server Error", 500)
            return
        }
    }
}
```

## HTML 模板入门

### Template 的命名结构

```go
func TestTemplate1(t *testing.T) {
    result := new(strings.Builder)
    ts := template.New("base")                                 // 创建一个名为 base 的空模板
    ts.Parse(`{{template "boy"}} love {{template "girl"}} ! `) // 使用 Parse 赋予 base 模板内容 ( 这里调用了其他模板 )

    ts.New("boy").Parse("i")       // Template 对象实际上是一个集合, 可以用 ts.New 添加更多模板
    ts.New("girl").Parse("homura") // 这里添加 boy 和 girl 两个模板, 供 base 模板调用

    ts.ExecuteTemplate(result, "base", nil) // 执行 template set 中的 base 模板
    ts.ExecuteTemplate(result, "girl", nil) // 执行 template set 中的 girl 模板
    fmt.Println(ts.DefinedTemplates())      // base boy girl
    fmt.Println(result, "!")

}

func TestTemplate2(t *testing.T) {
    // 添加 home.tmpl 模板, 但在它的内容中又定义了 home 模板, 所以会添加 home.tmpl 和 home 共两个模板
    ts, _ := template.New("home.tmpl").Parse(`12345 {{define "home"}}Home Page{{end}}`)
    fmt.Println(ts.DefinedTemplates())
    ts.ExecuteTemplate(os.Stdout, "home.tmpl", nil) // 只打印 12345, {{define}} 中的东西不算可显示内容
    ts.ExecuteTemplate(os.Stdout, "home", nil)      // 只打印 Home Page, 因为 home 模板的内容只包含这个
    fmt.Println()
}

func TestTemplate3(t *testing.T) {
    // ParseFiles() 和 ParseGlob() 会对每个文件调用 t.New(filename).Parse(content)
    // 然后 t.New(同一个名字) 会覆盖之前的内容, 所以 ParseFiles() 中的同名文件只有最后一个生效
    ts := template.Must(template.New("").Funcs(template.FuncMap{}).ParseGlob("templates/*"))
    t.Log(ts.DefinedTemplates())
}
```

### 首个 HTML 模板

Let’s start by creating a template file at ui/html/pages/home.tmpl

```bash
mkdir ui/html/pages             # 创建目录
touch ui/html/pages/home.tmpl   # 首个 html 模板
```

```html
<!doctype html>
<html lang='en'>
<head>
    <meta charset='utf-8'>
    <title>Home - Snippetbox</title>
</head>
<body>
<header>
    <h1><a href='/'>Snippetbox</a></h1>
</header>
<main>
    <h2>Latest Snippets</h2>
    <p>There's nothing to see here yet!</p>
</main>
<footer>Powered by <a href='https://golang.org/'>Go</a></footer>
</body>
</html>
```

```go
func testHTMLTemplate(w http.ResponseWriter, r *http.Request) {
    // Use the template.ParseFiles() function to read the template file into a template set.
    ts, err := template.ParseFiles("./web/home.tmpl")
    if err != nil {
        log.Print(err.Error())
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        return
    }
    // We then use the Execute() method on the template set to write the template content as the response body. The
    // last parameter to Execute() represents any dynamic data that we want to pass in, which for now we'll leave as nil
    err = ts.Execute(w, nil)
    if err != nil {
        log.Print(err.Error())
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        return
    }
}
```

### HTML 模板继承

As we add more pages to this web application there will be some shared, boilerplate, HTML markup that we want to include on every page — like the header, navigation and metadata inside the `<head>` HTML element. To save us typing and prevent duplication, it’s a good idea to create a base template which contains this shared content, which we can then compose with the page-specific markup for the individual pages.

```bash
touch ui/html/base.tmpl # 创建 base 模板
```

Here we’re using the `{{define "base"}}...{{end}}` action to define a distinct named template called base, which contains the content we want to appear on every page. Inside this we use the `{{template "title" .}}` and `{{template "main" .}}` actions to denote that we want to invoke other named templates (called `title` and `main`) at a particular point in the HTML. If you’re wondering, the dot at the end of the `{{template "title" .}}` action represents any dynamic data that you want to pass to the invoked template. ( `.` 表示全都传过去，`.Posts` 表示只传 `Posts` 字段过去 )

```html
{{define "base"}}
    <!doctype html>
    <html lang='en'>
    <head>
        <meta charset='utf-8'>
        <title>{{template "title" .}} - Snippetbox</title>
    </head>
    <body>
    <header>
        <h1><a href='/'>Snippetbox</a></h1>
    </header>
    <main>
        {{template "main" .}}
    </main>
    <footer>Powered by <a href='https://golang.org/'>Go</a></footer>
    </body>
    </html>
{{end}}
```

Now let’s go back to the ui/html/pages/home.tmpl file and update it to define `title` and `main` named templates containing the specific content for the home page.

```html
{{define "title"}}Home{{end}}

{{define "main"}}
    <h2>Latest Snippets</h2>
    <p>There's nothing to see here yet!</p>
{{end}}
```

The next step is to update the code in your handler so that it parses both template files. So now, our template set contains 3 named templates — `base`, `title` and `main`. We use the `ExecuteTemplate()` method to tell Go that we specifically want to respond using the content of the base template (which in turn invokes our `title` and `main` templates).

```go
func tesTemplateInheritance(w http.ResponseWriter, r *http.Request) {
    // Initialize a slice containing the paths to the two files. It's  important
    // to note that the file containing our base template must be the *first* file in the slice.
    files := []string{
        "./web/base.tmpl",
        "./web/home.tmpl",
    }
    // Use the template.ParseFiles() function to read the files and store the templates in a template set.
    ts, err := template.ParseFiles(files...)
    if err != nil {
        log.Print(err.Error())
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        return
    }
    // Use the ExecuteTemplate() method to write the content of the "base" template as the response body.
    err = ts.ExecuteTemplate(w, "base", nil)
    if err != nil {
        log.Print(err.Error())
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        return
    }
}
```

### 可重用 HTML 模板

For some applications you might want to break out certain bits of HTML into partials that can be reused in different pages or layouts. To illustrate, let’s create a partial containing the primary navigation bar for our web application.

Create a new ui/html/partials/nav.tmpl file containing a named template called "nav", like so:

```bash
mkdir ui/html/partials            # 创建 partials 目录
touch ui/html/partials/nav.tmpl   # 首个经常复用的模板
```

```html
{{define "nav"}}
    <nav>
        <a href='/'>Home</a>
    </nav>
{{end}}
```

Then update the `base` template so that it invokes the navigation partial using the `{{template "nav" .}}` action:

```html
    <!-- Invoke the navigation template -->
    {{template "nav" .}}
```

Finally, we need to update the handler to include the new ui/html/partials/nav.tmpl file when parsing the template files:

```go
    files := []string{
        "./web/base.tmpl",
        "./web/nav.tmpl",
        "./web/home.tmpl",
    }
```

### 定义和引用模板

使用 `{{define "base"}}...{{end}}` 定义的模板，可通过两种方式进行引用。

In the code above we’ve used the `{{template}}` action to invoke one template from another. But Go also provides a `{{block}}...{{end}}` action which you can use instead. This acts like the `{{template}}` action, except it allows you to specify some default content if the template being invoked doesn’t exist in the current template set.

In the context of a web application, this is useful when you want to provide some default content (such as a sidebar) which individual pages can override on a case-by-case basis if they need to. 

But — if you want — you don’t need to include any default content between the `{{block}}` and `{{end}}` actions. In that case, the invoked template acts like it’s ‘optional’. If the template exists in the template set, then it will be rendered. But if it doesn’t, then nothing will be displayed.

Syntactically you use it like this:

```html
{{define "base"}}
    <h1>An example template</h1>
    {{block "sidebar" .}}
        <p>My default sidebar content</p>
    {{end}}
{{end}}
```

### 打包 HTML 模板

Go also provides the [embed](https://pkg.go.dev/embed/) package which makes it possible to embed files into your Go program itself rather than reading them from disk.

## HTML 模板进阶

### 渲染动态内容

```go
func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {
    // 校验 id 是数字且为正数
    id, err := strconv.Atoi(r.URL.Query().Get("id"))
    if err != nil || id < 1 {
        app.notFound(w)
        return
    }
    // err 不为 nil 时可能是没有数据, 也可能是其他错误
    snippet, err := app.snippets.Get(id)
    if err != nil {
        if errors.Is(err, models.ErrNoRecord) {
            app.notFound(w)
        } else {
            app.serverError(w, err)
        }
        return
    }
    files := []string{
        "./web/ui/html/base.tmpl",
        "./web/ui/html/nav.tmpl",
        "./web/ui/html/pages/view.tmpl",
    }
    ts, err := template.ParseFiles(files...)
    if err != nil {
        app.serverError(w, err)
        return
    }
    // 这里要用 base 模板, base 是 HTML 的框架, base 会引用 main, 不同的页面只有 main 模板内容不同
    // ExecuteTemplate 的最后一个参数是交给模板渲染的动态数据, 把数据放到自定义的 templateData 结构里
    data := &templateData{
        Snippet: snippet,
    }
    err = ts.ExecuteTemplate(w, "base", data)
    if err != nil {
        app.serverError(w, err)
        return
    }
}
```

Next up we need to create the `view.tmpl` file containing the HTML markup for the page. Within your HTML templates, any dynamic data that you pass in is represented by the `.` character (referred to as dot). In this specific case, the underlying type of dot will be a `templateData` struct. We could yield the snippet title by writing `{{.Snippet.Title}} `in our templates.

```html
{{define "title"}}Snippet #{{.Snippet.ID}}{{end}}
{{define "main"}}
    <div class='snippet'>
        <div class='metadata'>
            <strong>{{.Snippet.Title}}</strong>
            <span>#{{.Snippet.ID}}</span>
        </div>
        <pre><code>{{.Snippet.Content}}</code></pre>
        <div class='metadata'>
            <time>Created: {{.Snippet.Created}}</time>
            <time>Expires: {{.Snippet.Expires}}</time>
        </div>
    </div>
{{end}}
```

### 渲染多个数据

An important thing to explain is that Go’s `html/template` package allows you to pass in one — and only one — item of dynamic data when rendering a template. A lightweight and type-safe way to achieve this is to wrap your dynamic data in a struct which acts like a single "holding structure" for your data. Let’s create a new `cmd/web/templates.go` file, containing a `templateData` struct to do exactly that.

```go
// 想让模板渲染的动态数据放到这个结构里面
type templateData struct {
    Snippet  *models.Snippet   // 详情页用到这个字段
    Snippets []*models.Snippet // 列表页用到这个字段
}
```

### 自动转义 HTML

The `html/template` package automatically escapes any data that is yielded between `{{ }}` tags. This behavior is hugely helpful in avoiding cross-site scripting (XSS) attacks, and is the reason that you should use the `html/template` package instead of the more generic `text/template` package that Go also provides. As an example of escaping:

```html
<!-- 使用 {{ }} 渲染这样一段内容 -->
{{"<script>alert('xss attack')</script>"}}
<!-- 经过 {{ }} 的自动转义, 内容会变成下面这样, 总之内容变成了纯文本而不是 HTML 源码 -->
&lt;script&gt;alert(&#39;xss attack&#39;)&lt;/script&gt;
```

The `html/template` package is also smart enough to make escaping context-dependent. It will use the appropriate escape sequences depending on whether the data is rendered in a part of the page that contains HTML, CSS, Javascript or a URI. Finally, the `html/template` package always strips out any HTML comments you include in your templates, including any conditional comments.

### 建立模板缓存

使用 map 作为模板缓存:

```go
func newTemplateCache() (map[string]*template.Template, error) {
    // 用 map 作为缓存
    cache := make(map[string]*template.Template)

    // 使用 filepath.Glob 获取以 .tmpl 结尾的文件路径, 返回值是 []string 类型
    pages, err := filepath.Glob("./web/ui/html/pages/*.tmpl")
    if err != nil {
        return nil, err
    }

    for _, page := range pages {
        // 从文件路径提取文件名, 比如 ./web/ui/html/pages/home.tmpl 变成 home.tmpl
        name := filepath.Base(page)

        // base + nav + page 构成一个完整的页面
        files := []string{
            "./web/ui/html/base.tmpl",
            "./web/ui/html/nav.tmpl",
            page,
        }

        // 用 home.tmpl 这样的文件名作为 key 在缓存中存储解析好的模板
        ts, err := template.ParseFiles(files...)
        if err != nil {
            return nil, err
        }

        cache[name] = ts
    }
    return cache, nil
}
```

在 main 函数中初始化:

```go
func main() {
    templateCache, err := newTemplateCache()
    if err != nil {
        errorLog.Fatal(err)
    }
    app := &application{
        templateCache: templateCache,
    }
}
```

在 handler 中使用模板缓存:

```go
// render 是渲染模板的辅助函数
func (app *application) render(w http.ResponseWriter, status int, page string, data *templateData) {
    ts, ok := app.templateCache[page]
    if !ok {
        err := fmt.Errorf("cannot find template: %s", page)
        app.serverError(w, err)
        return
    }
    w.WriteHeader(status)
    err := ts.ExecuteTemplate(w, "base", data) // 注意这里要用 base, base 是框架
    if err != nil {
        app.serverError(w, err)
        return
    }
}

func (app *application) homePage(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/" {
        app.notFound(w)
        return
    }

    snippets, err := app.snippets.Latest()
    if err != nil {
        app.serverError(w, err)
        return
    }

    // 非常地简洁易用
    app.render(w, http.StatusOK, "home.tmpl", &templateData{
        Snippets: snippets,
    })
}
```

### 自动注册 partials

在如上的模板缓存中，`nav.tmpl` 是写死的，如果以后加了 `footer.tmpl` 或 `comment.tmpl` 等页面组件，就得手动注册一下，所以可以改成自动注册 partials 文件夹中的页面组件:

```go
func newTemplateCache() (map[string]*template.Template, error) {
    cache := make(map[string]*template.Template)
    pages, err := filepath.Glob("./web/ui/html/pages/*.tmpl")
    if err != nil {
        return nil, err
    }

    for _, page := range pages {
        name := filepath.Base(page)

        // 先用 base.tmpl 创建 Template 对象
        ts, err := template.ParseFiles("./web/ui/html/base.tmpl")
        if err != nil {
            return nil, err
        }

        // 再往 Template 对象中添加更多内容, 也就是 partials 文件夹中的页面组件
        ts, err = ts.ParseGlob("./web/ui/html/partials/*.tmpl")
        if err != nil {
            return nil, err
        }

        // 最后添加页面相关的模板内容
        ts, err = ts.ParseFiles(page)
        if err != nil {
            return nil, err
        }

        cache[name] = ts
    }
    return cache, nil
}
```

### 模板中的运行时异常

如果渲染模板是发生了运行时异常，比如在某个位置添加一行 `{{len nil}}`，那么用户会看到渲染了一半的模板。

This is pretty bad. Our application has thrown an error, but the user has wrongly been sent a `200 OK` response. And even worse, they’ve received a half-complete HTML page. To fix this we need to make the template render a two-stage process. First, we should make a ‘trial’ render by writing the template into a buffer. If this fails, we can respond to the user with an error message. But if it works, we can then write the contents of the buffer to our `http.ResponseWriter`.

Let’s update the `render()` helper to use this approach instead:

```go
func (app *application) render(w http.ResponseWriter, status int, page string, data *templateData) {
    ts, ok := app.templateCache[page]
    if !ok {
        err := fmt.Errorf("cannot find template: %s", page)
        app.serverError(w, err)
        return
    }

    // 初始化 buffer, 对于不想拷贝的类型用 new(Type) 或 &Type{} 初始化比较好, 从源头杜绝拷贝结构体
    buf := new(bytes.Buffer)

    // 把模板渲染到 buffer 而不是 http.ResponseWriter, 这样才有机会检查, 渲染模板时是否发生运行时错误
    err := ts.ExecuteTemplate(buf, "base", data)
    if err != nil {
        app.serverError(w, err)
        return
    }

    // 如果没有发生错误, 再把 buffer 中的内容返回给用户
    w.WriteHeader(status)
    _, _ = buf.WriteTo(w)
}
```

### 添加通用的模板数据

In some web applications there may be common dynamic data that you want to include on more than one — or even every — webpage. For example, you might want to include the name and profile picture of the current user, or a CSRF token in all pages with forms.

In our case let’s begin with something simple, and say that we want to include the current year in the footer on every page. To do this we’ll begin by adding a new `CurrentYear` field to the `templateData` struct, like so:  

```go
type templateData struct {
    CurrentYear int               // 模板通用数据
    Snippet     *models.Snippet   // 详情页用到这个字段
    Snippets    []*models.Snippet // 列表页用到这个字段
}
```

The next step is to add a `newTemplateData()` helper method to our application, which will return a `templateData` struct initialized with the current year.  

```go
// 通过这个构造函数就能添加通用的模板数据
func (app *application) newTemplateData(r *http.Request) *templateData {
    return &templateData{
        CurrentYear: time.Now().Year(),
    }
}
```

Then let’s update our `home` and `snippetView` handlers to use the `newTemplateData()` helper, like so:

```go
func (app *application) homePage(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/" {
        app.notFound(w)
        return
    }
    snippets, err := app.snippets.Latest()
    if err != nil {
        app.serverError(w, err)
        return
    }

    // 初始化模板数据
    data := app.newTemplateData(r)
    data.Snippets = snippets
    app.render(w, http.StatusOK, "home.tmpl", data)
}
```

And then the final thing we need to do is update the `ui/html/base.tmpl` file to display the year in the footer:

```html
<footer>Powered by <a href='https://golang.org/'>Go</a> in {{.CurrentYear}}</footer>
```

### 添加自定义模板函数

During execution functions are found in two function maps: first in the template, then in the global function map. By default, no functions are defined in the template but the `Funcs` method can be used to add them. Let’s create a custom `humanDate()` function which outputs datetimes in a nice ‘humanized’ format like `02 Jan 2022 at 15:04`, instead of outputting dates in the default format of `2022-01-02 15:04:00 +0000 UTC` like we are currently. There are two main steps to doing this:

1. We need to create a `template.FuncMap` object containing the custom `humanDate()` function.
2. We need to use the `template.Funcs()` method to register this before parsing the templates.

```go
// template 函数只允许有一个返回值, 或一个返回值加一个 error
func humanDate(t time.Time) string {
    return t.Format(time.DateTime)
}

var templateFunctions = template.FuncMap{
    "humanDate": humanDate,
}

func main() {
    // 必须在 ParseFiles 解析模板内容之前把自定义函数注册上
    ts := template.New(name).Funcs(templateFunctions)
    ts, err := ts.ParseFiles(files...)
}
```

Now we can use our `humanDate()` function in the same way as the built-in template functions:

```html
<!-- 模板中使用自定义函数 -->
<td>{{humanDate .Created}}</td>
```

### 模板语法 - 条件和循环

- [详情可参考官方文档](https://pkg.go.dev/text/template#hdr-Actions)
- 注意 if、with、range 中的 else 块都是可选的，range 的对象必须是 array/slice/map/channel
- 注意 with 和 range 会修改句点 `.` 的含义，但不会修改它们的 else 分支中句点 `.` 的含义

> 注意 empty 的定义: The `{{if ...}}` action considers empty values (false, 0, any nil pointer or interface value, and any array, slice, map, or string of length zero) to be false.

```bash
{{if .Foo}} C1 {{else}} C2 {{end}}      # 如果 .Foo 不为空则展示 C1 否则展示 C2
{{with .Foo}} C1 {{else}} C2 {{end}}    # 如果 .Foo 不为空则展示 C1 并且把 . 设置为 .Foo, 否则展示 C2 并且不修改 .
{{range .Foo}} C1 {{else}} C2 {{end}}   # 如果 .Foo 长度大于 0 则展示 C1 并且把 . 设置为 .Foo 中的每一个元素
```

`with` 示例:

```html
{{with .Snippet}}
    <!-- 在这个上下文里 . 表示 .Snippet, 所以直接用 .Title 取字段 -->
    <span>{{.Title}}</span>
    <span>{{.Content}}</span>
{{end}}
```

`if-else` 和 `range` 示例:

```html
{{define "main"}}
    <h2>Latest Snippets</h2>
    <!-- 如果 .Snippets 不为空则渲染列表 -->
    {{if .Snippets}}
        <table>
            <tr>
                <th>Title</th>
                <th>Created</th>
                <th>ID</th>
            </tr>
            <!-- 如果 .Snippets 不为空则进行循环渲染 -->
            {{range .Snippets}}
                <!-- 在这个上下文中 . 表示 array/slice/map/channel 中的一个元素 -->
                <tr>
                    <td><a href='/snippet/view?id={{.ID}}'>{{.Title}}</a></td>
                    <td>{{.Created}}</td>
                    <td>#{{.ID}}</td>
                </tr>
            {{end}}
        </table>
    <!-- 否则展示展示一段文本 -->
    {{else}}
        <p>There's nothing to see here... yet!</p>
    {{end}}
{{end}}
```

`range` 配合 `break` 或 `continue` 示例:

```html
{{range .Foo}}
    <!-- 如果 .ID == 99 则用 continue 跳过这一元素, 直接进行下一轮循环 -->
    {{if eq .ID 99}}
        {{continue}}
    {{end}}
    <!-- 此处是循环渲染的内容 -->
    {{.ID}}
{{end}}

{{range .Foo}}
    <!-- 如果 .ID == 99 则用 break 跳出循环 -->
    {{if eq .ID 99}}
        {{break}}
    {{end}}
    <!-- 此处是循环渲染的内容 -->
    {{.ID}}
{{end}}
```

### 模板语法 - 内置的函数

The `html/template` package also provides some template functions which you can use to add extra logic to your templates and control what is rendered at runtime. You can find a complete listing of functions [here](https://pkg.go.dev/text/template/#hdr-Functions), but the most important ones are:

```bash
{{eq .Foo .Bar}}                # 计算 .Foo == .Bar, 结果为 true/false
{{if eq .Foo .Bar}}...{{end}}   # 可以和 if 连用
{{ne .Foo .Bar}}                # 计算 .Foo != .Bar, 结果为 true/false
{{not .Foo}}                    # 计算 !.Foo, 结果为 true/false
{{or .Foo .Bar}}                # 计算 .Foo || .Bar, 如果 .Foo 非空则返回 .Foo 否则返回 .Bar
{{index .Foo i}}                # 返回 .Foo 中索引为 i 的对象, 这里 .Foo 必须是 map/slice/array
{{printf "%s-%s" .Foo .Bar}}    # 格式化打印字符串, 类似于 fmt.Sprintf
{{len .Foo}}                    # 返回 .Foo 的长度
{{$bar := len .Foo}}            # 计算 .Foo 的长度并存储到 $bar 变量, 模板中的变量必须以 $ 开头
```

#### ➤ 模板语法 - 函数嵌套

It’s possible to combine multiple functions in your template tags, using the parentheses `()` to surround the functions and their arguments as necessary. For example, the following tag will render the content `C1` if the length of `Foo` is greater than 99:

```bash
{{if (gt (len .Foo) 99)}} C1 {{end}}               # 相当于 if ( len(.Foo) > 99 )
{{if (and (eq .Foo 1) (le .Bar 20))}} C1 {{end}}   # 相当于 if ( (.Foo == 1) && (.Bar <= 20) )
```

#### ➤ 模板语法 - 调用方法

If the type that you’re yielding between `{{ }}` tags has methods defined against it, you can call these methods (so long as they are exported and they return only a single value — or a single value and an error).

For example, if `.Snippet.Created` has the underlying type `time.Time` (which it does) you could render the name of the weekday by calling its `Weekday()` method like so:

```html
<span>{{.Snippet.Created.Weekday}}</span>
```

You can also pass parameters to methods. For example, you could use the `AddDate()` method to add six months to a time like so:

```html
<!-- 用一个空格而不是逗号分隔参数 -->
<span>{{.Snippet.Created.AddDate 0 6 0}}</span>
```

#### ➤ 模板语法 - Pipelining

In the code above, we called our custom template function like this:  

```html
<time>Created: {{humanDate .Created}}</time>
```

An alternative approach is to use the `|` character to pipeline values to a function.

```html
<time>Created: {{.Created | humanDate}}</time>
```

A nice feature of pipelining is that you can make an arbitrarily long chain of template functions which use the output from one as the input for the next.  

```html
<time>{{.Created | humanDate | printf "Created: %s"}}</time>
```

#### ➤ 建议看看官方文档: [template package - text/template](https://pkg.go.dev/text/template)

