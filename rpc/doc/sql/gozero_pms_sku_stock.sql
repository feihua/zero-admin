create table pms_sku_stock
(
    id              bigint auto_increment
        primary key,
    product_id      bigint         null,
    sku_code        varchar(64)    not null comment 'sku编码',
    price           decimal(10, 2) null,
    stock           int default 0  null comment '库存',
    low_stock       int            null comment '预警库存',
    pic             varchar(255)   null comment '展示图片',
    sale            int            null comment '销量',
    promotion_price decimal(10, 2) null comment '单品促销价格',
    lock_stock      int default 0  null comment '锁定库存',
    sp_data         varchar(500)   null comment '商品销售属性，json格式'
)
    comment 'sku的库存';

INSERT INTO gozero.pms_sku_stock (id, product_id, sku_code, price, stock, low_stock, pic, sale, promotion_price, lock_stock, sp_data) VALUES (98, 27, '201808270027001', 2699.00, 97, ' ', ' ', ' ', ' ', -24, '[{"key":"颜色","value":"黑色"},{"key":"容量","value":"32G"}]');
INSERT INTO gozero.pms_sku_stock (id, product_id, sku_code, price, stock, low_stock, pic, sale, promotion_price, lock_stock, sp_data) VALUES (99, 27, '201808270027002', 2999.00, 100, ' ', ' ', ' ', ' ', 0, '[{"key":"颜色","value":"黑色"},{"key":"容量","value":"64G"}]');
INSERT INTO gozero.pms_sku_stock (id, product_id, sku_code, price, stock, low_stock, pic, sale, promotion_price, lock_stock, sp_data) VALUES (100, 27, '201808270027003', 2699.00, 100, ' ', ' ', ' ', ' ', 0, '[{"key":"颜色","value":"蓝色"},{"key":"容量","value":"32G"}]');
INSERT INTO gozero.pms_sku_stock (id, product_id, sku_code, price, stock, low_stock, pic, sale, promotion_price, lock_stock, sp_data) VALUES (101, 27, '201808270027004', 2999.00, 100, ' ', ' ', ' ', ' ', 0, '[{"key":"颜色","value":"蓝色"},{"key":"容量","value":"64G"}]');
INSERT INTO gozero.pms_sku_stock (id, product_id, sku_code, price, stock, low_stock, pic, sale, promotion_price, lock_stock, sp_data) VALUES (102, 28, '201808270028001', 649.00, 99, ' ', ' ', ' ', ' ', -8, '[{"key":"颜色","value":"金色"},{"key":"容量","value":"16G"}]');
INSERT INTO gozero.pms_sku_stock (id, product_id, sku_code, price, stock, low_stock, pic, sale, promotion_price, lock_stock, sp_data) VALUES (103, 28, '201808270028002', 699.00, 99, ' ', ' ', ' ', ' ', -8, '[{"key":"颜色","value":"金色"},{"key":"容量","value":"32G"}]');
INSERT INTO gozero.pms_sku_stock (id, product_id, sku_code, price, stock, low_stock, pic, sale, promotion_price, lock_stock, sp_data) VALUES (104, 28, '201808270028003', 649.00, 100, ' ', ' ', ' ', ' ', 0, '[{"key":"颜色","value":"银色"},{"key":"容量","value":"16G"}]');
INSERT INTO gozero.pms_sku_stock (id, product_id, sku_code, price, stock, low_stock, pic, sale, promotion_price, lock_stock, sp_data) VALUES (105, 28, '201808270028004', 699.00, 100, ' ', ' ', ' ', ' ', 0, '[{"key":"颜色","value":"银色"},{"key":"容量","value":"32G"}]');
INSERT INTO gozero.pms_sku_stock (id, product_id, sku_code, price, stock, low_stock, pic, sale, promotion_price, lock_stock, sp_data) VALUES (106, 29, '201808270029001', 5499.00, 99, ' ', ' ', ' ', ' ', -8, '[{"key":"颜色","value":"金色"},{"key":"容量","value":"32G"}]');
INSERT INTO gozero.pms_sku_stock (id, product_id, sku_code, price, stock, low_stock, pic, sale, promotion_price, lock_stock, sp_data) VALUES (107, 29, '201808270029002', 6299.00, 100, ' ', ' ', ' ', ' ', 0, '[{"key":"颜色","value":"金色"},{"key":"容量","value":"64G"}]');
INSERT INTO gozero.pms_sku_stock (id, product_id, sku_code, price, stock, low_stock, pic, sale, promotion_price, lock_stock, sp_data) VALUES (108, 29, '201808270029003', 5499.00, 100, ' ', ' ', ' ', ' ', 0, '[{"key":"颜色","value":"银色"},{"key":"容量","value":"32G"}]');
INSERT INTO gozero.pms_sku_stock (id, product_id, sku_code, price, stock, low_stock, pic, sale, promotion_price, lock_stock, sp_data) VALUES (109, 29, '201808270029004', 6299.00, 100, ' ', ' ', ' ', ' ', 0, '[{"key":"颜色","value":"银色"},{"key":"容量","value":"64G"}]');
INSERT INTO gozero.pms_sku_stock (id, product_id, sku_code, price, stock, low_stock, pic, sale, promotion_price, lock_stock, sp_data) VALUES (110, 26, '201806070026001', 3788.00, 499, ' ', ' ', ' ', ' ', 0, '[{"key":"颜色","value":"金色"},{"key":"容量","value":"16G"}]');
INSERT INTO gozero.pms_sku_stock (id, product_id, sku_code, price, stock, low_stock, pic, sale, promotion_price, lock_stock, sp_data) VALUES (111, 26, '201806070026002', 3999.00, 500, ' ', ' ', ' ', ' ', 0, '[{"key":"颜色","value":"金色"},{"key":"容量","value":"32G"}]');
INSERT INTO gozero.pms_sku_stock (id, product_id, sku_code, price, stock, low_stock, pic, sale, promotion_price, lock_stock, sp_data) VALUES (112, 26, '201806070026003', 3788.00, 500, ' ', ' ', ' ', ' ', 0, '[{"key":"颜色","value":"银色"},{"key":"容量","value":"16G"}]');
INSERT INTO gozero.pms_sku_stock (id, product_id, sku_code, price, stock, low_stock, pic, sale, promotion_price, lock_stock, sp_data) VALUES (113, 26, '201806070026004', 3999.00, 500, ' ', ' ', ' ', ' ', 0, '[{"key":"颜色","value":"银色"},{"key":"容量","value":"32G"}]');
INSERT INTO gozero.pms_sku_stock (id, product_id, sku_code, price, stock, low_stock, pic, sale, promotion_price, lock_stock, sp_data) VALUES (163, 36, '202002210036001', 100.00, 100, 25, ' ', ' ', ' ', 9, '[{"key":"颜色","value":"红色"},{"key":"尺寸","value":"38"},{"key":"风格","value":"秋季"}]');
INSERT INTO gozero.pms_sku_stock (id, product_id, sku_code, price, stock, low_stock, pic, sale, promotion_price, lock_stock, sp_data) VALUES (164, 36, '202002210036002', 120.00, 98, 20, ' ', ' ', ' ', 6, '[{"key":"颜色","value":"红色"},{"key":"尺寸","value":"38"},{"key":"风格","value":"夏季"}]');
INSERT INTO gozero.pms_sku_stock (id, product_id, sku_code, price, stock, low_stock, pic, sale, promotion_price, lock_stock, sp_data) VALUES (165, 36, '202002210036003', 100.00, 100, 20, ' ', ' ', ' ', 0, '[{"key":"颜色","value":"红色"},{"key":"尺寸","value":"39"},{"key":"风格","value":"秋季"}]');
INSERT INTO gozero.pms_sku_stock (id, product_id, sku_code, price, stock, low_stock, pic, sale, promotion_price, lock_stock, sp_data) VALUES (166, 36, '202002210036004', 100.00, 100, 20, ' ', ' ', ' ', 0, '[{"key":"颜色","value":"红色"},{"key":"尺寸","value":"39"},{"key":"风格","value":"夏季"}]');
INSERT INTO gozero.pms_sku_stock (id, product_id, sku_code, price, stock, low_stock, pic, sale, promotion_price, lock_stock, sp_data) VALUES (167, 36, '202002210036005', 100.00, 100, 20, ' ', ' ', ' ', 0, '[{"key":"颜色","value":"蓝色"},{"key":"尺寸","value":"38"},{"key":"风格","value":"秋季"}]');
INSERT INTO gozero.pms_sku_stock (id, product_id, sku_code, price, stock, low_stock, pic, sale, promotion_price, lock_stock, sp_data) VALUES (168, 36, '202002210036006', 100.00, 100, 20, ' ', ' ', ' ', 0, '[{"key":"颜色","value":"蓝色"},{"key":"尺寸","value":"38"},{"key":"风格","value":"夏季"}]');
INSERT INTO gozero.pms_sku_stock (id, product_id, sku_code, price, stock, low_stock, pic, sale, promotion_price, lock_stock, sp_data) VALUES (169, 36, '202002210036007', 100.00, 100, 20, ' ', ' ', ' ', 0, '[{"key":"颜色","value":"蓝色"},{"key":"尺寸","value":"39"},{"key":"风格","value":"秋季"}]');
INSERT INTO gozero.pms_sku_stock (id, product_id, sku_code, price, stock, low_stock, pic, sale, promotion_price, lock_stock, sp_data) VALUES (170, 36, '202002210036008', 100.00, 100, 20, ' ', ' ', ' ', 0, '[{"key":"颜色","value":"蓝色"},{"key":"尺寸","value":"39"},{"key":"风格","value":"夏季"}]');
INSERT INTO gozero.pms_sku_stock (id, product_id, sku_code, price, stock, low_stock, pic, sale, promotion_price, lock_stock, sp_data) VALUES (171, 35, '202002250035001', 200.00, 100, 50, ' ', ' ', ' ', 0, '[{"key":"颜色","value":"红色"},{"key":"尺寸","value":"38"},{"key":"风格","value":"夏季"}]');
INSERT INTO gozero.pms_sku_stock (id, product_id, sku_code, price, stock, low_stock, pic, sale, promotion_price, lock_stock, sp_data) VALUES (172, 35, '202002250035002', 240.00, 100, 50, ' ', ' ', ' ', 0, '[{"key":"颜色","value":"红色"},{"key":"尺寸","value":"38"},{"key":"风格","value":"秋季"}]');
INSERT INTO gozero.pms_sku_stock (id, product_id, sku_code, price, stock, low_stock, pic, sale, promotion_price, lock_stock, sp_data) VALUES (173, 35, '202002250035003', 200.00, 100, 50, ' ', ' ', ' ', 0, '[{"key":"颜色","value":"红色"},{"key":"尺寸","value":"39"},{"key":"风格","value":"夏季"}]');
INSERT INTO gozero.pms_sku_stock (id, product_id, sku_code, price, stock, low_stock, pic, sale, promotion_price, lock_stock, sp_data) VALUES (174, 35, '202002250035004', 200.00, 100, 50, ' ', ' ', ' ', 0, '[{"key":"颜色","value":"红色"},{"key":"尺寸","value":"39"},{"key":"风格","value":"秋季"}]');
INSERT INTO gozero.pms_sku_stock (id, product_id, sku_code, price, stock, low_stock, pic, sale, promotion_price, lock_stock, sp_data) VALUES (175, 35, '202002250035005', 200.00, 100, 50, ' ', ' ', ' ', 0, '[{"key":"颜色","value":"蓝色"},{"key":"尺寸","value":"38"},{"key":"风格","value":"夏季"}]');
INSERT INTO gozero.pms_sku_stock (id, product_id, sku_code, price, stock, low_stock, pic, sale, promotion_price, lock_stock, sp_data) VALUES (176, 35, '202002250035006', 200.00, 100, 50, ' ', ' ', ' ', 0, '[{"key":"颜色","value":"蓝色"},{"key":"尺寸","value":"38"},{"key":"风格","value":"秋季"}]');
INSERT INTO gozero.pms_sku_stock (id, product_id, sku_code, price, stock, low_stock, pic, sale, promotion_price, lock_stock, sp_data) VALUES (177, 35, '202002250035007', 200.00, 100, 50, ' ', ' ', ' ', 0, '[{"key":"颜色","value":"蓝色"},{"key":"尺寸","value":"39"},{"key":"风格","value":"夏季"}]');
INSERT INTO gozero.pms_sku_stock (id, product_id, sku_code, price, stock, low_stock, pic, sale, promotion_price, lock_stock, sp_data) VALUES (178, 35, '202002250035008', 200.00, 100, 50, ' ', ' ', ' ', 0, '[{"key":"颜色","value":"蓝色"},{"key":"尺寸","value":"39"},{"key":"风格","value":"秋季"}]');