drop table if exists pms_product_spec_value;
create table pms_product_spec_value
(
    id          bigint auto_increment
        primary key,
    spec_id     bigint                             not null comment '规格ID',
    value       varchar(50)                        not null comment '规格值',
    sort        int      default 0                 not null comment '排序',
    create_by   bigint                             not null comment '创建人ID',
    create_time datetime default CURRENT_TIMESTAMP not null comment '创建时间',
    update_by   bigint                             null comment '更新人ID',
    update_time datetime                           null on update CURRENT_TIMESTAMP comment '更新时间',
    is_deleted  tinyint  default 0                 not null comment '是否删除'
)
    comment '商品规格值表';

create index idx_spec
    on pms_product_spec_value (spec_id, is_deleted);

-- 插入商品规格值数据
INSERT INTO pms_product_spec_value (spec_id, value, sort, create_by) VALUES
-- 手机颜色规格值
(1, '暗紫色', 1, 1),
(1, '自然色', 2, 1),
(1, '银色', 3, 1),
(1, '金色', 4, 1),

-- 手机容量规格值
(2, '128GB', 1, 1),
(2, '256GB', 2, 1),
(2, '512GB', 3, 1),
(2, '1TB', 4, 1),

-- 手机网络版本规格值
(3, '全网通版', 1, 1),
(3, '5G版', 2, 1),

-- 电脑颜色规格值
(4, '深空灰', 1, 1),
(4, '银色', 2, 1),
(4, '黑色', 3, 1),

-- 电脑内存规格值
(5, '8GB', 1, 1),
(5, '16GB', 2, 1),
(5, '32GB', 3, 1),

-- 电脑硬盘规格值
(6, '256GB', 1, 1),
(6, '512GB', 2, 1),
(6, '1TB', 3, 1),

-- 服装颜色规格值
(8, '黑色', 1, 1),
(8, '白色', 2, 1),
(8, '灰色', 3, 1),
(8, '藏青色', 4, 1),

-- 服装尺码规格值
(9, 'S', 1, 1),
(9, 'M', 2, 1),
(9, 'L', 3, 1),
(9, 'XL', 4, 1),
(9, 'XXL', 5, 1),

-- 食品口味规格值
(16, '原味', 1, 1),
(16, '辣味', 2, 1),
(16, '甜味', 3, 1),
(16, '咸味', 4, 1),

-- 食品包装规格值
(17, '单包装', 1, 1),
(17, '礼盒装', 2, 1),
(17, '家庭装', 3, 1),

-- 家电颜色规格值
(24, '白色', 1, 1),
(24, '银色', 2, 1),
(24, '黑色', 3, 1);