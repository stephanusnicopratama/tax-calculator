package boot

import (
	"github.com/spf13/viper"
	"log"
	"tax-calculator/core"
)

func BootConfig() {
	core.Globals.Config = viper.New()
	core.Globals.Config.SetConfigFile("docker-compose.yml")
	err := core.Globals.Config.ReadInConfig()

	if err != nil {
		log.Fatalln(err)
	}

}