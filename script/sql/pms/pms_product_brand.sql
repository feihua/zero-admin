drop table if exists pms_product_brand;
create table pms_product_brand
(
    id                    bigint auto_increment
        primary key,
    name                  varchar(64)                        not null comment '品牌名称',
    logo                  varchar(255)                       not null comment '品牌logo',
    big_pic               varchar(255)                       not null comment '专区大图',
    description           text                               not null comment '描述',
    first_letter          varchar(8)                         not null comment '首字母',
    sort                  int                                not null comment '排序',
    recommend_status      tinyint  default 0                 not null comment '推荐状态',
    product_count         int      default 0                 not null comment '产品数量',
    product_comment_count int      default 0                 not null comment '产品评论数量',
    is_enabled            tinyint  default 1                 not null comment '是否启用',
    create_by             bigint                             not null comment '创建人ID',
    create_time           datetime default CURRENT_TIMESTAMP not null comment '创建时间',
    update_by             bigint                             null comment '更新人ID',
    update_time           datetime                           null on update CURRENT_TIMESTAMP comment '更新时间',
    is_deleted            tinyint  default 0                 not null comment '是否删除'
)
    comment '商品品牌';

-- 插入商品品牌数据
insert into pms_product_brand (id, name, logo, big_pic, description, first_letter, sort, recommend_status, product_count, product_comment_count, is_enabled, create_by)
values (1, '苹果', 'http://129.204.203.29/brand_apple.png', 'http://129.204.203.29/brand_apple.png', '苹果的描述', 'A', 1, 1, 100, 50, 1, 1),
       (2, '方太', 'http://129.204.203.29/brand_fotile.png', 'http://129.204.203.29/brand_fotile.png', '方太的描述', 'F', 2, 0, 200, 80, 1, 2),
       (3, '格力', 'http://129.204.203.29/brand_geli.png', 'http://129.204.203.29/brand_geli.png', '格力的描述', 'G', 3, 1, 150, 60, 1, 3),
       (4, '海澜之家', 'http://129.204.203.29/brand_hailan.png', 'http://129.204.203.29/brand_hailan.png', '海澜之家的描述', 'H', 4, 1, 120, 40, 1, 4),
       (5, '华为', 'http://129.204.203.29/brand_huawei.png', 'http://129.204.203.29/brand_huawei.png', '华为的描述', 'H', 5, 1, 180, 70, 1, 5),
       (6, 'oppo', 'http://129.204.203.29/brand_oppo.png', 'http://129.204.203.29/brand_oppo.png', 'oppo的描述', 'O', 6, 0, 90, 30, 1, 6),
       (7, '三星', 'http://129.204.203.29/brand_sumsung.png', 'http://129.204.203.29/brand_sumsung.png', '三星的描述', 'S', 7, 1, 110, 50, 1, 7),
       (8, '万和', 'http://129.204.203.29/brand_wanhe.png', 'http://129.204.203.29/brand_wanhe.png', '万和的描述', 'W', 8, 0, 130, 60, 1, 8),
       (9, '小米', 'http://129.204.203.29/brand_xiaomi.png', 'http://129.204.203.29/brand_xiaomi.png', '小米的描述', 'X', 9, 1, 140, 55, 1, 9);