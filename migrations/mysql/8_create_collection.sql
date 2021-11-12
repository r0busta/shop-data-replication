-- +migrate Up
CREATE TABLE collection (
    id BIGINT NOT NULL PRIMARY KEY,
    image_id BIGINT,
    body_html TEXT,
    handle VARCHAR(255) NOT NULL,
    published_at DATETIME,
    published_scope VARCHAR(50) NOT NULL,
    sort_order VARCHAR(50) NOT NULL,
    title VARCHAR(255),
    template_suffix VARCHAR(255),
    updated_at DATETIME,
    FOREIGN KEY (image_id) REFERENCES image(id) ON DELETE SET NULL
);
CREATE TABLE collection_product (
    collection_id BIGINT NOT NULL,
    product_id BIGINT NOT NULL,
    FOREIGN KEY (collection_id) REFERENCES collection(id) ON DELETE CASCADE,
    FOREIGN KEY (product_id) REFERENCES product(id) ON DELETE CASCADE,
    CONSTRAINT PK_collection_product PRIMARY KEY (collection_id, product_id)
);
-- +migrate Down
DROP TABLE collection_product;
DROP TABLE collection;