drop table if exists sms_coupon_type;
create table sms_coupon_type
(
    id            bigint auto_increment comment '主键ID'
        primary key,
    name          varchar(50)                            not null comment '类型名称',
    code          varchar(32)                            not null comment '类型编码',
    description   varchar(500) default ''                not null comment '描述',
    discount_type tinyint                                not null comment '优惠方式：1-固定金额，2-折扣率，3-第N件特惠，4-买赠，5-特价，6-套装优惠，7-搭配优惠，8-积分抵现，9-积分倍率，10-免运费，11-运费减免，12-限时特权，13-会员特权',
    status        tinyint      default 1                 not null comment '是否启用',
    sort          int          default 0                 not null comment '排序',
    create_by     bigint                                 not null comment '创建人ID',
    create_time   datetime     default CURRENT_TIMESTAMP not null comment '创建时间',
    update_by     bigint                                 null comment '更新人ID',
    update_time   datetime                               null on update CURRENT_TIMESTAMP comment '更新时间',
    is_deleted    tinyint      default 0                 not null comment '是否删除'
)
    comment '优惠券类型表';

insert into sms_coupon_type (id, name, code, description, discount_type, status, sort, create_by, is_deleted)
values (1, '满减券', 'MANJIAN', '满足最低消费即可减去固定金额', 1, 1, 100, 1, 0),
       (2, '无门槛券', 'WUMENJIAN', '无最低消费限制的固定金额优惠', 1, 1, 90, 1, 0),
       (3, '新人券', 'XINREN', '新用户专享固定金额优惠', 1, 1, 80, 1, 0),
       (4, '会员券', 'HUIYUAN', '会员专享固定金额优惠', 1, 1, 70, 1, 0),
       (5, '折扣券', 'ZHEKOU', '按商品金额的比例进行折扣', 2, 1, 60, 1, 0),
       (6, '会员折扣券', 'HUIYUAN_ZHEKOU', '会员专享折扣优惠', 2, 1, 50, 1, 0),
       (7, '限时折扣券', 'XIANSHI_ZHEKOU', '限时特惠折扣', 2, 1, 40, 1, 0),
       (8, '品类折扣券', 'PINLEI_ZHEKOU', '特定品类商品的折扣优惠', 2, 1, 30, 1, 0),
       (9, '已停用满减券', 'DISABLE_MANJIAN', '已停用的满减优惠券类型', 1, 0, 10, 1, 0),
       (10, '已停用折扣券', 'DISABLE_ZHEKOU', '已停用的折扣优惠券类型', 2, 0, 20, 1, 0),
       (11, '第二件半价', 'SECOND_HALF', '第二件商品半价优惠', 3, 1, 95, 1, 0),
       (12, '买二送一', 'BUY_TWO_GET_ONE', '买两件赠送一件', 4, 1, 94, 1, 0),
       (13, '第三件1元', 'THIRD_ONE_YUAN', '第三件商品1元特惠', 5, 1, 93, 1, 0),
       (14, '套装优惠', 'BUNDLE_DISCOUNT', '购买指定套装商品享受优惠', 6, 1, 85, 1, 0),
       (15, '搭配优惠', 'COMBO_DISCOUNT', '指定商品搭配购买享受优惠', 7, 1, 84, 1, 0),
       (16, '积分抵现券', 'POINTS_CASH', '使用积分抵扣现金', 8, 1, 75, 1, 0),
       (17, '积分翻倍券', 'POINTS_MULTIPLY', '购物获得双倍积分', 9, 1, 74, 1, 0),
       (18, '免运费券', 'FREE_SHIPPING', '订单免运费', 10, 1, 65, 1, 0),
       (19, '运费减免券', 'SHIPPING_DISCOUNT', '运费优惠固定金额', 11, 1, 64, 1, 0),
       (20, '限时优先券', 'TIME_PRIORITY', '限时抢购优先购买权', 12, 1, 55, 1, 0),
       (21, '会员日特权券', 'VIP_DAY', '会员日专享双重优惠', 13, 1, 54, 1, 0);