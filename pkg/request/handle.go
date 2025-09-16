package request

import (
	"net/http"
	"projects/GoLinkStat/pkg/response"
)

func HandleBody[T any](resWriter *http.ResponseWriter, req *http.Request) (*T, error) {
	body, err := decode[T](req.Body)
	if err != nil {
		response.Json(err.Error(), *resWriter, 402)
		return nil, err
	}
	err = validatorBody(body)
	if err != nil {
		response.Json(err.Error(), *resWriter, 402)
		return nil, err
	}
	return &body, nil
}
