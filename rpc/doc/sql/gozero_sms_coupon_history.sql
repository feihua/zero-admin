create table sms_coupon_history
(
    id              bigint auto_increment
        primary key,
    coupon_id       bigint       null,
    member_id       bigint       null,
    coupon_code     varchar(64)  null,
    member_nickname varchar(64)  null comment '领取人昵称',
    get_type        int(1)       null comment '获取类型：0->后台赠送；1->主动获取',
    create_time     datetime     null,
    use_status      int(1)       null comment '使用状态：0->未使用；1->已使用；2->已过期',
    use_time        datetime     null comment '使用时间',
    order_id        bigint       null comment '订单编号',
    order_sn        varchar(100) null comment '订单号码'
)
    comment '优惠券使用、领取历史表';

create index idx_coupon_id
    on sms_coupon_history (coupon_id);

create index idx_member_id
    on sms_coupon_history (member_id);

INSERT INTO gozero.sms_coupon_history (id, coupon_id, member_id, coupon_code, member_nickname, get_type, create_time, use_status, use_time, order_id, order_sn) VALUES (2, 2, 1, '4931048380330002', 'windir', 1, '2018-08-29 14:04:12', 1, '2018-11-12 14:38:47', 12, '201809150101000001');
INSERT INTO gozero.sms_coupon_history (id, coupon_id, member_id, coupon_code, member_nickname, get_type, create_time, use_status, use_time, order_id, order_sn) VALUES (3, 3, 1, '4931048380330003', 'windir', 1, '2018-08-29 14:04:29', 0, null, null, null);
INSERT INTO gozero.sms_coupon_history (id, coupon_id, member_id, coupon_code, member_nickname, get_type, create_time, use_status, use_time, order_id, order_sn) VALUES (4, 4, 1, '4931048380330004', 'windir', 1, '2018-08-29 14:04:32', 0, null, null, null);
INSERT INTO gozero.sms_coupon_history (id, coupon_id, member_id, coupon_code, member_nickname, get_type, create_time, use_status, use_time, order_id, order_sn) VALUES (11, 7, 1, '4931048380330001', 'windir', 1, '2018-09-04 16:21:50', 0, null, null, null);
INSERT INTO gozero.sms_coupon_history (id, coupon_id, member_id, coupon_code, member_nickname, get_type, create_time, use_status, use_time, order_id, order_sn) VALUES (12, 2, 4, '0340981248320004', 'zhensan', 1, '2018-11-12 14:16:50', 0, null, null, null);
INSERT INTO gozero.sms_coupon_history (id, coupon_id, member_id, coupon_code, member_nickname, get_type, create_time, use_status, use_time, order_id, order_sn) VALUES (13, 3, 4, '0342977234360004', 'zhensan', 1, '2018-11-12 14:17:10', 0, null, null, null);
INSERT INTO gozero.sms_coupon_history (id, coupon_id, member_id, coupon_code, member_nickname, get_type, create_time, use_status, use_time, order_id, order_sn) VALUES (14, 4, 4, '0343342928830004', 'zhensan', 1, '2018-11-12 14:17:13', 0, null, null, null);
INSERT INTO gozero.sms_coupon_history (id, coupon_id, member_id, coupon_code, member_nickname, get_type, create_time, use_status, use_time, order_id, order_sn) VALUES (15, 2, 5, '0351883832180005', 'lisi', 1, '2018-11-12 14:18:39', 0, null, null, null);
INSERT INTO gozero.sms_coupon_history (id, coupon_id, member_id, coupon_code, member_nickname, get_type, create_time, use_status, use_time, order_id, order_sn) VALUES (16, 3, 5, '0352201672680005', 'lisi', 1, '2018-11-12 14:18:42', 0, null, null, null);
INSERT INTO gozero.sms_coupon_history (id, coupon_id, member_id, coupon_code, member_nickname, get_type, create_time, use_status, use_time, order_id, order_sn) VALUES (17, 4, 5, '0352505810180005', 'lisi', 1, '2018-11-12 14:18:45', 0, null, null, null);
INSERT INTO gozero.sms_coupon_history (id, coupon_id, member_id, coupon_code, member_nickname, get_type, create_time, use_status, use_time, order_id, order_sn) VALUES (18, 2, 6, '0356114588380006', 'wangwu', 1, '2018-11-12 14:19:21', 0, null, null, null);
INSERT INTO gozero.sms_coupon_history (id, coupon_id, member_id, coupon_code, member_nickname, get_type, create_time, use_status, use_time, order_id, order_sn) VALUES (19, 3, 6, '0356382856920006', 'wangwu', 1, '2018-11-12 14:19:24', 0, null, null, null);
INSERT INTO gozero.sms_coupon_history (id, coupon_id, member_id, coupon_code, member_nickname, get_type, create_time, use_status, use_time, order_id, order_sn) VALUES (20, 4, 6, '0356656798470006', 'wangwu', 1, '2018-11-12 14:19:27', 0, null, null, null);
INSERT INTO gozero.sms_coupon_history (id, coupon_id, member_id, coupon_code, member_nickname, get_type, create_time, use_status, use_time, order_id, order_sn) VALUES (21, 2, 3, '0363644984620003', 'windy', 1, '2018-11-12 14:20:36', 0, null, null, null);
INSERT INTO gozero.sms_coupon_history (id, coupon_id, member_id, coupon_code, member_nickname, get_type, create_time, use_status, use_time, order_id, order_sn) VALUES (22, 3, 3, '0363932820300003', 'windy', 1, '2018-11-12 14:20:39', 0, null, null, null);
INSERT INTO gozero.sms_coupon_history (id, coupon_id, member_id, coupon_code, member_nickname, get_type, create_time, use_status, use_time, order_id, order_sn) VALUES (23, 4, 3, '0364238275840003', 'windy', 1, '2018-11-12 14:20:42', 0, null, null, null);
INSERT INTO gozero.sms_coupon_history (id, coupon_id, member_id, coupon_code, member_nickname, get_type, create_time, use_status, use_time, order_id, order_sn) VALUES (24, 2, 7, '0385034833070007', 'lion', 1, '2018-11-12 14:24:10', 0, null, null, null);
INSERT INTO gozero.sms_coupon_history (id, coupon_id, member_id, coupon_code, member_nickname, get_type, create_time, use_status, use_time, order_id, order_sn) VALUES (25, 3, 7, '0385350208650007', 'lion', 1, '2018-11-12 14:24:13', 0, null, null, null);
INSERT INTO gozero.sms_coupon_history (id, coupon_id, member_id, coupon_code, member_nickname, get_type, create_time, use_status, use_time, order_id, order_sn) VALUES (26, 4, 7, '0385632733900007', 'lion', 1, '2018-11-12 14:24:16', 0, null, null, null);
INSERT INTO gozero.sms_coupon_history (id, coupon_id, member_id, coupon_code, member_nickname, get_type, create_time, use_status, use_time, order_id, order_sn) VALUES (27, 2, 8, '0388779132990008', 'shari', 1, '2018-11-12 14:24:48', 0, null, null, null);
INSERT INTO gozero.sms_coupon_history (id, coupon_id, member_id, coupon_code, member_nickname, get_type, create_time, use_status, use_time, order_id, order_sn) VALUES (28, 3, 8, '0388943658810008', 'shari', 1, '2018-11-12 14:24:49', 0, null, null, null);
INSERT INTO gozero.sms_coupon_history (id, coupon_id, member_id, coupon_code, member_nickname, get_type, create_time, use_status, use_time, order_id, order_sn) VALUES (29, 4, 8, '0389069398320008', 'shari', 1, '2018-11-12 14:24:51', 0, null, null, null);
INSERT INTO gozero.sms_coupon_history (id, coupon_id, member_id, coupon_code, member_nickname, get_type, create_time, use_status, use_time, order_id, order_sn) VALUES (30, 2, 9, '0390753935250009', 'aewen', 1, '2018-11-12 14:25:08', 0, null, null, null);
INSERT INTO gozero.sms_coupon_history (id, coupon_id, member_id, coupon_code, member_nickname, get_type, create_time, use_status, use_time, order_id, order_sn) VALUES (31, 3, 9, '0390882954470009', 'aewen', 1, '2018-11-12 14:25:09', 0, null, null, null);
INSERT INTO gozero.sms_coupon_history (id, coupon_id, member_id, coupon_code, member_nickname, get_type, create_time, use_status, use_time, order_id, order_sn) VALUES (32, 4, 9, '0391025542810009', 'aewen', 1, '2018-11-12 14:25:10', 0, null, null, null);