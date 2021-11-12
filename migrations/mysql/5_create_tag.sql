-- +migrate Up
create table tag (
    id bigint not null primary key,
    value varchar(255)
);
create table product_tag (
    tag_id bigint not null,
    product_id bigint not null,
    value varchar(255) not null,
    foreign key (tag_id) references tag(id) on delete cascade,
    foreign key (product_id) references product(id) on delete cascade,
    constraint pk_product_tag primary key (tag_id, product_id)
);
-- +migrate Down
drop table product_tag;
drop table tag;