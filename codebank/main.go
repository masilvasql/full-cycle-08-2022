package main

import (
	"codebank/domain"
	"codebank/infrastructure/repository"
	"codebank/usecase"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	db := setupDb()
	defer db.Close()

	cc := domain.NewCreditCard()
	cc.Number = "1234"
	cc.Name = "Marcelo"
	cc.ExpirationYear = 2021
	cc.ExpirationMonth = 11
	cc.CVV = 123
	cc.Limit = 1000
	cc.Balance = 0

	repo := repository.NewTransactionRepositoryDb(db)
	repo.CreateCreditCard(*cc)

}

func setupTransactionUseCase(db *sql.DB) usecase.UseCaseTransaction {
	transactionRepository := repository.NewTransactionRepositoryDb(db)
	useCase := usecase.NewUseCaseTransaction(transactionRepository)
	return useCase
}

func setupDb() *sql.DB {
	psql := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		"db",
		"5432",
		"postgres",
		"root",
		"codebank")
	db, err := sql.Open("postgres", psql)
	if err != nil {
		log.Fatal("error connection to database")
	}

	return db
}
