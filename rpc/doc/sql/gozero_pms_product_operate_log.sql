create table pms_product_operate_log
(
    id                  bigint auto_increment
        primary key,
    product_id          bigint         null,
    price_old           decimal(10, 2) null,
    price_new           decimal(10, 2) null,
    sale_price_old      decimal(10, 2) null,
    sale_price_new      decimal(10, 2) null,
    gift_point_old      int            null comment '赠送的积分',
    gift_point_new      int            null,
    use_point_limit_old int            null,
    use_point_limit_new int            null,
    operate_man         varchar(64)    null comment '操作人',
    create_time         datetime       null
);

