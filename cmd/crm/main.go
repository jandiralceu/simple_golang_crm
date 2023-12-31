package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jandiralceu/crm/configs"
	"github.com/jandiralceu/crm/internal/entities"
	"github.com/jandiralceu/crm/internal/infra/database"
	"github.com/jandiralceu/crm/internal/webserver/handlers"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"net/http"
)

func main() {
	r := chi.NewRouter()
	config, err := configs.LoadConfig("/")
	if err != nil {
		panic(err)
	}

	var db *gorm.DB

	if config.DBDns == "" {
		if config.Environment != "develop" {
			panic("failed to connect database")
		}

		db, err = gorm.Open(sqlite.Open("./app.db"), &gorm.Config{})
		if err != nil {
			panic("failed to connect database")
		}
	} else {
		db, err = gorm.Open(postgres.Open(config.DBDns), &gorm.Config{})
		if err != nil {
			panic("failed to connect database")
		}
	}

	if err = db.AutoMigrate(&entities.Customer{}); err != nil {
		panic("failed to migrate database")
	}

	customerHandler := handlers.NewCustomerHandler(database.CustomerDB(db))
	if config.Environment == "develop" {
		r.Use(middleware.Logger)
	}

	r.Use(middleware.Recoverer)

	r.Route("/api/v1/customers", func(r chi.Router) {
		r.Get("/", customerHandler.FindAll)
		r.Get("/{id}", customerHandler.FindByID)
		r.Post("/", customerHandler.Create)
		r.Put("/{id}", customerHandler.Update)
		r.Delete("/{id}", customerHandler.Delete)
	})

	r.Get("/", handlers.MainPageHandler)

	log.Fatal(http.ListenAndServe(":8080", r))
}
