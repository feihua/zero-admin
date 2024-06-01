create table sys_post
(
    id          bigint auto_increment comment '编号'
        primary key,
    post_name   varchar(50)                            not null comment '岗位名称',
    post_code   varchar(50)                            not null comment '岗位编码',
    post_status tinyint      default 1                 not null comment '岗位状态',
    post_sort   tinyint      default 0                 not null comment '岗位排序',
    remark      varchar(255) default ''                not null comment '备注信息',
    is_deleted    tinyint      default 0                 not null comment '是否删除  0：否  1：是',
    create_by   varchar(50)                            not null comment '创建者',
    create_time timestamp    default CURRENT_TIMESTAMP not null comment '创建时间',
    update_by   varchar(50)  default ''                not null comment '更新者',
    update_time datetime                               null on update CURRENT_TIMESTAMP comment '更新时间'
)
    comment '岗位管理';

INSERT INTO sys_post (id, post_name, post_code, post_status, post_sort, remark, is_deleted, create_by, create_time, update_by, update_time) VALUES (1, '董事长', 'ceo', 1, 1, 'ceo', 0, 'admin', '2024-06-01 12:10:09', '', null);
INSERT INTO sys_post (id, post_name, post_code, post_status, post_sort, remark, is_deleted, create_by, create_time, update_by, update_time) VALUES (2, '项目经理', 'pd', 1, 2, '项目经理', 0, 'admin', '2024-06-01 12:10:38', '', null);
INSERT INTO sys_post (id, post_name, post_code, post_status, post_sort, remark, is_deleted, create_by, create_time, update_by, update_time) VALUES (3, '人力资源', 'hr', 1, 3, '', 0, 'admin', '2024-06-01 12:13:04', '', null);
INSERT INTO sys_post (id, post_name, post_code, post_status, post_sort, remark, is_deleted, create_by, create_time, update_by, update_time) VALUES (4, '普通员工', 'user', 1, 4, '', 0, 'admin', '2024-06-01 12:13:21', '', null);
