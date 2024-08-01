package usecase

import (
	"context"

	"github.com/velvetriddles/fullstack-hidden-chapters/internal/domain"
)

type BookRepository interface {
	Create(ctx context.Context, book *domain.Book) error
	GetByID(ctx context.Context, id uint64) (*domain.Book, error)
	Update(ctx context.Context, id uint64, book *domain.Book) error
	Delete(ctx context.Context, id uint64) error
	// List(ctx context.Context, offset, limit int) ([]*domain.Book, error)
}

type bookUseCase struct{
	repo BookRepository
}

func NewBookUseCase(repo BookRepository) *bookUseCase{
	return &bookUseCase{repo: repo}
}
// почему по указателю?

func (buc *bookUseCase) GetBook(ctx context.Context, id uint64) (*domain.Book, error){
	return buc.repo.GetByID(ctx, id)
}

func (buc *bookUseCase) CreateBook(ctx context.Context, book *domain.Book) error{
	return buc.repo.Create(ctx, book)
}

func (buc *bookUseCase) UpdateBook(ctx context.Context,id uint64, book *domain.Book) error{
	return buc.repo.Update(ctx, id, book)
} 

func (buc *bookUseCase) DeleteBook(ctx context.Context, id uint64) error {
	return buc.repo.Delete(ctx, id)
}

