package handlers

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/jandiralceu/crm/internal/dto"
	"github.com/jandiralceu/crm/internal/entities"
	"github.com/jandiralceu/crm/internal/infra/database"
	"net/http"
	"strconv"
)

// CustomerHandler is a handler for Customer
type CustomerHandler struct {
	CategoryDB database.ICustomer
}

// NewCustomerHandler returns a new CustomerHandler instance
func NewCustomerHandler(db database.ICustomer) *CustomerHandler {
	return &CustomerHandler{CategoryDB: db}
}

// Create is a handler to create a Customer
func (h *CustomerHandler) Create(w http.ResponseWriter, r *http.Request) {
	var c dto.CreateCustomerDto

	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	category, err := entities.NewCustomer(c.Name, c.Role, c.Email, c.Phone)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := h.CategoryDB.Create(category); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// FindByID is a handler to find a Customer by ID
func (h *CustomerHandler) FindByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	category, err := h.CategoryDB.FindByID(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(category); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

// FindAll is a handler to find all Customers
func (h *CustomerHandler) FindAll(w http.ResponseWriter, r *http.Request) {
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		page = 0
	}

	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil {
		limit = 0
	}

	var sort string
	if sort = r.URL.Query().Get("sort"); sort == "" {
		sort = "asc"
	}

	customers, err := h.CategoryDB.FindAll(page, limit, sort)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(customers); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

// Update is a handler to update a Customer
func (h *CustomerHandler) Update(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var c dto.UpdateCustomerDto

	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	customer, err := h.CategoryDB.FindByID(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if customer == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if c.Name != "" && c.Name != customer.Name {
		customer.Name = c.Name
	}

	if c.Role != "" && c.Role != customer.Role {
		customer.Role = c.Role
	}

	if c.Phone != "" && c.Phone != customer.Phone {
		customer.Phone = c.Phone
	}

	if err := h.CategoryDB.Update(customer); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// Delete is a handler to delete a Customer
func (h *CustomerHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	_, err := h.CategoryDB.FindByID(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if err := h.CategoryDB.Delete(id); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
