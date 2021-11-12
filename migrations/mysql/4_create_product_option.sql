-- +migrate Up
CREATE TABLE product_option (
    id BIGINT NOT NULL PRIMARY KEY,
    product_id BIGINT NOT NULL,
    position INT NOT NULL,
    name VARCHAR(255) NOT NULL,
    FOREIGN KEY (product_id) REFERENCES product(id) ON DELETE CASCADE
);
CREATE TABLE product_option_value (
    option_id BIGINT NOT NULL PRIMARY KEY,
    value VARCHAR(255) NOT NULL,
    FOREIGN KEY (option_id) REFERENCES product_option(id) ON DELETE CASCADE
);
ALTER TABLE product_variant
ADD COLUMN option_1_value_id BIGINT,
    ADD FOREIGN KEY (option_1_value_id) REFERENCES product_option_value(option_id),
    ADD COLUMN option_2_value_id BIGINT,
    ADD FOREIGN KEY (option_2_value_id) REFERENCES product_option_value(option_id),
    ADD COLUMN option_3_value_id BIGINT,
    ADD FOREIGN KEY (option_3_value_id) REFERENCES product_option_value(option_id);
-- +migrate Down
ALTER TABLE product_variant DROP FOREIGN KEY product_variant_ibfk_2,
    DROP COLUMN option_1_value_id,
    DROP FOREIGN KEY product_variant_ibfk_3,
    DROP COLUMN option_2_value_id,
    DROP FOREIGN KEY product_variant_ibfk_4,
    DROP COLUMN option_3_value_id;
DROP TABLE product_option_value;
DROP TABLE product_option;