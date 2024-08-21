package models

type BookEntity struct {
	ID     int    `gorm:"primaryKey;autoIncrement" json:"id"`
	Title  string `gorm:"not null" json:"title" validate:"required"`
	Author string `gorm:"not null" json:"author" validate:"required"`
	Genre  string `gorm:"not null" json:"genre" validate:"required"`
	Read   bool   `gorm:"default:false" json:"read"`
}

func (BookEntity) TableName() string {
	return "books"
}
