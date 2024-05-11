create table sms_flash_promotion
(
    id          bigint auto_increment comment '编号'
        primary key,
    title       varchar(200) not null comment '标题',
    start_date  date         not null comment '开始日期',
    end_date    date         not null comment '结束日期',
    status tinyint not null comment '上下线状态',
    create_time datetime     not null comment '创建时间'
)
    comment '限时购表';

INSERT INTO sms_flash_promotion (id, title, start_date, end_date, status, create_time) VALUES (2, '春季家电家具疯狂秒杀1', '2022-11-12', '2025-11-12', 1, '2023-12-05 13:50:49');
INSERT INTO sms_flash_promotion (id, title, start_date, end_date, status, create_time) VALUES (3, '手机特卖', '2022-11-12', '2025-11-12', 1, '2023-12-05 13:50:49');
INSERT INTO sms_flash_promotion (id, title, start_date, end_date, status, create_time) VALUES (4, '春季家电家具疯狂秒杀3', '2022-11-12', '2025-11-12', 1, '2023-12-05 13:50:49');
INSERT INTO sms_flash_promotion (id, title, start_date, end_date, status, create_time) VALUES (5, '春季家电家具疯狂秒杀4', '2022-11-12', '2025-11-12', 1, '2023-12-05 13:50:49');
INSERT INTO sms_flash_promotion (id, title, start_date, end_date, status, create_time) VALUES (6, '春季家电家具疯狂秒杀5', '2022-11-12', '2025-11-12', 1, '2023-12-05 13:50:49');
INSERT INTO sms_flash_promotion (id, title, start_date, end_date, status, create_time) VALUES (7, '春季家电家具疯狂秒杀6', '2022-11-12', '2025-11-12', 1, '2023-12-05 13:50:49');
INSERT INTO sms_flash_promotion (id, title, start_date, end_date, status, create_time) VALUES (8, '春季家电家具疯狂秒杀7', '2022-11-12', '2025-11-12', 0, '2023-12-05 13:50:49');
INSERT INTO sms_flash_promotion (id, title, start_date, end_date, status, create_time) VALUES (9, '春季家电家具疯狂秒杀8', '2022-11-12', '2025-11-12', 0, '2023-12-05 13:50:49');
INSERT INTO sms_flash_promotion (id, title, start_date, end_date, status, create_time) VALUES (13, '测试', '2022-11-12', '2025-11-12', 1, '2023-12-05 13:50:49');
