-- +migrate Up
create table customer (
    id bigint not null primary key,
    accepts_marketing tinyint(1) not null,
    accepts_marketing_updated_at datetime not null,
    created_at datetime not null,
    email varchar(255),
    first_name varchar(255) not null,
    last_name varchar(255),
    marketing_opt_in_level varchar(50) not null,
    note varchar(255),
    order_count int not null,
    phone varchar(100) not null,
    state varchar(100) not null,
    tax_exempt tinyint(1) not null,
    total_spent decimal(13, 4) not null,
    updated_at datetime not null,
    verified_email tinyint(1) not null
);
-- +migrate Down
drop table customer;