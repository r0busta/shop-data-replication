-- +migrate Up
CREATE TABLE customer_address (
    id BIGINT NOT NULL PRIMARY KEY,
    customer_id BIGINT NOT NOT NULL,
    address1 VARCHAR(255),
    address2 VARCHAR(255),
    city VARCHAR(255),
    company VARCHAR(255),
    country VARCHAR(255),
    country_code VARCHAR(2),
    first_name VARCHAR(255),
    is_default TINYINT(1) NOT NULL,
    last_name VARCHAR(255),
    name VARCHAR(255),
    phone VARCHAR(255),
    province VARCHAR(255),
    province_code VARCHAR(2),
    zip VARCHAR(255),
    FOREIGN KEY (customer_id) REFERENCES customer(id) ON DELETE CASCADE,
);
ALTER TABLE customer
ADD COLUMN default_address_id BIGINT
AFTER id,
    ADD FOREIGN KEY (default_address_id) REFERENCES customer_address(id);
-- +migrate Down
ALTER TABLE customer DROP FOREIGN KEY customer_ibfk_0,
    DROP COLUMN default_address_id;
DROP TABLE customer_address;