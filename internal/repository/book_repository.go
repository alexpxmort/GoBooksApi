package repository

import (
	"gobooks/internal/config"
	"gobooks/internal/models"

	"gorm.io/gorm"
)

type BookRepository struct{}

func (r *BookRepository) CreateBook(book *models.BookEntity) error {
	return config.DB.Create(book).Error
}

func (r *BookRepository) GetAllBooks() ([]models.BookEntity, error) {
	var books []models.BookEntity
	result := config.DB.Find(&books)
	return books, result.Error
}

func (r *BookRepository) GetBookByID(id int) (*models.BookEntity, error) {
	var book models.BookEntity
	result := config.DB.First(&book, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &book, nil
}

func (r *BookRepository) UpdateBook(id int, updatedBook *models.BookEntity) (*models.BookEntity, error) {
	var book models.BookEntity

	// Encontra o livro existente pelo ID
	result := config.DB.First(&book, id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil // Livro não encontrado
		}
		return nil, result.Error
	}

	// Atualiza os campos do livro com os valores fornecidos
	// Note que você pode atualizar apenas os campos que são modificáveis
	book.Title = updatedBook.Title
	book.Author = updatedBook.Author
	book.Genre = updatedBook.Genre

	// Salva as mudanças no banco de dados
	result = config.DB.Save(&book)
	if result.Error != nil {
		return nil, result.Error
	}

	return &book, nil
}

func (r *BookRepository) MarkBookAsRead(id int) (*models.BookEntity, error) {
	var book models.BookEntity

	// Encontra o livro existente pelo ID
	result := config.DB.First(&book, id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil // Livro não encontrado
		}
		return nil, result.Error
	}

	// Atualiza o status de leitura do livro para lido
	book.Read = true

	// Salva as mudanças no banco de dados
	result = config.DB.Save(&book)
	if result.Error != nil {
		return nil, result.Error
	}

	return &book, nil
}
