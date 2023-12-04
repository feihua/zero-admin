create table ums_member_member_tag_relation
(
    id        bigint auto_increment
        primary key,
    member_id bigint not null comment '会员id',
    tag_id    bigint not null comment '标签id'
)
    comment '用户和标签关系表';

