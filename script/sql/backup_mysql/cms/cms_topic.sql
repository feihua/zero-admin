create table cms_topic
(
    id              bigint auto_increment
        primary key comment '主键ID',
    category_id     bigint                                not null comment '关联分类id',
    name            varchar(255)                          not null comment '话题名称',
    start_time      datetime                              not null comment '话题开始时间',
    end_time        datetime                              not null comment '话题结束时间',
    attend_count    int                                   not null comment '参与人数',
    attention_count int                                   not null comment '关注人数',
    read_count      int                                   not null comment '阅读数',
    award_name      varchar(100)                          not null comment '奖品名称',
    attend_type     varchar(100)                          not null comment '参与方式',
    content         text                                  not null comment '话题内容',
    create_by       varchar(50)                           not null comment '创建者',
    create_time     datetime    default CURRENT_TIMESTAMP not null comment '创建时间',
    update_by       varchar(50) default ''                not null comment '更新者',
    update_time     datetime                              null on update CURRENT_TIMESTAMP comment '更新时间'
)
    comment '话题表' charset = utf8;

