package failures

import (
	"encoding/json"
	"log"
	"net/http"
	"stewped-applet/common"
)

type (
	applicationError struct {
		Error      string `json:"error,omitempty"`
		Message    string `json:"message,omitempty"`
		StatusCode int    `json:"status,omitempty"`
	}
)

func WriteError(w http.ResponseWriter, handlerError error, message string, code int) {
	e := applicationError{
		Error:      handlerError.Error(),
		Message:    message,
		StatusCode: code,
	}
	log.Printf("application error occurred: %s\n", handlerError)
	w.Header().Set("Content-Type", common.ContentTypeJSON)
	w.WriteHeader(code)
	if j, err := json.Marshal(e); err == nil {
		w.Write(j)
	}
}

func WriteCustomError(w http.ResponseWriter, handlerError error, code int, body map[string]string) {
	log.Printf("application error occurred: %s\n", handlerError)
	w.Header().Set("Content-Type", common.ContentTypeJSON)
	w.WriteHeader(code)
	if j, err := json.Marshal(body); err == nil {
		w.Write(j)
	}
}
