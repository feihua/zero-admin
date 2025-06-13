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
    create_by     bigint                                 not null comment '创建人ID',
    create_time   datetime     default CURRENT_TIMESTAMP not null comment '创建时间',
    update_by     bigint                                 null comment '更新人ID',
    update_time   datetime                               null on update CURRENT_TIMESTAMP comment '更新时间',
    is_deleted    tinyint      default 0                 not null comment '是否删除'
)
    comment '商品属性表';

create index idx_group
    on pms_product_attribute (group_id, is_deleted);

