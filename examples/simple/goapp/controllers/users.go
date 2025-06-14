package controllers

import (
	"github.com/gin-gonic/gin"
)

type UserController struct {
}

func (c UserController) GetByID(ctx *gin.Context) {
	// TODO(manuelarte): to call axon server query
}

func NewUserController() UserController {
	return UserController{}
}
