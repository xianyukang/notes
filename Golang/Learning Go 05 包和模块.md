## Table of Contents
  - [Module](#Module)
    - [Repositories, Modules, Packages](#Repositories-Modules-Packages)
    - [创建 go.mod 文件](#%E5%88%9B%E5%BB%BA-gomod-%E6%96%87%E4%BB%B6)
  - [Package](#Package)
    - [创建和导入包](#%E5%88%9B%E5%BB%BA%E5%92%8C%E5%AF%BC%E5%85%A5%E5%8C%85)
    - [首字母大写表示导出](#%E9%A6%96%E5%AD%97%E6%AF%8D%E5%A4%A7%E5%86%99%E8%A1%A8%E7%A4%BA%E5%AF%BC%E5%87%BA)
    - [包的命名建议](#%E5%8C%85%E7%9A%84%E5%91%BD%E5%90%8D%E5%BB%BA%E8%AE%AE)
    - [处理包名冲突](#%E5%A4%84%E7%90%86%E5%8C%85%E5%90%8D%E5%86%B2%E7%AA%81)
    - [类型别名](#%E7%B1%BB%E5%9E%8B%E5%88%AB%E5%90%8D)
    - [调整 API 时保证兼容性](#%E8%B0%83%E6%95%B4-API-%E6%97%B6%E4%BF%9D%E8%AF%81%E5%85%BC%E5%AE%B9%E6%80%A7)
    - [Package Comments and godoc](#Package-Comments-and-godoc)
    - [The internal Package](#The-internal-Package)
    - [Circular Dependencies](#Circular-Dependencies)
  - [Working with Modules](#Working-with-Modules)
    - [用 decimal 包做精确计算](#%E7%94%A8-decimal-%E5%8C%85%E5%81%9A%E7%B2%BE%E7%A1%AE%E8%AE%A1%E7%AE%97)
    - [下载依赖模块](#%E4%B8%8B%E8%BD%BD%E4%BE%9D%E8%B5%96%E6%A8%A1%E5%9D%97)
    - [升级依赖模块](#%E5%8D%87%E7%BA%A7%E4%BE%9D%E8%B5%96%E6%A8%A1%E5%9D%97)
    - [间接依赖的模块](#%E9%97%B4%E6%8E%A5%E4%BE%9D%E8%B5%96%E7%9A%84%E6%A8%A1%E5%9D%97)
    - [查看依赖树](#%E6%9F%A5%E7%9C%8B%E4%BE%9D%E8%B5%96%E6%A0%91)
    - [理解版本号](#%E7%90%86%E8%A7%A3%E7%89%88%E6%9C%AC%E5%8F%B7)
    - [发布一个模块就是扔到 VCS](#%E5%8F%91%E5%B8%83%E4%B8%80%E4%B8%AA%E6%A8%A1%E5%9D%97%E5%B0%B1%E6%98%AF%E6%89%94%E5%88%B0-VCS)
    - [会选择同一个模块的最新版本](#%E4%BC%9A%E9%80%89%E6%8B%A9%E5%90%8C%E4%B8%80%E4%B8%AA%E6%A8%A1%E5%9D%97%E7%9A%84%E6%9C%80%E6%96%B0%E7%89%88%E6%9C%AC)
    - [Semantic Import Versioning](#Semantic-Import-Versioning)
    - [Module Proxy Servers](#Module-Proxy-Servers)

## Module

### Repositories, Modules, Packages

Most modern programming languages have a system for organizing code into namespaces and libraries, and Go is no exception. Go introduces some new approaches to this old idea. In this chapter, we’ll learn about organizing code with packages and modules, how to import them, how to work with third-party libraries, and how to create libraries of your own.  

A `repository` is familiar to all developers. It is a place in a version control system where the source code for a project is stored. A `module` is the root of a Go library or application, stored in a repository. Modules consist of one or more `packages`, which give the module organization and structure.  

Before we can use code from packages outside of the standard library, we need to make sure that we have declared that our project is a module. 通过创建一个 `go.mod` 文件.

### 创建 go.mod 文件

#### ➤ 示例

A collection of Go source code becomes a module when there’s a valid `go.mod` file in its root directory. Rather than create this file manually, we use the subcommands of the go mod command to manage modules. The command `go mod init MODULE_PATH` creates the go.mod file.

```go
// 首先 module path 可以随意, 但要确保唯一性( 指不和其他用到的模块冲突 ), 通常使用模块所在的 github 仓库地址
module xianyukang.com/learn_go

// Next, the go.mod file specifies the minimum compatible version of Go.
go 1.21

// Finally, the require section lists the modules that your module depends on.
require (
    github.com/goccy/go-json v0.9.7
    golang.org/x/net v0.10.0
    gopkg.in/yaml.v3 v3.0.1
)
```

When download a dependency with `go get`, then the exact version of the dependency will be recorded in the `go.mod` file. Because the exact version is known, this makes it much easier to ensure reproducible builds across different machines and environments. The `go.mod` file also defines the module path. This is essentially the identifier that will be used as the root import path for the packages in your project. It’s good practice to make the module path unique to you and your project. A common convention in the Go community is to base it on a URL that you own.

#### ➤ module path

Recall that a [module path](https://go.dev/ref/mod#module-path) consists of three parts: 

- a repository root path (corresponding to the repository root directory), 
- a module subdirectory, and 
- a major version suffix (only for modules released at `v2` or higher).

For most modules, the module path is equal to the repository root path, so the module’s root directory is the repository’s root directory. Modules are sometimes defined in repository subdirectories.  For example, suppose the module `example.com/monorepo/foo/bar` is in the repository with root path `example.com/monorepo`. Its `go.mod` file is in the `foo/bar` subdirectory.

Every module has a globally unique identifier. This is not unique to Go. Java uses globally unique package declarations like com.companyname.projectname.library. In Go, we usually use the path to the module repository where the module is found. For example, Proteus, a module I wrote to simplify relational database access in Go, can be found at GitHub. It has a module path of `github.com/jonbodner/proteus`. The module path is case-sensitive. To reduce confusion, do not use uppercase letters within it. 

You can pick almost any string as your module path, but the important thing to focus on is uniqueness. If you’re creating a project which can be downloaded and used by other people and programs, then it’s good practice for your module path to equal the location that the code can be downloaded from. For instance, if your package is hosted at `https://github.com/foo/bar` then the module path for the project should be `github.com/foo/bar`.

## Package

### 创建和导入包

如何创建包?  一个文件夹就是一个包，所以新建一个文件夹，并让其中的 .go 文件使用同一个包名  
如何导入包?  使用 `import "greenlight.xianyukang.com/internal/data"` 导入一个包

包中 .go 文件的文件名不重要，你可以将包的所有代码保存在一个文件里，或者把它们拆分在多个文件中，无论哪种方式，*它都将成为同一个包的一部分*。如果变量、函数或类型的名称以大写字母开头，则认为它是导出的，可以从当前包之外的包访问它。如果不以大写字母开头则认为该名称是未导出的，只能在当前包中使用。

You must specify an import path when importing from anywhere besides the standard library.  
The `import path` is built by **appending the path to the package within the module to the module path**.

As a general rule, you should make the name of the package match the name of the directory that contains the package.  
比如 `math/rand` 是包的导入路径， 那么用 `rand` 作为包名.  
但是 `导入路径` 和 `包名` 是两个概念，比如 `github.com/xxx/aaa` 这个导入路径下的包名可以为 `bbb`.  

While you can use a relative path to import a dependent package within the same module, don’t do this. Absolute import paths clarify what you are importing and make it easier to refactor your code.   

### 首字母大写表示导出

Go’s import statement allows you to access exported constants, variables, functions, and types in another package.  How do you export an identifier in Go? Rather than use a special keyword, Go uses capitalization to determine if a package-level identifier is visible outside of the package where it is declared. An identifier whose name starts with an uppercase letter is exported.  

Anything you export is part of your package’s API. Before you export an identifier, be sure that you intend to expose it to clients. Document all exported identifiers and keep them backward-compatible unless you are intentionally making a major version change.

### 包的命名建议

Having the package name as part of the name used to refer to items in the package has some implications. The first is that *package names should be descriptive*. Rather than have a package called `util`, create a package name that describes the functionality provided by the package.  

For example, say you have two helper functions: one to extract all names from a string and another to format names properly. Don’t create two functions in a `util` package called `ExtractNames` and `FormatNames`. If you do, every time you use these functions, they will be referred to as `util.ExtractNames` and `util.FormatNames`, and that `util` package tells you nothing about what the functions do.  

It’s better to create one function called `Names` in a package called `extract` and a second function called `Names` in a package called `format`. It’s OK for these two functions to have the same name, because they will always be disambiguated by their package names. The first will be referred to as `extract.Names` when imported, and the second will be referred to as `format.Names`.  

Don’t name your function `ExtractNames` when it is in the `extract` package.   

### 处理包名冲突

Sometimes you might find yourself importing two packages whose names collide. For example, the standard library includes two packages for generating random numbers; one is cryptographically secure (`crypto/rand`) and the other is not (`math/rand`). The regular generator is fine when you aren’t generating random numbers for encryption, but you need to seed it with an unpredictable value. A common pattern is to seed a regular random number generator with a value from a cryptographic generator. 

In Go, both packages have the same name (`rand`). When that happens, you provide an alternate name for one package within the current file:  `import crand "crypto/rand"`

The package name `.` places all the exported identifiers in the imported package into the current package’s namespace; you don’t need a prefix to refer to them. This is discouraged because it makes your source code less clear as you no longer know whether something is defined in the current package or an imported one by simply looking at its name.

You can also use `_` as the package name. We’ll explore what this does when we talk about the `init` function.  
比如可以使用 `import _ "github.com/lib/pq"` 引入包的副作用, 比如注册 postgres 数据库驱动.

### 类型别名

If we want to allow users to access `Foo` by the name `Bar`, all we need to do is: `type Bar = Foo`. The alias can even be assigned to a variable of the original type without a type conversion. One important point to remember: *an alias is just another name for a type*. If you want to add new methods or change the fields in an aliased struct, you must add them to the original type. You can alias a type that’s defined in the same package as the original type or in a different package. You can even alias a type from another module.

### 调整 API 时保证兼容性

After using a module for a while, you might realize that its API is not ideal. You might want to rename some of the exported identifiers or move them to another package within your module. To avoid a backward-breaking change, *don’t remove the original identifiers*; provide an alternate name instead.  

With a function or method, this is easy. You declare a function or method that calls the original. For a constant, simply declare a new constant with the same type and value, but a different name. When you want to rename or move an exported type, you have to use an alias.

### Package Comments and godoc

Go has its own format for writing comments that are automatically converted into documentation. It’s called `godoc` format and it’s very simple. There are no special symbols in a godoc comment. They just follow a convention. Here are the rules:  

- 在函数、类型、或变量前添加注释
- Start the comment with two forward slashes (//) followed by the name of the item.
- Use a blank comment to break your comment into multiple paragraphs.

Comments before the package declaration create package-level comments. If you have lengthy comments for the package (such as the extensive formatting documentation in the fmt package), the convention is to put the comments in a file in your package called doc.go.

```go
// Package money provides various utilities to make it easy to manage money.
package money

import "github.com/shopspring/decimal"

// Money represents the combination of an amount of money and the currency the money is in.
// Notice that the comment starts with the name of the struct.
type Money struct {
    Value    decimal.Decimal
    Currency string
}

// Convert converts the value of one currency to another.
//
// It has two parameters: a Money instance with the value to convert,
// and a string that represents the currency to convert to. Convert returns ...
func Convert(from Money, to string) (Money, error) {
    return Money{}, nil
}
```

### The internal Package

Sometimes you want to share a function, type, or constant between packages in your module, but you don’t want to make it part of your API. Go supports this via the special internal package name.  

When you create a package called `internal`, the exported identifiers in that package and its subpackages are only accessible to the direct parent package of internal and the sibling packages of internal. This is useful because it prevents other codebases from importing and relying on the packages in our internal directory.

### Circular Dependencies

Two of the goals of Go are a fast compiler and easy to understand source code. To support this, Go does not allow you to have a circular dependency between packages. This means that if package A imports package B, directly or indirectly, package B cannot import package A, directly or indirectly.   

If you find yourself with a circular dependency, you have a few options. 

1. 把两个包合并成一个包
2. 把导致 package 循环依赖的函数或类型移动到一个新包

## Working with Modules

### 用 decimal 包做精确计算

Unlike many other compiled languages, Go compiles all code for your application into a single binary, whether it was code you wrote or code from third parties. You should never use floating point numbers when you need an exact representation of a decimal number. If you do need an exact representation, one good library is [the decimal module from ShopSpring](https://github.com/shopspring/decimal).

```go
import (
    "github.com/shopspring/decimal"
    "testing"
)

func TestDecimal(t *testing.T) {
    amount, err := decimal.NewFromString("100")
    if err != nil {
        t.Fatal(err)
    }
    percent, err := decimal.NewFromString("1")
    if err != nil {
        t.Fatal(err)
    }

    percent = percent.Div(decimal.NewFromInt(100))
    total := amount.Add(amount.Mul(percent)).Round(2)
    t.Logf("%s 元加上 %s 的税等于 %s 元\n", amount, percent, total.StringFixed(2))
}
```

### 下载依赖模块

严格来说我们下载的是一个 go module (模块)，一个模块中包含若干个 package

```bash
go get github.com/go-sql-driver/mysql          # 可以用 go get 下载模块
go get github.com/go-sql-driver/mysql@v1.5.1   # 可以用 @v1.5.1 后缀下载指定版本
go get github.com/gdamore/tcell/v2             # 存在 v2 版本的模块会在 module path 中以 /v2 结尾
go get github.com/gdamore/tcell                # 不写 v2 后缀时下载的是 v1 的最新版
```

下载一个模块后，go.mod 中会多出一行，require github.com/go-sql-driver/mysql v1.6.0  
因为 go.mod 中记录了项目依赖哪些包，所以别人可以用 `go mod download` 一次性下载所有依赖包  
`go.sum` 文件记录了包的 hash，用于验证所下载的包有没有被修改过，确保下载了完全相同的包

When you download a dependency with `go get`, then the exact version of the dependency will be recorded in the go.mod file. When you run or build the code in your project, Go will use the exact dependencies listed in the go.mod file. If the necessary dependencies aren’t already on your local machine, then Go will automatically download them for you — along with any recursive dependencies too.

By default, Go picks the latest version of a dependency when you add it to your project. We can see what versions of the module are available with the `go list` command:

```bash
go list -m -versions github.com/gdamore/tcell/v2
```

#### ➤ 经常用到的 [golang.org/x](https://pkg.go.dev/golang.org/x) 中的包有什么来头吗?

> - 它们也是官方出品，有的已经进入了标准库 ( 例如 [golang.org/x/net/context](https://pkg.go.dev/golang.org/x/net/context) )，有的还没
> - 因为还没进入标准库，所以可以引入破坏式更新，版本号也是 `v0.22.0` 这样不太靠谱的样子

### 升级依赖模块

#### ➤ Run `go get -u ./... && go mod tidy`. 

[参考此处](https://stackoverflow.com/a/67202563/7549292):

- [`go get -u`](https://go.dev/ref/mod#go-get) (same as `go get -u .`) updates the package in the current directory and its dependencies to the *newer minor or patch releases when available.* In typical projects, running this in the module root is enough, as it likely imports everything else.
- `go get -u ./...` will expand to all packages rooted in the current directory, which effectively also updates everything.
- Following from the above, `go get -u ./foo/...` will update everything that is rooted in `./foo`
- `go get -u all` updates everything **including test dependencies**.

```bash
$ go get -u ./... && go mod tidy                    # 一般用这个升级所有包
$ go list -u -m all                                 # 查看哪些包有新版本，下面这行带过滤:
$ go list -u -m -f '{{if .Update}}{{.}}{{end}}' all # 参考 https://stackoverflow.com/a/55866702/7549292

$ go get -u github.com/foo/bar    # 升级单个包
$ go get -u .                     # 升级当前目录下的这个包，如果这个包引入了模块中所有的包，也就相当于升级所有包
$ go get -u ./...                 # 升级所有包，因为 ./... 会遍历所有子目录中的包

# If you want to upgrade to a specific version:  
$ go get -u github.com/foo/bar@v2.0.0

# You can remove any unused packages from your go.mod and go.sum files:  
$ go mod tidy -v

# To view available minor and patch upgrades for all direct and indirect dependencies, run 
$ go list -u -m all
```

#### ➤ all 的含义参考 [Package lists and patterns¶](https://pkg.go.dev/cmd/go#hdr-Package_lists_and_patterns)

> The main module is the module containing the directory where the `go` command is invoked. When using modules, "all" expands to all packages in the main module and their dependencies, including dependencies needed by tests of any of those. `go get -u all` updates modules from the [build list](https://go.dev/ref/mod#glos-build-list) from `go.mod`.

### 间接依赖的模块

A dependency of a module can be of two kinds:

- **Direct** - A direct dependency is a dependency which the module directly imports.
- **Indirect** – It is the dependency that is imported by the module’s direct dependencies. Also, any dependency that is mentioned in the go.mod file but not imported in any of the source files of the module is also treated as an indirect dependency.

### 查看依赖树

[build - How to view full dependency tree for nested Go dependencies](https://stackoverflow.com/questions/44521071/how-to-view-full-dependency-tree-for-nested-go-dependencies)

```bash
go mod graph                        # 列出各个模块的依赖模块
go mod graph | deptree -d 2 -t -a   # 安装 deptree 命令后，用它显示一棵树
go mod why golang.org/x/time/rate   # 反查为什么需要这个包
```

### 理解版本号

Semantic versioning divides a version number into three parts: the major version, the minor version, and the patch version, which are written as `major.minor.patch` and preceded by a `v`. The patch version number is incremented when fixing a bug, the minor version number is incremented (and the patch version is set back to 0) when a new, backward-compatible feature is added, and the major version number is incremented (and minor and patch are set back to 0) when making a change that breaks backward compatibility.

For projects that are still experimental — at major version `v0` — occasional breaking changes are expected by users. For projects which are declared stable — at major version `v1` or higher — breaking changes must be done in a new major version.

### 发布一个模块就是扔到 VCS

#### ➤ go get 可以从版本控制系统 (比如 Git) 下载源码

Making your module available to other people is as simple as putting it in a version control system. This is true whether you are releasing your project as open source on a public version control system like GitHub or a private one that’s hosted within your organization. Since Go programs build from source code and use a repository path to identify themselves, there’s no need to explicitly upload your module to a central library repository, like you do for Maven Central or npm. Make sure you check in both your go.mod file and your go.sum file.

#### ➤ 模块的版本就是仓库中的 tag

Whether your module is public or private, you must properly version your module so that it works correctly with Go’s module system. As long as you are adding functionality or patching bugs, the process is simple. Store your changes in your source code repository, then apply a tag that follows the semantic versioning rules.

### 会选择同一个模块的最新版本

At some point, your project will depend on two or more modules that all depend on the same module. As often happens, these modules declare that they depend on different minor or patch versions of that module. How does Go resolve this?

The module system uses the principle of minimum version selection. Let’s say that your module directly depends on modules A, B, and C. All three of these modules depend on module D. The go.mod file for module A declares that it depends on `v1.1.0`, module B declares that it depends on `v1.2.0`, and module C declares that it depends on `v1.2.3`. Go will import module D only once, and it will choose version `v1.2.3`, as that is the most recent specified version.

However, as sometimes happens, you might find that while module A works with version v1.1.0 of module D, it does not work with version v1.2.3. What do you do then? Go’s answer is that you need to contact the module authors to fix their incompatibilities. 

The import compatibility rule says that all minor and patch versions of a module must be backward compatible. If they aren’t, it’s a bug. In our example, either module D needs to be fixed because it broke backward compatibility, or module A needs to be fixed because it made a faulty assumption about the behavior of module D.  

This isn’t the most satisfying answer, but it’s the most honest. Some build systems, like npm, will include multiple versions of the same package. This can introduce its own set of bugs, especially when there is package-level state. It also increases the size of your application. In the end, some things are better solved by community than code.  

### Semantic Import Versioning

The import compatibility rule for Go: If an old package and a new package have the same import path,  the new package must be backwards compatible with the old package. 

The import compatibility rule dramatically simplifies the experience of using incompatible versions of a package. When each different version has a different import path, there is no ambiguity about the intended semantics of a given import statement. This makes it easier for both developers and tools to understand Go programs. 

总而言之,  导入路径相同的包应该是兼容的  
如果某包从 1.x.x 升级到了 2.x.x 那么需要在 module path 中加一个 v2, 例如 `import "my/thing/v2/sub/pkg"`.

![img](https://research.swtch.com/impver.png) 

### Module Proxy Servers

Every Go module is stored in a source code repository, like GitHub or GitLab. But by default, go get doesn’t fetch code directly from source code repositories. Instead, it sends requests to a [proxy server](https://proxy.golang.org/) run by Google. This server keeps copies of every version of virtually all public Go modules. If a module or a version of a module isn’t present on the proxy server, it downloads the module from the module’s repository, stores a copy, and returns the module.  

In addition to the proxy server, Google also maintains a sum database. It stores information on every version of every module. Just as the proxy server protects you from a module or a version of a module being removed from the internet, the sum database protects you against modifications to a version of a module. This could be malicious (someone has hijacked a module and slipped in malicious code), or it could be inadvertent (a module maintainer fixes a bug or adds a new feature and reuses an existing version tag). In either case, you don’t want to use a module version that has changed because you won’t be building the same binary and don’t know what the effects are on your application.  

Every time you download a module via go build, go test, or go get, the Go tools calculate a hash for the module and contact the sum database to compare the calculated hash to the hash stored for that module’s version. If they don’t match, the module isn’t installed.  

#### ➤ Specifying a Proxy Server

- If you don’t want to use Google’s, you can switch to GoCenter by setting the GOPROXY environment
  variable to https://gocenter.io,direct.  
- You can disable proxying entirely by setting the GOPROXY environment variable to `direct`. You’ll download modules directly from their repositories, but if you depend on a version that’s removed from the repository, you won’t be able to access it.
- You can run your own proxy server. The Athens Project provides an open source proxy server. Install one of these products on your network and then point GOPROXY to the URL.  

#### ➤ Private Repositories

Most organizations keep their code in private repositories. If you want to use a private module in another Go project, you can’t request it from Google’s proxy server. Go will fall back to checking the private repository directly, but you might not want to leak the names of private servers and repositories to external services. 

If you are using your own proxy server, or if you have disabled proxying, this isn’t an issue.  
If you are using a public proxy server, you can set the GOPRIVATE environment variable to a comma-separated list of your private repositories. For example, if you set GOPRIVATE to:  `GOPRIVATE=*.example.com,company.com/repo`. Any module stored in a repository that’s located at any subdomain of `example.com` or at a URL that starts with `company.com/repo` will be downloaded directly.  

