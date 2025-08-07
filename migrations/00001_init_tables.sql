-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE sellers
(
    id          SERIAL PRIMARY KEY,
    name        VARCHAR(100) NOT NULL,
    description TEXT
);

CREATE TABLE categories
(
    id   SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL
);

CREATE TABLE products
(
    id          SERIAL PRIMARY KEY,
    name        VARCHAR(50) NOT NULL,
    description TEXT,
    price       DECIMAL     NOT NULL DEFAULT 0,
    seller_id   INTEGER     NOT NULL,
    FOREIGN KEY (seller_id) REFERENCES sellers (id) ON DELETE CASCADE
);

CREATE TABLE warehouses
(
    id           SERIAL PRIMARY KEY,
    address      VARCHAR(100) NOT NULL,
    phone_number VARCHAR(12)  NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE warehouses;
DROP TABLE products;
DROP TABLE categories;
DROP TABLE sellers;
-- +goose StatementEnd
