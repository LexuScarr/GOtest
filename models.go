package main

type News struct {
	ID         uint       `gorm:"primaryKey"`
	Title      string     `gorm:"type:text;not null"`
	Content    string     `gorm:"type:text;not null"`
	Categories []Category `gorm:"many2many:news_categories;"`
}

type Category struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"type:text;not null"`
}

type NewsCategory struct {
	NewsID     uint
	CategoryID uint
}
