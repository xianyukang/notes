## Table of Contents
  - [GORM](#GORM)
    - [参考](#%E5%8F%82%E8%80%83)
    - [入门](#%E5%85%A5%E9%97%A8)
    - [日志配置](#%E6%97%A5%E5%BF%97%E9%85%8D%E7%BD%AE)
    - [GORM 配置](#GORM-%E9%85%8D%E7%BD%AE)
  - [声明模型](#%E5%A3%B0%E6%98%8E%E6%A8%A1%E5%9E%8B)
    - [例子](#%E4%BE%8B%E5%AD%90)
    - [创建或迁移表](#%E5%88%9B%E5%BB%BA%E6%88%96%E8%BF%81%E7%A7%BB%E8%A1%A8)
    - [一些 gorm tag](#%E4%B8%80%E4%BA%9B-gorm-tag)
  - [Create](#Create)
    - [插入新记录](#%E6%8F%92%E5%85%A5%E6%96%B0%E8%AE%B0%E5%BD%95)
    - [Create Hooks](#Create-Hooks)
    - [插入记录的各种方式](#%E6%8F%92%E5%85%A5%E8%AE%B0%E5%BD%95%E7%9A%84%E5%90%84%E7%A7%8D%E6%96%B9%E5%BC%8F)
  - [Query](#Query)
    - [根据 ID 查询数据](#%E6%A0%B9%E6%8D%AE-ID-%E6%9F%A5%E8%AF%A2%E6%95%B0%E6%8D%AE)
    - [使用 Where 条件](#%E4%BD%BF%E7%94%A8-Where-%E6%9D%A1%E4%BB%B6)
    - [使用分组条件](#%E4%BD%BF%E7%94%A8%E5%88%86%E7%BB%84%E6%9D%A1%E4%BB%B6)
    - [选择特定字段](#%E9%80%89%E6%8B%A9%E7%89%B9%E5%AE%9A%E5%AD%97%E6%AE%B5)
    - [Order、Limit、Group](#OrderLimitGroup)
    - [Distinct、Pluck、Count](#DistinctPluckCount)
    - [写一个分页器](#%E5%86%99%E4%B8%80%E4%B8%AA%E5%88%86%E9%A1%B5%E5%99%A8)
    - [一些高级查询](#%E4%B8%80%E4%BA%9B%E9%AB%98%E7%BA%A7%E6%9F%A5%E8%AF%A2)
  - [Update](#Update)
    - [各种更新方式](#%E5%90%84%E7%A7%8D%E6%9B%B4%E6%96%B0%E6%96%B9%E5%BC%8F)
    - [其他高级选项](#%E5%85%B6%E4%BB%96%E9%AB%98%E7%BA%A7%E9%80%89%E9%A1%B9)
    - [删除相关](#%E5%88%A0%E9%99%A4%E7%9B%B8%E5%85%B3)
  - [原生 SQL](#%E5%8E%9F%E7%94%9F-SQL)
    - [Raw 和 Exec](#Raw-%E5%92%8C-Exec)
    - [使用 sql.Row](#%E4%BD%BF%E7%94%A8-sqlRow)
  - [关联关系](#%E5%85%B3%E8%81%94%E5%85%B3%E7%B3%BB)
    - [Has Many](#Has-Many)
    - [Belongs To](#Belongs-To)
    - [Has One 与 Belongs To](#Has-One-%E4%B8%8E-Belongs-To)
    - [Many to Many](#Many-to-Many)
    - [关联查询与关联创建](#%E5%85%B3%E8%81%94%E6%9F%A5%E8%AF%A2%E4%B8%8E%E5%85%B3%E8%81%94%E5%88%9B%E5%BB%BA)
  - [使用相关](#%E4%BD%BF%E7%94%A8%E7%9B%B8%E5%85%B3)
    - [使用 Context](#%E4%BD%BF%E7%94%A8-Context)
    - [别忘了处理错误](#%E5%88%AB%E5%BF%98%E4%BA%86%E5%A4%84%E7%90%86%E9%94%99%E8%AF%AF)
    - [GORM Session 相关](#GORM-Session-%E7%9B%B8%E5%85%B3)
    - [钩子函数](#%E9%92%A9%E5%AD%90%E5%87%BD%E6%95%B0)
    - [自定义数据类型](#%E8%87%AA%E5%AE%9A%E4%B9%89%E6%95%B0%E6%8D%AE%E7%B1%BB%E5%9E%8B)
    - [用 Scopes 封装查询](#%E7%94%A8-Scopes-%E5%B0%81%E8%A3%85%E6%9F%A5%E8%AF%A2)
  - [事务操作](#%E4%BA%8B%E5%8A%A1%E6%93%8D%E4%BD%9C)
    - [禁用默认事务](#%E7%A6%81%E7%94%A8%E9%BB%98%E8%AE%A4%E4%BA%8B%E5%8A%A1)
    - [普通事务](#%E6%99%AE%E9%80%9A%E4%BA%8B%E5%8A%A1)
    - [嵌套事务](#%E5%B5%8C%E5%A5%97%E4%BA%8B%E5%8A%A1)
    - [手动事务](#%E6%89%8B%E5%8A%A8%E4%BA%8B%E5%8A%A1)

## GORM

### 参考

- [官方文档](https://gorm.io/zh_CN/docs/index.html)
- [视频教程](https://www.bilibili.com/video/BV1E64y1472a/)

### 入门

#### ➤ 安装

```bash
go get -u gorm.io/gorm
go get -u gorm.io/driver/mysql
```

#### ➤ 初始化 ORM

> 注意: 想要正确的处理 `time.Time` ，dsn 需要带上 `parseTime=True` 参数，要支持完整的 UTF-8 编码，需要加上 `charset=utf8mb4` 参数, 查看 [此文章](https://mathiasbynens.be/notes/mysql-utf8mb4) 理解 mysql 的字符串编码. 另外还需要添加 `loc=Local` ,  这样 mysql driver 从查询结果生成 time.Time 结构时,  生成的 time.Time 结构的 Location 字段才是 Local. [参数来源 mysql driver 文档](https://github.com/go-sql-driver/mysql#parameters)
>

```go
import gormMySql "gorm.io/driver/mysql"
dsn := "root:password@tcp(localhost)/snippetbox?charset=utf8mb4&parseTime=True&loc=Local"
db, err := gorm.Open(gormMySql.Open(dsn), &gorm.Config{
    Logger: logger.Default.LogMode(logger.Info),
})
if err != nil {
    panic("failed to connect database")
}
```

#### ➤ 定义 model

```go
// 其中 gorm.Model 定义了主键 ID、创建时间、更新时间等字段,  此外还有删除时间表示逻辑删除
type Product struct {
  gorm.Model
  Code  string
  Price uint
}
```

#### ➤ 几个 CRUD 例子

```go
  // Create
  db.Create(&Product{Code: "D42", Price: 100})

  // Read
  var product Product
  db.First(&product, 1)                  // find product with integer primary key
  db.First(&product, "code = ?", "D42")  // find product with code D42

  // Update
  // db.Model(&product) 表示用 product.ID 作为筛选条件
  db.Model(&product).Update("Price", 200)
  db.Model(&product).Updates(Product{Price: 200, Code: "F42"})
  db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

  // Delete - delete product
  db.Delete(&product, 1)
```

### 日志配置

#### [➤ 参考文档](https://gorm.io/docs/logger.html)

1. Gorm will print Slow SQL and happening errors by default.
2. 可以修改日志的输出目标、输出格式、如果把日志级别调成 info 能方便看 SQL 日志
4. 另外可以用 `db.Debug()` 把一个数据库操作的日志级别设为 info,  比如:  
   db.Debug().Where("name = ?", "jinzhu").First(&User{})

### GORM 配置

#### [➤ 参考文档](https://gorm.io/zh_CN/docs/gorm_config.html)

1. 修改创建表时的命名策略,  比如表名用什么前缀 (默认无)、使用单数还是复数 (默认复数, 推荐改成单数)

2. 禁用创建表时自动添加外键约束,  [大家设计数据库时使用外键吗？](https://www.zhihu.com/question/19600081)

在 `AutoMigrate` 或 `CreateTable` 时，GORM 会自动创建外键约束，可以禁用该特性.

#### ➤ 外键约束是什么?  

>
> 通过定义外键约束，关系数据库可以保证无法插入无效的数据。即如果 `classes` 表不存在 `id=99` 的记录，`students`表就无法插入`class_id=99`的记录。由于外键约束会降低数据库的性能，大部分互联网应用程序为了追求速度，并不设置外键约束，而是仅靠应用程序自身来保证逻辑的正确性。这种情况下，`class_id` 仅仅是一个普通的列，只是它起到了外键的作用而已。

3. 题外话, 为什么常常看到 `varchar(191)`、`varchar(255)` ?

(1) [mysql 使用字符串的前 N 个字符做索引](https://dba.stackexchange.com/a/122340)  
(2) InnoDB 存储引擎限制了字符串索引不能超过 767 Byte,  [如果用 utf8mb4 编码也就是 191 个字符](https://stackoverflow.com/a/1814594)  
(3) 所以一开始把字符串列设为 `varchar(191)`,  后面对这个 string column 建索引就不会报错,  比较方便

#### ➤ 例子如下

```go
func GORM配置() {
    var err error
    dsn := "root:password@tcp(localhost)/snippetbox?charset=utf8mb4&parseTime=True&loc=Local"
    db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
        Logger: logger.Default.LogMode(logger.Info), // 日志级别
        NamingStrategy: schema.NamingStrategy{
            TablePrefix:   "",   // 不加前缀
            SingularTable: true, // 单数表名
        },
        DisableForeignKeyConstraintWhenMigrating: true, // 创建表时禁用外键约束
    })
    CheckError(err)
}
```

## 声明模型

### 例子

Models are normal structs with basic Go types, pointers/alias of them or custom types implementing [Scanner](https://pkg.go.dev/database/sql/?tab=doc#Scanner) and [Valuer](https://pkg.go.dev/database/sql/driver#Valuer) interfaces. For Example:

```go
type User struct {
  ID           uint
  Name         string
  Email        *string 
  Age          uint8
  Birthday     *time.Time
  MemberNumber sql.NullString
  ActivatedAt  sql.NullTime
  CreatedAt    time.Time
  UpdatedAt    time.Time
}
```

#### ➤ GORM 约定的表名、列名、更新时间

GORM prefers convention over configuration. By default, GORM uses `ID` as primary key, pluralizes struct name to `snake_cases` as table name, `snake_case` as column name, and uses `CreatedAt`, `UpdatedAt` to track creating/updating time. If you follow the conventions adopted by GORM, you’ll need to write very little configuration/code. If convention doesn’t match your requirements, [GORM allows you to configure them](https://gorm.io/docs/conventions.html).

### 创建或迁移表

```go
type Character struct {
    gorm.Model                          // 嵌入 ID、CreatedAt、UpdatedAt、DeletedAt 等三个字段
    Age        uint8  `gorm:"not null"` // 如果用 int 对应到数据库则是 bigint,  有点浪费
    Name       string `gorm:"size:191"` // 注意设置 varchar(191),  否则默认是 longtext 类型
    From       string `gorm:"size:191"`
}

func 创建表() {
    _ = db.AutoMigrate(&Character{})

    // db.AutoMigrate() 能创建表或修改表定义、 另外可用 db.Migrator() 增删表、列、索引
    _ = db.Migrator().CreateTable(&Character{})
    fmt.Println(db.Migrator().GetTables())

    // 一般只在开发环境使用 GORM 的建表功能,  可以这样打印 GORM 生成的 SQL
    _ = db.Session(&gorm.Session{DryRun: true}).Debug().Migrator().CreateTable(&Character{})
}

//CREATE TABLE `character`
//(
//    `id`         bigint unsigned AUTO_INCREMENT,
//    `created_at` datetime(3) NULL,
//    `updated_at` datetime(3) NULL,
//    `deleted_at` datetime(3) NULL,
//    `age`        tinyint unsigned NOT NULL,
//    `name`       varchar(191),
//    `from`       varchar(191),
//    PRIMARY KEY (`id`),
//    INDEX `idx_character_deleted_at` (`deleted_at`)
//);
```

### 一些 gorm tag

#### [➤ 参考文档](https://gorm.io/docs/models.html#Fields-Tags)

```go
type GormTag struct {
    // 用 type 标签直接写 mysql 字段定义,  或者用 gorm 提供的 size/not null/default 等标签能兼容多种数据库
    Name  string `gorm:"type:varchar(191) not null default 'a b c'"`
    Name2 string `gorm:"size:191; not null; default:a b c"`

    // 把结构体中的 author 字段映射到表中的 writer 列
    Author string `gorm:"size:191; not null; column:writer; comment:作者"`

    // 如果数据库 title 字段允许为 NULL,  并且想插入 NULL,  可以用指针类型 *string
    // *string 与 sql.NullString 的区别: https://stackoverflow.com/q/40092155
    Title  string         `gorm:"size:191; not null"`
    Title2 *string        `gorm:"size:191"`
    Title3 sql.NullString `gorm:"size:191"`
}
// DeletedAt 字段在 GORM 中约定为逻辑删除
// 我想让字符串比较始终区分大小写、所以建表时设置了 collate utf8mb4_bin
// 若 Query 中想让字符串区分大小写, 参考: https://dev.mysql.com/doc/refman/8.0/en/case-sensitivity.html
type Character struct {
    ID        uint   `gorm:"primaryKey"`
    Name      string `gorm:"type:varchar(191) character set utf8mb4 collate utf8mb4_bin; uniqueIndex"`
    From      string `gorm:"type:varchar(191) character set utf8mb4 collate utf8mb4_bin;"`
    Age       uint8
    DeletedAt gorm.DeletedAt
}
```



## Create

### 插入新记录

```go
func 插入新记录() {
    c := Character{
        Name: "Homura",
        Age:  16,
        From: "Xenoblade 2",
    }
    result := db.Create(&c)
    // 用 result.Error 检查错误
    // 用 c.ID 获取自动生成的 ID
    // 用 result.RowsAffected 表示插入行数
    CheckError(result.Error)
    fmt.Println(c.ID, result.RowsAffected)
}

func 插入记录时选择需要的字段() {
    c := Character{
        Name: "Hikari",
        Age:  16,
        From: "Xenoblade 2",
    }
    // 选择 Name、From 字段、另外 CreatedAt、UpdatedAt 会被自动添加
    // INSERT INTO `character` (`created_at`,`updated_at`,`name`,`from`) VALUES (...)
    err := db.Select("Name", "From").Create(&c).Error
    CheckError(err)

    // 忽略三个时间字段
    // INSERT INTO `character` (`name`,`age`,`from`) VALUES (...)
    c.ID = 0
    err = db.Omit("UpdatedAt", "CreatedAt", "DeletedAt").Create(&c).Error
    CheckError(err)
}
```

### Create Hooks

GORM allows hooks `BeforeSave`, `BeforeCreate`, `AfterSave`, `AfterCreate`, those methods will be called when creating a record, refer [Hooks](https://gorm.io/docs/hooks.html) for details

```go
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
  u.UUID = uuid.New()
  if u.Role == "admin" {
    return errors.New("invalid role")
  }
  return
}
```

### 插入记录的各种方式

#### ➤ 批量插入

Pass slice data to method `Create`, GORM will generate a single SQL statement to insert all the data and backfill primary key values, hook methods will be invoked too.

```go
func 批量插入() {
    characters := []Character{{Name: "Rex"}, {Name: "Homura"}, {Name: "Hikari"}}
    err := db.Create(&characters).Error
    CheckError(err)
    for i := range characters {
        fmt.Println(characters[i].ID, characters[i].Name)
    }
}
```

#### ➤ [在模型中用 default 标签指定默认值](https://gorm.io/docs/create.html#Default-Values)

如果插入记录时 Age 字段是零值,  就会用 default 标签设置的默认值  
注意如果数据库中为 age 设了默认值,  那么模型中的 Age 字段也要加上 default 标签,  否则很容易插入零值

#### ➤ [Upsert](https://gorm.io/docs/create.html#Upsert-On-Conflict)

```go
func Upsert() {
    c := Character{
        Name: "Homura",
        Age:  16,
        From: "Xenoblade 2",
    }
    // 尝试插入新纪录,  如果违背了 name 字段的唯一索引,  则什么也不做
    // 注意 insert into ... on duplicate key update ... 中的 insert 总会让自增列加一导致主键 id 不连续
    result := db.Clauses(clause.OnConflict{DoNothing: true}).Create(&c)
    CheckError(result.Error)
    fmt.Println(c.ID, result.RowsAffected)

    // 尝试插入新纪录,  如果违背了 name 字段的唯一索引,  则更新 age、from、updated_at 字段
    columns := []string{"age", "from", "updated_at"}
    result = db.Clauses(clause.OnConflict{DoUpdates: clause.AssignmentColumns(columns)}).Create(&c)
    CheckError(result.Error)
    fmt.Println(c.ID, result.RowsAffected)
}
```

#### ➤ [返回首个匹配的记录、没有则创建一条新纪录](https://gorm.io/docs/advanced_query.html#FirstOrCreate)

```go
func FirstOrCreate() {
    c := Character{
        Name: "Cloud",
        Age:  21,
        From: "FF7",
    }
    // 先用一条 sql 查询数据库有没有 name 为 Hikari 的记录,  有则用查到的数据填充 c
    // 若没有 name 为 Hikari 的记录则执行第二条 sql,  用 c 中的数据插入一条记录
    result := db.FirstOrCreate(&c, Character{Name: c.Name})
    CheckError(result.Error)
    fmt.Println(c.ID, result.RowsAffected)
}
// FirstOrCreate 也支持用 Assign 做插入或更新,  这与 on duplicate key update ... 的区别是使用了两条 sql
```

#### ➤ 进阶内容

1. 可以不用结构体,  [而是用 Map 描述插入的记录](https://gorm.io/docs/create.html#Create-From-Map)
2. [如果插入记录时需要用到 uuid() 这样的 SQL 函数](https://gorm.io/docs/create.html#Create-From-SQL-Expr-Context-Valuer)
3. [insert record 的同时 insert 关联的 record](https://gorm.io/docs/create.html#Create-With-Associations)

## Query

### 根据 ID 查询数据

```go
func FirstAndFind() {
    // 按主键升序排序、First 返回第一个、Last 返回最后一个,  Take 则不排序
    // First、Last、Take 在没有找到匹配记录时会返回 gorm.ErrRecordNotFound
    var c Character
    err := db.First(&c, 1).Error
    if err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            // 不存在 id 为 1 的记录
        } else {
            panic(err)
        }
    }
    fmt.Println(c.ID, c.Name)

    // 根据主键查询多条数据, where id in (1, 2, 3)
    cs := make([]Character, 0)
    err = db.Find(&cs, []int{1, 2, 3}).Error
    CheckError(err)
    for _, c := range cs {
        fmt.Println(c.ID, c.Name)
    }
}
```

### 使用 Where 条件

#### ➤ String 条件

```go
func String_条件() {
   var c Character
   var cs []Character
   db.Where("age > ?", 20).Find(&cs)                          // where age > 20
   db.Where("age between ? AND ?", 20, 21).Find(&cs)          // where age between 20 AND 21
   db.Where("name like ?", "Ti%").Find(&cs)                   // where name like 'Ti%'
   db.Where("name in ?", []string{"Tifa", "Cloud"}).Find(&cs) // where name in ('Tifa', 'Cloud')
   db.Where("name = ?", "Tifa").First(&c)                     // where name = 'Tifa' order by id limit 1
   db.Where("name = ? AND age = ?", "Tifa", 20).First(&c)     // where name = 'Tifa' AND age = 20 order by id limit 1
}
```

#### ➤ Struct & Map 条件

```go
func Struct_Map_条件() {
  var c Character
  var cs []Character
  db.Where(&Character{Name: "Cloud", Age: 21}).First(&c)         // where name = 'Cloud' AND age = 21 order by id limit 1
  db.Where(map[string]any{"name": "Cloud", "age": 21}).Find(&cs) // where name = 'Cloud' AND age = 21

  db.Where(&Character{Name: "Cloud", Age: 0}).Find(&cs)         // 结构体零值字段不参与查询: where name = 'Cloud'
  db.Where(map[string]any{"name": "Cloud", "age": 0}).Find(&cs) // map 则都参与查询: where name = 'Cloud' AND age = 0
}
```

#### ➤ 条件可以直接写在 First、Find 中

```go
// 上面的提到的条件都可以直接写在 First、Last、Take、Find 中:  
db.Find(&cs, "name = ? AND `from` = ?", "Homura", "Xenoblade 2") // 这里用 `from` 转义 sql 关键字
```

#### ➤ Not 条件

```go
func Not_条件() {
   var cs []Character
   db.Not("`from` = ?", "FF7").Find(&cs)                               // where NOT `from` = 'FF7'
   db.Not(map[string]any{"name": []string{"Tifa", "Cloud"}}).Find(&cs) // where name NOT IN ('Tifa', 'Cloud')
   db.Not(Character{Name: "Ichigo", Age: 17}).Find(&cs)                // where (name != 'Ichigo' AND age != 17)
   db.Not([]int{1, 2, 3, 4, 5}).Find(&cs)                              // where id NOT IN (1,2,3,4,5)
   fmt.Println(cs)
}
```

#### ➤ Or 条件

```go
func Or_条件() {
   var cs []Character
   var orCondition = map[string]any{"name": "Ichigo", "age": 17}
   db.Where("id = ?", 6).Or("`from` = ?", "Bleach").Find(&cs) // where id = 6 OR `from` = 'Bleach'
   db.Where("name = ?", "Rukia").Or(orCondition).Find(&cs)    // where name = 'Rukia' OR (age = 17 AND name = 'Ichigo')
   fmt.Println(cs)
}
```

### 使用分组条件

#### ➤ `c1.Where(c2)` 会为 c2 加括号,  然后用 `AND` 连接 c1 和 c2

```go
func c1_Where_c2() {
   var cs []Character
   var c1 = db.Where("name = ?", "Rukia")           // name = 'Rukia'
   var c2 = db.Where("age = ?", 18).Or("age = 150") // age = 18 OR age = 150
   var c1_c2 = c1.Where(c2)                         // name = 'Rukia' AND (age = 18 OR age = 150)
   c1_c2.Find(&cs)
}
```

#### ➤ `db.Where(c1).Or(c2)` 会为 c1 和 c2 加上括号,  并用 `OR` 连接两个条件

```go
func Where_c1_c2() {
   var cs []Character
   var c1 = db.Where("name = ?", "Rukia")                       // name = 'Rukia'
   var c2 = db.Where("name = ?", "Ichigo").Where("age = ?", 17) // name = 'Ichigo' AND age = 17
   var c1_c2 = db.Where(c1).Or(c2)                              // (name = 'Rukia') OR (name = 'Ichigo' AND age = 17)
   c1_c2.Find(&cs)
}
```

### 选择特定字段

#### ➤ 默认会用 SELECT * 选择全部字段,  用 Select() 选择特定字段

```go
func 选择特定字段() {
   var cs []Character
   db.Select("name", "from").Find(&cs)                         // SELECT name, `from`
   db.Select([]string{"name", "from"}).Find(&cs)               // SELECT name, `from`
   db.Select("coalesce(name, ?) as name", "default").Find(&cs) // SELECT coalesce(name, 'default') as name
}
```

#### ➤ 可以用另一个结构体表示选择的字段

```go
type CharacterInfo struct {
   ID   uint
   Name string
   Age  uint8
}

func UseStructAsSelect() {
   var cs []CharacterInfo                     // 传 &Character{} 可以省点复制开销
   db.Model(&Character{}).Limit(10).Find(&cs) // SELECT id, name, age FROM `character` LIMIT 10
}
```

### Order、Limit、Group

#### ➤ Order、Limit、Offset

```go
func Order_Limit_Offset() {
   var cs []Character
   db.Order("age desc, name").Find(&cs)         // ORDER BY age desc, name
   db.Order("age desc").Order("name").Find(&cs) // ORDER BY age desc, name
   db.Limit(2).Find(&cs)                        // SELECT * FROM `character` LIMIT 2
   db.Limit(10).Offset(5).Find(&cs)             // SELECT * FROM `character` LIMIT 10 OFFSET 5
}

func Group_Having() {
   type Result struct {
      Count int
      From  string
   }
   var rs []Result
   // SELECT `from`, count(*) as `count` FROM `character` WHERE age > 16 GROUP BY `from` HAVING count > 1;
   db.Model(&Character{}).Select("`from`, count(*) as `count`").Where("age > ?", 16).
      Group("from").Having("count > ?", 1).Find(&rs)
}
```

### Distinct、Pluck、Count

```go
func Distinct_Pluck_Count() {
   var cs []Character
   var from []string
   var count int64
   db.Distinct("from", "age").Find(&cs)                            // SELECT DISTINCT `from`, age FROM `character`;
   db.Model(&Character{}).Distinct().Pluck("from", &from)          // Pluck 用于查询单列数据并将结果扫描到切片
   db.Model(&Character{}).Where("`from` = ?", "FF7").Count(&count) // Count 用于获取匹配的记录数

   // 去重计数: SELECT COUNT(DISTINCT(age)) FROM `character` WHERE id < 100;
   db.Model(&Character{}).Where("id < ?", 100).Distinct("age").Count(&count)
}
```

### 写一个分页器

```go
func Paginate(page, pageSize string) func(db *gorm.DB) *gorm.DB {
   return func(db *gorm.DB) *gorm.DB {
      page, _ := strconv.Atoi(page)
      if page == 0 {
         page = 1
      }

      pageSize, _ := strconv.Atoi(pageSize)
      switch {
      case pageSize > 100:
         pageSize = 100
      case pageSize <= 0:
         pageSize = 10
      }

      offset := (page - 1) * pageSize
      return db.Offset(offset).Limit(pageSize)
   }
}

func 分页器() {
   var cs []Character
   db.Scopes(Paginate("1", "2")).Find(&cs)     // SELECT * FROM `character` LIMIT 0,2
   db.Scopes(Paginate("2", "2")).Find(&cs)     // SELECT * FROM `character` LIMIT 2,2
   db.Scopes(Paginate("3", "2")).Find(&cs)     // SELECT * FROM `character` LIMIT 4,2
   db.Scopes(Paginate("???", "???")).Find(&cs) // 遇到非法字符串则默认为第 1 页、10 条数据
}
```

### 一些高级查询

- [使用 JOIN](https://gorm.io/zh_CN/docs/query.html#Joins)
- [使用 锁](https://gorm.io/zh_CN/docs/advanced_query.html#Locking-FOR-UPDATE)
- [使用 子查询](https://gorm.io/zh_CN/docs/advanced_query.html#%E5%AD%90%E6%9F%A5%E8%AF%A2)
- [使用命名参数 `@name` 代替占位符 `?`](https://gorm.io/zh_CN/docs/advanced_query.html#%E5%91%BD%E5%90%8D%E5%8F%82%E6%95%B0)
- [Find 至 Map](https://gorm.io/zh_CN/docs/advanced_query.html#Find-%E8%87%B3-map)
- [使用 Find 会触发查询钩子 AfterFind](https://gorm.io/zh_CN/docs/advanced_query.html#%E6%9F%A5%E8%AF%A2%E9%92%A9%E5%AD%90)
- [使用 Scopes 封装常用筛选条件](https://gorm.io/zh_CN/docs/advanced_query.html#Scopes)



## Update

### 各种更新方式

#### ➤ 保存所有字段、或插入新记录

```go
func db_Save() {
   // Save 会保存所有的字段，即使字段是零值
   // 如果 ID 对应的记录不存在, 那么 update 会变成 insert
   var c Character
   err := db.First(&c, 1).Error
   CheckError(err)
   c.Age = 15
   err = db.Save(&c).Error
   CheckError(err)
}
```

#### ➤ 更新单列、更新多列、批量更新

```go
func UpdateSingleColumn() {
   db.Model(&Character{}).Where("name = ?", "tifa").Update("name", "Tifa")      // where name = 'tifa'
   db.Model(&Character{ID: 1}).Where("name = ?", "tifa").Update("name", "Tifa") // 如果 ID 非零值, ID 也会用于筛选
}

func UpdateMultiColumn() {
   // 可以用 map 或 struct 更新多列,
   // struct 的零值字段不参与更新 (除非使用 Select),  map 则全都参与更新
   db.Model(&Character{ID: 233}).Updates(Character{Name: "Leon", From: "RE4", Age: 0})            // 不更新 age
   db.Model(&Character{ID: 233}).Updates(map[string]any{"name": "Leon", "from": "RE4", "age": 0}) // 更新 age
}

func 批量更新() {
    // 如果传给 Model() 的结构体的主键为零值，则 GORM 会执行批量更新:
    db.Model(&Character{}).Where("`from` = ?", "RE4").Update("from", "Resident Evil 4")

    // 如果在没有任何条件的情况下执行批量更新，GORM 不会执行该操作，并返回 ErrMissingWhereClause 错误
    // 对此，你必须加一些条件，或者使用原生 SQL:
    db.Model(&Character{}).Where("1 = 1").Update("name", "xxx")
    db.Exec("UPDATE `character` SET name = ?", "xxx")
}
```

#### ➤ 更新选定字段

```go
func UpdateSelected() {
   var c = Character{ID: 233}
   db.Model(&c).Select("name").Updates(map[string]any{"name": "Leon", "from": "RE4"}) // 只更新 name, 忽略其他字段
   db.Model(&c).Omit("name").Updates(map[string]any{"name": "Leon", "from": "RE4"})   // 不更新 name, 更新其他字段
   db.Model(&c).Select("name", "age").Updates(Character{Name: "Leon", Age: 0})        // 会更新 age 即使 age 是零值
}
```

#### ➤ 更新操作支持 BeforeSave、BeforeUpdate、AfterSave、AfterUpdate 等 [钩子函数](https://gorm.io/zh_CN/docs/hooks.html)

```go
func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
    if u.Role == "admin" {
        return errors.New("admin user not allowed to update")
    }
    return
}
```

### 其他高级选项

- [更新时使用表达式](https://gorm.io/zh_CN/docs/update.html#%E4%BD%BF%E7%94%A8-SQL-%E8%A1%A8%E8%BE%BE%E5%BC%8F%E6%9B%B4%E6%96%B0): Update("price", gorm.Expr("price * ? + ?", 2, 100))
- [根据子查询进行更新](https://gorm.io/zh_CN/docs/update.html#%E6%A0%B9%E6%8D%AE%E5%AD%90%E6%9F%A5%E8%AF%A2%E8%BF%9B%E8%A1%8C%E6%9B%B4%E6%96%B0)
- 跳过更新钩子函数、并且不修改 UpdateAt 字段,  可以用 [UpdateColumn](https://gorm.io/zh_CN/docs/update.html#%E4%B8%8D%E4%BD%BF%E7%94%A8-Hook-%E5%92%8C%E6%97%B6%E9%97%B4%E8%BF%BD%E8%B8%AA)
- `db.Model(&a).Updates(b)`, 可以在 [Before Update Hook](https://gorm.io/zh_CN/docs/update.html#%E6%A3%80%E6%9F%A5%E5%AD%97%E6%AE%B5%E6%98%AF%E5%90%A6%E6%9C%89%E5%8F%98%E6%9B%B4%EF%BC%9F) 里判断是否用 b 更新了 a
- [用钩子在保存和更新时自动把 password 转成 hash 后的密码](https://gorm.io/zh_CN/docs/update.html#%E5%9C%A8-Update-%E6%97%B6%E4%BF%AE%E6%94%B9%E5%80%BC)
- **可恶 gorm.cn 文档不是最新的, 害我一直在看过期的文档,  应该看 gorm.io**

### 删除相关

#### ➤ 例子

```go
func DeleteRecord() {
   // 删除时,  如果模型中存在 DeletedAt 字段则做逻辑删除 (软删除)
   var c = Character{ID: 233}
   db.Delete(&c)                             // 根据 id=233 删除记录
   db.Delete(&Character{}, 233)              // 根据 id=233 删除记录
   db.Delete(&Character{}, []int{233, 234})  // 根据 id 批量删除记录
   db.Where("name = ?", "Ichigo").Delete(&c) // 根据 id 和 name 删除记录
}
```

#### ➤ [查找被软删除的记录](https://gorm.io/zh_CN/docs/delete.html#%E6%9F%A5%E6%89%BE%E8%A2%AB%E8%BD%AF%E5%88%A0%E9%99%A4%E7%9A%84%E8%AE%B0%E5%BD%95)

#### ➤ [永久删除](https://gorm.io/zh_CN/docs/delete.html#%E6%B0%B8%E4%B9%85%E5%88%A0%E9%99%A4) 

## 原生 SQL

### Raw 和 Exec

#### ➤ 可以自行比较原生 SQL 和 db.Where(),  看哪一套语法更好

```go
func RawQuery() {
    var names []string
    var info CharacterInfo
    var infos []CharacterInfo
    var count int
    db.Raw("SELECT name FROM `character` WHERE id < ?", 100).Scan(&names)               // 扫描至字符串切片
    db.Raw("SELECT id,name,age FROM `character` WHERE id < ?", 100).Scan(&infos)        // 扫描至结构体切片
    db.Raw("SELECT id,name,age FROM `character` WHERE id < ? LIMIT 1", 100).Scan(&info) // 扫描至结构体
    db.Raw("SELECT COUNT(*) FROM `character`").Scan(&count)                             // 扫描至整数
    // Scan() 和 Find() 类似,  但 Scan() 不会触发查询钩子函数 AfterFind
}

func ExecUpdate() {
    db.Exec("UPDATE `character` SET name = ? WHERE id IN ? AND deleted_at IS NULL", "Ichigo", []int{233, 234})
    db.Exec("UPDATE `character` SET age = ? WHERE name = ?", gorm.Expr("age + ?", 1), "ichigo") // SET age = age + 1
}
```

### 使用 sql.Row

```go
func sqlRow() {
   var c Character
   row := db.Model(&Character{}).Select("id", "name").Where("id = ?", 1).Row()
   _ = row.Scan(&c.ID, &c.Name) // 注意选择了 id,name,  扫描时也要依次扫描到两个变量

   rows, err := db.Model(&Character{}).Select("id", "name").Where("`from` = ?", "FF7").Rows()
   CheckError(err)    // 检查错误
   defer rows.Close() // 没错就 defer 关闭资源
   for rows.Next() {
      err := rows.Scan(&c.ID, &c.Name) // 也可以 db.ScanRows(rows, &c)
      CheckError(err)
      fmt.Println(c.ID, c.Name)
   }
}
```

## 关联关系

### Has Many

一个公司拥有多个员工:

- 需要在被拥有者 (Employee) 中设置哪一个字段是外键,  默认为 CompanyID (拥有者类型名+拥有者主键字段名)
- 需要在拥有者 (Company) 中设置哪一个字段作为外键值,  默认为 ID (拥有者的主键字段名)

```go
type Company struct {
    ID   uint   `gorm:"primaryKey"`
    Name string `gorm:"size:191"`
    // foreignKey:CompanyID 表示 Employee 中的 CompanyID 是外键
    // references:ID        表示 db.Create(&c) 时会把公司的 ID 值设置到 c.Employees 中的外键
    Employees []Employee `gorm:"foreignKey:CompanyID; references:ID"`
}

type Employee struct {
    ID        uint   `gorm:"primaryKey"`
    Name      string `gorm:"size:191"`
    CompanyID uint   // 用来做关联的 company_id
    DeletedAt gorm.DeletedAt
}

func 重新建表() {
    _ = db.Exec("DROP TABLE IF EXISTS employee,company;").Error
    _ = db.Migrator().CreateTable(&Employee{}, &Company{})
}

func 关联创建() {
    db.Create(&Company{Name: "妖精的尾巴", Employees: []Employee{ // (1) 插入妖精的尾巴,  得到公司 id
        {Name: "纳兹"}, {Name: "露西"}, {Name: "艾露莎"},         // (2) 设置 CompanyID,  然后插入三个员工
    }})
    db.Create(&Company{Name: "护庭十三番队", Employees: []Employee{
        {Name: "一护"}, {Name: "露琪亚"},
    }})
}

func 关联查询() {
    // 下面的 Preload() 会为每一个公司找到关联的员工,  会执行两条查询:
    // (1) SELECT * FROM company WHERE id < 100
    // (2) SELECT * FROM employee WHERE company_id IN (1,2)
    var cs []Company
    db.Where("id < ?", 100).Preload("Employees").Find(&cs)

    for _, c := range cs {
        for _, e := range c.Employees {
            fmt.Println(c.Name, e.ID, e.Name)
        }
    }
}
```

### Belongs To

公司拥有多个员工,  但从 Employee 的角度来看,  他只属于一个 Company,  下面的类型表示 belongs to 关系  

```go
type Employee struct {
   ID        uint   `gorm:"primaryKey"`
   Name      string `gorm:"size:191"`
   CompanyID uint    // 用来做关联的 company_id
   Company   Company // 同时包含 CompanyID 和 Company
   DeletedAt gorm.DeletedAt
}

func 带关联的查询与创建() {
    // 查询 employee 的同时会把关联的 company 查出来:
    // (1) SELECT * FROM employee WHERE id = 5
    // (2) SELECT * FROM company WHERE id = 2
    var e Employee
    db.Preload("Company").Where("id = ?", 5).Find(&e)
    fmt.Println(e.Name, e.Company.Name)

    // 创建时会先后往 company、employee 插入一行记录,  并设置好员工的 CompanyID
    var e2 = Employee{Name: "Dante", Company: Company{Name: "恶魔五月哭"}}
    db.Create(&e2)
    fmt.Println(e2.Name, e2.CompanyID)
}
```

#### ➤ [如果外键名恰好在拥有者类型中存在，GORM 通常会错误的认为它是 has one 关系](https://gorm.io/zh_CN/docs/belongs_to.html#%E9%87%8D%E5%86%99%E5%BC%95%E7%94%A8) (好 TM 难看懂啊...)

### Has One 与 Belongs To

下面三种关系的共同点都是在 Employee 中放一个 CompanyID 字段

```go
type Employee struct { CompanyID uint; Company Company }   // 员工属于一个公司,  belongs to 关系
type Company struct { Employee Employee }                  // 公司包含一个员工,  has one 关系
type Company struct { Employees []Employee }               // 公司包含多个员工,  has many 关系
```

### Many to Many

```go
type Student struct {
   ID      uint     `gorm:"primaryKey"`
   Name    string   `gorm:"size:191"`
   Courses []Course `gorm:"many2many:student_course"`
}

type Course struct {
   ID       uint      `gorm:"primaryKey"`
   Name     string    `gorm:"size:191"`
   Students []Student `gorm:"many2many:student_course"`
}

func ManyToMany重新建表() {
   _ = db.Exec("DROP TABLE IF EXISTS student,course,student_course;").Error
   _ = db.Migrator().AutoMigrate(&Student{}, &Course{})
   courses := []Course{{ID: 1, Name: "魔法"}, {ID: 2, Name: "武技"}}
   students := []Student{{Name: "克劳德"}, {Name: "爱丽丝"}, {Name: "蒂法"}}
   db.Create(&courses)
   db.Create(&students)
   db.Exec("INSERT INTO student_course (student_id, course_id) VALUES (1,1), (1,2), (2,1), (3,2)")
}

func ManyToMany关联查询() {
   // 查找克劳德, 并加载他上了哪些课
   var s Student
   db.Preload("Courses").Where("name = ?", "克劳德").Find(&s)
   fmt.Println(s.Name, s.Courses)

   // 查找魔法课, 并加载它有哪些学生
   var c Course
   db.Preload("Students").Where("name = ?", "魔法").Find(&c)
   fmt.Println(c.Name, c.Students)
}
```

### 关联查询与关联创建

#### ➤ 会自动创建数据和关联

```go
func 自动创建数据和关联() {
   // INSERT INTO student (name) VALUES ('尤菲')
   // INSERT INTO course (`name`) VALUES ('忍术')
   // INSERT INTO student_course (student_id,course_id) VALUES (4,3)
   var s = Student{Name: "尤菲", Courses: []Course{{Name: "忍术"}}}
   db.Create(&s)
}
```

#### ➤ 可以用 `Preload`、`Joins`、`Association` 加载关联数据,  `Preload` 还支持 [自定义预加载 SQL](https://gorm.io/zh_CN/docs/preload.html#%E8%87%AA%E5%AE%9A%E4%B9%89%E9%A2%84%E5%8A%A0%E8%BD%BD-SQL)

```go
func 查找关联数据() {
    // Preload 使用多条 SQL 查询关联数据,  比如比下面的查询要用三条 SQL
    // 查询名为克劳德的学生、用 Preload 预加载他学过的、并且 id < 100 的课程
    var cloud Student
    db.Preload("Courses", "id < ?", 100).Where("name = ?", "克劳德").Find(&cloud)
    fmt.Println(cloud.Name, cloud.Courses)
    
    // 假如课程 has one 老师,  预加载课程列表的同时预加载 Course 中的 Teacher 字段
    db.Preload("Courses").Preload("Courses.Teacher").Where("name = ?", "克劳德").Find(&cloud)

    // Joins 使用 employee LEFT JOIN company 加载关联数据
    // Joins 只能用于 has one, belongs to 关系
    var ichigo Employee
    db.Joins("Company").Find(&ichigo, 4)
    fmt.Println(ichigo.Name, ichigo.Company.Name)

    // Association 使用 course INNER JOIN student_course ON (course_id = id AND student_id = 1)
    var cs []Course
    db.Model(&Student{ID: 1}).Association("Courses").Find(&cs) // 还可以加 Where() 来筛选 course
    fmt.Println(cs)

    // 题外话,  LEFT JOIN 中 ON 子句与 WHERE 子句的区别?
    // https://stackoverflow.com/a/354094
}
```

#### ➤ 添加关联数据

1. `Append()` 为 many to many、has many 添加新的关联数据；为 has one, belongs to 替换当前的关联数据

2. 参考这里的 [关联操作](https://gorm.io/zh_CN/docs/associations.html#%E6%9B%BF%E6%8D%A2%E5%85%B3%E8%81%94),  还可以删除关联数据、替换关联数据、统计有多少个关联数据、....

```go
func 添加关联数据() {
    var c Company
    db.First(&c, "name = ?", "妖精的尾巴")
    err := db.Model(&c).Association("Employees").Append([]Employee{{Name: "哈比"}})
    CheckError(err)
}
```

## 使用相关

### 使用 Context

#### ➤ 例子

```go
func UseContext() {
    // 设置两秒的超时时间
    ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
    defer cancel()

    // (1) 在单个操作中使用 Context
    var students []Student
    db.WithContext(ctx).Find(&students)
    fmt.Println(students)

    // (2) 在多个操作中使用同一 Context
    var c Character
    tx := db.WithContext(ctx)
    tx.First(&c, "name = ?", "Rukia")
    tx.Model(&c).Update("age", 150)
}
```

#### ➤ 可以在钩子函数例如 `BeforeCreate` 中[访问 Context 对象](https://gorm.io/zh_CN/docs/context.html#Hooks-x2F-Callbacks-%E4%B8%AD%E7%9A%84-Context)

#### ➤ [可以写一个中间件](https://gorm.io/zh_CN/docs/context.html#Chi-%E4%B8%AD%E9%97%B4%E4%BB%B6%E7%A4%BA%E4%BE%8B)、为一个请求中的所有 database 操作设置总的超时时间

### 别忘了处理错误

```go
func DontForgetErrorHandling() {
    var c Character
    // 当 First、Last、Take 方法找不到记录时，GORM 会返回 ErrRecordNotFound 错误
    if err := db.First(&c, "name = ?", "Cloud").Error; err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            fmt.Println("没有找到匹配的记录")
        } else {
            // 系统错误
        }
    }

    // Find 的错误处理则稍微简单些
    if result := db.Find(&c, "name = ?", "Tifa"); result.Error != nil {
        // 系统错误
    }
}
```

### GORM Session 相关

#### ➤ 如果想固定一部分查询条件,  需要新建 Session

```go
func NewGormSession() {
    // 可以用下面三个方法创建 Session
    tx := db.Where("name = ?", "Cloud").Session(&gorm.Session{})          // gorm.Session 有若干配置项
    tx = db.Where("name = ?", "Cloud").WithContext(context.Background())  // 设置超时时间
    tx = db.Where("name = ?", "Cloud").Debug()                            // 把日志级别改成 Info

    var c Character
    tx.Where("age < ?", 100).Find(&c) // name = 'Cloud' AND age < 100
    fmt.Println(c)

    // 下面是错误示例,  不要这么做!  db.Where() 不能复用
    tx = db.Where("name = ?", "Cloud")
    tx.Where("age > ?", 10).Find(&c) // 这条没问题:         name = 'Cloud' AND age > 10
    tx.Where("age > ?", 20).Find(&c) // 这条被上一条污染了: name = 'Cloud' AND age > 10 AND age > 20 
}
```

#### ➤ [GORM Session 有若干配置项](https://gorm.io/zh_CN/docs/session.html)

1. DryRun 用于检查生成的 SQL,  `stmt := db.Session(&gorm.Session{DryRun: true}).First(&user, 1).Statement`
2. 可以开启 PrepareStmt, 能提高一点效率,  [MySQL Prepared Statements](https://dev.mysql.com/doc/refman/5.7/en/sql-prepared-statements.html)
3. 自定义 Logger、跳过钩子函数、禁用嵌套事务、...

### 钩子函数

#### ➤ [参考文档](https://gorm.io/zh_CN/docs/hooks.html)

> 1. Hook 是在创建、查询、更新、删除等操作之前/之后调用的函数。  
>    如果您已经为模型定义了指定的方法，它会在创建、更新、查询、删除时自动被调用。
> 2. 如果任何回调返回错误，GORM 将停止后续的操作并回滚事务。  
> 3. 钩子方法的函数签名应该是 `func(*gorm.DB) error`

#### ➤ BeforeCreate 和 BeforeSave 的区别?

文档上也没写...,  Save/Create 这两个相近的词,  难道没有人好奇他们的区别吗?  
尝试找不同: BeforeSave 在创建和更新时都会调用,  BeforeCreate 只在创建时起作用

#### ➤ 可以用 `Set/Get` 方法[往钩子函数传值](https://gorm.io/zh_CN/docs/settings.html)

### 自定义数据类型

#### ➤ 注意 [gorm.io/datatypes](https://github.com/go-gorm/datatypes) 提供了 JSON、Time 数据类型的支持

#### ➤ 可以实现 Scanner 和 Valuer 两个接口,  [自定义一个类型如何序列化到 database](https://gorm.io/zh_CN/docs/data_types.html)

#### ➤ 注意 GORM 提供了 json 序列化器把结构体保存为 json 字符串

```go
type Name struct {
    First string
    Last  string
}

type Student2 struct {
    ID   uint
    Name Name `gorm:"type:varchar(100); serializer:json"`
}

func 使用json序列化器() {
    _ = db.AutoMigrate(&Student2{})
    db.Create(&Student2{Name: Name{First: "里昂", Last: "肯尼迪"}})
    // INSERT INTO `student2` (`name`) VALUES ('{"First":"里昂","Last":"肯尼迪"}')
}
```

### 用 Scopes 封装查询

#### ➤ [封装一些常用的查询条件以方便复用](https://gorm.io/zh_CN/docs/scopes.html#%E6%9F%A5%E8%AF%A2)

#### ➤ [用 Scopes 实现分页器、动态选择表名](https://gorm.io/zh_CN/docs/scopes.html#%E5%88%86%E9%A1%B5)

#### ➤ [用 Scopes 查到关联的数据, 然后添加筛选条件](https://gorm.io/zh_CN/docs/scopes.html#%E6%9B%B4%E6%96%B0)

## 事务操作

### 禁用默认事务

为了确保数据一致性，GORM 会在事务里执行写入操作（创建、更新、删除）。如果没有这方面的要求，您可以在初始化时禁用它，这将获得大约一点性能提升。(但个人觉得数据一致性比一点性能提升重要)

### 普通事务

```go
func UseTransaction() {
    err := db.Transaction(func(tx *gorm.DB) error {
        // 从这里开始，应该使用 'tx' 而不是 'db'
        if err := tx.Create(&Student{Name: "里昂"}).Error; err != nil {
            return err // 返回任何错误都会回滚事务
        }
        if err := tx.Create(&Student{Name: "吉尔"}).Error; err != nil {
            return err // 返回任何错误都会回滚事务
        }

        return errors.New("haha") // 返回故意的错误, 这会导致回滚之前的插入操作
        panic("haha panic")       // panic 也会触发回滚
    })
    CheckError(err) // db.Transaction() 返回事务中遇到的错误
}
```

### 嵌套事务

有一个坑,  若开启了 prepared statement 那么用嵌套事务时会报错,  mysql driver 说尚未支持 SAVEPOINT 命令
如果 GORM 全局开启了 prepared statement,  那么也无法在特定 Session 中关闭

```go
func 嵌套事务() {
    _ = db.Transaction(func(tx *gorm.DB) error {
        tx.Create(&Student{Name: "一护"})

        // SAVEPOINT 111
        err := tx.Transaction(func(tx *gorm.DB) error {
            tx.Create(&Student{Name: "蓝染"})
            return errors.New("bad guy") // 在嵌套事务中回滚蓝染
        })
        fmt.Println(err)
        // ROLLBACK TO SAVEPOINT 111

        // SAVEPOINT 222
        _ = tx.Transaction(func(tx *gorm.DB) error {
            tx.Create(&Student{Name: "织姬"})
            return nil
        })

        tx.Create(&Student{Name: "露琪亚"})
        return nil
    })
}
```

### 手动事务

#### ➤ 不完整的例子

```go
func 手动事务() {
    // 开始事务,  注意接下来要用 tx 而不是 db
    tx := db.Begin()

    // 在事务中执行一些 db 操作
    tx.Create(&Student{Name: "里昂"})
    tx.Create(&Student{Name: "但丁"})
    tx.Create(&Student{Name: "尼禄"})

    rand.Seed(time.Now().UnixNano())
    if n := rand.Intn(2); n == 0 {
        fmt.Println("遇到错误, 回滚")
        tx.Rollback()
    } else if n == 1 {
        fmt.Println("一切正常, 提交")
        tx.Commit()
    }
}
```

#### ➤ 完整的例子

```go
func 手动事务_完整例子() (err error) {
    // 注意接下来要用 tx 而不是 db
    tx := db.Begin()

    // 处理可能的 panic,  避免 panic 时忘了回滚
    defer func() {
        if r := recover(); r != nil {
            tx.Rollback()
            // 修改函数返回值,  避免 panic-recover 后返回 nil
            switch r := r.(type) {
            case error:
                err = r
            default:
                err = fmt.Errorf("error: %v", r)
            }
        }
    }()

    // 开启事务也会出错吗, 官网的例子说明有这种情况
    if err := tx.Error; err != nil {
        return err
    }

    // 执行 SQL 并检查错误,  遇到错误就回滚
    if err := tx.Create(&Student{Name: "Cloud"}).Error; err != nil {
        tx.Rollback()
        return err
    }
    if err := tx.Create(&Student{Name: "Tifa"}).Error; err != nil {
        tx.Rollback()
        return err
    }

    // 提交事务、并返回提交时遇到的错误
    return tx.Commit().Error
}
```

#### ➤ 另外[手动的嵌套事务](https://gorm.io/zh_CN/docs/transactions.html#SavePoint%E3%80%81RollbackTo)可以用 `SavePoint`、`RollbackTo` 两个方法
