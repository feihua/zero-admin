drop table if exists sys_user;
create table sys_user
(
    id              bigserial primary key,
    mobile          varchar(11)                            not null,
    user_name       varchar(50)                            not null,
    nick_name       varchar(30)                            not null,
    user_type       varchar(2)   default '00'              not null,
    avatar          varchar(100) default ''                not null,
    email           varchar(50)                            not null,
    password        varchar(64)                            not null,
    status          integer     default 1                 not null,
    dept_id         bigint       default 1                 not null,
    last_login_info      jsonb default '{}'                not null,
    pwd_update_date timestamptz,
    del_flag        integer     default 1                 not null,
    remark          varchar(255) default ''                not null,
    create_by       varchar(50)  default ''                not null,
    create_time     timestamptz  default current_timestamp not null,
    update_by       varchar(50)  default ''                not null,
    update_time     timestamptz,
    constraint ak_phone unique (mobile)
);

-- 添加表注释
comment on table sys_user is '用户信息';

-- 添加列注释
comment on column sys_user.id is '主键';
comment on column sys_user.mobile is '手机号码';
comment on column sys_user.user_name is '用户账号';
comment on column sys_user.nick_name is '用户昵称';
comment on column sys_user.user_type is '用户类型（00系统用户）';
comment on column sys_user.avatar is '头像路径';
comment on column sys_user.email is '用户邮箱';
comment on column sys_user.password is '密码';
comment on column sys_user.status is '状态(1:正常，0:禁用)';
comment on column sys_user.dept_id is '部门id';
comment on column sys_user.last_login_info is '最后登录信息';
comment on column sys_user.pwd_update_date is '密码最后更新时间';
comment on column sys_user.del_flag is '删除标志（0:删除,1:存在）';
comment on column sys_user.remark is '备注';
comment on column sys_user.create_by is '创建者';
comment on column sys_user.create_time is '创建时间';
comment on column sys_user.update_by is '更新者';
comment on column sys_user.update_time is '更新时间';


insert into sys_user (id, mobile, user_name, nick_name, avatar, email, password, status, remark) values (1, '18613030111', 'admin','admin', 'https://gw.alipayobjects.com/zos/antfincdn/xaosxunzyf/biazfanxmamnroxxvxka.png', 'xx@qq.com','123456', 1,  '超级管理员');
insert into sys_user (id, mobile, user_name, nick_name, avatar, email, password, status, remark) values (2, '18613030222', 'test', 'test', 'https://gw.alipayobjects.com/zos/antfincdn/xaosxunzyf/biazfanxmamnroxxvxka.png','123@qq.com','123456', 1, '演示权限');

drop table if exists sys_user_role;
create table sys_user_role
(
    id      bigserial primary key,
    user_id bigint default 0 not null,
    role_id bigint           not null
);

-- 添加表注释
comment on table sys_user_role is '用户角色关联表';

-- 添加列注释
comment on column sys_user_role.id is '主键';
comment on column sys_user_role.user_id is '用户id';
comment on column sys_user_role.role_id is '角色id';

insert into sys_user_role (id,user_id, role_id) values (1,1, 1);

drop table if exists sys_role;
create table sys_role
(
    id          bigserial primary key,
    role_name   varchar(50)                            not null,
    role_key    varchar(100) default ''                not null,
    data_scope  integer     default 1                 not null,
    status      integer     default 1                 not null,
    del_flag    integer     default 1                 not null,
    remark      varchar(255) default ''                not null,
    create_by   varchar(50)  default ''                not null,
    create_time timestamptz  default current_timestamp not null,
    update_by   varchar(50)  default ''                not null,
    update_time timestamptz,
    constraint role_name unique (role_name)
);

-- 创建索引
create index name_status_index on sys_role (role_name, status);

-- 添加表注释
comment on table sys_role is '角色信息';

-- 添加列注释
comment on column sys_role.id is '主键';
comment on column sys_role.role_name is '名称';
comment on column sys_role.role_key is '角色权限字符串';
comment on column sys_role.data_scope is '数据范围（1：全部数据权限 2：自定数据权限 3：本部门数据权限 4：本部门及以下数据权限）';
comment on column sys_role.status is '状态(1:正常，0:禁用)';
comment on column sys_role.del_flag is '删除标志（0:删除,1:存在）';
comment on column sys_role.remark is '备注';
comment on column sys_role.create_by is '创建者';
comment on column sys_role.create_time is '创建时间';
comment on column sys_role.update_by is '更新者';
comment on column sys_role.update_time is '更新时间';

insert into sys_role (id, role_name, role_key, status, remark) values (1, '超级管理员', 'admin',1, '全部权限');
insert into sys_role (id, role_name, role_key, status, remark) values (2, '演示角色', 'query',1,  '仅有查看功能');
insert into sys_role (id, role_name, role_key, status, remark) values (3, '121', 'dev',0, '121211');

drop table if exists sys_role_menu;
create table sys_role_menu
(
    id      bigserial primary key,
    role_id bigint not null,
    menu_id bigint not null
);

-- 添加表注释
comment on table sys_role_menu is '角色菜单关联表';

-- 添加列注释
comment on column sys_role_menu.id is '主键';
comment on column sys_role_menu.role_id is '角色id';
comment on column sys_role_menu.menu_id is '菜单id';

drop table if exists sys_menu;
create table sys_menu
(
    id            bigserial primary key,
    menu_name     varchar(50)                            not null,
    ancestors     varchar(50)  default ''                not null,
    menu_type     integer     default 1                 not null,
    menu_url      varchar(255) default ''                not null,
    menu_icon     varchar(255) default ''                not null,
    menu_sort     integer      default 1                 not null,
    parent_id     bigint       default 0                 not null,
    api_url       varchar(255) default ''                not null,
    visible       integer     default 1                 not null,
    status        integer     default 1                 not null,
    del_flag      integer     default 1                 not null,
    remark        varchar(255) default ''                not null,
    vue_path      varchar(64)  default ''                not null,
    vue_component varchar(64)  default ''                not null,
    vue_icon      varchar(64)  default ''                not null,
    vue_redirect  varchar(64)  default ''                not null,
    angular_icon  varchar(64)  default ''                not null,
    react_icon    varchar(64)  default ''                not null,
    create_by     varchar(50)  default ''                not null,
    create_time   timestamptz  default current_timestamp not null,
    update_by     varchar(50)  default ''                not null,
    update_time   timestamptz,
    constraint menu_name unique (menu_name)
);

-- 添加表注释
comment on table sys_menu is '菜单信息';

-- 添加列注释
comment on column sys_menu.id is '主键';
comment on column sys_menu.menu_name is '菜单名称';
comment on column sys_menu.ancestors is '祖级列表';
comment on column sys_menu.menu_type is '菜单类型(1:目录,2:菜单,3:按钮)';
comment on column sys_menu.visible is '显示状态（0:隐藏,显示:1）';
comment on column sys_menu.status is '菜单状态(1:正常，0:禁用)';
comment on column sys_menu.menu_sort is '排序';
comment on column sys_menu.del_flag is '删除标志（0:删除,1:存在）';
comment on column sys_menu.parent_id is '父id';
comment on column sys_menu.menu_url is '路由路径';
comment on column sys_menu.api_url is '接口url';
comment on column sys_menu.menu_icon is '菜单图标';
comment on column sys_menu.remark is '备注';
comment on column sys_menu.vue_path is 'vue的path';
comment on column sys_menu.vue_component is 'vue的页面';
comment on column sys_menu.vue_icon is 'vue的图标';
comment on column sys_menu.vue_redirect is 'vue的路由重定向';
comment on column sys_menu.angular_icon is 'angular的图标';
comment on column sys_menu.react_icon is 'antd react的图标';
comment on column sys_menu.create_by is '创建者';
comment on column sys_menu.create_time is '创建时间';
comment on column sys_menu.update_by is '更新者';
comment on column sys_menu.update_time is '更新时间';


insert into sys_menu (id, menu_name, menu_type, status, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) values (1, '目录', 1, 1, 1, 0, '/main', '', '', '目录');
insert into sys_menu (id, menu_name, menu_type, status, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) values (2, '欢迎', 1, 1, 2, 1, '/welcome', '', 'SmileOutlined', '欢迎');
insert into sys_menu (id, menu_name, menu_type, status, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) values (3, '系统管理', 1, 1, 3, 1, '/system', '', 'SettingOutlined', '系统管理');
insert into sys_menu (id, menu_name, menu_type, status, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) values (4, '日志管理', 1, 1, 4, 1, '/log', '', 'DeleteOutlined', '日志管理');
insert into sys_menu (id, menu_name, menu_type, status, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) values (5, '会员管理', 1, 1, 5, 1, '/ums', '', 'FrownOutlined', '会员管理');
insert into sys_menu (id, menu_name, menu_type, status, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) values (6, '商品管理', 1, 1, 6, 1, '/pms', '', 'GiftOutlined', '商品管理');
insert into sys_menu (id, menu_name, menu_type, status, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) values (7, '订单管理', 1, 1, 7, 1, '/oms', '', 'DollarCircleOutlined', '订单管理');
insert into sys_menu (id, menu_name, menu_type, status, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) values (8, '营销管理', 1, 1, 8, 1, '/sms', '', 'AlertOutlined', '营销管理');
insert into sys_menu (id, menu_name, menu_type, status, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) values (9, '内容管理', 1, 1, 9, 1, '/cms', '', 'SettingOutlined', '内容管理');

-- 配置用户信息权限
-- 添加列注释
comment on column sys_role_dept.id is '主键';
comment on column sys_role_dept.role_id is '角色id';
comment on column sys_role_dept.dept_id is '部门id';

drop table if exists sys_user_post;
create table sys_user_post
(
    id      bigserial primary key,
    user_id bigint not null,
    post_id bigint not null
);

-- 添加表注释
comment on table sys_user_post is '用户与岗位关联表';

-- 添加列注释
comment on column sys_user_post.id is '主键';
comment on column sys_user_post.user_id is '用户id';
comment on column sys_user_post.post_id is '岗位id';

insert into sys_user_post(user_id, post_id) values ('1', '1');
insert into sys_user_post(user_id, post_id) values ('2', '2');

drop table if exists sys_post;
create table sys_post
(
    id          bigserial primary key,
    post_code   varchar(64)             not null,
    post_name   varchar(50)             not null,
    sort        integer      default 0  not null,
    status      integer     default 1  not null,
    del_flag    integer     default 1  not null,
    remark      varchar(500) default '' not null,
    create_by   varchar(50)  default '' not null,
    create_time timestamptz  default current_timestamp  not null,
    update_by   varchar(50)  default ''                not null,
    update_time timestamptz
);

-- 添加表注释
comment on table sys_post is '岗位信息表';

-- 添加列注释
comment on column sys_post.id is '岗位id';
comment on column sys_post.post_code is '岗位编码';
comment on column sys_post.post_name is '岗位名称';
comment on column sys_post.sort is '显示顺序';
comment on column sys_post.status is '岗位状态（0：停用，1:正常）';
comment on column sys_post.del_flag is '删除标志（0:删除,1:存在）';
comment on column sys_post.remark is '备注';
comment on column sys_post.create_by is '创建者';
comment on column sys_post.create_time is '创建时间';
comment on column sys_post.update_by is '更新者';
comment on column sys_post.update_time is '更新时间';

insert into sys_post (id, post_code, post_name, sort, status, remark) values (1, 'ceo', '董事长', 1, 1, '');
insert into sys_post (id, post_code, post_name, sort, status, remark) values (2, 'se', '项目经理', 2, 1, '');
insert into sys_post (id, post_code, post_name, sort, status, remark) values (3, 'hr', '人力资源', 3, 1, '');
insert into sys_post (id, post_code, post_name, sort, status, remark) values (4, 'user', '普通员工', 1, 1, '');

drop table if exists sys_dept;
create table sys_dept
(
    id          bigserial primary key,
    parent_id   bigint      default 0                 not null,
    ancestors   varchar(50) default ''                not null,
    dept_name   varchar(30) default ''                not null,
    sort        integer     default 0                 not null,
    leader      varchar(20) default ''                not null,
    phone       varchar(11) default ''                not null,
    email       varchar(50) default ''                not null,
    status      integer    default 1                 not null,
    del_flag    integer    default 1                 not null,
    remark      varchar(255) default ''                not null,
    create_by   varchar(50) default ''                not null,
    create_time timestamptz default current_timestamp not null,
    update_by   varchar(50)  default ''                not null,
    update_time timestamptz
);

-- 添加表注释
comment on table sys_dept is '部门表';

-- 添加列注释
comment on column sys_dept.id is '部门id';
comment on column sys_dept.parent_id is '父部门id';
comment on column sys_dept.ancestors is '祖级列表';
comment on column sys_dept.dept_name is '部门名称';
comment on column sys_dept.sort is '显示顺序';
comment on column sys_dept.leader is '负责人';
comment on column sys_dept.phone is '联系电话';
comment on column sys_dept.email is '邮箱';
comment on column sys_dept.status is '状态0:停用,1:正常';
comment on column sys_dept.del_flag is '删除标志（0:删除,1:存在）';
comment on column sys_dept.remark is '备注';
comment on column sys_dept.create_by is '创建者';
comment on column sys_dept.create_time is '创建时间';
comment on column sys_dept.update_by is '更新者';
comment on column sys_dept.update_time is '更新时间';

insert into sys_dept (id, parent_id, ancestors, dept_name, sort, leader, phone, email, status) values (1,0, '0', '测试科技', 1, 'admin', '18613030352', '1002219331@qq.com', 1);
insert into sys_dept (id, parent_id, ancestors, dept_name, sort, leader, phone, email, status) values (2,1, '0,1', '深圳总公司', 1, '1', '1', 'xx@qq.com', 1);
insert into sys_dept (id, parent_id, ancestors, dept_name, sort, leader, phone, email, status) values (3,1, '0,1', '长沙分公司', 2, '1', '1', 'xx@qq.com', 1);
insert into sys_dept (id, parent_id, ancestors, dept_name, sort, leader, phone, email, status) values (4,2, '0,1,2', '研发部门', 1, '1', '1', 'xx@qq.com', 1);
insert into sys_dept (id, parent_id, ancestors, dept_name, sort, leader, phone, email, status) values (5,2, '0,1,2', '市场部门', 2, '1', '1', 'xx@qq.com', 1);
insert into sys_dept (id, parent_id, ancestors, dept_name, sort, leader, phone, email, status) values (6,2, '0,1,2', '测试部门', 3, '1', '1', 'xx@qq.com', 1);
insert into sys_dept (id, parent_id, ancestors, dept_name, sort, leader, phone, email, status) values (7,2, '0,1,2', '财务部门', 4, '1', '1', 'xx@qq.com', 1);
insert into sys_dept (id, parent_id, ancestors, dept_name, sort, leader, phone, email, status) values (8,2, '0,1,2', '运维部门', 5, '1', '1', 'xx@qq.com', 1);
insert into sys_dept (id, parent_id, ancestors, dept_name, sort, leader, phone, email, status) values (9,3, '0,1,3', '市场部门1', 6, '1', '1', 'xx@qq.com', 1);
insert into sys_dept (id, parent_id, ancestors, dept_name, sort, leader, phone, email, status) values (10,3, '0,1,3', '财务部门1', 1, '1', '1', 'xx@qq.com', 1);

-- select string_to_array(ancestors, ',') from sys_dept;
-- select id from sys_dept where 2 = any(string_to_array(ancestors, ',')::integer[])

drop table if exists sys_dict_data;
create table sys_dict_data
(
    id          bigserial primary key,
    dict_label  varchar(100) default ''                not null,
    dict_value  varchar(100) default ''                not null,
    dict_type   varchar(100) default ''                not null,
    css_class   varchar(100) default ''                not null,
    list_class  varchar(100) default ''                not null,
    is_default  varchar(1)   default 'n'               not null,
    sort        integer      default 0                 not null,
    status      integer     default 1                 not null,
    del_flag    integer     default 1                 not null,
    remark      varchar(500) default ''                not null,
    create_by   varchar(50)  default ''                not null,
    create_time timestamptz  default current_timestamp not null,
    update_by   varchar(50)  default ''                not null,
    update_time timestamptz
);

-- 添加表注释
comment on table sys_dict_data is '字典数据表';

-- 添加列注释
comment on column sys_dict_data.id is '字典编码';
comment on column sys_dict_data.dict_label is '字典标签';
comment on column sys_dict_data.dict_value is '字典键值';
comment on column sys_dict_data.dict_type is '字典类型';
comment on column sys_dict_data.css_class is '样式属性（其他样式扩展）';
comment on column sys_dict_data.list_class is '表格回显样式';
comment on column sys_dict_data.is_default is '是否默认（y是 n否）';
comment on column sys_dict_data.sort is '字典排序';
comment on column sys_dict_data.status is '状态（0：停用，1:正常）';
comment on column sys_dict_data.del_flag is '删除标志（0:删除,1:存在）';
comment on column sys_dict_data.remark is '备注';
comment on column sys_dict_data.create_by is '创建者';
comment on column sys_dict_data.create_time is '创建时间';
comment on column sys_dict_data.update_by is '更新者';
comment on column sys_dict_data.update_time is '更新时间';


insert into sys_dict_data (id, dict_label, dict_value, dict_type, css_class, list_class, is_default, sort, status, remark) values (1, '男', '0', 'sys_user_sex', '1', '1', 'n', 1, 1, '性别男');
insert into sys_dict_data (id, dict_label, dict_value, dict_type, css_class, list_class, is_default, sort, status, remark) values (2, '女', '1', 'sys_user_sex', '1', '1', 'n', 2, 1, '性别女');
insert into sys_dict_data (id, dict_label, dict_value, dict_type, css_class, list_class, is_default, sort, status, remark) values (3, '未知', '2', 'sys_user_sex', '1', '1', 'n', 3, 1, '性别未知');
insert into sys_dict_data (id, dict_label, dict_value, dict_type, css_class, list_class, is_default, sort, status, remark) values (4, '通知', '1', 'sys_notice_type', '1', '1', 'n', 1, 1, '通知');
insert into sys_dict_data (id, dict_label, dict_value, dict_type, css_class, list_class, is_default, sort, status, remark) values (5, '公告', '2', 'sys_notice_type', '1', '1', 'n', 2, 1, '公告');

drop table if exists sys_dict_type;
create table sys_dict_type
(
    id          bigserial primary key,
    dict_name   varchar(100) default ''                not null,
    dict_type   varchar(100) default ''                not null,
    status      integer     default 1                 not null,
    del_flag    integer     default 1                 not null,
    remark      varchar(500) default ''                not null,
    create_by   varchar(50)  default ''                not null,
    create_time timestamptz  default current_timestamp not null,
    update_by   varchar(50)  default ''                not null,
    update_time timestamptz,
    unique (dict_type)
);

-- 添加表注释
comment on table sys_dict_type is '字典类型表';

-- 添加列注释
comment on column sys_dict_type.id is '字典主键';
comment on column sys_dict_type.dict_name is '字典名称';
comment on column sys_dict_type.dict_type is '字典类型';
comment on column sys_dict_type.status is '状态（0：停用，1:正常）';
comment on column sys_dict_type.del_flag is '删除标志（0:删除,1:存在）';
comment on column sys_dict_type.remark is '备注';
comment on column sys_dict_type.create_by is '创建者';
comment on column sys_dict_type.create_time is '创建时间';
comment on column sys_dict_type.update_by is '更新者';
comment on column sys_dict_type.update_time is '更新时间';

insert into sys_dict_type (id, dict_name, dict_type, status, remark) values (1,'用户性别', 'sys_user_sex', 1, '用户性别列表');
insert into sys_dict_type (id, dict_name, dict_type, status, remark) values (2,'通知类型', 'sys_notice_type', 1, '通知类型列表');

drop table if exists sys_login_log;
create table sys_login_log
(
    id         bigserial primary key,
    login_name varchar(50)  default ''                not null,
    ip_addr    varchar(128) default ''                not null,
    extra      text         default ''                not null,
    status     integer     default 0                 not null,
    msg        varchar(255) default ''                not null,
    login_time timestamptz  default current_timestamp not null
);

-- 添加表注释
comment on table sys_login_log is '系统访问记录';

-- 添加列注释
comment on column sys_login_log.id is '访问id';
comment on column sys_login_log.login_name is '登录账号';
comment on column sys_login_log.ip_addr is '登录ip地址';
comment on column sys_login_log.extra is '其他信息（可选）';
comment on column sys_login_log.status is '登录状态(0:失败,1:成功)';
comment on column sys_login_log.msg is '提示消息';
comment on column sys_login_log.login_time is '访问时间';

drop table if exists sys_notice;
create table sys_notice
(
    id             bigserial primary key,
    notice_title   varchar(50)                            not null,
    notice_type    integer     default 1                 not null,
    notice_content varchar(255) default ''                not null,
    status         integer     default 1                 not null,
    remark         varchar(255) default ''                not null,
    del_flag       integer     default 1                 not null,
    create_by      varchar(50)  default ''                not null,
    create_time    timestamptz  default current_timestamp not null,
    update_by      varchar(50)  default ''                not null,
    update_time    timestamptz
);

-- 添加表注释
comment on table sys_notice is '通知公告表';

-- 添加列注释
comment on column sys_notice.id is '公告id';
comment on column sys_notice.notice_title is '公告标题';
comment on column sys_notice.notice_type is '公告类型（1:通知,2:公告）';
comment on column sys_notice.notice_content is '公告内容';
comment on column sys_notice.status is '公告状态（0:关闭,1:正常 ）';
comment on column sys_notice.del_flag is '删除标志（0:删除,1:存在）';
comment on column sys_notice.remark is '备注';
comment on column sys_notice.create_by is '创建者';
comment on column sys_notice.create_time is '创建时间';
comment on column sys_notice.update_by is '更新者';
comment on column sys_notice.update_time is '更新时间';


insert into sys_notice (id, notice_title, notice_type, notice_content, status) values (1,'测试通知1', 1, '这是一条测试通知内容', 1);
insert into sys_notice (id, notice_title, notice_type, notice_content, status) values (2,'测试公告2', 2, '这是一条测试公告内容', 1);


drop table if exists sys_operate_log;
create table sys_operate_log
(
    id            bigserial primary key,
    operate_name  varchar(50) default ''                not null,
    operate_ip    varchar(50) default ''                not null,
    operate_url   text        default ''                not null,
    operate_param text        default ''                not null,
    json_result   text        default ''                not null,
    extra         text        default ''                not null,
    status        integer    default 0                 not null,
    cost_time     bigint      default 0                 not null,
    operate_time  timestamptz default current_timestamp not null

);

-- 添加表注释
comment on table sys_operate_log is '操作日志记录';

-- 添加列注释
comment on column sys_operate_log.id is '日志主键';
comment on column sys_operate_log.operate_name is '操作人员';
comment on column sys_operate_log.operate_ip is '主机地址';
comment on column sys_operate_log.operate_url is '请求url';
comment on column sys_operate_log.operate_param is '请求参数';
comment on column sys_operate_log.json_result is '返回参数';
comment on column sys_operate_log.extra is '其他信息（可选）';
comment on column sys_operate_log.status is '操作状态(0:异常,正常)';
comment on column sys_operate_log.cost_time is '消耗时间';
comment on column sys_operate_log.operate_time is '操作时间';

-- 帮助表建语句
drop table if exists cms_help;
create table cms_help
(
    id          bigserial primary key,
    category_id bigint                                not null,
    icon        varchar                               not null,
    title       varchar                               not null,
    show_status integer     default 1                 not null,
    read_count  integer                               not null,
    content     text                                  not null,
    create_by   varchar     default ''                not null,
    create_time timestamptz default current_timestamp not null,
    update_by   varchar     default ''                not null,
    update_time timestamptz                           null
);

-- 添加帮助表注释
comment on table cms_help is '帮助表';

-- 添加帮助表列注释
comment on column cms_help.id is '主键id';
comment on column cms_help.category_id is '分类id';
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
    id          bigserial primary key,
    name        varchar                               not null,
    icon        varchar                               not null,
    help_count  integer                               not null,
    show_status integer     default 1                 not null,
    sort        integer                               not null,
    create_by   varchar     default ''                not null,
    create_time timestamptz default current_timestamp not null,
    update_by   varchar     default ''                not null,
    update_time timestamptz                           null
);

-- 添加帮助分类表注释
comment on table cms_help_category is '帮助分类表';

-- 添加帮助分类表列注释
comment on column cms_help_category.id is '主键id';
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
    id                 bigserial primary key,
    report_type        integer     default 1                 not null,
    report_member_name varchar                               not null,
    report_object      varchar                               not null,
    report_status      integer     default 1                 not null,
    handle_status      integer     default 1                 not null,
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
    id          bigserial primary key,
    name        varchar                               not null,
    sub_title   varchar                               not null,
    pic         varchar                               not null,
    sort        integer                               not null,
    show_status integer     default 1                 not null,
    create_by   varchar     default ''                not null,
    create_time timestamptz default current_timestamp not null,
    update_by   varchar     default ''                not null,
    update_time timestamptz                           null
);

-- 添加优选专区注释
comment on table cms_preferred_area is '优选专区';

-- 添加优选专区列注释
comment on column cms_preferred_area.id is '主键id';
comment on column cms_preferred_area.name is '专区名称';
comment on column cms_preferred_area.sub_title is '子标题';
comment on column cms_preferred_area.pic is '展示图片';
comment on column cms_preferred_area.sort is '排序';
comment on column cms_preferred_area.show_status is '显示状态：0->不显示；1->显示';
comment on column cms_preferred_area.create_by is '创建者';
comment on column cms_preferred_area.create_time is '创建时间';
comment on column cms_preferred_area.update_by is '更新者';
comment on column cms_preferred_area.update_time is '更新时间';

insert into cms_preferred_area (id, name, sub_title, pic, sort, show_status) values (1, '专区1', '专区1副标题', '', 1, 1);
insert into cms_preferred_area (id, name, sub_title, pic, sort, show_status) values (2, '专区2', '专区2副标题', '', 2, 1);
insert into cms_preferred_area (id, name, sub_title, pic, sort, show_status) values (3, '专区3', '专区3副标题', '', 3, 1);
insert into cms_preferred_area (id, name, sub_title, pic, sort, show_status) values (4, '专区4', '专区4副标题', '', 4, 1);
insert into cms_preferred_area (id, name, sub_title, pic, sort, show_status) values (5, '专区5', '专区5副标题', '', 4, 1);

-- 优选专区和产品关系表建语句
drop table if exists cms_preferred_area_product_relation;
create table cms_preferred_area_product_relation
(
    id                bigserial primary key,
    preferred_area_id bigint not null,
    product_id        bigint not null
);

-- 添加优选专区和产品关系表注释
comment on table cms_preferred_area_product_relation is '优选专区和产品关系表';

-- 添加优选专区和产品关系表列注释
comment on column cms_preferred_area_product_relation.id is '主键id';
comment on column cms_preferred_area_product_relation.preferred_area_id is '优选专区id';
comment on column cms_preferred_area_product_relation.product_id is '产品id';

