package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jmscatena/ecert-back-go/interfaces"
	"github.com/jmscatena/ecert-back-go/models"
	"github.com/jmscatena/ecert-back-go/services"
)

func NovoUsuario(c *gin.Context) {
	var usuario models.Usuario

	if err := c.ShouldBindJSON(&usuario); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusText(http.StatusBadRequest), "data": err})
	}
	var handler interfaces.PersistenceHandler[models.Usuario] = &usuario
	code, cerr := services.New[models.Usuario](handler)

	if cerr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusText(http.StatusBadRequest), "data": cerr})
	}
	c.JSON(http.StatusCreated, gin.H{"status": http.StatusText(http.StatusCreated), "data": code})
}

func IndexUsuario(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": "acesso " + http.StatusText(200)})

}
