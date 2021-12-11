package config

import (
	"net/http"

	"github.com/bozd4g/comparator/internal/services/configs"
	"github.com/gin-gonic/gin"
)

func New(service configs.Servicer) Controller {
	return Controller{service: service}
}

func (c Controller) Init(e *gin.Engine) {
	group := e.Group("api/configs")
	{
		group.GET("", c.GetAllHandler)
		group.GET(":name", c.GetByNameHandler)
	}
}

// GetAllConfigs godoc
// @Summary Get all configs
// @Description This method returns all configs
// @Accept json
// @Produce json
// @tags Configs
// @Success 200 {object} []configs.Dto "Success"
// @Router /api/configs [get]
func (c Controller) GetAllHandler(g *gin.Context) {
	dtos, err := c.service.GetAll()
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": "An error occured while retrieving the confings! Please try again later."})
		return
	}

	g.JSON(http.StatusOK, dtos)
}

// GetConfigsByName godoc
// @Summary Get configs by name
// @Description This method returns spesific config
// @Accept json
// @Produce json
// @tags Configs
// @Param name path string true "Config" Enums(Books)
// @Success 200 {object} configs.Dto "Success"
// @Router /api/configs/{name} [get]
func (c Controller) GetByNameHandler(g *gin.Context) {
	name := g.Param("name")
	if name == "" {
		g.JSON(http.StatusBadRequest, gin.H{"error": "Name cannot be empty!"})
		return
	}

	dto, err := c.service.GetByName(name)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	g.JSON(http.StatusOK, dto)
}
