package dto

import "encoding/json"

type ResponseData struct {
	meta Meta
	data interface{}
}

// Constructor cho Response thành công
func SuccessResponse(data interface{}) *ResponseData {
	return &ResponseData{
		meta: Meta{status: 200, message: "Success"},
		data: data,
	}
}

// Constructor cho Response lỗi
func ErrorResponse(code int, msg string) *ResponseData {
	return &ResponseData{
		meta: Meta{status: code, message: msg},
		data: nil,
	}
}

// MarshalJSON: Export ra ngoài khi trả về client
func (r ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Meta Meta        `json:"meta"`
		Data interface{} `json:"data,omitempty"`
	}{
		Meta: r.meta,
		Data: r.data,
	})
}
