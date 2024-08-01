CREATE TABLE cart_items (
    user_id BIGINT REFERENCES users(id),
    book_id BIGINT REFERENCES books(id),
    quantity BIGINT,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    PRIMARY KEY (user_id, book_id)
);