package errno

import "net/http"

type AppError struct {
	Code       int
	HTTPStatus int
}

func (e *AppError) Error() string {
	return "app error"
}

func New(code int, httpStatus int) *AppError {
	return &AppError{
		Code:       code,
		HTTPStatus: httpStatus,
	}
}

var (
	ErrInvalidParam = New(InvalidParam, http.StatusBadRequest)
	ErrUnauthorized = New(Unauthorized, http.StatusUnauthorized)
	ErrForbidden    = New(Forbidden, http.StatusForbidden)
	ErrInternal     = New(InternalErr, http.StatusInternalServerError)

	ErrUserNotExist = New(UserNotExist, http.StatusNotFound)
)
