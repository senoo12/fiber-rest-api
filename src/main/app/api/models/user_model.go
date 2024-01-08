package models

type User struct{
	ID uint `gorm:"primaryKey" json:"id"`
	Name string `gorm:"type:varchar(300)" json:"name"`
	Email string `gorm:"uniqueIndex" json:"email"`
	Password string `gorm:"type:varchar(300)" json:"password"`
}