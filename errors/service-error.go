package errors

// ServiceError should be used to retur business error messages
type ServiceError struct {
	Message string `json:"message"`
}
