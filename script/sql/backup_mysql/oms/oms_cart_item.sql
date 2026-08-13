DROP TABLE IF EXISTS oms_cart_item;
create table oms_cart_item
(
    id                  bigint auto_increment comment '主键ID'
        primary key,
    member_id           bigint                                 not null comment '会员ID',
    product_id          bigint                                 not null comment '商品ID',
    product_sku_id      bigint                                 not null comment '商品SKU ID',
    quantity            int          default 1                 not null comment '购买数量',
    price               decimal(10, 2)                         not null comment '添加到购物车时的价格',
    selected            tinyint      default 1                 not null comment '是否选中 0-未选中 1-选中',
    product_name        varchar(500)                           not null comment '商品名称',
    product_sub_title   varchar(500) default ''                not null comment '商品副标题',
    product_pic         text                                   not null comment '商品主图URL',
    product_sku_code    varchar(200)                           not null comment '商品SKU编码',
    product_sn          varchar(200)                           not null comment '商品货号',
    product_brand       varchar(200) default ''                not null comment '商品品牌',
    product_category_id bigint                                 not null comment '商品分类ID',
    product_attr        json                                   not null comment '商品销售属性JSON',
    member_nickname     varchar(500)                           not null comment '会员昵称',
    source              tinyint      default 1                 not null comment '来源 1-PC 2-H5 3-小程序 4-APP',
    delete_status       tinyint      default 0                 not null comment '删除状态 0-正常 1-删除',
    expire_time         datetime                               not null comment '过期时间',
    create_time         datetime     default CURRENT_TIMESTAMP not null comment '创建时间',
    update_time         datetime                               null on update CURRENT_TIMESTAMP comment '更新时间',
    constraint uk_member_sku_status
        unique (member_id, product_sku_id, delete_status)
)
    comment '购物车表' collate = utf8mb4_unicode_ci;

create index idx_create_time
    on oms_cart_item (create_time);

create index idx_expire_time
    on oms_cart_item (expire_time);

create index idx_member_status
    on oms_cart_item (member_id, delete_status);

create index idx_product_id
    on oms_cart_item (product_id);

create index idx_product_sku_id
    on oms_cart_item (product_sku_id);

create index idx_update_time
    on oms_cart_item (update_time);

