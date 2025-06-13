drop table if exists pms_product_attribute_group;
create table pms_product_attribute_group
(
    id          bigint auto_increment comment '主键id'
        primary key,
    category_id bigint                             not null comment '分类ID',
    name        varchar(50)                        not null comment '分组名称',
    sort        int      default 0                 not null comment '排序',
    create_by   bigint                             not null comment '创建人ID',
    create_time datetime default CURRENT_TIMESTAMP not null comment '创建时间',
    update_by   bigint                             null comment '更新人ID',
    update_time datetime                           null on update CURRENT_TIMESTAMP comment '更新时间',
    is_deleted  tinyint  default 0                 not null comment '是否删除'
)
    comment '商品属性分组表';

create index idx_category
    on pms_product_attribute_group (category_id, is_deleted);

