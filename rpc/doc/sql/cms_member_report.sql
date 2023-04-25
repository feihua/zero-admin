create table cms_member_report
(
    id                 bigint       not null
        primary key,
    report_type        int(1)       null comment '举报类型：0->商品评价；1->话题内容；2->用户评论',
    report_member_name varchar(100) null comment '举报人',
    create_time        datetime     null,
    report_object      varchar(100) null,
    report_status      int(1)       null comment '举报状态：0->未处理；1->已处理',
    handle_status      int(1)       null comment '处理结果：0->无效；1->有效；2->恶意',
    note               varchar(200) null
)
    comment '用户举报表' charset = utf8;

