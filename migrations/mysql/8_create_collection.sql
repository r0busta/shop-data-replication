-- +migrate Up
create table collection (
    id bigint not null primary key,
    image_id bigint,
    body_html text,
    handle varchar(255) not null,
    published_at datetime,
    published_scope varchar(50) not null,
    sort_order varchar(50) not null,
    title varchar(255),
    template_suffix varchar(255),
    updated_at datetime,
    foreign key (image_id) references image(id) on delete set null
);
create table collection_product (
    collection_id bigint not null,
    product_id bigint not null,
    foreign key (collection_id) references collection(id) on delete cascade,
    foreign key (product_id) references product(id) on delete cascade,
    constraint pk_collection_product primary key (collection_id, product_id)
);
-- +migrate Down
drop table collection_product;
drop table collection;