create table sys_menu
(
    id               bigint auto_increment comment '编号'
        primary key,
    name             varchar(50)                         null comment '菜单名称',
    parent_id        bigint                              null comment '父菜单ID，一级菜单为0',
    url              varchar(200)                        null comment '菜单URL,类型：1.普通页面（如用户管理， /sys/user） 2.嵌套完整外部页面，以http(s)开头的链接 3.嵌套服务器页面，使用iframe:前缀+目标URL(如SQL监控， iframe:/druid/login.html, iframe:前缀会替换成服务器地址)',
    perms            varchar(500)                        null comment '授权(多个用逗号分隔，如：sys:user:add,sys:user:edit)',
    type             int                                 null comment '类型   0：目录   1：菜单   2：按钮',
    icon             varchar(50)                         null comment '菜单图标',
    order_num        int                                 null comment '排序',
    create_by        varchar(50)                         null comment '创建人',
    create_time      timestamp default CURRENT_TIMESTAMP null comment '创建时间',
    last_update_by   varchar(50)                         null comment '更新人',
    last_update_time datetime                            null comment '更新时间',
    del_flag         tinyint   default 0                 null comment '是否删除  -1：已删除  0：正常'
)
    comment '菜单管理';

INSERT INTO gozero.sys_menu (id, name, parent_id, url, perms, type, icon, order_num, create_by, create_time, last_update_by, last_update_time, del_flag) VALUES (1, '欢迎', 0, '/welcome', '', 0, 'SmileOutlined', 1, 'liufeihua', '2021-02-26 14:45:04', 'liufeihua', '2021-02-26 14:45:04', 0);
INSERT INTO gozero.sys_menu (id, name, parent_id, url, perms, type, icon, order_num, create_by, create_time, last_update_by, last_update_time, del_flag) VALUES (2, '系统管理', 0, '/system', '', 0, 'SettingOutlined', 2, 'liufeihua', '2021-02-26 14:45:04', 'liufeihua', '2021-02-26 14:45:04', 0);
INSERT INTO gozero.sys_menu (id, name, parent_id, url, perms, type, icon, order_num, create_by, create_time, last_update_by, last_update_time, del_flag) VALUES (3, '用户列表', 2, '/system/user/list', '', 1, '', 1, 'liufeihua', '2021-02-26 14:45:04', 'liufeihua', '2021-02-26 14:45:04', 0);
INSERT INTO gozero.sys_menu (id, name, parent_id, url, perms, type, icon, order_num, create_by, create_time, last_update_by, last_update_time, del_flag) VALUES (4, '角色列表', 2, '/system/role/list', '', 1, '', 2, 'liufeihua', '2021-02-26 14:45:04', 'liufeihua', '2021-02-26 14:45:04', 0);
INSERT INTO gozero.sys_menu (id, name, parent_id, url, perms, type, icon, order_num, create_by, create_time, last_update_by, last_update_time, del_flag) VALUES (5, '菜单列表', 2, '/system/menu/list', '', 1, '', 3, 'liufeihua', '2021-02-26 14:45:04', 'liufeihua', '2021-02-26 14:45:04', 0);
INSERT INTO gozero.sys_menu (id, name, parent_id, url, perms, type, icon, order_num, create_by, create_time, last_update_by, last_update_time, del_flag) VALUES (6, '机构列表', 2, '/system/dept/list', '', 1, '', 4, 'liufeihua', '2021-02-26 14:45:04', 'liufeihua', '2021-02-26 14:45:04', 0);
INSERT INTO gozero.sys_menu (id, name, parent_id, url, perms, type, icon, order_num, create_by, create_time, last_update_by, last_update_time, del_flag) VALUES (7, '字典列表', 2, '/system/dict/list', '', 1, '', 5, 'liufeihua', '2021-02-26 14:45:05', 'liufeihua', '2021-02-26 14:45:05', 0);
INSERT INTO gozero.sys_menu (id, name, parent_id, url, perms, type, icon, order_num, create_by, create_time, last_update_by, last_update_time, del_flag) VALUES (8, '日志管理', 0, '/log', '', 0, 'DeleteOutlined', 3, 'liufeihua', '2021-02-26 14:45:05', 'liufeihua', '2021-02-26 14:45:05', 0);
INSERT INTO gozero.sys_menu (id, name, parent_id, url, perms, type, icon, order_num, create_by, create_time, last_update_by, last_update_time, del_flag) VALUES (9, '登录日志', 8, '/log/loginLog/list', '', 1, '', 3, 'liufeihua', '2021-02-26 14:45:05', 'liufeihua', '2021-02-26 14:45:05', 0);
INSERT INTO gozero.sys_menu (id, name, parent_id, url, perms, type, icon, order_num, create_by, create_time, last_update_by, last_update_time, del_flag) VALUES (10, '操作日志', 8, '/log/sysLog/list', '', 1, '', 3, 'liufeihua', '2021-02-26 14:45:05', 'liufeihua', '2021-02-26 14:45:05', 0);