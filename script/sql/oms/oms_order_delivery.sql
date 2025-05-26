drop table if exists oms_order_delivery;
create table oms_order_delivery
(
    id                bigint auto_increment
        primary key,
    order_id          bigint                                not null comment '订单ID',
    order_no          varchar(32)                           not null comment '订单编号',
    receiver_name     varchar(50)                           not null comment '收货人姓名',
    receiver_phone    varchar(20)                           not null comment '收货人电话',
    receiver_province varchar(50)                           not null comment '省份',
    receiver_city     varchar(50)                           not null comment '城市',
    receiver_district varchar(50)                           not null comment '区县',
    receiver_address  varchar(200)                          not null comment '详细地址',
    delivery_company  varchar(50) default ''                not null comment '物流公司',
    delivery_no       varchar(50) default ''                not null comment '物流单号',
    create_time       datetime    default CURRENT_TIMESTAMP not null comment '创建时间',
    update_time       datetime                              null on update CURRENT_TIMESTAMP comment '更新时间',
    is_deleted        tinyint     default 0                 not null comment '是否删除',
    constraint uk_order
        unique (order_id, is_deleted)
)
    comment '订单收货地址表';

create index idx_delivery
    on oms_order_delivery (delivery_no, is_deleted);


