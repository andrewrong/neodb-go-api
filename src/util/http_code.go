package util

type HttpStatus struct {
	Code    int
	Message string
	Data    interface{}
}

func NewHttpStatus(code int, msg string, data interface{}) *HttpStatus {
	return &HttpStatus{
		Code:    code,
		Message: msg,
		Data:    data,
	}
}

func NewHttpStatusWithCode(code int) *HttpStatus {
	return NewHttpStatus(code, "", nil)
}

var HttpStatusOK = NewHttpStatusWithCode(200)
var HttpAccept = NewHttpStatus(202, "accepted", nil)

var HttpStatus400 = NewHttpStatus(400, "bad request", nil)
var HttpStatus401 = NewHttpStatus(401, "Unauthorized", nil)
var HttpStatus403 = NewHttpStatus(403, "Forbidden", nil)
var HttpNotFound = NewHttpStatus(404, "Not Found", nil)
var HttpFound = NewHttpStatus(302, "Found", nil)

var HttpStatus500 = NewHttpStatus(500, "Internal Server Error", nil)
