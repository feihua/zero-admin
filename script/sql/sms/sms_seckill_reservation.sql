drop table if exists sms_seckill_reservation;
create table sms_seckill_reservation
(
    id          bigint auto_increment comment 'ID'
        primary key,
    user_id     bigint                             not null comment '用户ID',
    activity_id bigint                             not null comment '活动ID',
    product_id  bigint                             not null comment '秒杀商品ID',
    status      tinyint  default 0                 not null comment '状态：0-已预约，1-已参与，2-已取消',
    create_time datetime default CURRENT_TIMESTAMP not null,
    update_time datetime                           null on update CURRENT_TIMESTAMP,
    is_deleted  tinyint  default 0                 not null comment '是否删除'
)
    comment '秒杀预约表';


