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
				controllers.Add[models.Usuario](context, &obj)
			})
			user.GET("/", func(context *gin.Context) {
				controllers.GetAll[models.Usuario](context, &obj)
			})
			user.GET("/:id", func(context *gin.Context) {
				uid, _ := strconv.ParseInt(context.Param("id"), 10, 64)
				controllers.Get[models.Usuario](context, &obj, uint64(uid))
			})
			user.GET("/admin/", func(context *gin.Context) {
				//colocar as configuracoes para os params q virao do frontend
				params := "admin=?;ativo=?"
				controllers.GetBy[models.Usuario](context, &obj, params, false, true)
			})

			user.PATCH("/:id", func(context *gin.Context) {
				uid, _ := strconv.ParseInt(context.Param("id"), 10, 64)
				controllers.Modify[models.Usuario](context, &obj, uint64(uid))
			})
			user.DELETE("/:id", func(context *gin.Context) {
				uid, _ := strconv.ParseInt(context.Param("id"), 10, 64)
				controllers.Erase[models.Usuario](context, &obj, uint64(uid))
			})

		}
		/*
			inst := main.Group("instituicao")
			{
				var obj models.Instituicao
				inst.POST("/", func(context *gin.Context) {
					controllers.Add[models.Instituicao](context, &obj)
				})
				inst.GET("/", func(context *gin.Context) {
					controllers.GetAll[models.Instituicao](context, &obj)
				})
				inst.GET("/:id", func(context *gin.Context) {
					uid, _ := strconv.ParseInt(context.Param("id"), 10, 64)
					controllers.Get[models.Instituicao](context, &obj, uint64(uid))
				})
				inst.PATCH("/:id", func(context *gin.Context) {
					uid, _ := strconv.ParseInt(context.Param("id"), 10, 64)
					controllers.Modify[models.Instituicao](context, &obj, uint64(uid))
				})
				inst.DELETE("/:id", func(context *gin.Context) {
					uid, _ := strconv.ParseInt(context.Param("id"), 10, 64)
					controllers.Erase[models.Instituicao](context, &obj, uint64(uid))
				})
			}
			event := main.Group("evento")
			{
				var obj models.Evento
				event.POST("/", func(context *gin.Context) {
					controllers.Add[models.Evento](context, &obj)
				})
				event.GET("/", func(context *gin.Context) {
					controllers.GetAll[models.Evento](context, &obj)
				})
				event.GET("/:id", func(context *gin.Context) {
					uid, _ := strconv.ParseInt(context.Param("id"), 10, 64)
					controllers.Get[models.Evento](context, &obj, uint64(uid))
				})
				event.PATCH("/:id", func(context *gin.Context) {
					uid, _ := strconv.ParseInt(context.Param("id"), 10, 64)
					controllers.Modify[models.Evento](context, &obj, uint64(uid))
				})
				event.DELETE("/:id", func(context *gin.Context) {
					uid, _ := strconv.ParseInt(context.Param("id"), 10, 64)
					controllers.Erase[models.Evento](context, &obj, uint64(uid))
				})
			}
			cert := main.Group("cert")
			{
				var obj models.Certificado
				cert.POST("/", func(context *gin.Context) {
					controllers.Add[models.Certificado](context, &obj)
				})
				cert.GET("/", func(context *gin.Context) {
					controllers.GetAll[models.Certificado](context, &obj)
				})
				cert.GET("/:id", func(context *gin.Context) {
					uid, _ := strconv.ParseInt(context.Param("id"), 10, 64)
					controllers.Get[models.Certificado](context, &obj, uint64(uid))
				})
				cert.PATCH("/:id", func(context *gin.Context) {
					uid, _ := strconv.ParseInt(context.Param("id"), 10, 64)
					controllers.Modify[models.Certificado](context, &obj, uint64(uid))
				})
				cert.DELETE("/:id", func(context *gin.Context) {
					uid, _ := strconv.ParseInt(context.Param("id"), 10, 64)
					controllers.Erase[models.Certificado](context, &obj, uint64(uid))
				})
			}
			certval := main.Group("valida")
			{
				var obj models.CertVal
				certval.POST("/", func(context *gin.Context) {
					controllers.Add[models.CertVal](context, &obj)
				})
				certval.GET("/", func(context *gin.Context) {
					controllers.GetAll[models.CertVal](context, &obj)
				})
				certval.GET("/:id", func(context *gin.Context) {
					uid, _ := strconv.ParseInt(context.Param("id"), 10, 64)
					controllers.Get[models.CertVal](context, &obj, uint64(uid))
				})
				certval.PATCH("/", func(context *gin.Context) {
					uid, _ := strconv.ParseInt(context.Param("id"), 10, 64)
					controllers.Modify[models.CertVal](context, &obj, uint64(uid))
				})
				certval.DELETE("/:id", func(context *gin.Context) {
					uid, _ := strconv.ParseInt(context.Param("id"), 10, 64)
					controllers.Erase[models.CertVal](context, &obj, uint64(uid))
				})
			}
		*/
		//main.POST("login", controllers.Login)
	}

	return router
}
