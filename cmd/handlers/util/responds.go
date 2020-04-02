package util

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func Respond(w http.ResponseWriter, data map[string]interface{}) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func RespondMessage(w http.ResponseWriter, status bool, message string) {
	Respond(w, Message(status, message))
}

//Concrete responds with status codes
//

func RespondInvalidRequest(w http.ResponseWriter) {
	w.WriteHeader(http.StatusBadRequest)
	RespondMessage(w, false, "Invalid request")
}

func RespondInternalServerError(w http.ResponseWriter, message string) {
	w.WriteHeader(http.StatusInternalServerError)
	RespondMessage(w, false, message)
}

func RespondNotFound(w http.ResponseWriter, message string) {
	w.WriteHeader(http.StatusNotFound)
	RespondMessage(w, false, message)
}

func RespondNotAdmin(w http.ResponseWriter) {
	w.WriteHeader(http.StatusForbidden)
	RespondMessage(w, false, "This user is not an author")
}

func RespondUserNotFound(w http.ResponseWriter, username string) {
	RespondNotFound(w, fmt.Sprintf("User with name %q not found", username))
}
