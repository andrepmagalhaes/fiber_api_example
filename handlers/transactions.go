package handlers

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/andrepmagalhaes/q2bank_test/utils"
	"github.com/gofiber/fiber/v2"
)

type GetBalanceResponse struct {
	Message string `json:"message"`
	Balance float64 `json:"balance"`
}

func Balance(c *fiber.Ctx, db *sql.DB) error {
	response := GetBalanceResponse{}
	authorization := c.Get("Authorization")
	if authorization == "" {
		response.Message = "Unauthorized"
		return c.Status(401).JSON(response)
	}

	id, err := utils.VerifyJWT(authorization)

	if err != nil {
		log.Println(fmt.Sprintf("Error verifying JWT: %s", err.Error()))
		response.Message = "Internal Server Error"
		return c.Status(500).JSON(response)
	}

	if id == -1 {
		response.Message = "Unauthorized"
		return c.Status(401).JSON(response)
	}

	var balance uint64

	err = db.QueryRow("SELECT SUM(CASE WHEN transaction_type = 'credit' AND is_valid = true THEN amount WHEN transaction_type = 'debit' AND is_valid = true THEN -amount END) AS balance FROM public.\"Transactions\" WHERE payee = $1;", id).Scan(&balance)

	if err != nil {
		log.Println(fmt.Sprintf("Error getting balance: %s", err.Error()))
		response.Message = "Internal Server Error"
		return c.Status(500).JSON(response)
	}

	response.Balance = float64(balance) / 100
	return c.Status(200).JSON(response)
}