package utils

import "github.com/logo-user-management/app/errors"

func EqualPasswords(pass1, pass2 string) error {
	if pass1 != pass2 {
		return errors.PasswordsDoNotMatchError
	}
	return nil
}
