package entity

type User struct {
	ID       int64  `json:"id" gorm:"column:id; PRIMARY_KEY"`
	Username string `json:"username" validate:"required" gorm:"column:username"`
	Password string `json:"password" validate:"required" gorm:"column:password"`
}

func (User) TableName() string {
	return "user"
}
