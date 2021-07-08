package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/eduardosbcabral/codebank/domain"
	"github.com/eduardosbcabral/codebank/infrastructure/repository"
	"github.com/eduardosbcabral/codebank/usecase"
	_ "github.com/lib/pq"
)

func main() {
	db := setupDb()
	defer db.Close()

	cc := domain.NewCreditCard()
	cc.Number = "1234"
	cc.Name = "Eduardo"
	cc.ExpirationMonth = 12
	cc.ExpirationYear = 2028
	cc.CVV = 123
	cc.Limit = 1000
	cc.Balance = 0

	repo := repository.NewTransactionRepositoryDb(db)
	err := repo.CreateCreditCard(*cc)
	if err != nil {
		fmt.Println(err)
	}
}

func setupTransactionUseCase(db *sql.DB) usecase.UseCaseTransaction {
	transactionRepository := repository.NewTransactionRepositoryDb(db)
	usecase := usecase.NewUseCaseTransaction(transactionRepository)
	return usecase
}

func setupDb() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		"host.docker.internal",
		"5432",
		"postgres",
		"root",
		"codebank",
	)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal("error connecting to database")
	}

	return db
}
