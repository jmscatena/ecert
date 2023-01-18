package controllers

import (
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/jmscatena/ecert-back-go/interfaces"
	"github.com/jmscatena/ecert-back-go/services"
)

func Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": "acesso " + http.StatusText(200)})

}

func Add[T interfaces.Tables](c *gin.Context, o interfaces.PersistenceHandler[T]) {
	if reflect.TypeOf(o) != nil {
		if err := c.ShouldBindJSON(&o); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusText(http.StatusBadRequest), "data": err})
		}
		var handler interfaces.PersistenceHandler[T] = o
		code, cerr := services.New[T](handler)

		if cerr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusText(http.StatusBadRequest), "data": cerr})
		}
		c.JSON(http.StatusCreated, gin.H{"status": http.StatusText(http.StatusCreated), "data": code})
	}
}

func Get[T interfaces.Tables](c *gin.Context, o interfaces.PersistenceHandler[T], uid uint64) {
	if reflect.TypeOf(o) != nil {
		if uid == 0 {
			c.JSON(http.StatusNotFound, gin.H{"status": http.StatusText(http.StatusNotFound), "data": "No Data"})
		}
		var handler interfaces.PersistenceHandler[T] = o
		rec, cerr := services.Find[T](handler, uid)

		if cerr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusText(http.StatusBadRequest), "data": cerr})
		}
		c.JSON(http.StatusOK, gin.H{"data": rec, "status": http.StatusText(http.StatusOK)})
	}
}
