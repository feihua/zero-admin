drop table if exists ums_member_tag;
create table ums_member_tag
(
    id                  bigint auto_increment comment '主键ID'
        primary key,
    tag_name            varchar(100)                             not null comment '标签名称',
    description         varchar(255)   default ''                not null comment '标签描述',
    finish_order_count  int            default 0                 not null comment '自动打标签完成订单数量',
    finish_order_amount decimal(10, 2) default 0.00              not null comment '自动打标签完成订单金额',
    status              tinyint     default 1                 not null comment '状态：0-禁用，1-启用',
    create_by           bigint                                   not null comment '创建人ID',
    create_time         datetime       default CURRENT_TIMESTAMP not null comment '创建时间',
    update_by           bigint                                   null comment '更新人ID',
    update_time         datetime                                 null on update CURRENT_TIMESTAMP comment '更新时间',
    is_deleted          tinyint        default 0                 not null comment '是否删除',
    constraint uk_tag_name
        unique (tag_name, is_deleted)
)
    comment '用户标签表';


-- 插入用户标签数据
insert into ums_member_tag (id, tag_name, description, finish_order_count, finish_order_amount, status, create_by, is_deleted)
values (1, '新用户', '注册时间不超过一个月的用户', 0, 0.00, 1, 1, 0),
       (2, '忠实用户', '连续三个月每月都有订单的用户', 3, 300.00, 1, 1, 0),
       (3, '活跃用户', '最近一个月内完成5单以上的用户', 5, 500.00, 1, 1, 0),
       (4, '高消费用户', '累计消费金额超过1000元的用户', 10, 1000.00, 1, 1, 0);


drop table if exists ums_member_tag_relation;
create table ums_member_tag_relation
(
    id        bigint auto_increment comment '主键ID'
        primary key,
    member_id bigint not null comment '会员ID',
    tag_id    bigint not null comment '标签ID',
    create_time datetime default CURRENT_TIMESTAMP not null comment '创建时间'
)
    comment '会员标签关联表';

# 更新会员标签关联表的时机

# - 1.实时更新 ：
# - 当会员的行为（如下单、消费、签到等）发生变化时，立即更新会员的标签。

# - 2.定时任务 ：
# - 使用定时任务（如每天凌晨）批量更新会员的标签。

