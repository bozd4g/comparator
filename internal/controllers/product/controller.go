package product

import (
	"github.com/bozd4g/comparator/internal/services/products"
	"github.com/gin-gonic/gin"
	"net/http"
)

func New(service products.Service) Controller {
	return controller{service: service}
}

func (c controller) Init(e *gin.Engine) {
	group := e.Group("api/products")
	{
		group.GET(":name", c.getAllByNameHandler)
	}
}

// GetAllProductsByName godoc
// @Summary Get all prices
// @Description This method returns all prices
// @Accept json
// @Produce json
// @tags Products
// @Param name path string true "Product name"
// @Param enumstring query string false "Categories" Enums(Books)
// @Success 200 {object} []products.Dto "Success"
// @Router /api/products/{name} [get]
func (c controller) getAllByNameHandler(g *gin.Context) {
	name := g.Param("name")
	if name == "" {
		g.JSON(http.StatusBadRequest, gin.H{"error": "Name cannot be empty!"})
		return
	}

	var err error
	var dtos []products.Dto

	category := g.Query("category")
	if category == "" {
		dtos, err = c.service.GetAll(name)
	} else {
		dtos, err = c.service.GetAllByCategory(name, category)
	}

	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	g.JSON(http.StatusOK, dtos)
}
