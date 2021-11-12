-- +migrate Up
create table product (
    id bigint not null primary key,
    body_html text,
    title varchar(255),
    handle varchar(255) not null,
    product_type varchar(255),
    vendor varchar(255),
    created_at datetime,
    updated_at datetime,
    published_at datetime,
    published_scope varchar(50) not null,
    status tinyint(2) not null,
    template_suffix varchar(255)
);
-- +migrate Down
drop table product;