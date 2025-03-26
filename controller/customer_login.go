package controller

import (
	"final-exam/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// เชื่อมต่อกับฐานข้อมูล (ควรตั้งค่าผ่าน config)
var db *gorm.DB

// ฟังก์ชัน SetDatabase สำหรับการเชื่อมต่อฐานข้อมูล
func SetDatabase(database *gorm.DB) {
	db = database
}

// API: ล็อกอิน
func Login(c *gin.Context) {
	var loginData model.LoginRequest

	// อ่าน JSON request
	if err := c.BindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// ค้นหาข้อมูลลูกค้าจากอีเมล และรหัสผ่านในฐานข้อมูล
	var customer model.Customer
	err := db.Where("email = ? AND password = ?", loginData.Email, loginData.Password).First(&customer).Error
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	// สร้าง token หรือสามารถใช้ JWT สำหรับการยืนยันตัวตน
	token := "your-generated-jwt-token"

	// ส่งข้อมูลลูกค้ากลับ (ไม่รวมรหัสผ่าน)
	c.JSON(http.StatusOK, model.LoginResponse{
		Customer: customer,
		Token:    token,
	})
}

// API: อัปเดตที่อยู่ของลูกค้า
func UpdateAddress(c *gin.Context) {
	var updateData model.UpdateAddressRequest

	// อ่านข้อมูลที่อยู่ใหม่จาก JSON request
	if err := c.BindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// รับข้อมูลจาก path parameter (customer_id)
	customerID := c.Param("customer_id")

	// อัปเดตที่อยู่ในฐานข้อมูล
	var customer model.Customer
	if err := db.Where("customer_id = ?", customerID).First(&customer).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Customer not found"})
		return
	}

	// เปลี่ยนที่อยู่
	customer.Address = updateData.Address
	if err := db.Save(&customer).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update address"})
		return
	}

	// ส่งข้อมูลที่อัปเดตกลับไป
	c.JSON(http.StatusOK, gin.H{"message": "Address updated successfully"})
}

func Cus_login(router *gin.Engine) {
	// ตั้งเส้นทางสำหรับ Login
	router.POST("/auth/login", Login)

	// ตั้งเส้นทางสำหรับการอัปเดตที่อยู่ (ใช้ PUT)
	router.PUT("/customer/:customer_id/address", UpdateAddress)
}
