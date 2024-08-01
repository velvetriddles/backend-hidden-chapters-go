CREATE TABLE bookmarks (
    user_id BIGINT REFERENCES users(id),
    book_id BIGINT REFERENCES books(id),
    PRIMARY KEY (user_id, book_id)
);