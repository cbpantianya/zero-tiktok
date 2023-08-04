package error

import (
	"errors"
	"fmt"
)

// 统一错误处理

// Error 错误结构体
type Error struct {
	Code  int    `json:"code"` // 外部Code
	Msg   string `json:"msg"`  // 外部Msg
	Inner error  `json:"-"`    // 内部错误
}

func (e Error) Error() string {
	// code to string
	return fmt.Sprintf("code: %d, msg: %s, inner: %v", e.Code, e.Msg, e.Inner)
}

// NewError 新建错误
func NewError(code int, msg string, inner error) *Error {
	return &Error{Code: code, Msg: msg, Inner: inner}
}

// 错误列表
var (
	// ErrParam 参数错误
	ErrParam = NewError(40001, "参数错误", errors.New("parameter error"))
	// ErrDB 数据库错误（链接错误，查询失败，不包括空数据错误）
	ErrDB = NewError(50001, "数据库错误", errors.New("database error"))
)
