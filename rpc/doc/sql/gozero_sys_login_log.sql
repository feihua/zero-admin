create table sys_login_log
(
    id               bigint auto_increment comment '编号'
        primary key,
    user_name        varchar(50)                         null comment '用户名',
    status           varchar(50)                         null comment '登录状态（online:在线，登录初始状态，方便统计在线人数；login:退出登录后将online置为login；logout:退出登录）',
    ip               varchar(64)                         null comment 'IP地址',
    create_by        varchar(50)                         null comment '创建人',
    create_time      timestamp default CURRENT_TIMESTAMP null comment '创建时间',
    last_update_by   varchar(50)                         null comment '更新人',
    last_update_time datetime                            null comment '更新时间'
)
    comment '系统登录日志';

INSERT INTO gozero.sys_login_log (id, user_name, status, ip, create_by, create_time, last_update_by, last_update_time) VALUES (1, 'admin', 'login', '0:0:0:0:0:0:0:1', 'admin', '2018-09-23 19:54:16', 'admin', '2018-09-23 19:54:17');
INSERT INTO gozero.sys_login_log (id, user_name, status, ip, create_by, create_time, last_update_by, last_update_time) VALUES (2, 'admin', 'logout', '0:0:0:0:0:0:0:1', 'admin', '2018-09-23 19:54:17', 'admin', '2018-09-23 19:54:17');
INSERT INTO gozero.sys_login_log (id, user_name, status, ip, create_by, create_time, last_update_by, last_update_time) VALUES (3, 'admin', 'login', '0:0:0:0:0:0:0:1', 'admin', '2018-09-23 19:54:18', 'admin', '2018-09-23 19:54:17');