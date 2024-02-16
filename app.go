package main

import (
	"net/http"
	"strconv"

	"rinhadev/api"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	accountsService := api.AccountsService{}

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/clients/:id/extrato", func(c *gin.Context) {
		id_param := c.Param("id")
		accountId, err := strconv.ParseInt(id_param, 10, 64)

		if err != nil {
			c.Status(http.StatusNotFound)
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

		c.JSON(http.StatusOK, statement)
	})

	r.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
