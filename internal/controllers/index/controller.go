package index

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func New() Controller {
	return Controller{}
}

func (c Controller) Init(e *gin.Engine) {
	e.GET("/", c.IndexHandler)
}

// @Summary redirectToSwaggerUi
// @Description This method redirects to swagger ui
// @Accept json
// @Produce json
// @tags Index
// @Success 308 {string} string	"Redirect"
// @Router / [get]
func (c Controller) IndexHandler(g *gin.Context) {
	g.Redirect(http.StatusPermanentRedirect, "/swagger/index.html")
}
