drop table if exists oms_order_operation_log;
create table oms_order_operation_log
(
    id             bigint auto_increment
        primary key comment '主键ID',
    order_id       bigint                                 not null comment '订单ID',
    order_no       varchar(32)                            not null comment '订单编号',
    operation_type tinyint                                not null comment '操作类型：1-创建订单，2-支付订单，3-发货，4-确认收货，5-取消订单，6-退款',
    operator_id    bigint                                 not null comment '操作人ID',
    operator_type  tinyint                                not null comment '操作人类型：1-用户，2-系统，3-管理员',
    operator_note  varchar(500) default ''                not null comment '操作备注',
    create_time    datetime     default CURRENT_TIMESTAMP not null comment '操作时间'
)
    comment '订单操作记录表';

create index idx_operator
    on oms_order_operation_log (operator_id, operator_type);

create index idx_order
    on oms_order_operation_log (order_id);

