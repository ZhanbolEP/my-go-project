package services

import (
	"github.com/ZhanbolEP/my-go-project/models"
	"github.com/ZhanbolEP/my-go-project/repositories"
)

type OrderService struct {
	orderRepo repositories.OrderRepository
}

func NewOrderService(repo repositories.OrderRepository) *OrderService {
	return &OrderService{orderRepo: repo}
}

func (s *OrderService) CreateOrder(order *models.Order) (uint, error) {
	return s.orderRepo.CreateOrder(order)
}

func (s *OrderService) GetOrderById(id uint) (*models.Order, error) {
	return s.orderRepo.GetOrderById(id)
}

func (s *OrderService) GetAllOrders() ([]models.Order, error) {
	return s.orderRepo.GetAllOrders()
}

func (s *OrderService) UpdateOrder(order *models.Order) error {
	return s.orderRepo.UpdateOrder(order)
}

func (s *OrderService) DeleteOrder(id uint) error {
	return s.orderRepo.DeleteOrder(id)
}

func (s *OrderService) CreateOrderBook(orderId uint, bookId uint) error {
	return s.orderRepo.CreateOrderBook(orderId, bookId)
}

func(s *OrderService) GetOrdersForUser(userId uint) ([]models.Order, error) {
	return s.orderRepo.GetOrdersForUser(userId)
}