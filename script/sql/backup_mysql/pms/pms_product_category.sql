drop table if exists pms_product_category;
create table pms_product_category
(
    id            bigint auto_increment
        primary key,
    parent_id     bigint                                 not null comment '上级分类的编号：0表示一级分类',
    name          varchar(64)                            not null comment '商品分类名称',
    level         tinyint      default 1                 not null comment '分类级别：0->1级；1->2级',
    product_count int          default 0                 not null comment '商品数量',
    product_unit  varchar(64)  DEFAULT '件'              not null comment '商品单位',
    nav_status    tinyint      default 1                 not null comment '是否显示在导航栏：0->不显示；1->显示',
    sort          int          default 0                 not null comment '排序',
    icon          varchar(255) default ''                not null comment '图标',
    keywords      varchar(255) default ''                not null comment '关键字',
    description   varchar(512) default ''                not null comment '描述',
    is_enabled    tinyint      default 1                 not null comment '是否启用',
    create_by     bigint       default 1                 not null comment '创建人ID',
    create_time   datetime     default CURRENT_TIMESTAMP not null comment '创建时间',
    update_by     bigint                                 null comment '更新人ID',
    update_time   datetime                               null on update CURRENT_TIMESTAMP comment '更新时间',
    is_deleted    tinyint      default 0                 not null comment '是否删除'

)
    comment '产品分类';

-- 插入产品分类数据

insert into pms_product_category (id, parent_id, name, level, product_count, product_unit, nav_status, sort, icon)
values (1, 0, '手机', 0, 100, '台', 1, 1, ''),
       (2, 0, '电脑', 0, 200, '台', 1, 2, ''),
       (3, 0, '衣服', 0, 150, '件', 1, 3, ''),
       (4, 0, '电器', 0, 50, '台', 1, 4, ''),
       (5, 0, '空调', 0, 30, '件', 1, 5, ''),
       (6, 0, '厨房', 0, 810, '件', 1, 6, ''),
       (7, 0, '热水器', 0, 320, '件', 1, 7, ''),
       (8, 1, '小米手机', 1, 620, '台', 1, 8, 'http://129.204.203.29/xiaomi_s.jpg'),
       (9, 1, '苹果手机', 1, 90, '台', 1, 9, 'http://129.204.203.29/apple_s.jpg'),
       (10, 1, '华为手机', 1, 70, '台', 1, 10, 'http://129.204.203.29/hua_s.jpg'),
       (11, 1, '三星手机', 1, 340, '台', 1, 11, 'http://129.204.203.29/sumsang.jpg'),
       (12, 1, '荣耀手机', 1, 470, '台', 1, 12, 'http://129.204.203.29/rong_s.jpg'),
       (13, 2, '苹果电脑', 1, 140, '台', 1, 13, 'http://129.204.203.29/apple_c.jpg'),
       (14, 2, '小米电脑', 1, 410, '台', 1, 14, 'http://129.204.203.29/xiaomi_c.jpg'),
       (15, 2, '华为电脑', 1, 950, '台', 1, 15, 'http://129.204.203.29/hua_c.jpg'),
       (16, 3, '海澜之家', 1, 840, '件', 1, 16, 'http://129.204.203.29/hailan.jpg'),
       (17, 4, '电饭锅', 1, 540, '件', 1, 17, 'http://129.204.203.29/su.jpg'),
       (18, 5, '格力空调', 1, 400, '台', 1, 18, 'http://129.204.203.29/geli.jpg'),
       (19, 6, '方太', 1, 406, '件', 1, 19, 'http://129.204.203.29/fangtai.jpg'),
       (20, 7, '万和', 1, 70, '件', 1, 20, 'http://129.204.203.29/wanhe.png'),
       (21, 7, '万家乐', 1, 90, '件', 1, 21, 'http://129.204.203.29/wanjiale.jpg'),
       (22, 7, '海尔', 1, 112, '件', 1, 22, 'http://129.204.203.29/haier.jpg'),
       (23, 7, '美的', 1, 150, '件', 1, 23, 'http://129.204.203.29/meidi.png');