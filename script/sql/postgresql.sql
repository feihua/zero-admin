-- 帮助表建语句
drop table if exists cms_help;
create table cms_help
(
    id          bigserial primary KEY,
    category_id bigint                                not null,
    icon        varchar                               not null,
    title       varchar                               not null,
    show_status smallint    default 1                 not null,
    read_count  integer                               not null,
    content     text                                  not null,
    create_by   varchar     default ''                not null,
    create_time timestamptz default current_timestamp not null,
    update_by   varchar                               not null,
    update_time timestamptz                           not null
);

-- 添加帮助表注释
comment on table cms_help is '帮助表';

-- 添加帮助表列注释
comment on column cms_help.id is '主键ID';
comment on column cms_help.category_id is '分类ID';
comment on column cms_help.icon is '图标';
comment on column cms_help.title is '标题';
comment on column cms_help.show_status is '显示状态：0->不显示；1->显示';
comment on column cms_help.read_count is '阅读量';
comment on column cms_help.content is '内容';
comment on column cms_help.create_by is '创建者';
comment on column cms_help.create_time is '创建时间';
comment on column cms_help.update_by is '更新者';
comment on column cms_help.update_time is '更新时间';

-- 帮助分类表建语句
drop table if exists cms_help_category;
create table cms_help_category
(
    id          bigserial primary KEY,
    name        varchar                               not null,
    icon        varchar                               not null,
    help_count  integer                               not null,
    show_status smallint    default 1                 not null,
    sort        integer                               not null,
    create_by   varchar     default ''                not null,
    create_time timestamptz default current_timestamp not null,
    update_by   varchar                               not null,
    update_time timestamptz                           not null
);

-- 添加帮助分类表注释
comment on table cms_help_category is '帮助分类表';

-- 添加帮助分类表列注释
comment on column cms_help_category.id is '主键ID';
comment on column cms_help_category.name is '分类名称';
comment on column cms_help_category.icon is '分类图标';
comment on column cms_help_category.help_count is '专题数量';
comment on column cms_help_category.show_status is '显示状态：0->不显示；1->显示';
comment on column cms_help_category.sort is '排序';
comment on column cms_help_category.create_by is '创建者';
comment on column cms_help_category.create_time is '创建时间';
comment on column cms_help_category.update_by is '更新者';
comment on column cms_help_category.update_time is '更新时间';

-- 用户举报表建语句
drop table if exists cms_member_report;
create table cms_member_report
(
    id                 bigserial primary KEY,
    report_type        smallint    default 1                 not null,
    report_member_name varchar                               not null,
    report_object      varchar                               not null,
    report_status      smallint    default 1                 not null,
    handle_status      smallint    default 1                 not null,
    note               varchar                               not null,
    create_time        timestamptz default current_timestamp not null
);

-- 添加用户举报表注释
comment on table cms_member_report is '用户举报表';

-- 添加用户举报表列注释
comment on column cms_member_report.id is '编号';
comment on column cms_member_report.report_type is '举报类型：0->商品评价；1->话题内容；2->用户评论';
comment on column cms_member_report.report_member_name is '举报人';
comment on column cms_member_report.report_object is '被举报对象';
comment on column cms_member_report.report_status is '举报状态：0->未处理；1->已处理';
comment on column cms_member_report.handle_status is '处理结果：0->无效；1->有效；2->恶意';
comment on column cms_member_report.note is '备注';
comment on column cms_member_report.create_time is '创建时间';

-- 优选专区建语句
drop table if exists cms_preferred_area;
create table cms_preferred_area
(
    id          bigserial primary KEY,
    name        varchar                               not null,
    sub_title   varchar                               not null,
    pic         varchar                               not null,
    sort        integer                               not null,
    show_status smallint    default 1                 not null,
    create_by   varchar     default ''                not null,
    create_time timestamptz default current_timestamp not null,
    update_by   varchar                               not null,
    update_time timestamptz                           not null
);

-- 添加优选专区注释
comment on table cms_preferred_area is '优选专区';

-- 添加优选专区列注释
comment on column cms_preferred_area.id is '主键ID';
comment on column cms_preferred_area.name is '专区名称';
comment on column cms_preferred_area.sub_title is '子标题';
comment on column cms_preferred_area.pic is '展示图片';
comment on column cms_preferred_area.sort is '排序';
comment on column cms_preferred_area.show_status is '显示状态：0->不显示；1->显示';
comment on column cms_preferred_area.create_by is '创建者';
comment on column cms_preferred_area.create_time is '创建时间';
comment on column cms_preferred_area.update_by is '更新者';
comment on column cms_preferred_area.update_time is '更新时间';

-- 优选专区和产品关系表建语句
drop table if exists cms_preferred_area_product_relation;
create table cms_preferred_area_product_relation
(
    id                bigserial primary KEY,
    preferred_area_id bigint not null,
    product_id        bigint not null
);

-- 添加优选专区和产品关系表注释
comment on table cms_preferred_area_product_relation is '优选专区和产品关系表';

-- 添加优选专区和产品关系表列注释
comment on column cms_preferred_area_product_relation.id is '主键ID';
comment on column cms_preferred_area_product_relation.preferred_area_id is '优选专区ID';
comment on column cms_preferred_area_product_relation.product_id is '产品ID';

-- 专题表建语句
drop table if exists cms_subject;
create table cms_subject
(
    id               bigserial primary KEY,
    category_id      bigint                                not null,
    title            varchar                               not null,
    pic              varchar                               not null,
    product_count    integer                               not null,
    recommend_status smallint    default 1                 not null,
    collect_count    integer                               not null,
    read_count       integer                               not null,
    comment_count    integer                               not null,
    album_pics       varchar                               not null,
    description      varchar                               not null,
    show_status      smallint    default 1                 not null,
    content          text                                  not null,
    forward_count    integer                               not null,
    category_name    varchar                               not null,
    sort             integer                               not null,
    create_by        varchar     default ''                not null,
    create_time      timestamptz default current_timestamp not null,
    update_by        varchar                               not null,
    update_time      timestamptz                           not null
);

-- 添加专题表注释
comment on table cms_subject is '专题表';

-- 添加专题表列注释
comment on column cms_subject.id is '专题id';
comment on column cms_subject.category_id is '专题分类id';
comment on column cms_subject.title is '专题标题';
comment on column cms_subject.pic is '专题主图';
comment on column cms_subject.product_count is '关联产品数量';
comment on column cms_subject.recommend_status is '推荐状态：0->不推荐；1->推荐';
comment on column cms_subject.collect_count is '收藏数';
comment on column cms_subject.read_count is '阅读数';
comment on column cms_subject.comment_count is '评论数';
comment on column cms_subject.album_pics is '画册图片用逗号分割';
comment on column cms_subject.description is '专题内容';
comment on column cms_subject.show_status is '显示状态：0->不显示；1->显示';
comment on column cms_subject.content is '专题内容';
comment on column cms_subject.forward_count is '转发数';
comment on column cms_subject.category_name is '专题分类名称';
comment on column cms_subject.sort is '排序';
comment on column cms_subject.create_by is '创建者';
comment on column cms_subject.create_time is '创建时间';
comment on column cms_subject.update_by is '更新者';
comment on column cms_subject.update_time is '更新时间';

-- 专题分类表建语句
drop table if exists cms_subject_category;
create table cms_subject_category
(
    id            bigserial primary KEY,
    name          varchar                               not null,
    icon          varchar                               not null,
    subject_count integer                               not null,
    show_status   smallint    default 1                 not null,
    sort          integer                               not null,
    create_by     varchar     default ''                not null,
    create_time   timestamptz default current_timestamp not null,
    update_by     varchar                               not null,
    update_time   timestamptz                           not null
);

-- 添加专题分类表注释
comment on table cms_subject_category is '专题分类表';

-- 添加专题分类表列注释
comment on column cms_subject_category.id is '主键ID';
comment on column cms_subject_category.name is '专题分类名称';
comment on column cms_subject_category.icon is '分类图标';
comment on column cms_subject_category.subject_count is '专题数量';
comment on column cms_subject_category.show_status is '显示状态：0->不显示；1->显示';
comment on column cms_subject_category.sort is '排序';
comment on column cms_subject_category.create_by is '创建者';
comment on column cms_subject_category.create_time is '创建时间';
comment on column cms_subject_category.update_by is '更新者';
comment on column cms_subject_category.update_time is '更新时间';

-- 专题评论表建语句
drop table if exists cms_subject_comment;
create table cms_subject_comment
(
    id               bigserial primary KEY,
    subject_id       bigint                                not null,
    member_nick_name varchar                               not null,
    member_icon      varchar                               not null,
    content          varchar                               not null,
    create_time      timestamptz default current_timestamp not null,
    show_status      smallint    default 1                 not null
);

-- 添加专题评论表注释
comment on table cms_subject_comment is '专题评论表';

-- 添加专题评论表列注释
comment on column cms_subject_comment.id is '编号';
comment on column cms_subject_comment.subject_id is '关联专题id';
comment on column cms_subject_comment.member_nick_name is '关联会员昵称';
comment on column cms_subject_comment.member_icon is '会员头像';
comment on column cms_subject_comment.content is '评论内容';
comment on column cms_subject_comment.create_time is '创建时间';
comment on column cms_subject_comment.show_status is '是否显示，0->不显示；1->显示';

-- 专题商品关系表建语句
drop table if exists cms_subject_product_relation;
create table cms_subject_product_relation
(
    id         bigserial primary KEY,
    subject_id bigint not null,
    product_id bigint not null
);

-- 添加专题商品关系表注释
comment on table cms_subject_product_relation is '专题商品关系表';

-- 添加专题商品关系表列注释
comment on column cms_subject_product_relation.id is '主键ID';
comment on column cms_subject_product_relation.subject_id is '专题ID';
comment on column cms_subject_product_relation.product_id is '商品ID';

-- 话题表建语句
drop table if exists cms_topic;
create table cms_topic
(
    id              bigserial primary KEY,
    category_id     bigint                                not null,
    name            varchar                               not null,
    start_time      timestamptz                           not null,
    end_time        timestamptz                           not null,
    attend_count    integer                               not null,
    attention_count integer                               not null,
    read_count      integer                               not null,
    award_name      varchar                               not null,
    attend_type     varchar                               not null,
    content         text                                  not null,
    create_by       varchar     default ''                not null,
    create_time     timestamptz default current_timestamp not null,
    update_by       varchar                               not null,
    update_time     timestamptz                           not null
);

-- 添加话题表注释
comment on table cms_topic is '话题表';

-- 添加话题表列注释
comment on column cms_topic.id is '主键ID';
comment on column cms_topic.category_id is '关联分类id';
comment on column cms_topic.name is '话题名称';
comment on column cms_topic.start_time is '话题开始时间';
comment on column cms_topic.end_time is '话题结束时间';
comment on column cms_topic.attend_count is '参与人数';
comment on column cms_topic.attention_count is '关注人数';
comment on column cms_topic.read_count is '阅读数';
comment on column cms_topic.award_name is '奖品名称';
comment on column cms_topic.attend_type is '参与方式';
comment on column cms_topic.content is '话题内容';
comment on column cms_topic.create_by is '创建者';
comment on column cms_topic.create_time is '创建时间';
comment on column cms_topic.update_by is '更新者';
comment on column cms_topic.update_time is '更新时间';

-- 话题分类表建语句
drop table if exists cms_topic_category;
create table cms_topic_category
(
    id            bigserial primary KEY,
    name          varchar                               not null,
    icon          varchar                               not null,
    subject_count integer                               not null,
    show_status   smallint    default 1                 not null,
    sort          integer                               not null,
    create_by     varchar     default ''                not null,
    create_time   timestamptz default current_timestamp not null,
    update_by     varchar                               not null,
    update_time   timestamptz                           not null
);

