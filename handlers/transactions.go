package handlers

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/andrepmagalhaes/q2bank_test/utils"
	"github.com/gofiber/fiber/v2"
)

type BalanceResponse struct {
	Message string `json:"message"`
	Balance float64 `json:"balance"`
}

type TransactionBody struct {
	Value float64 `json:"value"`
	Payee int `json:"payee"`
	TransactionType string `json:"transaction_type"`
}

type TransactionResponse struct {
	Message string `json:"message"`
}


func balanceQuery(id int, db *sql.DB) (float64, error) {
	var balance uint64
	err := db.QueryRow("SELECT SUM(CASE WHEN transaction_type = 'credit' AND is_valid = true THEN amount WHEN transaction_type = 'debit' AND is_valid = true THEN -amount ELSE 0 END) FROM public.\"Transactions\" WHERE payee = $1", id).Scan(&balance)
	if err != nil {
		return 0, err
	}

	return float64(balance)/100, nil
}

func Balance(c *fiber.Ctx, db *sql.DB) error {
	response := BalanceResponse{}
	id, _, err := utils.VerifyJWT(c)

	if err != nil {
		log.Println(fmt.Sprintf("Error verifying JWT: %s", err.Error()))
		response.Message = "Internal Server Error"
		return c.Status(500).JSON(response)
	}

	if id == -1 {
		response.Message = "Unauthorized"
		return c.Status(401).JSON(response)
	}

	balance, err := balanceQuery(id, db)

	if err != nil {
		log.Println(fmt.Sprintf("Error getting balance: %s", err.Error()))
		response.Message = "Internal Server Error"
		return c.Status(500).JSON(response)
	}

	response.Balance = float64(balance) / 100
	return c.Status(200).JSON(response)
}

func findUserByIdQuery(id int, db *sql.DB) error {

	var userId int
	err := db.QueryRow("SELECT id FROM public.\"Users\" WHERE id = $1", id).Scan(&userId)
	if err != nil {
		return fmt.Errorf("User not found")
	}

	return nil

}

func Transaction(c *fiber.Ctx, db *sql.DB) error {
	response := TransactionResponse{}
	id, userType, err := utils.VerifyJWT(c)

	if err != nil {
		log.Println(fmt.Sprintf("Error verifying JWT: %s", err.Error()))
		response.Message = "Internal Server Error"
		return c.Status(500).JSON(response)
	}

	if id == -1 || userType == "" {
		response.Message = "Unauthorized"
		return c.Status(401).JSON(response)
	}

	if userType == "store" {
		response.Message = "User of type store cant make transactions"
		return c.Status(401).JSON(response)
	}

	auth, err := utils.GetTransactionAuth()

	if err != nil {
		log.Println(fmt.Sprintf("Error getting transaction auth: %s", err.Error()))
		response.Message = "Internal Server Error"
		return c.Status(500).JSON(response)
	}

	if auth == false {
		response.Message = "Transaction refused"
		return c.Status(401).JSON(response)
	}

	body := TransactionBody{}

	if err := c.BodyParser(&body); err != nil {
		response.Message = err.Error()
		return c.Status(400).JSON(response)
	}

	if body.Value <= 0 {
		response.Message = "Value must be greater than 0"
		return c.Status(400).JSON(response)
	}

	if body.TransactionType != "credit" && body.TransactionType != "debit" {
		response.Message = "Transaction type must be credit or debit"
		return c.Status(400).JSON(response)
	}

	if body.TransactionType == "debit" {
		balance, err := balanceQuery(id, db)

	// err = findUserByIdQuery(, db)



	return c.Status(200).JSON(response)
}

