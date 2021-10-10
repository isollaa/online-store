package entity

type Item struct {
	ID       int64   `json:"id" gorm:"column:id; PRIMARY_KEY"`
	Name     string  `json:"name" validate:"required" gorm:"column:name"`
	Quantity float64 `json:"quantity" validate:"required,numeric" gorm:"column:quantity"`
}

func (Item) TableName() string {
	return "item"
}
