drop table if exists pms_product_spec;
create table pms_product_spec
(
    id          bigint auto_increment
        primary key,
    category_id bigint                             not null comment '分类ID',
    name        varchar(50)                        not null comment '规格名称',
    sort        int      default 0                 not null comment '排序',
    status      tinyint  default 1                 not null comment '状态：0->禁用；1->启用',
    create_by   bigint   default 1                 not null comment '创建人ID',
    create_time datetime default CURRENT_TIMESTAMP not null comment '创建时间',
    update_by   bigint                             null comment '更新人ID',
    update_time datetime                           null on update CURRENT_TIMESTAMP comment '更新时间',
    is_deleted  tinyint  default 0                 not null comment '是否删除'
)
    comment '商品规格表';

create index idx_category
    on pms_product_spec (category_id, is_deleted);

-- 插入商品规格数据
INSERT INTO pms_product_spec (category_id, name, sort)
VALUES
-- 手机规格(category_id=1)
(1, '颜色', 1),
(1, '容量', 2),
(1, '网络版本', 3),
-- 小米手机规格(category_id=8)
(8, '颜色', 1),
(8, '容量', 2),
(8, '网络版本', 3),
-- 苹果手机规格(category_id=9)
(9, '颜色', 1),
(9, '容量', 2),
(9, '网络版本', 3),
-- 华为手机规格(category_id=10)
(10, '颜色', 1),
(10, '容量', 2),
(10, '网络版本', 3),
-- 三星手机规格(category_id=11)
(11, '颜色', 1),
(11, '容量', 2),
(11, '网络版本', 3),
-- 荣耀手机规格(category_id=12)
(12, '颜色', 1),
(12, '容量', 2),
(12, '网络版本', 3),

-- 电脑规格(category_id=2)
(2, '颜色', 1),
(2, '内存', 2),
(2, '硬盘', 3),
-- 苹果电脑规格(category_id=13)
(13, '颜色', 1),
(13, '内存', 2),
(13, '硬盘', 3),
-- 小米电脑规格(category_id=14)
(14, '颜色', 1),
(14, '内存', 2),
(14, '硬盘', 3),
-- 华为电脑规格(category_id=15)
(15, '颜色', 1),
(15, '内存', 2),
(15, '硬盘', 3),

-- 服装规格(category_id=3)
(3, '颜色', 1),
(3, '尺码', 2),
-- 海澜之家规格(category_id=16)
(16, '颜色', 1),
(16, '尺码', 2),


-- 电器规格(category_id=4)
(4, '颜色', 1),
(4, '容量', 2),
(4, '功率', 3),
-- 电饭锅规格(category_id=17)
(17, '颜色', 1),
(17, '容量', 2),
(17, '功率', 3),

-- 空调规格(category_id=5)
(5, '颜色', 1),
(5, '型号', 2),
(5, '功率', 3),
-- 格力空调规格(category_id=18)
(18, '颜色', 1),
(18, '型号', 2),
(18, '功率', 3),

-- 厨房规格(category_id=6)
(6, '颜色', 1),
(6, '火力', 2),
(6, '功率', 3),
-- 方太规格(category_id=19)
(19, '颜色', 1),
(19, '火力', 2),
(19, '功率', 3),

-- 热水器规格(category_id=7)
(7, '颜色', 1),
(7, '容量', 2),
(7, '款式', 3),
(7, '功率', 4),
-- 万和热水器规格(category_id=20)
(20, '颜色', 1),
(20, '容量', 2),
(20, '款式', 3),
(20, '功率', 4),
-- 万家乐热水器规格(category_id=21)
(21, '颜色', 1),
(21, '容量', 2),
(21, '款式', 3),
(21, '功率', 4),
-- 海尔热水器规格(category_id=22)
(22, '颜色', 1),
(22, '容量', 2),
(22, '款式', 3),
(22, '功率', 4),
-- 美的热水器规格(category_id=23)
(23, '颜色', 1),
(23, '容量', 2),
(23, '款式', 3),
(23, '功率', 4);