-- 添加话题分类表注释
comment on table cms_topic_category is '话题分类表';

-- 添加话题分类表列注释
comment on column cms_topic_category.id is '主键ID';
comment on column cms_topic_category.name is '分类名称';
comment on column cms_topic_category.icon is '分类图标';
comment on column cms_topic_category.subject_count is '专题数量';
comment on column cms_topic_category.show_status is '显示状态：0->不显示；1->显示';
comment on column cms_topic_category.sort is '排序';
comment on column cms_topic_category.create_by is '创建者';
comment on column cms_topic_category.create_time is '创建时间';
comment on column cms_topic_category.update_by is '更新者';
comment on column cms_topic_category.update_time is '更新时间';

-- 专题评论表建语句
drop table if exists cms_topic_comment;
create table cms_topic_comment
(
    id               bigserial primary KEY,
    member_nick_name varchar                               not null,
    topic_id         bigint                                not null,
    member_icon      varchar                               not null,
    content          varchar                               not null,
    create_time      timestamptz default current_timestamp not null,
    show_status      smallint    default 1                 not null
);

-- 添加专题评论表注释
comment on table cms_topic_comment is '专题评论表';

-- 添加专题评论表列注释
comment on column cms_topic_comment.id is '主键ID';
comment on column cms_topic_comment.member_nick_name is '评论人员昵称';
comment on column cms_topic_comment.topic_id is '专题ID';
comment on column cms_topic_comment.member_icon is '评论人员头像';
comment on column cms_topic_comment.content is '评论内容';
comment on column cms_topic_comment.create_time is '评论时间';
comment on column cms_topic_comment.show_status is '是否显示，0->不显示；1->显示';

-- 运费模版建语句
drop table if exists pms_feight_template;
create table pms_feight_template
(
    id              bigserial primary KEY,
    name            varchar                               not null,
    charge_type     smallint    default 1                 not null,
    first_weight    bigint                                not null,
    first_fee       bigint                                not null,
    continue_weight bigint                                not null,
    continue_fee    bigint                                not null,
    dest            varchar                               not null,
    create_time     timestamptz default current_timestamp not null,
    update_time     timestamptz                           not null
);

-- 添加运费模版注释
comment on table pms_feight_template is '运费模版';

-- 添加运费模版列注释
comment on column pms_feight_template.id is '';
comment on column pms_feight_template.name is '运费模版名称';
comment on column pms_feight_template.charge_type is '计费类型:0->按重量；1->按件数';
comment on column pms_feight_template.first_weight is '首重kg';
comment on column pms_feight_template.first_fee is '首费（元）';
comment on column pms_feight_template.continue_weight is '续重kg';
comment on column pms_feight_template.continue_fee is '续费（元）';
comment on column pms_feight_template.dest is '目的地（省、市）';
comment on column pms_feight_template.create_time is '创建时间';
comment on column pms_feight_template.update_time is '更新时间';

-- 商品会员价格表建语句
drop table if exists pms_member_price;
create table pms_member_price
(
    id                bigserial primary KEY,
    product_id        bigint  not null,
    member_level_id   bigint  not null,
    member_price      bigint  not null,
    member_level_name varchar not null
);

-- 添加商品会员价格表注释
comment on table pms_member_price is '商品会员价格表';

-- 添加商品会员价格表列注释
comment on column pms_member_price.id is '';
comment on column pms_member_price.product_id is '商品id';
comment on column pms_member_price.member_level_id is '会员等级id';
comment on column pms_member_price.member_price is '会员价格';
comment on column pms_member_price.member_level_name is '会员等级名称';

-- 商品属性表建语句
drop table if exists pms_product_attribute;
create table pms_product_attribute
(
    id            bigserial primary KEY,
    group_id      bigint                                not null,
    name          varchar                               not null,
    input_type    smallint    default 1                 not null,
    value_type    smallint    default 1                 not null,
    input_list    varchar                               not null,
    unit          varchar                               not null,
    is_required   smallint    default 1                 not null,
    is_searchable smallint    default 1                 not null,
    is_show       smallint    default 1                 not null,
    sort          integer                               not null,
    status        smallint    default 1                 not null,
    create_by     varchar     default ''                not null,
    create_time   timestamptz default current_timestamp not null,
    update_by     varchar                               not null,
    update_time   timestamptz                           not null,
    is_deleted    smallint    default 1                 not null
);

-- 添加商品属性表注释
comment on table pms_product_attribute is '商品属性表';

-- 添加商品属性表列注释
comment on column pms_product_attribute.id is '主键id';
comment on column pms_product_attribute.group_id is '属性分组ID';
comment on column pms_product_attribute.name is '属性名称';
comment on column pms_product_attribute.input_type is '输入类型：1-手动输入，2-单选，3-多选';
comment on column pms_product_attribute.value_type is '值类型：1-文本，2-数字，3-日期';
comment on column pms_product_attribute.input_list is '可选值列表，用逗号分隔';
comment on column pms_product_attribute.unit is '单位';
comment on column pms_product_attribute.is_required is '是否必填';
comment on column pms_product_attribute.is_searchable is '是否支持搜索';
comment on column pms_product_attribute.is_show is '是否显示';
comment on column pms_product_attribute.sort is '排序';
comment on column pms_product_attribute.status is '状态：0->禁用；1->启用';
comment on column pms_product_attribute.create_by is '创建人ID';
comment on column pms_product_attribute.create_time is '创建时间';
comment on column pms_product_attribute.update_by is '更新人ID';
comment on column pms_product_attribute.update_time is '更新时间';
comment on column pms_product_attribute.is_deleted is '是否删除';

-- 商品属性分组表建语句
drop table if exists pms_product_attribute_group;
create table pms_product_attribute_group
(
    id          bigserial primary KEY,
    category_id bigint                                not null,
    name        varchar                               not null,
    sort        integer                               not null,
    status      smallint    default 1                 not null,
    create_by   varchar     default ''                not null,
    create_time timestamptz default current_timestamp not null,
    update_by   varchar                               not null,
    update_time timestamptz                           not null,
    is_deleted  smallint    default 1                 not null
);

-- 添加商品属性分组表注释
comment on table pms_product_attribute_group is '商品属性分组表';

-- 添加商品属性分组表列注释
comment on column pms_product_attribute_group.id is '主键id';
comment on column pms_product_attribute_group.category_id is '分类ID';
comment on column pms_product_attribute_group.name is '分组名称';
comment on column pms_product_attribute_group.sort is '排序';
comment on column pms_product_attribute_group.status is '状态：0->禁用；1->启用';
comment on column pms_product_attribute_group.create_by is '创建人ID';
comment on column pms_product_attribute_group.create_time is '创建时间';
comment on column pms_product_attribute_group.update_by is '更新人ID';
comment on column pms_product_attribute_group.update_time is '更新时间';
comment on column pms_product_attribute_group.is_deleted is '是否删除';

-- 商品属性值表建语句
drop table if exists pms_product_attribute_value;
create table pms_product_attribute_value
(
    id           bigserial primary KEY,
    spu_id       bigint                                not null,
    attribute_id bigint                                not null,
    value        varchar                               not null,
    status       smallint    default 1                 not null,
    create_by    varchar     default ''                not null,
    create_time  timestamptz default current_timestamp not null,
    update_by    varchar                               not null,
    update_time  timestamptz                           not null,
    is_deleted   smallint    default 1                 not null
);

-- 添加商品属性值表注释
comment on table pms_product_attribute_value is '商品属性值表';

-- 添加商品属性值表列注释
comment on column pms_product_attribute_value.id is '主键id';
comment on column pms_product_attribute_value.spu_id is '商品SPU ID';
comment on column pms_product_attribute_value.attribute_id is '属性ID';
comment on column pms_product_attribute_value.value is '属性值';
comment on column pms_product_attribute_value.status is '状态：0->禁用；1->启用';
comment on column pms_product_attribute_value.create_by is '创建人ID';
comment on column pms_product_attribute_value.create_time is '创建时间';
comment on column pms_product_attribute_value.update_by is '更新人ID';
comment on column pms_product_attribute_value.update_time is '更新时间';
comment on column pms_product_attribute_value.is_deleted is '是否删除';

-- 商品品牌建语句
drop table if exists pms_product_brand;
create table pms_product_brand
(
    id                    bigserial primary KEY,
    name                  varchar                               not null,
    logo                  varchar                               not null,
    big_pic               varchar                               not null,
    description           text                                  not null,
    first_letter          varchar                               not null,
    sort                  integer                               not null,
    recommend_status      smallint    default 1                 not null,
    product_count         integer                               not null,
    product_comment_count integer                               not null,
    is_enabled            smallint    default 1                 not null,
    create_by             varchar     default ''                not null,
    create_time           timestamptz default current_timestamp not null,
    update_by             varchar                               not null,
    update_time           timestamptz                           not null,
    is_deleted            smallint    default 1                 not null

);

-- 添加商品品牌注释
comment on table pms_product_brand is '商品品牌';

-- 添加商品品牌列注释
comment on column pms_product_brand.id is '';
comment on column pms_product_brand.name is '品牌名称';
comment on column pms_product_brand.logo is '品牌logo';
comment on column pms_product_brand.big_pic is '专区大图';
comment on column pms_product_brand.description is '描述';
comment on column pms_product_brand.first_letter is '首字母';
comment on column pms_product_brand.sort is '排序';
comment on column pms_product_brand.recommend_status is '推荐状态';
comment on column pms_product_brand.product_count is '产品数量';
comment on column pms_product_brand.product_comment_count is '产品评论数量';
comment on column pms_product_brand.is_enabled is '是否启用';
comment on column pms_product_brand.create_by is '创建人ID';
comment on column pms_product_brand.create_time is '创建时间';
comment on column pms_product_brand.update_by is '更新人ID';
comment on column pms_product_brand.update_time is '更新时间';
comment on column pms_product_brand.is_deleted is '是否删除';

-- 产品分类建语句
drop table if exists pms_product_category;
create table pms_product_category
(
    id            bigserial primary KEY,
    parent_id     bigint                                not null,
    name          varchar                               not null,
    level         smallint    default 1                 not null,
    product_count integer                               not null,
    product_unit  varchar                               not null,
    nav_status    smallint    default 1                 not null,
    sort          integer                               not null,
    icon          varchar                               not null,
    keywords      varchar                               not null,
    description   varchar                               not null,
    is_enabled    smallint    default 1                 not null,
    create_by     varchar     default ''                not null,
    create_time   timestamptz default current_timestamp not null,
    update_by     varchar                               not null,
    update_time   timestamptz                           not null,
    is_deleted    smallint    default 1                 not null
);

-- 添加产品分类注释
comment on table pms_product_category is '产品分类';

-- 添加产品分类列注释
comment on column pms_product_category.id is '';
comment on column pms_product_category.parent_id is '上级分类的编号：0表示一级分类';
comment on column pms_product_category.name is '商品分类名称';
comment on column pms_product_category.level is '分类级别：0->1级；1->2级';
comment on column pms_product_category.product_count is '商品数量';
comment on column pms_product_category.product_unit is '商品单位';
comment on column pms_product_category.nav_status is '是否显示在导航栏：0->不显示；1->显示';
comment on column pms_product_category.sort is '排序';
comment on column pms_product_category.icon is '图标';
comment on column pms_product_category.keywords is '关键字';
comment on column pms_product_category.description is '描述';
comment on column pms_product_category.is_enabled is '是否启用';
comment on column pms_product_category.create_by is '创建人ID';
comment on column pms_product_category.create_time is '创建时间';
comment on column pms_product_category.update_by is '更新人ID';
comment on column pms_product_category.update_time is '更新时间';
comment on column pms_product_category.is_deleted is '是否删除';

