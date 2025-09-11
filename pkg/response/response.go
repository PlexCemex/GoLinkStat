package response

import (
	"encoding/json"
	"net/http"
)

func ResponseJson(res any, resWritter http.ResponseWriter, statusCode int16) {
	resWritter.Header().Set("Content-Type", "application/json")
	resWritter.WriteHeader(int(statusCode))
	json.NewEncoder(resWritter).Encode(res)
}
