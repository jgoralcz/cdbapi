package routes

import (
	"github.com/gin-gonic/gin"
	_ "github.com/jgoralcz/go_cdbapi/docs" // Swagger docs
	"github.com/jgoralcz/go_cdbapi/src/handlers"
	"github.com/jgoralcz/go_cdbapi/src/helpers"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	ginSwagger "github.com/swaggo/gin-swagger"
)

// Routes is a function that binds with http to handle particular routes
// It also includes special middleware when a panic occurs and logging each request.
// Because this uses gin, the endpoints are not REST.
func Routes() *gin.Engine {
	env := helpers.GetEnvOrDefault("ENV", "LOCAL")
	if env == "LOCAL" {
		gin.SetMode(gin.ReleaseMode)
	}
	gin.DisableConsoleColor()
	r := gin.Default()

	v1 := r.Group("/api/v1")
	{
		characters := v1.Group("/characters")
		{
			characters.GET(":id", handlers.CharacterByID)
			characters.GET("", handlers.Character)
		}
		// series := v1.Group("/series")
	}

	url := ginSwagger.URL("http://localhost:8443/swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	return r
}
