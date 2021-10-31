package entity

import (
	"time"

	"gorm.io/gorm"
)

type CheckIn struct {
	gorm.Model
	DateTime time.Time

	CustomerID *uint
	Customer   Customer `gorm:"references:id"`

	ReserveID *uint
	Reserve   Room `gorm:"references:id"`

	PaymentID *uint       `gorm:"uniqueIndex"`
	Payment   RoomPayment `gorm:"references:id"`

	// EmployeeID ทำหน้าที่เป็น FK
	EmployeeID *uint
	// เป็นข้อมูล employee เมื่อ join ตาราง
	Employee Employee `gorm:"references:id"`
}
type Customer struct {
	gorm.Model
	Name  string
	Email string `gorm:"uniqueIndex"`
	Tel   string

	Record []CheckIn `gorm:"foreignKey:CustomerID"`
}
type Employee struct {
	gorm.Model
	Name     string
	Tel      string
	Email    string `gorm:"uniqueIndex"`
	Password string
	Records  []CheckIn `gorm:"foreignKey:EmployeeID"`
}
type RoomPayment struct {
	gorm.Model
	Datepaid time.Time
	Amount   int
	//Record CheckIn `gorm:"references:id"`
	//RecordID *uint
	//Record []CheckIn `gorm:"foreignKey:PaymentID"`
}
type Room struct {
	gorm.Model
	Location   string
	Roomnumber string

	Records []CheckIn `gorm:"foreignKey:ReserveID"`
}
