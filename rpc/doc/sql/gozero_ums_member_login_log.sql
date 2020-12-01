create table ums_member_login_log
(
    id          bigint auto_increment
        primary key,
    member_id   bigint      null,
    create_time datetime    null,
    ip          varchar(64) null,
    city        varchar(64) null,
    login_type  int(1)      null comment '登录类型：0->PC；1->android;2->ios;3->小程序',
    province    varchar(64) null
)
    comment '会员登录记录';

