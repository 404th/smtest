package model

type User struct {
	Id        uint   `json:"id" gorm:"primaryKey"`
	Username  string `json:"username" gorm:"type:varchar(255);unique;not null"`
	Password  string `json:"password" gorm:"type:varchar(255);not null"`
	CreatedAt string `json:"createdAt" gorm:"type:varchar(255);default:CURRENT_TIMESTAMP()"`
}

type RegisterRequest struct {
	Username string `json:"username" gorm:"type:varchar(255);unique;not null"`
	Password string `json:"password" gorm:"type:varchar(255);not null"`
}

type RegisterResponse struct {
	Id uint `json:"id" gorm:"primaryKey"`
}

type LoginRequest struct {
	Username string `json:"username" gorm:"type:varchar(255);unique;not null"`
	Password string `json:"password" gorm:"type:varchar(255);not null"`
}

type LoginResponse struct {
	AccessToken string `json:"accessToken"`
}

type GetUserRequest struct {
	Id uint `json:"id" gorm:"primaryKey"`
}

func (User) TableName() string {
	return "users"
}
