package utils

import (
	"leobelini/cashly/internal/types/api"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleError(c *gin.Context, err error) {
	message, isInternal := GetErrorMessage(err)

	data := api.ErrorResponse{
		Message: message,
	}

	errorCode := http.StatusBadRequest
	if isInternal {
		errorCode = http.StatusInternalServerError
	}
	c.JSON(errorCode, data)
}
