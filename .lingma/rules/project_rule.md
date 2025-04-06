# Standard Operating Procedures (SOP) Manual

## 项目结构

```
cmd/
├── palace/
├── houyi/
├── rabbit/
├── deploy/
├── docs/
├── pkg/
├── proto/
├── .env.palace.example
├── .env.houyi.example
├── .env.rabbit.example
├── .gitignore
├── Makefile
├── README.md
├── go.mod
├── go.sum
```

服务器端语言是Go，入口在 `cmd/` 目录下

项目由多个服务构成，每个服务都有自己的 `main.go` 和 `wire.go`，所以如果你有新的服务需要到 `cmd/` 目录下创建对应的目录，然后创建 `main.go` 和 `wire.go` 文件，然后到 `Makefile` 中添加对应的命令。目前有`houyi`, `palace`, `rabbit` 三个服务

- `palace` 是主服务，用于管理整个moon系统，调度各个子服务， 并提供HTTP 和GRPC API接口
- `houyi` 是 `palace` 的子服务， 用于处理策略规则从而生成告警事件， 并推送到`palace` 服务
- `rabbit` 是 `palace` 的子服务， 消息服务， 用于发送 `palace` 的告警消息或者普通消息。

服务整体分为几个层：
1. `server`：包括gRPC和HTTP入口（HTTP-Gateway最终也是落到gRPC的）
2. `service`：gRPC之后的入口，主要做一些参数校验，参数从gRPC对象转成bo对象
    1. `build`: 存在`do` `bo` `proto` 的转换， 都在这里按照模块实现，出现数据转换的类型， 都需要来这里定义转换方法， 按模块区分文件
3. `biz`：业务逻辑层，主要是业务逻辑处理，比如调用外部服务、数据处理等，大部分逻辑都集中在这里
    1. `bo`: 业务对象， 用于和 `service` 层交互， 从 `service` 层接收参数， 并返回给 `service` 层, 方法入参数大于1个的， 需要用 `bo` 对象来封装
    2. `do`: 数据库对象， 用于和 `data` 层交互， 从 `data` 层接收参数， 并返回给 `data` 层
    3. `repository`: 仓库层, 纯接口， 用于封装 `data` 层的操作， 提供给 `biz` 层使用
    4. `vobj`: 值对象，表达业务状态、类型等
4. `data`：数据层，主要是数据的读写，包括数据库、外部服务等，这里接受 `bo`, `do` 对象（在 `internal/biz/bo/*.go`），但是返回的话需要返回 `do` 对象（在 `internal/biz/do/*.go`），数据库（目前gorm）和外部的服务的对象独立不可透传回 `biz` 和 `service` 层
    1. `impl`: 实现层， 用于实现 `repository` 层的操作， 并返回 `do`, `bo` 对象
    2. `query`: 查询层， 用于查询数据， 并返回 `do`, `bo` 对象, 由gorm gen生成， 例如生成`palace`的`query`文件， `make gen-palace`
5. `helper`: 辅助层， 用于辅助当前服务的逻辑， 例如一些内部的中间件、业务工具函数等

## Makefile
项目大量地方涉及自动生成代码的应用，所有涉及的命令都在 `Makefile` 里了，比如：
- `make all app=<palace|houyi|rabbit>`: 生成所有的代码
- `conf-<app>`: 生成配置
- `make api app=<palace|houyi|rabbit>` : 从 `proto/**/*.proto` 生成对应的 `*.go` 文件

## 配置
配置在 `config` 目录下，主要是 `*.yaml` 文件，配置里的一些敏感的是需要定义在 `.env.<app>.example` 里的。我们是用 `cmd/<app>/internal/conf/conf.proto` 定义的配置，然后通过 `make conf-<app>` 生成代码的

## API
全新创建一个API需要涉及这么几个步骤：
1. 到 `proto/api/<app>/*.proto` 中定义
2. 公共的API定义在 `proto/api/common/*.proto` 中
3. 然后运行 `make api app=<palace|houyi|rabbit>` 生成对应的代码
4. 在 `server` 层 `provider_set.go` 文件中注册对应的服务
5. 在 `service` 层 `provider_set.go` 文件中实现对应的方法
6. 在 `biz` 层 `provider_set.go` 文件中实现对应的逻辑
7. 在 `biz.repository` 层中实现定义对应接口
8. 在 `data.impl` 层 `provider_set.go` 文件中实现对应的数据操作（有可能原来就有存在的，需要自己判断）
9. 在 `data` 层 `provider_set.go` 文件中实现对应的查询操作（有可能原来就有存在的，需要自己判断）

## 错误
错误码在 `proto/merr/err.proto` 中定义，然后通过 `make errors` 生成对应的错误码。在 `biz` 层应该拦截所有的错误，并返回对应的 `merr` 错误给到 `service` 层，这样避免把错误信息暴露给到客户端

## 数据库
目前我们使用的是 `mysql` ，Go里用的是 `gorm` 这个ORM，所有的数据操作都是通过 `gorm` 这个库来进行的，所以如果有新的表需要创建，需要在 `internal/biz/do/<system|event|team>/` 下创建对应的 `*.go` 文件，然后运行 `make gen-palace` 生成对应的代码，然后在 `data` 层进行对应的实现。

`do` 里有一个 `base.go` 有需要的话可以用

## 测试
每次完成所有任务后请执行一下 `make all app=<palace|houyi|rabbit>` ，然后尝试编译一下 `make build app=<palace|houyi|rabbit>` ，确保都没有问题，如果有问题要解决

