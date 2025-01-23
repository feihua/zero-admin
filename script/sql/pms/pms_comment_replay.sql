create table pms_comment_replay
(
    id               bigint auto_increment
        primary key,
    comment_id       bigint        not null comment '评论id',
    member_nick_name varchar(255)  not null comment '评论人员昵称',
    member_icon      varchar(255)  not null comment '评论人员头像',
    content          varchar(1000) not null comment '内容',
    create_time      datetime      not null comment '评论时间',
    type             tinyint       not null comment '评论人员类型；0->会员；1->管理员'
)
    comment '产品评价回复表';

