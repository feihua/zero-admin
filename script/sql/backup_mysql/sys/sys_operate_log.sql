drop table if exists sys_operate_log;
create table sys_operate_log
(
    id               bigint auto_increment comment '操作日志id'
        primary key,
    title            varchar(50)   default ''                not null comment '模块标题',
    business_type    tinyint       default 0                 not null comment '业务类型（0其它 1新增 2修改 3删除）',
    method           varchar(200)  default ''                not null comment '方法名称',
    request_method   varchar(10)   default ''                not null comment '请求方式',
    operator_type    tinyint       default 0                 not null comment '操作类别（0其它 1后台用户 2手机端用户）',
    operate_name     varchar(50)   default ''                not null comment '操作人员',
    dept_name        varchar(50)   default ''                not null comment '部门名称',
    operate_url      varchar(255)  default ''                not null comment '请求URL',
    operate_ip       varchar(128)  default ''                not null comment '主机地址',
    operate_location varchar(255)  default ''                not null comment '操作地点',
    operate_param    varchar(2000) default ''                not null comment '请求参数',
    json_result      varchar(2000) default ''                not null comment '返回参数',
    platform         varchar(50)   default ''                not null comment '平台信息',
    browser          varchar(50)   default ''                not null comment '浏览器类型',
    version          varchar(50)   default ''                not null comment '浏览器版本',
    os               varchar(50)   default ''                not null comment '操作系统',
    arch             varchar(50)   default ''                not null comment '体系结构信息',
    engine           varchar(50)   default ''                not null comment '渲染引擎信息',
    engine_details   varchar(50)   default ''                not null comment '渲染引擎详细信息',
    extra            varchar(50)   default ''                not null comment '其他信息（可选）',
    status           tinyint       default 0                 not null comment ' 操作状态(0:异常,正常) ',
    error_msg        varchar(2000) default ''                not null comment ' 错误消息 ',
    operate_time     datetime      default CURRENT_TIMESTAMP not null comment ' 操作时间 ',
    cost_time        bigint(20)    default 0                 not null comment ' 消耗时间 '

) comment = '系统操作日志表';


