-- +migrate Up
create table image (
    id bigint not null primary key,
    created_at datetime,
    height int,
    src varchar(255) not null,
    updated_at datetime,
    width int
);
create table product_image (
    product_id bigint not null,
    image_id bigint not null,
    position smallint,
    foreign key (product_id) references product(id) on delete cascade,
    foreign key (image_id) references image(id) on delete cascade,
    constraint pk_product_image primary key (product_id, image_id)
);
alter table product_variant
add column image_id bigint
after product_id,
    add foreign key (image_id) references image(id);
-- +migrate Down
alter table product_variant drop foreign key product_variant_ibfk_5,
    drop column image_id;
drop table product_image;
drop table image;