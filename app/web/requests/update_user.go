package requests

import (
	"encoding/json"
	"github.com/go-chi/chi"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/logo-user-management/app/data"
	"github.com/logo-user-management/app/web"
	"github.com/pkg/errors"
	"net/http"
)

type UpdateUserRequest struct {
	Username string
	Data data.User `json:"data"`
}

func (r UpdateUserRequest) Validate() error {
	return validation.ValidateStruct(&r.Data,
		validation.Field(&r.Data.Email, validation.When(r.Data.Email != "", is.Email)),
	)
}

func NewUpdateUserRequest(r *http.Request) (*UpdateUserRequest, error) {
	req := UpdateUserRequest{
		Username: chi.URLParam(r, web.UsernameRequestKey),
	}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, errors.Wrap(err, "failed to decode request body")
	}

	return &req, req.Validate()
}
