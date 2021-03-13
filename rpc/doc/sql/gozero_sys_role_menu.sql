create table sys_role_menu
(
    id               bigint auto_increment comment '编号'
        primary key,
    role_id          bigint                              null comment '角色ID',
    menu_id          bigint                              null comment '菜单ID',
    create_by        varchar(50)                         null comment '创建人',
    create_time      timestamp default CURRENT_TIMESTAMP null comment '创建时间',
    last_update_by   varchar(50)                         null comment '更新人',
    last_update_time datetime                            null comment '更新时间'
)
    comment '角色菜单';

INSERT INTO gozero.sys_role_menu (id, role_id, menu_id, create_by, create_time, last_update_by, last_update_time) VALUES (638, 1, 1, 'admin', '2021-03-07 18:36:44', 'admin', '2021-03-07 18:36:44');
INSERT INTO gozero.sys_role_menu (id, role_id, menu_id, create_by, create_time, last_update_by, last_update_time) VALUES (639, 1, 2, 'admin', '2021-03-07 18:36:44', 'admin', '2021-03-07 18:36:44');
INSERT INTO gozero.sys_role_menu (id, role_id, menu_id, create_by, create_time, last_update_by, last_update_time) VALUES (640, 1, 3, 'admin', '2021-03-07 18:36:44', 'admin', '2021-03-07 18:36:44');
INSERT INTO gozero.sys_role_menu (id, role_id, menu_id, create_by, create_time, last_update_by, last_update_time) VALUES (641, 1, 4, 'admin', '2021-03-07 18:36:44', 'admin', '2021-03-07 18:36:45');
INSERT INTO gozero.sys_role_menu (id, role_id, menu_id, create_by, create_time, last_update_by, last_update_time) VALUES (642, 1, 5, 'admin', '2021-03-07 18:36:44', 'admin', '2021-03-07 18:36:45');
INSERT INTO gozero.sys_role_menu (id, role_id, menu_id, create_by, create_time, last_update_by, last_update_time) VALUES (643, 1, 6, 'admin', '2021-03-07 18:36:44', 'admin', '2021-03-07 18:36:45');
INSERT INTO gozero.sys_role_menu (id, role_id, menu_id, create_by, create_time, last_update_by, last_update_time) VALUES (644, 1, 7, 'admin', '2021-03-07 18:36:44', 'admin', '2021-03-07 18:36:45');
INSERT INTO gozero.sys_role_menu (id, role_id, menu_id, create_by, create_time, last_update_by, last_update_time) VALUES (645, 1, 8, 'admin', '2021-03-07 18:36:44', 'admin', '2021-03-07 18:36:45');
INSERT INTO gozero.sys_role_menu (id, role_id, menu_id, create_by, create_time, last_update_by, last_update_time) VALUES (646, 1, 9, 'admin', '2021-03-07 18:36:44', 'admin', '2021-03-07 18:36:45');
INSERT INTO gozero.sys_role_menu (id, role_id, menu_id, create_by, create_time, last_update_by, last_update_time) VALUES (647, 1, 10, 'admin', '2021-03-07 18:36:44', 'admin', '2021-03-07 18:36:45');