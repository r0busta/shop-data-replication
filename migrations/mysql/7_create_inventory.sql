-- +migrate Up
create table inventory_item (
    id bigint not null primary key,
    cost decimal(13, 4),
    country_code_of_origin varchar(2),
    created_at datetime,
    province_code_of_origin varchar(10),
    requires_shipping tinyint(1) not null,
    sku varchar(255),
    tracked tinyint(1) not null,
    updated_at datetime
);
create table inventory_level (
    inventory_item_id bigint not null,
    location_id bigint not null,
    available int not null,
    updated_at datetime,
    foreign key (inventory_item_id) references inventory_item(id) on delete cascade,
    foreign key (location_id) references location(id) on delete cascade,
    constraint pk_inventory_level primary key (inventory_item_id, location_id)
);
alter table product_variant
add column inventory_item_id bigint
after product_id,
    add foreign key (inventory_item_id) references image(id);
-- +migrate Down
alter table product_variant drop foreign key product_variant_ibfk_6,
    drop column inventory_item_id;
drop table inventory_level;
drop table inventory_item;