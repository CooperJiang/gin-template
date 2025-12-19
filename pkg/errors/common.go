package errors

import "net/http"

var (
	ErrUnauthorized      = New(CodeUnauthorized, "unauthorized")
	ErrForbidden         = New(CodeForbidden, "forbidden")
	ErrNotFound          = New(CodeNotFound, "not found")
	ErrValidation        = New(CodeValidationFailed, "validation failed")
	ErrInternalError     = New(CodeInternal, "internal error")
	ErrRateLimitExceeded = New(CodeRateLimited, "rate limit exceeded")
	ErrInvalidParameter  = New(CodeInvalidParameter, "invalid parameter")
)

var (
	ErrUserNotFound      = New(CodeUserNotFound, "user not found")
	ErrWrongPassword     = New(CodeWrongPassword, "wrong password")
	ErrUserDisabled      = New(CodeUserDisabled, "user disabled")
	ErrUserExists        = New(CodeUserExists, "user already exists")
	ErrInvalidAuthToken  = New(CodeInvalidAuthToken, "invalid auth token")
	ErrExpiredAuthToken  = New(CodeExpiredAuthToken, "auth token expired")
	ErrInvalidVerifyCode = New(CodeInvalidVerifyCode, "invalid verification code")
	ErrEmailExists       = New(CodeEmailExists, "email already exists")
)

var (
	ErrDBConnection = New(CodeDBConnectionFailed, "database connection failed")
	ErrQueryFailed  = New(CodeQueryFailed, "query failed")
	ErrDBNoRecord   = New(CodeDBNoRecord, "record not found")
	ErrDBDuplicate  = New(CodeDBDuplicate, "duplicate record")
)

func NewFromHTTPStatus(statusCode int, detail string) *Error {
	var code ErrorCode

	switch statusCode {
	case http.StatusBadRequest:
		code = CodeInvalidParameter
	case http.StatusUnauthorized:
		code = CodeUnauthorized
	case http.StatusForbidden:
		code = CodeForbidden
	case http.StatusNotFound:
		code = CodeNotFound
	case http.StatusMethodNotAllowed:
		code = CodeMethodNotAllowed
	case http.StatusConflict:
		code = CodeConflict
	case http.StatusRequestTimeout:
		code = CodeTimeout
	case http.StatusTooManyRequests:
		code = CodeRateLimited
	case http.StatusServiceUnavailable:
		code = CodeServiceUnavailable
	default:
		code = CodeInternal
	}

	return New(code, detail)
}

func NewBadRequest(detail string) *Error {
	return New(CodeInvalidParameter, detail)
}

func NewUnauthorized(detail string) *Error {
	return New(CodeUnauthorized, detail)
}

func NewForbidden(detail string) *Error {
	return New(CodeForbidden, detail)
}

func NewNotFound(detail string) *Error {
	return New(CodeNotFound, detail)
}

func NewConflict(detail string) *Error {
	return New(CodeConflict, detail)
}

func NewInternal(detail string) *Error {
	return New(CodeInternal, detail)
}
