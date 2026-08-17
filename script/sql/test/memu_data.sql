DO $$
    DECLARE
        parentId bigint;
    BEGIN
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
            constraint menu_name unique (menu_name, parent_id)
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


        insert into sys_menu (id, menu_name, menu_type, status, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) values (nextval('sys_menu_id_seq'), '目录', 1, 1, 1, 0, '/main', '', '', '目录');
        insert into sys_menu (id, menu_name, menu_type, status, menu_sort, parent_id, menu_url, api_url, menu_icon, remark, vue_icon) values (nextval('sys_menu_id_seq'), '欢迎', 1, 1, 2, 1, '/welcome', '', 'SmileOutlined', '欢迎','HomeFilled');
        insert into sys_menu (id, menu_name, menu_type, status, menu_sort, parent_id, menu_url, api_url, menu_icon, remark, vue_icon) values (nextval('sys_menu_id_seq'), '系统管理', 1, 1, 3, 1, '/system', '', 'SettingOutlined', '系统管理','Setting');
        insert into sys_menu (id, menu_name, menu_type, status, menu_sort, parent_id, menu_url, api_url, menu_icon, remark, vue_icon) values (nextval('sys_menu_id_seq'), '日志管理', 1, 1, 4, 1, '/log', '', 'DeleteOutlined', '日志管理','Delete');
        insert into sys_menu (id, menu_name, menu_type, status, menu_sort, parent_id, menu_url, api_url, menu_icon, remark, vue_icon) values (nextval('sys_menu_id_seq'), '会员管理', 1, 1, 5, 1, '/ums', '', 'FrownOutlined', '会员管理','Expand');
        insert into sys_menu (id, menu_name, menu_type, status, menu_sort, parent_id, menu_url, api_url, menu_icon, remark, vue_icon) values (nextval('sys_menu_id_seq'), '商品管理', 1, 1, 6, 1, '/pms', '', 'GiftOutlined', '商品管理','Coffee');
        insert into sys_menu (id, menu_name, menu_type, status, menu_sort, parent_id, menu_url, api_url, menu_icon, remark, vue_icon) values (nextval('sys_menu_id_seq'), '订单管理', 1, 1, 7, 1, '/oms', '', 'DollarCircleOutlined', '订单管理','ShoppingCart');
        insert into sys_menu (id, menu_name, menu_type, status, menu_sort, parent_id, menu_url, api_url, menu_icon, remark, vue_icon) values (nextval('sys_menu_id_seq'), '营销管理', 1, 1, 8, 1, '/sms', '', 'AlertOutlined', '营销管理','Postcard');
        insert into sys_menu (id, menu_name, menu_type, status, menu_sort, parent_id, menu_url, api_url, menu_icon, remark, vue_icon) values (nextval('sys_menu_id_seq'), '内容管理', 1, 1, 9, 1, '/cms', '', 'SettingOutlined', '内容管理','Files');


        -- 配置菜单信息权限
        INSERT INTO sys_menu (id, menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES (nextval('sys_menu_id_seq'), '菜单列表', 2, 3, 3, '/system/menu/list', '', '', '菜单信息管理');

        SELECT currval('sys_menu_id_seq') INTO parentId;

        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('添加菜单信息', 3, 1, parentId, '', '/api/sys/menu/addMenu', '', '添加菜单信息');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('删除菜单信息', 3, 2, parentId, '', '/api/sys/menu/deleteMenu', '', '删除菜单信息');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新菜单信息', 3, 3, parentId, '', '/api/sys/menu/updateMenu', '', '更新菜单信息');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新菜单信息状态', 3, 4, parentId, '', '/api/sys/menu/updateMenuStatus', '', '更新菜单信息状态');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询菜单信息详情', 3, 5, parentId, '', '/api/sys/menu/queryMenuDetail', '', '查询菜单信息详情');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询菜单信息列表', 3, 6, parentId, '', '/api/sys/menu/queryMenuList', '', '查询菜单信息列表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询菜单资源', 3, 6, parentId, '', '/api/sys/menu/queryMenuResourceList', '', '查询菜单资源');

        -- 配置用户信息权限
        INSERT INTO sys_menu (id, menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES (nextval('sys_menu_id_seq'), '用户列表', 2, 1, 3, '/system/user/list', '', '', '用户信息管理');

        SELECT currval('sys_menu_id_seq') INTO parentId;

        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('添加用户信息', 3, 1, parentId, '', '/api/sys/user/addUser,/api/sys/user/queryDeptAndPostList', '', '添加用户信息');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('删除用户信息', 3, 2, parentId, '', '/api/sys/user/deleteUser', '', '删除用户信息');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新用户信息', 3, 3, parentId, '', '/api/sys/user/updateUser,/api/sys/user/queryDeptAndPostList', '', '更新用户信息');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新用户信息状态', 3, 4, parentId, '', '/api/sys/user/updateUserStatus', '', '更新用户信息状态');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询用户信息详情', 3, 5, parentId, '', '/api/sys/user/queryUserDetail', '', '查询用户信息详情');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询用户信息列表', 3, 6, parentId, '', '/api/sys/user/queryUserList', '', '查询用户信息列表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('上传文件', 3, 6, parentId, '', '/api/sys/upload', '', '上传文件');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询用户角色关联', 3, 6, parentId, '', '/api/sys/user/queryUserRoleList', '', '查询用户角色关联');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新用户角色关联', 3, 6, parentId, '', '/api/sys/user/updateUserRoleList', '', '更新用户角色关联');



        -- 配置角色信息权限
        INSERT INTO sys_menu (id, menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES (nextval('sys_menu_id_seq'), '角色列表', 2, 2, 3, '/system/role/list', '', '', '角色信息管理');

        SELECT currval('sys_menu_id_seq') INTO parentId;

        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('添加角色信息', 3, 1, parentId, '', '/api/sys/role/addRole', '', '添加角色信息');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('删除角色信息', 3, 2, parentId, '', '/api/sys/role/deleteRole', '', '删除角色信息');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新角色信息', 3, 3, parentId, '', '/api/sys/role/updateRole', '', '更新角色信息');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新角色信息状态', 3, 4, parentId, '', '/api/sys/role/updateRoleStatus', '', '更新角色信息状态');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询角色信息详情', 3, 5, parentId, '', '/api/sys/role/queryRoleDetail', '', '查询角色信息详情');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询角色信息列表', 3, 6, parentId, '', '/api/sys/role/queryRoleList', '', '查询角色信息列表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('根据角色查询菜单', 3, 6, parentId, '', '/api/sys/role/queryRoleMenuList', '', '根据角色查询菜单');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('根据角色更新菜单', 3, 6, parentId, '', '/api/sys/role/updateRoleMenuList', '', '根据角色更新菜单');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询角色的用户关联', 3, 6, parentId, '', '/api/sys/role/queryRoleUserList', '', '查询角色的用户关联');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新角色的用户关联', 3, 6, parentId, '', '/api/sys/role/updateRoleUserList', '', '更新角色的用户关联');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('取消授权', 3, 6, parentId, '', '/api/sys/role/cancelAuthorization', '', '取消授权');
        -- 配置岗位信息表权限
        INSERT INTO sys_menu (id, menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES (nextval('sys_menu_id_seq'), '岗位管理', 2, 6, 3, '/system/post/list', '', '', '岗位信息');

        SELECT currval('sys_menu_id_seq') INTO parentId;

        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('添加岗位信息', 3, 1, parentId, '', '/api/sys/post/addPost', '', '添加岗位信息');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('删除岗位信息', 3, 2, parentId, '', '/api/sys/post/deletePost', '', '删除岗位信息');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新岗位信息', 3, 3, parentId, '', '/api/sys/post/updatePost', '', '更新岗位信息');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新岗位信息表状态', 3, 4, parentId, '', '/api/sys/post/updatePostStatus', '', '更新岗位信息表状态');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询岗位信息表详情', 3, 5, parentId, '', '/api/sys/post/queryPostDetail', '', '查询岗位信息表详情');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询岗位信息表列表', 3, 6, parentId, '', '/api/sys/post/queryPostList', '', '查询岗位信息表列表');

        -- 配置部门表权限
        INSERT INTO sys_menu (id, menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES (nextval('sys_menu_id_seq'), '部门管理', 2, 4, 3, '/system/dept/list', '', '', '部门');

        SELECT currval('sys_menu_id_seq') INTO parentId;

        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('添加部门', 3, 1, parentId, '', '/api/sys/dept/addDept', '', '添加部门');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('删除部门', 3, 2, parentId, '', '/api/sys/dept/deleteDept', '', '删除部门');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新部门', 3, 3, parentId, '', '/api/sys/dept/updateDept', '', '更新部门');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新部门表状态', 3, 4, parentId, '', '/api/sys/dept/updateDeptStatus', '', '更新部门表状态');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询部门表详情', 3, 5, parentId, '', '/api/sys/dept/queryDeptDetail', '', '查询部门表详情');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询部门表列表', 3, 6, parentId, '', '/api/sys/dept/queryDeptList', '', '查询部门表列表');

        -- 配置字典数据表权限
        INSERT INTO sys_menu (id, menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark, visible) VALUES (nextval('sys_menu_id_seq'), '字典数据', 2, 1, 3, '/sys/dictData', '', '', '字典数据',0);

        SELECT currval('sys_menu_id_seq') INTO parentId;

        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('添加字典数据', 3, 1, parentId, '', '/api/sys/dictData/addDictData', '', '添加字典数据');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('删除字典数据', 3, 2, parentId, '', '/api/sys/dictData/deleteDictData', '', '删除字典数据');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新字典数据', 3, 3, parentId, '', '/api/sys/dictData/updateDictData', '', '更新字典数据');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新字典数据表状态', 3, 4, parentId, '', '/api/sys/dictData/updateDictDataStatus', '', '更新字典数据表状态');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询字典数据表详情', 3, 5, parentId, '', '/api/sys/dictData/queryDictDataDetail', '', '查询字典数据表详情');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询字典数据表列表', 3, 6, parentId, '', '/api/sys/dictData/queryDictDataList', '', '查询字典数据表列表');

        -- 配置字典类型表权限
        INSERT INTO sys_menu (id, menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES (nextval('sys_menu_id_seq'), '字典管理', 2, 5, 3, '/system/dict/list', '', '', '字典类型');

        SELECT currval('sys_menu_id_seq') INTO parentId;

        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('添加字典类型', 3, 1, parentId, '', '/api/sys/dictType/addDictType', '', '添加字典类型');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('删除字典类型', 3, 2, parentId, '', '/api/sys/dictType/deleteDictType', '', '删除字典类型');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新字典类型', 3, 3, parentId, '', '/api/sys/dictType/updateDictType', '', '更新字典类型');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新字典类型表状态', 3, 4, parentId, '', '/api/sys/dictType/updateDictTypeStatus', '', '更新字典类型表状态');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询字典类型表详情', 3, 5, parentId, '', '/api/sys/dictType/queryDictTypeDetail', '', '查询字典类型表详情');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询字典类型表列表', 3, 6, parentId, '', '/api/sys/dictType/queryDictTypeList', '', '查询字典类型表列表');

        -- 配置系统访问记录权限
        INSERT INTO sys_menu (id, menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES (nextval('sys_menu_id_seq'), '登录日志', 2, 1, 4, '/log/loginLog/list', '', '', '系统访问记录管理');

        SELECT currval('sys_menu_id_seq') INTO parentId;

        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('删除系统访问记录', 3, 2, parentId, '', '/api/sys/log/deleteLoginLog', '', '删除系统访问记录');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询系统访问记录详情', 3, 5, parentId, '', '/api/sys/log/queryLoginLogDetail', '', '查询系统访问记录详情');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询系统访问记录列表', 3, 6, parentId, '', '/api/sys/log/queryLoginLogList', '', '查询系统访问记录列表');

        -- 配置通知公告表权限
        INSERT INTO sys_menu (id, menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES (nextval('sys_menu_id_seq'), '通知公告', 2, 7, 3, '/system/notice/list', '', '', '通知公告');

        SELECT currval('sys_menu_id_seq') INTO parentId;

        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('添加通知公告', 3, 1, parentId, '', '/api/sys/notice/addNotice', '', '添加通知公告');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('删除通知公告', 3, 2, parentId, '', '/api/sys/notice/deleteNotice', '', '删除通知公告');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新通知公告', 3, 3, parentId, '', '/api/sys/notice/updateNotice', '', '更新通知公告');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新通知公告表状态', 3, 4, parentId, '', '/api/sys/notice/updateNoticeStatus', '', '更新通知公告表状态');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询通知公告表详情', 3, 5, parentId, '', '/api/sys/notice/queryNoticeDetail', '', '查询通知公告表详情');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询通知公告表列表', 3, 6, parentId, '', '/api/sys/notice/queryNoticeList', '', '查询通知公告表列表');

        -- 配置操作日志记录权限
        INSERT INTO sys_menu (id, menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES (nextval('sys_menu_id_seq'), '操作日志', 2, 2, 4, '/log/sysLog/list', '', '', '操作日志记录管理');

        SELECT currval('sys_menu_id_seq') INTO parentId;

        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('删除操作日志记录', 3, 2, parentId, '', '/api/sys/log/deleteOperateLog', '', '删除操作日志记录');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询操作日志记录详情', 3, 5, parentId, '', '/api/sys/log/queryOperateLogDetail', '', '查询操作日志记录详情');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询操作日志记录列表', 3, 6, parentId, '', '/api/sys/log/queryOperateLogList', '', '查询操作日志记录列表');
    END $$;

