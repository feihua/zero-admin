drop table if exists sys_user;
create table sys_user
(
    id              bigint auto_increment comment '用户id'
        primary key,
    mobile          char(11)     default ''                not null comment '手机号码',
    user_name       varchar(50)                            not null comment '用户账号',
    nick_name       varchar(30)                            not null comment '用户昵称',
    user_type       varchar(2)   default '00'              not null comment '用户类型（00系统用户）',
    avatar          varchar(100) default ''                not null comment '头像路径',
    email           varchar(50)  default ''                not null comment '用户邮箱',
    password        varchar(64)                            not null comment '密码',
    status          tinyint      default 1                 not null comment '状态(1:正常，0:禁用)',
    dept_id         bigint       default 1                 not null comment '部门ID',
    login_ip        varchar(128) default ''                not null comment '最后登录IP',
    login_date      datetime comment '最后登录时间',
    login_browser   varchar(50)  default ''                not null comment '浏览器类型',
    login_os        varchar(50)  default ''                not null comment '操作系统',
    pwd_update_date datetime comment '密码最后更新时间',
    remark          varchar(255)                           not null comment '备注',
    del_flag        tinyint      default 1                 not null comment '删除标志（0代表删除 1代表存在）',
    create_by       varchar(50)  default 'admin'           not null comment '创建者',
    create_time     timestamp    default CURRENT_TIMESTAMP not null comment '创建时间',
    update_by       varchar(50)  default ''                not null comment '更新者',
    update_time     datetime                               null comment '更新时间',
    constraint AK_phone
        unique (mobile)
) comment '用户信息';


INSERT INTO sys_user (id, mobile, user_name, nick_name, avatar, email, password, status, remark) VALUES (1, '18613030111', 'admin','admin', 'https://gw.alipayobjects.com/zos/antfincdn/XAosXuNZyF/BiazfanxmamNRoxxVxka.png', '1002219331@qq.com','123456', 1,  '超级管理员');
INSERT INTO sys_user (id, mobile, user_name, nick_name, avatar, email, password, status, remark) VALUES (2, '18613030222', 'test', 'test', 'https://gw.alipayobjects.com/zos/antfincdn/XAosXuNZyF/BiazfanxmamNRoxxVxka.png', '123@qq.com','123456', 1, '演示权限');
