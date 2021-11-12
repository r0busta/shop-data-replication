-- +migrate Up
CREATE TABLE customer (
    id BIGINT NOT NULL PRIMARY KEY,
    accepts_marketing TINYINT(1) NOT NULL,
    accepts_marketing_updated_at DATETIME NOT NULL,
    created_at DATETIME NOT NULL,
    email VARCHAR(255),
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255),
    marketing_opt_in_level VARCHAR(50) NOT NULL,
    note VARCHAR(255),
    order_count INT NOT NULL,
    phone VARCHAR(100) NOT NULL,
    state VARCHAR(100) NOT NULL,
    tax_exempt TINYINT(1) NOT NULL,
    total_spent DECIMAL(13, 4) NOT NULL,
    updated_at DATETIME NOT NULL,
    verified_email TINYINT(1) NOT NULL
);
-- +migrate Down
DROP TABLE customer;