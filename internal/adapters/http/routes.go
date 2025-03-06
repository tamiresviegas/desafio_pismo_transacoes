package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/tamiresviegas/desafio_pismo_transacoes/cmd/docs"
	"github.com/tamiresviegas/desafio_pismo_transacoes/internal/adapters/handler"
)

// Ping godoc
// @Summary      Health check
// @Description  Retorna um status 200 se a API estiver online
// @Tags         health
// @Success      200  {string}  string  "pong"
// @Router       /ping [get]
func Ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}

func SetupRoutes(accountHandler *handler.AccountHandler, opTypesRepoHandler *handler.OperationTypesHandler, transactionHandler *handler.TransactionHandler) *gin.Engine {
	r := gin.Default()

	//Account
	r.POST("/accounts", accountHandler.CreateAccount)
	r.GET("/accounts", accountHandler.GetAllAccount)
	r.GET("/accounts/:accountId", accountHandler.GetAccountByID)
	r.DELETE("/account/:accountId", accountHandler.DeleteAccount)
	r.PATCH("/account/:accountId", accountHandler.UpdateAccount)

	//OperationType
	r.GET("/operationtypes", opTypesRepoHandler.GetAllOperationTypes)
	r.GET("/operationtypes/:operationtypeId", opTypesRepoHandler.GetOperationTypesByID)
	r.POST("/newoperationtypes", opTypesRepoHandler.CreateOperationTypes)
	r.DELETE("/operationtypes/:operationtypeId", opTypesRepoHandler.DeleteOperationTypes)
	r.PATCH("/operationtypes/:operationtypeId", opTypesRepoHandler.UpdateOperationTypes)

	//Transaction
	r.GET("/transactions", transactionHandler.GetAllTransaction)
	r.GET("/transaction/:transactions", transactionHandler.GetTransactionByID)
	r.POST("/transactions", transactionHandler.CreateTransaction)
	r.DELETE("/transaction/:transactions", transactionHandler.DeleteTransaction)
	r.PATCH("/transaction/:transactions", transactionHandler.UpdateTransaction)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
