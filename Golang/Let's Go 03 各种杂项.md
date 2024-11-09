## Table of Contents
  - [命令行配置](#%E5%91%BD%E4%BB%A4%E8%A1%8C%E9%85%8D%E7%BD%AE)
    - [用 flag 传入配置](#%E7%94%A8-flag-%E4%BC%A0%E5%85%A5%E9%85%8D%E7%BD%AE)
    - [可以读取环境变量](#%E5%8F%AF%E4%BB%A5%E8%AF%BB%E5%8F%96%E7%8E%AF%E5%A2%83%E5%8F%98%E9%87%8F)
    - [关于 Boolean flag](#%E5%85%B3%E4%BA%8E-Boolean-flag)
    - [可以解析到结构体](#%E5%8F%AF%E4%BB%A5%E8%A7%A3%E6%9E%90%E5%88%B0%E7%BB%93%E6%9E%84%E4%BD%93)
  - [日志相关](#%E6%97%A5%E5%BF%97%E7%9B%B8%E5%85%B3)
    - [日志分级](#%E6%97%A5%E5%BF%97%E5%88%86%E7%BA%A7)
    - [重定向日志](#%E9%87%8D%E5%AE%9A%E5%90%91%E6%97%A5%E5%BF%97)
    - [统一错误日志](#%E7%BB%9F%E4%B8%80%E9%94%99%E8%AF%AF%E6%97%A5%E5%BF%97)
    - [注意事项](#%E6%B3%A8%E6%84%8F%E4%BA%8B%E9%A1%B9)
  - [嵌入文件](#%E5%B5%8C%E5%85%A5%E6%96%87%E4%BB%B6)
    - [什么是 File Embedding](#%E4%BB%80%E4%B9%88%E6%98%AF-File-Embedding)
    - [embed.FS 会丢掉修改时间](#embedFS-%E4%BC%9A%E4%B8%A2%E6%8E%89%E4%BF%AE%E6%94%B9%E6%97%B6%E9%97%B4)
    - [使用嵌入的静态文件](#%E4%BD%BF%E7%94%A8%E5%B5%8C%E5%85%A5%E7%9A%84%E9%9D%99%E6%80%81%E6%96%87%E4%BB%B6)
    - [使用嵌入的模板文件](#%E4%BD%BF%E7%94%A8%E5%B5%8C%E5%85%A5%E7%9A%84%E6%A8%A1%E6%9D%BF%E6%96%87%E4%BB%B6)
  - [Testing](#Testing)
    - [测试入门](#%E6%B5%8B%E8%AF%95%E5%85%A5%E9%97%A8)
    - [表格测试](#%E8%A1%A8%E6%A0%BC%E6%B5%8B%E8%AF%95)
    - [封装 assert 函数](#%E5%B0%81%E8%A3%85-assert-%E5%87%BD%E6%95%B0)
    - [测试 HTTP Handler](#%E6%B5%8B%E8%AF%95-HTTP-Handler)
    - [测试 HTTP Middleware](#%E6%B5%8B%E8%AF%95-HTTP-Middleware)
    - [端到端测试 - 一个入门例子](#%E7%AB%AF%E5%88%B0%E7%AB%AF%E6%B5%8B%E8%AF%95--%E4%B8%80%E4%B8%AA%E5%85%A5%E9%97%A8%E4%BE%8B%E5%AD%90)
    - [封装可重用的 Test Helper](#%E5%B0%81%E8%A3%85%E5%8F%AF%E9%87%8D%E7%94%A8%E7%9A%84-Test-Helper)
    - [存储 Cookie 不跟随重定向](#%E5%AD%98%E5%82%A8-Cookie-%E4%B8%8D%E8%B7%9F%E9%9A%8F%E9%87%8D%E5%AE%9A%E5%90%91)
    - [端到端测试 - Mock 依赖项](#%E7%AB%AF%E5%88%B0%E7%AB%AF%E6%B5%8B%E8%AF%95--Mock-%E4%BE%9D%E8%B5%96%E9%A1%B9)
    - [端到端测试 - 发送 HTML 表单](#%E7%AB%AF%E5%88%B0%E7%AB%AF%E6%B5%8B%E8%AF%95--%E5%8F%91%E9%80%81-HTML-%E8%A1%A8%E5%8D%95)
    - [集成测试 - 与单元测试的区别](#%E9%9B%86%E6%88%90%E6%B5%8B%E8%AF%95--%E4%B8%8E%E5%8D%95%E5%85%83%E6%B5%8B%E8%AF%95%E7%9A%84%E5%8C%BA%E5%88%AB)
    - [集成测试 - 数据库集成测试](#%E9%9B%86%E6%88%90%E6%B5%8B%E8%AF%95--%E6%95%B0%E6%8D%AE%E5%BA%93%E9%9B%86%E6%88%90%E6%B5%8B%E8%AF%95)

## 命令行配置

### 用 flag 传入配置

Our web application’s `main.go` file currently contains a couple of hard-coded configuration settings:

- The network address for the server to listen on (currently `":4000"`)

- The file path for the static files directory (currently `"./ui/static"`)

Having these hard-coded isn’t ideal. There’s no separation between our configuration settings and code, and we can’t change the settings at runtime (which is important if you need different settings for development, testing and production environments).

In Go, a common and idiomatic way to manage configuration settings is to use command-line flags when starting an application. For example:

```bash
go run ./cmd/web -addr=":80"
```

The easiest way to accept and parse a command-line flag from your application is with a line of code like this. This essentially defines a new command-line flag with the name `addr`, a default value of `":4000"` and some short help text explaining what the flag controls. The value of the flag will be stored in the `addr` variable at runtime.

```go
addr := flag.String("addr", ":4000", "HTTP network address")
flag.Parse()
```

Let’s use this in our application and swap out the hard-coded network address in favor of a command-line flag instead:

```go
    // 定义名为 addr 的命令行 flag, 设置默认值和一小段解释
    addr := flag.String("addr", ":4000", "HTTP network address")
    
    // 调用 flag.Parse() 解析命令行 flag, 它会把解析结果存入 addr 变量
    // 必须在使用 addr 变量之前调用 flag.Parse(), 否则 addr 会保持默认值 :4000
    flag.Parse()

    // 注意 flag.String() 的返回值是指针, 所以要通过 *addr 获取值
    log.Printf("Starting server on %s", *addr)
```

In the code above we’ve used the `flag.String()` function to define the command-line flag. This has the benefit of converting whatever value the user provides at runtime to a `string` type. If the value can’t be converted to a string then the application will log an error and exit. Go also has a range of other functions including `flag.Int()`, `flag.Bool()` and `flag.Float64()`. 

Another great feature is that you can use the `-help` flag to list all the available command-line flags for an application and their accompanying help text. Give it a try:

```bash
go run ./cmd/web -help
```

### 可以读取环境变量

If you want, you can store your configuration settings in environment variables and access them directly from your application by using the `os.Getenv()` function like so:

```go
addr := os.Getenv("SNIPPETBOX_ADDR")
```

But this has some drawbacks compared to using command-line flags. You can’t specify a default setting (the return value from `os.Getenv()` is the empty string if the environment variable doesn’t exist), you don’t get the `-help` functionality that you do with command-line flags, and the return value from `os.Getenv()` is always a `string` — you don’t get automatic type conversions like you do with `flag.Int()` and the other command line flag functions.

Instead, you can get the best of both worlds by passing the environment variable as a command-line flag when starting the application. For example:

```bash
export SNIPPETBOX_ADDR=":9999"
go run ./cmd/web -addr=$SNIPPETBOX_ADDR
```

### 关于 Boolean flag

For flags defined with `flag.Bool()` omitting a value is the same as writing `-flag=true`. The following two commands are equivalent:

```bash
go run ./example -flag=true
go run ./example -flag
```

You must explicitly use `-flag=false` if you want to set a boolean flag value to false.

### 可以解析到结构体

It’s possible to parse command-line flag values into the memory addresses of pre-existing variables, using the `flag.StringVar()`, `flag.IntVar()`, `flag.BoolVar()` and other functions. This can be useful if you want to store all your configuration settings in a single struct. As a rough example:

```go
type config struct {
    addr      string
    staticDir string
}

func main() {
    var cfg config
    flag.StringVar(&cfg.addr, "addr", ":4000", "HTTP network address")
    flag.StringVar(&cfg.staticDir, "static-dir", "./ui/static", "Path to static assets")
    flag.Parse()
}
```

## 日志相关

### 日志分级

Both `log.Printf()` and `log.Fatal()` functions output messages via Go’s standard logger, which — by default — prefixes messages with the local date and time and writes them to the standard error stream (which should display in your terminal window).

In our application, we can break apart our log messages into two distinct types — or levels. The first type is informational messages (like "Starting server on :4000") and the second type is error messages. Let’s improve our application by adding some leveled logging capability, so that information and error messages are managed slightly differently. Specifically:

- We will prefix informational messages with `"INFO"` and output the message to standard out (stdout).

- We will prefix error messages with `"ERROR"` and output them to standard error (stderr), along with the relevant file name and line number that called the logger (to help with debugging).

There are a couple of different ways to do this, but a simple and clear approach is to use the log.New() function to create two new custom loggers. Open up your main.go file and update it as follows:

```go
    // 创建 Logger 时三个参数分别为，输出目标、日志前缀、额外信息
    infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

    // 这里用 log.Lshortfile flag 在错误日志中记录文件名和行号
    errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

    // 还是用 Print/Fatal 这些方法打日志，但换成了 infoLog 和 errorLog
    infoLog.Print("Starting server on ", server.Addr)
    err := server.ListenAndServe()
    if !errors.Is(err, http.ErrServerClosed) {
        errorLog.Fatal(err)
    }
```

### 重定向日志

During development, it’s easy to view the log output because the standard streams are displayed in the terminal. In staging or production environments, you can redirect the streams to a final destination for viewing and archival. This destination could be on-disk files, or a logging service such as Splunk. Either way, the final destination of the logs can be managed by your execution environment independently of the application.

For example, we could redirect the stdout and stderr streams to on-disk files when starting the application like so:

```bash
# Using the double arrow >> will append to an existing file,
# instead of truncating it when starting the application.
go run ./cmd/web >>/tmp/info.log 2>>/tmp/error.log
```

### 统一错误日志

By default, if Go’s HTTP server encounters an error it will log it using the standard logger. For consistency it’d be better to use our new `errorLog` logger instead. To make this happen we need to initialize a new `http.Server` struct containing the configuration settings for our server, instead of using the `http.ListenAndServe()` shortcut.

```go
    server := &http.Server{
        Addr:         *addr,
        Handler:      router,
        ErrorLog:     errorLog, // 让 Server 用这个 Logger
    }
```

### 注意事项

#### ➤ 不要在 main 函数之外用 panic, fatal

As a rule of thumb, you should avoid using the `Panic()` and `Fatal()` variations outside of your main() function — it’s good practice to return errors instead, and only panic or exit directly from `main()`.

#### ➤ Logger 是否并发安全

Custom loggers created by `log.New()` are concurrency-safe. You can share a single logger and use it across multiple goroutines and in your handlers without needing to worry about race conditions.

That said, *if you have multiple loggers writing to the same destination* then you need to **be careful and ensure that** the destination’s underlying `Write()` method is also safe for concurrent use.

总而言之:

- Logger 本身是并发安全的，如果每个 Logger 都有自己独享的输出文件，那么一切正常
- 如果两个 Logger 共享同一输出文件，就会出现两个线程同时写同一文件，此时必须在 `Write()` 方法中实现并发安全

## 嵌入文件

### 什么是 File Embedding

Go provides an `embed` package, which makes it possible to embed external files into your Go program itself. This feature is really nice because it makes it possible to create and distribute Go programs that are completely selfcontained. To illustrate how to use the `embed` package, we’ll update our application to embed and use the files in our existing `ui` directory (which contains our static CSS/JavaScript/image files and the HTML templates).

First create a new `ui/efs.go` file:

```go
//go:embed "html" "static"
var Files embed.FS
```

The important line here is `//go:embed "html" "static"`. This looks like a comment, but it is actually a special comment directive. When our application is compiled, this comment directive instructs Go to store the files from our `ui/html` and `ui/static` folders in an `embed.FS` embedded filesystem referenced by the global variable `Files`. 

There are a few important details about this which we need to explain.

- The `//go:embed` must be placed immediately above the variable in which you want to store the embedded files.
- The paths should be relative to the source code file containing the directive. So in our case, `go:embed "static" "html"` embeds the directories `ui/static` and `ui/html` from our project.
- You can only use the `go:embed` directive on global variables at package level, not within functions or methods.
- Paths cannot not contain `.` or `..` elements, nor may they begin or end with a `/`. This essentially restricts you to only embedding files that are contained in the same directory (or a subdirectory) as the source code which has the `go:embed` directive.
- If a path is to a directory, then all files in that directory are recursively embedded, except for files with names that begin with `.` or `_`. If you want to include these files you should use the `all:` prefix, like `go:embed "all:static"`.
- The path separator should always be a forward slash `/`, even on Windows machines.
- The embedded file system is always rooted in the directory which contains the `go:embed` directive. So, in the example above, our `Files` variable contains an `embed.FS` embedded filesystem and the root of that filesystem is our `ui` directory.

### embed.FS 会丢掉修改时间

- 使用 `//go:embed` 嵌入的文件会丢失所有 metadata，比如修改时间、权限信息
- 这会导致 `http.FileServer` 无法返回 `Last-Modified` 响应头，所以每次访问网站都会重新下载静态文件
  - 如果没有用 `Cache-Control` 响应头显式设置缓存行为，浏览器就会使用[启发式缓存](https://developer.mozilla.org/en-US/docs/Web/HTTP/Caching#heuristic_caching)
  - 启发式缓存的缓存时间为 (Date - Last-Modified) / 10，例如有 100 天没修改则缓存 10 天
  - 如果想使用 HTTP 缓存，推荐阅读相关资料，显式设置 `Cache-Control` 而不是依靠启发式缓存
- 用 Docker 打包应用的话，可以把静态文件放到某个路径，不用这个 embed.FS 也没关系

### 使用嵌入的静态文件

Let’s switch up our application so that it serves our static CSS, JavaScript and image files from the embedded file system — instead of reading them from the disk at runtime.

```go
// 这段代码要放在 ui/efs.go
//go:embed "html" "static"
var Files embed.FS

func routes(app *application) http.Handler {
    // 创建 File Server, 所有以 /static/ 开头的 URL 都交给它处理。因为需要把 GET /static/main.css 
    // 对应到文件目录中的 ./main.css, 所以用 http.StripPrefix 去掉 URL 中的 /static 前缀
    // fileServer := http.FileServer(http.Dir("./web/ui/static/"))
    // router.Handler(http.MethodGet, "/static/*filepath", http.StripPrefix("/static", fileServer))

    // 从 embed.FS 读取文件, 因为 URL /static/main.css 正好匹配路径 ./static/main.css 所以无需 http.StripPrefix
    fileServer := http.FileServer(http.FS(ui.Files))
    router.Handler(http.MethodGet, "/static/*filepath", fileServer)
}
```

### 使用嵌入的模板文件

Next let’s update the `cmd/web/templates.go` file so that our template cache uses the embedded HTML template files from `ui.Files`, instead of the ones on disk. We’ll need to leverage a couple of the special features that Go has for working with embedded filesystems:

- `fs.Glob()` returns a slice of filepaths matching a glob pattern. It’s effectively the same as the `filepath.Glob()` function that we used earlier in the book, except that it works on embedded filesystems.
- `Template.ParseFS()` can be used to parse the HTML templates from an embedded filesystem into a template set. This is effectively a replacement for both the `Template.ParseFiles()` and `Template.ParseGlob()` methods that we used earlier.

```go
// 这段代码要放在 ui/efs.go
//go:embed "html" "static"
var Files embed.FS

func newTemplateCache() (map[string]*template.Template, error) {
    // 用 map 作为缓存
    cache := make(map[string]*template.Template)

    // 使用 filepath.Glob 从磁盘获取以 .tmpl 结尾的文件路径, 返回值是 []string 类型
    // pages, err := filepath.Glob("./web/ui/html/pages/*.tmpl")

    // 使用 fs.Glob 从文件系统获取以 .tmpl 结尾的文件路径, 返回值是 []string 类型
    pages, err := fs.Glob(ui.Files, "html/pages/*.tmpl")
    if err != nil {
        return nil, err
    }

    for _, page := range pages {
        // 从文件路径提取文件名, 比如 ./web/ui/html/pages/home.tmpl 变成 home.tmpl
        name := filepath.Base(page)

        // base + nav + page 构成一个完整的页面
        patterns := []string{
            "html/base.tmpl",
            "html/partials/*.tmpl",
            page,
        }

        // 必须在 ParseFiles 解析模板内容之前把自定义函数注册上
        ts := template.New(name).Funcs(templateFunctions)

        // 使用 ParseFS 从 File System 读取文件, 而不是 ParseFiles 从磁盘读取文件
        ts, err := ts.ParseFS(ui.Files, patterns...)
        if err != nil {
            return nil, err
        }

        // 用 home.tmpl 这样的文件名作为 key 在缓存中存储解析好的模板
        cache[name] = ts
    }
    return cache, nil
}
```



## Testing

### 测试入门

- 测试文件和源文件放在同一目录，假如想测的函数在 `a.go` 里，那就新建名为 `a_test.go` 的测试文件

- 测试函数名以 `Test` 开头，用一个 `(t *testing.T)` 作为入参

- 执行测试的命令为: go test -v ./cmd/web/

```go
func humanDate(t time.Time) string {
    return t.Format("2006-01-02 15:04")
}

func TestHumanDate(t *testing.T) {
    want := "7777-07-07 07:07"
    got := humanDate(time.Date(7777, 7, 7, 7, 7, 0, 0, time.UTC))
    if got != want {
        t.Errorf("got %q; want %q", got, want)
    }
    // t.Errorf("用 t.Error 标记测试失败, 这不会终止程序")
    // t.Errorf("用 t.Error 标记测试失败, 程序将继续执行")
    // t.Fatalf("用 t.Fatal 终止当前测试or当前子测试, 其他测试继续执行")
}
```

### 表格测试

```go
func humanDate(t time.Time) string {
    if t.IsZero() {
        return ""
    }
    // 把时间转成 UTC+08:00 时区再进行格式化
    return t.In(time.FixedZone("CST", 8*3600)).Format("2006-01-02 15:04")
}

func TestHumanDate2(t *testing.T) {
    testCases := []struct {
        name string
        time time.Time
        want string
    }{
        {
            name: "Empty",
            time: time.Time{},
            want: "",
        },
        {
            name: "UTC",
            time: time.Date(2222, 2, 2, 0, 0, 0, 0, time.UTC),
            want: "2222-02-02 08:00", // 输入时间为 UTC 时区时, 格式化后应该加 8 个小时
        },
    }

    for _, tc := range testCases {
        // t.Run 默认同步运行, 所以此处无需使用 tc := tc, 除非在子测试中用了 t.Parallel()
        t.Run(tc.name, func(t *testing.T) {
            got := humanDate(tc.time)
            assert.Equal(t, tc.want, got)
        })
    }
}
```

It’s important to point out that you don’t need to use sub-tests in conjunction with table-driven tests (like we have done so far in this chapter). It’s perfectly valid to execute sub-tests by calling `t.Run()` consecutively in your test functions, similar to this:

```go
func TestExample(t *testing.T) {
    t.Run("Example sub-test 1", func(t *testing.T) {
        // Do a test.
    })
    t.Run("Example sub-test 2", func(t *testing.T) {
        // Do another test.
    })
    t.Run("Example sub-test 3", func(t *testing.T) {
        // And another...
    })
}
```

### 封装 assert 函数

```go
func Equal[T comparable](t *testing.T, expected, actual T) {
    // t.Errorf 会打印行号, 但是打印 t.Errorf 那一行的行号没有意义, 应该打印 assert.Equal() 对应的行号
    // 类似的问题在打印错误日志时也会遇到, 比如封装一个 app.serverError() 方法用于打印日志, 会用到 Output(2, ...)
    // func (app *application) serverError(w http.ResponseWriter, err error) {
    //     trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
    //     _ = app.errorLog.Output(2, trace) // 在日志中记录这一行没意义，应该记录调用 app.serverError 那一行
    //        http.Error(w, http.StatusText(500), 500)
    // }
    t.Helper()
    if actual != expected {
        t.Fatalf("got: %v; want %v", actual, expected)
    }
}

func NoError(t *testing.T, err error) {
    t.Helper()
    if err != nil {
        t.Fatal(err)
    }
}
```

### 测试 HTTP Handler

To assist in testing your HTTP handlers Go provides the `net/http/httptest` package, which contains a suite of useful tools. One of these tools is the `httptest.ResponseRecorder` type. This is essentially an implementation of `http.ResponseWriter` which records the response status code, headers and body instead of actually writing them to a HTTP connection. So an easy way to unit test your handlers is to create a new `httptest.ResponseRecorder` object, pass it to the handler function, and then examine it again after the handler returns.

```go
func ping(w http.ResponseWriter, r *http.Request) {
    _, _ = w.Write([]byte("OK"))
}

func TestPing(t *testing.T) {
    w := httptest.NewRecorder()
    r, err := httptest.NewRequest(http.MethodGet, "/", nil)
    if err != nil {
        t.Fatal(err)
    }
    // http handler 的单元测试, 就是创建请求、调用 handler、记录响应、检查响应
    // 注意 handler 通常包含很多依赖, 所以还得初始化依赖, 这个简单例子中没有体现这点
    ping(w, r)
    resp := w.Result()
    assert.Equal(t, http.StatusOK, resp.StatusCode)

    // 检查响应 Body
    //goland:noinspection GoUnhandledErrorResult
    defer resp.Body.Close()
    body, err := io.ReadAll(resp.Body)
    if err != nil {
        t.Fatal(err)
    }
    assert.Equal(t, "OK", string(body))
}
```

### 测试 HTTP Middleware

```go
func TestSecureHeaders(t *testing.T) {
    w := httptest.NewRecorder()
    r, err := httptest.NewRequest(http.MethodGet, "/", nil)
    assert.NoError(t, err)

    // 组合 Middleware 和 Handler 后调用它
    next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        _, _ = w.Write([]byte("OK"))
    })
    secureHeaders(next).ServeHTTP(w, r)
    resp := w.Result()

    // 检查有没有设置响应头
    expected := "default-src 'self'; style-src 'self' fonts.googleapis.com; font-src fonts.gstatic.com"
    assert.Equal(t, expected, resp.Header.Get("Content-Security-Policy"))
    assert.Equal(t, "origin-when-cross-origin", resp.Header.Get("Referrer-Policy"))
    assert.Equal(t, "nosniff", resp.Header.Get("X-Content-Type-Options"))
    assert.Equal(t, "deny", resp.Header.Get("X-Frame-Options"))
    assert.Equal(t, "0", resp.Header.Get("X-XSS-Protection"))

    // 检查 middleware 有没有调用 next
    assert.Equal(t, http.StatusOK, resp.StatusCode)
    defer resp.Body.Close()
    body, err := io.ReadAll(resp.Body)
    assert.NoError(t, err)
    assert.Equal(t, "OK", string(body))
}
```

### 端到端测试 - 一个入门例子

In the last chapter we talked through the general pattern for how to unit test your HTTP handlers in isolation. But — most of the time — your HTTP handlers aren’t actually used in isolation. So in this chapter we’re going to explain how to run endto-end tests on your web application that encompass your routing, middleware and handlers. Arguably, end-to-end testing should give you more confidence that your application is working correctly than unit testing in isolation.

- 端到端测试就是模拟 client 和 server 的交互

- 像这样启动 test server，然后用 client 访问它，就能让 router、middleware、handler 都被测到:

```go
func TestPing2(t *testing.T) {
    // 初始化依赖
    app := &application{
        errorLog: log.New(io.Discard, "", 0),
        infoLog:  log.New(io.Discard, "", 0),
    }

    // 启动 test server, 它会随机选一个可用端口, 用 ts.URL 获取访问地址, 另外别忘了用 defer 关掉服务器
    ts := httptest.NewTLSServer(routes(app))
    defer ts.Close()
    fmt.Println(ts.URL)

    // 使用 test server 提供的 client 访问测试服务器
    resp, err := ts.Client().Get(ts.URL + "/ping")
    assert.NoError(t, err)

    // 检查响应
    assert.Equal(t, http.StatusOK, resp.StatusCode)

    //goland:noinspection GoUnhandledErrorResult
    defer resp.Body.Close()
    body, err := io.ReadAll(resp.Body)
    assert.NoError(t, err)
    assert.Equal(t, "OK", string(body))
}
```

### 封装可重用的 Test Helper

Our TestPing test is now working nicely. But there’s a good opportunity to break out some of this code into helper functions, which we can reuse as we add more end-to-end tests to our project.  

```go
func newTestApplication(t *testing.T) *application {
    // 初始化依赖
    return &application{
        errorLog: log.New(io.Discard, "", 0),
        infoLog:  log.New(io.Discard, "", 0),
    }
}

type testServer struct {
    *httptest.Server
}

func newTestServer(t *testing.T, h http.Handler) *testServer {
    ts := httptest.NewTLSServer(h)
    return &testServer{
        Server: ts,
    }
}

// 向 test server 发出 get 请求并返回 status code, header, body
func (ts *testServer) get(t *testing.T, urlPath string) (int, http.Header, string) {
    resp, err := ts.Client().Get(ts.URL + urlPath)
    assert.NoError(t, err)

    //goland:noinspection GoUnhandledErrorResult
    defer resp.Body.Close()
    body, err := io.ReadAll(resp.Body)
    assert.NoError(t, err)

    return resp.StatusCode, resp.Header, string(body)
}

func TestPing3(t *testing.T) {
    // 所以封装一下, 代码更干净
    app := newTestApplication(t)
    ts := newTestServer(t, routes(app))
    defer ts.Close()

    code, header, body := ts.get(t, "/ping")
    assert.Equal(t, http.StatusOK, code)
    assert.Equal(t, "OK", body)
    assert.Equal(t, "text/plain; charset=utf-8", header.Get("Content-Type"))
}
```

### 存储 Cookie 不跟随重定向

- We want the client to automatically store any cookies sent in a HTTPS response, so that we can include them (if appropriate) in any subsequent requests back to the test server. This will come in handy later in the book when we need cookies to be supported across multiple requests in order to test our anti-CSRF measures.  
- We don’t want the client to automatically follow redirects. Instead we want it to return the first HTTPS response sent by our server so that we can test the response for that specific request.  

```go
func newTestServer(t *testing.T, h http.Handler) *testServer {
    ts := httptest.NewTLSServer(h)

    // 设置 cookie jar 让 client 像浏览器一样存储和发送 cookie
    jar, err := cookiejar.New(nil)
    assert.NoError(t, err)
    ts.Client().Jar = jar

    // 设置 client 不要跟随重定向, 因为我们要验证/检查的东西就是重定向响应本身
    // CheckRedirect 函数会在 client 收到 3xx 重定向时被执行
    ts.Client().CheckRedirect = func(req *http.Request, via []*http.Request) error {
        return http.ErrUseLastResponse // 不管什么类型的重定向都不跟随
    }

    return &testServer{
        Server: ts,
    }
}
```

### 端到端测试 - Mock 依赖项

Throughout this project we’ve injected dependencies into our handlers via the application struct, which currently looks like this:

```go
type application struct {
    errorLog       *log.Logger
    infoLog        *log.Logger
    snippets       models.SnippetModel          
    users          *models.UserModel             
    formDecoder    *form.Decoder                 
    templateCache  map[string]*template.Template 
    sessionManager *scs.SessionManager
}
```

When testing, it sometimes makes sense to mock these dependencies instead of using exactly the same ones that you do in your production application. For example, in the previous chapter we mocked the `errorLog` and `infoLog` dependencies with loggers that write messages to `io.Discard`, instead of the `os.Stdout` and `os.Stderr` streams like we do in our production application.

```go
func newTestApplication(t *testing.T) *application {
    return &application{
        errorLog: log.New(io.Discard, "", 0),
        infoLog:  log.New(io.Discard, "", 0),
    }
}
```

The reason for mocking these and writing to `io.Discard` is to avoid clogging up our test output with unnecessary log messages when we run `go test -v` (with verbose mode enabled). The other two dependencies that it makes sense for us to mock are the `models.SnippetModel` and `models.UserModel` database models. By creating mocks of these it’s possible for us to test the behavior of our handlers without needing to setup an entire test instance of the MySQL database.  

Create a new `internal/models/mocks` package:

```bash
mkdir internal/models/mocks
touch internal/models/mocks/snippets.go
touch internal/models/mocks/users.go
```

Let’s begin by creating a mock of our `models.SnippetModel`. To do this, we’re going to create a simple struct which implements the same methods as our production `models.SnippetModel`, but have the methods return some fixed dummy data instead.

```go
var mockSnippet = &models.Snippet{
    ID:      1,
    Title:   "I love homura",
    Content: "And I love hikari",
    Created: time.Now(),
    Expires: time.Now(),
}

type SnippetModel struct {
}

func (m *SnippetModel) Insert(title, content string, expires int) (int, error) {
    return 2, nil
}

func (m *SnippetModel) Get(id int) (*models.Snippet, error) {
    switch id {
    case 1:
        return mockSnippet, nil
    default:
        return nil, models.ErrNoRecord
    }
}

func (m *SnippetModel) Latest() ([]*models.Snippet, error) {
    return []*models.Snippet{mockSnippet}, nil
}
```

And let’s do the same for our `models.UserModel`, like so:  

```go
type UserModel struct {
}

func (m *UserModel) Insert(name, email, password string) error {
    switch email {
    case "nia@xb2.com":
        return models.ErrDuplicateEmail
    default:
        return nil
    }
}

func (m *UserModel) Authenticate(email, password string) (int, error) {
    if email == "nia@xb2.com" && password == "11111111" {
        return 1, nil
    }
    return 0, models.ErrInvalidCredentials
}

func (m *UserModel) Exists(id int) (bool, error) {
    switch id {
    case 1:
        return true, nil
    default:
        return false, nil
    }
}
```

Let’s head back to the `testutils_test.go` file and update the `newTestApplication()` function so that it creates an application struct with all the necessary dependencies for testing

```go
type application struct {
    snippets       models.SnippetModelInterface // 依赖接口而不是具体类型
    users          models.UserModelInterface    // GoLand 中可以一键从类型 Extract Interface
}

func newTestApplication(t *testing.T) *application {
    // 初始化依赖
    templateCache, err := newTemplateCache()
    assert.NoError(t, err)
    formDecoder := form.NewDecoder()
    sessionManager := scs.New() // 不设置 Store 则用内存作为 Store, 刚好适用于测试
    sessionManager.Lifetime = 12 * time.Hour
    sessionManager.Cookie.Secure = true
    return &application{
        errorLog:       log.New(io.Discard, "", 0),
        infoLog:        log.New(io.Discard, "", 0),
        snippets:       &mocks.SnippetModel{}, // 注意让 snippets, users 字段使用接口类型
        users:          &mocks.UserModel{},
        templateCache:  templateCache,
        formDecoder:    formDecoder,
        sessionManager: sessionManager,
    }
}
```

With that all now set up, let’s get stuck into writing an end-to-end test for our `snippetView` handler which uses these mocked dependencies.

```go
func TestSnippetView(t *testing.T) {
    app := newTestApplication(t)
    ts := newTestServer(t, routes(app))
    defer ts.Close()

    testCases := []struct {
        name     string
        urlPath  string
        wantCode int
        wantBody string
    }{
        {
            name:     "Valid ID",
            urlPath:  "/snippet/view/1",
            wantCode: http.StatusOK,
            wantBody: "And I love hikari",
        },
        {name: "Non-existent ID", urlPath: "/snippet/view/2", wantCode: http.StatusNotFound},
        {name: "Negative ID", urlPath: "/snippet/view/-1", wantCode: http.StatusNotFound},
        {name: "Decimal ID", urlPath: "/snippet/view/1.23", wantCode: http.StatusNotFound},
        {name: "String ID", urlPath: "/snippet/view/foo", wantCode: http.StatusNotFound},
        {name: "Empty ID", urlPath: "/snippet/view/", wantCode: http.StatusNotFound},
    }
    for _, tc := range testCases {
        tc := tc
        t.Run(tc.name, func(t *testing.T) {
            code, _, body := ts.get(t, tc.urlPath)
            assert.Equal(t, tc.wantCode, code)
            if tc.wantBody != "" {
                assert.StringContains(t, body, tc.wantBody)
            }
        })
    }
}
```

### 端到端测试 - 发送 HTML 表单

Testing this route is made a bit more complicated by the anti-CSRF check that our application does. Any request that we make to `POST /user/signup` will always receive a `400 Bad Request` response unless the request contains a valid CSRF token and cookie. To get around this we need to emulate the workflow of a real-life user as part of our test, like so:

1. Make a `GET /user/signup` request. This will return a response which contains a CSRF cookie in the response headers and the CSRF token for the signup page in the response body.

2. Extract the CSRF token from the HTML response body.
3. Make a `POST /user/signup` request, using the same `http.Client` that we used in step 1 (so it automatically passes the CSRF cookie with the POST request) and including the CSRF token alongside the other POST data that we want to test.

Let’s begin by adding a new helper function to our `cmd/web/testutils_test.go` file for extracting the CSRF token (if one exists) from a HTML response body:

```go
// 从 HTML 中提取 CSRF token 的正则表达式
var csrfTokenRX = regexp.MustCompile(`<input type='hidden' name='csrf_token' value='(.+)'>`)

func extractCSRFToken(t *testing.T, body string) string {
    // FindStringSubmatch 只返回首个匹配的子串, 可用 FindAllStringSubmatch 返回所有匹配的子串
    // 返回值是 []string 类型, 首个元素是与 regexp 匹配的子串, 剩下的元素与正则中的 capture group 一一对应
    matches := csrfTokenRX.FindStringSubmatch(body)
    if len(matches) < 2 {
        t.Fatal("no csrf token found in body")
    }
    // 我们在模板中用 <input ... value='{{.CSRFToken}}'> 设置表单的隐藏字段
    // token 中可能包含 '+' 字符, 会被 {{.CSRFToken}} 转义成 &#43; 所以取值时要还原
    return html.UnescapeString(matches[1])
}
```

Let’s go back to our cmd/web/handlers_test.go file and create a new TestUserSignup test.

```go
func TestUserSignup(t *testing.T) {
    app := newTestApplication(t)
    ts := newTestServer(t, routes(app))
    defer ts.Close()

    _, _, body := ts.get(t, "/user/signup")
    csrfToken := extractCSRFToken(t, body)
    t.Logf("CSRF token is: %q", csrfToken)
}
```

Let’s create a new postForm() method on our testServer type.

```go
// 向 test server 发出 get 请求并返回 status code, header, body
func (ts *testServer) postForm(t *testing.T, urlPath string, form url.Values) (int, http.Header, string) {
    resp, err := ts.Client().PostForm(ts.URL+urlPath, form)
    assert.NoError(t, err)

    //goland:noinspection GoUnhandledErrorResult
    defer resp.Body.Close()
    body, err := io.ReadAll(resp.Body)
    assert.NoError(t, err)
    return resp.StatusCode, resp.Header, string(body)
}
```

And now, at last, we’re ready to add some table-driven sub-tests to test the behavior of our application’s `POST /user/signup` route. Specifically, we want to test that:  

- A valid signup results in a `303 See Other` response.
- A form submission without a valid CSRF token results in a `400 Bad Request` response.
- A invalid from submission results in a `422 Unprocessable Entity` response and the signup form being redisplayed. This should happen when:
  - The name, email or password fields are empty.
  - The email is not in a valid format.
  - The password is less than 8 characters long.
  - The email address is already in use.

```go
func TestUserSignup(t *testing.T) {
    app := newTestApplication(t)
    ts := newTestServer(t, routes(app))
    defer ts.Close()

    _, _, body := ts.get(t, "/user/signup")
    validCSRFToken := extractCSRFToken(t, body)

    defaultForm := url.Values{}
    defaultForm.Add("name", "homura")
    defaultForm.Add("email", "homura@xb2.com")
    defaultForm.Add("password", "11111111")
    defaultForm.Add("csrf_token", validCSRFToken)

    formTag := "<form action='/user/signup' method='POST' novalidate>"
    testCases := []struct {
        name        string
        form        url.Values
        wantCode    int
        wantFormTag string
    }{
        {name: "Valid submission", wantCode: 303, form: formFields()},
        {name: "Invalid CSRF Token", wantCode: 400, form: formFields("csrf_token", "wrongToken")},
        {name: "Empty name", wantCode: 422, wantFormTag: formTag, form: formFields("name", "")},
        {name: "Empty email", wantCode: 422, wantFormTag: formTag, form: formFields("email", "")},
        {name: "Empty password", wantCode: 422, wantFormTag: formTag, form: formFields("password", "")},
        {name: "Invalid email", wantCode: 422, wantFormTag: formTag, form: formFields("email", "bob@example.")},
        {name: "Short password", wantCode: 422, wantFormTag: formTag, form: formFields("password", "pass")},
        {name: "Duplicate email", wantCode: 422, wantFormTag: formTag, form: formFields("email", "nia@xb2.com")},
    }
    for _, tc := range testCases {
        tc := tc
        t.Run(tc.name, func(t *testing.T) {
            // 除非设置过值, 否则使用 defaultForm 中的默认值
            form := url.Values{}
            for k, v := range defaultForm {
                form[k] = v
                if tc.form.Has(k) {
                    form[k] = tc.form[k]
                }
            }

            code, _, body := ts.postForm(t, "/user/signup", form)
            assert.Equal(t, tc.wantCode, code)
            if tc.wantFormTag != "" {
                assert.StringContains(t, body, tc.wantFormTag)
            }

        })
    }
}

func formFields(s ...string) url.Values {
    if len(s)%2 != 0 {
        panic("wrong number of arguments")
    }
    form := url.Values{}
    for i := 0; i+1 < len(s); i += 2 {
        form.Add(s[i], s[i+1])
    }
    return form
}
```

### 集成测试 - 与单元测试的区别

- [What's the difference between unit tests and integration tests?](https://stackoverflow.com/questions/5357601/whats-the-difference-between-unit-tests-and-integration-tests)

> A **unit test** is a test written by the programmer to verify that a relatively small piece of code is doing what it is intended to do. Unit tests should only depend on the tested implementation unit; they should not depend on external components such as databases, network services, web browser interaction. When such external elements are required, unit tests use mock objects. Unit tests test internal consistency as opposed to proving that they play nicely with some outside system.
>
> Why should a unit test have no external dependencies? Because in a unit test you only want to check the behavior of the tested function in a well-defined context. It should not be influenced by a potential bug in the dependencies. If you want to assert that the combination of the function and the dependencies works as expected, you are writing a integration test.

- 单元测试是为了验证一个较小的功能单元 ( 比如函数和类 ) 是否正确，因为外部依赖不是测试目标所以用 Mock 替代
- 集成测试是为了测试组件之间的配合，「连接数据库测试 CRUD 代码」就属于集成测试，它测试了 app 和 db 之间的配合
- 单元测试只测试一个工作单元，而集成测试要测试多个工作单元的协作。  
  单元测试用于验证零件的正确性，而集成测试用于验证组装的正确性。  

### 集成测试 - 数据库集成测试

Running end-to-end tests with mocked dependencies is a good thing to do, but we could improve confidence in our application even more if we also verify that our real MySQL database models are working as expected. To do this we can run integration tests against a test version our MySQL database. The first step is to create the test version of our MySQL database.

```mysql
# 创建测试数据库和用户
CREATE DATABASE test_snippetbox CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
CREATE USER 'test_web'@'%';
GRANT CREATE, DROP, ALTER, INDEX, SELECT, INSERT, UPDATE, DELETE ON test_snippetbox.* TO 'test_web'@'%';
ALTER USER 'test_web'@'%' IDENTIFIED BY 'pass';
```

Once that’s done, let’s make two SQL scripts:

1. A setup script to create the database tables (so that they mimic our production database) and insert a known set of test data than we can work with in our tests.

2. A teardown script which drops the database tables and any data.

The idea is that we’ll call these scripts at the start and end of each integration test, so that the test database is fully reset each time. This helps ensure that any changes we make during one test are not ‘leaking’ and affecting the results of another test. Let’s go ahead and create these scripts in a new `internal/models/testdata` directory like so:

```bash
# The Go tool ignores any directories called testdata, so these scripts will be ignored when compiling your 
# application (it also ignores any directories or files which have names that begin with an _ or . character)
mkdir internal/models/testdata
touch internal/models/testdata/setup.sql
touch internal/models/testdata/teardown.sql
```

File: internal/models/testdata/setup.sql

```mysql
CREATE TABLE snippets
(
    id      INTEGER      NOT NULL PRIMARY KEY AUTO_INCREMENT,
    title   VARCHAR(100) NOT NULL,
    content TEXT         NOT NULL,
    created DATETIME     NOT NULL,
    expires DATETIME     NOT NULL
);
CREATE INDEX idx_snippets_created ON snippets (created);
CREATE TABLE users
(
    id              INTEGER      NOT NULL PRIMARY KEY AUTO_INCREMENT,
    name            VARCHAR(255) NOT NULL,
    email           VARCHAR(255) NOT NULL,
    hashed_password CHAR(60)     NOT NULL,
    created         DATETIME     NOT NULL
);
ALTER TABLE users
    ADD CONSTRAINT users_uc_email UNIQUE (email);
INSERT INTO users (name, email, hashed_password, created)
VALUES ('homura',
        'homura@xb2.com',
        '$2a$12$vbuOtZQwRmA.jvg/8ZnGTuzqaG703sOuY9MHzCKg7xv4DmSLeSlp2',
        '2023-08-10 04:38:54');
```

File: internal/models/testdata/teardown.sql

```mysql
DROP TABLE users;
DROP TABLE snippets;
```

Let’s create a `internal/models/testutils_test.go` file and add to it a `newTestDB()` helper function which:

- Creates a new `*sql.DB` connection pool for the test database;
- Executes the `setup.sql` script to create the database tables and dummy data;
- Register a ‘cleanup’ function which executes the `teardown.sql` script and closes the connection pool.

```go
func newTestDB(t *testing.T) *sql.DB {
    // 使用 test user 连接 test database, 添加 multiStatements=true 参数让 db.Exec() 能执行多条 SQL
    dsn := "test_web:pass@tcp(localhost)/test_snippetbox?parseTime=true&multiStatements=true"
    db, err := sql.Open("mysql", dsn)
    assert.NoError(t, err)

    // 读取并执行 setup.sql
    script, err := os.ReadFile("./testdata/setup.sql")
    assert.NoError(t, err)
    _, err = db.Exec(string(script))
    assert.NoError(t, err)

    // 使用 t.Cleanup() 注册清理函数, 当调用 newTestDB(t) 的测试或子测试结束时, 清理函数就会执行
    // 在清理函数中读取并执行 teardown.sql, 然后关掉数据库连接池
    t.Cleanup(func() {
        script, err := os.ReadFile("./testdata/teardown.sql")
        assert.NoError(t, err)
        _, err = db.Exec(string(script))
        assert.NoError(t, err)
        _ = db.Close()
    })

    return db
}
```

We know that our `setup.sql` script creates a users table containing one record (which should have the user ID 1). So we want to test that: Calling `models.UserModel.Exists(1)` returns `(true, nil)`. Calling `models.UserModel.Exists()` with any other user ID returns `(false, nil)`

```go
// 使用了真实数据库, 所以这属于集成测试, 测试目的为「 验证 CRUD 功能是否正确 」
func TestUserModelExists(t *testing.T) {
    testCases := []struct {
        name   string
        userID int
        want   bool
    }{
        {name: "Valid ID", userID: 1, want: true},
        {name: "Zero ID", userID: 0, want: false},
        {name: "Nonexistent ID", userID: 2, want: false},
    }

    for _, tc := range testCases {
        tc := tc
        t.Run(tc.name, func(t *testing.T) {
            // 在这里调用 newTestDB(), 所以每次执行子测试都会创建并删掉数据库表
            db := newTestDB(t)
            m := UserModel{db}
            exists, err := m.Exists(tc.userID)
            assert.Equal(t, tc.want, exists)
            assert.NilError(t, err)
        })
    }
}
```