DO $$
    DECLARE
        parentId bigint;
    BEGIN

        -- 配置积分消费设置权限
        INSERT INTO sys_menu (id, menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES (nextval('sys_menu_id_seq'), '积分消费', 2, 7, 5, '/ums/integrationConsumeSetting/list', '', '', '积分消费设置管理');

        SELECT currval('sys_menu_id_seq') INTO parentId;

        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('添加积分消费设置', 3, 1, parentId, '', '/api/ums/consumeSetting/addMemberConsumeSetting', '', '添加积分消费设置');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('删除积分消费设置', 3, 2, parentId, '', '/api/ums/consumeSetting/deleteMemberConsumeSetting', '', '删除积分消费设置');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新积分消费设置', 3, 3, parentId, '', '/api/ums/consumeSetting/updateMemberConsumeSetting', '', '更新积分消费设置');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新积分消费设置状态', 3, 4, parentId, '', '/api/ums/consumeSetting/updateMemberConsumeSettingStatus', '', '更新积分消费设置状态');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询积分消费设置详情', 3, 5, parentId, '', '/api/ums/consumeSetting/queryMemberConsumeSettingDetail', '', '查询积分消费设置详情');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询积分消费设置列表', 3, 6, parentId, '', '/api/ums/consumeSetting/queryMemberConsumeSettingList', '', '查询积分消费设置列表');

        -- 配置会员信息表权限
        INSERT INTO sys_menu (id, menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES (nextval('sys_menu_id_seq'), '会员列表', 2, 1, 5, '/ums/member/list', '', '', '会员信息');

        SELECT currval('sys_menu_id_seq') INTO parentId;

        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('删除会员信息', 3, 2, parentId, '', '/api/ums/member/deleteMember', '', '删除会员信息');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新会员信息', 3, 3, parentId, '', '/api/ums/member/updateMember', '', '更新会员信息');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新会员信息表状态', 3, 4, parentId, '', '/api/ums/member/updateMemberStatus', '', '更新会员信息表状态');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询会员信息表详情', 3, 5, parentId, '', '/api/ums/member/queryMemberDetail', '', '查询会员信息表详情');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询会员信息表列表', 3, 6, parentId, '', '/api/ums/member/queryMemberList', '', '查询会员信息表列表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询会员会员地址列表', 3, 6, parentId, '', '/api/ums/address/queryMemberAddressList', '', '查询会员会员地址列表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询会员会员登录记录', 3, 6, parentId, '', '/api/ums/member/queryMemberLoginLogList', '', '查询会员会员登录记录');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询会员成长值', 3, 6, parentId, '', '/api/ums/growthLog/queryMemberGrowthLogList', '', '查询会员成长值变化');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询会员积分', 3, 6, parentId, '', '/api/ums/pointsLog/queryMemberPointsLogList', '', '查询会员积分变化');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询会员签到列表', 3, 6, parentId, '', '/api/ums/sign/queryMemberSignLogList', '', '查询会员签到列表');

        -- 配置会员等级表权限
        INSERT INTO sys_menu (id, menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES (nextval('sys_menu_id_seq'), '会员等级', 2, 2, 5, '/ums/memberLevel/list', '', '', '会员等级');

        SELECT currval('sys_menu_id_seq') INTO parentId;

        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('添加会员等级', 3, 1, parentId, '', '/api/ums/member/addMemberLevel', '', '添加会员等级');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('删除会员等级', 3, 2, parentId, '', '/api/ums/member/deleteMemberLevel', '', '删除会员等级');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新会员等级', 3, 3, parentId, '', '/api/ums/member/updateMemberLevel', '', '更新会员等级');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新会员等级表状态', 3, 4, parentId, '', '/api/ums/member/updateMemberLevelStatus', '', '更新会员等级表状态');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询会员等级表详情', 3, 5, parentId, '', '/api/ums/member/queryMemberLevelDetail', '', '查询会员等级表详情');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询会员等级表列表', 3, 6, parentId, '', '/api/ums/member/queryMemberLevelList', '', '查询会员等级表列表');

        -- 配置会员积分成长规则表权限
        INSERT INTO sys_menu (id, menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES (nextval('sys_menu_id_seq'), '成长规则', 2, 6, 5, '/ums/memberRule/list', '', '', '会员积分成长规则');

        SELECT currval('sys_menu_id_seq') INTO parentId;

        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('添加会员积分成长规则', 3, 1, parentId, '', '/api/ums/ruleSetting/addMemberRuleSetting', '', '添加会员积分成长规则');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('删除会员积分成长规则', 3, 2, parentId, '', '/api/ums/ruleSetting/deleteMemberRuleSetting', '', '删除会员积分成长规则');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新会员积分成长规则', 3, 3, parentId, '', '/api/ums/ruleSetting/updateMemberRuleSetting', '', '更新会员积分成长规则');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新会员积分成长规则表状态', 3, 4, parentId, '', '/api/ums/ruleSetting/updateMemberRuleSettingStatus', '', '更新会员积分成长规则表状态');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询会员积分成长规则表详情', 3, 5, parentId, '', '/api/ums/ruleSetting/queryMemberRuleSettingDetail', '', '查询会员积分成长规则表详情');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询会员积分成长规则表列表', 3, 6, parentId, '', '/api/ums/ruleSetting/queryMemberRuleSettingList', '', '查询会员积分成长规则表列表');

        -- 配置会员统计信息权限
        INSERT INTO sys_menu (id, menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES (nextval('sys_menu_id_seq'), '统计信息', 2, 5, 5, '/ums/statistics/list', '', '', '会员统计信息管理');

        SELECT currval('sys_menu_id_seq') INTO parentId;

        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询会员统计信息详情', 3, 5, parentId, '', '/api/ums/memberStatisticsInfo/queryMemberStatisticsInfoDetail', '', '查询会员统计信息详情');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询会员统计信息列表', 3, 6, parentId, '', '/api/ums/memberStatisticsInfo/queryMemberStatisticsInfoList', '', '查询会员统计信息列表');

        -- 配置用户标签表权限
        INSERT INTO sys_menu (id, menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES (nextval('sys_menu_id_seq'), '会员标签', 2, 4, 5, '/ums/memberTag/list', '', '', '用户标签');

        SELECT currval('sys_menu_id_seq') INTO parentId;

        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('添加用户标签', 3, 1, parentId, '', '/api/ums/tag/addMemberTag', '', '添加用户标签');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('删除用户标签', 3, 2, parentId, '', '/api/ums/tag/deleteMemberTag', '', '删除用户标签');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新用户标签', 3, 3, parentId, '', '/api/ums/tag/updateMemberTag', '', '更新用户标签');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新用户标签表状态', 3, 4, parentId, '', '/api/ums/tag/updateMemberTagStatus', '', '更新用户标签表状态');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询用户标签表详情', 3, 5, parentId, '', '/api/ums/tag/queryMemberTagDetail', '', '查询用户标签表详情');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询用户标签表列表', 3, 6, parentId, '', '/api/ums/tag/queryMemberTagList', '', '查询用户标签表列表');

        -- 配置会员任务表权限
        INSERT INTO sys_menu (id, menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES (nextval('sys_menu_id_seq'), '会员任务', 2, 3, 5, '/ums/memberTask/list', '', '', '会员任务');

        SELECT currval('sys_menu_id_seq') INTO parentId;

        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('添加会员任务', 3, 1, parentId, '', '/api/ums/task/addMemberTask', '', '添加会员任务');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('删除会员任务', 3, 2, parentId, '', '/api/ums/task/deleteMemberTask', '', '删除会员任务');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新会员任务', 3, 3, parentId, '', '/api/ums/task/updateMemberTask', '', '更新会员任务');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新会员任务表状态', 3, 4, parentId, '', '/api/ums/task/updateMemberTaskStatus', '', '更新会员任务表状态');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询会员任务表详情', 3, 5, parentId, '', '/api/ums/task/queryMemberTaskDetail', '', '查询会员任务表详情');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询会员任务表列表', 3, 6, parentId, '', '/api/ums/task/queryMemberTaskList', '', '查询会员任务表列表');


        END $$;


