package helpers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const (
	SUCCESS_MESSSAGE string = "Success"
)

func HandleResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

func GenerateRefCode() string {
	currentUnixTime := time.Now().Unix()

	refCode := fmt.Sprintf("REF%d", currentUnixTime)

	return refCode
}
