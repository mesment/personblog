package error


const (
	SUCCESS = 200
	ERROR = 500
	INVALID_PARAMS = 400

	ERROR_AUTH_CHECK_TOKEN_FAIL = 10001
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT = 10002
	ERROR_AUTH_TOKEN = 10003
	ERROR_AUTH = 10004
)

var MsgFlags = map[int]string {
	SUCCESS :"成功",
	ERROR : "交易失败",
	INVALID_PARAMS : "请求参数错误",
	ERROR_AUTH_CHECK_TOKEN_FAIL : "Token 鉴权失败",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT : "Token已超时",
	ERROR_AUTH_TOKEN : "生成Token失败",
	ERROR_AUTH : "鉴权失败",
}

//根据错误码返回信息
func GetMsg(code int) string  {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[ERROR]
}