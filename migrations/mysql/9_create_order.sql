-- +migrate Up
create table customer_order (
    id bigint not null primary key,
    order_number int not null,
    name varchar(255) not null,
    note varchar(255),
    email varchar(255) not null,
    financial_status varchar(50) not null,
    fulfillment_status varchar(50) not null,
    total_discounts decimal(13, 4) not null,
    total_line_items_price decimal(13, 4) not null,
    total_outstanding decimal(13, 4) not null,
    total_price decimal(13, 4) not null,
    total_tax decimal(13, 4) not null,
    created_at datetime,
    updated_at datetime,
    processed_at datetime,
    cancelled_at datetime,
    closed_at datetime
);
-- +migrate Down
drop table customer_order;