create table sms_home_new_product
(
    id               bigint auto_increment
        primary key,
    product_id       bigint      not null comment '商品id',
    product_name     varchar(64) not null comment '商品名称',
    recommend_status tinyint     not null comment '推荐状态：0->不推荐;1->推荐',
    sort             int         not null comment '排序'
)
    comment '新鲜好物表';

INSERT INTO sms_home_new_product (id, product_id, product_name, recommend_status, sort) VALUES (2, 27, '小米8 全面屏游戏智能手机 6GB+64GB 黑色 全网通4G 双卡双待', 1, 200);
INSERT INTO sms_home_new_product (id, product_id, product_name, recommend_status, sort) VALUES (8, 26, '华为 HUAWEI P20 ', 0, 0);
INSERT INTO sms_home_new_product (id, product_id, product_name, recommend_status, sort) VALUES (9, 27, '小米8 全面屏游戏智能手机 6GB+64GB 黑色 全网通4G 双卡双待', 1, 0);
INSERT INTO sms_home_new_product (id, product_id, product_name, recommend_status, sort) VALUES (10, 28, '小米 红米5A 全网通版 3GB+32GB 香槟金 移动联通电信4G手机 双卡双待', 1, 0);
INSERT INTO sms_home_new_product (id, product_id, product_name, recommend_status, sort) VALUES (11, 29, 'Apple iPhone 8 Plus 64GB 红色特别版 移动联通电信4G手机', 1, 0);
INSERT INTO sms_home_new_product (id, product_id, product_name, recommend_status, sort) VALUES (12, 30, 'HLA海澜之家简约动物印花短袖T恤', 1, 0);
INSERT INTO sms_home_new_product (id, product_id, product_name, recommend_status, sort) VALUES (13, 11, '女式超柔软拉毛运动开衫1', 0, 0);
