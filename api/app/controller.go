package app

import (
	"github.com/bozd4g/comparator/api/controllers/index"
	"github.com/bozd4g/comparator/api/controllers/product"
	"github.com/bozd4g/comparator/internal/application/comparators"
	"net/http"

	_ "github.com/bozd4g/comparator/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func (a *application) AddControllers() *application {
	a.InitIndexController()
	a.InitBookController()
	return a
}

func (a *application) InitIndexController() {
	index.New().Init(a.engine)
}

func (a *application) InitBookController() {
	service := comparators.New()
	product.New(service).Init(a.engine)
}

func (a *application) InitMiddlewares() *application {
	a.engine.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	return a
}

func (a *application) AddSwagger() *application {
	a.engine.GET("/swagger/*any", ginSwagger.DisablingWrapHandler(swaggerFiles.Handler, ""))
	a.engine.NoRoute(func(context *gin.Context) {
		context.JSON(http.StatusNotFound, gin.H{"error": "Not found"})
	})

	return a
}
