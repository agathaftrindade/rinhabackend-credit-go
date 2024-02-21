package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"rinhadev/api/presenters"
	"rinhadev/api/services"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func openDB() (*pgxpool.Pool, error) {
	db_url := os.Getenv("DATABASE_URL")

	return pgxpool.New(context.Background(), db_url)
}

func main() {

	err := godotenv.Load()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error is occurred  on .env file please check", err)
		os.Exit(1)
	}

	dbpool, err := openDB()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer dbpool.Close()

	r := gin.Default()

	accountsService := services.NewAccountsService(dbpool)
	transactionsService := services.NewTransactionsService(dbpool)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/clientes/:id/extrato", func(c *gin.Context) {

		id_param := c.Param("id")
		accountId, err := strconv.ParseInt(id_param, 10, 64)

		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}

		statement, err := accountsService.GetStatement(accountId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		if statement == nil {
			c.Status(http.StatusNotFound)
			return
		}

		c.IndentedJSON(http.StatusOK, presenters.NewStatementShowResponse(*statement))
	})

	r.POST("/clientes/:id/transacoes", func(c *gin.Context) {

		id_param := c.Param("id")
		accountId, err := strconv.ParseInt(id_param, 10, 64)

		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}

		var payload presenters.TransactionCreateRequest
		if err := c.ShouldBindJSON(&payload); err != nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
			return
		}

		balance, err := transactionsService.CreateTransaction(accountId, payload)
		if err != nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"error": err.Error(),
			})
			return
		}

		if balance == nil {
			c.Status(http.StatusNotFound)
			return
		}

		c.JSON(http.StatusOK, presenters.NewTransactionCreateResponse(*balance))
	})

	r.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

	// var greeting string
	// err = dbpool.QueryRow(context.Background(), "select 'Hello, world!'").Scan(&greeting)
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
	// 	os.Exit(1)
	// }

	// fmt.Println(greeting)
}
