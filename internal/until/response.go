package until

import "emaction/internal/model"

// OkWithData 成功响应带数据
func OkWithData(data interface{}) model.APIResponse {
	return model.APIResponse{
		Code: 0,
		Msg:  "success",
		Data: data,
	}
}

// Ok 成功响应不带数据
func Ok() model.APIResponse {
	return model.APIResponse{
		Code: 0,
		Msg:  "success",
	}
}

// FailWithData 失败响应带数据
func FailWithData(data interface{}) model.APIResponse {
	return model.APIResponse{
		Code: 1,
		Msg:  "fail",
		Data: data,
	}
}

// Fail 失败响应不带数据
func Fail() model.APIResponse {
	return model.APIResponse{
		Code: 1,
		Msg:  "fail",
	}
}

// FailWithMessage 失败响应带消息
func FailWithMessage(message string) model.APIResponse {
	return model.APIResponse{
		Code: 1,
		Msg:  message,
	}
}
