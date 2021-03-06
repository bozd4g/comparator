package app

import (
	"github.com/bozd4g/comparator/internal/controllers/config"
	"github.com/bozd4g/comparator/internal/controllers/index"
	"github.com/bozd4g/comparator/internal/controllers/product"
	"github.com/bozd4g/comparator/internal/services/products"
	"github.com/bozd4g/comparator/internal/services/configs"
	"net/http"

	_ "github.com/bozd4g/comparator/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func (a *application) AddControllers() *application {
	a.InitIndexController()
	a.InitBookController()
	a.InitConfigController()
	return a
}

func (a *application) InitIndexController() {
	index.New().Init(a.engine)
}

func (a *application) InitBookController() {
	configService := configs.New()
	service := products.New(configService)
	product.New(service).Init(a.engine)
}

func (a *application) InitConfigController() {
	service := configs.New()
	config.New(service).Init(a.engine)
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
