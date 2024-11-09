## Table of Contents
  - [Getting Started](#Getting-Started)
    - [初始化项目](#%E5%88%9D%E5%A7%8B%E5%8C%96%E9%A1%B9%E7%9B%AE)
    - [RESTful 风格的 API](#RESTful-%E9%A3%8E%E6%A0%BC%E7%9A%84-API)
    - [使用 httprouter](#%E4%BD%BF%E7%94%A8-httprouter)
  - [JSON 响应](#JSON-%E5%93%8D%E5%BA%94)
    - [返回 JSON 响应](#%E8%BF%94%E5%9B%9E-JSON-%E5%93%8D%E5%BA%94)
    - [一些注意事项](#%E4%B8%80%E4%BA%9B%E6%B3%A8%E6%84%8F%E4%BA%8B%E9%A1%B9)
    - [序列化结构体](#%E5%BA%8F%E5%88%97%E5%8C%96%E7%BB%93%E6%9E%84%E4%BD%93)
    - [自定义 JSON 序列化](#%E8%87%AA%E5%AE%9A%E4%B9%89-JSON-%E5%BA%8F%E5%88%97%E5%8C%96)
    - [自定义 JSON 反序列化](#%E8%87%AA%E5%AE%9A%E4%B9%89-JSON-%E5%8F%8D%E5%BA%8F%E5%88%97%E5%8C%96)
  - [JSON 错误](#JSON-%E9%94%99%E8%AF%AF)
    - [JSON 错误消息](#JSON-%E9%94%99%E8%AF%AF%E6%B6%88%E6%81%AF)
    - [Panic 异常恢复](#Panic-%E5%BC%82%E5%B8%B8%E6%81%A2%E5%A4%8D)
  - [JSON 请求](#JSON-%E8%AF%B7%E6%B1%82)
    - [解析 JSON 请求](#%E8%A7%A3%E6%9E%90-JSON-%E8%AF%B7%E6%B1%82)
    - [Create a Movie](#Create-a-Movie)
    - [校验 JSON 请求](#%E6%A0%A1%E9%AA%8C-JSON-%E8%AF%B7%E6%B1%82)

## Getting Started

### 初始化项目

#### ➤ 目录结构

```bash
go mod init api.xianyukang.com   # 初始化 go module
mkdir -p bin                     # 存放编译好的程序
mkdir -p cmd/api                 # 后端 api 接口，处理用户请求
mkdir -p internal                # 各种各样的工具包: database, validation, email, ...
mkdir -p migrations              # 使用 migration 可记录和复现数据库表结构的变动
mkdir -p remote                  # configuration files and setup scripts for our production server
touch Makefile                   # 构建程序时会用到的命令和脚本
touch cmd/api/main.go            # 
```

#### ➤ hello world !

```go
// 版本号, 可以在构建时设置 version 变量的值
const version = "1.0.0"

// 配置信息, 可以从 command line flag 读取
type config struct {
    port int    // 端口
    env  string // 环境 development, production
}

// 应用, 包含各种组件, 这些组件会被 HTTP handler 调用
type application struct {
    config config
    logger *log.Logger
}

func main() {
    var cfg config
    flag.IntVar(&cfg.port, "port", 4000, "API Server Port")
    flag.StringVar(&cfg.env, "env", "development", "Environment (development|production)")
    flag.Parse()

    // 初始化应用组件
    logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)
    app := &application{
        config: cfg,
        logger: logger,
    }

    // 创建 Router 和 Server
    mux := http.NewServeMux()
    mux.HandleFunc("/v1/healthcheck", app.healthcheckHandler)

    srv := &http.Server{
        Addr:         fmt.Sprintf(":%d", cfg.port),
        Handler:      mux,
        IdleTimeout:  time.Minute,
        ReadTimeout:  10 * time.Second,
        WriteTimeout: 30 * time.Second,
    }

    logger.Printf("starting %q server on %q", cfg.env, srv.Addr)
    err := srv.ListenAndServe()
    logger.Fatal(err)
}

func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
    _, _ = fmt.Fprintln(w, "status: available")
    _, _ = fmt.Fprintf(w, "environment: %s\n", app.config.env) // Handler 通过 app 调用应用组件
    _, _ = fmt.Fprintf(w, "version: %s\n", version)
}
```

#### ➤ 运行和测试

```bash
go run ./cmd/api                              # 以相对路径运行包, 注意要加 ./ 前缀
curl -i localhost:4000/v1/healthcheck         # -i 表示打印响应头
go run ./cmd/api -port=3030 -env=production   # 通过命令行修改服务器配置
```

### RESTful 风格的 API

The first thing is that requests with the same URL pattern will be routed to different handlers based on the HTTP request method.

- `GET` Use for actions that retrieve information only and don’t change the state of your application or any data.
- `POST` Use for non-idempotent actions that modify state. `POST` is generally used for actions that create a new resource.
- `PUT` Use for idempotent actions that modify the state of a resource at a specific URL. `PUT` is generally used for actions that replace or update an existing resource.
- `PATCH` Use for actions that partially update a resource at a specific URL. It’s OK for the action to be either idempotent or non-idempotent.
- `DELETE` Use for actions that delete a resource at a specific URL.

| Method | URL Pattern | Handler | Action |
| ----------------------- | --------------- | ------------------ | ---------------------------- |
| GET | /v1/movies | listMoviesHandler | Show the details of all movies |
| POST | /v1/movies | createMovieHandler | Create a new movie        |
| GET | /v1/movies/:id | showMovieHandler | Show the details of a specific movie |
| PUT | /v1/movies/:id | editMovieHandler | Update the details of a specific movie |
| DELETE | /v1/movies/:id | deleteMovieHandler | Delete a specific movie |

### 使用 httprouter

#### ➤ 标准库的 http.ServeMux 也很好，但你的需求超出了它的功能时，用三方库更合适

`http.ServeMux`, the router in the Go standard library — is quite limited in terms of its functionality. In particular it doesn’t allow you to route requests to different handlers based on the request method (GET, POST, etc.), nor does it provide support for clean URLs with interpolated parameters. Although you can [work-around](https://www.youtube.com/watch?v=yi5A3cK1LNA&t=704s) these limitations, generally it’s easier to use one of the many third-party routers that are available instead.

#### ➤ 安装 httprouter

```bash
go get github.com/julienschmidt/httprouter     # 著名的 Gin 框架使用了修改版的 httprouter
touch cmd/api/routes.go                        # router 相关代码放到单独的文件, 而不是全塞到 main.go
touch cmd/api/movies.go                        # movie 相关的 handler 放到单独的文件
```

#### ➤ 入门 httprouter

```go
func (app *application) routes() *httprouter.Router {
    router := httprouter.New()
    router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)
    router.HandlerFunc(http.MethodPost, "/v1/movies", app.createMovieHandler)
    router.HandlerFunc(http.MethodGet, "/v1/movies/:id", app.showMovieHandler)
    return router
}

func main() {
    // 创建 Server
    srv := &http.Server{
        Addr:         fmt.Sprintf(":%d", cfg.port),
        Handler:      app.routes(),
    }
}

func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
    _, _ = fmt.Fprintln(w, "create a new movie")
}

func (app *application) showMovieHandler(w http.ResponseWriter, r *http.Request) {
    // 从 url pattern "/v1/movies/:id" 中取出 id, 需要转成整数
    params := httprouter.ParamsFromContext(r.Context())
    id, err := strconv.ParseInt(params.ByName("id"), 10, 64)
    if err != nil || id < 1 {
        http.NotFound(w, r)
        return
    }
    _, _ = fmt.Fprintf(w, "show the details of movie %d\n", id)
}

```

测试一下:

```bash
go run ./cmd/api
curl -i localhost:4000/v1/healthcheck
curl -i -X POST localhost:4000/v1/movies
curl -i localhost:4000/v1/movies/123
curl -i localhost:4000/v1/movies/-1
```

#### ➤ 注意冲突的路由

- `GET /users/:name` 会和 `GET /users/groups` 发生冲突，原版 httprouter 不允许这么做
- 但是 Gin 中的修改版 httprouter [允许这么做](https://github.com/gin-gonic/gin/pull/2663)，其他的 router 例如 [chi](https://github.com/go-chi/chi) 也支持 conflict routes

#### ➤ 封装辅助函数: readIDParam()

The code to extract an `id` parameter from a URL like `/v1/movies/:id` is something that we’ll need repeatedly in our application, so let’s abstract the logic for this into a small reuseable helper method.  

```bash
touch cmd/api/helpers.go
```

And add a new `readIDParam()` method to the `application` struct, like so:

```go
func (app *application) readIDParam(r *http.Request) (int64, error) {
    // 从 url pattern "/v1/movies/:id" 中取出 id, 需要转成整数
    params := httprouter.ParamsFromContext(r.Context())
    id, err := strconv.ParseInt(params.ByName("id"), 10, 64)
    if err != nil || id < 1 {
        return 0, errors.New("invalid id parameter")
    }
    return id, nil
}

func (app *application) showMovieHandler(w http.ResponseWriter, r *http.Request) {
    id, err := app.readIDParam(r)
    if err != nil { // 无需判断 id < 1, 已经被 readIDParam 处理了
        http.NotFound(w, r)
        return
    }
    _, _ = fmt.Fprintf(w, "show the details of movie %d\n", id)
}
```

Let's create a reusable helper for sending JSON responses, which will ensure that all your API responses have a sensible and consistent structure.  

## JSON 响应

### 返回 JSON 响应

#### ➤ 最简单的做法

```go
func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
    // 首先把 Content-Type: text/plain; charset=utf-8 改成 Content-Type: application/json
    w.Header().Set("Content-Type", "application/json")

    // 然后往 Body 中写入 json 字符串, 就能返回 json 响应
    js := fmt.Sprintf(`{"status": "available", "environment": %q, "version": %q}`, app.config.env, version)
    _, _ = w.Write([]byte(js))
}
```

#### ➤ json.Marshal vs json.Encoder

At a high-level, Go’s `encoding/json` package provides two options for encoding things to JSON. You can either call the `json.Marshal()` function, or you can declare and use a `json.Encoder` type. There are some differences:

- You can set encoding options on the [json.Encoder](https://godoc.org/encoding/json#Encoder). Examples are [SetEscapeHTML](https://godoc.org/encoding/json#Encoder.SetEscapeHTML) and [SetIndent](https://godoc.org/encoding/json#Encoder.SetIndent).
- Use the one that's most convenient. If you want to play with string or bytes use `json.Marshal`.
- Use `json.Decoder` if your data is coming from an `io.Reader` stream, or you need to decode multiple values from a stream of data.

#### ➤ 用 json.Marshal() 把 Go 对象序列化成 JSON 字符串

```go
func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
    data := map[string]string{
        "status":      "available",
        "environment": app.config.env,
        "version":     version,
    }

    js, err := json.Marshal(data)
    if err != nil {
        app.logger.Print(err)
        http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    _, _ = w.Write(js)
}
```

#### ➤ 封装辅助函数: writeJSON()

As our API grows we’re going to be sending a lot of JSON responses, so it makes sense to move some of this logic into a reusable `writeJSON()` helper method.

```go
func (app *application) writeJSON(w http.ResponseWriter, status int, data any, headers http.Header) error {
    js, err := json.Marshal(data)
    if err != nil {
        return err
    }

    // 响应头必须在写入 body 前设置好, 否则不会起作用
    for key, value := range headers {
        w.Header()[key] = value
    }
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(status)

    _, err = w.Write(js)
    return err
}

func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
    data := map[string]string{
        "status":      "available",
        "environment": app.config.env,
        "version":     version,
    }

    err := app.writeJSON(w, http.StatusOK, data, nil)
    if err != nil {
        app.logger.Print(err)
        http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
    }
}
```

### 一些注意事项

- `[]byte` 类型的字段会被序列化成 base64 字符串 ( 所谓 base64 编码就是用一串可打印字符来表示任意二进制数据 )
- `map` 类型被序列化成 JSON 时，会对 key 进行字典序排序。结构体的字段顺序和 Go 源码书写顺序一致
- 序列化 `map[K]V` 时，`K` 必须是 string/int/time.Time/net.IP，或其他实现了 `encoding.TextMarshaler` 的类型
- 众所周知，nil slice 会被序列化成 `null`，而 empty slice 会被序列化成空数组 `[]`
- 序列化字符串时，其中的 `<>&` 字符会被分别转义成 `\u003c` `\u003e` `\u0026`，可通过 [SetEscapeHTML(false)](https://pkg.go.dev/encoding/json#Encoder.SetEscapeHTML) 关掉

#### ➤ 反序列化时，若以 interface{} 接收，数字默认解析为 float64 类型

```go
JSON numbers => float64
JSON arrays  => []interface{}
JSON objects => map[string]interface{}

func TestUseNumber(t *testing.T) {
    // 若不喜欢默认的 float64, 想解析成整数或字符串可以这么做
    js := `10`
    var n any
    dec := json.NewDecoder(strings.NewReader(js))
    dec.UseNumber() // Call the UseNumber() method on the decoder before using it.
    _ = dec.Decode(&n)
    nString := n.(json.Number).String()
    nInt64, _ := n.(json.Number).Int64()
    fmt.Println(nString, nInt64)
}
```

#### ➤ Using omitempty on a zero-valued struct doesn’t work

The `omitempty` struct tag directive never considers a struct to be empty — even if all the struct fields have their zero value and you use `omitempty` on those fields too. It will always appear as an object in the encoded JSON. For example, the following struct:  

```go
    s := struct {
        Foo struct {
            Bar string `json:",omitempty"`
        } `json:",omitempty"`
    }{}

// 对应的 JSON 输出为 {"Foo":{}}
// 这里空对象很碍眼，但无法用 omitempty 忽略掉 struct 类型的 empty 字段
// 解决办法是用指针, 因为 omitempty 能忽略掉指针类型的 empty 字段
```

#### ➤ 实现 MarshalJSON() 时记得用 value receiver

```go
func TestMarshalJSON(t *testing.T) {
    f := myFloat(1.0 / 3.0)

    // When encoding a pointer, the MarshalJSON method is used.
    js, _ := json.Marshal(&f)
    fmt.Printf("%s\n", js)

    // When encoding a value, the MarshalJSON method is ignored.
    js, _ = json.Marshal(f)
    fmt.Printf("%s\n", js)

    // 这种现象是因为, f 和 &f 拥有不同的 method set, f 的方法集不包含指针方法
}

type myFloat float64

// This has a pointer receiver. ( 别这样, 建议用 value receiver )
func (f *myFloat) MarshalJSON() ([]byte, error) {
    return []byte(fmt.Sprintf("%.2f", *f)), nil
}
```

#### ➤ Partial JSON decoding

If you have a lot of JSON input to process and only need a small part of it, it’s often possible to leverage the `json.RawMessage` type to help deal with this. For example:

```go
func TestPartialParse(t *testing.T) {
    // 假设只对其中的 genres 数组感兴趣
    js := `{"title": "Top Gun", "genres": ["action", "romance"], "year": 1986}`

    // 这个 json.RawMessage 中存储着一串字节, 尚未被解析
    var m map[string]json.RawMessage
    err := json.NewDecoder(strings.NewReader(js)).Decode(&m)
    if err != nil {
        log.Fatal(err)
    }

    // key 已经被解析了, 所以能访问 genres 字段, 此时再用 json.Unmarshal() 做进一步解析
    var genres []string
    err = json.Unmarshal(m["genres"], &genres)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("genres: %v\n", genres)
}
```

### 序列化结构体

#### ➤ 入门例子

Instead of encoding a map to create this JSON object (like we did in the previous chapter), this time we’re going to encode a custom Movie struct. So, first things first, we need to begin by defining a custom `Movie` struct. We’ll do this inside a new `internal/data` package, which later will grow to encapsulate all the custom data types for our project along with the logic for interacting with our database.

```bash
mkdir internal/data
touch internal/data/movies.go
```

```go
type Movie struct {
    // 一般用 int64, 没必要用 int32, 参考 https://www.reddit.com/r/golang/comments/sg8fyn/
    // 读写数据库/磁盘/网络等外部数据时, 不能用 int 因为它的大小是模糊的, 不能一会发 64 字节一会又发 32 字节
    ID        int64     `json:"id"`
    CreatedAt time.Time `json:"created_at"` // 这条数据何时添加到数据库, 并非电影的年份
    Title     string    `json:"title"`
    Year      int32     `json:"year"`
    Runtime   int32     `json:"runtime"` // 持续时间 (分钟)
    Genres    []string  `json:"genres"`  // 题材或艺术种类, 发音: https://www.bilibili.com/video/BV16j411z7zm/
    Version   int32     `json:"version"` // 从 1 开始, 表示这条数据的版本, 每次更新都会让版本号 +1
}

func (app *application) showMovieHandler(w http.ResponseWriter, r *http.Request) {
    id, err := app.readIDParam(r)
    if err != nil { // 无需判断 id < 1, 已经被 readIDParam 处理了
        http.NotFound(w, r)
        return
    }

    movie := data.Movie{
        ID:        id,
        CreatedAt: time.Now(),
        Title:     "Big Hero 6",
        Year:      2014,
        Runtime:   102,
        Genres:    []string{"Animation", "Adventure", "Comedy", "Sci-Fi"},
        Version:   1,
    }

    err = app.writeJSON(w, http.StatusOK, movie, nil)
    if err != nil {
        app.logger.Print(err)
        http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
    }
}
```

#### ➤ json key 用下划线还是驼峰?

- [json-api/json-api: A specification for building JSON APIs](https://github.com/json-api/json-api) ( 太长啦 )

```html
(1) 并没有世界统一的命名规范
(2) 例如 Google 用过驼峰，Twitter 和 Facebook 用过下划线
(3) Java 和 JavaScript 喜欢用驼峰，Python 和 PHP 喜欢用下划线
(4) 总之下划线和驼峰都行，重要的是统一性 ( 如果不知道怎么选，建议以身边的例子为准 )
```

#### ➤ json 相关的 struct tag

- 输出 json 时可用 `omitempty` 忽略 empty 字段，也就是 false/0/""/nil 以及空的 slice/array/map

```go
type User struct {
    Name      string    `json:"name"`
    Password  string    `json:"-"`                  // 符号 - 表示不管什么情况都忽略这个字段
    Age       int32     `json:"age,omitempty"`      // omitempty 用来忽略 empty 字段
    Birthday  time.Time `json:"birthday,omitempty"` // omitempty 无法用于结构体, 可用指针解决
    AgeString int32     `json:"age_string,string"`  // 用 string 表示输出为字符串类型
}

func TestStructTags(t *testing.T) {
    u := User{
        Name: "Homura",
    }
    _ = json.NewEncoder(os.Stdout).Encode(u)
}
```

#### ➤ Enveloping responses

Next let’s work on updating our responses so that the JSON data is always enveloped in a parent JSON object. Similar to this:

```json
{
  "movie": {
    "id": 123,
    "title": "Big Hero 6",
    "year": 2014,
    "runtime": 102,
    "genres": ["Animation", "Adventure", "Comedy", "Sci-Fi"],
    "version": 1
  }
}
```

Notice how the movie data is nested under the key `"movie"` here, rather than being the toplevel JSON object itself? Enveloping response data like this isn’t strictly necessary, and whether you choose to do so is partly a matter of style and taste. But there are a few tangible benefits:

1. Including a key name (like "movie") at the top-level of the JSON helps make the response more self-documenting. For any humans who see the response out of context, it is a bit easier to understand what the data relates to.
2. It reduces the risk of errors on the client side, because it’s harder to accidentally process one response thinking that it is something different. To get at the data, a client must explicitly reference it via the `"movie"` key.

```go
type envelop map[string]any
app.writeJSON(w, http.StatusOK, envelop{"movie": movie}, nil) // 数据放到 movie 这个键, 而不是直接作为顶层对象
```

### 自定义 JSON 序列化

#### ➤ 通过实现 json.Marshaler 接口，自定义如何序列化一个类型

When Go is encoding a particular type to JSON, it looks to see if the type has a `MarshalJSON()` method implemented on it. If it has, then Go will call this method to determine how to encode it. Strictly speaking, when Go is encoding a particular type to JSON it looks to see if the type satisfies the `json.Marshaler` interface, which looks like this:

```go
type Marshaler interface {
    MarshalJSON() ([]byte, error)
}
```

If the type does satisfy the interface, then Go will call its `MarshalJSON()` method and use the `[]byte` slice that it returns as the encoded JSON value. If the type doesn’t have a `MarshalJSON()` method, then Go will fall back to trying to encode it to JSON based on its own internal set of rules. So, if we want to customize how something is encoded, all we need to do is implement a `MarshalJSON()` method on it which returns a custom JSON representation of itself in a `[]byte` slice.  

> Behind the scenes `time.Time` is actually a struct, but it has a `MarshalJSON()` method which outputs a RFC 3339 format representation of itself. This is what gets called whenever a `time.Time` value is encoded to JSON.  

#### ➤ 例子，把整数 `102` 序列化成 `"102 mins"` 这种格式的字符串

A clean and simple approach is to create a custom type specifically for the `Runtime` field, and implement a `MarshalJSON()` method on this custom type. To prevent our `internal/data/movie.go` file from getting cluttered, let’s create a new file to hold the logic for the `Runtime` type:

```go
// touch internal/data/runtime.go

type Runtime int32

func (r Runtime) MarshalJSON() ([]byte, error) {
    str := fmt.Sprintf("%d mins", r)
    quoted := strconv.Quote(str) // 别忘了在 json 语法中, 字符串需要用双引号括起来
    return []byte(quoted), nil
}

type Movie struct {
    Runtime   Runtime   `json:"runtime,omitempty"` // 字段类型从 int32 改成 Runtime, omitempty 依然生效因为它不是结构体
}

func TestRuntime_MarshalJSON(t *testing.T) {
    // 注意实现 MarshalJSON 方法时要用 value receiver, 这样 r 和 &r 都会按我们的逻辑序列化
    // 如果用了 pointer receiver, 由于 r 是值, 它的方法集中不包含指针方法, json.Marshal(r) 会执行默认序列化
    r := Runtime(102)
    s, err := json.Marshal(r)
    fmt.Println(string(s), err)

    s, err = json.Marshal(&r)
    fmt.Println(string(s), err)
}
```

#### ➤ 这种做法的优缺点

All in all, this is a pretty nice way to generate custom JSON. Our code is succinct and clear, and we’ve got a custom `Runtime` type that we can use wherever and whenever we need it. But there is a downside. It’s important to be aware that using custom types can sometimes be awkward when integrating your code with other packages, and you may need to perform type conversions to change your custom type to and from a value that the other packages understand and accept.

#### ➤ 能实现同样效果的替代做法

(1) Customizing the `Movie` struct. Instead of creating a custom `Runtime` type, we could have implemented a `MarshalJSON()` method on our Movie struct and customized the whole thing.

```go
type Movie struct {
    ID        int64 // 注意 Movie 结构体不包含任何 struct tag
    CreatedAt time.Time
    Title     string
    Year      int32
    Runtime   int32
    Genres    []string
    Version   int32
}

func (m Movie) MarshalJSON() ([]byte, error) {
    // 在 Movie 上实现 json.Marshaler 而不是在各个字段上实现
    var runtime string
    if m.Runtime != 0 {
        runtime = fmt.Sprintf("%d mins", m.Runtime)
    }

    // 这个匿名结构体用来自定义 JSON 序列化
    s := struct {
        ID      int64    `json:"id"`
        Title   string   `json:"title"`
        Year    int32    `json:"year,omitempty"`
        Runtime string   `json:"runtime,omitempty"` // This is a string.
        Genres  []string `json:"genres,omitempty"`
        Version int32    `json:"version"`
    }{
        ID:      m.ID,
        Title:   m.Title,
        Year:    m.Year,  // 需要逐个字段填充数据
        Runtime: runtime, // 使用前面我们格式化好的字符串
        Genres:  m.Genres,
        Version: m.Version,
    }
    // Encode the anonymous struct to JSON, and return it.
    return json.Marshal(s)
}
```

(2) The downside of the approach above is that the code feels quite verbose and repetitive. To reduce duplication, instead of writing out all the struct fields long-hand it’s possible to embed an alias of the `Movie` struct in the anonymous struct. Like so:

```go
type Movie struct {
    ID        int64     `json:"id"`
    CreatedAt time.Time `json:"-"`
    Title     string    `json:"title"`
    Year      int32     `json:"year,omitempty"`
    Runtime   int32     `json:"-"` // 把需要特殊处理的 Runtime 字段直接排除了
    Genres    []string  `json:"genres,omitempty"`
    Version   int32     `json:"version"`
}

func (m Movie) MarshalJSON() ([]byte, error) {
    var runtime string
    if m.Runtime != 0 {
        runtime = fmt.Sprintf("%d mins", m.Runtime)
    }

    type MovieAlias Movie // 用来避免死循环

    s := struct {
        MovieAlias
        Runtime string `json:"runtime,omitempty"`
    }{
        MovieAlias: MovieAlias(m),
        Runtime:    runtime,
    }

    // 使用 type 搞出来的 MovieAlias 不包含任何方法, 所以序列化 s 中的 MovieAlias 字段会使用默认序列化
    // 也就是把 MovieAlias 逐字段序列化呗, 然后注意 MovieAlias 和 Movie 拥有相同的字段定义 ( 包括注解 )
    // 所以 MovieAlias 中的 Runtime 字段会被排除, 然后轮到匿名结构体中额外添加的 Runtime 字段登场
    return json.Marshal(s)
}
```

On one hand, this approach is nice because it drastically cuts the number of lines of code and reduces repetition. And if you have a large struct and only need to customize a couple of fields it can be a good option. But it’s not without some downsides.

- This technique feels like a bit of a 'trick'. It’s more clever and less clear than the first approach.
- You lose granular control over the ordering of fields in the JSON response.

### 自定义 JSON 反序列化

#### ➤ 目标: 把字符串 "107 mins" 解析成整数 107

When Go is decoding some JSON, it will check to see if the destination type satisfies the `json.Unmarshaler` interface. If it *does* satisfy the interface, then Go will call it’s `UnmarshalJSON()` method to determine how to decode the provided JSON into the target type. This is basically the reverse of the `json.Marshaler` interface that we used earlier to customize our JSON encoding behavior.

```go
var ErrInvalidRuntimeFormat = errors.New("invalid runtime format")

type Runtime int32

func (r *Runtime) MarshalJSON() ([]byte, error) {
    str := fmt.Sprintf("%d mins", r)
    quoted := strconv.Quote(str) // 别忘了在 JSON 语法中, 字符串需要用双引号括起来
    return []byte(quoted), nil
}

func (r *Runtime) UnmarshalJSON(bytes []byte) error {
    // 输入是 "107 mins" 所以先去掉两端双引号, 若失败则说明格式错误
    unquoted, err := strconv.Unquote(string(bytes))
    if err != nil {
        return ErrInvalidRuntimeFormat
    }

    // 进一步检查格式是否正确
    parts := strings.Split(unquoted, " ")
    if len(parts) != 2 || parts[1] != "mins" {
        return ErrInvalidRuntimeFormat
    }

    i, err := strconv.ParseInt(parts[0], 10, 32)
    if err != nil {
        return ErrInvalidRuntimeFormat
    }

    *r = Runtime(i)
    return nil
}
```

## JSON 错误

### JSON 错误消息

#### ➤ 遇到错误返回 plain-text 不太好，所以封装 helper 函数，用于返回 JSON 格式的错误响应

At this point our API is sending nicely formatted JSON responses for successful requests, but if a client makes a bad request — or something goes wrong in our application — we’re still sending them a plain-text error message from the `http.Error()` and `http.NotFound()` functions. In this chapter we’ll fix that by creating some additional helpers to manage errors and send the appropriate JSON responses to our clients. Create a new `cmd/api/errors.go` file:

```go
func (app *application) logError(r *http.Request, err error) {
    // 记录错误日志, 可添加额外信息, 比如请求的 URI 和 Method
    var (
        method = r.Method
        uri    = r.URL.RequestURI()
    )
    app.logger.Error(err.Error(), "method", method, "uri", uri)
}


func (app *application) errorResponse(w http.ResponseWriter, r *http.Request, status int, message any) {
    // 返回客户端 JSON 格式的错误响应
    data := envelop{"error": message}

    err := app.writeJSON(w, status, data, nil)
    if err != nil {
        // 如果 message 序列化错误, 或者进行网络 IO 时遇到未知错误, 返回空的 500 响应
        app.logError(r, err)
        w.WriteHeader(http.StatusInternalServerError)
    }
}

func (app *application) serverErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
    // 遇到服务器异常则记录错误日志, 返回客户端 JSON 响应
    app.logError(r, err)
    message := "the server encountered a problem and could not process your request"
    app.errorResponse(w, r, http.StatusInternalServerError, message)
}

func (app *application) notFoundResponse(w http.ResponseWriter, r *http.Request) {
    message := "the requested resource could not be found"
    app.errorResponse(w, r, http.StatusNotFound, message)
}

func (app *application) methodNotAllowedResponse(w http.ResponseWriter, r *http.Request) {
    message := fmt.Sprintf("the %s method is not supported for this resource", r.Method)
    app.errorResponse(w, r, http.StatusMethodNotAllowed, message)
}
```

Let’s update our API handlers to use these new helpers instead of the `http.Error()` and `http.NotFound()` functions. Like so:  

```go
func (app *application) showMovieHandler(w http.ResponseWriter, r *http.Request) {
    id, err := app.readIDParam(r)
    if err != nil {
        app.notFoundResponse(w, r) // Use the new notFoundResponse() helper
        return
    }

    movie := data.Movie{
        ID:        id,
        Title:     "Big Hero 6",
    }

    err = app.writeJSON(w, http.StatusOK, envelop{"movie": movie}, nil)
    if err != nil {
        app.serverErrorResponse(w, r, err) // Use the new serverErrorResponse() helper
    }
}
```

#### ➤ 自定义 httprouter 的 error hanlder

Any error messages that our own API handlers send will now be well-formed JSON responses. Which is great! But what about the error messages that `httprouter` automatically sends when it can’t find a matching route? By default, these will still be the same plain-text (non-JSON) responses that we saw earlier in the book. Fortunately, `httprouter` allows us to set our own custom error handlers when we initialize the router. These custom handlers must satisfy the `http.Handler` interface.

```go
func (app *application) routes() *httprouter.Router {
    router := httprouter.New()
    router.NotFound = http.HandlerFunc(app.notFoundResponse)
    router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)

    router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)
    router.HandlerFunc(http.MethodPost, "/v1/movies", app.createMovieHandler)
    router.HandlerFunc(http.MethodGet, "/v1/movies/:id", app.showMovieHandler)
    return router
}
```

#### ➤ 一些情况下 HTTP Server 还是会生成 plain-text 错误响应

Go’s `http.Server` may still automatically generate and send plain-text HTTP responses. These scenarios include when:

- The HTTP request specifies an unsupported HTTP protocol version.
- The HTTP request contains a missing or invalid `Host` header, or multiple `Host` headers.
- The HTTP request contains an invalid header name or value.
- The HTTP request contains an unsupported `Transfer-Encoding` header.
- The size of the HTTP request headers exceeds the server’s `MaxHeaderBytes` setting.
- The client makes a HTTP request to a HTTPS server.  

For example, if we try sending a request with an invalid `Host` header value we will get a response like this:  

```bash
# ~ curl -i -H "Host: 未知" http://localhost:4000/v1/healthcheck
HTTP/1.1 400 Bad Request: malformed Host header
Content-Type: text/plain; charset=utf-8
Connection: close

400 Bad Request: malformed Host header⏎
```

Unfortunately, these responses are hard-coded into the Go standard library, and there’s nothing we can do to customize them to use JSON instead. But while this is something to be aware of, it’s not necessarily something to worry about. ( 因为一般不会遇到，遇到了也可以忽略，算不上问题 )

### Panic 异常恢复

#### ➤ 避免用户啥也看不到，空响应 ( 显示该网页无法正常运作 )

At the moment any panics in our API handlers will be recovered automatically by Go’s `http.Server`. This will unwind the stack for the affected goroutine (calling any deferred functions along the way), close the underlying HTTP connection, and log an error message and stack trace. This behavior is *OK*, but it would be better for the client if we could also send a `500 Internal Server Error` response to explain that something has gone wrong — rather than just closing the HTTP connection with no context.

#### ➤ recoverPanic() 中间件

```go
func (app *application) recoverPanic(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        defer func() {
            if err := recover(); err != nil {
                // Make Go's HTTP server automatically close the current connection after a response has been sent.
                w.Header().Set("Connection", "close")

                // 这样不会打印异常堆栈, 如果想记录 stack trace 可参考这里: https://stackoverflow.com/a/52106370
                // Log the error using our custom Logger type at the ERROR level and send the client a 500 response.
                app.serverErrorResponse(w, r, fmt.Errorf("%s", err))
            }
        }()

        next.ServeHTTP(w, r)
    })
}

func (app *application) routes() http.Handler {
    router := httprouter.New()
    router.NotFound = http.HandlerFunc(app.notFoundResponse)
    router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)
    router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)

    // Wrap the router with the panic recovery middleware.
    return app.recoverPanic(router)
}
```

#### ➤ 注意其他 goroutine 的 panic 会让服务器挂掉

It’s really important to realize that our middleware will only recover panics that happen in the *same goroutine that executed the `recoverPanic()` middleware*. If, for example, you have a handler which spins up another goroutine (e.g. to do some background processing), then any panics that happen in the background goroutine will not be recovered — not by the `recoverPanic()` middleware… and not by the panic recovery built into `http.Server`. These panics will cause your application to exit and bring down the server.

So, if you are spinning up additional goroutines from within your handlers and there is any chance of a panic, you *must* make sure that you recover any panics from within those goroutines too.

```go
func myHandler(w http.ResponseWriter, r *http.Request) {
    // Spin up a new goroutine to do some background processing.
    go func() {
        defer func() { // 使用 panic recover 避免后台 goroutine 把服务搞挂掉
            if err := recover(); err != nil {
                log.Println(fmt.Errorf("%s\n%s", err, debug.Stack()))
            }
        }()
        doSomeBackgroundProcessing()
    }
    w.Write([]byte("OK"))
}
```

## JSON 请求

### 解析 JSON 请求

If a client wanted to add a record for the movie Moana to our API, they would send a request body similar to this:  

```http
### Create a Movie
POST http://localhost:4000/v1/movies
Content-Type: application/json

{
  "title": "Your Name",
  "year": 2016,
  "runtime": 106,
  "genres": [
    "Animation",
    "Drama",
    "Fantasy",
    "Romance"
  ]
}
```

#### ➤ 读取 request body 并解析到结构体

Just like JSON encoding, there are two approaches that you can take to decode JSON into a native Go object: using a `json.Decoder` type or using the `json.Unmarshal()` function. Both approaches have their pros and cons, but for the purpose of decoding JSON from a HTTP request body, using `json.Decoder` is generally the best choice. It’s more efficient than `json.Unmarshal()`, requires less code, and offers some helpful settings that you can use to tweak its behavior.  

```go
func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
    // 用匿名结构体存储 request body 的信息
    var input struct {
        Title   string   `json:"title"`
        Year    int32    `json:"year"`    // 与外界交换数据 ( 比如序列化和 IO )
        Runtime int32    `json:"runtime"` // 要用确定的 int32 而不是大小模糊的 int
        Genres  []string `json:"genres"`
    }

    // Decode 时的一些注意事项见下文
    err := json.NewDecoder(r.Body).Decode(&input)
    if err != nil {
        app.errorResponse(w, r, http.StatusBadRequest, err.Error())
        return
    }

    _, _ = fmt.Fprintf(w, "%+v\n", input)
}
```

```go
func (app *application) useUnmarshal(w http.ResponseWriter, r *http.Request) {
    // 可见 json.Unmarshal 比 json.NewDecoder 稍微多一两行代码:
    body, err := io.ReadAll(r.Body)
    if err != nil {
        app.serverErrorResponse(w, r, err)
        return
    }

    err = json.Unmarshal(body, &input)
    if err != nil {
        app.errorResponse(w, r, http.StatusBadRequest, err.Error())
        return
    }
}
```

> 反序列化 JSON 时的一些注意事项:
>
> 1. 调用 `Decode()` 时要传 non-nil pointer 否则会返回 json.InvalidUnmarshalError
> 2. 无论 encoding 或 decoding，结构体字段必须被导出 ( 首字母大写 ) 才能对 `encoding/json` 包可见
> 3. 使用 `json:"title"` 设置 JSON 字段名能方便重构，无论重命名 `Title` 还是 `json:"title"` 都不破坏映射关系
> 4. 一般来说别忘记关闭 `Body`，但 `http.Server` 会帮我们关闭 `r.Body` 所以不必自己关
>
> 几条反序列化规则:
>
> 1. [字段名匹配不区分大小写](https://stackoverflow.com/questions/49006073/json-unmarshal-struct-case-sensitively)，JSON 中的 `"title", "Title", "TITLE"` 字段都能匹配 `json:"title"` 或 `json:"TITLE"`
> 2. ( JSON 中的 ) 无法被映射到结构体的字段会被忽略，比如 `"foo": ...` 字段会被忽略
> 3. JSON 中不提供 `"year":` 字段时，结构体中的 `Year` 字段会保持零值 ( 想区分 `"year":0` 和不提供此字段可用 `*int32` 指针 )



#### ➤ 分别处理解析 request 时遇到的错误

- 请求可能包含各种错误 ( 一串乱码、格式不符合 JSON、类型不匹配 )，这些情况下客户端会收到 `Decode()` 返回的错误消息
- 如果是内部 API 这差不多足够了，如果是公开 API 则还不够好，错误消息有时太详细暴露实现细节，有时太模糊缺乏指导性

To improve this, we’re going to explain how to triage the errors returned by `Decode()` and replace them with clearer, easy-to-action, error messages to help the client debug exactly what is wrong with their JSON. At this point in our application build, the `Decode()` method could potentially return the following five types of error:  

- `json.SyntaxError`、`io.ErrUnexpectedEOF`，There is a syntax problem with the JSON being decoded.
- `json.UnmarshalTypeError`，A JSON value is not appropriate for the destination Go type.
- `json.InvalidUnmarshalError`，The decode destination is not valid (usually because it is not a pointer). This is actually a problem with our application code, not the JSON itself.
- `io.EOF`，The JSON being decoded is empty.

Triaging these potential errors (which we can do using Go’s `errors.Is()` and `errors.As()` functions) is going to make the code in our `createMovieHandler` a lot longer and more complicated. And the logic is something that we’ll need to duplicate in other handlers throughout this project too. So, to assist with this, let’s create a new `readJSON()` helper in the `cmd/api/helpers.go` file.

```go
func (app *application) readJSON(w http.ResponseWriter, r *http.Request, dst any) error {
    // 解析请求中的 JSON, 对可能的错误进行分类, 分别返回对应的错误消息
    err := json.NewDecoder(r.Body).Decode(dst)
    if err != nil {
        var syntaxError *json.SyntaxError
        var unmarshalTypeError *json.UnmarshalTypeError
        var invalidUnmarshalError *json.InvalidUnmarshalError

        switch {
        // 比如少了个逗号
        case errors.As(err, &syntaxError):
            return fmt.Errorf("body contains badly-formed JSON (at character %d)", syntaxError.Offset)

        // 比如缺少结尾的 "}"
        case errors.Is(err, io.ErrUnexpectedEOF):
            return errors.New("body contains badly-formed JSON")

        // 比如把 "abc" 解析成 int32
        case errors.As(err, &unmarshalTypeError):
            if unmarshalTypeError.Field != "" {
                return fmt.Errorf("body contains incorrect JSON type for field %q", unmarshalTypeError.Field)
            }
            return fmt.Errorf("body contains incorrect JSON type (at character %d)", unmarshalTypeError.Offset)

        // 比如 body 为空
        case errors.Is(err, io.EOF):
            return errors.New("body must not be empty")

        // 我们代码有错误, 比如传给 Decode() 的参数: ①不是指针 ②是 nil 指针
        // 为什么这里用 panic 是合适的? 因为我们代码有 bug 所以用 panic 很正常啊
        case errors.As(err, &invalidUnmarshalError):
            panic(err)

        // 其他未知错误原样返回
        default:
            return err
        }
    }
    return nil
}

func (app *application) badRequestResponse(w http.ResponseWriter, r *http.Request, err error) {
    // 这个很简单的函数有两个意义: ①统一的 bad request 处理 ②可读性比下面这行代码稍好一点
    app.errorResponse(w, r, http.StatusBadRequest, err.Error())
}


func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
    // 用匿名结构体存储 request body 的信息
    var input struct {
        Title   string   `json:"title"`
        Year    int32    `json:"year"`    // 与外界交换数据 ( 比如序列化和 IO )
        Runtime int32    `json:"runtime"` // 要用确定的 int32 而不是大小模糊的 int
        Genres  []string `json:"genres"`
    }

    // 使用封装好的 helper 函数解析 JSON
    err := app.readJSON(w, r, &input)
    if err != nil {
        app.badRequestResponse(w, r, err)
        return
    }

    _, _ = fmt.Fprintf(w, "%+v\n", input)
}
```

#### ➤ 限制 JSON 请求 ( 大小、未知字段、单个对象 )

- 默认会忽略掉未知字段，可以用 `json.Decoder` 的 `DisallowUnknownFields()` 禁用未知字段
- `json.Decoder` 支持解析 JSON 流，每调用一次 `Decode()` 则解析一个对象，用户可以像这样 `{}{}` 传多个对象，可限制成一个
- 目前没有限制 request body 的大小，可以轻松构造一个很大的请求来折磨服务器，可用 `http.MaxBytesReader()` 限制 body 大小  
  可以用命令 curl -d @/tmp/largefile.json localhost:4000/v1/movies 测试发送大 JSON

```go
func (app *application) readJSON(w http.ResponseWriter, r *http.Request, dst any) error {
    // 限制 request body 大小为 1 MB
    maxBytes := 1_048_576
    r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

    // 禁用 unknown fields
    dec := json.NewDecoder(r.Body)
    dec.DisallowUnknownFields()

    // 解析请求中的 JSON, 对可能的错误进行分类, 分别返回对应的错误消息
    err := dec.Decode(dst)
    if err != nil {
        var maxBytesError *http.MaxBytesError

        switch {
        // 略...

        // 请求中包含 unknown fields, 要从错误消息中提取字段名 json: unknown field "<name>"
        // 有点丑陋是吧, 理论上做成一个错误类型更好, 参见 https://github.com/golang/go/issues/29035
        case strings.HasPrefix(err.Error(), "json: unknown field "):
            fieldName := strings.TrimPrefix(err.Error(), "json: unknown field ")
            return fmt.Errorf("body contains unknown field %s", fieldName)

        // 请求 body 超过了 1 MB
        case errors.As(err, &maxBytesError):
            return fmt.Errorf("body must not be larger than %d bytes", maxBytesError.Limit)

        default:
            return err
        }
    }

    // 第二次调用 Decode() 没返回 io.EOF 说明请求中包含多个 JSON 对象
    err = dec.Decode(&struct{}{})
    if !errors.Is(err, io.EOF) {
        return errors.New("body must only contain a single JSON value")
    }
    return nil
}
```

### 校验 JSON 请求

#### ➤ 处理请求前，先校验请求数据，校验不通过返回 422 响应

In many cases, you’ll want to perform additional validation checks on the data from a client to make sure it meets your specific business rules before processing it.

- The movie title provided by the client is not empty and is not more than 500 bytes long.
- The movie year is not empty and is between 1888 and the current year.
- The movie runtime is not empty and is a positive integer.
- The movie has between one and five (unique) genres.

If any of those checks fail, we want to send the client a `422 Unprocessable Entity` response along with error messages which clearly describe the validation failures.

#### ➤ 编写 validator 包

- 使用 map 存储字段错误消息
- 可重用的校验逻辑则做成函数

```go
// mkdir internal/validator
// touch internal/validator/validator.go
// EmailRX from https://html.spec.whatwg.org/#valid-e-mail-address
var (
    EmailRX = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
)

type Validator struct {
    Errors map[string]string
}

func New() *Validator {
    return &Validator{Errors: make(map[string]string)}
}

func (v *Validator) Valid() bool {
    return len(v.Errors) == 0
}

func (v *Validator) AddError(key, message string) {
    if _, exists := v.Errors[key]; !exists {
        v.Errors[key] = message
    }
}

func (v *Validator) Check(ok bool, key, message string) {
    if !ok {
        v.AddError(key, message)
    }
}

func PermittedValue[T comparable](value T, permittedValues ...T) bool {
    return slices.Contains(permittedValues, value)
}

func Matches(value string, rx *regexp.Regexp) bool {
    return rx.MatchString(value)
}

func Unique[T comparable](values []T) bool {
    uniqueValues := make(map[T]bool)
    for _, value := range values {
        uniqueValues[value] = true
    }
    return len(values) == len(uniqueValues)
}
```

#### ➤ 使用 validator 包

The first thing we need to do is update our `cmd/api/errors.go` file to include a new `failedValidationResponse()` helper, which writes a `422 Unprocessable Entity` and the contents of the errors map from our new `Validator` type as a JSON response body.

```go
func (app *application) failedValidationResponse(w http.ResponseWriter, r *http.Request, errors map[string]string) {
    app.errorResponse(w, r, http.StatusUnprocessableEntity, errors)
}
```

Head back to your `createMovieHandler`:

```go
func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
    // 用匿名结构体存储 request body 的信息
    var input struct {
        Title   string       `json:"title"`
        Year    int32        `json:"year"`    // 与外界交换数据 ( 比如序列化和 IO )
        Runtime data.Runtime `json:"runtime"` // 要用确定的 int32 而不是大小模糊的 int
        Genres  []string     `json:"genres"`
    }

    // 使用封装好的 helper 函数解析 JSON
    err := app.readJSON(w, r, &input)
    if err != nil {
        app.badRequestResponse(w, r, err)
        return
    }

    // 校验请求
    v := validator.New()
    v.Check(input.Title != "", "title", "must be provided")
    v.Check(len(input.Title) <= 500, "title", "must not be more than 500 bytes long")

    v.Check(input.Year != 0, "year", "must be provided")
    v.Check(input.Year >= 1888, "year", "must be greater than 1888")
    v.Check(input.Year <= int32(time.Now().Year()), "year", "must not be in the future")

    v.Check(input.Runtime != 0, "runtime", "must be provided")
    v.Check(input.Runtime > 0, "runtime", "must be a positive integer")

    v.Check(input.Genres != nil, "genres", "must be provided")
    v.Check(len(input.Genres) >= 1, "genres", "must contain at least 1 genre")
    v.Check(len(input.Genres) <= 5, "genres", "must not contain more than 5 genres")
    v.Check(validator.Unique(input.Genres), "genres", "must not contain duplicate values")

    if !v.Valid() {
        app.failedValidationResponse(w, r, v.Errors)
        return
    }

    _, _ = fmt.Fprintf(w, "%+v\n", input)
}
```

#### ➤ 重用校验逻辑

In large projects it’s likely that you’ll want to reuse some of the same validation checks in multiple places. In our case — for example — we’ll want to use many of these same checks later when a client *edits* the movie data. To prevent duplication, we can collect the validation checks for a movie into a standalone `ValidateMovie()` function. In theory this function could live almost anywhere in our codebase — next to the handlers in the `cmd/api/movies.go` file. But personally, I like to keep the validation checks close to the relevant domain type in the `internal/data` package.

```go
func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {

    // 复制字段有点麻烦, 为啥不直接解码到 data.Movie ?
    // 因为用户可以设置 id/version 等无权编辑的字段, 另外 input 和 data.Movie 各有各的职责和结构
    movie := &data.Movie{
        Title:   input.Title,
        Year:    input.Year,
        Runtime: input.Runtime,
        Genres:  input.Genres,
    }

    // 校验请求
    v := validator.New()
    if data.ValidateMovie(v, movie); !v.Valid() {
        app.failedValidationResponse(w, r, v.Errors)
        return
    }

}

// 让 Movie 的校验靠近 Movie 的定义, 把这个函数放在 type Movie struct {...} 下面
func ValidateMovie(v *validator.Validator, movie *Movie) {
    v.Check(movie.Title != "", "title", "must be provided")
    v.Check(len(movie.Title) <= 500, "title", "must not be more than 500 bytes long")

    v.Check(movie.Year != 0, "year", "must be provided")
    v.Check(movie.Year >= 1888, "year", "must be greater than 1888")
    v.Check(movie.Year <= int32(time.Now().Year()), "year", "must not be in the future")

    v.Check(movie.Runtime != 0, "runtime", "must be provided")
    v.Check(movie.Runtime > 0, "runtime", "must be a positive integer")

    v.Check(movie.Genres != nil, "genres", "must be provided")
    v.Check(len(movie.Genres) >= 1, "genres", "must contain at least 1 genre")
    v.Check(len(movie.Genres) <= 5, "genres", "must not contain more than 5 genres")
    v.Check(validator.Unique(movie.Genres), "genres", "must not contain duplicate values")
}
```

