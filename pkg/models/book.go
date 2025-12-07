package models


type Book struct {
    ID     uint   `json:"id" gorm:"primaryKey"`
    Title  string `json:"title"`
    Author string `json:"author"`
	Publication string `json:"publication"`
}
