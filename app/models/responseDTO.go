package models

type ResponseDTO struct {
	Data  interface{} `json:"data"`
	Error *Error      `json:"error"`
}

type Error struct {
	ErrorCode    *int    `json:"code,omitempty"`
	ErrorMessage *string `json:"message,omitempty"`
}

func NewError(code int, message string) *Error {
	return &Error{ErrorCode: &code, ErrorMessage: &message}
}
