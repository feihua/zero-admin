drop table if exists sms_home_advertise;
create table sms_home_advertise
(
    id          bigint auto_increment comment '编号'
        primary key,
    name        varchar(100)                           not null comment '名称',
    type        tinyint                                not null comment '轮播位置：0->PC首页轮播；1->app首页轮播',
    pic         varchar(500)                           not null comment '图片地址',
    start_time  datetime                               not null comment '开始时间',
    end_time    datetime                               not null comment '结束时间',
    status      tinyint      default 1                 not null comment '上下线状态：0->下线；1->上线',
    click_count int          default 0                 not null comment '点击数',
    order_count int          default 0                 not null comment '下单数',
    url         varchar(500)                           not null comment '链接地址',
    remark      varchar(500) default ''                not null comment '备注',
    sort        int          default 0                 not null comment '排序',
    create_time datetime     default CURRENT_TIMESTAMP not null comment '创建时间',
    update_time datetime                               null on update CURRENT_TIMESTAMP comment '更新时间',
    is_deleted  tinyint      default 0                 not null comment '是否删除'
)
    comment '首页轮播广告表';

-- 插入首页轮播广告数据
insert into sms_home_advertise (id, name, type, pic, start_time, end_time, status, click_count, order_count, url, remark, sort)
values (1, '双十一大促', 0, 'http://example.com/ad1.jpg', '2023-11-01 00:00:00', '2029-11-11 23:59:59', 1, 0, 0, 'http://example.com/promo1', '双十一活动', 100),
       (2, '圣诞节特惠', 1, 'http://example.com/ad2.jpg', '2023-12-01 00:00:00', '2029-12-25 23:59:59', 1, 0, 0, 'http://example.com/promo2', '圣诞节活动', 90),
       (3, '新年狂欢', 0, 'http://example.com/ad3.jpg', '2023-12-26 00:00:00', '2029-01-01 23:59:59', 1, 0, 0, 'http://example.com/promo3', '新年活动', 80);