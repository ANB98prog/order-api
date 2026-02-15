package entity

type User struct {
	ID    uint    `gorm:"primary_key"`
	Phone string  `gorm:"unique"`
	Order []Order `gorm:"foreignkey:UserID"`
}
