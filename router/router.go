package router

import (
	"gin-quick-start/controllers"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	v1 := r.Group("/v1")
	{
		v1.GET("/test", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "test",
			})
		})
		v1.GET("/test_controller", controllers.TestController)
		//v1.POST("book", Controllers.AddNewBook)
		//v1.GET("book/:id", Controllers.GetOneBook)
		//v1.PUT("book/:id", Controllers.PutOneBook)
		//v1.DELETE("book/:id", Controllers.DeleteBook)
	}

	return r

}
