drop table if exists oms_order_return_reason;
create table oms_order_return_reason
(
    id          bigint auto_increment
        primary key comment '主键ID',
    name        varchar(100)                       not null comment '退货类型',
    sort        int                                not null comment '排序',
    status      tinyint                            not null comment '状态：0->不启用；1->启用',
    create_by   bigint                             not null comment '创建人ID',
    create_time datetime default CURRENT_TIMESTAMP not null comment '创建时间',
    update_by   bigint                             null comment '更新人ID',
    update_time datetime                           null on update CURRENT_TIMESTAMP comment '更新时间',
    is_deleted  tinyint  default 0                 not null comment '是否删除'
)
    comment '退货原因表';

-- 插入退货原因数据
INSERT INTO oms_order_return_reason (id, name, sort, status, create_by)
VALUES (1, '质量问题', 1, 1, 1),
       (2, '尺码太大', 1, 1, 1),
       (3, '颜色不喜欢', 1, 1, 1),
       (4, '7天无理由退货', 1, 1, 1),
       (5, '价格问题', 1, 0, 1),
       (12, '发票问题', 0, 1, 1),
       (13, '其他问题', 0, 1, 1),
       (14, '物流问题', 0, 1, 1),
       (15, '售后问题', 0, 1, 1);