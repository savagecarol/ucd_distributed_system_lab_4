package apis

import  "github.com/gin-gonic/gin"
import "net/http"
import "backend-service/model"

func first(c *gin.Context) {

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