drop table if exists pms_product_spu;
create table pms_product_spu
(
    id             bigint auto_increment
        primary key,
    name           varchar(200)                         not null comment '商品名称',
    category_id    bigint                               not null comment '商品分类ID',
    brand_id       bigint                               not null comment '品牌ID',
    unit           varchar(20)                          not null comment '单位',
    weight         decimal(10, 2)                       null comment '重量(kg)',
    keywords       varchar(255)                         null comment '关键词',
    brief          varchar(255)                         null comment '简介',
    description    text                                 null comment '详细描述',
    album_pics     varchar(1000)                        null comment '画册图片，最多8张，以逗号分割',
    main_pic       varchar(255)                         not null comment '主图',
    price_range    varchar(100)                         null comment '价格区间',
    publish_status tinyint(1) default 0                 not null comment '上架状态：0-下架，1-上架',
    sort           int        default 0                 not null comment '排序',
    sales          int        default 0                 not null comment '销量',
    create_by      bigint                               not null comment '创建人ID',
    create_time    datetime   default CURRENT_TIMESTAMP not null comment '创建时间',
    update_by      bigint                               null comment '更新人ID',
    update_time    datetime                             null on update CURRENT_TIMESTAMP comment '更新时间',
    is_deleted     tinyint    default 0                 not null comment '是否删除'
)
    comment '商品SPU表';

create index idx_brand
    on pms_product_spu (brand_id, is_deleted);

create index idx_category
    on pms_product_spu (category_id, is_deleted);

-- 插入商品SPU数据
INSERT INTO pms_product_spu (name, category_id, brand_id, unit, weight, keywords, brief, description, album_pics, main_pic, price_range, publish_status, sort, sales, create_by) VALUES
-- 手机
('iPhone 15 Pro', 4, 1, '台', 0.187, 'iPhone,苹果,手机,A17',
 'iPhone 15 Pro 支持USB-C快充，搭载A17 Pro芯片',
 'iPhone 15 Pro搭载创新的A17 Pro芯片，采用USB-C接口，支持快充。提供暗紫色、原色和银色三种颜色选择...',
 'http://example.com/album/iphone15-1.jpg,http://example.com/album/iphone15-2.jpg,http://example.com/album/iphone15-3.jpg',
 'http://example.com/main/iphone15.jpg',
 '7999-12999', 1, 1, 1000, 1),

('华为Mate60 Pro', 4, 2, '台', 0.225, '华为,Mate60,麒麟芯片',
 '华为Mate60 Pro 搭载麒麟芯片，支持卫星通讯',
 '华为Mate60 Pro采用创新的卫星通讯技术，搭载最新麒麟芯片，支持快充技术...',
 'http://example.com/album/mate60-1.jpg,http://example.com/album/mate60-2.jpg',
 'http://example.com/main/mate60.jpg',
 '6999-9999', 1, 2, 800, 1),

-- 电脑
('MacBook Pro 14', 5, 1, '台', 1.6, 'MacBook,苹果,笔记本,M2',
 'MacBook Pro 14寸 搭载M2 Pro芯片',
 'MacBook Pro 14寸版本搭载最新M2 Pro芯片，提供卓越性能和超长续航...',
 'http://example.com/album/macbook-1.jpg,http://example.com/album/macbook-2.jpg',
 'http://example.com/main/macbook.jpg',
 '14999-19999', 1, 3, 500, 1),

('联想小新Pro16', 5, 3, '台', 1.8, '联想,小新,笔记本,锐龙',
 '联想小新Pro16 搭载AMD锐龙7处理器',
 '联想小新Pro16搭载AMD锐龙7处理器，16寸大屏，高色域显示器...',
 'http://example.com/album/xiaoxin-1.jpg,http://example.com/album/xiaoxin-2.jpg',
 'http://example.com/main/xiaoxin.jpg',
 '5999-7999', 1, 4, 600, 1),

-- 家电
('美的智能空调', 10, 4, '台', 12.5, '美的,空调,智能',
 '美的智能变频空调 1.5匹',
 '美的智能变频空调，支持WiFi控制，节能省电，制冷制热双效合一...',
 'http://example.com/album/ac-1.jpg,http://example.com/album/ac-2.jpg',
 'http://example.com/main/ac.jpg',
 '2999-3999', 1, 5, 300, 1),

-- 厨房电器
('九阳破壁机', 11, 5, '台', 4.5, '九阳,破壁机,榨汁',
 '九阳破壁机 家用多功能',
 '九阳破壁机，支持热榨，可制作豆浆、果汁、养生饮品...',
 'http://example.com/album/juicer-1.jpg,http://example.com/album/juicer-2.jpg',
 'http://example.com/main/juicer.jpg',
 '799-999', 1, 6, 400, 1);