package requests

import (
	"encoding/json"
	"github.com/go-chi/chi"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/logo-user-management/app/web"
	"github.com/pkg/errors"
	"net/http"
	"strconv"
)

type GetUserByIdRequest struct {
	Id   int64        `json:"-"`
	Data PasswordData `json:"data"`
}

func (r GetUserByIdRequest) Validate() error {
	return nil
}

func NewGetUserByIdRequest(r *http.Request) (*GetUserByIdRequest, error) {
	rawId := chi.URLParam(r, web.UserIDRequestKey)
	v, err := strconv.ParseInt(rawId, 10, 64)
	if err == nil {
		return nil, validation.Errors{
			"user_id": errors.New("user id must be a 64bit integer"),
		}
	}
	req := GetUserByIdRequest{
		Id: v,
	}

	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, errors.Wrap(err, "failed to decode request body")
	}

	return &req, req.Validate()
}
