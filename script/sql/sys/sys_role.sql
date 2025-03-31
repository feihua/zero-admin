drop table if exists sys_role;
create table sys_role
(
    id          bigint auto_increment comment '角色id'
        primary key,
    role_name   varchar(50)                            not null comment '名称',
    role_key    varchar(100) default ''                not null comment '角色权限字符串',
    data_scope  tinyint      default 1                 not null comment '数据范围（1：全部数据权限 2：自定数据权限 3：本部门数据权限 4：本部门及以下数据权限）',
    status      tinyint      default 1                 not null comment '状态(1:正常，0:禁用)',
    remark      varchar(255)                           not null comment '备注',
    del_flag    tinyint      default 1                 not null comment '删除标志（0代表删除 1代表存在）',
    create_by   varchar(50)  default 'admin'           not null comment '创建者',
    create_time timestamp    default CURRENT_TIMESTAMP not null comment '创建时间',
    update_by   varchar(50)  default ''                not null comment '更新者',
    update_time datetime                               null on update CURRENT_TIMESTAMP comment '更新时间',
    constraint role_name
        unique (role_name)
) comment '角色信息';

create index name_status_index
    on sys_role (role_name, status);

INSERT INTO sys_role (id, role_name, role_key, status, remark) VALUES (1, '超级管理员', 'admin',1, '全部权限');
INSERT INTO sys_role (id, role_name, role_key, status, remark) VALUES (2, '演示角色', 'query',1,  '仅有查看功能');
INSERT INTO sys_role (id, role_name, role_key, status, remark) VALUES (3, '121', 'dev',0, '121211');
