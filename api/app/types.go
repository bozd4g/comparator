package app

import "github.com/gin-gonic/gin"

type Application interface {
	Build() Application
	Run() error
}

type application struct {
	engine *gin.Engine
}
