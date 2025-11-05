package apis

import (
	"backend-service/model"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func first(c *gin.Context) {
	log.Print("Request Cominfg to /first endpoint")
		rand.Seed(time.Now().UnixNano())

		if rand.Float32() < 0.4 {
			log.Print("Simulating delay...")
			delay := time.Duration(rand.Intn(2000)+1000) * time.Millisecond
			time.Sleep(delay)
		}
	
		if rand.Float32() < 0.3 {
			log.Print("Simulating error...")
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "simulated backend failure",
			})
			return
		}
	

	resp := make(map[string]string)
	resp["message"] = "Hello World 1"
	c.JSON(http.StatusOK, resp)	
}

func second(c *gin.Context) {

	var req model.SecondRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp := gin.H{
		"message": "Success",
		"name":    req.Name,
		"age":   req.Age,
	}

	c.JSON(http.StatusOK, resp)
}