create table pms_feight_template
(
    id              bigint auto_increment
        primary key,
    name            varchar(64)                        not null comment '运费模版名称',
    charge_type     tinyint                            not null comment '计费类型:0->按重量；1->按件数',
    first_weight    bigint                             not null comment '首重kg',
    first_fee       bigint                             not null comment '首费（元）',
    continue_weight bigint                             not null comment '续重kg',
    continue_fee    bigint                             not null comment '续费（元）',
    dest            varchar(255)                       not null comment '目的地（省、市）',
    create_time     datetime default CURRENT_TIMESTAMP not null comment '创建时间',
    update_time     datetime                           null on update CURRENT_TIMESTAMP comment '更新时间'

)
    comment '运费模版';

