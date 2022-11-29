package dataStructures

import (
	"time"
)

type Match struct {
	Id        uint      `json:"matchid" gorm:"primaryKey"`
	LikerId   int       `json:"likerId"`
	LikedId   int       `json:"likedId"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdatedTime"`
}

type Search struct {
	Id        uint      `json:"searchid" gorm:"primaryKey"`
	Name      string    `json:"name"`
	Topic     string    `json:"topic"`
	Level     string    `json:"level"`
	Gender    string    `json:"gender"`
	Radius    int       `json:"radius"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdatedTime"`
}

type Like struct {
	LikerId   int       `json:"likerId"`
	LikedId   int       `json:"likedId"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdatedTime"`
}

type UserLike struct {
	UserId int   `json:"userid"`
	Liked  []int `json:"liked"`
}
