package entity

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func DB() *gorm.DB {

	return db

}

func SetupDatabase() {

	database, err := gorm.Open(sqlite.Open("sapro.db"), &gorm.Config{})

	if err != nil {

		panic("failed to connect database")

	}

	// Migrate the schema

	database.AutoMigrate(
		&CheckIn{}, &Customer{}, &Employee{}, &RoomPayment{}, &Room{},
	)

	db = database
	password, err := bcrypt.GenerateFromPassword([]byte("123456"), 14)

	db.Model(&Employee{}).Create(&Employee{
		Name:     "นางสาวพร มณีวรรณ",
		Email:    "chanwit@gmail.com",
		Tel:      "0883322456",
		Password: string(password),
	})
	db.Model(&Employee{}).Create(&Employee{
		Name:     "นายสม จันทร์เพ็ญ",
		Email:    "som@gmail.com",
		Tel:      "0885548900",
		Password: string(password),
	})

	var porn Employee
	var som Employee
	db.Raw("SELECT * FROM employees WHERE email = ?", "chanwit@gmail.com").Scan(&porn)
	db.Raw("SELECT * FROM employees WHERE email = ?", "som@example.com").Scan(&som)

	// --- RoomPayment Data
	RP800 := RoomPayment{
		Datepaid: time.Now(),
		Amount:   800,
	}
	db.Model(&RoomPayment{}).Create(&RP800)

	RP1200 := RoomPayment{
		Datepaid: time.Now(),
		Amount:   1200,
	}
	db.Model(&RoomPayment{}).Create(&RP1200)

	// Customer Data
	Cmonkey := Customer{
		Name:  "นางมังกี้ ลูฟี้",
		Email: "monkey@gmail.com",
		Tel:   "0983322403",
	}
	db.Model(&Customer{}).Create(&Cmonkey)

	Cfish := Customer{
		Name:  "นางฟิซ เมอร์เมด",
		Email: "fish@gmail.com",
		Tel:   "0673322403",
	}
	db.Model(&Customer{}).Create(&Cfish)

	//Room Data
	R10012 := Room{
		Location:   "12.3433-08.4450",
		Roomnumber: "10012",
	}
	db.Model(&Room{}).Create(&R10012)

	R10023 := Room{
		Location:   "16.3433-08.4450",
		Roomnumber: "10023",
	}
	db.Model(&Room{}).Create(&R10023)

	//
	// === Query
	//

	// var target Employee
	// db.Model(&Employee{}).Find(&target, db.Where("email = ?", "chanwit@gmail.com"))

	// var employeeLogin Employee
	// db.Model(&Employee{}).Find(&employeeLogin, db.Where("email = ? and employee_id = ?", "พนักงานที่เข้าสู่ระบบ", target.ID))

	/*var watchedList []*WatchVideo
	db.Model(&WatchVideo{}).
		Joins("Playlist").
		Joins("Resolution").
		Joins("Video").
		Find(&watchedList, db.Where("playlist_id = ?", watchedPlaylist.ID))

	for _, wl := range watchedList {
		fmt.Printf("Watch Video: %v\n", wl.ID)
		fmt.Printf("%v\n", wl.Playlist.Title)
		fmt.Printf("%v\n", wl.Resolution.Value)
		fmt.Printf("%v\n", wl.Video.Name)
		fmt.Println("====")
	}*/

}
