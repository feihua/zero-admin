create table cms_topic
(
    id              bigint auto_increment
        primary key,
    category_id     bigint       not null,
    name            varchar(255) not null,
    create_time     datetime     not null,
    start_time      datetime     not null,
    end_time        datetime     not null,
    attend_count    int          not null comment '参与人数',
    attention_count int          not null comment '关注人数',
    read_count      int          not null,
    award_name      varchar(100) not null comment '奖品名称',
    attend_type     varchar(100) not null comment '参与方式',
    content         text         not null comment '话题内容'
)
    comment '话题表' charset = utf8;

