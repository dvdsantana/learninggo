package dberrors

type ConflictError struct{}

func (e *ConflictError) Error() string {
	return "conflict error"
}
