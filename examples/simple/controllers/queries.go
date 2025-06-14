package controllers

import (
	"github.com/gin-gonic/gin"
	"goapp/api"
)

type QueryController struct {
	UserReadProjection api.UserReadProjection
}

func (c QueryController) Post(ctx *gin.Context) {
	ctx.JSON(200, gin.H{})
}
