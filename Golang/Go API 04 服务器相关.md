## Table of Contents
  - [CORS](#CORS)
    - [概述](#%E6%A6%82%E8%BF%B0)
    - [处理简单请求](#%E5%A4%84%E7%90%86%E7%AE%80%E5%8D%95%E8%AF%B7%E6%B1%82)
    - [处理预检请求](#%E5%A4%84%E7%90%86%E9%A2%84%E6%A3%80%E8%AF%B7%E6%B1%82)
  - [限流](#%E9%99%90%E6%B5%81)
    - [意义](#%E6%84%8F%E4%B9%89)
    - [令牌桶算法](#%E4%BB%A4%E7%89%8C%E6%A1%B6%E7%AE%97%E6%B3%95)
    - [全局限流器](#%E5%85%A8%E5%B1%80%E9%99%90%E6%B5%81%E5%99%A8)
    - [基于 IP 限流](#%E5%9F%BA%E4%BA%8E-IP-%E9%99%90%E6%B5%81)
    - [支持命令行配置](#%E6%94%AF%E6%8C%81%E5%91%BD%E4%BB%A4%E8%A1%8C%E9%85%8D%E7%BD%AE)
  - [优雅关闭服务](#%E4%BC%98%E9%9B%85%E5%85%B3%E9%97%AD%E6%9C%8D%E5%8A%A1)
    - [意义](#%E6%84%8F%E4%B9%89)
    - [四个 Signal](#%E5%9B%9B%E4%B8%AA-Signal)
    - [拦截 Signal](#%E6%8B%A6%E6%88%AA-Signal)
    - [实现 Graceful Shutdown](#%E5%AE%9E%E7%8E%B0-Graceful-Shutdown)
  - [Metrics](#Metrics)
    - [概述](#%E6%A6%82%E8%BF%B0)
    - [使用 expvar](#%E4%BD%BF%E7%94%A8-expvar)
    - [Add Metrics](#Add-Metrics)
    - [Dynamic Metrics](#Dynamic-Metrics)
    - [Protecting the metrics endpoint](#Protecting-the-metrics-endpoint)
    - [Request-level Metrics](#Requestlevel-Metrics)
    - [Recording HTTP Status Codes](#Recording-HTTP-Status-Codes)

## CORS

### 概述

#### ➤ 参考 [跨域资源共享 CORS 详解](https://www.ruanyifeng.com/blog/2016/04/cors.html)

- 简单情况下，浏览器中的 js 只会请求同一个源中的内容，同源是指 「 协议+域名+端口号 」 都相同  
  例如 `http://` 和 `https://` 协议不同，然后 `bing.com` 和 `www.bing.com` 域名不同，所以都不算同源
- 跨源请求就是读写另一个应用/别人的资源，浏览器会拦截跨源的 js 请求，除非服务端明示浏览器我支持 CORS
- 整个 CORS 通信过程，都是浏览器自动完成，不需要用户参与。对于开发者来说，CORS通信与同源的AJAX通信没有差别，代码完全一样。浏览器一旦发现AJAX请求跨源，就会自动添加一些附加的头信息，有时还会多出一次附加的请求，但用户不会有感觉。因此，实现CORS通信的关键是服务器。只要服务器实现了CORS接口，就可以跨源通信。

#### ➤ 只要同时满足以下两大条件，就属于简单请求

> （1) 请求方法是 HEAD、GET、POST 三种之一
>
> （2）HTTP 请求头只包含无法用 js 修改的 [forbidden headers](https://developer.mozilla.org/en-US/docs/Glossary/Forbidden_header_name) 和这四个:
>
> - Accept、Accept-Language、Content-Language、Content-Type
> - 并且 Content-Type：只限于三个值 `application/x-www-form-urlencoded`、`multipart/form-data`、`text/plain`
>
> 这是为了兼容表单（form），因为历史上表单一直可以发出跨域请求。  
> AJAX 的跨域设计就是，只要表单可以发，AJAX 就可以直接发。凡是不同时满足上面两个条件，就属于非简单请求。

#### ➤ [简单请求和非简单请求的流程](https://www.ruanyifeng.com/blog/2016/04/cors.html#:~:text=%E4%B8%89%E3%80%81%E7%AE%80%E5%8D%95%E8%AF%B7%E6%B1%82-,3.1%20%E5%9F%BA%E6%9C%AC%E6%B5%81%E7%A8%8B,-%E5%AF%B9%E4%BA%8E%E7%AE%80%E5%8D%95%E8%AF%B7%E6%B1%82) ( 先看一遍原文 )

- 浏览器会自动为跨源 AJAX 请求添加 `Origin` 字段，表示本次请求来自哪个源（协议 + 域名 + 端口）
- 如果服务器返回的响应不包含 `Access-Control-Allow-Origin` 或不匹配，浏览器就知道出错了，会把这个响应禁掉
- 常见的 json 请求属于非简单请求，浏览器会先发一个预检请求，再根据预检响应判断要不要拦截这个跨源请求

#### ➤ [跨域的请求在服务端会不会真正执行？](https://cloud.tencent.com/developer/article/2009296)

- 简单请求：不管是否跨域，只要发出去了，一定会到达服务端并被执行，浏览器只会隐藏返回值
- 复杂请求：先发预检，预检不会真正执行业务逻辑，预检通过后才会发送真正请求并在服务端被执行

#### ➤ 浏览器的同源策略

- A webpage on one origin can *embed* certain types of resources from another origin in their HTML — including images, CSS, and JavaScript files. `<img src="http://anotherorigin.com/example.png" alt="example image">`
- A webpage on one origin can *send* data to a different origin. For example, it’s OK for an HTML form in a webpage to submit data to a different origin. ( 复杂请求要先通过预检才能发出去 )
- But a webpage on one origin is *not* allowed to *receive* data from a different origin.

This key thing here is the final bullet-point: the same-origin policy prevents a (potentially malicious) website on another origin from *reading* (possibly confidential) information from your website. ( 避免自己网站的信息/资源被其他网站用 js 读取 )

It’s important to emphasize that cross-origin *sending of data* is not prevented by the same-origin policy, despite also being dangerous. In fact, this is why CSRF attacks are possible and why we need to take additional steps to prevent them — like using `SameSite` cookies and CSRF tokens.

### 处理简单请求

#### ➤ 添加 Header

用响应头 `Access-Control-Allow-Origin` 指定当前响应可共享给哪些 origin，另外通配符 `*` 能匹配所有 origin

```http
Access-Control-Allow-Origin: *
```

最宽松的策略，共享给所有 origin:

```go
func (app *application) enableCORS(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Access-Control-Allow-Origin", "*")
        next.ServeHTTP(w, r)
    })
}

func (app *application) routes() http.Handler {

    // enableCORS 要放在最靠前的位置, 避免其他中间件拦截了请求, 然后响应中没有 CORS Header
    return app.recoverPanic(app.enableCORS(app.rateLimit(app.authenticate(router))))
}
```

#### ➤ 指定 Origin

Using a wildcard to allow cross-origin requests, like we are in the code above, can be useful in certain circumstances (like when you have a completely public API with no access control checks). But more often you’ll probably want to restrict CORS to a much smaller set of *trusted origins*. For example, if you want to only allow CORS from the origin `https://www.example.com` you could send the following header in your responses:

```http
Access-Control-Allow-Origin: https://www.example.com
```

But if you need to support multiple trusted origins, or you want the value to be configurable at runtime, then things get a bit more complex. One of the problems is that — in practice — you can only specify *exactly one origin* in the `Access-Control-Allow-Origin` header. You can’t include a list of multiple origin values, separated by spaces or commas like you might expect.

To work around this limitation, you’ll need to update your `enableCORS()` middleware to check if the value of the `Origin` header matches one of your trusted origins. If it does, then you can reflect (or *echo*) that value back in the `Access-Control-Allow-Origin` response header.

#### ➤ Supporting multiple dynamic origins

Let’s update our API so that cross-origin requests are restricted to a list of trusted origins, configurable at runtime. The first thing we’ll do is add a new `-cors-trusted-origins` command-line flag to our API application, which we can use to specify the list of trusted origins at runtime. We’ll set this up so that the origins must be separated by a space character — like so:

```bash
go run ./cmd/api -cors-trusted-origins="https://www.example.com https://staging.example.com"
```

添加命令行参数:

```go
type config struct {

    cors struct {
        trustedOrigins []string
    }
}

func main() {
    var cfg config

    // 把 cors 参数解析成字符串切片, 注意 val 为空时会得到 empty slice
    flag.Func("cors-trusted-origins", "Trusted CORS origins (space separated)", func(val string) error {
        cfg.cors.trustedOrigins = strings.Fields(val)
        return nil
    })

    flag.Parse()
}

```

#### ➤ 如果是列表中的 origin 则添加 Access-Control-Allow-Origin 响应头，否则不加

A side effect of this is that the *response will be different depending on the origin that the request is coming from*. Specifically, the value of the `Access-Control-Allow-Origin` header may be different in the response, or it may not even be included at all. So because of this we should make sure to always set a `Vary: Origin` response header to warn any caches that the response may be different. This is actually really important, and it can be the cause of subtle bugs [like this one](https://textslashplain.com/2018/08/02/cors-and-vary/) if you forget to do it.

```go
func (app *application) enableCORS(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // 响应会根据请求中的 Origin 头发生改变, 所以用 Vary 配合缓存服务器
        w.Header().Add("Vary", "Origin")

        // 没有 Origin 请求头或 Origin 不在列表内则不加 Access-Control-Allow-Origin 响应头
        origin := r.Header.Get("Origin")
        if origin != "" {
            for i := range app.config.cors.trustedOrigins {
                if origin == app.config.cors.trustedOrigins[i] {
                    w.Header().Set("Access-Control-Allow-Origin", origin)
                    break
                }
            }
        }

        next.ServeHTTP(w, r)
    })
}
```

### 处理预检请求

- 简单请求的请求方法: HEAD / GET / POST
- 简单请求的请求头: [forbidden headers](https://developer.mozilla.org/en-US/docs/Glossary/Forbidden_header_name)，另外 `Content-Type` 的取值范围和表单的 `enctype` 一样

When a cross-origin request doesn’t meet these conditions, then the web browser will trigger an initial ‘preflight’ request *before the real request*. The purpose of this preflight request is to determine whether the *real* cross-origin request will be permitted or not.

如下是预检请求的几个重要字段:

```http
OPTIONS /v1/tokens/authentication HTTP/1.1
Origin: http://localhost:8000
Access-Control-Request-Method: POST
Access-Control-Request-Headers: content-type
Host: localhost:4000
```

- `Origin` — As we saw previously, this lets our API know what origin the preflight request is coming from.
- `Access-Control-Request-Method` — This lets our API know what HTTP method will be used for the real request (in this case, we can see that the real request will be a `POST`).
- `Access-Control-Request-Headers` — This lets our API know what HTTP headers will be sent with the real request (in this case we can see that the real request will include a `content-type` header). 
- It’s important to note that `Access-Control-Request-Headers` won’t list *all* the headers that the real request will use. Only headers that are *not* [CORS-safe](https://developer.mozilla.org/en-US/docs/Glossary/CORS-safelisted_request_header) or [forbidden](https://developer.mozilla.org/en-US/docs/Glossary/Forbidden_header_name) will be listed. If there are no such headers, then `Access-Control-Request-Headers` may be omitted from the preflight request entirely.

#### ➤ 判定预检请求

In order to respond to a preflight request, the first thing we need to do is identify that it *is* a preflight request — rather than just a regular (possibly even cross-origin) `OPTIONS` request.

To do that, we can leverage the fact that preflight requests always have three components: the HTTP method `OPTIONS`, an `Origin` header, and an `Access-Control-Request-Method` header. If any one of these pieces is missing, we know that it is not a preflight request.

Once we identify that it is a preflight request, we need to send a `200 OK` response with some special headers to let the browser know whether or not it’s OK for the real request to proceed. These are:

- An `Access-Control-Allow-Origin` response header, which reflects the value of the preflight request’s `Origin` header
- An `Access-Control-Allow-Methods` header listing the HTTP methods that can be used in real cross-origin requests to the URL.
- An `Access-Control-Allow-Headers` header listing the request headers that can be included in real cross-origin requests to the URL.

In our case, we could set the following response headers to allow cross-origin requests for *all our endpoints*:

```http
Access-Control-Allow-Origin: <reflected trusted origin>
Access-Control-Allow-Methods: OPTIONS, PUT, PATCH, DELETE
Access-Control-Allow-Headers: Authorization, Content-Type
```

> **Important:** When responding to a preflight request it’s not necessary to include the CORS-safe methods `HEAD`, `GET` or `POST` in the `Access-Control-Allow-Methods` header. Likewise, it’s not necessary to include forbidden or CORS-safe headers in `Access-Control-Allow-Headers`.

#### ➤ 具体代码

1. Set a `Vary: Access-Control-Request-Method` header on all responses, as the response will be different depending on whether or not this header exists in the request.
2. Check whether the request is a preflight cross-origin request or not. If it’s not, then we should allow the request to proceed as normal.
3. Otherwise, if it is a preflight cross-origin request, then we should add the `Access-Control-Allow-Method` and `Access-Control-Allow-Headers` headers as described above.

```go
func (app *application) enableCORS(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // 响应会根据请求中的 Origin 头发生改变
        w.Header().Add("Vary", "Origin")
        w.Header().Add("Vary", "Access-Control-Request-Method")

        // 没有 Origin 请求头或 Origin 不在允许列表内则当成普通请求( 而不是跨源请求 )进行处理
        origin := r.Header.Get("Origin")
        if origin != "" {
            for i := range app.config.cors.trustedOrigins {
                if origin == app.config.cors.trustedOrigins[i] {
                    // 处理跨源请求
                    w.Header().Set("Access-Control-Allow-Origin", origin)
                    // 处理预检请求
                    if r.Method == http.MethodOptions && r.Header.Get("Access-Control-Request-Method") != "" {
                        w.Header().Set("Access-Control-Allow-Methods", "OPTIONS, PUT, PATCH, DELETE")
                        w.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type")
                        // 提前返回无需进一步处理, 另外上面为啥不加 GET/POST/HEAD 方法? 因为默认允许无需刻意声明
                        w.WriteHeader(http.StatusOK)
                        return
                    }
                    break
                }
            }
        }

        next.ServeHTTP(w, r)
    })
}
```

## 限流

### 意义

If you’re building an API for public use, then it’s quite likely that you’ll want to implement some form of *rate limiting* to prevent clients from making *too many requests too quickly*, and putting excessive strain on your server. Essentially, we want a middleware to check how many requests have been received in the last ‘N’ seconds and — if there have been too many — then it should send the client a `429 Too Many Requests` response. We’ll position this middleware before our main application handlers, so that it carries out this check *before* we do any expensive processing like decoding a JSON request body or querying our database.

### 令牌桶算法

The description from the official `x/time/rate` documentation says:

- A Limiter controls how frequently events are allowed to happen. 
- It implements a “token bucket” of size `b`, initially full and refilled at rate `r` tokens per second.

简而言之:

- 桶的容量是 b，最多只能装 b 个 token
- 一开始桶是满的，然后每秒生成 r 个 token 放到桶里
- 每个请求消耗一个 token，如果请求到来时桶是空的，那么返回 429 Too Many Requests 响应

In practice this means that our application would allow a maximum ‘burst’ of `b` HTTP requests in quick succession, but over time it would allow an average of `r` requests per second. If we want to create a rate limiter which allows an average of 2 requests per second, with a maximum of 4 requests in a single ‘burst’, we could do so with the following code:

```go
limiter := rate.NewLimiter(2, 4) // 限流每秒 2 个请求，突然来 4 个请求也行，但得缓一缓才能应对下一次突然 4 个
```

### 全局限流器

#### ➤ 下载 x/time/rate 包

Let’s start by creating a single *global rate limiter* for our application. This will consider *all the requests* that our API receives (rather than having separate rate limiters for every individual client/IP). Instead of writing our own rate-limiting logic from scratch, which would be quite complex and time-consuming, we can leverage the [`x/time/rate`](https://pkg.go.dev/golang.org/x/time/rate) package to help us here. This provides a tried-and-tested implementation of a [*token bucket*](https://en.wikipedia.org/wiki/Token_bucket#Algorithm) rate limiter.

```bash
go get golang.org/x/time/rate
```

#### ➤ middleware 套路模板

```go
func (app *application) exampleMiddleware(next http.Handler) http.Handler {
    
    // 这个位置的代码只会执行一次，用来初始化中间件

    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        
        // 被中间件拦截到的请求都会执行这段代码，这里是中间件的具体逻辑

        next.ServeHTTP(w, r)
    })
}
```

#### ➤ rateLimit() 中间件

```go
func (app *application) rateLimit(next http.Handler) http.Handler {
    // 限流每秒 2 个请求，突然来 4 个请求也行，但得缓一缓才能应对下一次突然 4 个
    limiter := rate.NewLimiter(2, 4)
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // 调用 Allow 获取一个 token, 如果触发限流那么返回 429 Too Many Requests 响应
        if !limiter.Allow() {
            app.rateLimitExceededResponse(w, r)
            return // 另外注意 limiter 是并发安全的对象
        }
        next.ServeHTTP(w, r)
    })
}

func (app *application) rateLimitExceededResponse(w http.ResponseWriter, r *http.Request) {
    message := "rate limit exceeded"
    app.errorResponse(w, r, http.StatusTooManyRequests, message)
}

func (app *application) routes() http.Handler {
    // 中间件执行顺序是从左往右: recoverPanic > rateLimit > router
    return app.recoverPanic(app.rateLimit(router))
}
```

同时发几个请求试试:

```bash
# 使用 fish shell
for i in (seq 7); curl http://localhost:4000/v1/healthcheck; end
```

### 基于 IP 限流

#### ➤ 每个 IP 对应一个限流器，注意 map 不是并发安全的，别忘了加锁啊

Using a global rate limiter can be useful when you want to enforce a strict limit on the total rate of requests to your API, and you don’t care where the requests are coming from. But it’s generally more common to want an individual rate limiter for each client, so that one bad client making too many requests doesn’t affect all the others.

A conceptually straightforward way to implement this is to create an in-memory *map of rate limiters*, using the IP address for each client as the map key. Each time a new client makes a request to our API, we will initialize a new rate limiter and add it to the map. For any subsequent requests, we will retrieve the client’s rate limiter from the map and check whether the request is permitted by calling its `Allow()` method, just like we did before.

But there’s one thing to be aware of: *by default, maps are not safe for concurrent use*. This is a problem for us because our `rateLimit()` middleware may be running in multiple goroutines at the same time (remember, Go’s `http.Server` handles each HTTP request in its own goroutine). *Maps are not safe for concurrent use: it’s not defined what happens when you read and write to them simultaneously.*

#### ➤ 把前面的全局限流，改成基于 IP 限流

```go
func (app *application) rateLimit(next http.Handler) http.Handler {
    // 记录 client 上次访问时间
    type client struct {
        limiter  *rate.Limiter
        lastSeen time.Time
    }

    // 经常有不同的 IP 看一看就离开, 属于写多读少, 所以用互斥锁而不是读写锁
    var (
        mu      sync.Mutex
        clients = make(map[string]*client)
    )

    // 注意每次新 IP 访问都会导致 map 增长, 如果不清理只会越来越大, 你也不想浪费几十上百 MB 内存吧
    go func() {
        for {
            time.Sleep(time.Minute)
            mu.Lock()
            for ip, client := range clients {
                if time.Since(client.lastSeen) > 3*time.Minute {
                    delete(clients, ip)
                }
            }
            mu.Unlock()
        }
    }()

    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // 提取 IP
        ip, _, err := net.SplitHostPort(r.RemoteAddr)
        if err != nil {
            app.serverErrorResponse(w, r, err)
            return
        }

        mu.Lock()

        // 检查 IP 是否已经在 map 里面, 如果不在则初始化对应的 rate limiter
        if _, found := clients[ip]; !found {
            clients[ip] = &client{limiter: rate.NewLimiter(2, 4)}
        }

        clients[ip].lastSeen = time.Now()

        if !clients[ip].limiter.Allow() {
            mu.Unlock()
            app.rateLimitExceededResponse(w, r)
            return
        }

        mu.Unlock() // 这里别习惯性 defer mu.Unlock, 否则要等后面一堆 http.Handler 处理完才会释放锁

        next.ServeHTTP(w, r)
    })
}
```

#### ➤ 可用 nginx 或 redis 实现分布式限流

Using this pattern for rate-limiting will only work if your API application is running on a single-machine. If your infrastructure is distributed, with your application running on multiple servers behind a load balancer, then you’ll need to use an alternative approach.

If you’re using HAProxy or Nginx as a load balancer or reverse proxy, both of these have built-in functionality for rate limiting that it would probably be sensible to use. Alternatively, you could use a fast database like Redis to maintain a request count for clients, running on a server which all your application servers can communicate with.

### 支持命令行配置

At the moment our requests-per-second and burst values are hard-coded into the `rateLimit()` middleware. This is OK, but it would be more flexible if they were configurable at runtime instead. Likewise, it would be useful to have an easy way to turn off rate limiting altogether (which is useful when you want to run benchmarks or carry out load testing, when all requests might be coming from a small number of IP addresses).

```go
type config struct {
    limiter struct {
        rps     float64
        burst   int
        enabled bool
    }
}

func main() {
    // 限流器配置
    flag.Float64Var(&cfg.limiter.rps, "limiter-rps", 2, "Rate limiter maximum requests per second")
    flag.IntVar(&cfg.limiter.burst, "limiter-burst", 4, "Rate limiter maximum burst")
    flag.BoolVar(&cfg.limiter.enabled, "limiter-enabled", true, "Enable rate limiter")
    flag.Parse()
}

func (app *application) rateLimit(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        if app.config.limiter.enabled {
            // 开了限流器才执行这里面的逻辑
        }

        next.ServeHTTP(w, r)
    })
}
```

## 优雅关闭服务

### 意义

At the moment, when we stop our API application (usually by pressing `Ctrl+C`) it is terminated *immediately* with no opportunity for in-flight HTTP requests to complete. This isn’t ideal for two reasons:

- It means that clients won’t receive responses to their in-flight requests  
  — all they will experience is a hard closure of the HTTP connection.
- Any work being carried out by our handlers may be left in an incomplete state.

### 四个 Signal

When our application is running, we can terminate it at any time by sending it a specific [signal](https://en.wikipedia.org/wiki/Signal_(IPC)#POSIX_signals). A common way to do this, which you’ve probably been using, is by pressing `Ctrl+C` on your keyboard to send an *interrupt* signal — also known as a `SIGINT`. 

| Signal    | Description                          | Keyboard shortcut | Catchable |
| :-------- | :----------------------------------- | :---------------- | :-------- |
| `SIGINT`  | Interrupt from keyboard              | `Ctrl+C`          | Yes       |
| `SIGQUIT` | Quit from keyboard                   | `Ctrl+\`          | Yes       |
| `SIGKILL` | Kill process (terminate immediately) | -                 | No        |
| `SIGTERM` | Terminate process in orderly manner  | -                 | Yes       |

Catachable signals can be intercepted by our application and either ignored, or used to trigger a certain action (such as a graceful shutdown). Other signals, like `SIGKILL`, are not catchable and cannot be intercepted.

```bash
go run ./cmd/api     # 运行服务
pgrep -l api         # 搜索名为 api 的进程，返回 pid 和进程名，可确认进程是否存在
pkill -SIGKILL api   # 向 api 进程发送 SIGKILL 信号, 会发现服务退出了并打印 signal: killed
```

### 拦截 Signal

The next thing that we want to do is update our application so that it ‘catches’ any `SIGINT` and `SIGTERM` signals. As we mentioned above, `SIGKILL` signals are not catchable (and will always cause the application to terminate immediately), and we’ll leave `SIGQUIT` with its default behavior (as it’s handy if you want to execute a non-graceful shutdown via a keyboard shortcut).

To catch the signals, we’ll need to spin up a background goroutine which runs for the lifetime of our application. In this background goroutine, we can use the [`signal.Notify()`](https://godoc.org/os/signal#Notify) function to listen for specific signals and relay them to a channel for further processing.

```go
func (app *application) serve() error {
    // 创建 Server
    srv := &http.Server{
        Addr:         fmt.Sprintf(":%d", app.config.port),
        Handler:      app.routes(),
        IdleTimeout:  time.Minute,
        ReadTimeout:  5 * time.Second,
        WriteTimeout: 10 * time.Second,
        ErrorLog:     slog.NewLogLogger(app.logger.Handler(), slog.LevelError),
    }

    go func() {
        quit := make(chan os.Signal, 1)                        // 使用 buffered channel 且容量应该为 1
        signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)   // 监听和拦截这两个 Signal
        s := <-quit                                            // 等待 Signal 出现
        app.logger.Info("caught signal", "signal", s.String())
        os.Exit(0)
    }()

    app.logger.Info("starting server", "addr", srv.Addr, "env", app.config.env)

    return srv.ListenAndServe()
}
```

### 实现 Graceful Shutdown

After receiving one of these signals we will call the [`Shutdown()`](https://golang.org/pkg/net/http/#Server.Shutdown) method on our HTTP server. `Shutdown()` gracefully shuts down the server without interrupting any active connections. Shutdown works by first closing all open listeners, then closing all idle connections, and then waiting indefinitely for connections to return to idle and then shut down.

```go
func (app *application) serve() error {

    srv := &http.Server{
        Addr:         fmt.Sprintf(":%d", app.config.port),
        Handler:      app.routes(),
        IdleTimeout:  time.Minute,
        ReadTimeout:  5 * time.Second,
        WriteTimeout: 10 * time.Second,
        ErrorLog:     slog.NewLogLogger(app.logger.Handler(), slog.LevelError),
    }

    shutdownError := make(chan error)

    go func() {
        quit := make(chan os.Signal, 1)                               // 使用 buffered channel 且容量应该为 1
        signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)          // 监听和拦截这两个 Signal
        s := <-quit                                                   // 等待 Signal 出现
        app.logger.Info("shutting down server", "signal", s.String())

        // 设置 30 秒超时, 祈祷进行中的 request 能在时限内处理完毕
        ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
        defer cancel()

        // 如果一切正常、能优雅结束, 那么 Shutdown 方法会返回 nil
        shutdownError <- srv.Shutdown(ctx)
    }()

    app.logger.Info("starting server", "addr", srv.Addr, "env", app.config.env)

    // 调用 Shutdown 会让 ListenAndServe 立刻返回 http.ErrServerClosed, 如果不是这个说明出现其他错误
    err := srv.ListenAndServe()
    if !errors.Is(err, http.ErrServerClosed) {
        return err
    }

    // 等待 Shutdown 执行结束并接收它的返回值
    err = <-shutdownError
    if err != nil {
        return err
    }

    // 走到这证明 server 能优雅结束, 打条日志庆祝下
    app.logger.Info("stopped server", "addr", srv.Addr)
    return nil
}
```

It’s important to be aware that the `Shutdown()` method does not wait for any background tasks to complete, nor does it close hijacked long-lived connections like WebSockets. Instead, you will need to implement your own logic to coordinate a graceful shutdown of these things. We’ll look at some techniques for doing this later in the book.

加个 5 秒的睡眠测一测:

```go
func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
    time.Sleep(5 * time.Second)
}
```

测试命令:

```bash
curl localhost:4000/v1/healthcheck & sleep 1; pkill -SIGTERM api
```

## Metrics

### 概述

When your application is running in production with real-life traffic — or if you’re carrying out targeted load testing — you might want some up-close insight into how it’s performing and what resources it is using.

For example, you might want to answer questions like:

- How much memory is my application using? How is this changing over time?
- How many goroutines are currently in use? How is this changing over time?
- How many database connections are in use and how many are idle? Do I need to change the connection pool settings?
- What is the ratio of successful HTTP responses to both client and server errors? Are error rates elevated above normal?

Having insight into these things can help inform your hardware and configuration setting choices, and act as an early warning sign of potential problems (such as memory leaks). To assist with this, Go’s standard library includes the [`expvar`](https://golang.org/pkg/expvar/) package which makes it easy to collate and view different application metrics at runtime.

### 使用 expvar

Viewing metrics for our application is made easy by the fact that the `expvar` package provides an [`expvar.Handler()`](https://golang.org/pkg/expvar/#Handler) function which returns a *HTTP handler exposing your application metrics*. By default this handler displays information about memory usage, along with a reminder of what command-line flags you used when starting the application, all outputted in JSON format.

So the first thing that we’re going to do is mount this handler at a new `GET /debug/vars` endpoint, like so:

```go
func (app *application) routes() http.Handler {

    router.Handler(http.MethodGet, "/debug/vars", expvar.Handler())
}
```

And if you visit [`http://localhost:4000/debug/vars`](http://localhost:4000/debug/vars) in your web browser, you should see a JSON response containing information about your running application. We can see that the JSON here currently contains two top-level items: `"cmdline"` and `"memstats"`. Let’s quickly talk through what these represent.

The `"cmdline"` item contains an array of the command-line arguments used to run the application, beginning with the program name. This is essentially a JSON representation of the [`os.Args`](https://golang.org/pkg/os/#pkg-variables) variable, and it’s useful if you want to see exactly what non-default settings were used when starting the application.

The `"memstats"` item contains a ‘moment-in-time’ snapshot of memory usage, as returned by the [`runtime.MemStats()`](https://golang.org/pkg/runtime/#MemStats) function. Documentation and descriptions for all of the values can be [found here](https://golang.org/pkg/runtime/#MemStats), but the most important ones are:

- `TotalAlloc` — Cumulative bytes allocated on the heap (will not decrease).
- `HeapAlloc` — Current number of bytes on the heap.
- `HeapObjects` — Current number of objects on the heap.
- `Sys` — Total bytes of memory obtained from the OS (i.e. total memory reserved by the Go runtime for the heap, stacks, and other internal data structures).
- `NumGC` — Number of completed garbage collector cycles.
- `NextGC` — The target heap size of the next garbage collector cycle (Go aims to keep `HeapAlloc` ≤ `NextGC`).

### Add Metrics

The default information exposed by the `expvar` handler is a good start, but we can make it even more useful by exposing some additional custom metrics in the JSON response. To illustrate this, we’ll start really simple and first expose our application version number in the JSON. The code to do this breaks down into two basic steps: first we need to register a custom variable with the `expvar` package, and then we need to set the value for the variable itself. In one line, the code looks roughly like this:

```go
expvar.NewString("version").Set(version)
```

The first part of this — `expvar.NewString("version")` — creates a new [`expvar.String`](https://golang.org/pkg/expvar/#String) type, then *publishes it* so it appears in the `expvar` handler’s JSON response with the name `"version"`, and then returns a pointer to it. Then we use the `Set()` method on it to assign an actual value to the pointer.

Two other things to note:

- The `expvar.String` type is safe for concurrent use. So — if you want to — it’s OK to manipulate this value at runtime from your application handlers.
- If you try to register two `expvar` variables with the same name, you’ll get a runtime panic when the duplicate variable is registered.

```go
    // Publish a new "version" variable in the expvar handler
    expvar.NewString("version").Set(version)

    app := &application{
    }
```

You should now see a `"version": "1.0.0"` item in the JSON.

```json
{
  "cmdline": [...],
  "memstats": {...},
  "version": "1.0.0"
}
```

### Dynamic Metrics

Occasionally you might want to publish metrics which require you to call other code — or do some kind of pre-processing — to generate the necessary information. To help with this there is the [`expvar.Publish()`](https://golang.org/pkg/expvar/#Publish) function, which allows you to publish the *result of a function* in the JSON output. For example, if you want to publish the number of currently active goroutines from Go’s [`runtime.NumGoroutine()`](https://golang.org/pkg/runtime/#NumGoroutine) function, you could write the following code:

```go
expvar.Publish("goroutines", expvar.Func(func() any {
    return runtime.NumGoroutine()
}))
```

> It’s important to point out here that the `any` value returned from this function *must* encode to JSON without any errors. If it can’t be encoded to JSON, then it will be omitted from the `expvar` output and the response from the `GET /debug/vars` endpoint will be malformed. Any error will be silently discarded. In the case of the code snippet above, `runtime.NumGoroutine()` returns a regular `int` type — which will encode to a JSON number. So there’s no problem with that here.

OK, let’s add this code to our `main()` function, along with two other functions which:

- Publish information about the state of our database connection pool (such as the number of idle and in-use connections) via the [`db.Stats()`](https://golang.org/pkg/database/sql/#DB.Stats) method.
- Publish the current Unix timestamp with second precision.

```go
func main() {
    // 添加静态的或动态的指标
    expvar.NewString("version").Set(version)
    expvar.Publish("goroutines", expvar.Func(func() any {
        return runtime.NumGoroutine()
    }))
    expvar.Publish("database", expvar.Func(func() any {
        return db.Stats()
    }))
    expvar.Publish("timestamp", expvar.Func(func() any {
        return time.Now().Unix()
    }))

    app := &application{
    }
}
```

If you like, you can use a tool like [hey](https://github.com/rakyll/hey) to generate some requests to your application and see how these figures change under load. For example, you can send a batch of requests to the `POST /v1/tokens/authentication` endpoint (which is slow and costly because it checks a bcrypt-hashed password) like so:

```bash
set BODY '{"email": "tifa@ff7.com", "password": "my waifu"}'
hey.exe -n 800 -d "$BODY" -m "POST" http://localhost:4000/v1/tokens/authentication
```

If you visit the `GET /debug/vars` endpoint while the `hey` tool is running, you should see that your application metrics now look quite different:

```json
{
    "cmdline": [
        "-cors-trusted-origins=http://localhost:8000 http://localhost:8000",
        "-limiter-enabled=false"
    ],
    "database": {
        "MaxOpenConnections": 25,
        "OpenConnections": 25,
        "InUse": 25,
        "Idle": 0,
        "WaitCount": 1609,
        "WaitDuration": 56642081300,
        "MaxIdleClosed": 0,
        "MaxIdleTimeClosed": 0,
        "MaxLifetimeClosed": 0
    },
    "goroutines": 137,
    "memstats": {},
    "timestamp": 1724764091,
    "version": "1.0.0"
}
```

At the moment I took this screenshot, we can see that my API application had `137` active goroutines, with `25` database connections in use and `0` connections sat idle. 

The database `WaitCount` figure of `1609` is the total number of times that our application had to wait for a database connection to become available in our `sql.DB` pool (because all connections were in-use). Likewise, `WaitCountDuration` is the cumulative amount of time (in nanoseconds) spent waiting for a connection. From these, it’s possible to calculate that *when our application did have to wait for a database connection, the average wait time was approximately 35 milliseconds*. Ideally, you want to be seeing zeroes or very low numbers for these two things under normal load in production.

Also, the `MaxIdleTimeClosed` figure is the total count of the number of connections that have been closed because they reached their `ConnMaxIdleTime` limit (which in our case is set to 15 minutes by default). If you leave the application running but don’t use it, and come back in 15 minutes time, you should see that the number of open connections has dropped to zero and the `MaxIdleTimeClosed` count has increased accordingly.

### Protecting the metrics endpoint

It’s important to be aware that these metrics provide very useful information to anyone who wants to perform a denial-of-service attack against your application, and that the `"cmdline"` values may also expose potentially sensitive information (like a database DSN). So you should make sure to restrict access to the `GET /debug/vars` endpoint when running in a production environment.

One option is to leverage our existing authentication process and create a `metrics:view` permission so that only certain trusted users can access the endpoint. Another option would be to use HTTP Basic Authentication to restrict access to the endpoint.

In our case, when we deploy our application in production later we will run it behind [Caddy](https://caddyserver.com/) as a reverse proxy. As part of our Caddy set up, we’ll restrict access to the `GET /debug/vars` endpoint so that it can only be accessed via connections from the local machine, rather than being exposed on the internet.

### Request-level Metrics

We’ll start by recording the following three things:

- The total number of requests received.
- The total number of responses sent.
- The total (cumulative) time taken to process all requests in [microseconds](https://en.wikipedia.org/wiki/Microsecond).

Let’s create a new `metrics()` middleware:

```go
func (app *application) metrics(next http.Handler) http.Handler {
    var (
        totalRequestsReceived           = expvar.NewInt("total_requests_received")
        totalResponsesSent              = expvar.NewInt("total_responses_sent")
        totalProcessingTimeMicroseconds = expvar.NewInt("total_processing_time_μs")
    )
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()
        totalRequestsReceived.Add(1)
        next.ServeHTTP(w, r)
        totalResponsesSent.Add(1)
        duration := time.Since(start).Microseconds()
        totalProcessingTimeMicroseconds.Add(duration)
    })
}

func (app *application) routes() http.Handler {

    return app.metrics(app.recoverPanic(app.enableCORS(app.rateLimit(app.authenticate(router)))))
}
```

And then use `hey` to generate the load like so:

```bash
set BODY '{"email": "tifa@ff7.com", "password": "my waifu"}'
hey.exe -n 800 -d "$BODY" -m "POST" http://localhost:4000/v1/tokens/authentication
```

Based on the `total_processing_time_μs` and `total_responses_sent` values, we can calculate that the average processing time per request was approximately `200 ms` (remember, this endpoint is deliberately slow and computationally expensive since we need to verify the bcrypted user password).

#### ➤ Calculating additional metrics

Based on this information in the `GET /debug/vars` response, you can also derive some additional interesting metrics. Such as…

- The number of ‘active’ in-flight requests:

  ```html
  total_requests_received - total_responses_sent
  ```

- The average number of requests received per second (between calls A and B to the `GET /debug/vars` endpoint):

  ```html
  (total_requests_received_B - total_requests_received_A) / (timestamp_B - timestamp_A)
  ```

- The average processing time per request (between calls A and B to the `GET /debug/vars` endpoint):

- ```html
  (total_processing_time_μs_B - total_processing_time_μs_A) / (total_requests_received_B - total_requests_received_A)
  ```

### Recording HTTP Status Codes

#### ➤ 实现自定义的 http.ResponseWriter

The tricky part of doing this is finding out *what HTTP status code a response has* in our `metrics()` middleware. Unfortunately Go doesn’t support this out-of-the-box — there is no built-in way to examine a `http.ResponseWriter` to see what status code is going to be sent to a client. To capture the response status code, we’ll instead need to create our own custom `http.ResponseWriter` that records a copy of the HTTP status code for future access.

Before Go 1.20 was released, doing this was brittle and [awkward to get right](https://github.com/felixge/httpsnoop#why-this-package-exists). But now Go has the new `http.ResponseController` type with support for `http.ResponseWriter` *unwrapping*, it’s suddenly become much more straightforward to implement yourself. If you need to work with an older version of Go, I highly recommend using the the third-party [`httpsnoop`](https://github.com/felixge/httpsnoop) package rather than implementing your own custom `http.ResponseWriter`.

Effectively, what we want to do is create a struct that wraps an existing `http.ResponseWriter` and has custom `Write()` and `WriteHeader()` methods implemented on it which record the response status code. Importantly, we should also implement an `Unwrap()` method on it, which returns the original, wrapped, `http.ResponseWriter` value.

```go
// 添加额外字段记录 status code, 然后重写相关方法
type metricsResponseWriter struct {
    wrapped       http.ResponseWriter
    statusCode    int
    headerWritten bool
}

func newMetricsResponseWriter(w http.ResponseWriter) *metricsResponseWriter {
    return &metricsResponseWriter{
        wrapped:    w,
        statusCode: http.StatusOK,
    }
}

func (m *metricsResponseWriter) Header() http.Header {
    return m.wrapped.Header()
}

func (m *metricsResponseWriter) WriteHeader(statusCode int) {
    m.wrapped.WriteHeader(statusCode)

    // WriteHeader 方法仅在第一次调用生效, 要遵守这个特性:
    if !m.headerWritten {
        m.statusCode = statusCode
        m.headerWritten = true
    }
}

func (m *metricsResponseWriter) Write(b []byte) (int, error) {
    // 没有调用 WriteHeader 就调用 Write 会自动 WriteHeader(200), 所以字段要设为 true
    m.headerWritten = true
    return m.wrapped.Write(b)
}

func (m *metricsResponseWriter) Unwrap() http.ResponseWriter {
    return m.wrapped
}
```

(1) The important thing to realize is that our `metricsResponseWriter` type *satisfies* the [`http.ResponseWriter` interface](https://pkg.go.dev/net/http#ResponseWriter). It has `Header()`, `WriteHeader()` and `Write()` methods with the appropriate signature, so we can use it in our handlers just like normal.

(2) Also, notice that we don’t record the status code until *after* the ‘pass through’ call in the `WriteHeader()` method. This is because a panic in that operation (potentially due to an invalid status code) may mean that a different status code is ultimately sent to the client.

(3) Lastly, we also set a default status code of `200 OK` in the `newMetricsResponseWriter()` function. It’s important that we set this default here, in case a handler doesn’t ever call `Write()` or `WriteHeader()`.

#### ➤ 统计各个状态码对应的响应数

```go
func (app *application) metrics(next http.Handler) http.Handler {
    var (
        totalRequestsReceived           = expvar.NewInt("total_requests_received")
        totalResponsesSent              = expvar.NewInt("total_responses_sent")
        totalProcessingTimeMicroseconds = expvar.NewInt("total_processing_time_μs")
        totalResponsesSentByStatus      = expvar.NewMap("total_responses_sent_by_status") // 并发安全的计数器
    )
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()
        totalRequestsReceived.Add(1)
        mw := newMetricsResponseWriter(w) // 替换 ResponseWriter
        next.ServeHTTP(mw, r)
        totalResponsesSent.Add(1)
        totalResponsesSentByStatus.Add(strconv.Itoa(mw.statusCode), 1) // 现在知道状态码是啥了
        duration := time.Since(start).Microseconds()
        totalProcessingTimeMicroseconds.Add(duration)
    })
}
```

再次运行上面的 hey 测试，并且开启限流，可得:

```json
{
    "total_responses_sent_by_status": {
        "200": 4,
        "201": 4,
        "404": 1,
        "429": 796
    }
}
```

#### ➤ 题外话，类型嵌入和显式 wrap 哪个好?

```go
// 这是手动 wrap
// 调用方法时长这样: m.wrapped.WriteHeader(statusCode)
type metricsResponseWriter struct {
    wrapped       http.ResponseWriter
    statusCode    int
    headerWritten bool
}

// 这是类型嵌入
// 调用方法时长这样: m.ResponseWriter.WriteHeader(statusCode)
type metricsResponseWriter struct {
    http.ResponseWriter
    statusCode    int
    headerWritten bool
}
```

Embedded `http.ResponseWriter` will give you the same end result as the original approach. However, the gain is that you don’t need to write a `Header()` method for the `metricsResponseWriter` struct (it’s automatically promoted from the embedded `http.ResponseWriter`). Whereas the loss — at least in my eyes — is that it’s a bit less clear and explicit than using a `wrapped` field. Either approach is fine, it’s really just a matter of taste which one you prefer.
