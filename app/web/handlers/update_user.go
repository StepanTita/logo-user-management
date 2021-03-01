package handlers

import (
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/logo-user-management/app/ctx"
	"github.com/logo-user-management/app/data"
	"github.com/logo-user-management/app/render"
	"github.com/logo-user-management/app/web/requests"
	"net/http"
)

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	log := ctx.Log(r)

	request, err := requests.NewUpdateUserRequest(r)
	if err != nil {
		if verr, ok := err.(validation.Errors); ok {
			log.WithError(verr).Debug("failed to parse update user request")
			render.Respond(w, http.StatusBadRequest, render.Message(fmt.Sprintf("request was invalid in some way: %s", verr.Error())))
			return
		}
		log.WithError(err).Error("failed to parse update user request")
		render.Respond(w, http.StatusInternalServerError, render.Message("something bad happened parsing the request"))
		return
	}

	user, err := ctx.Users(r).GetUser(request.Username)
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

	user, err = updateUser(user, &request.Data)
	if err != nil {
		log.WithError(err).Debug("user has wrong update data")
		render.Respond(w, http.StatusBadRequest, render.Message(fmt.Sprintf("request was invalid in some way: %s", err.Error())))
		return
	}

	err = ctx.Users(r).UpdateUser(*user)
	if err != nil {
		log.WithError(err).Error("failed to create user")
		render.Respond(w, http.StatusInternalServerError, render.Message("something bad happened creating the user"))
		return
	}

	render.Respond(w, http.StatusOK, render.Message(user.ToMap()))
}

func updateUser(oldUser *data.User, newUser *data.User) (*data.User, error) {
	if newUser.Username != "" {
		oldUser.Username = newUser.Username
	}

	if newUser.Name != "" {
		oldUser.Name = newUser.Name
	}

	if newUser.Surname != "" {
		oldUser.Surname = newUser.Surname
	}

	if newUser.Email != "" {
		if err := validation.Validate(newUser.Email, is.Email); err != nil {
			return nil, err
		}
		oldUser.Email = newUser.Email
	}

	if newUser.Password != "" {
		oldUser.Password = newUser.Password
	}

	return oldUser, nil
}
