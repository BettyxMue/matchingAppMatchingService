package dataStructures

import (
	"time"
)

type Match struct {
	Id           uint      `json:"matchid" gorm:"primaryKey"`
	SearchId     Search    `json:"match_searchid"`
	UserId1      int       `json:"userid1"`
	UserId2      int       `json:"userid2"`
	ConfirmUser1 bool      `json:"confirm_user1"`
	ConfirmUser2 bool      `json:"confirm_user2"`
	CreatedAt    time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"autoUpdatedTime"`
	DeletedAt    time.Time `json:"deleted_at"`
}

type Search struct {
	Id        uint      `json:"searchid" gorm:"primaryKey"`
	Name	  string	`json:"name"`
	Topic     string    `json:"topic"`
	Level     string    `json:"level"`
	Gender    string    `json:"gender"`
	Radius    int       `json:"radius"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdatedTime"`
}
