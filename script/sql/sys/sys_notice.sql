create table sys_notice
(
    id             bigint auto_increment comment '公告ID'
        primary key,
    notice_title   varchar(50)                            not null comment '公告标题',
    notice_type    tinyint      default 1                 not null comment '公告类型（1:通知,2:公告）',
    notice_content varchar(255) default ''                not null comment '公告内容',
    status         tinyint      default 0                 not null comment '公告状态（0:关闭,1:正常 ）',
    remark         varchar(255) default ''                not null comment '备注',
    create_by      varchar(50)  default ''                not null comment '创建者',
    create_time    timestamp    default CURRENT_TIMESTAMP not null comment '创建时间',
    update_by      varchar(50)  default ''                not null comment '更新者',
    update_time    datetime                               null on update CURRENT_TIMESTAMP comment '更新时间'
) comment '通知公告表';


INSERT INTO sys_notice (notice_title, notice_type, notice_content, status) VALUES ('测试通知1', 1, '这是一条测试通知内容', 1);
INSERT INTO sys_notice (notice_title, notice_type, notice_content, status) VALUES ('测试公告2', 2, '这是一条测试公告内容', 1);


