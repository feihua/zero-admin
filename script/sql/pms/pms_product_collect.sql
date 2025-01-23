create table pms_product_collect
(
    id           int auto_increment
        primary key,
    user_id      int      default 0                 not null comment '用户表的用户ID',
    value_id     int      default 0                 not null comment '如果type=0，则是商品ID；如果type=1，则是专题ID',
    collect_type tinyint  default 0                 not null comment '收藏类型，如果type=0，则是商品ID；如果type=1，则是专题ID',
    add_time     datetime default CURRENT_TIMESTAMP not null comment '创建时间',
    update_time  datetime                           null on update CURRENT_TIMESTAMP comment '更新时间',
    deleted      tinyint  default 0                 not null comment '逻辑删除'
)
    comment '收藏表';

create index goods_id
    on pms_product_collect (value_id);

create index user_id
    on pms_product_collect (user_id);

