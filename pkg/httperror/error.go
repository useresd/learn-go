package httperror

type httpError struct {
	err  error
	code int
}

func New(err error, code int) *httpError {
	return &httpError{err: err, code: code}
}

func (e *httpError) Error() string {
	return e.err.Error()
}
