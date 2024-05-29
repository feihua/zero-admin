create table sys_user_post
(
    id      bigint auto_increment comment '编号'
        primary key,
    user_id bigint not null comment '用户Id',
    post_id bigint not null comment '岗位Id'
) comment '用户岗位关联';

