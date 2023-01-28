package handlers

import (
	"database/sql"
	"testing"

	_ "github.com/lib/pq"
)

func TestValidateUserQuery(t *testing.T) {

	dbConnStr := "postgresql://postgres:123456@localhost:5432/q2bank?sslmode=disable"
	db, err := sql.Open("postgres", dbConnStr)

	if err != nil {
		t.Errorf("Error creating connection with db: %s", err.Error())
	}
	defer db.Close()

	err = validateUserQuery("person1@email.com", "00011111111111", db)
	if err == nil {
		t.Errorf("Error validating user: %s | User already exists on db when test says it doesnt", err.Error())
	}

	err = validateUserQuery("person1@email.com", "000000000000", db)
	if err == nil {
		t.Errorf("Error validating user: %s | User already exists on db when test says it doesnt", err.Error())
	}

	err = validateUserQuery("person3@email.com", "00011111111111", db)
	if err == nil {
		t.Errorf("Error validating user: %s | cpf_cnpj already exists on db when test says it doesnt", err.Error())
	}

	err = validateUserQuery("person3@email.com", "000000000000", db)
	if err != nil {
		t.Errorf("Error validating user: %s | User shouldnt exist on db when test says it does", err.Error())
	}

}

func TestFindUserQuery(t *testing.T) {
	
	dbConnStr := "postgresql://postgres:123456@localhost:5432/q2bank?sslmode=disable"
	db, err := sql.Open("postgres", dbConnStr)

	if err != nil {
		t.Errorf("Error creating connection with db: %s", err.Error())
	}
	defer db.Close()

	id, password, userType, err := findUserQuery("person1@email.com", db)
	if err != nil {
		t.Errorf("Error finding user: %s | User should exist on db when test says it doesnt", err.Error())
	}
	if id == -1 {
		t.Errorf("Error finding user: %s | User's id should be different than -1 since it exists", err.Error())
	}
	if password == "" {
		t.Errorf("Error finding user: %s | User's password should be different than empty string since it exists", err.Error())
	}
	if userType == "" {
		t.Errorf("Error finding user: %s | User's type should be different than empty string since it exists", err.Error())
	}

	id, password, userType, err = findUserQuery("person123123@email.com", db)
	if err == nil {
		t.Errorf("Error finding user: %s | User shouldnt exist on db when test says it does", err.Error())
	}
	if id != -1 {
		t.Errorf("Error finding user: %s | User's id should be -1 since it doesnt exist", err.Error())
	}
	if password != "" {
		t.Errorf("Error finding user: %s | User's password should be empty string since it doesnt exist", err.Error())
	}
	if userType != "" {
		t.Errorf("Error finding user: %s | User's type should be empty string since it doesnt exist", err.Error())
	}
}

func TestCreateUserQuery(t *testing.T) {
	dbConnStr := "postgresql://postgres:123456@localhost:5432/q2bank?sslmode=disable"
	db, err := sql.Open("postgres", dbConnStr)

	if err != nil {
		t.Errorf("Error creating connection with db: %s", err.Error())
	}
	defer db.Close()
	
	err = insertUserQuery(CreateBody{Email: "email@email.com", Password: "password", CpfCnpj: "00011111111111", Name: "test", UserType: "person"}, db)

	if err == nil {
		t.Errorf("Error creating user: %s | User shouldnt be created since the given cpf_cnpj already exists on db and it should be unique", err.Error())
	}

	err = insertUserQuery(CreateBody{Email: "person1@email.com", Password: "password", CpfCnpj: "123123123123", Name: "test", UserType: "person"}, db)

	if err == nil {
		t.Errorf("Error creating user: %s | User shouldnt be created since the given email already exists on db and it should be unique", err.Error())
	}

	err = insertUserQuery(CreateBody{Email: "person123@email.com", Password: "password", CpfCnpj: "123123123123", Name: "test", UserType: "asdf"}, db)

	if err == nil {
		t.Errorf("Error creating user: %s | User shouldnt be created since the given user_type doesnt exists on user_type enum", err.Error())
	}

	err = insertUserQuery(CreateBody{Email: "person123@email.com", Password: "password", CpfCnpj: "123123123123", Name: "test", UserType: "bank"}, db)

	if err == nil {
		t.Errorf("Error creating user: %s | User shouldnt be created since the user_type bank can only be created manually through sql queries in db", err.Error())
	}

	err = insertUserQuery(CreateBody{Email: "person123@email.com", Password: "password", CpfCnpj: "123123123123", Name: "test", UserType: "person"}, db)

	if err != nil {
		t.Errorf("Error creating user: %s | User should be created since all its parameters are correct", err.Error())
	}

	_, err = db.Exec("DELETE FROM public.\"Users\" WHERE email = 'person123@email.com'")

	if err != nil {
		t.Errorf("Error deleting user: %s | User should be deleted since it was created for this test", err.Error())
	}

}