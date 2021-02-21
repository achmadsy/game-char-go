package models

type Character struct {
	ID            uint    `json:"id" gorm:"primary_key"`
	Name          string  `json:"name"`
	CharacterCode int     `json:"char_code"`
	Power         int     `json:"power"`
	Value         float32 `json:"value"`
	UserID        uint    `json:"userId"`
}

type User struct {
	ID    uint        `json:"id" gorm:"primary_key"`
	Chars []Character `json:"chars"`
}
