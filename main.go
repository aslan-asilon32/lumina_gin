// main.go
package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql" // Import the MySQL driver
)

func main() {
	router := gin.Default()
	router.GET("/users", getUsers)
	router.GET("/users/:id", getUserByID)
	router.GET("/products", getProducts)
	router.GET("/products/:id", getProductByID)

	router.Run("localhost:8585")
}
