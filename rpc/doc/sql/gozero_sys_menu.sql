create table sys_menu
(
    id               bigint auto_increment comment '编号'
        primary key,
    name             varchar(50)                         null comment '菜单名称',
    parent_id        bigint                              null comment '父菜单ID，一级菜单为0',
    url              varchar(200)                        null,
    perms            varchar(500)                        null comment '授权(多个用逗号分隔，如：sys:user:add,sys:user:edit)',
    type             int                                 null comment '类型   0：目录   1：菜单   2：按钮',
    icon             varchar(50)                         null comment '菜单图标',
    order_num        int                                 null comment '排序',
    create_by        varchar(50)                         null comment '创建人',
    create_time      timestamp default CURRENT_TIMESTAMP null comment '创建时间',
    last_update_by   varchar(50)                         null comment '更新人',
    last_update_time datetime                            null comment '更新时间',
    del_flag         tinyint   default 0                 null comment '是否删除  -1：已删除  0：正常',
    vue_path         varchar(64)                         null comment 'vue系统的path',
    vue_component    varchar(64)                         null comment 'vue的页面',
    vue_icon         varchar(64)                         null comment 'vue的图标',
    vue_redirect     varchar(64)                         null comment 'vue的路由重定向'
)
    comment '菜单管理';

INSERT INTO gozero.sys_menu (id, name, parent_id, url, perms, type, icon, order_num, create_by, create_time, last_update_by, last_update_time, del_flag, vue_path, vue_component, vue_icon, vue_redirect) VALUES (1, '欢迎', 0, '/welcome', '', 0, 'SmileOutlined', 1, 'liufeihua', '2021-02-26 14:45:04', 'liufeihua', '2021-02-26 14:45:04', 0, '', null, null, null);
INSERT INTO gozero.sys_menu (id, name, parent_id, url, perms, type, icon, order_num, create_by, create_time, last_update_by, last_update_time, del_flag, vue_path, vue_component, vue_icon, vue_redirect) VALUES (2, '系统管理', 0, '/system', '', 0, 'SettingOutlined', 2, 'liufeihua', '2021-02-26 14:45:04', 'liufeihua', '2021-02-26 14:45:04', 0, '/sys', 'Layout', 'el-icon-setting', '/sys/userList');
INSERT INTO gozero.sys_menu (id, name, parent_id, url, perms, type, icon, order_num, create_by, create_time, last_update_by, last_update_time, del_flag, vue_path, vue_component, vue_icon, vue_redirect) VALUES (3, '用户列表', 2, '/system/user/list', '', 1, '', 1, 'liufeihua', '2021-02-26 14:45:04', 'liufeihua', '2021-02-26 14:45:04', 0, 'userList', 'system/user/index', 'el-icon-user', null);
INSERT INTO gozero.sys_menu (id, name, parent_id, url, perms, type, icon, order_num, create_by, create_time, last_update_by, last_update_time, del_flag, vue_path, vue_component, vue_icon, vue_redirect) VALUES (4, '角色列表', 2, '/system/role/list', '', 1, '', 2, 'liufeihua', '2021-02-26 14:45:04', 'liufeihua', '2021-02-26 14:45:04', 0, 'roleList', 'system/role/index', 'el-icon-female', null);
INSERT INTO gozero.sys_menu (id, name, parent_id, url, perms, type, icon, order_num, create_by, create_time, last_update_by, last_update_time, del_flag, vue_path, vue_component, vue_icon, vue_redirect) VALUES (5, '菜单列表', 2, '/system/menu/list', '', 1, '', 3, 'liufeihua', '2021-02-26 14:45:04', 'liufeihua', '2021-02-26 14:45:04', 0, 'resourcesList', 'system/resource/index', 'el-icon-postcard', null);
INSERT INTO gozero.sys_menu (id, name, parent_id, url, perms, type, icon, order_num, create_by, create_time, last_update_by, last_update_time, del_flag, vue_path, vue_component, vue_icon, vue_redirect) VALUES (6, '机构列表', 2, '/system/dept/list', '', 1, '', 4, 'liufeihua', '2021-02-26 14:45:04', 'liufeihua', '2021-02-26 14:45:04', 0, 'deptList', 'system/dept/index', 'el-icon-bangzhu', null);
INSERT INTO gozero.sys_menu (id, name, parent_id, url, perms, type, icon, order_num, create_by, create_time, last_update_by, last_update_time, del_flag, vue_path, vue_component, vue_icon, vue_redirect) VALUES (7, '字典列表', 2, '/system/dict/list', '', 1, '', 5, 'liufeihua', '2021-02-26 14:45:05', 'liufeihua', '2021-02-26 14:45:05', 0, 'dictList', 'system/dict/index', 'el-icon-delete-location', null);
INSERT INTO gozero.sys_menu (id, name, parent_id, url, perms, type, icon, order_num, create_by, create_time, last_update_by, last_update_time, del_flag, vue_path, vue_component, vue_icon, vue_redirect) VALUES (8, '日志管理', 0, '/log', '', 0, 'DeleteOutlined', 3, 'liufeihua', '2021-02-26 14:45:05', 'liufeihua', '2021-02-26 14:45:05', 0, '/log', 'Layout', 'el-icon-delete', '/log/loginLogList');
INSERT INTO gozero.sys_menu (id, name, parent_id, url, perms, type, icon, order_num, create_by, create_time, last_update_by, last_update_time, del_flag, vue_path, vue_component, vue_icon, vue_redirect) VALUES (9, '登录日志', 8, '/log/loginLog/list', '', 1, '', 3, 'liufeihua', '2021-02-26 14:45:05', 'liufeihua', '2021-02-26 14:45:05', 0, 'loginLogList', 'log/loginlog/index', 'el-icon-remove-outline', null);
INSERT INTO gozero.sys_menu (id, name, parent_id, url, perms, type, icon, order_num, create_by, create_time, last_update_by, last_update_time, del_flag, vue_path, vue_component, vue_icon, vue_redirect) VALUES (10, '操作日志', 8, '/log/sysLog/list', '', 1, '', 3, 'liufeihua', '2021-02-26 14:45:05', 'liufeihua', '2021-02-26 14:45:05', 0, 'sysLogList', 'log/syslog/index', 'el-icon-camera', null);