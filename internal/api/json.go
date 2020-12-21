package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/var512/ddda-save-corrupter/internal/logger"
)

// ErrorJSON replies to the request with the specified error message and HTTP code using JSON.
func ErrorJSON(w http.ResponseWriter, error string, code int) {
	logger.Default.Println("== Error:", error)

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)

	j, err := json.Marshal(error)
	if err != nil {
		log.Fatal(err)
	}
	_, err = w.Write(j)
	if err != nil {
		log.Fatal(err)
	}
}

// ResponseJSON replies to the request with the specified message using JSON.
func ResponseJSON(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	j, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}
	_, err = w.Write(j)
	if err != nil {
		log.Fatal(err)
	}
}
