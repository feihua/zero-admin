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

INSERT INTO gozero.pms_member_price (id, product_id, member_level_id, member_price, member_level_name) VALUES (26, 7, 1, 500.00, null);
INSERT INTO gozero.pms_member_price (id, product_id, member_level_id, member_price, member_level_name) VALUES (27, 8, 1, 500.00, null);
INSERT INTO gozero.pms_member_price (id, product_id, member_level_id, member_price, member_level_name) VALUES (28, 9, 1, 500.00, null);
INSERT INTO gozero.pms_member_price (id, product_id, member_level_id, member_price, member_level_name) VALUES (29, 10, 1, 500.00, null);
INSERT INTO gozero.pms_member_price (id, product_id, member_level_id, member_price, member_level_name) VALUES (30, 11, 1, 500.00, null);
INSERT INTO gozero.pms_member_price (id, product_id, member_level_id, member_price, member_level_name) VALUES (31, 12, 1, 500.00, null);
INSERT INTO gozero.pms_member_price (id, product_id, member_level_id, member_price, member_level_name) VALUES (32, 13, 1, 500.00, null);
INSERT INTO gozero.pms_member_price (id, product_id, member_level_id, member_price, member_level_name) VALUES (33, 14, 1, 500.00, null);
INSERT INTO gozero.pms_member_price (id, product_id, member_level_id, member_price, member_level_name) VALUES (37, 18, 1, 500.00, null);
INSERT INTO gozero.pms_member_price (id, product_id, member_level_id, member_price, member_level_name) VALUES (44, 7, 2, 480.00, null);
INSERT INTO gozero.pms_member_price (id, product_id, member_level_id, member_price, member_level_name) VALUES (45, 7, 3, 450.00, null);
INSERT INTO gozero.pms_member_price (id, product_id, member_level_id, member_price, member_level_name) VALUES (52, 22, 1, null, null);
INSERT INTO gozero.pms_member_price (id, product_id, member_level_id, member_price, member_level_name) VALUES (53, 22, 2, null, null);
INSERT INTO gozero.pms_member_price (id, product_id, member_level_id, member_price, member_level_name) VALUES (54, 22, 3, null, null);
INSERT INTO gozero.pms_member_price (id, product_id, member_level_id, member_price, member_level_name) VALUES (58, 24, 1, null, null);
INSERT INTO gozero.pms_member_price (id, product_id, member_level_id, member_price, member_level_name) VALUES (59, 24, 2, null, null);
INSERT INTO gozero.pms_member_price (id, product_id, member_level_id, member_price, member_level_name) VALUES (60, 24, 3, null, null);
INSERT INTO gozero.pms_member_price (id, product_id, member_level_id, member_price, member_level_name) VALUES (112, 23, 1, 88.00, '黄金会员');
INSERT INTO gozero.pms_member_price (id, product_id, member_level_id, member_price, member_level_name) VALUES (113, 23, 2, 88.00, '白金会员');
INSERT INTO gozero.pms_member_price (id, product_id, member_level_id, member_price, member_level_name) VALUES (114, 23, 3, 66.00, '钻石会员');
INSERT INTO gozero.pms_member_price (id, product_id, member_level_id, member_price, member_level_name) VALUES (142, 31, 1, null, '黄金会员');
INSERT INTO gozero.pms_member_price (id, product_id, member_level_id, member_price, member_level_name) VALUES (143, 31, 2, null, '白金会员');
INSERT INTO gozero.pms_member_price (id, product_id, member_level_id, member_price, member_level_name) VALUES (144, 31, 3, null, '钻石会员');
INSERT INTO gozero.pms_member_price (id, product_id, member_level_id, member_price, member_level_name) VALUES (148, 32, 1, null, '黄金会员');
INSERT INTO gozero.pms_member_price (id, product_id, member_level_id, member_price, member_level_name) VALUES (149, 32, 2, null, '白金会员');
INSERT INTO gozero.pms_member_price (id, product_id, member_level_id, member_price, member_level_name) VALUES (150, 32, 3, null, '钻石会员');
INSERT INTO gozero.pms_member_price (id, product_id, member_level_id, member_price, member_level_name) VALUES (154, 33, 1, null, '黄金会员');
INSERT INTO gozero.pms_member_price (id, product_id, member_level_id, member_price, member_level_name) VALUES (155, 33, 2, null, '白金会员');
INSERT INTO gozero.pms_member_price (id, product_id, member_level_id, member_price, member_level_name) VALUES (156, 33, 3, null, '钻石会员');
INSERT INTO gozero.pms_member_price (id, product_id, member_level_id, member_price, member_level_name) VALUES (175, 34, 1, null, '黄金会员');
INSERT INTO gozero.pms_member_price (id, product_id, member_level_id, member_price, member_level_name) VALUES (176, 34, 2, null, '白金会员');
INSERT INTO gozero.pms_member_price (id, product_id, member_level_id, member_price, member_level_name) VALUES (177, 34, 3, null, '钻石会员');
INSERT INTO gozero.pms_member_price (id, product_id, member_level_id, member_price, member_level_name) VALUES (178, 30, 1, null, '黄金会员');
INSERT INTO gozero.pms_member_price (id, product_id, member_level_id, member_price, member_level_name) VALUES (179, 30, 2, null, '白金会员');
INSERT INTO gozero.pms_member_price (id, product_id, member_level_id, member_price, member_level_name) VALUES (180, 30, 3, null, '钻石会员');
INSERT INTO gozero.pms_member_price (id, product_id, member_level_id, member_price, member_level_name) VALUES (192, 27, 1, null, '黄金会员');
INSERT INTO gozero.pms_member_price (id, product_id, member_level_id, member_price, member_level_name) VALUES (193, 27, 2, null, '白金会员');
INSERT INTO gozero.pms_member_price (id, product_id, member_level_id, member_price, member_level_name) VALUES (194, 27, 3, null, '钻石会员');
INSERT INTO gozero.pms_member_price (id, product_id, member_level_id, member_price, member_level_name) VALUES (195, 28, 1, null, '黄金会员');
INSERT INTO gozero.pms_member_price (id, product_id, member_level_id, member_price, member_level_name) VALUES (196, 28, 2, null, '白金会员');
INSERT INTO gozero.pms_member_price (id, product_id, member_level_id, member_price, member_level_name) VALUES (197, 28, 3, null, '钻石会员');
INSERT INTO gozero.pms_member_price (id, product_id, member_level_id, member_price, member_level_name) VALUES (198, 29, 1, null, '黄金会员');
INSERT INTO gozero.pms_member_price (id, product_id, member_level_id, member_price, member_level_name) VALUES (199, 29, 2, null, '白金会员');
INSERT INTO gozero.pms_member_price (id, product_id, member_level_id, member_price, member_level_name) VALUES (200, 29, 3, null, '钻石会员');
INSERT INTO gozero.pms_member_price (id, product_id, member_level_id, member_price, member_level_name) VALUES (201, 26, 1, null, '黄金会员');
INSERT INTO gozero.pms_member_price (id, product_id, member_level_id, member_price, member_level_name) VALUES (202, 26, 2, null, '白金会员');
INSERT INTO gozero.pms_member_price (id, product_id, member_level_id, member_price, member_level_name) VALUES (203, 26, 3, null, '钻石会员');
INSERT INTO gozero.pms_member_price (id, product_id, member_level_id, member_price, member_level_name) VALUES (246, 36, 1, null, '黄金会员');
INSERT INTO gozero.pms_member_price (id, product_id, member_level_id, member_price, member_level_name) VALUES (247, 36, 2, null, '白金会员');
INSERT INTO gozero.pms_member_price (id, product_id, member_level_id, member_price, member_level_name) VALUES (248, 36, 3, null, '钻石会员');
INSERT INTO gozero.pms_member_price (id, product_id, member_level_id, member_price, member_level_name) VALUES (249, 35, 1, null, '黄金会员');
INSERT INTO gozero.pms_member_price (id, product_id, member_level_id, member_price, member_level_name) VALUES (250, 35, 2, null, '白金会员');
INSERT INTO gozero.pms_member_price (id, product_id, member_level_id, member_price, member_level_name) VALUES (251, 35, 3, null, '钻石会员');