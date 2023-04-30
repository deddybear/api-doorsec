package helper

import (
	"encoding/json"
	"net/http"
)

func ReadFromRequestBody(req *http.Request, body interface{}) {
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(body)
	PanicIfError(err)
}

func WriteToResponseBody(res http.ResponseWriter, response interface{}) {
	res.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(res)
	err := encoder.Encode(response)
	PanicIfError(err)
}
