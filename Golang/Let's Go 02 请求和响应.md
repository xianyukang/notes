## Table of Contents
  - [Response Header](#Response-Header)
    - [ResponseWriter](#ResponseWriter)
    - [设置 Header](#%E8%AE%BE%E7%BD%AE-Header)
    - [自动生成的 Header](#%E8%87%AA%E5%8A%A8%E7%94%9F%E6%88%90%E7%9A%84-Header)
    - [自动标准化 Header 名](#%E8%87%AA%E5%8A%A8%E6%A0%87%E5%87%86%E5%8C%96-Header-%E5%90%8D)
    - [处理 Query String](#%E5%A4%84%E7%90%86-Query-String)
  - [读取表单](#%E8%AF%BB%E5%8F%96%E8%A1%A8%E5%8D%95)
    - [添加表单页](#%E6%B7%BB%E5%8A%A0%E8%A1%A8%E5%8D%95%E9%A1%B5)
    - [解析表单数据](#%E8%A7%A3%E6%9E%90%E8%A1%A8%E5%8D%95%E6%95%B0%E6%8D%AE)
    - [用三方库解析表单](#%E7%94%A8%E4%B8%89%E6%96%B9%E5%BA%93%E8%A7%A3%E6%9E%90%E8%A1%A8%E5%8D%95)
    - [进一步化简表单解析](#%E8%BF%9B%E4%B8%80%E6%AD%A5%E5%8C%96%E7%AE%80%E8%A1%A8%E5%8D%95%E8%A7%A3%E6%9E%90)
    - [关于 r.Form](#%E5%85%B3%E4%BA%8E-rForm)
    - [别用 r.FormValue](#%E5%88%AB%E7%94%A8-rFormValue)
    - [别用 Get 处理多值字段](#%E5%88%AB%E7%94%A8-Get-%E5%A4%84%E7%90%86%E5%A4%9A%E5%80%BC%E5%AD%97%E6%AE%B5)
  - [校验表单数据](#%E6%A0%A1%E9%AA%8C%E8%A1%A8%E5%8D%95%E6%95%B0%E6%8D%AE)
    - [校验表单数据](#%E6%A0%A1%E9%AA%8C%E8%A1%A8%E5%8D%95%E6%95%B0%E6%8D%AE)
    - [常用校验逻辑](#%E5%B8%B8%E7%94%A8%E6%A0%A1%E9%AA%8C%E9%80%BB%E8%BE%91)
    - [展示错误消息](#%E5%B1%95%E7%A4%BA%E9%94%99%E8%AF%AF%E6%B6%88%E6%81%AF)
    - [封装 Validator](#%E5%B0%81%E8%A3%85-Validator)
  - [JSON 序列化](#JSON-%E5%BA%8F%E5%88%97%E5%8C%96)
    - [Encoder vs Marshal](#Encoder-vs-Marshal)
    - [自定义 json 序列化](#%E8%87%AA%E5%AE%9A%E4%B9%89-json-%E5%BA%8F%E5%88%97%E5%8C%96)
    - [几条序列化规则](#%E5%87%A0%E6%9D%A1%E5%BA%8F%E5%88%97%E5%8C%96%E8%A7%84%E5%88%99)
    - [隐藏敏感或空字段](#%E9%9A%90%E8%97%8F%E6%95%8F%E6%84%9F%E6%88%96%E7%A9%BA%E5%AD%97%E6%AE%B5)
    - [json 序列化陷阱](#json-%E5%BA%8F%E5%88%97%E5%8C%96%E9%99%B7%E9%98%B1)
  - [JSON 反序列化](#JSON-%E5%8F%8D%E5%BA%8F%E5%88%97%E5%8C%96)
    - [解析 json 请求](#%E8%A7%A3%E6%9E%90-json-%E8%AF%B7%E6%B1%82)
    - [更健壮的请求解析](#%E6%9B%B4%E5%81%A5%E5%A3%AE%E7%9A%84%E8%AF%B7%E6%B1%82%E8%A7%A3%E6%9E%90)
    - [自定义 json 反序列化](#%E8%87%AA%E5%AE%9A%E4%B9%89-json-%E5%8F%8D%E5%BA%8F%E5%88%97%E5%8C%96)
    - [如何区分零值和字段不存在](#%E5%A6%82%E4%BD%95%E5%8C%BA%E5%88%86%E9%9B%B6%E5%80%BC%E5%92%8C%E5%AD%97%E6%AE%B5%E4%B8%8D%E5%AD%98%E5%9C%A8)

## Response Header

### ResponseWriter

```go
type ResponseWriter interface {
    Header() Header              // (1) 设置响应头
    WriteHeader(statusCode int)  // (2) 写入状态码
    Write([]byte) (int, error)   // (3) 写入响应数据
}
```

*These methods must be called in a specific order*. 

1. First, call `w.Header()` to get an instance of http.Header and set any response headers you need. If you don’t need to set any headers, you don’t need to call it. 
   
2. Next, call `w.WriteHeader()` with the HTTP status code for your response. If you are sending a response that has a 200 status code, you can skip WriteHeader. 
   
3. Finally, call the `w.Write()` method to set the body for the response.

### 设置 Header

#### ➤ 设置 Status Code

It’s only possible to call `w.WriteHeader()` once per response, and after the status code has been written it can’t be changed. If you don’t call w.WriteHeader() explicitly, then the first call to w.Write() will automatically send a 200 status code to the user. So, if you want to send a non-200 status code, you must call w.WriteHeader() before any call to w.Write()

#### ➤ 调用 w.WriteHeader() 或 w.Write() 后无法继续设置响应头

Important: Changing the response header map after a call to `w.WriteHeader()` or `w.Write()` will have no effect. 
You need to make sure that your response header map contains all the headers you want before you call these methods.

#### ➤ 读写 Response Header

We can add a new header to the response header map by using the `w.Header().Set()` method. here’s also `Add()`, `Del()`, `Get()` and `Values()` methods that you can use to read and manipulate the header map too.

```go
w.Header().Set("Cache-Control", "public, max-age=36000")   // 已经存在的同名 Header 会被覆盖
w.Header().Add("Cache-Control", "public")                  // 同名 Header 会被保留
w.Header().Add("Cache-Control", "max-age=36000")           // 会出现两行 Cache-Control
w.Header().Del("Cache-Control")                            // 删掉所有这个名字的 Header
w.Header().Get("Cache-Control")                            // 若存在多个同名 Header，Get() 只返回首个值
w.Header().Values("Set-Cookie")                            // HTTP 允许多个同名 Header，Values() 返回一个列表
```

### 自动生成的 Header

When sending a response Go will automatically set `Date` and `Content-Length` and `Content-Type` headers for you. Go will attempt to set the correct one for you by content sniffing the response body with the `http.DetectContentType()` function. If this function can’t guess the content type, Go will fall back to setting the header `Content-Type: application/octet-stream` instead.

The `http.DetectContentType()` function generally works quite well, but a common gotcha for web developers new to Go is that it can’t distinguish JSON from plain text. So, by default, JSON responses will be sent with a `Content-Type: text/plain; charset=utf-8` header. You can prevent this from happening by setting the correct header manually like so

```go
w.Header().Set("Content-Type", "application/json")
w.Write([]byte(`{"name":"Alex"}`))
```

The `Del()` method doesn’t remove system-generated headers. To suppress these, you need to access the underlying header map directly and set the value to nil. If you want to suppress the Date header, for example, you need to write: `w.Header()["Date"] = nil`

### 自动标准化 Header 名

When you’re using the Set(), Add(), Del(), Get() and Values() methods on the header map, the header name will always be canonicalized using the `textproto.CanonicalMIMEHeaderKey()` function. Canonical form means that

1. first character and after that any character following a hyphen is in uppercase. 

2. All other characters will be in lowercase.

总而言之:

- 如果服务器收到 content-type: applcation/json，那么 http 库会把请求头标准化为 Content-Type: applcation/json
- 如果用 w.Header().Set("X-XSS-Protection", "...") 设置响应头，那么 http 库会把响应头标准化为 X-Xss-Protection
- 如果不想 Header 名发生变化，可以直接访问底层 map: w.Header()["X-XSS-Protection"] = []string{"1; mode=block"}

在 HTTP/2 标准中，A request or response containing uppercase header field names MUST be treated as malformed. If a HTTP/2 connection is being used, Go will always automatically convert the header names and values to lowercase for you as per [the HTTP/2 specifications](https://tools.ietf.org/html/rfc7540#section-8.1.2).

### 处理 Query String

- `r.URL.Query().Get()` will always return a string value for a parameter, or the empty string "" if no matching parameter exists.

```go
func testQueryString(w http.ResponseWriter, r *http.Request) {
    // 校验 id 是数字且为正数
    id, err := strconv.Atoi(r.URL.Query().Get("id"))
    if err != nil || id < 1 {
        http.NotFound(w, r)
        return
    }
    _, _ = fmt.Fprintf(w, "id is %d", id)
}
```

## 读取表单

### 添加表单页

Let’s begin by making a new `ui/html/pages/create.tmpl` file to hold the HTML for the form:

```html
{{define "title"}}Create a New Snippet{{end}}
{{define "main"}}
    <form action='/snippet/create' method='POST'>
        <div>
            <label>Title:</label>
            <input type='text' name='title'>
        </div>
        <div>
            <label>Content:</label>
            <textarea name='content'></textarea>
        </div>
        <div>
            <label>Delete in:</label>
            <input type='radio' name='expires' value='365' checked> One Year
            <input type='radio' name='expires' value='7'> One Week
            <input type='radio' name='expires' value='1'> One Day
        </div>
        <div>
            <input type='submit' value='Publish snippet'>
        </div>
    </form>
{{end}}
```

We need to update the `snippetCreateForm` handler so that it renders our new page like so:

```go
func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {
    data := app.newTemplateData(r)
    app.render(w, http.StatusOK, "create.tmpl", data)
}
```

### 解析表单数据

At a high-level we can break this down into two distinct steps.

1. First, we need to use the `r.ParseForm()` method to parse the request body. This checks that the request body is well-formed, and then stores the form data in the request’s `r.PostForm` map. If there are any errors encountered when parsing the body (like there is no body, or it’s too large to process) then it will return an error. The `r.ParseForm()` method is also idempotent; it can safely be called multiple times on the same request without any side-effects
2. We can then get to the form data contained in `r.PostForm` by using the `r.PostForm.Get()` method. For example, we can retrieve the value of the `title` field with `r.PostForm.Get("title")`. If there is no matching field name in the form this will return the empty string `""`, similar to the way that query string parameters worked earlier in the book.  

```go
func (app *application) snippetCreatePost(w http.ResponseWriter, r *http.Request) {
    // 先调用 r.ParseForm() 把表单数据解析到 r.PostForm
    err := r.ParseForm()
    if err != nil {
        app.clientError(w, http.StatusBadRequest)
        return
    }

    // 通过 r.PostForm.Get() 方法读取字段值, 表单中未提供的字段会返回空字符串 ""
    title := r.PostForm.Get("title")
    content := r.PostForm.Get("content")

    // 注意 r.PostForm.Get() 始终返回字符串类型, 所以有时候要转一下类型
    expires, err := strconv.Atoi(r.PostForm.Get("expires"))
    if err != nil {
        app.clientError(w, http.StatusBadRequest)
        return
    }

    // 插入记录
    id, err := app.snippets.Insert(title, content, expires)
    if err != nil {
        app.serverError(w, err)
        return
    }
    
    // Redirect the user to the relevant page for the snippet.
    // 补充阅读: https://stackoverflow.com/questions/4764297/difference-between-http-redirect-codes
    http.Redirect(w, r, fmt.Sprintf("/snippet/view/%d", id), http.StatusSeeOther)
}
```

### 用三方库解析表单

Another thing we can do to simplify our handlers is use a third-party package like `go-playground/form` to automatically decode the form data into the `createSnippetForm` struct. Using an automatic decoder is totally optional, but it can help to save you time and typing — especially if your application has lots of forms.

Please go ahead and install it like so: `go get github.com/go-playground/form/v4@v4`. To get this working, the first thing that we need to do is initialize a new `*form.Decoder` instance in our `main.go` file and make it available to our handlers as a dependency.

```go
type application struct {
    formDecoder   *form.Decoder                 // 添加 decoder 字段供 handler 使用
}
func main() {
    formDecoder := form.NewDecoder()
    app := &application{
        formDecoder:   formDecoder,
    }
}
```

Next let’s go to our `cmd/web/handlers.go` file and update it to use this new decoder, like so:

```go
type snippetCreateForm struct {
    Title               string `form:"title"`
    Content             string `form:"content"` // 通过 struct tag 告诉 decoder
    Expires             int    `form:"expires"` // 如何把 HTML 表单映射或对应到结构体字段
    validator.Validator `form:"-"`              // 嵌入 Validator, 并且在解析表单时忽略此字段
}

func (app *application) snippetCreatePost(w http.ResponseWriter, r *http.Request) {
    // 先调用 r.ParseForm() 把表单数据解析到 r.PostForm
    err := r.ParseForm()
    if err != nil {
        app.clientError(w, http.StatusBadRequest)
        return
    }

    var form snippetCreateForm                      // 字段很多时用 Decoder 解析更方便
    err = app.formDecoder.Decode(&form, r.PostForm) // 解析到结构体时, 一定要记得传指针
    if err != nil {
        app.clientError(w, http.StatusBadRequest)
        return
    }
}
```

### 进一步化简表单解析

When we call `app.formDecoder.Decode()` it requires a non-nil pointer as the target decode destination. If we try to pass in something that isn’t a non-nil pointer, then `Decode()` will return a `form.InvalidDecoderError` error. To assist with this, let’s create a new `decodePostForm()` helper which does three things:

- Calls `r.ParseForm()` on the current request.  
- Calls `app.formDecoder.Decode()` to unpack the HTML form data to a target destination.
- Checks for a `form.InvalidDecoderError` error and triggers a panic if we ever see it.  

```go
func (app *application) decodePostForm(r *http.Request, dst any) error {
    err := r.ParseForm()
    if err != nil {
        return err
    }

    // 如果 dst 不是 non-nil 指针, Decode 方法的返回值是 *form.InvalidDecoderError 类型
    err = app.formDecoder.Decode(dst, r.PostForm)
    if err != nil {
        var invalidDecoderError *form.InvalidDecoderError
        if errors.As(err, &invalidDecoderError) {
            panic(err) // 出现这种异常, 说明代码写错了, 所以 panic
        }
        return err // 是其他解析错误则直接返回
    }
    return nil
}

func (app *application) snippetCreatePost(w http.ResponseWriter, r *http.Request) {
    var form snippetCreateForm
    err := app.decodePostForm(r, &form) // 代码更漂亮了, 重复更少了 !
    if err != nil {
        app.clientError(w, http.StatusBadRequest)
        return
    }
}
```

### 关于 r.Form

In our code above, we accessed the form values via the `r.PostForm` map. But an alternative approach is to use the (subtly different) `r.Form` map. The `r.PostForm` map is populated only for `POST`, `PATCH` and `PUT` requests, and contains the form data from the request body.

In contrast, the `r.Form` map is populated for all requests (irrespective of their HTTP method), and contains the form data from request body and query string parameters. So, if our form was submitted to `/snippet/create?foo=bar`, we could also get the value of the `foo` parameter by calling `r.Form.Get("foo")`. Note that in the event of a conflict, the request body value will take precedent over the query string parameter.

Using the `r.Form` map can be useful if your application sends data in a HTML form and in the URL, or you have an application that is agnostic about how parameters are passed. But in our case those things aren’t applicable. We expect our form data to be sent in the request body only, so it’s sensible for us to access it via `r.PostForm`.

### 别用 r.FormValue

The `net/http` package also provides the methods `r.FormValue()` and `r.PostFormValue()`. These are essentially shortcut functions that call `r.ParseForm()` for you, and then fetch the appropriate field value from `r.Form` or `r.PostForm` respectively.

I recommend avoiding these shortcuts because they silently ignore any errors returned by `r.ParseForm()`. That’s not ideal — it means our application could be encountering errors and failing for users, but there’s no feedback mechanism to let them know.

### 别用 Get 处理多值字段

Strictly speaking, the `r.PostForm.Get()` method that we’ve used above only returns the first value for a specific form field. This means you can’t use it with form fields which potentially send multiple values, such as a group of checkboxes.

```html
<input type="checkbox" name="items" value="foo"> Foo
<input type="checkbox" name="items" value="bar"> Bar
<input type="checkbox" name="items" value="baz"> Baz
```

In this case you’ll need to work with the `r.PostForm` map directly. The underlying type of the `r.PostForm` map is `url.Values`, which in turn has the underlying type `map[string][]string`. So, for fields with multiple values you can loop over the underlying map to access them like so: 

```go
for i, item := range r.PostForm["items"] {
    fmt.Fprintf(w, "%d: Item %s\n", i, item)
} 
```

## 校验表单数据

### 校验表单数据

Right now there’s a glaring problem with our code: we’re not validating the (untrusted) user input from the form in any way. Specifically for this form we want to:

- Check that the `title` and `content` fields are not empty.
- Check that the `title` field is not more than 100 characters long.
- Check that the `expires` value exactly matches one of our permitted values (1, 7 or 365 days).  

```go
func (app *application) snippetCreatePost(w http.ResponseWriter, r *http.Request) {
    fieldErrors := make(map[string]string)
    if strings.TrimSpace(title) == "" {
        fieldErrors["title"] = "This field cannot be blank"
    } else if utf8.RuneCountInString(title) > 100 {
        fieldErrors["title"] = "This field cannot be more than 100 characters long"
    }

    if strings.TrimSpace(content) == "" {
        fieldErrors["content"] = "This field cannot be blank"
    }

    if expires != 1 && expires != 7 && expires != 365 {
        fieldErrors["expires"] = "This field must equal 1, 7 or 365"
    }

    if len(fieldErrors) > 0 {
        _, _ = fmt.Fprint(w, fieldErrors)
        return
    }
}
```

### 常用校验逻辑

```go
func TestInputValidation(t *testing.T) {
    r := http.Request{}
    title := r.PostForm.Get("title")
    _ = title == ""                         // 检查 empty 字符串
    _ = strings.TrimSpace(title) == ""      // 检查 blank 字符串
    _ = utf8.RuneCountInString(title) > 100 // 检查字符串长度
    strings.HasPrefix(title, "abc")         // 字符串 starts with
    strings.HasSuffix(title, "abc")         // 字符串 ends with
    strings.Contains(title, "abc")          // 字符串 contains
}

func TestRegexp(t *testing.T) {
    var _ = regexp.MustCompile(`abc.com`)                             // 一般要保留编译结果, 效率更高
    println(regexp.MustCompile(`^abc.com$`).MatchString("abc_com"))   // 字符 . 在正则中有特殊含义
    println(regexp.MustCompile(`^abc\.com$`).MatchString("abc.com"))  // 使用 \. 转义一下
    println(regexp.MustCompile("^abc\\.com$").MatchString("abc.com")) // 使用 "" 则需要 \\ 表示一个 \

    println("\u5149\u002f\u7130")                         // 使用 16 进制 unicode 码点表示 "光/焰"
    regexp.MustCompile("^[\u0400-\u04FF\u0500-\u052F]+$") // 检查字符串位于这两个 unicode range
}

func TestCheckEmail(t *testing.T) {
    // 怎样才算完美的邮箱正则, 这是个有争议的话题, https://stackoverflow.com/questions/201323
    // 如下是 WHATWG 推荐的正则, https://html.spec.whatwg.org/multipage/input.html#valid-e-mail-address
    //goland:noinspection ALL 此处 GoLand 会提示说: "正则表达式中没必要用 \/ 转义 / 字符", 但我乐意这么干
    var rxEmail = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

    email := "123@abc.com"
    // 邮箱长度应该 <= 254 个字节, https://stackoverflow.com/questions/386294
    if len(email) > 254 || !rxEmail.MatchString(email) {
        fmt.Println("error: foo is not a valid email address")
    }
}
func TestCheckURL(t *testing.T) {
    // 校验 URL
    u, err := url.Parse("http://localhost:8080")
    if err != nil {
        fmt.Println("error: not a valid URL")
    } else if u.Scheme == "" || u.Host == "" {
        fmt.Println("error: must be an absolute URL")
    } else if u.Scheme != "http" && u.Scheme != "https" {
        fmt.Println("error: must begin with http or https")
    }
}

func TestNumber(t *testing.T) {
    // 校验整数
    r := &http.Request{}
    n, err := strconv.Atoi(r.Form.Get("foo"))
    if err != nil {
        fmt.Println("error: foo must be an integer")
    } else if n < 0 || n > 10 {
        fmt.Printf("error: foo must be between 0 and 100")
    }

    // 校验浮点数
    f, err := strconv.ParseFloat(r.Form.Get("foo"), 64)
    if err != nil {
        fmt.Println("error: foo must be a float")
    } else if f < 0 || f > 1 {
        fmt.Printf("error: foo must be between 0 and 1")
    }
}

func TestDateTime(t *testing.T) {
    // 校验日期
    r := &http.Request{}
    d, err := time.Parse("2006-01-02", r.Form.Get("foo"))
    if err != nil {
        fmt.Printf("error: foo is not a valid date")
    } else if d.Year() != 2017 {
        fmt.Printf("error: foo is not between 2017-01-01 and 2017-12-31")
    }

    // 校验用户提交了合法的时间, 虽然日常生活中不会强调时区, 但现实中的时间都是带时区的
    // is valid and between 2017-01-01 00:00 and 2017-12-31 23:59:
    loc, err := time.LoadLocation("Europe/Vienna")
    if err != nil {
        // Windows 用户可能需要自己安装个 zoneinfo.zip 文件
    }
    dt, err := time.ParseInLocation("2006-01-02T15:04:05", r.Form.Get("foo"), loc)
    if err != nil {
        fmt.Printf("error: foo is not a valid datetime")
    } else if dt.Year() != 2017 {
        fmt.Printf("error: foo is not between 2017-01-01 00:00:00 and 2017-12-31 23:59:00")
    }
    fmt.Println()
}
```

### 展示错误消息

If there are any validation errors we want to re-display the HTML form, highlighting the fields which failed validation and automatically re-populating any previously submitted data (so that the user doesn’t need to enter it again). To do this, let’s begin by adding a new `Form` field to our `templateData` struct:

```go
type templateData struct {
    CurrentYear int               //  
    Snippet     *models.Snippet   // 
    Snippets    []*models.Snippet // 
    Form        any               // 把表单数据返回给用户
}
```

Next let’s head back to our `cmd/web/handlers.go` file and define a new `snippetCreateForm` struct type to hold the form data and any validation errors, and update our `snippetCreatePost` handler to use this.

```go
type snippetCreateForm struct {
    Title       string
    Content     string
    Expires     int
    FieldErrors map[string]string
}

func (app *application) snippetCreatePost(w http.ResponseWriter, r *http.Request) {
    err := r.ParseForm()
    if err != nil {
        app.clientError(w, http.StatusBadRequest)
        return
    }
    expires, err := strconv.Atoi(r.PostForm.Get("expires"))
    if err != nil {
        app.clientError(w, http.StatusBadRequest)
        return
    }
    form := snippetCreateForm{
        Title:       r.PostForm.Get("title"),
        Content:     r.PostForm.Get("content"),
        Expires:     expires,
        FieldErrors: map[string]string{},
    }

    // 校验表单数据
    if strings.TrimSpace(form.Title) == "" {
        form.FieldErrors["title"] = "This field cannot be blank"
    } else if utf8.RuneCountInString(form.Title) > 100 {
        form.FieldErrors["title"] = "This field cannot be more than 100 characters long"
    }
    if strings.TrimSpace(form.Content) == "" {
        form.FieldErrors["content"] = "This field cannot be blank"
    }
    if form.Expires != 1 && form.Expires != 7 && form.Expires != 365 {
        form.FieldErrors["expires"] = "This field must equal 1, 7 or 365"
    }

    // 如果有校验错误, 重新展示 create.tmpl 模板, 填充用户提交的数据
    if len(form.FieldErrors) > 0 {
        data := app.newTemplateData(r)
        data.Form = form
        app.render(w, http.StatusUnprocessableEntity, "create.tmpl", data) // 用 422 状态表示校验不通过
        return
    }

    id, err := app.snippets.Insert(form.Title, form.Content, form.Expires)
    if err != nil {
        app.serverError(w, err)
        return
    }
    http.Redirect(w, r, fmt.Sprintf("/snippet/view/%d", id), http.StatusSeeOther)
}
```

The next thing that we need to do is update our `create.tmpl` template to display the validation errors and re-populate any previous data. For the validation errors, the underlying type of our `FieldErrors` field
is a `map[string]string`, which uses the form field names as keys. For maps, it’s possible to access the value for a given key by simply chaining the key name. So, for example, to render a validation error for the `title` field we can use the tag `{{.Form.FieldErrors.title}}` in our template.

```html
{{define "title"}}Create a New Snippet{{end}}
{{define "main"}}
    <form action='/snippet/create' method='POST'>
        <div>
            <label>Title:</label>
            <!-- 通过 with 动作, 有则展示错误消息 -->
            {{with .Form.FieldErrors.title}}
                <label class="error">{{.}}</label>
            {{end}}
            <!-- 填充用户提交的表单数据,  -->
            <input type='text' name='title' value="{{.Form.Title}}">
        </div>
        <div>
            <label>Content:</label>
            {{with .Form.FieldErrors.content}}
                <label class="error">{{.}}</label>
            {{end}}
            <textarea name='content'>{{.Form.Content}}</textarea>
        </div>
        <div>
            <label>Delete in:</label>
            {{with .Form.FieldErrors.expires}}
                <label class="error">{{.}}</label>
            {{end}}
            <!-- 填充单选框 -->
            <input type='radio' name='expires' value='365' {{if (eq .Form.Expires 365)}}checked{{end}}> One Year
            <input type='radio' name='expires' value='7' {{if (eq .Form.Expires 7)}}checked{{end}}> One Week
            <input type='radio' name='expires' value='1' {{if (eq .Form.Expires 1)}}checked{{end}}> One Day
        </div>
        <div>
            <input type='submit' value='Publish snippet'>
        </div>
    </form>
{{end}}
```

Our `snippetCreate` handler currently doesn’t set a value for the `templateData.Form` field, meaning that when Go tries to evaluate a template tag like `{{with .Form.FieldErrors.title}}` it would result in an error because `Form` is `nil`.

```go
func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {
    // 模板中用到了 .Form.FieldErrors.title 所以应该确保 .Form 字段不为 nil
    // 一般还要确保 .Form.FieldErrors 也不为 nil, 但它是 map 类型, 读取 nil map 返回零值, 写入 nil map 则 panic
    data := app.newTemplateData(r)
    data.Form = snippetCreateForm{
        Expires: 365, // 设置表单字段的默认值
    }
    app.render(w, http.StatusOK, "create.tmpl", data)
}
```

### 封装 Validator

And while the approach we’ve taken is fine as a one-off, if your application has many forms then you can end up with quite a lot of repetition in your code and validation rules. So to help us with validation throughout the rest of this project, we’ll create our own small `internal/validator` package to abstract some of this behavior and reduce the boilerplate code in our handlers.

```go
type Validator struct {
    FieldErrors map[string]string
}

func (v *Validator) Valid() bool {
    return len(v.FieldErrors) == 0
}

func (v *Validator) AddFieldError(key, message string) {
    // 如果内部 map 不存在则初始化
    if v.FieldErrors == nil {
        v.FieldErrors = make(map[string]string)
    }

    // 已经存在则不添加
    if _, exists := v.FieldErrors[key]; !exists {
        v.FieldErrors[key] = message
    }
}

func (v *Validator) CheckField(ok bool, key, message string) {
    if !ok {
        v.AddFieldError(key, message)
    }
}

// 常用的校验函数如下:
func NotBlank(value string) bool {
    return strings.TrimSpace(value) != ""
}

func MaxChars(value string, n int) bool {
    return utf8.RuneCountInString(value) <= n
}

func PermittedInt(value int, permittedValues ...int) bool {
    for i := range permittedValues {
        if value == permittedValues[i] {
            return true
        }
    }
    return false
}
```

We’ll embed a `Validator` instance in our `snippetCreateForm` struct:

```go
type snippetCreateForm struct {
    Title               string
    Content             string
    Expires             int
    validator.Validator // 嵌入 Validator
}

func (app *application) snippetCreatePost(w http.ResponseWriter, r *http.Request) {
    // 校验表单数据 ( 封装 Validator 后代码看起来更漂亮了 ! )
    form.CheckField(validator.NotBlank(form.Title), "title", "This field cannot be blank")
    form.CheckField(validator.MaxChars(form.Title, 100), "title", "This field cannot be more than 100 characters long")
    form.CheckField(validator.NotBlank(form.Content), "content", "This field cannot be blank")
    form.CheckField(validator.PermittedInt(form.Expires, 1, 7, 365), "expires", "This field must equal 1, 7 or 365")

    // 如果有校验错误, 重新展示 create.tmpl 模板, 填充用户提交的数据
    if !form.Valid() {
        data := app.newTemplateData(r)
        data.Form = form
        app.render(w, http.StatusUnprocessableEntity, "create.tmpl", data) // 用 422 状态表示校验不通过
        return
    }
}

```

We’ve now got an `internal/validator` package with validation rules and logic that can be reused across our application, and it can easily be extended to include additional rules in the future. Both form data and errors are neatly encapsulated in a single `snippetCreateForm` struct — which we can easily pass to our templates.

## JSON 序列化

#### ➤ 返回 JSON 响应

At this stage, the thing I’d like to emphasize is that JSON is just text. In fact, the only special thing we need to do is set a `Content-Type: application/json` header on the response, so that the client knows it’s receiving JSON and can interpret it accordingly.

#### ➤ 这有一个 JSON API 的规范: [JSON:API — A specification for building APIs in JSON](https://jsonapi.org/)

It’s important to emphasize that there’s no single right or wrong way to structure your JSONresponses. But whatever you do, it is valuable to think about formatting upfront and to maintain a clear and consistent response structure across your different API endpoints.

### Encoder vs Marshal

#### ➤ Marshal / Unmarshal

```markdown
- 需要和字符串打交道
- 作用是把对象序列化成 JSON 字符串
```

#### ➤ Encoder / Decoder

```markdown
- 需要和 Reader/Writer 打交道
- 作用是把序列化结果写入一个流
- 支持定制 JSON 格式，例如开启缩进、关掉HTML转义

- Encoder 有个优点是写到输出流，无需在内存中保留完整的序列化结果
- 但这同时也是缺点，如果序列化到一半时遇到错误，接收端会收到格式坏掉的 JSON，想要数据一致性就得分两步
```

### 自定义 json 序列化

#### ➤ 给字段建个类型，然后实现 json.Marshaler

The key thing to understand is this: When Go is encoding a particular type to JSON, it looks to see if the type has a `MarshalJSON()` method implemented on it. If it has, then Go will call this method to determine how to encode it. If the type satisfies the `json.Marshaler` interface, which looks like this, then Go will call its `MarshalJSON()` method and use the `[]byte` slice that it returns as the encoded JSON value.

```go
// 例如 time.Time 就实现了这个接口，每次需要序列化一个 time.Time 字段时，它的 MarshalJSON 方法就会被调用
type Marshaler interface {
    MarshalJSON() ([]byte, error)
}
```

If the type doesn’t have a `MarshalJSON()` method, then Go will fall back to trying to encode it to JSON based on its own internal set of rules. So, if we want to customize how something is encoded, all we need to do is implement a `MarshalJSON()` method on it which returns a custom JSON representation of itself in a `[]byte` slice.

```go
// 自定义序列化来实现格式转换,  比如 Movie 有个名为 runtime 的 int 字段,  我们想把 120 变成 "120 mins"
type Runtime int32

func (r Runtime) MarshalJSON() ([]byte, error) {
    jsonValue := fmt.Sprintf("%d mins", r)
    quotedJSONValue := strconv.Quote(jsonValue) // 注意给字符串加上两端的双引号,  这样才能作为 json 字符串
    return []byte(quotedJSONValue), nil
}
```

#### ➤ Movie 包含 Runtime 字段，在 Movie 中实现 json.Marshaler 也能达成相同效果

```go
// Note that there are no struct tags on the Movie struct itself.
type Movie struct {
    ID        int64
    CreatedAt time.Time
    Title     string
    Year      int32
    Runtime   int32
    Genres    []string
    Version   int32
}

func (m Movie) MarshalJSON() ([]byte, error) {
    // Declare a variable to hold the custom runtime string (this will be the empty string "" by default).
    var runtime string

    if m.Runtime != 0 {
        runtime = fmt.Sprintf("%d mins", m.Runtime)
    }

    // 上面的 Movie 结构体是源数据，这个匿名结构体负责处理序列化
    // 加上了 json tag，去掉了不需要显示的 CreatedAt 字段
    aux := struct {
        ID      int64    `json:"id"`
        Title   string   `json:"title"`
        Year    int32    `json:"year,omitempty"`
        Runtime string   `json:"runtime,omitempty"` // This is a string.
        Genres  []string `json:"genres,omitempty"`
        Version int32    `json:"version"`
    }{
        ID:      m.ID,
        Title:   m.Title,
        Year:    m.Year,
        Runtime: runtime, // Note that we assign the value from the runtime variable here.
        Genres:  m.Genres,
        Version: m.Version,
    }

    // Encode the anonymous struct to JSON, and return it.
    return json.Marshal(aux)
}
```

#### ➤ 对上述代码的优化

The downside of the approach above is that the code feels quite verbose and repetitive. You might be wondering: *is there a better way?* To reduce duplication, instead of writing out all the struct fields long-hand it’s possible to embed an alias of the `Movie` struct in the anonymous struct. Like so:

```go
// Notice that we use the - directive on the Runtime field, so that it never appears in the JSON output.
type Movie struct {
    ID        int64     `json:"id"`
    CreatedAt time.Time `json:"-"`
    Title     string    `json:"title"`
    Year      int32     `json:"year,omitempty"`
    Runtime   int32     `json:"-"`
    Genres    []string  `json:"genres,omitempty"`
    Version   int32     `json:"version"`
}

func (m Movie) MarshalJSON() ([]byte, error) {
    var runtime string
    if m.Runtime != 0 {
        runtime = fmt.Sprintf("%d mins", m.Runtime)
    }

    // MovieAlias contains all the fields that Movie has but, importantly, none of the methods. 
    type MovieAlias Movie

    // Embed the MovieAlias type inside the anonymous struct, along with a Runtime field 
    // that has the type string and the necessary struct tags. It's important that we 
    // embed the MovieAlias type here, rather than the Movie type directly, to avoid 
    // inheriting the MarshalJSON() method of the Movie type (which would result in an 
    // infinite loop during encoding).
    aux := struct {
        MovieAlias
        Runtime string `json:"runtime,omitempty"`
    }{
        MovieAlias: MovieAlias(m),
        Runtime:    runtime, // 相当于重定义 Runtime 字段，注意 MovieAlias.Runtime 会被 json:"-" 标签排除
    }

    return json.Marshal(aux)
}
// 虽然相比上一种方式减少了重复，但也减少了优点: 代码使用了更多的 trick，Runtime 字段的顺序也无法细粒度控制
```


### 几条序列化规则

- `time.Time` will be encoded as a JSON string in RFC 3339 format like *"2020-11-08T06:27:59+01:00"*
- A *[]byte slice* will be encoded as a base64-encoded JSON string, rather than as a JSON array. So, for example, a byte slice of `[]byte{'h','e','l','l','o'}` would appear as `"aGVsbG8="` in the JSON output. The base64 encoding uses padding and the standard character set.
- *Channels, functions and complex number types cannot be encoded*.  
  If you try to do so, you’ll get a `json.UnsupportedTypeError` error at runtime.  
- *Any pointer values will encode as the value pointed to*.  
  Likewise, `interface{}` values will encode as the value contained in the interface.
- When we call `json.NewEncoder(w).Encode(data)` the JSON is created and written to the `http.ResponseWriter` in a single step, which means there’s no opportunity to set HTTP response headers conditionally based on whether the `Encode()` method returns an error or not.

#### ➤ 不直接用作顶层对象，而是再包一层

例如把 `{name: "ichigo"}` 变成 `{ user: { name: "ichigo" } }`

Enveloping response data like this isn’t strictly necessary, and whether you choose to do so is partly a matter of style and taste. But there are a few tangible benefits: Including a key name (like "movie") at the top-level of the JSON helps make the response more *self-documenting*.

### 隐藏敏感或空字段

It’s possible to control the *visibility of individual struct fields* in the JSON by using the `omitempty` and `-` struct tag.

The *- (hyphen)* directive can be used when you *never want* a particular struct field to appear in the JSON output. This is useful for fields that contain internal *system information* that isn’t relevant to your users, or *sensitive information* that you don’t want to expose (like the hash of a password).

In contrast the `omitempty` directive hides a field in the JSON output if and *only if the struct field value is empty*, where empty is defined as being:

```markdown
- Equal to false, 0, or ""
- An empty array, slice or map
- A nil pointer or a nil interface value  
```

You can also prevent a struct field from appearing in the JSON output by simply making it unexported. But using the `json:"-"` struct tag is generally a better choice: it helps prevents problems if someone changes the field to be exported in the future without realizing the consequences.

Hint: If you want to *use omitempty and not change the key name* then you can leave it blank in the struct tag — like this: `json:",omitempty"`. Notice that the leading comma is still required.

### json 序列化陷阱

#### ➤ Encoder 会加换行符

```go
func main() {
   var b bytes.Buffer
   var c = Character{Name: "Cloud", From: "FF7", Age: 21}
   var _ = json.NewEncoder(&b).Encode(c)     // encoder 为 stream 设计
   fmt.Println(b.String())                   // 会在一个对象结束后加个 \n
   fmt.Println(b.Bytes()[b.Len()-1] == '\n') // 能避免 1 2 3 三个对象混合成 123
}
```

#### ➤ 默认会进行 HTML 转义，不会原样输出 <>& 等字符

The `Encoder.SetEscapeHTML` method description talks about the default encoding behavior for the and, less than and greater than characters. You can't disable this behavior for the `json.Marshal` calls. It's bad because it assumes that the primary use case for JSON is a web page, which breaks the configuration libraries and the REST/HTTP APIs by default.

```go
func main() {
    var m = map[string]string{
        "<KEY>": "<VALUE>",
    }

    data, err := json.Marshal(m)
    if err != nil {
        panic(err)
    }
    fmt.Println(string(data)) // 会被自动转义: {"\u003cKEY\u003e":"\u003cVALUE\u003e"}

    var b bytes.Buffer
    var encoder = json.NewEncoder(&b)
    encoder.SetEscapeHTML(false) // 关掉 <>& 这三个字符的自动转义
    err = encoder.Encode(m)
    if err != nil {
        panic(err)
    }
    fmt.Println(b.String()) // 原样输出: {"<KEY>":"<VALUE>"}
}
```

#### ➤ 反序列化成 interface 类型，数字的实际类型是 float64，而不是 int

By default, Go treats numeric values in JSON as `float64` numbers when you decode/unmarshal JSON data into an interface. This means the following code will fail with a panic:

```go
func main() {
    var data = []byte(`{"status": 200}`)

    var result map[string]interface{}
    if err := json.Unmarshal(data, &result); err != nil {
        fmt.Println("error:", err)
        return
    }

    var status = result["status"].(int) // panic
    fmt.Println("status value:", status)

    // #### ➤ Option 1: use the float value as-is :-)

    // #### ➤ Option 2: convert the float value to the integer type you need.
    var _ = uint64(result["status"].(float64))

    // #### ➤ Option 3: use a Decoder and tell it to represent JSON numbers using the Number interface type.
    var decoder = json.NewDecoder(bytes.NewReader(data))
    decoder.UseNumber()
    if err := decoder.Decode(&result); err != nil {
        fmt.Println("error:", err)
        return
    }
    var state, _ = result["status"].(json.Number).Int64() // Int64 returns the number as an int64.
    fmt.Println("status value:", state)

    // #### ➤ Option 4: use a struct type that maps your numeric value to the numeric type you need.
    var result2 struct {
        Status uint64 `json:"status"`
    }
    if err := json.NewDecoder(bytes.NewReader(data)).Decode(&result2); err != nil {
        fmt.Println("error:", err)
        return
    }
    fmt.Printf("result2 => %+v", result2)
}
```

## JSON 反序列化

### 解析 json 请求

#### ➤ 入门例子

```go
// 注意事项:
//    ①input 指针不能为 nil
//    ②input 结构体中的字段需要导出
//    ③可以用 struct tag 映射字段
//    ④If there is no matching struct tag, Go will attempt to decode the value into a field that matches the key name
//      (exact matches are preferred, but it will fall back to a case-insensitive match). 
//      Any JSON key/value pairs which cannot be successfully mapped to the struct fields will be silently ignored.
err := json.NewDecoder(r.Body).Decode(&input)
```

#### ➤ 解析时类型要对应

假设请求是 `{year: "2022"}`，然后结构体中 Year 字段是 int32 类型，那么解析会报错，因为字符串不能解析成数字

```markdown
JSON boolean ⇒ bool
JSON string ⇒ string
JSON number ⇒ int, uint, float, rune
JSON array  ⇒ array, slice
JSON object ⇒ struct, map
```

### 更健壮的请求解析

#### ➤ (1) 限制请求大小为 1MB

#### ➤ (2) 不允许 unknown fields

```markdown
请求中可能包含 xxx 这样的未定义的字段，一种处理方式是忽略它们
另一种方式是通知客户端 xxx 字段未定义，不允许未定义的字段
对于第二种情况，可以开启 json.Decoder 的 DisallowUnknownFields() 选项
```

#### ➤ (3) 处理 bad requests

How to deal with bad requests from clients and invalid JSON, and return clear, actionable, error messages. What if the client sends something that isn’t JSON, like XML or some random bytes? What happens if the JSON is malformed or contains an error? What if the JSON types don’t match the types we are trying to decode into?

```go
// 这是最简单的处理办法，遇到解析错误时，返回一个 error json
err := json.NewDecoder(r.Body).Decode(&input)
if err != nil {
   app.errorResponse(w, r, http.StatusBadRequest, err.Error())
   return
}
```

#### ➤ (4) 更友好的错误消息

For a private API which won’t be used by members of the public, then this behavior is probably fine and you needn’t do anything else. But for a public-facing API, the error messages themselves aren’t ideal. *Some are too detailed and expose information about the underlying API implementation*. To improve this, we’re going to explain how to *triage the errors*  returned by `Decode()` and replace them with clearer, easy-to-action, error messages to help the client debug exactly what is wrong with their JSON.

| Decode() 可能返回的错误    | 原因                                                         |
| -------------------------- | ------------------------------------------------------------ |
| io.EOF                     | The JSON being decoded is empty                              |
| io.ErrUnexpectedEOF        | There is a syntax problem with the JSON being decoded        |
| json.SyntaxError           | There is a syntax problem with the JSON being decoded        |
| json.UnmarshalTypeError    | A JSON value is not appropriate for the destination Go type  |
| json.InvalidUnmarshalError | The decode destination is not valid (usually because it is not a pointer)<br />This is actually a problem with our application code, not the JSON itself |

#### ➤ 代码如下:

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
        var syntaxError *json.SyntaxError
        var unmarshalTypeError *json.UnmarshalTypeError
        var invalidUnmarshalError *json.InvalidUnmarshalError
        var maxBytesError *http.MaxBytesError

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

        // 请求中包含 unknown fields, 要从错误消息中提取字段名 json: unknown field "<name>"
        // 有点丑陋是吧, 理论上做成一个错误类型更好, 参见 https://github.com/golang/go/issues/29035
        case strings.HasPrefix(err.Error(), "json: unknown field "):
            fieldName := strings.TrimPrefix(err.Error(), "json: unknown field ")
            return fmt.Errorf("body contains unknown field %s", fieldName)

        // 请求 body 超过了 1 MB
        case errors.As(err, &maxBytesError):
            return fmt.Errorf("body must not be larger than %d bytes", maxBytesError.Limit)

        // 我们代码有错误, 比如传给 Decode() 的参数: ①不是指针 ②是 nil 指针
        // 为什么这里用 panic 是合适的? 因为我们代码有 bug 所以用 panic 很正常啊
        case errors.As(err, &invalidUnmarshalError):
            panic(err)

        // 其他未知错误原样返回
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

### 自定义 json 反序列化

- 要求把 `duration: "123 mins"` 解析成 `duration: 123`，这种*格式转换、类型转换*需要自定义 JSON 解析
- 自定义 JSON 解析的关键是实现这个接口: type Unmarshaler interface { UnmarshalJSON([]byte) error }

示例:

```go
var ErrInvalidRuntimeFormat = errors.New("invalid runtime format")

type Runtime int32

//goland:noinspection GoMixedReceiverTypes
func (r Runtime) MarshalJSON() ([]byte, error) {
    str := fmt.Sprintf("%d mins", r)
    quoted := strconv.Quote(str) // 别忘了在 json 语法中, 字符串需要用双引号括起来
    return []byte(quoted), nil
}

//goland:noinspection GoMixedReceiverTypes
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

### 如何区分零值和字段不存在

#### ➤ 对于 json 请求，如何区分 year:0 和不提供 year 字段?

如果请求中没有提供 year 字段，那么解析后结构体中的 Year 字段值为 int32 的零值，也就是两种情况对应相同的解析结果

In the context of partial update this causes a problem. How do we tell the difference between:

- A client providing a key/value pair which has a zero-value value — like {"title": ""} —
  in which case we want to return a validation error.  
- A client not providing a key/value pair in their JSON at all — in which case we want to
  ‘skip’ updating the field but not send a validation error.  

#### ➤ 可利用指针作区分，因为它的零值是 nil

So we could change the fields in our input struct to be pointers. Then to see if a client has provided a particular key/value pair in the JSON, we can simply check whether the corresponding field in the input struct equals nil or not.  

```go
    // Declare an input struct to hold the expected data from the client.
    // Use pointers for the Title, Year and Runtime fields.
    var input struct {
        Title   *string       `json:"title"`   // This will be nil if there is no corresponding key in the JSON
        Year    *int32        `json:"year"`    // This will be nil if there is no corresponding key in the JSON
        Runtime *data.Runtime `json:"runtime"` // This will be nil if there is no corresponding key in the JSON
        Genres  []string      `json:"genres"`  // We don't need to change this because slices already have the zero-value nil
    }
```

To summarize this: we’ve changed our input struct so that all the fields now have the zerovalue nil.  After parsing the JSON request, we then go through the input struct fields and only update the movie record if the new value is not nil.  

One special-case to be aware of is when the client explicitly supplies a field in the JSON request with the value `null`. In this case, our handler will ignore the field and treat it like it hasn’t been supplied.
