-- +migrate Up
CREATE TABLE product_variant (
    id BIGINT NOT NULL PRIMARY KEY,
    product_id BIGINT NOT NULL,
    barcode VARCHAR(255),
    compare_at_price DECIMAL(13, 4),
    fulfillment_service VARCHAR(255),
    grams DOUBLE,
    inventory_management VARCHAR(50) NOT NULL,
    inventory_policy VARCHAR(50) NOT NULL,
    inventory_quantity INT NOT NULL,
    position SMALLINT NOT NULL,
    price DECIMAL(13, 4) NOT NULL,
    sku VARCHAR(255),
    taxable TINYINT(1) NOT NULL,
    title VARCHAR(255),
    weight DOUBLE,
    weight_unit VARCHAR(4) NOT NULL,
    created_at DATETIME,
    updated_at DATETIME,
    FOREIGN KEY (product_id) REFERENCES product(id) ON DELETE CASCADE
);
-- +migrate Down
DROP TABLE product_variant;