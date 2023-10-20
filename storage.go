package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Storage interface {
	CreateAccount(*Account) error
	DeleteAccount(int) error
	UpdateAccount(*Account) error
	GetAccounts() ([]*Account, error)
	GetAccountByID(int) (*Account, error)
	GetAccountByNumber(int) (*Account, error)
}

type PostgresStore struct {
	db *sql.DB
}

// docker run --name gobank -e POSTGRES_PASSWORD=password -p 5432:5432 -d postgres
func NewPostgresStore() (*PostgresStore, error) {
	connStr := "user=postgres dbname=postgres password=password sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	// Failed to open db
	if err != nil {
		return nil, err
	}

	// Checking db connection
	if err := db.Ping(); err != nil {
		return nil, err
	}
	fmt.Println("connected to postgres db")

	return &PostgresStore{
		db: db,
	}, nil
}

func (s *PostgresStore) Init() error {
	return s.createAccountTable()
}

func (s *PostgresStore) createAccountTable() error {
	// Create table if it doesn't already exist
	query := `create table if not exists account (
		id serial primary key,
		first_name varchar(50),
		last_name varchar(50),
		number serial,
		encrypted_password varchar(50),
		balance serial,
		created_at timestamp
	)`

	// Response not needed, query creates table not querying from db
	_, err := s.db.Exec(query)
	return err
}

func (s *PostgresStore) CreateAccount(account *Account) error {
	query := `
	insert into account (first_name, last_name, number, encrypted_password, balance, created_at)
	values ($1, $2, $3, $4, $5, $6)
	`

	// Insertion query, no meaningful response
	_, err := s.db.Exec(query, account.FirstName, account.LastName, account.Number, account.EncryptedPassword, account.Balance, account.CreatedAt)

	return err
}

func (s *PostgresStore) UpdateAccount(*Account) error {
	return nil
}

func (s *PostgresStore) DeleteAccount(id int) error {
	_, err := s.db.Query("delete from account where id = $1", id)
	return err
}

func (s *PostgresStore) GetAccountByNumber(number int) (*Account, error) {
	response := s.db.QueryRow("select * from account where number = $1", number)

	// Decoding row response into empy Account
	account := new(Account)
	err := response.Scan(
		&account.ID,
		&account.FirstName,
		&account.LastName,
		&account.Number,
		&account.EncryptedPassword,
		&account.Balance,
		&account.CreatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("account %d not found", number)
	}

	return account, nil
}

func (s *PostgresStore) GetAccountByID(id int) (*Account, error) {
	response := s.db.QueryRow("select * from account where id = $1", id)

	// Decoding row response into empy Account
	account := new(Account)
	err := response.Scan(
		&account.ID,
		&account.FirstName,
		&account.LastName,
		&account.Number,
		&account.EncryptedPassword,
		&account.Balance,
		&account.CreatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("account %d not found", id)
	}

	return account, nil
}

func (s *PostgresStore) GetAccounts() ([]*Account, error) {
	response, err := s.db.Query("select * from account")
	if err != nil {
		return nil, err
	}

	accounts := []*Account{}

	// Preps response to be scanned into an empty Account
	for response.Next() {
		// Decoding account from db response
		account := new(Account)
		err := response.Scan(
			&account.ID,
			&account.FirstName,
			&account.LastName,
			&account.Number,
			&account.EncryptedPassword,
			&account.Balance,
			&account.CreatedAt)
		if err != nil {
			return nil, err
		}
		// Adding to account slice
		accounts = append(accounts, account)
	}

	// Successfully fetched accounts
	return accounts, nil
}
