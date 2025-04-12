package handler

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/404th/smtest/internal/repository/model"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgconn"
	"gorm.io/gorm"
)

// custom err resp
func handleErrorResponse(c *gin.Context, err error) {
	resp := model.Response{
		Data:    "internal_server_error",
		Message: "Serverda ichki xatolik yuzaga keldi",
	}
	statusCode := http.StatusInternalServerError

	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		resp.Data = "not_found"
		resp.Message = "Ma'lumot topilmadi"
		statusCode = http.StatusNotFound

	case errors.Is(err, gorm.ErrInvalidTransaction):
		resp.Data = "transaction_error"
		resp.Message = "Xato ma'lumot yuborilgan"
		statusCode = http.StatusBadRequest

	case errors.Is(err, gorm.ErrNotImplemented):
		resp.Data = "not_implemented"
		resp.Message = "To'liq ko'tarilmagan"
		statusCode = http.StatusNotImplemented

	case isPgUniqueViolation(err):
		resp.Data = "conflict"
		resp.Message = "Kiritilgan ma'lumot tizimda mavjud"
		statusCode = http.StatusConflict

	case isPgForeignKeyViolation(err):
		resp.Data = "bad_request"
		resp.Message = "Foreign key noto'g'ri kiritilgan"
		statusCode = http.StatusBadRequest

	case errors.Is(err, gorm.ErrInvalidData):
		resp.Data = "bad_request"
		resp.Message = "Notog'ri ma'lumot kiritilgan"
		statusCode = http.StatusBadRequest

	case err.Error() == "invalid input": // INFO: maxsus error create qilish uchun template
		resp.Data = "bad_request"
		resp.Message = "Xato ma'lumot kiritilgan"
		statusCode = http.StatusBadRequest

	case err.Error() == "unauthorized":
		resp.Data = "unauthorized"
		resp.Message = "Tizimga qayta kiring"
		statusCode = http.StatusUnauthorized

	case err.Error() == "forbidden":
		resp.Data = "forbidden"
		resp.Message = "Ruxsat etilmagan"
		statusCode = http.StatusForbidden

	default:
		fmt.Printf("Unhandled error: %v\n", err)
	}

	c.JSON(statusCode, resp)
}

func isPgUniqueViolation(err error) bool {
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		return pgErr.Code == "23505"
	}
	return false
}

func isPgForeignKeyViolation(err error) bool {
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		return pgErr.Code == "23503"
	}
	return false
}
