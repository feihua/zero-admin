drop table if exists oms_order_payment;
create table oms_order_payment
(
    id             bigint auto_increment
        primary key comment '主键ID',
    order_id       bigint                                not null comment '订单ID',
    order_no       varchar(32)                           not null comment '订单编号',
    pay_type       tinyint                               not null comment '支付方式：1-支付宝，2-微信，3-银联',
    transaction_id varchar(50) default ''                not null comment '支付流水号',
    total_amount   decimal(10, 2)                        not null comment '订单金额',
    pay_amount     decimal(10, 2)                        not null comment '支付金额',
    pay_status     tinyint     default 0                 not null comment '支付状态：0-待支付，1-支付成功，2-支付失败',
    pay_time       datetime                              null comment '支付时间',
    create_time    datetime    default CURRENT_TIMESTAMP not null comment '创建时间',
    update_time    datetime                              null on update CURRENT_TIMESTAMP,
    is_deleted     tinyint     default 0                 not null comment '是否删除'
)
    comment '订单支付表';

create index idx_transaction
    on oms_order_payment (transaction_id, is_deleted);

INSERT INTO oms_order_payment (
    order_id, order_no, pay_type, transaction_id, total_amount,
    pay_amount, pay_status, pay_time, create_time, update_time
) VALUES
-- Payment info for order ORD20231201001 (pending payment)
(1, 'ORD20231201001', 1, 'ALI2023120209159876541', 299.00,
 299.00, 2, NULL, '2023-12-01 10:30:00', NULL),
(1, 'ORD20231201001', 2, 'WX20231203164511223311', 299.00,
 299.00, 0, NULL, '2023-12-01 10:30:00', NULL),

-- Payment info for order ORD20231201002 (paid via WeChat)
(2, 'ORD20231201002', 2, 'WX20231201150012345678', 159.90,
 139.90, 1, '2023-12-01 15:00:00', '2023-12-01 14:20:00', '2023-12-01 15:00:00'),

-- Payment info for order ORD20231202001 (paid via Alipay)
(3, 'ORD20231202001', 1, 'ALI20231202091598765432', 89.50,
 99.50, 1, '2023-12-02 09:15:00', '2023-12-02 09:15:00', '2023-12-02 09:15:00'),

-- Payment info for order ORD20231203001 (paid via WeChat)
(4, 'ORD20231203001', 2, 'WX20231203164511223344', 199.99,
 179.99, 1, '2023-12-03 16:45:00', '2023-12-03 16:45:00', '2023-12-03 16:45:00'),

-- Payment info for order ORD20231204001 (cancelled, no payment)
(5, 'ORD20231204001', 0, '', 459.00,
 459.00, 0, NULL, '2023-12-04 11:20:00', NULL),

-- Payment info for order ORD20231205001 (paid via Alipay)
(6, 'ORD20231205001', 1, 'ALI20231205153055667788', 1299.00,
 1199.00, 1, '2023-12-05 15:30:00', '2023-12-05 15:30:00', '2023-12-08 10:00:00'),

-- Payment info for order ORD20231206001 (paid via UnionPay)
(7, 'ORD20231206001', 3, 'UNION20231206111599887766', 299.00,
 299.00, 1, '2023-12-06 11:15:00', '2023-12-06 11:15:00', '2023-12-06 11:15:00');
