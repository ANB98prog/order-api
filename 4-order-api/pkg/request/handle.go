package request

import (
	"github.com/ANB98prog/purple-school-homeworks/4-order-api/pkg/json"
	"github.com/ANB98prog/purple-school-homeworks/4-order-api/pkg/response"
	"net/http"
)

func HandleBody[T any](w *http.ResponseWriter, r *http.Request) (*T, error) {
	body, err := json.Decode[T](r.Body)
	if err != nil {
		response.BadRequest(*w, response.NewErrorMessage(err))
		return nil, err
	}

	err = json.IsValid(body)
	if err != nil {
		response.BadRequest(*w, response.NewErrorMessage(err))
		return nil, err
	}

	return &body, nil
}
