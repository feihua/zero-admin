create table cms_subject_comment
(
    id               bigint auto_increment
        primary key comment '编号',
    subject_id       bigint        not null comment '关联专题id',
    member_nick_name varchar(255)  not null comment '关联会员昵称',
    member_icon      varchar(255)  not null comment '会员头像',
    content          varchar(1000) not null comment '评论内容',
    create_time      datetime      not null comment '创建时间',
    show_status      tinyint       not null comment '是否显示，0->不显示；1->显示'
)
    comment '专题评论表' charset = utf8;

