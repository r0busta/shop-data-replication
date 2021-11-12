-- +migrate Up
CREATE TABLE location(
    id BIGINT NOT NULL PRIMARY KEY,
    active TINYINT(1) NOT NULL,
    address1 VARCHAR(255),
    address2 VARCHAR(255),
    city VARCHAR(255),
    created_at DATETIME,
    name VARCHAR(255) NOT NULL,
    phone VARCHAR(100),
    updated_at DATETIME,
    zip VARCHAR(100),
    country_code VARCHAR(2),
    province_code VARCHAR(10)
);
-- +migrate Down
DROP TABLE location;