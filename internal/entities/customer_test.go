package entities

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewCustomer(t *testing.T) {
	test := assert.New(t)
	customer, err := NewCustomer(
		"Michael Jordan",
		"Player",
		"mj@email.com",
		"+1-202-555-0131",
	)

	test.Nil(err)
	test.NotNil(customer)
	test.NotEmpty(customer.ID)
	test.Equal("Michael Jordan", customer.Name)
	test.Equal("Player", customer.Role)
	test.Equal("+1-202-555-0131", customer.Phone)
	test.Equal("mj@email.com", customer.Email)
}

func TestNewCustomerWithEmptyEmail(t *testing.T) {
	test := assert.New(t)
	customer, err := NewCustomer(
		"Michael Jordan",
		"Player",
		"",
		"+1-202-555-0131",
	)

	test.NotNil(err)
	test.Nil(customer)
	test.Equal("email is required", err.Error())
}

func TestNewCustomerWithEmptyName(t *testing.T) {
	test := assert.New(t)
	customer, err := NewCustomer(
		"",
		"Player",
		"mj@email.com",
		"+1-202-555-0131",
	)

	test.NotNil(err)
	test.Nil(customer)
	test.Equal("name is required", err.Error())
}

func TestNewCustomerWithEmptyRole(t *testing.T) {
	test := assert.New(t)
	customer, err := NewCustomer(
		"Michael Jordan",
		"",
		"mj@email.com",
		"+1-202-555-0131",
	)

	test.NotNil(err)
	test.Nil(customer)
	test.Equal("role is required", err.Error())
}

func TestNewCustomerWithEmptyPhone(t *testing.T) {
	test := assert.New(t)
	customer, err := NewCustomer(
		"Michael Jordan",
		"Player",
		"mj@email.com",
		"",
	)

	test.NotNil(err)
	test.Nil(customer)
	test.Equal("phone is required", err.Error())
}
