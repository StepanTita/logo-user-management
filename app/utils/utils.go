package utils

import (
	logoerrors "github.com/logo-user-management/app/errors"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

func EqualPasswords(hashedPwd, plainPwd string) error {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, []byte(plainPwd))
	if err != nil {
		return logoerrors.PasswordsDoNotMatchError
	}

	return nil
}

func HashAndSalt(pwd string) (string, error) {

	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.MinCost)
	if err != nil {
		return "", errors.Wrap(err, "failed to generate hash and salt from pwd")
	}
	return string(hash), nil
}
