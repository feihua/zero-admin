create table sys_role_dept
(
    id          bigint auto_increment comment '编号'
        primary key,
    role_id     bigint                              not null comment '角色ID',
    dept_id     bigint                              not null comment '机构ID',
    create_by   varchar(50)                         not null comment '创建人',
    create_time timestamp default CURRENT_TIMESTAMP not null comment '创建时间'
)
    comment '角色机构';
