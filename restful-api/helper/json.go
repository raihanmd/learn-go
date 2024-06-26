package helper

import (
	"encoding/json"
	"net/http"
)

func ReadFromReqBody(req *http.Request, result any) {
	err := json.NewDecoder(req.Body).Decode(result)
	PanicIfError(err)
}

func WriteToResBody(w http.ResponseWriter, response any) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(response)
	PanicIfError(err)
}
