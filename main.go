package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"final-exam/controller"
	"final-exam/model"
)

func main() {
	// โหลดค่า config จากไฟล์
	viper.SetConfigName("config") // ชื่อไฟล์ config
	viper.AddConfigPath(".")      // ที่อยู่ไฟล์ config
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error reading config file:", err)
	}

	// เชื่อมต่อฐานข้อมูล
	dsn := viper.GetString("mysql.dsn")             // ดึงค่า dsn จาก config
	dialector := mysql.Open(dsn)                    // ใช้ dsn เชื่อมต่อกับ MySQL
	db, err := gorm.Open(dialector, &gorm.Config{}) // เชื่อมต่อกับฐานข้อมูล
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}
	fmt.Println("Database connection successful")

	// ดึงข้อมูลจากฐานข้อมูล (เช่นข้อมูลจากตาราง customer)
	var customers []model.Customer
	if err := db.Find(&customers).Error; err != nil {
		log.Fatalf("Error fetching customers: %v", err)
	}

	// แสดงข้อมูลลูกค้าในฐานข้อมูล
	fmt.Println("List of customers:")
	for _, customer := range customers {
		fmt.Printf("ID: %d, Name: %s %s, Email: %s\n", customer.CustomerID, customer.FirstName, customer.LastName, customer.Email)
	}

	// ตั้งค่า database ให้กับ controller
	controller.SetDatabase(db)

	// ตั้งค่า Gin router
	router := gin.Default()

	// ลงทะเบียนเส้นทาง API
	router.POST("/auth/login", controller.Login)                           // เส้นทางสำหรับ Login
	router.PUT("/customer/:customer_id/address", controller.UpdateAddress) // เส้นทางสำหรับอัปเดตที่อยู่

	// เริ่มต้นเซิร์ฟเวอร์
	port := viper.GetString("server.port") // ดึงค่าพอร์ตจาก config
	if port == "" {
		port = "8081" // ค่าเริ่มต้นถ้าไม่ได้กำหนดใน config
	}
	router.Run(":" + port) // เริ่มต้นเซิร์ฟเวอร์ที่พอร์ตที่กำหนด
}
