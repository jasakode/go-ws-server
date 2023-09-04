package middleware

import (
	"encoding/json"
	"net/http"
)

func Message(status bool, message string) map[string]interface{} {
	response := make(map[string]interface{})
	response["status"] = status
	response["message"] = message
	return response
}

func Response(w http.ResponseWriter, body map[string]interface{}) {
	resByte, _ := json.Marshal(body)
	w.Write(resByte)
}
