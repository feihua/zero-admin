create table ums_member_rule_setting
(
    id                  bigint auto_increment
        primary key,
    continue_sign_day   int                                   not null comment '连续签到天数',
    continue_sign_point int                                   not null comment '连续签到赠送数量',
    consume_per_point   bigint                                not null comment '每消费多少元获取1个点',
    low_order_amount    bigint                                not null comment '最低获取点数的订单金额',
    max_point_per_order int                                   not null comment '每笔订单最高获取点数',
    rule_type           tinyint                               not null comment '类型：0->积分规则；1->成长值规则',
    status              tinyint     default 1                 not null comment '状态：0->禁用；1->启用',
    create_by           varchar(50)                           not null comment '创建者',
    create_time         datetime    default CURRENT_TIMESTAMP not null comment '创建时间',
    update_by           varchar(50) default ''                not null comment '更新者',
    update_time         datetime                              null on update CURRENT_TIMESTAMP comment '更新时间'
)
    comment '会员积分成长规则表';

