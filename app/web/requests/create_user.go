package requests

import (
	"encoding/json"
	"github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/logo-user-management/app/data"
	"github.com/pkg/errors"
	"net/http"
)

type CreateUserRequest struct {
	Data data.User `json:"data"`
}

// Basic data validation
func (r CreateUserRequest) Validate() error {
	return validation.ValidateStruct(&r.Data,
		validation.Field(&r.Data.Username, validation.Required, validation.Length(6, 50)),
		validation.Field(&r.Data.Email, validation.Required, is.Email),
		validation.Field(&r.Data.Password, validation.Required, validation.Length(6, 50)),
	)
}

func NewCreateUserRequest(r *http.Request) (*CreateUserRequest, error) {
	req := CreateUserRequest{}

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, errors.Wrap(err, "failed to decode request body")
	}

	return &req, req.Validate()
}
