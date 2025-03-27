drop table if exists oms_cart_item;
create table oms_cart_item
(
    id                  bigint auto_increment
        primary key,
    product_id          bigint                             not null comment '商品id',
    product_sku_id      bigint                             not null comment '商品库存id',
    member_id           bigint                             not null comment '会员id',
    quantity            int                                not null comment '购买数量',
    price               bigint                             not null comment '添加到购物车的价格',
    product_pic         varchar(1000)                      not null comment '商品主图',
    product_name        varchar(500)                       not null comment '商品名称',
    product_sub_title   varchar(500)                       not null comment '商品副标题（卖点）',
    product_sku_code    varchar(200)                       not null comment '商品sku条码',
    member_nickname     varchar(500)                       not null comment '会员昵称',
    delete_status       tinyint  default 0                 not null comment '是否删除',
    product_category_id bigint                             not null comment '商品分类',
    product_brand       varchar(200)                       not null comment '商品品牌',
    product_sn          varchar(200)                       not null comment '货号',
    product_attr        varchar(500)                       not null comment '商品销售属性:[{"key":"颜色","value":"颜色"},{"key":"容量","value":"4G"}]',
    create_time         datetime default CURRENT_TIMESTAMP not null comment '创建时间',
    update_time         datetime                           null on update CURRENT_TIMESTAMP comment '更新时间'
)
    comment '购物车';

INSERT INTO oms_cart_item (product_id, product_sku_id, member_id, quantity, price, product_pic, product_name, product_sub_title, product_sku_code, member_nickname, delete_status, product_category_id, product_brand, product_sn, product_attr) VALUES (27, 163, 2, 1, 2999, 'http://macro-oss.oss-cn-shenzhen.aliyuncs.com/mall/images/20221104/xiaomi_12_pro_01.jpg', '小米12 Pro 天玑版 天玑9000+处理器 5000万疾速影像 2K超视感屏 120Hz高刷 67W快充', '天玑9000+处理器、5160mAh大电量、2KAmoled超视感屏【点击购买小米11Ultra，戳】 ', '202211040040001', '', 0, 19, '小米', '100027789721', '[{"key":"颜色","value":"黑色"},{"key":"容量","value":"128G"}]');
INSERT INTO oms_cart_item (product_id, product_sku_id, member_id, quantity, price, product_pic, product_name, product_sub_title, product_sku_code, member_nickname, delete_status, product_category_id, product_brand, product_sn, product_attr) VALUES (27, 163, 1, 3, 2999, 'http://macro-oss.oss-cn-shenzhen.aliyuncs.com/mall/images/20221104/xiaomi_12_pro_01.jpg', '小米12 Pro 天玑版 天玑9000+处理器 5000万疾速影像 2K超视感屏 120Hz高刷 67W快充', '天玑9000+处理器、5160mAh大电量、2KAmoled超视感屏【点击购买小米11Ultra，戳】 ', '202211040040001', '', 0, 19, '小米', '100027789721', '[{"key":"颜色","value":"黑色"},{"key":"容量","value":"128G"}]');
