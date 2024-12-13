package models

type Address struct {
	City    string `json:"city" gorm:"type:varchar(255);not null"`
	Country string `json:"country" gorm:"type:varchar(255)"`
	State   string `json:"state" gorm:"type:varchar(255)"`
	Zipcode string `json:"zipcode" gorm:"type:varchar(255)"`
}

type Order struct {
	BaseModel
	Name       string  `json:"name" gorm:"type:varchar(255);not null"`
	Email      string  `json:"email" gorm:"type:varchar(255);not null"`
	Address    Address `json:"address" gorm:"embedded"` //embed another table
	Phone      string  `json:"phone" gorm:"type:varchar(20);not null"`
	TotalPrice float64 `json:"total_price" gorm:"column:total_price;not null"` //gorm: column giving name to it
	UserId     uint    `json:"user_id" gorm:"column:user_id;not null"`

	Books []Book `json:"books" gorm:"many2many:order_books;"` // many-to-many relationship - for join table (array in table)
	User  User   `json:"user" gorm:"foreignKey:user_id"`
}

func (o *Order) TableName() string {
	return "orders"
}

type OrderBook struct {
	BaseModel
	OrderID uint `json:"order_id" gorm:"column:order_id;not null"`
	BookID  uint `json:"book_id" gorm:"column:book_id;not null"`
}

func (ob *OrderBook) TableName() string {
	return "order_books"
}
