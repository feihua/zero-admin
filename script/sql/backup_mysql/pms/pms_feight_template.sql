drop table if exists pms_feight_template;
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

INSERT INTO pms_feight_template (name, charge_type, first_weight, first_fee, continue_weight, continue_fee, dest)
VALUES
    ('北京标准运费模板', 0, 1000, 2000, 500, 500, '北京'),
    ('上海快递专用模板', 1, 1, 1500, 1, 300, '上海'),
    ('广东经济型运费模板', 0, 2000, 3000, 1000, 800, '广东'),
    ('浙江自提免运费模板', 1, 1, 1000, 1, 200, '浙江');
