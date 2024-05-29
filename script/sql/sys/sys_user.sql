create table sys_user
(
    id          bigint auto_increment comment '编号'
        primary key,
    user_name   varchar(50)                            not null comment '用户名',
    nick_name   varchar(150) default ''                not null comment '昵称',
    avatar      varchar(150) default ''                not null comment '头像',
    password    varchar(100)                           not null comment '密码',
    salt        varchar(40)                            not null comment '加密盐',
    email       varchar(100) default ''                not null comment '邮箱',
    mobile      varchar(100) default ''                not null comment '手机号',
    user_status tinyint      default 1                 not null comment '帐号状态（0正常 1停用）',
    dept_id     bigint                                 not null comment '部门id',
    remark      varchar(255) default ''                not null comment '备注信息',
    is_deleted  tinyint      default 0                 not null comment '是否删除  0：否  1：是',
    login_time  datetime                               null comment '登录时间',
    login_ip    varchar(64)  default ''                not null comment '登录ip',
    create_by   varchar(50)                            not null comment '创建者',
    create_time timestamp    default CURRENT_TIMESTAMP not null comment '创建时间',
    update_by   varchar(50)  default ''                not null comment '更新者',
    update_time datetime                               null on update CURRENT_TIMESTAMP comment '更新时间',

    constraint name
        unique (user_name)
)
    comment '用户管理';
