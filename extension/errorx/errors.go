package errorx

import (
	"fmt"
	"strings"

	stderr "github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	// DefaultSuccessCode 默认成功响应码
	DefaultSuccessCode = 0
	// DefaultErrorCode 默认错误码（如果没有具体错误码code定义，则使用该值）
	DefaultErrorCode = -1

	LevelDebug Level = iota
	LevelInfo
	LevelWarning
	LevelError
	LevelCritical
)

type Level int

type Error struct {
	// http response code, will be used in http response body.
	//
	// example: 404
	// if not set, default to DefaultSuccessCode
	Code int // 错误响应码

	// a human-readable reason of the error.
	//
	// example: "the user does not exist"
	// required: true
	Reason string // 通用错误信息，展示给用户

	// the error details (the error's message).
	//
	// example: "cannot create the user: the user already exists"
	Message string // 错误详情

	GRPCCode codes.Code // [可选]自定义grpc Status Code
	err      error      // 错误原始信息
	level    Level      // 错误级别
}

func (e Error) Error() string {
	buffer := strings.Builder{}
	buffer.WriteString(fmt.Sprintf("code: %d, reason: %s", e.Code, e.Reason))

	if e.Message != "" {
		buffer.WriteString(fmt.Sprintf(", message: %s", e.Message))
	}
	if e.err != nil {
		buffer.WriteString(fmt.Sprintf(", detail: %s", e.err.Error()))
	}

	return buffer.String()
}

func (e Error) Level() Level {
	return e.level
}

// Unwrap returns the underlying error.
func (e Error) Unwrap() error {
	return e.err
}

func (e Error) Wrap(err error) error {
	e.err = err
	return WithStack(e)
}

func (e Error) WithWrap(err error) error {
	if err == nil {
		return WithStack(&e)
	}

	e.err = err
	e.Message = err.Error()
	return WithStack(&e)
}

// WithMessage fork a new Error object and add Detail message
func (e Error) WithMessage(detail string) Error {
	e.Message = detail
	return e
}

// WithMessageF fork a new Error object and add Detail message
func (e Error) WithMessageF(format string, args ...interface{}) Error {
	e.Message = fmt.Sprintf(format, args...)
	return e
}

func (e Error) WithReason(reason string) Error {
	e.Reason = reason
	return e
}

func (e Error) WithReasonF(format string, args ...interface{}) Error {
	e.Reason = fmt.Sprintf(format, args...)
	return e
}

func (e Error) WithError(err error) Error {
	if err == nil {
		return e
	}

	e.Message = err.Error()
	e.err = WithStack(err)
	return e
}

func (e *Error) SetGRPCCode(code codes.Code) *Error {
	e.GRPCCode = code
	return e
}

func (e *Error) SetReason(reason string) *Error {
	e.Reason = reason
	return e
}

func (e *Error) SetMessage(message string) *Error {
	e.Message = message
	return e
}

func (e Error) Cause() error {
	return e.err
}

func (e Error) Is(err error) bool {
	if err == nil {
		return false
	}

	switch te := err.(type) {
	case *Error:
		return e.Code == te.Code && e.GRPCCode == te.GRPCCode
	case Error:
		return e.Code == te.Code && e.GRPCCode == te.GRPCCode
	case grpcError:
		return te.GRPCStatus().Code() == e.GRPCCode
	default:
		// unwrap the error
		for {
			err = stderr.Unwrap(err)
			if err == nil {
				return false
			}
			if te, ok := err.(*Error); ok {
				return e.Code == te.Code && e.GRPCCode == te.GRPCCode
			}
			if te, ok := err.(Error); ok {
				return e.Code == te.Code && e.GRPCCode == te.GRPCCode
			}
			if te, ok := err.(grpcError); ok {
				return te.GRPCStatus().Code() == e.GRPCCode
			}
		}
	}
}

func (e Error) As(value any) bool {
	switch v := value.(type) {
	case *Error:
		*v = e
		return true
	default:
		return false
	}
}

// NewError return a new Error object with error level
func NewError(code int, reason string) Error {
	return NewErrorWithLevel(code, reason, LevelError)
}

func NewErrorWithLevel(code int, reason string, level Level) Error {
	return Error{Code: code, Reason: reason, GRPCCode: codes.Code(code), level: level}
}

func NewErrorWithMessage(code int, msg string, message string) Error {
	return Error{Code: code, Reason: msg, Message: message, GRPCCode: codes.Code(code)}
}

// Explode 分解指定的错误
// 返回错误码和错误信息
func Explode(err error) (code int, grpcCode codes.Code, message string, detail string) {
	if err == nil {
		return DefaultSuccessCode, codes.OK, "success", ""
	}

	if e, ok := err.(*Error); ok {
		return e.Code, e.GRPCCode, e.Reason, e.Message
	}
	if e, ok := err.(Error); ok {
		return e.Code, e.GRPCCode, e.Reason, e.Message
	}

	// grpc status code
	s, ok := status.FromError(err)
	if ok {
		return int(s.Code()), s.Code(), s.Message(), err.Error()
	}

	detail = err.Error()
	for {
		err = stderr.Unwrap(err)
		if err == nil {
			return ErrUnknown.Code, ErrUnknown.GRPCCode, ErrUnknown.Reason, detail
		}
		if te, ok := err.(*Error); ok {
			return te.Code, te.GRPCCode, te.Reason, te.Message
		}
	}
}

// IsBizFault 是否业务错误
func IsBizFault(err error) bool {
	if err == nil {
		return false
	}

	switch te := err.(type) {
	case *Error:
		return te.err == nil
	case Error:
		return te.err == nil
	default:
		return true
	}
}

// WithStack mirrors the WithStack method of the stderr package.
// It adds a stack trace to the error if it does not already have one.
func WithStack(err error) error {
	if e, ok := err.(StackTracer); ok && len(e.StackTrace()) > 0 {
		return err
	}

	return stderr.WithStack(err)
}

func WithMessage(err error, msg string) error {
	return stderr.WithMessage(err, msg)
}

func WithMessageF(err error, format string, args ...interface{}) error {
	return stderr.WithMessagef(err, format, args...)
}

type StackTracer interface {
	StackTrace() stderr.StackTrace
}

// grpcError GRPC错误类型
type grpcError interface {
	GRPCStatus() *status.Status
}

type As interface {
	As(any) bool
}

type LevelContainer interface {
	Level() Level
}

func GetErrorLevel(err error, defaultLevel Level) Level {
	if err == nil {
		return defaultLevel
	}

	if e, ok := err.(LevelContainer); ok {
		return e.Level()
	}

	return defaultLevel
}
