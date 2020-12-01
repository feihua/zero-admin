create table sms_home_recommend_subject
(
    id               bigint auto_increment
        primary key,
    subject_id       bigint      null,
    subject_name     varchar(64) null,
    recommend_status int(1)      null,
    sort             int         null
)
    comment '首页推荐专题表';

INSERT INTO gozero.sms_home_recommend_subject (id, subject_id, subject_name, recommend_status, sort) VALUES (14, 1, 'polo衬衫的也时尚', 1, 0);
INSERT INTO gozero.sms_home_recommend_subject (id, subject_id, subject_name, recommend_status, sort) VALUES (15, 2, '大牌手机低价秒', 1, 0);
INSERT INTO gozero.sms_home_recommend_subject (id, subject_id, subject_name, recommend_status, sort) VALUES (16, 3, '晓龙845新品上市', 1, 0);
INSERT INTO gozero.sms_home_recommend_subject (id, subject_id, subject_name, recommend_status, sort) VALUES (17, 4, '夏天应该穿什么', 1, 0);
INSERT INTO gozero.sms_home_recommend_subject (id, subject_id, subject_name, recommend_status, sort) VALUES (18, 5, '夏季精选', 1, 100);
INSERT INTO gozero.sms_home_recommend_subject (id, subject_id, subject_name, recommend_status, sort) VALUES (19, 6, '品牌手机降价', 1, 0);