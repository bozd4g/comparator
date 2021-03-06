package product

import (
	"github.com/bozd4g/comparator/internal/services/products"
	"github.com/gin-gonic/gin"
)

type Controller interface {
	Init(e *gin.Engine)
}

type controller struct {
	service products.Service
}
