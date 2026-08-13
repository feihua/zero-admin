drop table if exists sys_dept;
create table sys_dept
(
    id          bigint auto_increment comment '部门id'
        primary key,
    parent_id   bigint(20)   default 0                 not null comment '上级部门id',
    ancestors   varchar(50)  default ''                not null comment '祖级列表',
    dept_name   varchar(30)  default ''                not null comment '部门名称',
    sort        int(4)       default 0                 not null comment '显示顺序',
    leader      varchar(20)  default ''                not null comment '负责人',
    phone       varchar(11)  default ''                not null comment '联系电话',
    email       varchar(50)  default ''                not null comment '邮箱',
    status      tinyint      default 0                 not null comment '部门状态（0：停用，1:正常）',
    del_flag    tinyint      default 1                 not null comment '删除标志（0代表删除 1代表存在）',
    remark      varchar(255) default ''                not null comment '备注信息',
    create_by   varchar(50)  default 'admin'           not null comment '创建者',
    create_time timestamp    default CURRENT_TIMESTAMP not null comment '创建时间',
    update_by   varchar(50)  default ''                not null comment '更新者',
    update_time datetime                               null on update CURRENT_TIMESTAMP comment '更新时间'
) comment = '部门表';

INSERT INTO sys_dept (id, parent_id, ancestors, dept_name, sort, leader, phone, email, status, del_flag) VALUES (1, 0, '0', '测试科技', 1, 'admin', '18613030352', '1002219331@qq.com', 1, 1);
INSERT INTO sys_dept (id, parent_id, ancestors, dept_name, sort, leader, phone, email, status, del_flag) VALUES (2, 1, '0,1', '深圳总公司', 1, '1', '1', 'xx@qq.com', 1, 1);
INSERT INTO sys_dept (id, parent_id, ancestors, dept_name, sort, leader, phone, email, status, del_flag) VALUES (3, 1, '0,1', '长沙分公司', 2, '1', '1', 'xx@qq.com', 1, 1);
INSERT INTO sys_dept (id, parent_id, ancestors, dept_name, sort, leader, phone, email, status, del_flag) VALUES (4, 2, '0,1,2', '研发部门', 1, '1', '1', 'xx@qq.com', 1, 1);
INSERT INTO sys_dept (id, parent_id, ancestors, dept_name, sort, leader, phone, email, status, del_flag) VALUES (5, 2, '0,1,2', '市场部门', 2, '1', '1', 'xx@qq.com', 1, 1);
INSERT INTO sys_dept (id, parent_id, ancestors, dept_name, sort, leader, phone, email, status, del_flag) VALUES (6, 2, '0,1,2', '测试部门', 3, '1', '1', 'xx@qq.com', 1, 1);
INSERT INTO sys_dept (id, parent_id, ancestors, dept_name, sort, leader, phone, email, status, del_flag) VALUES (7, 2, '0,1,2', '财务部门', 4, '1', '1', 'xx@qq.com', 1, 1);
INSERT INTO sys_dept (id, parent_id, ancestors, dept_name, sort, leader, phone, email, status, del_flag) VALUES (8, 2, '0,1,2', '运维部门', 5, '1', '1', 'xx@qq.com', 1, 1);
INSERT INTO sys_dept (id, parent_id, ancestors, dept_name, sort, leader, phone, email, status, del_flag) VALUES (9, 3, '0,1,3', '市场部门1', 6, '1', '1', 'xx@qq.com', 1, 1);
INSERT INTO sys_dept (id, parent_id, ancestors, dept_name, sort, leader, phone, email, status, del_flag) VALUES (10, 3, '0,1,3', '财务部门1', 1, '1', '1', 'xx@qq.com', 1, 1);
