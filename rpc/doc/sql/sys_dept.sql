create table sys_dept
(
    id          bigint auto_increment comment '编号'
        primary key,
    name        varchar(50)                         not null comment '机构名称',
    parent_id   bigint                              not null comment '上级机构ID，一级机构为0',
    order_num   int                                 not null comment '排序',
    create_by   varchar(50)                         not null comment '创建者',
    create_time timestamp default CURRENT_TIMESTAMP null comment '创建时间',
    update_by   varchar(50)                         null comment '更新者',
    update_time datetime  default CURRENT_TIMESTAMP null on update CURRENT_TIMESTAMP comment '更新时间',
    del_flag    tinyint   default 0                 not null comment '是否删除  0：已删除  1：正常',
    parent_ids  varchar(64)                         not null comment '上级机构IDs，一级机构为0'
)
    comment '部门管理';

INSERT INTO gozero.sys_dept (id, name, parent_id, order_num, create_by, create_time, update_by, update_time, del_flag, parent_ids) VALUES (1, '轻尘集团', 0, 0, 'admin', current_time, 'admin', current_time, 1, '0');
INSERT INTO gozero.sys_dept (id, name, parent_id, order_num, create_by, create_time, update_by, update_time, del_flag, parent_ids) VALUES (2, '牧尘集团', 0, 1, 'admin', current_time, 'admin', current_time, 1, '0');
INSERT INTO gozero.sys_dept (id, name, parent_id, order_num, create_by, create_time, update_by, update_time, del_flag, parent_ids) VALUES (3, '三国集团', 0, 2, 'admin', current_time, 'admin', current_time, 1, '0');
INSERT INTO gozero.sys_dept (id, name, parent_id, order_num, create_by, create_time, update_by, update_time, del_flag, parent_ids) VALUES (4, '上海分公司', 2, 0, 'admin', current_time, 'admin', current_time, 1, '2');
INSERT INTO gozero.sys_dept (id, name, parent_id, order_num, create_by, create_time, update_by, update_time, del_flag, parent_ids) VALUES (5, '北京分公司', 1, 1, 'admin', current_time, 'admin', current_time, 1, '1');
INSERT INTO gozero.sys_dept (id, name, parent_id, order_num, create_by, create_time, update_by, update_time, del_flag, parent_ids) VALUES (6, '北京分公司', 2, 1, 'admin', current_time, 'admin', current_time, 1, '2');
INSERT INTO gozero.sys_dept (id, name, parent_id, order_num, create_by, create_time, update_by, update_time, del_flag, parent_ids) VALUES (7, '技术部', 5, 0, 'admin', current_time, 'admin', current_time, 1, '1,5');
INSERT INTO gozero.sys_dept (id, name, parent_id, order_num, create_by, create_time, update_by, update_time, del_flag, parent_ids) VALUES (8, '技术部', 4, 0, 'admin', current_time, 'admin', current_time, 1, '2,4');
INSERT INTO gozero.sys_dept (id, name, parent_id, order_num, create_by, create_time, update_by, update_time, del_flag, parent_ids) VALUES (9, '技术部', 6, 0, 'admin', current_time, 'admin', current_time, 1, '2,6');
INSERT INTO gozero.sys_dept (id, name, parent_id, order_num, create_by, create_time, update_by, update_time, del_flag, parent_ids) VALUES (10, '市场部', 5, 0, 'admin', current_time, 'admin', current_time, 1, '1,5');
INSERT INTO gozero.sys_dept (id, name, parent_id, order_num, create_by, create_time, update_by, update_time, del_flag, parent_ids) VALUES (11, '市场部', 6, 0, 'admin', current_time, 'admin', current_time, 1, '2,6');
INSERT INTO gozero.sys_dept (id, name, parent_id, order_num, create_by, create_time, update_by, update_time, del_flag, parent_ids) VALUES (12, '魏国', 3, 0, 'admin', current_time, 'admin', current_time, 1, '3');
INSERT INTO gozero.sys_dept (id, name, parent_id, order_num, create_by, create_time, update_by, update_time, del_flag, parent_ids) VALUES (13, '蜀国', 3, 1, 'admin', current_time, 'admin', current_time, 1, '3');
INSERT INTO gozero.sys_dept (id, name, parent_id, order_num, create_by, create_time, update_by, update_time, del_flag, parent_ids) VALUES (14, '吴国', 3, 2, 'admin', current_time, 'admin', current_time, 1, '3');
