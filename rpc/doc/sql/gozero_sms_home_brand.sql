create table sms_home_brand
(
    id               bigint auto_increment
        primary key,
    brand_id         bigint      null,
    brand_name       varchar(64) null,
    recommend_status int(1)      null,
    sort             int         null
)
    comment '首页推荐品牌表';

INSERT INTO gozero.sms_home_brand (id, brand_id, brand_name, recommend_status, sort) VALUES (1, 1, '万和', 1, 200);
INSERT INTO gozero.sms_home_brand (id, brand_id, brand_name, recommend_status, sort) VALUES (2, 2, '三星', 1, 0);
INSERT INTO gozero.sms_home_brand (id, brand_id, brand_name, recommend_status, sort) VALUES (6, 6, '小米', 1, 300);
INSERT INTO gozero.sms_home_brand (id, brand_id, brand_name, recommend_status, sort) VALUES (8, 5, '方太', 1, 100);
INSERT INTO gozero.sms_home_brand (id, brand_id, brand_name, recommend_status, sort) VALUES (31, 49, '七匹狼', 0, 0);
INSERT INTO gozero.sms_home_brand (id, brand_id, brand_name, recommend_status, sort) VALUES (32, 50, '海澜之家', 1, 0);
INSERT INTO gozero.sms_home_brand (id, brand_id, brand_name, recommend_status, sort) VALUES (33, 51, '苹果', 1, 0);
INSERT INTO gozero.sms_home_brand (id, brand_id, brand_name, recommend_status, sort) VALUES (34, 2, '三星', 0, 0);
INSERT INTO gozero.sms_home_brand (id, brand_id, brand_name, recommend_status, sort) VALUES (35, 3, '华为', 1, 0);
INSERT INTO gozero.sms_home_brand (id, brand_id, brand_name, recommend_status, sort) VALUES (36, 4, '格力', 1, 0);
INSERT INTO gozero.sms_home_brand (id, brand_id, brand_name, recommend_status, sort) VALUES (37, 5, '方太', 1, 0);
INSERT INTO gozero.sms_home_brand (id, brand_id, brand_name, recommend_status, sort) VALUES (38, 1, '万和', 1, 0);
INSERT INTO gozero.sms_home_brand (id, brand_id, brand_name, recommend_status, sort) VALUES (39, 21, 'OPPO', 1, 0);