-- 专题表建语句
drop table if exists cms_subject;
create table cms_subject
(
    id               bigserial primary key,
    category_id      bigint                                not null,
    title            varchar                               not null,
    pic              varchar                               not null,
    product_count    integer                               not null,
    recommend_status integer     default 1                 not null,
    collect_count    integer                               not null,
    read_count       integer                               not null,
    comment_count    integer                               not null,
    album_pics       varchar                               not null,
    description      varchar                               not null,
    show_status      integer     default 1                 not null,
    content          text                                  not null,
    forward_count    integer                               not null,
    category_name    varchar                               not null,
    sort             integer     default 1                 not null,
    create_by        varchar     default ''                not null,
    create_time      timestamptz default current_timestamp not null,
    update_by        varchar     default ''                not null,
    update_time      timestamptz                           null
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

insert into cms_subject (id, category_id, title, pic, product_count, collect_count,read_count, comment_count, album_pics, description, show_status, content, forward_count,category_name, sort) values (1, 1, '电磁炉', '1', 1, 1, 1, 1, '1', ' ', 1, '1', 1, '厨房专题',1);
insert into cms_subject (id, category_id, title, pic, product_count, collect_count,read_count, comment_count, album_pics, description, show_status, content, forward_count,category_name, sort) values (2, 2, '衣服', '1', 1, 1, 1, 1, '1', ' ', 1, '1', 1, '生活专题',2);

-- 专题分类表建语句
drop table if exists cms_subject_category;
create table cms_subject_category
(
    id            bigserial primary key,
    name          varchar                               not null,
    icon          varchar                               not null,
    subject_count integer                               not null,
    show_status   integer     default 1                 not null,
    sort          integer                               not null,
    create_by     varchar     default ''                not null,
    create_time   timestamptz default current_timestamp not null,
    update_by     varchar     default ''                not null,
    update_time   timestamptz                           null
);

-- 添加专题分类表注释
comment on table cms_subject_category is '专题分类表';

-- 添加专题分类表列注释
comment on column cms_subject_category.id is '主键id';
comment on column cms_subject_category.name is '专题分类名称';
comment on column cms_subject_category.icon is '分类图标';
comment on column cms_subject_category.subject_count is '专题数量';
comment on column cms_subject_category.show_status is '显示状态：0->不显示；1->显示';
comment on column cms_subject_category.sort is '排序';
comment on column cms_subject_category.create_by is '创建者';
comment on column cms_subject_category.create_time is '创建时间';
comment on column cms_subject_category.update_by is '更新者';
comment on column cms_subject_category.update_time is '更新时间';

insert into cms_subject_category (id, name, icon, subject_count, show_status, sort) values (1, '厨房专题', ' ', 1, 1, 1);
insert into cms_subject_category (id, name, icon, subject_count, show_status, sort) values (2, '生活专题', ' ', 1, 1, 1);


-- 专题评论表建语句
drop table if exists cms_subject_comment;
create table cms_subject_comment
(
    id               bigserial primary key,
    subject_id       bigint                                not null,
    member_nick_name varchar                               not null,
    member_icon      varchar                               not null,
    content          varchar                               not null,
    create_time      timestamptz default current_timestamp not null,
    show_status      integer     default 1                 not null
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
    id         bigserial primary key,
    subject_id bigint not null,
    product_id bigint not null
);

-- 添加专题商品关系表注释
comment on table cms_subject_product_relation is '专题商品关系表';

-- 添加专题商品关系表列注释
comment on column cms_subject_product_relation.id is '主键id';
comment on column cms_subject_product_relation.subject_id is '专题id';
comment on column cms_subject_product_relation.product_id is '商品id';

-- 话题表建语句
drop table if exists cms_topic;
create table cms_topic
(
    id              bigserial primary key,
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
    update_time     timestamptz                           null
);

-- 添加话题表注释
comment on table cms_topic is '话题表';

-- 添加话题表列注释
comment on column cms_topic.id is '主键id';
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
    id            bigserial primary key,
    name          varchar                               not null,
    icon          varchar                               not null,
    subject_count integer                               not null,
    show_status   integer     default 1                 not null,
    sort          integer                               not null,
    create_by     varchar     default ''                not null,
    create_time   timestamptz default current_timestamp not null,
    update_by     varchar                               not null,
    update_time   timestamptz                           null
);

-- 添加话题分类表注释
comment on table cms_topic_category is '话题分类表';

-- 添加话题分类表列注释
comment on column cms_topic_category.id is '主键id';
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
    id               bigserial primary key,
    member_nick_name varchar                               not null,
    topic_id         bigint                                not null,
    member_icon      varchar                               not null,
    content          varchar                               not null,
    create_time      timestamptz default current_timestamp not null,
    show_status      integer     default 1                 not null
);

-- 添加专题评论表注释
comment on table cms_topic_comment is '专题评论表';

-- 添加专题评论表列注释
comment on column cms_topic_comment.id is '主键id';
comment on column cms_topic_comment.member_nick_name is '评论人员昵称';
comment on column cms_topic_comment.topic_id is '专题id';
comment on column cms_topic_comment.member_icon is '评论人员头像';
comment on column cms_topic_comment.content is '评论内容';
comment on column cms_topic_comment.create_time is '评论时间';
comment on column cms_topic_comment.show_status is '是否显示，0->不显示；1->显示';

