-- +migrate Up
CREATE TABLE product (
    id BIGINT NOT NULL,
    title VARCHAR(255),
    handle VARCHAR(255) NOT NULL,
    product_type VARCHAR(255),
    vendor VARCHAR(255),
    created_at DATETIME,
    updated_at DATETIME,
    status INT NOT NULL,
    PRIMARY KEY (id)
);
-- +migrate Down
DROP TABLE product;