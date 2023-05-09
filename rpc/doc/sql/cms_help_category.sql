create table cms_help_category
(
    id          bigint auto_increment
        primary key,
    name        varchar(100) not null,
    icon        varchar(500) not null comment '分类图标',
    help_count  int          not null comment '专题数量',
    show_status int(2)       not null,
    sort        int          not null
)
    comment '帮助分类表' charset = utf8;

