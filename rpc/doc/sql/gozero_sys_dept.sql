create table sys_dept
(
    id               bigint auto_increment comment '编号'
        primary key,
    name             varchar(50)                         null comment '机构名称',
    parent_id        bigint                              null comment '上级机构ID，一级机构为0',
    order_num        int                                 null comment '排序',
    create_by        varchar(50)                         null comment '创建人',
    create_time      timestamp default CURRENT_TIMESTAMP null comment '创建时间',
    last_update_by   varchar(50)                         null comment '更新人',
    last_update_time datetime                            null comment '更新时间',
    del_flag         tinyint   default 0                 null comment '是否删除  -1：已删除  0：正常'
)
    comment '机构管理';

INSERT INTO gozero.sys_dept (id, name, parent_id, order_num, create_by, create_time, last_update_by, last_update_time, del_flag) VALUES (1, '轻尘集团', 0, 0, 'admin', '2018-09-23 19:35:22', 'admin', '2018-09-23 19:35:22', 0);
INSERT INTO gozero.sys_dept (id, name, parent_id, order_num, create_by, create_time, last_update_by, last_update_time, del_flag) VALUES (2, '牧尘集团', 0, 1, 'admin', '2018-09-23 19:35:55', 'admin', '2018-09-23 19:35:22', 0);
INSERT INTO gozero.sys_dept (id, name, parent_id, order_num, create_by, create_time, last_update_by, last_update_time, del_flag) VALUES (3, '三国集团', 0, 2, 'admin', '2018-09-23 19:36:24', 'admin', '2018-09-23 19:35:22', 0);
INSERT INTO gozero.sys_dept (id, name, parent_id, order_num, create_by, create_time, last_update_by, last_update_time, del_flag) VALUES (4, '上海分公司', 2, 0, 'admin', '2018-09-23 19:37:03', 'admin', '2018-09-23 19:35:22', 0);
INSERT INTO gozero.sys_dept (id, name, parent_id, order_num, create_by, create_time, last_update_by, last_update_time, del_flag) VALUES (5, '北京分公司', 1, 1, 'admin', '2018-09-23 19:37:17', 'admin', '2018-09-23 19:35:22', 0);
INSERT INTO gozero.sys_dept (id, name, parent_id, order_num, create_by, create_time, last_update_by, last_update_time, del_flag) VALUES (6, '北京分公司', 2, 1, 'admin', '2018-09-23 19:37:28', 'admin', '2018-09-23 19:35:22', 0);
INSERT INTO gozero.sys_dept (id, name, parent_id, order_num, create_by, create_time, last_update_by, last_update_time, del_flag) VALUES (7, '技术部', 5, 0, 'admin', '2018-09-23 19:38:00', 'admin', '2018-09-23 19:35:22', 0);
INSERT INTO gozero.sys_dept (id, name, parent_id, order_num, create_by, create_time, last_update_by, last_update_time, del_flag) VALUES (8, '技术部', 4, 0, 'admin', '2018-09-23 19:38:10', 'admin', '2018-09-23 19:35:22', 0);
INSERT INTO gozero.sys_dept (id, name, parent_id, order_num, create_by, create_time, last_update_by, last_update_time, del_flag) VALUES (9, '技术部', 6, 0, 'admin', '2018-09-23 19:38:17', 'admin', '2018-09-23 19:35:22', 0);
INSERT INTO gozero.sys_dept (id, name, parent_id, order_num, create_by, create_time, last_update_by, last_update_time, del_flag) VALUES (10, '市场部', 5, 0, 'admin', '2018-09-23 19:38:45', 'admin', '2018-09-23 19:35:22', 0);
INSERT INTO gozero.sys_dept (id, name, parent_id, order_num, create_by, create_time, last_update_by, last_update_time, del_flag) VALUES (11, '市场部', 6, 0, 'admin', '2018-09-23 19:39:01', 'admin', '2018-09-23 19:35:22', 0);
INSERT INTO gozero.sys_dept (id, name, parent_id, order_num, create_by, create_time, last_update_by, last_update_time, del_flag) VALUES (12, '魏国', 3, 0, 'admin', '2018-09-23 19:40:42', 'admin', '2018-09-23 19:35:22', 0);
INSERT INTO gozero.sys_dept (id, name, parent_id, order_num, create_by, create_time, last_update_by, last_update_time, del_flag) VALUES (13, '蜀国', 3, 1, 'admin', '2018-09-23 19:40:54', 'admin', '2018-09-23 19:35:22', 0);
INSERT INTO gozero.sys_dept (id, name, parent_id, order_num, create_by, create_time, last_update_by, last_update_time, del_flag) VALUES (14, '吴国', 3, 2, 'admin', '2018-09-23 19:41:04', 'admin', '2018-09-23 19:35:22', 0);