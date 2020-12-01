create table sms_flash_promotion_session
(
    id          bigint auto_increment comment '编号'
        primary key,
    name        varchar(200) null comment '场次名称',
    start_time  time         null comment '每日开始时间',
    end_time    time         null comment '每日结束时间',
    status      int(1)       null comment '启用状态：0->不启用；1->启用',
    create_time datetime     null comment '创建时间'
)
    comment '限时购场次表';

INSERT INTO gozero.sms_flash_promotion_session (id, name, start_time, end_time, status, create_time) VALUES (1, '8:00', '08:00:29', '09:00:33', 1, '2018-11-16 13:22:17');
INSERT INTO gozero.sms_flash_promotion_session (id, name, start_time, end_time, status, create_time) VALUES (2, '10:00', '10:00:33', '11:00:33', 1, '2018-11-16 13:22:34');
INSERT INTO gozero.sms_flash_promotion_session (id, name, start_time, end_time, status, create_time) VALUES (3, '12:00', '12:00:00', '13:00:22', 1, '2018-11-16 13:22:37');
INSERT INTO gozero.sms_flash_promotion_session (id, name, start_time, end_time, status, create_time) VALUES (4, '14:00', '14:00:00', '15:00:00', 1, '2018-11-16 13:22:41');
INSERT INTO gozero.sms_flash_promotion_session (id, name, start_time, end_time, status, create_time) VALUES (5, '16:00', '16:00:00', '17:00:00', 1, '2018-11-16 13:22:45');
INSERT INTO gozero.sms_flash_promotion_session (id, name, start_time, end_time, status, create_time) VALUES (6, '18:00', '18:00:00', '19:00:00', 1, '2018-11-16 13:21:34');
INSERT INTO gozero.sms_flash_promotion_session (id, name, start_time, end_time, status, create_time) VALUES (7, '20:00', '20:00:33', '21:00:33', 0, '2018-11-16 13:22:55');