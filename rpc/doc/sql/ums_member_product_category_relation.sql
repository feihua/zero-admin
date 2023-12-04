create table ums_member_product_category_relation
(
    id                  bigint auto_increment
        primary key,
    member_id           bigint not null comment '会员id',
    product_category_id bigint not null comment '商品分类id'
)
    comment '会员与产品分类关系表（用户喜欢的分类）';

