package database

import (
	"github.com/jandiralceu/crm/internal/entities"
	"gorm.io/gorm"
)

type Customer struct {
	DB *gorm.DB
}

// CustomerDB  is used to access the customer database
func CustomerDB(db *gorm.DB) *Customer {
	return &Customer{DB: db}
}

// Create inserts a new customer into the database
func (c *Customer) Create(customer *entities.Customer) error {
	return c.DB.Create(customer).Error
}

// FindAll retrieves all customers from the database
func (c *Customer) FindAll(page, limit int, sort string) ([]entities.Customer, error) {
	var customers []entities.Customer
	var err error

	if sort == "" || !(sort == "asc" || sort == "desc") {
		sort = "asc"
	}

	if page != 0 && limit != 0 {
		err = c.DB.Limit(limit).Offset((page - 1) * limit).Order("created_at " + sort).Find(&customers).Error
	} else {
		err = c.DB.Order("created_at " + sort).Find(&customers).Error
	}

	return customers, err
}

// FindByID retrieves a customer by its ID from the database
func (c *Customer) FindByID(id string) (*entities.Customer, error) {
	var customer entities.Customer
	err := c.DB.First(&customer, id).Error
	return &customer, err
}

// Update updates a customer in the database
func (c *Customer) Update(customer *entities.Customer) error {
	if _, err := c.FindByID(customer.ID.String()); err != nil {
		return err
	}
	return c.DB.Save(customer).Error
}

// Delete deletes a customer from the database
func (c *Customer) Delete(id string) error {
	customer, err := c.FindByID(id)
	if err != nil {
		return err
	}

	return c.DB.Delete(customer).Error
}
