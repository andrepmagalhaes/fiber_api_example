package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/andrepmagalhaes/q2bank_test/handlers"
	"github.com/andrepmagalhaes/q2bank_test/utils"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)

func main() {

	err := utils.SetupDotEnv()
	if err != nil {
		log.Fatal(fmt.Printf("Error loading .env file: %s", err.Error()))
	}

	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
			log.Fatal(err)
	}
	defer file.Close()
	log.SetOutput(file)

	dbConnStr := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable", os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_PORT"), os.Getenv("POSTGRES_DB"))

	fmt.Println(dbConnStr)

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

	app.Listen(":" + os.Getenv("API_PORT"))

}
