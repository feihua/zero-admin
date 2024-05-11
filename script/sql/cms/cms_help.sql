create table cms_help
(
    id          bigint auto_increment
        primary key,
    category_id bigint       not null,
    icon        varchar(500) not null,
    title       varchar(100) not null,
    show_status int(1)       not null,
    create_time datetime     not null,
    read_count  int(1)       not null,
    content     text         not null
)
    comment '帮助表' charset = utf8;

