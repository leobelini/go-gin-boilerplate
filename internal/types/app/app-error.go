package app

type AppError struct {
	Message    string `json:"message"`
	IsInternal bool   `json:"isInternal"`
}

func (e *AppError) Error() string {
	return e.Message
}
