drop table if exists pms_product_sku;
create table pms_product_sku
(
    id             bigint auto_increment
        primary key,
    spu_id         bigint                               not null comment 'SPU ID',
    name           varchar(200)                         not null comment 'SKU名称',
    sku_code       varchar(50)                          not null comment 'SKU编码',
    main_pic       varchar(255)                         not null comment '主图',
    album_pics     varchar(1000)                        null comment '图片集',
    price          decimal(10, 2)                       not null comment '价格',
    stock          int        default 0                 not null comment '库存',
    low_stock      int        default 0                 not null comment '预警库存',
    spec_data      json                                 not null comment '规格数据',
    weight         decimal(10, 2)                       null comment '重量(kg)',
    publish_status tinyint(1) default 0                 not null comment '上架状态：0-下架，1-上架',
    verify_status  tinyint(1) default 0                 not null comment '审核状态：0-未审核，1-审核通过，2-审核不通过',
    sort           int        default 0                 not null comment '排序',
    sales          int        default 0                 not null comment '销量',
    create_by      bigint                               not null comment '创建人ID',
    create_time    datetime   default CURRENT_TIMESTAMP not null comment '创建时间',
    update_by      bigint                               null comment '更新人ID',
    update_time    datetime                             null on update CURRENT_TIMESTAMP comment '更新时间',
    is_deleted     tinyint    default 0                 not null comment '是否删除',
    constraint uk_sku_code
        unique (sku_code, is_deleted)
)
    comment '商品SKU表';

create index idx_spu_id
    on pms_product_sku (spu_id, is_deleted);

-- 插入商品SKU数据
INSERT INTO pms_product_sku (spu_id, name, sku_code, main_pic, album_pics, price, stock, low_stock, spec_data, weight, publish_status, verify_status, sort, sales, create_by) VALUES
-- iPhone 15 Pro的不同规格组合
(1, 'iPhone 15 Pro 暗紫色 128GB 全网通版', 'IP15P-128-PURPLE',
 'http://example.com/sku/iphone15-purple-128.jpg',
 'http://example.com/sku/iphone15-p-1.jpg,http://example.com/sku/iphone15-p-2.jpg',
 7999.00, 1000, 100,
 '{"颜色": "暗紫色", "容量": "128GB", "网络版本": "全网通版"}',
 0.187, 1, 1, 1, 300, 1),

(1, 'iPhone 15 Pro 暗紫色 256GB 全网通版', 'IP15P-256-PURPLE',
 'http://example.com/sku/iphone15-purple-256.jpg',
 'http://example.com/sku/iphone15-p-1.jpg,http://example.com/sku/iphone15-p-2.jpg',
 8999.00, 800, 80,
 '{"颜色": "暗紫色", "容量": "256GB", "网络版本": "全网通版"}',
 0.187, 1, 1, 2, 200, 1),

(1, 'iPhone 15 Pro 自然色 128GB 全网通版', 'IP15P-128-NATURAL',
 'http://example.com/sku/iphone15-natural-128.jpg',
 'http://example.com/sku/iphone15-n-1.jpg,http://example.com/sku/iphone15-n-2.jpg',
 7999.00, 1200, 100,
 '{"颜色": "自然色", "容量": "128GB", "网络版本": "全网通版"}',
 0.187, 1, 1, 3, 250, 1),

(1, 'iPhone 15 Pro 自然色 256GB 全网通版', 'IP15P-256-NATURAL',
 'http://example.com/sku/iphone15-natural-256.jpg',
 'http://example.com/sku/iphone15-n-1.jpg,http://example.com/sku/iphone15-n-2.jpg',
 8999.00, 900, 80,
 '{"颜色": "自然色", "容量": "256GB", "网络版本": "全网通版"}',
 0.187, 1, 1, 4, 180, 1),

-- MacBook Pro的不同规格组合
(3, 'MacBook Pro 14 深空灰 16GB+512GB', 'MBP14-16-512-GRAY',
 'http://example.com/sku/macbook-gray-512.jpg',
 'http://example.com/sku/macbook-g-1.jpg,http://example.com/sku/macbook-g-2.jpg',
 14999.00, 500, 50,
 '{"颜色": "深空灰", "内存": "16GB", "硬盘": "512GB"}',
 1.6, 1, 1, 5, 100, 1),

(3, 'MacBook Pro 14 深空灰 16GB+1TB', 'MBP14-16-1T-GRAY',
 'http://example.com/sku/macbook-gray-1t.jpg',
 'http://example.com/sku/macbook-g-1.jpg,http://example.com/sku/macbook-g-2.jpg',
 16999.00, 300, 30,
 '{"颜色": "深空灰", "内存": "16GB", "硬盘": "1TB"}',
 1.6, 1, 1, 6, 80, 1);