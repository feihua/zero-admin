drop table if exists pms_product_attribute;
create table pms_product_attribute
(
    id            bigint auto_increment comment '主键id'
        primary key,
    group_id      bigint                                 not null comment '属性分组ID',
    name          varchar(50)                            not null comment '属性名称',
    input_type    tinyint                                not null comment '输入类型：1-手动输入，2-单选，3-多选',
    value_type    tinyint                                not null comment '值类型：1-文本，2-数字，3-日期',
    input_list    varchar(500) default ''                not null comment '可选值列表，用逗号分隔',
    unit          varchar(20)  default ''                not null comment '单位',
    is_required   tinyint      default 0                 not null comment '是否必填',
    is_searchable tinyint      default 0                 not null comment '是否支持搜索',
    is_show       tinyint      default 1                 not null comment '是否显示',
    sort          int          default 0                 not null comment '排序',
    status        tinyint      default 0                 not null comment '状态：0->禁用；1->启用',
    create_by     bigint                                 not null comment '创建人ID',
    create_time   datetime     default CURRENT_TIMESTAMP not null comment '创建时间',
    update_by     bigint                                 null comment '更新人ID',
    update_time   datetime                               null on update CURRENT_TIMESTAMP comment '更新时间',
    is_deleted    tinyint      default 0                 not null comment '是否删除'
)
    comment '商品属性表';

create index idx_group
    on pms_product_attribute (group_id, is_deleted);

-- 插入商品属性数据
INSERT INTO pms_product_attribute (group_id, name, input_type, value_type, input_list, unit, is_required, is_searchable, is_show, sort, create_by) VALUES
-- 手机属性 (对应group_id为手机基本信息、主体参数等分组的ID)
(8, '品牌', 2, 1, 'Apple,Samsung,Huawei,Xiaomi', '', 1, 1, 1, 1, 1),
(8, '型号', 1, 1, '', '', 1, 1, 1, 2, 1),
(8, '上市年份', 1, 2, '', '年', 1, 1, 1, 3, 1),
(8, '机身颜色', 2, 1, '黑色,白色,金色,银色', '', 1, 1, 1, 4, 1),
(8, '运行内存', 2, 2, '4GB,6GB,8GB,12GB', 'GB', 1, 1, 1, 5, 1),
(8, '机身存储', 2, 2, '64GB,128GB,256GB,512GB', 'GB', 1, 1, 1, 6, 1),
(8, '屏幕尺寸', 1, 2, '', '英寸', 1, 1, 1, 7, 1),
(8, '电池容量', 1, 2, '', 'mAh', 1, 0, 1, 8, 1);

# -- 电脑属性
# (5, '品牌', 2, 1, 'Lenovo,HP,Dell,Apple', '', 1, 1, 1, 1, 1),
# (5, 'CPU型号', 2, 1, 'Intel i5,Intel i7,Intel i9,AMD', '', 1, 1, 1, 2, 1),
# (5, '内存容量', 2, 2, '8GB,16GB,32GB,64GB', 'GB', 1, 1, 1, 3, 1),
# (5, '硬盘容量', 2, 2, '256GB,512GB,1TB,2TB', 'GB', 1, 1, 1, 4, 1),
# (5, '显卡型号', 1, 1, '', '', 1, 1, 1, 5, 1),
#
# -- 服装通用属性
# (2, '品牌', 2, 1, 'Nike,Adidas,Uniqlo,H&M', '', 1, 1, 1, 1, 1),
# (2, '适用季节', 3, 1, '春季,夏季,秋季,冬季', '', 1, 1, 1, 2, 1),
# (2, '材质成分', 1, 1, '', '', 1, 0, 1, 3, 1),
# (2, '尺码', 2, 1, 'S,M,L,XL,XXL', '', 1, 1, 1, 4, 1),
#
# -- 食品通用属性
# (3, '品牌', 2, 1, '', '', 1, 1, 1, 1, 1),
# (3, '保质期', 1, 2, '', '天', 1, 0, 1, 2, 1),
# (3, '储存方式', 2, 1, '常温,冷藏,冷冻', '', 1, 0, 1, 3, 1),
# (3, '生产日期', 1, 3, '', '', 1, 0, 1, 4, 1),
# (3, '产地', 1, 1, '', '', 1, 1, 1, 5, 1),
#
# -- 家电通用属性
# (10, '品牌', 2, 1, 'Midea,Haier,Gree,Siemens', '', 1, 1, 1, 1, 1),
# (10, '型号', 1, 1, '', '', 1, 1, 1, 2, 1),
# (10, '功率', 1, 2, '', 'W', 1, 0, 1, 3, 1),
# (10, '能效等级', 2, 1, 'A+++,A++,A+,A,B', '', 1, 1, 1, 4, 1),
# (10, '质保期', 1, 2, '', '年', 1, 0, 1, 5, 1),
#
# -- 厨房电器属性
# (11, '容量', 1, 2, '', 'L', 1, 1, 1, 1, 1),
# (11, '控制方式', 2, 1, '按键式,触控式,旋钮式,智能控制', '', 1, 1, 1, 2, 1),
# (11, '特色功能', 3, 1, '预约,保温,除菌,智能控温', '', 0, 1, 1, 3, 1);
#
#
#