-- 产品的分类和属性的关系表，用于设置分类筛选条件（只支持一级分类）建语句
drop table if exists pms_product_category_attribute_relation;
create table pms_product_category_attribute_relation
(
    id                   bigserial primary KEY,
    product_category_id  bigint not null,
    product_attribute_id bigint not null
);

-- 添加产品的分类和属性的关系表，用于设置分类筛选条件（只支持一级分类）注释
comment on table pms_product_category_attribute_relation is '产品的分类和属性的关系表，用于设置分类筛选条件（只支持一级分类）';

-- 添加产品的分类和属性的关系表，用于设置分类筛选条件（只支持一级分类）列注释
comment on column pms_product_category_attribute_relation.id is '';
comment on column pms_product_category_attribute_relation.product_category_id is '商品分类id';
comment on column pms_product_category_attribute_relation.product_attribute_id is '商品属性id';

-- 产品满减表(只针对同商品)建语句
drop table if exists pms_product_full_reduction;
create table pms_product_full_reduction
(
    id           bigserial primary KEY,
    product_id   bigint not null,
    full_price   bigint not null,
    reduce_price bigint not null
);

-- 添加产品满减表(只针对同商品)注释
comment on table pms_product_full_reduction is '产品满减表(只针对同商品)';

-- 添加产品满减表(只针对同商品)列注释
comment on column pms_product_full_reduction.id is '';
comment on column pms_product_full_reduction.product_id is '商品id';
comment on column pms_product_full_reduction.full_price is '商品满多少';
comment on column pms_product_full_reduction.reduce_price is '商品减多少';

-- 产品阶梯价格表(只针对同商品)建语句
drop table if exists pms_product_ladder;
create table pms_product_ladder
(
    id         bigserial primary KEY,
    product_id bigint  not null,
    count      integer not null,
    discount   bigint  not null,
    price      bigint  not null
);

-- 添加产品阶梯价格表(只针对同商品)注释
comment on table pms_product_ladder is '产品阶梯价格表(只针对同商品)';

-- 添加产品阶梯价格表(只针对同商品)列注释
comment on column pms_product_ladder.id is '';
comment on column pms_product_ladder.product_id is '商品id';
comment on column pms_product_ladder.count is '满足的商品数量';
comment on column pms_product_ladder.discount is '折扣';
comment on column pms_product_ladder.price is '折后价格';

-- 商品SKU表建语句
drop table if exists pms_product_sku;
create table pms_product_sku
(
    id                   bigserial primary KEY,
    spu_id               bigint                                not null,
    name                 varchar                               not null,
    sku_code             varchar                               not null,
    main_pic             varchar                               not null,
    album_pics           varchar                               not null,
    price                numeric                               not null,
    promotion_price      numeric                               not null,
    promotion_start_time timestamptz                           not null,
    promotion_end_time   timestamptz                           not null,
    stock                integer                               not null,
    low_stock            integer                               not null,
    spec_data            jsonb                                 not null,
    weight               numeric                               not null,
    publish_status       smallint    default 1                 not null,
    verify_status        smallint    default 1                 not null,
    sort                 integer                               not null,
    sales                integer                               not null,
    create_by            varchar     default ''                not null,
    create_time          timestamptz default current_timestamp not null,
    update_by            varchar                               not null,
    update_time          timestamptz                           not null,
    is_deleted           smallint    default 1                 not null
);

-- 添加商品SKU表注释
comment on table pms_product_sku is '商品SKU表';

-- 添加商品SKU表列注释
comment on column pms_product_sku.id is '商品SkuId';
comment on column pms_product_sku.spu_id is '商品SpuId';
comment on column pms_product_sku.name is 'SKU名称';
comment on column pms_product_sku.sku_code is 'SKU编码';
comment on column pms_product_sku.main_pic is '主图';
comment on column pms_product_sku.album_pics is '图片集';
comment on column pms_product_sku.price is '价格';
comment on column pms_product_sku.promotion_price is '单品促销价格';
comment on column pms_product_sku.promotion_start_time is '促销开始时间';
comment on column pms_product_sku.promotion_end_time is '促销结束时间';
comment on column pms_product_sku.stock is '库存';
comment on column pms_product_sku.low_stock is '预警库存';
comment on column pms_product_sku.spec_data is '规格数据';
comment on column pms_product_sku.weight is '重量(kg)';
comment on column pms_product_sku.publish_status is '上架状态：0-下架，1-上架';
comment on column pms_product_sku.verify_status is '审核状态：0-未审核，1-审核通过，2-审核不通过';
comment on column pms_product_sku.sort is '排序';
comment on column pms_product_sku.sales is '销量';
comment on column pms_product_sku.create_by is '创建人ID';
comment on column pms_product_sku.create_time is '创建时间';
comment on column pms_product_sku.update_by is '更新人ID';
comment on column pms_product_sku.update_time is '更新时间';
comment on column pms_product_sku.is_deleted is '是否删除';

-- 商品规格表建语句
drop table if exists pms_product_spec;
create table pms_product_spec
(
    id          bigserial primary KEY,
    category_id bigint                                not null,
    name        varchar                               not null,
    sort        integer                               not null,
    status      smallint    default 1                 not null,
    create_by   varchar     default ''                not null,
    create_time timestamptz default current_timestamp not null,
    update_by   varchar                               not null,
    update_time timestamptz                           not null,
    is_deleted  smallint    default 1                 not null
);

-- 添加商品规格表注释
comment on table pms_product_spec is '商品规格表';

-- 添加商品规格表列注释
comment on column pms_product_spec.id is '';
comment on column pms_product_spec.category_id is '分类ID';
comment on column pms_product_spec.name is '规格名称';
comment on column pms_product_spec.sort is '排序';
comment on column pms_product_spec.status is '状态：0->禁用；1->启用';
comment on column pms_product_spec.create_by is '创建人ID';
comment on column pms_product_spec.create_time is '创建时间';
comment on column pms_product_spec.update_by is '更新人ID';
comment on column pms_product_spec.update_time is '更新时间';
comment on column pms_product_spec.is_deleted is '是否删除';

-- 商品规格值表建语句
drop table if exists pms_product_spec_value;
create table pms_product_spec_value
(
    id          bigserial primary KEY,
    spec_id     bigint                                not null,
    value       varchar                               not null,
    sort        integer                               not null,
    status      smallint    default 1                 not null,
    create_by   varchar     default ''                not null,
    create_time timestamptz default current_timestamp not null,
    update_by   varchar                               not null,
    update_time timestamptz                           not null,
    is_deleted  smallint    default 1                 not null
);

-- 添加商品规格值表注释
comment on table pms_product_spec_value is '商品规格值表';

-- 添加商品规格值表列注释
comment on column pms_product_spec_value.id is '';
comment on column pms_product_spec_value.spec_id is '规格ID';
comment on column pms_product_spec_value.value is '规格值';
comment on column pms_product_spec_value.sort is '排序';
comment on column pms_product_spec_value.status is '状态：0->禁用；1->启用';
comment on column pms_product_spec_value.create_by is '创建人ID';
comment on column pms_product_spec_value.create_time is '创建时间';
comment on column pms_product_spec_value.update_by is '更新人ID';
comment on column pms_product_spec_value.update_time is '更新时间';
comment on column pms_product_spec_value.is_deleted is '是否删除';

-- 商品SPU表建语句
drop table if exists pms_product_spu;
create table pms_product_spu
(
    id                    bigserial primary KEY,
    name                  varchar                               not null,
    subTitle              varchar                               not null,
    product_sn            varchar                               not null,
    category_id           bigint                                not null,
    category_ids          varchar                               not null,
    category_name         varchar                               not null,
    brand_id              bigint                                not null,
    brand_name            varchar                               not null,
    unit                  varchar                               not null,
    weight                numeric                               not null,
    keywords              varchar                               not null,
    album_pics            varchar                               not null,
    main_pic              varchar                               not null,
    price_range           varchar                               not null,
    publish_status        smallint    default 1                 not null,
    new_status            smallint    default 1                 not null,
    recommend_status      smallint    default 1                 not null,
    verify_status         smallint    default 1                 not null,
    preview_status        smallint    default 1                 not null,
    sort                  integer                               not null,
    new_status_sort       integer                               not null,
    recommend_status_sort integer                               not null,
    sales                 integer                               not null,
    stock                 integer                               not null,
    low_stock             integer                               not null,
    promotion_type        smallint    default 1                 not null,
    detail_html           text                                  not null,
    detail_mobile_html    text                                  not null,
    create_by             varchar     default ''                not null,
    create_time           timestamptz default current_timestamp not null,
    update_by             varchar                               not null,
    update_time           timestamptz                           not null,
    is_deleted            smallint    default 1                 not null
);

-- 添加商品SPU表注释
comment on table pms_product_spu is '商品SPU表';

-- 添加商品SPU表列注释
comment on column pms_product_spu.id is '商品SpuId';
comment on column pms_product_spu.name is '商品名称';
comment on column pms_product_spu.subTitle is '副标题';
comment on column pms_product_spu.product_sn is '商品货号';
comment on column pms_product_spu.category_id is '商品分类ID';
comment on column pms_product_spu.category_ids is '商品分类ID集合';
comment on column pms_product_spu.category_name is '商品分类名称';
comment on column pms_product_spu.brand_id is '品牌ID';
comment on column pms_product_spu.brand_name is '品牌名称';
comment on column pms_product_spu.unit is '单位';
comment on column pms_product_spu.weight is '重量(kg)';
comment on column pms_product_spu.keywords is '关键词';
comment on column pms_product_spu.album_pics is '画册图片，最多8张，以逗号分割';
comment on column pms_product_spu.main_pic is '主图';
comment on column pms_product_spu.price_range is '价格区间';
comment on column pms_product_spu.publish_status is '上架状态：0-下架，1-上架';
comment on column pms_product_spu.new_status is '新品状态:0->不是新品；1->新品';
comment on column pms_product_spu.recommend_status is '推荐状态；0->不推荐；1->推荐';
comment on column pms_product_spu.verify_status is '审核状态：0->未审核；1->审核通过';
comment on column pms_product_spu.preview_status is '是否为预告商品：0->不是；1->是';
comment on column pms_product_spu.sort is '排序';
comment on column pms_product_spu.new_status_sort is '新品排序';
comment on column pms_product_spu.recommend_status_sort is '推荐排序';
comment on column pms_product_spu.sales is '销量';
comment on column pms_product_spu.stock is '库存';
comment on column pms_product_spu.low_stock is '预警库存';
comment on column pms_product_spu.promotion_type is '促销类型：0->没有促销使用原价;1->使用促销价；2->使用会员价；3->使用阶梯价格；4->使用满减价格；5->秒杀';
comment on column pms_product_spu.detail_html is '网页详情';
comment on column pms_product_spu.detail_mobile_html is '移动端详情';
comment on column pms_product_spu.create_by is '创建人ID';
comment on column pms_product_spu.create_time is '创建时间';
comment on column pms_product_spu.update_by is '更新人ID';
comment on column pms_product_spu.update_time is '更新时间';
comment on column pms_product_spu.is_deleted is '是否删除';

-- 购物车表建语句
drop table if exists oms_cart_item;
create table oms_cart_item
(
    id                  bigserial primary KEY,
    member_id           bigint                                not null,
    product_id          bigint                                not null,
    product_sku_id      bigint                                not null,
    quantity            integer                               not null,
    price               numeric                               not null,
    selected            smallint    default 1                 not null,
    product_name        varchar                               not null,
    product_sub_title   varchar                               not null,
    product_pic         text                                  not null,
    product_sku_code    varchar                               not null,
    product_sn          varchar                               not null,
    product_brand       varchar                               not null,
    product_category_id bigint                                not null,
    product_attr        jsonb                                 not null,
    member_nickname     varchar                               not null,
    source              smallint    default 1                 not null,
    delete_status       smallint    default 1                 not null,
    expire_time         timestamptz                           not null,
    create_time         timestamptz default current_timestamp not null,
    update_time         timestamptz                           not null
);

