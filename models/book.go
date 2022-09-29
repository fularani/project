package models

type Book struct {
    Id     int    `json:"id" gorm:"primaryKey,not null,unique"`
    Title  string `json:"title"`
    Author string `json:"author"`
    Desc   string `json:"desc"`
}

func (e *Book) TableName() string {
	return "Book"
}