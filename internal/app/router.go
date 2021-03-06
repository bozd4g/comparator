package app

import "github.com/gin-gonic/gin"

func (a *application) AddRouter() {
	a.engine = gin.Default()
}
