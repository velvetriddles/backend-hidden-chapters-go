CREATE TABLE books (
    id BIGSERIAL PRIMARY KEY,
    language_id BIGINT,
    genre_id BIGINT,
    price BIGINT,
    name VARCHAR(255),
    author VARCHAR(255),
    description TEXT,
    pages BIGINT,
    quantity BIGINT,
    year_of_publishing BIGINT,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);