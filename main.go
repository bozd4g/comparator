package main

import (
	"fmt"
	"github.com/bozd4g/comparator/internal/app"
)

// @title Comparator API
// @version 1.0
// @description This is a price comparator application.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email me@furkanbozdag.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath /
func main() {
	application := app.New()
	err := application.Build().Run()
	if err != nil {
		panic(fmt.Sprintf("App cannot be started! Error %+v", err))
	}
}
