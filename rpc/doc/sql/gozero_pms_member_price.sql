create table pms_member_price
(
    id                bigint auto_increment
        primary key,
    product_id        bigint         null,
    member_level_id   bigint         null,
    member_price      decimal(10, 2) null comment '会员价格',
    member_level_name varchar(100)   null
)
    comment '商品会员价格表';


INSERT INTO gozero.pms_member_price (id, product_id, member_level_id, member_price, member_level_name) VALUES (112, 23, 1, 88.00, '黄金会员');
INSERT INTO gozero.pms_member_price (id, product_id, member_level_id, member_price, member_level_name) VALUES (113, 23, 2, 88.00, '白金会员');
INSERT INTO gozero.pms_member_price (id, product_id, member_level_id, member_price, member_level_name) VALUES (114, 23, 3, 66.00, '钻石会员');
INSERT INTO gozero.pms_member_price (id, product_id, member_level_id, member_price, member_level_name) VALUES (142, 31, 1, 66.00, '黄金会员');
