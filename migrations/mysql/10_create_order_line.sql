-- +migrate Up
create table order_line (
    id bigint not null primary key,
    order_id bigint not null,
    product_id bigint,
    variant_id bigint,
    sku varchar(255),
    name varchar(255),
    title varchar(255),
    variant_title varchar(255),
    vendor varchar(255),
    quantity int not null,
    price decimal(13, 4) not null,
    total_discount decimal(13, 4) not null,
    foreign key (order_id) references customer_order(id) on delete cascade,
    foreign key (product_id) references product(id) on delete set null,
    foreign key (variant_id) references product_variant(id) on delete set null
);
-- +migrate Down
drop table order_line;