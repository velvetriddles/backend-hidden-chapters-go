CREATE TABLE meta_users (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT REFERENCES users(id),
    balance BIGINT,
    country VARCHAR(100),
    city VARCHAR(100),
    phone VARCHAR(50),
    first_name VARCHAR(100),
    last_name VARCHAR(100),
    address TEXT
);