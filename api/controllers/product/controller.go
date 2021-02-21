package product

import (
	"github.com/bozd4g/comparator/internal/application/comparators"
	"github.com/gin-gonic/gin"
	"net/http"
)

func New(service comparators.Service) Controller {
	return controller{service: service}
}

func (c controller) Init(e *gin.Engine) {
	group := e.Group("api/products")
	{
		group.GET(":name/list", c.getAllHandler)
	}
}

// @Summary Get all prices
// @Description This method returns all prices
// @Accept json
// @Produce json
// @tags Products
// @Param name path string true "product"
// @Success 200 {object} []comparators.ProductDto "Success"
// @Router /api/products/{name}/list [get]
func (c controller) getAllHandler(g *gin.Context) {
	name := g.Param("name")
	if name == "" {
		g.JSON(http.StatusBadRequest, gin.H{"error": "Name cannot be empty!"})
		return
	}

	g.JSON(200, gin.H{"success": name})
}
