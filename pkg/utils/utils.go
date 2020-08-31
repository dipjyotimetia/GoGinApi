package utils

import (
	"github.com/GoGinApi/v2/entity"
	"regexp"
)

const (
	emailRegex = `^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`
)

// ValidateUser returns a slice of string of validation errors
func ValidateUser(user entity.Register, err []string) []string {
	emailCheck := regexp.MustCompile(emailRegex).MatchString(user.Email)
	if !emailCheck {
		err = append(err, "Invalid email")
	}
	if len(user.Password) < 4 {
		err = append(err, "Invalid password, Password should be more than 4 characters")
	}
	if len(user.Name) < 1 {
		err = append(err, "Invalid name, please enter a name")
	}

	return err
}

func ValidatePasswordReset(resetPassword entity.ResetPassword) (bool, string) {
	if len(resetPassword.Password) < 4 {
		return false, "Invalid password, password should be more than 4 characters"
	}
	if resetPassword.Password != resetPassword.ConfirmPassword {
		return false, "Password reset failed, passwords must match"
	}
	return true, ""
}
