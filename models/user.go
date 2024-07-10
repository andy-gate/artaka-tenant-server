package models

type User struct {
	Id         		uint			`json:"id"`
	Username		string		`json:"username"`
	Referral_code	string		`json:"referral_code"`
}

type UserResponse struct {
	Success			bool		`json:"success"`
	Data			User		`json:"data"`
}

type QueryUser struct {
	Username    string   `json:"username"`
	Password    string   `json:"Password"`
}

