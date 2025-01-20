create table sys_operate_log
(
    id                 bigint auto_increment comment '编号'
        primary key,
    title              varchar(50)                           not null comment '系统模块',
    operation_type     varchar(50)                           not null comment '操作类型',
    operation_name     varchar(50)                           not null comment '操作人员',
    request_method     varchar(200)                          not null comment '请求方式',
    operation_url      varchar(50)                           not null comment '操作方法',
    operation_params   text                                  not null comment '请求参数',
    operation_response text                                  not null comment '响应参数',
    operation_status   tinyint                               not null comment '操作状态',
    dept_name          varchar(50)                           not null comment '部门名称',
    use_time           bigint                                not null comment '执行时长(毫秒)',
    browser            varchar(64)                           not null comment '浏览器',
    os                 varchar(64)                           not null comment '操作系统',
    operation_ip       varchar(64) default ''                not null comment '操作地址',
    operation_time     timestamp   default CURRENT_TIMESTAMP not null comment '操作时间'
)
    comment '系统操作日志表';


