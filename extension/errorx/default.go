package errorx

// WIP错误码区段一览：https://confluence.dustess.com/pages/viewpage.action?pageId=13612568
// 计费错误码区段一览：https://doc.weixin.qq.com/sheet/e3_APQAxwbKANooa0I91wUTy6SbebrKo?scode=ACUArAcBAAkjDgb7DgAPQAxwbKANo

var (
	// WebGateWayCodeTable 定义了一些特殊的错误码，用于网关层面的错误码映射。
	// Note: grpc的错误码是以uint32类型存储的，所以在定义默认类型时不能使用负数。
	WebGateWayCodeTable = map[int32]int32{
		10000: -10000, // 通用错误码
		11001: -11001, // MySQL数据库错误
		12001: -12001, // Mongo数据库错误
		16001: -16001, // 对象存储错误
	}

	// ErrUnknown 通用系统错误，用于在接口返回时兜底（包括非错误码及负错误码），不向外暴露系统错误细节
	ErrUnknown = Error{
		Code:     10000, // 会自动转换为-10000暴露给前端
		Reason:   "出现了未知错误",
		GRPCCode: 10000, // 会自动转换为-10000暴露给前端
		level:    LevelError,
	}
	// ErrInvalidParam 通用参数错误
	ErrInvalidParam = Error{
		Code:     10001,
		Reason:   "参数错误",
		GRPCCode: 10001,
		level:    LevelInfo,
	}
	ErrMySQL = Error{
		Code:     11001,
		Reason:   "数据库异常",
		GRPCCode: 11001,
		level:    LevelError,
	}
	ErrMongoDB = Error{
		Code:     12001,
		Reason:   "MongoDB异常",
		GRPCCode: 12001,
		level:    LevelError,
	}
	ErrRPCFailed = Error{
		Code:     14001,
		Reason:   "RPC调用失败",
		GRPCCode: 14001,
		level:    LevelError,
	}
	ErrObjectStorageService = Error{
		Code:     16001,
		Reason:   "对象存储服务异常",
		GRPCCode: 16001,
		level:    LevelError,
	}
	ErrGRPCInterceptor = Error{
		Code:     17001,
		Reason:   "GRPC拦截器异常",
		GRPCCode: 17001,
		level:    LevelError,
	}
	ErrUnauthorized = Error{
		Code:     10105,
		Reason:   "未授权的用户",
		GRPCCode: 10105,
		level:    LevelInfo,
	}
	ErrUnauthenticated = Error{
		Code:     10106,
		Reason:   "未认证的请求",
		GRPCCode: 10106,
		level:    LevelInfo,
	}
)

// http 相关错误
var (
	ErrRequestParmas = Error{
		Code:     15001,
		Reason:   "请求参数错误",
		GRPCCode: 15001,
		level:    LevelInfo,
	}
)

// 用户域错误码（60010000 - 600019999），放到这里的目的是让拦截器能够对用户域的
// 错误码进行拦截，以便在拦截器中进行特殊处理。
//
// 大多数情况下，用户域的错误都应该以 INFO 级别记录，不应该记录一个 ERROR 级别的
// 错误，因为这样会导致不必要的错误告警，从而影响开发人员的判断。比如用户认证失败、
// 用户未被授权等，都应该以 INFO 级别记录，而不是 ERROR 级别。
var (
	ErrInvalidAuthenticationCode = Error{
		Code:     60010001,
		Reason:   "授权凭证无效或已过期",
		GRPCCode: 60010001,
		level:    LevelInfo,
	}
	ErrRequiredAuthenticationCode = Error{
		Code:     60010002,
		Reason:   "授权凭证不能为空",
		GRPCCode: 60010002,
		level:    LevelInfo,
	}
	ErrInvalidAuthenticationState = Error{
		Code:     60010003,
		Reason:   "授权状态无效",
		GRPCCode: 60010003,
		level:    LevelInfo,
	}
	ErrUnsupportedAuthenticationMethod = Error{
		Code:     60010004,
		Reason:   "不支持的授权方式",
		GRPCCode: 60010004,
		level:    LevelInfo,
	}
	ErrOAuth2AuthorizationFailed = Error{
		Code:     60010005,
		Reason:   "OAuth2认证失败",
		GRPCCode: 60010005,
		level:    LevelWarning,
	}
	ErrOAuth2UserInfoFailed = Error{
		Code:     60010006,
		Reason:   "OAuth2获取用户信息失败",
		GRPCCode: 60010006,
		level:    LevelWarning,
	}
	ErrUserDisabled = Error{
		Code:     60010011,
		Reason:   "用户已禁用",
		GRPCCode: 60010011,
		level:    LevelInfo,
	}
	ErrInvalidCaptcha = Error{
		Code:     60010012,
		Reason:   "验证码错误",
		GRPCCode: 60010012,
		level:    LevelInfo,
	}
	ErrInvalidUsernameOrPassword = Error{
		Code:     60010013,
		Reason:   "用户名或密码错误",
		GRPCCode: 60010013,
		level:    LevelInfo,
	}
)

func NewErrMySQL(err error) error {
	return ErrMySQL.WithWrap(err)
}

func NewErrRPCFailed(err error) error {
	return ErrRPCFailed.WithWrap(err)
}

func NewErrObjectStorageService(err error) error {
	return ErrObjectStorageService.WithWrap(err)
}

func NewInvalidParam(err error) error {
	return ErrInvalidParam.WithWrap(err)
}

func NewErrUnknown(err error) error {
	return ErrUnknown.WithWrap(err)
}

func NewErrMongoDB(err error) error {
	return ErrMongoDB.WithWrap(err)
}
