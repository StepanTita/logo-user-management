package errors

import "github.com/pkg/errors"

var (
	PasswordsDoNotMatchError = errors.New("passwords do not match")
)
