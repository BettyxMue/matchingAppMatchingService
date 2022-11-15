package dataStructures

type User struct {
	Id          int    `json:"id"`
	City        string `json:"city"`
	Email       string `json:"email"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"name"`
	Password    string `json:"password"`
	Street      string `json:"street"`
	HouseNumber string `json:"houseNumber"`
	Username    string `json:"username"`
}

type Match struct {
	Id      string `json:"id"`
	UserId1 int    `json:"userid1"`
	UserId2 int    `json:"userid2"`
}
