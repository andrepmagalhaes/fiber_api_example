package handlers

import (
	"database/sql"
	"testing"

	_ "github.com/lib/pq"
)

func TestBalanceQuery(t *testing.T) {
	
	dbConnStr := "postgresql://postgres:123456@localhost:5432/q2bank?sslmode=disable"
	db, err := sql.Open("postgres", dbConnStr)

	if err != nil {
		t.Errorf("Error creating connection with db: %s", err.Error())
	}
	defer db.Close()

	balance, err := balanceQuery(2, db)
	if err != nil {
		t.Errorf("Error getting users balance: %s | The given user should have a positive balance, got %f", err.Error(), balance)
	}

	balance, err = balanceQuery(120381203, db)
	if balance != 0 || err != nil {
		t.Errorf("Error getting users balance | the given balance should be 0, got %f", balance)
	}
}

func TestFindUserByIdQuery(t *testing.T){

	dbConnStr := "postgresql://postgres:123456@localhost:5432/q2bank?sslmode=disable"
	db, err := sql.Open("postgres", dbConnStr)
	if err != nil {
		t.Errorf("Error creating connection with db: %s", err.Error())
	}
	defer db.Close()
	
	err = findUserByIdQuery(2, db)
	if err != nil {
		t.Errorf("Error finding user by id: %s | The given user should exist", err.Error())
	}

	err = findUserByIdQuery(120381203, db)
	if err == nil {
		t.Errorf("Error finding user by id: %s | The given user shouldnt exist", err.Error())
	}
	
}

func TestInsertTransactionQuery(t *testing.T){

	dbConnStr := "postgresql://postgres:123456@localhost:5432/q2bank?sslmode=disable"
	db, err := sql.Open("postgres", dbConnStr)
	if err != nil {
		t.Errorf("Error creating connection with db: %s", err.Error())
	}
	defer db.Close()

	err = insertTransactionQuery(2, 120381203, 100, db)
	if err == nil {
		t.Errorf("Error inserting transaction: The given transaction should be invalid since payee doesnt exists")
	}

	err = insertTransactionQuery(120381203, 3, 100, db)
	if err == nil {
		t.Errorf("Error inserting transaction: The given transaction should be invalid since payer doesnt exists")
	}

	err = insertTransactionQuery(2, 3, 100, db)
	if err != nil {
		t.Errorf("Error inserting transaction: %s | The given transaction should be valid", err.Error())
	}

	_, err = db.Exec("DELETE FROM public.\"Transactions\" WHERE payer = 2 AND payee = 3 AND amount = 100")
	if err != nil {
		t.Errorf("Error deleting test transaction: %s", err.Error())
	}



}