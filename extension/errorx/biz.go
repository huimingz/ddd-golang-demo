package errorx

var (
	// 系统公共错误（-1， 1000-9999）
	ErrUnknownException = NewErrorWithLevel(20000, "遇到了未知错误", LevelError)
	ErrDBOperation      = NewErrorWithLevel(20001, "数据库出现异常错误", LevelError)
	ErrInternalServer   = NewErrorWithLevel(20002, "服务器内部错误", LevelError) // Internal server error.
	ErrIllegalArgument  = NewErrorWithLevel(20003, "请求参数错误", LevelInfo)
	ErrServerBusy       = NewErrorWithLevel(20004, "服务器繁忙，请稍后重试", LevelError)
	ErrForbidden        = NewErrorWithLevel(20005, "无权访问", LevelInfo)
	ErrInvalidSession   = NewErrorWithLevel(20008, "无效的会话", LevelInfo) // Invalid sessionstore.

	ErrLoginRequired             = NewErrorWithLevel(20100, "需要登录", LevelInfo)
	ErrNotSupportedAuthorization = NewErrorWithLevel(20101, "未支持的认证类型", LevelInfo)
	ErrInvalidAuthorizationCode  = NewErrorWithLevel(20102, "授权码已失效", LevelInfo)
	ErrInvalidAccessToken        = NewErrorWithLevel(20104, "无效的access token", LevelInfo) // Failed to get access token,
	ErrInvalidToken              = NewErrorWithLevel(20105, "无效的令牌", LevelInfo)           // Invalid token provided.
	ErrTokenRequired             = NewErrorWithLevel(20106, "需要令牌", LevelInfo)            // Token required.

	ErrUserInactive                = NewErrorWithLevel(20200, "您的账户已被禁用", LevelInfo)
	ErrIncorrectUsernameOrPassword = NewErrorWithLevel(20201, "用户名或密码错误", LevelInfo) // The username and/or the password you entered is incorrect.

	ErrResourceNotFound      = NewErrorWithLevel(20300, "资源未找到", LevelInfo)     // Resource not found.
	ErrResourceAlreadyExists = NewErrorWithLevel(20301, "资源已存在", LevelInfo)     // Resource already exists.
	ErrResourceSingular      = NewErrorWithLevel(20302, "资源不唯一", LevelInfo)     // Resource is not singular.
	ErrResourceNotLoaded     = NewErrorWithLevel(20303, "资源未加载", LevelInfo)     // Resource is not loaded.
	ErrResourceConstraint    = NewErrorWithLevel(20304, "资源约束错误", LevelInfo)    // Resource constraint error.
	ErrTaskQueueFull         = NewErrorWithLevel(20305, "任务队列已满", LevelWarning) // Task queue is full.

	ErrExternalServiceError = NewErrorWithLevel(20400, "外部服务错误", LevelError) // External service error.

	ErrRejected = NewErrorWithLevel(20500, "请求被拒绝", LevelInfo)
)
