-- +migrate Up
create table location(
    id bigint not null primary key,
    active tinyint(1) not null,
    address1 varchar(255),
    address2 varchar(255),
    city varchar(255),
    created_at datetime,
    name varchar(255) not null,
    phone varchar(100),
    updated_at datetime,
    zip varchar(100),
    country_code varchar(2),
    province_code varchar(10)
);
-- +migrate Down
drop table location;