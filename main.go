package main

import (
	"database/sql"
	"ecomm-go/app/handlers"
	"ecomm-go/app/handlers/auth"
	"ecomm-go/db/goqueries"
	"ecomm-go/types"
	"encoding/gob"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	"github.com/stripe/stripe-go/v72"
	"log"
	"os"
	"time"
)

func main() {
	gob.Register(types.CookieUser{})

	connStr := "postgres://postgres:secret@localhost:5432/ecomm?sslmode=disable"
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(time.Minute)

	err = godotenv.Load(".env")
	if err != nil {
		log.Fatal("HandleError loading .env file")
	}

	stripe.Key = os.Getenv("STRIPE_KEY")

	queries := goqueries.New(db)
	app := echo.New()
	app.Validator = &handlers.CustomValidator{Validator: validator.New()}

	h := handlers.DataHandler{
		Queries:         queries,
		Store:           auth.NewAuth(),
		DefaultURL:      os.Getenv("DEFAULT_URL"),
		StripePublicKey: os.Getenv("STRIPE_PUBLIC_KEY"),
	}

	handlers.RegisterRoutes(app, h)
	go handlers.BroadcastMessages(queries)

	app.Logger.Fatal(app.Start(":4000"))

}
