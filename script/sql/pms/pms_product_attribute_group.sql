drop table if exists pms_product_attribute_group;
create table pms_product_attribute_group
(
    id          bigint auto_increment comment '主键id'
        primary key,
    category_id bigint                             not null comment '分类ID',
    name        varchar(50)                        not null comment '分组名称',
    sort        int      default 0                 not null comment '排序',
    status      tinyint  default 0                 not null comment '状态：0->禁用；1->启用',
    create_by   bigint                             not null comment '创建人ID',
    create_time datetime default CURRENT_TIMESTAMP not null comment '创建时间',
    update_by   bigint                             null comment '更新人ID',
    update_time datetime                           null on update CURRENT_TIMESTAMP comment '更新时间',
    is_deleted  tinyint  default 0                 not null comment '是否删除'
)
    comment '商品属性分组表';

create index idx_category
    on pms_product_attribute_group (category_id, is_deleted);

-- 插入商品属性分组数据
INSERT INTO pms_product_attribute_group (category_id, name, sort, create_by, create_time) VALUES
-- 电子产品(id=1)通用属性组
(1, '基本信息', 1, 1, NOW()),
(1, '规格参数', 2, 1, NOW()),
(1, '包装信息', 3, 1, NOW()),

-- 手机(id=4)属性组
(4, '基本信息', 1, 1, NOW()),
(4, '主体参数', 2, 1, NOW()),
(4, '网络参数', 3, 1, NOW()),
(4, '显示参数', 4, 1, NOW()),
(4, '摄像功能', 5, 1, NOW()),

-- 电脑(id=5)属性组
(5, '基本信息', 1, 1, NOW()),
(5, '处理器', 2, 1, NOW()),
(5, '内存容量', 3, 1, NOW()),
(5, '显卡性能', 4, 1, NOW()),
(5, '存储容量', 5, 1, NOW()),

-- 服装(id=2)通用属性组
(2, '基本信息', 1, 1, NOW()),
(2, '材质参数', 2, 1, NOW()),
(2, '尺码参数', 3, 1, NOW()),
(2, '洗涤说明', 4, 1, NOW()),

-- 男装(id=6)和女装(id=7)属性组
(6, '基本信息', 1, 1, NOW()),
(6, '版型参数', 2, 1, NOW()),
(6, '面料说明', 3, 1, NOW()),
(7, '基本信息', 1, 1, NOW()),
(7, '版型参数', 2, 1, NOW()),
(7, '面料说明', 3, 1, NOW()),

-- 食品(id=3)通用属性组
(3, '基本信息', 1, 1, NOW()),
(3, '营养成分', 2, 1, NOW()),
(3, '储存条件', 3, 1, NOW()),

-- 零食(id=8)和饮料(id=9)属性组
(8, '基本信息', 1, 1, NOW()),
(8, '口味参数', 2, 1, NOW()),
(8, '保质期限', 3, 1, NOW()),
(9, '基本信息', 1, 1, NOW()),
(9, '配料说明', 2, 1, NOW()),
(9, '保质期限', 3, 1, NOW()),

-- 家电(id=10)通用属性组
(10, '基本信息', 1, 1, NOW()),
(10, '规格参数', 2, 1, NOW()),
(10, '功能特点', 3, 1, NOW()),

-- 厨房电器(id=11)属性组
(11, '基本信息', 1, 1, NOW()),
(11, '功率参数', 2, 1, NOW()),
(11, '特色功能', 3, 1, NOW()),
(11, '安全参数', 4, 1, NOW());