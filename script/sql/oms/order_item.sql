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

