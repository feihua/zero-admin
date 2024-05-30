create table sys_user_role
(
    id      bigint auto_increment comment '编号'
        primary key,
    user_id bigint not null comment '用户Id',
    role_id bigint not null comment '角色Id'
) comment '用户角色关联表';

INSERT INTO sys_user_role (id, user_id, role_id) VALUES (1, 1, 1);
