CREATE TABLE product
(
    id           serial primary key not null unique,
    track_number varchar(255)       not null unique,
    price        int                not null,
    rid          varchar(455)       not null,
    name         text               not null,
    sale         int,
    size         varchar(255),
    total_price  int                not null,
    nm_id        int                not null,
    brand        varchar(255)       not null,
    status       int                not null
);

CREATE TABLE payment
(
    id            serial primary key not null unique,
    transaction   varchar(455)       not null,
    request_id    varchar(455),
    currency      varchar(100),
    provider      varchar(255)       not null,
    amount        int                not null,
    payment_dt    int                not null,
    bank          varchar(255)       not null,
    delivery_cost int                not null,
    goods_total   int                not null,
    custom_fee    int                not null
);

CREATE TABLE delivery
(
    id      serial primary key not null unique,
    name    varchar(455)       not null,
    phone   varchar(100)       not null,
    zip     varchar(255)       not null,
    city    varchar(255)       not null,
    address varchar(555)       not null,
    region  varchar(355)       not null,
    email   varchar(555)       not null
);

CREATE TABLE "order"
(
    id                 serial primary key                                               not null unique,
    order_uid          text                                                             not null unique,
    track_number       varchar(255)                                                     not null,
    entry              varchar(255)                                                     not null,
    locale             varchar(100)                                                     not null,
    internal_signature varchar(555),
    customer_id        varchar(255)                                                     not null,
    delivery_service   varchar(255),
    shardkey           varchar(255)                                                     not null,
    sm_id              int                                                              not null,
    oof_shard          varchar(255)                                                     not null,
    delivery_id        int references delivery (id) on update cascade on delete cascade not null,
    payment            int references payment (id) on update cascade on delete cascade  not null
);

CREATE TABLE order_product
(
    id         serial primary key                                              not null unique,
    order_id   int references "order" (id) on update cascade on delete cascade   not null,
    product_id int references product (id) on update cascade on delete cascade not null
);
