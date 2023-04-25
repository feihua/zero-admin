create table cms_subject_comment
(
    id               bigint auto_increment
        primary key,
    subject_id       bigint        null,
    member_nick_name varchar(255)  null,
    member_icon      varchar(255)  null,
    content          varchar(1000) null,
    create_time      datetime      null,
    show_status      int(1)        null
)
    comment '专题评论表' charset = utf8;

