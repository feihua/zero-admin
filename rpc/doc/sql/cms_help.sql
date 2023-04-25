create table cms_help
(
    id          bigint auto_increment
        primary key,
    category_id bigint       null,
    icon        varchar(500) null,
    title       varchar(100) null,
    show_status int(1)       null,
    create_time datetime     null,
    read_count  int(1)       null,
    content     text         null
)
    comment '帮助表' charset = utf8;

