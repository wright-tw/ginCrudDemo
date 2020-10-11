package server

type Response struct {
	Code    int32                  `json:"code"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
}

func MakeResponse(code int32, message string, data map[string]interface{}) Response {
	var response Response
	response.Data = data
	response.Message = message
	response.Code = code
	return response
}

func EmptyData() map[string]interface{} {
	return map[string]interface{}{}
}
