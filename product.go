// product.go
package main

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type product struct {
	ProductID int    `json:"Product_id"`
	Name      string `json:"name"`
}

func getProducts(c *gin.Context) {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/lumina")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer db.Close()

	rows, err := db.Query("SELECT product_id, name FROM products") // Only select the relevant columns
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var products []product
	for rows.Next() {
		var u product
		err := rows.Scan(&u.ProductID, &u.Name)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		products = append(products, u)
	}

	c.IndentedJSON(http.StatusOK, products)
}

func getProductByID(c *gin.Context) {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/lumina")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer db.Close()

	productIDStr := c.Param("id")
	productID, err := strconv.Atoi(productIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	row := db.QueryRow("SELECT product_id, name FROM products WHERE product_id = ?", productID)

	var p product
	err = row.Scan(&p.ProductID, &p.Name)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, p)
}
