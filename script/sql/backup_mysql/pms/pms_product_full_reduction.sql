drop table if exists pms_product_full_reduction;
create table pms_product_full_reduction
(
    id           bigint(11) auto_increment
        primary key,
    product_id   bigint not null comment '商品id',
    full_price   bigint not null comment '商品满多少',
    reduce_price bigint not null comment '商品减多少'
)
    comment '产品满减表(只针对同商品)';


INSERT INTO pms_product_full_reduction (id, product_id, full_price, reduce_price) VALUES (1, 1, 100.00, 20.00);
INSERT INTO pms_product_full_reduction (id, product_id, full_price, reduce_price) VALUES (2, 2, 100.00, 20.00);
INSERT INTO pms_product_full_reduction (id, product_id, full_price, reduce_price) VALUES (3, 3, 100.00, 20.00);
INSERT INTO pms_product_full_reduction (id, product_id, full_price, reduce_price) VALUES (4, 4, 100.00, 20.00);
INSERT INTO pms_product_full_reduction (id, product_id, full_price, reduce_price) VALUES (5, 5, 100.00, 20.00);
INSERT INTO pms_product_full_reduction (id, product_id, full_price, reduce_price) VALUES (6, 6, 100.00, 20.00);
INSERT INTO pms_product_full_reduction (id, product_id, full_price, reduce_price) VALUES (7, 7, 100.00, 20.00);
INSERT INTO pms_product_full_reduction (id, product_id, full_price, reduce_price) VALUES (8, 8, 100.00, 20.00);
INSERT INTO pms_product_full_reduction (id, product_id, full_price, reduce_price) VALUES (9, 9, 100.00, 20.00);
INSERT INTO pms_product_full_reduction (id, product_id, full_price, reduce_price) VALUES (10, 10, 200.00, 50.00);
INSERT INTO pms_product_full_reduction (id, product_id, full_price, reduce_price) VALUES (11, 11, 300.00, 100.00);
INSERT INTO pms_product_full_reduction (id, product_id, full_price, reduce_price) VALUES (14, 12, 0.00, 0.00);
INSERT INTO pms_product_full_reduction (id, product_id, full_price, reduce_price) VALUES (16, 13, 0.00, 0.00);
INSERT INTO pms_product_full_reduction (id, product_id, full_price, reduce_price) VALUES (34, 14, 0.00, 0.00);
INSERT INTO pms_product_full_reduction (id, product_id, full_price, reduce_price) VALUES (44, 15, 0.00, 0.00);
INSERT INTO pms_product_full_reduction (id, product_id, full_price, reduce_price) VALUES (46, 16, 0.00, 0.00);
