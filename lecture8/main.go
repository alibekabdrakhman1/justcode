package main

import (
	"github.com/alibekabdrakhman/justcode/lecture8/model"
	"github.com/alibekabdrakhman/justcode/lecture8/storage"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id, err := repo.User.CreateUser(c, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, id)

}
func getUser(c *gin.Context) {
	repo, err := storage.NewStorage(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error creating repo in getUser handler": err.Error()})
		return
	}
	id, _ := strconv.Atoi(c.Param("id"))

	user, err := repo.User.GetUser(c, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func deleteUser(c *gin.Context) {
	repo, err := storage.NewStorage(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error creating repo in getUser handler": err.Error()})
		return
	}
	id, _ := strconv.Atoi(c.Param("id"))

	err = repo.User.DeleteUser(c, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user deleted successfully"})
}
func getAllUsers(c *gin.Context) {
	repo, err := storage.NewStorage(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error creating repo in getAllUsers handler": err.Error()})
		return
	}
	users, err := repo.User.GetAllUsers(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch users"})
		return
	}

	c.JSON(http.StatusOK, users)
}
