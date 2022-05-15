// Package errmsg 错误码信息
package errmsg

const (
	SUCCESS = 200
	ERROR   = 500

	// ErrorUsernameUsed code = 100... 开头的是用户模块的错误
	ErrorUsernameUsed   = 1001
	ErrorPasswordWrong  = 1002
	ErrorUserNotExist   = 1003
	ErrorTokenNotExist  = 1004
	ErrorTokenExpired   = 1005
	ErrorTokenWrong     = 1006
	ErrorTokenTypeWrong = 1007
	// code = 200... 开头的是文章模块的错误

	// code = 300... 开头的是分类模块的错误

)

var codeMsg = map[int]string{
	SUCCESS:             "OK",
	ERROR:               "FAIL",
	ErrorUsernameUsed:   "用户名已经被使用",
	ErrorPasswordWrong:  "密码错误",
	ErrorUserNotExist:   "用户不存在",
	ErrorTokenNotExist:  "token不存在",
	ErrorTokenExpired:   "token已经过期",
	ErrorTokenWrong:     "token错误",
	ErrorTokenTypeWrong: "token格式错误",
}

func GetErrMsg(code int) string {
	return codeMsg[code]
}
