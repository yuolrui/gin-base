package errno

const (
	OK = 0

	// common
	InvalidParam = 10001
	Unauthorized = 10002
	Forbidden    = 10003
	InternalErr  = 10005

	// user
	UserNotExist = 20001

	ParamErr = 20002
)

// AllCodes 返回所有错误码列表，用于统一校验 i18n
func AllCodes() []int {
	return []int{
		OK,
		InvalidParam,
		Unauthorized,
		Forbidden,
		InternalErr,
		UserNotExist,
	}
}
