package response

import (
	"encoding/json"
	"net/http"
)

func ResponseJson(res any, resWriter http.ResponseWriter, statusCode int16) {
	resWriter.Header().Set("Content-Type", "application/json")
	resWriter.WriteHeader(int(statusCode))
	json.NewEncoder(resWriter).Encode(res)
}
