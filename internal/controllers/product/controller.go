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
		group.POST("", c.getAllByMultipleNameHandler)
	}
}

// GetAllProductsByName godoc
// @Summary Get all prices
// @Description This method returns all prices
// @Accept json
// @Produce json
// @tags Products
// @Param name path string true "Product name"
// @Param category query string false "Category" Enums(Books)
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

// GetAllProductsByMultipleName godoc
// @Summary Get all prices with their total
// @Description This method returns all prices with their total price
// @Accept json
// @Produce json
// @tags Products
// @Param request body []string true "Products"
// @Param category query string true "Category" Enums(Books)
// @Success 200 {object} []products.MultipleDto "Success"
// @Router /api/products [post]
func (c controller) getAllByMultipleNameHandler(g *gin.Context) {
	var names []string
	if err := g.ShouldBind(&names); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": "The request data is invalid!"})
		return
	}

	category := g.Query("category")
	dtos, err := c.service.GetAllMultipleByCategory(names, category)

	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	g.JSON(http.StatusOK, dtos)
}