-- 添加购物车表注释
comment on table oms_cart_item is '购物车表';

-- 添加购物车表列注释
comment on column oms_cart_item.id is '主键ID';
comment on column oms_cart_item.member_id is '会员ID';
comment on column oms_cart_item.product_id is '商品ID';
comment on column oms_cart_item.product_sku_id is '商品SKU ID';
comment on column oms_cart_item.quantity is '购买数量';
comment on column oms_cart_item.price is '添加到购物车时的价格';
comment on column oms_cart_item.selected is '是否选中 0-未选中 1-选中';
comment on column oms_cart_item.product_name is '商品名称';
comment on column oms_cart_item.product_sub_title is '商品副标题';
comment on column oms_cart_item.product_pic is '商品主图URL';
comment on column oms_cart_item.product_sku_code is '商品SKU编码';
comment on column oms_cart_item.product_sn is '商品货号';
comment on column oms_cart_item.product_brand is '商品品牌';
comment on column oms_cart_item.product_category_id is '商品分类ID';
comment on column oms_cart_item.product_attr is '商品销售属性JSON';
comment on column oms_cart_item.member_nickname is '会员昵称';
comment on column oms_cart_item.source is '来源 1-PC 2-H5 3-小程序 4-APP';
comment on column oms_cart_item.delete_status is '删除状态 0-正常 1-删除';
comment on column oms_cart_item.expire_time is '过期时间';
comment on column oms_cart_item.create_time is '创建时间';
comment on column oms_cart_item.update_time is '更新时间';

-- 公司收发货地址表建语句
drop table if exists oms_company_address;
create table oms_company_address
(
    id             bigserial primary KEY,
    address_name   varchar                               not null,
    name           varchar                               not null,
    phone          varchar                               not null,
    province       varchar                               not null,
    city           varchar                               not null,
    region         varchar                               not null,
    detail_address varchar                               not null,
    send_status    smallint    default 1                 not null,
    receive_status smallint    default 1                 not null,
    create_by      varchar     default ''                not null,
    create_time    timestamptz default current_timestamp not null,
    update_by      varchar                               not null,
    update_time    timestamptz                           not null,
    is_deleted     smallint    default 1                 not null
);

-- 添加公司收发货地址表注释
comment on table oms_company_address is '公司收发货地址表';

-- 添加公司收发货地址表列注释
comment on column oms_company_address.id is '主键ID';
comment on column oms_company_address.address_name is '地址名称';
comment on column oms_company_address.name is '收发货人姓名';
comment on column oms_company_address.phone is '收货人电话';
comment on column oms_company_address.province is '省/直辖市';
comment on column oms_company_address.city is '市';
comment on column oms_company_address.region is '区';
comment on column oms_company_address.detail_address is '详细地址';
comment on column oms_company_address.send_status is '默认发货地址：0->否；1->是';
comment on column oms_company_address.receive_status is '默认收货地址：0->否；1->是';
comment on column oms_company_address.create_by is '创建人ID';
comment on column oms_company_address.create_time is '创建时间';
comment on column oms_company_address.update_by is '更新人ID';
comment on column oms_company_address.update_time is '更新时间';
comment on column oms_company_address.is_deleted is '是否删除';

-- 订单表建语句
drop table if exists oms_order;
create table oms_order
(
    id                      bigserial primary KEY,
    member_id               bigint                                not null,
    coupon_id               bigint                                not null,
    order_sn                varchar                               not null,
    create_time             timestamptz default current_timestamp not null,
    member_username         varchar                               not null,
    total_amount            bigint                                not null,
    pay_amount              bigint                                not null,
    freight_amount          bigint                                not null,
    promotion_amount        bigint                                not null,
    integration_amount      bigint                                not null,
    coupon_amount           bigint                                not null,
    discount_amount         bigint                                not null,
    pay_type                smallint    default 1                 not null,
    source_type             smallint    default 1                 not null,
    status                  smallint    default 1                 not null,
    order_type              smallint    default 1                 not null,
    delivery_company        varchar                               not null,
    delivery_sn             varchar                               not null,
    auto_confirm_day        integer                               not null,
    integration             integer                               not null,
    growth                  integer                               not null,
    promotion_info          varchar                               not null,
    bill_type               smallint    default 1                 not null,
    bill_header             varchar                               not null,
    bill_content            varchar                               not null,
    bill_receiver_phone     varchar                               not null,
    bill_receiver_email     varchar                               not null,
    receiver_name           varchar                               not null,
    receiver_phone          varchar                               not null,
    receiver_post_code      varchar                               not null,
    receiver_province       varchar                               not null,
    receiver_city           varchar                               not null,
    receiver_region         varchar                               not null,
    receiver_detail_address varchar                               not null,
    note                    varchar                               not null,
    confirm_status          smallint    default 1                 not null,
    delete_status           smallint    default 1                 not null,
    use_integration         integer                               not null,
    payment_time            timestamptz                           not null,
    delivery_time           timestamptz                           not null,
    receive_time            timestamptz                           not null,
    comment_time            timestamptz                           not null,
    modify_time             timestamptz                           not null
);

-- 添加订单表注释
comment on table oms_order is '订单表';

-- 添加订单表列注释
comment on column oms_order.id is '订单id';
comment on column oms_order.member_id is '会员id';
comment on column oms_order.coupon_id is '优惠券id';
comment on column oms_order.order_sn is '订单编号';
comment on column oms_order.create_time is '提交时间';
comment on column oms_order.member_username is '用户帐号';
comment on column oms_order.total_amount is '订单总金额';
comment on column oms_order.pay_amount is '应付金额（实际支付金额）';
comment on column oms_order.freight_amount is '运费金额';
comment on column oms_order.promotion_amount is '促销优化金额（促销价、满减、阶梯价）';
comment on column oms_order.integration_amount is '积分抵扣金额';
comment on column oms_order.coupon_amount is '优惠券抵扣金额';
comment on column oms_order.discount_amount is '管理员后台调整订单使用的折扣金额';
comment on column oms_order.pay_type is '支付方式：0->未支付；1->支付宝；2->微信';
comment on column oms_order.source_type is '订单来源：0->PC订单；1->app订单';
comment on column oms_order.status is '订单状态：0->待付款；1->待发货；2->已发货；3->已完成；4->已关闭；5->无效订单';
comment on column oms_order.order_type is '订单类型：0->正常订单；1->秒杀订单';
comment on column oms_order.delivery_company is '物流公司(配送方式)';
comment on column oms_order.delivery_sn is '物流单号';
comment on column oms_order.auto_confirm_day is '自动确认时间（天）';
comment on column oms_order.integration is '可以获得的积分';
comment on column oms_order.growth is '可以活动的成长值';
comment on column oms_order.promotion_info is '活动信息';
comment on column oms_order.bill_type is '发票类型：0->不开发票；1->电子发票；2->纸质发票';
comment on column oms_order.bill_header is '发票抬头';
comment on column oms_order.bill_content is '发票内容';
comment on column oms_order.bill_receiver_phone is '收票人电话';
comment on column oms_order.bill_receiver_email is '收票人邮箱';
comment on column oms_order.receiver_name is '收货人姓名';
comment on column oms_order.receiver_phone is '收货人电话';
comment on column oms_order.receiver_post_code is '收货人邮编';
comment on column oms_order.receiver_province is '省份/直辖市';
comment on column oms_order.receiver_city is '城市';
comment on column oms_order.receiver_region is '区';
comment on column oms_order.receiver_detail_address is '详细地址';
comment on column oms_order.note is '订单备注';
comment on column oms_order.confirm_status is '确认收货状态：0->未确认；1->已确认';
comment on column oms_order.delete_status is '删除状态：0->未删除；1->已删除';
comment on column oms_order.use_integration is '下单时使用的积分';
comment on column oms_order.payment_time is '支付时间';
comment on column oms_order.delivery_time is '发货时间';
comment on column oms_order.receive_time is '确认收货时间';
comment on column oms_order.comment_time is '评价时间';
comment on column oms_order.modify_time is '修改时间';

-- 订单收货地址表建语句
drop table if exists oms_order_delivery;
create table oms_order_delivery
(
    id                bigserial primary KEY,
    order_id          bigint                                not null,
    order_no          varchar                               not null,
    receiver_name     varchar                               not null,
    receiver_phone    varchar                               not null,
    receiver_province varchar                               not null,
    receiver_city     varchar                               not null,
    receiver_district varchar                               not null,
    receiver_address  varchar                               not null,
    delivery_company  varchar                               not null,
    delivery_no       varchar                               not null,
    create_time       timestamptz default current_timestamp not null,
    update_time       timestamptz                           not null,
    is_deleted        smallint    default 1                 not null
);

-- 添加订单收货地址表注释
comment on table oms_order_delivery is '订单收货地址表';

-- 添加订单收货地址表列注释
comment on column oms_order_delivery.id is '';
comment on column oms_order_delivery.order_id is '订单ID';
comment on column oms_order_delivery.order_no is '订单编号';
comment on column oms_order_delivery.receiver_name is '收货人姓名';
comment on column oms_order_delivery.receiver_phone is '收货人电话';
comment on column oms_order_delivery.receiver_province is '省份';
comment on column oms_order_delivery.receiver_city is '城市';
comment on column oms_order_delivery.receiver_district is '区县';
comment on column oms_order_delivery.receiver_address is '详细地址';
comment on column oms_order_delivery.delivery_company is '物流公司';
comment on column oms_order_delivery.delivery_no is '物流单号';
comment on column oms_order_delivery.create_time is '创建时间';
comment on column oms_order_delivery.update_time is '更新时间';
comment on column oms_order_delivery.is_deleted is '是否删除';

-- 订单商品表建语句
drop table if exists oms_order_item;
create table oms_order_item
(
    id                bigserial primary KEY,
    order_id          bigint                                not null,
    order_no          varchar                               not null,
    order_item_status smallint    default 1                 not null,
    sku_id            bigint                                not null,
    sku_name          varchar                               not null,
    sku_pic           varchar                               not null,
    sku_price         numeric                               not null,
    sku_quantity      integer                               not null,
    spec_data         jsonb                                 not null,
    sku_total_amount  numeric                               not null,
    promotion_amount  numeric                               not null,
    coupon_amount     numeric                               not null,
    points_amount     numeric                               not null,
    discount_amount   numeric                               not null,
    real_amount       numeric                               not null,
    create_time       timestamptz default current_timestamp not null,
    is_deleted        smallint    default 1                 not null
);

-- 添加订单商品表注释
comment on table oms_order_item is '订单商品表';

-- 添加订单商品表列注释
comment on column oms_order_item.id is '';
comment on column oms_order_item.order_id is '订单ID';
comment on column oms_order_item.order_no is '订单编号';
comment on column oms_order_item.order_item_status is '订单商品状态：1-正常,2-退货申请中,3-已退货,4-已拒绝';
comment on column oms_order_item.sku_id is '商品SKU ID';
comment on column oms_order_item.sku_name is '商品名称';
comment on column oms_order_item.sku_pic is '商品图片';
comment on column oms_order_item.sku_price is '商品单价';
comment on column oms_order_item.sku_quantity is '商品数量';
comment on column oms_order_item.spec_data is '规格数据';
comment on column oms_order_item.sku_total_amount is '商品总金额';
comment on column oms_order_item.promotion_amount is '促销分摊金额';
comment on column oms_order_item.coupon_amount is '优惠券分摊金额';
comment on column oms_order_item.points_amount is '积分分摊金额';
comment on column oms_order_item.discount_amount is '优惠分摊金额';
comment on column oms_order_item.real_amount is '实付金额';
comment on column oms_order_item.create_time is '创建时间';
comment on column oms_order_item.is_deleted is '是否删除';

