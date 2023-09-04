package ws

import (
	"encoding/json"
	"time"
)

type Actions struct {
	Action  string      `json:"action"`
	Message string      `json:"message"`
	Time    int64       `json:"time"`
	Data    interface{} `json:"data"`
}

type CreateActionsResponse struct {
	ToString func() (string, error)
	ToByte   func() ([]byte, error)
	Action   Actions
}

func CreateActions[DATA interface{}](action, message string, data ...DATA) CreateActionsResponse {
	response := CreateActionsResponse{}
	response.Action.Action = action
	response.Action.Message = message
	if len(data) > 0 {
		response.Action.Data = data[0]
	}
	response.Action.Time = time.Now().Unix()
	response.ToString = func() (string, error) {
		actionStr, errMarsal := json.Marshal(response.Action)
		if errMarsal != nil {
			return "", errMarsal
		}
		return string(actionStr), nil
	}
	response.ToByte = func() ([]byte, error) {
		actionByte, errMarsal := json.Marshal(response.Action)
		if errMarsal != nil {
			return []byte{}, errMarsal
		}
		return actionByte, nil
	}
	return response
}
