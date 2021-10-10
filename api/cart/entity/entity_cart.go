package entity

type Cart struct {
	ID       int64   `json:"id" gorm:"column:id; PRIMARY_KEY"`
	UserID   int64   `json:"user_id" gorm:"column:user_id"`
	ItemID   int64   `json:"item_id" validate:"required,numeric" gorm:"column:item_id"`
	ItemName string  `json:"item_name" gorm:"column:item_name"`
	Quantity float64 `json:"quantity" validate:"required,numeric" gorm:"column:quantity"`
}

func (Cart) TableName() string {
	return "cart"
}
