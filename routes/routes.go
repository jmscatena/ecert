package routes

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jmscatena/ecert-back-go/controllers"
	"github.com/jmscatena/ecert-back-go/models"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("/")
	{
		user := main.Group("user")
		{
			var obj models.Usuario
			user.POST("/", func(context *gin.Context) {
				//var obj models.Usuario
				controllers.Add[models.Usuario](context, &obj)
			})
			user.GET("/:id", func(context *gin.Context) {
				//var rec models.Usuario
				uid, _ := strconv.ParseInt(context.Param("id"), 10, 64)
				controllers.Get[models.Usuario](context, &obj, uint64(uid))
			})

			user.DELETE("/", controllers.Index)
			user.PATCH("/", controllers.Index)
		}
		inst := main.Group("instituicao")
		{
			inst.POST("/", func(context *gin.Context) {
				var obj models.Instituicao
				controllers.Add[models.Instituicao](context, &obj)
			})
			inst.GET("/", controllers.Index)
			inst.DELETE("/", controllers.Index)
			inst.PATCH("/", controllers.Index)
		}
		event := main.Group("evento")
		{
			event.POST("/", func(context *gin.Context) {
				var obj models.Evento
				controllers.Add[models.Evento](context, &obj)
			})
			event.GET("/:id", func(context *gin.Context) {
				var obj models.Evento
				controllers.Add[models.Evento](context, &obj)
			})
			event.DELETE("/", controllers.Index)
			event.PATCH("/", controllers.Index)
		}
		cert := main.Group("cert")
		{
			cert.POST("/", func(context *gin.Context) {
				var obj models.Certificado
				controllers.Add[models.Certificado](context, &obj)
			})
			cert.GET("/", controllers.Index)
			cert.DELETE("/", controllers.Index)
			cert.PATCH("/", controllers.Index)
		}
		certval := main.Group("valida")
		{
			certval.POST("/", func(context *gin.Context) {
				var obj models.CertVal
				controllers.Add[models.CertVal](context, &obj)
			})
			certval.GET("/", controllers.Index)
			certval.DELETE("/", controllers.Index)
			certval.PATCH("/", controllers.Index)
		}

		//main.POST("login", controllers.Login)
	}

	return router
}
