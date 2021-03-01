package requests

import (
	"github.com/go-chi/chi"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/logo-user-management/app/web"
	"net/http"
)

type DeleteUserRequest struct {
	Username string
}

func (r DeleteUserRequest) Validate() error {
	return validation.Errors{
		"username": validation.Validate(r.Username, validation.Required),
	}.Filter()
}

func NewDeleteUserRequest(r *http.Request) (*GetUserRequest, error) {
	req := GetUserRequest{
		Username: chi.URLParam(r, web.UsernameRequestKey),
	}

	return &req, req.Validate()
}
