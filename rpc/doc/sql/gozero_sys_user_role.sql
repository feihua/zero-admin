create table sys_user_role
(
    id               bigint auto_increment comment '编号'
        primary key,
    user_id          bigint      null comment '用户ID',
    role_id          bigint      null comment '角色ID',
    create_by        varchar(50) null comment '创建人',
    create_time      datetime    null comment '创建时间',
    last_update_by   varchar(50) null comment '更新人',
    last_update_time datetime    null comment '更新时间'
)
    comment '用户角色';

INSERT INTO gozero.sys_user_role (id, user_id, role_id, create_by, create_time, last_update_by, last_update_time) VALUES (1, 1, 1, null, null, null, null);
INSERT INTO gozero.sys_user_role (id, user_id, role_id, create_by, create_time, last_update_by, last_update_time) VALUES (2, 2, 1, null, null, null, null);
INSERT INTO gozero.sys_user_role (id, user_id, role_id, create_by, create_time, last_update_by, last_update_time) VALUES (26, 5, 3, null, null, null, null);
INSERT INTO gozero.sys_user_role (id, user_id, role_id, create_by, create_time, last_update_by, last_update_time) VALUES (33, 6, 2, null, null, null, null);
INSERT INTO gozero.sys_user_role (id, user_id, role_id, create_by, create_time, last_update_by, last_update_time) VALUES (34, 4, 2, null, null, null, null);
INSERT INTO gozero.sys_user_role (id, user_id, role_id, create_by, create_time, last_update_by, last_update_time) VALUES (35, 9, 2, null, null, null, null);
INSERT INTO gozero.sys_user_role (id, user_id, role_id, create_by, create_time, last_update_by, last_update_time) VALUES (36, 10, 3, null, null, null, null);
INSERT INTO gozero.sys_user_role (id, user_id, role_id, create_by, create_time, last_update_by, last_update_time) VALUES (37, 11, 2, null, null, null, null);
INSERT INTO gozero.sys_user_role (id, user_id, role_id, create_by, create_time, last_update_by, last_update_time) VALUES (38, 12, 3, null, null, null, null);
INSERT INTO gozero.sys_user_role (id, user_id, role_id, create_by, create_time, last_update_by, last_update_time) VALUES (39, 15, 2, null, null, null, null);
INSERT INTO gozero.sys_user_role (id, user_id, role_id, create_by, create_time, last_update_by, last_update_time) VALUES (41, 16, 3, null, null, null, null);
INSERT INTO gozero.sys_user_role (id, user_id, role_id, create_by, create_time, last_update_by, last_update_time) VALUES (42, 8, 2, null, null, null, null);
INSERT INTO gozero.sys_user_role (id, user_id, role_id, create_by, create_time, last_update_by, last_update_time) VALUES (43, 7, 4, null, null, null, null);
INSERT INTO gozero.sys_user_role (id, user_id, role_id, create_by, create_time, last_update_by, last_update_time) VALUES (45, 18, 2, null, null, null, null);
INSERT INTO gozero.sys_user_role (id, user_id, role_id, create_by, create_time, last_update_by, last_update_time) VALUES (46, 17, 3, null, null, null, null);
INSERT INTO gozero.sys_user_role (id, user_id, role_id, create_by, create_time, last_update_by, last_update_time) VALUES (47, 3, 4, null, null, null, null);
INSERT INTO gozero.sys_user_role (id, user_id, role_id, create_by, create_time, last_update_by, last_update_time) VALUES (48, 21, 2, null, null, null, null);
INSERT INTO gozero.sys_user_role (id, user_id, role_id, create_by, create_time, last_update_by, last_update_time) VALUES (57, 31, 2, null, null, null, null);
INSERT INTO gozero.sys_user_role (id, user_id, role_id, create_by, create_time, last_update_by, last_update_time) VALUES (58, 30, 2, null, null, null, null);
INSERT INTO gozero.sys_user_role (id, user_id, role_id, create_by, create_time, last_update_by, last_update_time) VALUES (59, 32, 3, null, null, null, null);
INSERT INTO gozero.sys_user_role (id, user_id, role_id, create_by, create_time, last_update_by, last_update_time) VALUES (73, 33, 8, null, null, null, null);
INSERT INTO gozero.sys_user_role (id, user_id, role_id, create_by, create_time, last_update_by, last_update_time) VALUES (74, 25, 8, null, null, null, null);
INSERT INTO gozero.sys_user_role (id, user_id, role_id, create_by, create_time, last_update_by, last_update_time) VALUES (75, 25, 2, null, null, null, null);
INSERT INTO gozero.sys_user_role (id, user_id, role_id, create_by, create_time, last_update_by, last_update_time) VALUES (80, 22, 2, null, null, null, null);
INSERT INTO gozero.sys_user_role (id, user_id, role_id, create_by, create_time, last_update_by, last_update_time) VALUES (81, 23, 3, null, null, null, null);
INSERT INTO gozero.sys_user_role (id, user_id, role_id, create_by, create_time, last_update_by, last_update_time) VALUES (82, 24, 4, null, null, null, null);
INSERT INTO gozero.sys_user_role (id, user_id, role_id, create_by, create_time, last_update_by, last_update_time) VALUES (83, 26, 3, null, null, null, null);
INSERT INTO gozero.sys_user_role (id, user_id, role_id, create_by, create_time, last_update_by, last_update_time) VALUES (85, 29, 2, null, null, null, null);
INSERT INTO gozero.sys_user_role (id, user_id, role_id, create_by, create_time, last_update_by, last_update_time) VALUES (86, 28, 4, null, null, null, null);
INSERT INTO gozero.sys_user_role (id, user_id, role_id, create_by, create_time, last_update_by, last_update_time) VALUES (87, 27, 3, null, null, null, null);