package response

import (
	"encoding/json"
	"net/http"
)

type ResponsePayload struct {
	Success    bool   `json:"success"`
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	Data       any    `json:"data,omitempty"`
}

func (response *ResponsePayload) Write() []byte {
	encoded, err := json.Marshal(response)
	if err != nil {
		panic(err)
	}
	return encoded
}

func (response *ResponsePayload) SetSuccess(success bool) *ResponsePayload {
	response.Success = success
	return response
}

func (response *ResponsePayload) SetStatus(statusCode int) *ResponsePayload {
	response.StatusCode = statusCode
	return response
}

func (response *ResponsePayload) SetMessage(message string) *ResponsePayload {
	response.Message = message
	return response
}

func (response *ResponsePayload) SetData(data any) *ResponsePayload {
	response.Data = data
	return response
}

func (response *ResponsePayload) FromError(err error) *ResponsePayload {
	response = FromError(err)
	return response
}

func FromError(err error) *ResponsePayload {
	return &ResponsePayload{
		Success:    false,
		StatusCode: http.StatusInternalServerError,
		Message:    err.Error(),
		Data:       nil,
	}
}
