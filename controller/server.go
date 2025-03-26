package controller

import "github.com/gin-gonic/gin"

// เริ่มต้นเซิร์ฟเวอร์
func StartServer() {
	// ตั้งค่าให้เซิร์ฟเวอร์ทำงานในโหมด Release
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	// API ทดสอบเพื่อให้มั่นใจว่า API ทำงานได้
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "API is now working",
		})
	})

	// เรียกใช้ DemoController หรือ Login Controller
	Cus_login(router)

	// เริ่มเซิร์ฟเวอร์
	router.Run(":8081")
}
