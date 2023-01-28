package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/andrepmagalhaes/q2bank_test/handlers"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)

func main() {

	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
			log.Fatal(err)
	}
	defer file.Close()
	log.SetOutput(file)

	dbConnStr := "postgresql://postgres:123456@localhost:5432/q2bank?sslmode=disable"
	db, err := sql.Open("postgres", dbConnStr)

	if err != nil {
		log.Fatal(fmt.Printf("Error opening database: %q", err))
	}
	defer db.Close()

	app:= fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("ok")
	})

	app.Post("/create", func(c *fiber.Ctx) error {
		return handlers.Create(c, db)
	})

	app.Post("/login", func(c *fiber.Ctx) error {
		return handlers.Login(c, db)
	})

	app.Get("/balance", func(c *fiber.Ctx) error {
		return handlers.Balance(c, db)
	})

	app.Post("/transaction", func(c *fiber.Ctx) error {
		return handlers.Transaction(c, db)
	})

	app.Listen(":3000")

}
