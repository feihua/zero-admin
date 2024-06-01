create table sys_dict_type
(
    id          bigint auto_increment comment '编号'
        primary key,
    dict_name   varchar(100)                           not null comment '字典名称',
    dict_type   varchar(100)                           not null comment '字典类型',
    dict_status tinyint      default 1                 not null comment '字典状态',
    remark      varchar(255) default ''                not null comment '备注信息',
    is_system   tinyint      default 0                 not null comment '是否系统预留  0：否  1：是',
    is_deleted  tinyint      default 0                 not null comment '是否删除  0：否  1：是',
    create_by   varchar(50)                            not null comment '创建者',
    create_time timestamp    default CURRENT_TIMESTAMP not null comment '创建时间',
    update_by   varchar(50)  default ''                not null comment '更新者',
    update_time datetime                               null on update CURRENT_TIMESTAMP comment '更新时间'

)
    comment '字典类型表';

INSERT INTO sys_dict_type (id, dict_name, dict_type, dict_status, remark, is_system, is_deleted, create_by, create_time, update_by, update_time) VALUES (1, '操作类型', 'sys_oper_type', 1, '操作类型列表', 1, 0, 'admin', '2024-06-01 13:18:56', '', null);
INSERT INTO sys_dict_type (id, dict_name, dict_type, dict_status, remark, is_system, is_deleted, create_by, create_time, update_by, update_time) VALUES (2, 'test', 'test', 1, 'test', 1, 0, 'admin', '2024-06-01 14:44:53', '', null);


create table sys_dict_item
(
    id          bigint auto_increment comment '编号'
        primary key,
    dict_type   varchar(100)                           not null comment '字典类型',
    dict_label  varchar(100)                           not null comment '字典标签',
    dict_value  varchar(100)                           not null comment '字典键值',
    dict_status tinyint      default 1                 not null comment '字典状态',
    dict_sort   tinyint      default 0                 not null comment '排序',
    remark      varchar(255) default ''                not null comment '备注信息',
    is_default  tinyint      default 0                 not null comment '是否默认  0：否  1：是',
    is_deleted  tinyint      default 1                 not null comment '是否删除  0：否  1：是',
    create_by   varchar(50)                            not null comment '创建者',
    create_time timestamp    default CURRENT_TIMESTAMP not null comment '创建时间',
    update_by   varchar(50)  default ''                not null comment '更新者',
    update_time datetime                               null on update CURRENT_TIMESTAMP comment '更新时间'

)
    comment '字典数据表';

INSERT INTO sys_dict_item (id, dict_type, dict_label, dict_value, dict_status, dict_sort, remark, is_default, is_deleted, create_by, create_time, update_by, update_time) VALUES (1, 'sys_oper_type', '新增', '1', 1, 1, '新增操作', 0, 1, 'admin', '2024-06-01 15:26:38', '', null);
INSERT INTO sys_dict_item (id, dict_type, dict_label, dict_value, dict_status, dict_sort, remark, is_default, is_deleted, create_by, create_time, update_by, update_time) VALUES (2, 'sys_oper_type', '修改', '2', 1, 2, '修改操作', 0, 1, 'admin', '2024-06-01 15:30:58', '', null);
INSERT INTO sys_dict_item (id, dict_type, dict_label, dict_value, dict_status, dict_sort, remark, is_default, is_deleted, create_by, create_time, update_by, update_time) VALUES (3, 'sys_oper_type', '删除', '3', 1, 3, '删除操作', 0, 1, 'admin', '2024-06-01 15:31:18', 'admin', '2024-06-01 15:33:59');
INSERT INTO sys_dict_item (id, dict_type, dict_label, dict_value, dict_status, dict_sort, remark, is_default, is_deleted, create_by, create_time, update_by, update_time) VALUES (4, 'sys_oper_type', '导出', '4', 1, 4, '导出操作', 0, 1, 'admin', '2024-06-01 15:32:13', 'admin', '2024-06-01 15:34:18');
INSERT INTO sys_dict_item (id, dict_type, dict_label, dict_value, dict_status, dict_sort, remark, is_default, is_deleted, create_by, create_time, update_by, update_time) VALUES (5, 'sys_oper_type', '导入', '5', 1, 5, '导入操作', 0, 1, 'admin', '2024-06-01 15:32:43', '', null);
INSERT INTO sys_dict_item (id, dict_type, dict_label, dict_value, dict_status, dict_sort, remark, is_default, is_deleted, create_by, create_time, update_by, update_time) VALUES (6, 'sys_oper_type', '清空数据', '6', 1, 6, '清空操作', 0, 1, 'admin', '2024-06-01 15:33:08', '', null);
INSERT INTO sys_dict_item (id, dict_type, dict_label, dict_value, dict_status, dict_sort, remark, is_default, is_deleted, create_by, create_time, update_by, update_time) VALUES (7, 'sys_oper_type', '其它', '7', 1, 7, '其它操作', 0, 1, 'admin', '2024-06-01 15:33:28', '', null);
