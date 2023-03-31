package util

import (
	"encoding/json"
	CONSTANT "hdfc-backend/constant"
	"net/http"
)

// Meta - meta data
// status - HTTP status codes like 200,201,400,500,503
// message - Any message which would be used by app to display or take action
// message_type - 0 : no dialog, 1 : show dialog, 2 : show toast,
// dev_message - developer message
type Meta struct {
	Status      string `json:"status"`
	Message     string `json:"message"`
	MessageType string `json:"message_type"`
	DevMessage  string `json:"dev_message"`
}

// Response - response
type Response struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

// SetReponse - set request response with status, message etc
func SetReponse(w http.ResponseWriter, status, msg, msgType, devMessage string, resp interface{}) {
	if status == CONSTANT.StatusCodeOk {
		w.WriteHeader(http.StatusOK)
	} else if status == CONSTANT.StatusCodeBadRequest {
		w.WriteHeader(http.StatusBadRequest)
	} else if status == CONSTANT.StatusCodeServerError {
		w.WriteHeader(http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(Response{
		Meta: setMeta(status, msg, msgType, devMessage),
		Data: resp,
	})
}

// setMeta - set meta data
func setMeta(status, message, msgType, devMessage string) Meta {
	if len(message) == 0 {
		if status == CONSTANT.StatusCodeBadRequest {
			message = "Bad Request"
		} else if status == CONSTANT.StatusCodeServerError {
			message = "Server Error"
		}
	}
	return Meta{
		Status:      status,
		Message:     message,
		MessageType: msgType,
		DevMessage:  devMessage,
	}
}
