create table cms_help
(
    id          bigint auto_increment
        primary key comment '主键ID',
    category_id bigint                                not null comment '分类ID',
    icon        varchar(500)                          not null comment '图标',
    title       varchar(100)                          not null comment '标题',
    show_status tinyint                               not null comment '显示状态：0->不显示；1->显示',
    read_count  int                                   not null comment '阅读量',
    content     text                                  not null comment '内容',
    create_by   varchar(50)                           not null comment '创建者',
    create_time datetime    default CURRENT_TIMESTAMP not null comment '创建时间',
    update_by   varchar(50) default ''                not null comment '更新者',
    update_time datetime                              null on update CURRENT_TIMESTAMP comment '更新时间'
)
    comment '帮助表' charset = utf8;

