drop table if exists ums_member_rule_setting;
create table ums_member_rule_setting
(
    id                  bigint auto_increment
        primary key,
    consume_per_point   bigint                             not null comment '每消费多少元获取1个点',
    low_order_amount    bigint                             not null comment '最低获取点数的订单金额',
    max_point_per_order int                                not null comment '每笔订单最高获取点数',
    rule_type           tinyint                            not null comment '类型：0->积分规则；1->成长值规则',
    status              tinyint  default 1                 not null comment '状态：0->禁用；1->启用',
    create_by           bigint                             not null comment '创建人ID',
    create_time         datetime default CURRENT_TIMESTAMP not null comment '创建时间',
    update_by           bigint                             null comment '更新人ID',
    update_time         datetime                           null on update CURRENT_TIMESTAMP comment '更新时间'
)
    comment '会员积分成长规则表';


-- 向会员积分成长规则表添加模拟数据
INSERT INTO ums_member_rule_setting (consume_per_point, low_order_amount, max_point_per_order, rule_type, status,
                                     create_by)
VALUES
-- 积分规则
(100, 1000, 50, 0, 1, 1),  -- 每消费1元获取1积分，最低10元订单，每单最多50积分
(200, 2000, 100, 0, 1, 1), -- 每消费2元获取1积分，最低20元订单，每单最多100积分
(500, 5000, 200, 0, 0, 1), -- 每消费5元获取1积分，最低50元订单，每单最多200积分（禁用状态）

-- 成长值规则
(300, 3000, 30, 1, 1, 1),  -- 每消费3元获取1成长值，最低30元订单，每单最多30成长值
(500, 10000, 100, 1, 1, 1); -- 每消费5元获取1成长值，最低100元订单，每单最多100成长值