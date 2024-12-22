package services

import (
	"github.com/ZhanbolEP/my-go-project/models"
	"github.com/ZhanbolEP/my-go-project/repositories"
)

type BookService struct {
	bookRepo repositories.BookRepository
}

func NewBookService(repo repositories.BookRepository) *BookService {
	return &BookService{bookRepo: repo}
}

func (s *BookService) CreateBook(book *models.Book) error {
	return s.bookRepo.CreateBook(book)
}

func (s *BookService) GetBookById(id uint) (*models.Book, error) {
	return s.bookRepo.GetBookById(id)
}

func (s *BookService) GetAllBooks() ([]models.Book, error) {
	return s.bookRepo.GetAllBooks()
}

func(s *BookService) UpdateBook(book *models.Book) error {
	return s.bookRepo.UpdateBook(book) 
}

func (s *BookService) DeleteBook(id uint) error {
	return s.bookRepo.DeleteBook(id)
}

func (s *BookService) GetHomeBooks() ([]models.Book, []models.Book, error) {
	return s.bookRepo.GetHomeBooks()
}
