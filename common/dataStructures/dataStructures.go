package dataStructures

import (
	"time"
)

type Match struct {
	Id        int       `json:"matchid" gorm:"primaryKey"`
	LikerId   int       `json:"likerId"`
	LikedId   int       `json:"likedId"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdatedTime"`
}

type Search struct {
	Id        int       `json:"searchid" gorm:"primaryKey"`
	Name      string    `json:"name"`
	Skill     int       `json:"skill"`
	Level     string    `json:"level"`
	Gender    int       `json:"gender"`
	Radius    int       `json:"radius"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdatedTime"`
	CreatedBy int       `json:"created_by"`
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

type Dislike struct {
	DislikerId int       `json:"dislikerId"`
	DislikedId int       `json:"dislikedId"`
	CreatedAt  time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"autoUpdatedTime"`
}

type User struct {
	ID              int       `json:"id" gorm:"primaryKey"`
	CreatedAt       time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt       time.Time `json:"updated_at" gorm:"autoUpdatedTime"`
	First_name      string    `json:"firstName"`
	Name            string    `json:"name"`
	Gender          int       `json:"gender"`
	Username        string    `json:"username"`
	Email           string    `json:"email"`
	Street          string    `json:"street"`
	HouseNumber     string    `json:"houseNumber"`
	TelephoneNumber string    `json:"telephoneNumber"`
	ProfilPicture   []byte    `json:"profilePicture"`
	Confirmed       bool      `json:"confirmed"`
	Active          bool      `json:"active"`
	Password        string    `json:"password"`
	SearchedSkills  []*Skill  `json:"searchedSkills" gorm:"many2many:user_searchedSkills"`
	AchievedSkills  []*Skill  `json:"achievedSkills" gorm:"many2many:user_achievedSkills"`
	CityIdentifier  int
	City            *City `json:"city" gorm:"foreignKey:CityIdentifier"`
}

type Skill struct {
	ID             uint      `json:"id" gorm:"primaryKey"`
	CreatedAt      time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt      time.Time `json:"updated_at" gorm:"autoUpdatedTime"`
	Name           string    `json:"name"`
	Level          string    `json:"level"`
	UsersSearching []*User   `json:"usersSearching" gorm:"many2many:user_searchedSkills"`
	UsersAchieved  []*User   `json:"usersAchieved" gorm:"many2many:user_achievedSkills"`
}

type City struct {
	PLZ       uint      `json:"plz" gorm:"primaryKey"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdatedTime"`
	Place     string    `json:"place"`
}
