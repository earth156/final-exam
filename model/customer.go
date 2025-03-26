package model

import (
	"time"
)

type Customer struct {
	CustomerID  int       `gorm:"column:customer_id;AUTO_INCREMENT;primary_key" json:"customer_id"`
	FirstName   string    `gorm:"column:first_name;NOT NULL" json:"first_name"`
	LastName    string    `gorm:"column:last_name;NOT NULL" json:"last_name"`
	Email       string    `gorm:"column:email;unique;NOT NULL" json:"email"`
	PhoneNumber string    `gorm:"column:phone_number" json:"phone_number"`
	Address     string    `gorm:"column:address" json:"address"`
	Password    string    `gorm:"column:password;NOT NULL" json:"-"` // ไม่ส่งรหัสผ่านใน JSON
	CreatedAt   time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at;default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (m *Customer) TableName() string {
	return "customer"
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UpdateAddressRequest struct {
	Address string `json:"address" binding:"required"`
}

type LoginResponse struct {
	Customer Customer `json:"customer"`
	Token    string   `json:"token"`
}
