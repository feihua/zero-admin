drop table if exists pms_product_category;
create table pms_product_category
(
    id            bigint auto_increment
        primary key,
    parent_id     bigint                                not null comment '上级分类的编号：0表示一级分类',
    name          varchar(64)                           not null comment '商品分类名称',
    level         tinyint                               not null comment '分类级别：0->1级；1->2级',
    product_count int                                   not null comment '商品数量',
    product_unit  varchar(64) DEFAULT '件'              not null comment '商品单位',
    nav_status    tinyint                               not null comment '是否显示在导航栏：0->不显示；1->显示',
    sort          int                                   not null comment '排序',
    icon          varchar(255)                          not null comment '图标',
    keywords      varchar(255)                          not null comment '关键字',
    description   text                                  not null comment '描述',
    is_enabled    tinyint     default 1                 not null comment '是否启用',
    create_by     bigint                                not null comment '创建人ID',
    create_time   datetime    default CURRENT_TIMESTAMP not null comment '创建时间',
    update_by     bigint                                null comment '更新人ID',
    update_time   datetime                              null on update CURRENT_TIMESTAMP comment '更新时间',
    is_deleted    tinyint     default 0                 not null comment '是否删除'

)
    comment '产品分类';

-- 插入产品分类数据
insert into pms_product_category (id, parent_id, name, level, product_count, product_unit, nav_status, sort, icon, keywords, description, is_enabled, create_by, is_deleted)
values  (1, 0, '电子产品', 0, 100, '件', 1, 1, 'http://129.204.203.29/big.png', '电子,科技', '电子产品分类', 1, 1, 0),
        (2, 0, '服装', 0, 200, '件', 1, 2, 'http://129.204.203.29/big.png', '时尚,服饰', '服装分类', 1, 2, 0),
        (3, 0, '食品', 0, 150, '件', 1, 3, 'http://129.204.203.29/big.png', '食品,饮料', '食品分类', 1, 3, 0),
        (4, 1, '手机', 1, 50, '件', 1, 4, 'http://129.204.203.29/big.png', '手机,智能', '手机分类', 1, 4, 0),
        (5, 1, '电脑', 1, 30, '件', 1, 5, 'http://129.204.203.29/big.png', '电脑,笔记本', '电脑分类', 1, 5, 0),
        (6, 2, '男装', 1, 80, '件', 1, 6, 'http://129.204.203.29/big.png', '男装,时尚', '男装分类', 1, 6, 0),
        (7, 2, '女装', 1, 120, '件', 1, 7, 'http://129.204.203.29/big.png', '女装,时尚', '女装分类', 1, 7, 0),
        (8, 3, '零食', 1, 60, '件', 1, 8, 'http://129.204.203.29/big.png', '零食,小吃', '零食分类', 1, 8, 0),
        (9, 3, '饮料', 1, 90, '件', 1, 9, 'http://129.204.203.29/big.png', '饮料,饮品', '饮料分类', 1, 9, 0),
        (10, 0, '家电', 0, 70, '件', 1, 10, 'http://129.204.203.29/big.png', '家电,电器', '家电分类', 1, 10, 0),
        (11, 10, '厨房电器', 1, 40, '件', 1, 11, 'http://129.204.203.29/big.png', '厨房,电器', '厨房电器分类', 1, 11, 0);