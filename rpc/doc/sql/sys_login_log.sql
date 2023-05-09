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

INSERT INTO sys_login_log (id, user_name, status, ip, create_by, create_time) VALUES (1, 'admin', 'login', '0:0:0:0:0:0:0:1', 'admin', '2018-09-23 19:54:16');
INSERT INTO sys_login_log (id, user_name, status, ip, create_by, create_time) VALUES (2, 'admin', 'logout', '0:0:0:0:0:0:0:1', 'admin', '2018-09-23 19:54:17');
INSERT INTO sys_login_log (id, user_name, status, ip, create_by, create_time) VALUES (3, 'admin', 'login', '0:0:0:0:0:0:0:1', 'admin', '2018-09-23 19:54:18');
INSERT INTO sys_login_log (id, user_name, status, ip, create_by, create_time) VALUES (4, 'admin', 'login', '[::1]:50065', 'admin', '2022-06-22 20:30:34');
INSERT INTO sys_login_log (id, user_name, status, ip, create_by, create_time) VALUES (5, 'admin', 'login', '[::1]:50065', 'admin', '2022-06-22 20:32:05');
INSERT INTO sys_login_log (id, user_name, status, ip, create_by, create_time) VALUES (6, 'admin', 'login', '127.0.0.1:51810', 'admin', '2023-04-14 10:47:41');
INSERT INTO sys_login_log (id, user_name, status, ip, create_by, create_time) VALUES (7, 'admin', 'login', '127.0.0.1:51301', 'admin', '2023-04-14 15:17:11');
INSERT INTO sys_login_log (id, user_name, status, ip, create_by, create_time) VALUES (8, 'admin', 'login', '127.0.0.1:60061', 'admin', '2023-04-17 16:06:05');
INSERT INTO sys_login_log (id, user_name, status, ip, create_by, create_time) VALUES (9, 'admin', 'login', '127.0.0.1:60755', 'admin', '2023-04-17 16:13:35');
INSERT INTO sys_login_log (id, user_name, status, ip, create_by, create_time) VALUES (10, 'admin', 'login', '127.0.0.1:54293', 'admin', '2023-04-20 14:40:27');
INSERT INTO sys_login_log (id, user_name, status, ip, create_by, create_time) VALUES (11, 'admin', 'login', '127.0.0.1:49585', 'admin', '2023-04-21 15:45:50');
INSERT INTO sys_login_log (id, user_name, status, ip, create_by, create_time) VALUES (12, 'admin', 'login', '127.0.0.1:59816', 'admin', '2023-04-23 09:25:20');
INSERT INTO sys_login_log (id, user_name, status, ip, create_by, create_time) VALUES (13, 'admin', 'login', '127.0.0.1:60292', 'admin', '2023-04-24 14:58:02');
INSERT INTO sys_login_log (id, user_name, status, ip, create_by, create_time) VALUES (14, 'admin', 'login', '127.0.0.1:62269', 'admin', '2023-04-25 15:24:23');
