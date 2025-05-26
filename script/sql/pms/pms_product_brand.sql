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
values (1, '华为', 'http://129.204.203.29/big.png', 'http://129.204.203.29/big.png', '华为的描述', 'H', 1, 1, 100, 50, 1, 1),
       (2, '苹果', 'http://129.204.203.29/big.png', 'http://129.204.203.29/big.png', '苹果的描述', 'A', 2, 0, 200, 80, 1, 2),
       (3, '耐克', 'http://129.204.203.29/big.png', 'http://129.204.203.29/big.png', '耐克的描述', 'N', 3, 1, 150, 60, 1, 3),
       (4, '阿迪达斯', 'http://129.204.203.29/big.png', 'http://129.204.203.29/big.png', '阿迪达斯的描述', 'A', 4, 0, 120, 40, 1, 4),
       (5, '雀巢', 'http://129.204.203.29/big.png', 'http://129.204.203.29/big.png', '雀巢的描述', 'Q', 5, 1, 180, 70, 1, 5),
       (6, '可口可乐', 'http://129.204.203.29/big.png', 'http://129.204.203.29/big.png', '可口可乐的描述', 'K', 6, 0, 90, 30, 1, 6),
       (7, '索尼', 'http://129.204.203.29/big.png', 'http://129.204.203.29/big.png', '索尼的描述', 'S', 7, 1, 110, 50, 1, 7),
       (8, '海尔', 'http://129.204.203.29/big.png', 'http://129.204.203.29/big.png', '海尔的描述', 'H', 8, 0, 130, 60, 1, 8),
       (9, '欧莱雅', 'http://129.204.203.29/big.png', 'http://129.204.203.29/big.png', '欧莱雅的描述', 'O', 9, 1, 140, 55, 1, 9),
       (10, '雅诗兰黛', 'http://129.204.203.29/big.png', 'http://129.204.203.29/big.png', '雅诗兰黛的描述', 'Y', 10, 0, 160, 65, 1, 10),
       (11, '特斯拉', 'http://129.204.203.29/big.png', 'http://129.204.203.29/big.png', '特斯拉的描述', 'T', 11, 1, 170, 75, 1, 11);