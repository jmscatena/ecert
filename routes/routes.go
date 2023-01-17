package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jmscatena/ecert-back-go/controllers"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("/")
	{
		user := main.Group("user")
		{
			user.POST("/", controllers.NovoUsuario)
			user.GET("/", controllers.IndexUsuario)
			user.DELETE("/", controllers.IndexUsuario)
			user.PATCH("/", controllers.IndexUsuario)
		}

		/*books := main.Group("books", middlewares.Auth())
		{
			books.GET("/", controllers.ShowAllBooks)
			books.GET("/:id", controllers.ShowBook)
			books.POST("/", controllers.CreateBook)
			books.PUT("/", controllers.EditBook)
			books.DELETE("/:id", controllers.DeleteBook)
		}
		*/
		//main.POST("login", controllers.Login)
	}

	return router
}
