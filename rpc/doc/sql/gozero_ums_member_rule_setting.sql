create table ums_member_rule_setting
(
    id                  bigint auto_increment
        primary key,
    continue_sign_day   int            null comment '连续签到天数',
    continue_sign_point int            null comment '连续签到赠送数量',
    consume_per_point   decimal(10, 2) null comment '每消费多少元获取1个点',
    low_order_amount    decimal(10, 2) null comment '最低获取点数的订单金额',
    max_point_per_order int            null comment '每笔订单最高获取点数',
    type                int(1)         null comment '类型：0->积分规则；1->成长值规则'
)
    comment '会员积分成长规则表';

