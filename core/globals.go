package core

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

type MGlobals struct {
	Router    *gin.Engine
	Config    *viper.Viper
	DB				*gorm.DB
}

var Globals *MGlobals

func init() {
	if Globals == nil {
		Globals = &MGlobals{}
	}
}