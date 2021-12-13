package validator

// FieldErr is struct for validation error
type FieldErr struct {
	Field string `json:"field"`
	Error string `json:"error"`
}
