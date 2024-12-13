package repositories

import (
	"github.com/kamva/mgm/v3"
	"github.com/ZhanbolEP/my-go-project/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log/slog"
)

type BookRepository interface {
	CreateBook(book *models.Book) error
	GetBookById(id string) (*models.Book, error)
	GetAllBooks() ([]models.Book, error)
	UpdateBook(book *models.Book) error
	DeleteBook(id string) error
	GetHomeBooks() ([]models.Book, []models.Book, error)
}

type bookRepository struct{}

func NewBookRepository() BookRepository {
	return &bookRepository{}
}

// CreateBook adds a new book to the database
func (r *bookRepository) CreateBook(book *models.Book) error {
	return mgm.Coll(book).Create(book)
}

// GetBookById fetches a book by its ID
func (r *bookRepository) GetBookById(id string) (*models.Book, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	book := &models.Book{}
	err = mgm.Coll(book).FindByID(objID, book)
	return book, err
}

// GetAllBooks fetches all books from the database
func (r *bookRepository) GetAllBooks() ([]models.Book, error) {
	var books []models.Book
	err := mgm.Coll(&models.Book{}).SimpleFind(&books, bson.M{})
	return books, err
}

// UpdateBook updates an existing book
func (r *bookRepository) UpdateBook(book *models.Book) error {
	return mgm.Coll(book).Update(book)
}

func (r *bookRepository) DeleteBook(id string) error {
	// Convert the string ID to a MongoDB ObjectID
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	// Use DeleteOne to remove the book with the specified ID
	_, err = mgm.Coll(&models.Book{}).DeleteOne(mgm.Ctx(), bson.M{"_id": objID})
	return err
}

// GetHomeBooks fetches trending and recommended books
func (r *bookRepository) GetHomeBooks() ([]models.Book, []models.Book, error) {
	var topSellerBooks []models.Book
	var recommendedBooks []models.Book

	// Get trending books
	err := mgm.Coll(&models.Book{}).SimpleFind(&topSellerBooks, bson.M{"trending": true})
	if err != nil {
		slog.Error("Error getting top seller books", "error", err.Error())
		return nil, nil, err
	}

	if len(topSellerBooks) > 0 {
		// Get recommended books except for top sellers
		topSellerIDs := []primitive.ObjectID{}
		for _, book := range topSellerBooks {
			topSellerIDs = append(topSellerIDs, book.ID)
		}
		err = mgm.Coll(&models.Book{}).SimpleFind(&recommendedBooks, bson.M{
			"trending": false,
			"_id":      bson.M{"$nin": topSellerIDs},
		})
		if err != nil {
			slog.Error("Error getting recommended books", "error", err.Error())
		}
	} else {
		// Get all non-trending books
		err = mgm.Coll(&models.Book{}).SimpleFind(&recommendedBooks, bson.M{"trending": false})
		if err != nil {
			slog.Error("Error getting recommended books", "error", err.Error())
		}
	}

	return topSellerBooks, recommendedBooks, err
}
