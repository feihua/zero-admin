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