-- 订单主表建语句
drop table if exists oms_order_main;
create table oms_order_main
(
    id                   bigserial primary KEY,
    order_no             varchar                               not null,
    user_id              bigint                                not null,
    order_status         smallint    default 1                 not null,
    total_amount         numeric                               not null,
    promotion_amount     numeric                               not null,
    coupon_amount        numeric                               not null,
    points_amount        numeric                               not null,
    discount_amount      numeric                               not null,
    freight_amount       numeric                               not null,
    pay_amount           numeric                               not null,
    pay_type             smallint    default 1                 not null,
    pay_time             timestamptz                           not null,
    delivery_time        timestamptz                           not null,
    receive_time         timestamptz                           not null,
    comment_time         timestamptz                           not null,
    source_type          smallint    default 1                 not null,
    express_order_number varchar                               not null,
    use_points           integer                               not null,
    receive_status       smallint    default 1                 not null,
    remark               varchar                               not null,
    create_time          timestamptz default current_timestamp not null,
    update_time          timestamptz                           not null,
    is_deleted           smallint    default 1                 not null
);

-- 添加订单主表注释
comment on table oms_order_main is '订单主表';

-- 添加订单主表列注释
comment on column oms_order_main.id is '';
comment on column oms_order_main.order_no is '订单编号';
comment on column oms_order_main.user_id is '用户ID';
comment on column oms_order_main.order_status is '订单状态：1-待支付,2-已支付,3-已发货,4-已完成,5-已取消,6-已退款,7-售后中';
comment on column oms_order_main.total_amount is '订单总金额';
comment on column oms_order_main.promotion_amount is '促销金额';
comment on column oms_order_main.coupon_amount is '优惠券金额';
comment on column oms_order_main.points_amount is '积分金额';
comment on column oms_order_main.discount_amount is '优惠金额';
comment on column oms_order_main.freight_amount is '运费金额';
comment on column oms_order_main.pay_amount is '实付金额';
comment on column oms_order_main.pay_type is '支付方式：1-支付宝,2-微信,3-银联';
comment on column oms_order_main.pay_time is '支付时间';
comment on column oms_order_main.delivery_time is '发货时间';
comment on column oms_order_main.receive_time is '收货时间';
comment on column oms_order_main.comment_time is '评价时间';
comment on column oms_order_main.source_type is '订单来源：1-APP,2-PC,3-小程序';
comment on column oms_order_main.express_order_number is '快递单号';
comment on column oms_order_main.use_points is '下单时使用的积分';
comment on column oms_order_main.receive_status is '是否确认收货：0->否,1->是';
comment on column oms_order_main.remark is '订单备注';
comment on column oms_order_main.create_time is '提交时间';
comment on column oms_order_main.update_time is '';
comment on column oms_order_main.is_deleted is '是否删除';

-- 订单操作记录表建语句
drop table if exists oms_order_operation_log;
create table oms_order_operation_log
(
    id             bigserial primary KEY,
    order_id       bigint                                not null,
    order_no       varchar                               not null,
    operation_type smallint    default 1                 not null,
    operator_id    bigint                                not null,
    operator_type  smallint    default 1                 not null,
    operator_note  varchar                               not null,
    create_time    timestamptz default current_timestamp not null
);

-- 添加订单操作记录表注释
comment on table oms_order_operation_log is '订单操作记录表';

-- 添加订单操作记录表列注释
comment on column oms_order_operation_log.id is '主键ID';
comment on column oms_order_operation_log.order_id is '订单ID';
comment on column oms_order_operation_log.order_no is '订单编号';
comment on column oms_order_operation_log.operation_type is '操作类型：1-创建订单，2-支付订单，3-发货，4-确认收货，5-取消订单，6-退款';
comment on column oms_order_operation_log.operator_id is '操作人ID';
comment on column oms_order_operation_log.operator_type is '操作人类型：1-用户，2-系统，3-管理员';
comment on column oms_order_operation_log.operator_note is '操作备注';
comment on column oms_order_operation_log.create_time is '操作时间';

-- 订单支付表建语句
drop table if exists oms_order_payment;
create table oms_order_payment
(
    id             bigserial primary KEY,
    order_id       bigint                                not null,
    order_no       varchar                               not null,
    pay_type       smallint    default 1                 not null,
    transaction_id varchar                               not null,
    total_amount   numeric                               not null,
    pay_amount     numeric                               not null,
    pay_status     smallint    default 1                 not null,
    pay_time       timestamptz                           not null,
    create_time    timestamptz default current_timestamp not null,
    update_time    timestamptz                           not null,
    is_deleted     smallint    default 1                 not null
);

-- 添加订单支付表注释
comment on table oms_order_payment is '订单支付表';

-- 添加订单支付表列注释
comment on column oms_order_payment.id is '主键ID';
comment on column oms_order_payment.order_id is '订单ID';
comment on column oms_order_payment.order_no is '订单编号';
comment on column oms_order_payment.pay_type is '支付方式：1-支付宝，2-微信，3-银联';
comment on column oms_order_payment.transaction_id is '支付流水号';
comment on column oms_order_payment.total_amount is '订单金额';
comment on column oms_order_payment.pay_amount is '支付金额';
comment on column oms_order_payment.pay_status is '支付状态：0-待支付，1-支付成功，2-支付失败';
comment on column oms_order_payment.pay_time is '支付时间';
comment on column oms_order_payment.create_time is '创建时间';
comment on column oms_order_payment.update_time is '';
comment on column oms_order_payment.is_deleted is '是否删除';

-- 订单优惠信息表建语句
drop table if exists oms_order_promotion;
create table oms_order_promotion
(
    id              bigserial primary KEY,
    order_id        bigint                                not null,
    order_no        varchar                               not null,
    promotion_type  smallint    default 1                 not null,
    promotion_id    bigint                                not null,
    promotion_name  varchar                               not null,
    discount_amount numeric                               not null,
    create_time     timestamptz default current_timestamp not null,
    is_deleted      smallint    default 1                 not null
);

-- 添加订单优惠信息表注释
comment on table oms_order_promotion is '订单优惠信息表';

-- 添加订单优惠信息表列注释
comment on column oms_order_promotion.id is '主键ID';
comment on column oms_order_promotion.order_id is '订单ID';
comment on column oms_order_promotion.order_no is '订单编号';
comment on column oms_order_promotion.promotion_type is '优惠类型：1-优惠券，2-积分抵扣，3-会员折扣，4-促销活动';
comment on column oms_order_promotion.promotion_id is '优惠ID';
comment on column oms_order_promotion.promotion_name is '优惠名称';
comment on column oms_order_promotion.discount_amount is '优惠金额';
comment on column oms_order_promotion.create_time is '';
comment on column oms_order_promotion.is_deleted is '是否删除';

-- 退货/售后主表建语句
drop table if exists oms_order_return;
create table oms_order_return
(
    id              bigserial primary KEY,
    order_id        bigint                                not null,
    return_no       varchar                               not null,
    member_id       bigint                                not null,
    status          smallint    default 1                 not null,
    type            smallint    default 1                 not null,
    reason          varchar                               not null,
    description     varchar                               not null,
    proof_pic       varchar                               not null,
    refund_amount   numeric                               not null,
    return_name     varchar                               not null,
    return_phone    varchar                               not null,
    company_address varchar                               not null,
    create_time     timestamptz default current_timestamp not null,
    handle_time     timestamptz                           not null,
    handle_note     varchar                               not null,
    handle_man      varchar                               not null,
    receive_time    timestamptz                           not null,
    receive_note    varchar                               not null,
    receive_man     varchar                               not null,
    refund_time     timestamptz                           not null,
    close_time      timestamptz                           not null,
    remark          varchar                               not null
);

-- 添加退货/售后主表注释
comment on table oms_order_return is '退货/售后主表';

-- 添加退货/售后主表列注释
comment on column oms_order_return.id is '主键ID';
comment on column oms_order_return.order_id is '关联订单ID';
comment on column oms_order_return.return_no is '退货单号';
comment on column oms_order_return.member_id is '会员ID';
comment on column oms_order_return.status is '退货状态（0待审核 1审核通过 2已收货 3已退款 4已拒绝 5已关闭）';
comment on column oms_order_return.type is '售后类型（0退货退款 1仅退款 2换货）';
comment on column oms_order_return.reason is '退货原因';
comment on column oms_order_return.description is '问题描述';
comment on column oms_order_return.proof_pic is '凭证图片，逗号分隔';
comment on column oms_order_return.refund_amount is '退款金额';
comment on column oms_order_return.return_name is '退货人姓名';
comment on column oms_order_return.return_phone is '退货人电话';
comment on column oms_order_return.company_address is '退货收货地址';
comment on column oms_order_return.create_time is '申请时间';
comment on column oms_order_return.handle_time is '处理时间';
comment on column oms_order_return.handle_note is '处理备注';
comment on column oms_order_return.handle_man is '处理人员';
comment on column oms_order_return.receive_time is '收货时间';
comment on column oms_order_return.receive_note is '收货备注';
comment on column oms_order_return.receive_man is '收货人';
comment on column oms_order_return.refund_time is '退款时间';
comment on column oms_order_return.close_time is '关闭时间';
comment on column oms_order_return.remark is '备注';

-- 退货/售后明细表建语句
drop table if exists oms_order_return_item;
create table oms_order_return_item
(
    id            bigserial primary KEY,
    return_id     bigint  not null,
    order_id      bigint  not null,
    order_item_id bigint  not null,
    sku_id        bigint  not null,
    sku_name      varchar not null,
    sku_pic       varchar not null,
    sku_attrs     varchar not null,
    quantity      integer not null,
    product_price numeric not null,
    real_amount   numeric not null,
    reason        varchar not null,
    remark        varchar not null
);

-- 添加退货/售后明细表注释
comment on table oms_order_return_item is '退货/售后明细表';

-- 添加退货/售后明细表列注释
comment on column oms_order_return_item.id is '主键ID';
comment on column oms_order_return_item.return_id is '退货单ID（关联oms_order_return.id）';
comment on column oms_order_return_item.order_id is '订单ID';
comment on column oms_order_return_item.order_item_id is '订单明细ID';
comment on column oms_order_return_item.sku_id is '商品SKU ID';
comment on column oms_order_return_item.sku_name is '商品名称';
comment on column oms_order_return_item.sku_pic is '商品图片';
comment on column oms_order_return_item.sku_attrs is '商品销售属性';
comment on column oms_order_return_item.quantity is '退货数量';
comment on column oms_order_return_item.product_price is '商品单价';
comment on column oms_order_return_item.real_amount is '实际退款金额';
comment on column oms_order_return_item.reason is '退货原因';
comment on column oms_order_return_item.remark is '备注';

-- 退货原因表建语句
drop table if exists oms_order_return_reason;
create table oms_order_return_reason
(
    id          bigserial primary KEY,
    name        varchar                               not null,
    sort        integer                               not null,
    status      smallint    default 1                 not null,
    create_by   varchar     default ''                not null,
    create_time timestamptz default current_timestamp not null,
    update_by   varchar                               not null,
    update_time timestamptz                           not null,
    is_deleted  smallint    default 1                 not null
);

-- 添加退货原因表注释
comment on table oms_order_return_reason is '退货原因表';

-- 添加退货原因表列注释
comment on column oms_order_return_reason.id is '主键ID';
comment on column oms_order_return_reason.name is '退货类型';
comment on column oms_order_return_reason.sort is '排序';
comment on column oms_order_return_reason.status is '状态：0->不启用；1->启用';
comment on column oms_order_return_reason.create_by is '创建人ID';
comment on column oms_order_return_reason.create_time is '创建时间';
comment on column oms_order_return_reason.update_by is '更新人ID';
comment on column oms_order_return_reason.update_time is '更新时间';
comment on column oms_order_return_reason.is_deleted is '是否删除';

