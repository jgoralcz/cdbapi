package routes

import (
	"github.com/jgoralcz/cdbapi/src/handlers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	// _ "github.com/jgoralcz/cdbapi/docs" // Swagger docs
	// "github.com/swaggo/gin-swagger/swaggerFiles"
	// ginSwagger "github.com/swaggo/gin-swagger"
)

// Routes is a function that binds with http to handle particular routes
// It also includes special middleware when a panic occurs and logging each request.
// Because this uses gin, the endpoints are not REST.
func Routes() *echo.Echo {
	// env := helpers.GetEnvOrDefault("ENV", "LOCAL")
	// if env == "LOCAL" {
	// 	gin.SetMode(gin.ReleaseMode)
	// }

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	v1 := e.Group("/api/v1")
	{
		characters := v1.Group("/characters")
		{
			characters.GET("/:id", handlers.CharacterByID)
			characters.GET("/", handlers.Character)
		}
		// series := v1.Group("/series")
	}

	// url := ginSwagger.URL("http://localhost:8443/swagger/doc.json")
	// r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	return e
}
