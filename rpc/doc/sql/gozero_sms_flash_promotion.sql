create table sms_flash_promotion
(
    id          bigint auto_increment
        primary key,
    title       varchar(200) null,
    start_date  date         null comment '开始日期',
    end_date    date         null comment '结束日期',
    status      int(1)       null comment '上下线状态',
    create_time datetime     null comment '秒杀时间段名称'
)
    comment '限时购表';

INSERT INTO gozero.sms_flash_promotion (id, title, start_date, end_date, status, create_time) VALUES (2, '春季家电家具疯狂秒杀1', '2018-11-12', '2018-11-23', 1, '2018-11-16 11:12:13');
INSERT INTO gozero.sms_flash_promotion (id, title, start_date, end_date, status, create_time) VALUES (3, '手机特卖', '2018-11-03', '2018-11-10', 1, '2018-11-16 11:11:31');
INSERT INTO gozero.sms_flash_promotion (id, title, start_date, end_date, status, create_time) VALUES (4, '春季家电家具疯狂秒杀3', '2018-11-24', '2018-11-25', 1, '2018-11-16 11:12:19');
INSERT INTO gozero.sms_flash_promotion (id, title, start_date, end_date, status, create_time) VALUES (5, '春季家电家具疯狂秒杀4', '2018-11-16', '2018-11-16', 1, '2018-11-16 11:12:24');
INSERT INTO gozero.sms_flash_promotion (id, title, start_date, end_date, status, create_time) VALUES (6, '春季家电家具疯狂秒杀5', '2018-11-16', '2018-11-16', 1, '2018-11-16 11:12:31');
INSERT INTO gozero.sms_flash_promotion (id, title, start_date, end_date, status, create_time) VALUES (7, '春季家电家具疯狂秒杀6', '2018-11-16', '2018-11-16', 1, '2018-11-16 11:12:35');
INSERT INTO gozero.sms_flash_promotion (id, title, start_date, end_date, status, create_time) VALUES (8, '春季家电家具疯狂秒杀7', '2018-11-16', '2018-11-16', 0, '2018-11-16 11:12:39');
INSERT INTO gozero.sms_flash_promotion (id, title, start_date, end_date, status, create_time) VALUES (9, '春季家电家具疯狂秒杀8', '2018-11-16', '2018-11-16', 0, '2018-11-16 11:12:42');
INSERT INTO gozero.sms_flash_promotion (id, title, start_date, end_date, status, create_time) VALUES (13, '测试', '2018-11-01', '2018-11-30', 0, '2018-11-19 10:34:24');