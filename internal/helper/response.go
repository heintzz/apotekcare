package helper

import (
	"encoding/json"
	"net/http"
)

type APIResponse struct {
	Success  	bool        `json:"success"`
	Status 		int					`json:"status"`
	Message	 	string      `json:"message"`
	Error    	interface{} `json:"error,omitempty"`
	ErrorCode int 				`json:"errorCode,omitempty"`
}

func (resp APIResponse) WriteJsonResponse(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Agent")
	
	w.WriteHeader(resp.Status)
	json.NewEncoder(w).Encode(resp)
}