create table sys_config
(
    id               bigint auto_increment comment '编号'
        primary key,
    value            varchar(100)                        not null comment '数据值',
    label            varchar(100)                        not null comment '标签名',
    type             varchar(100)                        not null comment '类型',
    description      varchar(100)                        not null comment '描述',
    sort             decimal                             not null comment '排序（升序）',
    create_by        varchar(50)                         null comment '创建人',
    create_time      timestamp default CURRENT_TIMESTAMP null comment '创建时间',
    last_update_by   varchar(50)                         null comment '更新人',
    last_update_time datetime                            null comment '更新时间',
    remarks          varchar(255)                        null comment '备注信息',
    del_flag         tinyint   default 0                 null comment '是否删除  -1：已删除  0：正常'
)
    comment '系统配置表';

INSERT INTO gozero.sys_config (id, value, label, type, description, sort, create_by, create_time, last_update_by, last_update_time, remarks, del_flag) VALUES (1, '#14889A', 'theme', 'color', '主题色', 0, 'admin', '2018-09-23 19:52:54', 'admin', '2020-11-19 18:14:02', '主题色', 0);