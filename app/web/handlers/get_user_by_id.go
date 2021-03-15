package handlers

import (
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/logo-user-management/app/ctx"
	"github.com/logo-user-management/app/render"
	"github.com/logo-user-management/app/web/requests"
	"net/http"
)

func GetUserByID(w http.ResponseWriter, r *http.Request) {
	log := ctx.Log(r)

	request, err := requests.NewGetUserByIdRequest(r)
	if err != nil {
		if verr, ok := err.(validation.Errors); ok {
			log.WithError(verr).Debug("failed to parse get user request")
			render.Respond(w, http.StatusBadRequest, render.Message(fmt.Sprintf("request was invalid in some way: %s", verr.Error())))
			return
		}
		log.WithError(err).Error("something bad happened")
		render.Respond(w, http.StatusInternalServerError, render.Message("something bad happened parsing the request"))
		return
	}

	user, err := ctx.Users(r).GetUserById(request.Id)
	if err != nil {
		log.WithError(err).Error("failed to get user")
		render.Respond(w, http.StatusInternalServerError, render.Message("something bad happened"))
		return
	}
	if user == nil {
		log.WithError(err).Debug("specified user not found")
		render.Respond(w, http.StatusNotFound, render.Message("specified user not found"))
		return
	}

	render.Respond(w, http.StatusOK, render.Message(user.ToReturn()))
}
