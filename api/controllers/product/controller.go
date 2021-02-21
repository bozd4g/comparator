package product

import (
	"github.com/bozd4g/comparator/internal/application/comparators"
	"github.com/gin-gonic/gin"
)

func New(service comparators.Service) Controller {
	return controller{service: service}
}

func (c controller) Init(e *gin.Engine) {
	group := e.Group("api/products")
	{
		group.GET("", c.getAllHandler)
	}
}

// @Summary Get all prices
// @Description This method returns all prices
// @Accept json
// @Produce json
// @tags Products
// @Success 200 {object} []comparators.ProductDto "Success"
// @Router /api/products/{name}/list [get]
func (c controller) getAllHandler(g *gin.Context) {
	g.JSON(200, gin.H{"success": true})
}
