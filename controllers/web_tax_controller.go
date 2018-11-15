package controllers

import (
	"github.com/gin-gonic/gin"
	"tax-calculator/services"
)

func TaxViewBillController(ctx *gin.Context) {
	taxServices := &services.TaxServices{}
	data := taxServices.FindAll()
	ctx.HTML(200, "show_bill.tmpl", gin.H{
		"data": data,
	})
}

func TaxViewObjectController(ctx *gin.Context) {
	taxServices := &services.TaxServices{}
	result := taxServices.Load(ctx.Param("taxId"))
	ctx.HTML(200, "tax_object.tmpl", gin.H{
		"data": result,
	})
}

func TaxViewCreateController(ctx *gin.Context) {
	ctx.HTML(200, "create_tax.tmpl", gin.H{
		"has_status": false,
	})
}