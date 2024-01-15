// user.go
package main

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type user struct {
	UserID int    `json:"user_id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
}

func getUsers(c *gin.Context) {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/lumina")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer db.Close()

	rows, err := db.Query("SELECT user_id, name, email FROM users") // Only select the relevant columns
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var users []user
	for rows.Next() {
		var u user
		err := rows.Scan(&u.UserID, &u.Name, &u.Email)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		users = append(users, u)
	}

	c.IndentedJSON(http.StatusOK, users)
}

func getUserByID(c *gin.Context) {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/lumina")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer db.Close()

	userIDStr := c.Param("id")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	row := db.QueryRow("SELECT user_id, name, email FROM users WHERE user_id = ?", userID)

	var u user
	err = row.Scan(&u.UserID, &u.Name, &u.Email)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, u)
}
