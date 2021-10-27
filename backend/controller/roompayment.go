package controller

import (
	"net/http"

	"github.com/TPncy/sapro/entity"
	"github.com/gin-gonic/gin"
)

// POST /roompayment
func CreateRoomPayment(c *gin.Context) {
	var roompayment entity.RoomPayment
	if err := c.ShouldBindJSON(&roompayment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&roompayment).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": roompayment})
}

// GET /roompayment/:id
func GetRoomPayment(c *gin.Context) {
	var roompayment entity.RoomPayment
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM room_payments WHERE id = ?", id).Scan(&roompayment).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": roompayment})
}

// GET /roompayments
func ListRoomPayments(c *gin.Context) {
	var roompayments []entity.RoomPayment
	if err := entity.DB().Raw("SELECT * FROM room_payments").Scan(&roompayments).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": roompayments})
}

// DELETE /roompayments/:id
func DeleteRoomPayment(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM room_payments WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "roompayment not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /roompayments
func UpdateRoomPayment(c *gin.Context) {
	var roompayment entity.RoomPayment
	if err := c.ShouldBindJSON(&roompayment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", roompayment.ID).First(&roompayment); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "roompayment not found"})
		return
	}

	if err := entity.DB().Save(&roompayment).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": roompayment})
}
