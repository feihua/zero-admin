create table sys_login_log
(
    id          bigint auto_increment comment '编号'
        primary key,
    user_name   varchar(50)                         not null comment '用户名',
    status      varchar(50)                         not null comment '登录状态（online:在线，登录初始状态，方便统计在线人数；login:退出登录后将online置为login；logout:退出登录）',
    ip          varchar(64)                         not null comment 'IP地址',
    create_by   varchar(50)                         not null comment '创建人',
    create_time timestamp default CURRENT_TIMESTAMP not null comment '创建时间'
)
    comment '系统登录日志';
