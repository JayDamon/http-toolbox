package httptoolbox

import (
	"encoding/json"
	"net/http"
)

func ReadJsonBodyToVariable(w http.ResponseWriter, r *http.Request, result any) {
	err := json.NewDecoder(r.Body).Decode(&result)
	if err != nil {
		RespondError(w, http.StatusInternalServerError, err.Error())
	}
}

func Respond(w http.ResponseWriter, status int, payload any) {

	if payload == nil {
		createJsonResponse(w, status)
		return
	}

	response, err := json.Marshal(payload)
	if err != nil {
		RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	createJsonResponse(w, status)
	w.Write([]byte(response))
}

func RespondNoBody(w http.ResponseWriter, status int) {
	createJsonResponse(w, status)
}

// RespondError forms message and code as http error response
func RespondError(w http.ResponseWriter, code int, message string) {
	Respond(w, code, map[string]string{"error": message})
}

func createJsonResponse(w http.ResponseWriter, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
}
