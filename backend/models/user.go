package models

type User struct {
	ID       uint   `json:"id" gorm:"primary_key;autoIncrement"`
	Username string `json:"username" gorm:"unique;not null"`
	Password string `json:"password" gorm:"not null"`
}

// type User struct {
//     ID       uint   `json:"id" gorm:"primary_key;autoIncrement"`
//     Username string `json:"username" gorm:"not null"`
//     Password string `json:"password" gorm:"not null"`
// }
