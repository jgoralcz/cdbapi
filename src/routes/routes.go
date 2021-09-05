package routes

import (
	_ "github.com/jgoralcz/cdbapi/docs" // docs is generated by Swag CLI, you have to import it.
	"github.com/jgoralcz/cdbapi/src/handlers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// Routes is a function that binds with http to handle particular routes.
// It also includes special middleware when a panic occurs and logging each request.
func Routes() *echo.Echo {
	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	v1 := e.Group("/v1")
	{
		characters := v1.Group("/characters")
		{
			characters.GET("/:id/images", handlers.CharacterImages)
			characters.GET("/random", handlers.CharacterRandom)
			characters.GET("/:id", handlers.CharacterByID)
			characters.GET("", handlers.Character)
		}

		images := v1.Group("/images")
		{
			images.GET("/:id", handlers.ImageByID)
		}

		series := v1.Group("/series")
		{
			series.GET("/random", handlers.SeriesRandom)
			series.GET("/:id", handlers.SeriesByID)
			series.GET("", handlers.Series)
		}

		search := v1.Group("/search")
		{
			search.GET("", handlers.Search)
		}
	}

	e.GET("/swagger/*", echoSwagger.WrapHandler)
	return e
}