DO $$
    DECLARE
        parentId bigint;
    BEGIN

        -- 配置商品属性表权限
        INSERT INTO sys_menu (id, menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark, visible) VALUES (nextval('sys_menu_id_seq'), '商品属性', 2, 1, 6, '/pms/attribute/list', '', '', '商品属性',0);

        SELECT currval('sys_menu_id_seq') INTO parentId;

        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('添加商品属性', 3, 1, parentId, '', '/api/pms/attribute/addProductAttribute', '', '添加商品属性');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('删除商品属性', 3, 2, parentId, '', '/api/pms/attribute/deleteProductAttribute', '', '删除商品属性');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新商品属性', 3, 3, parentId, '', '/api/pms/attribute/updateProductAttribute', '', '更新商品属性');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新商品属性表状态', 3, 4, parentId, '', '/api/pms/attribute/updateProductAttributeStatus', '', '更新商品属性表状态');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询商品属性表详情', 3, 5, parentId, '', '/api/pms/attribute/queryProductAttributeDetail', '', '查询商品属性表详情');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询商品属性表列表', 3, 6, parentId, '', '/api/pms/attribute/queryProductAttributeList', '', '查询商品属性表列表');

        -- 配置商品属性分组表权限
        INSERT INTO sys_menu (id, menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES (nextval('sys_menu_id_seq'), '属性分组', 2, 3, 6, '/pms/attributeGroup/list', '', '', '商品属性分组');

        SELECT currval('sys_menu_id_seq') INTO parentId;

        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('添加商品属性分组', 3, 1, parentId, '', '/api/pms/attributeGroup/addProductAttributeGroup', '', '添加商品属性分组');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('删除商品属性分组', 3, 2, parentId, '', '/api/pms/attributeGroup/deleteProductAttributeGroup', '', '删除商品属性分组');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新商品属性分组', 3, 3, parentId, '', '/api/pms/attributeGroup/updateProductAttributeGroup', '', '更新商品属性分组');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新商品属性分组表状态', 3, 4, parentId, '', '/api/pms/attributeGroup/updateProductAttributeGroupStatus', '', '更新商品属性分组表状态');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询商品属性分组表详情', 3, 5, parentId, '', '/api/pms/attributeGroup/queryProductAttributeGroupDetail', '', '查询商品属性分组表详情');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询商品属性分组表列表', 3, 6, parentId, '', '/api/pms/attributeGroup/queryProductAttributeGroupList', '', '查询商品属性分组表列表');

        -- 配置商品品牌权限
        INSERT INTO sys_menu (id, menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES (nextval('sys_menu_id_seq'), '商品品牌', 2, 2, 6, '/pms/productBrand/list', '', '', '商品品牌管理');

        SELECT currval('sys_menu_id_seq') INTO parentId;

        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('添加商品品牌', 3, 1, parentId, '', '/api/pms/brand/addProductBrand', '', '添加商品品牌');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('删除商品品牌', 3, 2, parentId, '', '/api/pms/brand/deleteProductBrand', '', '删除商品品牌');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新商品品牌', 3, 3, parentId, '', '/api/pms/brand/updateProductBrand', '', '更新商品品牌');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新商品品牌状态', 3, 4, parentId, '', '/api/pms/brand/updateProductBrandStatus', '', '更新商品品牌状态');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询商品品牌详情', 3, 5, parentId, '', '/api/pms/brand/queryProductBrandDetail', '', '查询商品品牌详情');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询商品品牌列表', 3, 6, parentId, '', '/api/pms/brand/queryProductBrandList', '', '查询商品品牌列表');

        -- 配置产品分类权限
        INSERT INTO sys_menu (id, menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES (nextval('sys_menu_id_seq'), '商品分类', 2, 1, 6, '/pms/productCategory/list', '', '', '产品分类管理');

        SELECT currval('sys_menu_id_seq') INTO parentId;

        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('添加产品分类', 3, 1, parentId, '', '/api/pms/category/addProductCategory', '', '添加产品分类');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('删除产品分类', 3, 2, parentId, '', '/api/pms/category/deleteProductCategory', '', '删除产品分类');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新产品分类', 3, 3, parentId, '', '/api/pms/category/updateProductCategory', '', '更新产品分类');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新产品分类状态', 3, 4, parentId, '', '/api/pms/category/updateProductCategoryStatus', '', '更新产品分类状态');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询产品分类详情', 3, 5, parentId, '', '/api/pms/category/queryProductCategoryDetail', '', '查询产品分类详情');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询产品分类列表', 3, 6, parentId, '', '/api/pms/category/queryProductCategoryList', '', '查询产品分类列表');


        -- 配置运费模版权限
        INSERT INTO sys_menu (id, menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES (nextval('sys_menu_id_seq'), '运费模版', 2, 9, 6, '/pms/feightTemplate', '', '', '运费模版管理');

        SELECT currval('sys_menu_id_seq') INTO parentId;

        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('添加运费模版', 3, 1, parentId, '', '/api/pms/feightTemplate/addFeightTemplate', '', '添加运费模版');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('删除运费模版', 3, 2, parentId, '', '/api/pms/feightTemplate/deleteFeightTemplate', '', '删除运费模版');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新运费模版', 3, 3, parentId, '', '/api/pms/feightTemplate/updateFeightTemplate', '', '更新运费模版');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新运费模版状态', 3, 4, parentId, '', '/api/pms/feightTemplate/updateFeightTemplateStatus', '', '更新运费模版状态');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询运费模版详情', 3, 5, parentId, '', '/api/pms/feightTemplate/queryFeightTemplateDetail', '', '查询运费模版详情');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询运费模版列表', 3, 6, parentId, '', '/api/pms/feightTemplate/queryFeightTemplateList', '', '查询运费模版列表');

        -- 配置商品sku表权限
        INSERT INTO sys_menu (id, menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark, visible) VALUES (nextval('sys_menu_id_seq'), '商品sku', 2, 8, 6, '/pms/productSku/list', '', '', '商品sku',0);

        SELECT currval('sys_menu_id_seq') INTO parentId;

        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('添加商品sku', 3, 1, parentId, '', '/api/pms/product/addProductSku', '', '添加商品sku');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('删除商品sku', 3, 2, parentId, '', '/api/pms/product/deleteProductSku', '', '删除商品sku');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新商品sku', 3, 3, parentId, '', '/api/pms/product/updateProductSku', '', '更新商品sku');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新商品sku表状态', 3, 4, parentId, '', '/api/pms/product/updateProductSkuStatus', '', '更新商品sku表状态');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询商品sku表详情', 3, 5, parentId, '', '/api/pms/product/queryProductSkuDetail', '', '查询商品sku表详情');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询商品sku表列表', 3, 6, parentId, '', '/api/pms/product/queryProductSkuList', '', '查询商品sku表列表');

        -- 配置商品规格表权限
        INSERT INTO sys_menu (id, menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES (nextval('sys_menu_id_seq'), '商品规格', 2, 5, 6, '/pms/productSpec/list', '', '', '商品规格');

        SELECT currval('sys_menu_id_seq') INTO parentId;

        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('添加商品规格', 3, 1, parentId, '', '/api/pms/productSpec/addSpec', '', '添加商品规格');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('删除商品规格', 3, 2, parentId, '', '/api/pms/productSpec/deleteSpec', '', '删除商品规格');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新商品规格', 3, 3, parentId, '', '/api/pms/productSpec/updateSpec', '', '更新商品规格');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新商品规格表状态', 3, 4, parentId, '', '/api/pms/productSpec/updateSpecStatus', '', '更新商品规格表状态');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询商品规格表详情', 3, 5, parentId, '', '/api/pms/productSpec/querySpecDetail', '', '查询商品规格表详情');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询商品规格表列表', 3, 6, parentId, '', '/api/pms/productSpec/querySpecList', '', '查询商品规格表列表');

        -- 配置商品规格值表权限
        INSERT INTO sys_menu (id, menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark,visible) VALUES (nextval('sys_menu_id_seq'), '商品规格值', 2, 6, 6, '/pms/productSpecValue', '', '', '商品规格值',0);

        SELECT currval('sys_menu_id_seq') INTO parentId;

        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('添加商品规格值', 3, 1, parentId, '', '/api/pms/productSpecValue/addSpecValue', '', '添加商品规格值');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('删除商品规格值', 3, 2, parentId, '', '/api/pms/productSpecValue/deleteSpecValue', '', '删除商品规格值');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新商品规格值', 3, 3, parentId, '', '/api/pms/productSpecValue/updateSpecValue', '', '更新商品规格值');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新商品规格值表状态', 3, 4, parentId, '', '/api/pms/productSpecValue/updateSpecValueStatus', '', '更新商品规格值表状态');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询商品规格值表详情', 3, 5, parentId, '', '/api/pms/productSpecValue/querySpecValueDetail', '', '查询商品规格值表详情');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询商品规格值表列表', 3, 6, parentId, '', '/api/pms/productSpecValue/querySpecValueList', '', '查询商品规格值表列表');

        -- 配置商品spu表权限
        INSERT INTO sys_menu (id, menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES (nextval('sys_menu_id_seq'), '商品列表', 2, 7, 6, '/pms/productSpu/list', '', '', '商品spu');

        SELECT currval('sys_menu_id_seq') INTO parentId;

        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('添加商品spu', 3, 1, parentId, '', '/api/pms/product/addProductSpu', '', '添加商品spu');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('删除商品spu', 3, 2, parentId, '', '/api/pms/product/deleteProductSpu', '', '删除商品spu');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新商品spu', 3, 3, parentId, '', '/api/pms/product/updateProductSpu', '', '更新商品spu');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新商品spu表状态', 3, 4, parentId, '', '/api/pms/product/updateProductSpuStatus', '', '更新商品spu表状态');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询商品spu表详情', 3, 5, parentId, '', '/api/pms/product/queryProductSpuDetail', '', '查询商品spu表详情');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询商品spu表列表', 3, 6, parentId, '', '/api/pms/product/queryProductSpuList', '', '查询商品spu表列表');
    END $$;


