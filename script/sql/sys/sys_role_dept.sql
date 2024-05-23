create table sys_role_dept
(
    id      bigint auto_increment comment '编号'
        primary key,
    role_id bigint not null comment '角色ID',
    dept_id bigint not null comment '机构ID'
)
    comment '角色机构';
