create table pms_comment_replay
(
    id               bigint auto_increment
        primary key,
    comment_id       bigint        null,
    member_nick_name varchar(255)  null,
    member_icon      varchar(255)  null,
    content          varchar(1000) null,
    create_time      datetime      null,
    type             int(1)        null comment '评论人员类型；0->会员；1->管理员'
)
    comment '产品评价回复表';

