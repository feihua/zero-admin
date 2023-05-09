create table sms_home_recommend_subject
(
    id               bigint auto_increment
        primary key,
    subject_id       bigint      not null comment '专题id',
    subject_name     varchar(64) not null comment '专题名称',
    recommend_status int(1)      not null comment '推荐状态：0->不推荐;1->推荐',
    sort             int         not null comment '排序'
)
    comment '首页推荐专题表';

INSERT INTO sms_home_recommend_subject (id, subject_id, subject_name, recommend_status, sort) VALUES (14, 1, 'polo衬衫的也时尚', 1, 0);
INSERT INTO sms_home_recommend_subject (id, subject_id, subject_name, recommend_status, sort) VALUES (15, 2, '大牌手机低价秒', 1, 0);
INSERT INTO sms_home_recommend_subject (id, subject_id, subject_name, recommend_status, sort) VALUES (16, 3, '晓龙845新品上市', 1, 0);
INSERT INTO sms_home_recommend_subject (id, subject_id, subject_name, recommend_status, sort) VALUES (17, 4, '夏天应该穿什么', 1, 0);
INSERT INTO sms_home_recommend_subject (id, subject_id, subject_name, recommend_status, sort) VALUES (18, 5, '夏季精选', 1, 99);
INSERT INTO sms_home_recommend_subject (id, subject_id, subject_name, recommend_status, sort) VALUES (21, 6, '品牌手机降价', 0, 0);
