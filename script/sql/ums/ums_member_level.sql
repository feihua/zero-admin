create table ums_member_level
(
    id                   bigint auto_increment
        primary key,
    level_name           varchar(100)   not null comment '等级名称',
    growth_point         int            not null comment '成长点',
    default_status       tinyint        not null comment '是否为默认等级：0->不是；1->是',
    free_freight_point   decimal(10, 2) not null comment '免运费标准',
    comment_growth_point int            not null comment '每次评价获取的成长值',
    is_free_freight      tinyint        not null comment '是否有免邮特权',
    is_sign_in           tinyint        not null comment '是否有签到特权',
    is_comment           tinyint        not null comment '是否有评论获奖励特权',
    is_promotion         tinyint        not null comment '是否有专享活动特权',
    is_member_price      tinyint        not null comment '是否有会员价格特权',
    is_birthday          tinyint        not null comment '是否有生日特权',
    remark               varchar(200)   null comment '备注'
)
    comment '会员等级表';

INSERT INTO ums_member_level (id, level_name, growth_point, default_status, free_freight_point, comment_growth_point,
                              is_free_freight, is_sign_in, is_comment, is_promotion, is_member_price, is_birthday,
                              remark)
VALUES (1, '黄金会员', 1000, 0, 199.00, 5, 1, 1, 1, 1, 1, 1, 'remarks');
INSERT INTO ums_member_level (id, level_name, growth_point, default_status, free_freight_point, comment_growth_point,
                              is_free_freight, is_sign_in, is_comment, is_promotion, is_member_price, is_birthday,
                              remark)
VALUES (2, '白金会员', 5000, 0, 99.00, 10, 1, 1, 1, 1, 1, 1, 'remarks');
INSERT INTO ums_member_level (id, level_name, growth_point, default_status, free_freight_point, comment_growth_point,
                              is_free_freight, is_sign_in, is_comment, is_promotion, is_member_price, is_birthday,
                              remark)
VALUES (3, '钻石会员', 15000, 0, 69.00, 15, 1, 1, 1, 1, 1, 1, 'remarks');
INSERT INTO ums_member_level (id, level_name, growth_point, default_status, free_freight_point, comment_growth_point,
                              is_free_freight, is_sign_in, is_comment, is_promotion, is_member_price, is_birthday,
                              remark)
VALUES (4, '普通会员', 1, 1, 199.00, 20, 1, 1, 1, 1, 0, 0, 'remarks');
