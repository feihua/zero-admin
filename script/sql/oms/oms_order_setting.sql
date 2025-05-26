drop table if exists oms_order_setting;
create table oms_order_setting
(
    id                    bigint auto_increment
        primary key        comment '主键ID',
    flash_order_overtime  int                                not null comment '秒杀订单超时关闭时间(分)',
    normal_order_overtime int                                not null comment '正常订单超时时间(分)',
    confirm_overtime      int                                not null comment '发货后自动确认收货时间（天）',
    finish_overtime       int                                not null comment '自动完成交易时间，不能申请售后（天）',
    status                tinyint  default 1                 not null comment '状态：0->禁用；1->启用',
    is_default            tinyint  default 0                 not null comment '是否默认：0->否；1->是',
    comment_overtime      int                                not null comment '订单完成后自动好评时间（天）',
    create_by             bigint                             not null comment '创建人ID',
    create_time           datetime default CURRENT_TIMESTAMP not null comment '创建时间',
    update_by             bigint                             null comment '更新人ID',
    update_time           datetime                           null on update CURRENT_TIMESTAMP comment '更新时间',
    is_deleted            tinyint  default 0                 not null comment '是否删除'

)
    comment '订单设置表';

-- 插入订单设置数据
insert into oms_order_setting (id, flash_order_overtime, normal_order_overtime, confirm_overtime, finish_overtime, status, is_default, comment_overtime, create_by, is_deleted)
values (1, 30, 60, 7, 15, 1, 1, 7, 1, 0),
       (2, 45, 90, 10, 20, 1, 0, 10, 2, 0),
       (3, 60, 120, 5, 10, 0, 0, 5, 3, 0);