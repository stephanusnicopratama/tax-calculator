package boot

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"tax-calculator/core"
)

func BootDatabase() {
	var err error
	username := core.Globals.Config.GetString("services.database.environment.MYSQL_USER")
	password := core.Globals.Config.GetString("services.database.environment.MYSQL_PASSWORD")
	database := core.Globals.Config.GetString("services.database.environment.MYSQL_DATABASE")
	dataSourceName := fmt.Sprintf("%s:%s@tcp(database:3306)/%s", username, password, database)
	core.Globals.DB, err = gorm.Open("mysql", dataSourceName)
	if err != nil {
		fmt.Println(err)
	}
	core.Globals.DB.LogMode(true)
}
