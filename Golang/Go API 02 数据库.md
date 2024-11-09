## Table of Contents
  - [连接数据库](#%E8%BF%9E%E6%8E%A5%E6%95%B0%E6%8D%AE%E5%BA%93)
    - [例子](#%E4%BE%8B%E5%AD%90)
    - [连接池概述](#%E8%BF%9E%E6%8E%A5%E6%B1%A0%E6%A6%82%E8%BF%B0)
    - [连接池配置](#%E8%BF%9E%E6%8E%A5%E6%B1%A0%E9%85%8D%E7%BD%AE)
  - [SQL Migrations](#SQL-Migrations)
    - [有啥用](#%E6%9C%89%E5%95%A5%E7%94%A8)
    - [使用流程](#%E4%BD%BF%E7%94%A8%E6%B5%81%E7%A8%8B)
    - [处理错误](#%E5%A4%84%E7%90%86%E9%94%99%E8%AF%AF)
  - [CRUD](#CRUD)
    - [创建 Model](#%E5%88%9B%E5%BB%BA-Model)
    - [实现 Insert()](#%E5%AE%9E%E7%8E%B0-Insert)
    - [实现 Get()](#%E5%AE%9E%E7%8E%B0-Get)
    - [实现 Update()](#%E5%AE%9E%E7%8E%B0-Update)
    - [实现 Delete()](#%E5%AE%9E%E7%8E%B0-Delete)
    - [部分更新](#%E9%83%A8%E5%88%86%E6%9B%B4%E6%96%B0)
    - [版本号乐观锁](#%E7%89%88%E6%9C%AC%E5%8F%B7%E4%B9%90%E8%A7%82%E9%94%81)
    - [设置查询超时](#%E8%AE%BE%E7%BD%AE%E6%9F%A5%E8%AF%A2%E8%B6%85%E6%97%B6)
  - [CRUD - 列表](#CRUD--%E5%88%97%E8%A1%A8)
    - [查询参数](#%E6%9F%A5%E8%AF%A2%E5%8F%82%E6%95%B0)
    - [返回列表](#%E8%BF%94%E5%9B%9E%E5%88%97%E8%A1%A8)
    - [筛选电影](#%E7%AD%9B%E9%80%89%E7%94%B5%E5%BD%B1)
    - [全文搜索](#%E5%85%A8%E6%96%87%E6%90%9C%E7%B4%A2)
    - [排序](#%E6%8E%92%E5%BA%8F)
    - [分页](#%E5%88%86%E9%A1%B5)

## 连接数据库

### 例子

To work with a SQL database we need to use a database driver to act as a ‘middleman’ between Go and the database itself.

```bash
go get github.com/lib/pq@v1
```

We’ll also need a *data source name* (DSN), which is basically a string that contains the necessary connection parameters. The exact format of the DSN will depend on which database driver you’re using (and should be described in the driver documentation), but when using `pq` you should be able to connect to your local `greenlight` database as the `greenlight` user with the following DSN:

```bash
postgres://greenlight:pa55word@localhost/greenlight
```

#### ➤ 代码例子

```go
func main() {
    var cfg config
    flag.IntVar(&cfg.port, "port", 4000, "API Server Port")
    flag.StringVar(&cfg.env, "env", "development", "Environment (development|production)")
    // 推荐在 .bashrc 中配置 export GREENLIGHT_DB_DSN='postgres://greenlight:pa55word@localhost/greenlight?sslmode=disable'
    flag.StringVar(&cfg.db.dsn, "db-dsn", os.Getenv("GREENLIGHT_DB_DSN"), "PostgreSQL DSN")
    flag.Parse()

    // 初始化日志
    logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

    // 初始化数据库
    db, err := openDB(cfg)
    if err != nil {
        logger.Error(err.Error())
        os.Exit(1)
    }
    defer db.Close() // 关闭数据库连接池
    logger.Info("database connection pool established")

    app := &application{
        config: cfg,
        logger: logger,
    }

    // 创建 Server
    srv := &http.Server{
        Addr:         fmt.Sprintf(":%d", cfg.port),
        Handler:      app.routes(),
        IdleTimeout:  time.Minute,
        ReadTimeout:  10 * time.Second,
        WriteTimeout: 30 * time.Second,
        ErrorLog:     slog.NewLogLogger(logger.Handler(), slog.LevelError),
    }

    logger.Info("starting server", "addr", srv.Addr, "env", cfg.env)
    err = srv.ListenAndServe()
    logger.Error(err.Error())
    os.Exit(1)
}

func openDB(cfg config) (*sql.DB, error) {
    db, err := sql.Open("postgres", cfg.db.dsn)
    if err != nil {
        return nil, err
    }

    // 连接池是惰性的, 首次用到连接时才会建立连接, 所以调用 Ping 建个连接试试有没有错
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    err = db.PingContext(ctx)
    if err != nil {
        _ = db.Close()
        return nil, err
    }

    return db, nil
}
```

### 连接池概述

A `sql.DB` pool contains two types of connections — ‘in-use’ connections and ‘idle’ connections. 

When you instruct Go to perform a database task, it will first check if any idle connections are available in the pool. If one is available, then Go will reuse this existing connection and mark it as in-use for the duration of the task. If there are no idle connections in the pool when you need one, then Go will create a new additional connection.  

When Go reuses an idle connection from the pool, any problems with the connection are handled gracefully. Bad connections will automatically be re-tried twice before giving up, at which point Go will remove the bad connection from the pool and create a new one to carry out the task.  

### 连接池配置

#### ➤ SetMaxOpenConns(): 调低能限流，调高增加并发查询上限

The `SetMaxOpenConns()` method allows you to set an upper `MaxOpenConns` limit on the number of ‘open’ connections (in-use + idle connections) in the pool. By default, the number of open connections is unlimited. Broadly speaking, the higher that you set the `MaxOpenConns` limit, the more database queries can be performed concurrently and the lower the risk is that the *connection pool itself* will be a bottleneck in your application.

But leaving it unlimited isn’t necessarily the best thing to do. *By default PostgreSQL has a hard limit of 100 open connections* and, if this hard limit is hit under heavy load, it will cause our `pq` driver to return a "sorry, too many clients already" error. To avoid this error, it makes sense limit the number of open connections in our pool to comfortably *below 100* — leaving enough headroom for any other applications or sessions that also need to use PostgreSQL.

But setting a limit comes with an important caveat. If the `MaxOpenConns` limit is reached, and all connections are in-use, then any further database tasks will be forced to wait until a connection becomes free and marked as idle. In the context of our API, the user’s HTTP request could *‘hang’ indefinitely* while waiting for a free connection. So to mitigate this, it’s important to *always set a timeout on database tasks* using a `context.Context` object.

#### ➤ SetMaxIdleConns(): 适当保留连接能提升性能，保留太多连接会浪费内存

The `SetMaxIdleConns()` method sets an upper `MaxIdleConns` limit on the number of idle connections in the pool. By default, the maximum number of idle connections is 2. In theory, allowing a higher number of idle connections in the pool will improve performance because it makes it less likely that a new connection needs to be established from scratch.

But it’s also important to realize that keeping an idle connection alive comes at a cost. It takes up memory which can otherwise be used for your application and database, and it’s also possible that if a connection is idle for too long then it may become unusable. For example, by default MySQL will automatically close any connections which haven’t been used for 8 hours. So, potentially, setting `MaxIdleConns` too high may result in more connections becoming unusable and more memory resources being used than if you had a smaller idle connection pool.

#### ➤ SetConnMaxLifetime(): 例如让连接在 1h 后不会再被复用

The `SetConnMaxLifetime()` method sets the `ConnMaxLifetime` limit — the maximum length of time that a connection can be reused for. By default, there’s no maximum lifetime and connections will be reused forever. If we set `ConnMaxLifetime` to one hour, for example, it means that all connections will be marked as ‘expired’ one hour after they were first created, and cannot be reused after they’ve expired.

#### ➤ SetConnMaxIdleTime(): 例如让连接在空闲 1h 后被清理回收资源

The `SetConnMaxIdleTime()` method sets the `ConnMaxIdleTime` limit. By default there’s no limit.  If we set `ConnMaxIdleTime` to 1 hour, for example, any connections that have sat idle in the pool for 1 hour since last being used will be marked as expired and removed by the background cleanup operation.

#### ➤ 怎么配这些参数?

For this project we’ll set a `MaxOpenConns` limit of 25 connections. I’ve found this to be a reasonable starting point for small-to-medium web applications and APIs, but ideally you should tweak this value for your hardware depending on the results of *benchmarking and load-testing*.

```go
func main() {
    // db 连接数上限设为 25, 连接的空闲回收时间设为 15 min
    flag.IntVar(&cfg.db.maxOpenConns, "db-max-open-conns", 25, "PostgreSQL max open connections")
    flag.IntVar(&cfg.db.maxIdleConns, "db-max-idle-conns", 25, "PostgreSQL max idle connections")
    flag.DurationVar(&cfg.db.maxIdleTime, "db-max-idle-time", 15*time.Minute, "PostgreSQL max connection idle time")
    flag.Parse()
}

func openDB(cfg config) (*sql.DB, error) {
    db, err := sql.Open("postgres", cfg.db.dsn)
    if err != nil {
        return nil, err
    }
    // 设置连接池参数
    db.SetMaxOpenConns(cfg.db.maxOpenConns)
    db.SetMaxIdleConns(cfg.db.maxIdleConns)
    db.SetConnMaxIdleTime(cfg.db.maxIdleTime)

    // 连接池是惰性的, 首次用到连接时才会建立连接, 所以调用 Ping 建个连接试试有没有错
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    err = db.PingContext(ctx)
    if err != nil {
        _ = db.Close()
        return nil, err
    }

    return db, nil
}
```

## SQL Migrations

### 有啥用

We could simply use the `psql` tool again and run the necessary `CREATE TABLE` statement against our database. But instead, we’re going to explore how to use SQL migrations to create the table (and more generally, manage database schema changes throughout the project). In case you’re not familiar with the idea of SQL migrations, at a very high-level the concept works like this: 

1. For every change that you want to make to your database schema (like creating a table, adding a column, or removing an unused index) you create a pair of migration files. One file is the ‘up’ migration which contains the SQL statements necessary to implement the change, and the other is a ‘down’ migration which contains the SQL statements to reverse (or roll-back) the change.
2. Each pair of migration files is numbered sequentially, usually `0001, 0002, 0003...` or with a [Unix timestamp](https://en.wikipedia.org/wiki/Unix_time), to indicate the order in which migrations should be applied to a database.
3. You use some kind of tool or script to execute or rollback the SQL statements in the sequential migration files against your database. The tool keeps track of which migrations have already been applied, so that only the necessary SQL statements are actually executed.

Using migrations to manage your database schema, rather than manually executing the SQL statements yourself, has a few benefits:  

- The database schema (along with its evolution and changes) is completely described by the ‘up’ and ‘down’ SQL migration files. And because these are just regular files containing some SQL statements, they can be *included and tracked* alongside the rest of your code in a version control system.
- It’s possible to replicate the current database schema precisely on another machine by running the necessary ‘up’ migrations. This is a big help when you need to manage and synchronize database schemas in different environments (development, testing, production, etc.)
- It’s possible to roll-back database schema changes if necessary by applying the appropriate ‘down’ migrations.


### 使用流程

#### ➤ 安装 golang-migrate

1. [去这里下载程序](https://github.com/golang-migrate/migrate/releases)，然后把文件移动到 $GOPATH/bin/
2. 另外 Linux 中可以这样做

```bash
curl -L https://github.com/golang-migrate/migrate/releases/download/v4.17.0/migrate.linux-amd64.tar.gz | tar xvz
mv ./migrate $GOPATH/bin/migrate
```

#### ➤ 首先生成一对 up/down 文件

每次执行数据库修改，都用 `create` 命令创建一对 sql 文件，分别包含修改/撤销代码

```bash
# The -seq flag indicates that we want to use sequential numbering like 0001, 0002, ...
# for the migration files (instead of a Unix timestamp, which is the default).  
migrate create -seq -ext=.sql -dir=./migrations create_movies_table

# 一个文件中包含多个建表语句，也是可以的
migrate create -seq -ext=.sql -dir=./migrations create_tables
```

如下是建表 SQL 和 撤销 SQL:

```postgresql
-- file: 000001_create_movies_table.up.sql
CREATE TABLE IF NOT EXISTS movies
(
    id         bigserial PRIMARY KEY,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    title      text                        NOT NULL,
    year       integer                     NOT NULL,
    runtime    integer                     NOT NULL,
    genres     text[]                      NOT NULL,
    version    integer                     NOT NULL DEFAULT 1
);
-- file: 000001_create_movies_table.down.sql
DROP TABLE IF EXISTS movies;
```

#### ➤ 执行 migration

```bash
set DSN 'postgres://greenlight:pa55word@localhost/greenlight?sslmode=disable'
migrate -path=./migrations -database=$DSN up
```

You should see that the movies table has been created, along with a `schema_migrations` table, both of which are owned by the greenlight user. The `schema_migrations` table is automatically generated by the migrate tool and used to keep track of which migrations have been applied.

```sql
select * from schema_migrations;
# version | dirty
#---------+-------
#       2 | f
```

The `version` column here indicates that our migration files up to (and including) number `2` in the sequence have been executed against the database. The value of the `dirty` column is `false`, which indicates that the migration files were cleanly executed without any errors and the SQL statements they contain were successfully applied in full.  

```bash
migrate -path=./migrations -database=$DSN version   # 查看数据库的 migration version
migrate -path=./migrations -database =$DSN down 1   # 回退一个版本
migrate -path=./migrations -database =$DSN up 1     # 上升一个版本
migrate -path=./migrations -database=$DSN goto 1    # 设为指定版本, 推荐用这个, 比 up/down 更清晰

migrate -path=./migrations -database=$DSN up        # apply all up migrations
migrate -path=./migrations -database=$DSN down      # apply all down migrations
```

### 处理错误

When you run a migration that contains an error, all SQL statements up to the erroneous one will be applied and then the `migrate` tool will exit with a message describing the error.

#### ➤ 需要手动诊断文件中的哪些 SQL 已经执行，然后手动回滚这些 SQL，最后改一下 version

If the migration file which failed contained multiple SQL statements, then it’s possible that the migration file was *partially* applied before the error was encountered. In turn, this means that the database is in an unknown state as far as the `migrate` tool is concerned.

Accordingly, the `version` field in the `schema_migrations` field will contain the number for the *failed migration* and the `dirty` field will be set to `true`. At this point, if you run another migration (even a ‘down’ migration) you will get an error message similar to this:

```bash
Dirty database version {X}. Fix and force version.
```

What you need to do is investigate the original error and figure out if the migration file which failed was partially applied. If it was, then you need to manually roll-back the partially applied migration. Once that’s done, then you must also ‘force’ the `version` number in the `schema_migrations` table to the correct value. For example, to force the database version number to `1` you should use the `force` command like so:

```bash
$ migrate -path=./migrations -database=$DSN force 1
```

Once you force the version, the database is considered ‘clean’ and you should be able to run migrations again without any problem.

## CRUD

### 创建 Model

If you don’t like the term *model* then you might want to think of this as your *data access* or *storage* layer instead. But whatever you prefer to call it, the principle is the same — it will *encapsulate all the code for reading and writing movie data to and from our PostgreSQL database*.

#### ➤ 怎样组织代码

```go
// 为了直观地调用各类 Model 方法，例如 app.models.Movies.Insert(...)
type application struct {
    models data.Models
}

// 注册各类 Model, 例如 Movies, Users, ...
type Models struct {
    Movies MovieModel
}

func NewModels(db *sql.DB) Models {
    return Models{
        Movies: MovieModel{DB: db},
    }
}

// 各类 Model 封装了相应的 CRUD 代码
type MovieModel struct {
    DB *sql.DB
}

func (m MovieModel) Insert(movie *Movie) error {
    return nil
}
```

### 实现 Insert()

We want it to execute the following SQL query:

```postgresql
INSERT INTO movies (title, year, runtime, genres) 
VALUES ($1, $2, $3, $4)
RETURNING id, created_at, version
```

There are few things about this query which warrant a bit of explanation.

- It uses `$N` notation to represent *placeholder parameters* for the data that we want to insert in the `movies` table. Every time that you pass untrusted input data from a client to a SQL database it’s important to use placeholder parameters to help prevent SQL injection attacks, unless you have a very specific reason for not using them.
- At the end of the query we have a [`RETURNING`](https://www.postgresql.org/docs/current/dml-returning.html) clause. This is a PostgreSQL-specific clause (it’s not part of the SQL standard) that you can use to return values from any record that is being manipulated by an `INSERT`, `UPDATE` or `DELETE` statement. In this query we’re using it to return the system-generated `id`, `created_at` and `version` values.

Throughout this project we’ll stick with using Go’s [`database/sql`](https://golang.org/pkg/database/sql/) package to execute our database queries, rather than using a third-party [ORM](https://github.com/avelino/awesome-go#orm) or [other tool](https://github.com/avelino/awesome-go#database). Normally, you would use Go’s [`Exec()`](https://golang.org/pkg/database/sql/#DB.Exec) method to execute an `INSERT` statement against a database table. But because our SQL query is returning a single row of data (thanks to the `RETURNING` clause), we’ll need to use the [`QueryRow()`](https://golang.org/pkg/database/sql/#DB.QueryRow) method here instead.

```go
func (m MovieModel) Insert(movie *Movie) error {
    // 编写 SQL 并填充参数
    query := `
        INSERT INTO movies (title, year, runtime, genres) 
        VALUES ($1, $2, $3, $4)
        RETURNING id, created_at, version`
    
    args := []any{movie.Title, movie.Year, movie.Runtime, pq.Array(movie.Genres)}

    // 执行 QueryRow() 并用 Scan() 读取返回的数据
    return m.DB.QueryRow(query, args...).Scan(&movie.ID, &movie.CreatedAt, &movie.Version)
}
```

- Because the `Insert()` method signature takes a `*Movie` pointer as the parameter, when we call `Scan()` to read in the system-generated data we’re updating the values *at the location the parameter points to*. Essentially, our `Insert()` method *mutates* the `Movie` struct that we pass to it and adds the system-generated values to it.
- Storing the inputs in a slice isn’t strictly necessary, but as mentioned in the code comments above it’s a nice pattern that can help the clarity of your code. 
- Also, notice the final value in the slice? In order to store our `movie.Genres` value (which is a `[]string` slice) in the database, we need to pass it through the [`pq.Array()`](https://pkg.go.dev/github.com/lib/pq?tab=doc#Array) adapter function before executing the SQL query. 
- Behind the scenes, the `pq.Array()` adapter takes our `[]string` slice and converts it to a [`pq.StringArray`](https://pkg.go.dev/github.com/lib/pq#StringArray) type. In turn, the `pq.StringArray` type implements the [`driver.Valuer`](https://pkg.go.dev/database/sql/driver?tab=doc#Valuer) and [`sql.Scanner`](https://pkg.go.dev/database/sql?tab=doc#Scanner) interfaces necessary to translate our native `[]string` slice to and from a value that our PostgreSQL database can understand and store in a `text[]` array column.
- You can also use the `pq.Array()` adapter function in the same way with `[]bool`, `[]byte`, `[]int32`, `[]int64`, `[]float32` and `[]float64` slices in your Go code.

Let’s hook up the `Insert()` method to our `createMovieHandler` so that our `POST /v1/movies` endpoint works in full:

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

    // 执行 CRUD, 注意 Insert() 会用数据库生成的 id, created_at, version 更新 movie 结构
    err = app.models.Movies.Insert(movie)
    if err != nil {
        app.serverErrorResponse(w, r, err)
        return
    }

    // 资源创建成功后, 发个 201 响应, 用 Location 表示资源所在的 URL
    headers := make(http.Header)
    headers.Set("Location", fmt.Sprintf("/v1/movies/%d", movie.ID))

    // 返回创建的资源
    err = app.writeJSON(w, http.StatusCreated, envelop{"movie": movie}, headers)
    if err != nil {
        app.serverErrorResponse(w, r, err)
        return
    }
}
```

### 实现 Get()

Model:

```go
func (m MovieModel) Get(id int64) (*Movie, error) {
    query := `
        SELECT id, created_at, title, year, runtime, genres, version
        FROM movies
        WHERE id = $1`

    var movie Movie

    err := m.DB.QueryRow(query, id).Scan(
        &movie.ID,
        &movie.CreatedAt,
        &movie.Title,
        &movie.Year,
        &movie.Runtime,
        pq.Array(&movie.Genres), // 查询和解析数组类型, 都要通过 pq.Array 处理
        &movie.Version,
    )

    // 返回自定义错误, 而不是底层数据库的错误, 方便解耦, 切换数据库实现
    if err != nil {
        switch {
        case errors.Is(err, sql.ErrNoRows):
            return nil, ErrRecordNotFound
        default:
            return nil, err
        }
    }

    return &movie, nil
}
```

Handler:

```go
func (app *application) showMovieHandler(w http.ResponseWriter, r *http.Request) {
    id, err := app.readIDParam(r)
    if err != nil { // 无需判断 id < 1, 已经被 readIDParam 处理了
        app.notFoundResponse(w, r)
        return
    }

    // 判断错误类型
    movie, err := app.models.Movies.Get(id)
    if err != nil {
        switch {
        case errors.Is(err, data.ErrRecordNotFound):
            app.notFoundResponse(w, r)
        default:
            app.serverErrorResponse(w, r, err)
        }
        return
    }

    err = app.writeJSON(w, http.StatusOK, envelop{"movie": movie}, nil)
    if err != nil {
        app.serverErrorResponse(w, r, err)
    }
}
```

### 实现 Update()

Model:

```go
func (m MovieModel) Update(movie *Movie) error {
    // 注意让 version 加一, 并返回更新后的 version
    query := `
        UPDATE movies 
        SET title = $1, year = $2, runtime = $3, genres = $4, version = version + 1
        WHERE id = $5
        RETURNING version`

    args := []any{movie.Title, movie.Year, movie.Runtime, pq.Array(movie.Genres), movie.ID}

    return m.DB.QueryRow(query, args...).Scan(&movie.Version)
}
```

Handler:

```go
func (app *application) updateMovieHandler(w http.ResponseWriter, r *http.Request) {
    // 解析 ID
    id, err := app.readIDParam(r)
    if err != nil {
        app.notFoundResponse(w, r)
        return
    }

    // 先得取出来, 才能更新
    movie, err := app.models.Movies.Get(id)
    if err != nil {
        switch {
        case errors.Is(err, data.ErrRecordNotFound):
            app.notFoundResponse(w, r)
        default:
            app.serverErrorResponse(w, r, err)
        }
        return
    }

    // 解析请求
    var input struct {
        Title   string       `json:"title"`
        Year    int32        `json:"year"`
        Runtime data.Runtime `json:"runtime"`
        Genres  []string     `json:"genres"`
    }
    err = app.readJSON(w, r, &input)
    if err != nil {
        app.badRequestResponse(w, r, err)
        return
    }

    // 校验请求
    movie.Title = input.Title
    movie.Year = input.Year
    movie.Runtime = input.Runtime
    movie.Genres = input.Genres

    v := validator.New()
    if data.ValidateMovie(v, movie); !v.Valid() {
        app.failedValidationResponse(w, r, v.Errors)
        return
    }

    // CRUD
    err = app.models.Movies.Update(movie)
    if err != nil {
        app.serverErrorResponse(w, r, err)
        return
    }

    err = app.writeJSON(w, http.StatusOK, envelope{"movie": movie}, nil)
    if err != nil {
        app.serverErrorResponse(w, r, err)
        return
    }

}
```

### 实现 Delete()

Model:

```go
func (m MovieModel) Delete(id int64) error {
    if id < 1 {
        return ErrRecordNotFound
    }

    query := `DELETE FROM movies WHERE id = $1`
    result, err := m.DB.Exec(query, id)
    if err != nil {
        return err
    }

    rowsAffected, err := result.RowsAffected()
    if err != nil {
        return err
    }

    // 未能删除任何记录
    if rowsAffected == 0 {
        return ErrRecordNotFound
    }

    return nil
}
```

Handler:

```go
func (app *application) deleteMovieHandler(w http.ResponseWriter, r *http.Request) {
    id, err := app.readIDParam(r)
    if err != nil {
        app.notFoundResponse(w, r)
        return
    }

    err = app.models.Movies.Delete(id)
    if err != nil {
        switch {
        case errors.Is(err, data.ErrRecordNotFound):
            app.notFoundResponse(w, r)
        default:
            app.serverErrorResponse(w, r, err)
        }
        return
    }

    err = app.writeJSON(w, http.StatusOK, envelope{"message": "movie successfully deleted"}, nil)
    if err != nil {
        app.serverErrorResponse(w, r, err)
        return
    }
}
```

### 部分更新

#### ➤ 如何区分 {"year": 0}, {"year": null}, {}

1. 用指针可以区分 `{year: 0}` 和 `{}`，一般把 `{year: null}` 和 `{}` 视为等价，不刻意区分
2. 注意用指针无法区分 `{year: null}` 和 `{}`，[想区分可以试试 json.RawMessage](https://github.com/golang/go/issues/50360#issuecomment-1262376093)，相当于自定义 JSON Parser

```go
    var input struct {
        Title   *string       `json:"title"`
        Year    *int32        `json:"year"`    // 为了区分 {} 和 {year:0} 改用指针
        Runtime *data.Runtime `json:"runtime"` //
        Genres  []string      `json:"genres"`  // 不用动, 因为能区分 {} 和 {genres:[]}
    }
```

#### ➤ 用全量更新实现部分更新

```go
    // 先把更新目标 movie 查出来, 然后只修改请求中存在的字段, 然后调用全量更新, 就实现了部分更新
    if input.Title != nil {
        movie.Title = *input.Title
    }
    if input.Year != nil {
        movie.Year = *input.Year
    }
    if input.Runtime != nil {
        movie.Runtime = *input.Runtime
    }
    if input.Genres != nil {
        movie.Genres = input.Genres
    }
```

#### ➤ 全量替换用 PUT，部分更新用 PATCH，但后者的能力包含前者

```go
router.HandlerFunc(http.MethodPatch, "/v1/movies/:id", app.updateMovieHandler)
```

### 版本号乐观锁

Go’s `http.Server` handles each HTTP request in its own goroutine. In our `updateMovieHandler` — there is a race condition if two clients  try to update the same movie record at exactly the same time. There are a couple of options, but the simplest and cleanest approach in this case is to use a form of optimistic locking based on the version number in our movie record.

The fix works like this:

1. Alice and Bob’s goroutines both call app.models.Movies.Get() to retrieve a copy of the movie record. Both of these records have the version number N.
2. Alice and Bob’s goroutines make their respective changes to the movie record.
3. Alice and Bob’s goroutines call app.models.Movies.Update() with their copies of the movie record. *But the update is only executed if the version number in the database is still N*. If it has changed, then we don’t execute the update and send the client an error message instead.
4. This means that the first update request that reaches our database will succeed, and whoever is making the second update will receive an error message instead of having their change applied.

We’ll need to change the SQL statement for updating a movie so that it looks like this:

```sql
UPDATE movies SET title = $1, year = $2, version = version + 1
WHERE id = $5 AND version = $6
RETURNING version
```

If no matching record can be found, this query will result in a `sql.ErrNoRows` error and we know that the version number has been changed (or the record has been deleted completely). Either way, it’s a form of edit conflict and we can use this as a trigger to send the client an appropriate error response.  

#### ➤ 在 bash 中执行下述命令测试并发更新

```bash
# <(command) 的作用是把 command 的输出保存到临时文件，然后返回文件描述符
file <(seq 3)         # 适用于 bash
file (seq 3 | psub)   # 适用于 fish

# xargs 可利用文件内容作为参数，然后批量执行命令
echo 'one two three' | xargs mkdir             # 默认使用空白字符作为分隔符，这里会执行 1 次 mkdir 且传入 3 个参数
echo 'one,two,three' | xargs -d ',' -L1 echo   # 使用 , 作为分隔符，每次命令至多使用 1 个参数，这里会执行 3 次 echo
seq 3 | xargs -I % echo id:%                   # -I % 表示用 % 作为参数占位符

# 测试并发更新，-P8 表示同时运行 8 个进程
xargs -I % -P8 curl -X PATCH -d '{"runtime": "97 mins"}' "localhost:4000/v1/movies/4" < <(printf '%s\n' {1..16})
```

#### ➤ 避免客户端基于已过期的数据做更新

One of the nice things about the optimistic locking pattern that we’ve used here is that you can extend it so the client passes the version number that they expect in an `X-Expected-Version` header. In certain applications, this can be useful to help the client ensure they are not sending their update request based on outdated information.

```go
    // If the request contains a X-Expected-Version header, verify that the movie
    // version in the database matches the expected version specified in the header.
    if r.Header.Get("X-Expected-Version") != "" {
        if strconv.FormatInt(int64(movie.Version), 32) != r.Header.Get("X-Expected-Version") {
            app.editConflictResponse(w, r)
            return
        }
    }
```

#### ➤ 代码如下

注意 `updateMovieHandler()` 存在并发问题，先把 moive 查出来再更新，会有 data race，可用乐观锁解决:

```go
var (
    ErrRecordNotFound = errors.New("record not found")
    ErrEditConflict   = errors.New("edit conflict")
)

func (m MovieModel) Update(movie *Movie) error {
    // 注意让 version 加一并返回更新后的 version, 此处 AND version = $6 是乐观锁
    // 如果更新时发现 version 变了, 意味着数据已经被其他线程改过了
    // 那么 AND version = $6 会匹配 0 行记录, 这个查询会返回 sql.ErrNoRows 错误
    query := `
        UPDATE movies 
        SET title = $1, year = $2, runtime = $3, genres = $4, version = version + 1
        WHERE id = $5 AND version = $6
        RETURNING version`

    args := []any{movie.Title, movie.Year, movie.Runtime, pq.Array(movie.Genres), movie.ID, movie.Version}

    err := m.DB.QueryRow(query, args...).Scan(&movie.Version)
    if err != nil {
        switch {
        case errors.Is(err, sql.ErrNoRows):
            // 说明 version 发生改变, 或者对应的记录被删了, 但不管哪种情况都属于「 编辑冲突 」
            return ErrEditConflict
        default:
            return err
        }
    }
    return nil
}
```

Handler:

```go
func (app *application) editConflictResponse(w http.ResponseWriter, r *http.Request) {
    message := "unable to update the record due to an edit conflict, please try again"
    app.errorResponse(w, r, http.StatusConflict, message) // 409 Conflict
}

func (app *application) updateMovieHandler(w http.ResponseWriter, r *http.Request) {
    // 不同错误返回不同响应
    err = app.models.Movies.Update(movie)
    if err != nil {
        switch {
        case errors.Is(err, data.ErrEditConflict):
            app.editConflictResponse(w, r)
        default:
            app.serverErrorResponse(w, r, err)
        }
        return
    }
}
```

测试一下并发更新:

```bash
xargs -I % -P8 curl -X PATCH -d '{"runtime": "106 mins"}' "localhost:4000/v1/movies/1" < <(printf '%s\n' {1..8}) | jq
# 有的成功了，有的出现下述错误
# { "error": "unable to update the record due to an edit conflict, please try again" }
```

### 设置查询超时

This feature can be useful when *you have a SQL query that is taking longer to run than expected*. When this happens, it suggests a problem — either with that particular query or your database or application more generally — and you probably want to cancel the query (in order to free up resources), log an error for further investigation, and return a `500 Internal Server Error` response to the client.

#### ➤ 模拟慢查询

```go
func (m MovieModel) Get(id int64) (*Movie, error) {
    // pg_sleep(10) 会等待 10 秒再返回数据
    query := `
        SELECT pg_sleep(10), id, created_at, title, year, runtime, genres, version
        FROM movies
        WHERE id = $1`

    var movie Movie

    // 因为多 select 了一列, 对应的 Scan 也要改
    err := m.DB.QueryRow(query, id).Scan(
        &[]byte{},
        &movie.ID,
        &movie.CreatedAt,
        &movie.Title,
        &movie.Year,
        &movie.Runtime,
        pq.Array(&movie.Genres), // 查询和解析数组类型, 都要通过 pq.Array 处理
        &movie.Version,
    )
}
```

#### ➤ 设置查询超时

```go
func (m MovieModel) Get(id int64) (*Movie, error) {
    // pg_sleep(10) 会等待 10 秒再返回数据
    query := `
        SELECT pg_sleep(10), id, created_at, title, year, runtime, genres, version
        FROM movies
        WHERE id = $1`

    var movie Movie

    // 设置 3 秒的查询超时
    ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
    defer cancel() // 在函数返回前调用 cancel(), 既能提前释放资源, 也能避免资源泄露

    // 改用 QueryRowContext() 方法
    err := m.DB.QueryRowContext(ctx, query, id).Scan(
        &[]byte{}, // 因为多 select 了一列, 对应的 Scan 也要改
        &movie.ID,
        &movie.CreatedAt,
        &movie.Title,
        &movie.Year,
        &movie.Runtime,
        pq.Array(&movie.Genres), // 查询和解析数组类型, 都要通过 pq.Array 处理
        &movie.Version,
    )

    // 返回自定义错误, 而不是底层数据库的错误, 方便解耦, 切换数据库实现
    if err != nil {
        switch {
        case errors.Is(err, sql.ErrNoRows):
            return nil, ErrRecordNotFound
        default:
            return nil, err
        }
    }

    return &movie, nil
}
```

- The `defer cancel()` line is necessary because it ensures that the resources associated with our context will always be released before the `Get()` method returns, thereby preventing a memory leak. Without it, the resources won’t be released until either the 3-second timeout is hit or the parent context (which in this specific example is [`context.Background()`](https://golang.org/pkg/context/#Background)) is canceled.
- The timeout countdown begins from the moment that the context is created using `context.WithTimeout()`. Any time spent executing code between creating the context and calling `QueryRowContext()` will count towards the timeout.



#### ➤ 原理是有个后台线程会在 3 秒后告诉 DB 要取消某个已经超时的查询

After 3 seconds, the context timeout is reached and our `pq` database driver sends a cancellation signal to PostgreSQL. More precisely, our context (the one with the 3-second timeout) has a `Done` channel, and when the timeout is reached the `Done` channel will be closed. While the SQL query is running, our database driver `pq` is also running a background goroutine which listens on this `Done` channel. If the channel gets closed, then `pq` sends a cancellation signal to PostgreSQL. PostgreSQL terminates the query, and then sends the error message that we see above as a response to the original `pq` goroutine. That error message is then returned to our database model’s `Get()` method.

```bash
time=2024-04-09T12:38:18.639+08:00 level=ERROR msg="pq: canceling statement due to user request" method=GET uri=/v1/movies/1
```

#### ➤ 执行 SQL 查询前可能超时，解析查询结果也可能超时

It’s possible that the timeout deadline will be hit before the PostgreSQL query even starts. In a similar vein, it’s also possible that the timeout deadline will be hit later on when the data returned from the query is being processed with `Scan()`.

- 例如并发请求很多，连接池用光了，没能等到空闲连接就已经超时，那么 `QueryRowContext()` 会返回  `context.DeadlineExceeded`
- 遍历查询结果的每一行并调用 `Scan()`，这个耗时可能很长，在超时后调用 `Scan()`，它也会返回 `context.DeadlineExceeded`

```bash
time=2024-04-09T12:43:33.453+08:00 level=ERROR msg="context deadline exceeded" method=GET uri=/v1/movies/1
```

#### ➤ 记得为所有 CRUD 操作设置 3 秒超时

```go
    // 都添加如下 3 行代码
    ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
    defer cancel()
    return m.DB.QueryRowContext(ctx, ...)
```

#### ➤ 可以把整个请求的 timeout 设为 3 秒，而不是某个 CRUD 操作

```go
func (app *application) exampleHandler(w http.ResponseWriter, r *http.Request) {

    // 整个请求的 timeout 是 3 秒
    ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
    defer cancel()

    // 传入 ctx 而不是在 CRUD 操作中新建 ctx
    example, err := app.models.Example.Get(ctx, id)
    if err != nil {
        switch {
        case errors.Is(err, data.ErrNotFound):
            app.notFoundResponse(w, r)
        default:
            app.serverErrorResponse(w, r, err)
        }
        return
    }
}
```

#### ➤ 要不要使用 r.Context() 作为 Parent

```go
func (app *application) exampleHandler(w http.ResponseWriter, r *http.Request) {
    // 基于 r.Context() 设置超时
    ctx, cancel := context.WithTimeout(r.Context(), 3*time.Second)
    defer cancel()

    // 什么时候 r.Context() 会被取消?
    // the request context is canceled when the client’s connection closes, ( 例如常见的关闭标签页，或关闭浏览器 )
    // the request is canceled (with HTTP/2),
    // or when the ServeHTTP method returns.
}
```

- 好处:
  - 可以访问 r.Context() 中的数据
  - 如果用户取消了请求 ( 例如断开连接 )，那么 server 可以提前返回，避免继续做无意义的计算

- 但也有缺点不要随便滥用: ( 那些开销大，并且容易被取消请求的接口，则比较合适 )
  - 错误处理更麻烦，取消的请求和超时的请求，错误都是 `pq: canceling statement due to user request`，可用 `ctx.Err()` 区分
  - 增加心智负担，例如开个 goroutine 后台跑 SQL，那么切记不要用 `r.Context()`，否则请求结束时，后台任务也会被取消
- 折中的办法:

```go
// Go 1.21 加了 WithoutCancel() 方法，用于忽略来自 parent context 的取消
// 这样既能访问 r.Context() 中的数据，又避免了直接使用 r.Context() 作为 parent 导致的复杂性和错误处理
ctx, cancel := context.WithTimeout(context.WithoutCancel(r.Context()), time.Second)
defer cancel()
```

## CRUD - 列表

### 查询参数

#### ➤ 请求示例

We’re going configure the `GET /v1/movies` endpoint so that a client can control which movie records are returned via *query string parameters*. For example:

```html
/v1/movies?title=godfather&genres=crime,drama&page=1&page_size=5&sort=-year
```

If a client sends a query string like this, it is essentially saying to our API: *“please return the first 5 records where the movie name includes `godfather` and the genres include `crime` and `drama`, sorted by descending release year”*.

#### ➤ 一些注意事项

1. `sort=-year` 表示按年份降序，其中的 `-` 号是降序的意思，不加这个符号则是升序
2. `genres=crime,drama` 中的 "crime,drama" 要从字符串转成 `[]string` 类型
3. `page` 和 `page_size` 要转成 `int` 类型，并且校验是否负数，另外 page_size 有上限
4. 如果用户不提供 `page`, `page_size`, `sort` 等参数，这些参数应该有个合理的默认值

#### ➤ 封装一些解析查询参数的 helper 函数

```go
func (app *application) readString(qs url.Values, key string, defaultValue string) string {
    s := qs.Get(key)
    if s == "" {
        return defaultValue
    }
    return s
}

func (app *application) readCSV(qs url.Values, key string, defaultValue []string) []string {
    csv := qs.Get(key)
    if csv == "" {
        return defaultValue
    }
    return strings.Split(csv, ",")
}

func (app *application) readInt(qs url.Values, key string, defaultValue int, v *validator.Validator) int {
    s := qs.Get(key)
    if s == "" {
        return defaultValue
    }

    // 如果 string 不能转成 int, 那么添加错误消息, 并返回默认值
    i, err := strconv.Atoi(s)
    if err != nil {
        v.AddError(key, "must be an integer value")
        return defaultValue
    }
    return i
}
```

#### ➤ 添加 listMovieHandler

```go
func (app *application) listMovieHandler(w http.ResponseWriter, r *http.Request) {
    var input struct {
        Title    string
        Genres   []string
        Page     int
        PageSize int
        Sort     string
    }

    v := validator.New()
    qs := r.URL.Query()
    input.Title = app.readString(qs, "title", "")
    input.Genres = app.readCSV(qs, "genres", []string{})
    input.Page = app.readInt(qs, "page", 1, v)
    input.PageSize = app.readInt(qs, "page_size", 20, v)
    input.Sort = app.readString(qs, "sort", "id")
    if !v.Valid() {
        app.failedValidationResponse(w, r, v.Errors)
        return
    }

    _, _ = fmt.Fprintf(w, "%+v\n", input)
}
```

#### ➤ 可重用 Filters

The `page`, `page_size` and `sort` query string parameters are things that you’ll potentially want to use on other endpoints in your API too. Let’s quickly split them out into a reusable `Filters` struct.

```go
type Filters struct {
    Page     int
    PageSize int
    Sort     string
}

func (app *application) listMovieHandler(w http.ResponseWriter, r *http.Request) {
    var input struct {
        Title        string
        Genres       []string
        data.Filters // 嵌入 Filters 结构
    }
}
```

#### ➤ 校验 Filters 相关数据

- The `page` value is between 1 and 10,000,000.
- The `page_size` value is between 1 and 100.
- The `sort` parameter contains a known and supported value for our movies table. Specifically, we’ll allow `"id"`, `"title"`, `"year"`, `"runtime"`, `"-id"`, `"-title"`, `"-year"` or `"-runtime"`.

Let’s open up the `internal/data/filters.go` file and create a new `ValidateFilters()` function which conducts these checks on the values. We’ll follow the same pattern that we used for the `ValidateMovie()` function earlier on to do this, like so:

```go
type Filters struct {
    Page         int
    PageSize     int
    Sort         string
    SortSafelist []string
}

func ValidateFilters(v *validator.Validator, f Filters) {
    v.Check(f.Page > 0, "page", "must be greater than zero")
    v.Check(f.Page <= 10_000_000, "page", "must be a maximum of 10 million")
    v.Check(f.PageSize > 0, "page_size", "must be greater than zero")
    v.Check(f.PageSize <= 100, "page_size", "must be a maximum of 100")

    // Check that the sort parameter matches a value in the safelist.
    v.Check(validator.PermittedValue(f.Sort, f.SortSafelist...), "sort", "invalid sort value")
}

func (app *application) listMovieHandler(w http.ResponseWriter, r *http.Request) {

    // 当前接口支持的排序字段
    input.Filters.SortSafelist = []string{"id", "title", "year", "runtime", "-id", "-title", "-year", "-runtime"}

    // 校验 Filters 相关数据
    if data.ValidateFilters(v, input.Filters); !v.Valid() {
        app.failedValidationResponse(w, r, v.Errors)
        return
    }
}
```

### 返回列表

Model:

```go
func (m MovieModel) GetAll(title string, genres []string, filters Filters) ([]*Movie, error) {
    query := `
        SELECT id, created_at, title, year, runtime, genres, version
        FROM movies
        ORDER BY id`

    ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
    defer cancel()

    // QueryRow 只返回第一行, 如果查不到数据会返回错误 sql.ErrNoRows
    // Query    查不到数据时, 不会返回错误, 但调用 rows.Next() 会返回 false
    rows, err := m.DB.QueryContext(ctx, query)
    if err != nil {
        return nil, err
    }
    defer rows.Close() // 现在 rows 可用了, 所以别忘了关闭 rows, 否则资源泄露

    movies := []*Movie{} // 查不到数据时返回空切片, 而不是 nil
    for rows.Next() {
        var m Movie
        err := rows.Scan(&m.ID, &m.CreatedAt, &m.Title, &m.Year, &m.Runtime, pq.Array(&m.Genres), &m.Version)
        if err != nil {
            return nil, err
        }
        movies = append(movies, &m)
    }

    // 上面的 for 循环不一定能遍历所有行, 在取第 1000 行时可能遇到错误, 所以别忘了检查 rows.Err()
    if err = rows.Err(); err != nil {
        return nil, err
    }

    return movies, nil
}
```

Handler:

```go
func (app *application) listMovieHandler(w http.ResponseWriter, r *http.Request) {

    movies, err := app.models.Movies.GetAll(input.Title, input.Genres, input.Filters)
    if err != nil {
        app.serverErrorResponse(w, r, err)
        return
    }

    err = app.writeJSON(w, http.StatusOK, envelope{"movies": movies}, nil)
    if err != nil {
        app.serverErrorResponse(w, r, err)
        return
    }
}
```

### 筛选电影

#### ➤ 请求示例

```bash
GET /v1/movies                                          # 返回所有 moive
GET /v1/movies?title=black+panther                      # 返回标题是 black panther 的电影, 要求完全匹配, 但忽略大小写
GET /v1/movies?genres=adventure                         # 返回 genres 包含 adventure 的电影
GET /v1/movies?title=moana&genres=animation,adventure   # 标题是 moana 且类型包含 animation,adventure
```

#### ➤ 动态拼接 SQL

The hardest part of building a dynamic filtering feature like this is the SQL query to retrieve the data — we need it to work with no filters, filters on both `title` and `genres`, or a filter on only one of them. To deal with this, one option is to build up the SQL query dynamically at runtime… with the necessary SQL for each filter concatenated or interpolated into the `WHERE` clause. But this approach can make your code messy and difficult to understand, especially for large queries which need to support lots of filter options.

#### ➤ 可以用一个固定 SQL

```sql
SELECT id, created_at, title, year, runtime, genres, version FROM movies
# 不提供 title 时, 表达式 $1 = '' 的值是 true, 然后 OR true 会让 title 这个筛选条件失效
WHERE (LOWER(title) = LOWER($1) OR $1 = '') 

# 注意 @> 是 PostgreSQL 中的 contains 操作符, 当一个 array 包含另一个 array 时返回 true
# 这里 '{}' 表示空数组, 然后 $2 的默认值记得设为空切片: input.Genres = app.readCSV(qs, "genres", []string{})
AND (genres @> $2 OR $2 = '{}')
ORDER BY id
```

#### ➤ 代码

```go
func (m MovieModel) GetAll(title string, genres []string, filters Filters) ([]*Movie, error) {
    query := `
        SELECT id, created_at, title, year, runtime, genres, version
        FROM movies
        WHERE (LOWER(title) = LOWER($1) OR $1 = '') 
        AND (genres @> $2 OR $2 = '{}')     
        ORDER BY id`

    ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
    defer cancel()

    rows, err := m.DB.QueryContext(ctx, query, title, pq.Array(genres))
}
```

### 全文搜索

#### ➤ [先了解一下什么是倒排索引](https://zhuanlan.zhihu.com/p/94181307)

#### ➤ [再了解一下 tsvector 是什么](https://www.postgresql.org/docs/current/datatype-textsearch.html)

#### ➤ [关于 PostgreSQL 全文检索的实战](https://www.skypyb.com/2020/12/jishu/1705/)

In this chapter we’re going to make our movie title filter easier to use by adapting it to support *partial matches*, rather than requiring a match on the full title. There are a few different ways we could implement this feature in our codebase, but an effective and intuitive method (from a client point of view) is to leverage PostgreSQL’s *full-text search* functionality, which allows you to perform ‘natural language’ searches on text fields in your database. ( 默认不支持中文分词 )

To implement a basic full-text search on our `title` field, we’re going to update our SQL query to look like this:

```sql
SELECT id, created_at, title, year, runtime, genres, version
FROM movies
WHERE (to_tsvector('simple', title) @@ plainto_tsquery('simple', $1) OR $1 = '') 
AND (genres @> $2 OR $2 = '{}')     
ORDER BY id
```

(1) to_tsvector('simple', title)

The `to_tsvector('simple', title)` function takes a movie title and splits it into *lexemes*. We specify the `simple` configuration, which means that the lexemes are just lowercase versions of the words in the title. For example, the movie title `"The Breakfast Club"` would be split into the lexemes `'breakfast' 'club' 'the'`. Other ‘non-simple’ configurations may apply additional rules to the lexemes, such as the removal of common words or applying language-specific stemming.

(2) plainto_tsquery('simple', $1)

The `plainto_tsquery('simple', $1)` function takes a search value and turns it into a formatted *query term* that PostgreSQL full-text search can understand. It normalizes the search value (again using the `simple` configuration), strips any special characters, and inserts the *and operator* `&` between the words. As an example, the search value `"The Club"` would result in the query term `'the' & 'club'`.

(3) @@

The `@@` operator is the matches operator. In our statement we are using it to check whether the generated *query term matches the lexemes*. To continue the example, the query term `'the' & 'club'` will match rows which contain *both* lexemes `'the'` and `'club'`.

#### ➤ Let’s go ahead and put this into action.

```go
func (m MovieModel) GetAll(title string, genres []string, filters Filters) ([]*Movie, error) {
    query := `
        SELECT id, created_at, title, year, runtime, genres, version
        FROM movies
        WHERE (to_tsvector('simple', title) @@ plainto_tsquery('simple', $1) OR $1 = '') 
        AND (genres @> $2 OR $2 = '{}')     
        ORDER BY id`
    // 剩下的代码不用改
}
```

#### ➤ 添加索引

To keep our SQL query performing quickly as the dataset grows, it’s sensible to use indexes to help avoid full table scans and avoid generating the lexemes for the `title` field every time the query is run. In our case it makes sense to create [GIN indexes](https://www.postgresql.org/docs/current/textsearch-indexes.html) on both the `genres` field and the lexemes generated by `to_tsvector()`, both which are used in the `WHERE` clause of our SQL query.

Create a new pair of migration files:

```postgresql
-- 创建 migration 文件, 然后执行 up
-- migrate create -seq -ext .sql -dir ./migrations add_movies_indexes
-- migrate -path ./migrations -database $GREENLIGHT_DB_DSN up

-- up file
CREATE INDEX IF NOT EXISTS movies_title_idx ON movies USING GIN (to_tsvector('simple', title));
CREATE INDEX IF NOT EXISTS movies_genres_idx ON movies USING GIN (genres);

-- down file
DROP INDEX IF EXISTS movies_title_idx;
DROP INDEX IF EXISTS movies_genres_idx;
```

### 排序

#### ➤ 格式化 ORDER BY 要注意 SQL 注入问题

As we briefly explained earlier, we want to let the client control the sort order via a query string parameter in the format `sort={-}{field_name}`, where the optional `-` character is used to indicate a descending sort order. Behind the scenes we will want to translate this into an [`ORDER BY year DESC`](https://www.postgresql.org/docs/current/queries-order.html) clause in our SQL query. The difficulty here is that the values for the `ORDER BY` clause will need to be generated at runtime based on the query string values from the client. Ideally we’d use placeholder parameters to insert these dynamic values into our query, but unfortunately it’s *not possible to use placeholder parameters for column names or SQL keywords* (including `ASC` and `DESC`).

So instead, we’ll need to interpolate these dynamic values into our query using `fmt.Sprintf()` — making sure that the values are checked against a strict safelist first to prevent a SQL injection attack. 

#### ➤ 必须加 ORDER BY，且根据唯一字段排序，否则每次查都随机顺序返回

When working with PostgreSQL, it’s also important to be aware that the order of returned rows *is only guaranteed by the rules that your `ORDER BY` clause imposes*. Likewise, in our database multiple movies will have the same `year` value. If we order based on the `year` column, then the movies are guaranteed be ordered by year, but the movies *for* a particular year could appear in any order at any time. *If sorting is not chosen, the rows will be returned in an unspecified order. The actual order in that case will depend on the scan and join plan types and the order on disk,*

This point is particularly important in the context of an endpoint which provides pagination. We need to make sure that the order of movies is perfectly consistent between requests to prevent items in the list ‘jumping’ between the pages. Fortunately, guaranteeing the order is simple — we just need to ensure that the `ORDER BY` clause always includes a primary key column (or another column with a unique constraint on it). So, in our case, we can apply a secondary sort on the `id` column to ensure an always-consistent order. Like so:

```sql
ORDER BY year DESC, id ASC
```

#### ➤ 实现排序

Let’s begin by updating our `Filters` struct to include some `sortColumn()` and `sortDirection()` helpers that transform a query string value (like `-year`) into values we can use in our SQL query.

```go
func (f Filters) sortColumn() string {
    // 检查请求中的 Sort 字段是否符合白名单, 去掉 -year 中的减号
    for _, safeValue := range f.SortSafelist {
        if f.Sort == safeValue {
            return strings.TrimPrefix(f.Sort, "-")
        }
    }

    // 正常不会走到这, 因为 ValidateFilters 会校验 f.Sort, 不管怎样这个函数是避免 SQL 注入的保险
    panic("unsafe sort parameter: " + f.Sort)
}

func (f Filters) sortDirection() string {
    if strings.HasPrefix(f.Sort, "-") {
        return "DESC"
    }
    return "ASC"
}
```

Notice that the `sortColumn()` function is constructed in such a way that it will panic if the client-provided `Sort` value doesn’t match one of the entries in our safelist. In theory this shouldn’t happen — the `Sort` value should have already been checked by calling the `ValidateFilters()` function — but this is a sensible failsafe to help stop a SQL injection attack occurring.

```go
func (m MovieModel) GetAll(title string, genres []string, filters Filters) ([]*Movie, error) {
    // 格式化 ORDER BY 子句, 其中 id ASC 能让 year 相同的电影以固定顺序返回
    query := fmt.Sprintf(`
        SELECT id, created_at, title, year, runtime, genres, version
        FROM movies
        WHERE (to_tsvector('simple', title) @@ plainto_tsquery('simple', $1) OR $1 = '') 
        AND (genres @> $2 OR $2 = '{}')     
        ORDER BY %s %s, id ASC`, filters.sortColumn(), filters.sortDirection())
}
```

### 分页

#### ➤ 使用 LIMIT 和 OFFSET 分页

The `LIMIT` clause allows you to set the maximum number of records that a SQL query should return, and `OFFSET` allows you to ‘skip’ a specific number of rows before starting to return records from the query. The math is pretty straightforward:

```sql
LIMIT = page_size
OFFSET = (page - 1) * page_size
```

Let’s start by adding some helper methods to our `Filters` struct for calculating the appropriate `LIMIT` and `OFFSET` values.

```go
func (f Filters) limit() int {
    return f.PageSize
}

func (f Filters) offset() int {
    // 两个 int 相乘有 overflow 的风险, 但校验过 PageSize <= 100, Page <= 1000_0000 所以不会溢出
    return (f.Page - 1) * f.PageSize
}
```

We need to update our database model’s `GetAll()` method:

```go
func (m MovieModel) GetAll(title string, genres []string, filters Filters) ([]*Movie, error) {

    query := fmt.Sprintf(`
        SELECT id, created_at, title, year, runtime, genres, version
        FROM movies
        WHERE (to_tsvector('simple', title) @@ plainto_tsquery('simple', $1) OR $1 = '') 
        AND (genres @> $2 OR $2 = '{}')     
        ORDER BY %s %s, id ASC
        LIMIT $3 OFFSET $4`, filters.sortColumn(), filters.sortDirection())

    args := []any{title, pq.Array(genres), filters.limit(), filters.offset()}
}
```

#### ➤ 返回分页信息 ( 怎么计算 total_records 字段 )

```json
{
    "metadata": {
        "current_page": 1,
        "page_size": 20,
        "first_page": 1,
        "last_page": 42,
        "total_records": 832
    },
    "movies": [...]
}
```

The challenging part of doing this is generating the `total_records` figure. We want this to reflect the total number of available records *given the `title` and `genres` filters that are applied* — not the absolute total of records in the `movies` table. A neat way to do this is to adapt our existing SQL query to include a [window function](https://www.postgresql.org/docs/current/tutorial-window.html) which counts the total number of filtered rows, like so:

```postgresql
SELECT count(*) OVER(), id, created_at, title, year, runtime, genres, version
FROM movies
WHERE (to_tsvector('simple', title) @@ plainto_tsquery('simple', $1) OR $1 = '') 
AND (genres @> $2 OR $2 = '{}')     
ORDER BY %s %s, id ASC
LIMIT $3 OFFSET $4
```

The inclusion of the `count(*) OVER()` expression at the start of the query will result in the filtered record count being included as the first value in each row. A bit like this:

```html
 count | id |     title     | year
-------+----+---------------+------
     4 |  1 | Your Name     | 2016
     4 |  2 | Big Hero 6    | 2014
     4 |  3 | Summer Wars   | 2009
     4 |  4 | Spirited Away | 2001
```

When PostgreSQL executes this SQL query, the (very simplified) sequence of events runs broadly like this:

1. The `WHERE` clause is used to filter the data in the `movies` table and get the *qualifying rows*.
2. The window function `count(*) OVER()` is applied, which counts all the qualifying rows.
3. The `ORDER BY` rules are applied and the qualifying rows are sorted.
4. The `LIMIT` and `OFFSET` rules are applied and the appropriate sub-set of sorted qualifying rows is returned.

#### ➤ 返回分页信息

```go
type Metadata struct {
    CurrentPage  int `json:"current_page,omitempty"`
    PageSize     int `json:"page_size,omitempty"`
    FirstPage    int `json:"first_page,omitempty"`
    LastPage     int `json:"last_page,omitempty"`
    TotalRecords int `json:"total_records,omitempty"`
}

func calculateMetadata(totalRecords, page, pageSize int) Metadata {
    if totalRecords == 0 {
        return Metadata{}
    }
    return Metadata{
        CurrentPage:  page,
        PageSize:     pageSize,
        FirstPage:    1,
        LastPage:     int(math.Ceil(float64(totalRecords) / float64(pageSize))), // 向上取整
        TotalRecords: totalRecords,
    }
}

func (m MovieModel) GetAll(title string, genres []string, filters Filters) ([]*Movie, Metadata, error) {
    query := fmt.Sprintf(`
        SELECT count(*) OVER(), id, created_at, title, year, runtime, genres, version
        FROM movies
        WHERE (to_tsvector('simple', title) @@ plainto_tsquery('simple', $1) OR $1 = '') 
        AND (genres @> $2 OR $2 = '{}')     
        ORDER BY %s %s, id ASC
        LIMIT $3 OFFSET $4`, filters.sortColumn(), filters.sortDirection())

    args := []any{title, pq.Array(genres), filters.limit(), filters.offset()}

    ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
    defer cancel()

    rows, err := m.DB.QueryContext(ctx, query, args...)
    if err != nil {
        return nil, Metadata{}, err
    }
    defer rows.Close()

    totalRecords := 0
    movies := []*Movie{}
    for rows.Next() {
        var m Movie
        err := rows.Scan(&totalRecords, &m.ID, &m.CreatedAt, &m.Title, &m.Year, &m.Runtime, pq.Array(&m.Genres), &m.Version)
        if err != nil {
            return nil, Metadata{}, err
        }
        movies = append(movies, &m)
    }

    if err = rows.Err(); err != nil {
        return nil, Metadata{}, err
    }

    metadata := calculateMetadata(totalRecords, filters.Page, filters.PageSize)
    return movies, metadata, nil
}
```

