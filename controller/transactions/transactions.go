package transactions

import (
	"fmt"
	"net/http"
	"strconv"

	"finance/models"

	"github.com/gin-gonic/gin"
)

func GetAll(ctx *gin.Context) {
	transactions, err := models.GetAll()
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "could not parse request data"})
		return
	}
	ctx.JSON(http.StatusOK, transactions)

}

func GetOne(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "could not parse id"})
		return
	}

	transaction, err := models.GetOne(id)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch transaction"})
		return
	}
	ctx.JSON(http.StatusOK, transaction)

}

func Create(ctx *gin.Context) {

	var transaction models.Transaction
	err := ctx.ShouldBindJSON(&transaction)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data", "error": err})

		return
	}
	transaction.Save()
	ctx.JSON(http.StatusCreated, gin.H{"message": "transaction created", "transaction": transaction})

}

func Update(ctx *gin.Context) {

	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "could not parse id"})
		return
	}

	_, err = models.GetOne(id)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch transaction"})
		return
	}

	var updateTransaction models.Transaction
	err = ctx.ShouldBindJSON(&updateTransaction)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data", "error": err})
		return
	}

	updateTransaction.ID = id
	err = updateTransaction.Update()
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not update transaction", "error": err})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "event updated", "event": updateTransaction})

}

func Delete(ctx *gin.Context) {

	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "could not parse id"})
		return
	}

	transaction, err := models.GetOne(id)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch transaction"})
		return
	}

	err = transaction.Delete()
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not delete transaction", "error": err})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "transaction deleted", "transaction": transaction})

}