-- 订单设置表建语句
drop table if exists oms_order_setting;
create table oms_order_setting
(
    id                    bigserial primary KEY,
    flash_order_overtime  integer                               not null,
    normal_order_overtime integer                               not null,
    confirm_overtime      integer                               not null,
    finish_overtime       integer                               not null,
    status                smallint    default 1                 not null,
    is_default            smallint    default 1                 not null,
    comment_overtime      integer                               not null,
    create_by             varchar     default ''                not null,
    create_time           timestamptz default current_timestamp not null,
    update_by             varchar                               not null,
    update_time           timestamptz                           not null,
    is_deleted            smallint    default 1                 not null
);

-- 添加订单设置表注释
comment on table oms_order_setting is '订单设置表';

-- 添加订单设置表列注释
comment on column oms_order_setting.id is '主键ID';
comment on column oms_order_setting.flash_order_overtime is '秒杀订单超时关闭时间(分)';
comment on column oms_order_setting.normal_order_overtime is '正常订单超时时间(分)';
comment on column oms_order_setting.confirm_overtime is '发货后自动确认收货时间（天）';
comment on column oms_order_setting.finish_overtime is '自动完成交易时间，不能申请售后（天）';
comment on column oms_order_setting.status is '状态：0->禁用；1->启用';
comment on column oms_order_setting.is_default is '是否默认：0->否；1->是';
comment on column oms_order_setting.comment_overtime is '订单完成后自动好评时间（天）';
comment on column oms_order_setting.create_by is '创建人ID';
comment on column oms_order_setting.create_time is '创建时间';
comment on column oms_order_setting.update_by is '更新人ID';
comment on column oms_order_setting.update_time is '更新时间';
comment on column oms_order_setting.is_deleted is '是否删除';

-- 优惠券表建语句
drop table if exists sms_coupon;
create table sms_coupon
(
    id             bigserial primary KEY,
    type_id        bigint                                not null,
    name           varchar                               not null,
    code           varchar                               not null,
    amount         numeric                               not null,
    min_amount     numeric                               not null,
    start_time     timestamptz                           not null,
    end_time       timestamptz                           not null,
    total_count    integer                               not null,
    received_count integer                               not null,
    used_count     integer                               not null,
    per_limit      integer                               not null,
    status         smallint    default 1                 not null,
    is_enabled     smallint    default 1                 not null,
    description    varchar                               not null,
    create_by      varchar     default ''                not null,
    create_time    timestamptz default current_timestamp not null,
    update_by      varchar                               not null,
    update_time    timestamptz                           not null,
    is_deleted     smallint    default 1                 not null
);

-- 添加优惠券表注释
comment on table sms_coupon is '优惠券表';

-- 添加优惠券表列注释
comment on column sms_coupon.id is '优惠券ID';
comment on column sms_coupon.type_id is '优惠券类型ID';
comment on column sms_coupon.name is '优惠券名称';
comment on column sms_coupon.code is '优惠券码';
comment on column sms_coupon.amount is '优惠金额/折扣率';
comment on column sms_coupon.min_amount is '最低使用金额';
comment on column sms_coupon.start_time is '生效时间';
comment on column sms_coupon.end_time is '失效时间';
comment on column sms_coupon.total_count is '发放总量';
comment on column sms_coupon.received_count is '已领取数量';
comment on column sms_coupon.used_count is '已使用数量';
comment on column sms_coupon.per_limit is '每人限领数量';
comment on column sms_coupon.status is '状态：0-未开始，1-进行中，2-已结束，3-已取消';
comment on column sms_coupon.is_enabled is '是否启用';
comment on column sms_coupon.description is '使用说明';
comment on column sms_coupon.create_by is '创建人ID';
comment on column sms_coupon.create_time is '创建时间';
comment on column sms_coupon.update_by is '更新人ID';
comment on column sms_coupon.update_time is '更新时间';
comment on column sms_coupon.is_deleted is '是否删除';

-- 优惠券领取记录表建语句
drop table if exists sms_coupon_record;
create table sms_coupon_record
(
    id              bigserial primary KEY,
    coupon_id       bigint                                not null,
    member_id       bigint                                not null,
    get_time        timestamptz                           not null,
    get_type        smallint    default 1                 not null,
    use_time        timestamptz                           not null,
    order_id        bigint                                not null,
    order_amount    numeric                               not null,
    discount_amount numeric                               not null,
    status          smallint    default 1                 not null,
    invalid_time    timestamptz                           not null,
    invalid_reason  varchar                               not null,
    create_time     timestamptz default current_timestamp not null,
    is_deleted      smallint    default 1                 not null
);

-- 添加优惠券领取记录表注释
comment on table sms_coupon_record is '优惠券领取记录表';

-- 添加优惠券领取记录表列注释
comment on column sms_coupon_record.id is '主键ID';
comment on column sms_coupon_record.coupon_id is '优惠券ID';
comment on column sms_coupon_record.member_id is '用户ID';
comment on column sms_coupon_record.get_time is '领取时间';
comment on column sms_coupon_record.get_type is '获取类型：0->后台赠送；1->主动获取';
comment on column sms_coupon_record.use_time is '使用时间';
comment on column sms_coupon_record.order_id is '使用订单ID';
comment on column sms_coupon_record.order_amount is '订单金额';
comment on column sms_coupon_record.discount_amount is '优惠金额';
comment on column sms_coupon_record.status is '状态：0-未使用，1-已使用，2-已过期，3-已失效';
comment on column sms_coupon_record.invalid_time is '失效时间';
comment on column sms_coupon_record.invalid_reason is '失效原因';
comment on column sms_coupon_record.create_time is '创建时间';
comment on column sms_coupon_record.is_deleted is '是否删除';

-- 优惠券使用范围表建语句
drop table if exists sms_coupon_scope;
create table sms_coupon_scope
(
    id          bigserial primary KEY,
    coupon_id   bigint                                not null,
    scope_type  smallint    default 1                 not null,
    scope_id    bigint                                not null,
    create_time timestamptz default current_timestamp not null
);

-- 添加优惠券使用范围表注释
comment on table sms_coupon_scope is '优惠券使用范围表';

-- 添加优惠券使用范围表列注释
comment on column sms_coupon_scope.id is '主键ID';
comment on column sms_coupon_scope.coupon_id is '优惠券ID';
comment on column sms_coupon_scope.scope_type is '范围类型：0-全场，1-分类，2-商品';
comment on column sms_coupon_scope.scope_id is '范围ID（分类ID或商品ID）';
comment on column sms_coupon_scope.create_time is '创建时间';

-- 优惠券类型表建语句
drop table if exists sms_coupon_type;
create table sms_coupon_type
(
    id            bigserial primary KEY,
    name          varchar                               not null,
    code          varchar                               not null,
    description   varchar                               not null,
    discount_type smallint    default 1                 not null,
    status        smallint    default 1                 not null,
    sort          integer                               not null,
    create_by     varchar     default ''                not null,
    create_time   timestamptz default current_timestamp not null,
    update_by     varchar                               not null,
    update_time   timestamptz                           not null,
    is_deleted    smallint    default 1                 not null
);

-- 添加优惠券类型表注释
comment on table sms_coupon_type is '优惠券类型表';

-- 添加优惠券类型表列注释
comment on column sms_coupon_type.id is '主键ID';
comment on column sms_coupon_type.name is '类型名称';
comment on column sms_coupon_type.code is '类型编码';
comment on column sms_coupon_type.description is '描述';
comment on column sms_coupon_type.discount_type is '优惠方式：1-固定金额，2-折扣率，3-第N件特惠，4-买赠，5-特价，6-套装优惠，7-搭配优惠，8-积分抵现，9-积分倍率，10-免运费，11-运费减免，12-限时特权，13-会员特权';
comment on column sms_coupon_type.status is '是否启用';
comment on column sms_coupon_type.sort is '排序';
comment on column sms_coupon_type.create_by is '创建人ID';
comment on column sms_coupon_type.create_time is '创建时间';
comment on column sms_coupon_type.update_by is '更新人ID';
comment on column sms_coupon_type.update_time is '更新时间';
comment on column sms_coupon_type.is_deleted is '是否删除';

-- 首页轮播广告表建语句
drop table if exists sms_home_advertise;
create table sms_home_advertise
(
    id          bigserial primary KEY,
    name        varchar                               not null,
    type        smallint    default 1                 not null,
    pic         varchar                               not null,
    start_time  timestamptz                           not null,
    end_time    timestamptz                           not null,
    status      smallint    default 1                 not null,
    click_count integer                               not null,
    order_count integer                               not null,
    url         varchar                               not null,
    remark      varchar                               not null,
    sort        integer                               not null,
    create_time timestamptz default current_timestamp not null,
    update_time timestamptz                           not null,
    is_deleted  smallint    default 1                 not null
);

-- 添加首页轮播广告表注释
comment on table sms_home_advertise is '首页轮播广告表';

-- 添加首页轮播广告表列注释
comment on column sms_home_advertise.id is '编号';
comment on column sms_home_advertise.name is '名称';
comment on column sms_home_advertise.type is '轮播位置：0->PC首页轮播；1->app首页轮播';
comment on column sms_home_advertise.pic is '图片地址';
comment on column sms_home_advertise.start_time is '开始时间';
comment on column sms_home_advertise.end_time is '结束时间';
comment on column sms_home_advertise.status is '上下线状态：0->下线；1->上线';
comment on column sms_home_advertise.click_count is '点击数';
comment on column sms_home_advertise.order_count is '下单数';
comment on column sms_home_advertise.url is '链接地址';
comment on column sms_home_advertise.remark is '备注';
comment on column sms_home_advertise.sort is '排序';
comment on column sms_home_advertise.create_time is '创建时间';
comment on column sms_home_advertise.update_time is '更新时间';
comment on column sms_home_advertise.is_deleted is '是否删除';

-- 秒杀活动表建语句
drop table if exists sms_seckill_activity;
create table sms_seckill_activity
(
    id          bigserial primary KEY,
    name        varchar                               not null,
    description varchar                               not null,
    start_time  timestamptz                           not null,
    end_time    timestamptz                           not null,
    status      smallint    default 1                 not null,
    is_enabled  smallint    default 1                 not null,
    create_by   varchar     default ''                not null,
    create_time timestamptz default current_timestamp not null,
    update_by   varchar                               not null,
    update_time timestamptz                           not null,
    is_deleted  smallint    default 1                 not null
);

-- 添加秒杀活动表注释
comment on table sms_seckill_activity is '秒杀活动表';

-- 添加秒杀活动表列注释
comment on column sms_seckill_activity.id is '编号';
comment on column sms_seckill_activity.name is '活动名称';
comment on column sms_seckill_activity.description is '活动描述';
comment on column sms_seckill_activity.start_time is '开始时间';
comment on column sms_seckill_activity.end_time is '结束时间';
comment on column sms_seckill_activity.status is '状态:0-上线,1-下线';
comment on column sms_seckill_activity.is_enabled is '是否启用';
comment on column sms_seckill_activity.create_by is '创建人ID';
comment on column sms_seckill_activity.create_time is '创建时间';
comment on column sms_seckill_activity.update_by is '更新人ID';
comment on column sms_seckill_activity.update_time is '更新时间';
comment on column sms_seckill_activity.is_deleted is '是否删除';

-- 秒杀商品表建语句
drop table if exists sms_seckill_product;
create table sms_seckill_product
(
    id            bigserial primary KEY,
    activity_id   bigint                                not null,
    session_id    bigint                                not null,
    sku_id        bigint                                not null,
    sku_name      varchar                               not null,
    seckill_price numeric                               not null,
    seckill_stock integer                               not null,
    stock_locked  integer                               not null,
    per_limit     integer                               not null,
    sort          integer                               not null,
    status        smallint    default 1                 not null,
    create_by     varchar     default ''                not null,
    create_time   timestamptz default current_timestamp not null,
    update_by     varchar                               not null,
    update_time   timestamptz                           not null,
    is_deleted    smallint    default 1                 not null
);

