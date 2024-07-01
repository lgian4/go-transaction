package routes

import (
	transactionController "finance/controller/transactions"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	transactions := server.Group("/transactions")

	transactions.GET("/", transactionController.GetAll)
	transactions.POST("/", transactionController.Create)
	transactions.GET("/:id", transactionController.GetOne)
	transactions.PATCH("/:id", transactionController.Update)
	transactions.DELETE("/:id", transactionController.Delete)

}
