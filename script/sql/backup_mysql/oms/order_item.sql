drop table if exists oms_order_item;
create table oms_order_item
(
    id                bigint auto_increment
        primary key,
    order_id          bigint                                      not null comment '订单ID',
    order_no          varchar(32)                                 not null comment '订单编号',
    order_item_status tinyint        default 1                    not null comment '订单商品状态：1-正常,2-退货申请中,3-已退货,4-已拒绝',
    sku_id            bigint                                      not null comment '商品SKU ID',
    sku_name          varchar(200)                                not null comment '商品名称',
    sku_pic           varchar(500)                                not null comment '商品图片',
    sku_price         decimal(10, 2)                              not null comment '商品单价',
    sku_quantity      int                                         not null comment '商品数量',
    spec_data         json                                        not null comment '规格数据',
    sku_total_amount  decimal(10, 2)                              not null comment '商品总金额',
    promotion_amount  decimal(10, 2) default 0.00                 not null comment '促销分摊金额',
    coupon_amount     decimal(10, 2) default 0.00                 not null comment '优惠券分摊金额',
    points_amount     decimal(10, 2) default 0.00                 not null comment '积分分摊金额',
    discount_amount   decimal(10, 2) default 0.00                 not null comment '优惠分摊金额',
    real_amount       decimal(10, 2)                              not null comment '实付金额',
    create_time       datetime(3)    default CURRENT_TIMESTAMP(3) not null comment '创建时间',
    is_deleted        tinyint(1)     default 0                    not null comment '是否删除'
)
    comment '订单商品表';

create index idx_order
    on oms_order_item (order_id, is_deleted);

create index idx_sku
    on oms_order_item (sku_id, is_deleted);

INSERT INTO oms_order_item (
    order_id, order_no, order_item_status, sku_id, sku_name, sku_pic,
    sku_price, sku_quantity, spec_data, sku_total_amount, promotion_amount,
    coupon_amount, points_amount, discount_amount, real_amount, create_time
) VALUES
-- Order items for ORD20231201001
(1, 'ORD20231201001', 1, 3001, 'Apple iPhone 15 Pro', 'http://129.204.203.29/big.png',
 299.00, 1, '{"color": "Black", "storage": "256GB"}', 299.00, 0.00,
 0.00, 0.00, 0.00, 299.00, '2023-12-01 10:30:00.000'),

(1, 'ORD20231201001', 1, 3001, 'Apple iPhone 15 Pro', 'http://129.204.203.29/big.png',
 299.00, 1, '{"color": "Black", "storage": "256GB"}', 299.00, 0.00,
 0.00, 0.00, 0.00, 299.00, '2023-12-01 10:30:00.000'),

-- Order items for ORD20231201002
(2, 'ORD20231201002', 1, 3002, 'Nike Air Max 270', 'http://129.204.203.29/big.png',
 159.90, 1, '{"size": "42", "color": "White"}', 159.90, 10.00,
 5.00, 0.00, 5.00, 139.90, '2023-12-01 14:20:00.000'),

-- Order items for ORD20231202001
(3, 'ORD20231202001', 1, 3003, 'Samsung Galaxy Watch5', 'http://129.204.203.29/big.png',
 89.50, 1, '{"color": "Silver", "size": "44mm"}', 89.50, 0.00,
 0.00, 0.00, 0.00, 99.50, '2023-12-02 09:15:00.000'),

-- Multiple items for ORD20231203001
(4, 'ORD20231203001', 1, 3004, 'Sony WH-1000XM4 Headphones', 'http://129.204.203.29/big.png',
 199.99, 1, '{"color": "Black"}', 199.99, 0.00,
 20.00, 0.00, 0.00, 179.99, '2023-12-03 16:45:00.000'),

-- Order items for ORD20231204001
(5, 'ORD20231204001', 1, 3006, 'iPad Air 5', 'http://129.204.203.29/big.png',
 459.00, 1, '{"color": "Blue", "storage": "64GB"}', 459.00, 0.00,
 0.00, 0.00, 0.00, 459.00, '2023-12-04 11:20:00.000'),

-- Order items for ORD20231205001
(6, 'ORD20231205001', 1, 3007, 'MacBook Pro 13"', 'http://129.204.203.29/big.png',
 1299.00, 1, '{"color": "Space Gray", "storage": "256GB"}', 1299.00, 100.00,
 50.00, 0.00, 50.00, 1199.00, '2023-12-05 15:30:00.000'),

-- Multiple items for ORD20231206001
(7, 'ORD20231206001', 1, 3001, 'Apple iPhone 15 Pro', 'http://129.204.203.29/big.png',
 299.00, 1, '{"color": "Black", "storage": "256GB"}', 299.00, 0.00,
 0.00, 0.00, 0.00, 299.00, '2023-12-06 11:15:00.000'),

(7, 'ORD20231206001', 2, 3008, 'iPhone Screen Protector', 'http://129.204.203.29/big.png',
 15.00, 1, '{"type": "Tempered Glass"}', 15.00, 0.00,
 0.00, 0.00, 0.00, 0.00, '2023-12-06 11:15:00.000');
