package controllers

import "github.com/gin-gonic/gin"

type ActuatorControllers struct{}

func (c ActuatorControllers) Info(ctx *gin.Context) {
	ctx.JSON(200, gin.H{})
}
