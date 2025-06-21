drop table if exists pms_product_sku;
create table pms_product_sku
(
    id                   bigint auto_increment
        primary key comment '商品SkuId',
    spu_id               bigint                             not null comment '商品SpuId',
    name                 varchar(200)                       not null comment 'SKU名称',
    sku_code             varchar(50)                        not null comment 'SKU编码',
    main_pic             varchar(255)                       not null comment '主图',
    album_pics           varchar(1000)                      not null comment '图片集',
    price                decimal(10, 2)                     not null comment '价格',
    promotion_price      decimal(10, 2)                     not null comment '单品促销价格',
    promotion_start_time datetime                           null comment '促销开始时间',
    promotion_end_time   datetime                           null comment '促销结束时间',
    stock                int                                not null comment '库存',
    low_stock            int                                not null comment '预警库存',
    spec_data            json                               not null comment '规格数据',
    weight               decimal(10, 2)                     not null comment '重量(kg)',
    publish_status       tinyint  default 0                 not null comment '上架状态：0-下架，1-上架',
    verify_status        tinyint  default 0                 not null comment '审核状态：0-未审核，1-审核通过，2-审核不通过',
    sort                 int      default 1                 not null comment '排序',
    sales                int      default 0                 not null comment '销量',
    create_by            bigint                             not null comment '创建人ID',
    create_time          datetime default CURRENT_TIMESTAMP not null comment '创建时间',
    update_by            bigint                             null comment '更新人ID',
    update_time          datetime                           null on update CURRENT_TIMESTAMP comment '更新时间',
    is_deleted           tinyint  default 0                 not null comment '是否删除',
    constraint uk_sku_code
        unique (sku_code, is_deleted)
)
    comment '商品SKU表';

create index idx_spu_id
    on pms_product_sku (spu_id, is_deleted);

-- 插入商品SKU数据
insert into pms_product_sku (id, spu_id, name, sku_code, main_pic, album_pics, price, promotion_price, stock, low_stock, spec_data, weight, publish_status, verify_status, sort, sales, create_by)
values  (1, 1, 'iPhone 15 Pro 暗紫色 128GB 全网通版', 'IP15P-128-PURPLE', 'http://129.204.203.29/big.png', 'http://example.com/sku/iphone15-p-1.jpg,http://example.com/sku/iphone15-p-2.jpg', 7999.00, 7899.00, 1000, 100, '{"容量": "128GB", "颜色": "暗紫色", "网络版本": "全网通版"}', 0.19, 1, 1, 1, 300, 1),
        (2, 1, 'iPhone 15 Pro 暗紫色 256GB 全网通版', 'IP15P-256-PURPLE', 'http://129.204.203.29/big.png', 'http://example.com/sku/iphone15-p-1.jpg,http://example.com/sku/iphone15-p-2.jpg', 8999.00, 8899.00, 800, 80, '{"容量": "256GB", "颜色": "暗紫色", "网络版本": "全网通版"}', 0.19, 1, 1, 2, 200, 1),
        (3, 1, 'iPhone 15 Pro 自然色 128GB 全网通版', 'IP15P-128-NATURAL', 'http://129.204.203.29/big.png', 'http://example.com/sku/iphone15-n-1.jpg,http://example.com/sku/iphone15-n-2.jpg', 7999.00, 7899.00, 1200, 100, '{"容量": "128GB", "颜色": "自然色", "网络版本": "全网通版"}', 0.19, 1, 1, 3, 250, 1),
        (4, 1, 'iPhone 15 Pro 自然色 256GB 全网通版', 'IP15P-256-NATURAL', 'http://129.204.203.29/big.png', 'http://example.com/sku/iphone15-n-1.jpg,http://example.com/sku/iphone15-n-2.jpg', 8999.00, 8899.00, 900, 80, '{"容量": "256GB", "颜色": "自然色", "网络版本": "全网通版"}', 0.19, 1, 1, 4, 180, 1),
        (5, 3, 'MacBook Pro 14 深空灰 16GB+512GB', 'MBP14-16-512-GRAY', 'http://129.204.203.29/big.png', 'http://example.com/sku/macbook-g-1.jpg,http://example.com/sku/macbook-g-2.jpg', 14999.00, 13999.00, 500, 50, '{"内存": "16GB", "硬盘": "512GB", "颜色": "深空灰"}', 1.60, 1, 1, 5, 100, 1),
        (6, 3, 'MacBook Pro 14 深空灰 16GB+1TB', 'MBP14-16-1T-GRAY', 'http://129.204.203.29/big.png', 'http://example.com/sku/macbook-g-1.jpg,http://example.com/sku/macbook-g-2.jpg', 16999.00, 15999.00, 300, 30, '{"内存": "16GB", "硬盘": "1TB", "颜色": "深空灰"}', 1.60, 1, 1, 6, 80, 1);
