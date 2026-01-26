drop table if exists pms_product_spu;
create table pms_product_spu
(
    id                    bigint auto_increment
        primary key comment '商品SpuId',
    name                  varchar(200)                           not null comment '商品名称',
    subTitle              varchar(200)                           not null comment '副标题',
    product_sn            varchar(64)                            not null comment '商品货号',
    category_id           bigint                                 not null comment '商品分类ID',
    category_ids          varchar(20)                            not null comment '商品分类ID集合',
    category_name         varchar(255)                           not null comment '商品分类名称',
    brand_id              bigint                                 not null comment '品牌ID',
    brand_name            varchar(255)                           not null comment '品牌名称',
    unit                  varchar(20)                            not null comment '单位',
    weight                decimal(10, 2)                         not null comment '重量(kg)',
    keywords              varchar(255) default ''                not null comment '关键词',
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
    detail_html           text                                   not null comment '网页详情',
    detail_mobile_html    text                                   not null comment '移动端详情',
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
insert into gozero.pms_product_spu (id, name, subTitle, product_sn, category_id, category_ids, category_name, brand_id, brand_name, unit, weight, keywords, album_pics, main_pic, price_range, publish_status, new_status, recommend_status, verify_status, preview_status, sort, new_status_sort, recommend_status_sort, sales, stock, low_stock, promotion_type, detail_html, detail_mobile_html, create_by, create_time, update_by, update_time, is_deleted)
values  (1, '小米（MI）REDMI Note15 Pro 天玑7400-Ultra 7000mAh 龙晶玻璃十倍抗摔 IP68 8+256 子夜黑 红米 5G手机', '新年狂欢购，限时特惠价，品质好物带回家', 'fsef', 8, '', '小米手机', 9, 'test', '台', 0.19, '小米', 'http://129.204.203.29/xiaomi_s.jpg', 'http://129.204.203.29/xiaomi_s.jpg', '7999-12999', 1, 1, 1, 1, 0, 1, 1, 1, 1000, 0, 0, 0, 'test', 'test', 1, '2026-01-26 10:49:06', null, '2026-01-26 11:22:25', 0),
        (2, 'Apple/苹果 iPhone 17 256GB 薰衣草紫色 支持移动联通电信5G 双卡双待手机', '春节大促销，满减送好礼，实惠享不停', 'fsef', 9, '', '苹果手机', 1, 'test', '台', 0.19, 'Apple', 'http://129.204.203.29/apple_s.jpg', 'http://129.204.203.29/apple_s.jpg', '7999-12999', 1, 1, 1, 1, 0, 1, 1, 1, 1000, 0, 0, 0, 'test', 'test', 1, '2026-01-26 10:49:06', null, '2026-01-26 11:22:25', 0),
        (3, '华为（HUAWEI）Mate X5典藏版X3折叠屏手机全网通正品特北斗卫星华为大折叠 青山黛【Mate X3】 【12G+256G】 赠运费险|详情咨询客服', '年货节来袭，超值折扣季，好货提前囤', 'fsef', 10, '', '华为手机', 5, 'test', '台', 0.19, '华为', 'http://129.204.203.29/hua_s.jpg', 'http://129.204.203.29/hua_s.jpg', '7999-12999', 1, 1, 1, 1, 0, 1, 1, 1, 1000, 0, 0, 0, 'test', 'test', 1, '2026-01-26 10:49:06', null, '2026-01-26 11:22:25', 0),
        (4, '三星（SAMSUNG）W25 心系天下 折叠屏手机2亿像素 Galaxy Ai新品商务高端智能手机 陶瓷黑 16GB+1TB 现货速发|国行正品', '新春钜惠节，买一送一活动，惊喜连连', 'fsef', 11, '', '三星手机', 7, 'test', '台', 0.19, '三星', 'http://129.204.203.29/sumsang.jpg', 'http://129.204.203.29/sumsang.jpg', '7999-12999', 1, 1, 1, 1, 0, 1, 1, 1, 1000, 0, 0, 0, 'test', 'test', 1, '2026-01-26 10:49:06', null, '2026-01-26 11:22:25', 0),
        (5, '荣耀X70 8+128 朱砂红 金标十面抗摔 8300mAh青海湖电池 IP69防水 荣耀绿洲护眼屏 AI手机 国家补贴', '新年新气象，商品大降价，优惠享不停', 'fsef', 12, '', '荣耀手机', 5, 'test', '台', 0.19, '荣耀X70', 'http://129.204.203.29/rong_s.jpg', 'http://129.204.203.29/rong_s.jpg', '7999-12999', 1, 1, 1, 1, 0, 1, 1, 1, 1000, 0, 0, 0, 'test', 'test', 1, '2026-01-26 10:49:06', null, '2026-01-26 11:22:25', 0),
        (6, '苹果（Apple）MacBook Pro Air 超薄商务办公 剪辑设计学生游戏99新13寸Pro丨i5+8G+512G固态丨', '春节特卖会，精选好物五折起，不容错过', 'fsef', 13, '', '苹果电脑', 1, 'test', '台', 0.19, 'MacBook Pro Air', 'http://129.204.203.29/apple_c.jpg', 'http://129.204.203.29/apple_c.jpg', '7999-12999', 1, 1, 1, 1, 0, 1, 1, 1, 1000, 0, 0, 0, 'test', 'test', 1, '2026-01-26 10:49:06', null, '2026-01-26 11:22:25', 0),
        (7, '小米笔记本电脑 红米 REDMI Book 16 酷睿i5标压 16英寸 1TB办公学生轻薄本', '岁末大清仓，年终特价优惠，省钱购物', 'fsef', 14, '', '小米电脑', 9, 'test', '台', 0.19, '红米 REDMI Book 16', 'http://129.204.203.29/xiaomi_c.jpg', 'http://129.204.203.29/xiaomi_c.jpg', '7999-12999', 1, 1, 1, 1, 0, 1, 1, 1, 1000, 0, 0, 0, 'test', 'test', 1, '2026-01-26 10:49:06', null, '2026-01-26 11:22:25', 0),
        (8, '华为（HUAWEI）展机MatebookXPro轻薄旗舰酷睿Ultra9微绒典藏版商务办公电脑 热销新款 Ultra7 32G 1T 触屏', '新年开门红，限时秒杀活动，低至三折', 'fsef', 15, '', '华为电脑', 5, 'test', '台', 0.19, '华为（HUAWEI）展机MatebookXPro', 'http://129.204.203.29/hua_c.jpg', 'http://129.204.203.29/hua_c.jpg', '7999-12999', 1, 1, 1, 1, 0, 1, 1, 1, 1000, 0, 0, 0, 'test', 'test', 1, '2026-01-26 10:49:06', null, '2026-01-26 11:22:25', 0),
        (9, 'HLA海澜之家长袖针织衫男轻商务时尚系列圆领毛衣冬季男', '年味十足购，年货大促销，温馨过新年', 'fsef', 16, '', '海澜之家', 4, 'test', '台', 0.19, 'HLA海澜之家', 'http://129.204.203.29/hailan.jpg', 'http://129.204.203.29/hailan.jpg', '7999-12999', 1, 1, 1, 1, 0, 1, 1, 1, 1000, 0, 0, 0, 'test', 'test', 1, '2026-01-26 10:49:06', null, '2026-01-26 11:22:25', 0),
        (10, '苏泊尔SUPOR好帮手铝合金压力锅4.5L带蒸格20cm高压锅燃气专用', '新春购物节，满额立减优惠，实惠多多', 'fsef', 17, '', '电饭锅', 1, 'test', '台', 0.19, '苏泊尔SUPOR', 'http://129.204.203.29/su.jpg', 'http://129.204.203.29/su.jpg', '7999-12999', 1, 1, 1, 1, 0, 1, 1, 1, 1000, 0, 0, 0, 'test', 'test', 1, '2026-01-26 10:49:06', null, '2026-01-26 11:22:25', 0),
        (11, '格力出品 晶弘空调 小凉神 大1匹 新一级能效变频 纯铜管卧室省电挂机 国家补贴KFR-26GW/JHFNhAa1Bj', '新年新体验，品质商品特价，享受生活', 'fsef', 18, '', '格力空调', 3, 'test', '台', 0.19, '格力出品', 'http://129.204.203.29/geli.jpg', 'http://129.204.203.29/geli.jpg', '7999-12999', 1, 1, 1, 1, 0, 1, 1, 1, 1000, 0, 0, 0, 'test', 'test', 1, '2026-01-26 10:49:06', null, '2026-01-26 11:22:25', 0),
        (12, '方太燃气灶天然气 家用嵌入式5.2kW* 猛火燃气灶 易清洁可调节 TX22', '春节福利季，超值好物优惠，全家共享', 'fsef', 19, '', '方太', 2, 'test', '台', 0.19, '方太燃气灶天然气', 'http://129.204.203.29/fangtai.jpg', 'http://129.204.203.29/fangtai.jpg', '7999-12999', 1, 1, 1, 1, 0, 1, 1, 1, 1000, 0, 0, 0, 'test', 'test', 1, '2026-01-26 10:49:06', null, '2026-01-26 11:22:25', 0),
        (13, '万和【T2】40升2000W速热省电双盾防护专利防电墙租房首选大水量家用储水式电热水器', '年货大集会，折扣狂欢节，欢欢喜喜过大年', 'fsef', 20, '', '万和', 8, 'test', '台', 0.19, '万和', 'http://129.204.203.29/wanhe.png', 'http://129.204.203.29/wanhe.png', '7999-12999', 1, 1, 1, 1, 0, 1, 1, 1, 1000, 0, 0, 0, 'test', 'test', 1, '2026-01-26 10:49:06', null, '2026-01-26 11:22:25', 0),
        (14, '万家乐【安睡洗S9·四驱增压水伺服】16升天然气超一级恒温一级静音零冷感防冻以旧换新S9MAX燃气热水器', '新年大放价，限时抢购活动，先到先得', 'fsef', 21, '', '万家乐', 8, 'test', '台', 0.19, '万家乐', 'http://129.204.203.29/wanjiale.jpg', 'http://129.204.203.29/wanjiale.jpg', '7999-12999', 1, 1, 1, 1, 0, 1, 1, 1, 1000, 0, 0, 0, 'test', 'test', 1, '2026-01-26 10:49:06', null, '2026-01-26 11:22:25', 0),
        (15, '海尔出品统帅60升电热水器京东自营上门安装 国家补贴 2200W节能速热一级能效家用储水式LEC6001-LD5金', '新春嘉年华，特价商品琳琅满目，任君挑选', 'fsef', 22, '', '海尔', 3, 'test', '台', 0.19, '海尔出品统帅', 'http://129.204.203.29/haier.jpg', 'http://129.204.203.29/haier.jpg', '7999-12999', 1, 1, 1, 1, 0, 1, 1, 1, 1000, 0, 0, 0, 'test', 'test', 1, '2026-01-26 10:49:06', null, '2026-01-26 11:22:25', 0),
        (16, '美的（Midea）【整机8年质保】60升电热水器2000W家用出租屋节能抑菌安全防电墙以旧换新门店同款F60-20F1(H)', '春节大回馈，买就送好礼，购物更划算', 'fsef', 23, '', '美的', 3, 'test', '台', 0.19, '美的', 'http://129.204.203.29/meidi.png', 'http://129.204.203.29/meidi.png', '7999-12999', 1, 1, 1, 1, 0, 1, 1, 1, 1000, 0, 0, 0, 'test', 'test', 1, '2026-01-26 10:49:06', null, '2026-01-26 11:22:25', 0);