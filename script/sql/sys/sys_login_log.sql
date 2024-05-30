create table sys_login_log
(
    id           bigint auto_increment comment '编号'
        primary key,
    user_name   varchar(50)                         not null comment '用户名',
    login_status varchar(50)                         not null comment '登录状态',
    ip_address    varchar(64)                         not null comment 'IP地址',
    browser     varchar(64)                         not null comment '浏览器',
    os     varchar(64)                         not null comment '操作信息',
    login_time   timestamp default CURRENT_TIMESTAMP not null comment '登录时间'
)
    comment '系统登录日志表';
