create table cms_subject_comment
(
    id               bigint auto_increment
        primary key,
    subject_id       bigint        not null,
    member_nick_name varchar(255)  not null,
    member_icon      varchar(255)  not null,
    content          varchar(1000) not null,
    create_time      datetime      not null,
    show_status tinyint not null
)
    comment '专题评论表' charset = utf8;

