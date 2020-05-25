package routers

import (
	"log"
	"net/http"

	swaggerFiles "github.com/swaggo/files" // swaggo embed files
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
	"github.com/opentracing-contrib/go-gin/ginhttp"
	"github.com/opentracing/opentracing-go"

	"github.com/hashfyre/sample-go-app/app/handlers"
	"github.com/hashfyre/sample-go-app/app/middlewares"
	"github.com/hashfyre/sample-go-app/app/tracer"
	"github.com/hashfyre/sample-go-app/app/types"
	_ "github.com/hashfyre/sample-go-app/docs" // Swaggo dependency
)

var router *gin.Engine

// Setup - sets up gin routes with specific handlers
// chains middlewares
func Setup() *gin.Engine {
	router := gin.New()
	router.GET("/healthz/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "sample-go-app live"})
	})

	// Setting up middlewares

	// Defaults - need to make this optional only on test mode
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// Context Metadata
	router.Use(middleware.AppContext())

	// tracing
	tracer, closer, err := tracer.InitTracer()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("initialized tracer: ", tracer)
	opentracing.SetGlobalTracer(tracer)
	defer closer.Close()
	router.Use(ginhttp.Middleware(tracer))

	// Setting up api group
	// Setup Controllers for each
	api := router.Group("api")
	v1 := api.Group("v1")

	user := v1.Group("user")
	handlers.RegisterNoAuthUserRoutes(user)

	users := v1.Group("users")
	v1.Use(middleware.BasicAuth())
	handlers.RegisterUsersRoutes(users)

	// Set up Swaggo path here
	// Currently commened will need to enable this only in test mode
	// router.Static("/swagger", "./swagger")
	// docs := admin.Group("docs")
	router.GET("/docs/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.NoRoute(func(c *gin.Context) {
		c.AbortWithStatusJSON(http.StatusNotFound, types.ResponseError{
			Code:    "ERR_PAGE_NOT_FOUND",
			Message: "Page not found",
		})
	})

	return router
}
