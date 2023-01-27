package handlers

import (
	"database/sql"
	"fmt"
	"testing"

	_ "github.com/lib/pq"
)

func TestBalanceQuery(t *testing.T) {
	
	dbConnStr := "postgresql://postgres:123456@localhost:5432/q2bank?sslmode=disable"
	db, err := sql.Open("postgres", dbConnStr)

	if err != nil {
		t.Error(fmt.Sprintf("Error creating connection with db: %s", err.Error()))
	}
	defer db.Close()

	balance, err := balanceQuery(2, db)
	if err != nil {
		t.Error(fmt.Sprintf("Error getting users balance: %s | The given user should have a positive balance, got %f", err.Error(), balance))
	}

	balance, err = balanceQuery(120381203, db)
	if err == nil {
		t.Error(fmt.Sprintf("Error getting users balance: %s | the given balance shouldnt exists", err.Error()))
	}
}

func TestFindUserByIdQuery(t *testing.T){

	dbConnStr := "postgresql://postgres:123456@localhost:5432/q2bank?sslmode=disable"
	db, err := sql.Open("postgres", dbConnStr)
	if err != nil {
		t.Error(fmt.Sprintf("Error creating connection with db: %s", err.Error()))
	}
	defer db.Close()
	
	err = findUserByIdQuery(2, db)
	if err != nil {
		t.Error(fmt.Sprintf("Error finding user by id: %s | The given user should exist", err.Error()))
	}

	err = findUserByIdQuery(120381203, db)
	if err == nil {
		t.Error(fmt.Sprintf("Error finding user by id: %s | The given user shouldnt exist", err.Error()))
	}
	
}