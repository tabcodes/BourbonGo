package main

import (
	"log"
	"net/http"
	"strconv"

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
	var json models.Bourbon

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Invalid ID supplied"})
	}

	bourbon, err := models.GetBourbonById(strconv.Itoa(id))
	checkErr(err)
	if bourbon == (models.Bourbon{}) {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}

	success, err := models.UpdateBourbon(id, json)
	if success {
		c.JSON(http.StatusOK, gin.H{"status": "updated"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}
}

func deleteBourbon(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	bourbon, err := models.GetBourbonById(strconv.Itoa(id))
	checkErr(err)
	if bourbon == (models.Bourbon{}) {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}

	success, err := models.DeleteBourbon(id)
	if success {
		c.JSON(http.StatusNoContent, gin.H{"status": "deleted"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}
}
