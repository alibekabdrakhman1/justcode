package main

import (
	"github.com/alibekabdrakhman/justcode/lecture8/storage"
	"github.com/gin-gonic/gin"
	"net/http"
)

const authToken = "tokenXcxzcasdKLDSAdxc"

func main() {

	r := gin.Default()

	r.Use(func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token != authToken {
			c.JSON(http.StatusForbidden, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}
		c.Next()
	})
	v1 := r.Group("/api/v1")
	{
		v1.POST("/user", createUser)
		v1.GET("/user/:id", getUser)
		v1.DELETE("/user/:id", deleteUser)
		v1.GET("/users", getAllUsers)
	}

	r.Run(":3000")
}
func createUser(c *gin.Context) {
	repo, err := storage.NewStorage(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error creating repo in createUser handler": err.Error()})
		return
	}

}
func getUser(c *gin.Context) {
	repo, err := storage.NewStorage(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error creating repo in getUser handler": err.Error()})
		return
	}

}
func deleteUser(c *gin.Context) {
	repo, err := storage.NewStorage(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error creating repo in deleteUser handler": err.Error()})
		return
	}
}
func getAllUsers(c *gin.Context) {
	repo, err := storage.NewStorage(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error creating repo in getAllUsers handler": err.Error()})
		return
	}
}
