package product

import (
	"net/http"

	"github.com/bozd4g/comparator/internal/services/products"
	"github.com/gin-gonic/gin"
)

func New(service products.Servicer) Controller {
	return Controller{service: service}
}

func (c Controller) Init(e *gin.Engine) {
	group := e.Group("api/products")
	{
		group.GET(":name", c.GetAllByNameHandler)
		group.POST("", c.GetAllByMultipleNameHandler)
	}
}

// GetAllProductsByName godoc
// @Summary Get all prices
// @Description This method returns all prices
// @Accept json
// @Produce json
// @tags Products
// @Param name path string true "Product name"
// @Param config query string true "Config" Enums(Books)
// @Success 200 {object} []products.Dto "Success"
// @Router /api/products/{name} [get]
func (c Controller) GetAllByNameHandler(g *gin.Context) {
	name := g.Param("name")
	if name == "" {
		g.JSON(http.StatusBadRequest, gin.H{"error": "Name cannot be empty!"})
		return
	}

	var err error
	var dtos []products.Dto

	config := g.Query("config")
	if config == "" {
		dtos, err = c.service.GetAll(name)
	} else {
		dtos, err = c.service.GetAllByConfig(name, config)
	}

	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	g.JSON(http.StatusOK, dtos)
}

// GetAllProductsByMultipleName godoc
// @Summary Get all prices with their total
// @Description This method returns all prices with their total price
// @Accept json
// @Produce json
// @tags Products
// @Param request body []string true "Products"
// @Param config query string true "Config" Enums(Books)
// @Success 200 {object} []products.MultipleDto "Success"
// @Router /api/products [post]
func (c Controller) GetAllByMultipleNameHandler(g *gin.Context) {
	var names []string
	if err := g.ShouldBind(&names); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": "The request data is invalid!"})
		return
	}

	config := g.Query("config")
	dtos, err := c.service.GetAllMultipleByConfig(names, config)

	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	g.JSON(http.StatusOK, dtos)
}
