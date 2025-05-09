-- +goose Up
alter table users add constraint unique_user_name unique (name);

-- +goose Down
alter table users drop constraint unique_user_name;