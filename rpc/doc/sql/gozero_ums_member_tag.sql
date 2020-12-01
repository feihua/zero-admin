create table ums_member_tag
(
    id                  bigint auto_increment
        primary key,
    name                varchar(100)   null,
    finish_order_count  int            null comment '自动打标签完成订单数量',
    finish_order_amount decimal(10, 2) null comment '自动打标签完成订单金额'
)
    comment '用户标签表';

