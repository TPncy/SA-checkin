package controller

import (
	"net/http"

	"github.com/TPncy/sapro/entity"
	"github.com/gin-gonic/gin"
)

// POST /check-ins
func CreateCheckIn(c *gin.Context) {

	var room entity.Room
	var customer entity.Customer
	var roompayment entity.RoomPayment
	var employee entity.Employee
	var checkin entity.CheckIn

	// ผลลัพธ์ที่ได้จากขั้นตอนที่ 8 จะถูก bind เข้าตัวแปร CheckIn
	if err := c.ShouldBindJSON(&checkin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 9: ค้นหา RoomPayment ด้วย id
	if tx := entity.DB().Where("id = ?", checkin.PaymentID).First(&roompayment); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "roomplayment not found"})
		return
	}

	// 10: ค้นหา Room ด้วย id
	if tx := entity.DB().Where("id = ?", checkin.ReserveID).First(&room); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "room not found"})
		return
	}

	// 11: ค้นหา Customer ด้วย id
	if tx := entity.DB().Where("id = ?", checkin.CustomerID).First(&customer); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "customer not found"})
		return
	}

	// 12: ค้นหา Employee ด้วย id
	if tx := entity.DB().Where("id = ?", checkin.EmployeeID).First(&employee); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "employee not found"})
		return
	}
	// 13: สร้าง CheckIn
	ci := entity.CheckIn{
		Reserve:  room,             // โยงความสัมพันธ์กับ Entity Room
		Customer: customer,         // โยงความสัมพันธ์กับ Entity Customer
		Payment:  roompayment,      // โยงความสัมพันธ์กับ Entity RoomPayment
		Employee: employee,         // โยงความสัมพันธ์กับ Entity Employee
		DateTime: checkin.DateTime, // ตั้งค่าฟิลด์ Date_time
	}

	// 14: บันทึก
	if err := entity.DB().Create(&ci).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": ci})
}

// GET /checkin/:id
func GetCheckIn(c *gin.Context) {
	var checkin entity.CheckIn
	id := c.Param("id")
	if err := entity.DB().Preload("room").Preload("customer").Preload("roompayment").Preload("employee").Raw("SELECT * FROM check_ins WHERE id = ?", id).Find(&checkin).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": checkin})
}

// GET /check_ins
func ListCheckIns(c *gin.Context) {
	var checkins []entity.CheckIn
	if err := entity.DB().Preload("room").Preload("customer").Preload("roompayment").Preload("employee").Raw("SELECT * FROM check_ins").Find(&checkins).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": checkins})
}

// DELETE /check_ins/:id
func DeleteCheckIn(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM check_ins WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "check_ins not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /check_ins
func UpdateCheckIn(c *gin.Context) {
	var checkin entity.CheckIn
	if err := c.ShouldBindJSON(&checkin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", checkin.ID).First(&checkin); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "check_ins not found"})
		return
	}

	if err := entity.DB().Save(&checkin).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": checkin})
}
