create table sys_operate_log
(
    id              bigint auto_increment comment '编号'
        primary key,
    user_name       varchar(50)                           not null comment '用户名',
    operation       varchar(50)                           not null comment '用户操作',
    request_method  varchar(200)                          not null comment '请求方法',
    request_params  text                                  not null comment '请求参数',
    response_params text                                  not null comment '响应参数',
    use_time        bigint                                not null comment '执行时长(毫秒)',
    operation_ip    varchar(64) default ''                not null comment 'IP地址',
    operation_time  timestamp   default CURRENT_TIMESTAMP not null comment '操作时间'
)
    comment '系统操作日志';


