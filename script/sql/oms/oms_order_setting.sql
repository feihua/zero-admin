create table oms_order_setting
(
    id                    bigint auto_increment
        primary key,
    flash_order_overtime  int               not null comment '秒杀订单超时关闭时间(分)',
    normal_order_overtime int               not null comment '正常订单超时时间(分)',
    confirm_overtime      int               not null comment '发货后自动确认收货时间（天）',
    finish_overtime       int               not null comment '自动完成交易时间，不能申请售后（天）',
    status                tinyint default 1 not null comment '状态：0->禁用；1->启用',
    is_default            tinyint default 0 not null comment '是否默认：0->否；1->是',
    comment_overtime      int               not null comment '订单完成后自动好评时间（天）'

)
    comment '订单设置表';

INSERT INTO oms_order_setting (id, flash_order_overtime, normal_order_overtime, confirm_overtime, finish_overtime, comment_overtime) VALUES (1, 60, 120, 15, 7, 7);
