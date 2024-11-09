## Table of Contents
  - [优先用 Caddy/Nginx](#%E4%BC%98%E5%85%88%E7%94%A8-CaddyNginx)
    - [注意事项](#%E6%B3%A8%E6%84%8F%E4%BA%8B%E9%A1%B9)
  - [静态文件服务器](#%E9%9D%99%E6%80%81%E6%96%87%E4%BB%B6%E6%9C%8D%E5%8A%A1%E5%99%A8)
    - [首个 File Server](#%E9%A6%96%E4%B8%AA-File-Server)
    - [返回单个文件](#%E8%BF%94%E5%9B%9E%E5%8D%95%E4%B8%AA%E6%96%87%E4%BB%B6)
    - [关闭文件目录](#%E5%85%B3%E9%97%AD%E6%96%87%E4%BB%B6%E7%9B%AE%E5%BD%95)
    - [关于 IO 性能](#%E5%85%B3%E4%BA%8E-IO-%E6%80%A7%E8%83%BD)
    - [一些注意事项](#%E4%B8%80%E4%BA%9B%E6%B3%A8%E6%84%8F%E4%BA%8B%E9%A1%B9)
  - [HTTPS 相关](#HTTPS-%E7%9B%B8%E5%85%B3)
    - [生成自签证书](#%E7%94%9F%E6%88%90%E8%87%AA%E7%AD%BE%E8%AF%81%E4%B9%A6)
    - [HTTPS 服务器](#HTTPS-%E6%9C%8D%E5%8A%A1%E5%99%A8)
    - [修改 TLS 配置](#%E4%BF%AE%E6%94%B9-TLS-%E9%85%8D%E7%BD%AE)
    - [设置连接超时](#%E8%AE%BE%E7%BD%AE%E8%BF%9E%E6%8E%A5%E8%B6%85%E6%97%B6)

## 优先用 Caddy/Nginx

### 注意事项

静态文件服务器和 HTTPS 都可以用 Nginx / Caddy 实现，不必在 Golang 中折腾

Caddy 支持 HTTPS 特别方便，连免费证书都会帮你申请和续签

## 静态文件服务器

### 首个 File Server

Now let’s improve the look and feel of the home page by adding some static CSS and image files to our project, along with a tiny bit of JavaScript to highlight the active navigation item.

```bash
curl https://www.alexedwards.net/static/sb-v2.tar.gz # 下载静态文件, 自行解压到 ui/static
```

Go’s net/http package ships with a built-in `http.FileServer` handler which you can use to serve files over HTTP from a specific directory. Let’s add a new route to our application so that all requests which begin with `"/static/"` are handled using this, like so:

```go
fileServer := http.FileServer(http.Dir("./ui/static/"))
```

When this handler receives a request, it will remove the leading slash from the URL path and then search the `./ui/static` directory for the corresponding file to send to the user. So, for this to work correctly, we must strip the leading `"/static"` from the URL path before passing it to `http.FileServer`. Otherwise it will be looking for a file which doesn’t exist and the user will receive a 404 page not found response. Fortunately Go includes a `http.StripPrefix()` helper specifically for this task.

```go
    // 创建 File Server
    fileServer := http.FileServer(http.Dir("./web/static/"))
    
    // 所有以 /static/ 开头的 URL 都交给 File Server 处理
    // 把请求交给 File Server 前, 去掉请求 URL 中的 /static 前缀
    router.Handle("/static/", http.StripPrefix("/static", fileServer))
```

更新 ui/html/base.tmpl 模板以使用静态文件

```html
    <head>
        <!-- Link to the CSS stylesheet and favicon -->
        <link rel='stylesheet' href='/static/css/main.css'>
        <link rel='shortcut icon' href='/static/img/favicon.ico' type='image/x-icon'>
        <!-- Also link to some fonts hosted by Google -->
        <link rel='stylesheet' href='https://fonts.googleapis.com/css?family=Ubuntu+Mono:400,700'>
    </head>
    <!-- And include the JavaScript file -->
    <script src="/static/js/main.js" type="text/javascript"></script>
```

### 返回单个文件

Sometimes you might want to serve a single file from within a handler. For this there’s the `http.ServeFile()` function, which you can use like so:

```go
// Warning: http.ServeFile() does not automatically sanitize the file path. 
// If you’re constructing a file path from untrusted user input, to avoid directory 
// traversal attacks you must sanitize the input with filepath.Clean() before using it.
func downloadHandler(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "./ui/static/file.zip") // 这里的文件路径并非用户提供的，所以无需 clean
}
```

### 关闭文件目录

If you want to disable directory listings there are a few different approaches you can take. The simplest way? Add a blank `index.html` file to the specific directory that you want to disable listings for. This will then be served instead of the directory listing, and the user will get a `200 OK` response with no body. If you want to do this for all directories under `./ui/static` you can use the command:

```bash
find ./ui/static -type d -exec touch {}/index.html \;
```

A more complicated (but arguably better) solution is to create a custom implementation of [http.FileSystem](https://pkg.go.dev/net/http/#FileSystem), and have it return an `os.ErrNotExist` error for any directories. A full explanation and sample code can be found in [this blog post](https://www.alexedwards.net/blog/disable-http-fileserver-directory-listings).

### 关于 IO 性能

In the code above we’ve set up our file server so that it serves files out of the `./ui/static` directory on your hard disk. But it’s important to note that, once the application is up-and-running, `http.FileServer` probably won’t be reading these files from disk. Both Windows and Unix-based operating systems cache recently-used files in RAM, so (for frequently-served files at least) it’s likely that `http.FileServer` will be serving them from RAM rather than making the relatively slow round-trip to your hard disk.

### 一些注意事项

Go’s file server has a few really nice features that are worth mentioning:

- It sanitizes all request paths by running them through the [path.Clean()](https://pkg.go.dev/path/#Clean) function before searching for a file. This removes any `.` and `..` elements from the URL path, which helps to stop directory traversal attacks. 

- [Range requests](https://benramsey.com/blog/2008/05/206-partial-content-and-range-requests) are fully supported. This is great if your application is serving large files and you want to support resumable downloads. You can see this functionality in action if you use curl to request bytes 100-199 of the `logo.png` file, like so:

  ```bash
  curl -i -H "Range: bytes=100-199" --output - http://localhost:4000/static/img/logo.png
  ```

  

- The `Last-Modified` and `If-Modified-Since` headers are transparently supported. If a file hasn’t changed since the user last requested it, then `http.FileServer` will send a `304 Not Modified` status code instead of the file itself. This helps reduce latency and processing overhead for both the client and server

- The `Content-Type` is automatically set from the file extension using the `mime.TypeByExtension()` function. You can add your own custom extensions and content types using the `mime.AddExtensionType()` function if necessary.



## HTTPS 相关

### 生成自签证书

HTTPS is essentially HTTP sent across a TLS (Transport Layer Security) connection. Because it’s sent over a TLS connection the data is encrypted and signed, which helps ensure its privacy and integrity during transit. Before our server can start using HTTPS, we need to generate a TLS certificate. For production servers I recommend using [Let’s Encrypt ( 免费好用的 HTTPS 证书 )](https://letsencrypt.org/) to create your TLS certificates, but for development purposes the simplest thing to do is to generate your own self-signed certificate.  

Handily, the `crypto/tls` package in Go’s standard library includes a `generate_cert.go` tool that we can use to easily create our own self-signed certificate. To run the `generate_cert.go` tool, you’ll need to know the place on your computer where the source code for the Go standard library is installed.

```bash
mkdir tls
cd tls
go run /usr/local/go/src/crypto/tls/generate_cert.go --rsa-bits=2048 --host=localhost
```

Behind the scenes the `generate_cert.go` tool works in two stages:

1. First it generates a 2048-bit RSA key pair, which is a cryptographically secure public key and private key.
2. It then stores the private key in a `key.pem` file, and generates a self-signed TLS certificate for the host localhost containing the public key — which it stores in a `cert.pem` file. Both the private key and certificate are PEM encoded, which is the standard format used by most TLS implementations.

By default, the `generate_cert.go` tool grants read permission to all users for the `cert.pem` file but read permission only to the owner of the `key.pem` file. If you’re using a version control system (like Git) you may want to add an ignore rule so the contents of the tls directory are not accidentally committed. For instance:

```bash
echo 'tls/' >> .gitignore
```

### HTTPS 服务器

Starting a HTTPS web server is lovely and simple — we just need open the `main.go` file and swap the `srv.ListenAndServe()` method for `srv.ListenAndServeTLS()` instead. If you open up your web browser and visit https://localhost:4000/ you will probably get a browser warning similar to the screenshot below.  

```go
func main() {
    // 让 session cookie 只在 HTTPS 中发送 ( 如果是不安全的 HTTP 则不发送 session cookie )
    sessionManager.Cookie.Secure = true
    err = server.ListenAndServeTLS("./web/tls/cert.pem", "./web/tls/key.pem")
}
```

- It’s important to note that our HTTPS server only supports HTTPS. If you try making a regular HTTP request to it, the server will send the user a 400 Bad Request status.

- A big plus of using HTTPS is that — if a client supports HTTP/2 connections — Go’s HTTPS server will automatically upgrade the connection to use HTTP/2. This is good because it means that, ultimately, our pages will load faster for users. If you’re not familiar with HTTP/2 you can watch this [GoSF meetup talk](https://www.youtube.com/watch?v=FARQMJndUn0).

### 修改 TLS 配置

Go has good default settings for its HTTPS server, but it’s possible to optimize and customize how the server behaves. Go supports a few elliptic curves, but as of Go 1.20 only `tls.CurveP256` and `tls.X25519` [have assembly implementations](https://github.com/golang/go/tree/8488309192b0ed4b393e2f7b2a93491139ff8ad0/src/crypto/internal/nistec). The others are very CPU intensive, so omitting them helps ensure that our server will remain performant under heavy loads.

For some applications, it may be desirable to limit your HTTPS server to only support some of these cipher suites. For example, you might want to only support cipher suites which use ECDHE (forward secrecy) and not support weak cipher suites. Important: Restricting the supported cipher suites to only include strong, modern, ciphers can mean that users with certain older browsers won’t be able to use your website. There’s a balance to be struck between security and backwards-compatibility and the right decision for you will depend on the technology typically used by your user base. [Mozilla’s recommended configurations](https://wiki.mozilla.org/Security/Server_Side_TLS) for modern, intermediate and old browsers may assist you in making a decision here.

```go
    tlsConfig := &tls.Config{
        CurvePreferences: []tls.CurveID{tls.X25519, tls.CurveP256}, // 这两个曲线有优化, 其他曲线较慢
        MinVersion:       tls.VersionTLS12,                         // 比版本 1.2 低的 TLS 有安全漏洞
        CipherSuites: []uint16{
            tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384, // 这些是推荐的 cipher suite 加密套件
            tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,   // 注意 TLS 1.3 会忽略 CipherSuites 设置
            tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305,  // 因为它的加密套件全都是安全的, 无需修改
            tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305,
            tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
            tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
        },
    }
    server := &http.Server{
        TLSConfig:    tlsConfig,
    }
```

### 设置连接超时

By default, Go enables `keep-alives` on all accepted connections. This helps reduce latency (especially for HTTPS connections) because a client can reuse the same connection for multiple requests without having to repeat the handshake. In our case, we’ve set `IdleTimeout` to 1 minute, which means that all `keep-alive` connections will be automatically closed after 1 minute of inactivity.  

In our code we’ve also set the `ReadTimeout` setting to 5 seconds. This means that if the request headers or body are still being read 5 seconds after the request is first accepted, then Go will close the underlying connection, then the user won’t receive any HTTP response. Setting a short `ReadTimeout` period helps to mitigate the risk from
slow-client attacks — such as Slowloris — which could otherwise keep a connection open indefinitely by sending partial, incomplete, HTTP requests.

The `WriteTimeout` setting will close the underlying connection if our server attempts to write to the connection after a given period (in our code, 10 seconds). For HTTPS connections, if some data is written to the connection more than 10 seconds after the request is first accepted, Go will close the underlying connection instead of writing the data. This means that if you’re using HTTPS (like we are) it’s sensible to set `WriteTimeout` to a value greater than `ReadTimeout`.  

> It’s important to bear in mind that writes made by a handler are buffered and written to the connection as one when the handler returns. Therefore, the idea of `WriteTimeout` is generally not to prevent long-running handlers, but to prevent the data that the handler returns from taking too long to write.  

```go
    server := &http.Server{
        ReadTimeout:  5 * time.Second,
        WriteTimeout: 10 * time.Second,
        IdleTimeout:  60 * time.Second,
    }
```