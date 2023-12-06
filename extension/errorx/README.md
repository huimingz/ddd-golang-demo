# errors

通用错误处理

## Usage

```golang
package xxx

import (
	"git.dustess.com/operation-billing/kit/errors"
)

// errors 中定义了常用的OSS系统标准错误，比如参数错误、数据库错误等。
// 当然你可以在自己的服务内定义自己的错误，定义方式如下：
var (
	// ErrUnsupportedAuthorizationType 不支持的授权类型
	ErrUnsupportedAuthorizationType = errors.Error{
		Code: 6001006,
		GRPCCode: codes.Code(6001006), 
		Reason: "不支持的授权类型",
	}
)
```

业务逻辑中使用自定义错误

```go
package biz

import (
	"git.dustess.com/operation-billing/kit/errors"
)

// 使用
if len(dto.Name) > 20 {
    // 可以在通用错误定义的基础上补充业务相关的详情信息，通过WithDetail()的链式调用方式
    return nil, errors.ErrInvalidRequests.WithMessage("跟进人角色名称（Code）长度应小于20bytes")
}

role, err := s.tenantFollowerRepository.AddRole(ctx, entity.NewTenantFollowerRole(dto.Code, dto.Name, dto.Desc))
if err != nil {
	// 强制要求在原始错误出现的地方使用 errors.WithStack 对错误进行封装，保证错误处理的拦
	// 截器能够正确打印堆栈信息，以便事后故障分析。
	//
	// 使用 errors.ErrDBOperation.WithWrap(err) 可将错误封装为标准错误（推荐这么做）。
	// 当然，有时候我们想给客户返回明确的错误描述，可以使用 SetReason 方法，示例如下：
	//  errors.ErrDBOperation.WithReason("详细描述").Wrap(err)
	//
	// 注意：当你捕捉到了一个错误，并且能够处理它，这是只需要使用 warn 级别记录信息。
    return nil, errors.ErrDBOperation.WithWrap(err)
}
```

下面给出了一个完整的错误处理示例：
```go
package main

import (
	"context"
	"git.dustess.com/operation-billing/kit/errors"
	"git.dustess.com/operation-billing/kit/logs"
	"github.com/go-resty/resty/v2"
	"google.golang.org/grpc/codes"
	"time"
)

var ErrExternalHTTP = errors.Error{
	Code:     400,
	Reason:   "请求外部资源失败",
	GRPCCode: codes.Internal,
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	if err := doSomething(ctx); err != nil {
		logs.Error(ctx, "打印错误堆栈信息", logs.Any("error", err))

		// 断言是否为 ErrExternalHTTP 错误类型
		if ErrExternalHTTP.Is(err) {
			logs.Info(ctx, "这是一个外部请求失败时抛出的错误")
		}
	}
}

func doSomething(ctx context.Context) error {
	v, err := FetchSomethingFromInternet(ctx)
	if err != nil {
		// 也可以返回 errors.WithStack(err)
		return errors.WithStack(err)
	}

	logs.Info(ctx, "记录结果", logs.Any("value", v))
	return nil
}

func FetchSomethingFromInternet(ctx context.Context) (string, error) {
	logs.Info(ctx, "fetching something from internet")
	client := resty.New()
	resp, err := client.R().
		SetContext(ctx).
		Get("http://localhost:8881/")
	if err != nil {
		return "", ErrExternalHTTP.WithReason("发生了一些错误").Wrap(err)
	}
	return resp.String(), nil
}
```

下面是输出结果：
```
2022-04-15T14:32:24.060+0800    ERROR   打印错误堆栈信息        {"error": "error: code = 400, reason = 发生了一些错误", "errorVerbose": "error: code = 400, reason = 发生了一些错误\ngit.dustess.com/operation-rors.WithStack\n\t/Users/huimingz/Developer/go/src/dustess.com/kit/errors/errors.go:155\nmain.FetchSomethingFromInternet\n\t/Users/huimingz/Developer/go/src/dustess.com/kit/scratch/log.go:49\nmain.doSomething\n\t/Users/huimingz/Developer/go/src/dustess.com/kit/scratch/log.go:33\nmain.main\n\t/Users/huimingz/Developer/go/src/dustess.com/kit/scratch/log.go:22\nruntime.main\n\t/Users/huimingz/.gvm/gos/go1.18/src/runtime/proc.go:250\nruntime.goexit\n\t/Users/huimingz/.gvm/gos/go1.18/src/runtime/asm_arm64.s:1259"}
2022-04-15T14:32:24.060+0800    INFO    这是一个外部请求失败时抛出的错误
```

