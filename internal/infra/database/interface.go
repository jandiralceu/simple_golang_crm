package database

import "github.com/jandiralceu/crm/internal/entities"

// ICustomer interface is the contract to implement a Customer repository
type ICustomer interface {
	Create(c *entities.Customer) error
	FindAll(page, limit int, sort string) ([]entities.Customer, error)
	FindByID(id string) (*entities.Customer, error)
	Update(c *entities.Customer) error
	Delete(id string) error
}
