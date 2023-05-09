create table pms_sku_stock
(
    id              bigint auto_increment
        primary key,
    product_id      bigint         not null,
    sku_code        varchar(64)    not null comment 'sku编码',
    price           decimal(10, 2) not null,
    stock           int default 0  not null comment '库存',
    low_stock       int            not null comment '预警库存',
    pic             varchar(255)   not null comment '展示图片',
    sale            int            not null comment '销量',
    promotion_price decimal(10, 2) not null comment '单品促销价格',
    lock_stock      int default 0  not null comment '锁定库存',
    sp_data         varchar(500)   not null comment '商品销售属性，json格式'
)
    comment 'sku的库存';

INSERT INTO pms_sku_stock (id, product_id, sku_code, price, stock, low_stock, pic, sale, promotion_price, lock_stock, sp_data) VALUES (163, 36, '202002210036001', 100.00, 100, 25, ' ', 0, 0.00, 9, '[{"key":"颜色","value":"红色"},{"key":"尺寸","value":"38"},{"key":"风格","value":"秋季"}]');
INSERT INTO pms_sku_stock (id, product_id, sku_code, price, stock, low_stock, pic, sale, promotion_price, lock_stock, sp_data) VALUES (164, 36, '202002210036002', 120.00, 98, 20, ' ', 0, 0.00, 6, '[{"key":"颜色","value":"红色"},{"key":"尺寸","value":"38"},{"key":"风格","value":"夏季"}]');
INSERT INTO pms_sku_stock (id, product_id, sku_code, price, stock, low_stock, pic, sale, promotion_price, lock_stock, sp_data) VALUES (165, 36, '202002210036003', 100.00, 100, 20, ' ', 0, 0.00, 0, '[{"key":"颜色","value":"红色"},{"key":"尺寸","value":"39"},{"key":"风格","value":"秋季"}]');
INSERT INTO pms_sku_stock (id, product_id, sku_code, price, stock, low_stock, pic, sale, promotion_price, lock_stock, sp_data) VALUES (166, 36, '202002210036004', 100.00, 100, 20, ' ', 0, 0.00, 0, '[{"key":"颜色","value":"红色"},{"key":"尺寸","value":"39"},{"key":"风格","value":"夏季"}]');
INSERT INTO pms_sku_stock (id, product_id, sku_code, price, stock, low_stock, pic, sale, promotion_price, lock_stock, sp_data) VALUES (167, 36, '202002210036005', 100.00, 100, 20, ' ', 0, 0.00, 0, '[{"key":"颜色","value":"蓝色"},{"key":"尺寸","value":"38"},{"key":"风格","value":"秋季"}]');
INSERT INTO pms_sku_stock (id, product_id, sku_code, price, stock, low_stock, pic, sale, promotion_price, lock_stock, sp_data) VALUES (168, 36, '202002210036006', 100.00, 100, 20, ' ', 0, 0.00, 0, '[{"key":"颜色","value":"蓝色"},{"key":"尺寸","value":"38"},{"key":"风格","value":"夏季"}]');
INSERT INTO pms_sku_stock (id, product_id, sku_code, price, stock, low_stock, pic, sale, promotion_price, lock_stock, sp_data) VALUES (169, 36, '202002210036007', 100.00, 100, 20, ' ', 0, 0.00, 0, '[{"key":"颜色","value":"蓝色"},{"key":"尺寸","value":"39"},{"key":"风格","value":"秋季"}]');
INSERT INTO pms_sku_stock (id, product_id, sku_code, price, stock, low_stock, pic, sale, promotion_price, lock_stock, sp_data) VALUES (170, 36, '202002210036008', 100.00, 100, 20, ' ', 0, 0.00, 0, '[{"key":"颜色","value":"蓝色"},{"key":"尺寸","value":"39"},{"key":"风格","value":"夏季"}]');
INSERT INTO pms_sku_stock (id, product_id, sku_code, price, stock, low_stock, pic, sale, promotion_price, lock_stock, sp_data) VALUES (171, 35, '202002250035001', 200.00, 100, 50, ' ', 0, 0.00, 0, '[{"key":"颜色","value":"红色"},{"key":"尺寸","value":"38"},{"key":"风格","value":"夏季"}]');
INSERT INTO pms_sku_stock (id, product_id, sku_code, price, stock, low_stock, pic, sale, promotion_price, lock_stock, sp_data) VALUES (172, 35, '202002250035002', 240.00, 100, 50, ' ', 0, 0.00, 0, '[{"key":"颜色","value":"红色"},{"key":"尺寸","value":"38"},{"key":"风格","value":"秋季"}]');
INSERT INTO pms_sku_stock (id, product_id, sku_code, price, stock, low_stock, pic, sale, promotion_price, lock_stock, sp_data) VALUES (173, 35, '202002250035003', 200.00, 100, 50, ' ', 0, 0.00, 0, '[{"key":"颜色","value":"红色"},{"key":"尺寸","value":"39"},{"key":"风格","value":"夏季"}]');
INSERT INTO pms_sku_stock (id, product_id, sku_code, price, stock, low_stock, pic, sale, promotion_price, lock_stock, sp_data) VALUES (174, 35, '202002250035004', 200.00, 100, 50, ' ', 0, 0.00, 0, '[{"key":"颜色","value":"红色"},{"key":"尺寸","value":"39"},{"key":"风格","value":"秋季"}]');
INSERT INTO pms_sku_stock (id, product_id, sku_code, price, stock, low_stock, pic, sale, promotion_price, lock_stock, sp_data) VALUES (175, 35, '202002250035005', 200.00, 100, 50, ' ', 0, 0.00, 0, '[{"key":"颜色","value":"蓝色"},{"key":"尺寸","value":"38"},{"key":"风格","value":"夏季"}]');
INSERT INTO pms_sku_stock (id, product_id, sku_code, price, stock, low_stock, pic, sale, promotion_price, lock_stock, sp_data) VALUES (176, 35, '202002250035006', 200.00, 100, 50, ' ', 0, 0.00, 0, '[{"key":"颜色","value":"蓝色"},{"key":"尺寸","value":"38"},{"key":"风格","value":"秋季"}]');
INSERT INTO pms_sku_stock (id, product_id, sku_code, price, stock, low_stock, pic, sale, promotion_price, lock_stock, sp_data) VALUES (177, 35, '202002250035007', 200.00, 100, 50, ' ', 0, 0.00, 0, '[{"key":"颜色","value":"蓝色"},{"key":"尺寸","value":"39"},{"key":"风格","value":"夏季"}]');
INSERT INTO pms_sku_stock (id, product_id, sku_code, price, stock, low_stock, pic, sale, promotion_price, lock_stock, sp_data) VALUES (178, 35, '202002250035008', 200.00, 100, 50, ' ', 0, 0.00, 0, '[{"key":"颜色","value":"蓝色"},{"key":"尺寸","value":"39"},{"key":"风格","value":"秋季"}]');
