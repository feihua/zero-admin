create table sys_log
(
    id               bigint auto_increment comment '编号'
        primary key,
    user_name        varchar(50)                         null comment '用户名',
    operation        varchar(50)                         null comment '用户操作',
    method           varchar(200)                        null comment '请求方法',
    params           varchar(5000)                       null comment '请求参数',
    time             bigint                              not null comment '执行时长(毫秒)',
    ip               varchar(64)                         null comment 'IP地址',
    create_by        varchar(50)                         null comment '创建人',
    create_time      timestamp default CURRENT_TIMESTAMP null comment '创建时间',
    last_update_by   varchar(50)                         null comment '更新人',
    last_update_time datetime                            null comment '更新时间'
)
    comment '系统操作日志';

INSERT INTO gozero.sys_log (id, user_name, operation, method, params, time, ip, create_by, create_time, last_update_by, last_update_time) VALUES (1, 'admin', '', 'com.louis.kitty.admin.sevice.impl.SysDictServiceImpl.findPage()', '{"columnFilters":{"label":{"name":"label","value":""}},"pageNum":1,"pageSize":8}', 4, '0:0:0:0:0:0:0:1', 'admin', '2018-09-23 19:54:16', '', '2018-09-23 19:54:16');
INSERT INTO gozero.sys_log (id, user_name, operation, method, params, time, ip, create_by, create_time, last_update_by, last_update_time) VALUES (2, 'admin', '', 'com.louis.kitty.admin.sevice.impl.SysRoleServiceImpl.findPage()', '{"columnFilters":{"name":{"name":"name","value":""}},"pageNum":1,"pageSize":8}', 4, '0:0:0:0:0:0:0:1', 'admin', '2018-09-23 19:54:17', '', '2018-09-23 19:54:16');
INSERT INTO gozero.sys_log (id, user_name, operation, method, params, time, ip, create_by, create_time, last_update_by, last_update_time) VALUES (3, 'admin', '', 'com.louis.kitty.admin.sevice.impl.SysUserServiceImpl.findPage()', '{"columnFilters":{"name":{"name":"name","value":""}},"pageNum":1,"pageSize":8}', 36, '0:0:0:0:0:0:0:1', 'admin', '2018-09-23 19:54:18', '', '2018-09-23 19:54:16');
INSERT INTO gozero.sys_log (id, user_name, operation, method, params, time, ip, create_by, create_time, last_update_by, last_update_time) VALUES (4, 'admin', '', 'com.louis.kitty.admin.sevice.impl.SysDictServiceImpl.findPage()', '{"columnFilters":{"label":{"name":"label","value":""}},"pageNum":1,"pageSize":8}', 4, '0:0:0:0:0:0:0:1', 'admin', '2018-09-23 19:54:20', '', '2018-09-23 19:54:16');
INSERT INTO gozero.sys_log (id, user_name, operation, method, params, time, ip, create_by, create_time, last_update_by, last_update_time) VALUES (5, 'admin', '', 'com.louis.kitty.admin.sevice.impl.SysRoleServiceImpl.findPage()', '{"columnFilters":{"name":{"name":"name","value":""}},"pageNum":1,"pageSize":8}', 4, '0:0:0:0:0:0:0:1', 'admin', '2018-09-23 19:54:20', '', '2018-09-23 19:54:16');
INSERT INTO gozero.sys_log (id, user_name, operation, method, params, time, ip, create_by, create_time, last_update_by, last_update_time) VALUES (6, 'admin', '', 'com.louis.kitty.admin.sevice.impl.SysUserServiceImpl.findPage()', '{"columnFilters":{"name":{"name":"name","value":""}},"pageNum":1,"pageSize":8}', 27, '0:0:0:0:0:0:0:1', 'admin', '2018-09-23 19:54:21', '', '2018-09-23 19:54:16');