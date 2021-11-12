-- +migrate Up
create table product_variant (
    id bigint not null primary key,
    product_id bigint not null,
    barcode varchar(255),
    compare_at_price decimal(13, 4),
    fulfillment_service varchar(255),
    grams double,
    inventory_management varchar(50) not null,
    inventory_policy varchar(50) not null,
    inventory_quantity int not null,
    position smallint not null,
    price decimal(13, 4) not null,
    sku varchar(255),
    taxable tinyint(1) not null,
    title varchar(255),
    weight double,
    weight_unit varchar(4) not null,
    created_at datetime,
    updated_at datetime,
    foreign key (product_id) references product(id) on delete cascade
);
-- +migrate Down
drop table product_variant;