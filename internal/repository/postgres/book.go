package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"

	_ "github.com/lib/pq"
	"github.com/velvetriddles/fullstack-hidden-chapters/internal/domain"
)

type bookRepository struct {
	db *sql.DB
}

func NewBookRepository(db *sql.DB) *bookRepository {
	return &bookRepository{db: db}
}

func (r *bookRepository) Create(ctx context.Context, book *domain.Book) error {
	query := `
        INSERT INTO books (
            language_id, 
            genre_id,
            price,
            name,
            author,
            description,
            pages,
            quantity,
            year_of_publishing
        )
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
        RETURNING id, created_at, updated_at`

	err := r.db.QueryRowContext(ctx, query,
		book.LanguageID,
		book.GenreID,
		book.Price,
		book.Name,
		book.Author,
		book.Description,
		book.Pages,
		book.Quantity,
		book.YearOfPublishing).Scan(&book.ID, &book.CreatedAt, &book.UpdatedAt)

	if err != nil {
		return fmt.Errorf("failed to create book: %w", err)
	}

	return nil
}
func (r *bookRepository) GetByID(ctx context.Context, id uint64) (*domain.Book, error) {
	query := `
        SELECT id, language_id, genre_id, price, name, author, description, 
               pages, quantity, year_of_publishing, created_at, updated_at
        FROM books
        WHERE id = $1
    `

	var book domain.Book
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&book.ID,
		&book.LanguageID,
		&book.GenreID,
		&book.Price,
		&book.Name,
		&book.Author,
		&book.Description,
		&book.Pages,
		&book.Quantity,
		&book.YearOfPublishing,
		&book.CreatedAt,
		&book.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("book not found")
		}
		return nil, err
	}

	return &book, nil
}
func (r *bookRepository) Update(ctx context.Context, id uint64, book *domain.Book) error {
	var args []interface{}
	var fieldsForUpdate []string
	argID := 1
	if book.LanguageID != nil {
		fieldsForUpdate = append(fieldsForUpdate, fmt.Sprintf("language_id = $%d", argID))
		args = append(args, *book.LanguageID)
		argID++
	}
	if book.GenreID != nil {
		fieldsForUpdate = append(fieldsForUpdate, fmt.Sprintf("genre_id = $%d", argID))
		args = append(args, *book.GenreID)
		argID++
	}
	if book.Price != nil {
		fieldsForUpdate = append(fieldsForUpdate, fmt.Sprintf("price = $%d", argID))
		args = append(args, *book.Price)
		argID++
	}
	if book.Name != nil {
		fieldsForUpdate = append(fieldsForUpdate, fmt.Sprintf("name = $%d", argID))
		args = append(args, *book.Name)
		argID++
	}
	if book.Author != nil {
		fieldsForUpdate = append(fieldsForUpdate, fmt.Sprintf("author = $%d", argID))
		args = append(args, *book.Author)
		argID++
	}
	if book.Description != nil {
		fieldsForUpdate = append(fieldsForUpdate, fmt.Sprintf("description = $%d", argID))
		args = append(args, *book.Description)
		argID++
	}
	if book.Pages != nil {
		fieldsForUpdate = append(fieldsForUpdate, fmt.Sprintf("pages = $%d", argID))
		args = append(args, *book.Pages)
		argID++
	}
	if book.Quantity != nil {
		fieldsForUpdate = append(fieldsForUpdate, fmt.Sprintf("quantity = $%d", argID))
		args = append(args, *book.Quantity)
		argID++
	}
	if book.YearOfPublishing != nil {
		fieldsForUpdate = append(fieldsForUpdate, fmt.Sprintf("year_of_publishing = $%d", argID))
		args = append(args, *book.YearOfPublishing)
		argID++
	}

	fieldsForUpdate = append(fieldsForUpdate, fmt.Sprintf("updated_at = $%d", argID))
	args = append(args, time.Now())
	argID++

	if len(fieldsForUpdate) == 0 {
		return errors.New("no fields to update")
	}

	query := fmt.Sprintf("UPDATE books SET %s WHERE id = $%d", strings.Join(fieldsForUpdate, ","), argID)
	args = append(args, id)
	result, err := r.db.ExecContext(ctx, query, args...)
	if err != nil {
		return errors.New("failed to update book")
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}
	if rows == 0 {
		return errors.New("books not found or no changes")
	}
	return nil

}
func (r *bookRepository) Delete(ctx context.Context, id uint64) error {
	query := `DELETE FROM books WHERE id = $1`
	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to update book %w", err)
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rows == 0 {
		return fmt.Errorf("book not found")
	}
	return nil
}
