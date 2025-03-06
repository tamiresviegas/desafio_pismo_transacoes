package main

import (
	"log"

	"github.com/tamiresviegas/desafio_pismo_transacoes/config"
	"github.com/tamiresviegas/desafio_pismo_transacoes/internal/adapters/handler"
	"github.com/tamiresviegas/desafio_pismo_transacoes/internal/adapters/http"
	"github.com/tamiresviegas/desafio_pismo_transacoes/internal/adapters/repository"
	"github.com/tamiresviegas/desafio_pismo_transacoes/internal/core/service"
)

func main() {

	db, err := config.ConnectBD()
	if err != nil {
		log.Fatal("Erro ao conectar ao banco de dados:", err)
	}
	

	accountRepo := repository.NewAccountRepository(db.DB)
	accountService := service.NewAccountService(accountRepo)
	accountHandler := handler.NewAccountHandler(accountService)

	opTypesRepo := repository.NewOperationTypesRepository(db.DB)
	opTypesRepoService := service.NewOperationTypesService(opTypesRepo)
	opTypesRepoHandler := handler.NewOperationTypesHandler(opTypesRepoService)

	transactionRepo := repository.NewTransactionRepository(db.DB)
	transactionService := service.NewTransactionService(transactionRepo, accountRepo, opTypesRepo)
	transactionHandler := handler.NewTransactionHandler(transactionService)

	r := http.SetupRoutes(accountHandler, opTypesRepoHandler, transactionHandler)
	r.Run(":8080")
}
