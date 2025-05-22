package controllers

import "github.com/gin-gonic/gin"

type QueryController struct{}

func (c QueryController) Get(ctx *gin.Context) {
	ctx.JSON(200, gin.H{})
}

func (c QueryController) GetUserByIDQuery(ctx *gin.Context) {
	ctx.JSON(200, gin.H{})
}
