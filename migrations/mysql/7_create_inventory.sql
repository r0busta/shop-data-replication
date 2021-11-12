-- +migrate Up
CREATE TABLE inventory_item (
    id BIGINT NOT NULL PRIMARY KEY,
    cost DECIMAL(13, 4),
    country_code_of_origin VARCHAR(2),
    created_at DATETIME,
    province_code_of_origin VARCHAR(10),
    requires_shipping TINYINT(1) NOT NULL,
    sku VARCHAR(255),
    tracked TINYINT(1) NOT NULL,
    updated_at DATETIME
);
CREATE TABLE inventory_level (
    inventory_item_id BIGINT NOT NULL,
    location_id BIGINT NOT NULL,
    available INT NOT NULL,
    updated_at DATETIME,
    FOREIGN KEY (inventory_item_id) REFERENCES inventory_item(id) ON DELETE CASCADE,
    FOREIGN KEY (location_id) REFERENCES location(id) ON DELETE CASCADE,
    CONSTRAINT PK_inventory_level PRIMARY KEY (inventory_item_id, location_id)
);
ALTER TABLE product_variant
ADD COLUMN inventory_item_id BIGINT
AFTER product_id,
    ADD FOREIGN KEY (inventory_item_id) REFERENCES image(id);
-- +migrate Down
ALTER TABLE product_variant DROP FOREIGN KEY product_variant_ibfk_6,
    DROP COLUMN inventory_item_id;
DROP TABLE inventory_level;
DROP TABLE inventory_item;