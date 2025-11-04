drop table if exists oms_order_return_item;
create table oms_order_return_item
(
    id            bigint auto_increment comment '主键ID'
        primary key,
    return_id     bigint                      not null comment '退货单ID（关联oms_order_return.id）',
    order_id      bigint                      not null comment '订单ID',
    order_item_id bigint                      not null comment '订单明细ID',
    sku_id        bigint                      not null comment '商品SKU ID',
    sku_name      varchar(255)                not null comment '商品名称',
    sku_pic       varchar(500)   default ''   not null comment '商品图片',
    sku_attrs     varchar(500)   default ''   not null comment '商品销售属性',
    quantity      int            default 1    not null comment '退货数量',
    product_price decimal(10, 2) default 0.00 not null comment '商品单价',
    real_amount   decimal(10, 2) default 0.00 not null comment '实际退款金额',
    reason        varchar(255)   default ''   not null comment '退货原因',
    remark        varchar(255)   default ''   not null comment '备注'
)
    comment '退货/售后明细表';

create index idx_order_id
    on oms_order_return_item (order_id);

create index idx_order_item_id
    on oms_order_return_item (order_item_id);

create index idx_return_id
    on oms_order_return_item (return_id);

INSERT INTO oms_order_return_item (
    return_id, order_id, order_item_id, sku_id, sku_name, sku_pic,
    sku_attrs, quantity, product_price, real_amount, reason, remark
) VALUES
-- Return items for return order RTN20231201001
(1, 1001, 2001, 3001, 'Apple iPhone 15 Pro', 'http://129.204.203.29/big.png',
 'Color:Black,Storage:256GB', 1, 299.00, 299.00, '商品质量问题', '屏幕有划痕'),

-- Return items for return order RTN20231201002
(2, 1002, 2002, 3002, 'Nike Air Max 270', 'http://129.204.203.29/big.png',
 'Size:42,Color:White', 1, 159.90, 159.90, '尺寸不合适', '需要换小一号'),

-- Return items for return order RTN20231202001
(3, 1003, 2003, 3003, 'Samsung Galaxy Watch5', 'http://129.204.203.29/big.png',
 'Color:Silver,Size:44mm', 1, 89.50, 89.50, '发错商品', '收到的是40mm版本'),

-- Multiple items for return order RTN20231203001
(4, 1004, 2004, 3004, 'Sony WH-1000XM4 Headphones', 'http://129.204.203.29/big.png',
 'Color:Black', 1, 199.99, 199.99, '商品不喜欢', '颜色与网站图片差异较大'),

(4, 1004, 2005, 3005, 'Sony Headphone Case', 'http://129.204.203.29/big.png',
 'Color:Black', 1, 25.00, 0.00, '配套商品退货', '主商品退货，配件一并退回'),

-- Return items for return order RTN20231204001
(5, 1005, 2006, 3006, 'iPad Air 5', 'http://129.204.203.29/big.png',
 'Color:Blue,Storage:64GB', 1, 459.00, 0.00, '人为损坏', '屏幕有裂痕'),

-- Return items for return order RTN20231205001
(6, 1006, 2007, 3007, 'MacBook Pro 13"', 'http://129.204.203.29/big.png',
 'Color:Space Gray,Storage:256GB', 1, 1299.00, 0.00, '其他原因', '客户主动取消');
