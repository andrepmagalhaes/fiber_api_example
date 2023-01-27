package handlers

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
)

func Create(c *fiber.Ctx, db *sql.DB) error {
	return c.SendString("createAccount")
}

func Login(c *fiber.Ctx, db *sql.DB) error {
	return c.SendString("login")
}