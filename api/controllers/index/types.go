package index

import (
	"github.com/gin-gonic/gin"
)

type Controller interface {
	Init(e *gin.Engine)
}

type controller struct {
}
