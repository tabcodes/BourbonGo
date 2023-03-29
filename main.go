package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	v1 := r.Group("/v1")
	{

		v1.GET("bourbons", getBourbons)
		v1.GET("bourbon/:id", getBourbonById)

		v1.POST("bourbon", addBourbon)
		v1.PUT("bourbon/:id", updateBourbon)
		v1.DELETE("bourbon/:id", deleteBourbon)

	}

	r.Run()
}

func getBourbons(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "getBourbons Called"})
}

func getBourbonById(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{"message": "getBourbonById " + id + " Called"})
}

func addBourbon(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "addBourbon Called"})
}

func updateBourbon(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "updateBourbon Called"})
}

func deleteBourbon(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{"message": "deleteBourbon " + id + " Called"})
}
