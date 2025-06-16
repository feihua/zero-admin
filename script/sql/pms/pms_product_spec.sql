drop table if exists pms_product_spec;
create table pms_product_spec
(
    id          bigint auto_increment
        primary key,
    category_id bigint                             not null comment '分类ID',
    name        varchar(50)                        not null comment '规格名称',
    sort        int      default 0                 not null comment '排序',
    status      tinyint  default 0                 not null comment '状态：0->禁用；1->启用',
    create_by   bigint                             not null comment '创建人ID',
    create_time datetime default CURRENT_TIMESTAMP not null comment '创建时间',
    update_by   bigint                             null comment '更新人ID',
    update_time datetime                           null on update CURRENT_TIMESTAMP comment '更新时间',
    is_deleted  tinyint  default 0                 not null comment '是否删除'
)
    comment '商品规格表';

create index idx_category
    on pms_product_spec (category_id, is_deleted);

-- 插入商品规格数据
INSERT INTO pms_product_spec (category_id, name, sort, create_by) VALUES
-- 手机规格(category_id=4)
(4, '颜色', 1, 1),
(4, '容量', 2, 1),
(4, '网络版本', 3, 1),

-- 电脑规格(category_id=5)
(5, '颜色', 1, 1),
(5, '内存', 2, 1),
(5, '硬盘', 3, 1),
(5, '处理器', 4, 1),

-- 服装规格(category_id=2)
(2, '颜色', 1, 1),
(2, '尺码', 2, 1),

-- 男装规格(category_id=6)
(6, '颜色', 1, 1),
(6, '尺码', 2, 1),
(6, '款式', 3, 1),

-- 女装规格(category_id=7)
(7, '颜色', 1, 1),
(7, '尺码', 2, 1),
(7, '款式', 3, 1),

-- 食品规格(category_id=3)
(3, '口味', 1, 1),
(3, '包装规格', 2, 1),

-- 零食规格(category_id=8)
(8, '口味', 1, 1),
(8, '包装规格', 2, 1),
(8, '净含量', 3, 1),

-- 饮料规格(category_id=9)
(9, '口味', 1, 1),
(9, '包装规格', 2, 1),
(9, '容量', 3, 1),

-- 家电规格(category_id=10)
(10, '颜色', 1, 1),
(10, '型号', 2, 1),

-- 厨房电器规格(category_id=11)
(11, '颜色', 1, 1),
(11, '容量', 2, 1),
(11, '款式', 3, 1);