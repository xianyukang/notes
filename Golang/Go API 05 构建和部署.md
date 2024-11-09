## Table of Contents
  - [Makefile](#Makefile)
    - [入门](#%E5%85%A5%E9%97%A8)
    - [使用变量](#%E4%BD%BF%E7%94%A8%E5%8F%98%E9%87%8F)
    - [传递参数](#%E4%BC%A0%E9%80%92%E5%8F%82%E6%95%B0)
    - [命名空间](#%E5%91%BD%E5%90%8D%E7%A9%BA%E9%97%B4)
    - [前置条件](#%E5%89%8D%E7%BD%AE%E6%9D%A1%E4%BB%B6)
    - [显示帮助信息](#%E6%98%BE%E7%A4%BA%E5%B8%AE%E5%8A%A9%E4%BF%A1%E6%81%AF)
  - [help: print this help message](#help-print-this-help-message)
  - [run/api: run the cmd/api application](#runapi-run-the-cmdapi-application)
    - [Phony targets](#Phony-targets)
    - [管理环境变量](#%E7%AE%A1%E7%90%86%E7%8E%AF%E5%A2%83%E5%8F%98%E9%87%8F)
  - [Building](#Building)
    - [Quality Controlling Code](#Quality-Controlling-Code)
  - [audit: tidy dependencies and format, vet and test all code](#audit-tidy-dependencies-and-format-vet-and-test-all-code)
    - [Module Proxies](#Module-Proxies)
    - [and Vendoring](#and-Vendoring)
  - [vendor: tidy and vendor dependencies](#vendor-tidy-and-vendor-dependencies)
    - [Building Binaries](#Building-Binaries)
  - [build/api: build the cmd/api application](#buildapi-build-the-cmdapi-application)
    - [Version Numbers](#Version-Numbers)
  - [Deployment and Hosting](#Deployment-and-Hosting)
    - [Digital Ocean](#Digital-Ocean)
    - [Server Configuration](#Server-Configuration)
    - [Deployment](#Deployment)
  - [production/connect: connect to the production server](#productionconnect-connect-to-the-production-server)
  - [production/deploy/api: deploy the api to production](#productiondeployapi-deploy-the-api-to-production)
    - [Background Service 1](#Background-Service-1)
    - [Background Service 2](#Background-Service-2)
    - [Reverse Proxy](#Reverse-Proxy)

## Makefile

### 入门

A makefile is essentially a text file which contains one or more *rules* that the `make` utility can run. Each rule has a *target* and contains a sequence of sequential *commands* which are executed when the rule is run. Generally speaking, makefile rules have the following structure:

```makefile
# comment (optional)
target: 
    command_1
    command_2
    @echo 注意命令的缩进符是 tab 不能用空格
    @echo 注意符号 @ 的作用为不显示当前这行命令，只显示命令的输出
```

Let’s create a rule which executes the `go run ./cmd/api` command to run our API application.

```makefile
dsn = postgres://greenlight:pa55word@localhost/greenlight?sslmode=disable
trusted_origins = http://localhost:8000 http://localhost:9000

run:
    go run ./cmd/api -db-dsn='$(dsn)' -cors-trusted-origins='$(trusted_origins)' -limiter-enabled=true
```

You can execute a specific rule by running `$ make <target>` from your terminal. When we type `make run`, the make utility looks for a file called `Makefile` or `makefile` in the current directory and then executes the commands associated with the `run` target.

### 使用变量

When we execute a `make` rule, every environment variable that is available to `make` when it starts is transformed into a *make variable* with the same name and value. We can then access these variables using the syntax `${VARIABLE_NAME}` in our makefile.

```makefile
# 因为 make 会读取环境变量, 所以 dsn 也可以移到环境变量中设置
dsn = postgres://greenlight:pa55word@localhost/greenlight?sslmode=disable

up:
    @echo 'Running up migrations...'
    migrate -path ./migrations -database ${dsn} up
```

### 传递参数

可以用 `make hello name=tifa` 设置 `name` 的值

```makefile
hello:
    @echo Hello ${name}!

migration:
    @echo 'Creating migration files for ${name}...'
    migrate create -seq -ext=.sql -dir=./migrations ${name}
```

### 命名空间

As your makefile continues to grow, you might want to start *namespacing* your target names to provide some differentiation between rules and help organize the file. For example, in a large makefile rather than having the target name `up` it would be clearer to give it the name `db/migrations/up` instead. I recommend using the `/` character as a namespace separator, rather than a period, hyphen or the `:` character.

```makefile
run/api:
    go run ./cmd/api -db-dsn='$(dsn)' -cors-trusted-origins='$(trusted_origins)' -limiter-enabled=true

db/psql:
    docker exec -it postgres psql -U postgres -d greenlight

db/migrations/new:
    @echo 'Creating migration files for ${name}...'
    migrate create -seq -ext=.sql -dir=./migrations ${name}

db/migrations/up:
    @echo 'Running up migrations...'
    migrate -path ./migrations -database ${dsn} up
```

### 前置条件

When you specify a prerequisite target for a rule, the corresponding commands for the prerequisite targets will be run *before* executing the actual target commands. Let’s leverage this functionality to ask the user for *confirmation to continue* before executing our `db/migrations/up` rule. To do this, we’ll create a new `confirm` target which asks the user `Are you sure? [y/N]` and exits with an error if they do not enter `y`. Then we’ll use this new `confirm` target as a prerequisite for `db/migrations/up`.

注意一下:

- `echo -n` 表示不打印末尾的换行符
- `read ans` 表示读取一行输入存入 `ans` 变量
- `[ 1 = 1 ]` 用来测试条件, 这和 `test 1 = 1` 是同一个命令
- `${ans:-N}` 能取出 `ans` 变量的值, 如果变量为空则使用默认值 `N`，  
  两个表达式 `${var:-default}` 和 `${var:=default}` 的结果相同，但变量为空时，`${var:=default}` 还会把默认值写入空变量
- 在 Makefile 中要用 `$$` 表示字符 `$`
- 另外 `test 1 = 0 && echo ok` 因条件为 false 所以不会执行 echo 命令

```bash
confirm:
    @echo -n 'Are you sure? [y/N]' && read ans && [ $${ans:-N} = y ]
```

Essentially, what happens here is that we ask the user `Are you sure? [y/N]` and then read the response. We then use the code `[ $${ans:-N} = y ]` to evaluate the response — this will return `true` if the user enters `y` and `false` if they enter anything else. If a command in a makefile returns `false`, then `make` will stop running the rule and exit with an error message — essentially stopping the rule in its tracks.

```makefile
confirm:
    @echo -n 'Are you sure? [y/N]' && read ans && [ $${ans:-N} = y ]

# 先确认一下, 再执行 up 操作
db/migrations/up: confirm
    @echo 'Running up migrations...'
    migrate -path ./migrations -database ${dsn} up
```

### 显示帮助信息

- 思路是在每条 rule 前加一行 `## target: comment`，然后用 `sed` 命令从 Makefile 文件中提取帮助信息
- 注意把 `help` 作为第一条 rule 有特殊用意，因为单独输入 `make` 命令时，会默认执行第一条 rule

```makefile
## help: print this help message
help:
    @echo 'Usage:'
    @sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

## run/api: run the cmd/api application
run/api:
    go run ./cmd/api -db-dsn='$(dsn)' -cors-trusted-origins='$(trusted_origins)' -limiter-enabled=true
```

### Phony targets

In this chapter we’ve been using `make` to execute *actions*, but another (and arguably, the primary) purpose of `make` is to help create files on disk where *the name of a target is the name of a file* being created by the rule. 

If you’re using `make` primarily to execute actions, like we are, then this can cause a problem if there is a file in your project directory *with the same path* as a target name. If you want, you can demonstrate this problem by creating a file called `./run/api` in the root of your project directory, like so:

```bash
$ mkdir run && touch run/api
$ make run/api
make: 'run/api' is up to date. 
```

Because we already have a file on disk at `./run/api`, the `make` tool considers this rule to have already been executed and so returns the message that we see above without taking any further action. To work around this, we can declare our makefile targets to be [phony targets](https://www.gnu.org/software/make/manual/html_node/Phony-Targets.html#Phony-Targets):

> A phony target is one that is not really the name of a file; rather it is just a name for a rule to be executed.

To declare a target as phony, you can make it prerequisite of the special [`.PHONY`](https://www.gnu.org/software/make/manual/html_node/Special-Targets.html#Special-Targets) target. The syntax looks like this:

```makefile
.PHONY: help confirm run/api
```

You might think that it’s only necessary to declare targets as phony if you have a conflicting file name, but in practice *not declaring a target as phony when it actually is* can lead to bugs or confusing behavior. For example, imagine if in the future someone unknowingly creates a file called `confirm` in the root of the project directory. This would mean that our `confirm` rule is never executed, which in turn would lead to dangerous or destructive rules being executed without confirmation.

### 管理环境变量

Using the `make run/api` command to run our API application opens up an opportunity to tweak our command-line flags, and remove the default value for our database DSN from the `main.go` file. Like so:

```go
// 默认值改成 ""，不通过 Go 读取环境变量
// flag.StringVar(&cfg.db.dsn, "db-dsn", os.Getenv("GREENLIGHT_DB_DSN"), "PostgreSQL DSN")
flag.StringVar(&cfg.db.dsn, "db-dsn", "", "PostgreSQL DSN")
```

Instead, we can update our makefile so that the DSN value from the `GREENLIGHT_DB_DSN` environment variable is passed in as part of the rule. If you’re following along, please go ahead and update the `run/api` rule as follows:

```go
run/api:
    go run ./cmd/api -db-dsn=${GREENLIGHT_DB_DSN}
```

This is a small change but a really nice one, because it means that the *default configuration values for our application no longer change depending on the operating environment*. The command-line flag values passed at runtime are the *sole* mechanism for configuring our application settings, and there are still no secrets hard-coded in our project files.

#### ➤ 使用 .envrc 文件

If you like, you could also *remove* the `dsn` environment variable from your `$HOME/.profile` or `$HOME/.bashrc` files, and store it in a `.envrc` file in the root of your project directory instead.

```bash
export dsn='postgres://greenlight:pa55word@localhost/greenlight?sslmode=disable'
export trusted_origins='http://localhost:8000 http://localhost:9000'
```

You can then use a tool like [direnv](https://direnv.net/) to automatically load the variables from the `.envrc` file into your current shell, or alternatively, you can add an `include` command at the top of your `Makefile` to load them instead. Like so:

```makefile
# 读取此文件中的环境变量
include .envrc
```

This approach is particularly convenient in projects where you need to make frequent changes to your environment variables, because it means that you can just edit the `.envrc` file without needing to reboot your computer or run `source` after each change. Another nice benefit of this approach is that it provides a degree of separation between variables if you’re working on multiple projects on the same machine.

> **Important:** If you use this approach and your `.envrc` file contains any secrets, you must be careful to not commit the file into a version control system (like Git or Mercurial).
>
> echo '.envrc' >> .gitignore

## Building

### Quality Controlling Code

In this chapter we’re going to focus on adding an `audit` rule to our `Makefile` to check, test and tidy up our codebase automatically. In particular, the rule will:

- Use the `go mod tidy` command to prune any unused dependencies from the `go.mod` and `go.sum` files, and add any missing dependencies.

- Use the `go mod verify` command to check that the dependencies on your computer (located in your module cache located at `$GOPATH/pkg/mod`) haven’t been changed since they were downloaded and that they match the cryptographic hashes in your `go.sum` file. Running this helps ensure that the dependencies being used are the exact ones that you expect.
- Use the `go fmt ./...` command to format all `.go` files in the project directory, according to the Go standard. This will reformat files ‘in place’ and output the names of any changed files.
- Use the `go vet ./...` command to check all `.go` files in the project directory. The [`go vet`](https://golang.org/cmd/vet/) tool runs a variety of *analyzers* which carry out static analysis of your code and warn you about things which might be wrong but won’t be picked up by the compiler — such as unreachable code, unnecessary assignments, and badly-formed build tags.
- Use the `go test -race -vet=off ./...` command to run all tests in the project directory. By default, `go test` automatically executes a small subset of the `go vet` checks before running any tests, so to avoid duplication we’ll use the `-vet=off` flag to turn this off. The `-race` flag enables Go’s [race detector](https://golang.org/doc/articles/race_detector.html), which can help pick up certain classes of race conditions while tests are running.
- Use the third-party [`staticcheck`](https://staticcheck.io/) tool to carry out some [additional static analysis checks](https://staticcheck.io/docs/checks).

If you’re following along, you’ll need to install the `staticcheck` tool on your machine at this point. The simplest way to do this is by running the `go install` command. This will install the latest version of Staticcheck to `$GOPATH/bin`:

```bash
go install honnef.co/go/tools/cmd/staticcheck@latest
which staticcheck
```

Once that’s installed, let’s go ahead and create a new `audit` rule in our makefile.

```makefile
## audit: tidy dependencies and format, vet and test all code
audit:
    @echo 'Tidying and verifying module dependencies...'
    go mod tidy
    go mod verify
    @echo 'Formatting code...'
    go fmt ./...
    @echo 'Vetting code'
    go vet ./...
    staticcheck ./...
    @echo 'Running tests...'
    go test -race -vet=off ./...
```

Now that’s done, all you need to do is type `make audit` to run these checks before you commit any code changes into your version control system or build any binaries.

### Module Proxies

#### ➤ 使用 module proxy 解决什么问题?

One of the risks of using third-party packages in your Go code is that the package repository may cease to be available. For example, the `httprouter` package plays a central part in our application, and if the author ever decided to delete it from GitHub it would cause us quite a headache to scramble and replace it with an alternative.

Go supports *module proxies* (also known as *module mirrors*) by default. These are services which mirror source code from the original, authoritative, repositories (such as those hosted on GitHub, GitLab or BitBucket).

Go ahead and run the `go env` command on your machine to print out the settings for your Go operating environment.

```bash
$ go env
GOPATH='/home/kang/go'
GOPROXY='https://goproxy.cn,direct'
...
```

The important thing to look at here is the `GOPROXY` setting, which contains a comma-separated list of module mirrors. By default it has the following value:

```bash
GOPROXY="https://proxy.golang.org,direct"
```

#### ➤ 粗略的工作流程

The URL `https://proxy.golang.org` that we see here points to a *module mirror* maintained by the Go team at Google, containing copies of the source code from tens of thousands of open-source Go packages. Whenever you fetch a package using the `go` command — either with `go get` or one of the `go mod *` commands — it will first attempt to retrieve the source code from this mirror.

If the mirror already has a stored copy of the source code for the required package and version number, then it will return this code immediately in a zip file. Otherwise, if it’s not already stored, then the mirror will attempt to fetch the code from the authoritative repository, proxy it onwards to you, and store it for future use.

If the mirror can’t fetch the code at all, then it will return an error response and the `go` tool will fall back to fetching a copy directly from the authoritative repository (thanks to the `direct` directive in the `GOPROXY` setting).

#### ➤ 三个好处

Using a module mirror as the first fetch location has a few benefits:

- The `https://proxy.golang.org` module mirror *typically* stores packages long-term, thereby providing a degree of protection in case the original repository disappears from the internet.
- It’s not possible to override or delete a package once it’s stored in the `https://proxy.golang.org` module mirror. This can help prevent any bugs or problems which might arise if a package author (or an attacker) releases an edited version of the package *with the same version number*.
- Fetching modules from the `https://proxy.golang.org` mirror can be [much faster](https://twitter.com/sajma/status/1155006281263923201?s=21) than getting them from the authoritative repositories.

#### ➤ 可配置多个 module proxy

For example, if you wanted to switch to using `https://goproxy.io` as the primary mirror, then fall back to using `https://proxy.golang.org` as a secondary mirror, then fall back to a direct fetch, you could update your `GOPROXY` setting like so:

```bash
export GOPROXY=https://goproxy.io,https://proxy.golang.org,direct
```

### and Vendoring

#### ➤ 用不了或不想用 module proxy，总之想把依赖模块保存到项目中

Go’s module mirror functionality is great, and I recommend using it. But it isn’t a silver bullet for all developers and all projects. For example, perhaps you don’t want to use a module mirror provided by Google or another third-party, but you also don’t want the overhead of hosting your own mirror. Or maybe you need to routinely work in an environment without network access. In those scenarios you probably still want to mitigate the risk of a disappearing dependency, but using a module mirror isn’t possible or appealing. Additionally, if you need to come back to a ‘cold’ codebase in 5 or 10 years’ time, will the `proxy.golang.org` module mirror still be available? Hopefully it will — but it’s hard to say for sure.

So, for these reasons, it can still be sensible to *vendor* your project dependencies using the `go mod vendor` command. Vendoring dependencies in this way basically stores a complete copy of the source code for third-party packages in a `vendor` folder in your project.

#### ➤ 使用 go mod vendor 命令

Let's add a `vendor` rule to our Makefile:

```makefile
## vendor: tidy and vendor dependencies
vendor:
    @echo 'Tidying and verifying module dependencies...'
    go mod tidy
    go mod verify
    @echo 'Vendoring dependencies...'
    go mod vendor
```

Let’s quickly step through what will happen when we run `make vendor`:

- The `go mod tidy` command will make sure the `go.mod` and `go.sum` files list all the necessary dependencies for our project (and no unnecessary ones).
- The `go mod verify` command will verify that the dependencies stored in your module cache (located on your machine at `$GOPATH/pkg/mod`) match the cryptographic hashes in the `go.sum` file.
- The `go mod vendor` command will then copy the necessary source code from your module cache into a new `vendor` directory in your project root.

Let’s try this out and run the new `vendor` rule like so:

```bash
make vendor
```

Once that’s completed, you should see that a new `vendor` directory has been created containing copies of all the source code along with a `modules.txt` file. Now, when you run a command such as `go run`, `go test` or `go build`, the `go` tool will recognize the presence of a `vendor` folder and *the dependency code in the vendor folder* will be used — rather than the code in the module cache on your local machine.

> **Note:** If you want to confirm that it’s really the vendored dependencies being used, you can run `go clean -modcache` to remove *everything* from your local module cache. When you run the API again, you should find that it still starts up correctly *without* needing to re-fetch the dependencies from the Go module mirror.

#### ➤ 依赖模块也会成为项目代码的一部分

Because all the dependency source code is now stored in your project repository itself, it’s easy to check it into Git (or an alternative version control system) alongside the rest of your code. This is reassuring because it gives you *complete ownership* of all the code used to build and run your applications, kept under version control.

The downside of this, of course, is that it adds size and bloat to your project repository. This is of particular concern in projects that have a lot of dependencies and the repository will be cloned *a lot*, such as projects where a CI/CD system clones the repository with each new commit.

#### ➤ 这个 vendor/modules.txt 的作用

This `vendor/modules.txt` file is essentially a *manifest* of the vendored packages and their version numbers. When vendoring is being used, the `go` tool will check that the module version numbers in `modules.txt` are consistent with the version numbers in the `go.mod` file. If there’s any inconsistency, then the `go` tool will report an error.

> **Note:** It’s important to point out that there’s no easy way to verify that the *checksums of the vendored dependencies* match the checksums in the `go.sum` file. Or, in other words, there’s no equivalent to `go mod verify` which works *directly* on the contents of the `vendor` folder.
>
> To mitigate that, it’s a good idea to run *both* `go mod verify` and `go mod vendor` regularly. Using `go mod verify` will verify that the dependencies in your module cache match the `go.sum` file, and `go mod vendor` will copy those same dependencies from the module cache into your `vendor` folder.

#### ➤ 不要修改 vendor 目录中的代码

Lastly, you should avoid making any changes to the code in the `vendor` directory. Doing so can potentially cause confusion (because the code would no longer be consistent with the original version of the source code) and — besides — running `go mod vendor` will overwrite any changes you make each time you run it. If you need to change the code for a dependency, it’s far better to fork it and import the forked version instead.

#### ➤ 添加新依赖后，要重新执行 vendor

In the next section of the book we’re going to deploy our API application to the internet with [Caddy](https://caddyserver.com/) as a reverse-proxy in-front of it. This means that, as far as our API is concerned, all the requests it receives will be coming from a single IP address (the one running the Caddy instance). In turn, that will cause problems for our rate limiter middleware which limits access based on IP address.

Fortunately, like most other reverse proxies, Caddy adds an [`X-Forwarded-For`](https://caddyserver.com/docs/caddyfile/directives/reverse_proxy#headers) header to each request. This header will contain the *real IP address* for the client.

Although we could write the logic to check for the presence of an `X-Forwarded-For` header and handle it ourselves, I recommend using the [`realip`](https://github.com/tomasen/realip) package to help with this. This package retrieves the client IP address from any `X-Forwarded-For` or `X-Real-IP` headers, falling back to use `r.RemoteAddr` if neither of them are present.

If you’re following along, go ahead and install the latest version of `realip` using the `go get` command:

```bash
go get github.com/tomasen/realip@latest
```

Then open up the `cmd/api/middleware.go` file and update the `rateLimit()` middleware to use this package like so:

```go
// ip, _, err := net.SplitHostPort(r.RemoteAddr)
ip := realip.FromRequest(r)
```

If you try to run the API application again now, you should receive an error message similar to this:

```bash
$ make run/api
go: inconsistent vendoring in /mnt/d/project/greenlight:
        github.com/tomasen/realip: is explicitly required in go.mod, but not marked as explicit in vendor/modules.txt
```

Essentially what’s happening here is that Go is looking for the `github.com/tomasen/realip` package in our `vendor` directory, but at the moment that package doesn’t exist in there. To solve this, you’ll need to manually run either the `make vendor` or `make audit` commands. Once that’s done, everything should work correctly again.

### Building Binaries

To build a binary we need to use the `go build` command. As a simple example, usage looks like this:

```bash
go build -o=./bin/api ./cmd/api
```

When we run this command, `go build` will *compile* the `cmd/api` package (and any dependent packages) into files containing machine code, and then *link* these together to an form executable binary. In the command above, the executable binary will be output to `./bin/api`. For convenience, let’s add a new `build/api` rule to our makefile which runs this command, like so:

```makefile
## build/api: build the cmd/api application
build/api:
    @echo 'Building cmd/api...'
    go build -o ./bin/api ./cmd/api
```

#### ➤ Reducing binary size

If you take a closer look at the executable binary you’ll see that it weighs in at 12MB. It’s possible to reduce the binary size by around 25% by instructing the Go linker to strip the [DWARF](https://golang.org/pkg/debug/dwarf/) debugging information and [symbol table](https://medium.com/a-journey-with-go/go-how-to-take-advantage-of-the-symbols-table-360dd52269e5) from the binary. We can do this as part of the `go build` command by using the *linker flag* `-ldflags="-s"` as follows:

```makefile
build/api:
    go build -ldflags='-s' -o ./bin/api ./cmd/api
```

If you run `make build/api` again, you should now find that the size of the binary is reduced to under 8 MB. It’s important to be aware that stripping the DWARF information and symbol table will make it harder to debug an executable using a tool like [Delve](https://github.com/go-delve/delve) or [gdb](https://www.gnu.org/software/gdb/). But, generally, it’s not often that you’ll need to do this — and there’s even an [open proposal](https://github.com/golang/go/issues/26074) from Rob Pike to make omitting DWARF information the default behavior of the linker in the future.

#### ➤ Cross-compilation

By default, the `go build` command will output a binary suitable for use on your *local machine’s operating system and architecture*. But it also supports cross-compilation, so you can generate a binary suitable for use on a different machine. This is particularly useful if you’re developing on one operating system and deploying on another.

To see a list of all the operating system/architecture combinations that Go supports, you can run this command:

```bash
$ go tool dist list
android/arm64
ios/arm64
linux/amd64
linux/arm64
windows/amd64
windows/arm64
```

And you can specify the operating system and architecture that you want to create the binary for by setting `GOOS` and `GOARCH` environment variables when running `go build`. For example:

```bash
$ GOOS=linux GOARCH=amd64 go build {args}
```

Let’s update our `make build/api` rule so that it creates two binaries — one for use on your local machine, and another for deploying to the Ubuntu Linux server.

```makefile
build/api:
    go build -ldflags='-s' -o ./bin/api ./cmd/api
    GOOS=linux GOARCH=amd64 go build -ldflags='-s' -o=./bin/linux_amd64/api ./cmd/api
```

As a general rule, you probably don’t want to commit your Go binaries into version control alongside your source code as they will significantly inflate the size of your repository. Let’s add an additional rule to the `.gitignore` file:

```bash
echo 'bin/' >> .gitignore
```

#### ➤ Build caching

It’s important to note that the `go build` command caches build output in the Go *build cache*. This cached output will be reused again in future builds where appropriate, which can significantly speed up the overall build time for your application.

If you’re not sure where your build cache is, you can check by running the `go env GOCACHE` command:

```bash
$ go env GOCACHE
/home/kang/.cache/go-build
```

You should also be aware that the build cache does not automatically detect any changes to C libraries that your code imports with [`cgo`](https://golang.org/cmd/cgo/). So, if you’ve changed a C library since the last build, you’ll need to use the `-a` flag to force all packages to be rebuilt when running `go build`. Alternatively, you could use `go clean` to purge the cache:

```bash
$ go build -a -o=/bin/foo ./cmd/foo        # Force all packages to be rebuilt
$ go clean -cache                          # Remove everything from the build cache
```

### Version Numbers

#### ➤ 支持 -version 查看版本号

Let’s start by updating our application so that we can easily check the version number by running the binary with a `-version` command-line flag. We need to define a boolean `version` command-line flag, check for this flag on startup, and then print out the version number and exit the application if necessary.

```go
func main() {
    displayVersion := flag.Bool("version", false, "Display version and exit")
    
    flag.Parse()
    
    if *displayVersion {
        fmt.Printf("Version:\t%s\n", version)
        os.Exit(0)
    }
}
```

#### ➤ 使用构建信息中的 vcs commit hash 作为版本号

Since version 1.18, Go now embeds *version control* information in your executable binaries when you run `go build` on a `main` package that is tracked with Git, Mercurial, Fossil, or Bazaar. There are two ways to access this version control information — either by using the `go version -m` command on your binary, or from within your application code itself by calling [`debug.ReadBuildInfo()`](https://pkg.go.dev/runtime/debug#ReadBuildInfo).

```bash
$ go version -m ./bin/api
./bin/api: go1.22.1
        dep     github.com/lib/pq       v1.10.9 h1:YXG7RB+JIjhP29X+OtkiDnYaXQwpS4JEWq7dtCCRUEw=
        dep     golang.org/x/crypto     v0.27.0 h1:GXm2NjJrPaiv/h1tb2UH8QfgC/hOf/+z0p6PT8o1w7A=
        dep     golang.org/x/time       v0.6.0  h1:eTDhh4ZXt5Qf0augr54TN6suAUudPcawVZeIAPU7D4U=

        build   GOARCH=amd64
        build   GOOS=linux

        build   vcs=git
        build   vcs.revision=e76eafc92e9ddaf584a73446b28dc0d6ba5669cd
        build   vcs.time=2024-09-09T06:20:24Z
        build   vcs.modified=true
```

We can see the version of Go that it was built with (`go1.22.1` in my case), the module dependencies, and information about the build settings — including the linker flags used and the OS and architecture it was built for.

- `vcs=git` tells us that the version control system being used is Git.
- `vcs.revision` is the hash for the latest Git commit.
- `vcs.time` is the time that this commit was made.

`vcs.modified` tells us whether the code tracked by the Git repository has been modified since the commit was made. A value of `false` indicates that the code has not been modified, meaning that the binary was built using the *exact code from the `vcs.revision` commit*. A value of `true` indicates that the version control repository was ‘dirty’ when the binary was built — and the code used to build the binary may not be the exact code from the `vcs.revision` commit.

As I mentioned briefly above, all the information that you see in the `go version -m` output is also available to you at runtime. Let’s leverage this and adapt our `main.go` file so that the `version` value is set to the Git commit hash, rather than the hardcoded constant `"1.0.0"`.

To assist with this, we’ll create a small `internal/vcs` package which generates a version number for our application based on the commit hash from `vcs.revision` plus an optional `-dirty` suffix if `vcs.modified=true`.

```bash
mkdir internal/vcs
touch internal/vcs/vcs.go
```

Code:

```go
func Version() string {
    var time string
    var revision string
    var modified bool

    bi, ok := debug.ReadBuildInfo()
    if ok {
        // 遍历键值对, key 就是 go version -m ./bin/api 所显示的 key
        for _, s := range bi.Settings {
            switch s.Key {
            case "vcs.time":
                time = s.Value
            case "vcs.revision":
                revision = s.Value
            case "vcs.modified":
                if s.Value == "true" {
                    modified = true
                }
            }
        }
    }
    if modified {
        return fmt.Sprintf("%s  %s  dirty", time, revision)
    }
    return fmt.Sprintf("%s  %s", time, revision)
}
```

Let’s head back to our `main.go` file and update it to set the version number using this new `vcs.Version()` function:

```go
var (
    version = vcs.Version()
)
```

Alright, let’s try this out. Go ahead and rebuild the binary again.

```bash
$ make build/api
$ ./bin/api -version
Version:        2024-09-09T17:02:59Z  6add593ff4bf4d24b4a9ecfa8f809879d536bfff  dirty
```

#### ➤ 现在确定 go 应用是哪个版本的代码就很轻松了

Our application version number now aligns with the commit history in our Git repository, meaning that it’s easy for us to identify exactly what code a particular binary contains or a running application is using. All we need to do is run the binary with the `-version` flag, or call the healthcheck endpoint, and then cross-reference the version number against the Git repository history.

> **Important:** Version control information is only embedded by default when you run `go build`. It is never embedded when you use `go run`, and is only embedded when using `go test` if you use the `-buildvcs=true` flag. You can see this behavior in action if you use `go run` in conjunction with our `-version` flag; no version control information is embedded and the `version` string remains blank.
>
> ```bash
> $ go run ./cmd/api -version
> Version:
> ```

#### ➤ 也可以用 linker flag 实现，

Prior to Go 1.18 the idiomatic way to manage version numbers automatically was to ‘burn-in’ the version number when building the binary using the `-X` linker flag. Using `debug.ReadBuildInfo()` is now the preferred method, but the old approach can still be useful if you need to set the version number to something that isn’t available via `debug.ReadBuildInfo()`.

For example, if you wanted to set the version number to the value of a `VERSION` environment variable *on the machine building the binary*, you could use the `-X` linker flag to ‘burn-in’ this value to the `main.version` variable. Like so:

```makefile
# Note: "ld" stands for linker. Linkers in Linux were originally called loaders.
build/api:
    @echo 'Building cmd/api...'
    go build -ldflags='-s -X main.version=${VERSION}' -o=./bin/api ./cmd/api
    GOOS=linux GOARCH=amd64 go build -ldflags='-s -X main.version=${VERSION}' -o=./bin/linux_amd64/api ./cmd/api
```

## Deployment and Hosting

We’ll run everything on a single Ubuntu Linux server. Our stack will consist of a PostgreSQL database and the executable binary for our Greenlight API. But in addition to this, we’ll also run [Caddy](https://caddyserver.com/) as a *reverse proxy* in front of the Greenlight API.

Using Caddy has a couple of benefits. It will automatically handle and terminate HTTPS connections for us — including automatically generating and managing TLS certificates via [Let’s Encrypt](https://letsencrypt.org/) and [ZeroSSL](https://zerossl.com/) — and we can also use Caddy to easily restrict internet access to our metrics endpoint.

### Digital Ocean

1. 首次注册有 5 刀的预付款，然后账户里就有 5 刀余额了
2. Digital Ocean 中的 VPS/VM/cloud server 叫做 droplet 水滴

#### ➤ 创建和添加 SSH key

In order to log in to droplets in your Digital Ocean account you’ll need a *SSH keypair*.

```bash
# 备注的格式可以用 username@hostname
ssh-keygen -t rsa -b 4096 -C "greenlight@example.com" -f $HOME/.ssh/id_rsa_greenlight
```

Head back to your Digital Ocean control panel and navigate to the **Settings › Security** tab. Click the **Add SSH Key** button, then in the popup window that appears paste in the text contents from your `$HOME/.ssh/id_rsa_greenlight.pub` public key file.

#### ➤ 创建 droplet

Now that you have a valid SSH key added to your account, it’s time to actually create a droplet. There are a couple of ways that you can do this. It’s possible to do so programmatically via the [Digital Ocean API](https://developers.digitalocean.com/documentation/v2/) or using the official [command-line tool](https://www.digitalocean.com/community/tutorials/how-to-use-doctl-the-official-digitalocean-command-line-client), and if you need to create or manage a lot of servers then I recommend using these.

Or alternatively, it’s possible to create a droplet manually via your control panel on the Digital Ocean website. This is the approach we’ll take in this book, partly because it’s simple enough to do as a one-off, and partly because it helps give overview of the available droplet settings if you haven’t used Digital Ocean before.

Go ahead and click the green **Create** button in the top right corner and select **Droplets** from the dropdown menu:

1. 选择地区，没有较近的 Hong Kong 那就选 San Francisco
2. 选择最新的 Ubuntu LTS 版本，有 5 年的支持时间
3. 选择规格，可以选最便宜的 4 刀每月，每月 500GB 出口流量 ( 数据进入服务器是免费的，但出来要计费 )
4. 认证方法选择 SSH Key，免费的指标监控功能勾选上，然后填个容易识别的主机名如 greenlight-production
5. 等服务器启动后，把 IP 地址复制下来，然后用 SSH 连上去，例如 ssh root@1.2.3.4

> 首次连接服务器会有警告 "WARNING: REMOTE HOST IDENTIFICATION HAS CHANGED!"
>
> 这是因为如果服务器的 ID 无故发生改变，肯定要引起重视，连接目标可能是黑客伪装的服务器
>
> 但首次连接必定会触发此警告，因为之前记录的 ID 为空
>
> 当你输入 yes 后，SSH 客户端会记住这个 hostname/IP 对应的 ID

### Server Configuration

Now that our Ubuntu Linux droplet has been successfully commissioned, we need to do some housekeeping to secure the server and get it ready-to-use. Rather than do this manually, in this chapter we’re going to create a reusable script to automate these setup tasks. At a high level we want our setup script to do the following things:

- Update all packages on the server, including applying security updates.
- Set the server timezone (in my case I’ll set it to `Europe/Berlin`) and install support for all [locales](https://help.ubuntu.com/community/Locale).
- Create a `greenlight` user on the server, which we can use for day-to-day maintenance and for running our API application (rather than using the all-powerful `root` user account). We should also add the `greenlight` user to the `sudo` group, so that it can perform actions as `root` if necessary.
- Copy the `root` user’s `$HOME/.ssh` directory into the `greenlight` users home directory. This will enable us to authenticate as the `greenlight` user using the same SSH key pair that we used to authenticate as the `root` user. We should also force the `greenlight` user to set a new password the first time they log in.
- Configure firewall settings to only permit traffic on ports `22` (SSH), `80` (HTTP) and `443` (HTTPS). We’ll also install [fail2ban](https://www.fail2ban.org/wiki/index.php/Main_Page) to automatically temporarily ban an IP address if it makes too many failed SSH login attempts.
- Install PostgreSQL. We’ll also create the `greenlight` database and user, and create a system-wide `GREENLIGHT_DB_DSN` environment variable for connecting to the database.
- Install the `migrate` tool, using the [pre-built binaries](https://github.com/golang-migrate/migrate/releases) from GitHub.
- Install Caddy by following the [official installation instructions](https://caddyserver.com/docs/install#debian-ubuntu-raspbian) for Ubuntu.
- Reboot the droplet.

> #### ➤ root group vs sudoer 有什么区别?
>
> The 'root group' as in what you would specify in /etc/group is about unix permissions. Every file has user, group, and 'other' permissions. If a file is set such that users in group root can read it, then you can grant a user the ability to read that file by putting the user in group root. Of course then that user can read *every* file which has the read bit set for group root.
>
> The "sudo" command lets you execute commands with superuser privileges as long as your user id is in the `sudoers` file, giving you the necessary authorization. So, e.g. `sudo vi /etc/hosts` would allow you to edit the hosts file as if you were running as root. You don't even need the root password, just your own login password. And of course, `sudo su` would allow you to simply become root. 
>
> The sudoers file is about running commands with the effective ID of other users. You have more granular control over what commands each user can run, and as as whom. So if you want a user to only be able to run one specific command as root, then you would set that in the sudoers file. 
>
> Effectively, your administrators are root, plus everybody listed in the sudoers file. `sudo` is, fundamentally, a *usability* tool, not a security tool. Prior to the invention of the "sudo" command, if you wanted to perform administrative tasks, you had to login as root, either by getting a login prompt somehow, or with the `su` command. That's a bit of a hassle, and also doesn't let you give users partial administrative powers.

#### ➤ 自动化配置脚本

```bash
mkdir -p remote/setup
touch remote/setup/01.sh
```

代码:

```bash
#!/usr/bin/env bash
# 上面这行 shebang 比 /bin/bash 要好, 因为会从 PATH 搜索 bash 命令的位置而不是写死一个位置

set -u # 使用未定义的变量视为错误
set -e # 遇到错误退出脚本, 不再执行下一行命令

TIMEZONE=Asia/Shanghai
USERNAME=greenlight

# 让用户输入一个密码
read -p "Enter password for greenlight DB user: " DB_PASSWORD

# Force all output to be presented in en_US for the duration of this script.
export LC_ALL=en_US.UTF-8

# ==================================================================================== #
# SCRIPT LOGIC
# ==================================================================================== #

apt update

# 设置时区
timedatectl set-timezone ${TIMEZONE}
apt --yes install locales-all

# 创建用户并添加到 sudo group
useradd --create-home --shell "/bin/bash" --groups sudo "${USERNAME}"

# 要求用户首次登陆时设置密码
passwd --delete "${USERNAME}"
chage --lastday 0 "${USERNAME}"

# 把 root 的 SSH keys 复制到新用户, 所以用同一对秘钥能登陆这两个用户
rsync --archive --chown=${USERNAME}:${USERNAME} /root/.ssh /home/${USERNAME}

# 开启防火墙放开 SSH/HTTP/HTTPS 端口
ufw allow 22
ufw allow 80/tcp
ufw allow 443/tcp
ufw --force enable

# 安装 fail2ban
apt --yes install fail2ban

# 安装 migrate 的最新版
curl -L https://github.com/golang-migrate/migrate/releases/latest/download/migrate.linux-amd64.tar.gz | tar xvz
mv migrate /usr/local/bin/migrate
rm LICENSE README.md

# 安装 PostgreSQL
apt --yes install postgresql

# 初始化数据库: 创建db、安装扩展、创建用户
# Set up the greenlight DB and create a user account with the password entered earlier.
# 从 PostgreSQL 15 起默认没有建表权限, 需要改 db owner, 参考 https://stackoverflow.com/questions/67276391/
sudo -i -u postgres psql -c "CREATE DATABASE greenlight"
sudo -i -u postgres psql -d greenlight -c "CREATE EXTENSION IF NOT EXISTS citext"
sudo -i -u postgres psql -d greenlight -c "CREATE ROLE greenlight WITH LOGIN PASSWORD '${DB_PASSWORD}'"
sudo -i -u postgres psql -d greenlight -c "ALTER DATABASE greenlight OWNER TO greenlight"

# 添加一个系统级的环境变量, 方便连接数据库
echo "GREENLIGHT_DB_DSN='postgres://greenlight:${DB_PASSWORD}@localhost/greenlight'" >> /etc/environment

# 安装 Caddy
apt --yes install debian-keyring debian-archive-keyring apt-transport-https
curl -1sLf 'https://dl.cloudsmith.io/public/caddy/stable/gpg.key' | sudo gpg --dearmor -o /usr/share/keyrings/caddy-stable-archive-keyring.gpg
curl -1sLf 'https://dl.cloudsmith.io/public/caddy/stable/debian.deb.txt' | sudo tee /etc/apt/sources.list.d/caddy-stable.list
apt update
apt --yes install caddy

# 升级所有包. 选项 --force-confnew 表示如果配置文件有新版本, 那么用新版本
apt --yes -o Dpkg::Options::="--force-confnew" upgrade

echo "Script complete! Rebooting..."
reboot
```

#### ➤ 上传脚本，然后运行

```bash
# 把 setup 文件夹同步到服务器的 /root 目录下
rsync -rP --delete ./remote/setup digital-ocean:/root
```

> **Note:** In this `rsync` command the `-r` flag indicates that we want to copy the contents of `./remote/setup` recursively, the `-P` flag indicates that we want to display progress of the transfer, and the `--delete` flag indicates that we want to delete any extraneous files from destination directory on the droplet.

Now that a copy of our setup script is on the droplet, let’s use the `ssh` command to execute the script on the remote machine as the `root` user. We’ll use the `-t` flag to force *pseudo-terminal allocation*, which is useful when executing screen-based programs on a remote machine. 

```bash
ssh -t digital-ocean "bash /root/setup/01.sh"
```

验证方法:

1. 使用 greenlight 用户重连服务器，按提示设置密码
2. 使用 psql $GREENLIGHT_DB_DSN 验证 PostgreSQL 服务已运行
3. 打开浏览器然后输入服务器的 ip 地址，验证 Caddy 服务已运行

以后如果有新的初始化脚本，可以建个 02.sh 然后执行:

```bash
rsync -rP --delete ./remote/setup greenlight@1.2.3.4:~
ssh -t greenlight@1.2.3.4 "sudo bash /home/greenlight/setup/02.sh"
```

### Deployment

At a very high-level, our deployment process will consist of three actions:

1. Copying the application binary and SQL migration files to the droplet.
2. Executing the migrations against the PostgreSQL database on the droplet.
3. Starting the application binary as a *background service*.

Let’s begin by creating a new `make production/deploy/api` rule in our makefile:

```makefile
production_host_ip = 1.2.3.4

## production/connect: connect to the production server
production/connect:
    @echo 记得启动 ssh-agent 并添加 key 否则连不上
    ssh greenlight@${production_host_ip}

## production/deploy/api: deploy the api to production
production/deploy/api:
    rsync -P ./bin/linux_amd64/api greenlight@${production_host_ip}:~
    rsync -rP --delete ./migrations greenlight@${production_host_ip}:~
    ssh -t greenlight@${production_host_ip} 'migrate -path ~/migrations -database $$GREENLIGHT_DB_DSN up'
```

验证方法:

```bash
make production/connect   # 连接服务器
ls -R                     # 看下文件同步过去没
psql $GREENLIGHT_DB_DSN   # 进数据库
\dt                       # 看看表建好了没
```

试试运行 api:

```bash
sudo ufw allow 4000/tcp                                       # 放开 4000 端口
./api -port=4000 -db-dsn=$GREENLIGHT_DB_DSN -env=production   # 运行 api
curl --silent http://1.2.3.4:4000/v1/healthcheck | jq         # 试试访问接口
```

### Background Service 1

Now that we know our API executable works fine on our production droplet, the next step is to configure it to run as a *background service*, including starting up automatically when the droplet is rebooted. There are a few different tools we could use to do this, but in this book we will use [systemd](https://www.freedesktop.org/wiki/Software/systemd/) — a collection of tools for managing services that ships with Ubuntu (and many other Linux distributions).

In order to run our API application as a background service, the first thing we need to do is make a [unit file](https://www.digitalocean.com/community/tutorials/understanding-systemd-units-and-unit-files), which informs systemd how and when to run the service.

```bash
mkdir remote/production
touch remote/production/api.service
```

Unit file:

```bash
[Unit]
# Description is a human-readable name for the service.
Description=Greenlight API service

# Wait until PostgreSQL is running and the network is "up" before starting the service.
After=postgresql.service
After=network-online.target
Wants=network-online.target

# Configure service start rate limiting. If the service is (re)started more than 5 times
# in 600 seconds then don't permit it to start anymore.
StartLimitIntervalSec=600
StartLimitBurst=5

[Service]
# Execute the API binary as the greenlight user, loading the environment variables from
# /etc/environment and using the working directory /home/greenlight.
Type=exec
User=greenlight
Group=greenlight
EnvironmentFile=/etc/environment
WorkingDirectory=/home/greenlight
ExecStart=/home/greenlight/api -port=4000 -db-dsn=${GREENLIGHT_DB_DSN} -env=production

# Automatically restart the service after a 5-second wait if it exits with a non-zero
# exit code. If it restarts more than 5 times in 600 seconds, then the rate limit we
# configured above will be hit and it won't be restarted anymore.
Restart=on-failure
RestartSec=5

[Install]
# Start the service automatically at boot time (the 'multi-user.target' describes a boot
# state when the system will accept logins).
WantedBy=multi-user.target
```

The next step is to install this unit file on our droplet and start up the service.

1. To install the unit file, we need to copy it into the `/etc/systemd/system/` folder on our droplet. 
2. Then we need to run the `systemctl enable api` command on our droplet to make `systemd` aware of the new unit file and automatically enable the service when the droplet is rebooted.
3. Finally, we need to run `systemctl restart api` to start (or restart) the service.

```makefile
production/deploy/api:
    rsync -P ./bin/linux_amd64/api greenlight@${production_host_ip}:~
    rsync -rP --delete ./migrations greenlight@${production_host_ip}:~
    rsync -P ./remote/production/api.service greenlight@${production_host_ip}:~
    ssh -t greenlight@${production_host_ip} '\
        migrate -path ~/migrations -database $$GREENLIGHT_DB_DSN up \
        && sudo mv ~/api.service /etc/systemd/system/ \
        && sudo systemctl enable api \
        && sudo systemctl restart api \
    '
```

验证方法:

```bash
make production/connect                                 # 连接服务器
sudo systemctl status api                               # 查看服务状态
ps -U greenlight                                        # 看看这个用户有哪些进程
sudo reboot                                             # 重启服务器
curl --silent http://1.2.3.4:4000/v1/healthcheck | jq   # 验证 api 服务有没有自启
```

关闭之前打开的 4000 端口:

```
sudo ufw delete allow 4000/tcp
sudo ufw status
```

### Background Service 2

#### ➤ Listening on a restricted port

If you’re *not* planning to run your application behind a reverse proxy, and want to listen for requests directly on port 80 or 443, you’ll need to set up your unit file so that the service has the `CAP_NET_BIND_SERVICE` capability (which will allow it to bind to a restricted port). For example:

```bash
[Unit]
Description=Greenlight API service

After=postgresql.service
After=network-online.target
Wants=network-online.target

StartLimitIntervalSec=600
StartLimitBurst=5

[Service]
Type=exec
User=greenlight
Group=greenlight
# 加上这两行可以绑定 80 或 443 之类的端口
CapabilityBoundingSet=CAP_NET_BIND_SERVICE
AmbientCapabilities=CAP_NET_BIND_SERVICE
EnvironmentFile=/etc/environment
WorkingDirectory=/home/greenlight
ExecStart=/home/greenlight/api -port=80 -db-dsn=${GREENLIGHT_DB_DSN} -env=production

Restart=on-failure
RestartSec=5

[Install]
WantedBy=multi-user.target
```

#### ➤ Viewing logs

It’s possible to view the logs for your background service using the `journalctl` command, like so:

```bash
sudo journalctl -u api
```

The `journalctl` command is really powerful and offers a wide variety of parameters that you can use to filter your log messages and customize the formatting. [This article](https://www.digitalocean.com/community/tutorials/how-to-use-journalctl-to-view-and-manipulate-systemd-logs) provides a great introduction to the details of `journalctl` and is well worth a read.

#### ➤ Configuring the SMTP provider

Our unit file is currently set up to start our API with the following command:

```bash
ExecStart=/home/greenlight/api -port=4000 -db-dsn=${GREENLIGHT_DB_DSN} -env=production
```

It’s important to remember that apart from the `port`, `db-dsn` and `env` flags that we’re specifying here, our application will still be using the default values for the other settings which are hardcoded into the `cmd/api/main.go` file — including the SMTP credentials for your Mailtrap inbox. Under normal circumstances, you would want to set your production SMTP credentials as part of this command in the unit file too.

### Reverse Proxy

#### ➤ 入门

We’re now in the position where we have our Greenlight API application running as a background service on our droplet, and listening for HTTP requests on port `4000`. And we also have Caddy running as a background service and listening for HTTP requests on port `80`. So the next step in setting up our production environment is to configure Caddy to act as a *reverse proxy* and forward any HTTP requests that it receives onward to our API. The simplest way to configure Caddy is to create a *Caddyfile*:

```bash
touch remote/production/Caddyfile
```

And then add the following content:

```nginx
http://1.2.3.4 {
  reverse_proxy localhost:4000
}
```

This rule tells Caddy that we want it to listen for HTTP requests to `161.35.71.158` and then act a reverse proxy, forwarding the request to port `4000` on the local machine. ( If you plan on using Caddy in your own production projects, then I recommend reading through the [official Caddyfile documentation](https://caddyserver.com/docs/caddyfile), which is excellent. )

Let’s update our `production/deploy/api` rule:

```makefile
production/deploy/api:
    rsync -P ./bin/linux_amd64/api greenlight@${production_host_ip}:~
    rsync -rP --delete ./migrations greenlight@${production_host_ip}:~
    rsync -P ./remote/production/api.service greenlight@${production_host_ip}:~
    rsync -P ./remote/production/Caddyfile greenlight@${production_host_ip}:~
    ssh -t greenlight@${production_host_ip} '\
        migrate -path ~/migrations -database $$GREENLIGHT_DB_DSN up \
        && sudo mv ~/api.service /etc/systemd/system/ \
        && sudo systemctl enable api \
        && sudo systemctl restart api \
        && sudo mv ~/Caddyfile /etc/caddy/ \
        && sudo systemctl reload caddy \
    '
```

Go ahead and run the rule to deploy the Caddyfile into production. At this point you can visit `http://<your_droplet_ip>/v1/healthcheck` in a web browser, and you should find that the request is successfully forwarded on from Caddy to our API.

#### ➤ 限制敏感接口例如 /debug/vars 仅允许本地访问

Fortunately, it’s very easy to block access to this by adding a new [`respond`](https://caddyserver.com/docs/caddyfile/directives/respond) directive to our Caddyfile. With this new directive we’re instructing Caddy to send a `403 Forbidden` response for all requests which have a URL path beginning `/debug/`.

```bash
http://143.198.134.0 {
  respond /debug/* "Not Permitted" 403
  reverse_proxy localhost:4000
}
```

Although the metrics are no longer publicly accessible, you can still access them by connecting to your droplet via SSH and making a request to `http://localhost:4000/debug/vars`.

```bash
make production/connect
curl http://localhost:4000/debug/vars
```

Or alternatively, you can open a SSH tunnel to the droplet and view them using a web browser on your *local machine*.

```bash
# 把本地的 9999 端口转发/映射到远程主机上的 localhost:4000 端口
ssh -L :9999:localhost:4000 user@host
```

#### ➤ 使用域名

The first thing you’ll need to do is configure the DNS records for your domain name so that they contain an `A` record pointing to the IP address for your droplet. Once you’ve got the DNS record in place, the next task is to update the Caddyfile to use your domain name instead of your droplet’s IP address.

```bash
http://homura233.fun {
  respond /debug/* "Not Permitted" 403
  reverse_proxy localhost:4000
}
```

#### ➤ 使用 HTTPS

Now that we have a domain name set up we can utilize one of Caddy’s headline features: *automatic HTTPS*. Caddy will automatically handle provisioning and renewing TLS certificates for your domain via Let’s Encrypt or ZeroSSL (depending on availability), as well as redirecting all HTTP requests to HTTPS. It’s simple to set up, very robust, and saves you the overhead of needing to keep track of certificate renewals manually.

```bash
# 留个邮箱好在 tls 证书有问题时联系你
{
  email you@example.com
}

# 删掉了 http:// 前缀
homura233.fun {
  respond /debug/* "Not Permitted" 403
  reverse_proxy localhost:4000
}
```

#### ➤ 一种配置升级路线

Before launching a new service, it’s often useful to do a thought experiment and ask yourself: *What happens as traffic to the service increases? How would I manage it?*. Very roughly, the path looks something like this:

- Use single low-powered droplet running Caddy, PostgreSQL, and the Go application. This is what we currently have.
- ↓ Upgrade the droplet to have more CPU and/or memory as necessary.
- ↓ Move PostgreSQL to a separate droplet, or use a [managed database](https://www.digitalocean.com/products/managed-databases/).
- ↓ Upgrade droplets/managed databases to have more CPU and/or memory as necessary.
- If the droplet running the Go application is a bottleneck:
  - ↓ Profile and optimize your Go application code.
  - ↓ Run the Go application on multiple droplets, behind a load balancer.
- If the PostgreSQL database is a bottleneck:
  - ↓ Profile and optimize the database settings and database queries.
  - ↓ If appropriate, cache the results of expensive/frequent database queries.
  - ↓ If appropriate, move some operations to an in-memory database such as Redis.
  - ↓ Start using read-only database replicas for queries where possible.
  - ↓ Start sharding the database.