-- 运费模版建语句
drop table if exists pms_feight_template;
create table pms_feight_template
(
    id              bigserial primary key,
    name            varchar                               not null,
    charge_type     integer     default 1                 not null,
    first_weight    bigint                                not null,
    first_fee       bigint                                not null,
    continue_weight bigint                                not null,
    continue_fee    bigint                                not null,
    dest            varchar                               not null,
    create_time     timestamptz default current_timestamp not null,
    update_time     timestamptz                           null
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

insert into pms_feight_template (name, charge_type, first_weight, first_fee, continue_weight, continue_fee, dest)
values
    ('北京标准运费模板', 0, 1000, 2000, 500, 500, '北京'),
    ('上海快递专用模板', 1, 1, 1500, 1, 300, '上海'),
    ('广东经济型运费模板', 0, 2000, 3000, 1000, 800, '广东'),
    ('浙江自提免运费模板', 1, 1, 1000, 1, 200, '浙江');


-- 商品会员价格表建语句
drop table if exists pms_member_price;
create table pms_member_price
(
    id                bigserial primary key,
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

insert into pms_member_price (id, product_id, member_level_id, member_price, member_level_name) values (1, 1, 1, 88.00, '黄金会员');
insert into pms_member_price (id, product_id, member_level_id, member_price, member_level_name) values (2, 2, 2, 88.00, '白金会员');
insert into pms_member_price (id, product_id, member_level_id, member_price, member_level_name) values (3, 3, 3, 66.00, '钻石会员');
insert into pms_member_price (id, product_id, member_level_id, member_price, member_level_name) values (4, 4, 1, 66.00, '黄金会员');
insert into pms_member_price (id, product_id, member_level_id, member_price, member_level_name) values (5, 5, 1, 66.00, '黄金会员');
insert into pms_member_price (id, product_id, member_level_id, member_price, member_level_name) values (6, 6, 1, 66.00, '黄金会员');
insert into pms_member_price (id, product_id, member_level_id, member_price, member_level_name) values (7, 7, 1, 66.00, '黄金会员');
insert into pms_member_price (id, product_id, member_level_id, member_price, member_level_name) values (8, 8, 1, 66.00, '黄金会员');
insert into pms_member_price (id, product_id, member_level_id, member_price, member_level_name) values (9, 9, 1, 66.00, '黄金会员');
insert into pms_member_price (id, product_id, member_level_id, member_price, member_level_name) values (10, 10, 1, 66.00, '黄金会员');
insert into pms_member_price (id, product_id, member_level_id, member_price, member_level_name) values (11, 11, 1, 66.00, '黄金会员');
insert into pms_member_price (id, product_id, member_level_id, member_price, member_level_name) values (12, 12, 1, 66.00, '黄金会员');
insert into pms_member_price (id, product_id, member_level_id, member_price, member_level_name) values (13, 13, 1, 66.00, '黄金会员');
insert into pms_member_price (id, product_id, member_level_id, member_price, member_level_name) values (14, 14, 1, 66.00, '黄金会员');
insert into pms_member_price (id, product_id, member_level_id, member_price, member_level_name) values (15, 15, 1, 66.00, '黄金会员');
insert into pms_member_price (id, product_id, member_level_id, member_price, member_level_name) values (16, 16, 1, 66.00, '黄金会员');

-- 商品属性表建语句
drop table if exists pms_product_attribute;
create table pms_product_attribute
(
    id            bigserial primary key,
    group_id      bigint                                not null,
    name          varchar                               not null,
    input_type    integer     default 1                 not null,
    value_type    integer     default 1                 not null,
    input_list    varchar                               not null,
    unit          varchar                               not null,
    is_required   integer     default 1                 not null,
    is_searchable integer     default 1                 not null,
    is_show       integer     default 1                 not null,
    sort          integer                               not null,
    status        integer     default 1                 not null,
    create_by     varchar     default ''                not null,
    create_time   timestamptz default current_timestamp not null,
    update_by     varchar     default ''                not null,
    update_time   timestamptz                           null,
    is_deleted    integer     default 1                 not null
);

-- 添加商品属性表注释
comment on table pms_product_attribute is '商品属性表';

-- 添加商品属性表列注释
comment on column pms_product_attribute.id is '主键id';
comment on column pms_product_attribute.group_id is '属性分组id';
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
comment on column pms_product_attribute.create_by is '创建人id';
comment on column pms_product_attribute.create_time is '创建时间';
comment on column pms_product_attribute.update_by is '更新人id';
comment on column pms_product_attribute.update_time is '更新时间';
comment on column pms_product_attribute.is_deleted is '是否删除(0:否,1:是)';

-- 插入商品属性数据
insert into pms_product_attribute (group_id, name, input_type, value_type, input_list, unit, is_required, is_searchable, is_show, sort)
values
-- 手机属性 (对应group_id为手机基本信息、主体参数等分组的id)
(1, '品牌', 2, 1, 'apple,samsung,huawei,xiaomi', '', 1, 1, 1, 1),
(1, '型号', 1, 1, '', '', 1, 1, 1, 2),
(1, '上市年份', 2, 2, '2025,2026', '年', 1, 1, 1, 3),
(1, '机身颜色', 2, 1, '黑色,白色,金色,银色', '', 1, 1, 1, 4),
-- 主体参数属性值
(2, '处理器', 2, 1, 'a16,a17,a18,a19', 'gb', 1, 1, 1, 5),
(2, '运行内存', 2, 1, '4gb,6gb,8gb,12gb', 'gb', 1, 1, 1, 6),
(2, '机身存储', 2, 1, '64gb,128gb,256gb,512gb', 'gb', 1, 1, 1, 7),
(2, '电池容量', 2, 1, '2400,3600', 'mah', 1, 0, 1, 8),
(2, '机身重量', 2, 1, '167g,187g', 'g', 1, 0, 1, 9),
-- 网络参数属性值
(3, '5g网络', 2, 1, '支持,不支持', '', 1, 1, 1, 10),
(3, '网络制式', 3, 1, 'gsm,wcdma,lte,5g', '', 1, 1, 1, 11),
(3, 'sim卡类型', 2, 1, '单卡,双卡双待', '', 1, 1, 1, 12),
(3, 'wifi', 2, 1, 'wifi 5,wifi 6e', '', 1, 1, 1, 13),
-- 显示参数属性值
(4, '屏幕尺寸', 2, 1, '6.1,6.3,5.5', '英寸', 1, 1, 1, 14),
(4, '屏幕分辨率', 2, 1, '1080x1920,2556x1179', '', 1, 1, 1, 15),
(4, '屏幕刷新率', 2, 1, '60,120', '', 1, 1, 1, 16),
(4, '屏幕类型', 2, 1, 'oled,icd', '', 1, 1, 1, 17),
(4, '屏幕亮度', 2, 1, '2000,3000', '', 1, 1, 1, 18),
-- 摄像功能属性值
(5, '主摄像头', 2, 1, '4800,3600,2400', '', 1, 1, 1, 19),
(5, '前置摄像头', 2, 1, '1200,1600', '', 1, 1, 1, 20),
(5, '摄像头数量', 2, 1, '1,2,3', '', 1, 1, 1, 21),
(5, '视频拍摄', 3, 1, '4k,1080p,780p', '', 1, 1, 1, 22),
(5, '防抖功能', 2, 1, '不防抖,光学防抖', '', 1, 1, 1, 23),

-- 小米手机属性 (对应group_id为手机基本信息、主体参数等分组的id)
(6, '品牌', 2, 1, 'xiaomi,红米', '', 1, 1, 1, 1),
(6, '型号', 1, 1, '', '', 1, 1, 1, 2),
(6, '上市年份', 2, 2, '2025,2026', '年', 1, 1, 1, 3),
(6, '机身颜色', 2, 1, '黑色,白色,金色,银色', '', 1, 1, 1, 4),
-- 小米手机主体参数属性值
(7, '处理器', 2, 1, 'a16,a17,a18,a19', 'gb', 1, 1, 1, 5),
(7, '运行内存', 2, 1, '4gb,6gb,8gb,12gb', 'gb', 1, 1, 1, 6),
(7, '机身存储', 2, 1, '64gb,128gb,256gb,512gb', 'gb', 1, 1, 1, 7),
(7, '电池容量', 2, 1, '2400,3600', 'mah', 1, 0, 1, 8),
(7, '机身重量', 2, 1, '167g,187g', 'g', 1, 0, 1, 9),
-- 小米手机网络参数属性值
(8, '5g网络', 2, 1, '支持,不支持', '', 1, 1, 1, 10),
(8, '网络制式', 3, 1, 'gsm,wcdma,lte,5g', '', 1, 1, 1, 11),
(8, 'sim卡类型', 2, 1, '单卡,双卡双待', '', 1, 1, 1, 12),
(8, 'wifi', 2, 1, 'wifi 5,wifi 6e', '', 1, 1, 1, 13),
-- 小米手机显示参数属性值
(9, '屏幕尺寸', 2, 1, '6.1,6.3,5.5', '英寸', 1, 1, 1, 14),
(9, '屏幕分辨率', 2, 1, '1080x1920,2556x1179', '', 1, 1, 1, 15),
(9, '屏幕刷新率', 2, 1, '60,120', '', 1, 1, 1, 16),
(9, '屏幕类型', 2, 1, 'oled,icd', '', 1, 1, 1, 17),
(9, '屏幕亮度', 2, 1, '2000,3000', '', 1, 1, 1, 18),
-- 小米手机摄像功能属性值
(10, '主摄像头', 2, 1, '4800,3600,2400', '', 1, 1, 1, 19),
(10, '前置摄像头', 2, 1, '1200,1600', '', 1, 1, 1, 20),
(10, '摄像头数量', 2, 1, '1,2,3', '', 1, 1, 1, 21),
(10, '视频拍摄', 3, 1, '4k,1080p,780p', '', 1, 1, 1, 22),
(10, '防抖功能', 2, 1, '不防抖,光学防抖', '', 1, 1, 1, 23),

-- 苹果手机属性 (对应group_id为手机基本信息、主体参数等分组的id)
(11, '品牌', 2, 1, 'iphone17,iphone17 pro, iphone17 pro max', '', 1, 1, 1, 1),
(11, '型号', 1, 1, '', '', 1, 1, 1, 2),
(11, '上市年份', 2, 2, '2025,2026', '年', 1, 1, 1, 3),
(11, '机身颜色', 2, 1, '黑色,白色,金色,银色', '', 1, 1, 1, 4),
-- 苹果手机主体参数属性值
(12, '处理器', 2, 1, 'a16,a17,a18,a19', 'gb', 1, 1, 1, 5),
(12, '运行内存', 2, 1, '4gb,6gb,8gb,12gb', 'gb', 1, 1, 1, 6),
(12, '机身存储', 2, 1, '64gb,128gb,256gb,512gb', 'gb', 1, 1, 1, 7),
(12, '电池容量', 2, 1, '2400,3600', 'mah', 1, 0, 1, 8),
(12, '机身重量', 2, 1, '167g,187g', 'g', 1, 0, 1, 9),
-- 苹果手机网络参数属性值
(13, '5g网络', 2, 1, '支持,不支持', '', 1, 1, 1, 10),
(13, '网络制式', 3, 1, 'gsm,wcdma,lte,5g', '', 1, 1, 1, 11),
(13, 'sim卡类型', 2, 1, '单卡,双卡双待', '', 1, 1, 1, 12),
(13, 'wifi', 2, 1, 'wifi 5,wifi 6e', '', 1, 1, 1, 13),
-- 苹果手机显示参数属性值
(14, '屏幕尺寸', 2, 1, '6.1,6.3,5.5', '英寸', 1, 1, 1, 14),
(14, '屏幕分辨率', 2, 1, '1080x1920,2556x1179', '', 1, 1, 1, 15),
(14, '屏幕刷新率', 2, 1, '60,120', '', 1, 1, 1, 16),
(14, '屏幕类型', 2, 1, 'oled,icd', '', 1, 1, 1, 17),
(14, '屏幕亮度', 2, 1, '2000,3000', '', 1, 1, 1, 18),
-- 苹果手机摄像功能属性值
(15, '主摄像头', 2, 1, '4800,3600,2400', '', 1, 1, 1, 19),
(15, '前置摄像头', 2, 1, '1200,1600', '', 1, 1, 1, 20),
(15, '摄像头数量', 2, 1, '1,2,3', '', 1, 1, 1, 21),
(15, '视频拍摄', 3, 1, '4k,1080p,780p', '', 1, 1, 1, 22),
(15, '防抖功能', 2, 1, '不防抖,光学防抖', '', 1, 1, 1, 23),

-- 华为手机属性 (对应group_id为手机基本信息、主体参数等分组的id)
(16, '品牌', 2, 1, 'nova 15,mate 80', '', 1, 1, 1, 1),
(16, '型号', 1, 1, '', '', 1, 1, 1, 2),
(16, '上市年份', 2, 2, '2025,2026', '年', 1, 1, 1, 3),
(16, '机身颜色', 2, 1, '黑色,白色,金色,银色', '', 1, 1, 1, 4),
-- 华为手机主体参数属性值
(17, '处理器', 2, 1, 'a16,a17,a18,a19', 'gb', 1, 1, 1, 5),
(17, '运行内存', 2, 1, '4gb,6gb,8gb,12gb', 'gb', 1, 1, 1, 6),
(17, '机身存储', 2, 1, '64gb,128gb,256gb,512gb', 'gb', 1, 1, 1, 7),
(17, '电池容量', 2, 1, '2400,3600', 'mah', 1, 0, 1, 8),
(17, '机身重量', 2, 1, '167g,187g', 'g', 1, 0, 1, 9),
-- 华为手机网络参数属性值
(18, '5g网络', 2, 1, '支持,不支持', '', 1, 1, 1, 10),
(18, '网络制式', 3, 1, 'gsm,wcdma,lte,5g', '', 1, 1, 1, 11),
(18, 'sim卡类型', 2, 1, '单卡,双卡双待', '', 1, 1, 1, 12),
(18, 'wifi', 2, 1, 'wifi 5,wifi 6e', '', 1, 1, 1, 13),
-- 华为手机显示参数属性值
(19, '屏幕尺寸', 2, 1, '6.1,6.3,5.5', '英寸', 1, 1, 1, 14),
(19, '屏幕分辨率', 2, 1, '1080x1920,2556x1179', '', 1, 1, 1, 15),
(19, '屏幕刷新率', 2, 1, '60,120', '', 1, 1, 1, 16),
(19, '屏幕类型', 2, 1, 'oled,icd', '', 1, 1, 1, 17),
(19, '屏幕亮度', 2, 1, '2000,3000', '', 1, 1, 1, 18),
-- 华为手机摄像功能属性值
(20, '主摄像头', 2, 1, '4800,3600,2400', '', 1, 1, 1, 19),
(20, '前置摄像头', 2, 1, '1200,1600', '', 1, 1, 1, 20),
(20, '摄像头数量', 2, 1, '1,2,3', '', 1, 1, 1, 21),
(20, '视频拍摄', 3, 1, '4k,1080p,780p', '', 1, 1, 1, 22),
(20, '防抖功能', 2, 1, '不防抖,光学防抖', '', 1, 1, 1, 23),

-- 三星手机属性 (对应group_id为手机基本信息、主体参数等分组的id)
(21, '品牌', 2, 1, 'ultra,galaxy', '', 1, 1, 1, 1),
(21, '型号', 1, 1, '', '', 1, 1, 1, 2),
(21, '上市年份', 2, 2, '2025,2026', '年', 1, 1, 1, 3),
(21, '机身颜色', 2, 1, '黑色,白色,金色,银色', '', 1, 1, 1, 4),
-- 三星手机主体参数属性值
(22, '处理器', 2, 1, 'a16,a17,a18,a19', 'gb', 1, 1, 1, 5),
(22, '运行内存', 2, 1, '4gb,6gb,8gb,12gb', 'gb', 1, 1, 1, 6),
(22, '机身存储', 2, 1, '64gb,128gb,256gb,512gb', 'gb', 1, 1, 1, 7),
(22, '电池容量', 2, 1, '2400,3600', 'mah', 1, 0, 1, 8),
(22, '机身重量', 2, 1, '167g,187g', 'g', 1, 0, 1, 9),
-- 三星手机网络参数属性值
(23, '5g网络', 2, 1, '支持,不支持', '', 1, 1, 1, 10),
(23, '网络制式', 3, 1, 'gsm,wcdma,lte,5g', '', 1, 1, 1, 11),
(23, 'sim卡类型', 2, 1, '单卡,双卡双待', '', 1, 1, 1, 12),
(23, 'wifi', 2, 1, 'wifi 5,wifi 6e', '', 1, 1, 1, 13),
-- 三星手机显示参数属性值
(24, '屏幕尺寸', 2, 1, '6.1,6.3,5.5', '英寸', 1, 1, 1, 14),
(24, '屏幕分辨率', 2, 1, '1080x1920,2556x1179', '', 1, 1, 1, 15),
(24, '屏幕刷新率', 2, 1, '60,120', '', 1, 1, 1, 16),
(24, '屏幕类型', 2, 1, 'oled,icd', '', 1, 1, 1, 17),
(24, '屏幕亮度', 2, 1, '2000,3000', '', 1, 1, 1, 18),
-- 三星手机摄像功能属性值
(25, '主摄像头', 2, 1, '4800,3600,2400', '', 1, 1, 1, 19),
(25, '前置摄像头', 2, 1, '1200,1600', '', 1, 1, 1, 20),
(25, '摄像头数量', 2, 1, '1,2,3', '', 1, 1, 1, 21),
(25, '视频拍摄', 3, 1, '4k,1080p,780p', '', 1, 1, 1, 22),
(25, '防抖功能', 2, 1, '不防抖,光学防抖', '', 1, 1, 1, 23),

-- 荣耀手机属性 (对应group_id为手机基本信息、主体参数等分组的id)
(26, '品牌', 2, 1, '荣耀400,荣耀power2', '', 1, 1, 1, 1),
(26, '型号', 1, 1, '', '', 1, 1, 1, 2),
(26, '上市年份', 2, 2, '2025,2026', '年', 1, 1, 1, 3),
(26, '机身颜色', 2, 1, '黑色,白色,金色,银色', '', 1, 1, 1, 4),
-- 荣耀手机主体参数属性值
(27, '处理器', 2, 1, 'a16,a17,a18,a19', 'gb', 1, 1, 1, 5),
(27, '运行内存', 2, 1, '4gb,6gb,8gb,12gb', 'gb', 1, 1, 1, 6),
(27, '机身存储', 2, 1, '64gb,128gb,256gb,512gb', 'gb', 1, 1, 1, 7),
(27, '电池容量', 2, 1, '2400,3600', 'mah', 1, 0, 1, 8),
(27, '机身重量', 2, 1, '167g,187g', 'g', 1, 0, 1, 9),
-- 荣耀手机网络参数属性值
(28, '5g网络', 2, 1, '支持,不支持', '', 1, 1, 1, 10),
(28, '网络制式', 3, 1, 'gsm,wcdma,lte,5g', '', 1, 1, 1, 11),
(28, 'sim卡类型', 2, 1, '单卡,双卡双待', '', 1, 1, 1, 12),
(28, 'wifi', 2, 1, 'wifi 5,wifi 6e', '', 1, 1, 1, 13),
-- 荣耀手机显示参数属性值
(29, '屏幕尺寸', 2, 1, '6.1,6.3,5.5', '英寸', 1, 1, 1, 14),
(29, '屏幕分辨率', 2, 1, '1080x1920,2556x1179', '', 1, 1, 1, 15),
(29, '屏幕刷新率', 2, 1, '60,120', '', 1, 1, 1, 16),
(29, '屏幕类型', 2, 1, 'oled,icd', '', 1, 1, 1, 17),
(29, '屏幕亮度', 2, 1, '2000,3000', '', 1, 1, 1, 18),
-- 荣耀手机摄像功能属性值
(30, '主摄像头', 2, 1, '4800,3600,2400', '', 1, 1, 1, 19),
(30, '前置摄像头', 2, 1, '1200,1600', '', 1, 1, 1, 20),
(30, '摄像头数量', 2, 1, '1,2,3', '', 1, 1, 1, 21),
(30, '视频拍摄', 3, 1, '4k,1080p,780p', '', 1, 1, 1, 22),
(30, '防抖功能', 2, 1, '不防抖,光学防抖', '', 1, 1, 1, 23),

-- 电脑属性
(31, '品牌', 2, 1, 'lenovo,hp,dell,apple', '', 1, 1, 1, 1),
(32, 'cpu型号', 2, 1, 'intel i5,intel i7,intel i9,amd', '', 1, 1, 1, 2),
(33, '内存容量', 2, 2, '8gb,16gb,32gb,64gb', 'gb', 1, 1, 1, 3),
(34, '显卡型号', 1, 1, '', '', 1, 1, 1, 4),
(35, '硬盘容量', 2, 2, '256gb,512gb,1tb,2tb', 'gb', 1, 1, 1, 5),
-- 苹果电脑属性
(36, '品牌', 2, 1, 'mac mini, macbook air, macbook pro', '', 1, 1, 1, 1),
(37, 'cpu型号', 2, 1, 'intel i5,intel i7,intel i9,amd', '', 1, 1, 1, 2),
(38, '内存容量', 2, 2, '8gb,16gb,32gb,64gb', 'gb', 1, 1, 1, 3),
(39, '显卡型号', 1, 1, '', '', 1, 1, 1, 4),
(40, '硬盘容量', 2, 2, '256gb,512gb,1tb,2tb', 'gb', 1, 1, 1, 5),
-- 小米电脑属性
(41, '品牌', 2, 1, 'xiaomi plus, xiaomi pro', '', 1, 1, 1, 1),
(42, 'cpu型号', 2, 1, 'intel i5,intel i7,intel i9,amd', '', 1, 1, 1, 2),
(43, '内存容量', 2, 2, '8gb,16gb,32gb,64gb', 'gb', 1, 1, 1, 3),
(44, '显卡型号', 1, 1, '', '', 1, 1, 1, 4),
(45, '硬盘容量', 2, 2, '256gb,512gb,1tb,2tb', 'gb', 1, 1, 1, 5),
-- 华为电脑属性
(46, '品牌', 2, 1, 'huawei pro, huawei plus', '', 1, 1, 1, 1),
(47, 'cpu型号', 2, 1, 'intel i5,intel i7,intel i9,amd', '', 1, 1, 1, 2),
(48, '内存容量', 2, 2, '8gb,16gb,32gb,64gb', 'gb', 1, 1, 1, 3),
(49, '显卡型号', 1, 1, '', '', 1, 1, 1, 4),
(50, '硬盘容量', 2, 2, '256gb,512gb,1tb,2tb', 'gb', 1, 1, 1, 5),

-- 服装通用属性
(51, '品牌', 2, 1, 'nike,adidas,uniqlo,h&m', '', 1, 1, 1, 1),
(52, '材质成分', 1, 1, '', '', 1, 0, 1, 3),
(53, '尺码', 2, 1, 's,m,l,xl,xxl', '', 1, 1, 1, 4),
(54, '适用季节', 3, 1, '春季,夏季,秋季,冬季', '', 1, 1, 1, 2),
-- 海澜之家服装通用属性
(55, '品牌', 2, 1, '海澜之家', '', 1, 1, 1, 1),
(56, '材质成分', 1, 1, '', '', 1, 0, 1, 3),
(57, '尺码', 2, 1, 's,m,l,xl,xxl', '', 1, 1, 1, 4),
(58, '适用季节', 3, 1, '春季,夏季,秋季,冬季', '', 1, 1, 1, 2),

-- 厨房电器属性
(59, '品牌', 2, 1, 'midea,haier,gree,siemens', '', 1, 1, 1, 1),
(59, '型号', 1, 1, '', '', 1, 1, 1, 2),
(60, '功率', 2, 1, '2000,3000', 'w', 1, 0, 1, 3),
(61, '能效等级', 2, 1, 'a+++,a++,a+,a,b', '', 1, 1, 1, 4),
-- 电饭锅电器属性
(62, '品牌', 2, 1, 'midea,haier,gree,siemens', '', 1, 1, 1, 1),
(62, '型号', 1, 1, '', '', 1, 1, 1, 2),
(63, '功率', 2, 1, '2000,3000', 'w', 1, 0, 1, 3),
(64, '能效等级', 2, 1, 'a+++,a++,a+,a,b', '', 1, 1, 1, 4),


-- 空调属性
(65, '品牌', 2, 1, 'midea,haier,gree,siemens', '', 1, 1, 1, 1),
(65, '型号', 1, 1, '', '', 1, 1, 1, 2),
(66, '功率', 2, 1, '2000,3000', 'w', 1, 0, 1, 3),
(67, '能效等级', 2, 1, 'a+++,a++,a+,a,b', '', 1, 1, 1, 4),
-- 格力空调属性
(68, '品牌', 2, 1, 'gree', '', 1, 1, 1, 1),
(68, '型号', 1, 1, '', '', 1, 1, 1, 2),
(69, '功率', 2, 1, '2000,3000', 'w', 1, 0, 1, 3),
(70, '能效等级', 2, 1, 'a+++,a++,a+,a,b', '', 1, 1, 1, 4),

-- 厨房电器属性
(71, '控制方式', 2, 1, '按键式,触控式,旋钮式,智能控制', '', 1, 1, 1, 2),
(72, '功率', 2, 1, '2000,3000', 'w', 1, 0, 1, 3),
(73, '特色功能', 3, 1, '预约,保温,除菌,智能控温', '', 0, 1, 1, 3),
-- 方太厨房电器属性
(74, '控制方式', 2, 1, '按键式,触控式,旋钮式,智能控制', '', 1, 1, 1, 2),
(75, '功率', 2, 1, '2000,3000', 'w', 1, 0, 1, 3),
(76, '特色功能', 3, 1, '预约,保温,除菌,智能控温', '', 0, 1, 1, 3),

-- 热水器电器属性
(77, '容量', 2, 1, '60,80,100', 'l', 1, 1, 1, 1),
(77, '控制方式', 2, 1, '按键式,触控式,旋钮式,智能控制', '', 1, 1, 1, 2),
(78, '功率', 2, 1, '2000,3000', 'w', 1, 0, 1, 3),
(79, '特色功能', 3, 1, '预约,保温,除菌,智能控温', '', 0, 1, 5, 4),
(80, '安全参数', 2, 1, '出水断电,出水不断电', '', 0, 1, 1, 6),

-- 万和热水器属性
(81, '容量', 2, 1, '60,80,100', 'l', 1, 1, 1, 1),
(81, '控制方式', 2, 1, '按键式,触控式,旋钮式,智能控制', '', 1, 1, 1, 2),
(82, '功率', 2, 1, '2000,3000', 'w', 1, 0, 1, 3),
(83, '特色功能', 3, 1, '预约,保温,除菌,智能控温', '', 0, 1, 5, 4),
(84, '安全参数', 2, 1, '出水断电,出水不断电', '', 0, 1, 1, 6),

-- 万家乐热水器属性
(85, '容量', 2, 1, '60,80,100', 'l', 1, 1, 1, 1),
(85, '控制方式', 2, 1, '按键式,触控式,旋钮式,智能控制', '', 1, 1, 1, 2),
(86, '功率', 2, 1, '2000,3000', 'w', 1, 0, 1, 3),
(87, '特色功能', 3, 1, '预约,保温,除菌,智能控温', '', 0, 1, 5, 4),
(88, '安全参数', 2, 1, '出水断电,出水不断电', '', 0, 1, 1, 6),

-- 海尔热水器属性
(89, '容量', 2, 1, '60,80,100', 'l', 1, 1, 1, 1),
(89, '控制方式', 2, 1, '按键式,触控式,旋钮式,智能控制', '', 1, 1, 1, 2),
(90, '功率', 2, 1, '2000,3000', 'w', 1, 0, 1, 3),
(91, '特色功能', 3, 1, '预约,保温,除菌,智能控温', '', 0, 1, 5, 4),
(92, '安全参数', 2, 1, '出水断电,出水不断电', '', 0, 1, 1, 6),

-- 美的热水器属性
(93, '容量', 2, 1, '60,80,100', 'l', 1, 1, 1, 1),
(93, '控制方式', 2, 1, '按键式,触控式,旋钮式,智能控制', '', 1, 1, 1, 2),
(94, '功率', 2, 1, '2000,3000', 'w', 1, 0, 1, 3),
(95, '特色功能', 3, 1, '预约,保温,除菌,智能控温', '', 0, 1, 5, 4),
(96, '安全参数', 2, 1, '出水断电,出水不断电', '', 0, 1, 1, 6);

-- 商品属性分组表建语句
drop table if exists pms_product_attribute_group;
create table pms_product_attribute_group
(
    id          bigserial primary key,
    category_id bigint                                not null,
    name        varchar                               not null,
    sort        integer                               not null,
    status      integer     default 1                 not null,
    create_by   varchar     default ''                not null,
    create_time timestamptz default current_timestamp not null,
    update_by   varchar     default ''                not null,
    update_time timestamptz                           null,
    is_deleted  integer     default 1                 not null
);

-- 添加商品属性分组表注释
comment on table pms_product_attribute_group is '商品属性分组表';

-- 添加商品属性分组表列注释
comment on column pms_product_attribute_group.id is '主键id';
comment on column pms_product_attribute_group.category_id is '分类id';
comment on column pms_product_attribute_group.name is '分组名称';
comment on column pms_product_attribute_group.sort is '排序';
comment on column pms_product_attribute_group.status is '状态：0->禁用；1->启用';
comment on column pms_product_attribute_group.create_by is '创建人id';
comment on column pms_product_attribute_group.create_time is '创建时间';
comment on column pms_product_attribute_group.update_by is '更新人id';
comment on column pms_product_attribute_group.update_time is '更新时间';
comment on column pms_product_attribute_group.is_deleted is '是否删除(0:否,1:是)';

-- 插入商品属性分组数据
insert into pms_product_attribute_group (category_id, name, sort)
values

-- 手机(id=1)属性组
(1, '基本信息', 1),
(1, '主体参数', 2),
(1, '网络参数', 3),
(1, '显示参数', 4),
(1, '摄像功能', 5),
-- 小米手机(id=8)属性组
(8, '基本信息', 1),
(8, '主体参数', 2),
(8, '网络参数', 3),
(8, '显示参数', 4),
(8, '摄像功能', 5),
-- 苹果手机(id=9)属性组
(9, '基本信息', 1),
(9, '主体参数', 2),
(9, '网络参数', 3),
(9, '显示参数', 4),
(9, '摄像功能', 5),
-- 华为手机(id=10)属性组
(10, '基本信息', 1),
(10, '主体参数', 2),
(10, '网络参数', 3),
(10, '显示参数', 4),
(10, '摄像功能', 5),
-- 三星手机(id=11)属性组
(11, '基本信息', 1),
(11, '主体参数', 2),
(11, '网络参数', 3),
(11, '显示参数', 4),
(11, '摄像功能', 5),
-- 荣耀手机(id=12)属性组
(12, '基本信息', 1),
(12, '主体参数', 2),
(12, '网络参数', 3),
(12, '显示参数', 4),
(12, '摄像功能', 5),

    -- 电脑(id=2)属性组
(2, '基本信息', 1),
(2, '处理器', 2),
(2, '内存容量', 3),
(2, '显卡性能', 4),
(2, '存储容量', 5),
-- 苹果电脑(id=13)属性组
(13, '基本信息', 1),
(13, '处理器', 2),
(13, '内存容量', 3),
(13, '显卡性能', 4),
(13, '存储容量', 5),
-- 小米电脑(id=14)属性组
(14, '基本信息', 1),
(14, '处理器', 2),
(14, '内存容量', 3),
(14, '显卡性能', 4),
(14, '存储容量', 5),
-- 华为电脑(id=15)属性组
(15, '基本信息', 1),
(15, '处理器', 2),
(15, '内存容量', 3),
(15, '显卡性能', 4),
(15, '存储容量', 5),

-- 服装(id=3)通用属性组
(3, '基本信息', 1),
(3, '材质参数', 2),
(3, '尺码参数', 3),
(3, '季节', 4),
-- 海澜之家(id=16)通用属性组
(16, '基本信息', 1),
(16, '材质参数', 2),
(16, '尺码参数', 3),
(16, '季节', 4),

-- 电器(id=4)通用属性组
(4, '基本信息', 1),
(4, '规格参数', 2),
(4, '功能特点', 3),
-- 电饭锅(id=17)通用属性组
(17, '基本信息', 1),
(17, '规格参数', 2),
(17, '功能特点', 3),

-- 空调(id=5)通用属性组
(5, '基本信息', 1),
(5, '规格参数', 2),
(5, '功能特点', 3),
-- 格力空调(id=17)通用属性组
(18, '基本信息', 1),
(18, '规格参数', 2),
(18, '功能特点', 3),

-- 厨房电器(id=6)属性组
(6, '基本信息', 1),
(6, '功率参数', 2),
(6, '特色功能', 3),
-- 方太电器(id=19)属性组
(19, '基本信息', 1),
(19, '功率参数', 2),
(19, '特色功能', 3),

-- 热水器(id=7)属性组
(7, '基本信息', 1),
(7, '功率参数', 2),
(7, '特色功能', 3),
(7, '安全参数', 4),
-- 万和热水器(id=20)属性组
(20, '基本信息', 1),
(20, '功率参数', 2),
(20, '特色功能', 3),
(20, '安全参数', 4),
-- 万家乐热水器(id=21)属性组
(21, '基本信息', 1),
(21, '功率参数', 2),
(21, '特色功能', 3),
(21, '安全参数', 4),
-- 海尔热水器(id=22)属性组
(22, '基本信息', 1),
(22, '功率参数', 2),
(22, '特色功能', 3),
(22, '安全参数', 4),
-- 美的热水器(id=23)属性组
(23, '基本信息', 1),
(23, '功率参数', 2),
(23, '特色功能', 3),
(23, '安全参数', 4);

-- 商品属性值表建语句
drop table if exists pms_product_attribute_value;
create table pms_product_attribute_value
(
    id           bigserial primary key,
    spu_id       bigint                                not null,
    attribute_id bigint                                not null,
    value        varchar                               not null,
    status       integer     default 1                 not null,
    create_by    varchar     default ''                not null,
    create_time  timestamptz default current_timestamp not null,
    update_by    varchar     default ''                not null,
    update_time  timestamptz                           null,
    is_deleted   integer     default 1                 not null
);

-- 添加商品属性值表注释
comment on table pms_product_attribute_value is '商品属性值表';

-- 添加商品属性值表列注释
comment on column pms_product_attribute_value.id is '主键id';
comment on column pms_product_attribute_value.spu_id is '商品spu id';
comment on column pms_product_attribute_value.attribute_id is '属性id';
comment on column pms_product_attribute_value.value is '属性值';
comment on column pms_product_attribute_value.status is '状态：0->禁用；1->启用';
comment on column pms_product_attribute_value.create_by is '创建人id';
comment on column pms_product_attribute_value.create_time is '创建时间';
comment on column pms_product_attribute_value.update_by is '更新人id';
comment on column pms_product_attribute_value.update_time is '更新时间';
comment on column pms_product_attribute_value.is_deleted is '是否删除(0:否,1:是)';

-- 插入商品属性值数据（假设iphone 15 pro的spu_id=1）
insert into pms_product_attribute_value (spu_id, attribute_id, value, create_by)
values
-- 小米手机基本信息属性值
(1, 24, 'xiaomi', 1),                    -- 品牌
(1, 25, '1888', 1),                      -- 型号
(1, 26, '2026', 1),                      -- 上市年份
(1, 27, '自然色', 1),                    -- 机身颜色
-- 小米手机主体参数属性值
(1, 28, 'a17', 1),                       -- 处理器
(1, 29, '8gb', 1),                       -- 运行内存
(1, 30, '256gb', 1),                     -- 机身存储
(1, 31, '3274mah', 1),                   -- 电池容量
(1, 32, '187g', 1),                      -- 机身重量
-- 小米手机网络参数属性值
(1, 33, '支持', 1),                      -- 5g网络
(1, 34, 'gsm,wcdma,lte,5g', 1),          -- 网络制式
(1, 35, '双卡双待', 1),                  -- sim卡类型
(1, 36, 'wifi 6e', 1),                   -- wifi
-- 小米手机显示参数属性值
(1, 37, '6.9', 1),                       -- 屏幕尺寸
(1, 38, '2556x1179', 1),                 -- 屏幕分辨率
(1, 39, '120', 1),                       -- 屏幕刷新率
(1, 40, 'oled', 1),                      -- 屏幕类型
(1, 41, '2000', 1),                      -- 屏幕亮度
-- 小米手机摄像功能属性值
(1, 42, '4800', 1),                      -- 主摄像头
(1, 43, '1200', 1),                      -- 前置摄像头
(1, 44, '3', 1),                         -- 摄像头数量
(1, 45, '4k,1080p', 1),                  -- 视频拍摄
(1, 46, '光学防抖', 1),                  -- 防抖功能

-- 苹果手机基本信息属性值
(2, 47, 'apple', 1),                     -- 品牌
(2, 48, 'iphone 15 pro', 1),             -- 型号
(2, 49, '2023', 1),                      -- 上市年份
(2, 20, '自然色', 1),                    -- 机身颜色
-- 苹果手机主体参数属性值
(2, 51, 'a17', 1),                       -- 处理器
(2, 52, '8gb', 1),                       -- 运行内存
(2, 53, '256gb', 1),                     -- 机身存储
(2, 54, '3274mah', 1),                   -- 电池容量
(2, 55, '187g', 1),                      -- 机身重量
-- 苹果手机网络参数属性值
(2, 56, '支持', 1),                      -- 5g网络
(2, 57, 'gsm,wcdma,lte,5g', 1),          -- 网络制式
(2, 58, '双卡双待', 1),                  -- sim卡类型
(2, 59, 'wifi 6e', 1),                   -- wifi
-- 苹果手机显示参数属性值
(2, 60, '6.1', 1),                       -- 屏幕尺寸
(2, 61, '2556x1179', 1),                 -- 屏幕分辨率
(2, 62, '120', 1),                       -- 屏幕刷新率
(2, 63, 'oled', 1),                      -- 屏幕类型
(2, 64, '2000', 1),                      -- 屏幕亮度
-- 苹果手机摄像功能属性值
(2, 65, '4800', 1),                      -- 主摄像头
(2, 66, '1200', 1),                      -- 前置摄像头
(2, 67, '3', 1),                         -- 摄像头数量
(2, 68, '4k,1080p', 1),                  -- 视频拍摄
(2, 69, '光学防抖', 1),                  -- 防抖功能

-- 华为手机基本信息属性值
(3, 70, 'huawei', 1),                    -- 品牌
(3, 71, 'mate 80', 1),                   -- 型号
(3, 72, '2023', 1),                      -- 上市年份
(3, 73, '自然色', 1),                    -- 机身颜色
-- 华为手机主体参数属性值
(3, 74, 'a17 pro', 1),                   -- 处理器
(3, 75, '8gb', 1),                       -- 运行内存
(3, 76, '256gb', 1),                     -- 机身存储
(3, 77, '3274mah', 1),                   -- 电池容量
(3, 78, '187g', 1),                      -- 机身重量
-- 华为手机网络参数属性值
(3, 79, '支持', 1),                      -- 5g网络
(3, 80, 'gsm,wcdma,lte,5g', 1),          -- 网络制式
(3, 81, '双卡双待', 1),                  -- sim卡类型
(3, 82, 'wifi 6e', 1),                   -- wifi
-- 华为手机显示参数属性值
(3, 83, '6.1', 1),                       -- 屏幕尺寸
(3, 84, '2556x1179', 1),                 -- 屏幕分辨率
(3, 85, '120', 1),                       -- 屏幕刷新率
(3, 86, 'oled', 1),                      -- 屏幕类型
(3, 87, '2000', 1),                      -- 屏幕亮度
-- 华为手机摄像功能属性值
(3, 88, '4800', 1),                      -- 主摄像头
(3, 89, '1200', 1),                      -- 前置摄像头
(3, 90, '3', 1),                         -- 摄像头数量
(3, 91, '4k,1080p', 1),                  -- 视频拍摄
(3, 92, '光学防抖', 1),                  -- 防抖功能

-- 三星手机基本信息属性值
(4, 93, 'ultra', 1),                     -- 品牌
(4, 94, 'ultra', 1),                     -- 型号
(4, 95, '2023', 1),                      -- 上市年份
(4, 96, '自然色', 1),                    -- 机身颜色
-- 三星手机主体参数属性值
(4, 97, 'a17 pro', 1),                   -- 处理器
(4, 98, '8gb', 1),                       -- 运行内存
(4, 99, '256gb', 1),                     -- 机身存储
(4, 100, '3274mah', 1),                  -- 电池容量
(4, 101, '187g', 1),                     -- 机身重量
-- 三星手机网络参数属性值
(4, 102, '支持', 1),                     -- 5g网络
(4, 103, 'gsm,wcdma,lte,5g', 1),         -- 网络制式
(4, 104, '双卡双待', 1),                 -- sim卡类型
(4, 105, 'wifi 6e', 1),                  -- wifi
-- 三星手机显示参数属性值
(4, 106, '6.1', 1),                      -- 屏幕尺寸
(4, 107, '2556x1179', 1),                -- 屏幕分辨率
(4, 108, '120', 1),                      -- 屏幕刷新率
(4, 109, 'oled', 1),                     -- 屏幕类型
(4, 110, '2000', 1),                     -- 屏幕亮度
-- 三星手机摄像功能属性值
(4, 111, '4800', 1),                     -- 主摄像头
(4, 112, '1200', 1),                     -- 前置摄像头
(4, 113, '3', 1),                        -- 摄像头数量
(4, 114, '4k,1080p', 1),                 -- 视频拍摄
(4, 115, '光学防抖', 1),                 -- 防抖功能

-- 荣耀手机基本信息属性值
(5, 116, '荣耀', 1),                     -- 品牌
(5, 117, '荣耀power2', 1),               -- 型号
(5, 118, '2023', 1),                     -- 上市年份
(5, 119, '自然色', 1),                   -- 机身颜色
-- 荣耀手机主体参数属性值
(5, 120, 'a17 pro', 1),                  -- 处理器
(5, 121, '8gb', 1),                      -- 运行内存
(5, 122, '256gb', 1),                    -- 机身存储
(5, 123, '3274mah', 1),                  -- 电池容量
(5, 124, '187g', 1),                     -- 机身重量
-- 荣耀手机网络参数属性值
(5, 125, '支持', 1),                     -- 5g网络
(5, 126, 'gsm,wcdma,lte,5g', 1),         -- 网络制式
(5, 127, '双卡双待', 1),                 -- sim卡类型
(5, 128, 'wifi 6e', 1),                  -- wifi
-- 荣耀手机显示参数属性值
(5, 129, '6.1', 1),                      -- 屏幕尺寸
(5, 130, '2556x1179', 1),                -- 屏幕分辨率
(5, 131, '120', 1),                      -- 屏幕刷新率
(5, 132, 'oled', 1),                     -- 屏幕类型
(5, 133, '2000', 1),                     -- 屏幕亮度
-- 荣耀手机摄像功能属性值
(5, 134, '4800', 1),                     -- 主摄像头
(5, 135, '1200', 1),                     -- 前置摄像头
(5, 136, '3', 1),                        -- 摄像头数量
(5, 137, '4k,1080p', 1),                 -- 视频拍摄
(5, 138, '光学防抖', 1),                 -- 防抖功能

-- 苹果电脑属性值
(6, 144, 'macbook air', 1),              -- 品牌
(6, 145, 'a19', 1),                      -- cpu型号
(6, 146, '32g', 1),                      -- 内存容量
(6, 147, '123456', 1),                   -- 显卡型号
(6, 148, '512g', 1),                     -- 硬盘容量

-- 小米电脑属性值
(7, 149, 'xiaobook air', 1),             -- 品牌
(7, 150, 'a19', 1),                      -- cpu型号
(7, 151, '32g', 1),                      -- 内存容量
(7, 152, '123456', 1),                   -- 显卡型号
(7, 153, '512g', 1),                     -- 硬盘容量

-- 华为电脑属性值
(8, 154, 'matebook', 1),                 -- 品牌
(8, 155, 'a19', 1),                      -- cpu型号
(8, 156, '32g', 1),                      -- 内存容量
(8, 157, '123456', 1),                   -- 显卡型号
(8, 158, '512g', 1),                     -- 硬盘容量

-- 海澜之家服装通用属性值
(9, 163, '海澜之家', 1),                 -- 品牌
(9, 164, '棉', 1),                       -- 材质成分
(9, 165, 'xl', 1),                       -- 尺码
(9, 166, '春季,夏季,秋季,冬季', 1),      -- 适用季节

-- 电饭锅电器属性值
(10, 171, 'super', 1),                   -- 品牌
(10, 172, '2026', 1),                    -- 型号
(10, 173, '2000', 1),                    -- 功率
(10, 174, 'a+++', 1),                    -- 能效等级


-- 格力空调属性值
(11, 179, 'gree', 1),                    -- 品牌
(11, 180, '2026', 1),                    -- 型号
(11, 181, '2000', 1),                    -- 功率
(11, 182, 'a+++', 1),                    -- 能效等级

-- 方太属性值
(12, 186, '按键式', 1),                  -- 控制方式
(12, 187, '2026', 1),                    -- 功率
(12, 188, '预约,保温,除菌,智能控温', 1), -- 特色功能


-- 万和热水器属性值
(13, 194, '100', 1),                     -- 容量
(13, 195, '按键式', 1),                  -- 控制方式
(13, 196, '2000', 1),                    -- 功率
(13, 197, '预约,保温,除菌,智能控温', 1), -- 特色功能
(13, 198, '出水断电', 1),                -- 安全参数

-- 万家乐热水器属性值
(14, 199, '100', 1),                     -- 容量
(14, 200, '触控式', 1),                  -- 控制方式
(14, 201, '2000', 1),                    -- 功率
(14, 202, '预约,保温,除菌,智能控温', 1), -- 特色功能
(14, 203, '出水断电', 1),                -- 安全参数

-- 海尔热水器属性值
(15, 204, '80', 1),                      -- 容量
(15, 205, '智能控制', 1),                -- 控制方式
(15, 206, '2000', 1),                    -- 功率
(15, 207, '预约,保温,除菌,智能控温', 1), -- 特色功能
(15, 208, '出水断电', 1),                -- 安全参数

-- 美的热水器属性值
(16, 209, '60', 1),                      -- 容量
(16, 210, '旋钮式', 1),                  -- 控制方式
(16, 211, '3000', 1),                    -- 功率
(16, 212, '预约,保温,除菌,智能控温', 1), -- 特色功能
(16, 213, '出水断电', 1); -- 安全参数

-- 商品品牌建语句
drop table if exists pms_product_brand;
create table pms_product_brand
(
    id                    bigserial primary key,
    name                  varchar                               not null,
    logo                  varchar                               not null,
    big_pic               varchar                               not null,
    description           text                                  not null,
    first_letter          varchar                               not null,
    sort                  integer                               not null,
    recommend_status      integer     default 1                 not null,
    product_count         integer                               not null,
    product_comment_count integer                               not null,
    is_enabled            integer     default 1                 not null,
    create_by             varchar     default ''                not null,
    create_time           timestamptz default current_timestamp not null,
    update_by             varchar     default ''                not null,
    update_time           timestamptz                           null,
    is_deleted            integer     default 1                 not null

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
comment on column pms_product_brand.create_by is '创建人id';
comment on column pms_product_brand.create_time is '创建时间';
comment on column pms_product_brand.update_by is '更新人id';
comment on column pms_product_brand.update_time is '更新时间';
comment on column pms_product_brand.is_deleted is '是否删除(0:否,1:是)';

insert into pms_product_brand (id, name, logo, big_pic, description, first_letter, sort, recommend_status, product_count, product_comment_count, is_enabled, create_by)
values (1, '苹果', 'http://129.204.203.29/brand_apple.png', 'http://129.204.203.29/brand_apple.png', '苹果的描述', 'a', 1, 1, 100, 50, 1, 1),
       (2, '方太', 'http://129.204.203.29/brand_fotile.png', 'http://129.204.203.29/brand_fotile.png', '方太的描述', 'f', 2, 0, 200, 80, 1, 2),
       (3, '格力', 'http://129.204.203.29/brand_geli.png', 'http://129.204.203.29/brand_geli.png', '格力的描述', 'g', 3, 1, 150, 60, 1, 3),
       (4, '海澜之家', 'http://129.204.203.29/brand_hailan.png', 'http://129.204.203.29/brand_hailan.png', '海澜之家的描述', 'h', 4, 1, 120, 40, 1, 4),
       (5, '华为', 'http://129.204.203.29/brand_huawei.png', 'http://129.204.203.29/brand_huawei.png', '华为的描述', 'h', 5, 1, 180, 70, 1, 5),
       (6, 'oppo', 'http://129.204.203.29/brand_oppo.png', 'http://129.204.203.29/brand_oppo.png', 'oppo的描述', 'o', 6, 0, 90, 30, 1, 6),
       (7, '三星', 'http://129.204.203.29/brand_sumsung.png', 'http://129.204.203.29/brand_sumsung.png', '三星的描述', 's', 7, 1, 110, 50, 1, 7),
       (8, '万和', 'http://129.204.203.29/brand_wanhe.png', 'http://129.204.203.29/brand_wanhe.png', '万和的描述', 'w', 8, 0, 130, 60, 1, 8),
       (9, '小米', 'http://129.204.203.29/brand_xiaomi.png', 'http://129.204.203.29/brand_xiaomi.png', '小米的描述', 'x', 9, 1, 140, 55, 1, 9);


-- 产品分类建语句
drop table if exists pms_product_category;
create table pms_product_category
(
    id            bigserial primary key,
    parent_id     bigint                                not null,
    name          varchar                               not null,
    level         integer     default 1                 not null,
    product_count integer                               not null,
    product_unit  varchar                               not null,
    nav_status    integer     default 1                 not null,
    sort          integer                               not null,
    icon          varchar                               not null,
    keywords      varchar     default ''                not null,
    description   varchar     default ''                not null,
    is_enabled    integer     default 1                 not null,
    create_by     varchar     default ''                not null,
    create_time   timestamptz default current_timestamp not null,
    update_by     varchar     default ''                not null,
    update_time   timestamptz                           null,
    is_deleted    integer     default 1                 not null
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
comment on column pms_product_category.create_by is '创建人id';
comment on column pms_product_category.create_time is '创建时间';
comment on column pms_product_category.update_by is '更新人id';
comment on column pms_product_category.update_time is '更新时间';
comment on column pms_product_category.is_deleted is '是否删除(0:否,1:是)';

insert into pms_product_category (id, parent_id, name, level, product_count, product_unit, nav_status, sort, icon)
values (1, 0, '手机', 0, 100, '台', 1, 1, ''),
       (2, 0, '电脑', 0, 200, '台', 1, 2, ''),
       (3, 0, '衣服', 0, 150, '件', 1, 3, ''),
       (4, 0, '电器', 0, 50, '台', 1, 4, ''),
       (5, 0, '空调', 0, 30, '件', 1, 5, ''),
       (6, 0, '厨房', 0, 810, '件', 1, 6, ''),
       (7, 0, '热水器', 0, 320, '件', 1, 7, ''),
       (8, 1, '小米手机', 1, 620, '台', 1, 8, 'http://129.204.203.29/xiaomi_s.jpg'),
       (9, 1, '苹果手机', 1, 90, '台', 1, 9, 'http://129.204.203.29/apple_s.jpg'),
       (10, 1, '华为手机', 1, 70, '台', 1, 10, 'http://129.204.203.29/hua_s.jpg'),
       (11, 1, '三星手机', 1, 340, '台', 1, 11, 'http://129.204.203.29/sumsang.jpg'),
       (12, 1, '荣耀手机', 1, 470, '台', 1, 12, 'http://129.204.203.29/rong_s.jpg'),
       (13, 2, '苹果电脑', 1, 140, '台', 1, 13, 'http://129.204.203.29/apple_c.jpg'),
       (14, 2, '小米电脑', 1, 410, '台', 1, 14, 'http://129.204.203.29/xiaomi_c.jpg'),
       (15, 2, '华为电脑', 1, 950, '台', 1, 15, 'http://129.204.203.29/hua_c.jpg'),
       (16, 3, '海澜之家', 1, 840, '件', 1, 16, 'http://129.204.203.29/hailan.jpg'),
       (17, 4, '电饭锅', 1, 540, '件', 1, 17, 'http://129.204.203.29/su.jpg'),
       (18, 5, '格力空调', 1, 400, '台', 1, 18, 'http://129.204.203.29/geli.jpg'),
       (19, 6, '方太', 1, 406, '件', 1, 19, 'http://129.204.203.29/fangtai.jpg'),
       (20, 7, '万和', 1, 70, '件', 1, 20, 'http://129.204.203.29/wanhe.png'),
       (21, 7, '万家乐', 1, 90, '件', 1, 21, 'http://129.204.203.29/wanjiale.jpg'),
       (22, 7, '海尔', 1, 112, '件', 1, 22, 'http://129.204.203.29/haier.jpg'),
       (23, 7, '美的', 1, 150, '件', 1, 23, 'http://129.204.203.29/meidi.png');


-- 产品的分类和属性的关系表，用于设置分类筛选条件（只支持一级分类）建语句
drop table if exists pms_product_category_attribute_relation;
create table pms_product_category_attribute_relation
(
    id                   bigserial primary key,
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
    id           bigserial primary key,
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

insert into pms_product_full_reduction (id, product_id, full_price, reduce_price) values (1, 1, 100.00, 20.00);
insert into pms_product_full_reduction (id, product_id, full_price, reduce_price) values (2, 2, 100.00, 20.00);
insert into pms_product_full_reduction (id, product_id, full_price, reduce_price) values (3, 3, 100.00, 20.00);
insert into pms_product_full_reduction (id, product_id, full_price, reduce_price) values (4, 4, 100.00, 20.00);
insert into pms_product_full_reduction (id, product_id, full_price, reduce_price) values (5, 5, 100.00, 20.00);
insert into pms_product_full_reduction (id, product_id, full_price, reduce_price) values (6, 6, 100.00, 20.00);
insert into pms_product_full_reduction (id, product_id, full_price, reduce_price) values (7, 7, 100.00, 20.00);
insert into pms_product_full_reduction (id, product_id, full_price, reduce_price) values (8, 8, 100.00, 20.00);
insert into pms_product_full_reduction (id, product_id, full_price, reduce_price) values (9, 9, 100.00, 20.00);
insert into pms_product_full_reduction (id, product_id, full_price, reduce_price) values (10, 10, 200.00, 50.00);
insert into pms_product_full_reduction (id, product_id, full_price, reduce_price) values (11, 11, 300.00, 100.00);
insert into pms_product_full_reduction (id, product_id, full_price, reduce_price) values (14, 12, 0.00, 0.00);
insert into pms_product_full_reduction (id, product_id, full_price, reduce_price) values (16, 13, 0.00, 0.00);
insert into pms_product_full_reduction (id, product_id, full_price, reduce_price) values (34, 14, 0.00, 0.00);
insert into pms_product_full_reduction (id, product_id, full_price, reduce_price) values (44, 15, 0.00, 0.00);
insert into pms_product_full_reduction (id, product_id, full_price, reduce_price) values (46, 16, 0.00, 0.00);

-- 产品阶梯价格表(只针对同商品)建语句
drop table if exists pms_product_ladder;
create table pms_product_ladder
(
    id         bigserial primary key,
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

insert into pms_product_ladder (id, product_id, count, discount, price) values (1, 1, 3, 0.70, 0.00);
insert into pms_product_ladder (id, product_id, count, discount, price) values (2, 2, 3, 0.70, 0.00);
insert into pms_product_ladder (id, product_id, count, discount, price) values (3, 3, 3, 0.70, 0.00);
insert into pms_product_ladder (id, product_id, count, discount, price) values (4, 4, 3, 0.70, 0.00);
insert into pms_product_ladder (id, product_id, count, discount, price) values (5, 5, 3, 0.70, 0.00);
insert into pms_product_ladder (id, product_id, count, discount, price) values (6, 6, 3, 0.70, 0.00);
insert into pms_product_ladder (id, product_id, count, discount, price) values (7, 7, 3, 0.70, 0.00);
insert into pms_product_ladder (id, product_id, count, discount, price) values (8, 8, 3, 0.70, 0.00);
insert into pms_product_ladder (id, product_id, count, discount, price) values (12, 9, 3, 0.70, 0.00);
insert into pms_product_ladder (id, product_id, count, discount, price) values (14, 10, 4, 0.60, 0.00);
insert into pms_product_ladder (id, product_id, count, discount, price) values (15, 11, 5, 0.50, 0.00);
insert into pms_product_ladder (id, product_id, count, discount, price) values (18, 12, 0, 0.00, 0.00);
insert into pms_product_ladder (id, product_id, count, discount, price) values (20, 13, 0, 0.00, 0.00);
insert into pms_product_ladder (id, product_id, count, discount, price) values (38, 14, 0, 0.00, 0.00);
insert into pms_product_ladder (id, product_id, count, discount, price) values (48, 15, 0, 0.00, 0.00);
insert into pms_product_ladder (id, product_id, count, discount, price) values (50, 16, 0, 0.00, 0.00);

-- 商品sku表建语句
drop table if exists pms_product_sku;
create table pms_product_sku
(
    id                   bigserial primary key,
    spu_id               bigint                                not null,
    name                 varchar                               not null,
    sku_code             varchar                               not null,
    main_pic             varchar                               not null,
    album_pics           varchar                               not null,
    price                numeric                               not null,
    promotion_price      numeric                               not null,
    promotion_start_time timestamptz                            null,
    promotion_end_time   timestamptz                            null,
    stock                integer                               not null,
    low_stock            integer                               not null,
    spec_data            jsonb                                 not null,
    weight               numeric                               not null,
    publish_status       integer     default 1                 not null,
    verify_status        integer     default 1                 not null,
    sort                 integer                               not null,
    sales                integer                               not null,
    create_by            varchar     default ''                not null,
    create_time          timestamptz default current_timestamp not null,
    update_by            varchar     default ''                not null,
    update_time          timestamptz                           null,
    is_deleted           integer     default 1                 not null
);

-- 添加商品sku表注释
comment on table pms_product_sku is '商品sku表';

-- 添加商品sku表列注释
comment on column pms_product_sku.id is '商品skuid';
comment on column pms_product_sku.spu_id is '商品spuid';
comment on column pms_product_sku.name is 'sku名称';
comment on column pms_product_sku.sku_code is 'sku编码';
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
comment on column pms_product_sku.create_by is '创建人id';
comment on column pms_product_sku.create_time is '创建时间';
comment on column pms_product_sku.update_by is '更新人id';
comment on column pms_product_sku.update_time is '更新时间';
comment on column pms_product_sku.is_deleted is '是否删除(0:否,1:是)';

insert into pms_product_sku ( spu_id, name, sku_code, main_pic, album_pics, price, promotion_price, stock, low_stock, spec_data, weight, publish_status, verify_status, sort, sales)
values
--小米手机
( 1, '小米手机 金色 128gb 全网通版', '123', 'http://129.204.203.29/xiaomi_s.jpg', '', 7999.00, 7899.00, 1000, 100, '{"颜色": "金色", "容量": "128gb", "网络版本": "全网通版"}', 0.19, 1, 1, 1, 300),
( 1, '小米手机 金色 256gb 全网通版', '1231', 'http://129.204.203.29/xiaomi_s.jpg', '', 8999.00, 8899.00, 800, 80, '{"颜色": "金色", "容量": "256gb", "网络版本": "全网通版"}', 0.19, 1, 1, 2, 200),
--苹果手机
( 2, '苹果手机 金色 128gb 全网通版', '1223', 'http://129.204.203.29/apple_s.jpg', '', 7999.00, 7899.00, 1000, 100, '{"颜色": "金色", "容量": "128gb", "网络版本": "全网通版"}', 0.19, 1, 1, 1, 300),
( 2, '苹果手机 金色 256gb 全网通版', '1233', 'http://129.204.203.29/apple_s.jpg', '', 8999.00, 8899.00, 800, 80, '{"颜色": "金色", "容量": "256gb", "网络版本": "全网通版"}', 0.19, 1, 1, 2, 200),
--华为手机
( 3, '华为手机 金色 128gb 全网通版', '1213', 'http://129.204.203.29/hua_s.jpg', '', 7999.00, 7899.00, 1000, 100, '{"颜色": "金色", "容量": "128gb", "网络版本": "全网通版"}', 0.19, 1, 1, 1, 300),
( 3, '华为手机 金色 256gb 全网通版', '1253', 'http://129.204.203.29/hua_s.jpg', '', 8999.00, 8899.00, 800, 80, '{"颜色": "金色", "容量": "256gb", "网络版本": "全网通版"}', 0.19, 1, 1, 2, 200),
--三星手机
( 4, '三星手机 金色 128gb 全网通版', '1623', 'http://129.204.203.29/sumsang.jpg', '', 7999.00, 7899.00, 1000, 100, '{"颜色": "金色", "容量": "128gb", "网络版本": "全网通版"}', 0.19, 1, 1, 1, 300),
( 4, '三星手机 金色 256gb 全网通版', '1723', 'http://129.204.203.29/sumsang.jpg', '', 8999.00, 8899.00, 800, 80, '{"颜色": "金色", "容量": "256gb", "网络版本": "全网通版"}', 0.19, 1, 1, 2, 200),
--荣耀手机
( 5, '荣耀手机 金色 128gb 全网通版', '12833', 'http://129.204.203.29/rong_s.jpg', '', 7999.00, 7899.00, 1000, 100, '{"颜色": "金色", "容量": "128gb", "网络版本": "全网通版"}', 0.19, 1, 1, 1, 300),
( 5, '荣耀手机 金色 256gb 全网通版', '12831', 'http://129.204.203.29/rong_s.jpg', '', 8999.00, 8899.00, 800, 80, '{"颜色": "金色", "容量": "256gb", "网络版本": "全网通版"}', 0.19, 1, 1, 2, 200),

--苹果电脑
( 6, '苹果电脑 金色 16g 256gb', '12193', 'http://129.204.203.29/apple_c.jpg', '', 7999.00, 7899.00, 1000, 100, '{"颜色": "金色", "内存": "16g", "硬盘": "256gb"}', 0.19, 1, 1, 1, 300),
( 6, '苹果电脑 金色 32g 512gb', '1293', 'http://129.204.203.29/apple_c.jpg', '', 8999.00, 8899.00, 800, 80, '{"颜色": "金色", "内存": "32g", "硬盘": "512gb"}', 0.19, 1, 1, 2, 200),
--小米电脑
( 7, '小米电脑 金色 16g 256gb', '10323', 'http://129.204.203.29/xiaomi_c.jpg', '', 7999.00, 7899.00, 1000, 100, '{"颜色": "金色", "内存": "16g", "硬盘": "256gb"}', 0.19, 1, 1, 1, 300),
( 7, '小米电脑 金色 32g 512gb', '1023', 'http://129.204.203.29/xiaomi_c.jpg', '', 8999.00, 8899.00, 800, 80, '{"颜色": "金色", "内存": "32g", "硬盘": "512gb"}', 0.19, 1, 1, 2, 200),
--华为电脑
( 8, '华为电脑 金色 16g 256gb', '1203', 'http://129.204.203.29/hua_c.jpg', '', 7999.00, 7899.00, 1000, 100, '{"颜色": "金色", "内存": "16g", "硬盘": "256gb"}', 0.19, 1, 1, 1, 300),
( 8, '华为电脑 金色 32g 512gb', '12013', 'http://129.204.203.29/hua_c.jpg', '', 8999.00, 8899.00, 800, 80, '{"颜色": "金色", "内存": "32g", "硬盘": "512gb"}', 0.19, 1, 1, 2, 200),

--海澜之家
( 9, '海澜之家 黑色 l', '12543', 'http://129.204.203.29/hailan.jpg', '', 7999.00, 7899.00, 1000, 100, '{"颜色": "黑色", "尺码": "l"}', 0.19, 1, 1, 1, 300),
( 9, '海澜之家 黑色 xl', '12311', 'http://129.204.203.29/hailan.jpg', '', 8999.00, 8899.00, 800, 80, '{"颜色": "黑色", "尺码": "xl"}', 0.19, 1, 1, 2, 200),

--苏泊尔
( 10, '苏泊尔 银色 3l 1500w', '12163', 'http://129.204.203.29/su.jpg', '', 7999.00, 7899.00, 1000, 100, '{"颜色": "黑色", "容量": "3l", "功率": "1500w"}', 0.19, 1, 1, 1, 300),
( 10, '苏泊尔 银色 4l 2000w ', '12113', 'http://129.204.203.29/su.jpg', '', 8999.00, 8899.00, 800, 80, '{"颜色": "黑色", "容量": "4l", "功率": "2000w"}', 0.19, 1, 1, 2, 200),


--格力空调
( 11, '格力空调 银色 2024 1500w', '123231', 'http://129.204.203.29/geli.jpg', '', 7999.00, 7899.00, 1000, 100, '{"颜色": "黑色", "型号": "2024", "功率": "1500w"}', 0.19, 1, 1, 1, 300),
( 11, '格力空调 银色 2025 2000w ', '1143', 'http://129.204.203.29/geli.jpg', '', 8999.00, 8899.00, 800, 80, '{"颜色": "黑色", "型号": "2025", "功率": "2000w"}', 0.19, 1, 1, 2, 200),


--方太燃气灶
( 12, '方太燃气灶 黑色 大 1500w', '12351', 'http://129.204.203.29/fangtai.jpg', '', 7999.00, 7899.00, 1000, 100, '{"颜色": "黑色", "火力": "大", "功率": "1500w"}', 0.19, 1, 1, 1, 300),
( 12, '方太燃气灶 黑色 小 2000w ', '12323', 'http://129.204.203.29/fangtai.jpg', '', 8999.00, 8899.00, 800, 80, '{"颜色": "黑色", "火力": "小", "功率": "2000w"}', 0.19, 1, 1, 2, 200),

--万和热水器
( 13, '万和热水器 黑色 60l 2024 1500w', '12322', 'http://129.204.203.29/wanhe.png', '', 7999.00, 7899.00, 1000, 100, '{"颜色": "黑色", "容量": "60l", "款式": "2024", "功率": "1500w"}', 0.19, 1, 1, 1, 300),
( 13, '万和热水器 银色 100l 2025 3000w ', '12223', 'http://129.204.203.29/wanhe.png', '', 8999.00, 8899.00, 800, 80, '{"颜色": "银色", "容量": "100l", "款式": "2025", "功率": "3000w"}', 0.19, 1, 1, 2, 200),
--万家乐热水器
( 14, '万家乐热水器 黑色 60l 2024 3000w', '12333', 'http://129.204.203.29/wanjiale.jpg', '', 7999.00, 7899.00, 1000, 100, '{"颜色": "黑色", "容量": "60l", "款式": "2024", "功率": "3000w"}', 0.19, 1, 1, 1, 300),
( 14, '万家乐热水器 银色 80l 2026 2000w ', '11123', 'http://129.204.203.29/wanjiale.jpg', '', 8999.00, 8899.00, 800, 80, '{"颜色": "银色", "容量": "80l", "款式": "2026", "功率": "2000w"}', 0.19, 1, 1, 2, 200),
--海尔热水器
( 15, '海尔热水器 黑色 80l 2024 1500w', '12553', 'http://129.204.203.29/haier.jpg', '', 7999.00, 7899.00, 1000, 100, '{"颜色": "黑色", "容量": "80l", "款式": "2024", "功率": "1500w"}', 0.19, 1, 1, 1, 300),
( 15, '海尔热水器 银色 100l 2025 2000w ', '16623', 'http://129.204.203.29/haier.jpg', '', 8999.00, 8899.00, 800, 80, '{"颜色": "银色", "容量": "100l", "款式": "2025", "功率": "2000w"}', 0.19, 1, 1, 2, 200),
--美的热水器
( 16, '美的热水器 黑色 60l 2024 1500w', '12387', 'http://129.204.203.29/meidi.png', '', 7999.00, 7899.00, 1000, 100, '{"颜色": "黑色", "容量": "60l", "款式": "2024", "功率": "1500w"}', 0.19, 1, 1, 1, 300),
( 16, '美的热水器 银色 100l 2026 3000w ', '18231', 'http://129.204.203.29/meidi.png', '', 8999.00, 8899.00, 800, 80, '{"颜色": "银色", "容量": "100l", "款式": "2026", "功率": "3000w"}', 0.19, 1, 1, 2, 200);

-- 商品规格表建语句
drop table if exists pms_product_spec;
create table pms_product_spec
(
    id          bigserial primary key,
    category_id bigint                                not null,
    name        varchar                               not null,
    sort        integer                               not null,
    status      integer     default 1                 not null,
    create_by   varchar     default ''                not null,
    create_time timestamptz default current_timestamp not null,
    update_by   varchar     default ''                not null,
    update_time timestamptz                           null,
    is_deleted  integer     default 1                 not null
);

-- 添加商品规格表注释
comment on table pms_product_spec is '商品规格表';

-- 添加商品规格表列注释
comment on column pms_product_spec.id is '';
comment on column pms_product_spec.category_id is '分类id';
comment on column pms_product_spec.name is '规格名称';
comment on column pms_product_spec.sort is '排序';
comment on column pms_product_spec.status is '状态：0->禁用；1->启用';
comment on column pms_product_spec.create_by is '创建人id';
comment on column pms_product_spec.create_time is '创建时间';
comment on column pms_product_spec.update_by is '更新人id';
comment on column pms_product_spec.update_time is '更新时间';
comment on column pms_product_spec.is_deleted is '是否删除(0:否,1:是)';

insert into pms_product_spec (category_id, name, sort)
values
-- 手机规格(category_id=1)
(1, '颜色', 1),
(1, '容量', 2),
(1, '网络版本', 3),
-- 小米手机规格(category_id=8)
(8, '颜色', 1),
(8, '容量', 2),
(8, '网络版本', 3),
-- 苹果手机规格(category_id=9)
(9, '颜色', 1),
(9, '容量', 2),
(9, '网络版本', 3),
-- 华为手机规格(category_id=10)
(10, '颜色', 1),
(10, '容量', 2),
(10, '网络版本', 3),
-- 三星手机规格(category_id=11)
(11, '颜色', 1),
(11, '容量', 2),
(11, '网络版本', 3),
-- 荣耀手机规格(category_id=12)
(12, '颜色', 1),
(12, '容量', 2),
(12, '网络版本', 3),

-- 电脑规格(category_id=2)
(2, '颜色', 1),
(2, '内存', 2),
(2, '硬盘', 3),
-- 苹果电脑规格(category_id=13)
(13, '颜色', 1),
(13, '内存', 2),
(13, '硬盘', 3),
-- 小米电脑规格(category_id=14)
(14, '颜色', 1),
(14, '内存', 2),
(14, '硬盘', 3),
-- 华为电脑规格(category_id=15)
(15, '颜色', 1),
(15, '内存', 2),
(15, '硬盘', 3),

-- 服装规格(category_id=3)
(3, '颜色', 1),
(3, '尺码', 2),
-- 海澜之家规格(category_id=16)
(16, '颜色', 1),
(16, '尺码', 2),


-- 电器规格(category_id=4)
(4, '颜色', 1),
(4, '容量', 2),
(4, '功率', 3),
-- 电饭锅规格(category_id=17)
(17, '颜色', 1),
(17, '容量', 2),
(17, '功率', 3),

-- 空调规格(category_id=5)
(5, '颜色', 1),
(5, '型号', 2),
(5, '功率', 3),
-- 格力空调规格(category_id=18)
(18, '颜色', 1),
(18, '型号', 2),
(18, '功率', 3),

-- 厨房规格(category_id=6)
(6, '颜色', 1),
(6, '火力', 2),
(6, '功率', 3),
-- 方太规格(category_id=19)
(19, '颜色', 1),
(19, '火力', 2),
(19, '功率', 3),

-- 热水器规格(category_id=7)
(7, '颜色', 1),
(7, '容量', 2),
(7, '款式', 3),
(7, '功率', 4),
-- 万和热水器规格(category_id=20)
(20, '颜色', 1),
(20, '容量', 2),
(20, '款式', 3),
(20, '功率', 4),
-- 万家乐热水器规格(category_id=21)
(21, '颜色', 1),
(21, '容量', 2),
(21, '款式', 3),
(21, '功率', 4),
-- 海尔热水器规格(category_id=22)
(22, '颜色', 1),
(22, '容量', 2),
(22, '款式', 3),
(22, '功率', 4),
-- 美的热水器规格(category_id=23)
(23, '颜色', 1),
(23, '容量', 2),
(23, '款式', 3),
(23, '功率', 4);
-- 商品规格值表建语句
drop table if exists pms_product_spec_value;
create table pms_product_spec_value
(
    id          bigserial primary key,
    spec_id     bigint                                not null,
    value       varchar                               not null,
    sort        integer                               not null,
    status      integer     default 1                 not null,
    create_by   varchar     default ''                not null,
    create_time timestamptz default current_timestamp not null,
    update_by   varchar     default ''                not null,
    update_time timestamptz                           null,
    is_deleted  integer     default 1                 not null
);

-- 添加商品规格值表注释
comment on table pms_product_spec_value is '商品规格值表';

-- 添加商品规格值表列注释
comment on column pms_product_spec_value.id is '';
comment on column pms_product_spec_value.spec_id is '规格id';
comment on column pms_product_spec_value.value is '规格值';
comment on column pms_product_spec_value.sort is '排序';
comment on column pms_product_spec_value.status is '状态：0->禁用；1->启用';
comment on column pms_product_spec_value.create_by is '创建人id';
comment on column pms_product_spec_value.create_time is '创建时间';
comment on column pms_product_spec_value.update_by is '更新人id';
comment on column pms_product_spec_value.update_time is '更新时间';
comment on column pms_product_spec_value.is_deleted is '是否删除(0:否,1:是)';

-- 插入商品规格值数据
insert into pms_product_spec_value (spec_id, value, sort)
values
-- 手机颜色规格值
(1, '暗紫色', 1),
(1, '自然色', 2),
(1, '银色', 3),
(1, '金色', 4),
-- 手机容量规格值
(2, '128gb', 1),
(2, '256gb', 2),
(2, '512gb', 3),
(2, '1tb', 4),
-- 手机网络版本规格值
(3, '全网通版', 1),
(3, '5g版', 2),

-- 小米手机颜色规格值
(4, '暗紫色', 1),
(4, '自然色', 2),
(4, '银色', 3),
(4, '金色', 4),
-- 小米手机容量规格值
(5, '128gb', 1),
(5, '256gb', 2),
(5, '512gb', 3),
(5, '1tb', 4),
-- 小米手机网络版本规格值
(6, '全网通版', 1),
(6, '5g版', 2),

-- 苹果手机颜色规格值
(7, '暗紫色', 1),
(7, '自然色', 2),
(7, '银色', 3),
(7, '金色', 4),
-- 苹果手机容量规格值
(8, '128gb', 1),
(8, '256gb', 2),
(8, '512gb', 3),
(8, '1tb', 4),
-- 苹果手机网络版本规格值
(9, '全网通版', 1),
(9, '5g版', 2),

-- 华为手机颜色规格值
(10, '暗紫色', 1),
(10, '自然色', 2),
(10, '银色', 3),
(10, '金色', 4),
-- 华为手机容量规格值
(11, '128gb', 1),
(11, '256gb', 2),
(11, '512gb', 3),
(11, '1tb', 4),
-- 华为手机网络版本规格值
(12, '全网通版', 1),
(12, '5g版', 2),

-- 三星手机颜色规格值
(13, '暗紫色', 1),
(13, '自然色', 2),
(13, '银色', 3),
(13, '金色', 4),
-- 三星手机容量规格值
(14, '128gb', 1),
(14, '256gb', 2),
(14, '512gb', 3),
(14, '1tb', 4),
-- 三星网络版本规格值
(15, '全网通版', 1),
(15, '5g版', 2),

-- 荣耀手机颜色规格值
(16, '暗紫色', 1),
(16, '自然色', 2),
(16, '银色', 3),
(16, '金色', 4),
-- 荣耀手机容量规格值
(17, '128gb', 1),
(17, '256gb', 2),
(17, '512gb', 3),
(17, '1tb', 4),
-- 荣耀手机网络版本规格值
(18, '全网通版', 1),
(18, '5g版', 2),

-- 电脑颜色规格值
(19, '深空灰', 1),
(19, '银色', 2),
(19, '黑色', 3),
-- 电脑内存规格值
(20, '8gb', 1),
(20, '16gb', 2),
(20, '32gb', 3),
-- 电脑硬盘规格值
(21, '256gb', 1),
(21, '512gb', 2),
(21, '1tb', 3),

-- 苹果电脑颜色规格值
(22, '深空灰', 1),
(22, '银色', 2),
(22, '黑色', 3),
-- 苹果电脑内存规格值
(23, '8gb', 1),
(23, '16gb', 2),
(23, '32gb', 3),
-- 苹果电脑硬盘规格值
(24, '256gb', 1),
(24, '512gb', 2),
(24, '1tb', 3),

-- 小米电脑颜色规格值
(25, '深空灰', 1),
(25, '银色', 2),
(25, '黑色', 3),
-- 小米电脑内存规格值
(26, '8gb', 1),
(26, '16gb', 2),
(26, '32gb', 3),
-- 小米电脑硬盘规格值
(27, '256gb', 1),
(27, '512gb', 2),
(27, '1tb', 3),

-- 华为电脑颜色规格值
(28, '深空灰', 1),
(28, '银色', 2),
(28, '黑色', 3),
-- 华为电脑内存规格值
(29, '8gb', 1),
(29, '16gb', 2),
(29, '32gb', 3),
-- 华为电脑硬盘规格值
(30, '256gb', 1),
(30, '512gb', 2),
(30, '1tb', 3),

-- 服装颜色规格值
(31, '黑色', 1),
(31, '白色', 2),
(31, '灰色', 3),
(31, '藏青色', 4),
-- 服装尺码规格值
(32, 's', 1),
(32, 'm', 2),
(32, 'l', 3),
(32, 'xl', 4),
(32, 'xxl', 5),

-- 海澜之家服装颜色规格值
(33, '黑色', 1),
(33, '白色', 2),
(33, '灰色', 3),
(33, '藏青色', 4),
-- 海澜之家服装尺码规格值
(34, 's', 1),
(34, 'm', 2),
(34, 'l', 3),
(34, 'xl', 4),
(34, 'xxl', 5),

-- 电器颜色规格值
(35, '白色', 1),
(35, '银色', 2),
(35, '黑色', 3),
-- 电器容量规格值
(36, '3l', 1),
(36, '4l', 2),
(36, '5l', 3),
-- 电器功率规格值
(37, '1500w', 1),
(37, '2000w', 2),
(37, '3000w', 3),

-- 电饭锅颜色规格值
(38, '白色', 1),
(38, '银色', 2),
(38, '黑色', 3),
-- 电饭锅容量规格值
(39, '3l', 1),
(39, '4l', 2),
(39, '5l', 3),
-- 电饭锅功率规格值
(40, '1500w', 1),
(40, '2000w', 2),
(40, '3000w', 3),

-- 空调颜色规格值
(41, '白色', 1),
(41, '银色', 2),
(41, '黑色', 3),
-- 空调型号规格值
(42, '2024', 1),
(42, '2025', 2),
(42, '2026', 3),
-- 空调功率规格值
(43, '1500w', 1),
(43, '2000w', 2),
(43, '3000w', 3),

-- 格力空调颜色规格值
(44, '白色', 1),
(44, '银色', 2),
(44, '黑色', 3),
-- 格力空调型号规格值
(45, '2024', 1),
(45, '2025', 2),
(45, '2026', 3),
-- 格力空调功率规格值
(46, '1500w', 1),
(46, '2000w', 2),
(46, '3000w', 3),


-- 厨房颜色规格值
(47, '白色', 1),
(47, '银色', 2),
(47, '黑色', 3),
-- 厨房火力规格值
(48, '大', 1),
(48, '中', 2),
(48, '小', 3),
-- 厨房功率规格值
(49, '1500w', 1),
(49, '2000w', 2),
(49, '3000w', 3),

-- 方太颜色规格值
(50, '白色', 1),
(50, '银色', 2),
(50, '黑色', 3),
-- 方太火力规格值
(51, '大', 1),
(51, '中', 2),
(51, '小', 3),
-- 方太功率规格值
(52, '1500w', 1),
(52, '2000w', 2),
(52, '3000w', 3),

-- 热水器颜色规格值
(53, '白色', 1),
(53, '银色', 2),
(53, '黑色', 3),
-- 热水器容量规格值
(54, '60l', 1),
(54, '80l', 2),
(54, '100l', 3),
-- 热水器款式规格值
(55, '2024', 1),
(55, '2025', 2),
(55, '2026', 3),
-- 热水器功率规格值
(56, '1500w', 1),
(56, '2000w', 2),
(56, '3000w', 3),

-- 万和热水器颜色规格值
(57, '白色', 1),
(57, '银色', 2),
(57, '黑色', 3),
-- 万和热水器容量规格值
(58, '60l', 1),
(58, '80l', 2),
(58, '100l', 3),
-- 万和 热水器款式规格值
(59, '2024', 1),
(59, '2025', 2),
(59, '2026', 3),
-- 万和热水器功率规格值
(60, '1500w', 1),
(60, '2000w', 2),
(60, '3000w', 3),

-- 万家乐热水器颜色规格值
(61, '白色', 1),
(61, '银色', 2),
(61, '黑色', 3),
-- 万家乐热水器容量规格值
(62, '60l', 1),
(62, '80l', 2),
(62, '100l', 3),
-- 万家乐热水器款式规格值
(63, '2024', 1),
(63, '2025', 2),
(63, '2026', 3),
-- 万家乐热水器功率规格值
(64, '1500w', 1),
(64, '2000w', 2),
(64, '3000w', 3),

-- 海尔热水器颜色规格值
(65, '白色', 1),
(65, '银色', 2),
(65, '黑色', 3),
-- 海尔热水器容量规格值
(66, '60l', 1),
(66, '80l', 2),
(66, '100l', 3),
-- 海尔热水器款式规格值
(67, '2024', 1),
(67, '2025', 2),
(67, '2026', 3),
-- 海尔热水器功率规格值
(68, '1500w', 1),
(68, '2000w', 2),
(68, '3000w', 3),

-- 美的热水器颜色规格值
(69, '白色', 1),
(69, '银色', 2),
(69, '黑色', 3),
-- 美的热水器容量规格值
(70, '60l', 1),
(70, '80l', 2),
(70, '100l', 3),
-- 美的热水器款式规格值
(71, '2024', 1),
(71, '2025', 2),
(71, '2026', 3),
-- 美的热水器功率规格值
(72, '1500w', 1),
(72, '2000w', 2),
(72, '3000w', 3);

-- 商品spu表建语句
drop table if exists pms_product_spu;
create table pms_product_spu
(
    id                    bigserial primary key,
    name                  varchar                               not null,
    sub_title              varchar                               not null,
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
    publish_status        integer     default 1                 not null,
    new_status            integer     default 1                 not null,
    recommend_status      integer     default 1                 not null,
    verify_status         integer     default 1                 not null,
    preview_status        integer     default 1                 not null,
    sort                  integer                               not null,
    new_status_sort       integer                               not null,
    recommend_status_sort integer                               not null,
    sales                 integer                               not null,
    stock                 integer                               not null,
    low_stock             integer                               not null,
    promotion_type        integer     default 1                 not null,
    detail_html           text                                  not null,
    detail_mobile_html    text                                  not null,
    create_by             varchar     default ''                not null,
    create_time           timestamptz default current_timestamp not null,
    update_by             varchar     default ''                not null,
    update_time           timestamptz                           null,
    is_deleted            integer     default 1                 not null
);

-- 添加商品spu表注释
comment on table pms_product_spu is '商品spu表';

-- 添加商品spu表列注释
comment on column pms_product_spu.id is '商品spuid';
comment on column pms_product_spu.name is '商品名称';
comment on column pms_product_spu.sub_title is '副标题';
comment on column pms_product_spu.product_sn is '商品货号';
comment on column pms_product_spu.category_id is '商品分类id';
comment on column pms_product_spu.category_ids is '商品分类id集合';
comment on column pms_product_spu.category_name is '商品分类名称';
comment on column pms_product_spu.brand_id is '品牌id';
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
comment on column pms_product_spu.create_by is '创建人id';
comment on column pms_product_spu.create_time is '创建时间';
comment on column pms_product_spu.update_by is '更新人id';
comment on column pms_product_spu.update_time is '更新时间';
comment on column pms_product_spu.is_deleted is '是否删除(0:否,1:是)';

insert into pms_product_spu (id, name, sub_title, product_sn, category_id, category_ids, category_name, brand_id, brand_name, unit, weight, keywords, album_pics, main_pic, price_range, publish_status, new_status, recommend_status, verify_status, preview_status, sort, new_status_sort, recommend_status_sort, sales, stock, low_stock, promotion_type, detail_html, detail_mobile_html)
values  (1, '小米（mi）redmi note15 pro 天玑7400-ultra 7000mah 龙晶玻璃十倍抗摔 ip68 8+256 子夜黑 红米 5g手机', '新年狂欢购，限时特惠价，品质好物带回家', 'fsef', 8, '', '小米手机', 9, 'test', '台', 0.19, '小米', 'http://129.204.203.29/xiaomi_s.jpg', 'http://129.204.203.29/xiaomi_s.jpg', '7999-12999', 1, 1, 1, 1, 0, 1, 1, 1, 1000, 0, 0, 0, 'test', 'test'),
        (2, 'apple/苹果 iphone 17 256gb 薰衣草紫色 支持移动联通电信5g 双卡双待手机', '春节大促销，满减送好礼，实惠享不停', 'fsef', 9, '', '苹果手机', 1, 'test', '台', 0.19, 'apple', 'http://129.204.203.29/apple_s.jpg', 'http://129.204.203.29/apple_s.jpg', '7999-12999', 1, 1, 1, 1, 0, 1, 1, 1, 1000, 0, 0, 0, 'test', 'test'),
        (3, '华为（huawei）mate x5典藏版x3折叠屏手机全网通正品特北斗卫星华为大折叠 青山黛【mate x3】 【12g+256g】 赠运费险|详情咨询客服', '年货节来袭，超值折扣季，好货提前囤', 'fsef', 10, '', '华为手机', 5, 'test', '台', 0.19, '华为', 'http://129.204.203.29/hua_s.jpg', 'http://129.204.203.29/hua_s.jpg', '7999-12999', 1, 1, 1, 1, 0, 1, 1, 1, 1000, 0, 0, 0, 'test', 'test'),
        (4, '三星（samsung）w25 心系天下 折叠屏手机2亿像素 galaxy ai新品商务高端智能手机 陶瓷黑 16gb+1tb 现货速发|国行正品', '新春钜惠节，买一送一活动，惊喜连连', 'fsef', 11, '', '三星手机', 7, 'test', '台', 0.19, '三星', 'http://129.204.203.29/sumsang.jpg', 'http://129.204.203.29/sumsang.jpg', '7999-12999', 1, 1, 1, 1, 0, 1, 1, 1, 1000, 0, 0, 0, 'test', 'test'),
        (5, '荣耀x70 8+128 朱砂红 金标十面抗摔 8300mah青海湖电池 ip69防水 荣耀绿洲护眼屏 ai手机 国家补贴', '新年新气象，商品大降价，优惠享不停', 'fsef', 12, '', '荣耀手机', 5, 'test', '台', 0.19, '荣耀x70', 'http://129.204.203.29/rong_s.jpg', 'http://129.204.203.29/rong_s.jpg', '7999-12999', 1, 1, 1, 1, 0, 1, 1, 1, 1000, 0, 0, 0, 'test', 'test'),
        (6, '苹果（apple）macbook pro air 超薄商务办公 剪辑设计学生游戏99新13寸pro丨i5+8g+512g固态丨', '春节特卖会，精选好物五折起，不容错过', 'fsef', 13, '', '苹果电脑', 1, 'test', '台', 0.19, 'macbook pro air', 'http://129.204.203.29/apple_c.jpg', 'http://129.204.203.29/apple_c.jpg', '7999-12999', 1, 1, 1, 1, 0, 1, 1, 1, 1000, 0, 0, 0, 'test', 'test'),
        (7, '小米笔记本电脑 红米 redmi book 16 酷睿i5标压 16英寸 1tb办公学生轻薄本', '岁末大清仓，年终特价优惠，省钱购物', 'fsef', 14, '', '小米电脑', 9, 'test', '台', 0.19, '红米 redmi book 16', 'http://129.204.203.29/xiaomi_c.jpg', 'http://129.204.203.29/xiaomi_c.jpg', '7999-12999', 1, 1, 1, 1, 0, 1, 1, 1, 1000, 0, 0, 0, 'test', 'test'),
        (8, '华为（huawei）展机matebookxpro轻薄旗舰酷睿ultra9微绒典藏版商务办公电脑 热销新款 ultra7 32g 1t 触屏', '新年开门红，限时秒杀活动，低至三折', 'fsef', 15, '', '华为电脑', 5, 'test', '台', 0.19, '华为（huawei）展机matebookxpro', 'http://129.204.203.29/hua_c.jpg', 'http://129.204.203.29/hua_c.jpg', '7999-12999', 1, 1, 1, 1, 0, 1, 1, 1, 1000, 0, 0, 0, 'test', 'test'),
        (9, 'hla海澜之家长袖针织衫男轻商务时尚系列圆领毛衣冬季男', '年味十足购，年货大促销，温馨过新年', 'fsef', 16, '', '海澜之家', 4, 'test', '台', 0.19, 'hla海澜之家', 'http://129.204.203.29/hailan.jpg', 'http://129.204.203.29/hailan.jpg', '7999-12999', 1, 1, 1, 1, 0, 1, 1, 1, 1000, 0, 0, 0, 'test', 'test'),
        (10, '苏泊尔supor好帮手铝合金压力锅4.5l带蒸格20cm高压锅燃气专用', '新春购物节，满额立减优惠，实惠多多', 'fsef', 17, '', '电饭锅', 1, 'test', '台', 0.19, '苏泊尔supor', 'http://129.204.203.29/su.jpg', 'http://129.204.203.29/su.jpg', '7999-12999', 1, 1, 1, 1, 0, 1, 1, 1, 1000, 0, 0, 0, 'test', 'test'),
        (11, '格力出品 晶弘空调 小凉神 大1匹 新一级能效变频 纯铜管卧室省电挂机 国家补贴kfr-26gw/jhfnhaa1bj', '新年新体验，品质商品特价，享受生活', 'fsef', 18, '', '格力空调', 3, 'test', '台', 0.19, '格力出品', 'http://129.204.203.29/geli.jpg', 'http://129.204.203.29/geli.jpg', '7999-12999', 1, 1, 1, 1, 0, 1, 1, 1, 1000, 0, 0, 0, 'test', 'test'),
        (12, '方太燃气灶天然气 家用嵌入式5.2kw* 猛火燃气灶 易清洁可调节 tx22', '春节福利季，超值好物优惠，全家共享', 'fsef', 19, '', '方太', 2, 'test', '台', 0.19, '方太燃气灶天然气', 'http://129.204.203.29/fangtai.jpg', 'http://129.204.203.29/fangtai.jpg', '7999-12999', 1, 1, 1, 1, 0, 1, 1, 1, 1000, 0, 0, 0, 'test', 'test'),
        (13, '万和【t2】40升2000w速热省电双盾防护专利防电墙租房首选大水量家用储水式电热水器', '年货大集会，折扣狂欢节，欢欢喜喜过大年', 'fsef', 20, '', '万和', 8, 'test', '台', 0.19, '万和', 'http://129.204.203.29/wanhe.png', 'http://129.204.203.29/wanhe.png', '7999-12999', 1, 1, 1, 1, 0, 1, 1, 1, 1000, 0, 0, 0, 'test', 'test'),
        (14, '万家乐【安睡洗s9·四驱增压水伺服】16升天然气超一级恒温一级静音零冷感防冻以旧换新s9max燃气热水器', '新年大放价，限时抢购活动，先到先得', 'fsef', 21, '', '万家乐', 8, 'test', '台', 0.19, '万家乐', 'http://129.204.203.29/wanjiale.jpg', 'http://129.204.203.29/wanjiale.jpg', '7999-12999', 1, 1, 1, 1, 0, 1, 1, 1, 1000, 0, 0, 0, 'test', 'test'),
        (15, '海尔出品统帅60升电热水器京东自营上门安装 国家补贴 2200w节能速热一级能效家用储水式lec6001-ld5金', '新春嘉年华，特价商品琳琅满目，任君挑选', 'fsef', 22, '', '海尔', 3, 'test', '台', 0.19, '海尔出品统帅', 'http://129.204.203.29/haier.jpg', 'http://129.204.203.29/haier.jpg', '7999-12999', 1, 1, 1, 1, 0, 1, 1, 1, 1000, 0, 0, 0, 'test', 'test'),
        (16, '美的（midea）【整机8年质保】60升电热水器2000w家用出租屋节能抑菌安全防电墙以旧换新门店同款f60-20f1(h)', '春节大回馈，买就送好礼，购物更划算', 'fsef', 23, '', '美的', 3, 'test', '台', 0.19, '美的', 'http://129.204.203.29/meidi.png', 'http://129.204.203.29/meidi.png', '7999-12999', 1, 1, 1, 1, 0, 1, 1, 1, 1000, 0, 0, 0, 'test', 'test');


-- 购物车表建语句
drop table if exists oms_cart_item;
create table oms_cart_item
(
    id                  bigserial primary key,
    member_id           bigint                                not null,
    product_id          bigint                                not null,
    product_sku_id      bigint                                not null,
    quantity            integer                               not null,
    price               numeric                               not null,
    selected            integer     default 1                 not null,
    product_name        varchar                               not null,
    product_sub_title   varchar                               not null,
    product_pic         text                                  not null,
    product_sku_code    varchar                               not null,
    product_sn          varchar                               not null,
    product_brand       varchar                               not null,
    product_category_id bigint                                not null,
    product_attr        jsonb                                 not null,
    member_nickname     varchar                               not null,
    source              integer     default 1                 not null,
    delete_status       integer     default 1                 not null,
    expire_time         timestamptz                           not null,
    create_time         timestamptz default current_timestamp not null,
    update_time         timestamptz                           null
);

-- 添加购物车表注释
comment on table oms_cart_item is '购物车表';

-- 添加购物车表列注释
comment on column oms_cart_item.id is '主键id';
comment on column oms_cart_item.member_id is '会员id';
comment on column oms_cart_item.product_id is '商品id';
comment on column oms_cart_item.product_sku_id is '商品sku id';
comment on column oms_cart_item.quantity is '购买数量';
comment on column oms_cart_item.price is '添加到购物车时的价格';
comment on column oms_cart_item.selected is '是否选中 0-未选中 1-选中';
comment on column oms_cart_item.product_name is '商品名称';
comment on column oms_cart_item.product_sub_title is '商品副标题';
comment on column oms_cart_item.product_pic is '商品主图url';
comment on column oms_cart_item.product_sku_code is '商品sku编码';
comment on column oms_cart_item.product_sn is '商品货号';
comment on column oms_cart_item.product_brand is '商品品牌';
comment on column oms_cart_item.product_category_id is '商品分类id';
comment on column oms_cart_item.product_attr is '商品销售属性json';
comment on column oms_cart_item.member_nickname is '会员昵称';
comment on column oms_cart_item.source is '来源 1-pc 2-h5 3-小程序 4-app';
comment on column oms_cart_item.delete_status is '删除状态 0-正常 1-删除';
comment on column oms_cart_item.expire_time is '过期时间';
comment on column oms_cart_item.create_time is '创建时间';
comment on column oms_cart_item.update_time is '更新时间';

-- 公司收发货地址表建语句
drop table if exists oms_company_address;
create table oms_company_address
(
    id             bigserial primary key,
    address_name   varchar                               not null,
    name           varchar                               not null,
    phone          varchar                               not null,
    province       varchar                               not null,
    city           varchar                               not null,
    region         varchar                               not null,
    detail_address varchar                               not null,
    send_status    integer     default 1                 not null,
    receive_status integer     default 1                 not null,
    create_by      varchar     default ''                not null,
    create_time    timestamptz default current_timestamp not null,
    update_by      varchar     default ''                not null,
    update_time    timestamptz                           null,
    is_deleted     integer     default 0                 not null
);

-- 添加公司收发货地址表注释
comment on table oms_company_address is '公司收发货地址表';

-- 添加公司收发货地址表列注释
comment on column oms_company_address.id is '主键id';
comment on column oms_company_address.address_name is '地址名称';
comment on column oms_company_address.name is '收发货人姓名';
comment on column oms_company_address.phone is '收货人电话';
comment on column oms_company_address.province is '省/直辖市';
comment on column oms_company_address.city is '市';
comment on column oms_company_address.region is '区';
comment on column oms_company_address.detail_address is '详细地址';
comment on column oms_company_address.send_status is '默认发货地址：0->否；1->是';
comment on column oms_company_address.receive_status is '默认收货地址：0->否；1->是';
comment on column oms_company_address.create_by is '创建人id';
comment on column oms_company_address.create_time is '创建时间';
comment on column oms_company_address.update_by is '更新人id';
comment on column oms_company_address.update_time is '更新时间';
comment on column oms_company_address.is_deleted is '是否删除(0:否,1:是)';

insert into oms_company_address (id, address_name, name, phone, province, city, region, detail_address, send_status, receive_status, create_by)
values (1, '总部地址', '张三', '13800138000', '北京市', '北京市', '朝阳区', '建国路88号', 1, 0, 1),
       (2, '分公司地址', '李四', '13900139000', '上海市', '上海市', '浦东新区', '世纪大道100号', 0, 1, 2),
       (3, '仓库地址', '王五', '13700137000', '广东省', '广州市', '天河区', '体育西路200号', 1, 1, 3),
       (4, '配送中心', '赵六', '13600136000', '浙江省', '杭州市', '西湖区', '文三路300号', 0, 0, 4),
       (5, '备用地址', '孙七', '13500135000', '江苏省', '南京市', '玄武区', '中山路400号', 0, 0, 5);

-- 订单表建语句
drop table if exists oms_order;
create table oms_order
(
    id                      bigserial primary key,
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
    pay_type                integer     default 1                 not null,
    source_type             integer     default 1                 not null,
    status                  integer     default 1                 not null,
    order_type              integer     default 1                 not null,
    delivery_company        varchar                               not null,
    delivery_sn             varchar                               not null,
    auto_confirm_day        integer                               not null,
    integration             integer                               not null,
    growth                  integer                               not null,
    promotion_info          varchar                               not null,
    bill_type               integer     default 1                 not null,
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
    confirm_status          integer     default 1                 not null,
    delete_status           integer     default 1                 not null,
    use_integration         integer                               not null,
    payment_time            timestamptz                           ,
    delivery_time           timestamptz                           ,
    receive_time            timestamptz                           ,
    comment_time            timestamptz                           ,
    modify_time             timestamptz
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
comment on column oms_order.source_type is '订单来源：0->pc订单；1->app订单';
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

insert into oms_order (id, member_id, coupon_id, order_sn, create_time, member_username, total_amount, pay_amount, freight_amount, promotion_amount, integration_amount, coupon_amount, discount_amount, pay_type, source_type, status, order_type, delivery_company, delivery_sn, auto_confirm_day, integration, growth, promotion_info, bill_type, bill_header, bill_content, bill_receiver_phone, bill_receiver_email, receiver_name, receiver_phone, receiver_post_code, receiver_province, receiver_city, receiver_region, receiver_detail_address, note, confirm_status, delete_status, use_integration) values (12, 1, 2, '201809150101000001', '2025-01-23 14:50:23', 'test', 18732, 16378, 20, 2344, 0, 10, 10, 0, 1, 4, 0, '顺丰快递', '201707196398345', 15, 13284, 13284, '单品促销,打折优惠：满3件，打7.50折,满减优惠：满1000.00元，减120.00元,满减优惠：满1000.00元，减120.00元,无优惠', 1, '1111', '1111', '18613030352', '1002219331@qq.com', '大梨', '18033441849', '518000', '江苏省', '常州市', '天宁区', '东晓街道', '111', 0, 0, 1000);
insert into oms_order (id, member_id, coupon_id, order_sn, create_time, member_username, total_amount, pay_amount, freight_amount, promotion_amount, integration_amount, coupon_amount, discount_amount, pay_type, source_type, status, order_type, delivery_company, delivery_sn, auto_confirm_day, integration, growth, promotion_info, bill_type, bill_header, bill_content, bill_receiver_phone, bill_receiver_email, receiver_name, receiver_phone, receiver_post_code, receiver_province, receiver_city, receiver_region, receiver_detail_address, note, confirm_status, delete_status, use_integration) values (13, 1, 2, '201809150102000002', '2025-01-23 14:50:23', 'test', 18732, 16378, 0, 2344, 0, 10, 0, 1, 1, 1, 0, '顺丰快递', '201707196398345', 15, 13284, 13284, '单品促销,打折优惠：满3件，打7.50折,满减优惠：满1000.00元，减120.00元,满减优惠：满1000.00元，减120.00元,无优惠', 1, '1111', '1111', '18613030352', '1002219331@qq.com', '大梨', '18033441849', '518000', '广东省', '深圳市', '福田区', '东晓街道', '111', 0, 0, 1000);
insert into oms_order (id, member_id, coupon_id, order_sn, create_time, member_username, total_amount, pay_amount, freight_amount, promotion_amount, integration_amount, coupon_amount, discount_amount, pay_type, source_type, status, order_type, delivery_company, delivery_sn, auto_confirm_day, integration, growth, promotion_info, bill_type, bill_header, bill_content, bill_receiver_phone, bill_receiver_email, receiver_name, receiver_phone, receiver_post_code, receiver_province, receiver_city, receiver_region, receiver_detail_address, note, confirm_status, delete_status, use_integration) values (14, 1, 2, '201809130101000001', '2025-01-23 14:50:23', 'test', 18732, 16378, 0, 2344, 0, 10, 0, 2, 1, 2, 0, '顺丰快递', '201707196398345', 15, 13284, 13284, '单品促销,打折优惠：满3件，打7.50折,满减优惠：满1000.00元，减120.00元,满减优惠：满1000.00元，减120.00元,无优惠', 1, '1111', '1111', '18613030352', '1002219331@qq.com', '大梨', '18033441849', '518000', '广东省', '深圳市', '福田区', '东晓街道', '111', 0, 0, 1000);
insert into oms_order (id, member_id, coupon_id, order_sn, create_time, member_username, total_amount, pay_amount, freight_amount, promotion_amount, integration_amount, coupon_amount, discount_amount, pay_type, source_type, status, order_type, delivery_company, delivery_sn, auto_confirm_day, integration, growth, promotion_info, bill_type, bill_header, bill_content, bill_receiver_phone, bill_receiver_email, receiver_name, receiver_phone, receiver_post_code, receiver_province, receiver_city, receiver_region, receiver_detail_address, note, confirm_status, delete_status, use_integration) values (15, 1, 2, '201809130102000002', '2025-01-23 14:50:23', 'test', 18732, 16378, 0, 2344, 0, 10, 0, 1, 1, 3, 0, '顺丰快递', '201707196398345', 15, 13284, 13284, '单品促销,打折优惠：满3件，打7.50折,满减优惠：满1000.00元，减120.00元,满减优惠：满1000.00元，减120.00元,无优惠', 1, '1111', '1111', '18613030352', '1002219331@qq.com', '大梨', '18033441849', '518000', '广东省', '深圳市', '福田区', '东晓街道', '111', 1, 0, 1000);
insert into oms_order (id, member_id, coupon_id, order_sn, create_time, member_username, total_amount, pay_amount, freight_amount, promotion_amount, integration_amount, coupon_amount, discount_amount, pay_type, source_type, status, order_type, delivery_company, delivery_sn, auto_confirm_day, integration, growth, promotion_info, bill_type, bill_header, bill_content, bill_receiver_phone, bill_receiver_email, receiver_name, receiver_phone, receiver_post_code, receiver_province, receiver_city, receiver_region, receiver_detail_address, note, confirm_status, delete_status, use_integration) values (16, 1, 2, '201809140101000001', '2025-01-23 14:50:23', 'test', 18732, 16378, 0, 2344, 0, 10, 0, 2, 1, 4, 0, '顺丰快递', '201707196398345', 15, 13284, 13284, '单品促销,打折优惠：满3件，打7.50折,满减优惠：满1000.00元，减120.00元,满减优惠：满1000.00元，减120.00元,无优惠', 1, '1111', '1111', '18613030352', '1002219331@qq.com', '大梨', '18033441849', '518000', '广东省', '深圳市', '福田区', '东晓街道', '111', 0, 0, 1000);

insert into oms_order (id, member_id, coupon_id, order_sn, create_time, member_username, total_amount, pay_amount, freight_amount, promotion_amount, integration_amount, coupon_amount, discount_amount, pay_type, source_type, status, order_type, delivery_company, delivery_sn, auto_confirm_day, integration, growth, promotion_info, bill_type, bill_header, bill_content, bill_receiver_phone, bill_receiver_email, receiver_name, receiver_phone, receiver_post_code, receiver_province, receiver_city, receiver_region, receiver_detail_address, note, confirm_status, delete_status, use_integration) values (17, 2, 2, '201809150101000001', '2025-01-23 14:50:23', 'koobe', 18732, 16378, 20, 2344, 0, 10, 10, 0, 1, 4, 0, '顺丰快递', '201707196398345', 15, 13284, 13284, '单品促销,打折优惠：满3件，打7.50折,满减优惠：满1000.00元，减120.00元,满减优惠：满1000.00元，减120.00元,无优惠', 1, '1111', '1111', '18613030352', '1002219331@qq.com', '大梨', '18033441849', '518000', '江苏省', '常州市', '天宁区', '东晓街道', '111', 0, 0, 1000);
insert into oms_order (id, member_id, coupon_id, order_sn, create_time, member_username, total_amount, pay_amount, freight_amount, promotion_amount, integration_amount, coupon_amount, discount_amount, pay_type, source_type, status, order_type, delivery_company, delivery_sn, auto_confirm_day, integration, growth, promotion_info, bill_type, bill_header, bill_content, bill_receiver_phone, bill_receiver_email, receiver_name, receiver_phone, receiver_post_code, receiver_province, receiver_city, receiver_region, receiver_detail_address, note, confirm_status, delete_status, use_integration) values (18, 2, 2, '201809150102000002', '2025-01-23 14:50:23', 'koobe', 18732, 16378, 0, 2344, 0, 10, 0, 1, 1, 1, 0, '顺丰快递', '201707196398345', 15, 13284, 13284, '单品促销,打折优惠：满3件，打7.50折,满减优惠：满1000.00元，减120.00元,满减优惠：满1000.00元，减120.00元,无优惠', 1, '1111', '1111', '18613030352', '1002219331@qq.com', '大梨', '18033441849', '518000', '广东省', '深圳市', '福田区', '东晓街道', '111', 0, 0, 1000);
insert into oms_order (id, member_id, coupon_id, order_sn, create_time, member_username, total_amount, pay_amount, freight_amount, promotion_amount, integration_amount, coupon_amount, discount_amount, pay_type, source_type, status, order_type, delivery_company, delivery_sn, auto_confirm_day, integration, growth, promotion_info, bill_type, bill_header, bill_content, bill_receiver_phone, bill_receiver_email, receiver_name, receiver_phone, receiver_post_code, receiver_province, receiver_city, receiver_region, receiver_detail_address, note, confirm_status, delete_status, use_integration) values (19, 2, 2, '201809130101000001', '2025-01-23 14:50:23', 'koobe', 18732, 16378, 0, 2344, 0, 10, 0, 2, 1, 2, 0, '顺丰快递', '201707196398345', 15, 13284, 13284, '单品促销,打折优惠：满3件，打7.50折,满减优惠：满1000.00元，减120.00元,满减优惠：满1000.00元，减120.00元,无优惠', 1, '1111', '1111', '18613030352', '1002219331@qq.com', '大梨', '18033441849', '518000', '广东省', '深圳市', '福田区', '东晓街道', '111', 0, 0, 1000);
insert into oms_order (id, member_id, coupon_id, order_sn, create_time, member_username, total_amount, pay_amount, freight_amount, promotion_amount, integration_amount, coupon_amount, discount_amount, pay_type, source_type, status, order_type, delivery_company, delivery_sn, auto_confirm_day, integration, growth, promotion_info, bill_type, bill_header, bill_content, bill_receiver_phone, bill_receiver_email, receiver_name, receiver_phone, receiver_post_code, receiver_province, receiver_city, receiver_region, receiver_detail_address, note, confirm_status, delete_status, use_integration) values (20, 2, 2, '201809130102000002', '2025-01-23 14:50:23', 'koobe', 18732, 16378, 0, 2344, 0, 10, 0, 1, 1, 3, 0, '顺丰快递', '201707196398345', 15, 13284, 13284, '单品促销,打折优惠：满3件，打7.50折,满减优惠：满1000.00元，减120.00元,满减优惠：满1000.00元，减120.00元,无优惠', 1, '1111', '1111', '18613030352', '1002219331@qq.com', '大梨', '18033441849', '518000', '广东省', '深圳市', '福田区', '东晓街道', '111', 1, 0, 1000);
insert into oms_order (id, member_id, coupon_id, order_sn, create_time, member_username, total_amount, pay_amount, freight_amount, promotion_amount, integration_amount, coupon_amount, discount_amount, pay_type, source_type, status, order_type, delivery_company, delivery_sn, auto_confirm_day, integration, growth, promotion_info, bill_type, bill_header, bill_content, bill_receiver_phone, bill_receiver_email, receiver_name, receiver_phone, receiver_post_code, receiver_province, receiver_city, receiver_region, receiver_detail_address, note, confirm_status, delete_status, use_integration) values (21, 2, 2, '201809140101000001', '2025-01-23 14:50:23', 'koobe', 18732, 16378, 0, 2344, 0, 10, 0, 2, 1, 4, 0, '顺丰快递', '201707196398345', 15, 13284, 13284, '单品促销,打折优惠：满3件，打7.50折,满减优惠：满1000.00元，减120.00元,满减优惠：满1000.00元，减120.00元,无优惠', 1, '1111', '1111', '18613030352', '1002219331@qq.com', '大梨', '18033441849', '518000', '广东省', '深圳市', '福田区', '东晓街道', '111', 0, 0, 1000);

-- 订单收货地址表建语句
drop table if exists oms_order_delivery;
create table oms_order_delivery
(
    id                bigserial primary key,
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
    update_time       timestamptz null,
    is_deleted        integer     default 1                 not null
);

-- 添加订单收货地址表注释
comment on table oms_order_delivery is '订单收货地址表';

-- 添加订单收货地址表列注释
comment on column oms_order_delivery.id is '';
comment on column oms_order_delivery.order_id is '订单id';
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
comment on column oms_order_delivery.is_deleted is '是否删除(0:否,1:是)';

insert into oms_order_delivery (order_id, order_no, receiver_name, receiver_phone, receiver_province,receiver_city, receiver_district, receiver_address, delivery_company,delivery_no
) values
-- delivery info for order ord20231201001
(1, 'ord20231201001', '张三', '13800138001', '北京市',
 '北京市', '朝阳区', '某某街道101号', '',
 ''),

-- delivery info for order ord20231201002
(2, 'ord20231201002', '李四', '13800138002', '上海市',
 '上海市', '浦东新区', '某某路202号', '',
 ''),

-- delivery info for order ord20231202001
(3, 'ord20231202001', '王五', '13800138003', '广东省',
 '广州市', '天河区', '某某大道303号', '顺丰速运',
 'sf123456789cn'),

-- delivery info for order ord20231203001
(4, 'ord20231203001', '赵六', '13800138004', '广东省',
 '深圳市', '南山区', '某某科技园404号', '圆通速递',
 'yt987654321cn'),

-- delivery info for order ord20231204001
(5, 'ord20231204001', '钱七', '13800138005', '浙江省',
 '杭州市', '西湖区', '某某广场505号', '',
 ''),

-- delivery info for order ord20231205001
(6, 'ord20231205001', '孙八', '13800138006', '江苏省',
 '南京市', '鼓楼区', '某某大厦606号', 'ems',
 'ems1122334455'),

-- delivery info for order ord20231206001
(7, 'ord20231206001', '周九', '13800138007', '浙江省',
 '宁波市', '鄞州区', '某某商业街707号', '中通快递',
 'zt5566778899');

-- 订单商品表建语句
drop table if exists oms_order_item;
create table oms_order_item
(
    id                bigserial primary key,
    order_id          bigint                                not null,
    order_no          varchar                               not null,
    order_item_status integer     default 1                 not null,
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
    is_deleted        integer     default 1                 not null
);

-- 添加订单商品表注释
comment on table oms_order_item is '订单商品表';

-- 添加订单商品表列注释
comment on column oms_order_item.id is '';
comment on column oms_order_item.order_id is '订单id';
comment on column oms_order_item.order_no is '订单编号';
comment on column oms_order_item.order_item_status is '订单商品状态：1-正常,2-退货申请中,3-已退货,4-已拒绝';
comment on column oms_order_item.sku_id is '商品sku id';
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
comment on column oms_order_item.is_deleted is '是否删除(0:否,1:是)';

insert into oms_order_item (
    order_id, order_no, order_item_status, sku_id, sku_name, sku_pic,
    sku_price, sku_quantity, spec_data, sku_total_amount, promotion_amount,
    coupon_amount, points_amount, discount_amount, real_amount, create_time
) values
-- order items for ord20231201001
(1, 'ord20231201001', 1, 3001, 'apple iphone 15 pro', 'http://129.204.203.29/big.png',
 299.00, 1, '{"color": "black", "storage": "256gb"}', 299.00, 0.00,
 0.00, 0.00, 0.00, 299.00, '2023-12-01 10:30:00.000'),

(1, 'ord20231201001', 1, 3001, 'apple iphone 15 pro', 'http://129.204.203.29/big.png',
 299.00, 1, '{"color": "black", "storage": "256gb"}', 299.00, 0.00,
 0.00, 0.00, 0.00, 299.00, '2023-12-01 10:30:00.000'),

-- order items for ord20231201002
(2, 'ord20231201002', 1, 3002, 'nike air max 270', 'http://129.204.203.29/big.png',
 159.90, 1, '{"size": "42", "color": "white"}', 159.90, 10.00,
 5.00, 0.00, 5.00, 139.90, '2023-12-01 14:20:00.000'),

-- order items for ord20231202001
(3, 'ord20231202001', 1, 3003, 'samsung galaxy watch5', 'http://129.204.203.29/big.png',
 89.50, 1, '{"color": "silver", "size": "44mm"}', 89.50, 0.00,
 0.00, 0.00, 0.00, 99.50, '2023-12-02 09:15:00.000'),

-- multiple items for ord20231203001
(4, 'ord20231203001', 1, 3004, 'sony wh-1000xm4 headphones', 'http://129.204.203.29/big.png',
 199.99, 1, '{"color": "black"}', 199.99, 0.00,
 20.00, 0.00, 0.00, 179.99, '2023-12-03 16:45:00.000'),

-- order items for ord20231204001
(5, 'ord20231204001', 1, 3006, 'ipad air 5', 'http://129.204.203.29/big.png',
 459.00, 1, '{"color": "blue", "storage": "64gb"}', 459.00, 0.00,
 0.00, 0.00, 0.00, 459.00, '2023-12-04 11:20:00.000'),

-- order items for ord20231205001
(6, 'ord20231205001', 1, 3007, 'macbook pro 13"', 'http://129.204.203.29/big.png',
 1299.00, 1, '{"color": "space gray", "storage": "256gb"}', 1299.00, 100.00,
 50.00, 0.00, 50.00, 1199.00, '2023-12-05 15:30:00.000'),

-- multiple items for ord20231206001
(7, 'ord20231206001', 1, 3001, 'apple iphone 15 pro', 'http://129.204.203.29/big.png',
 299.00, 1, '{"color": "black", "storage": "256gb"}', 299.00, 0.00,
 0.00, 0.00, 0.00, 299.00, '2023-12-06 11:15:00.000'),

(7, 'ord20231206001', 2, 3008, 'iphone screen protector', 'http://129.204.203.29/big.png',
 15.00, 1, '{"type": "tempered glass"}', 15.00, 0.00,
 0.00, 0.00, 0.00, 0.00, '2023-12-06 11:15:00.000');

-- 订单主表建语句
drop table if exists oms_order_main;
create table oms_order_main
(
    id                   bigserial primary key,
    order_no             varchar                               not null,
    user_id              bigint                                not null,
    order_status         integer     default 1                 not null,
    total_amount         numeric                               not null,
    promotion_amount     numeric                               not null,
    coupon_amount        numeric                               not null,
    points_amount        numeric                               not null,
    discount_amount      numeric                               not null,
    freight_amount       numeric                               not null,
    pay_amount           numeric                               not null,
    pay_type             integer     default 0                 not null,
    pay_time             timestamptz                           null,
    delivery_time        timestamptz                           null,
    receive_time         timestamptz                           null,
    comment_time         timestamptz                           null,
    source_type          integer     default 1                 not null,
    express_order_number varchar                               not null,
    use_points           integer                               not null,
    receive_status       integer     default 1                 not null,
    remark               varchar                               not null,
    create_time          timestamptz default current_timestamp not null,
    update_time          timestamptz null,
    is_deleted           integer     default 1                 not null
);

-- 添加订单主表注释
comment on table oms_order_main is '订单主表';

-- 添加订单主表列注释
comment on column oms_order_main.id is '';
comment on column oms_order_main.order_no is '订单编号';
comment on column oms_order_main.user_id is '用户id';
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
comment on column oms_order_main.source_type is '订单来源：1-app,2-pc,3-小程序';
comment on column oms_order_main.express_order_number is '快递单号';
comment on column oms_order_main.use_points is '下单时使用的积分';
comment on column oms_order_main.receive_status is '是否确认收货：0->否,1->是';
comment on column oms_order_main.remark is '订单备注';
comment on column oms_order_main.create_time is '提交时间';
comment on column oms_order_main.update_time is '';
comment on column oms_order_main.is_deleted is '是否删除(0:否,1:是)';

insert into oms_order_main (
    order_no, user_id, order_status, total_amount, promotion_amount, coupon_amount,
    points_amount, discount_amount, freight_amount, pay_amount, pay_type, pay_time,
    delivery_time, receive_time, comment_time, source_type, express_order_number,
    use_points, receive_status, remark
) values
-- pending payment order
('ord20231201001', 10001, 1, 299.00, 0.00, 0.00,
 0.00, 0.00, 0.00, 299.00, 0, null,
 null, null, null, 1, '',
 0, 0, '新订单'),

-- paid order
('ord20231201002', 10002, 2, 159.90, 10.00, 5.00,
 0.00, 5.00, 0.00, 139.90, 2, '2023-12-01 15:00:00',
 null, null, null, 2, '',
 0, 0, '已支付'),

-- shipped order
('ord20231202001', 10003, 3, 89.50, 0.00, 0.00,
 0.00, 0.00, 10.00, 99.50, 1, '2023-12-02 09:15:00',
 '2023-12-02 14:30:00', null, null, 3, 'sf123456789cn',
 100, 0, '已发货'),

-- completed order
('ord20231203001', 10004, 4, 199.99, 0.00, 20.00,
 0.00, 0.00, 0.00, 179.99, 2, '2023-12-03 16:45:00',
 '2023-12-04 09:00:00', '2023-12-05 10:00:00', '2023-12-05 11:00:00', 1, 'yt987654321cn',
 0, 1, '已完成'),

-- cancelled order
('ord20231204001', 10005, 5, 459.00, 0.00, 0.00,
 0.00, 0.00, 0.00, 459.00, 0, null,
 null, null, null, 2, '',
 0, 0, '用户取消'),

-- refunded order
('ord20231205001', 10006, 6, 1299.00, 100.00, 50.00,
 0.00, 50.00, 0.00, 1199.00, 1, '2023-12-05 15:30:00',
 '2023-12-06 10:00:00', '2023-12-07 14:00:00', null, 1, 'ems1122334455',
 0, 1, '已退款'),

-- after-sales order
('ord20231206001', 10007, 7, 299.00, 0.00, 0.00,
 0.00, 0.00, 0.00, 299.00, 2, '2023-12-06 11:15:00',
 '2023-12-07 15:00:00', '2023-12-08 09:00:00', null, 3, 'zt5566778899',
 0, 1, '售后处理中');


-- 订单操作记录表建语句
drop table if exists oms_order_operation_log;
create table oms_order_operation_log
(
    id             bigserial primary key,
    order_id       bigint                                not null,
    order_no       varchar                               not null,
    operation_type integer     default 1                 not null,
    operator_id    bigint                                not null,
    operator_type  integer     default 1                 not null,
    operator_note  varchar                               not null,
    create_time    timestamptz default current_timestamp not null
);

-- 添加订单操作记录表注释
comment on table oms_order_operation_log is '订单操作记录表';

-- 添加订单操作记录表列注释
comment on column oms_order_operation_log.id is '主键id';
comment on column oms_order_operation_log.order_id is '订单id';
comment on column oms_order_operation_log.order_no is '订单编号';
comment on column oms_order_operation_log.operation_type is '操作类型：1-创建订单，2-支付订单，3-发货，4-确认收货，5-取消订单，6-退款';
comment on column oms_order_operation_log.operator_id is '操作人id';
comment on column oms_order_operation_log.operator_type is '操作人类型：1-用户，2-系统，3-管理员';
comment on column oms_order_operation_log.operator_note is '操作备注';
comment on column oms_order_operation_log.create_time is '操作时间';

insert into oms_order_operation_log (
    order_id, order_no, operation_type, operator_id, operator_type, operator_note, create_time
) values
-- order creation logs
(1, 'ord20231201001', 1, 10001, 1, '用户创建订单', '2023-12-01 10:30:00'),
(2, 'ord20231201002', 1, 10002, 1, '用户创建订单', '2023-12-01 14:20:00'),
(3, 'ord20231202001', 1, 10003, 1, '用户创建订单', '2023-12-02 09:15:00'),
(4, 'ord20231203001', 1, 10004, 1, '用户创建订单', '2023-12-03 16:45:00'),
(5, 'ord20231204001', 1, 10005, 1, '用户创建订单', '2023-12-04 11:20:00'),
(6, 'ord20231205001', 1, 10006, 1, '用户创建订单', '2023-12-05 15:30:00'),
(7, 'ord20231206001', 1, 10007, 1, '用户创建订单', '2023-12-06 11:15:00'),

-- payment logs
(1, 'ord20231201001', 2, 10002, 1, '用户完成支付', '2023-12-01 15:00:00'),
(2, 'ord20231201002', 2, 10002, 1, '用户完成支付', '2023-12-01 15:00:00'),
(3, 'ord20231202001', 2, 10003, 1, '用户完成支付', '2023-12-02 09:15:00'),
(4, 'ord20231203001', 2, 10004, 1, '用户完成支付', '2023-12-03 16:45:00'),
(6, 'ord20231205001', 2, 10006, 1, '用户完成支付', '2023-12-05 15:30:00'),
(7, 'ord20231206001', 2, 10007, 1, '用户完成支付', '2023-12-06 11:15:00'),

-- shipping logs
(3, 'ord20231202001', 3, 1, 3, '管理员发货', '2023-12-02 14:30:00'),
(4, 'ord20231203001', 3, 1, 3, '管理员发货', '2023-12-04 09:00:00'),
(6, 'ord20231205001', 3, 1, 3, '管理员发货', '2023-12-06 10:00:00'),
(7, 'ord20231206001', 3, 1, 3, '管理员发货', '2023-12-07 15:00:00'),

-- receipt logs
(4, 'ord20231203001', 4, 10004, 1, '用户确认收货', '2023-12-05 10:00:00'),
(6, 'ord20231205001', 4, 10006, 1, '用户确认收货', '2023-12-07 14:00:00'),
(7, 'ord20231206001', 4, 10007, 1, '用户确认收货', '2023-12-08 09:00:00'),

-- cancellation logs
(5, 'ord20231204001', 5, 10005, 1, '用户取消订单', '2023-12-04 13:00:00'),

-- refund logs
(6, 'ord20231205001', 6, 1, 3, '管理员处理退款', '2023-12-08 10:00:00');

-- 订单支付表建语句
drop table if exists oms_order_payment;
create table oms_order_payment
(
    id             bigserial primary key,
    order_id       bigint                                not null,
    order_no       varchar                               not null,
    pay_type       integer     default 1                 not null,
    transaction_id varchar                               not null,
    total_amount   numeric                               not null,
    pay_amount     numeric                               not null,
    pay_status     integer     default 1                 not null,
    pay_time       timestamptz                           null,
    create_time    timestamptz default current_timestamp not null,
    update_time    timestamptz                           null,
    is_deleted     integer     default 1                 not null
);

-- 添加订单支付表注释
comment on table oms_order_payment is '订单支付表';

-- 添加订单支付表列注释
comment on column oms_order_payment.id is '主键id';
comment on column oms_order_payment.order_id is '订单id';
comment on column oms_order_payment.order_no is '订单编号';
comment on column oms_order_payment.pay_type is '支付方式：1-支付宝，2-微信，3-银联';
comment on column oms_order_payment.transaction_id is '支付流水号';
comment on column oms_order_payment.total_amount is '订单金额';
comment on column oms_order_payment.pay_amount is '支付金额';
comment on column oms_order_payment.pay_status is '支付状态：0-待支付，1-支付成功，2-支付失败';
comment on column oms_order_payment.pay_time is '支付时间';
comment on column oms_order_payment.create_time is '创建时间';
comment on column oms_order_payment.update_time is '';
comment on column oms_order_payment.is_deleted is '是否删除(0:否,1:是)';

insert into oms_order_payment (
    order_id, order_no, pay_type, transaction_id, total_amount,
    pay_amount, pay_status, pay_time, create_time, update_time
) values
-- payment info for order ord20231201001 (pending payment)
(1, 'ord20231201001', 1, 'ali2023120209159876541', 299.00,
 299.00, 2, null, '2023-12-01 10:30:00', null),
(1, 'ord20231201001', 2, 'wx20231203164511223311', 299.00,
 299.00, 0, null, '2023-12-01 10:30:00', null),

-- payment info for order ord20231201002 (paid via wechat)
(2, 'ord20231201002', 2, 'wx20231201150012345678', 159.90,
 139.90, 1, '2023-12-01 15:00:00', '2023-12-01 14:20:00', '2023-12-01 15:00:00'),

-- payment info for order ord20231202001 (paid via alipay)
(3, 'ord20231202001', 1, 'ali20231202091598765432', 89.50,
 99.50, 1, '2023-12-02 09:15:00', '2023-12-02 09:15:00', '2023-12-02 09:15:00'),

-- payment info for order ord20231203001 (paid via wechat)
(4, 'ord20231203001', 2, 'wx20231203164511223344', 199.99,
 179.99, 1, '2023-12-03 16:45:00', '2023-12-03 16:45:00', '2023-12-03 16:45:00'),

-- payment info for order ord20231204001 (cancelled, no payment)
(5, 'ord20231204001', 0, '', 459.00,
 459.00, 0, null, '2023-12-04 11:20:00', null),

-- payment info for order ord20231205001 (paid via alipay)
(6, 'ord20231205001', 1, 'ali20231205153055667788', 1299.00,
 1199.00, 1, '2023-12-05 15:30:00', '2023-12-05 15:30:00', '2023-12-08 10:00:00'),

-- payment info for order ord20231206001 (paid via unionpay)
(7, 'ord20231206001', 3, 'union20231206111599887766', 299.00,
 299.00, 1, '2023-12-06 11:15:00', '2023-12-06 11:15:00', '2023-12-06 11:15:00');

-- 订单优惠信息表建语句
drop table if exists oms_order_promotion;
create table oms_order_promotion
(
    id              bigserial primary key,
    order_id        bigint                                not null,
    order_no        varchar                               not null,
    promotion_type  integer     default 1                 not null,
    promotion_id    bigint                                null,
    promotion_name  varchar                               not null,
    discount_amount numeric                               not null,
    create_time     timestamptz default current_timestamp not null,
    is_deleted      integer     default 1                 not null
);

-- 添加订单优惠信息表注释
comment on table oms_order_promotion is '订单优惠信息表';

-- 添加订单优惠信息表列注释
comment on column oms_order_promotion.id is '主键id';
comment on column oms_order_promotion.order_id is '订单id';
comment on column oms_order_promotion.order_no is '订单编号';
comment on column oms_order_promotion.promotion_type is '优惠类型：1-优惠券，2-积分抵扣，3-会员折扣，4-促销活动';
comment on column oms_order_promotion.promotion_id is '优惠id';
comment on column oms_order_promotion.promotion_name is '优惠名称';
comment on column oms_order_promotion.discount_amount is '优惠金额';
comment on column oms_order_promotion.create_time is '';
comment on column oms_order_promotion.is_deleted is '是否删除(0:否,1:是)';

insert into oms_order_promotion (id, order_id, order_no, promotion_type, promotion_id, promotion_name, discount_amount, is_deleted)
values (1, 1, 'ord20240001', 1, 2001, '新人优惠券', 20.00, 0),
       (2, 1, 'ord20240002', 2, null, '积分抵扣', 15.50, 0),
       (3, 2, 'ord20240003', 3, null, '会员折扣', 10.00, 0),
       (4, 3, 'ord20240004', 4, 3001, '618大促', 50.00, 0),
       (5, 4, 'ord20240005', 1, 2002, '满减券', 30.00, 0);

-- 退货/售后主表建语句
drop table if exists oms_order_return;
create table oms_order_return
(
    id              bigserial primary key,
    order_id        bigint                                not null,
    return_no       varchar                               not null,
    member_id       bigint                                not null,
    status          integer     default 1                 not null,
    type            integer     default 1                 not null,
    reason          varchar                               not null,
    description     varchar                               not null,
    proof_pic       varchar                               not null,
    refund_amount   numeric                               not null,
    return_name     varchar                               not null,
    return_phone    varchar                               not null,
    company_address varchar                               not null,
    create_time     timestamptz default current_timestamp not null,
    handle_time     timestamptz                           null,
    handle_note     varchar                               not null,
    handle_man      varchar                               not null,
    receive_time    timestamptz                           null,
    receive_note    varchar                               not null,
    receive_man     varchar                               not null,
    refund_time     timestamptz                           null,
    close_time      timestamptz                           null,
    remark          varchar                               not null
);

-- 添加退货/售后主表注释
comment on table oms_order_return is '退货/售后主表';

-- 添加退货/售后主表列注释
comment on column oms_order_return.id is '主键id';
comment on column oms_order_return.order_id is '关联订单id';
comment on column oms_order_return.return_no is '退货单号';
comment on column oms_order_return.member_id is '会员id';
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

insert into oms_order_return (
    order_id, return_no, member_id, status, type, reason, description,
    proof_pic, refund_amount, return_name, return_phone, company_address,
    create_time, handle_time, handle_note, handle_man, receive_time,
    receive_note, receive_man, refund_time, close_time, remark
) values
-- pending review return request
(1001, 'rtn20231201001', 10001, 0, 0, '商品质量问题', '收到的商品有明显划痕和损坏',
 'pic1.jpg,pic2.jpg', 299.00, '张三', '13800138001', '北京市朝阳区某某街道101号',
 '2023-12-01 10:30:00', null, '', '', null,
 '', '', null, null, '客户 urgent'),

-- approved return
(1002, 'rtn20231201002', 10002, 1, 0, '尺寸不合适', '买大了，需要换小一号',
 'pic3.jpg', 159.90, '李四', '13800138002', '上海市浦东新区某某路202号',
 '2023-12-01 14:20:00', '2023-12-01 15:00:00', '已审核通过', '客服小王', null,
 '', '', null, null, '普通客户'),

-- received return
(1003, 'rtn20231202001', 10003, 2, 1, '发错商品', '收到的不是订购的商品',
 '', 89.50, '王五', '13800138003', '广州市天河区某某大道303号',
 '2023-12-02 09:15:00', '2023-12-02 10:00:00', '审核通过', '客服小李', '2023-12-03 14:30:00',
 '已收到退货商品', '仓库小赵', null, null, 'vip客户'),

-- refunded return
(1004, 'rtn20231203001', 10004, 3, 0, '商品不喜欢', '颜色与描述不符',
 'pic4.jpg,pic5.jpg', 199.99, '赵六', '13800138004', '深圳市南山区某某科技园404号',
 '2023-12-03 16:45:00', '2023-12-03 17:00:00', '同意退款', '客服小陈', '2023-12-04 09:00:00',
 '商品已验收', '仓库小孙', '2023-12-04 10:30:00', null, '需加快处理'),

-- rejected return
(1005, 'rtn20231204001', 10005, 4, 2, '人为损坏', '商品有明显人为使用痕迹',
 'pic6.jpg', 0.00, '钱七', '13800138005', '杭州市西湖区某某广场505号',
 '2023-12-04 11:20:00', '2023-12-04 13:00:00', '不符合退货条件，已拒绝', '客服小周', null,
 '', '', null, null, '注意客户态度'),

-- closed return
(1006, 'rtn20231205001', 10006, 5, 1, '其他原因', '客户主动取消退货申请',
 '', 0.00, '孙八', '13800138006', '南京市鼓楼区某某大厦606号',
 '2023-12-05 15:30:00', null, '', '', null,
 '', '', null, '2023-12-05 16:00:00', '客户自行解决');

-- 退货/售后明细表建语句
drop table if exists oms_order_return_item;
create table oms_order_return_item
(
    id            bigserial primary key,
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
comment on column oms_order_return_item.id is '主键id';
comment on column oms_order_return_item.return_id is '退货单id（关联oms_order_return.id）';
comment on column oms_order_return_item.order_id is '订单id';
comment on column oms_order_return_item.order_item_id is '订单明细id';
comment on column oms_order_return_item.sku_id is '商品sku id';
comment on column oms_order_return_item.sku_name is '商品名称';
comment on column oms_order_return_item.sku_pic is '商品图片';
comment on column oms_order_return_item.sku_attrs is '商品销售属性';
comment on column oms_order_return_item.quantity is '退货数量';
comment on column oms_order_return_item.product_price is '商品单价';
comment on column oms_order_return_item.real_amount is '实际退款金额';
comment on column oms_order_return_item.reason is '退货原因';
comment on column oms_order_return_item.remark is '备注';

insert into oms_order_return_item (
    return_id, order_id, order_item_id, sku_id, sku_name, sku_pic,
    sku_attrs, quantity, product_price, real_amount, reason, remark
) values
-- return items for return order rtn20231201001
(1, 1001, 2001, 3001, 'apple iphone 15 pro', 'http://129.204.203.29/big.png',
 'color:black,storage:256gb', 1, 299.00, 299.00, '商品质量问题', '屏幕有划痕'),

-- return items for return order rtn20231201002
(2, 1002, 2002, 3002, 'nike air max 270', 'http://129.204.203.29/big.png',
 'size:42,color:white', 1, 159.90, 159.90, '尺寸不合适', '需要换小一号'),

-- return items for return order rtn20231202001
(3, 1003, 2003, 3003, 'samsung galaxy watch5', 'http://129.204.203.29/big.png',
 'color:silver,size:44mm', 1, 89.50, 89.50, '发错商品', '收到的是40mm版本'),

-- multiple items for return order rtn20231203001
(4, 1004, 2004, 3004, 'sony wh-1000xm4 headphones', 'http://129.204.203.29/big.png',
 'color:black', 1, 199.99, 199.99, '商品不喜欢', '颜色与网站图片差异较大'),

(4, 1004, 2005, 3005, 'sony headphone case', 'http://129.204.203.29/big.png',
 'color:black', 1, 25.00, 0.00, '配套商品退货', '主商品退货，配件一并退回'),

-- return items for return order rtn20231204001
(5, 1005, 2006, 3006, 'ipad air 5', 'http://129.204.203.29/big.png',
 'color:blue,storage:64gb', 1, 459.00, 0.00, '人为损坏', '屏幕有裂痕'),

-- return items for return order rtn20231205001
(6, 1006, 2007, 3007, 'macbook pro 13"', 'http://129.204.203.29/big.png',
 'color:space gray,storage:256gb', 1, 1299.00, 0.00, '其他原因', '客户主动取消');

-- 退货原因表建语句
drop table if exists oms_order_return_reason;
create table oms_order_return_reason
(
    id          bigserial primary key,
    name        varchar                               not null,
    sort        integer                               not null,
    status      integer     default 1                 not null,
    create_by   varchar     default ''                not null,
    create_time timestamptz default current_timestamp not null,
    update_by   varchar     default ''                not null,
    update_time timestamptz                           null,
    is_deleted  integer     default 1                 not null
);
insert into oms_order_return_reason (id, name, sort, status, create_by)
values (1, '质量问题', 1, 1, 1),
       (2, '尺码太大', 1, 1, 1),
       (3, '颜色不喜欢', 1, 1, 1),
       (4, '7天无理由退货', 1, 1, 1),
       (5, '价格问题', 1, 0, 1),
       (12, '发票问题', 0, 1, 1),
       (13, '其他问题', 0, 1, 1),
       (14, '物流问题', 0, 1, 1),
       (15, '售后问题', 0, 1, 1);

-- 添加退货原因表注释
comment on table oms_order_return_reason is '退货原因表';

-- 添加退货原因表列注释
comment on column oms_order_return_reason.id is '主键id';
comment on column oms_order_return_reason.name is '退货类型';
comment on column oms_order_return_reason.sort is '排序';
comment on column oms_order_return_reason.status is '状态：0->不启用；1->启用';
comment on column oms_order_return_reason.create_by is '创建人id';
comment on column oms_order_return_reason.create_time is '创建时间';
comment on column oms_order_return_reason.update_by is '更新人id';
comment on column oms_order_return_reason.update_time is '更新时间';
comment on column oms_order_return_reason.is_deleted is '是否删除(0:否,1:是)';

-- 订单设置表建语句
drop table if exists oms_order_setting;
create table oms_order_setting
(
    id                    bigserial primary key,
    flash_order_overtime  integer                               not null,
    normal_order_overtime integer                               not null,
    confirm_overtime      integer                               not null,
    finish_overtime       integer                               not null,
    status                integer     default 1                 not null,
    is_default            integer     default 1                 not null,
    comment_overtime      integer                               not null,
    create_by             varchar     default ''                not null,
    create_time           timestamptz default current_timestamp not null,
    update_by             varchar     default ''                not null,
    update_time           timestamptz  null,
    is_deleted            integer     default 1                 not null
);

-- 添加订单设置表注释
comment on table oms_order_setting is '订单设置表';

-- 添加订单设置表列注释
comment on column oms_order_setting.id is '主键id';
comment on column oms_order_setting.flash_order_overtime is '秒杀订单超时关闭时间(分)';
comment on column oms_order_setting.normal_order_overtime is '正常订单超时时间(分)';
comment on column oms_order_setting.confirm_overtime is '发货后自动确认收货时间（天）';
comment on column oms_order_setting.finish_overtime is '自动完成交易时间，不能申请售后（天）';
comment on column oms_order_setting.status is '状态：0->禁用；1->启用';
comment on column oms_order_setting.is_default is '是否默认：0->否；1->是';
comment on column oms_order_setting.comment_overtime is '订单完成后自动好评时间（天）';
comment on column oms_order_setting.create_by is '创建人id';
comment on column oms_order_setting.create_time is '创建时间';
comment on column oms_order_setting.update_by is '更新人id';
comment on column oms_order_setting.update_time is '更新时间';
comment on column oms_order_setting.is_deleted is '是否删除(0:否,1:是)';

insert into oms_order_setting (id, flash_order_overtime, normal_order_overtime, confirm_overtime, finish_overtime, status, is_default, comment_overtime, create_by, is_deleted)
values (1, 30, 60, 7, 15, 1, 1, 7, 1, 0),
       (2, 45, 90, 10, 20, 1, 0, 10, 2, 0),
       (3, 60, 120, 5, 10, 0, 0, 5, 3, 0);

-- 优惠券表建语句
drop table if exists sms_coupon;
create table sms_coupon
(
    id             bigserial primary key,
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
    status         integer     default 1                 not null,
    is_enabled     integer     default 1                 not null,
    description    varchar                               not null,
    create_by      varchar     default ''                not null,
    create_time    timestamptz default current_timestamp not null,
    update_by      varchar     default ''                not null,
    update_time    timestamptz                            null,
    is_deleted     integer     default 1                 not null
);

-- 添加优惠券表注释
comment on table sms_coupon is '优惠券表';

-- 添加优惠券表列注释
comment on column sms_coupon.id is '优惠券id';
comment on column sms_coupon.type_id is '优惠券类型id';
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
comment on column sms_coupon.create_by is '创建人id';
comment on column sms_coupon.create_time is '创建时间';
comment on column sms_coupon.update_by is '更新人id';
comment on column sms_coupon.update_time is '更新时间';
comment on column sms_coupon.is_deleted is '是否删除(0:否,1:是)';

insert into sms_coupon (type_id, name, code, amount, min_amount, start_time, end_time, total_count, received_count, used_count, per_limit, status, is_enabled, description, create_by, is_deleted)
values (1, '满减优惠券', 'fullreduce', 20.00, 150.00, '2023-11-01 00:00:00', '2029-11-30 23:59:59', 300, 0, 0, 1, 1, 1, '满150减20', 3, 0),
       (3, '新用户优惠券', 'newuser2023', 50.00, 200.00, '2023-11-01 00:00:00', '2029-12-31 23:59:59', 1000, 0, 0, 1, 0, 1, '适用于新用户首次购物', 1, 0),
       (5, '双十一折扣券', 'double11', 10.00, 100.00, '2023-11-11 00:00:00', '2029-11-11 23:59:59', 500, 0, 0, 2, 1, 1, '双十一当天使用', 2, 0);

-- 优惠券领取记录表建语句
drop table if exists sms_coupon_record;
create table sms_coupon_record
(
    id              bigserial primary key,
    coupon_id       bigint                                not null,
    member_id       bigint                                not null,
    get_time        timestamptz                           not null,
    get_type        integer     default 1                 not null,
    use_time        timestamptz                            null,
    order_id        bigint     default 0                  not null,
    order_amount    numeric    default 0                  not null,
    discount_amount numeric    default 0                  not null,
    status          integer     default 0                 not null,
    invalid_time    timestamptz                            null,
    invalid_reason  varchar    default ''                 not null,
    create_time     timestamptz default current_timestamp not null,
    is_deleted      integer     default 0                 not null
);

-- 添加优惠券领取记录表注释
comment on table sms_coupon_record is '优惠券领取记录表';

-- 添加优惠券领取记录表列注释
comment on column sms_coupon_record.id is '主键id';
comment on column sms_coupon_record.coupon_id is '优惠券id';
comment on column sms_coupon_record.member_id is '用户id';
comment on column sms_coupon_record.get_time is '领取时间';
comment on column sms_coupon_record.get_type is '获取类型：0->后台赠送；1->主动获取';
comment on column sms_coupon_record.use_time is '使用时间';
comment on column sms_coupon_record.order_id is '使用订单id';
comment on column sms_coupon_record.order_amount is '订单金额';
comment on column sms_coupon_record.discount_amount is '优惠金额';
comment on column sms_coupon_record.status is '状态：0-未使用，1-已使用，2-已过期，3-已失效';
comment on column sms_coupon_record.invalid_time is '失效时间';
comment on column sms_coupon_record.invalid_reason is '失效原因';
comment on column sms_coupon_record.create_time is '创建时间';
comment on column sms_coupon_record.is_deleted is '是否删除(0:否,1:是)';

insert into sms_coupon_record (coupon_id, member_id, get_time, get_type, use_time, invalid_time)
values (1, 1001, current_timestamp, 0, null, null),
       (2, 1001, current_timestamp, 1, null, null),
       (3, 1001, current_timestamp, 1, null, null);

-- 优惠券使用范围表建语句
drop table if exists sms_coupon_scope;
create table sms_coupon_scope
(
    id          bigserial primary key,
    coupon_id   bigint                                not null,
    scope_type  integer     default 1                 not null,
    scope_id    bigint                                not null,
    create_time timestamptz default current_timestamp not null
);

-- 添加优惠券使用范围表注释
comment on table sms_coupon_scope is '优惠券使用范围表';

-- 添加优惠券使用范围表列注释
comment on column sms_coupon_scope.id is '主键id';
comment on column sms_coupon_scope.coupon_id is '优惠券id';
comment on column sms_coupon_scope.scope_type is '范围类型：0-全场，1-分类，2-商品';
comment on column sms_coupon_scope.scope_id is '范围id（分类id或商品id）';
comment on column sms_coupon_scope.create_time is '创建时间';

insert into sms_coupon_scope (coupon_id, scope_type, scope_id)
values (1, 0, 0),
       (2, 0, 0),
       (3, 0, 0);
-- 优惠券类型表建语句
drop table if exists sms_coupon_type;
create table sms_coupon_type
(
    id            bigserial primary key,
    name          varchar                               not null,
    code          varchar                               not null,
    description   varchar                               not null,
    discount_type integer     default 1                 not null,
    status        integer     default 1                 not null,
    sort          integer                               not null,
    create_by     varchar     default ''                not null,
    create_time   timestamptz default current_timestamp not null,
    update_by     varchar     default ''                not null,
    update_time   timestamptz                           null,
    is_deleted    integer     default 1                 not null
);

-- 添加优惠券类型表注释
comment on table sms_coupon_type is '优惠券类型表';

-- 添加优惠券类型表列注释
comment on column sms_coupon_type.id is '主键id';
comment on column sms_coupon_type.name is '类型名称';
comment on column sms_coupon_type.code is '类型编码';
comment on column sms_coupon_type.description is '描述';
comment on column sms_coupon_type.discount_type is '优惠方式：1-固定金额，2-折扣率，3-第n件特惠，4-买赠，5-特价，6-套装优惠，7-搭配优惠，8-积分抵现，9-积分倍率，10-免运费，11-运费减免，12-限时特权，13-会员特权';
comment on column sms_coupon_type.status is '是否启用';
comment on column sms_coupon_type.sort is '排序';
comment on column sms_coupon_type.create_by is '创建人id';
comment on column sms_coupon_type.create_time is '创建时间';
comment on column sms_coupon_type.update_by is '更新人id';
comment on column sms_coupon_type.update_time is '更新时间';
comment on column sms_coupon_type.is_deleted is '是否删除(0:否,1:是)';

insert into sms_coupon_type (id, name, code, description, discount_type, status, sort, create_by, is_deleted)
values (1, '满减券', 'manjian', '满足最低消费即可减去固定金额', 1, 1, 100, 1, 0),
       (2, '无门槛券', 'wumenjian', '无最低消费限制的固定金额优惠', 1, 1, 90, 1, 0),
       (3, '新人券', 'xinren', '新用户专享固定金额优惠', 1, 1, 80, 1, 0),
       (4, '会员券', 'huiyuan', '会员专享固定金额优惠', 1, 1, 70, 1, 0),
       (5, '折扣券', 'zhekou', '按商品金额的比例进行折扣', 2, 1, 60, 1, 0),
       (6, '会员折扣券', 'huiyuan_zhekou', '会员专享折扣优惠', 2, 1, 50, 1, 0),
       (7, '限时折扣券', 'xianshi_zhekou', '限时特惠折扣', 2, 1, 40, 1, 0),
       (8, '品类折扣券', 'pinlei_zhekou', '特定品类商品的折扣优惠', 2, 1, 30, 1, 0),
       (9, '已停用满减券', 'disable_manjian', '已停用的满减优惠券类型', 1, 0, 10, 1, 0),
       (10, '已停用折扣券', 'disable_zhekou', '已停用的折扣优惠券类型', 2, 0, 20, 1, 0),
       (11, '第二件半价', 'second_half', '第二件商品半价优惠', 3, 1, 95, 1, 0),
       (12, '买二送一', 'buy_two_get_one', '买两件赠送一件', 4, 1, 94, 1, 0),
       (13, '第三件1元', 'third_one_yuan', '第三件商品1元特惠', 5, 1, 93, 1, 0),
       (14, '套装优惠', 'bundle_discount', '购买指定套装商品享受优惠', 6, 1, 85, 1, 0),
       (15, '搭配优惠', 'combo_discount', '指定商品搭配购买享受优惠', 7, 1, 84, 1, 0),
       (16, '积分抵现券', 'points_cash', '使用积分抵扣现金', 8, 1, 75, 1, 0),
       (17, '积分翻倍券', 'points_multiply', '购物获得双倍积分', 9, 1, 74, 1, 0),
       (18, '免运费券', 'free_shipping', '订单免运费', 10, 1, 65, 1, 0),
       (19, '运费减免券', 'shipping_discount', '运费优惠固定金额', 11, 1, 64, 1, 0),
       (20, '限时优先券', 'time_priority', '限时抢购优先购买权', 12, 1, 55, 1, 0),
       (21, '会员日特权券', 'vip_day', '会员日专享双重优惠', 13, 1, 54, 1, 0);

-- 首页轮播广告表建语句
drop table if exists sms_home_advertise;
create table sms_home_advertise
(
    id          bigserial primary key,
    name        varchar                               not null,
    type        integer     default 1                 not null,
    pic         varchar                               not null,
    start_time  timestamptz                           not null,
    end_time    timestamptz                           not null,
    status      integer     default 1                 not null,
    click_count integer                               not null,
    order_count integer                               not null,
    url         varchar                               not null,
    remark      varchar                               not null,
    sort        integer                               not null,
    create_time timestamptz default current_timestamp not null,
    update_time timestamptz                           null,
    is_deleted  integer     default 1                 not null
);

-- 添加首页轮播广告表注释
comment on table sms_home_advertise is '首页轮播广告表';

-- 添加首页轮播广告表列注释
comment on column sms_home_advertise.id is '编号';
comment on column sms_home_advertise.name is '名称';
comment on column sms_home_advertise.type is '轮播位置：0->pc首页轮播；1->app首页轮播';
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
comment on column sms_home_advertise.is_deleted is '是否删除(0:否,1:是)';

insert into sms_home_advertise (id, name, type, pic, start_time, end_time, status, click_count, order_count, url, remark, sort)
values (1, '美食5折限时购', 1, 'http://129.204.203.29/banner01.png', '2023-11-01 00:00:00', '2029-11-11 23:59:59', 1, 0, 0, 'http://example.com/promo1', '双十一活动', 100),
       (2, '特惠全场享8折', 1, 'http://129.204.203.29/banner02.png', '2023-12-01 00:00:00', '2029-12-25 23:59:59', 1, 0, 0, 'http://example.com/promo2', '圣诞节活动', 90),
       (3, '邀请好友送优惠券', 1, 'http://129.204.203.29/banner03.png', '2023-12-26 00:00:00', '2029-01-01 23:59:59', 1, 0, 0, 'http://example.com/promo3', '新年活动', 80);

-- 秒杀活动表建语句
drop table if exists sms_seckill_activity;
create table sms_seckill_activity
(
    id          bigserial primary key,
    name        varchar                               not null,
    description varchar                               not null,
    start_time  timestamptz                           not null,
    end_time    timestamptz                           not null,
    status      integer     default 1                 not null,
    is_enabled  integer     default 1                 not null,
    create_by   varchar     default ''                not null,
    create_time timestamptz default current_timestamp not null,
    update_by   varchar     default ''                not null,
    update_time timestamptz                           null,
    is_deleted  integer     default 1                 not null
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
comment on column sms_seckill_activity.create_by is '创建人id';
comment on column sms_seckill_activity.create_time is '创建时间';
comment on column sms_seckill_activity.update_by is '更新人id';
comment on column sms_seckill_activity.update_time is '更新时间';
comment on column sms_seckill_activity.is_deleted is '是否删除(0:否,1:是)';

-- 插入秒杀活动数据
insert into sms_seckill_activity (id, name, description, start_time, end_time, status, is_enabled, create_by, is_deleted)
values (1, '双十一秒杀活动', '双十一期间的限时秒杀活动', '2025-11-11 00:00:00', '2025-11-11 23:59:59', 0, 1, 1, 0),
       (2, '黑色星期五秒杀', '黑色星期五特惠秒杀活动', '2025-11-24 00:00:00', '2025-11-24 23:59:59', 0, 1, 2, 0),
       (3, '圣诞节秒杀', '圣诞节期间的限时秒杀活动', '2025-12-25 00:00:00', '2025-12-25 23:59:59', 0, 1, 3, 0);

-- 秒杀商品表建语句
drop table if exists sms_seckill_product;
create table sms_seckill_product
(
    id            bigserial primary key,
    activity_id   bigint                                not null,
    session_id    bigint                                not null,
    sku_id        bigint                                not null,
    sku_name      varchar                               not null,
    seckill_price numeric                               not null,
    seckill_stock integer                               not null,
    stock_locked  integer                               not null,
    per_limit     integer                               not null,
    sort          integer                               not null,
    status        integer     default 1                 not null,
    create_by     varchar     default ''                not null,
    create_time   timestamptz default current_timestamp not null,
    update_by     varchar     default ''                not null,
    update_time   timestamptz                           null,
    is_deleted    integer     default 1                 not null
);

-- 添加秒杀商品表注释
comment on table sms_seckill_product is '秒杀商品表';

-- 添加秒杀商品表列注释
comment on column sms_seckill_product.id is 'id';
comment on column sms_seckill_product.activity_id is '活动id';
comment on column sms_seckill_product.session_id is '秒杀场次id';
comment on column sms_seckill_product.sku_id is '商品sku id';
comment on column sms_seckill_product.sku_name is '商品名称';
comment on column sms_seckill_product.seckill_price is '秒杀价格';
comment on column sms_seckill_product.seckill_stock is '秒杀库存';
comment on column sms_seckill_product.stock_locked is '锁定库存';
comment on column sms_seckill_product.per_limit is '每人限购数量';
comment on column sms_seckill_product.sort is '排序';
comment on column sms_seckill_product.status is '状态：0-未上架，1-已上架';
comment on column sms_seckill_product.create_by is '创建人id';
comment on column sms_seckill_product.create_time is '创建时间';
comment on column sms_seckill_product.update_by is '更新人id';
comment on column sms_seckill_product.update_time is '更新时间';
comment on column sms_seckill_product.is_deleted is '是否删除(0:否,1:是)';

insert into sms_seckill_product (id, activity_id, session_id, sku_id, sku_name, seckill_price, seckill_stock, stock_locked, per_limit, sort, status, create_by, is_deleted)
values (1, 1, 1, 101, '华为手机',99.99, 100, 0, 1, 100, 1, 1, 0),
       (2, 1, 2, 102, '苹果手机',49.99, 200, 0, 2, 90, 1, 2, 0),
       (3, 2, 3, 103, '小米手机',29.99, 150, 0, 1, 80, 1, 3, 0);

-- 秒杀预约表建语句
drop table if exists sms_seckill_reservation;
create table sms_seckill_reservation
(
    id          bigserial primary key,
    user_id     bigint                                not null,
    activity_id bigint                                not null,
    product_id  bigint                                not null,
    status      integer     default 1                 not null,
    create_time timestamptz default current_timestamp not null,
    update_time timestamptz                           null,
    is_deleted  integer     default 1                 not null
);

-- 添加秒杀预约表注释
comment on table sms_seckill_reservation is '秒杀预约表';

-- 添加秒杀预约表列注释
comment on column sms_seckill_reservation.id is 'id';
comment on column sms_seckill_reservation.user_id is '用户id';
comment on column sms_seckill_reservation.activity_id is '活动id';
comment on column sms_seckill_reservation.product_id is '秒杀商品id';
comment on column sms_seckill_reservation.status is '状态：0-已预约，1-已参与，2-已取消';
comment on column sms_seckill_reservation.create_time is '';
comment on column sms_seckill_reservation.update_time is '';
comment on column sms_seckill_reservation.is_deleted is '是否删除(0:否,1:是)';

-- 秒杀场次表建语句
drop table if exists sms_seckill_session;
create table sms_seckill_session
(
    id          bigserial primary key,
    name        varchar                               not null,
    start_time  varchar                               not null,
    end_time    varchar                               not null,
    status      integer     default 1                 not null,
    sort        integer                               not null,
    create_by   varchar     default ''                not null,
    create_time timestamptz default current_timestamp not null,
    update_by   varchar     default ''                not null,
    update_time timestamptz                           null,
    is_deleted  integer     default 1                 not null
);

-- 添加秒杀场次表注释
comment on table sms_seckill_session is '秒杀场次表';

-- 添加秒杀场次表列注释
comment on column sms_seckill_session.id is '秒杀场次id';
comment on column sms_seckill_session.name is '场次名称';
comment on column sms_seckill_session.start_time is '开始时间';
comment on column sms_seckill_session.end_time is '结束时间';
comment on column sms_seckill_session.status is '状态：0-禁用，1-启用';
comment on column sms_seckill_session.sort is '排序';
comment on column sms_seckill_session.create_by is '创建人id';
comment on column sms_seckill_session.create_time is '创建时间';
comment on column sms_seckill_session.update_by is '更新人id';
comment on column sms_seckill_session.update_time is '更新时间';
comment on column sms_seckill_session.is_deleted is '是否删除(0:否,1:是)';

insert into sms_seckill_session (name, start_time, end_time, status, sort, create_by)
values ('早场', '10:00:00', '12:00:00', 1, 100, 1),
       ('午场', '14:00:00', '16:00:00', 1, 90, 1),
       ('晚场', '19:00:00', '21:00:00', 1, 80, 1);

-- 会员收货地址表建语句
drop table if exists ums_member_address;
create table ums_member_address
(
    id             bigserial primary key,
    member_id      bigint                                not null,
    receiver_name  varchar                               not null,
    receiver_phone varchar                               not null,
    province       varchar                               not null,
    city           varchar                               not null,
    district       varchar                               not null,
    detail_address varchar                               not null,
    postal_code    varchar                               not null,
    tag            varchar                               not null,
    is_default     integer     default 1                 not null,
    create_time    timestamptz default current_timestamp not null,
    update_time    timestamptz                           null,
    is_deleted     integer     default 1                 not null
);

-- 添加会员收货地址表注释
comment on table ums_member_address is '会员收货地址表';

-- 添加会员收货地址表列注释
comment on column ums_member_address.id is '主键id';
comment on column ums_member_address.member_id is '会员id';
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
comment on column ums_member_address.is_deleted is '是否删除(0:否,1:是)';

insert into ums_member_address (id, member_id, receiver_name, receiver_phone, province, city, district, detail_address, postal_code, tag, is_default)
values (1, 1001, '张三', '13800138001', '广东省', '深圳市', '南山区', '科技园科兴科学园b座', '518057', '公司', 0),
       (2, 1001, '张三', '13800138001', '广东省', '深圳市', '福田区', '福中三路1006号诺德中心', '518048', '家', 1),
       (3, 1002, '李四', '13800138002', '广东省', '广州市', '天河区', '珠江新城花城大道85号高德置地广场', '510623', '公司', 1),
       (4, 1002, '李四妈妈', '13800138003', '广东省', '广州市', '越秀区', '解放北路928号', '510030', '父母家', 0),
       (5, 1003, '王五', '13800138004', '广东省', '珠海市', '香洲区', '情侣南路399号', '519000', '家', 1),
       (6, 1003, '王五', '13800138004', '广东省', '珠海市', '横琴新区', '环岛东路2000号', '519031', '公司', 0),
       (7, 1004, '赵六', '13800138005', '广东省', '东莞市', '南城区', '鸿福路200号第一国际', '523000', '公司', 1),
       (8, 1005, '钱七', '13800138006', '广东省', '深圳市', '宝安区', '新安街道创业二路腾讯滨海大厦', '518101', '公司', 1),
       (9, 1005, '钱七', '13800138006', '广东省', '深圳市', '罗湖区', '东门南路金光华广场', '518001', '家', 0);


-- 积分消费设置建语句
drop table if exists ums_member_consume_setting;
create table ums_member_consume_setting
(
    id                    bigserial primary key,
    deduction_per_amount  integer                               not null,
    max_percent_per_order integer                               not null,
    use_unit              integer                               not null,
    coupon_status         integer     default 1                 not null,
    status                integer     default 1                 not null,
    create_by             varchar     default ''                not null,
    create_time           timestamptz default current_timestamp not null,
    update_by             varchar     default ''                not null,
    update_time           timestamptz                           null
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
comment on column ums_member_consume_setting.create_by is '创建人id';
comment on column ums_member_consume_setting.create_time is '创建时间';
comment on column ums_member_consume_setting.update_by is '更新人id';
comment on column ums_member_consume_setting.update_time is '更新时间';

insert into ums_member_consume_setting
(deduction_per_amount, max_percent_per_order, use_unit, coupon_status, status, create_by)
values (10, 50, 100, 0, 1, 1), -- 每1元抵扣10积分，最高抵扣订单金额的50%，最小使用单位100积分，不可与优惠券同用
       (20, 70, 100, 1, 0, 1), -- 每1元抵扣20积分，最高抵扣订单金额的70%，最小使用单位100积分，可与优惠券同用
       (5, 30, 50, 1, 0, 1); -- 每1元抵扣5积分，最高抵扣订单金额的30%，最小使用单位50积分，可与优惠券同用

-- 会员成长值记录表建语句
drop table if exists ums_member_growth_log;
create table ums_member_growth_log
(
    id            bigserial primary key,
    member_id     bigint                                not null,
    change_type   integer     default 1                 not null,
    change_growth integer                               not null,
    source_type   integer     default 1                 not null,
    description   varchar                               not null,
    operate_man   varchar                               not null,
    operate_note  varchar                               not null,
    create_time   timestamptz default current_timestamp not null
);

-- 添加会员成长值记录表注释
comment on table ums_member_growth_log is '会员成长值记录表';

-- 添加会员成长值记录表列注释
comment on column ums_member_growth_log.id is '';
comment on column ums_member_growth_log.member_id is '会员id';
comment on column ums_member_growth_log.change_type is '变更类型：1-添加成长值，2-减少成长值';
comment on column ums_member_growth_log.change_growth is '变更成长值';
comment on column ums_member_growth_log.source_type is '来源类型：0-其他，1-订单，2-活动，3-签到，4-管理员修改';
comment on column ums_member_growth_log.description is '描述';
comment on column ums_member_growth_log.operate_man is '操作人员';
comment on column ums_member_growth_log.operate_note is '操作备注';
comment on column ums_member_growth_log.create_time is '创建时间';

insert into ums_member_growth_log (id, member_id, change_type, change_growth, source_type, description, operate_man, operate_note)
values (1, 1001, 1, 100, 1, '订单完成增加成长值', '管理员a', '订单id: 12345'),
       (2, 1002, 2, 50, 2, '活动参与减少成长值', '管理员b', '活动id: 67890'),
       (3, 1003, 1, 200, 3, '签到增加成长值', '管理员c', '签到id: 11223'),
       (4, 1004, 2, 30, 4, '管理员修改减少成长值', '管理员d', '手动调整'),
       (5, 1005, 1, 150, 0, '其他来源增加成长值', '管理员e', '特殊奖励');

-- 会员信息表建语句
drop table if exists ums_member_info;
create table ums_member_info
(
    id                 bigserial primary key,
    member_id          bigint                                not null,
    wx_openid          varchar    default ''                not null,
    level_id           bigint                                not null,
    nickname           varchar                               not null,
    mobile             varchar                               not null,
    source             integer     default 1                 not null,
    password           varchar                               not null,
    avatar             varchar                               not null,
    signature          varchar                               not null,
    gender             integer     default 1                 not null,
    birthday           date                                  null,
    growth_point       integer                               not null,
    points             integer                               not null,
    total_points       integer                               not null,
    spend_amount       numeric                               not null,
    order_count        integer                               not null,
    coupon_count       integer                               not null,
    comment_count      integer                               not null,
    return_count       integer                               not null,
    lottery_times      integer                               not null,
    first_login_status integer     default 1                 not null,
    last_login         timestamptz                           null,
    is_enabled         integer     default 1                 not null,
    create_time        timestamptz default current_timestamp not null,
    update_time        timestamptz                           null,
    is_deleted         integer     default 1                 not null
);

-- 添加会员信息表注释
comment on table ums_member_info is '会员信息表';

-- 添加会员信息表列注释
comment on column ums_member_info.id is '主键id';
comment on column ums_member_info.member_id is '会员id';
comment on column ums_member_info.wx_openid is '微信openid';
comment on column ums_member_info.level_id is '等级id';
comment on column ums_member_info.nickname is '昵称';
comment on column ums_member_info.mobile is '手机号码';
comment on column ums_member_info.source is '注册来源：0-pc，1-app，2-小程序';
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
comment on column ums_member_info.is_deleted is '是否删除(0:否,1:是)';


insert into ums_member_info (id, member_id, level_id, nickname, mobile, source, password, avatar, signature, gender, birthday, growth_point, points, total_points, spend_amount, order_count, coupon_count, comment_count, return_count, lottery_times, last_login, is_enabled)
values (1, 1001, 1, '张三', '13800138001', 0, '123456', 'https://example.com/avatar/001.jpg', '生活就是购物~', 1, '1990-01-15', 100, 500, 1000, 999.99, 10, 3, 8, 1, 2, '2024-01-15 08:30:00', 1),
       (2, 1002, 2, '李四', '13800138002', 1, '123456', 'https://example.com/avatar/002.jpg', '潮流时尚，品质生活', 2, '1995-03-20', 500, 1200, 2000, 2999.99, 25, 8, 20, 0, 5, '2024-01-15 12:45:00', 1),
       (3, 1003, 3, '王五', '13800138003', 2, '123456', 'https://example.com/avatar/003.jpg', 'vip专享生活', 1, '1988-07-08', 1000, 3000, 5000, 8999.99, 50, 15, 45, 2, 10, '2024-01-15 15:20:00', 1),
       (4, 1004, 4, '赵六', '13800138004', 1, '123456', 'https://example.com/avatar/004.jpg', '享受生活每一天', 2, '1992-12-25', 2000, 5000, 8000, 15999.99, 80, 20, 75, 3, 15, '2024-01-15 18:10:00', 1),
       (5, 1005, 5, '钱七', '13800138005', 0, '123456', 'https://example.com/avatar/005.jpg', '至尊购物体验', 1, '1985-05-01', 5000, 10000, 15000, 29999.99, 150, 30, 120, 5, 20, '2024-01-15 20:30:00', 1);


-- 会员等级表建语句
drop table if exists ums_member_level;
create table ums_member_level
(
    id            bigserial primary key,
    name          varchar                               not null,
    level         integer                               not null,
    growth_point  integer                               not null,
    discount_rate numeric                               not null,
    free_freight  integer     default 1                 not null,
    comment_extra integer     default 1                 not null,
    privileges    varchar                               not null,
    remark        varchar                               not null,
    is_enabled    integer     default 1                 not null,
    create_by     varchar     default ''                not null,
    create_time   timestamptz default current_timestamp not null,
    update_by     varchar     default ''                not null,
    update_time   timestamptz                           null,
    is_deleted    integer     default 1                 not null
);

-- 添加会员等级表注释
comment on table ums_member_level is '会员等级表';

-- 添加会员等级表列注释
comment on column ums_member_level.id is '主键id';
comment on column ums_member_level.name is '等级名称';
comment on column ums_member_level.level is '等级';
comment on column ums_member_level.growth_point is '升级所需成长值';
comment on column ums_member_level.discount_rate is '折扣率(0-100)';
comment on column ums_member_level.free_freight is '是否免运费';
comment on column ums_member_level.comment_extra is '是否可评论获取奖励';
comment on column ums_member_level.privileges is '会员特权json';
comment on column ums_member_level.remark is '备注';
comment on column ums_member_level.is_enabled is '是否启用';
comment on column ums_member_level.create_by is '创建人id';
comment on column ums_member_level.create_time is '创建时间';
comment on column ums_member_level.update_by is '更新人id';
comment on column ums_member_level.update_time is '更新时间';
comment on column ums_member_level.is_deleted is '是否删除(0:否,1:是)';

insert into ums_member_level (id, name, level, growth_point, discount_rate, free_freight, comment_extra, privileges, remark, is_enabled, create_by)
values (1, '普通会员', 1, 0, 100.00, 0, 0, '{
    "priority_service": 0,
    "birthday_gift": 0,
    "exclusive_price": 0,
    "vip_room": 0,
    "free_return": 0,
    "personal_butler": 0,
    "early_access": 0,
    "point_rate": 1.0
 }', '新注册会员默认等级', 1, 1),
       (2, '银卡会员', 2, 1000, 98.00, 0, 1, '{
    "priority_service": 1,
    "birthday_gift": 1,
    "exclusive_price": 0,
    "vip_room": 0,
    "free_return": 0,
    "personal_butler": 0,
    "early_access": 1,
    "point_rate": 1.2
 }', '消费满1000成长值可升级', 1, 1),
       (3, '金卡会员', 3, 3000, 95.00, 1, 1, '{
    "priority_service": 2,
    "birthday_gift": 2,
    "exclusive_price": 1,
    "vip_room": 0,
    "free_return": 1,
    "personal_butler": 0,
    "early_access": 2,
    "point_rate": 1.5
 }', '消费满3000成长值可升级', 1, 1),
       (4, '钻石会员', 4, 10000, 92.00, 1, 1, '{
    "priority_service": 2,
    "birthday_gift": 2,
    "exclusive_price": 1,
    "vip_room": 1,
    "free_return": 1,
    "personal_butler": 0,
    "early_access": 2,
    "point_rate": 2.0
 }', '消费满10000成长值可升级', 1, 1),
       (5, '黑金会员', 5, 50000, 88.00, 1, 1, '{
    "priority_service": 3,
    "birthday_gift": 3,
    "exclusive_price": 1,
    "vip_room": 1,
    "free_return": 1,
    "personal_butler": 1,
    "early_access": 3,
    "point_rate": 3.0
 }', '年消费达到50000可升级', 1, 1);

-- 会员登录记录建语句
drop table if exists ums_member_login_log;
create table ums_member_login_log
(
    id          bigserial primary key,
    member_id   bigint                                not null,
    member_ip   varchar                               not null,
    city        varchar                               not null,
    login_type  integer     default 1                 not null,
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
comment on column ums_member_login_log.login_type is '登录类型：0->pc；1->android;2->ios;3->小程序';
comment on column ums_member_login_log.province is '登录省份';
comment on column ums_member_login_log.create_time is '登录时间';

-- 会员积分记录表建语句
drop table if exists ums_member_points_log;
create table ums_member_points_log
(
    id            bigserial primary key,
    member_id     bigint                                not null,
    change_type   integer     default 1                 not null,
    change_points integer                               not null,
    source_type   integer     default 1                 not null,
    description   varchar                               not null,
    operate_man   varchar                               not null,
    operate_note  varchar                               not null,
    create_time   timestamptz default current_timestamp not null
);

-- 添加会员积分记录表注释
comment on table ums_member_points_log is '会员积分记录表';

-- 添加会员积分记录表列注释
comment on column ums_member_points_log.id is '';
comment on column ums_member_points_log.member_id is '会员id';
comment on column ums_member_points_log.change_type is '变更类型：1-添加积分，2-减少积分';
comment on column ums_member_points_log.change_points is '变更积分';
comment on column ums_member_points_log.source_type is '来源类型：0-其他，1-订单，2-活动，3-签到，4-管理员修改';
comment on column ums_member_points_log.description is '描述';
comment on column ums_member_points_log.operate_man is '操作人员';
comment on column ums_member_points_log.operate_note is '操作备注';
comment on column ums_member_points_log.create_time is '创建时间';

insert into ums_member_points_log (id, member_id, change_type, change_points, source_type, description, operate_man, operate_note)
values (1, 1001, 1, 100, 1, '订单完成增加积分', '管理员a', '订单id: 12345'),
       (2, 1002, 2, 50, 2, '活动参与减少积分', '管理员b', '活动id: 67890'),
       (3, 1003, 1, 200, 3, '签到增加积分', '管理员c', '签到id: 11223'),
       (4, 1004, 2, 30, 4, '管理员修改减少积分', '管理员d', '手动调整'),
       (5, 1005, 1, 150, 0, '其他来源增加积分', '管理员e', '特殊奖励');

-- 会员积分成长规则表建语句
drop table if exists ums_member_rule_setting;
create table ums_member_rule_setting
(
    id                  bigserial primary key,
    consume_per_point   bigint                                not null,
    low_order_amount    bigint                                not null,
    max_point_per_order integer                               not null,
    rule_type           integer     default 1                 not null,
    status              integer     default 1                 not null,
    create_by           varchar     default ''                not null,
    create_time         timestamptz default current_timestamp not null,
    update_by           varchar     default ''                not null,
    update_time         timestamptz                           null
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
comment on column ums_member_rule_setting.create_by is '创建人id';
comment on column ums_member_rule_setting.create_time is '创建时间';
comment on column ums_member_rule_setting.update_by is '更新人id';
comment on column ums_member_rule_setting.update_time is '更新时间';

-- 向会员积分成长规则表添加模拟数据
insert into ums_member_rule_setting (consume_per_point, low_order_amount, max_point_per_order, rule_type, status,create_by)
values
-- 积分规则
(100, 1000, 50, 0, 1, 1),  -- 每消费1元获取1积分，最低10元订单，每单最多50积分
(200, 2000, 100, 0, 1, 1), -- 每消费2元获取1积分，最低20元订单，每单最多100积分
(500, 5000, 200, 0, 0, 1), -- 每消费5元获取1积分，最低50元订单，每单最多200积分（禁用状态）
-- 成长值规则
(300, 3000, 30, 1, 1, 1),  -- 每消费3元获取1成长值，最低30元订单，每单最多30成长值
(500, 10000, 100, 1, 1, 1); -- 每消费5元获取1成长值，最低100元订单，每单最多100成长值

-- 会员签到记录表建语句
drop table if exists ums_member_sign_log;
create table ums_member_sign_log
(
    id            bigserial primary key,
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
comment on column ums_member_sign_log.member_id is '会员id';
comment on column ums_member_sign_log.sign_date is '签到日期';
comment on column ums_member_sign_log.continue_days is '连续签到天数';
comment on column ums_member_sign_log.points is '获得积分';
comment on column ums_member_sign_log.create_time is '';

insert into ums_member_sign_log (id, member_id, sign_date, continue_days, points)
values (1, 1001, '2024-01-01', 1, 10),
       (2, 1002, '2024-01-02', 2, 20),
       (3, 1003, '2024-01-03', 3, 30),
       (4, 1004, '2024-01-04', 4, 40),
       (5, 1005, '2024-01-05', 5, 50);

-- 会员统计信息建语句
drop table if exists ums_member_statistics_info;
create table ums_member_statistics_info
(
    id                    bigserial primary key,
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

insert into ums_member_statistics_info (id, member_id, consume_amount, order_count, coupon_count, comment_count, return_order_count, login_count, attend_count, fans_count, collect_product_count, collect_subject_count, collect_topic_count, collect_comment_count, invite_friend_count, recent_order_time)
values (1, 1001, 5000, 10, 5, 20, 2, 50, 10, 100, 15, 5, 3, 8, 2, '2024-01-01 10:00:00'),
       (2, 1002, 3000, 8, 3, 15, 1, 30, 5, 50, 10, 3, 2, 5, 1, '2024-01-02 11:00:00'),
       (3, 1003, 7000, 15, 7, 25, 3, 70, 20, 150, 20, 10, 5, 12, 4, '2024-01-03 12:00:00'),
       (4, 1004, 2000, 5, 2, 10, 0, 20, 3, 30, 5, 2, 1, 3, 0, '2024-01-04 13:00:00'),
       (5, 1005, 10000, 20, 10, 30, 5, 100, 25, 200, 25, 15, 8, 20, 5, '2024-01-05 14:00:00');


-- 用户标签表建语句
drop table if exists ums_member_tag;
create table ums_member_tag
(
    id                  bigserial primary key,
    tag_name            varchar                               not null,
    description         varchar                               not null,
    finish_order_count  integer                               not null,
    finish_order_amount numeric                               not null,
    status              integer     default 1                 not null,
    create_by           varchar     default ''                not null,
    create_time         timestamptz default current_timestamp not null,
    update_by           varchar     default ''                not null,
    update_time         timestamptz                           null,
    is_deleted          integer     default 1                 not null
);

-- 添加用户标签表注释
comment on table ums_member_tag is '用户标签表';

-- 添加用户标签表列注释
comment on column ums_member_tag.id is '主键id';
comment on column ums_member_tag.tag_name is '标签名称';
comment on column ums_member_tag.description is '标签描述';
comment on column ums_member_tag.finish_order_count is '自动打标签完成订单数量';
comment on column ums_member_tag.finish_order_amount is '自动打标签完成订单金额';
comment on column ums_member_tag.status is '状态：0-禁用，1-启用';
comment on column ums_member_tag.create_by is '创建人id';
comment on column ums_member_tag.create_time is '创建时间';
comment on column ums_member_tag.update_by is '更新人id';
comment on column ums_member_tag.update_time is '更新时间';
comment on column ums_member_tag.is_deleted is '是否删除(0:否,1:是)';

insert into ums_member_tag (id, tag_name, description, finish_order_count, finish_order_amount, status, create_by, is_deleted)
values (1, '新用户', '注册时间不超过一个月的用户', 0, 0.00, 1, 1, 0),
       (2, '忠实用户', '连续三个月每月都有订单的用户', 3, 300.00, 1, 1, 0),
       (3, '活跃用户', '最近一个月内完成5单以上的用户', 5, 500.00, 1, 1, 0),
       (4, '高消费用户', '累计消费金额超过1000元的用户', 10, 1000.00, 1, 1, 0);


-- 会员标签关联表建语句
drop table if exists ums_member_tag_relation;
create table ums_member_tag_relation
(
    id          bigserial primary key,
    member_id   bigint                                not null,
    tag_id      bigint                                not null,
    create_time timestamptz default current_timestamp not null
);

-- 添加会员标签关联表注释
comment on table ums_member_tag_relation is '会员标签关联表';

-- 添加会员标签关联表列注释
comment on column ums_member_tag_relation.id is '主键id';
comment on column ums_member_tag_relation.member_id is '会员id';
comment on column ums_member_tag_relation.tag_id is '标签id';
comment on column ums_member_tag_relation.create_time is '创建时间';
-- # 更新会员标签关联表的时机
--
-- # - 1.实时更新 ：
-- # - 当会员的行为（如下单、消费、签到等）发生变化时，立即更新会员的标签。
--
-- # - 2.定时任务 ：
-- # - 使用定时任务（如每天凌晨）批量更新会员的标签。

-- 会员任务表建语句
drop table if exists ums_member_task;
create table ums_member_task
(
    id             bigserial primary key,
    task_name      varchar                               not null,
    task_desc      varchar                               not null,
    task_growth    integer                               not null,
    task_integral  integer                               not null,
    task_type      integer     default 1                 not null,
    complete_count integer                               not null,
    reward_type    integer     default 1                 not null,
    reward_params  varchar                               not null,
    start_time     timestamptz                           not null,
    end_time       timestamptz                           not null,
    status         integer     default 1                 not null,
    sort           integer                               not null,
    create_by      varchar     default ''                not null,
    create_time    timestamptz default current_timestamp not null,
    update_by      varchar      default ''                not null,
    update_time    timestamptz                           null,
    is_deleted     integer     default 1                 not null
);

-- 添加会员任务表注释
comment on table ums_member_task is '会员任务表';

-- 添加会员任务表列注释
comment on column ums_member_task.id is '主键id';
comment on column ums_member_task.task_name is '任务名称';
comment on column ums_member_task.task_desc is '任务描述';
comment on column ums_member_task.task_growth is '赠送成长值';
comment on column ums_member_task.task_integral is '赠送积分';
comment on column ums_member_task.task_type is '任务类型：0-新手任务，1-日常任务，2-周常任务，3-月常任务';
comment on column ums_member_task.complete_count is '需要完成次数';
comment on column ums_member_task.reward_type is '奖励类型：0-积分成长值，1-优惠券，2-抽奖次数';
comment on column ums_member_task.reward_params is '奖励参数json';
comment on column ums_member_task.start_time is '任务开始时间';
comment on column ums_member_task.end_time is '任务结束时间';
comment on column ums_member_task.status is '状态：0-禁用，1-启用';
comment on column ums_member_task.sort is '排序';
comment on column ums_member_task.create_by is '创建人id';
comment on column ums_member_task.create_time is '创建时间';
comment on column ums_member_task.update_by is '更新人id';
comment on column ums_member_task.update_time is '更新时间';
comment on column ums_member_task.is_deleted is '是否删除(0:否,1:是)';

-- 插入会员任务数据
insert into ums_member_task (id, task_name, task_desc, task_growth, task_integral, task_type, complete_count, reward_type, reward_params, start_time, end_time, status, sort, create_by, is_deleted)
values (1, '注册奖励', '完成注册即可获得奖励', 100, 50, 0, 1, 0, '{"coupon_id": null}', '2024-01-01 00:00:00', '2024-12-31 23:59:59', 1, 1, 1, 0),
       (2, '每日签到', '每日签到可获得成长值和积分', 10, 5, 1, 1, 1, '{"coupon_id": null}', '2024-01-01 00:00:00', '2024-12-31 23:59:59', 1, 2, 1, 0),
       (3, '每周购物', '每周购物满100元可获得奖励', 50, 20, 2, 1, 2, '{"coupon_id": null}', '2024-01-01 00:00:00', '2024-12-31 23:59:59', 1, 3, 1, 0),
       (4, '每月评价', '每月评价满5次可获得奖励', 30, 15, 3, 5, 1, '{"coupon_id": null}', '2024-01-01 00:00:00', '2024-12-31 23:59:59', 1, 4, 1, 0);

-- 会员任务关联表建语句
drop table if exists ums_member_task_relation;
create table ums_member_task_relation
(
    id          bigserial primary key,
    member_id   bigint                                not null,
    task_id     bigint                                not null,
    create_time timestamptz default current_timestamp not null
);

-- 添加会员任务关联表注释
comment on table ums_member_task_relation is '会员任务关联表';

-- 添加会员任务关联表列注释
comment on column ums_member_task_relation.id is '主键id';
comment on column ums_member_task_relation.member_id is '会员id';
comment on column ums_member_task_relation.task_id is '任务id';
comment on column ums_member_task_relation.create_time is '创建时间';
-- # 注册奖励和每日签到的任务，触发时机：
-- #  1.实时触发 ：
-- #    - 当会员注册奖励和每日签到时，立即检查该会员是否达到任务要求。如果达到，则立即发放奖励。
--
--
-- # 每周购物和每月评价的任务，触发时机：
-- #
-- # 1. 每周定时任务 ：
-- #    - 设置一个每周定时任务（例如每周一凌晨），检查所有会员在过去一周内的购物情况。如果会员在过去一周内的购物金额或订单数量达到任务要求，则立即发放奖励。
--
-- # 2. 实时触发 ：
-- #    - 当会员完成一次购物时，立即检查该会员在当前周的购物总额或订单数量是否达到任务要求。如果达到，则立即发放奖励。

