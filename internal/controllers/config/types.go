package config

import (
	"github.com/bozd4g/comparator/internal/services/configs"
	"github.com/gin-gonic/gin"
)

type Controller interface {
	Init(e *gin.Engine)
}

type controller struct {
	service configs.Service
}
