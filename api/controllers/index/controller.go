package index

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func New() Controller {
	return &controller{}
}

func (c controller) Init(e *gin.Engine) {
	e.GET("/", c.indexHandler)
}

// @Summary redirectToSwaggerUi
// @Description This method redirects to swagger ui
// @Accept json
// @Produce json
// @tags IndexController
// @Success 308 {string} string	"Redirect"
// @Router / [get]
func (c controller) indexHandler(g *gin.Context) {
	g.Redirect(http.StatusPermanentRedirect, "/swagger/index.html")
}