-- 添加秒杀商品表注释
comment on table sms_seckill_product is '秒杀商品表';

-- 添加秒杀商品表列注释
comment on column sms_seckill_product.id is 'ID';
comment on column sms_seckill_product.activity_id is '活动ID';
comment on column sms_seckill_product.session_id is '秒杀场次ID';
comment on column sms_seckill_product.sku_id is '商品SKU ID';
comment on column sms_seckill_product.sku_name is '商品名称';
comment on column sms_seckill_product.seckill_price is '秒杀价格';
comment on column sms_seckill_product.seckill_stock is '秒杀库存';
comment on column sms_seckill_product.stock_locked is '锁定库存';
comment on column sms_seckill_product.per_limit is '每人限购数量';
comment on column sms_seckill_product.sort is '排序';
comment on column sms_seckill_product.status is '状态：0-未上架，1-已上架';
comment on column sms_seckill_product.create_by is '创建人ID';
comment on column sms_seckill_product.create_time is '创建时间';
comment on column sms_seckill_product.update_by is '更新人ID';
comment on column sms_seckill_product.update_time is '更新时间';
comment on column sms_seckill_product.is_deleted is '是否删除';

-- 秒杀预约表建语句
drop table if exists sms_seckill_reservation;
create table sms_seckill_reservation
(
    id          bigserial primary KEY,
    user_id     bigint                                not null,
    activity_id bigint                                not null,
    product_id  bigint                                not null,
    status      smallint    default 1                 not null,
    create_time timestamptz default current_timestamp not null,
    update_time timestamptz                           not null,
    is_deleted  smallint    default 1                 not null
);

-- 添加秒杀预约表注释
comment on table sms_seckill_reservation is '秒杀预约表';

-- 添加秒杀预约表列注释
comment on column sms_seckill_reservation.id is 'ID';
comment on column sms_seckill_reservation.user_id is '用户ID';
comment on column sms_seckill_reservation.activity_id is '活动ID';
comment on column sms_seckill_reservation.product_id is '秒杀商品ID';
comment on column sms_seckill_reservation.status is '状态：0-已预约，1-已参与，2-已取消';
comment on column sms_seckill_reservation.create_time is '';
comment on column sms_seckill_reservation.update_time is '';
comment on column sms_seckill_reservation.is_deleted is '是否删除';

-- 秒杀场次表建语句
drop table if exists sms_seckill_session;
create table sms_seckill_session
(
    id          bigserial primary KEY,
    name        varchar                               not null,
    start_time  varchar                               not null,
    end_time    varchar                               not null,
    status      smallint    default 1                 not null,
    sort        integer                               not null,
    create_by   varchar     default ''                not null,
    create_time timestamptz default current_timestamp not null,
    update_by   varchar                               not null,
    update_time timestamptz                           not null,
    is_deleted  smallint    default 1                 not null
);

-- 添加秒杀场次表注释
comment on table sms_seckill_session is '秒杀场次表';

-- 添加秒杀场次表列注释
comment on column sms_seckill_session.id is '秒杀场次ID';
comment on column sms_seckill_session.name is '场次名称';
comment on column sms_seckill_session.start_time is '开始时间';
comment on column sms_seckill_session.end_time is '结束时间';
comment on column sms_seckill_session.status is '状态：0-禁用，1-启用';
comment on column sms_seckill_session.sort is '排序';
comment on column sms_seckill_session.create_by is '创建人ID';
comment on column sms_seckill_session.create_time is '创建时间';
comment on column sms_seckill_session.update_by is '更新人ID';
comment on column sms_seckill_session.update_time is '更新时间';
comment on column sms_seckill_session.is_deleted is '是否删除';

-- 会员收货地址表建语句
drop table if exists ums_member_address;
create table ums_member_address
(
    id             bigserial primary KEY,
    member_id      bigint                                not null,
    receiver_name  varchar                               not null,
    receiver_phone varchar                               not null,
    province       varchar                               not null,
    city           varchar                               not null,
    district       varchar                               not null,
    detail_address varchar                               not null,
    postal_code    varchar                               not null,
    tag            varchar                               not null,
    is_default     smallint    default 1                 not null,
    create_time    timestamptz default current_timestamp not null,
    update_time    timestamptz                           not null,
    is_deleted     smallint    default 1                 not null
);

-- 添加会员收货地址表注释
comment on table ums_member_address is '会员收货地址表';

-- 添加会员收货地址表列注释
comment on column ums_member_address.id is '主键ID';
comment on column ums_member_address.member_id is '会员ID';
comment on column ums_member_address.receiver_name is '收货人姓名';
comment on column ums_member_address.receiver_phone is '收货人电话';
comment on column ums_member_address.province is '省份';
comment on column ums_member_address.city is '城市';
comment on column ums_member_address.district is '区县';
comment on column ums_member_address.detail_address is '详细地址';
comment on column ums_member_address.postal_code is '邮政编码';
comment on column ums_member_address.tag is '地址标签：家、公司等';
comment on column ums_member_address.is_default is '是否默认地址';
comment on column ums_member_address.create_time is '创建时间';
comment on column ums_member_address.update_time is '更新时间';
comment on column ums_member_address.is_deleted is '是否删除';

-- 积分消费设置建语句
drop table if exists ums_member_consume_setting;
create table ums_member_consume_setting
(
    id                    bigserial primary KEY,
    deduction_per_amount  integer                               not null,
    max_percent_per_order integer                               not null,
    use_unit              integer                               not null,
    coupon_status         smallint    default 1                 not null,
    status                smallint    default 1                 not null,
    create_by             varchar     default ''                not null,
    create_time           timestamptz default current_timestamp not null,
    update_by             varchar                               not null,
    update_time           timestamptz                           not null
);

-- 添加积分消费设置注释
comment on table ums_member_consume_setting is '积分消费设置';

-- 添加积分消费设置列注释
comment on column ums_member_consume_setting.id is '';
comment on column ums_member_consume_setting.deduction_per_amount is '每一元需要抵扣的积分数量';
comment on column ums_member_consume_setting.max_percent_per_order is '每笔订单最高抵用百分比';
comment on column ums_member_consume_setting.use_unit is '每次使用积分最小单位100';
comment on column ums_member_consume_setting.coupon_status is '是否可以和优惠券同用；0->不可以；1->可以';
comment on column ums_member_consume_setting.status is '状态：0->禁用；1->启用';
comment on column ums_member_consume_setting.create_by is '创建人ID';
comment on column ums_member_consume_setting.create_time is '创建时间';
comment on column ums_member_consume_setting.update_by is '更新人ID';
comment on column ums_member_consume_setting.update_time is '更新时间';

-- 会员成长值记录表建语句
drop table if exists ums_member_growth_log;
create table ums_member_growth_log
(
    id            bigserial primary KEY,
    member_id     bigint                                not null,
    change_type   smallint    default 1                 not null,
    change_growth integer                               not null,
    source_type   smallint    default 1                 not null,
    description   varchar                               not null,
    operate_man   varchar                               not null,
    operate_note  varchar                               not null,
    create_time   timestamptz default current_timestamp not null
);

-- 添加会员成长值记录表注释
comment on table ums_member_growth_log is '会员成长值记录表';

-- 添加会员成长值记录表列注释
comment on column ums_member_growth_log.id is '';
comment on column ums_member_growth_log.member_id is '会员ID';
comment on column ums_member_growth_log.change_type is '变更类型：1-添加成长值，2-减少成长值';
comment on column ums_member_growth_log.change_growth is '变更成长值';
comment on column ums_member_growth_log.source_type is '来源类型：0-其他，1-订单，2-活动，3-签到，4-管理员修改';
comment on column ums_member_growth_log.description is '描述';
comment on column ums_member_growth_log.operate_man is '操作人员';
comment on column ums_member_growth_log.operate_note is '操作备注';
comment on column ums_member_growth_log.create_time is '创建时间';

-- 会员信息表建语句
drop table if exists ums_member_info;
create table ums_member_info
(
    id                 bigserial primary KEY,
    member_id          bigint                                not null,
    wx_openid          varchar                               not null,
    level_id           bigint                                not null,
    nickname           varchar                               not null,
    mobile             varchar                               not null,
    source             smallint    default 1                 not null,
    password           varchar                               not null,
    avatar             varchar                               not null,
    signature          varchar                               not null,
    gender             smallint    default 1                 not null,
    birthday           date                                  not null,
    growth_point       integer                               not null,
    points             integer                               not null,
    total_points       integer                               not null,
    spend_amount       numeric                               not null,
    order_count        integer                               not null,
    coupon_count       integer                               not null,
    comment_count      integer                               not null,
    return_count       integer                               not null,
    lottery_times      integer                               not null,
    first_login_status smallint    default 1                 not null,
    last_login         timestamptz                           not null,
    is_enabled         smallint    default 1                 not null,
    create_time        timestamptz default current_timestamp not null,
    update_time        timestamptz                           not null,
    is_deleted         smallint    default 1                 not null
);

-- 添加会员信息表注释
comment on table ums_member_info is '会员信息表';

-- 添加会员信息表列注释
comment on column ums_member_info.id is '主键ID';
comment on column ums_member_info.member_id is '会员ID';
comment on column ums_member_info.wx_openid is '微信openid';
comment on column ums_member_info.level_id is '等级ID';
comment on column ums_member_info.nickname is '昵称';
comment on column ums_member_info.mobile is '手机号码';
comment on column ums_member_info.source is '注册来源：0-PC，1-APP，2-小程序';
comment on column ums_member_info.password is '密码';
comment on column ums_member_info.avatar is '头像';
comment on column ums_member_info.signature is '个性签名';
comment on column ums_member_info.gender is '性别：0-未知，1-男，2-女';
comment on column ums_member_info.birthday is '生日';
comment on column ums_member_info.growth_point is '成长值';
comment on column ums_member_info.points is '积分';
comment on column ums_member_info.total_points is '累计获得积分';
comment on column ums_member_info.spend_amount is '累计消费金额';
comment on column ums_member_info.order_count is '订单数';
comment on column ums_member_info.coupon_count is '优惠券数量';
comment on column ums_member_info.comment_count is '评价数';
comment on column ums_member_info.return_count is '退货数';
comment on column ums_member_info.lottery_times is '剩余抽奖次数';
comment on column ums_member_info.first_login_status is '1:未登录过,2:已登录过';
comment on column ums_member_info.last_login is '最后登录';
comment on column ums_member_info.is_enabled is '是否启用：0-禁用，1-启用';
comment on column ums_member_info.create_time is '创建时间';
comment on column ums_member_info.update_time is '更新时间';
comment on column ums_member_info.is_deleted is '是否删除';

-- 会员等级表建语句
drop table if exists ums_member_level;
create table ums_member_level
(
    id            bigserial primary KEY,
    name          varchar                               not null,
    level         integer                               not null,
    growth_point  integer                               not null,
    discount_rate numeric                               not null,
    free_freight  smallint    default 1                 not null,
    comment_extra smallint    default 1                 not null,
    privileges    varchar                               not null,
    remark        varchar                               not null,
    is_enabled    smallint    default 1                 not null,
    create_by     varchar     default ''                not null,
    create_time   timestamptz default current_timestamp not null,
    update_by     varchar                               not null,
    update_time   timestamptz                           not null,
    is_deleted    smallint    default 1                 not null
);

-- 添加会员等级表注释
comment on table ums_member_level is '会员等级表';

