-- +goose Up
CREATE TABLE posts (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    title TEXT,
    url TEXT UNIQUE,
    description TEXT,
    published_at TIMESTAMP NOT NULL,
    feed_id UUID REFERENCES feeds(id)
);

-- +goose Down
DROP TABLE posts;