package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"BourbonGo/models"
)

func main() {

	err := models.ConnectDatabase()
	checkErr(err)

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

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func getBourbons(c *gin.Context) {
	bourbons, err := models.GetBourbons(20)
	checkErr(err)

	if bourbons == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No bourbons found!"})
	} else {
		c.JSON(http.StatusOK, gin.H{"data": bourbons})
	}
}

func getBourbonById(c *gin.Context) {
	id := c.Param("id")

	bourbon, err := models.GetBourbonById(id)
	checkErr(err)
	if bourbon == (models.Bourbon{}) {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
	} else {
		c.JSON(http.StatusOK, gin.H{"data": bourbon})
	}

}

func addBourbon(c *gin.Context) {
	var json models.Bourbon

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	success, err := models.CreateBourbon(json)

	if success {
		c.JSON(http.StatusCreated, gin.H{"status": "created"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}

}

func updateBourbon(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "updateBourbon Called"})
}

func deleteBourbon(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{"message": "deleteBourbon " + id + " Called"})
}
