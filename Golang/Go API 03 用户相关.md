## Table of Contents
  - [注册用户](#%E6%B3%A8%E5%86%8C%E7%94%A8%E6%88%B7)
    - [创建表](#%E5%88%9B%E5%BB%BA%E8%A1%A8)
    - [CRUD 😀](#CRUD-)
    - [注册用户](#%E6%B3%A8%E5%86%8C%E7%94%A8%E6%88%B7)
  - [发邮件](#%E5%8F%91%E9%82%AE%E4%BB%B6)
    - [SMTP 服务器](#SMTP-%E6%9C%8D%E5%8A%A1%E5%99%A8)
    - [Email 模板](#Email-%E6%A8%A1%E6%9D%BF)
    - [发送 Email](#%E5%8F%91%E9%80%81-Email)
    - [在后台发邮件](#%E5%9C%A8%E5%90%8E%E5%8F%B0%E5%8F%91%E9%82%AE%E4%BB%B6)
    - [优雅关闭后台任务](#%E4%BC%98%E9%9B%85%E5%85%B3%E9%97%AD%E5%90%8E%E5%8F%B0%E4%BB%BB%E5%8A%A1)
  - [激活用户](#%E6%BF%80%E6%B4%BB%E7%94%A8%E6%88%B7)
    - [激活流程](#%E6%BF%80%E6%B4%BB%E6%B5%81%E7%A8%8B)
    - [创建表](#%E5%88%9B%E5%BB%BA%E8%A1%A8)
    - [生成 token](#%E7%94%9F%E6%88%90-token)
    - [CRUD](#CRUD)
    - [发送激活邮件](#%E5%8F%91%E9%80%81%E6%BF%80%E6%B4%BB%E9%82%AE%E4%BB%B6)
    - [重发激活邮件](#%E9%87%8D%E5%8F%91%E6%BF%80%E6%B4%BB%E9%82%AE%E4%BB%B6)
    - [激活用户](#%E6%BF%80%E6%B4%BB%E7%94%A8%E6%88%B7)
  - [各种认证方式](#%E5%90%84%E7%A7%8D%E8%AE%A4%E8%AF%81%E6%96%B9%E5%BC%8F)
    - [HTTP Basic Authentication](#HTTP-Basic-Authentication)
    - [Token Authentication](#Token-Authentication)
    - [API-key Authentication](#APIkey-Authentication)
    - [OAuth 2.0 / OpenID Connect](#OAuth-20--OpenID-Connect)
    - [Which one should I use?](#Which-one-should-I-use)
  - [用户认证](#%E7%94%A8%E6%88%B7%E8%AE%A4%E8%AF%81)
    - [生成 Authentication token](#%E7%94%9F%E6%88%90-Authentication-token)
    - [用 Middleware 认证请求者是谁](#%E7%94%A8-Middleware-%E8%AE%A4%E8%AF%81%E8%AF%B7%E6%B1%82%E8%80%85%E6%98%AF%E8%B0%81)
  - [权限验证](#%E6%9D%83%E9%99%90%E9%AA%8C%E8%AF%81)
    - [只允许已激活用户访问](#%E5%8F%AA%E5%85%81%E8%AE%B8%E5%B7%B2%E6%BF%80%E6%B4%BB%E7%94%A8%E6%88%B7%E8%AE%BF%E9%97%AE)
    - [需要更细粒度的权限控制](#%E9%9C%80%E8%A6%81%E6%9B%B4%E7%BB%86%E7%B2%92%E5%BA%A6%E7%9A%84%E6%9D%83%E9%99%90%E6%8E%A7%E5%88%B6)
    - [用户和权限是多对多关系](#%E7%94%A8%E6%88%B7%E5%92%8C%E6%9D%83%E9%99%90%E6%98%AF%E5%A4%9A%E5%AF%B9%E5%A4%9A%E5%85%B3%E7%B3%BB)
    - [创建表](#%E5%88%9B%E5%BB%BA%E8%A1%A8)
    - [CRUD](#CRUD)
    - [检查用户权限](#%E6%A3%80%E6%9F%A5%E7%94%A8%E6%88%B7%E6%9D%83%E9%99%90)
    - [添加用户权限](#%E6%B7%BB%E5%8A%A0%E7%94%A8%E6%88%B7%E6%9D%83%E9%99%90)

## 注册用户

### 创建表

创建 SQL migration 文件:

```bash
migrate create -seq -ext=.sql -dir=./migrations create_users_table
```

填入如下代码:

```postgresql
-- 000004_create_users_table.up.sql
-- timestamp(p) 中的 p 表示精度, 取值可以从 0 到 6 表示精确到秒和微秒
CREATE TABLE IF NOT EXISTS users
(
    id            bigserial PRIMARY KEY,
    created_at    timestamp(0) with time zone NOT NULL DEFAULT now(),
    name          text                        NOT NULL,
    email         citext UNIQUE               NOT NULL,          -- 比较邮箱时忽略大小写, 邮箱不允许重复
    password_hash bytea                       NOT NULL,          -- 用 byte array 存储加密后的密码
    activated     bool                        NOT NULL,          -- 需要通过邮箱激活用户
    version       integer                     NOT NULL DEFAULT 1 -- 乐观锁处理并发更新
);

-- 000004_create_users_table.down.sql
DROP TABLE IF EXISTS users;
```

最后执行如下命令:

```bash
set DSN 'postgres://greenlight:pa55word@localhost/greenlight?sslmode=disable'
migrate -path=./migrations -database=$DSN up
```

### CRUD 😀

#### ➤ 创建文件 & 安装依赖:

```bash
touch internal/data/users.go
go get golang.org/x/crypto/bcrypt
```

#### ➤ 添加 User 结构:

```go
type User struct {
    ID        int64     `json:"id"`
    CreatedAt time.Time `json:"created_at"`
    Name      string    `json:"name"`
    Email     string    `json:"email"`
    Password  password  `json:"-"`
    Activated bool      `json:"activated"`
    Version   int       `json:"-"`
}

type password struct {
    plaintext *string // 用 *string 指针区分「没有」和「空字符串」
    hash      []byte
}

func (p *password) Set(plaintextPassword string) error {
    // 有的语言默认 cost 是 12, 但咱用目前的默认值 10 就好, 这能减少服务器和黑客的计算压力
    hash, err := bcrypt.GenerateFromPassword([]byte(plaintextPassword), bcrypt.DefaultCost)
    if err != nil {
        return err
    }
    p.plaintext = &plaintextPassword
    p.hash = hash
    return nil
}

func (p *password) Matches(plaintextPassword string) (bool, error) {
    // 用 bcrypt hash 中记录的 salt 和 cost 对明文密码再加密一次, 如果是相同的结果, 证明密码匹配
    err := bcrypt.CompareHashAndPassword(p.hash, []byte(plaintextPassword))
    if err != nil {
        switch {
        case errors.Is(err, bcrypt.ErrMismatchedHashAndPassword):
            return false, nil
        default:
            return false, err
        }
    }

    return true, nil
}
```

#### ➤ 添加校验函数:

```go
func ValidateEmail(v *validator.Validator, email string) {
    v.Check(email != "", "email", "must be provided")
    v.Check(validator.Matches(email, validator.EmailRX), "email", "must be a valid email address")
}

func ValidatePassword(v *validator.Validator, password string) {
    v.Check(password != "", "password", "must be provided")
    v.Check(len(password) >= 8, "password", "must be at least 8 bytes long")
    v.Check(len(password) <= 72, "password", "must not be more than 72 bytes long") // bcrypt 输入上限 72 bytes
}

func ValidateUser(v *validator.Validator, user *User) {
    v.Check(user.Name != "", "name", "must be provided")
    v.Check(len(user.Name) <= 500, "name", "must not be more than 500 bytes long")
    ValidateEmail(v, user.Email)
    if user.Password.plaintext != nil {
        ValidatePassword(v, *user.Password.plaintext)
    }
    // 完整性检查: 如果 hash 为空意味着代码中存在逻辑错误, 例如忘记设置 hash
    if user.Password.hash == nil {
        panic("missing password hash for user")
    }
}
```

#### ➤ 添加 UserModel

```go
var (
    ErrDuplicateEmail = errors.New("duplicate email")
)

type UserModel struct {
    DB *sql.DB
}

func (m UserModel) Insert(user *User) error {
    // 使用 RETURNING 返回数据库生成的 id, created_at, version 再填充回 User 结构
    query := `
        INSERT INTO users (name, email, password_hash, activated)
        VALUES ($1, $2, $3, $4)
        RETURNING id, created_at, version`

    args := []any{user.Name, user.Email, user.Password.hash, user.Activated}

    ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
    defer cancel()

    // 如果插入会导致 email 重复, 数据库会返回如下错误消息 ( 感觉这样比较字符串有点丑陋啊 )
    err := m.DB.QueryRowContext(ctx, query, args...).Scan(&user.ID, &user.CreatedAt, &user.Version)
    if err != nil {
        switch {
        case err.Error() == `pq: duplicate key value violates unique constraint "users_email_key"`:
            return ErrDuplicateEmail
        default:
            return err
        }
    }
    return nil
}

func (m UserModel) GetByEmail(email string) (*User, error) {
    query := `
        SELECT id, created_at, name, email, password_hash, activated, version
        FROM users
        WHERE email = $1`

    var user User

    ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
    defer cancel()

    err := m.DB.QueryRowContext(ctx, query, email).Scan(
        &user.ID,
        &user.CreatedAt,
        &user.Name,
        &user.Email,
        &user.Password.hash,
        &user.Activated,
        &user.Version,
    )
    if err != nil {
        switch {
        case errors.Is(err, sql.ErrNoRows):
            return nil, ErrRecordNotFound
        default:
            return nil, err
        }
    }

    return &user, nil
}

func (m UserModel) Update(user *User) error {
    // 使用 version 乐观锁处理并发更新, 返回新的 version
    query := `
        UPDATE users
        SET name = $1, email = $2, password_hash = $3, activated = $4, version = version + 1
        WHERE id = $5 AND version = $6
        RETURNING version`

    args := []any{
        user.Name,
        user.Email,
        user.Password.hash,
        user.Activated,
        user.ID,
        user.Version,
    }

    ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
    defer cancel()

    // 更新可能导致 email 重复所以要检查这种情况
    err := m.DB.QueryRowContext(ctx, query, args...).Scan(&user.Version)
    if err != nil {
        switch {
        case err.Error() == `pq: duplicate key value violates unique constraint "users_email_key"`:
            return ErrDuplicateEmail
        case errors.Is(err, sql.ErrNoRows):
            return ErrEditConflict
        default:
            return err
        }
    }
    return nil
}
```

### 注册用户

#### ➤ 请求示例

```http
POST localhost:4000/v1/users
Content-Type: application/json

{
  "name": "焰",
  "email": "homura@xb2.com",
  "password": "my waifu"
}
```

#### ➤ Handler

```go
func (app *application) registerUserHandler(w http.ResponseWriter, r *http.Request) {
    // 解析请求
    var input struct {
        Name     string `json:"name"`
        Email    string `json:"email"`
        Password string `json:"password"`
    }

    err := app.readJSON(w, r, &input)
    if err != nil {
        app.badRequestResponse(w, r, err)
        return
    }

    // 拷贝数据
    user := &data.User{
        Name:      input.Name,
        Email:     input.Email,
        Activated: false,
    }

    err = user.Password.Set(input.Password)
    if err != nil {
        app.serverErrorResponse(w, r, err)
        return
    }

    // 校验数据
    v := validator.New()
    if data.ValidateUser(v, user); !v.Valid() {
        app.failedValidationResponse(w, r, v.Errors)
        return
    }

    // CRUD
    err = app.models.Users.Insert(user)
    if err != nil {
        switch {
        case errors.Is(err, data.ErrDuplicateEmail):
            v.AddError("email", "a user with this email already exists")
            app.failedValidationResponse(w, r, v.Errors)
        default:
            app.serverErrorResponse(w, r, err)
        }
        return
    }

    // 返回 201 Created 和创建的用户
    err = app.writeJSON(w, http.StatusCreated, envelope{"user": user}, nil)
    if err != nil {
        app.serverErrorResponse(w, r, err)
        return
    }
}
```

## 发邮件

### SMTP 服务器

#### ➤ 先看看 [What is SMTP](https://www.youtube.com/watch?v=PJo5yOtu7o8) ( 你需要 SMTP Server 即邮差，帮你把邮件送到目标的邮箱 )

In order to develop our email sending functionality, we’ll need access to a SMTP (Simple Mail Transfer Protocol) server that we can safely use for testing purposes. But in this book we’re going to use [Mailtrap](https://mailtrap.io/). The reason for using Mailtrap is because it’s a specialist service for *sending emails during development and testing*. Each inbox has its own set of SMTP credentials, which you can display by clicking the **Show Credentials** link.

```js
// 配置邮件客户端的 SMTP Server 时会用到如下信息:
var transport = nodemailer.createTransport({
  host: "sandbox.smtp.mailtrap.io",
  port: 2525,
  auth: {
    user: "...",
    pass: "..."
  }
});
```

### Email 模板

创建文件:

```bash
mkdir -p internal/mailer/templates
touch internal/mailer/templates/user_welcome.tmpl
```

模板文件:

```html
{{define "htmlBody"}}
    <!doctype html>
    <html lang="en">
    <head>
        <meta charset="utf-8">
        <title>blah</title>
    </head>
    <body>
    <div>
        <p>欢迎注册 xxx，你的用户 ID 是 {{.ID}}</p>
    </div>
    </body>
    </html>
{{end}}

{{define "plainBody"}}
    欢迎注册 xxx，你的用户 ID 是 {{.ID}}
{{end}}

{{define "subject"}}Welcome to xxx!{{end}}
```

### 发送 Email

#### ➤ 标准库有 net/smtp 但三方库 [go-mail/mail](https://github.com/go-mail/mail) 更好用:

```bash
go get gopkg.in/mail.v2
touch internal/mailer/mailer.go
```

#### ➤ 封装 Mailer:

```go
import (
    "github.com/go-mail/mail/v2"
    "html/template"
)

//go:embed "templates"
var templateFS embed.FS // 使用特殊注释嵌入 templates 文件夹, 能方便部署

type Mailer struct {
    dialer *mail.Dialer // 用来发邮件
    sender string       // 表示发件人
}

func New(host string, port int, username, password, sender string) Mailer {
    dialer := mail.NewDialer(host, port, username, password)
    dialer.Timeout = 5 * time.Second
    dialer.StartTLSPolicy = mail.MandatoryStartTLS
    return Mailer{
        dialer: dialer,
        sender: sender,
    }
}

func (m Mailer) Send(recipient, templateFile string, data any) error {
    // 从嵌入式文件系统读取模板
    tmpl, err := template.New("email").ParseFS(templateFS, "templates/"+templateFile)
    if err != nil {
        return err
    }

    // 渲染 subject, plainBody, htmlBody 模板到各自的 buffer
    subject := new(bytes.Buffer)
    err = tmpl.ExecuteTemplate(subject, "subject", data)
    if err != nil {
        return err
    }

    plainBody := new(bytes.Buffer)
    err = tmpl.ExecuteTemplate(plainBody, "plainBody", data)
    if err != nil {
        return err
    }

    htmlBody := new(bytes.Buffer)
    err = tmpl.ExecuteTemplate(htmlBody, "htmlBody", data)
    if err != nil {
        return err
    }

    // 初始化邮件
    msg := mail.NewMessage()
    msg.SetHeader("To", recipient)
    msg.SetHeader("From", m.sender)
    msg.SetHeader("Subject", subject.String())
    msg.SetBody("text/plain", plainBody.String())
    msg.AddAlternative("text/html", htmlBody.String()) // 注意这行只能在 SetBody() 的后面执行

    // 发送邮件 ( 如果超时会返回 "dial tcp: i/o timeout" 错误
    err = m.dialer.DialAndSend(msg)
    if err != nil {
        return err
    }

    return nil
}
```

#### ➤ 从命令行读取 SMTP Server 配置

```go
type config struct {

    smtp struct {
        host     string
        port     int
        username string
        password string
        sender   string
    }
}

type application struct {

    mailer mailer.Mailer
}

func main() {

    // SMTP Server 配置
    flag.StringVar(&cfg.smtp.host, "smtp-host", "sandbox.smtp.mailtrap.io", "SMTP host")
    flag.IntVar(&cfg.smtp.port, "smtp-port", 587, "SMTP port")
    flag.StringVar(&cfg.smtp.username, "smtp-username", "...", "SMTP username")
    flag.StringVar(&cfg.smtp.password, "smtp-password", "...", "SMTP password")
    flag.StringVar(&cfg.smtp.sender, "smtp-sender", "Greenlight <no-reply@xxx.com>", "SMTP sender")

    app := &application{

        mailer: mailer.New(cfg.smtp.host, cfg.smtp.port, cfg.smtp.username, cfg.smtp.password, cfg.smtp.sender),
    }
}

```

#### ➤ 注册成功后发邮件

```go
func (app *application) registerUserHandler(w http.ResponseWriter, r *http.Request) {

    // CRUD

    // 注册成功后发邮件
    err = app.mailer.Send(user.Email, "user_welcome.tmpl", user)
    if err != nil {
        app.serverErrorResponse(w, r, err)
        return
    }

    // 返回 201 Created 和创建的用户
}
```

### 在后台发邮件

Sending the welcome email from the `registerUserHandler` method adds quite a lot of latency ( 要 5 秒甚至更多 ) to the total request/response round-trip for the client. One way we could reduce this latency is by sending the email in a *background goroutine*. This would effectively ‘decouple’ the task of sending an email from the rest of the code in our `registerUseHandler`, and means that we could return a HTTP response to the client without waiting for the email sending to complete.

```go
func (app *application) registerUserHandler(w http.ResponseWriter, r *http.Request) {
    // 注册成功后发邮件 ( 在后台执行 )
    app.background(func() {
        err := app.mailer.Send(user.Email, "user_welcome.tmpl", user)
        if err != nil {
            app.logger.Error(err.Error())
        }
    })
}

func (app *application) background(fn func()) {
    // 启动后台 goroutine
    go func() {
        // 避免 panic 把服务搞挂
        defer func() {
            if err := recover(); err != nil {
                app.logger.Error(fmt.Sprintf("%v", err))
            }
        }()

        fn()
    }()
}
```

### 优雅关闭后台任务

Sending our welcome email in the background is working well, but there’s still an issue we need to address. When we initiate a graceful shutdown of our application, it *won’t wait for any background goroutines that we’ve launched to complete*. So — if we happen to shutdown our server at an unlucky moment — it’s possible that a new client will be created on our system but they will never be sent their welcome email.

```go
type application struct {

    wg     sync.WaitGroup
}

func (app *application) background(fn func()) {
    // 启动后台 goroutine 之前把 wg + 1
    app.wg.Add(1)
    go func() {
        // 确保 wg - 1, 这行 defer 放最上面所以最后执行, 能代表 goroutine 的结束
        defer app.wg.Done()

        // 避免 panic 把服务搞挂
        defer func() {
            if err := recover(); err != nil {
                app.logger.Error(fmt.Sprintf("%v", err))
            }
        }()

        fn()
    }()
}

        // 如果一切正常、能优雅结束, 那么 Shutdown 方法会返回 nil
        err := srv.Shutdown(ctx)
        if err != nil {
            shutdownError <- err
        }

        // 等待后台任务完成
        app.logger.Info("completing background tasks", "addr", srv.Addr)
        app.wg.Wait()
        shutdownError <- nil

```

## 激活用户

### 激活流程

1、创建用户时生成一个随机 token

2、在数据库存储 token 的 hash 和过期时间

3、把 token 发到用户的邮箱

4、用户提交 token 到 `PUT /v1/users/activated` 接口

5、服务器检查提交的 token 是否存在，是否过期，都正常则把用户的 activated 字段设为 true

6、删除数据库中的 token

### 创建表

Migration:

```bash
migrate create -seq -ext .sql -dir ./migrations create_tokens_table
migrate -path=./migrations -database=$DSN up
```

SQL:

```postgresql
-- 000005_create_tokens_table.up.sql
CREATE TABLE If NOT EXISTS tokens
(
    hash    bytea PRIMARY KEY,
    user_id bigint                      NOT NULL REFERENCES users ON DELETE CASCADE,
    expiry  timestamp(0) with time zone NOT NULL,
    scope   text                        NOT NULL
);

-- 000005_create_tokens_table.down.sql
DROP TABLE IF EXISTS tokens;
```

注意事项:

- 不直接存 token 而是存 token 的 hash ( [有这么重要吗?](https://security.stackexchange.com/questions/99602/activate-user-account-through-email) )  
  假设 token 明文存储且泄露，那么黑客可以激活用户，绕过邮箱所有权检查，伪造 god@heaven.com 的账户
- 使用外键，在删除  user 时级联删除所有 token (似乎触犯了什么禁忌，哎呦不要那么迷信，具体场景具体分析)
- `scope` 字段表示 token 的用途，后面还会创建 authentication token 也存储在这个表

### 生成 token

If the token is easy to guess or can be brute-forced, then it would be possible for an attacker to activate a user’s account even if they don’t have access to the user’s email inbox. Because of this, we want the token to be generated by a *cryptographically secure random number generator* (CSPRNG) and have enough entropy (or *randomness*) that it is impossible to guess. In our case, we’ll create our activation tokens using Go’s [`crypto/rand`](https://golang.org/pkg/crypto/rand/) package and 128-bits (16 bytes) of entropy.

```bash
touch internal/data/tokens.go
```

Code:

```go
const (
    ScopeActivation = "activation"
)

type Token struct {
    Plaintext string
    Hash      []byte
    UserID    int64
    Expiry    time.Time
    Scope     string
}

func generateToken(userID int64, ttl time.Duration, scope string) (*Token, error) {
    token := &Token{
        UserID: userID,
        Expiry: time.Now().Add(ttl),
        Scope:  scope,
    }

    // 生成一串随机字节填满 randomBytes
    // 一般都用 crypto/rand, 如果想要固定的随机序列或非常性能敏感那么可以用 math/rand
    randomBytes := make([]byte, 16)
    _, err := rand.Read(randomBytes)
    if err != nil {
        return nil, err
    }

    // 使用 base32 进行编码, 毕竟是一串随机字节, 直接当成 UTF-8 字符串使用估计会乱码
    // base64 编码效率高, base32 更易于人类阅读, 此处用不到 padding 所以去掉末尾 = 号
    token.Plaintext = base32.StdEncoding.WithPadding(base32.NoPadding).EncodeToString(randomBytes)

    // 使用 SHA-256 计算 hash ( SHA 全称 Secure Hash Algorithm
    // 另外返回值是长度为 32 的 array 而不是 slice 所以转一下类型
    hash := sha256.Sum256([]byte(token.Plaintext))
    token.Hash = hash[:]

    return token, nil
}
```

### CRUD

```go
func ValidateTokenPlaintext(v *validator.Validator, tokenPlaintext string) {
    v.Check(tokenPlaintext != "", "token", "must be provided")
    v.Check(len(tokenPlaintext) == 26, "token", "must be 26 bytes long") // 向上取整 16*8/5 = 25.6
}

type TokenModel struct {
    DB *sql.DB
}

func (m TokenModel) New(userID int64, ttl time.Duration, scope string) (*Token, error) {
    token, err := generateToken(userID, ttl, scope)
    if err != nil {
        return nil, err
    }

    err = m.Insert(token)
    return token, err
}

func (m TokenModel) Insert(token *Token) error {
    query := `
        INSERT INTO tokens (hash, user_id, expiry, scope) 
        VALUES ($1, $2, $3, $4)`

    args := []any{token.Hash, token.UserID, token.Expiry, token.Scope}

    ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
    defer cancel()

    _, err := m.DB.ExecContext(ctx, query, args...)
    return err
}

func (m TokenModel) DeleteAllForUser(scope string, userID int64) error {
    query := `
        DELETE FROM tokens
        WHERE scope = $1 AND user_id = $2`

    ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
    defer cancel()

    _, err := m.DB.ExecContext(ctx, query, scope, userID)
    return err
}
```

### 发送激活邮件

```go
Please send a request to the `PUT /v1/users/activated` endpoint with the following JSON
body to activate your account:

{"token": "Y3QMGX3PJ3WLRL2YRTQGQ6KRHU"}
```

The most important thing about this email is that we’re instructing the user to activate by issuing a `PUT` request to our API — not by *clicking a link* which contains the token as part of the URL path or query string. Having a user click a link to activate via a `GET` request (which is used by default when clicking a link) would certainly be more convenient, but in the case of our API it has some big drawbacks. In particular:

- It would violate the HTTP principle that the `GET` method should only be used for ‘safe’ requests which retrieve resources — not for requests that modify something (like a user’s activation status).
- It’s possible that the user’s web browser or antivirus will pre-fetch the link URL in the background, inadvertently activating the account. [This Stack Overflow comment](https://security.stackexchange.com/a/197005) explains the risk of this nicely.
- 可以提供链接让用户跳转到网站，然后用户复制 token 点击激活 ( 若在 URL 中记录 token 要设置 [Referrer-Policy: Origin](https://medium.com/@shahjerry33/password-reset-token-leak-via-referrer-2e622500c2c1)

```go
    // CRUD
    err = app.models.Users.Insert(user)
    if err != nil {
    }

    token, err := app.models.Tokens.New(user.ID, 3*24*time.Hour, data.ScopeActivation)
    if err != nil {
        app.serverErrorResponse(w, r, err)
        return
    }

    // 注册成功后发邮件 ( 在后台执行 )
    app.background(func() {
        m := map[string]any{
            "activationToken": token.Plaintext,
            "userID":          user.ID,
        }
        err := app.mailer.Send(user.Email, "user_welcome.tmpl", m)
        if err != nil {
            app.logger.Error(err.Error())
        }
    })

```

### 重发激活邮件

You may also want to provide a standalone endpoint for generating and sending activation tokens to your users. This can be useful if you need to re-send an activation token, such as when a user doesn’t activate their account within the 3-day time limit, or they never receive their welcome email.

The code to implement this endpoint is a mix of patterns that we’ve talked about already, so rather than repeating them in the main flow of the book the instructions are included in [this appendix](http://localhost:8000/21.02-creating-additional-activation-tokens.html).

### 激活用户

#### ➤ 一对多关系的两个常用查询，放在各自的 Model 中

One user may have many tokens, but a token can only belong to one user. When you have a one-to-many relationship like this, you’ll potentially want to execute queries against the relationship from two different sides.

```html
UserModel.GetForToken(token)   → Retrieve the user associated with a token
TokenModel.GetAllForUser(user) → Retrieve all tokens associated with a user
```

The nice thing about this approach is that the entities being returned align with the main responsibility of the models: the `UserModel` method is returning a user, and the `TokenModel` method is returning tokens.

#### ➤ 激活流程

```go
func (app *application) activateUserHandler(w http.ResponseWriter, r *http.Request) {
    // 解析请求
    var input struct {
        TokenPlaintext string `json:"token"`
    }

    err := app.readJSON(w, r, &input)
    if err != nil {
        app.badRequestResponse(w, r, err)
        return
    }

    // 校验请求
    v := validator.New()
    if data.ValidateTokenPlaintext(v, input.TokenPlaintext); !v.Valid() {
        app.failedValidationResponse(w, r, v.Errors)
        return
    }

    // 获取 token 关联的用户
    user, err := app.models.Users.GetForToken(data.ScopeActivation, input.TokenPlaintext)
    if err != nil {
        switch {
        case errors.Is(err, data.ErrRecordNotFound):
            v.AddError("token", "invalid or expired activation token")
            app.failedValidationResponse(w, r, v.Errors)
        default:
            app.serverErrorResponse(w, r, err)
        }
        return
    }

    // 激活用户
    user.Activated = true
    err = app.models.Users.Update(user)
    if err != nil {
        switch {
        case errors.Is(err, data.ErrEditConflict):
            app.editConflictResponse(w, r)
        default:
            app.serverErrorResponse(w, r, err)
        }
        return
    }
    
    // 删除这个用户所有的 activation token
    err = app.models.Tokens.DeleteAllForUser(data.ScopeActivation, user.ID)
    if err != nil {
        app.serverErrorResponse(w, r, err)
        return
    }
    
    // 返回 user details
    err = app.writeJSON(w, http.StatusOK, envelope{"user": user}, nil)
    if err != nil {
        app.serverErrorResponse(w, r, err)
        return
    }
}
```

#### ➤ 实现 UserModel.GetForToken()

```go
func (m UserModel) GetForToken(tokenScope, tokenPlaintext string) (*User, error) {
    // 即便是一对多关系, 因为 token hash 是主键且为 WHERE 条件, 所以只能找到 0/1 个记录
    query := `
        SELECT users.id, users.created_at, users.name, users.email, users.password_hash, users.activated, users.version
        FROM users INNER JOIN tokens ON users.id = tokens.user_id
        WHERE tokens.hash = $1
        AND tokens.scope = $2
        AND tokens.expiry > $3`

    tokenHash := sha256.Sum256([]byte(tokenPlaintext))
    args := []any{tokenHash[:], tokenScope, time.Now()} // 因为 pq driver 不支持 array 所以用 [:] 转成切片

    ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
    defer cancel()

    var user User
    err := m.DB.QueryRowContext(ctx, query, args...).Scan(
        &user.ID,
        &user.CreatedAt,
        &user.Name,
        &user.Email,
        &user.Password.hash,
        &user.Activated,
        &user.Version,
    )
    if err != nil {
        switch {
        case errors.Is(err, sql.ErrNoRows):
            return nil, ErrRecordNotFound
        default:
            return nil, err
        }
    }

    return &user, nil
}
```

#### ➤ 为啥用 PUT 而不是 POST

```go
router.HandlerFunc(http.MethodPut, "/v1/users/activated", app.activateUserHandler)
```

If a client sends the same `PUT /v1/users/activated` request multiple times, the first will succeed (assuming the token is valid) and then any subsequent requests will result in an error being sent to the client. But the important thing is that *nothing in our application state (i.e. database) changes after that first request*. Basically, there are no application state side-effects from the client sending the same request multiple times, which means that the endpoint is idempotent and using `PUT` is more appropriate than `POST`.

## 各种认证方式

( 下述所有方式都是在使用 HTTPS 的前提下

### HTTP Basic Authentication

Perhaps the simplest way to determine *who is making a request* to your API is to use HTTP basic authentication. With this method, the client includes an `Authorization` header with every request containing their credentials. The credentials need to be in the format `username:password` and base-64 encoded. So, for example, to authenticate as `alice@example.com:pa55word` the client would send the following header:

```http
Authorization: Basic YWxpY2VAZXhhbXBsZS5jb206cGE1NXdvcmQ=
```

In your API, you can then extract the credentials from this header using Go’s [`Request.BasicAuth()`](https://golang.org/pkg/net/http/#Request.BasicAuth) method, and verify that they’re correct before continuing to process the request. A big plus of HTTP basic authentication is how simple it is for clients. They can just send the same header with every request — and HTTP basic authentication is supported out-of-the-box by most programming languages, web browsers, and tools such as `curl` and `wget`.

For APIs with ‘real’ user accounts and — in particular — hashed passwords, it’s not such a great fit. Comparing the password provided by a client against a (slow) hashed password is a deliberately costly operation, and when using HTTP basic authentication you need to do that check for every request. That will create a lot of extra work for your API server and add significant latency to responses.

It’s often useful in the scenario where your API doesn’t have ‘real’ user accounts, but you want a quick and easy way to restrict access to it or protect it from prying eyes. Basic authentication can still be a good choice if traffic to your API is very low and response speed is not important to you.

### Token Authentication

The high-level idea behind *token authentication* (also sometimes known as *bearer token authentication*) works like this:

1. The client sends a request to your API containing their credentials (typically username or email address, and password).
2. The API verifies that the credentials are correct, generates a *bearer token* which represents the user, and sends it back to the user. The token expires after a set period of time, after which the user will need to resubmit their credentials again to get a new token.
3. For subsequent requests to the API, the client includes the token in an `Authorization` header like this:

```http
Authorization: Bearer <token>
```

4. When your API receives this request, it checks that the token hasn’t expired and examines the token value to determine who the user is.

For APIs where user passwords are hashed (like ours), this approach is better than basic authentication because it means that the slow password check only has to be done periodically — either when creating a token for the first time or after a token has expired. 

The downside is that managing tokens can be complicated for clients — they will need to implement the necessary logic for caching tokens, monitoring and managing token expiry, and periodically generating new tokens. We can break down token authentication further into two sub-types: *stateful* and *stateless* token authentication.

#### ➤ Stateful token authentication

In a stateful token approach, the value of the token is a high-entropy cryptographically-secure random string. This token — or a fast hash of it — is stored server-side in a database, alongside the user ID and an expiry time for the token. When the client sends back the token in subsequent requests, your application can look up the token in the database, check that it hasn’t expired, and retrieve the corresponding user ID to find out who the request is coming from.

The big advantage of this is that your API maintains control over the tokens — it’s straightforward to revoke tokens on a per-token or per-user basis by deleting them from the database or marking them as expired. Conceptually it’s also simple and robust — the security is provided by the token being ‘unguessable’, which is why it’s important to use a high-entropy cryptographically-secure random value for the token.

So, what are the downsides? Beyond the complexity for clients that is inherent with token authentication generally, it’s difficult to find much to criticize about this approach. Perhaps the fact that it requires a database lookup is a negative — but in most cases you will need to make a database lookup to check the user’s activation status or retrieve additional information about them *anyway*.

#### ➤ Stateless token authentication

In contrast, stateless tokens encode the user ID and expiry time *in the token itself*. The token is cryptographically signed to prevent tampering and (in some cases) encrypted to prevent the contents being read. 

There are a few different technologies that you can use to create stateless tokens. Encoding the information in a [JWT](https://en.wikipedia.org/wiki/JSON_Web_Token) (JSON Web Token) is probably the most well-known approach, but [PASETO](https://developer.okta.com/blog/2019/10/17/a-thorough-introduction-to-paseto), [Branca](https://branca.io/) and [nacl/secretbox](https://pkg.go.dev/golang.org/x/crypto/nacl/secretbox) are viable alternatives too. Although the implementation details of these technologies are different, the overarching pros and cons in terms of authentication are similar.

The main selling point of using stateless tokens for authentication is that the work to encode and decode the token can be done in memory, and all the information required to identify the user is contained within the token itself. There’s no need to perform a database lookup to find out who a request is coming from.

The primary downside of stateless tokens is that they can’t easily be revoked once they are issued. In an emergency, you could effectively revoke *all* tokens by changing the secret used for signing your tokens (forcing all users to re-authenticate), or another workaround is to maintain a blocklist of revoked tokens in a database (although that defeats the ‘stateless’ aspect of having stateless tokens).

> **Note:** You should generally avoid storing additional information in a stateless token, such as a user’s activation status or permissions, and using that as the basis for *authorization* checks. During the lifetime of the token, the information encoded into it will potentially become stale and out-of-sync with the real data in your system — and relying on stale data for authorization checks can easily lead to unexpected behavior for users and various security issues.

Finally, with JWTs in particular, the fact that they’re highly configurable means that there are *lots of things you can get wrong*. The [Critical vulnerabilities in JSON Web Token libraries](https://auth0.com/blog/critical-vulnerabilities-in-json-web-token-libraries/) and [JWT Security Best Practices](https://curity.io/resources/learn/jwt-best-practices/) articles provide a good introduction to the type of things you need to be careful of here.

Because of these downsides, stateless tokens — and JWTs in particular — are generally not the best choice for managing authentication in most API applications. But they *can* be very useful in a scenario where you need *delegated authentication* — where the application *creating* the authentication token is different to the application *consuming* it, and those applications don’t share any state (which means that using stateful tokens isn’t an option). For instance, if you’re building a system which has a microservice-style architecture behind the scenes, then a stateless token created by an ‘authentication’ service can subsequently be passed to other services to identify the user.

### API-key Authentication

The idea behind API-key authentication is that a user has a non-expiring secret ‘key’ associated with their account. This key should be a high-entropy cryptographically-secure random string, and a fast hash of the key (SHA256 or SHA512) should be stored alongside the corresponding user ID in your database. The user then passes their key with each request to your API in a header like this:

```http
Authorization: Key <key>
```

On receiving it, your API can regenerate the fast hash of the key and use it to lookup the corresponding user ID from your database. Conceptually, this isn’t a million miles away from the stateful token approach — the main difference is that the keys are permanent keys, rather than temporary tokens.

On one hand, this is nice for the client as they can use the same key for every request and they don’t need to write code to manage tokens or expiry. On the other hand, the user now has two long-lived secrets to manage which can potentially compromise their account: their password, and their API key.

Supporting API keys also adds additional complexity to your API application — you’ll need a way for users to regenerate their API key if they lose it or the key is compromised, and you may also wish to support multiple API keys for the same user, so they can use different keys for different purposes.

It’s also important to note that API keys themselves should only ever be communicated to users over a secure channel, and you should treat them with the same level of care that you would a user’s password.

### OAuth 2.0 / OpenID Connect

Another option is to leverage OAuth 2.0 for authentication. With this approach, information about your users (and their passwords) is stored by a third-party *identity provider* like Google or Facebook rather than yourself.

The first thing to mention here is that *OAuth 2.0 is not an authentication protocol*, and you shouldn’t really use it for authenticating users. The oauth.net website has a great article [explaining this](https://oauth.net/articles/authentication/), and I highly recommend reading it.

If you want to implement authentication checks against a third-party identity provider, you should use [OpenID Connect](https://openid.net/connect/) (which is built directly on top of OAuth 2.0). There’s a comprehensive overview of OpenID Connect [here](https://connect2id.com/learn/openid-connect), but at a very, very, high level it works like this:

- When you want to authenticate a request, you redirect the user to an ‘authentication and consent’ form hosted by the identity provider.
- If the user consents, then the identity provider sends your API an *authorization code*.
- Your API then sends the authorization code to another endpoint provided by the identity provider. They verify the authorization code, and if it’s valid they will send you a JSON response containing an *ID token*.
- This ID token is itself a JWT. You need to validate and decode this JWT to get the actual user information, which includes things like their email address, name, birth date, timezone etc.
- Now that you know who the user is, you can then implement a stateful or stateless authentication token pattern so that you don’t have to go through the whole process for every subsequent request.

Like all the other options we’ve looked at, there are pros and cons to using OpenID Connect. The big plus is that you don’t need to persistently store user information or passwords yourself. The big downside is that it’s quite complex — although there are some helper packages like [`coreos/go-oidc`](https://github.com/coreos/go-oidc) which do a good job of masking that complexity and providing a simple interface for the OpenID Connect workflow that you can hook in to.

It’s also important to point out that using OpenID Connect requires all your users to have an account with the identity provider, and the ‘authentication and consent’ step requires human interaction via a web browser — which is probably fine if your API is the back-end for a website, but not ideal if it is a ‘standalone’ API with other computer programs as clients.

### Which one should I use?

It’s difficult to give blanket guidance on what authentication approach is best to use for your API. As with most things in programming, different tools are appropriate for different jobs. But as simple, rough, rules-of-thumb:

- If your API doesn’t have ‘real’ user accounts with slow password hashes, then HTTP basic authentication can be a good — and often overlooked — fit.
- If you don’t want to store user passwords yourself, all your users have accounts with a third-party identity provider that supports OpenID Connect, and your API is the back-end for a website… then use OpenID Connect.
- If you require delegated authentication, such as when your API has a microservice architecture with different services for performing authentication and performing other tasks, then use stateless authentication tokens.
- Otherwise use API keys or stateful authentication tokens. In general:
  - Stateful authentication tokens are a nice fit for APIs that act as the back-end for a website or single-page application, as there is a natural moment when the user logs-in where they can be exchanged for user credentials.
  - In contrast, API keys can be better for more ‘general purpose’ APIs because they’re permanent and simpler for developers to use in their applications and scripts.

In the rest of this book, we’re going to implement authentication using the *stateful authentication token* pattern. In our case we’ve already built a lot of the necessary logic for this as part of our *activation tokens* work.

## 用户认证

### 生成 Authentication token

In this chapter we’re going to focus on building up the code for a new `POST/v1/tokens/authentication` endpoint, which will allow a client to exchange their credentials (email address and password) for a stateful authentication token. At a high level, the process for exchanging a user’s credentials for an authentication token will work like this:

1. 用户发送 credentials 即邮箱和密码
2. 服务器根据 email 查找用户，如果用户存在，那么比对密码
3. 如果密码匹配，那么生成 token，有效时间为 24h

#### ➤ 只返回 token 和 expiry 两个字段:

```go
const (
    ScopeActivation     = "activation"
    ScopeAuthentication = "authentication"
)

type Token struct {
    Plaintext string    `json:"token"`
    Hash      []byte    `json:"-"`
    UserID    int64     `json:"-"`
    Expiry    time.Time `json:"expiry"`
    Scope     string    `json:"-"`
}

func (app *application) invalidCredentialsResponse(w http.ResponseWriter, r *http.Request) {
    message := "invalid authentication credentials"
    app.errorResponse(w, r, http.StatusUnauthorized, message)
}
```

#### ➤ 生成 authentication token:

```go
func (app *application) createAuthenticationTokenHandler(w http.ResponseWriter, r *http.Request) {
    var input struct {
        Email    string `json:"email"`
        Password string `json:"password"`
    }
    err := app.readJSON(w, r, &input)
    if err != nil {
        app.badRequestResponse(w, r, err)
        return
    }

    v := validator.New()
    data.ValidateEmail(v, input.Email)
    data.ValidatePasswordPlaintext(v, input.Password)
    if !v.Valid() {
        app.failedValidationResponse(w, r, v.Errors)
        return
    }

    // 用 email 找不到对应用户则返回 401 Unauthorized
    user, err := app.models.Users.GetByEmail(input.Email)
    if err != nil {
        switch {
        case errors.Is(err, data.ErrRecordNotFound):
            app.invalidCredentialsResponse(w, r)
        default:
            app.serverErrorResponse(w, r, err)
        }
        return
    }

    // 比对密码
    match, err := user.Password.Matches(input.Password)
    if err != nil {
        app.serverErrorResponse(w, r, err)
        return
    }
    if !match {
        app.invalidCredentialsResponse(w, r)
        return
    }

    // 生成 token
    token, err := app.models.Tokens.New(user.ID, 24*time.Hour, data.ScopeAuthentication)
    if err != nil {
        app.serverErrorResponse(w, r, err)
        return
    }

    // 创建成功所以返回 201 Created
    err = app.writeJSON(w, http.StatusCreated, envelope{"authentication_token": token}, nil)
    if err != nil {
        app.serverErrorResponse(w, r, err)
        return
    }
}
```

### 用 Middleware 认证请求者是谁

#### ➤ 概述

Now that our clients have a way to exchange their credentials for an authentication token, let’s look at how we can use that token to *authenticate them*, so we know exactly which user a request is coming from. Essentially, once a client has an authentication token we will expect them to include it with all subsequent requests in an `Authorization` header, like so:

```http
Authorization: Bearer IEYZQUBEMPPAKPOAWTPV6YJ6RM
```

When we receive these requests, we’ll use a new `authenticate()` middleware method to execute the following logic:

- If the authentication token is not valid, then we will send the client a `401 Unauthorized` response
- If the authentication token is valid, we will look up the user details and add their details to the *request context*.
- If no `Authorization` header was provided, we will add the details for an *anonymous user* to the request context instead.

#### ➤ 创建匿名用户

```go
var AnonymousUser = &User{}

func (u *User) IsAnonymous() bool {
    return u == AnonymousUser
}
```

#### ➤ 读写 Request Context

- Every `http.Request` that our application processes has a [`context.Context`](https://golang.org/pkg/context/#Context) embedded in it, which we can use to store key/value pairs containing arbitrary data during the lifetime of the request. In this case we want to store a `User` struct containing the current user’s information.
- Any values stored in the request context have the type `any`. This means that after retrieving a value from the request context you need to assert it back to its original type before using it.
- It’s good practice to use your own custom type for the request context keys. This helps prevent naming collisions between your code and any third-party packages which are also using the request context to store information.

```go
type contextKey string

const userContextKey = contextKey("user")

func (app *application) contextSetUser(r *http.Request, user *data.User) *http.Request {
    ctx := context.WithValue(r.Context(), userContextKey, user)
    return r.WithContext(ctx)
}

func (app *application) contextGetUser(r *http.Request) *data.User {
    user, ok := r.Context().Value(userContextKey).(*data.User)
    if !ok {
        // 明明没 user 你还用这个方法, 这是逻辑错误
        panic("missing user value in request context")
    }
    return user
}
```

#### ➤ Authentication Middleware

```go
func (app *application) authenticate(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        authorizationHeader := r.Header.Get("Authorization")

        // 为空则视为匿名用户
        if authorizationHeader == "" {
            r = app.contextSetUser(r, data.AnonymousUser)
            next.ServeHTTP(w, r)
            return
        }

        // 检查格式是否为 "Bearer <token>"
        headerParts := strings.Split(authorizationHeader, " ")
        if len(headerParts) != 2 || headerParts[0] != "Bearer" {
            app.invalidAuthenticationTokenResponse(w, r)
            return
        }

        // 校验 token
        token := headerParts[1]
        v := validator.New()
        if data.ValidateTokenPlaintext(v, token); !v.Valid() {
            app.invalidAuthenticationTokenResponse(w, r)
            return
        }

        // 取出 token 对应的用户, 然后放到 request context
        user, err := app.models.Users.GetForToken(data.ScopeAuthentication, token)
        if err != nil {
            switch {
            case errors.Is(err, data.ErrRecordNotFound):
                app.invalidAuthenticationTokenResponse(w, r)
            default:
                app.serverErrorResponse(w, r, err)
            }
            return
        }

        r = app.contextSetUser(r, user)
        next.ServeHTTP(w, r)
    })
}

func (app *application) routes() http.Handler {

    return app.recoverPanic(app.rateLimit(app.authenticate(router)))
}

func (app *application) invalidAuthenticationTokenResponse(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("WWW-Authenticate", "Bearer") // 提醒客户端需要 bearer token
    message := "invalid or missing authentication token"
    app.errorResponse(w, r, http.StatusUnauthorized, message)
}
```

## 权限验证

### 只允许已激活用户访问

The first thing we’re going to do in terms of authorization is restrict access to our `/v1/movies**` endpoints — so that they can only be accessed by users who are authenticated (not anonymous), and who have activated their account. Carrying out these kinds of checks is an ideal task for some middleware, so let’s make a new `requireActivatedUser()` middleware method to handle this.

- 对匿名用户返回 `401 Unauthorized` 和错误消息 `“you must be authenticated to access this resource”`
- 对未激活用户返回 `403 Forbidden` 和错误消息 `“your user account must be activated to access this resource”`

> **Remember:** A `401 Unauthorized` response should be used when you have missing or bad authentication, and a `403 Forbidden` response should be used afterwards, when the user is authenticated but isn’t allowed to perform the requested operation.

```go
func (app *application) authenticationRequiredResponse(w http.ResponseWriter, r *http.Request) {
    message := "you must be authenticated to access this resource"
    app.errorResponse(w, r, http.StatusUnauthorized, message)
}

func (app *application) inactiveAccountResponse(w http.ResponseWriter, r *http.Request) {
    message := "your user account must be activated to access this resource"
    app.errorResponse(w, r, http.StatusForbidden, message)
}

func (app *application) requireActivatedUser(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        user := app.contextGetUser(r)
        if user.IsAnonymous() {
            app.authenticationRequiredResponse(w, r)
            return
        }

        if !user.Activated {
            app.inactiveAccountResponse(w, r)
            return
        }

        next.ServeHTTP(w, r)
    }
}

func (app *application) routes() http.Handler {

    // 这些接口只允许已激活用户访问:
    router.HandlerFunc(http.MethodGet, "/v1/movies", app.requireActivatedUser(app.listMovieHandler))
    router.HandlerFunc(http.MethodPost, "/v1/movies", app.requireActivatedUser(app.createMovieHandler))
    router.HandlerFunc(http.MethodGet, "/v1/movies/:id", app.requireActivatedUser(app.showMovieHandler))
    router.HandlerFunc(http.MethodPatch, "/v1/movies/:id", app.requireActivatedUser(app.updateMovieHandler))
    router.HandlerFunc(http.MethodDelete, "/v1/movies/:id", app.requireActivatedUser(app.deleteMovieHandler))

    return app.recoverPanic(app.rateLimit(app.authenticate(router)))
}
```

#### ➤ 拆成两个中间件

At the moment we have one piece of middleware doing two checks: first it checks that the user is authenticated (not anonymous), and second it checks that they are activated. But it’s possible to imagine a scenario where you *only* want to check that a user is authenticated, and you don’t care whether they are activated or not.

```go
func (app *application) requireAuthenticatedUser(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        user := app.contextGetUser(r)
        if user.IsAnonymous() {
            app.authenticationRequiredResponse(w, r)
            return
        }
        next.ServeHTTP(w, r)
    }
}

func (app *application) requireActivatedUser(next http.HandlerFunc) http.HandlerFunc {
    fn := func(w http.ResponseWriter, r *http.Request) {
        user := app.contextGetUser(r)
        if !user.Activated {
            app.inactiveAccountResponse(w, r)
            return
        }
        next.ServeHTTP(w, r)
    }
    
    // 可以像这样融合两个中间件
    return app.requireAuthenticatedUser(fn)
}
```

### 需要更细粒度的权限控制

#### ➤ 例如只允许部分用户修改数据

Restricting our API so that movie data can only be accessed and edited by activated users is useful, but sometimes you might need a more granular level of control. For example, in our case we might be happy for ‘regular’ users of our API to *read* the movie data (so long as they are activated), but we want to restrict *write* access to a smaller subset of trusted users.

In this chapter we’re going to introduce the concept of *permissions* to our application, so that only users who have a specific permission can perform specific operations. In our case, we’re going to create two permissions: a `movies:read` permission which will allow a user to fetch and filter movies, and a `movies:write` permission which will allow users to create, edit and delete movies.

The required permissions will align with our API endpoints like so:

| Method | URL Pattern               | Required permission |
| :----- | :------------------------ | :------------------ |
| GET    | /v1/healthcheck           | —                   |
| GET    | /v1/movies                | `movies:read`       |
| POST   | /v1/movies                | `movies:write`      |
| GET    | /v1/movies/:id            | `movies:read`       |
| PATCH  | /v1/movies/:id            | `movies:write`      |
| DELETE | /v1/movies/:id            | `movies:write`      |
| POST   | /v1/users                 | —                   |
| PUT    | /v1/users/activated       | —                   |
| POST   | /v1/tokens/authentication | —                   |

### 用户和权限是多对多关系

The classic way to manage a many-to-many relationship in a relational database like PostgreSQL is to create a *joining table* between the two entities. 

#### ➤ 权限表

| id   | code         |
| :--- | :----------- |
| 1    | movies:read  |
| 2    | movies:write |

#### ➤ joining table: users_permissions

| user_id | permission_id |
| :------ | :------------ |
| 1       | 1             |
| 2       | 1             |
| 2       | 2             |

所以用户 2 有 `movies:read` 和 `movies:write` 权限

### 创建表

```bash
migrate create -seq -ext .sql -dir ./migrations add_permissions
migrate -path ./migrations -database $DSN up
```

SQL:

```postgresql
CREATE TABLE IF NOT EXISTS permissions
(
    id   bigserial PRIMARY KEY,
    code text NOT NULL
);

CREATE TABLE IF NOT EXISTS users_permissions
(
    user_id       bigint NOT NULL REFERENCES users ON DELETE CASCADE,
    permission_id bigint NOT NULL REFERENCES permissions ON DELETE CASCADE,
    PRIMARY KEY (user_id, permission_id) -- 联合主键, 可用来去重
);

-- 添加两个权限
INSERT INTO permissions (code)
VALUES ('movies:read'),
       ('movies:write');
```

### CRUD

#### ➤ 查询用户的权限

```go
type Permissions []string

func (p Permissions) Include(code string) bool {
    for i := range p {
        if p[i] == code {
            return true
        }
    }
    return false
}

type PermissionModel struct {
    DB *sql.DB
}

func (m PermissionModel) GetAllForUser(userID int64) (Permissions, error) {
    // JOIN 两个表和三个表的语义完全不同哦, JOIN 三个表是多对多关系, 并且确保 user_id 一定在 users 表中
    query := `
        SELECT permissions.code
        FROM permissions
        INNER JOIN users_permissions ON users_permissions.permission_id = permissions.id
        INNER JOIN users ON users_permissions.user_id = users.id
        WHERE users.id = $1`

    ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
    defer cancel()

    // 别忘了关闭 rows
    rows, err := m.DB.QueryContext(ctx, query, userID)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    // 别忘了检查 rows.Err()
    var permissions Permissions
    for rows.Next() {
        var permission string
        err := rows.Scan(&permission)
        if err != nil {
            return nil, err
        }
        permissions = append(permissions, permission)
    }
    if err = rows.Err(); err != nil {
        return nil, err
    }

    return permissions, nil
}
```

### 检查用户权限

Conceptually, what we need to do here isn’t too complicated.

- We’ll make a new `requirePermission()` middleware which accepts a specific permission code like `"movies:read"` as an argument.
- In this middleware we’ll retrieve the current user from the request context, and call the `app.models.Permissions.GetAllForUser()` method (which we just made) to get a slice of their permissions.
- Then we can check to see if the slice contains the specific permission code needed. If it doesn’t, we should send the client a `403 Forbidden` response.

```go
func (app *application) notPermittedResponse(w http.ResponseWriter, r *http.Request) {
    message := "your user account doesn't have the necessary permissions to access this resource"
    app.errorResponse(w, r, http.StatusForbidden, message)
}

func (app *application) requirePermission(code string, next http.HandlerFunc) http.HandlerFunc {
    fn := func(w http.ResponseWriter, r *http.Request) {
        user := app.contextGetUser(r)
        permissions, err := app.models.Permissions.GetAllForUser(user.ID)
        if err != nil {
            app.serverErrorResponse(w, r, err)
            return
        }

        if !permissions.Include(code) {
            app.notPermittedResponse(w, r)
            return
        }

        next.ServeHTTP(w, r)
    }
    // 融合三个检查: 非匿名, 已激活, 具有特定权限
    return app.requireActivatedUser(fn)
}

func (app *application) routes() http.Handler {

    router.HandlerFunc(http.MethodGet, "/v1/movies", app.requirePermission("movies:read", app.listMovieHandler))
    router.HandlerFunc(http.MethodPost, "/v1/movies", app.requirePermission("movies:write", app.createMovieHandler))
    router.HandlerFunc(http.MethodGet, "/v1/movies/:id", app.requirePermission("movies:read", app.showMovieHandler))
    router.HandlerFunc(http.MethodPatch, "/v1/movies/:id", app.requirePermission("movies:write", app.updateMovieHandler))
    router.HandlerFunc(http.MethodDelete, "/v1/movies/:id", app.requirePermission("movies:write", app.deleteMovieHandler))

}
```

### 添加用户权限

When a new user registers an account they *don’t have any permissions*. In order to grant permissions to a user we’ll need to update our `PermissionModel` to include an `AddForUser()` method, which adds *one or more permission codes for a specific user* to our database.

```go
// Add the "movies:read" and "movies:write" permissions for the user with ID = 2.
app.models.Permissions.AddForUser(2, "movies:read", "movies:write")
```

The SQL statement that we need to insert this data looks like this:

```postgresql
INSERT INTO users_permissions
SELECT $1, permissions.id FROM permissions WHERE permissions.code = ANY($2)
```

In this query the `$1` parameter will be the user’s ID, and the `$2` parameter will be *a PostgreSQL array* of the permission codes that we want to add for the user, like `{'movies:read', 'movies:write'}`. So what’s happening here is that the `SELECT ...` statement on the second line creates an ‘interim’ table with rows made up of the user ID *and the corresponding IDs for the permission codes in the array*. Then we insert the contents of this interim table into our `user_permissions` table.

```go
func (m PermissionModel) AddForUser(userID int64, codes ...string) error {
    query := `
        INSERT INTO users_permissions
        SELECT $1, permissions.id FROM permissions WHERE permissions.code = ANY($2)`
    
    ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
    defer cancel()
    
    _, err := m.DB.ExecContext(ctx, query, userID, pq.Array(codes))
    return err
}
```

Let’s update our `registerUserHandler` so that new users are automatically granted the `movies:read` permission when they register.

```go
func (app *application) registerUserHandler(w http.ResponseWriter, r *http.Request) {

    // CRUD
    err = app.models.Users.Insert(user)
    if err != nil {

    }

    err = app.models.Permissions.AddForUser(user.ID, "movies:read")
    if err != nil {
        app.serverErrorResponse(w, r, err)
        return
    }

}
```

