package repositories

import (
	"github.com/kamva/mgm/v3"
	"github.com/ZhanbolEP/my-go-project/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OrderRepository interface {
	CreateOrder(order *models.Order) (primitive.ObjectID, error)
	GetOrderById(id primitive.ObjectID) (*models.Order, error)
	GetAllOrders() ([]models.Order, error)
	UpdateOrder(order *models.Order) error
	DeleteOrder(id primitive.ObjectID) error
	CreateOrderBook(orderId primitive.ObjectID, bookId primitive.ObjectID) error
	GetOrdersForUser(userId primitive.ObjectID) ([]models.Order, error)
}

type orderRepository struct{}

func NewOrderRepository() OrderRepository {
	return &orderRepository{}
}

func (r *orderRepository) CreateOrder(order *models.Order) (primitive.ObjectID, error) {
	err := mgm.Coll(order).Create(order) // Create order document in MongoDB
	if err != nil {
		return primitive.NilObjectID, err
	}
	return order.ID, nil
}

func (r *orderRepository) GetOrderById(id primitive.ObjectID) (*models.Order, error) {
	var order models.Order
	err := mgm.Coll(&order).FindByID(id, &order) // Find order by ID
	return &order, err
}

func (r *orderRepository) GetAllOrders() ([]models.Order, error) {
	var orders []models.Order
	err := mgm.Coll(&models.Order{}).SimpleFind(&orders, bson.M{}) // Find all orders
	return orders, err
}

func (r *orderRepository) UpdateOrder(order *models.Order) error {
	return mgm.Coll(order).Update(order) // Update the order document
}

func (r *orderRepository) DeleteOrder(id primitive.ObjectID) error {
	return mgm.Coll(&models.Order{}).DeleteById(id) // Delete order by ID
}

func (r *orderRepository) CreateOrderBook(orderId primitive.ObjectID, bookId primitive.ObjectID) error {
	// Create an order-book relationship (in MongoDB, you may just keep references as you did in the Order model)
	orderBook := models.OrderBook{
		OrderID: orderId,
		BookID:  bookId,
	}
	return mgm.Coll(&models.OrderBook{}).Create(&orderBook)
}

func (r *orderRepository) GetOrdersForUser(userId primitive.ObjectID) ([]models.Order, error) {
	var orders []models.Order
	err := mgm.Coll(&models.Order{}).SimpleFind(&orders, bson.M{"user_id": userId}) // Find orders for a user
	return orders, err
}