一些其它示例：
```golang

func foo(ctx context.Context) (resp interface{}, error) {
	result, err := Invoke(ctx)
	if err != nil {
		// 记录错误栈信息
		return nil, errors.WithStack(err)
		
		// 记录错误栈信息 + 添加自定义信息
		return nil, errors.WithMessage(errors.WithStack(err), "自定义内容") // 不区分前后顺序
		
		// 注意上面的错误都会被识别为默认错误，
		// 带特定标识状态码的错误，需要使用预定义错误。
		// 数据库错误 + 记录错误栈信息
		return nil, errors.NewErrMySQL(err, [reason])
    }
	return result, nil 
}
```

## 日志规范

### 哪些场景需要打印日志
- 系统初始化流程
- 编程语言提示异常
- 业务流程预期不符
- 系统核心角色，组件关键动作
- 作为服务调用方（打印入参、出参）
- 作为服务提供方（打印入参）
- 定时任务运行相关记录

### 日志级别

#### ERROR
影响到程序正常运行、当前请求正常运行的异常情况。如：
1. 打开配置文件失败；
2. 所有第三方对接的异常(包括第三方返回错误码)；
3. 所有影响功能使用的异常，包括:SQLException 、空指针异常以及业务异常之外的所有异常。

#### WARN
告警日志。不应该出现但是不影响程序、当前请求正常运行的异常情况。。

但是一旦出现了也需要关注，因此一般该级别的日志达到一定的阈值之后，就得提示给用户或者需要关注的人了。如：
1. 有容错机制的时候出现的错误情况；
2. 找不到配置文件，但是系统能自动创建配置文件；
3. 即将接近临界值的时候，如 cpu 存储已经超过阈值的 80%，就需要打印 warn 级别的日志。

#### INFO
记录输入输出、程序关键节点等必要信息。平时无需关注，但出问题时可根据INFO日志诊断出问题。
1. Service方法中对于系统/业务状态的变更；
2. 调用第三方时的调用参数和调用结果（入参和出参）；
3. 提供方，需要记录入参；
4. 定时任务的开始执行与结束执行。

#### DEBUG
调试信息，对系统每一步的运行状态进行精确的记录。

#### TRACE
特别详细的服务调用流程各个节点信息。业务代码中，除非涉及到多级服务的调用，否则不要使用(除非有特殊用意，否则请使用DEBUG级别替代)

### 日志打印规范
1. **增删改** 操作需要打印参数日志，方便定位异常业务问题；
2. **条件分支** 需要打印日志：包括条件值以及重要参数；
3. 明确日志打印 **级别** 与包含的 **信息**;
4. **提供方** 服务，建议以 **INFO** 级别记录入参，出参可选;
5. **消费队列消息**，务必打印消息内容;
6. **调用方服务**，建议以 **INFO** 级别记录入参和出参;
7. **运行环境问题**，如网络错误、建议以 WARN 级别记录错误堆栈;
8. **定时任务**，务必打印任务开始时间、结束时间。涉及扫描数据的任务，务必打印扫描范围;
9. 异常信息应该包括两类信息：**案发现场信息**和**异常堆栈信息**。如果不处理，使用 `errorx.WithStack` 对错误进行封装后抛出；
10. 谨慎地记录日志；
11. 生产环境禁止输出 debug 日志;
12. 有选择地输出 info 日志;
13. 可以使用 warn 日志级别来记录用户输入参数错误的情况，避免用户投诉时，无所适从。

    注意日志输出的级别，error 级别只记录系统逻辑出错、异常等重要的错误信息。如非必要，请不要在此场景打出 error 级别。(上述已经说明了 error 与 warn 级别日志的区别)