-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE product_images
(
    id         SERIAL PRIMARY KEY,
    product_id INTEGER      NOT NULL,
    file_path  VARCHAR(255) NOT NULL,
    FOREIGN KEY (product_id) REFERENCES products (id) ON DELETE CASCADE
);

CREATE TABLE product_warehouse
(
    id           SERIAL PRIMARY KEY,
    product_id   INTEGER NOT NULL,
    warehouse_id INTEGER NOT NULL,
    FOREIGN KEY (product_id) REFERENCES products (id) ON DELETE CASCADE,
    FOREIGN KEY (warehouse_id) REFERENCES warehouses (id) ON DELETE CASCADE
);

CREATE TABLE category_product
(
    id          SERIAL PRIMARY KEY,
    product_id  INTEGER NOT NULL,
    category_id INTEGER NOT NULL,
    FOREIGN KEY (product_id) REFERENCES products (id) ON DELETE CASCADE,
    FOREIGN KEY (category_id) REFERENCES categories (id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE category_product;
DROP TABLE product_warehouse;
DROP TABLE product_images;
-- +goose StatementEnd
