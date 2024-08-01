CREATE TABLE child_genres (
    id BIGSERIAL PRIMARY KEY,
    name_genre_child VARCHAR(100),
    id_genre_parent BIGINT REFERENCES parent_genres(id)
);