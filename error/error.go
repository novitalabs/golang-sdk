package error

import (
	"fmt"
)

const (
	CodeNormal               = 0  // normal return code
	CodeInternalError        = -1 // internal server error
	CodeInvalidParam         = 1  // invalid parameter in request
	CodeModelNotExist        = 2  // model(e.g. checkpoint) not exist
	CodeTaskIdNotExist       = 3  // task_id not exist
	CodeInvalidAuth          = 4  // key is invalid
	CodeParamRangeOutOfLimit = 6  // parameter(e.g. height\size) out of range
	CodeCostBalanceFailure   = 7  // balance is not enought
	CodeSamplerNotExist      = 8  // sampler not found
	CodeNotSupport           = 10 // feature not supported
	CodePromptIllegal        = 11 // prompt is illegal(e.g. child porn)
)

type Error struct {
	Code int
	Msg  string
}

func New(code int, msg string) *Error {
	return &Error{
		Code: code,
		Msg:  msg,
	}
}

func (o Error) Error() string {
	return fmt.Sprintf("API returns error, code = %d, msg = %s, please refer to error/error.go for more details",
		o.Code, o.Msg)
}
