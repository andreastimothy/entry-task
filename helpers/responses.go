package helpers

type SuccessResponses struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type FailedResponses struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (err FailedResponses) Error() string {
	return err.Message
}

func SuccessResponse(code int, message string, data interface{}) SuccessResponses {
	response := SuccessResponses{
		Code:    code,
		Message: message,
		Data:    data,
	}
	return response
}

func FailedResponse(code int, message string) FailedResponses {
	response := FailedResponses{
		Code:    code,
		Message: message,
	}
	return response
}