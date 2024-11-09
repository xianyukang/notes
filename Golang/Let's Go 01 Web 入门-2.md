## Table of Contents
  - [连接 MySQL](#%E8%BF%9E%E6%8E%A5-MySQL)
    - [设置 MySQL](#%E8%AE%BE%E7%BD%AE-MySQL)
    - [连接 MySQL](#%E8%BF%9E%E6%8E%A5-MySQL)
    - [连接 MySQL 示例](#%E8%BF%9E%E6%8E%A5-MySQL-%E7%A4%BA%E4%BE%8B)
    - [设置连接池](#%E8%AE%BE%E7%BD%AE%E8%BF%9E%E6%8E%A5%E6%B1%A0)
  - [用标准库 CRUD](#%E7%94%A8%E6%A0%87%E5%87%86%E5%BA%93-CRUD)
    - [创建 models 包](#%E5%88%9B%E5%BB%BA-models-%E5%8C%85)
    - [查询时的类型转换](#%E6%9F%A5%E8%AF%A2%E6%97%B6%E7%9A%84%E7%B1%BB%E5%9E%8B%E8%BD%AC%E6%8D%A2)
    - [CRUD - 插入记录](#CRUD--%E6%8F%92%E5%85%A5%E8%AE%B0%E5%BD%95)
    - [CRUD - 查询单行](#CRUD--%E6%9F%A5%E8%AF%A2%E5%8D%95%E8%A1%8C)
    - [CRUD - 查询多行](#CRUD--%E6%9F%A5%E8%AF%A2%E5%A4%9A%E8%A1%8C)
    - [CRUD - 使用事务](#CRUD--%E4%BD%BF%E7%94%A8%E4%BA%8B%E5%8A%A1)
  - [Middleware](#Middleware)
    - [中间件是什么](#%E4%B8%AD%E9%97%B4%E4%BB%B6%E6%98%AF%E4%BB%80%E4%B9%88)
    - [中间件概述](#%E4%B8%AD%E9%97%B4%E4%BB%B6%E6%A6%82%E8%BF%B0)
    - [更安全的响应头](#%E6%9B%B4%E5%AE%89%E5%85%A8%E7%9A%84%E5%93%8D%E5%BA%94%E5%A4%B4)
    - [中间件的控制流](#%E4%B8%AD%E9%97%B4%E4%BB%B6%E7%9A%84%E6%8E%A7%E5%88%B6%E6%B5%81)
    - [记录请求日志](#%E8%AE%B0%E5%BD%95%E8%AF%B7%E6%B1%82%E6%97%A5%E5%BF%97)
    - [恢复本协程的 Panic](#%E6%81%A2%E5%A4%8D%E6%9C%AC%E5%8D%8F%E7%A8%8B%E7%9A%84-Panic)
    - [注意其他协程的 Panic](#%E6%B3%A8%E6%84%8F%E5%85%B6%E4%BB%96%E5%8D%8F%E7%A8%8B%E7%9A%84-Panic)
    - [换种写法 Middleware Chain](#%E6%8D%A2%E7%A7%8D%E5%86%99%E6%B3%95-Middleware-Chain)
  - [处理 Session](#%E5%A4%84%E7%90%86-Session)
    - [会话有啥用](#%E4%BC%9A%E8%AF%9D%E6%9C%89%E5%95%A5%E7%94%A8)
    - [选择 Session 包](#%E9%80%89%E6%8B%A9-Session-%E5%8C%85)
    - [设置 Session Manager](#%E8%AE%BE%E7%BD%AE-Session-Manager)
    - [读写 Session Data](#%E8%AF%BB%E5%86%99-Session-Data)
    - [关于 Session 的原理](#%E5%85%B3%E4%BA%8E-Session-%E7%9A%84%E5%8E%9F%E7%90%86)
  - [User Authentication](#User-Authentication)
    - [数据库相关](#%E6%95%B0%E6%8D%AE%E5%BA%93%E7%9B%B8%E5%85%B3)
    - [展示用户注册页](#%E5%B1%95%E7%A4%BA%E7%94%A8%E6%88%B7%E6%B3%A8%E5%86%8C%E9%A1%B5)
    - [校验用户注册表单](#%E6%A0%A1%E9%AA%8C%E7%94%A8%E6%88%B7%E6%B3%A8%E5%86%8C%E8%A1%A8%E5%8D%95)
    - [加密密码后存储用户](#%E5%8A%A0%E5%AF%86%E5%AF%86%E7%A0%81%E5%90%8E%E5%AD%98%E5%82%A8%E7%94%A8%E6%88%B7)
    - [展示用户登录页](#%E5%B1%95%E7%A4%BA%E7%94%A8%E6%88%B7%E7%99%BB%E5%BD%95%E9%A1%B5)
    - [处理用户登录请求](#%E5%A4%84%E7%90%86%E7%94%A8%E6%88%B7%E7%99%BB%E5%BD%95%E8%AF%B7%E6%B1%82)
    - [处理用户登出请求](#%E5%A4%84%E7%90%86%E7%94%A8%E6%88%B7%E7%99%BB%E5%87%BA%E8%AF%B7%E6%B1%82)
  - [User Authorization](#User-Authorization)
    - [认证和验权的区别](#%E8%AE%A4%E8%AF%81%E5%92%8C%E9%AA%8C%E6%9D%83%E7%9A%84%E5%8C%BA%E5%88%AB)
    - [判断用户是否已登录](#%E5%88%A4%E6%96%AD%E7%94%A8%E6%88%B7%E6%98%AF%E5%90%A6%E5%B7%B2%E7%99%BB%E5%BD%95)
    - [用中间件设置访问权限校验](#%E7%94%A8%E4%B8%AD%E9%97%B4%E4%BB%B6%E8%AE%BE%E7%BD%AE%E8%AE%BF%E9%97%AE%E6%9D%83%E9%99%90%E6%A0%A1%E9%AA%8C)
    - [是时候加 CSRF 防护了](#%E6%98%AF%E6%97%B6%E5%80%99%E5%8A%A0-CSRF-%E9%98%B2%E6%8A%A4%E4%BA%86)
  - [请求上下文](#%E8%AF%B7%E6%B1%82%E4%B8%8A%E4%B8%8B%E6%96%87)
    - [为何需要 Request Context](#%E4%B8%BA%E4%BD%95%E9%9C%80%E8%A6%81-Request-Context)
    - [如何使用 Request Context](#%E5%A6%82%E4%BD%95%E4%BD%BF%E7%94%A8-Request-Context)
    - [如何避免 Key Collisions  ](#%E5%A6%82%E4%BD%95%E9%81%BF%E5%85%8D-Key-Collisions)
    - [中间件: 判断用户是否登录](#%E4%B8%AD%E9%97%B4%E4%BB%B6-%E5%88%A4%E6%96%AD%E7%94%A8%E6%88%B7%E6%98%AF%E5%90%A6%E7%99%BB%E5%BD%95)
    - [不要滥用 Request Context](#%E4%B8%8D%E8%A6%81%E6%BB%A5%E7%94%A8-Request-Context)

## 连接 MySQL

### 设置 MySQL

安装 MySQL 并创建数据库、表、MySQL用户:

```mysql
-- Create a new UTF-8 `snippetbox` database.
CREATE DATABASE snippetbox CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
-- Switch to using the `snippetbox` database.
USE snippetbox;
-- Create a `snippets` table.
CREATE TABLE snippets
(
    id      INTEGER      NOT NULL PRIMARY KEY AUTO_INCREMENT,
    title   VARCHAR(100) NOT NULL,
    content TEXT         NOT NULL,
    created DATETIME     NOT NULL,
    expires DATETIME     NOT NULL
);
-- Add an index on the created column.
CREATE INDEX idx_snippets_created ON snippets (created);

-- Add some dummy records (which we'll use in the next couple of chapters).
INSERT INTO snippets (title, content, created, expires)
VALUES ('This is 抛瓦', '抛瓦', UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL 365 DAY)),
       ('I need 抛瓦', 'need 抛瓦', UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL 365 DAY)),
       ('I need more 抛瓦', 'more 抛瓦', UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL 365 DAY));

-- Create a new web user with SELECT and INSERT privileges only on the database.
CREATE USER 'web'@'%'; # 我需要通过局域网访问 mysql 所以用通配符 % 而不是 localhost
GRANT SELECT, INSERT, UPDATE, DELETE ON snippetbox.* TO 'web'@'%';
ALTER USER 'web'@'%' IDENTIFIED BY 'pass';
```

### 连接 MySQL

```bash
go get -u github.com/go-sql-driver/mysql@v1 # 安装 MySQL 驱动
```

Now that the MySQL database is all set up and we’ve got a driver installed, the natural next step is to connect to the database from our web application. To do this we need Go’s `sql.Open()` function, which you use a bit like this:

```go
// 这叫匿名引入，能引入包的副作用，package 中的 init 函数会注册驱动
import _ "github.com/go-sql-driver/mysql"
// sql.Open() initializes a new sql.DB object, which is essentially a pool of database connections.
db, err := sql.Open("mysql", "web:pass@tcp(localhost)/snippetbox?parseTime=true")
```

There are a few things about this code to explain and emphasize:

- The first parameter to `sql.Open()` is the driver name and the second parameter is the data source name (sometimes also called a connection string or DSN) which describes how to connect to your database.
- The format of the data source name will depend on which database and driver you’re using.
- The `parseTime=true` part of the DSN above is a driver-specific parameter which instructs our driver to convert SQL `TIME` and `DATE` fields to Go `time.Time` objects.
- The `sql.Open()` function returns a `sql.DB` object. This isn’t a database connection — it’s a pool of many connections. This is an important difference to understand. Go manages the connections in this pool as needed, automatically opening and closing connections to the database via the driver.
- The connection pool is safe for concurrent access, so you can use it from web application handlers safely.
- The connection pool is intended to be long-lived. In a web application it’s normal to initialize the connection pool in your `main()` function and then pass the pool to your handlers. You shouldn’t call `sql.Open()` in a short-lived handler itself — it would be a waste of memory and network resources.

### 连接 MySQL 示例

Open up your main.go file and add the following code:

```go
import _ "github.com/go-sql-driver/mysql"
func main() {
    dsn := flag.String("dsn", "web:pass@tcp(localhost)/snippetbox?parseTime=true", "MySQL data source name")
    flag.Parse()

    // 初始化数据库连接池
    db, err := openDB(*dsn)
    if err != nil {
        errorLog.Fatal(err)
    }
    defer db.Close()
}
func openDB(dsn string) (*sql.DB, error) {
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        return nil, err
    }
    // Ping 一下测试是否连上了
    if err := db.Ping(); err != nil {
        return nil, err
    }
    return db, nil
}
```

There’re a few things about this code which are interesting:

- The `sql.Open()` function doesn’t actually create any connections, all it does is initialize the pool for future use. Actual connections to the database are established lazily, as and when needed for the first time. So to verify that everything is set up correctly we need to use the `db.Ping()` method to create a connection and check for any errors.
- At this moment in time, the call to `defer db.Close()` is a bit superfluous. Our application is only ever terminated by a signal interrupt (i.e. `Ctrl+c`) or by `errorLog.Fatal()`. In both of those cases, the program exits immediately and deferred functions are never run. But including `db.Close()` is a good habit to get into and it could be beneficial later in the future if you add a graceful shutdown to your application.

### 设置连接池

<font color='#D05'>By default, there is no limit on the maximum number of open connections</font> (idle + in-use) at one time, but the default maximum number of idle connections in the pool is 2. You can change these defaults with the `db.SetMaxOpenConns(100)` and `db.SetMaxIdleConns(5)` methods.

If the maximum number of open connections is reached, Go will wait until one of the connections is freed and becomes idle. From a user perspective, this means their HTTP request will hang until a connection is freed.

*Your database itself probably has a hard limit on the maximum number of connections*. For example, the default limit for MySQL is 151. So leaving `SetMaxOpenConns()` totally unlimited or setting it to greater than 151 may result in your database returning  a "too many connections" error under high load.

For some applications waiting for a idle connection might be fine, but in a web application *it’s arguably better* to immediately log the `"too many connections"` error message and send a 500 Internal Server Error to the user, rather than having their HTTP request hang and potentially timeout while waiting for a free connection. That’s why I haven’t used the `SetMaxOpenConns()` and `SetMaxIdleConns()` methods in our application, and left the behavior of `sql.DB` as the default settings.

## 用标准库 CRUD

The `database/sql` package essentially provides a standard interface between your Go application and the world of SQL databases. So long as you use the `database/sql` package, the Go code you write will generally be portable and will work with any kind of SQL database — whether it’s MySQL, PostgreSQL, SQLite or something else ( driver-specific quirks and SQL implementations aside ).

If the verbosity really is starting to grate on you, you might want to consider trying the `jmoiron/sqlx` package. It’s well designed and provides some good extensions that make working with SQL queries quicker and easier. 

### 创建 models 包

Benefits of this structure:

- There’s a clean separation of concerns. Our database logic isn’t tied to our handlers which means that handler responsibilities are limited to HTTP stuff (i.e. validating requests and writing responses). This will make it easier to write tight, focused, unit tests in the future.
- Because the model actions are defined as methods on an object — in our case `SnippetModel` — there’s the opportunity to create an interface and mock it for unit testing purposes.

Let’s open the `internal/models/snippets.go` file and add a new `Snippet` struct to represent the data for an individual snippet, along with a `SnippetModel` type with methods on it to access and manipulate the snippets in our database. Like so:

```go
// 定义一个结构体表示 Snippet
type Snippet struct {
    ID      int
    Title   string
    Content string
    Created time.Time
    Expires time.Time
}

// 定义一个 SnippetModel 封装 CRUD
type SnippetModel struct {
    DB *sql.DB
}

func (m *SnippetModel) Insert(title string, content string, expires int) (int, error) {
    return 0, nil
}
func (m *SnippetModel) Get(id int) (*Snippet, error) {
    return nil, nil
}
func (m *SnippetModel) Latest() ([]*Snippet, error) {
    return nil, nil
}
```

We need to establish a new `SnippetModel` struct in our `main()` function and then inject it as a dependency via the `application` struct — just like we have with our other dependencies.

```go
type application struct {
    errorLog *log.Logger
    infoLog  *log.Logger
    snippets *models.SnippetModel // 让 Handler 可以通过 snippets 进行 CRUD
}
func main() {
    app := &application{
        errorLog: errorLog,
        infoLog:  infoLog,
        snippets: &models.SnippetModel{DB: db},
    }
}
```

### 查询时的类型转换

Behind the scenes of `rows.Scan()` your driver will automatically convert the raw output from the SQL database to the required native Go types. Usually:

- CHAR, VARCHAR and TEXT map to `string`.
- BOOLEAN maps to `bool`.
- INT maps to `int`; BIGINT maps to `int64`.
- DECIMAL and NUMERIC map to `float64`.
- TIME, DATE and TIMESTAMP map to `time.Time`.

可以用 `*string` 或 `sql.NullString` 解决可能为 `NULL` 的字段, 但最好还是设计表时不允许 `NULL` 字段

Go doesn’t do very well is managing `NULL` values in database records. Let’s pretend that the title column in our snippets table contains a `NULL` value in a particular row. If we queried that row, then `rows.Scan()` would return an error because it can’t convert `NULL` into a `string`.

A quirk of our MySQL driver is that we need to use the `parseTime=true` parameter in our DSN to force it to convert TIME and DATE fields to `time.Time`. Otherwise it returns these as `[]byte` objects.

### CRUD - 插入记录

Go provides three different methods for executing database queries:

- `DB.Query()` is used for `SELECT` queries which return multiple rows.

- `DB.QueryRow()` is used for `SELECT` queries which return a single row.

- `DB.Exec()` is used for statements which don’t return rows (like `INSERT` and `DELETE`)

Now let’s update the `SnippetModel.Insert()` method so that it creates a new record in our `snippets` table and then returns the integer `id` for the new record. Because the data we’ll be using will ultimately be untrusted user input from a form, it’s good practice to use placeholder parameters (`?`) instead of interpolating data in the SQL query.

```go
func (m *SnippetModel) Insert(title string, content string, expires int) (int, error) {
    stmt := `INSERT INTO snippets (title, content, created, expires) 
VALUES(?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))`

    // 按照占位符的顺序依次传入 title, content, expires
    // 返回值是 sql.Result 类型, 它包含 SQL 的执行结果
    result, err := m.DB.Exec(stmt, title, content, expires)
    if err != nil {
        return 0, err
    }

    // 这样获取刚刚插入的记录对应的主键 ID
    id, err := result.LastInsertId()
    if err != nil {
        return 0, err
    }

    // 返回的 ID 是 int64 类型所以要转一下
    return int(id), nil
}
```

Let’s quickly discuss the `sql.Result` type returned by `DB.Exec()`. This provides two methods:

- `LastInsertId()` — which returns the integer (an `int64`) generated by the database in response to a command.
- `RowsAffected()` — which returns the number of rows (as an `int64`) affected by the statement.
- 如果用不到可以忽略返回的 sql.Result: `_, err := m.DB.Exec("INSERT INTO ...", ...)`
- 另外 PostgreSQL [不支持](https://github.com/lib/pq/issues/24) `LastInsertId()` 方法

Using the model in our handlers:

```go
func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        w.Header().Set("Allow", http.MethodPost)
        app.clientError(w, http.StatusMethodNotAllowed)
        return
    }
    // 调用 models 插入记录
    id, err := app.snippets.Insert("爱你的猫抛瓦", "抛瓦", 7)
    if err != nil {
        app.serverError(w, err)
        return
    }
    // Redirect the user to the relevant page for the snippet.
    // https://stackoverflow.com/questions/4764297/difference-between-http-redirect-codes
    http.Redirect(w, r, fmt.Sprintf("/snippet/view?id=%d", id), http.StatusSeeOther)
}
```

可用 curl 做测试:

```bash
# -i 表示显示 Response Header，-L 表示跟随重定向，-X 指定请求方法为 POST
curl -iL -X POST http://localhost:4000/snippet/create
```

### CRUD - 查询单行

```go
var ErrNoRecord = errors.New("models: no matching record found")
func (m *SnippetModel) Get(id int) (*Snippet, error) {
    stmt := `SELECT id, title, content, created, expires FROM snippets WHERE expires > UTC_TIMESTAMP() AND id = ?`
    // 按占位符顺序依次传入参数, 返回值是 *sql.Row 类型
    row := m.DB.QueryRow(stmt, id)
    // 初始化一个结构体用于存查到的数据
    s := &Snippet{}
    // 用 Scan 方法把查询结果序列化到结构体, 这里的参数顺序和参数个数必须和 SELECT 一一对应
    err := row.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
    // 如果没有查到任何记录, row.Scan 会返回 sql.ErrNoRows 错误
    if err != nil {
        if errors.Is(err, sql.ErrNoRows) {
            return nil, ErrNoRecord // 返回自定义的错误类型, 方便和具体数据库解耦
        } else {
            return nil, err
        }
    }
    return s, nil
}
```

As an aside, you might be wondering why we’re returning the `ErrNoRecord` error from our `SnippetModel.Get()` method, instead of `sql.ErrNoRows` directly. The reason is to help encapsulate the model completely, so that our application isn’t concerned with the underlying datastore or reliant on datastore-specific errors for its behavior. Alright, let’s put the `SnippetModel.Get()` method into action:

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
    _, _ = fmt.Fprintf(w, "%+v", snippet)
}
```

### CRUD - 查询多行

```go
func (m *SnippetModel) Latest() ([]*Snippet, error) {
    stmt := `SELECT id, title, content, created, expires 
FROM snippets WHERE expires > UTC_TIMESTAMP() ORDER BY id DESC LIMIT 10`
    // 使用 Query 方法查询多行数据, 返回值是 *sql.Rows 类型
    rows, err := m.DB.Query(stmt)
    if err != nil {
        return nil, err
    }
    // 在检查错误后使用 defer 函数释放资源 ( 忘了 Close 会导致一直占用, 连接无法回到连接池 )
    defer rows.Close()

    // 创建空切片用来存储结果, 注意序列化成 json 时, 空切片变成 "[]" 而 nil 切片变成 "null"
    snippets := make([]*Snippet, 0)

    // 遍历查询结果集, 用 Scan 把每一行结果序列化成结构体
    for rows.Next() {
        s := &Snippet{}
        err = rows.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
        if err != nil {
            return nil, err
        }
        snippets = append(snippets, s)
    }
    // 退出上面的循环, 并不意味着 result set 被完整遍历, 有可能因错提前退出, 别忘了检查错误
    if err := rows.Err(); err != nil {
        return nil, err
    }

    return snippets, nil
}
```

### CRUD - 使用事务

It’s important to realize that calls to `Exec()`, `Query()` and `QueryRow()` can use any connection from the `sql.DB` pool. Even if you have two calls to `Exec()` immediately next to each other in your code, there is no guarantee that they will use the same database connection. Sometimes this isn’t acceptable. For instance, if you lock a table with MySQL’s `LOCK TABLES` command you must call `UNLOCK TABLES` on exactly the same connection to avoid a deadlock.

To guarantee that the same connection is used you can wrap multiple statements in a transaction. Here’s the basic pattern:

```go
func (m *SnippetModel) ExampleTransaction() error {
    // 调用 Begin 创建新的 sql.Tx 对象, 它代表一个数据库事务
    tx, err := m.DB.Begin()
    if err != nil {
        return err
    }
    // 通过 defer 让 Rollback 无论如何都会执行, 对已成功提交的事务 Rollback 不会有任何效果
    defer tx.Rollback()

    // 使用事务执行 CRUD
    _, err = tx.Exec("UPDATE snippets SET content = 'example1' WHERE id = 1;")
    if err != nil {
        return err
    }
    _, err = tx.Exec("UPDATE snippets SET content = 'example2' WHERE id = 2;")
    if err != nil {
        return err
    }

    // 最后 Commit 事务, 就算 Commit 失败也会被上面的 defer tx.Rollback() 回滚
    err = tx.Commit()
    return err
}
// 不论任何代码路径，都别忘了最后调用 Commit 或 Rollback
// 否则 connection 无法回到连接池，造成资源泄露
```

## Middleware

### 中间件是什么

- 中间件就是在 Handler 前后做点什么，比较通用/业务无关的需求适合用中间件实现，例如日志、限流、压缩、...
- [Middleware 如何工作,  以及用 20 行代码写一个中间件框架](https://chai2010.cn/advanced-go-programming-book/ch5-web/ch5-03-middleware.html#:~:text=%E4%BB%80%E4%B9%88%E5%8F%AF%E8%AF%BB%E6%80%A7%E3%80%82-,5.3.3%20%E6%9B%B4%E4%BC%98%E9%9B%85%E7%9A%84%E4%B8%AD%E9%97%B4%E4%BB%B6%E5%86%99%E6%B3%95,-%E4%B8%8A%E4%B8%80%E8%8A%82%E4%B8%AD)
- 每一个 Web 框架都会有对应的编写 Middleware 的方法, [这里有社区贡献的、写好的、适用于 gin 的中间件](https://github.com/gin-gonic/contrib)

The `http.StripPrefix()` function is a middleware, which removes a specific prefix from the request’s URL path before passing the request on to the file server.

### 中间件概述

When you’re building a web application there’s probably some shared functionality that you want to use for many (or even all) HTTP requests. For example, you might want to log every request, compress every response, or check a cache before passing the request to your handlers. Middleware is essentially some self-contained code which
independently acts on a request before or after your normal application handlers. 

You can think of a Go web application as a chain of `ServeHTTP()` methods being called one after another. The basic idea of middleware is to insert another handler into this chain. The middleware handler executes some logic, like logging a request, and then calls the `ServeHTTP()` method of the next handler in the chain.

The standard pattern for creating your own middleware looks like this:  

```go
// myMiddleware() is a function that accepts the next handler in a chain as a parameter. 
// It returns a handler which executes some logic and then calls the next handler.
func myMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Any code here will execute on the way down the chain.
        next.ServeHTTP(w, r)
        // Any code here will execute on the way back up the chain.
    })
}
```

It’s important to explain that where you position the middleware in the chain of handlers will affect the behavior of your application. If you position your middleware before the servemux in the chain then it will act on every request that your application receives. A good example of where this would be useful is middleware to log requests — as that’s typically something you would want to do for all requests.  

```bash
myMiddleware → servemux → application handler
```

Alternatively, you can position the middleware after the servemux in the chain — by wrapping a specific application handler. This would cause your middleware to only be executed for a specific route. An example of this would be something like authorization middleware, which you may only want to run on specific routes.

```bash
servemux → myMiddleware → application handler
```

### 更安全的响应头

Let's make our own middleware which automatically adds the following HTTP security headers to every response  

```http
Content-Security-Policy: default-src 'self'; style-src 'self' fonts.googleapis.com; font-src fonts.gstatic.com
Referrer-Policy: origin-when-cross-origin
X-Content-Type-Options: nosniff
X-Frame-Options: deny
X-XSS-Protection: 0
```

- `Content-Security-Policy` (often abbreviated to CSP) headers are used to restrict where the resources for your web page (e.g. JavaScript, images, fonts etc) can be loaded from. Setting a strict CSP policy helps prevent a variety of cross-site scripting, clickjacking, and other code-injection attacks.  
- CSP headers and how they work is a big topic, and I recommend reading this [primer](https://developer.mozilla.org/en-US/docs/Web/HTTP/CSP) if you haven’t come across them before. But, in our case, the header tells the browser that it’s OK to load fonts from `fonts.gstatic.com`, stylesheets from `fonts.googleapis.com` and `self` (our own origin), and then everything else only from `self`. Inline JavaScript is blocked by default.  
- [Referrer-Policy](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Referrer-Policy) is used to control what information is included in a `Referer` header when a user navigates away from your web page. In our case, we’ll set the value to `origin-when-cross-origin`, which means that the full URL will be included for same-origin requests, but for all other requests information like the URL path and any query string values will be stripped out.
- `X-Content-Type-Options: nosniff` instructs browsers to not MIME-type sniff the content-type of the response, which in turn helps to prevent content-sniffing attacks. ( 深究下去也不是我能掌握的东西，那就照抄吧 )
- `X-Frame-Options: deny` is used to help prevent clickjacking attacks in older browsers that don’t support CSP headers. Sites can use this to avoid [click-jacking](https://developer.mozilla.org/en-US/docs/Web/Security/Types_of_attacks#click-jacking) attacks, by ensuring that their content is not embedded into other sites. The [`Content-Security-Policy`](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Content-Security-Policy) HTTP header has a [`frame-ancestors`](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Content-Security-Policy/frame-ancestors) directive which [obsoletes](https://w3c.github.io/webappsec-csp/#frame-ancestors-and-frame-options) this header for supporting browsers.
- `X-XSS-Protection: 0` is used to disable the blocking of cross-site scripting attacks. Previously it was good practice to set this header to `X-XSS-Protection: 1; mode=block`, but when you’re using CSP headers like we are [the recommendation](https://owasp.org/www-project-secure-headers/#x-xss-protection) is to disable this feature altogether.

OK, let’s get back to our Go code and begin by creating a new `middleware.go` file.

```go
func secureHeaders(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Security-Policy", "default-src 'self'; style-src 'self' fonts.googleapis.com; font-src fonts.gstatic.com")
        w.Header().Set("Referrer-Policy", "origin-when-cross-origin")
        w.Header().Set("X-Content-Type-Options", "nosniff")
        w.Header().Set("X-Frame-Options", "deny")
        w.Header().Set("X-XSS-Protection", "0")
        next.ServeHTTP(w, r)
    })
}
```

Because we want this middleware to act on every request that is received, we need it to be executed before a request hits our servemux. We want the flow of control through our application to look like:  

```bash
secureHeaders → servemux → application handler
```

To do this we’ll need the `secureHeaders` middleware function to wrap our servemux. Let’s update the `routes.go` file to do exactly that:

```go
func routes(app *application) http.Handler {
    router := http.NewServeMux()
    router.HandleFunc("/", app.homePage)
    return secureHeaders(router) // 用中间件 wrap 一下 router, 就能让中间件对所有请求都生效
}
```

### 中间件的控制流

It’s important to know that when the last handler in the chain returns, control is passed back up the chain in the reverse direction. So when our code is being executed the flow of control actually looks like this:  

````go
// middleware → servemux → application handler → servemux → middleware
func myMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Any code here will execute on the way down the chain.
        next.ServeHTTP(w, r)
        // Any code here will execute on the way back up the chain.
    })
}
````

Another thing to mention is that if you call `return` in your middleware function before you call `next.ServeHTTP()`, then the chain will stop being executed and control will flow back upstream. As an example, a common use-case for early returns is authentication middleware which only allows execution of the chain to continue if a particular check is passed. For instance:  

```go
func myMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // If the user isn't authorized send a 403 Forbidden status and return to stop executing the chain.
        if !isAuthorized(r) {
            w.WriteHeader(http.StatusForbidden)
            return
        }
        // Otherwise, call the next handler in the chain.
        next.ServeHTTP(w, r)
    })
}
```

### 记录请求日志

Let’s continue in the same vein and add some middleware to log HTTP requests. Specifically, we’re going to use the information logger that we created earlier to record the IP address of the user, and which URL and method are being requested.

```go
// Middleware 就是 ServeHTTP 的 wrapper 函数, 这里还用到了闭包, 所以函数能访问 app.infoLog
func (app *application) logRequest(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        app.infoLog.Printf("%s - %s %s %s", r.RemoteAddr, r.Proto, r.Method, r.URL.RequestURI())
        next.ServeHTTP(w, r)
    })
}
```

Now let’s update our `routes.go` file so that the `logRequest` middleware is executed first, and for all requests, so that the flow of control looks like this: 

```go
// logRequest → secureHeaders → servemux → application handler
func routes(app *application) http.Handler {
    return app.logRequest(secureHeaders(router)) // 越外层的越先执行
}
```

### 恢复本协程的 Panic

In a simple Go application, when your code panics it will result in the application being terminated straight away. But our web application is a bit more sophisticated. Go’s HTTP server assumes that the effect of any panic is isolated to the goroutine serving the active HTTP request (remember, every request is handled in it’s own goroutine).  

Specifically, following a panic our server will log a stack trace to the server error log, unwind the stack for the affected goroutine (calling any deferred functions along the way) and close the underlying HTTP connection. But it won’t terminate the application, so importantly, any panic in your handlers won’t bring down your server.

```go
func (app *application) homePage(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/" {
        app.notFound(w)
        return
    }
    panic("oops")
}
```

But if a panic does happen in one of our handlers, what will the user see? Unfortunately, all we get is an empty response due to Go closing the underlying HTTP connection following the panic. This isn’t a great experience for the user. It would be more appropriate and meaningful to send them a proper HTTP response with a `500 Internal Server Error` status instead. Open up your `middleware.go` file and add the following code:

```go
func (app *application) recoverPanic(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // 安排 defer 函数, 其中检查是否出现 panic
        defer func() {
            if err := recover(); err != nil {
                w.Header().Set("Connection", "close")     // 让 HTTP Server 关掉连接 ( 在发完响应之后 )
                app.serverError(w, fmt.Errorf("%s", err)) // 不管 err 是什么, 利用它的字符串创建 error
            }
        }()
        next.ServeHTTP(w, r)
    })
}
```

- Setting the `Connection: close` header on the response acts as a trigger to make Go’s HTTP server automatically close the current connection after a response has been sent. It also informs the user that the connection will be closed. Note: If the protocol being used is HTTP/2, Go will automatically strip the `Connection: close` header from the response (so it is not malformed) and send a `GOAWAY` frame.
- The value returned by the builtin `recover()` function has the type `any`, and its underlying type could be `string`, `error`, or something else — whatever the parameter passed to `panic()` was. In the code above, we normalize this into an `error` by using the `fmt.Errorf()` function to create a new `error` object containing the default textual representation of the any value.

Let’s now put this to use in the `routes.go` file, so that it is the first thing in our chain to be executed (so that it covers panics in all subsequent middleware and handlers). 

```go
func routes(app *application) http.Handler {
    return app.recoverPanic(app.logRequest(secureHeaders(router))) // 越外层的越先执行
}
```

### 注意其他协程的 Panic

It’s important to realise that our middleware will only recover panics that happen in the same goroutine that executed the `recoverPanic()` middleware. If, for example, you have a handler which spins up another goroutine
(e.g. to do some background processing), then any panics that happen in the second goroutine will not be recovered — not by the `recoverPanic()` middleware… and not by the panic recovery built into Go HTTP server. **They will cause your application to exit and bring down the server**.

So, if you are spinning up additional goroutines from within your web application and there is any chance of a panic, you must make sure that you recover any panics from within those too. For example:  

```go
func myHandler(w http.ResponseWriter, r *http.Request) {
    // Spin up a new goroutine to do some background processing.
    go func() {
        defer func() {
            if err := recover(); err != nil {
                log.Print(fmt.Errorf("%s\n%s", err, debug.Stack()))
            }
        }()
        doSomeBackgroundProcessing()
    }()
    w.Write([]byte("OK"))
}
```

### 换种写法 Middleware Chain

You don’t need to use [github.com/justinas/alice](https://github.com/justinas/alice), but the reason I recommend it is because it makes it easy to create composable, reusable, middleware chains — and that can be a real help as your application grows and your routes become more complex. It allows you to rewrite a handler chain like this:

```go
return myMiddleware1(myMiddleware2(myMiddleware3(myHandler)))                   // 旧写法
return alice.New(myMiddleware1, myMiddleware2, myMiddleware3).Then(myHandler)   // 新写法
```

But the real power lies in the fact that you can use it to create middleware chains that can be assigned to variables, appended to, and reused. For example:

```go
myChain := alice.New(myMiddlewareOne, myMiddlewareTwo) // 使用旧写法就不好创建可重用的 Chain
myOtherChain := myChain.Append(myMiddleware3)          // 以及往 Chain 中继续 Append 中间件
return myOtherChain.Then(myHandler)
```

Install the package using: `go get github.com/justinas/alice@v1`. OK, let’s update our `routes.go` file as follows:

```go
func routes(app *application) http.Handler {
    // 创建 standard middleware chain, 并把它用于每一个请求
    standard := alice.New(app.recoverPanic, app.logRequest, secureHeaders)
    return standard.Then(router)
}
```





## 处理 Session

### 会话有啥用

A nice touch to improve our user experience would be to display a one-time confirmation message which the user sees after they’ve added a new snippet. A confirmation message like this should only show up for the user once (immediately after creating the snippet) and no other users should ever see the message. If you’ve been programming for a while already, you might know this type of functionality as a flash message or a toast.

To make this work, we need to start sharing data (or state) between HTTP requests for the same user. The most common way to do that is to implement a session for the user.  

### 选择 Session 包

- Until recently [gorilla/sessions](https://github.com/gorilla/sessions) was a very popular option for this. Unfortunately it’s no longer maintained and is in archive mode on GitHub, meaning that any bugs in the package — including security bugs — will not be fixed. ( 似乎又复活了 )

- So for most applications I now recommend using [alexedwards/scs](https://github.com/alexedwards/scs). It has a fairly simple API, supports automatic loading and saving of session data via middleware, and lets you store session data serverside in variety of databases (including MySQL, PostgreSQL and Redis).  
- For this project we’ve already got a MySQL database set up, so we’ll opt to use `alexedwards/scs` and store the session data in MySQL.

```bash
go get github.com/alexedwards/scs/v2@v2
go get github.com/alexedwards/scs/mysqlstore@latest
```

### 设置 Session Manager

The first thing we need to do is create a sessions table in our MySQL database to hold the session data for our users. Execute the following SQL statement to setup the `sessions` table:  

```mysql
CREATE TABLE sessions
(
    token  CHAR(43) PRIMARY KEY,
    data   BLOB         NOT NULL,
    expiry TIMESTAMP(6) NOT NULL
);
CREATE INDEX sessions_expiry_idx ON sessions (expiry);
```

- The `token` field will contain a unique, randomly-generated, identifier for each session.
- The `data` field will contain the actual session data that you want to share between HTTP requests. This is stored as binary data in a `BLOB` (binary large object) type.

- The `expiry` field will contain an expiry time for the session. The `scs` package will automatically delete expired sessions from the `sessions` table so that it doesn’t grow too large.  

The next thing we need to do is establish a session manager in our `main.go` file and make it available to our handlers via the `application` struct.

```go
type application struct {
    sessionManager *scs.SessionManager
}
func main() {
    sessionManager := scs.New()
    sessionManager.Store = mysqlstore.New(db)
    sessionManager.Lifetime = 12 * time.Hour

    app := &application{
        sessionManager: sessionManager,
    }
}
```

For the sessions to work, we also need to wrap our application routes with the middleware provided by the `SessionManager.LoadAndSave()` method. This middleware automatically loads and saves session data with every HTTP request and response. It’s important to note that we don’t need this middleware to act on all our application routes. Specifically, we don’t need it on the `/static/*filepath` route, because all this does is serve static files and there is no need for any stateful behavior.

```go
func routes(app *application) http.Handler {
    router := httprouter.New()
    // 这是为动态内容创建的 middleware chain, 而静态内容不需要使用 session
    dynamic := alice.New(app.sessionManager.LoadAndSave)
    router.Handler(http.MethodGet, "/", dynamic.ThenFunc(app.homePage))                         
    router.Handler(http.MethodGet, "/snippet/view/:id", dynamic.ThenFunc(app.snippetViewPage))  
    router.Handler(http.MethodGet, "/snippet/create", dynamic.ThenFunc(app.snippetCreate))      
    router.Handler(http.MethodPost, "/snippet/create", dynamic.ThenFunc(app.snippetCreatePost)) 
}
```

### 读写 Session Data

```go
func (app *application) snippetCreatePost(w http.ResponseWriter, r *http.Request) {
    // 使用 flash 作为 key 往 session 中添加数据
    app.sessionManager.Put(r.Context(), "flash", "Snippet successfully created!")
}

func (app *application) snippetViewPage(w http.ResponseWriter, r *http.Request) {
    // 因为 flash 消息只应该展示一次, 所以不用 GetString() 方法
    // 而是用 PopString() 方法读取并移除 flash 消息, 不存在时 GetString/PopString 返回空串
    flash := app.sessionManager.PopString(r.Context(), "flash")

    data := app.newTemplateData(r)
    data.Snippet = snippet
    data.Flash = flash
    app.render(w, http.StatusOK, "view.tmpl", data)
}
```

And now, we can update our `base.tmpl` file to display the flash message, if one exists.

```html
    <main>
        <!-- 如果 .Flash 不为空则展示 -->
        {{with .Flash}}
            <div class="flash">{{.}}</div>
        {{end}}
        {{template "main" .}}
    </main>
```

A little improvement we can make (which will save us some work later in the project) is to automate the display of flash messages, so that any message is automatically included the next time any page is rendered.

```go
func (app *application) newTemplateData(r *http.Request) *templateData {
    return &templateData{
        CurrentYear: time.Now().Year(),
        Flash:       app.sessionManager.PopString(r.Context(), "flash"),
    }
}
```

### 关于 Session 的原理

1. 服务器往 Cookie 中写入一个 Session ID
2. 后续浏览器请求都会带上这个 Session ID，这就实现了「 记住用户 / 创建会话 」
3. Session 中间件读取请求中的 Session ID，然后去 Session Store ( MySQL / Redis / ... ) 加载会话数据
4. 每次更新会话数据，Session 中间件都会把修改结果，保存到 Session Store

So, what happens in our application is that the `LoadAndSave()` middleware checks each incoming request for a session cookie. If a session cookie is present, it reads the session token and retrieves the corresponding session data from the database (while also checking that the session hasn’t expired). It then adds the session data to the request context so it can be used in your handlers.

Any changes that you make to the session data in your handlers are updated in the request context, and then the `LoadAndSave()` middleware updates the database with any changes to the session data before it returns.  



## User Authentication

### 数据库相关

建立 users 表:

```mysql
CREATE TABLE users
(
    id              INTEGER      NOT NULL PRIMARY KEY AUTO_INCREMENT,
    name            VARCHAR(255) NOT NULL,
    email           VARCHAR(255) NOT NULL,
    hashed_password CHAR(60)     NOT NULL,
    created         DATETIME     NOT NULL
);
ALTER TABLE users ADD CONSTRAINT users_uc_email UNIQUE (email);
```

定义错误接口:

```go
var (
    // 使用自定义错误而不是底层数据库返回的错误, 方便和数据库解耦, 切换数据库实现
    // 例如 MySQL 中查不到记录时会得到 sql.ErrNoRows, 而 MongoDB 会得到 mongo.ErrNoDocuments
    // 总之让 MysqlStore 和 MongoStore 遵守同样的接口, 在查不到记录时统一返回 ErrNoRecord
    ErrNoRecord       = errors.New("models: no matching record found")
    ErrDuplicateEmail = errors.New("models: duplicate email")

    // We'll use this later if a user tries to login with an incorrect email address or password.
    ErrInvalidCredentials = errors.New("models: invalid credentials")
)
```

定义结构体和 CURD 方法:

```go
type User struct {
    ID             int
    Name           string
    Email          string
    HashedPassword []byte
    Created        time.Time
}

type UserModel struct {
    DB *sql.DB
}

// TODO: 添加新用户
func (m *UserModel) Insert(name, email, password string) error {
    return nil
}
// TODO: 如果用户名和密码正确, 则返回用户 ID
func (m *UserModel) Authenticate(email, password string) (int, error) {
    return 0, nil
}
// TODO: 检查用户 ID 是否存在
func (m *UserModel) Exists(id int) (bool, error) {
    return false, nil
}
```

### 展示用户注册页

添加路由:

```go
func routes(app *application) http.Handler {
    // 添加 user authentication 相关路由
    router.Handler(http.MethodGet, "/user/signup", dynamic.ThenFunc(app.userSignup))
    router.Handler(http.MethodPost, "/user/signup", dynamic.ThenFunc(app.userSignupPost))
    router.Handler(http.MethodGet, "/user/login", dynamic.ThenFunc(app.userLogin))
    router.Handler(http.MethodPost, "/user/login", dynamic.ThenFunc(app.userLoginPost))
    router.Handler(http.MethodPost, "/user/logout", dynamic.ThenFunc(app.userLogoutPost))
}
```

添加注册页模板:

```html
{{define "title"}}Signup{{end}}
{{define "main"}}
    <form action='/user/signup' method='POST' novalidate>
        <div>
            <label>Name:</label>
            {{with .Form.FieldErrors.name}}
                <label class='error'>{{.}}</label>
            {{end}}
            <input type='text' name='name' value='{{.Form.Name}}'>
        </div>
        <div>
            <label>Email:</label>
            {{with .Form.FieldErrors.email}}
                <label class='error'>{{.}}</label>
            {{end}}
            <input type='email' name='email' value='{{.Form.Email}}'>
        </div>
        <div>
            <label>Password:</label>
            {{with .Form.FieldErrors.password}}
                <label class='error'>{{.}}</label>
            {{end}}
            <!-- 邮箱可以重新展示, 但密码不能, 因为让密码成为 HTML 内容, 就存在被缓存的可能 -->
            <!-- https://ux.stackexchange.com/questions/20418 -->
            <input type='password' name='password'>
        </div>
        <div>
            <input type='submit' value='Signup'>
        </div>
    </form>
{{end}}
```

①定义注册页的表单 ②渲染注册页模板:

```go
type userSignupForm struct {
    Name                string `form:"name"`
    Email               string `form:"email"`
    Password            string `form:"password"`
    validator.Validator `form:"-"`
}

func (app *application) userSignup(w http.ResponseWriter, r *http.Request) {
    data := app.newTemplateData(r)
    data.Form = userSignupForm{}
    app.render(w, http.StatusOK, "signup.tmpl", data)
}
```

### 校验用户注册表单

封装一些校验函数:

```go
//goland:noinspection RegExpRedundantEscape 我就是要用 \\/ 转义 /
var EmailRX = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

func MinChars(value string, n int) bool {
    return utf8.RuneCountInString(value) >= n
}

func Matches(value string, rx *regexp.Regexp) bool {
    return rx.MatchString(value)
}
```

校验逻辑如下:

```go
func (app *application) userSignupPost(w http.ResponseWriter, r *http.Request) {
    var form userSignupForm
    err := app.decodePostForm(r, &form)
    if err != nil {
        app.clientError(w, http.StatusBadRequest)
        return
    }
    // ①要求 name, email, password 非空 ②要求 email 的格式符合邮箱格式
    // ③要求密码长度至少为 8 位         ④要求邮箱不重复
    form.CheckField(validator.NotBlank(form.Name), "name", "This field cannot be blank")
    form.CheckField(validator.NotBlank(form.Email), "email", "This field cannot be blank")
    form.CheckField(validator.Matches(form.Email, validator.EmailRX), "email", "This field must be a valid email address")
    form.CheckField(validator.NotBlank(form.Password), "password", "This field cannot be blank")
    form.CheckField(validator.MinChars(form.Password, 8), "password", "This field must be at least 8 characters long")

    if !form.Valid() {
        data := app.newTemplateData(r)
        data.Form = form
        app.render(w, http.StatusUnprocessableEntity, "signup.tmpl", data)
        return
    }

    fmt.Fprintln(w, "Create a new user...")
}
```

### 加密密码后存储用户

Please download the latest version of the `bcrypt` package: `go get golang.org/x/crypto/bcrypt@latest`

```go
func main() {
    // 第二个参数是 cost, 越高则越难破解, 但执行 GenerateFromPassword 的时间也会变长
    // cost 10 要几十毫秒, cost 12 要几百毫秒, bcrypt.DefaultCost 是 10, 一般情况用 12
    hash, err := bcrypt.GenerateFromPassword([]byte(form.Password), 12)

    // 使用如下函数, 比较密码的 hash 和登录请求中的密码
    hash := []byte("$2a$12$NuTjWXm3KKntReFwyBVHyuf/to.HEwTy.eS206TNfkGfr6GzGJSWG")
    err := bcrypt.CompareHashAndPassword(hash, []byte("the password in a request"))
    if err != nil {
        // 密码不匹配
    }
}
```

Next, update the `UserModel.Insert()` method so that it creates a new record in our `users` table. We also need to manage the potential error caused by a duplicate email violating the `UNIQUE` constraint that we added to the table.  

> I understand that the code in our `UserModel.Insert()` method isn’t very pretty, and that checking the error returned by MySQL feels a bit flaky. What if future versions of MySQL change their error numbers? Or the format of their error messages?
>
> An alternative (but also imperfect) option would be to add an `UserModel.EmailTaken()` method to our model which checks to see if a user with a specific email already exists. We could call this before we try to insert a new record. However, this would introduce a race condition to our application. If two users try to sign up with the same email address at exactly the same time, both submissions will pass the validation check but ultimately only one `INSERT` into the MySQL database will succeed. The other will violate our `UNIQUE` constraint and the user would end up receiving a `500 Internal Server Error` response.  
>
> The outcome of this particular race condition is fairly benign, and some people would advise you to simply not worry about it. But thinking critically about your application logic and writing code which avoids race conditions is a good habit to get into.

```go
func (m *UserModel) Insert(name, email, password string) error {
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
    if err != nil {
        return err
    }
    stmt := `INSERT INTO users (name, email, hashed_password, created) VALUES (?, ?, ?, UTC_TIMESTAMP())`
    _, err = m.DB.Exec(stmt, name, email, string(hashedPassword))
    if err != nil {
        // 检查错误是否为 *mysql.MySQLError 类型, 如果是则继续检查
        // 错误码是否为 1062 ( ER_DUP_ENTRY ) 以及错误消息是否包含 users_uc_email 索引
        var mySQLError *mysql.MySQLError
        if errors.As(err, &mySQLError) {
            if mySQLError.Number == 1062 && strings.Contains(mySQLError.Message, "users_uc_email") {
                return ErrDuplicateEmail // 说明 email 字段重复了
            }
        }
        return err // 是其他错误则直接返回
    }
    return nil
}
```

We can then finish this all off by updating the `userSignup` handler like so:

```go
func (app *application) userSignupPost(w http.ResponseWriter, r *http.Request) {
    var form userSignupForm
    err := app.decodePostForm(r, &form)
    if err != nil {
        app.clientError(w, http.StatusBadRequest)
        return
    }

    form.CheckField(validator.NotBlank(form.Name), "name", "This field cannot be blank")
    form.CheckField(validator.NotBlank(form.Email), "email", "This field cannot be blank")
    form.CheckField(validator.Matches(form.Email, validator.EmailRX), "email", "This field must be a valid email address")
    form.CheckField(validator.NotBlank(form.Password), "password", "This field cannot be blank")
    form.CheckField(validator.MinChars(form.Password, 8), "password", "This field must be at least 8 characters long")

    if !form.Valid() {
        data := app.newTemplateData(r)
        data.Form = form
        app.render(w, http.StatusUnprocessableEntity, "signup.tmpl", data)
        return
    }

    // 尝试存储用户, 如果出现 email 重复则展示错误消息
    err = app.users.Insert(form.Name, form.Email, form.Password)
    if err != nil {
        if errors.Is(err, models.ErrDuplicateEmail) {
            form.AddFieldError("email", "Email address is already in use")
            data := app.newTemplateData(r)
            data.Form = form
            app.render(w, http.StatusUnprocessableEntity, "signup.tmpl", data)
        } else {
            app.serverError(w, err) // 未知错误
        }
        return
    }

    // 用户创建成功则展示一下 flash 消息, 并重定向到登录页
    app.sessionManager.Put(r.Context(), "flash", "Your signup was successful. Please log in.")
    http.Redirect(w, r, "/user/login", http.StatusSeeOther)
}
```

### 展示用户登录页

To support validation errors which aren’t associated with one specific form field:

```go
type Validator struct {
    FieldErrors    map[string]string
    NonFieldErrors []string // 与具体字段无关的错误
}

func (v *Validator) Valid() bool {
    return len(v.FieldErrors) == 0 && len(v.NonFieldErrors) == 0
}

func (v *Validator) AddNonFieldError(message string) {
    v.NonFieldErrors = append(v.NonFieldErrors, message)
}
```

Next let’s create a new `ui/html/pages/login.tmpl` template containing the markup for our login page.

```html
{{define "title"}}Login{{end}}
{{define "main"}}
    <form action='/user/login' method='POST' novalidate>
        <!-- Notice that here we are looping over the NonFieldErrors and
        displaying
        them, if any exist -->
        {{range .Form.NonFieldErrors}}
            <div class='error'>{{.}}</div>
        {{end}}
        <div>
            <label>Email:</label>
            {{with .Form.FieldErrors.email}}
                <label class='error'>{{.}}</label>
            {{end}}
            <input type='email' name='email' value='{{.Form.Email}}'>
        </div>
        <div>
            <label>Password:</label>
            {{with .Form.FieldErrors.password}}
                <label class='error'>{{.}}</label>
            {{end}}
            <input type='password' name='password'>
        </div>
        <div>
            <input type='submit' value='Login'>
        </div>
    </form>
{{end}}
```

Create a new `userLoginForm` struct (to represent and hold the form data):

```go
type userLoginForm struct {
    Email               string `form:"email"`
    Password            string `form:"password"`
    validator.Validator `form:"-"`
}

func (app *application) userLogin(w http.ResponseWriter, r *http.Request) {
    data := app.newTemplateData(r)
    data.Form = userLoginForm{}
    app.render(w, http.StatusOK, "login.tmpl", data)
}
```

### 处理用户登录请求

```go
func (m *UserModel) Authenticate(email, password string) (int, error) {
    // 先从数据库取出用户, 如果用户不存在或无效则返回 ErrInvalidCredentials
    // 然后判断 hash 和 password 是否匹配, 不匹配返回 ErrInvalidCredentials, 匹配则返回用户 ID
    var id int
    var hashedPassword []byte
    stmt := "SELECT id, hashed_password FROM users WHERE email = ?"
    err := m.DB.QueryRow(stmt, email).Scan(&id, &hashedPassword)
    if err != nil {
        if errors.Is(err, sql.ErrNoRows) {
            return 0, ErrInvalidCredentials
        } else {
            return 0, err
        }
    }
    err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(password))
    if err != nil {
        // 判断一下是否为 bcrypt.ErrMismatchedHashAndPassword 错误
        if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
            return 0, ErrInvalidCredentials
        } else {
            return 0, err
        }
    }
    return id, nil
}

func (app *application) userLoginPost(w http.ResponseWriter, r *http.Request) {
    var form userLoginForm
    err := app.decodePostForm(r, &form)
    if err != nil {
        app.clientError(w, http.StatusBadRequest)
        return
    }
    // 别忘了校验 email 和 password 是否为空, 以及邮箱格式是否合法
    form.CheckField(validator.NotBlank(form.Email), "email", "This field cannot be blank")
    form.CheckField(validator.Matches(form.Email, validator.EmailRX), "email", "This field must be a valid email address")
    form.CheckField(validator.NotBlank(form.Password), "password", "This field cannot be blank")
    if !form.Valid() {
        data := app.newTemplateData(r)
        data.Form = form
        app.render(w, http.StatusUnprocessableEntity, "login.tmpl", data)
        return
    }
    // 检查 email 和 password 是否正确, 不正确则展示错误消息
    userID, err := app.users.Authenticate(form.Email, form.Password)
    if err != nil {
        if errors.Is(err, models.ErrInvalidCredentials) {
            form.AddNonFieldError("Email or password is incorrect")
            data := app.newTemplateData(r)
            data.Form = form
            app.render(w, http.StatusUnprocessableEntity, "login.tmpl", data)
        } else {
            app.serverError(w, err)
        }
        return
    }
    // 把用户 ID 存 Session 里面表示用户已经登录了, 然后重定向到首页
    // 注意登录后要换一个新的 Session ID ( 而会话数据会保留 ), 以便防御 Session Fixation 攻击
    err = app.sessionManager.RenewToken(r.Context())
    if err != nil {
        app.serverError(w, err)
        return
    }
    app.sessionManager.Put(r.Context(), "authenticatedUserID", userID)
    http.Redirect(w, r, "/", http.StatusSeeOther)
}
```

### 处理用户登出请求

```go
func (app *application) userLogoutPost(w http.ResponseWriter, r *http.Request) {
    // Renew the Session ID After Any Privilege Level Change. ( https://tinyurl.com/t5vx62ju )
    // Common scenarios to consider include; password changes, permission changes, or
    // switching from a regular user role to an administrator role within the web application.
    err := app.sessionManager.RenewToken(r.Context())
    if err != nil {
        app.serverError(w, err)
        return
    }
    // 移除 authenticatedUserID, 因为判断用户有没有登录就是判断 session 中有没有 authenticatedUserID
    app.sessionManager.Remove(r.Context(), "authenticatedUserID")
    app.sessionManager.Put(r.Context(), "flash", "You've been logged out successfully!")
    http.Redirect(w, r, "/", http.StatusSeeOther)
}
```

## User Authorization

### 认证和验权的区别

- **Authentication** = login + password *(who you are)*
- **Authorization** = permissions *(what you are allowed to do)*

Being able to authenticate the users of our application is all well and good, but now we need to do something useful with that information.

1. Only authenticated (i.e. logged in) users can create a new snippet.

2. Authenticated users should see links to ‘Home’, ‘Create snippet’ and ‘Logout’. 
3. Unauthenticated users should see links to ‘Home’, ‘Signup’ and ‘Login’.

### 判断用户是否已登录

We can check whether a request is being made by an authenticated user or not by checking for the existence of an `"authenticatedUserID"` value in their session data.

```go
func (app *application) isAuthenticated(r *http.Request) bool {
    return app.sessionManager.Exists(r.Context(), "authenticatedUserID")
}
```

The next step is to find a way to pass this information to our HTML templates, so that we can toggle the contents of the navigation bar appropriately.

```go
type templateData struct {
    IsAuthenticated bool // 加个字段, 表示用户是否登录
}

// 这是模板数据的构造函数, 在这里初始化一些通用的模板数据
func (app *application) newTemplateData(r *http.Request) *templateData {
    return &templateData{
        CurrentYear:     time.Now().Year(),
        Flash:           app.sessionManager.PopString(r.Context(), "flash"),
        IsAuthenticated: app.isAuthenticated(r),
    }
}
```

We can update the ui/html/partials/nav.tmpl file to toggle the navigation links like so:

```html
{{define "nav"}}
    <nav>
        <div>
            <a href='/'>Home</a>
            <!-- 已登录用户展示 Create 按钮 -->
            {{if .IsAuthenticated}}
                <a href='/snippet/create'>Create snippet</a>
            {{end}}
        </div>
        <div>
            <!-- 已登录用户展示 Logout 按钮, 未登录用户展示 Signup 和 Login -->
            {{if .IsAuthenticated}}
                <form action='/user/logout' method='POST'>
                    <button>Logout</button>
                </form>
            {{else}}
                <a href='/user/signup'>Signup</a>
                <a href='/user/login'>Login</a>
            {{end}}
        </div>
    </nav>
{{end}}
```

### 用中间件设置访问权限校验

We’re hiding the ‘Create snippet’ navigation link for any user that isn’t logged in. But an unauthenticated user could still create a new snippet by visiting the `/snippet/create` page directly. Let’s fix that, so that if an unauthenticated user tries to visit any routes with the URL path `/snippet/create` they are redirected to `/user/login` instead.

```go
func (app *application) requireAuthentication(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        if !app.isAuthenticated(r) {
            http.Redirect(w, r, "/user/login", http.StatusSeeOther) // 未登录则重定向到登录页
            return                                                  // 不让后续 handler 被执行
        }
        // Set the "Cache-Control: no-store" header so that pages require authentication
        // are not stored in the users browser cache (or other intermediary cache).
        w.Header().Add("Cache-Control", "no-store")
        next.ServeHTTP(w, r)
    })
}

func routes(app *application) http.Handler {
    router := httprouter.New()
    // 这是为动态内容创建的 middleware chain, 而静态内容不需要使用 session
    dynamic := alice.New(app.sessionManager.LoadAndSave)
    router.Handler(http.MethodGet, "/", dynamic.ThenFunc(app.homePage))                        // 列表页
    router.Handler(http.MethodGet, "/snippet/view/:id", dynamic.ThenFunc(app.snippetViewPage)) // 详情页
    router.Handler(http.MethodGet, "/user/signup", dynamic.ThenFunc(app.userSignup))           // signup page
    router.Handler(http.MethodPost, "/user/signup", dynamic.ThenFunc(app.userSignupPost))      // signup api
    router.Handler(http.MethodGet, "/user/login", dynamic.ThenFunc(app.userLogin))             // login page
    router.Handler(http.MethodPost, "/user/login", dynamic.ThenFunc(app.userLoginPost))        // login api

    // 在 session 中间件后面加一个检查是否登录的中间件, 下面的东西要登录后才能访问
    protected := dynamic.Append(app.requireAuthentication)
    router.Handler(http.MethodGet, "/snippet/create", protected.ThenFunc(app.snippetCreate))      // create page
    router.Handler(http.MethodPost, "/snippet/create", protected.ThenFunc(app.snippetCreatePost)) // create api
    router.Handler(http.MethodPost, "/user/logout", protected.ThenFunc(app.userLogoutPost))       // logout api

    // 自定义 404 Handler, 统一用 app.NotFound() 处理
    router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        app.notFound(w)
    })

    // 创建 File Server, 所有以 /static/ 开头的 URL 都交给它处理
    // 把请求交给 File Server 处理前, 先去掉请求 URL 中的 /static 前缀
    fileServer := http.FileServer(http.Dir("./web/ui/static/"))
    router.Handler(http.MethodGet, "/static/*filepath", http.StripPrefix("/static", fileServer))

    // 创建 standard middleware chain, 并把它用于每一个请求
    standard := alice.New(app.recoverPanic, app.logRequest, secureHeaders)
    return standard.Then(router)
}
```

### 是时候加 CSRF 防护了

One mitigation that we can take to prevent CSRF attacks is to make sure that the `SameSite` attribute is appropriately set on our session cookie. By default the `alexedwards/scs` package that we’re using always sets `SameSite=Lax` on the session cookie. 

- [看看这里了解一下 Cookie 的 SameSite 属性，设置 SameSite=Lax 基本就杜绝了 CSRF 攻击](https://www.ruanyifeng.com/blog/2019/09/cookie-samesite.html#:~:text=%E4%BA%86%E4%BB%80%E4%B9%88%E7%BD%91%E7%AB%99%E3%80%82-,%E4%BA%8C%E3%80%81SameSite%20%E5%B1%9E%E6%80%A7,-Cookie%20%E7%9A%84)
- `SameSite=Lax` 会让 GET 型跨站请求会带上 cookie，但 POST/PUT/... 等修改状态的请求则不发送 cookie

So long as our application uses the `POST` method for any state-changing HTTP requests (like our login, signup, logout and create snippet form submissions), it means that the session cookie won’t be sent for these requests if they are cross-site — which in turn means that they should be safe from CSRF attacks. However, the `SameSite` attribute is still relatively new and only fully supported by 94% of browsers worldwide. So, although it’s something that we can (and should) use as a defensive measure, we can’t rely on it for all users.

A good and popular choice is `justinas/nosurf`, which uses the double-submit cookie pattern to prevent CSRF attacks. In this pattern a random CSRF token is generated and sent to the user in a CSRF cookie. This CSRF token is then added to a hidden field in each HTML form that’s vulnerable to CSRF. When the form is submitted, it uses some middleware to check that the hidden field value and cookie value match. 

You can install the latest version like so:  `go get github.com/justinas/nosurf@v1`

```go
func noSurf(next http.Handler) http.Handler {
    csrfHandler := nosurf.New(next)
    csrfHandler.SetBaseCookie(http.Cookie{
        HttpOnly: true,
        Path:     "/",
        Secure:   true,
    })
    return csrfHandler
}
```

One of the forms that we need to protect from CSRF attacks is our logout form, which is included in our `nav.tmpl` partial and could potentially appear on any page of our application. So, because of this, we need to use our `noSurf()` middleware on all of our application routes (apart from `/static/*filepath`).

```go
func routes(app *application) http.Handler {
    router := httprouter.New()
    // session、csrf 中间件会应用于所有后端接口 ( 接口中可选用不用这些功能 )
    dynamic := alice.New(app.sessionManager.LoadAndSave, noSurf)
}
```

At this point, you might like to fire up the application and try submitting one of the forms. When you do, the request should be intercepted by the `noSurf()` middleware and you should receive a `400 Bad Request` response. To make the form submissions work, we need to use the `nosurf.Token()` function to get the CSRF token and add it to a hidden `csrf_token` field in each of our forms. So the next step is to add a new `CSRFToken` field to our `templateData` struct:

```go
type templateData struct {
    CSRFToken string
}
```

And because the logout form can potentially appear on every page, it makes sense to add the CSRF token to the template data automatically via our `newTemplateData()` helper. This will mean that it’s available to our templates each time we render a page.

```go
func (app *application) newTemplateData(r *http.Request) *templateData {
    return &templateData{
        CurrentYear:     time.Now().Year(),
        Flash:           app.sessionManager.PopString(r.Context(), "flash"),
        IsAuthenticated: app.isAuthenticated(r),
        CSRFToken:       nosurf.Token(r), // Add the CSRF token
    }
}
```

Finally, we need to update all the forms in our application to include this CSRF token in a hidden field.

```html
<!-- 在模板文件中搜索 form 标签, 然后全都加上 csrf_token 隐藏字段 -->
<form action='/user/logout' method='POST'>
    <input type='hidden' name='csrf_token' value='{{.CSRFToken}}'>
    <button>Logout</button>
</form>
```

什么时候要加 CSRF 防护 ? 

1. 所有表单都要加 CSRF Token 
2. You need CSRF protection if you use implicit authentication (typically cookies) 
3. but not if you use explicit authentication (a header added by javascript).

如果满足这些条件，能完全依靠 SameSite 属性做 CSRF 防护，无需像上面一样使用 csrf_token:

1. Make TLS 1.3 the minimum supported version in the TLS config for your server, then all browsers able to use your application will support `SameSite` cookies ( due to the fact that there is no browser that exists which supports TLS 1.3 and does not support `SameSite` cookies ).
2. Set `SameSite=Lax` or `SameSite=Strict` on the session cookie.
3. Use an ‘unsafe’ HTTP method (i.e. `POST`, `PUT` or `DELETE`) for any state-changing requests.

## 请求上下文

### 为何需要 Request Context

At the moment our logic for authenticating a user consists of simply checking whether a `"authenticatedUserID"` value exists in their session data, like so:

```go
func (app *application) isAuthenticated(r *http.Request) bool {
    return app.sessionManager.Exists(r.Context(), "authenticatedUserID")
}
```

We could make this check more robust by querying our users database table to make sure that the `"authenticatedUserID"` value is a real, valid, value (i.e we haven’t deleted the user’s account since they last logged in). 

But there is a slight problem with doing this additional database check. Our `isAuthenticated()` helper can potentially be called multiple times in each request cycle. Currently we use it twice — once in the `requireAuthentication()` middleware and again in the `newTemplateData()` helper. So, if we query the database from the `isAuthenticated()` helper directly, we would end up making duplicated round-trips to the database during every request. And that’s not very efficient.

A better approach would be to carry out this check in some middleware to determine whether the current request is from an authenticated user or not, and then pass that information down to all subsequent handlers in the chain.

### 如何使用 Request Context

Every `http.Request` that our middleware and handlers process has a `context.Context` object embedded in it, which we can use to store information during the lifetime of the request. A common use-case for this is to pass information between your pieces of middleware and other handlers. In our case, we want to use it to check if a user is authenticated once in some middleware, and if they are, then make this information available to all our other middleware and handlers.

How to add data to a request’s context:

- First, we use the `r.Context()` method to retrieve the existing context from a request.

- Then we use the `context.WithValue()` method to create a new copy of the existing context, containing the key
  `"isAuthenticated"` and a value of `true`.

- Then finally we use the `r.WithContext()` method to create a copy of the request containing our new context. Notice that we don’t actually update the context for a request directly. What we’re doing is creating a new copy of the `http.Request` object with our new context in it.

```go
ctx := r.Context()
ctx = context.WithValue(ctx, "isAuthenticated", true) // 复制一下 ctx 并额外添加 isAuthenticated 字段
r = r.WithContext(ctx)                                // 复制一下 r 并把 Context 设为 ctx, 然后把结果赋值回 r 变量

// for clarity, I made the code a bit more verbose than it needs to be. It’s more typical to write it like this:
ctx = context.WithValue(r.Context(), "isAuthenticated", true)
r = r.WithContext(ctx)
```

How to retrieve data? The important thing to explain here is that, behind the scenes, request context values are stored with the type `any`. And that means that, after retrieving them from the context, you’ll need to assert them to their original type before you use them. To retrieve a value we need to use the `r.Context().Value()` method, like so:

```go
isAuthenticated, ok := r.Context().Value("isAuthenticated").(bool) // 因为返回 any 类型所以要进行 type assert
if !ok {
    return errors.New("could not convert value to bool")
}
```

### 如何避免 Key Collisions  

In the code samples above, I’ve used the string `"isAuthenticated"` as the key for storing and retrieving the data from a request’s context. But this isn’t recommended because there’s a risk that other third-party packages used by your application will also want to store data using the key `"isAuthenticated"` — and that would cause a naming collision. To avoid this, it’s good practice to create your own custom type which you can use for your context keys. 

自定义类型能解决 key 冲突，因为判断相等性时要看类型，比如 `hour(1) != minute(1)`，然后 `hour` 类型又不导出只有我能用:

```go
// Declare a custom "contextKey" type for your context keys.
type contextKey string
// Create a constant with the type contextKey that we can use.
const isAuthenticatedContextKey = contextKey("isAuthenticated")

// Set the value in the request context, using our isAuthenticatedContextKey constant as the key.
ctx = context.WithValue(r.Context(), isAuthenticatedContextKey, true)
r = r.WithContext(ctx)

// Retrieve the value from the request context using our constant as the key.
isAuthenticated, ok := r.Context().Value(isAuthenticatedContextKey).(bool)
if !ok {
    return errors.New("could not convert value to bool")
}
```

### 中间件: 判断用户是否登录

Let’s start to use the request context functionality in our application. We’ll begin by updating the `UserModel.Exists()` method, so that it returns `true` if a user with a specific ID exists in our `users` table.

```go
func (m *UserModel) Exists(id int) (bool, error) {
    var exists bool
    stmt := "SELECT EXISTS(SELECT true FROM users WHERE id = ?)" // 判断用户是否存在
    err := m.DB.QueryRow(stmt, id).Scan(&exists)
    return exists, err
}
```

Let’s create a new `authenticate()` middleware method which:

1. Retrieves the user’s ID from their session data.
2. Checks the database to see if the ID corresponds to a valid user using the `UserModel.Exists()` method.
3. Updates the request context to include an `isAuthenticatedContextKey` key with the value `true`.  
   ( We create a copy of the request with a `isAuthenticatedContextKey` key and `true` value stored in the request context. We then pass this copy of the `*http.Request` to the next handler in the chain. )

```go
type contextKey string                                          // 定义自定义类型, 防止 key collision
const isAuthenticatedContextKey = contextKey("isAuthenticated") // 定义常量, 用于存取 request context

func (app *application) authenticate(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // 从 session 中取 authenticatedUserID,  如果不存在 GetInt 方法会返回零值
        id := app.sessionManager.GetInt(r.Context(), "authenticatedUserID")
        if id == 0 {
            next.ServeHTTP(w, r) // 用户未登录, 无需做什么特殊处理, 直接调用下一个 handler
            return
        }
        // 检查 authenticatedUserID 对应的用户是否合法 ( 可能被删了或禁用了 )
        exists, err := app.users.Exists(id)
        if err != nil {
            app.serverError(w, err)
            return
        }
        // 若合法则在 request context 存储结果, 方便后续 handler 判断用户是否登录, 避免再次查询
        if exists {
            ctx := context.WithValue(r.Context(), isAuthenticatedContextKey, true)
            r = r.WithContext(ctx) // 下面调用 ServeHTTP 时 r 被替换成了带 Context 的
        }
        next.ServeHTTP(w, r)
    })
}

func (app *application) isAuthenticated(r *http.Request) bool {
    // 像这样只判断 authenticatedUserID 在不在 session 中不够严谨 ( 比如 session 未过期但用户被禁用了 )
    // return app.sessionManager.Exists(r.Context(), "authenticatedUserID")

    // 由 authenticate 中间件判断用户是否登录, 这个 helper 函数负责取出结果
    isAuthenticated, ok := r.Context().Value(isAuthenticatedContextKey).(bool)

    // 此处 type assert 返回 !ok 只可能是 nil interface ( 类型一定是 bool, 因为自己就是这么存的 )
    if !ok {
        return false // 说明设置过 isAuthenticatedContextKey, 所以返回 false
    }
    return isAuthenticated
}
```

Let’s update the `cmd/web/routes.go` file to include the `authenticate()` middleware in our `dynamic` middleware chain:

```go
func routes(app *application) http.Handler {
    router := httprouter.New()
    // session、csrf、authenticate 中间件会应用于所有后端接口 ( 接口中可选用不用这些功能 )
    dynamic := alice.New(app.sessionManager.LoadAndSave, noSurf, app.authenticate)
}
```

Then, if you want, open MySQL and delete the record for the user that you’re logged in as from the database. And when you go back to your browser and refresh the page, the application is now smart enough to recognize that the user has been deleted, and you’ll find yourself treated as an unauthenticated (logged-out) user.  

### 不要滥用 Request Context

It’s important to emphasize that request context should only be used to store information relevant to the lifetime of a specific request. The Go documentation for `context.Context` warns:

> Use context Values only for request-scoped data that transits processes and APIs.  

That means you should not use it to pass dependencies that exist outside of the lifetime of a request — like loggers, template caches and your database connection pool — to your middleware and handlers.

For reasons of type-safety and clarity of code, it’s almost always better to make these dependencies ( loggers, template caches and your database connection pool ) available to your handlers explicitly, by either making your handlers methods against an `application` struct (like we have in this book) or passing them in a closure.

