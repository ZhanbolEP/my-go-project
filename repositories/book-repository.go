package repositories

import (
	"github.com/ZhanbolEP/my-go-project/models"
	"gorm.io/gorm"
	"log/slog"
)

type BookRepository interface {
	CreateBook(book *models.Book) error
	GetBookById(id uint) (*models.Book, error)
	GetAllBooks() ([]models.Book, error)
	UpdateBook(book *models.Book) error
	DeleteBook(id uint) error
	GetHomeBooks() ([]models.Book, []models.Book, error)
}

type bookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) BookRepository {
	return &bookRepository{db}
}

func (r *bookRepository) CreateBook(book *models.Book) error {
	return r.db.Create(book).Error
}

func (r *bookRepository) GetBookById(id uint) (*models.Book, error) {
	var book models.Book
	err := r.db.First(&book, id).Error
	return &book, err
}

func (r *bookRepository) GetAllBooks() ([]models.Book, error) {
	var books []models.Book
	err := r.db.Find(&books).Error
	return books, err
}

func (r *bookRepository) UpdateBook(book *models.Book) error {
	return r.db.Save(book).Error
}

func (r *bookRepository) DeleteBook(id uint) error {
	return r.db.Delete(&models.Book{}, id).Error
}

func (r *bookRepository) GetHomeBooks() ([]models.Book, []models.Book, error) {
	var topSellerBooks []models.Book
	var recommendedBooks []models.Book
	err := r.db.Where("trending = ?", true).Find(&topSellerBooks).Error
	if err != nil {
		slog.Error("Error getting top seller books", "error", err.Error())
	}
	if len(topSellerBooks) > 0 {
		// Get recommended books except top seller books
		ids := make([]uint, len(topSellerBooks))
		for i, book := range topSellerBooks {
			ids[i] = book.ID
		}
		err = r.db.Where("trending = ? AND id NOT IN (?)", false, ids).Find(&recommendedBooks).Error
		return topSellerBooks, recommendedBooks, err
	}
	err = r.db.Where("trending = ?", false).Find(&recommendedBooks).Error
	if err != nil {
		slog.Error("Error getting recommended books", "error", err.Error())
	}
	return topSellerBooks, recommendedBooks, err
}
