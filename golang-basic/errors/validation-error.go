package errors

type ValidationError struct {
	Message string
}

func (error *ValidationError) Error() string {
	return error.Message
}
