package handlers

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/andrepmagalhaes/q2bank_test/utils"
	"github.com/gofiber/fiber/v2"
)

type CreateBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	CpfCnpj	string `json:"cpf_cnpj"`
	Name string `json:"name"`
	UserType string `json:"user_type"`
}

type CreateResponse struct {
	Message string `json:"message"`
}

func Create(c *fiber.Ctx, db *sql.DB) error {
	body := CreateBody{}
	response := CreateResponse{}

	if err := c.BodyParser(&body); err != nil {
		response.Message = err.Error()
		return c.Status(400).JSON(response)
	}

	if body.Email == "" {
		response.Message = "email is required"
		return c.Status(400).JSON(response)
	}

	if body.Password == "" {
		response.Message = "password is required"
		return c.Status(400).JSON(response)
	}

	if body.CpfCnpj == "" {
		response.Message = "cpf_cnpj is required"
		return c.Status(400).JSON(response)
	}

	if body.Name == "" {
		response.Message = "name is required"
		return c.Status(400).JSON(response)
	}

	if body.UserType == "" {
		response.Message = "user_type is required"
		return c.Status(400).JSON(response)
	}

	if body.UserType != "person" && body.UserType != "store" {
		response.Message = "user_type must be either person or store"
		return c.Status(400).JSON(response)
	}

	pwValidationMessage, pwValidationValidity := utils.ValidatePassword(body.Password)
	
	if !pwValidationValidity {
		response.Message = pwValidationMessage
		return c.Status(400).JSON(response)
	}

	hashedPassword, err := utils.HashPassword(body.Password)

	if err != nil {
		log.Println(fmt.Sprintf("Error hashing password: %s", err.Error()))
		response.Message = "Internal Server Error"
		return c.Status(500).JSON(response)
	}

	rows, err := db.Query("SELECT ID FROM public.\"Users\" WHERE email = $1 OR cpf_cnpj = $2", body.Email, body.CpfCnpj)

	if err != nil {
		log.Println(fmt.Sprintf("Error checking if user exists: %s", err.Error()))
		response.Message = "Internal Server Error"
		return c.Status(500).JSON(response)
	}

	if rows.Next() {
		response.Message = "User with this email or cpf_cnpj already exists"
		return c.Status(400).JSON(response)
	}

	_, err = db.Exec("INSERT INTO public.\"Users\" (email, password, cpf_cnpj, name, user_type) VALUES ($1, $2, $3, $4, $5)", body.Email, hashedPassword, body.CpfCnpj, body.Name, body.UserType)

	if err != nil {
		log.Println(fmt.Sprintf("Error inserting user: %s", err.Error()))
		response.Message = "Internal Server Error"
		return c.Status(500).JSON(response)
	}

	response.Message = "created"
	return c.Status(200).JSON(response)
}

type LoginBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Message string `json:"message"`
	Token string `json:"token"`
}

func Login(c *fiber.Ctx, db *sql.DB) error {
	body := LoginBody{}
	response := LoginResponse{}

	if err := c.BodyParser(&body); err != nil {
		response.Message = err.Error()
		return c.Status(400).JSON(response)
	}

	if body.Email == "" {
		response.Message = "email is required"
		return c.Status(400).JSON(response)
	}

	if body.Password == "" {
		response.Message = "password is required"
		return c.Status(400).JSON(response)
	}

	rows, err := db.Query("SELECT id, password FROM public.\"Users\" WHERE email = $1", body.Email)

	if err != nil {
		log.Println(fmt.Sprintf("Error checking if user exists: %s", err.Error()))
		response.Message = "Internal Server Error"
		return c.Status(500).JSON(response)
	}

	if !rows.Next() {
		response.Message = "User with this email does not exist"
		return c.Status(400).JSON(response)
	}

	var id int
	var password string

	err = rows.Scan(&id, &password)

	if err != nil {
		log.Println(fmt.Sprintf("Error scanning user: %s", err.Error()))
		response.Message = "Internal Server Error"
		return c.Status(500).JSON(response)
	}

	if !utils.CheckPasswordHash(body.Password, password) {
		response.Message = "Wrong password"
		return c.Status(400).JSON(response)
	}

	response.Token, err = utils.CreateJWT(id)

	if err != nil {
		log.Println(fmt.Sprintf("Error creating JWT: %s", err.Error()))
		response.Message = "Internal Server Error"
		return c.Status(500).JSON(response)
	}

	response.Message = "logged in"
	return c.Status(200).JSON(response)
}