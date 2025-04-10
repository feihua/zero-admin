create table pms_product_category
(
    id            bigint auto_increment
        primary key,
    parent_id     bigint                             not null comment '上机分类的编号：0表示一级分类',
    name          varchar(64)                        not null comment '商品分类名称',
    level         tinyint                            not null comment '分类级别：0->1级；1->2级',
    product_count int                                not null comment '商品数量',
    product_unit  varchar(64)                        not null comment '商品数量',
    nav_status    tinyint                            not null comment '是否显示在导航栏：0->不显示；1->显示',
    show_status   tinyint                            not null comment '显示状态：0->不显示；1->显示',
    sort          int                                not null comment '排序',
    icon          varchar(255)                       not null comment '图标',
    keywords      varchar(255)                       not null comment '关键字',
    description   text                               not null comment '描述',
    create_time   datetime default CURRENT_TIMESTAMP not null comment '创建时间',
    update_time   datetime                           null on update CURRENT_TIMESTAMP comment '更新时间'

)
    comment '产品分类';


INSERT INTO pms_product_category (id, parent_id, name, level, product_count, product_unit, nav_status, show_status, sort, icon, keywords, description) VALUES (1, 0, '服装', 0, 100, '件', 1, 1, 1, 'http://macro-oss.oss-cn-shenzhen.aliyuncs.com/mall/images/20180522/web.png', '服装', '服装分类');
INSERT INTO pms_product_category (id, parent_id, name, level, product_count, product_unit, nav_status, show_status, sort, icon, keywords, description) VALUES (2, 0, '手机数码', 0, 100, '件', 1, 1, 1, 'http://macro-oss.oss-cn-shenzhen.aliyuncs.com/mall/images/20180522/web.png', '手机数码', '手机数码');
INSERT INTO pms_product_category (id, parent_id, name, level, product_count, product_unit, nav_status, show_status, sort, icon, keywords, description) VALUES (3, 0, '家用电器', 0, 100, '件', 1, 1, 1, 'http://macro-oss.oss-cn-shenzhen.aliyuncs.com/mall/images/20180522/web.png', '家用电器', '家用电器');
INSERT INTO pms_product_category (id, parent_id, name, level, product_count, product_unit, nav_status, show_status, sort, icon, keywords, description) VALUES (4, 0, '家具家装', 0, 100, '件', 1, 1, 1, 'http://macro-oss.oss-cn-shenzhen.aliyuncs.com/mall/images/20180522/web.png', '家具家装', '家具家装');
INSERT INTO pms_product_category (id, parent_id, name, level, product_count, product_unit, nav_status, show_status, sort, icon, keywords, description) VALUES (5, 0, '汽车用品', 0, 100, '件', 1, 1, 1, 'http://macro-oss.oss-cn-shenzhen.aliyuncs.com/mall/images/20180522/web.png', '汽车用品', '汽车用品');
INSERT INTO pms_product_category (id, parent_id, name, level, product_count, product_unit, nav_status, show_status, sort, icon, keywords, description) VALUES (7, 1, '外套', 1, 100, '件', 1, 1, 0, 'http://macro-oss.oss-cn-shenzhen.aliyuncs.com/mall/images/20190519/5ac4780cN6087feb5.jpg', '外套', '外套');
INSERT INTO pms_product_category (id, parent_id, name, level, product_count, product_unit, nav_status, show_status, sort, icon, keywords, description) VALUES (8, 1, 'T恤', 1, 100, '件', 1, 1, 0, 'http://macro-oss.oss-cn-shenzhen.aliyuncs.com/mall/images/20190519/5ac47ffaN8a7b2e14.png', 'T恤', 'T恤');
INSERT INTO pms_product_category (id, parent_id, name, level, product_count, product_unit, nav_status, show_status, sort, icon, keywords, description) VALUES (9, 1, '休闲裤', 1, 100, '件', 1, 1, 0, 'http://macro-oss.oss-cn-shenzhen.aliyuncs.com/mall/images/20190519/5ac47845N7374a31d.jpg', '休闲裤', '休闲裤');
INSERT INTO pms_product_category (id, parent_id, name, level, product_count, product_unit, nav_status, show_status, sort, icon, keywords, description) VALUES (10, 1, '牛仔裤', 1, 100, '件', 1, 1, 0, 'http://macro-oss.oss-cn-shenzhen.aliyuncs.com/mall/images/20190519/5ac47841Nff658599.jpg', '牛仔裤', '牛仔裤');
INSERT INTO pms_product_category (id, parent_id, name, level, product_count, product_unit, nav_status, show_status, sort, icon, keywords, description) VALUES (11, 1, '衬衫', 1, 100, '件', 1, 1, 0, 'http://macro-oss.oss-cn-shenzhen.aliyuncs.com/mall/images/20190519/5ac48007Nb30b2118.jpg', '衬衫', '衬衫分类');
INSERT INTO pms_product_category (id, parent_id, name, level, product_count, product_unit, nav_status, show_status, sort, icon, keywords, description) VALUES (13, 12, '家电子分类1', 1, 1, 'string', 0, 1, 0, 'http://macro-oss.oss-cn-shenzhen.aliyuncs.com/mall/images/20180522/web.png', 'string', 'string');
INSERT INTO pms_product_category (id, parent_id, name, level, product_count, product_unit, nav_status, show_status, sort, icon, keywords, description) VALUES (14, 12, '家电子分类2', 1, 1, 'string', 0, 1, 0, 'http://macro-oss.oss-cn-shenzhen.aliyuncs.com/mall/images/20180522/web.png', 'string', 'string');
INSERT INTO pms_product_category (id, parent_id, name, level, product_count, product_unit, nav_status, show_status, sort, icon, keywords, description) VALUES (19, 2, '手机通讯', 1, 0, '件', 0, 1, 0, 'http://macro-oss.oss-cn-shenzhen.aliyuncs.com/mall/images/20190519/5ac48d27N3f5bb821.jpg', '手机通讯', '手机通讯');
INSERT INTO pms_product_category (id, parent_id, name, level, product_count, product_unit, nav_status, show_status, sort, icon, keywords, description) VALUES (29, 1, '男鞋', 1, 0, '', 0, 1, 0, 'http://macro-oss.oss-cn-shenzhen.aliyuncs.com/mall/images/20180522/web.png', '', '');
INSERT INTO pms_product_category (id, parent_id, name, level, product_count, product_unit, nav_status, show_status, sort, icon, keywords, description) VALUES (30, 2, '手机配件', 1, 0, '', 0, 1, 0, 'http://macro-oss.oss-cn-shenzhen.aliyuncs.com/mall/images/20190519/5ac48672N11cf61fe.jpg', '手机配件', '手机配件');
INSERT INTO pms_product_category (id, parent_id, name, level, product_count, product_unit, nav_status, show_status, sort, icon, keywords, description) VALUES (31, 2, '摄影摄像', 1, 0, '', 0, 1, 0, 'http://macro-oss.oss-cn-shenzhen.aliyuncs.com/mall/images/20190519/5a1679f2Nc2f659b6.jpg', '', '');
INSERT INTO pms_product_category (id, parent_id, name, level, product_count, product_unit, nav_status, show_status, sort, icon, keywords, description) VALUES (32, 2, '影音娱乐', 1, 0, '', 0, 1, 0, 'http://macro-oss.oss-cn-shenzhen.aliyuncs.com/mall/images/20190519/5a167859N01d8198b.jpg', '', '');
INSERT INTO pms_product_category (id, parent_id, name, level, product_count, product_unit, nav_status, show_status, sort, icon, keywords, description) VALUES (33, 2, '数码配件', 1, 0, '', 0, 1, 0, 'http://macro-oss.oss-cn-shenzhen.aliyuncs.com/mall/images/20190519/5a1676e9N1ba70a81.jpg', '', '');
INSERT INTO pms_product_category (id, parent_id, name, level, product_count, product_unit, nav_status, show_status, sort, icon, keywords, description) VALUES (34, 2, '智能设备', 1, 0, '', 0, 1, 0, 'http://macro-oss.oss-cn-shenzhen.aliyuncs.com/mall/images/20180522/web.png', '', '');
INSERT INTO pms_product_category (id, parent_id, name, level, product_count, product_unit, nav_status, show_status, sort, icon, keywords, description) VALUES (35, 3, '电视', 1, 0, '', 0, 1, 0, 'http://macro-oss.oss-cn-shenzhen.aliyuncs.com/mall/images/20180522/web.png', '', '');
INSERT INTO pms_product_category (id, parent_id, name, level, product_count, product_unit, nav_status, show_status, sort, icon, keywords, description) VALUES (36, 3, '空调', 1, 0, '', 0, 1, 0, 'http://macro-oss.oss-cn-shenzhen.aliyuncs.com/mall/images/20180522/web.png', '', '');
INSERT INTO pms_product_category (id, parent_id, name, level, product_count, product_unit, nav_status, show_status, sort, icon, keywords, description) VALUES (37, 3, '洗衣机', 1, 0, '', 0, 1, 0, 'http://macro-oss.oss-cn-shenzhen.aliyuncs.com/mall/images/20180522/web.png', '', '');
INSERT INTO pms_product_category (id, parent_id, name, level, product_count, product_unit, nav_status, show_status, sort, icon, keywords, description) VALUES (38, 3, '冰箱', 1, 0, '', 0, 1, 0, 'http://macro-oss.oss-cn-shenzhen.aliyuncs.com/mall/images/20180522/web.png', '', '');
INSERT INTO pms_product_category (id, parent_id, name, level, product_count, product_unit, nav_status, show_status, sort, icon, keywords, description) VALUES (39, 3, '厨卫大电', 1, 0, '', 0, 1, 0, 'http://macro-oss.oss-cn-shenzhen.aliyuncs.com/mall/images/20180522/web.png', '', '');
INSERT INTO pms_product_category (id, parent_id, name, level, product_count, product_unit, nav_status, show_status, sort, icon, keywords, description) VALUES (40, 3, '厨房小电', 1, 0, '', 0, 1, 0, 'http://macro-oss.oss-cn-shenzhen.aliyuncs.com/mall/images/20180522/web.png', '', '');
INSERT INTO pms_product_category (id, parent_id, name, level, product_count, product_unit, nav_status, show_status, sort, icon, keywords, description) VALUES (41, 3, '生活电器', 1, 0, '', 0, 1, 0, 'http://macro-oss.oss-cn-shenzhen.aliyuncs.com/mall/images/20180522/web.png', '', '');
INSERT INTO pms_product_category (id, parent_id, name, level, product_count, product_unit, nav_status, show_status, sort, icon, keywords, description) VALUES (42, 3, '个护健康', 1, 0, '', 0, 1, 0, 'http://macro-oss.oss-cn-shenzhen.aliyuncs.com/mall/images/20180522/web.png', '', '');
INSERT INTO pms_product_category (id, parent_id, name, level, product_count, product_unit, nav_status, show_status, sort, icon, keywords, description) VALUES (43, 4, '厨房卫浴', 1, 0, '', 0, 1, 0, 'http://macro-oss.oss-cn-shenzhen.aliyuncs.com/mall/images/20180522/web.png', '', '');
INSERT INTO pms_product_category (id, parent_id, name, level, product_count, product_unit, nav_status, show_status, sort, icon, keywords, description) VALUES (44, 4, '灯饰照明', 1, 0, '', 0, 1, 0, 'http://macro-oss.oss-cn-shenzhen.aliyuncs.com/mall/images/20180522/web.png', '', '');
INSERT INTO pms_product_category (id, parent_id, name, level, product_count, product_unit, nav_status, show_status, sort, icon, keywords, description) VALUES (45, 4, '五金工具', 1, 0, '', 0, 1, 0, 'http://macro-oss.oss-cn-shenzhen.aliyuncs.com/mall/images/20180522/web.png', '', '');
INSERT INTO pms_product_category (id, parent_id, name, level, product_count, product_unit, nav_status, show_status, sort, icon, keywords, description) VALUES (46, 4, '卧室家具', 1, 0, '', 0, 1, 0, 'http://macro-oss.oss-cn-shenzhen.aliyuncs.com/mall/images/20180522/web.png', '', '');
INSERT INTO pms_product_category (id, parent_id, name, level, product_count, product_unit, nav_status, show_status, sort, icon, keywords, description) VALUES (47, 4, '客厅家具', 1, 0, '', 0, 1, 0, 'http://macro-oss.oss-cn-shenzhen.aliyuncs.com/mall/images/20180522/web.png', '', '');
INSERT INTO pms_product_category (id, parent_id, name, level, product_count, product_unit, nav_status, show_status, sort, icon, keywords, description) VALUES (48, 5, '全新整车', 1, 0, '', 0, 1, 0, 'http://macro-oss.oss-cn-shenzhen.aliyuncs.com/mall/images/20180522/web.png', '', '');
INSERT INTO pms_product_category (id, parent_id, name, level, product_count, product_unit, nav_status, show_status, sort, icon, keywords, description) VALUES (49, 5, '车载电器', 1, 0, '', 0, 1, 0, 'http://macro-oss.oss-cn-shenzhen.aliyuncs.com/mall/images/20180522/web.png', '', '');
INSERT INTO pms_product_category (id, parent_id, name, level, product_count, product_unit, nav_status, show_status, sort, icon, keywords, description) VALUES (50, 5, '维修保养', 1, 0, '', 0, 1, 0, 'http://macro-oss.oss-cn-shenzhen.aliyuncs.com/mall/images/20180522/web.png', '', '');
INSERT INTO pms_product_category (id, parent_id, name, level, product_count, product_unit, nav_status, show_status, sort, icon, keywords, description) VALUES (51, 5, '汽车装饰', 1, 0, '', 0, 1, 0, 'http://macro-oss.oss-cn-shenzhen.aliyuncs.com/mall/images/20180522/web.png', '', '');
