create table cms_help
(
    id          bigint auto_increment
        primary key,
    category_id bigint                             not null,
    icon        varchar(500)                       not null,
    title       varchar(100)                       not null,
    show_status tinyint                            not null,
    read_count  int                                not null,
    content     text                               not null,
    create_by   varchar(50)                        not null comment '创建者',
    create_time datetime default CURRENT_TIMESTAMP not null comment '创建时间',
    update_by   varchar(50)                        null comment '更新者',
    update_time datetime                           null on update CURRENT_TIMESTAMP comment '更新时间'
)
    comment '帮助表' charset = utf8;

