-- +goose Up
-- +goose StatementBegin
ALTER TABLE posts
ADD COLUMN complete BOOLEAN DEFAULT 0;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

-- +goose StatementEnd
