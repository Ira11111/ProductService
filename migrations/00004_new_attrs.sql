-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
ALTER TABLE products
    ADD COLUMN order_count INTEGER NOT NULL DEFAULT 0;
ALTER TABLE sellers
    ADD COLUMN user_id INTEGER NOT NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
ALTER TABLE products
    DROP COLUMN order_count
ALTER TABLE sellers
    DROP COLUMN user_id;
-- +goose StatementEnd
