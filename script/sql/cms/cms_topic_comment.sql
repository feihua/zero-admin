create table cms_topic_comment
(
    id               bigint auto_increment
        primary key comment '主键ID',
    member_nick_name varchar(255)  not null comment '评论人员昵称',
    topic_id         bigint        not null comment '专题ID',
    member_icon      varchar(255)  not null comment '评论人员头像',
    content          varchar(1000) not null comment '评论内容',
    create_time      datetime      not null comment '评论时间',
    show_status      tinyint       not null comment '是否显示，0->不显示；1->显示'
)
    comment '专题评论表' charset = utf8;

