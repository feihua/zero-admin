create table cms_member_report
(
    id                 bigint                                 not null
        primary key comment '编号',
    report_type        tinyint                                not null comment '举报类型：0->商品评价；1->话题内容；2->用户评论',
    report_member_name varchar(100)                           not null comment '举报人',
    report_object      varchar(100)                           not null comment '被举报对象',
    report_status      tinyint                                not null comment '举报状态：0->未处理；1->已处理',
    handle_status      tinyint                                not null comment '处理结果：0->无效；1->有效；2->恶意',
    note               varchar(200) default ''                not null comment '备注',
    create_time        datetime     default CURRENT_TIMESTAMP not null comment '创建时间'
)
    comment '用户举报表' charset = utf8;

