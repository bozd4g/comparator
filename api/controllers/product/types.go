package product

import (
	"github.com/bozd4g/comparator/internal/application/comparators"
	"github.com/gin-gonic/gin"
)

type Controller interface {
	Init(e *gin.Engine)
}

type controller struct {
	service comparators.Service
}