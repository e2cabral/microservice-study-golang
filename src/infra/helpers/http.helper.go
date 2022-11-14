package helpers

import "net/http"

type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Status  int         `json:"status"`
}

func Ok(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	response := Response{
		Message: "",
		Data:    data,
		Status:  http.StatusOK,
	}
	ToJSON(w, response)
}

func Forbidden(w http.ResponseWriter, message string) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusForbidden)
	response := Response{
		Message: message,
		Data:    nil,
		Status:  http.StatusNotFound,
	}
	ToJSON(w, response)
}

func NotFound(w http.ResponseWriter, message string) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotFound)
	response := Response{
		Message: message,
		Data:    nil,
		Status:  http.StatusNotFound,
	}
	ToJSON(w, response)
}

func InternalServerError(w http.ResponseWriter, message string) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusInternalServerError)
	response := Response{
		Message: message,
		Data:    nil,
		Status:  http.StatusInternalServerError,
	}
	ToJSON(w, response)
}