DO $$
    DECLARE
        parentId bigint;
    BEGIN

        -- 配置公司收发货地址表权限
        INSERT INTO sys_menu (id, menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES (nextval('sys_menu_id_seq'), '地址管理', 2, 3, 7, '/oms/companyAddress/list', '', '', '公司收发货地址');

        SELECT currval('sys_menu_id_seq') INTO parentId;

        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('添加公司收发货地址', 3, 1, parentId, '', '/api/oms/companyAddress/addCompanyAddress', '', '添加公司收发货地址');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('删除公司收发货地址', 3, 2, parentId, '', '/api/oms/companyAddress/deleteCompanyAddress', '', '删除公司收发货地址');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新公司收发货地址', 3, 3, parentId, '', '/api/oms/companyAddress/updateCompanyAddress', '', '更新公司收发货地址');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新公司收发货地址表状态', 3, 4, parentId, '', '/api/oms/companyAddress/updateCompanyAddressStatus', '', '更新公司收发货地址表状态');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询公司收发货地址表详情', 3, 5, parentId, '', '/api/oms/companyAddress/queryCompanyAddressDetail', '', '查询公司收发货地址表详情');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询公司收发货地址表列表', 3, 6, parentId, '', '/api/oms/companyAddress/queryCompanyAddressList', '', '查询公司收发货地址表列表');


        -- 配置订单主表权限
        INSERT INTO sys_menu (id, menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES (nextval('sys_menu_id_seq'), '订单列表', 2, 5, 7, '/oms/order/list', '', '', '订单主');

        SELECT currval('sys_menu_id_seq') INTO parentId;

        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新订单主', 3, 3, parentId, '', '/api/oms/order/updateOrderMain', '', '更新订单主');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新订单主表状态', 3, 4, parentId, '', '/api/oms/order/updateOrderMainStatus', '', '更新订单主表状态');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询订单主表详情', 3, 5, parentId, '', '/api/oms/order/queryOrderMainDetail', '', '查询订单主表详情');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询订单主表列表', 3, 6, parentId, '', '/api/oms/order/queryOrderMainList', '', '查询订单主表列表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新订单收货地址', 3, 6, parentId, '', '/api/oms/orderDelivery/updateOrderDelivery', '', '更新订单收货地址');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询订单收货信息', 3, 6, parentId, '', '/api/oms/orderDelivery/queryOrderDeliveryDetail', '', '查询订单收货信息');

        -- 配置退货/售后主表权限
        INSERT INTO sys_menu (id, menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES (nextval('sys_menu_id_seq'), '退货列表', 2, 4, 7, '/oms/orderReturnApply/list', '', '', '退货/售后主');

        SELECT currval('sys_menu_id_seq') INTO parentId;

        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('添加退货/售后主', 3, 1, parentId, '', '/api/oms/orderReturn/addOrderReturn', '', '添加退货/售后主');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('删除退货/售后主', 3, 2, parentId, '', '/api/oms/orderReturn/deleteOrderReturn', '', '删除退货/售后主');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新退货/售后主', 3, 3, parentId, '', '/api/oms/orderReturn/updateOrderReturn', '', '更新退货/售后主');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新退货/售后主表状态', 3, 4, parentId, '', '/api/oms/orderReturn/updateOrderReturnStatus', '', '更新退货/售后主表状态');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询退货/售后主表详情', 3, 5, parentId, '', '/api/oms/orderReturn/queryOrderReturnDetail', '', '查询退货/售后主表详情');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询退货/售后主表列表', 3, 6, parentId, '', '/api/oms/orderReturn/queryOrderReturnList', '', '查询退货/售后主表列表');

        -- 配置退货原因表权限
        INSERT INTO sys_menu (id, menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES (nextval('sys_menu_id_seq'), '退货原因', 2, 2, 7, '/oms/orderReturnReason/list', '', '', '退货原因');

        SELECT currval('sys_menu_id_seq') INTO parentId;

        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('添加退货原因', 3, 1, parentId, '', '/api/oms/returnReason/addOrderReturnReason', '', '添加退货原因');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('删除退货原因', 3, 2, parentId, '', '/api/oms/returnReason/deleteOrderReturnReason', '', '删除退货原因');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新退货原因', 3, 3, parentId, '', '/api/oms/returnReason/updateOrderReturnReason', '', '更新退货原因');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新退货原因表状态', 3, 4, parentId, '', '/api/oms/returnReason/updateOrderReturnReasonStatus', '', '更新退货原因表状态');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询退货原因表详情', 3, 5, parentId, '', '/api/oms/returnReason/queryOrderReturnReasonDetail', '', '查询退货原因表详情');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询退货原因表列表', 3, 6, parentId, '', '/api/oms/returnReason/queryOrderReturnReasonList', '', '查询退货原因表列表');

        -- 配置订单设置表权限
        INSERT INTO sys_menu (id, menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES (nextval('sys_menu_id_seq'), '订单设置', 2, 1, 7, '/oms/orderSetting/list', '', '', '订单设置');

        SELECT currval('sys_menu_id_seq') INTO parentId;

        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('添加订单设置', 3, 1, parentId, '', '/api/oms/setting/addOrderSetting', '', '添加订单设置');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('删除订单设置', 3, 2, parentId, '', '/api/oms/setting/deleteOrderSetting', '', '删除订单设置');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新订单设置', 3, 3, parentId, '', '/api/oms/setting/updateOrderSetting', '', '更新订单设置');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新订单设置表状态', 3, 4, parentId, '', '/api/oms/setting/updateOrderSettingStatus', '', '更新订单设置表状态');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询订单设置表详情', 3, 5, parentId, '', '/api/oms/setting/queryOrderSettingDetail', '', '查询订单设置表详情');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询订单设置表列表', 3, 6, parentId, '', '/api/oms/setting/queryOrderSettingList', '', '查询订单设置表列表');
    END $$;

