package helper

import (
	"encoding/json"
	"net/http"
)

func DecodeJSONFromRequest(request *http.Request, result interface{}) {
	decoder := json.NewDecoder(request.Body)
	errDecode := decoder.Decode(result)
	PanicError(errDecode)
}

func EncodeJSONFromResponse(writer http.ResponseWriter, response interface{}) {
	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	errEncode := encoder.Encode(response)
	PanicError(errEncode)

}
