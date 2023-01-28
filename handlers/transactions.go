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
}

type TransactionResponse struct {
	Message string `json:"message"`
}


func balanceQuery(id int, db *sql.DB) (float64, error) {
	var balance float64
	err := db.QueryRow("SELECT SUM(CASE WHEN payee = $1 AND is_valid = true THEN amount WHEN payer = $1 AND is_valid = true THEN -amount ELSE 0 END) as balance FROM public.\"Transactions\"", id).Scan(&balance)
	
	if err != nil {
		log.Printf("Error getting balance: %s", err.Error())
		return 0, err
	}

	return balance, nil
}

func Balance(c *fiber.Ctx, db *sql.DB) error {
	response := BalanceResponse{}
	id, _, err := utils.VerifyJWT(c)

	if err != nil || id == -1 {
		response.Message = "Unauthorized"
		return c.Status(401).JSON(response)
	}

	balance, err := balanceQuery(id, db)

	if err != nil {
		log.Printf("Error getting balance: %s", err.Error())
		response.Message = "Internal Server Error"
		return c.Status(500).JSON(response)
	}

	response.Balance = balance
	return c.Status(200).JSON(response)
}

func findUserByIdQuery(id int, db *sql.DB) error {

	var userId int
	err := db.QueryRow("SELECT id FROM public.\"Users\" WHERE id = $1", id).Scan(&userId)
	if err != nil {
		return fmt.Errorf("user not found")
	}

	return nil

}

func insertTransactionQuery(payer int, payee int, value float64, db *sql.DB) error {
	_, err := db.Exec("INSERT INTO public.\"Transactions\" (payer, payee, amount) VALUES ($1, $2, $3)", payer, payee, value)
	if err != nil {
		log.Printf("Error inserting transaction: %s", err.Error())
		return fmt.Errorf("internal server error")
	}

	return nil
}

func Transaction(c *fiber.Ctx, db *sql.DB) error {
	response := TransactionResponse{}
	id, userType, err := utils.VerifyJWT(c)

	if err != nil || id == -1 || userType == "" {
		response.Message = "Unauthorized"
		return c.Status(401).JSON(response)
	}

	if userType == "store" {
		response.Message = "User of type store cant make transactions"
		return c.Status(401).JSON(response)
	}

	auth, err := utils.GetTransactionAuth()

	if err != nil {
		log.Printf("Error getting transaction auth: %s", err.Error())
		response.Message = "Internal Server Error"
		return c.Status(500).JSON(response)
	}

	if !auth {
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

	if body.Payee == id {
		response.Message = "Payee must be different from payer"
		return c.Status(400).JSON(response)
	}

	if body.Payee == 0 {
		response.Message = "Payee must be different from 0"
		return c.Status(400).JSON(response)
	}

	err = findUserByIdQuery(body.Payee, db)

	if err != nil {
		response.Message = "Payee not found"
		return c.Status(400).JSON(response)
	}

	balance, err := balanceQuery(id, db)

	if err != nil {
		log.Printf("Error getting balance: %s", err.Error())
		response.Message = "Internal Server Error"
		return c.Status(500).JSON(response)
	}

	if balance < body.Value {
		response.Message = "Insufficient funds"
		return c.Status(400).JSON(response)
	}

	err = insertTransactionQuery(id, body.Payee, body.Value, db)
	if err != nil {
		response.Message = err.Error()
		return c.Status(500).JSON(response)
	}

	response.Message = "Transaction successful"

	return c.Status(200).JSON(response)
}

