package server

import (
	item "item-stock/controller"

	"github.com/gin-gonic/gin"
)



func Init()  {
	r := router()
	r.Run()
}

func router() *gin.Engine {
	r := gin.Default()

	i := r.Group("/items")
	{
		ctrl := item.Controller{}
		i.GET("", ctrl.Index)
		i.GET("/:id", ctrl.Show)
		i.POST("", ctrl.Create)
		i.PUT("/:id", ctrl.Update)
		i.DELETE("/:id", ctrl.Delete)
	}

	return r
}