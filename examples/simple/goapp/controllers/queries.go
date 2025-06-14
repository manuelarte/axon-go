package controllers

import (
	"github.com/gin-gonic/gin"
	"goapp/api"
	"goapp/constants"
	"net/http"
)

type QueryController struct {
	UserReadProjection api.UserReadProjection
}

func NewQueryController(p api.UserReadProjection) QueryController {
	return QueryController{UserReadProjection: p}
}

func (c QueryController) Post(ctx *gin.Context) {
	queryName := ctx.GetHeader("Axoniq-Queryname")
	switch queryName {
	case constants.GetUserByIDQueryType:
		c.getByID(ctx)
		return
	default:
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "query not found"})
	}
}

func (c QueryController) getByID(ctx *gin.Context) {
	var query api.GetUserByIDQuery
	if err := ctx.ShouldBindBodyWithJSON(&query); err != nil {
		ctx.JSON(400, gin.H{"msg": err.Error()})
		return
	}
	user, err := c.UserReadProjection.GetUserByID(ctx, query)
	if err != nil {
		ctx.JSON(400, gin.H{"msg": err.Error()})
		return
	}
	// TODO(manuelarte): I need to set the metadata
	ctx.Header("AxonIQ-PayloadType", constants.UserReadType)
	ctx.JSON(200, user)
}
