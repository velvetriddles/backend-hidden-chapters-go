package domain

import "time"

type Book struct {
	ID               uint64
	LanguageID       *uint64
	GenreID          *uint64
	Price            *int64
	Name             *string
	Author           *string
	Description      *string
	Pages            *uint64
	Quantity         *uint64
	YearOfPublishing *uint64
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

type BookPrice struct {
	BookID   uint64
	FormatID *uint64
	Price    *int64
}

type ParentGenre struct {
	ID              uint64
	NameGenreParent *string
}

type ChildGenre struct {
	ID             uint64
	NameGenreChild *string
	IDGenreParent  *uint64
}

type Format struct {
	ID   uint64
	Name *string
}

type Rating struct {
	UserID    uint64
	BookID    uint64
	Grade     *uint8
	GradeDate time.Time
}