DO $$
    DECLARE
        parentId bigint;
    BEGIN


        -- 配置优惠券表权限
        INSERT INTO sys_menu (id, menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES (nextval('sys_menu_id_seq'), '优惠券', 2, 2, 8, '/sms/coupon/list', '', '', '优惠券');

        SELECT currval('sys_menu_id_seq') INTO parentId;

        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('添加优惠券', 3, 1, parentId, '', '/api/sms/coupon/addCoupon', '', '添加优惠券');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('删除优惠券', 3, 2, parentId, '', '/api/sms/coupon/deleteCoupon', '', '删除优惠券');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新优惠券', 3, 3, parentId, '', '/api/sms/coupon/updateCoupon', '', '更新优惠券');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新优惠券表状态', 3, 4, parentId, '', '/api/sms/coupon/updateCouponStatus', '', '更新优惠券表状态');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询优惠券表详情', 3, 5, parentId, '', '/api/sms/coupon/queryCouponDetail', '', '查询优惠券表详情');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询优惠券表列表', 3, 6, parentId, '', '/api/sms/coupon/queryCouponList', '', '查询优惠券表列表');

        -- 配置优惠券领取记录表权限
        INSERT INTO sys_menu (id, menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark,visible) VALUES (nextval('sys_menu_id_seq'), '优惠券领取记录', 2, 1, 8, '/sms/couponRecord', '', '', '优惠券领取记录',0);

        SELECT currval('sys_menu_id_seq') INTO parentId;

        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新优惠券领取记录', 3, 3, parentId, '', '/api/sms/couponRecord/updateCouponRecord', '', '更新优惠券领取记录');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询优惠券领取记录表详情', 3, 5, parentId, '', '/api/sms/couponRecord/queryCouponRecordDetail', '', '查询优惠券领取记录表详情');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询优惠券领取记录表列表', 3, 6, parentId, '', '/api/sms/couponRecord/queryCouponRecordList', '', '查询优惠券领取记录表列表');

        -- 配置优惠券使用范围表权限
        INSERT INTO sys_menu (id, menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark, visible) VALUES (nextval('sys_menu_id_seq'), '优惠券使用范围', 2, 6, 8, '/sms/couponScope', '', '', '优惠券使用范围',0);

        SELECT currval('sys_menu_id_seq') INTO parentId;

        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('添加优惠券使用范围', 3, 1, parentId, '', '/api/sms/couponScope/addCouponScope', '', '添加优惠券使用范围');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('删除优惠券使用范围', 3, 2, parentId, '', '/api/sms/couponScope/deleteCouponScope', '', '删除优惠券使用范围');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新优惠券使用范围', 3, 3, parentId, '', '/api/sms/couponScope/updateCouponScope', '', '更新优惠券使用范围');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新优惠券使用范围表状态', 3, 4, parentId, '', '/api/sms/couponScope/updateCouponScopeStatus', '', '更新优惠券使用范围表状态');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询优惠券使用范围表详情', 3, 5, parentId, '', '/api/sms/couponScope/queryCouponScopeDetail', '', '查询优惠券使用范围表详情');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询优惠券使用范围表列表', 3, 6, parentId, '', '/api/sms/couponScope/queryCouponScopeList', '', '查询优惠券使用范围表列表');

        -- 配置优惠券类型表权限
        INSERT INTO sys_menu (id, menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES (nextval('sys_menu_id_seq'), '优惠券类型', 2, 5, 8, '/sms/couponType', '', '', '优惠券类型');

        SELECT currval('sys_menu_id_seq') INTO parentId;

        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('添加优惠券类型', 3, 1, parentId, '', '/api/sms/couponType/addCouponType', '', '添加优惠券类型');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('删除优惠券类型', 3, 2, parentId, '', '/api/sms/couponType/deleteCouponType', '', '删除优惠券类型');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新优惠券类型', 3, 3, parentId, '', '/api/sms/couponType/updateCouponType', '', '更新优惠券类型');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新优惠券类型表状态', 3, 4, parentId, '', '/api/sms/couponType/updateCouponTypeStatus', '', '更新优惠券类型表状态');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询优惠券类型表详情', 3, 5, parentId, '', '/api/sms/couponType/queryCouponTypeDetail', '', '查询优惠券类型表详情');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询优惠券类型表列表', 3, 6, parentId, '', '/api/sms/couponType/queryCouponTypeList', '', '查询优惠券类型表列表');

        -- 配置首页轮播广告表权限
        INSERT INTO sys_menu (id, menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES (nextval('sys_menu_id_seq'), '广告列表', 2, 3, 8, '/sms/homeAdvertise/list', '', '', '首页轮播广告');

        SELECT currval('sys_menu_id_seq') INTO parentId;

        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('添加首页轮播广告', 3, 1, parentId, '', '/api/sms/homeAdvertise/addHomeAdvertise', '', '添加首页轮播广告');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('删除首页轮播广告', 3, 2, parentId, '', '/api/sms/homeAdvertise/deleteHomeAdvertise', '', '删除首页轮播广告');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新首页轮播广告', 3, 3, parentId, '', '/api/sms/homeAdvertise/updateHomeAdvertise', '', '更新首页轮播广告');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新首页轮播广告表状态', 3, 4, parentId, '', '/api/sms/homeAdvertise/updateHomeAdvertiseStatus', '', '更新首页轮播广告表状态');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询首页轮播广告表详情', 3, 5, parentId, '', '/api/sms/homeAdvertise/queryHomeAdvertiseDetail', '', '查询首页轮播广告表详情');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询首页轮播广告表列表', 3, 6, parentId, '', '/api/sms/homeAdvertise/queryHomeAdvertiseList', '', '查询首页轮播广告表列表');

        -- 配置秒杀活动表权限
        INSERT INTO sys_menu (id, menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES (nextval('sys_menu_id_seq'), '秒杀活动', 2, 1, 8, '/sms/seckillActivity/list', '', '', '秒杀活动');

        SELECT currval('sys_menu_id_seq') INTO parentId;

        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('添加秒杀活动', 3, 1, parentId, '', '/api/sms/seckillActivity/addSeckillActivity', '', '添加秒杀活动');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('删除秒杀活动', 3, 2, parentId, '', '/api/sms/seckillActivity/deleteSeckillActivity', '', '删除秒杀活动');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新秒杀活动', 3, 3, parentId, '', '/api/sms/seckillActivity/updateSeckillActivity', '', '更新秒杀活动');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新秒杀活动表状态', 3, 4, parentId, '', '/api/sms/seckillActivity/updateSeckillActivityStatus', '', '更新秒杀活动表状态');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询秒杀活动表详情', 3, 5, parentId, '', '/api/sms/seckillActivity/querySeckillActivityDetail', '', '查询秒杀活动表详情');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询秒杀活动表列表', 3, 6, parentId, '', '/api/sms/seckillActivity/querySeckillActivityList', '', '查询秒杀活动表列表');

        -- 配置秒杀商品表权限
        INSERT INTO sys_menu (id, menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark,visible) VALUES (nextval('sys_menu_id_seq'), '秒杀商品', 2, 8, 8, '/sms/seckillProduct', '', '', '秒杀商品',0);

        SELECT currval('sys_menu_id_seq') INTO parentId;

        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('添加秒杀商品', 3, 1, parentId, '', '/api/sms/seckillProduct/addSeckillProduct', '', '添加秒杀商品');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('删除秒杀商品', 3, 2, parentId, '', '/api/sms/seckillProduct/deleteSeckillProduct', '', '删除秒杀商品');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新秒杀商品', 3, 3, parentId, '', '/api/sms/seckillProduct/updateSeckillProduct', '', '更新秒杀商品');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新秒杀商品表状态', 3, 4, parentId, '', '/api/sms/seckillProduct/updateSeckillProductStatus', '', '更新秒杀商品表状态');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询秒杀商品表详情', 3, 5, parentId, '', '/api/sms/seckillProduct/querySeckillProductDetail', '', '查询秒杀商品表详情');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询秒杀商品表列表', 3, 6, parentId, '', '/api/sms/seckillProduct/querySeckillProductList', '', '查询秒杀商品表列表');

        -- 配置秒杀预约表权限
        INSERT INTO sys_menu (id, menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark,visible) VALUES (nextval('sys_menu_id_seq'), '秒杀预约', 2, 9, 8, '/sms/seckillReservation', '', '', '秒杀预约',0);

        SELECT currval('sys_menu_id_seq') INTO parentId;

        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('添加秒杀预约', 3, 1, parentId, '', '/api/sms/seckillReservation/addSeckillReservation', '', '添加秒杀预约');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('删除秒杀预约', 3, 2, parentId, '', '/api/sms/seckillReservation/deleteSeckillReservation', '', '删除秒杀预约');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新秒杀预约', 3, 3, parentId, '', '/api/sms/seckillReservation/updateSeckillReservation', '', '更新秒杀预约');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新秒杀预约表状态', 3, 4, parentId, '', '/api/sms/seckillReservation/updateSeckillReservationStatus', '', '更新秒杀预约表状态');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询秒杀预约表详情', 3, 5, parentId, '', '/api/sms/seckillReservation/querySeckillReservationDetail', '', '查询秒杀预约表详情');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询秒杀预约表列表', 3, 6, parentId, '', '/api/sms/seckillReservation/querySeckillReservationList', '', '查询秒杀预约表列表');

        -- 配置秒杀场次表权限
        INSERT INTO sys_menu (id, menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark, visible) VALUES (nextval('sys_menu_id_seq'), '秒杀场次', 2, 4, 8, '/sms/seckillSession', '', '', '秒杀场次',0);

        SELECT currval('sys_menu_id_seq') INTO parentId;

        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('添加秒杀场次', 3, 1, parentId, '', '/api/sms/seckillSession/addSeckillSession', '', '添加秒杀场次');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('删除秒杀场次', 3, 2, parentId, '', '/api/sms/seckillSession/deleteSeckillSession', '', '删除秒杀场次');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新秒杀场次', 3, 3, parentId, '', '/api/sms/seckillSession/updateSeckillSession', '', '更新秒杀场次');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新秒杀场次表状态', 3, 4, parentId, '', '/api/sms/seckillSession/updateSeckillSessionStatus', '', '更新秒杀场次表状态');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询秒杀场次表详情', 3, 5, parentId, '', '/api/sms/seckillSession/querySeckillSessionDetail', '', '查询秒杀场次表详情');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询秒杀场次表列表', 3, 6, parentId, '', '/api/sms/seckillSession/querySeckillSessionList', '', '查询秒杀场次表列表');

        -- 配置品牌推荐权限
        INSERT INTO sys_menu (id, menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark, visible) VALUES (nextval('sys_menu_id_seq'), '品牌推荐', 2, 4, 8, '/sms/homeBrand/list', '', '', '品牌推荐',1);

        SELECT currval('sys_menu_id_seq') INTO parentId;

        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('添加品牌推荐', 3, 1, parentId, '', '/api/sms/homeBrand/addHomeBrand', '', '添加品牌推荐');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('删除品牌推荐', 3, 2, parentId, '', '/api/sms/homeBrand/deleteHomeBrand', '', '删除品牌推荐');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新品牌推荐', 3, 3, parentId, '', '/api/sms/homeBrand/updateHomeBrand', '', '更新品牌推荐');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新品牌推荐表状态', 3, 4, parentId, '', '/api/sms/homeBrand/updateHomeBrandStatus', '', '更新品牌推荐表状态');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询品牌推荐表详情', 3, 5, parentId, '', '/api/sms/homeBrand/queryHomeBrandDetail', '', '查询品牌推荐表详情');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询品牌推荐表列表', 3, 6, parentId, '', '/api/sms/homeBrand/queryHomeBrandList', '', '查询品牌推荐表列表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('修改推荐品牌排序', 3, 6, parentId, '', '/api/sms/homeBrand/updateHomeBrandSort', '', '修改推荐品牌排序');

        -- 配置新品推荐表权限
        INSERT INTO sys_menu (id, menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark, visible) VALUES (nextval('sys_menu_id_seq'), '新品推荐', 2, 4, 8, '/sms/homeNewProduct/list', '', '', '新品推荐',1);

        SELECT currval('sys_menu_id_seq') INTO parentId;

        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('添加新品推荐', 3, 1, parentId, '', '/api/sms/homeNewProduct/addHomeNewProduct', '', '添加新品推荐');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('删除新品推荐', 3, 2, parentId, '', '/api/sms/homeNewProduct/deleteHomeNewProduct', '', '删除新品推荐');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新新品推荐', 3, 3, parentId, '', '/api/sms/homeNewProduct/updateHomeNewProduct', '', '更新新品推荐');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新新品推荐表状态', 3, 4, parentId, '', '/api/sms/homeNewProduct/updateHomeNewProductStatus', '', '更新新品推荐表状态');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询新品推荐表详情', 3, 5, parentId, '', '/api/sms/homeNewProduct/queryHomeNewProductDetail', '', '查询新品推荐表详情');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询新品推荐表列表', 3, 6, parentId, '', '/api/sms/homeNewProduct/queryHomeNewProductList', '', '查询新品推荐表列表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新新品推荐排序', 3, 6, parentId, '', '/api/sms/homeNewProduct/updateNewProductSort', '', '更新新品推荐排序');

        -- 配置人气推荐表权限
        INSERT INTO sys_menu (id, menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark, visible) VALUES (nextval('sys_menu_id_seq'), '人气推荐', 2, 4, 8, '/sms/homeRecommendProduct/list', '', '', '人气推荐',1);

        SELECT currval('sys_menu_id_seq') INTO parentId;

        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('添加人气推荐', 3, 1, parentId, '', '/api/sms/homeRecommendProduct/addHomeRecommendProduct', '', '添加人气推荐');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('删除人气推荐', 3, 2, parentId, '', '/api/sms/homeRecommendProduct/deleteHomeRecommendProduct', '', '删除人气推荐');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新人气推荐', 3, 3, parentId, '', '/api/sms/homeRecommendProduct/updateHomeRecommendProduct', '', '更新人气推荐');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新人气推荐表状态', 3, 4, parentId, '', '/api/sms/homeRecommendProduct/updateHomeRecommendProductStatus', '', '更新人气推荐表状态');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询人气推荐表详情', 3, 5, parentId, '', '/api/sms/homeRecommendProduct/queryHomeRecommendProductDetail', '', '查询人气推荐表详情');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询人气推荐表列表', 3, 6, parentId, '', '/api/sms/homeRecommendProduct/queryHomeRecommendProductList', '', '查询人气推荐表列表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新人气推荐排序', 3, 6, parentId, '', '/api/sms/homeRecommendProduct/updateRecommendProductSort', '', '更新人气推荐排序');

        -- 配置专题推荐表权限
        INSERT INTO sys_menu (id, menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark, visible) VALUES (nextval('sys_menu_id_seq'), '专题推荐', 2, 4, 8, '/sms/homeRecommendSubject/list', '', '', '专题推荐',1);

        SELECT currval('sys_menu_id_seq') INTO parentId;

        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('添加专题推荐', 3, 1, parentId, '', '/api/sms/homeRecommendSubject/addHomeRecommendSubject', '', '添加专题推荐');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('删除专题推荐', 3, 2, parentId, '', '/api/sms/homeRecommendSubject/deleteHomeRecommendSubject', '', '删除专题推荐');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新专题推荐', 3, 3, parentId, '', '/api/sms/homeRecommendSubject/updateHomeRecommendSubject', '', '更新专题推荐');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新专题推荐表状态', 3, 4, parentId, '', '/api/sms/homeRecommendSubject/updateHomeRecommendSubjectStatus', '', '更新专题推荐表状态');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询专题推荐表详情', 3, 5, parentId, '', '/api/sms/homeRecommendSubject/queryHomeRecommendSubjectDetail', '', '查询专题推荐表详情');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询专题推荐表列表', 3, 6, parentId, '', '/api/sms/homeRecommendSubject/queryHomeRecommendSubjectList', '', '查询专题推荐表列表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新专题推荐排序', 3, 6, parentId, '', '/api/sms/homeRecommendSubject/updateRecommendSubjectSort', '', '更新专题推荐排序');

    END $$;

