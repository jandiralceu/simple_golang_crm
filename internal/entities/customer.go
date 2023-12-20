package entities

import (
	"errors"
	"github.com/jandiralceu/crm/pkg/entities"
)

// Customer represents a customer
type Customer struct {
	ID        entities.ID `json:"id"`
	Name      string      `json:"name"`
	Role      string      `json:"role"`
	Email     string      `json:"email"`
	Phone     string      `json:"phone"`
	Contacted bool        `json:"contacted"`
}

// NewCustomer creates a new Customer
func NewCustomer(name, role, email, phone string) (*Customer, error) {
	category := &Customer{
		ID:        entities.GenerateUUID(),
		Name:      name,
		Role:      role,
		Email:     email,
		Phone:     phone,
		Contacted: false,
	}

	if err := category.isValid(); err != nil {
		return nil, err
	}

	return category, nil
}

// isValid checks whether the Customer fields are valid or not
func (p *Customer) isValid() error {
	if p.ID.String() == "" {
		return errors.New("id is required")
	}

	if p.Name == "" {
		return errors.New("name is required")
	}

	if p.Role == "" {
		return errors.New("role is required")
	}

	if p.Email == "" {
		return errors.New("email is required")
	}

	if p.Phone == "" {
		return errors.New("phone is required")
	}

	return nil
}
