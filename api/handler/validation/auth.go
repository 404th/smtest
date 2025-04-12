package validation

import "github.com/404th/smtest/internal/repository/model"

func ValidateRegister(req *model.User) (bool, string) {
	if len(req.Username) < 4 {
		return false, "Username must be longer than 4 symbols"
	}

	if len(req.Password) < 1 {
		return false, "Password must be provided"
	}

	return true, "OK"
}
