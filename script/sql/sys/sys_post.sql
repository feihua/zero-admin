drop table if exists sys_post;
create table sys_post
(
    id          bigint auto_increment comment '岗位id'
        primary key,
    post_code   varchar(64)                            not null comment '岗位编码',
    post_name   varchar(50)                            not null comment '岗位名称',
    sort        int          default 0                 not null comment '显示顺序',
    status      tinyint      default 0                 not null comment '岗位状态（0：停用，1:正常）',
    remark      varchar(500) default ''                not null comment '备注',
    create_by   varchar(50)  default 'admin'           not null comment '创建者',
    create_time timestamp    default CURRENT_TIMESTAMP not null comment '创建时间',
    update_by   varchar(50)  default ''                not null comment '更新者',
    update_time datetime                               null on update CURRENT_TIMESTAMP comment '更新时间'
) comment = '岗位信息表';

INSERT INTO sys_post (post_code, post_name, sort, status, remark) VALUES ('ceo', '董事长', 1, 1, '');
INSERT INTO sys_post (post_code, post_name, sort, status, remark) VALUES ('pd', '项目经理', 2, 1, '');
INSERT INTO sys_post (post_code, post_name, sort, status, remark) VALUES ('hr', '人力资源', 3, 1, '');
INSERT INTO sys_post (post_code, post_name, sort, status, remark) VALUES ('user', '普通员工', 4, 1, '');
