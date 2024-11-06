package entity

type ValidationErrorEntity struct {
	Message     string                         `json:"message"`
	FieldErrors map[string][]*FieldErrorEntity `json:"fieldErrors"`
}

func NewValidationErrorEntity(message string, fieldErrors map[string][]*FieldErrorEntity) *ValidationErrorEntity {
	return &ValidationErrorEntity{Message: message, FieldErrors: fieldErrors}
}
