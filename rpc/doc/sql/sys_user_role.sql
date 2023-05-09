create table sys_user_role
(
    id          bigint auto_increment comment '编号'
        primary key,
    user_id     bigint      not null comment '用户ID',
    role_id     bigint      not null comment '角色ID',
    create_by   varchar(50) not null comment '创建人',
    create_time datetime    not null comment '创建时间'
)
    comment '用户角色';

INSERT INTO sys_user_role (id, user_id, role_id, create_by, create_time) VALUES (1, 1, 1, ' ', '2023-05-09 15:24:12');
INSERT INTO sys_user_role (id, user_id, role_id, create_by, create_time) VALUES (2, 2, 1, ' ', '2023-05-09 15:24:17');
INSERT INTO sys_user_role (id, user_id, role_id, create_by, create_time) VALUES (26, 5, 3, ' ', '2023-05-09 15:24:19');
INSERT INTO sys_user_role (id, user_id, role_id, create_by, create_time) VALUES (33, 6, 2, ' ', '2023-05-09 15:24:21');
INSERT INTO sys_user_role (id, user_id, role_id, create_by, create_time) VALUES (34, 4, 2, ' ', '2023-05-09 15:24:21');
INSERT INTO sys_user_role (id, user_id, role_id, create_by, create_time) VALUES (35, 9, 2, ' ', '2023-05-09 15:24:21');
INSERT INTO sys_user_role (id, user_id, role_id, create_by, create_time) VALUES (36, 10, 3, ' ', '2023-05-09 15:24:21');
INSERT INTO sys_user_role (id, user_id, role_id, create_by, create_time) VALUES (37, 11, 2, ' ', '2023-05-09 15:24:21');
INSERT INTO sys_user_role (id, user_id, role_id, create_by, create_time) VALUES (38, 12, 3, ' ', '2023-05-09 15:24:21');
INSERT INTO sys_user_role (id, user_id, role_id, create_by, create_time) VALUES (39, 15, 2, ' ', '2023-05-09 15:24:21');
INSERT INTO sys_user_role (id, user_id, role_id, create_by, create_time) VALUES (41, 16, 3, ' ', '2023-05-09 15:24:21');
INSERT INTO sys_user_role (id, user_id, role_id, create_by, create_time) VALUES (42, 8, 2, ' ', '2023-05-09 15:24:21');
INSERT INTO sys_user_role (id, user_id, role_id, create_by, create_time) VALUES (43, 7, 4, ' ', '2023-05-09 15:24:21');
INSERT INTO sys_user_role (id, user_id, role_id, create_by, create_time) VALUES (45, 18, 2, ' ', '2023-05-09 15:24:21');
INSERT INTO sys_user_role (id, user_id, role_id, create_by, create_time) VALUES (46, 17, 3, ' ', '2023-05-09 15:24:21');
INSERT INTO sys_user_role (id, user_id, role_id, create_by, create_time) VALUES (47, 3, 4, ' ', '2023-05-09 15:24:21');
INSERT INTO sys_user_role (id, user_id, role_id, create_by, create_time) VALUES (48, 21, 2, ' ', '2023-05-09 15:24:21');
INSERT INTO sys_user_role (id, user_id, role_id, create_by, create_time) VALUES (57, 31, 2, ' ', '2023-05-09 15:24:21');
INSERT INTO sys_user_role (id, user_id, role_id, create_by, create_time) VALUES (58, 30, 2, ' ', '2023-05-09 15:24:21');
INSERT INTO sys_user_role (id, user_id, role_id, create_by, create_time) VALUES (59, 32, 3, ' ', '2023-05-09 15:24:21');
INSERT INTO sys_user_role (id, user_id, role_id, create_by, create_time) VALUES (73, 33, 8, ' ', '2023-05-09 15:24:21');
INSERT INTO sys_user_role (id, user_id, role_id, create_by, create_time) VALUES (74, 25, 8, ' ', '2023-05-09 15:24:21');
INSERT INTO sys_user_role (id, user_id, role_id, create_by, create_time) VALUES (75, 25, 2, ' ', '2023-05-09 15:24:21');
INSERT INTO sys_user_role (id, user_id, role_id, create_by, create_time) VALUES (80, 22, 2, ' ', '2023-05-09 15:24:21');
INSERT INTO sys_user_role (id, user_id, role_id, create_by, create_time) VALUES (81, 23, 3, ' ', '2023-05-09 15:24:21');
INSERT INTO sys_user_role (id, user_id, role_id, create_by, create_time) VALUES (82, 24, 4, ' ', '2023-05-09 15:24:21');
INSERT INTO sys_user_role (id, user_id, role_id, create_by, create_time) VALUES (83, 26, 3, ' ', '2023-05-09 15:24:21');
INSERT INTO sys_user_role (id, user_id, role_id, create_by, create_time) VALUES (85, 29, 2, ' ', '2023-05-09 15:24:21');
INSERT INTO sys_user_role (id, user_id, role_id, create_by, create_time) VALUES (86, 28, 4, ' ', '2023-05-09 15:24:21');
INSERT INTO sys_user_role (id, user_id, role_id, create_by, create_time) VALUES (87, 27, 3, ' ', '2023-05-09 15:24:21');
