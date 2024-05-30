create table sys_role
(
    id          bigint auto_increment comment '编号'
        primary key,
    role_name   varchar(100)                           not null comment '角色名称',
    role_key    varchar(100)                           not null comment '权限字符',
    role_status tinyint      default 1                 not null comment '角色状态',
    role_sort   tinyint      default 0                 not null comment '角色排序',
    data_scope  tinyint      default 0                 not null comment '数据权限',
    is_deleted  tinyint      default 0                 not null comment '是否删除  0：否  1：是',
    is_admin    tinyint      default 0                 not null comment '是否超级管理员:  0：否  1：是',
    remark      varchar(100) default ''                not null comment '备注',
    create_by   varchar(50)                            not null comment '创建者',
    create_time timestamp    default CURRENT_TIMESTAMP not null comment '创建时间',
    update_by   varchar(50)  default ''                not null comment '更新者',
    update_time datetime                               null on update CURRENT_TIMESTAMP comment '更新时间'

)
    comment '角色信息表';


