drop table if exists sms_seckill_product;
create table sms_seckill_product
(
    id            bigint auto_increment comment 'ID'
        primary key,
    activity_id   bigint                             not null comment '活动ID',
    session_id    bigint                             not null comment '秒杀场次ID',
    sku_id        bigint                             not null comment '商品SKU ID',
    sku_name      varchar(255)                       not null comment '商品名称',
    seckill_price decimal(10, 2)                     not null comment '秒杀价格',
    seckill_stock int                                not null comment '秒杀库存',
    stock_locked  int      default 0                 not null comment '锁定库存',
    per_limit     int      default 1                 not null comment '每人限购数量',
    sort          int      default 0                 not null comment '排序',
    status        tinyint  default 1                 not null comment '状态：0-未上架，1-已上架',
    create_by     bigint                             not null comment '创建人ID',
    create_time   datetime default CURRENT_TIMESTAMP not null comment '创建时间',
    update_by     bigint                             null comment '更新人ID',
    update_time   datetime                           null on update CURRENT_TIMESTAMP comment '更新时间',
    is_deleted    tinyint  default 0                 not null comment '是否删除'
)
    comment '秒杀商品表';

-- 插入秒杀商品数据
insert into sms_seckill_product (id, activity_id, session_id, sku_id, sku_name, seckill_price, seckill_stock, stock_locked, per_limit, sort, status, create_by, is_deleted)
values (1, 1, 1, 101, '华为手机',99.99, 100, 0, 1, 100, 1, 1, 0),
       (2, 1, 2, 102, '苹果手机',49.99, 200, 0, 2, 90, 1, 2, 0),
       (3, 2, 3, 103, '小米手机',29.99, 150, 0, 1, 80, 1, 3, 0);