-- 添加会员等级表列注释
comment on column ums_member_level.id is '主键ID';
comment on column ums_member_level.name is '等级名称';
comment on column ums_member_level.level is '等级';
comment on column ums_member_level.growth_point is '升级所需成长值';
comment on column ums_member_level.discount_rate is '折扣率(0-100)';
comment on column ums_member_level.free_freight is '是否免运费';
comment on column ums_member_level.comment_extra is '是否可评论获取奖励';
comment on column ums_member_level.privileges is '会员特权JSON';
comment on column ums_member_level.remark is '备注';
comment on column ums_member_level.is_enabled is '是否启用';
comment on column ums_member_level.create_by is '创建人ID';
comment on column ums_member_level.create_time is '创建时间';
comment on column ums_member_level.update_by is '更新人ID';
comment on column ums_member_level.update_time is '更新时间';
comment on column ums_member_level.is_deleted is '是否删除';

-- 会员登录记录建语句
drop table if exists ums_member_login_log;
create table ums_member_login_log
(
    id          bigserial primary KEY,
    member_id   bigint                                not null,
    member_ip   varchar                               not null,
    city        varchar                               not null,
    login_type  smallint    default 1                 not null,
    province    varchar                               not null,
    create_time timestamptz default current_timestamp not null
);

-- 添加会员登录记录注释
comment on table ums_member_login_log is '会员登录记录';

-- 添加会员登录记录列注释
comment on column ums_member_login_log.id is '';
comment on column ums_member_login_log.member_id is '会员id';
comment on column ums_member_login_log.member_ip is '登录ip';
comment on column ums_member_login_log.city is '登录城市';
comment on column ums_member_login_log.login_type is '登录类型：0->PC；1->android;2->ios;3->小程序';
comment on column ums_member_login_log.province is '登录省份';
comment on column ums_member_login_log.create_time is '登录时间';

-- 会员积分记录表建语句
drop table if exists ums_member_points_log;
create table ums_member_points_log
(
    id            bigserial primary KEY,
    member_id     bigint                                not null,
    change_type   smallint    default 1                 not null,
    change_points integer                               not null,
    source_type   smallint    default 1                 not null,
    description   varchar                               not null,
    operate_man   varchar                               not null,
    operate_note  varchar                               not null,
    create_time   timestamptz default current_timestamp not null
);

-- 添加会员积分记录表注释
comment on table ums_member_points_log is '会员积分记录表';

-- 添加会员积分记录表列注释
comment on column ums_member_points_log.id is '';
comment on column ums_member_points_log.member_id is '会员ID';
comment on column ums_member_points_log.change_type is '变更类型：1-添加积分，2-减少积分';
comment on column ums_member_points_log.change_points is '变更积分';
comment on column ums_member_points_log.source_type is '来源类型：0-其他，1-订单，2-活动，3-签到，4-管理员修改';
comment on column ums_member_points_log.description is '描述';
comment on column ums_member_points_log.operate_man is '操作人员';
comment on column ums_member_points_log.operate_note is '操作备注';
comment on column ums_member_points_log.create_time is '创建时间';

-- 会员积分成长规则表建语句
drop table if exists ums_member_rule_setting;
create table ums_member_rule_setting
(
    id                  bigserial primary KEY,
    consume_per_point   bigint                                not null,
    low_order_amount    bigint                                not null,
    max_point_per_order integer                               not null,
    rule_type           smallint    default 1                 not null,
    status              smallint    default 1                 not null,
    create_by           varchar     default ''                not null,
    create_time         timestamptz default current_timestamp not null,
    update_by           varchar                               not null,
    update_time         timestamptz                           not null
);

-- 添加会员积分成长规则表注释
comment on table ums_member_rule_setting is '会员积分成长规则表';

-- 添加会员积分成长规则表列注释
comment on column ums_member_rule_setting.id is '';
comment on column ums_member_rule_setting.consume_per_point is '每消费多少元获取1个点';
comment on column ums_member_rule_setting.low_order_amount is '最低获取点数的订单金额';
comment on column ums_member_rule_setting.max_point_per_order is '每笔订单最高获取点数';
comment on column ums_member_rule_setting.rule_type is '类型：0->积分规则；1->成长值规则';
comment on column ums_member_rule_setting.status is '状态：0->禁用；1->启用';
comment on column ums_member_rule_setting.create_by is '创建人ID';
comment on column ums_member_rule_setting.create_time is '创建时间';
comment on column ums_member_rule_setting.update_by is '更新人ID';
comment on column ums_member_rule_setting.update_time is '更新时间';

-- 会员签到记录表建语句
drop table if exists ums_member_sign_log;
create table ums_member_sign_log
(
    id            bigserial primary KEY,
    member_id     bigint                                not null,
    sign_date     date                                  not null,
    continue_days integer                               not null,
    points        integer                               not null,
    create_time   timestamptz default current_timestamp not null
);

-- 添加会员签到记录表注释
comment on table ums_member_sign_log is '会员签到记录表';

-- 添加会员签到记录表列注释
comment on column ums_member_sign_log.id is '';
comment on column ums_member_sign_log.member_id is '会员ID';
comment on column ums_member_sign_log.sign_date is '签到日期';
comment on column ums_member_sign_log.continue_days is '连续签到天数';
comment on column ums_member_sign_log.points is '获得积分';
comment on column ums_member_sign_log.create_time is '';

-- 会员统计信息建语句
drop table if exists ums_member_statistics_info;
create table ums_member_statistics_info
(
    id                    bigserial primary KEY,
    member_id             bigint      not null,
    consume_amount        bigint      not null,
    order_count           integer     not null,
    coupon_count          integer     not null,
    comment_count         integer     not null,
    return_order_count    integer     not null,
    login_count           integer     not null,
    attend_count          integer     not null,
    fans_count            integer     not null,
    collect_product_count integer     not null,
    collect_subject_count integer     not null,
    collect_topic_count   integer     not null,
    collect_comment_count integer     not null,
    invite_friend_count   integer     not null,
    recent_order_time     timestamptz not null
);

-- 添加会员统计信息注释
comment on table ums_member_statistics_info is '会员统计信息';

-- 添加会员统计信息列注释
comment on column ums_member_statistics_info.id is '';
comment on column ums_member_statistics_info.member_id is '会员id';
comment on column ums_member_statistics_info.consume_amount is '累计消费金额';
comment on column ums_member_statistics_info.order_count is '订单数量';
comment on column ums_member_statistics_info.coupon_count is '优惠券数量';
comment on column ums_member_statistics_info.comment_count is '评价数';
comment on column ums_member_statistics_info.return_order_count is '退货数量';
comment on column ums_member_statistics_info.login_count is '登录次数';
comment on column ums_member_statistics_info.attend_count is '关注数量';
comment on column ums_member_statistics_info.fans_count is '粉丝数量';
comment on column ums_member_statistics_info.collect_product_count is '收藏的商品数量';
comment on column ums_member_statistics_info.collect_subject_count is '收藏的专题活动数量';
comment on column ums_member_statistics_info.collect_topic_count is '收藏的评论数量';
comment on column ums_member_statistics_info.collect_comment_count is '收藏的专题活动数量';
comment on column ums_member_statistics_info.invite_friend_count is '邀请好友数';
comment on column ums_member_statistics_info.recent_order_time is '最后一次下订单时间';

-- 用户标签表建语句
drop table if exists ums_member_tag;
create table ums_member_tag
(
    id                  bigserial primary KEY,
    tag_name            varchar                               not null,
    description         varchar                               not null,
    finish_order_count  integer                               not null,
    finish_order_amount numeric                               not null,
    status              smallint    default 1                 not null,
    create_by           varchar     default ''                not null,
    create_time         timestamptz default current_timestamp not null,
    update_by           varchar                               not null,
    update_time         timestamptz                           not null,
    is_deleted          smallint    default 1                 not null
);

-- 添加用户标签表注释
comment on table ums_member_tag is '用户标签表';

-- 添加用户标签表列注释
comment on column ums_member_tag.id is '主键ID';
comment on column ums_member_tag.tag_name is '标签名称';
comment on column ums_member_tag.description is '标签描述';
comment on column ums_member_tag.finish_order_count is '自动打标签完成订单数量';
comment on column ums_member_tag.finish_order_amount is '自动打标签完成订单金额';
comment on column ums_member_tag.status is '状态：0-禁用，1-启用';
comment on column ums_member_tag.create_by is '创建人ID';
comment on column ums_member_tag.create_time is '创建时间';
comment on column ums_member_tag.update_by is '更新人ID';
comment on column ums_member_tag.update_time is '更新时间';
comment on column ums_member_tag.is_deleted is '是否删除';

-- 会员标签关联表建语句
drop table if exists ums_member_tag_relation;
create table ums_member_tag_relation
(
    id          bigserial primary KEY,
    member_id   bigint                                not null,
    tag_id      bigint                                not null,
    create_time timestamptz default current_timestamp not null
);

-- 添加会员标签关联表注释
comment on table ums_member_tag_relation is '会员标签关联表';

-- 添加会员标签关联表列注释
comment on column ums_member_tag_relation.id is '主键ID';
comment on column ums_member_tag_relation.member_id is '会员ID';
comment on column ums_member_tag_relation.tag_id is '标签ID';
comment on column ums_member_tag_relation.create_time is '创建时间';

-- 会员任务表建语句
drop table if exists ums_member_task;
create table ums_member_task
(
    id             bigserial primary KEY,
    task_name      varchar                               not null,
    task_desc      varchar                               not null,
    task_growth    integer                               not null,
    task_integral  integer                               not null,
    task_type      smallint    default 1                 not null,
    complete_count integer                               not null,
    reward_type    smallint    default 1                 not null,
    reward_params  varchar                               not null,
    start_time     timestamptz                           not null,
    end_time       timestamptz                           not null,
    status         smallint    default 1                 not null,
    sort           integer                               not null,
    create_by      varchar     default ''                not null,
    create_time    timestamptz default current_timestamp not null,
    update_by      bigint                                not null,
    update_time    timestamptz                           not null,
    is_deleted     smallint    default 1                 not null
);

-- 添加会员任务表注释
comment on table ums_member_task is '会员任务表';

-- 添加会员任务表列注释
comment on column ums_member_task.id is '主键ID';
comment on column ums_member_task.task_name is '任务名称';
comment on column ums_member_task.task_desc is '任务描述';
comment on column ums_member_task.task_growth is '赠送成长值';
comment on column ums_member_task.task_integral is '赠送积分';
comment on column ums_member_task.task_type is '任务类型：0-新手任务，1-日常任务，2-周常任务，3-月常任务';
comment on column ums_member_task.complete_count is '需要完成次数';
comment on column ums_member_task.reward_type is '奖励类型：0-积分成长值，1-优惠券，2-抽奖次数';
comment on column ums_member_task.reward_params is '奖励参数JSON';
comment on column ums_member_task.start_time is '任务开始时间';
comment on column ums_member_task.end_time is '任务结束时间';
comment on column ums_member_task.status is '状态：0-禁用，1-启用';
comment on column ums_member_task.sort is '排序';
comment on column ums_member_task.create_by is '创建人ID';
comment on column ums_member_task.create_time is '创建时间';
comment on column ums_member_task.update_by is '更新人ID';
comment on column ums_member_task.update_time is '更新时间';
comment on column ums_member_task.is_deleted is '是否删除';

-- 会员任务关联表建语句
drop table if exists ums_member_task_relation;
create table ums_member_task_relation
(
    id          bigserial primary KEY,
    member_id   bigint                                not null,
    task_id     bigint                                not null,
    create_time timestamptz default current_timestamp not null
);

-- 添加会员任务关联表注释
comment on table ums_member_task_relation is '会员任务关联表';

-- 添加会员任务关联表列注释
comment on column ums_member_task_relation.id is '主键ID';
comment on column ums_member_task_relation.member_id is '会员ID';
comment on column ums_member_task_relation.task_id is '任务ID';
comment on column ums_member_task_relation.create_time is '创建时间';
