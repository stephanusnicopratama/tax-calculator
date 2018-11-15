package boot

import (
	"github.com/gin-gonic/gin"
	"tax-calculator/controllers"
	"tax-calculator/core"
)

func BootServer() {
	core.Globals.Router = gin.Default()
	core.Globals.Router.LoadHTMLGlob("views/*")
	core.Globals.Router.GET("/", func(context *gin.Context) { controllers.CommonController(context) })
	core.Globals.Router.GET("/create-tax", func(context *gin.Context) { controllers.TaxViewCreateController(context) })
	core.Globals.Router.GET("/tax/:taxId", func(context *gin.Context) { controllers.TaxViewObjectController(context) })
	core.Globals.Router.GET("/tax", func(context *gin.Context) { controllers.TaxViewBillController(context) })
	api := core.Globals.Router.Group("/api")
	{
		api.GET("/tax", func(context *gin.Context) { controllers.TaxAPIFindAllController(context) })
		api.GET("/tax/:taxId", func(context *gin.Context) { controllers.TaxAPILoadController(context) })
		api.POST("/tax", func(context *gin.Context) { controllers.TaxAPICreateController(context) })
	}
}
