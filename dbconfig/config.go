package dbconfig

type Users struct {

	id int
	email string
	password string
	cpf string
	name string
	customer_type string
	created_at string
	is_deleted bool

}

type Transactions struct {

	id int
	payer int
	payee int
	transaction_type string
	amount int64
	created_at string
	is_valid bool

}

const PostgresDriver = "postgres"

const User = "postgres"

const Host = "localhost"

const Port = "5432"

const Password = "123456"

const DbName = "postgres"

