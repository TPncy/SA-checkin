package controller

import (
	"net/http"

	"github.com/TPncy/sapro/entity"
	"github.com/gin-gonic/gin"
)

// POST /customers
func CreateCustomer(c *gin.Context) {
	var customer entity.Customer
	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&customer).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": customer})
}

// GET /customer/:id
func GetCustomer(c *gin.Context) {
	var customer entity.Customer
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM customers WHERE id = ?", id).Scan(&customer).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": customer})
}

// GET /customers
func ListCustomers(c *gin.Context) {
	var customers []entity.Customer
	if err := entity.DB().Raw("SELECT * FROM customers").Scan(&customers).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": customers})
}

// DELETE /customers/:id
func DeleteCustomer(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM customers WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "customer not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /customers
func UpdateCustomer(c *gin.Context) {
	var customer entity.Customer
	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", customer.ID).First(&customer); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "customers not found"})
		return
	}

	if err := entity.DB().Save(&customer).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": customer})
}
