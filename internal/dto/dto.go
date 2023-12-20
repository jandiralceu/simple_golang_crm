package dto

// CreateCustomerDto is a model that client use when creating a new Customer
type CreateCustomerDto struct {
	Name  string `json:"name"`
	Role  string `json:"role"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

// UpdateCustomerDto is a model that client use when updating a Customer
type UpdateCustomerDto struct {
	Name  string `json:"name"`
	Role  string `json:"role"`
	Phone string `json:"phone"`
}

// ErrorResponseDto is a struct to return an error
type ErrorResponseDto struct {
	Message string `json:"message"`
}
