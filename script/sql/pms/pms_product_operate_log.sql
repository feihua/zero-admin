create table pms_product_operate_log
(
    id                  bigint auto_increment
        primary key,
    product_id          bigint                             not null comment '产品id',
    price_old           bigint                             not null comment '原价',
    price_new           bigint                             not null comment '新价格',
    sale_price_old      bigint                             not null comment '原售价',
    sale_price_new      bigint                             not null comment '新售价',
    gift_point_old      int                                not null comment '赠送的积分',
    gift_point_new      int                                not null comment '新的积分',
    use_point_limit_old int                                not null comment '使用积分限制',
    use_point_limit_new int                                not null comment '新的使用积分限制',
    operate_man         varchar(64)                        not null comment '操作人',
    create_time         datetime default CURRENT_TIMESTAMP not null comment '创建时间'
) comment '商品操作日志'

