package boot

import (
	"fmt"
	"strings"
	"tax-calculator/core"
)

func Bootstrap() {
	BootConfig()
	BootDatabase()
	BootServer()
}

func Run() {
	configPort := core.Globals.Config.GetStringSlice("services.app.ports")
	ports := strings.Split(configPort[0],":")
	port := fmt.Sprintf(":%s", ports[0])
	core.Globals.Router.Run(port)
}