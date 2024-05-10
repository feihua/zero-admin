create table sms_flash_promotion_session
(
    id          bigint auto_increment comment '编号'
        primary key,
    name        varchar(200) not null comment '场次名称',
    start_time  time         not null comment '每日开始时间',
    end_time    time         not null comment '每日结束时间',
    status tinyint not null comment '启用状态：0->不启用；1->启用',
    create_time datetime     not null comment '创建时间'
)
    comment '限时购场次表';

INSERT INTO sms_flash_promotion_session (id, name, start_time, end_time, status, create_time) VALUES (1, '8:00', '08:00:29', '09:00:33', 1, current_time);
INSERT INTO sms_flash_promotion_session (id, name, start_time, end_time, status, create_time) VALUES (2, '10:00', '10:00:33', '11:00:33', 1, current_time);
INSERT INTO sms_flash_promotion_session (id, name, start_time, end_time, status, create_time) VALUES (3, '12:00', '12:00:00', '13:00:22', 1, current_time);
INSERT INTO sms_flash_promotion_session (id, name, start_time, end_time, status, create_time) VALUES (4, '14:00', '14:00:00', '15:00:00', 1, current_time);
INSERT INTO sms_flash_promotion_session (id, name, start_time, end_time, status, create_time) VALUES (5, '16:00', '16:00:00', '17:00:00', 1, current_time);
INSERT INTO sms_flash_promotion_session (id, name, start_time, end_time, status, create_time) VALUES (6, '18:00', '18:00:00', '19:00:00', 1, current_time);
INSERT INTO sms_flash_promotion_session (id, name, start_time, end_time, status, create_time) VALUES (7, '20:00', '20:00:33', '21:00:33', 0, current_time);
