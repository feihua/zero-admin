drop table if exists pms_product_attribute_value;

create table pms_product_attribute_value
(
    id           bigint auto_increment comment '主键id'
        primary key,
    spu_id       bigint                             not null comment '商品SPU ID',
    attribute_id bigint                             not null comment '属性ID',
    value        varchar(500)                       not null comment '属性值',
    create_by    bigint                             not null comment '创建人ID',
    create_time  datetime default CURRENT_TIMESTAMP not null comment '创建时间',
    update_by    bigint                             null comment '更新人ID',
    update_time  datetime                           null on update CURRENT_TIMESTAMP comment '更新时间',
    is_deleted   tinyint  default 0                 not null comment '是否删除'
)
    comment '商品属性值表';

create index idx_attribute
    on pms_product_attribute_value (attribute_id, is_deleted);

create index idx_spu
    on pms_product_attribute_value (spu_id, is_deleted);


-- 插入商品属性值数据（假设iPhone 15 Pro的spu_id=1）
INSERT INTO pms_product_attribute_value (spu_id, attribute_id, value, create_by) VALUES
-- 基本信息属性值
(1, 1, 'Apple', 1),                    -- 品牌
(1, 2, 'iPhone 15 Pro', 1),           -- 型号
(1, 3, '2023', 1),                    -- 上市年份
(1, 4, '自然色', 1),                   -- 机身颜色
(1, 5, '8GB', 1),                     -- 运行内存
(1, 6, '256GB', 1),                   -- 机身存储

-- 主体参数属性值
(1, 7, 'A17 Pro', 1),                 -- 处理器
(1, 8, '8GB', 1),                     -- 运行内存
(1, 9, '256GB', 1),                   -- 机身存储
(1, 10, '3274mAh', 1),                -- 电池容量
(1, 11, '187g', 1),                   -- 机身重量

-- 网络参数属性值
(1, 12, '支持', 1),                    -- 5G网络
(1, 13, 'GSM,WCDMA,LTE,5G', 1),       -- 网络制式
(1, 14, '双卡双待', 1),                -- SIM卡类型
(1, 15, 'WiFi 6E', 1),                -- WiFi

-- 显示参数属性值
(1, 16, '6.1', 1),                    -- 屏幕尺寸
(1, 17, '2556x1179', 1),              -- 屏幕分辨率
(1, 18, '120', 1),                    -- 屏幕刷新率
(1, 19, 'OLED', 1),                   -- 屏幕类型
(1, 20, '2000', 1),                   -- 屏幕亮度

-- 摄像功能属性值
(1, 21, '4800', 1),                   -- 主摄像头
(1, 22, '1200', 1),                   -- 前置摄像头
(1, 23, '3', 1),                      -- 摄像头数量
(1, 24, '4K,1080P', 1),               -- 视频拍摄
(1, 25, '光学防抖', 1);                -- 防抖功能