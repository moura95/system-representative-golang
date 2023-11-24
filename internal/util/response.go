package util

type Response struct {
	Code  interface{} `json:"code"`
	Error interface{} `json:"error"`
	Data  interface{} `json:"data"`
}

func SuccessResponse(code, data interface{}, error string) Response {
	return Response{
		Code:  code,
		Data:  data,
		Error: error,
	}
}

func ErrorResponse(code int32, data interface{}, error string) Response {
	return Response{
		Code:  code,
		Data:  data,
		Error: error,
	}
}
