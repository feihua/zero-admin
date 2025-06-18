drop table if exists pms_product_spu;
create table pms_product_spu
(
    id                    bigint auto_increment
        primary key comment '商品SpuId',
    name                  varchar(200)                           not null comment '商品名称',
    category_id           bigint                                 not null comment '商品分类ID',
    category_ids          varchar(20)                            not null comment '商品分类ID集合',
    category_name         varchar(255)                           not null comment '商品分类名称',
    brand_id              bigint                                 not null comment '品牌ID',
    brand_name            varchar(255)                           not null comment '品牌名称',
    unit                  varchar(20)                            not null comment '单位',
    weight                decimal(10, 2)                         not null comment '重量(kg)',
    keywords              varchar(255) default ''                not null comment '关键词',
    brief                 varchar(255)                           not null comment '简介',
    description           text                                   not null comment '详细描述',
    album_pics            varchar(1000)                          not null comment '画册图片，最多8张，以逗号分割',
    main_pic              varchar(255)                           not null comment '主图',
    price_range           varchar(100)                           not null comment '价格区间',
    publish_status        tinyint      default 0                 not null comment '上架状态：0-下架，1-上架',
    new_status            tinyint      default 0                 not null comment '新品状态:0->不是新品；1->新品',
    recommend_status      tinyint      default 0                 not null comment '推荐状态；0->不推荐；1->推荐',
    verify_status         tinyint      default 0                 not null comment '审核状态：0->未审核；1->审核通过',
    preview_status        tinyint      default 0                 not null comment '是否为预告商品：0->不是；1->是',
    sort                  int          default 1                 not null comment '排序',
    new_status_sort       int          default 1                 not null comment '新品排序',
    recommend_status_sort int          default 1                 not null comment '推荐排序',
    sales                 int          default 0                 not null comment '销量',
    stock                 int                                    not null comment '库存',
    low_stock             int                                    not null comment '预警库存',
    promotion_type        tinyint      default 0                 not null comment '促销类型：0->没有促销使用原价;1->使用促销价；2->使用会员价；3->使用阶梯价格；4->使用满减价格；5->秒杀',
    detail_title          varchar(255)                           not null comment '详情标题',
    detail_desc           text                                   not null comment '详情描述',
    detail_html           text                                   not null comment '产品详情网页内容',
    detail_mobile_html    text                                   not null comment '移动端网页详情',
    create_by             bigint                                 not null comment '创建人ID',
    create_time           datetime     default CURRENT_TIMESTAMP not null comment '创建时间',
    update_by             bigint                                 null comment '更新人ID',
    update_time           datetime                               null on update CURRENT_TIMESTAMP comment '更新时间',
    is_deleted            tinyint      default 0                 not null comment '是否删除'
)
    comment '商品SPU表';

create index idx_brand
    on pms_product_spu (brand_id, is_deleted);

create index idx_category
    on pms_product_spu (category_id, is_deleted);

-- 插入商品SPU数据
insert into pms_product_spu (id, name, category_id, category_ids, category_name, brand_id, brand_name, unit, weight, keywords, brief, description, album_pics, main_pic, price_range, publish_status, new_status, recommend_status, verify_status, preview_status, sort, sales, stock, low_stock, promotion_type, detail_title, detail_desc, detail_html, detail_mobile_html, create_by)
values  (1, 'iPhone 15 Pro', 4, '', 'test', 1, 'test', '台', 0.19, 'iPhone,苹果,手机,A17', 'iPhone 15 Pro 支持USB-C快充，搭载A17 Pro芯片', 'iPhone 15 Pro搭载创新的A17 Pro芯片，采用USB-C接口，支持快充。提供暗紫色、原色和银色三种颜色选择...', 'http://example.com/album/iphone15-1.jpg,http://example.com/album/iphone15-2.jpg,http://example.com/album/iphone15-3.jpg', 'http://example.com/main/iphone15.jpg', '7999-12999', 1, 0, 0, 0, 0, 1, 1000, 0, 0, 0, 'test', 'test', 'test', 'test', 1),
        (2, '华为Mate60 Pro', 4, '', 'test', 2, 'test', '台', 0.23, '华为,Mate60,麒麟芯片', '华为Mate60 Pro 搭载麒麟芯片，支持卫星通讯', '华为Mate60 Pro采用创新的卫星通讯技术，搭载最新麒麟芯片，支持快充技术...', 'http://example.com/album/mate60-1.jpg,http://example.com/album/mate60-2.jpg', 'http://example.com/main/mate60.jpg', '6999-9999', 1, 0, 0, 0, 0, 2, 800, 0, 0, 0, 'test', 'test', 'test', 'test', 1),
        (3, 'MacBook Pro 14', 5, '', 'test', 1, 'test', '台', 1.60, 'MacBook,苹果,笔记本,M2', 'MacBook Pro 14寸 搭载M2 Pro芯片', 'MacBook Pro 14寸版本搭载最新M2 Pro芯片，提供卓越性能和超长续航...', 'http://example.com/album/macbook-1.jpg,http://example.com/album/macbook-2.jpg', 'http://example.com/main/macbook.jpg', '14999-19999', 1, 0, 0, 0, 0, 3, 500, 0, 0, 0, 'test', 'test', 'test', 'test', 1),
        (4, '联想小新Pro16', 5, '', 'test', 3, 'test', '台', 1.80, '联想,小新,笔记本,锐龙', '联想小新Pro16 搭载AMD锐龙7处理器', '联想小新Pro16搭载AMD锐龙7处理器，16寸大屏，高色域显示器...', 'http://example.com/album/xiaoxin-1.jpg,http://example.com/album/xiaoxin-2.jpg', 'http://example.com/main/xiaoxin.jpg', '5999-7999', 1, 0, 0, 0, 0, 4, 600, 0, 0, 0, 'test', 'test', 'test', 'test', 1),
        (5, '美的智能空调', 10, '', 'test', 4, 'test', '台', 12.50, '美的,空调,智能', '美的智能变频空调 1.5匹', '美的智能变频空调，支持WiFi控制，节能省电，制冷制热双效合一...', 'http://example.com/album/ac-1.jpg,http://example.com/album/ac-2.jpg', 'http://example.com/main/ac.jpg', '2999-3999', 1, 0, 0, 0, 0, 5, 300, 0, 0, 0, 'test', 'test', 'test', 'test', 1),
        (6, '九阳破壁机', 11, '', 'test', 5, 'test', '台', 4.50, '九阳,破壁机,榨汁', '九阳破壁机 家用多功能', '九阳破壁机，支持热榨，可制作豆浆、果汁、养生饮品...', 'http://example.com/album/juicer-1.jpg,http://example.com/album/juicer-2.jpg', 'http://example.com/main/juicer.jpg', '799-999', 1, 0, 0, 0, 0, 6, 400, 0, 0, 0, 'test', 'test', 'test', 'test', 1);