## 代码注释
通常不需要代码注释，更多是用规范的命令来表现意图，不过有一些地方必须要有注释的**请用英文注释**

# 开发篇

本项目采用mini [DDD](https://www.google.com/search?q=DDD) 设计思想，总共分为以下几个模块

* API: 接口层，包含了HTTP、gRPC等接口定义
    * proto
    * pb
* service: 服务入口，承接了API(Server)和biz层，通常做一些参数校验、转换等工作
* biz: 业务逻辑层，包含了业务逻辑的实现
    * bo
    * do
    * repository
    * vobj
* data: 数据访问层，负责包括数据库、外部数据和外部服务等资源的处理
    * cache
    * db
    * impl

## 依赖倒置

正确的调用关系是：

`service` -> `biz` -> `repository` -> `impl` -> `data`

需要注意的是，每一层都会涉及到数据结构的转置，也就是每层只关心自己收到的数据结构，如：

- `service` 层只关心 `pb` 包中的数据结构，也就是 `proto` 生成的 `req` 和 `resp`，当 `service` 调用 `biz` 时，会将 `pb`
  转成 `bo`

## 事务管理

> biz 层使用事务管理

* repository 定义

```go
type Transaction interface {
  MainExec(ctx context.Context, fn func (ctx context.Context) error) error
  BizExec(ctx context.Context, fn func (ctx context.Context) error) error
  EventExec(ctx context.Context, fn func (ctx context.Context) error) error
}
```

# 开发者必看

## 开发规范

### 1. 包导入规范

引入包的时候， 不管是否只有一个外部包还是多个外部包，统一使用如下写法

```go
import (
	"fmt"
)
```

在同时包含多方包导入时候， 导入顺序按照以下顺序导入， 中间用空行分隔

```go
import (
  // 空白标识符导入
  _ "gorm.io/driver/sqlite"
  
  // 标准库
  "fmt"
  
  // 第三方包
  "github.com/xxxx/xxxx"
  
  // 项目内部包
  "github.com/moon-monitor/moon/internal/biz"
)
```

如果有多方包命名冲突， 可以使用如下写法

```go
import (
    // 内部http包
    nethttp "net/http"
	
    // 第三方http包
    transporthttp "github.com/go-kratos/kratos/v2/transport/http"
)
```

### 2. 命名规范

* 统一按照Go官方命名规则（原则是Goland编辑器不能有警告⚠️）

* 对于不确定是否导出的变量， 统一使用小写字母开头，只有明确导出的，才使用大写字母开头

* 对于实现接口，对应的New方法应返回该接口，从而屏蔽实现细节，对应实现采用结构体采用小写字母开头方式

    ```go
    // InterfaceName 接口注释
    type InterfaceName interface {
        // MethodName 方法注释
        MethodName(param1, param2 string) (string, error)
    }
  
  
    type interfaceNameImpl struct {}
  
    // NewInterfaceName 创建接口实现
    func NewInterfaceName() InterfaceName {
        return &interfaceNameImpl{}
    }
    ```

* 对于出现频率超过`2次`以上的常量（`数字`、`字符串`）等，应统一定义为常量，可借鉴`vobj`定义方式


### 3. 注释规范

* 函数注释

```go
// FuncName 函数注释
func FuncName(param1, param2 string) (string, error) {
    return "", nil
}
```

多行的场景, 后面行缩进一字符， 第一行与后面空行间隔

```go
// FuncName 函数注释
// 
//  param1: 参数1注释
//  param2: 参数2注释
//  return1: 返回值1注释
//  return2: 返回值2注释
func FuncName(param1, param2 string) (string, error) {
    return "", nil
}
```

* 变量注释

```go
var (
	// varName 变量注释
	varName string
	// varName2 变量注释
	varName2 string
)
```

* 结构体注释

```go
// StructName 结构体注释
type StructName struct {
	// fieldName 字段注释
	fieldName string
	// fieldName2 字段注释
	fieldName2 string
}
```

* 接口注释

```go
// InterfaceName 接口注释
type InterfaceName interface {
	// MethodName 方法注释
	MethodName(param1, param2 string) (string, error)
}
```

* 方法注释

```go
// MethodName 方法注释
func (s *StructName) MethodName(param1, param2 string) (string, error) {
    return "", nil
}
```

* 常量注释

```go
// ConstName 常量注释
const (
  // ConstName 常量注释
  ConstName = "constName"
  // ConstName2 常量注释
  ConstName2 = "constName2"
  // ConstName3 常量注释
  ConstName3 = "constName3"
)

```

**特殊说明， 项目中vobj里面的常量， 注释是放在行尾，因为会通过stringer工具生成，对应的方法**

### 4. 函数定义规范

原则上函数或者方法，第一参数都是context（不涉及外部调用或者调用链中涉及上下文管理的工具类方法除外），如果有return error， error放在最后

```go
func FuncName(ctx context.Context, param1) (string, error) {
	return "", nil
}
```

### 5. 结构体方法

接收器名一般为结构体名的首字母小写，比如结构体为User，方法为GetName，接收器为u，需为单字符，不可以结构体名称各个单词缩写命名, 也不可使用this、self等

```go
type User struct {}

func (u *User) GetName() string {
	return ""
}
```

## 外部库引入规范

在开发某些功能时， 需要引入对应的外部实现来完成， 需要团队评估该外部依赖，同意后方可引入

对于简单功能的外部包，建议是内部自己实现, 减少外部依赖，减少维护成本

