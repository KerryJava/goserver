package base

import "github.com/gorilla/rpc/v2/json2"

const (
	E_LOGIN = 9000
	E_SQL
	E_PARAMS
	E_SYSTEM
)

var (
	ErrSystem   = &json2.Error{Code: E_SYSTEM, Message: "系统错误"}
	ErrParams   = &json2.Error{Code: E_PARAMS, Message: "参数错误"}
	ErrSQLQuery = &json2.Error{Code: E_SQL, Message: "database error"}
	ErrLogin    = &json2.Error{Code: E_LOGIN, Message: "account error"}
)
