package domain

import "time"

type User struct {
	ID       uint64
	Username *string
	Email    *string
	Password *string
	Role     *string
}

type MetaUser struct {
	ID        uint64
	UserID    *uint64
	Balance   *int64
	Country   *string
	City      *string
	Phone     *string
	FirstName *string
	LastName  *string
	Address   *string
}

type Bookmark struct {
	UserID uint64
	BookID uint64
}

type CartItem struct {
	UserID    uint64
	BookID    uint64
	Quantity  *uint64
	CreatedAt time.Time
}
