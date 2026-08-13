drop table if exists sms_seckill_session;
create table sms_seckill_session
(
    id          bigint auto_increment comment '秒杀场次ID'
        primary key,
    name        varchar(50)                        not null comment '场次名称',
    start_time  varchar(12)                        not null comment '开始时间',
    end_time    varchar(12)                        not null comment '结束时间',
    status      tinyint  default 0                 not null comment '状态：0-禁用，1-启用',
    sort        int      default 0                 not null comment '排序',
    create_by   bigint                             not null comment '创建人ID',
    create_time datetime default CURRENT_TIMESTAMP not null comment '创建时间',
    update_by   bigint                             null comment '更新人ID',
    update_time datetime                           null on update CURRENT_TIMESTAMP comment '更新时间',
    is_deleted  tinyint  default 0                 not null comment '是否删除'
)
    comment '秒杀场次表';


INSERT INTO sms_seckill_session (name, start_time, end_time, status, sort, create_by)
VALUES ('早场', '10:00:00', '12:00:00', 1, 100, 1),
       ('午场', '14:00:00', '16:00:00', 1, 90, 1),
       ('晚场', '19:00:00', '21:00:00', 1, 80, 1);

