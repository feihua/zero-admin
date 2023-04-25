create table cms_help_category
(
    id          bigint auto_increment
        primary key,
    name        varchar(100) null,
    icon        varchar(500) null comment '分类图标',
    help_count  int          null comment '专题数量',
    show_status int(2)       null,
    sort        int          null
)
    comment '帮助分类表' charset = utf8;

