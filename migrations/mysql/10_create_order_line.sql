-- +migrate Up
CREATE TABLE order_line (
    id BIGINT NOT NULL PRIMARY KEY,
    order_id BIGINT NOT NULL,
    product_id BIGINT,
    variant_id BIGINT,
    sku VARCHAR(255),
    name VARCHAR(255),
    title VARCHAR(255),
    variant_title VARCHAR(255),
    vendor VARCHAR(255),
    quantity INT NOT NULL,
    price DECIMAL(13, 4) NOT NULL,
    total_discount DECIMAL(13, 4) NOT NULL,
    FOREIGN KEY (order_id) REFERENCES order(id) ON DELETE CASCADE,
    FOREIGN KEY (product_id) REFERENCES product(id) ON DELETE SET NULL,
    FOREIGN KEY (variant_id) REFERENCES product_variant(id) ON DELETE SET NULL
);
-- +migrate Down
DROP TABLE order_line;