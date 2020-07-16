package controllers

type BaseController struct{}

type Response struct {
	Code    int32                  `json:"code" default=1`
	Message string                 `json:"message" default=""`
	Data    map[string]interface{} `json:"data"`
}

func (this *BaseController) MakeResponse(code int32, message string, data map[string]interface{}) Response {
	var response Response
	response.Data = data
	response.Message = message
	response.Code = code
	return response
}

func (this *BaseController) EmptyData() map[string]interface{} {
	return map[string]interface{}{}
}
