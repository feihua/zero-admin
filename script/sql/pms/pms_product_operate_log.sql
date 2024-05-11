create table pms_product_operate_log
(
    id                  bigint auto_increment
        primary key,
    product_id          bigint         not null,
    price_old           decimal(10, 2) not null,
    price_new           decimal(10, 2) not null,
    sale_price_old      decimal(10, 2) not null,
    sale_price_new      decimal(10, 2) not null,
    gift_point_old      int            not null comment '赠送的积分',
    gift_point_new      int            not null,
    use_point_limit_old int            not null,
    use_point_limit_new int            not null,
    operate_man         varchar(64)    not null comment '操作人',
    create_time         datetime       not null
);

