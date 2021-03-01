package requests

import (
	"encoding/json"
	"github.com/go-chi/chi"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/logo-user-management/app/web"
	"github.com/pkg/errors"
	"net/http"
)

type PasswordData struct {
	Password string `json:"password"`
}

type GetUserRequest struct {
	Username string       `json:"-"`
	Data     PasswordData `json:"data"`
}

func (r GetUserRequest) Validate() error {
	return validation.Errors{
		"username": validation.Validate(r.Username, validation.Required),
	}.Filter()
}

func NewGetUserRequest(r *http.Request) (*GetUserRequest, error) {
	req := GetUserRequest{
		Username: chi.URLParam(r, web.UsernameRequestKey),
	}

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, errors.Wrap(err, "failed to decode request body")
	}

	return &req, req.Validate()
}
