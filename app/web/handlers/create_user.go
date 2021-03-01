package handlers

import (
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/logo-user-management/app/ctx"
	"github.com/logo-user-management/app/render"
	"github.com/logo-user-management/app/web/requests"
	"net/http"
)

// todo add salt
// todo add encryption
func CreateUser(w http.ResponseWriter, r *http.Request) {
	log := ctx.Log(r)

	request, err := requests.NewCreateUserRequest(r)
	if err != nil {
		if verr, ok := err.(validation.Errors); ok {
			log.WithError(verr).Debug("failed to parse create user request")
			render.Respond(w, http.StatusBadRequest, render.Message(fmt.Sprintf("request was invalid in some way: %s", verr.Error())))
			return
		}
		log.WithError(err).Error("failed to parse create user request")
		render.Respond(w, http.StatusInternalServerError, render.Message("something bad happened parsing the request"))
		return
	}

	err = ctx.Users(r).CreateUser(request.Data)
	if err != nil {
		log.WithError(err).Error("failed to create user")
		render.Respond(w, http.StatusInternalServerError, render.Message("something bad happened creating the user"))
		return
	}

	user, err := ctx.Users(r).GetUser(request.Data.Username)
	if err != nil {
		log.WithError(err).Error("failed to find user")
		render.Respond(w, http.StatusInternalServerError, render.Message("something bad happened trying to find the user"))
		return
	}

	render.Respond(w, http.StatusOK, render.Message(user.ToMap()))
}
