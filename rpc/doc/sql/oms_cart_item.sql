create table oms_cart_item
(
    id                  bigint auto_increment
        primary key,
    product_id          bigint           not null comment '商品id',
    product_sku_id      bigint           not null comment '商品库存id',
    member_id           bigint           not null comment '会员id',
    quantity            int              not null comment '购买数量',
    price               decimal(10, 2)   not null comment '添加到购物车的价格',
    product_pic         varchar(1000)    not null comment '商品主图',
    product_name        varchar(500)     not null comment '商品名称',
    product_sub_title   varchar(500)     not null comment '商品副标题（卖点）',
    product_sku_code    varchar(200)     not null comment '商品sku条码',
    member_nickname     varchar(500)     not null comment '会员昵称',
    create_date         datetime         not null comment '创建时间',
    modify_date         datetime         not null comment '修改时间',
    delete_status       int(1) default 0 not null comment '是否删除',
    product_category_id bigint           not null comment '商品分类',
    product_brand       varchar(200)     not null comment '商品品牌',
    product_sn          varchar(200)     not null comment '货号',
    product_attr        varchar(500)     not null comment '商品销售属性:[{"key":"颜色","value":"颜色"},{"key":"容量","value":"4G"}]'
)
    comment '购物车表';
