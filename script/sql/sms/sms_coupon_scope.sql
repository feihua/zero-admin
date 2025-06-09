drop table if exists sms_coupon_scope;
create table sms_coupon_scope
(
    id          bigint auto_increment  comment '主键ID'
        primary key,
    coupon_id   bigint                             not null comment '优惠券ID',
    scope_type  tinyint                            not null comment '范围类型：0-全场，1-分类，2-商品',
    scope_id    bigint                             not null comment '范围ID（分类ID或商品ID）',
    create_time datetime default CURRENT_TIMESTAMP not null comment '创建时间'

)
    comment '优惠券使用范围表';

