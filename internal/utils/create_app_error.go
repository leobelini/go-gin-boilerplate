package utils

import "leobelini/cashly/internal/types/app"

func CreateAppError(message string, isInternal bool) error {
	return &app.AppError{Message: message, IsInternal: isInternal}
}
