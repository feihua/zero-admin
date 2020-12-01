create table pms_product_full_reduction
(
    id           bigint(11) auto_increment
        primary key,
    product_id   bigint         null,
    full_price   decimal(10, 2) null,
    reduce_price decimal(10, 2) null
)
    comment '产品满减表(只针对同商品)';

INSERT INTO gozero.pms_product_full_reduction (id, product_id, full_price, reduce_price) VALUES (1, 7, 100.00, 20.00);
INSERT INTO gozero.pms_product_full_reduction (id, product_id, full_price, reduce_price) VALUES (2, 8, 100.00, 20.00);
INSERT INTO gozero.pms_product_full_reduction (id, product_id, full_price, reduce_price) VALUES (3, 9, 100.00, 20.00);
INSERT INTO gozero.pms_product_full_reduction (id, product_id, full_price, reduce_price) VALUES (4, 10, 100.00, 20.00);
INSERT INTO gozero.pms_product_full_reduction (id, product_id, full_price, reduce_price) VALUES (5, 11, 100.00, 20.00);
INSERT INTO gozero.pms_product_full_reduction (id, product_id, full_price, reduce_price) VALUES (6, 12, 100.00, 20.00);
INSERT INTO gozero.pms_product_full_reduction (id, product_id, full_price, reduce_price) VALUES (7, 13, 100.00, 20.00);
INSERT INTO gozero.pms_product_full_reduction (id, product_id, full_price, reduce_price) VALUES (8, 14, 100.00, 20.00);
INSERT INTO gozero.pms_product_full_reduction (id, product_id, full_price, reduce_price) VALUES (9, 18, 100.00, 20.00);
INSERT INTO gozero.pms_product_full_reduction (id, product_id, full_price, reduce_price) VALUES (10, 7, 200.00, 50.00);
INSERT INTO gozero.pms_product_full_reduction (id, product_id, full_price, reduce_price) VALUES (11, 7, 300.00, 100.00);
INSERT INTO gozero.pms_product_full_reduction (id, product_id, full_price, reduce_price) VALUES (14, 22, 0.00, 0.00);
INSERT INTO gozero.pms_product_full_reduction (id, product_id, full_price, reduce_price) VALUES (16, 24, 0.00, 0.00);
INSERT INTO gozero.pms_product_full_reduction (id, product_id, full_price, reduce_price) VALUES (34, 23, 0.00, 0.00);
INSERT INTO gozero.pms_product_full_reduction (id, product_id, full_price, reduce_price) VALUES (44, 31, 0.00, 0.00);
INSERT INTO gozero.pms_product_full_reduction (id, product_id, full_price, reduce_price) VALUES (46, 32, 0.00, 0.00);
INSERT INTO gozero.pms_product_full_reduction (id, product_id, full_price, reduce_price) VALUES (48, 33, 0.00, 0.00);
INSERT INTO gozero.pms_product_full_reduction (id, product_id, full_price, reduce_price) VALUES (55, 34, 0.00, 0.00);
INSERT INTO gozero.pms_product_full_reduction (id, product_id, full_price, reduce_price) VALUES (56, 30, 0.00, 0.00);
INSERT INTO gozero.pms_product_full_reduction (id, product_id, full_price, reduce_price) VALUES (59, 27, 0.00, 0.00);
INSERT INTO gozero.pms_product_full_reduction (id, product_id, full_price, reduce_price) VALUES (60, 28, 500.00, 50.00);
INSERT INTO gozero.pms_product_full_reduction (id, product_id, full_price, reduce_price) VALUES (61, 28, 1000.00, 120.00);
INSERT INTO gozero.pms_product_full_reduction (id, product_id, full_price, reduce_price) VALUES (62, 29, 0.00, 0.00);
INSERT INTO gozero.pms_product_full_reduction (id, product_id, full_price, reduce_price) VALUES (63, 26, 0.00, 0.00);
INSERT INTO gozero.pms_product_full_reduction (id, product_id, full_price, reduce_price) VALUES (78, 36, 0.00, 0.00);
INSERT INTO gozero.pms_product_full_reduction (id, product_id, full_price, reduce_price) VALUES (79, 35, 0.00, 0.00);