DO $$
    DECLARE
        parentId bigint;
    BEGIN


        -- 配置帮助表权限
        INSERT INTO sys_menu (id, menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark,visible) VALUES (nextval('sys_menu_id_seq'), '帮助', 2, 1, 9, '/cms/help', '', '', '帮助',0);

        SELECT currval('sys_menu_id_seq') INTO parentId;

        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('添加帮助', 3, 1, parentId, '', '/api/cms/help/addHelp', '', '添加帮助');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('删除帮助', 3, 2, parentId, '', '/api/cms/help/deleteHelp', '', '删除帮助');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新帮助', 3, 3, parentId, '', '/api/cms/help/updateHelp', '', '更新帮助');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新帮助表状态', 3, 4, parentId, '', '/api/cms/help/updateHelpStatus', '', '更新帮助表状态');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询帮助表详情', 3, 5, parentId, '', '/api/cms/help/queryHelpDetail', '', '查询帮助表详情');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询帮助表列表', 3, 6, parentId, '', '/api/cms/help/queryHelpList', '', '查询帮助表列表');

        -- 配置帮助分类表权限
        INSERT INTO sys_menu (id, menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES (nextval('sys_menu_id_seq'), '帮助分类', 2, 1, 9, '/cms/helpCategory', '', '', '帮助分类');

        SELECT currval('sys_menu_id_seq') INTO parentId;

        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('添加帮助分类', 3, 1, parentId, '', '/api/cms/helpCategory/addHelpCategory', '', '添加帮助分类');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('删除帮助分类', 3, 2, parentId, '', '/api/cms/helpCategory/deleteHelpCategory', '', '删除帮助分类');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新帮助分类', 3, 3, parentId, '', '/api/cms/helpCategory/updateHelpCategory', '', '更新帮助分类');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新帮助分类表状态', 3, 4, parentId, '', '/api/cms/helpCategory/updateHelpCategoryStatus', '', '更新帮助分类表状态');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询帮助分类表详情', 3, 5, parentId, '', '/api/cms/helpCategory/queryHelpCategoryDetail', '', '查询帮助分类表详情');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询帮助分类表列表', 3, 6, parentId, '', '/api/cms/helpCategory/queryHelpCategoryList', '', '查询帮助分类表列表');

        -- 配置用户举报表权限
        INSERT INTO sys_menu (id, menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES (nextval('sys_menu_id_seq'), '用户举报', 2, 1, 9, '/cms/memberReport', '', '', '用户举报');

        SELECT currval('sys_menu_id_seq') INTO parentId;

        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('删除用户举报', 3, 2, parentId, '', '/api/cms/memberReport/deleteMemberReport', '', '删除用户举报');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新用户举报', 3, 3, parentId, '', '/api/cms/memberReport/updateMemberReport', '', '更新用户举报');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新用户举报表状态', 3, 4, parentId, '', '/api/cms/memberReport/updateMemberReportStatus', '', '更新用户举报表状态');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询用户举报表详情', 3, 5, parentId, '', '/api/cms/memberReport/queryMemberReportDetail', '', '查询用户举报表详情');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询用户举报表列表', 3, 6, parentId, '', '/api/cms/memberReport/queryMemberReportList', '', '查询用户举报表列表');

        -- 配置优选专区权限
        INSERT INTO sys_menu (id, menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES (nextval('sys_menu_id_seq'), '优选专区', 2, 1, 9, '/cms/preferredArea', '', '', '优选专区管理');

        SELECT currval('sys_menu_id_seq') INTO parentId;

        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('添加优选专区', 3, 1, parentId, '', '/api/cms/preferredArea/addPreferredArea', '', '添加优选专区');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('删除优选专区', 3, 2, parentId, '', '/api/cms/preferredArea/deletePreferredArea', '', '删除优选专区');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新优选专区', 3, 3, parentId, '', '/api/cms/preferredArea/updatePreferredArea', '', '更新优选专区');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新优选专区状态', 3, 4, parentId, '', '/api/cms/preferredArea/updatePreferredAreaStatus', '', '更新优选专区状态');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询优选专区详情', 3, 5, parentId, '', '/api/cms/preferredArea/queryPreferredAreaDetail', '', '查询优选专区详情');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询优选专区列表', 3, 6, parentId, '', '/api/cms/prefrenceArea/queryPreferredAreaList', '', '查询优选专区列表');

        -- 配置专题表权限
        INSERT INTO sys_menu (id, menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES (nextval('sys_menu_id_seq'), '专题专区', 2, 1, 9, '/cms/subject', '', '', '专题');

        SELECT currval('sys_menu_id_seq') INTO parentId;

        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('添加专题', 3, 1, parentId, '', '/api/cms/subject/addSubject', '', '添加专题');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('删除专题', 3, 2, parentId, '', '/api/cms/subject/deleteSubject', '', '删除专题');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新专题', 3, 3, parentId, '', '/api/cms/subject/updateSubject', '', '更新专题');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新专题表状态', 3, 4, parentId, '', '/api/cms/subject/updateSubjectStatus', '', '更新专题表状态');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询专题表详情', 3, 5, parentId, '', '/api/cms/subject/querySubjectDetail', '', '查询专题表详情');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询专题表列表', 3, 6, parentId, '', '/api/cms/subject/querySubjectList', '', '查询专题表列表');

        -- 配置专题分类表权限
        INSERT INTO sys_menu (id, menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES (nextval('sys_menu_id_seq'), '专题分类', 2, 1, 9, '/cms/subjectCategory', '', '', '专题分类');

        SELECT currval('sys_menu_id_seq') INTO parentId;

        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('添加专题分类', 3, 1, parentId, '', '/api/cms/subjectCategory/addSubjectCategory', '', '添加专题分类');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('删除专题分类', 3, 2, parentId, '', '/api/cms/subjectCategory/deleteSubjectCategory', '', '删除专题分类');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新专题分类', 3, 3, parentId, '', '/api/cms/subjectCategory/updateSubjectCategory', '', '更新专题分类');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新专题分类表状态', 3, 4, parentId, '', '/api/cms/subjectCategory/updateSubjectCategoryStatus', '', '更新专题分类表状态');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询专题分类表详情', 3, 5, parentId, '', '/api/cms/subjectCategory/querySubjectCategoryDetail', '', '查询专题分类表详情');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询专题分类表列表', 3, 6, parentId, '', '/api/cms/subjectCategory/querySubjectCategoryList', '', '查询专题分类表列表');

        -- 配置专题评论表权限
        INSERT INTO sys_menu (id, menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark,visible) VALUES (nextval('sys_menu_id_seq'), '专题评论', 2, 1, 9, '/cms/subjectComment', '', '', '专题评论',0);

        SELECT currval('sys_menu_id_seq') INTO parentId;

        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('删除专题评论', 3, 2, parentId, '', '/api/cms/subjectComment/deleteSubjectComment', '', '删除专题评论');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新专题评论', 3, 3, parentId, '', '/api/cms/subjectComment/updateSubjectComment', '', '更新专题评论');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新专题评论表状态', 3, 4, parentId, '', '/api/cms/subjectComment/updateSubjectCommentStatus', '', '更新专题评论表状态');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询专题评论表详情', 3, 5, parentId, '', '/api/cms/subjectComment/querySubjectCommentDetail', '', '查询专题评论表详情');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询专题评论表列表', 3, 6, parentId, '', '/api/cms/subjectComment/querySubjectCommentList', '', '查询专题评论表列表');

        -- 配置话题表权限
        INSERT INTO sys_menu (id, menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES (nextval('sys_menu_id_seq'), '话题', 2, 1, 9, '/cms/topic', '', '', '话题');

        SELECT currval('sys_menu_id_seq') INTO parentId;

        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('添加话题', 3, 1, parentId, '', '/api/cms/topic/addTopic', '', '添加话题');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('删除话题', 3, 2, parentId, '', '/api/cms/topic/deleteTopic', '', '删除话题');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新话题', 3, 3, parentId, '', '/api/cms/topic/updateTopic', '', '更新话题');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新话题表状态', 3, 4, parentId, '', '/api/cms/topic/updateTopicStatus', '', '更新话题表状态');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询话题表详情', 3, 5, parentId, '', '/api/cms/topic/queryTopicDetail', '', '查询话题表详情');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询话题表列表', 3, 6, parentId, '', '/api/cms/topic/queryTopicList', '', '查询话题表列表');

        -- 配置话题分类表权限
        INSERT INTO sys_menu (id, menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES (nextval('sys_menu_id_seq'), '话题分类', 2, 1, 9, '/cms/topicCategory', '', '', '话题分类');

        SELECT currval('sys_menu_id_seq') INTO parentId;

        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('添加话题分类', 3, 1, parentId, '', '/api/cms/topicCategory/addTopicCategory', '', '添加话题分类');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('删除话题分类', 3, 2, parentId, '', '/api/cms/topicCategory/deleteTopicCategory', '', '删除话题分类');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新话题分类', 3, 3, parentId, '', '/api/cms/topicCategory/updateTopicCategory', '', '更新话题分类');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新话题分类表状态', 3, 4, parentId, '', '/api/cms/topicCategory/updateTopicCategoryStatus', '', '更新话题分类表状态');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询话题分类表详情', 3, 5, parentId, '', '/api/cms/topicCategory/queryTopicCategoryDetail', '', '查询话题分类表详情');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询话题分类表列表', 3, 6, parentId, '', '/api/cms/topicCategory/queryTopicCategoryList', '', '查询话题分类表列表');

        -- 配置话题评论表权限
        INSERT INTO sys_menu (id, menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark,visible) VALUES (nextval('sys_menu_id_seq'), '话题评论', 2, 1, 9, '/cms/topicComment', '', '', '话题评论',0);

        SELECT currval('sys_menu_id_seq') INTO parentId;

        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('删除话题评论', 3, 2, parentId, '', '/api/cms/topicComment/deleteTopicComment', '', '删除话题评论');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新话题评论', 3, 3, parentId, '', '/api/cms/topicComment/updateTopicComment', '', '更新话题评论');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新话题评论表状态', 3, 4, parentId, '', '/api/cms/topicComment/updateTopicCommentStatus', '', '更新话题评论表状态');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询话题评论表详情', 3, 5, parentId, '', '/api/cms/topicComment/queryTopicCommentDetail', '', '查询话题评论表详情');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询话题评论表列表', 3, 6, parentId, '', '/api/cms/topicComment/queryTopicCommentList', '', '查询话题评论表列表');
    END $$;



