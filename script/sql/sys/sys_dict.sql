create table sys_dict
(
    id          bigint auto_increment comment '编号'
        primary key,
    dict_name   varchar(100)                           not null comment '字典名称',
    dict_type   varchar(100)                           not null comment '字典类型',
    dict_status tinyint                                not null comment '字典状态',
    remark      varchar(255) default ''                not null comment '备注信息',
    is_system   tinyint      default 0                 not null comment '是否系统预留  0：否  1：是',
    del_flag    tinyint      default 0                 not null comment '是否删除  1：已删除  0：正常',
    create_by   varchar(50)                            not null comment '创建者',
    create_time timestamp    default CURRENT_TIMESTAMP not null comment '创建时间',
    update_by   varchar(50)  default ''                not null comment '更新者',
    update_time datetime                               null on update CURRENT_TIMESTAMP comment '更新时间'

)
    comment '字典表';

create table sys_dict_item
(
    id          bigint auto_increment comment '编号'
        primary key,
    dict_type   varchar(100)                           not null comment '字典类型',
    dict_label  varchar(100)                           not null comment '字典标签',
    dict_value  varchar(100)                           not null comment '字典键值',
    dict_status tinyint                                not null comment '字典状态',
    dict_sort   tinyint      default 0                 not null comment '排序',
    remark      varchar(255) default ''                not null comment '备注信息',
    is_default  tinyint      default 0                 not null comment '是否默认  0：否  1：是',
    del_flag    tinyint      default 0                 not null comment '是否删除  1：已删除  0：正常',
    create_by   varchar(50)                            not null comment '创建者',
    create_time timestamp    default CURRENT_TIMESTAMP not null comment '创建时间',
    update_by   varchar(50)  default ''                not null comment '更新者',
    update_time datetime                               null on update CURRENT_TIMESTAMP comment '更新时间'

)
    comment '字典项表';

