package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, gin.H{"message": "User is successfully created."})
}
