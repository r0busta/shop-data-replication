-- +migrate Up
CREATE TABLE image (
    id BIGINT NOT NULL PRIMARY KEY,
    created_at DATETIME,
    height INT,
    src VARCHAR(255) NOT NULL,
    updated_at DATETIME,
    width INT
);
CREATE TABLE product_image (
    product_id BIGINT NOT NULL,
    image_id BIGINT NOT NULL,
    position SMALLINT,
    FOREIGN KEY (product_id) REFERENCES product(id) ON DELETE CASCADE,
    FOREIGN KEY (image_id) REFERENCES image(id) ON DELETE CASCADE,
    CONSTRAINT PK_product_image PRIMARY KEY (product_id, image_id)
);
ALTER TABLE product_variant
ADD COLUMN image_id BIGINT
AFTER product_id,
    ADD FOREIGN KEY (image_id) REFERENCES image(id);
-- +migrate Down
ALTER TABLE product_variant DROP FOREIGN KEY product_variant_ibfk_5,
    DROP COLUMN image_id;
DROP TABLE product_image;
DROP TABLE image;