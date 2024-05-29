create table sys_dept
(
    id          bigint auto_increment comment '编号'
        primary key,
    dept_name   varchar(50)                            not null comment '部门名称',
    dept_status tinyint      default 1                 not null comment '部门状态',
    dept_sort   tinyint      default 0                 not null comment '部门排序',
    parent_id   bigint                                 not null comment '上级机构ID，一级机构为0',
    leader      varchar(64)  default ''                not null comment '负责人',
    phone       varchar(11)  default ''                not null comment '电话号码',
    email       varchar(24)  default ''                not null comment '邮箱',
    remark      varchar(255) default ''                not null comment '备注信息',
    is_deleted  tinyint      default 0                 not null comment '是否删除  0：否  1：是',
    parent_ids  varchar(64)                            not null comment '上级机构IDs，一级机构为0',
    create_by   varchar(50)                            not null comment '创建者',
    create_time datetime     default CURRENT_TIMESTAMP not null comment '创建时间',
    update_by   varchar(50)  default ''                not null comment '更新者',
    update_time datetime                               null on update CURRENT_TIMESTAMP comment '更新时间'
)
    comment '部门管理';

