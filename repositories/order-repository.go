package repositories

import (
	"github.com/ZhanbolEP/my-go-project/models"
	"gorm.io/gorm"
)

type OrderRepository interface {
	CreateOrder(order *models.Order) (uint, error)
	GetOrderById(id uint) (*models.Order, error)
	GetAllOrders() ([]models.Order, error)
	UpdateOrder(order *models.Order) error
	DeleteOrder(id uint) error
	CreateOrderBook(uint, uint) error
	GetOrdersForUser(uint) ([]models.Order, error)
}

type orderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderRepository{db}
}

func (r *orderRepository) CreateOrder(order *models.Order) (uint, error) {
	if err := r.db.Create(order).Error; err != nil {
		return 0, err
	}
	return order.ID, nil
}

func (r *orderRepository) GetOrderById(id uint) (*models.Order, error) {
	var order models.Order
	err := r.db.Preload("Books").First(&order, id).Error
	return &order, err
}

func (r *orderRepository) GetAllOrders() ([]models.Order, error) {
	var orders []models.Order
	err := r.db.Preload("Books").Find(&orders).Error
	return orders, err
}

func (r *orderRepository) UpdateOrder(order *models.Order) error {
	return r.db.Save(order).Error
}

func (r *orderRepository) DeleteOrder(id uint) error {
	return r.db.Delete(&models.Order{}, id).Error
}

func (r *orderRepository) CreateOrderBook(orderId uint, bookId uint) error {
	orderBook := models.OrderBook{OrderID: orderId, BookID: bookId}
	return r.db.Create(&orderBook).Error
}

func (r *orderRepository) GetOrdersForUser(userId uint) ([]models.Order, error) {
	var orders []models.Order
	err := r.db.Where("user_id = ?", userId).Preload("Books").Find(&orders).Error
	return orders, err
}
