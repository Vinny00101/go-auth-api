package err

type HttpError struct {
	Status  int
	Message string
}

func (e *HttpError) Error() string {
	return e.Message
}

func (e *HttpError) Code() int {
	return e.Status
}

func NewErrorHttp(status int, msg string) *HttpError {
	return &HttpError{
		Status:  status,
		Message: msg,
	}
}
