create table sys_login_log
(
    id           bigint auto_increment comment '编号'
        primary key,
    login_name   varchar(50)                         not null comment '用户名',
    login_status varchar(50)                         not null comment '登录状态',
    login_ip     varchar(64)                         not null comment 'IP地址',
    login_time   timestamp default CURRENT_TIMESTAMP not null comment '登录时间'
)
    comment '系统登录日志';
