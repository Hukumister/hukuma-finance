package api

import (
	// "net/http"

	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/hukuma-finance/docs"
	"github.com/hukuma-finance/internal/api/v1"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

// InitRouter initialize routing information
func InitRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())

	router.LoadHTMLFiles("web/index.html")
	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", nil)
	})

	apiv1 := router.Group("/api/v1")
	apiv1.GET("/currencies", v1.GetArticle)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
