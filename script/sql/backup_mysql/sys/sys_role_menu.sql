create table sys_role_menu
(
    id      bigint auto_increment comment '编号'
        primary key,
    role_id bigint not null comment '角色Id',
    menu_id bigint not null comment '菜单Id'
)
    comment '角色菜单关联表';

