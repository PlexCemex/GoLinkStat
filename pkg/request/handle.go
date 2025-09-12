package request

import (
	"net/http"
	"projects/GoLinkStat/pkg/response"
)

func HandleBody[T any](resWriter *http.ResponseWriter, req *http.Request) (*T, error) {
	body, err := Decode[T](req.Body)
	if err != nil {
		response.ResponseJson(err.Error(), *resWriter, 402)
		return nil, err
	}
	err = ValidatorBody(body)
	if err != nil {
		response.ResponseJson(err.Error(), *resWriter, 402)
		return nil, err
	}
	return &body, nil
}
