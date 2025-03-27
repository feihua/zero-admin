drop table if exists sms_home_brand;
create table sms_home_brand
(
    id               bigint auto_increment
        primary key,
    brand_id         bigint      not null comment '商品品牌id',
    brand_name       varchar(64) not null comment '商品品牌名称',
    recommend_status tinyint     not null comment '推荐状态：0->不推荐;1->推荐',
    sort             int         not null comment '排序'
)
    comment '首页推荐品牌';

INSERT INTO sms_home_brand (brand_id, brand_name, recommend_status, sort) VALUES (1, '万和', 1, 1);
INSERT INTO sms_home_brand (brand_id, brand_name, recommend_status, sort) VALUES (2, '三星', 1, 2);
INSERT INTO sms_home_brand (brand_id, brand_name, recommend_status, sort) VALUES (3, '华为', 1, 3);
INSERT INTO sms_home_brand (brand_id, brand_name, recommend_status, sort) VALUES (4, '格力', 1, 4);
INSERT INTO sms_home_brand (brand_id, brand_name, recommend_status, sort) VALUES (5, '方太', 1, 5);
INSERT INTO sms_home_brand (brand_id, brand_name, recommend_status, sort) VALUES (6, '小米', 1, 6);
INSERT INTO sms_home_brand (brand_id, brand_name, recommend_status, sort) VALUES (7, 'OPPO', 1, 7);
INSERT INTO sms_home_brand (brand_id, brand_name, recommend_status, sort) VALUES (8, '七匹狼', 1, 8);
INSERT INTO sms_home_brand (brand_id, brand_name, recommend_status, sort) VALUES (9, '海澜之家', 1, 9);
INSERT INTO sms_home_brand (brand_id, brand_name, recommend_status, sort) VALUES (10, '苹果', 1, 10);
INSERT INTO sms_home_brand (brand_id, brand_name, recommend_status, sort) VALUES (11, 'NIKE', 1, 11);


