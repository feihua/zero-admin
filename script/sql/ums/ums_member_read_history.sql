drop table if exists ums_member_read_history;
create table ums_member_read_history
(
    id                bigint auto_increment
        primary key,
    member_id         bigint                                 not null comment '会员id',
    member_nick_name  varchar(50)                            not null comment '会员姓名',
    member_icon       varchar(200)                           not null comment '会员头像',
    product_id        bigint                                 not null comment '商品id',
    product_name      varchar(50)                            not null comment '商品名称',
    product_pic       varchar(200)                           not null comment '商品图片',
    product_sub_title varchar(128) default ''                not null comment '商品标题',
    product_price     bigint                                 not null comment '商品价格',
    create_time       timestamp    default CURRENT_TIMESTAMP not null comment '浏览时间'
)
    comment '用户商品浏览历史记录';

INSERT INTO ums_member_read_history (member_id, member_nick_name, member_icon, product_id, product_name, product_pic, product_sub_title, product_price) VALUES (1, 'test', '', 13, '华为 HUAWEI P20', 'https://macro-oss.oss-cn-shenzhen.aliyuncs.com/mall/images/20221028/iphone14_001.jpg', 'AI智慧全面屏 6GB +64GB 亮黑色 全网通版 移动联通电信4G手机 双卡双待手机 双卡双待', 3788);
INSERT INTO ums_member_read_history (member_id, member_nick_name, member_icon, product_id, product_name, product_pic, product_sub_title, product_price) VALUES (2, 'koobe', '', 13, '华为 HUAWEI P20', 'https://macro-oss.oss-cn-shenzhen.aliyuncs.com/mall/images/20221028/iphone14_001.jpg', 'AI智慧全面屏 6GB +64GB 亮黑色 全网通版 移动联通电信4G手机 双卡双待手机 双卡双待', 3788);
