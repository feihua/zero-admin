create table sys_role
(
    id          bigint auto_increment comment '编号'
        primary key,
    name        varchar(100)                          not null comment '角色名称',
    remark      varchar(100) default ''                not null comment '备注',
    create_by   varchar(50)                           not null comment '创建者',
    create_time timestamp   default CURRENT_TIMESTAMP not null comment '创建时间',
    update_by   varchar(50) default ''                not null comment '更新者',
    update_time datetime                              null on update CURRENT_TIMESTAMP comment '更新时间',
    del_flag    tinyint     default 1                 not null comment '是否删除  0：已删除  1：正常',
    status      tinyint     default 1                 not null comment '状态  1:启用,0:禁用'
)
    comment '角色管理';

INSERT INTO sys_role (id, name, remark, create_by, create_time, update_by, update_time, del_flag, status) VALUES (1, 'admin', '超级管理员', 'admin', current_time, 'admin', current_time, 1, 1);
INSERT INTO sys_role (id, name, remark, create_by, create_time, update_by, update_time, del_flag, status) VALUES (2, 'mng', '项目经理', 'admin', current_time, 'admin', current_time, 1, 1);
INSERT INTO sys_role (id, name, remark, create_by, create_time, update_by, update_time, del_flag, status) VALUES (3, 'dev', '开发人员', 'admin', current_time, 'admin', current_time, 1, 1);
INSERT INTO sys_role (id, name, remark, create_by, create_time, update_by, update_time, del_flag, status) VALUES (4, 'test', '测试人员', 'admin', current_time, 'admin', current_time, 10, 1);
