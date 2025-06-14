package controllers

import (
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	axongo "github.com/manuelarte/axon-go"
	"goapp/api"
	"goapp/constants"
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
	if err := ctx.ShouldBindBodyWithJSON(&query); err != nil {
		ctx.JSON(400, gin.H{"msg": err.Error()})
		return
	}
	user, err := c.userReadProjection.GetUserByID(ctx, query)
	if err != nil {
		ctx.JSON(400, gin.H{"msg": err.Error()})
		return
	}
	body := axongo.QueryResponse{
		Id:          constants.Ptr(uuid.NewString()),
		Payload:     constants.Ptr(structs.Map(user)),
		PayloadType: constants.Ptr(user.GetType()),
	}
	ctx.JSON(200, body)
}
