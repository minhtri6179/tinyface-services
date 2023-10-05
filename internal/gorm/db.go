package db

type User struct {
	ID           uint   `gorm:"primary_key"`
	Username     string `gorm:"unique_index"`
	HashPassword string `gorm:"not null"`
	Salt         string `gorm:"not null"`
	FirstName    string `gorm:"not null"`
	LastName     string `gorm:"not null"`
	DataOfBirth  string `gorm:"not null"`
	Email        string `gorm:"unique_index"`
}

func (User) TableName() string {
	return "users"
}
