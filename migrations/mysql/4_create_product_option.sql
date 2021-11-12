-- +migrate Up
create table product_option (
    id bigint not null primary key,
    product_id bigint not null,
    position int not null,
    name varchar(255) not null,
    foreign key (product_id) references product(id) on delete cascade
);
create table product_option_value (
    option_id bigint not null primary key,
    value varchar(255) not null,
    foreign key (option_id) references product_option(id) on delete cascade
);
alter table product_variant
add column option_1_value_id bigint,
    add foreign key (option_1_value_id) references product_option_value(option_id),
    add column option_2_value_id bigint,
    add foreign key (option_2_value_id) references product_option_value(option_id),
    add column option_3_value_id bigint,
    add foreign key (option_3_value_id) references product_option_value(option_id);
-- +migrate Down
alter table product_variant drop foreign key product_variant_ibfk_2,
    drop column option_1_value_id,
    drop foreign key product_variant_ibfk_3,
    drop column option_2_value_id,
    drop foreign key product_variant_ibfk_4,
    drop column option_3_value_id;
drop table product_option_value;
drop table product_option;