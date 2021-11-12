-- +migrate Up
CREATE TABLE tag (
    id BIGINT NOT NULL PRIMARY KEY,
    value VARCHAR(255)
);
CREATE TABLE product_tag (
    tag_id BIGINT NOT NULL,
    product_id BIGINT NOT NULL,
    value VARCHAR(255) NOT NULL,
    FOREIGN KEY (tag_id) REFERENCES tag(id) ON DELETE CASCADE,
    FOREIGN KEY (product_id) REFERENCES product(id) ON DELETE CASCADE,
    CONSTRAINT PK_product_tag PRIMARY KEY (tag_id, product_id)
);
-- +migrate Down
DROP TABLE product_tag;
DROP TABLE tag;