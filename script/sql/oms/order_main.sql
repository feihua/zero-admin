drop table if exists oms_order_main;
create table oms_order_main
(
    id                   bigint auto_increment
        primary key,
    order_no             varchar(32)                              not null comment '订单编号',
    user_id              bigint                                   not null comment '用户ID',
    order_status         tinyint        default 1                 not null comment '订单状态：1-待支付,2-已支付,3-已发货,4-已完成,5-已取消,6-已退款,7-售后中',
    total_amount         decimal(10, 2)                           not null comment '订单总金额',
    promotion_amount     decimal(10, 2) default 0.00              not null comment '促销金额',
    coupon_amount        decimal(10, 2) default 0.00              not null comment '优惠券金额',
    points_amount        decimal(10, 2) default 0.00              not null comment '积分金额',
    discount_amount      decimal(10, 2) default 0.00              not null comment '优惠金额',
    freight_amount       decimal(10, 2) default 0.00              not null comment '运费金额',
    pay_amount           decimal(10, 2)                           not null comment '实付金额',
    pay_type             tinyint                                  null comment '支付方式：1-支付宝,2-微信,3-银联',
    pay_time             datetime                                 null comment '支付时间',
    delivery_time        datetime                                 null comment '发货时间',
    receive_time         datetime                                 null comment '收货时间',
    comment_time         datetime                                 null comment '评价时间',
    source_type          tinyint        default 1                 not null comment '订单来源：1-APP,2-PC,3-小程序',
    express_order_number varchar(64)                              not null comment '快递单号',
    use_points           int            default 0                 not null comment '下单时使用的积分',
    receive_status       tinyint                                  not null comment '是否确认收货：0->否,1->是',
    remark               varchar(500)   default ''                not null comment '订单备注',
    create_time          datetime       default CURRENT_TIMESTAMP not null comment '提交时间',
    update_time          datetime                                 null on update CURRENT_TIMESTAMP,
    is_deleted           tinyint        default 0                 not null comment '是否删除',
    constraint uk_order_no
        unique (order_no, is_deleted)
)
    comment '订单主表';

create index idx_time
    on oms_order_main (create_time, is_deleted);

create index idx_user
    on oms_order_main (user_id, order_status, is_deleted);

INSERT INTO oms_order_main (
    order_no, user_id, order_status, total_amount, promotion_amount, coupon_amount,
    points_amount, discount_amount, freight_amount, pay_amount, pay_type, pay_time,
    delivery_time, receive_time, comment_time, source_type, express_order_number,
    use_points, receive_status, remark, create_time, update_time
) VALUES
-- Pending payment order
('ORD20231201001', 10001, 1, 299.00, 0.00, 0.00,
 0.00, 0.00, 0.00, 299.00, NULL, NULL,
 NULL, NULL, NULL, 1, '',
 0, 0, '新订单', '2023-12-01 10:30:00', NULL),

-- Paid order
('ORD20231201002', 10002, 2, 159.90, 10.00, 5.00,
 0.00, 5.00, 0.00, 139.90, 2, '2023-12-01 15:00:00',
 NULL, NULL, NULL, 2, '',
 0, 0, '已支付', '2023-12-01 14:20:00', '2023-12-01 15:00:00'),

-- Shipped order
('ORD20231202001', 10003, 3, 89.50, 0.00, 0.00,
 0.00, 0.00, 10.00, 99.50, 1, '2023-12-02 09:15:00',
 '2023-12-02 14:30:00', NULL, NULL, 3, 'SF123456789CN',
 100, 0, '已发货', '2023-12-02 09:15:00', '2023-12-02 14:30:00'),

-- Completed order
('ORD20231203001', 10004, 4, 199.99, 0.00, 20.00,
 0.00, 0.00, 0.00, 179.99, 2, '2023-12-03 16:45:00',
 '2023-12-04 09:00:00', '2023-12-05 10:00:00', '2023-12-05 11:00:00', 1, 'YT987654321CN',
 0, 1, '已完成', '2023-12-03 16:45:00', '2023-12-05 11:00:00'),

-- Cancelled order
('ORD20231204001', 10005, 5, 459.00, 0.00, 0.00,
 0.00, 0.00, 0.00, 459.00, NULL, NULL,
 NULL, NULL, NULL, 2, '',
 0, 0, '用户取消', '2023-12-04 11:20:00', '2023-12-04 13:00:00'),

-- Refunded order
('ORD20231205001', 10006, 6, 1299.00, 100.00, 50.00,
 0.00, 50.00, 0.00, 1199.00, 1, '2023-12-05 15:30:00',
 '2023-12-06 10:00:00', '2023-12-07 14:00:00', NULL, 1, 'EMS1122334455',
 0, 1, '已退款', '2023-12-05 15:30:00', '2023-12-08 10:00:00'),

-- After-sales order
('ORD20231206001', 10007, 7, 299.00, 0.00, 0.00,
 0.00, 0.00, 0.00, 299.00, 2, '2023-12-06 11:15:00',
 '2023-12-07 15:00:00', '2023-12-08 09:00:00', NULL, 3, 'ZT5566778899',
 0, 1, '售后处理中', '2023-12-06 11:15:00', '2023-12-08 14:00:00');
