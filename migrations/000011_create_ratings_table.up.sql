CREATE TABLE ratings (
    user_id BIGINT REFERENCES users(id),
    book_id BIGINT REFERENCES books(id),
    grade SMALLINT,
    grade_date TIMESTAMP NOT NULL DEFAULT NOW(),
    PRIMARY KEY (user_id, book_id)
);