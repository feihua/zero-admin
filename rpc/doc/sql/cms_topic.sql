create table cms_topic
(
    id              bigint auto_increment
        primary key,
    category_id     bigint       null,
    name            varchar(255) null,
    create_time     datetime     null,
    start_time      datetime     null,
    end_time        datetime     null,
    attend_count    int          null comment '参与人数',
    attention_count int          null comment '关注人数',
    read_count      int          null,
    award_name      varchar(100) null comment '奖品名称',
    attend_type     varchar(100) null comment '参与方式',
    content         text         null comment '话题内容'
)
    comment '话题表' charset = utf8;

