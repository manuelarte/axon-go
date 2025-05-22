package controllers

import (
	"github.com/gin-gonic/gin"
	"goapp/api"
)

type UserController struct {
	userReadProjection api.UserReadProjection
}

func NewUserController(p api.UserReadProjection) UserController {
	return UserController{
		userReadProjection: p,
	}
}

func (c UserController) GetByID(ctx *gin.Context) {
	var query api.GetUserByIDQuery
	if err := ctx.ShouldBindUri(&query); err != nil {
		ctx.JSON(400, gin.H{"msg": err.Error()})
		return
	}
	// TODO(manuelarte): THIS IS NOT CORRECT, I NEED TO SEND A QUERY
	user, err := c.userReadProjection.GetUserByID(ctx, query)
	if err != nil {
		ctx.JSON(400, gin.H{"msg": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"user": user})
}
