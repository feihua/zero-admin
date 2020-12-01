create table sms_coupon_product_relation
(
    id           bigint auto_increment
        primary key,
    coupon_id    bigint       null,
    product_id   bigint       null,
    product_name varchar(500) null comment '商品名称',
    product_sn   varchar(200) null comment '商品编码'
)
    comment '优惠券和产品的关系表';

INSERT INTO gozero.sms_coupon_product_relation (id, coupon_id, product_id, product_name, product_sn) VALUES (9, 21, 33, '小米（MI）小米电视4A ', '4609652');