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

INSERT INTO oms_order_operation_log (
    order_id, order_no, operation_type, operator_id, operator_type, operator_note, create_time
) VALUES
-- Order creation logs
(1, 'ORD20231201001', 1, 10001, 1, '用户创建订单', '2023-12-01 10:30:00'),
(2, 'ORD20231201002', 1, 10002, 1, '用户创建订单', '2023-12-01 14:20:00'),
(3, 'ORD20231202001', 1, 10003, 1, '用户创建订单', '2023-12-02 09:15:00'),
(4, 'ORD20231203001', 1, 10004, 1, '用户创建订单', '2023-12-03 16:45:00'),
(5, 'ORD20231204001', 1, 10005, 1, '用户创建订单', '2023-12-04 11:20:00'),
(6, 'ORD20231205001', 1, 10006, 1, '用户创建订单', '2023-12-05 15:30:00'),
(7, 'ORD20231206001', 1, 10007, 1, '用户创建订单', '2023-12-06 11:15:00'),

-- Payment logs
(1, 'ORD20231201001', 2, 10002, 1, '用户完成支付', '2023-12-01 15:00:00'),
(2, 'ORD20231201002', 2, 10002, 1, '用户完成支付', '2023-12-01 15:00:00'),
(3, 'ORD20231202001', 2, 10003, 1, '用户完成支付', '2023-12-02 09:15:00'),
(4, 'ORD20231203001', 2, 10004, 1, '用户完成支付', '2023-12-03 16:45:00'),
(6, 'ORD20231205001', 2, 10006, 1, '用户完成支付', '2023-12-05 15:30:00'),
(7, 'ORD20231206001', 2, 10007, 1, '用户完成支付', '2023-12-06 11:15:00'),

-- Shipping logs
(3, 'ORD20231202001', 3, 1, 3, '管理员发货', '2023-12-02 14:30:00'),
(4, 'ORD20231203001', 3, 1, 3, '管理员发货', '2023-12-04 09:00:00'),
(6, 'ORD20231205001', 3, 1, 3, '管理员发货', '2023-12-06 10:00:00'),
(7, 'ORD20231206001', 3, 1, 3, '管理员发货', '2023-12-07 15:00:00'),

-- Receipt logs
(4, 'ORD20231203001', 4, 10004, 1, '用户确认收货', '2023-12-05 10:00:00'),
(6, 'ORD20231205001', 4, 10006, 1, '用户确认收货', '2023-12-07 14:00:00'),
(7, 'ORD20231206001', 4, 10007, 1, '用户确认收货', '2023-12-08 09:00:00'),

-- Cancellation logs
(5, 'ORD20231204001', 5, 10005, 1, '用户取消订单', '2023-12-04 13:00:00'),

-- Refund logs
(6, 'ORD20231205001', 6, 1, 3, '管理员处理退款', '2023-12-08 10:00:00');
