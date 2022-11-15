package dataStructures

type User struct {
	Id          int    `json:"id"`
	City        string `json:"city"`
	Email       string `json:"email"`
	First_name  string `json:"first_name"`
	Name        string `json:"name"`
	Password    string `json:"password"`
	Street      string `json:"street"`
	HouseNumber string `json:"houseNumber"`
	Username    string `json:"username"`
}

type Match struct {
	Id    string `json:"id"`
	User1 User   `json:"user1"`
	User2 User   `json:"user2"`
}
