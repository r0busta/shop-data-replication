-- +migrate Up
create table customer_address (
    id bigint not null primary key,
    customer_id bigint not null,
    address1 varchar(255),
    address2 varchar(255),
    city varchar(255),
    company varchar(255),
    country varchar(255),
    country_code varchar(2),
    first_name varchar(255),
    is_default tinyint(1) not null,
    last_name varchar(255),
    name varchar(255),
    phone varchar(255),
    province varchar(255),
    province_code varchar(2),
    zip varchar(255),
    foreign key (customer_id) references customer(id) on delete cascade
);

alter table customer
    add column default_address_id bigint
    after id,
        add foreign key (default_address_id) references customer_address(id);

-- +migrate Down
alter table customer drop foreign key customer_ibfk_1,
    drop column default_address_id;
drop table customer_address;