drop table if exists sys_login_log;
create table sys_login_log
(
    id             bigint auto_increment comment '登录日志id'
        primary key,
    login_name     varchar(50)  default ''                not null comment '登录账号',
    ipaddr         varchar(128) default ''                not null comment '登录IP地址',
    login_location varchar(255) default ''                not null comment '登录地点',
    platform       varchar(50)  default ''                not null comment '平台信息',
    browser        varchar(50)  default ''                not null comment '浏览器类型',
    version        varchar(50)  default ''                not null comment '浏览器版本',
    os             varchar(50)  default ''                not null comment '操作系统',
    arch           varchar(50)  default ''                not null comment '体系结构信息',
    engine         varchar(50)  default ''                not null comment '渲染引擎信息',
    engine_details varchar(50)  default ''                not null comment '渲染引擎详细信息',
    extra          varchar(50)  default ''                not null comment '其他信息（可选）',
    status         tinyint      default 0                 not null comment '登录状态(0:失败,1:成功)',
    msg            varchar(255) default ''                not null comment '提示消息',
    login_time     datetime     default CURRENT_TIMESTAMP not null comment '访问时间'
) comment = '系统登录日志表';
