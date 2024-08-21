package service

import (
	"gobooks/internal/models"
	"gobooks/internal/repository"
)

type BookService struct {
	repo *repository.BookRepository
}

func (s *BookService) CreateBook(book *models.BookEntity) error {
	return s.repo.CreateBook(book)
}

func (s *BookService) GeAll() ([]models.BookEntity, error) {
	return s.repo.GetAllBooks()
}

func (s *BookService) GetById(bookId int) (*models.BookEntity, error) {
	return s.repo.GetBookByID(bookId)
}

func (s *BookService) UpdateBook(id int, updatedBook *models.BookEntity) (*models.BookEntity, error) {
	return s.repo.UpdateBook(id, updatedBook)
}

// UpdateBookReadStatus atualiza o status de leitura de um livro pra lido
func (s *BookService) MarkBookAsRead(id int) (*models.BookEntity, error) {
	return s.repo.MarkBookAsRead(id)
}

func NewBookService(repo *repository.BookRepository) *BookService {
	return &BookService{repo: repo}
}
