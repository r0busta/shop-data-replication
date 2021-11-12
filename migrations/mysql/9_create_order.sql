-- +migrate Up
CREATE TABLE customer_order (
    id BIGINT NOT NULL PRIMARY KEY,
    order_number INT NOT NULL,
    name VARCHAR(255) NOT NULL,
    note VARCHAR(255),
    email VARCHAR(255) NOT NULL,
    financial_status VARCHAR(50) NOT NULL,
    fulfillment_status VARCHAR(50) NOT NULL,
    total_discounts DECIMAL(13, 4) NOT NULL,
    total_line_items_price DECIMAL(13, 4) NOT NULL,
    total_outstanding DECIMAL(13, 4) NOT NULL,
    total_price DECIMAL(13, 4) NOT NULL,
    total_tax DECIMAL(13, 4) NOT NULL,
    created_at DATETIME,
    updated_at DATETIME,
    processed_at DATETIME,
    cancelled_at DATETIME,
    closed_at DATETIME
);
-- +migrate Down
DROP TABLE customer_order;