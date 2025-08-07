-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
ALTER TABLE product_warehouse
    ADD COLUMN amount integer NOT NULL DEFAULT 0;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
ALTER TABLE product_warehouse
    DROP COLUMN amount;
-- +goose StatementEnd
