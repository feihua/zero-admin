create table pms_feight_template
(
    id              bigint auto_increment
        primary key,
    name            varchar(64)  not null comment '运费模版名称',
    charge_type     tinyint      not null comment '计费类型:0->按重量；1->按件数',
    first_weight    bigint       not null comment '首重kg',
    first_fee       bigint       not null comment '首费（元）',
    continue_weight bigint       not null,
    continue_fee    bigint       not null,
    dest            varchar(255) not null comment '目的地（省、市）'
)
    comment '运费模版';

