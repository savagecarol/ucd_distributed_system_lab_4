package apis

import  "github.com/gin-gonic/gin"
import "net/http"
import "backend-service/model"
import "math/rand"
import "time"

func first(c *gin.Context) {

		rand.Seed(time.Now().UnixNano())

		if rand.Float32() < 0.3 {
			delay := time.Duration(rand.Intn(2000)+1000) * time.Millisecond // 1–3 sec
			time.Sleep(delay)
		}
	
		// Random error to simulate failure (20% chance)
		if rand.Float32() < 0.2 {
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