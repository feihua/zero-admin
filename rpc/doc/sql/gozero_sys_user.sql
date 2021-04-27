create table sys_user
(
    id               bigint auto_increment comment '编号'
        primary key,
    name             varchar(50)                         not null comment '用户名',
    nick_name        varchar(150)                        null comment '昵称',
    avatar           varchar(150)                        null comment '头像',
    password         varchar(100)                        null comment '密码',
    salt             varchar(40)                         null comment '加密盐',
    email            varchar(100)                        null comment '邮箱',
    mobile           varchar(100)                        null comment '手机号',
    status           tinyint                             null comment '状态  0：禁用   1：正常',
    dept_id          bigint                              null comment '机构ID',
    create_by        varchar(50)                         null comment '创建人',
    create_time      timestamp default CURRENT_TIMESTAMP null comment '创建时间',
    last_update_by   varchar(50)                         null comment '更新人',
    last_update_time datetime                            null comment '更新时间',
    del_flag         tinyint   default 0                 null comment '是否删除  -1：已删除  0：正常',
    job_id           int                                 null comment '岗位Id',
    constraint name
        unique (name)
)
    comment '用户管理';

INSERT INTO gozero.sys_user (id, name, nick_name, avatar, password, salt, email, mobile, status, dept_id, create_by, create_time, last_update_by, last_update_time, del_flag, job_id) VALUES (1, 'admin', '超管', '', '123456', 'YzcmCZNvbXocrsz9dm8e', 'admin@qq.com', '13612345678', 1, 4, 'admin', '2018-08-14 11:11:11', 'admin', '2018-08-14 11:11:11', 0, 1);
INSERT INTO gozero.sys_user (id, name, nick_name, avatar, password, salt, email, mobile, status, dept_id, create_by, create_time, last_update_by, last_update_time, del_flag, job_id) VALUES (22, 'liubei', '刘备', '', '123456', 'YzcmCZNvbXocrsz9dm8e', 'test@qq.com', '13889700023', 1, 7, 'admin', '2018-09-23 19:43:00', 'admin', '2019-01-10 11:41:13', 0, 1);
INSERT INTO gozero.sys_user (id, name, nick_name, avatar, password, salt, email, mobile, status, dept_id, create_by, create_time, last_update_by, last_update_time, del_flag, job_id) VALUES (23, 'zhaoyun', '赵云', '', '123456', 'YzcmCZNvbXocrsz9dm8e', 'test@qq.com', '13889700023', 1, 7, 'admin', '2018-09-23 19:43:44', 'admin', '2018-09-23 19:43:52', 0, 1);
INSERT INTO gozero.sys_user (id, name, nick_name, avatar, password, salt, email, mobile, status, dept_id, create_by, create_time, last_update_by, last_update_time, del_flag, job_id) VALUES (24, 'zhugeliang', '诸葛亮', '', '123456', 'YzcmCZNvbXocrsz9dm8e', 'test@qq.com', '13889700023', 7, 11, 'admin', '2018-09-23 19:44:23', 'admin', '2018-09-23 19:44:29', 0, 1);
INSERT INTO gozero.sys_user (id, name, nick_name, avatar, password, salt, email, mobile, status, dept_id, create_by, create_time, last_update_by, last_update_time, del_flag, job_id) VALUES (25, 'caocao', '曹操', '', '123456', 'YzcmCZNvbXocrsz9dm8e', 'test@qq.com', '13889700023', 1, 8, 'admin', '2018-09-23 19:45:32', 'admin', '2019-01-10 17:59:14', 0, 1);
INSERT INTO gozero.sys_user (id, name, nick_name, avatar, password, salt, email, mobile, status, dept_id, create_by, create_time, last_update_by, last_update_time, del_flag, job_id) VALUES (26, 'dianwei', '典韦', '', '123456', 'YzcmCZNvbXocrsz9dm8e', 'test@qq.com', '13889700023', 1, 10, 'admin', '2018-09-23 19:45:48', 'admin', '2018-09-23 19:45:57', 0, 1);
INSERT INTO gozero.sys_user (id, name, nick_name, avatar, password, salt, email, mobile, status, dept_id, create_by, create_time, last_update_by, last_update_time, del_flag, job_id) VALUES (27, 'xiahoudun', '夏侯惇', '', '123456', 'YzcmCZNvbXocrsz9dm8e', 'test@qq.com', '13889700023', 1, 8, 'admin', '2018-09-23 19:46:09', 'admin', '2018-09-23 19:46:17', 0, 1);
INSERT INTO gozero.sys_user (id, name, nick_name, avatar, password, salt, email, mobile, status, dept_id, create_by, create_time, last_update_by, last_update_time, del_flag, job_id) VALUES (28, 'xunyu', '荀彧', '', '123456', 'YzcmCZNvbXocrsz9dm8e', 'test@qq.com', '13889700023', 1, 10, 'admin', '2018-09-23 19:46:38', 'admin', '2018-11-04 15:33:17', 0, 1);
INSERT INTO gozero.sys_user (id, name, nick_name, avatar, password, salt, email, mobile, status, dept_id, create_by, create_time, last_update_by, last_update_time, del_flag, job_id) VALUES (29, 'sunquan', '孙权', '', '123456', 'YzcmCZNvbXocrsz9dm8e', 'test@qq.com', '13889700023', 1, 10, 'admin', '2018-09-23 19:46:54', 'admin', '2018-09-23 19:47:03', 0, 1);
INSERT INTO gozero.sys_user (id, name, nick_name, avatar, password, salt, email, mobile, status, dept_id, create_by, create_time, last_update_by, last_update_time, del_flag, job_id) VALUES (30, 'zhouyu', '周瑜', '', '123456', 'YzcmCZNvbXocrsz9dm8e', 'test@qq.com', '13889700023', 1, 11, 'admin', '2018-09-23 19:47:28', 'admin', '2018-09-23 19:48:04', 0, 1);
INSERT INTO gozero.sys_user (id, name, nick_name, avatar, password, salt, email, mobile, status, dept_id, create_by, create_time, last_update_by, last_update_time, del_flag, job_id) VALUES (31, 'luxun', '陆逊', '', '123456', 'YzcmCZNvbXocrsz9dm8e', 'test@qq.com', '13889700023', 1, 11, 'admin', '2018-09-23 19:47:44', 'admin', '2018-09-23 19:47:58', 0, 1);
INSERT INTO gozero.sys_user (id, name, nick_name, avatar, password, salt, email, mobile, status, dept_id, create_by, create_time, last_update_by, last_update_time, del_flag, job_id) VALUES (32, 'huanggai', '黄盖', '', '', '', 'test@qq.com', '13889700023', 1, 11, '', '2018-09-23 19:48:38', 'admin', '2021-04-03 15:43:52', 0, 1);
INSERT INTO gozero.sys_user (id, name, nick_name, avatar, password, salt, email, mobile, status, dept_id, create_by, create_time, last_update_by, last_update_time, del_flag, job_id) VALUES (33, '1', '1', '', '123456', '123456', '1', '1', 1, 2, 'admin', '2021-04-26 17:57:50', 'admin', '2021-04-26 17:57:50', 0, 1);
INSERT INTO gozero.sys_user (id, name, nick_name, avatar, password, salt, email, mobile, status, dept_id, create_by, create_time, last_update_by, last_update_time, del_flag, job_id) VALUES (35, '12', '1', '', '123456', '123456', '1', '1', 1, 2, 'admin', '2021-04-26 18:01:53', 'admin', '2021-04-26 18:01:54', 0, 1);
INSERT INTO gozero.sys_user (id, name, nick_name, avatar, password, salt, email, mobile, status, dept_id, create_by, create_time, last_update_by, last_update_time, del_flag, job_id) VALUES (36, '12313', '12', '', '123456', '123456', '1', '1', 1, 2, 'admin', '2021-04-26 18:03:07', 'admin', '2021-04-26 18:03:07', 0, 1);
INSERT INTO gozero.sys_user (id, name, nick_name, avatar, password, salt, email, mobile, status, dept_id, create_by, create_time, last_update_by, last_update_time, del_flag, job_id) VALUES (37, '324', '1', '', '123456', '123456', '1', '1', 1, 3, 'admin', '2021-04-26 18:07:31', 'admin', '2021-04-26 18:07:32', 0, 1);
INSERT INTO gozero.sys_user (id, name, nick_name, avatar, password, salt, email, mobile, status, dept_id, create_by, create_time, last_update_by, last_update_time, del_flag, job_id) VALUES (38, 'aa', 'aa', '', '123456', '123456', 'a', 'a', 1, 7, 'admin', '2021-04-27 11:24:14', 'admin', '2021-04-27 11:24:14', 0, 4);
INSERT INTO gozero.sys_user (id, name, nick_name, avatar, password, salt, email, mobile, status, dept_id, create_by, create_time, last_update_by, last_update_time, del_flag, job_id) VALUES (39, '133', '133', '', '', '', '1121', '1', 1, 2, 'admin', '2021-04-27 12:30:15', 'admin', '2021-04-27 13:53:40', 0, 1);
INSERT INTO gozero.sys_user (id, name, nick_name, avatar, password, salt, email, mobile, status, dept_id, create_by, create_time, last_update_by, last_update_time, del_flag, job_id) VALUES (40, 'liu', 'liu', '', '123456', '123456', '1002219331@qq.com', '18613030352', 1, 8, 'admin', '2021-04-27 13:47:42', 'admin', '2021-04-27 13:47:42', 0, 4);