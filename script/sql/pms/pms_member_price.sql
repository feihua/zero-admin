drop table if exists pms_member_price;
create table pms_member_price
(
    id                bigint auto_increment
        primary key,
    product_id        bigint       not null comment '商品id',
    member_level_id   bigint       not null comment '会员等级id',
    member_price      bigint       not null comment '会员价格',
    member_level_name varchar(100) not null comment '会员等级名称'
)
    comment '商品会员价格表';

INSERT INTO pms_member_price (id, product_id, member_level_id, member_price, member_level_name) VALUES (1, 1, 1, 88.00, '黄金会员');
INSERT INTO pms_member_price (id, product_id, member_level_id, member_price, member_level_name) VALUES (2, 2, 2, 88.00, '白金会员');
INSERT INTO pms_member_price (id, product_id, member_level_id, member_price, member_level_name) VALUES (3, 3, 3, 66.00, '钻石会员');
INSERT INTO pms_member_price (id, product_id, member_level_id, member_price, member_level_name) VALUES (4, 4, 1, 66.00, '黄金会员');
INSERT INTO pms_member_price (id, product_id, member_level_id, member_price, member_level_name) VALUES (5, 5, 1, 66.00, '黄金会员');
INSERT INTO pms_member_price (id, product_id, member_level_id, member_price, member_level_name) VALUES (6, 6, 1, 66.00, '黄金会员');
INSERT INTO pms_member_price (id, product_id, member_level_id, member_price, member_level_name) VALUES (7, 7, 1, 66.00, '黄金会员');
INSERT INTO pms_member_price (id, product_id, member_level_id, member_price, member_level_name) VALUES (8, 8, 1, 66.00, '黄金会员');
INSERT INTO pms_member_price (id, product_id, member_level_id, member_price, member_level_name) VALUES (9, 9, 1, 66.00, '黄金会员');
INSERT INTO pms_member_price (id, product_id, member_level_id, member_price, member_level_name) VALUES (10, 10, 1, 66.00, '黄金会员');
INSERT INTO pms_member_price (id, product_id, member_level_id, member_price, member_level_name) VALUES (11, 11, 1, 66.00, '黄金会员');
INSERT INTO pms_member_price (id, product_id, member_level_id, member_price, member_level_name) VALUES (12, 12, 1, 66.00, '黄金会员');
INSERT INTO pms_member_price (id, product_id, member_level_id, member_price, member_level_name) VALUES (13, 13, 1, 66.00, '黄金会员');
INSERT INTO pms_member_price (id, product_id, member_level_id, member_price, member_level_name) VALUES (14, 14, 1, 66.00, '黄金会员');
INSERT INTO pms_member_price (id, product_id, member_level_id, member_price, member_level_name) VALUES (15, 15, 1, 66.00, '黄金会员');
INSERT INTO pms_member_price (id, product_id, member_level_id, member_price, member_level_name) VALUES (16, 16, 1, 66.00, '黄金会员');
