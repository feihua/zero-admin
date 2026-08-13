drop table if exists pms_product_ladder;
create table pms_product_ladder
(
    id         bigint auto_increment
        primary key,
    product_id bigint not null comment '商品id',
    count      int    not null comment '满足的商品数量',
    discount   bigint not null comment '折扣',
    price      bigint not null comment '折后价格'
)
    comment '产品阶梯价格表(只针对同商品)';


INSERT INTO pms_product_ladder (id, product_id, count, discount, price) VALUES (1, 1, 3, 0.70, 0.00);
INSERT INTO pms_product_ladder (id, product_id, count, discount, price) VALUES (2, 2, 3, 0.70, 0.00);
INSERT INTO pms_product_ladder (id, product_id, count, discount, price) VALUES (3, 3, 3, 0.70, 0.00);
INSERT INTO pms_product_ladder (id, product_id, count, discount, price) VALUES (4, 4, 3, 0.70, 0.00);
INSERT INTO pms_product_ladder (id, product_id, count, discount, price) VALUES (5, 5, 3, 0.70, 0.00);
INSERT INTO pms_product_ladder (id, product_id, count, discount, price) VALUES (6, 6, 3, 0.70, 0.00);
INSERT INTO pms_product_ladder (id, product_id, count, discount, price) VALUES (7, 7, 3, 0.70, 0.00);
INSERT INTO pms_product_ladder (id, product_id, count, discount, price) VALUES (8, 8, 3, 0.70, 0.00);
INSERT INTO pms_product_ladder (id, product_id, count, discount, price) VALUES (12, 9, 3, 0.70, 0.00);
INSERT INTO pms_product_ladder (id, product_id, count, discount, price) VALUES (14, 10, 4, 0.60, 0.00);
INSERT INTO pms_product_ladder (id, product_id, count, discount, price) VALUES (15, 11, 5, 0.50, 0.00);
INSERT INTO pms_product_ladder (id, product_id, count, discount, price) VALUES (18, 12, 0, 0.00, 0.00);
INSERT INTO pms_product_ladder (id, product_id, count, discount, price) VALUES (20, 13, 0, 0.00, 0.00);
INSERT INTO pms_product_ladder (id, product_id, count, discount, price) VALUES (38, 14, 0, 0.00, 0.00);
INSERT INTO pms_product_ladder (id, product_id, count, discount, price) VALUES (48, 15, 0, 0.00, 0.00);
INSERT INTO pms_product_ladder (id, product_id, count, discount, price) VALUES (50, 16, 0, 0.00, 0.00);
