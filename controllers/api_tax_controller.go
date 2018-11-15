package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"tax-calculator/models"
	"tax-calculator/services"
)

var taxServices services.ITaxServices

func TaxAPIFindAllController(ctx *gin.Context) {
	taxServices = &services.TaxServices{}
	ctx.JSON(200, taxServices.FindAll())
}

func TaxAPILoadController(ctx *gin.Context) {
	taxServices = &services.TaxServices{}
	ctx.JSON(200, taxServices.Load(ctx.Param("taxId")))
}

func TaxAPICreateController(ctx *gin.Context) {
	taxServices = &services.TaxServices{}
	payload := &models.Tax{
		TaxCode: cast.ToInt(ctx.PostForm("tax_code")),
		Price:   cast.ToFloat64(ctx.PostForm("price")),
		Name:    ctx.PostForm("name"),
	}
	status := taxServices.Create(payload)
	ctx.HTML(200, "create_tax.tmpl", gin.H{
		"status": status,
		"has_status": true,
	})
}
