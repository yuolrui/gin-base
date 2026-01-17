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
)
