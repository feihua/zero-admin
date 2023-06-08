create table sys_user_role
(
    id          bigint auto_increment comment '编号'
        primary key,
    user_id     bigint                              not null comment '用户ID',
    role_id     bigint                              not null comment '角色ID',
    create_by   varchar(50)                         not null comment '创建人',
    create_time timestamp default CURRENT_TIMESTAMP not null comment '创建时间',
) comment '用户角色';
