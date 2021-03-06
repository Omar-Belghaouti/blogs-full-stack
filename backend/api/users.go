package api

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// User struct
type User struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Blogs     []int  `json:"blogs"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

// createUserRequest struct
type createUserRequest struct {
	Name string `json:"name"`
}

// updateUserRequest struct
type updateUserRequest struct {
	Name string `json:"name"`
}

// getUsers returns all users
func (s *Server) getUsers(c *gin.Context) {
	c.JSON(200, s.db.Users)
}

// createUser creates a new user
func (s *Server) createUser(c *gin.Context) {
	var request createUserRequest
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := User{
		ID:        len(s.db.Users) + 1,
		Name:      request.Name,
		Blogs:     []int{},
		CreatedAt: time.Now().UTC().Format(time.RFC3339),
		UpdatedAt: time.Now().UTC().Format(time.RFC3339),
	}
	s.db.Users = append(s.db.Users, user)

	c.JSON(http.StatusCreated, user)
}

// getUser returns a user by id
func (s *Server) getUser(c *gin.Context) {
	_id := c.Param("id")
	id, err := strconv.Atoi(_id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// find user by id
	found := false
	for _, user := range s.db.Users {
		if user.ID == id {
			found = true
		}
	}
	if !found {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}
	c.JSON(http.StatusOK, s.db.Users[id-1])
}

// updateUser updates a user by id
func (s *Server) updateUser(c *gin.Context) {
	_id := c.Param("id")
	id, err := strconv.Atoi(_id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var request updateUserRequest
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// find user by id
	found := false
	for _, user := range s.db.Users {
		if user.ID == id {
			found = true
		}
	}
	if !found {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}
	s.db.Users[id-1].Name = request.Name
	s.db.Users[id-1].UpdatedAt = time.Now().UTC().Format(time.RFC3339)
	c.JSON(http.StatusOK, s.db.Users[id-1])
}

// deleteUser deletes a user by id
func (s *Server) deleteUser(c *gin.Context) {
	_id := c.Param("id")
	id, err := strconv.Atoi(_id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// find user by id
	found := false
	for _, user := range s.db.Users {
		if user.ID == id {
			found = true
		}
	}
	if !found {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}
	// delete user
	s.db.Users = append(s.db.Users[:id-1], s.db.Users[id:]...)
	c.JSON(http.StatusOK, gin.H{"message": "user deleted"})
}
