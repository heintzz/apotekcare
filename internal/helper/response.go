package helper

import (
	"encoding/json"
	"net/http"
)

type APIResponse struct {
	HttpCode 	int					`json:"-"`
	Success  	bool        `json:"success"`
	Message	 	string      `json:"message"`
	Error    	interface{} `json:"error,omitempty"`
	Payload 	interface{} `json:"payload,omitempty"`
	ErrorCode string	  	`json:"error_code,omitempty"`
}

func (resp APIResponse) WriteJsonResponse(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Agent")
	
	w.WriteHeader(resp.HttpCode)
	json.NewEncoder(w).Encode(resp)
}

