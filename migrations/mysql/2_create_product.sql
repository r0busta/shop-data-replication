-- +migrate Up
CREATE TABLE product (
    id BIGINT NOT NULL PRIMARY KEY,
    body_html TEXT,
    title VARCHAR(255),
    handle VARCHAR(255) NOT NULL,
    product_type VARCHAR(255),
    vendor VARCHAR(255),
    created_at DATETIME,
    updated_at DATETIME,
    published_at DATETIME,
    published_scope VARCHAR(50) NOT NULL,
    status TINYINT(2) NOT NULL,
    template_suffix VARCHAR(255)
);
-- +migrate Down
DROP TABLE product;