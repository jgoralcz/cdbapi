package routes

import (
	"github.com/gin-gonic/gin"
)

// Routes is a function that binds with http to handle particular routes
// It also includes special middleware when a panic occurs and logging each request.
// Because this uses gin, the endpoints are not REST.
func Routes() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/api/v1")
	{
		characters := v1.Group("/characters")
		{
			characters.GET(":id", CharacterByIDHandler)
			characters.GET("", CharacterHandler)
		}
	}

	return r
}
