create table sys_role_menu
(
    id          bigint auto_increment comment '编号'
        primary key,
    role_id     bigint                              not null comment '角色ID',
    menu_id     bigint                              not null comment '菜单ID',
    create_by   varchar(50)                         not null comment '创建者',
    create_time timestamp default CURRENT_TIMESTAMP not null comment '创建时间'
)
    comment '角色菜单';

