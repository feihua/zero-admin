drop table if exists pms_product_attribute_value;

create table pms_product_attribute_value
(
    id           bigint auto_increment comment '主键id'
        primary key,
    spu_id       bigint                             not null comment '商品SPU ID',
    attribute_id bigint                             not null comment '属性ID',
    value        varchar(500)                       not null comment '属性值',
    create_by    bigint                             not null comment '创建人ID',
    create_time  datetime default CURRENT_TIMESTAMP not null comment '创建时间',
    update_by    bigint                             null comment '更新人ID',
    update_time  datetime                           null on update CURRENT_TIMESTAMP comment '更新时间',
    is_deleted   tinyint  default 0                 not null comment '是否删除'
)
    comment '商品属性值表';

create index idx_attribute
    on pms_product_attribute_value (attribute_id, is_deleted);

create index idx_spu
    on pms_product_attribute_value (spu_id, is_deleted);

