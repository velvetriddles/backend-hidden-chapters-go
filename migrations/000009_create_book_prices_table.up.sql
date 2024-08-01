CREATE TABLE book_prices (
    book_id BIGINT REFERENCES books(id),
    format_id BIGINT,
    price BIGINT,
    PRIMARY KEY (book_id, format_id)
);