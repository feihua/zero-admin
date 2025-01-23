create table sms_home_advertise
(
    id          bigint auto_increment
        primary key,
    name        varchar(100)            not null comment '名称',
    type        tinyint                 not null comment '轮播位置：0->PC首页轮播；1->app首页轮播',
    pic         varchar(500)            not null comment '图片地址',
    start_time  datetime                not null comment '开始时间',
    end_time    datetime                not null comment '结束时间',
    status      tinyint                 not null comment '上下线状态：0->下线；1->上线',
    click_count int                     not null comment '点击数',
    order_count int                     not null comment '下单数',
    url         varchar(500)            not null comment '链接地址',
    note        varchar(500) default '' not null comment '备注',
    sort        int          default 0  not null comment '排序'
)
    comment '首页轮播广告表';

INSERT INTO sms_home_advertise (id, name, type, pic, start_time, end_time, status, click_count, order_count, url, note, sort) VALUES (2, '小米推荐广告', 1, 'http://macro-oss.oss-cn-shenzhen.aliyuncs.com/mall/images/20221108/xiaomi_banner_01.png', current_time, current_time, 1, 0, 0, 'www.baidu.com', '夏季大热促销', 0);
INSERT INTO sms_home_advertise (id, name, type, pic, start_time, end_time, status, click_count, order_count, url, note, sort) VALUES (3, '华为推荐广告', 1, 'http://macro-oss.oss-cn-shenzhen.aliyuncs.com/mall/images/20221108/huawei_banner_01.png', current_time, current_time, 1, 0, 0, 'www.baidu.com', '夏季大热促销1', 0);
INSERT INTO sms_home_advertise (id, name, type, pic, start_time, end_time, status, click_count, order_count, url, note, sort) VALUES (4, '苹果推荐广告', 1, 'http://macro-oss.oss-cn-shenzhen.aliyuncs.com/mall/images/20221108/apple_banner_01.png', current_time, current_time, 1, 0, 0, 'www.baidu.com', '夏季大热促销2', 0);
INSERT INTO sms_home_advertise (id, name, type, pic, start_time, end_time, status, click_count, order_count, url, note, sort) VALUES (9, '三星推荐广告', 1, 'http://macro-oss.oss-cn-shenzhen.aliyuncs.com/mall/images/20221108/sanxing_banner_01.png', current_time, current_time, 1, 0, 0, 'www.baidu.com', '电影推荐广告', 100);
INSERT INTO sms_home_advertise (id, name, type, pic, start_time, end_time, status, click_count, order_count, url, note, sort) VALUES (10, 'OPPO推荐广告', 1, 'http://macro-oss.oss-cn-shenzhen.aliyuncs.com/mall/images/20221108/oppo_banner_01.png', current_time, current_time, 1, 0, 0, 'www.baidu.com', '电影推荐广告', 99);
