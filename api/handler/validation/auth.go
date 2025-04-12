package validation

import "github.com/404th/smtest/internal/repository/model"

func ValidateRegister(req *model.User) (bool, string) {
	if len(req.Username) < 4 {
		return false, "Username kamida 4 ta belgidan iborat bo'lishi kerak"
	}

	if len(req.Password) < 1 {
		return false, "Password bo'sh bo'lishi mumkin emas"
	}

	return true, "OK"
}
