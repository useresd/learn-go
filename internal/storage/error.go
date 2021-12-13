package storage

type ErrorRecordNotFound struct {
	Message string
}

func (e *ErrorRecordNotFound) Error() string {
	return e.Message
}
