package controllers

import "github.com/gin-gonic/gin"

func CommonController(ctx *gin.Context) {
	ctx.String(200, "Welcome to Tax Calculator API")
}
