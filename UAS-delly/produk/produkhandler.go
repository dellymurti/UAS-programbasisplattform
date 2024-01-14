package produkhandler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Produk struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price float64 `json:"price"`
}

var produkList []Produk

func Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   produkList,
	})
}

func Show(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

		for _, p := range produkList {
		if p.ID == id {
			product = p
			break
		}
	}

	if product.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   product,
	})
}

func Create(c *gin.Context) {
	var newProduct Produk
	if err := c.ShouldBindJSON(&newProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	NewProduct.ID = len(produkList) + 1

	produkList = append(produkList, newProduct)

	c.JSON(http.StatusCreated, gin.H{
		"status": "success",
		"data":   newProduct,
	})
}

func Update(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var existingProduct Produk
	for i, p := range produkList {
		if p.ID == id {
			existingProduct = p

			var updatedProduct Produk
			if err := c.ShouldBindJSON(&updatedProduct); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			updatedProduct.ID = existingProduct.ID
			produkList[i] = updatedProduct

			c.JSON(http.StatusOK, gin.H{
				"status": "success",
				"data":   updatedProduct,
			})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
}

func Delete(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	for i, p := range produkList {
		if p.ID == id {
			produkList = append(produkList[:i], produkList[i+1:]...)

			c.JSON(http.StatusOK, gin.H{"status": "success"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
}
