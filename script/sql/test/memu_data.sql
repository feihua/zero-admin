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
        insert into sys_menu (id, menu_name, menu_type, status, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) values (nextval('sys_menu_id_seq'), '欢迎', 1, 1, 2, 1, '/welcome', '', 'SmileOutlined', '欢迎');
        insert into sys_menu (id, menu_name, menu_type, status, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) values (nextval('sys_menu_id_seq'), '系统管理', 1, 1, 3, 1, '/system', '', 'SettingOutlined', '系统管理');
        insert into sys_menu (id, menu_name, menu_type, status, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) values (nextval('sys_menu_id_seq'), '日志管理', 1, 1, 4, 1, '/log', '', 'DeleteOutlined', '日志管理');
        insert into sys_menu (id, menu_name, menu_type, status, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) values (nextval('sys_menu_id_seq'), '会员管理', 1, 1, 5, 1, '/ums', '', 'FrownOutlined', '会员管理');
        insert into sys_menu (id, menu_name, menu_type, status, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) values (nextval('sys_menu_id_seq'), '商品管理', 1, 1, 6, 1, '/pms', '', 'GiftOutlined', '商品管理');
        insert into sys_menu (id, menu_name, menu_type, status, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) values (nextval('sys_menu_id_seq'), '订单管理', 1, 1, 7, 1, '/oms', '', 'DollarCircleOutlined', '订单管理');
        insert into sys_menu (id, menu_name, menu_type, status, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) values (nextval('sys_menu_id_seq'), '营销管理', 1, 1, 8, 1, '/sms', '', 'AlertOutlined', '营销管理');
        insert into sys_menu (id, menu_name, menu_type, status, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) values (nextval('sys_menu_id_seq'), '内容管理', 1, 1, 9, 1, '/cms', '', 'SettingOutlined', '内容管理');


        -- 配置菜单信息权限
        INSERT INTO sys_menu (id, menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES (nextval('sys_menu_id_seq'), '菜单信息管理', 2, 1, 3, '/sys/menu', '', '', '菜单信息管理');

        SELECT currval('sys_menu_id_seq') INTO parentId;

        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('添加菜单信息', 3, 1, parentId, '', '/api/sys/menu/addMenu', '', '添加菜单信息');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('删除菜单信息', 3, 2, parentId, '', '/api/sys/menu/deleteMenu', '', '删除菜单信息');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新菜单信息', 3, 3, parentId, '', '/api/sys/menu/updateMenu', '', '更新菜单信息');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新菜单信息状态', 3, 4, parentId, '', '/api/sys/menu/updateMenuStatus', '', '更新菜单信息状态');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询菜单信息详情', 3, 5, parentId, '', '/api/sys/menu/queryMenuDetail', '', '查询菜单信息详情');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询菜单信息列表', 3, 6, parentId, '', '/api/sys/menu/queryMenuList', '', '查询菜单信息列表');

        -- 配置用户信息权限
        INSERT INTO sys_menu (id, menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES (nextval('sys_menu_id_seq'), '用户信息管理', 2, 1, 3, '/sys/user', '', '', '用户信息管理');

        SELECT currval('sys_menu_id_seq') INTO parentId;

        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('添加用户信息', 3, 1, parentId, '', '/api/sys/user/addUser', '', '添加用户信息');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('删除用户信息', 3, 2, parentId, '', '/api/sys/user/deleteUser', '', '删除用户信息');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新用户信息', 3, 3, parentId, '', '/api/sys/user/updateUser', '', '更新用户信息');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新用户信息状态', 3, 4, parentId, '', '/api/sys/user/updateUserStatus', '', '更新用户信息状态');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询用户信息详情', 3, 5, parentId, '', '/api/sys/user/queryUserDetail', '', '查询用户信息详情');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询用户信息列表', 3, 6, parentId, '', '/api/sys/user/queryUserList', '', '查询用户信息列表');

        -- 配置用户角色关联表权限
        INSERT INTO sys_menu (id, menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES (nextval('sys_menu_id_seq'), '用户角色关联表管理', 2, 1, 3, '/sys/userRole', '', '', '用户角色关联表管理');

        SELECT currval('sys_menu_id_seq') INTO parentId;

        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('添加用户角色关联表', 3, 1, parentId, '', '/api/sys/userRole/addUserRole', '', '添加用户角色关联表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('删除用户角色关联表', 3, 2, parentId, '', '/api/sys/userRole/deleteUserRole', '', '删除用户角色关联表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新用户角色关联表', 3, 3, parentId, '', '/api/sys/userRole/updateUserRole', '', '更新用户角色关联表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新用户角色关联表状态', 3, 4, parentId, '', '/api/sys/userRole/updateUserRoleStatus', '', '更新用户角色关联表状态');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询用户角色关联表详情', 3, 5, parentId, '', '/api/sys/userRole/queryUserRoleDetail', '', '查询用户角色关联表详情');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询用户角色关联表列表', 3, 6, parentId, '', '/api/sys/userRole/queryUserRoleList', '', '查询用户角色关联表列表');

        -- 配置角色和部门关联表权限
        INSERT INTO sys_menu (id, menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES (nextval('sys_menu_id_seq'), '角色和部门关联表管理', 2, 1, 3, '/sys/roleDept', '', '', '角色和部门关联表管理');

        SELECT currval('sys_menu_id_seq') INTO parentId;

        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('添加角色和部门关联表', 3, 1, parentId, '', '/api/sys/roleDept/addRoleDept', '', '添加角色和部门关联表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('删除角色和部门关联表', 3, 2, parentId, '', '/api/sys/roleDept/deleteRoleDept', '', '删除角色和部门关联表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新角色和部门关联表', 3, 3, parentId, '', '/api/sys/roleDept/updateRoleDept', '', '更新角色和部门关联表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新角色和部门关联表状态', 3, 4, parentId, '', '/api/sys/roleDept/updateRoleDeptStatus', '', '更新角色和部门关联表状态');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询角色和部门关联表详情', 3, 5, parentId, '', '/api/sys/roleDept/queryRoleDeptDetail', '', '查询角色和部门关联表详情');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询角色和部门关联表列表', 3, 6, parentId, '', '/api/sys/roleDept/queryRoleDeptList', '', '查询角色和部门关联表列表');

        -- 配置角色信息权限
        INSERT INTO sys_menu (id, menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES (nextval('sys_menu_id_seq'), '角色信息管理', 2, 1, 3, '/sys/role', '', '', '角色信息管理');

        SELECT currval('sys_menu_id_seq') INTO parentId;

        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('添加角色信息', 3, 1, parentId, '', '/api/sys/role/addRole', '', '添加角色信息');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('删除角色信息', 3, 2, parentId, '', '/api/sys/role/deleteRole', '', '删除角色信息');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新角色信息', 3, 3, parentId, '', '/api/sys/role/updateRole', '', '更新角色信息');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新角色信息状态', 3, 4, parentId, '', '/api/sys/role/updateRoleStatus', '', '更新角色信息状态');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询角色信息详情', 3, 5, parentId, '', '/api/sys/role/queryRoleDetail', '', '查询角色信息详情');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询角色信息列表', 3, 6, parentId, '', '/api/sys/role/queryRoleList', '', '查询角色信息列表');

        -- 配置角色菜单关联表权限
        INSERT INTO sys_menu (id, menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES (nextval('sys_menu_id_seq'), '角色菜单关联表管理', 2, 1, 3, '/sys/roleMenu', '', '', '角色菜单关联表管理');

        SELECT currval('sys_menu_id_seq') INTO parentId;

        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('添加角色菜单关联表', 3, 1, parentId, '', '/api/sys/roleMenu/addRoleMenu', '', '添加角色菜单关联表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('删除角色菜单关联表', 3, 2, parentId, '', '/api/sys/roleMenu/deleteRoleMenu', '', '删除角色菜单关联表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新角色菜单关联表', 3, 3, parentId, '', '/api/sys/roleMenu/updateRoleMenu', '', '更新角色菜单关联表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新角色菜单关联表状态', 3, 4, parentId, '', '/api/sys/roleMenu/updateRoleMenuStatus', '', '更新角色菜单关联表状态');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询角色菜单关联表详情', 3, 5, parentId, '', '/api/sys/roleMenu/queryRoleMenuDetail', '', '查询角色菜单关联表详情');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询角色菜单关联表列表', 3, 6, parentId, '', '/api/sys/roleMenu/queryRoleMenuList', '', '查询角色菜单关联表列表');

        -- 配置用户与岗位关联表权限
        INSERT INTO sys_menu (id, menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES (nextval('sys_menu_id_seq'), '用户与岗位关联表管理', 2, 1, 3, '/sys/userPost', '', '', '用户与岗位关联表管理');

        SELECT currval('sys_menu_id_seq') INTO parentId;

        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('添加用户与岗位关联表', 3, 1, parentId, '', '/api/sys/userPost/addUserPost', '', '添加用户与岗位关联表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('删除用户与岗位关联表', 3, 2, parentId, '', '/api/sys/userPost/deleteUserPost', '', '删除用户与岗位关联表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新用户与岗位关联表', 3, 3, parentId, '', '/api/sys/userPost/updateUserPost', '', '更新用户与岗位关联表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新用户与岗位关联表状态', 3, 4, parentId, '', '/api/sys/userPost/updateUserPostStatus', '', '更新用户与岗位关联表状态');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询用户与岗位关联表详情', 3, 5, parentId, '', '/api/sys/userPost/queryUserPostDetail', '', '查询用户与岗位关联表详情');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询用户与岗位关联表列表', 3, 6, parentId, '', '/api/sys/userPost/queryUserPostList', '', '查询用户与岗位关联表列表');

        -- 配置岗位信息表权限
        INSERT INTO sys_menu (id, menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES (nextval('sys_menu_id_seq'), '岗位信息表管理', 2, 1, 3, '/sys/post', '', '', '岗位信息表管理');

        SELECT currval('sys_menu_id_seq') INTO parentId;

        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('添加岗位信息表', 3, 1, parentId, '', '/api/sys/post/addPost', '', '添加岗位信息表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('删除岗位信息表', 3, 2, parentId, '', '/api/sys/post/deletePost', '', '删除岗位信息表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新岗位信息表', 3, 3, parentId, '', '/api/sys/post/updatePost', '', '更新岗位信息表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新岗位信息表状态', 3, 4, parentId, '', '/api/sys/post/updatePostStatus', '', '更新岗位信息表状态');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询岗位信息表详情', 3, 5, parentId, '', '/api/sys/post/queryPostDetail', '', '查询岗位信息表详情');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询岗位信息表列表', 3, 6, parentId, '', '/api/sys/post/queryPostList', '', '查询岗位信息表列表');

        -- 配置部门表权限
        INSERT INTO sys_menu (id, menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES (nextval('sys_menu_id_seq'), '部门表管理', 2, 1, 3, '/sys/dept', '', '', '部门表管理');

        SELECT currval('sys_menu_id_seq') INTO parentId;

        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('添加部门表', 3, 1, parentId, '', '/api/sys/dept/addDept', '', '添加部门表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('删除部门表', 3, 2, parentId, '', '/api/sys/dept/deleteDept', '', '删除部门表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新部门表', 3, 3, parentId, '', '/api/sys/dept/updateDept', '', '更新部门表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新部门表状态', 3, 4, parentId, '', '/api/sys/dept/updateDeptStatus', '', '更新部门表状态');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询部门表详情', 3, 5, parentId, '', '/api/sys/dept/queryDeptDetail', '', '查询部门表详情');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询部门表列表', 3, 6, parentId, '', '/api/sys/dept/queryDeptList', '', '查询部门表列表');

        -- 配置字典数据表权限
        INSERT INTO sys_menu (id, menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES (nextval('sys_menu_id_seq'), '字典数据表管理', 2, 1, 3, '/sys/dictData', '', '', '字典数据表管理');

        SELECT currval('sys_menu_id_seq') INTO parentId;

        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('添加字典数据表', 3, 1, parentId, '', '/api/sys/dictData/addDictData', '', '添加字典数据表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('删除字典数据表', 3, 2, parentId, '', '/api/sys/dictData/deleteDictData', '', '删除字典数据表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新字典数据表', 3, 3, parentId, '', '/api/sys/dictData/updateDictData', '', '更新字典数据表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新字典数据表状态', 3, 4, parentId, '', '/api/sys/dictData/updateDictDataStatus', '', '更新字典数据表状态');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询字典数据表详情', 3, 5, parentId, '', '/api/sys/dictData/queryDictDataDetail', '', '查询字典数据表详情');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询字典数据表列表', 3, 6, parentId, '', '/api/sys/dictData/queryDictDataList', '', '查询字典数据表列表');

        -- 配置字典类型表权限
        INSERT INTO sys_menu (id, menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES (nextval('sys_menu_id_seq'), '字典类型表管理', 2, 1, 3, '/sys/dictType', '', '', '字典类型表管理');

        SELECT currval('sys_menu_id_seq') INTO parentId;

        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('添加字典类型表', 3, 1, parentId, '', '/api/sys/dictType/addDictType', '', '添加字典类型表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('删除字典类型表', 3, 2, parentId, '', '/api/sys/dictType/deleteDictType', '', '删除字典类型表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新字典类型表', 3, 3, parentId, '', '/api/sys/dictType/updateDictType', '', '更新字典类型表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新字典类型表状态', 3, 4, parentId, '', '/api/sys/dictType/updateDictTypeStatus', '', '更新字典类型表状态');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询字典类型表详情', 3, 5, parentId, '', '/api/sys/dictType/queryDictTypeDetail', '', '查询字典类型表详情');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询字典类型表列表', 3, 6, parentId, '', '/api/sys/dictType/queryDictTypeList', '', '查询字典类型表列表');

        -- 配置系统访问记录权限
        INSERT INTO sys_menu (id, menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES (nextval('sys_menu_id_seq'), '系统访问记录管理', 2, 1, 4, '/sys/loginLog', '', '', '系统访问记录管理');

        SELECT currval('sys_menu_id_seq') INTO parentId;

        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('添加系统访问记录', 3, 1, parentId, '', '/api/sys/loginLog/addLoginLog', '', '添加系统访问记录');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('删除系统访问记录', 3, 2, parentId, '', '/api/sys/loginLog/deleteLoginLog', '', '删除系统访问记录');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新系统访问记录', 3, 3, parentId, '', '/api/sys/loginLog/updateLoginLog', '', '更新系统访问记录');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新系统访问记录状态', 3, 4, parentId, '', '/api/sys/loginLog/updateLoginLogStatus', '', '更新系统访问记录状态');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询系统访问记录详情', 3, 5, parentId, '', '/api/sys/loginLog/queryLoginLogDetail', '', '查询系统访问记录详情');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询系统访问记录列表', 3, 6, parentId, '', '/api/sys/loginLog/queryLoginLogList', '', '查询系统访问记录列表');

        -- 配置通知公告表权限
        INSERT INTO sys_menu (id, menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES (nextval('sys_menu_id_seq'), '通知公告表管理', 2, 1, 3, '/sys/notice', '', '', '通知公告表管理');

        SELECT currval('sys_menu_id_seq') INTO parentId;

        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('添加通知公告表', 3, 1, parentId, '', '/api/sys/notice/addNotice', '', '添加通知公告表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('删除通知公告表', 3, 2, parentId, '', '/api/sys/notice/deleteNotice', '', '删除通知公告表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新通知公告表', 3, 3, parentId, '', '/api/sys/notice/updateNotice', '', '更新通知公告表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新通知公告表状态', 3, 4, parentId, '', '/api/sys/notice/updateNoticeStatus', '', '更新通知公告表状态');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询通知公告表详情', 3, 5, parentId, '', '/api/sys/notice/queryNoticeDetail', '', '查询通知公告表详情');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询通知公告表列表', 3, 6, parentId, '', '/api/sys/notice/queryNoticeList', '', '查询通知公告表列表');

        -- 配置操作日志记录权限
        INSERT INTO sys_menu (id, menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES (nextval('sys_menu_id_seq'), '操作日志记录管理', 2, 1, 4, '/sys/operateLog', '', '', '操作日志记录管理');

        SELECT currval('sys_menu_id_seq') INTO parentId;

        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('添加操作日志记录', 3, 1, parentId, '', '/api/sys/operateLog/addOperateLog', '', '添加操作日志记录');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('删除操作日志记录', 3, 2, parentId, '', '/api/sys/operateLog/deleteOperateLog', '', '删除操作日志记录');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新操作日志记录', 3, 3, parentId, '', '/api/sys/operateLog/updateOperateLog', '', '更新操作日志记录');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新操作日志记录状态', 3, 4, parentId, '', '/api/sys/operateLog/updateOperateLogStatus', '', '更新操作日志记录状态');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询操作日志记录详情', 3, 5, parentId, '', '/api/sys/operateLog/queryOperateLogDetail', '', '查询操作日志记录详情');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询操作日志记录列表', 3, 6, parentId, '', '/api/sys/operateLog/queryOperateLogList', '', '查询操作日志记录列表');
    END $$;

DO $$
    DECLARE
        parentId bigint;
    BEGIN


        -- 配置会员收货地址表权限
        INSERT INTO sys_menu (id, menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES (nextval('sys_menu_id_seq'), '会员收货地址表管理', 2, 1, 5, '/ums/memberAddress', '', '', '会员收货地址表管理');

        SELECT currval('sys_menu_id_seq') INTO parentId;

        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('添加会员收货地址表', 3, 1, parentId, '', '/api/ums/memberAddress/addMemberAddress', '', '添加会员收货地址表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('删除会员收货地址表', 3, 2, parentId, '', '/api/ums/memberAddress/deleteMemberAddress', '', '删除会员收货地址表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新会员收货地址表', 3, 3, parentId, '', '/api/ums/memberAddress/updateMemberAddress', '', '更新会员收货地址表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新会员收货地址表状态', 3, 4, parentId, '', '/api/ums/memberAddress/updateMemberAddressStatus', '', '更新会员收货地址表状态');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询会员收货地址表详情', 3, 5, parentId, '', '/api/ums/memberAddress/queryMemberAddressDetail', '', '查询会员收货地址表详情');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询会员收货地址表列表', 3, 6, parentId, '', '/api/ums/memberAddress/queryMemberAddressList', '', '查询会员收货地址表列表');

        -- 配置积分消费设置权限
        INSERT INTO sys_menu (id, menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES (nextval('sys_menu_id_seq'), '积分消费设置管理', 2, 1, 5, '/ums/memberConsumeSetting', '', '', '积分消费设置管理');

        SELECT currval('sys_menu_id_seq') INTO parentId;

        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('添加积分消费设置', 3, 1, parentId, '', '/api/ums/memberConsumeSetting/addMemberConsumeSetting', '', '添加积分消费设置');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('删除积分消费设置', 3, 2, parentId, '', '/api/ums/memberConsumeSetting/deleteMemberConsumeSetting', '', '删除积分消费设置');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新积分消费设置', 3, 3, parentId, '', '/api/ums/memberConsumeSetting/updateMemberConsumeSetting', '', '更新积分消费设置');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新积分消费设置状态', 3, 4, parentId, '', '/api/ums/memberConsumeSetting/updateMemberConsumeSettingStatus', '', '更新积分消费设置状态');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询积分消费设置详情', 3, 5, parentId, '', '/api/ums/memberConsumeSetting/queryMemberConsumeSettingDetail', '', '查询积分消费设置详情');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询积分消费设置列表', 3, 6, parentId, '', '/api/ums/memberConsumeSetting/queryMemberConsumeSettingList', '', '查询积分消费设置列表');

        -- 配置会员成长值记录表权限
        INSERT INTO sys_menu (id, menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES (nextval('sys_menu_id_seq'), '会员成长值记录表管理', 2, 1, 5, '/ums/memberGrowthLog', '', '', '会员成长值记录表管理');

        SELECT currval('sys_menu_id_seq') INTO parentId;

        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('添加会员成长值记录表', 3, 1, parentId, '', '/api/ums/memberGrowthLog/addMemberGrowthLog', '', '添加会员成长值记录表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('删除会员成长值记录表', 3, 2, parentId, '', '/api/ums/memberGrowthLog/deleteMemberGrowthLog', '', '删除会员成长值记录表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新会员成长值记录表', 3, 3, parentId, '', '/api/ums/memberGrowthLog/updateMemberGrowthLog', '', '更新会员成长值记录表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新会员成长值记录表状态', 3, 4, parentId, '', '/api/ums/memberGrowthLog/updateMemberGrowthLogStatus', '', '更新会员成长值记录表状态');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询会员成长值记录表详情', 3, 5, parentId, '', '/api/ums/memberGrowthLog/queryMemberGrowthLogDetail', '', '查询会员成长值记录表详情');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询会员成长值记录表列表', 3, 6, parentId, '', '/api/ums/memberGrowthLog/queryMemberGrowthLogList', '', '查询会员成长值记录表列表');

        -- 配置会员信息表权限
        INSERT INTO sys_menu (id, menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES (nextval('sys_menu_id_seq'), '会员信息表管理', 2, 1, 5, '/ums/memberInfo', '', '', '会员信息表管理');

        SELECT currval('sys_menu_id_seq') INTO parentId;

        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('添加会员信息表', 3, 1, parentId, '', '/api/ums/memberInfo/addMemberInfo', '', '添加会员信息表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('删除会员信息表', 3, 2, parentId, '', '/api/ums/memberInfo/deleteMemberInfo', '', '删除会员信息表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新会员信息表', 3, 3, parentId, '', '/api/ums/memberInfo/updateMemberInfo', '', '更新会员信息表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新会员信息表状态', 3, 4, parentId, '', '/api/ums/memberInfo/updateMemberInfoStatus', '', '更新会员信息表状态');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询会员信息表详情', 3, 5, parentId, '', '/api/ums/memberInfo/queryMemberInfoDetail', '', '查询会员信息表详情');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询会员信息表列表', 3, 6, parentId, '', '/api/ums/memberInfo/queryMemberInfoList', '', '查询会员信息表列表');

        -- 配置会员等级表权限
        INSERT INTO sys_menu (id, menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES (nextval('sys_menu_id_seq'), '会员等级表管理', 2, 1, 5, '/ums/memberLevel', '', '', '会员等级表管理');

        SELECT currval('sys_menu_id_seq') INTO parentId;

        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('添加会员等级表', 3, 1, parentId, '', '/api/ums/memberLevel/addMemberLevel', '', '添加会员等级表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('删除会员等级表', 3, 2, parentId, '', '/api/ums/memberLevel/deleteMemberLevel', '', '删除会员等级表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新会员等级表', 3, 3, parentId, '', '/api/ums/memberLevel/updateMemberLevel', '', '更新会员等级表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新会员等级表状态', 3, 4, parentId, '', '/api/ums/memberLevel/updateMemberLevelStatus', '', '更新会员等级表状态');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询会员等级表详情', 3, 5, parentId, '', '/api/ums/memberLevel/queryMemberLevelDetail', '', '查询会员等级表详情');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询会员等级表列表', 3, 6, parentId, '', '/api/ums/memberLevel/queryMemberLevelList', '', '查询会员等级表列表');

        -- 配置会员登录记录权限
        INSERT INTO sys_menu (id, menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES (nextval('sys_menu_id_seq'), '会员登录记录管理', 2, 1, 5, '/ums/memberLoginLog', '', '', '会员登录记录管理');

        SELECT currval('sys_menu_id_seq') INTO parentId;

        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('添加会员登录记录', 3, 1, parentId, '', '/api/ums/memberLoginLog/addMemberLoginLog', '', '添加会员登录记录');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('删除会员登录记录', 3, 2, parentId, '', '/api/ums/memberLoginLog/deleteMemberLoginLog', '', '删除会员登录记录');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新会员登录记录', 3, 3, parentId, '', '/api/ums/memberLoginLog/updateMemberLoginLog', '', '更新会员登录记录');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新会员登录记录状态', 3, 4, parentId, '', '/api/ums/memberLoginLog/updateMemberLoginLogStatus', '', '更新会员登录记录状态');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询会员登录记录详情', 3, 5, parentId, '', '/api/ums/memberLoginLog/queryMemberLoginLogDetail', '', '查询会员登录记录详情');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询会员登录记录列表', 3, 6, parentId, '', '/api/ums/memberLoginLog/queryMemberLoginLogList', '', '查询会员登录记录列表');

        -- 配置会员积分记录表权限
        INSERT INTO sys_menu (id, menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES (nextval('sys_menu_id_seq'), '会员积分记录表管理', 2, 1, 5, '/ums/memberPointsLog', '', '', '会员积分记录表管理');

        SELECT currval('sys_menu_id_seq') INTO parentId;

        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('添加会员积分记录表', 3, 1, parentId, '', '/api/ums/memberPointsLog/addMemberPointsLog', '', '添加会员积分记录表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('删除会员积分记录表', 3, 2, parentId, '', '/api/ums/memberPointsLog/deleteMemberPointsLog', '', '删除会员积分记录表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新会员积分记录表', 3, 3, parentId, '', '/api/ums/memberPointsLog/updateMemberPointsLog', '', '更新会员积分记录表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新会员积分记录表状态', 3, 4, parentId, '', '/api/ums/memberPointsLog/updateMemberPointsLogStatus', '', '更新会员积分记录表状态');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询会员积分记录表详情', 3, 5, parentId, '', '/api/ums/memberPointsLog/queryMemberPointsLogDetail', '', '查询会员积分记录表详情');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询会员积分记录表列表', 3, 6, parentId, '', '/api/ums/memberPointsLog/queryMemberPointsLogList', '', '查询会员积分记录表列表');

        -- 配置会员积分成长规则表权限
        INSERT INTO sys_menu (id, menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES (nextval('sys_menu_id_seq'), '会员积分成长规则表管理', 2, 1, 5, '/ums/memberRuleSetting', '', '', '会员积分成长规则表管理');

        SELECT currval('sys_menu_id_seq') INTO parentId;

        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('添加会员积分成长规则表', 3, 1, parentId, '', '/api/ums/memberRuleSetting/addMemberRuleSetting', '', '添加会员积分成长规则表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('删除会员积分成长规则表', 3, 2, parentId, '', '/api/ums/memberRuleSetting/deleteMemberRuleSetting', '', '删除会员积分成长规则表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新会员积分成长规则表', 3, 3, parentId, '', '/api/ums/memberRuleSetting/updateMemberRuleSetting', '', '更新会员积分成长规则表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新会员积分成长规则表状态', 3, 4, parentId, '', '/api/ums/memberRuleSetting/updateMemberRuleSettingStatus', '', '更新会员积分成长规则表状态');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询会员积分成长规则表详情', 3, 5, parentId, '', '/api/ums/memberRuleSetting/queryMemberRuleSettingDetail', '', '查询会员积分成长规则表详情');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询会员积分成长规则表列表', 3, 6, parentId, '', '/api/ums/memberRuleSetting/queryMemberRuleSettingList', '', '查询会员积分成长规则表列表');

        -- 配置会员签到记录表权限
        INSERT INTO sys_menu (id, menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES (nextval('sys_menu_id_seq'), '会员签到记录表管理', 2, 1, 5, '/ums/memberSignLog', '', '', '会员签到记录表管理');

        SELECT currval('sys_menu_id_seq') INTO parentId;

        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('添加会员签到记录表', 3, 1, parentId, '', '/api/ums/memberSignLog/addMemberSignLog', '', '添加会员签到记录表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('删除会员签到记录表', 3, 2, parentId, '', '/api/ums/memberSignLog/deleteMemberSignLog', '', '删除会员签到记录表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新会员签到记录表', 3, 3, parentId, '', '/api/ums/memberSignLog/updateMemberSignLog', '', '更新会员签到记录表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新会员签到记录表状态', 3, 4, parentId, '', '/api/ums/memberSignLog/updateMemberSignLogStatus', '', '更新会员签到记录表状态');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询会员签到记录表详情', 3, 5, parentId, '', '/api/ums/memberSignLog/queryMemberSignLogDetail', '', '查询会员签到记录表详情');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询会员签到记录表列表', 3, 6, parentId, '', '/api/ums/memberSignLog/queryMemberSignLogList', '', '查询会员签到记录表列表');

        -- 配置会员统计信息权限
        INSERT INTO sys_menu (id, menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES (nextval('sys_menu_id_seq'), '会员统计信息管理', 2, 1, 5, '/ums/memberStatisticsInfo', '', '', '会员统计信息管理');

        SELECT currval('sys_menu_id_seq') INTO parentId;

        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('添加会员统计信息', 3, 1, parentId, '', '/api/ums/memberStatisticsInfo/addMemberStatisticsInfo', '', '添加会员统计信息');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('删除会员统计信息', 3, 2, parentId, '', '/api/ums/memberStatisticsInfo/deleteMemberStatisticsInfo', '', '删除会员统计信息');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新会员统计信息', 3, 3, parentId, '', '/api/ums/memberStatisticsInfo/updateMemberStatisticsInfo', '', '更新会员统计信息');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新会员统计信息状态', 3, 4, parentId, '', '/api/ums/memberStatisticsInfo/updateMemberStatisticsInfoStatus', '', '更新会员统计信息状态');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询会员统计信息详情', 3, 5, parentId, '', '/api/ums/memberStatisticsInfo/queryMemberStatisticsInfoDetail', '', '查询会员统计信息详情');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询会员统计信息列表', 3, 6, parentId, '', '/api/ums/memberStatisticsInfo/queryMemberStatisticsInfoList', '', '查询会员统计信息列表');

        -- 配置用户标签表权限
        INSERT INTO sys_menu (id, menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES (nextval('sys_menu_id_seq'), '用户标签表管理', 2, 1, 5, '/ums/memberTag', '', '', '用户标签表管理');

        SELECT currval('sys_menu_id_seq') INTO parentId;

        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('添加用户标签表', 3, 1, parentId, '', '/api/ums/memberTag/addMemberTag', '', '添加用户标签表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('删除用户标签表', 3, 2, parentId, '', '/api/ums/memberTag/deleteMemberTag', '', '删除用户标签表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新用户标签表', 3, 3, parentId, '', '/api/ums/memberTag/updateMemberTag', '', '更新用户标签表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新用户标签表状态', 3, 4, parentId, '', '/api/ums/memberTag/updateMemberTagStatus', '', '更新用户标签表状态');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询用户标签表详情', 3, 5, parentId, '', '/api/ums/memberTag/queryMemberTagDetail', '', '查询用户标签表详情');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询用户标签表列表', 3, 6, parentId, '', '/api/ums/memberTag/queryMemberTagList', '', '查询用户标签表列表');

        -- 配置会员标签关联表权限
        INSERT INTO sys_menu (id, menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES (nextval('sys_menu_id_seq'), '会员标签关联表管理', 2, 1, 5, '/ums/memberTagRelation', '', '', '会员标签关联表管理');

        SELECT currval('sys_menu_id_seq') INTO parentId;

        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('添加会员标签关联表', 3, 1, parentId, '', '/api/ums/memberTagRelation/addMemberTagRelation', '', '添加会员标签关联表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('删除会员标签关联表', 3, 2, parentId, '', '/api/ums/memberTagRelation/deleteMemberTagRelation', '', '删除会员标签关联表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新会员标签关联表', 3, 3, parentId, '', '/api/ums/memberTagRelation/updateMemberTagRelation', '', '更新会员标签关联表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新会员标签关联表状态', 3, 4, parentId, '', '/api/ums/memberTagRelation/updateMemberTagRelationStatus', '', '更新会员标签关联表状态');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询会员标签关联表详情', 3, 5, parentId, '', '/api/ums/memberTagRelation/queryMemberTagRelationDetail', '', '查询会员标签关联表详情');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询会员标签关联表列表', 3, 6, parentId, '', '/api/ums/memberTagRelation/queryMemberTagRelationList', '', '查询会员标签关联表列表');

        -- 配置会员任务表权限
        INSERT INTO sys_menu (id, menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES (nextval('sys_menu_id_seq'), '会员任务表管理', 2, 1, 5, '/ums/memberTask', '', '', '会员任务表管理');

        SELECT currval('sys_menu_id_seq') INTO parentId;

        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('添加会员任务表', 3, 1, parentId, '', '/api/ums/memberTask/addMemberTask', '', '添加会员任务表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('删除会员任务表', 3, 2, parentId, '', '/api/ums/memberTask/deleteMemberTask', '', '删除会员任务表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新会员任务表', 3, 3, parentId, '', '/api/ums/memberTask/updateMemberTask', '', '更新会员任务表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新会员任务表状态', 3, 4, parentId, '', '/api/ums/memberTask/updateMemberTaskStatus', '', '更新会员任务表状态');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询会员任务表详情', 3, 5, parentId, '', '/api/ums/memberTask/queryMemberTaskDetail', '', '查询会员任务表详情');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询会员任务表列表', 3, 6, parentId, '', '/api/ums/memberTask/queryMemberTaskList', '', '查询会员任务表列表');

        -- 配置会员任务关联表权限
        INSERT INTO sys_menu (id, menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES (nextval('sys_menu_id_seq'), '会员任务关联表管理', 2, 1, 5, '/ums/memberTaskRelation', '', '', '会员任务关联表管理');

        SELECT currval('sys_menu_id_seq') INTO parentId;

        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('添加会员任务关联表', 3, 1, parentId, '', '/api/ums/memberTaskRelation/addMemberTaskRelation', '', '添加会员任务关联表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('删除会员任务关联表', 3, 2, parentId, '', '/api/ums/memberTaskRelation/deleteMemberTaskRelation', '', '删除会员任务关联表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新会员任务关联表', 3, 3, parentId, '', '/api/ums/memberTaskRelation/updateMemberTaskRelation', '', '更新会员任务关联表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新会员任务关联表状态', 3, 4, parentId, '', '/api/ums/memberTaskRelation/updateMemberTaskRelationStatus', '', '更新会员任务关联表状态');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询会员任务关联表详情', 3, 5, parentId, '', '/api/ums/memberTaskRelation/queryMemberTaskRelationDetail', '', '查询会员任务关联表详情');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询会员任务关联表列表', 3, 6, parentId, '', '/api/ums/memberTaskRelation/queryMemberTaskRelationList', '', '查询会员任务关联表列表');
    END $$;



DO $$
    DECLARE
        parentId bigint;
    BEGIN


        -- 配置商品会员价格表权限
        INSERT INTO sys_menu (id, menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES (nextval('sys_menu_id_seq'), '商品会员价格表管理', 2, 1, 6, '/pms/memberPrice', '', '', '商品会员价格表管理');

        SELECT currval('sys_menu_id_seq') INTO parentId;

        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('添加商品会员价格表', 3, 1, parentId, '', '/api/pms/memberPrice/addMemberPrice', '', '添加商品会员价格表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('删除商品会员价格表', 3, 2, parentId, '', '/api/pms/memberPrice/deleteMemberPrice', '', '删除商品会员价格表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新商品会员价格表', 3, 3, parentId, '', '/api/pms/memberPrice/updateMemberPrice', '', '更新商品会员价格表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新商品会员价格表状态', 3, 4, parentId, '', '/api/pms/memberPrice/updateMemberPriceStatus', '', '更新商品会员价格表状态');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询商品会员价格表详情', 3, 5, parentId, '', '/api/pms/memberPrice/queryMemberPriceDetail', '', '查询商品会员价格表详情');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询商品会员价格表列表', 3, 6, parentId, '', '/api/pms/memberPrice/queryMemberPriceList', '', '查询商品会员价格表列表');

        -- 配置商品属性表权限
        INSERT INTO sys_menu (id, menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES (nextval('sys_menu_id_seq'), '商品属性表管理', 2, 1, 6, '/pms/productAttribute', '', '', '商品属性表管理');

        SELECT currval('sys_menu_id_seq') INTO parentId;

        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('添加商品属性表', 3, 1, parentId, '', '/api/pms/productAttribute/addProductAttribute', '', '添加商品属性表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('删除商品属性表', 3, 2, parentId, '', '/api/pms/productAttribute/deleteProductAttribute', '', '删除商品属性表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新商品属性表', 3, 3, parentId, '', '/api/pms/productAttribute/updateProductAttribute', '', '更新商品属性表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新商品属性表状态', 3, 4, parentId, '', '/api/pms/productAttribute/updateProductAttributeStatus', '', '更新商品属性表状态');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询商品属性表详情', 3, 5, parentId, '', '/api/pms/productAttribute/queryProductAttributeDetail', '', '查询商品属性表详情');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询商品属性表列表', 3, 6, parentId, '', '/api/pms/productAttribute/queryProductAttributeList', '', '查询商品属性表列表');

        -- 配置商品属性分组表权限
        INSERT INTO sys_menu (id, menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES (nextval('sys_menu_id_seq'), '商品属性分组表管理', 2, 1, 6, '/pms/productAttributeGroup', '', '', '商品属性分组表管理');

        SELECT currval('sys_menu_id_seq') INTO parentId;

        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('添加商品属性分组表', 3, 1, parentId, '', '/api/pms/productAttributeGroup/addProductAttributeGroup', '', '添加商品属性分组表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('删除商品属性分组表', 3, 2, parentId, '', '/api/pms/productAttributeGroup/deleteProductAttributeGroup', '', '删除商品属性分组表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新商品属性分组表', 3, 3, parentId, '', '/api/pms/productAttributeGroup/updateProductAttributeGroup', '', '更新商品属性分组表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新商品属性分组表状态', 3, 4, parentId, '', '/api/pms/productAttributeGroup/updateProductAttributeGroupStatus', '', '更新商品属性分组表状态');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询商品属性分组表详情', 3, 5, parentId, '', '/api/pms/productAttributeGroup/queryProductAttributeGroupDetail', '', '查询商品属性分组表详情');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询商品属性分组表列表', 3, 6, parentId, '', '/api/pms/productAttributeGroup/queryProductAttributeGroupList', '', '查询商品属性分组表列表');

        -- 配置商品属性值表权限
        INSERT INTO sys_menu (id, menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES (nextval('sys_menu_id_seq'), '商品属性值表管理', 2, 1, 6, '/pms/productAttributeValue', '', '', '商品属性值表管理');

        SELECT currval('sys_menu_id_seq') INTO parentId;

        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('添加商品属性值表', 3, 1, parentId, '', '/api/pms/productAttributeValue/addProductAttributeValue', '', '添加商品属性值表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('删除商品属性值表', 3, 2, parentId, '', '/api/pms/productAttributeValue/deleteProductAttributeValue', '', '删除商品属性值表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新商品属性值表', 3, 3, parentId, '', '/api/pms/productAttributeValue/updateProductAttributeValue', '', '更新商品属性值表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新商品属性值表状态', 3, 4, parentId, '', '/api/pms/productAttributeValue/updateProductAttributeValueStatus', '', '更新商品属性值表状态');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询商品属性值表详情', 3, 5, parentId, '', '/api/pms/productAttributeValue/queryProductAttributeValueDetail', '', '查询商品属性值表详情');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询商品属性值表列表', 3, 6, parentId, '', '/api/pms/productAttributeValue/queryProductAttributeValueList', '', '查询商品属性值表列表');

        -- 配置商品品牌权限
        INSERT INTO sys_menu (id, menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES (nextval('sys_menu_id_seq'), '商品品牌管理', 2, 1, 6, '/pms/productBrand', '', '', '商品品牌管理');

        SELECT currval('sys_menu_id_seq') INTO parentId;

        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('添加商品品牌', 3, 1, parentId, '', '/api/pms/productBrand/addProductBrand', '', '添加商品品牌');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('删除商品品牌', 3, 2, parentId, '', '/api/pms/productBrand/deleteProductBrand', '', '删除商品品牌');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新商品品牌', 3, 3, parentId, '', '/api/pms/productBrand/updateProductBrand', '', '更新商品品牌');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新商品品牌状态', 3, 4, parentId, '', '/api/pms/productBrand/updateProductBrandStatus', '', '更新商品品牌状态');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询商品品牌详情', 3, 5, parentId, '', '/api/pms/productBrand/queryProductBrandDetail', '', '查询商品品牌详情');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询商品品牌列表', 3, 6, parentId, '', '/api/pms/productBrand/queryProductBrandList', '', '查询商品品牌列表');

        -- 配置产品分类权限
        INSERT INTO sys_menu (id, menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES (nextval('sys_menu_id_seq'), '产品分类管理', 2, 1, 6, '/pms/productCategory', '', '', '产品分类管理');

        SELECT currval('sys_menu_id_seq') INTO parentId;

        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('添加产品分类', 3, 1, parentId, '', '/api/pms/productCategory/addProductCategory', '', '添加产品分类');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('删除产品分类', 3, 2, parentId, '', '/api/pms/productCategory/deleteProductCategory', '', '删除产品分类');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新产品分类', 3, 3, parentId, '', '/api/pms/productCategory/updateProductCategory', '', '更新产品分类');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新产品分类状态', 3, 4, parentId, '', '/api/pms/productCategory/updateProductCategoryStatus', '', '更新产品分类状态');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询产品分类详情', 3, 5, parentId, '', '/api/pms/productCategory/queryProductCategoryDetail', '', '查询产品分类详情');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询产品分类列表', 3, 6, parentId, '', '/api/pms/productCategory/queryProductCategoryList', '', '查询产品分类列表');

        -- 配置产品的分类和属性的关系表，用于设置分类筛选条件（只支持一级分类）权限
        INSERT INTO sys_menu (id, menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES (nextval('sys_menu_id_seq'), '产品的分类和属性的关系表，用于设置分类筛选条件（只支持一级分类）管理', 2, 1, 6, '/pms/productCategoryAttributeRelation', '', '', '产品的分类和属性的关系表，用于设置分类筛选条件（只支持一级分类）管理');

        SELECT currval('sys_menu_id_seq') INTO parentId;

        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('添加产品的分类和属性的关系表，用于设置分类筛选条件（只支持一级分类）', 3, 1, parentId, '', '/api/pms/productCategoryAttributeRelation/addProductCategoryAttributeRelation', '', '添加产品的分类和属性的关系表，用于设置分类筛选条件（只支持一级分类）');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('删除产品的分类和属性的关系表，用于设置分类筛选条件（只支持一级分类）', 3, 2, parentId, '', '/api/pms/productCategoryAttributeRelation/deleteProductCategoryAttributeRelation', '', '删除产品的分类和属性的关系表，用于设置分类筛选条件（只支持一级分类）');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新产品的分类和属性的关系表，用于设置分类筛选条件（只支持一级分类）', 3, 3, parentId, '', '/api/pms/productCategoryAttributeRelation/updateProductCategoryAttributeRelation', '', '更新产品的分类和属性的关系表，用于设置分类筛选条件（只支持一级分类）');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新产品的分类和属性的关系表，用于设置分类筛选条件（只支持一级分类）状态', 3, 4, parentId, '', '/api/pms/productCategoryAttributeRelation/updateProductCategoryAttributeRelationStatus', '', '更新产品的分类和属性的关系表，用于设置分类筛选条件（只支持一级分类）状态');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询产品的分类和属性的关系表，用于设置分类筛选条件（只支持一级分类）详情', 3, 5, parentId, '', '/api/pms/productCategoryAttributeRelation/queryProductCategoryAttributeRelationDetail', '', '查询产品的分类和属性的关系表，用于设置分类筛选条件（只支持一级分类）详情');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询产品的分类和属性的关系表，用于设置分类筛选条件（只支持一级分类）列表', 3, 6, parentId, '', '/api/pms/productCategoryAttributeRelation/queryProductCategoryAttributeRelationList', '', '查询产品的分类和属性的关系表，用于设置分类筛选条件（只支持一级分类）列表');

        -- 配置产品满减表(只针对同商品)权限
        INSERT INTO sys_menu (id, menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES (nextval('sys_menu_id_seq'), '产品满减表(只针对同商品)管理', 2, 1, 6, '/pms/productFullReduction', '', '', '产品满减表(只针对同商品)管理');

        SELECT currval('sys_menu_id_seq') INTO parentId;

        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('添加产品满减表(只针对同商品)', 3, 1, parentId, '', '/api/pms/productFullReduction/addProductFullReduction', '', '添加产品满减表(只针对同商品)');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('删除产品满减表(只针对同商品)', 3, 2, parentId, '', '/api/pms/productFullReduction/deleteProductFullReduction', '', '删除产品满减表(只针对同商品)');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新产品满减表(只针对同商品)', 3, 3, parentId, '', '/api/pms/productFullReduction/updateProductFullReduction', '', '更新产品满减表(只针对同商品)');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新产品满减表(只针对同商品)状态', 3, 4, parentId, '', '/api/pms/productFullReduction/updateProductFullReductionStatus', '', '更新产品满减表(只针对同商品)状态');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询产品满减表(只针对同商品)详情', 3, 5, parentId, '', '/api/pms/productFullReduction/queryProductFullReductionDetail', '', '查询产品满减表(只针对同商品)详情');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询产品满减表(只针对同商品)列表', 3, 6, parentId, '', '/api/pms/productFullReduction/queryProductFullReductionList', '', '查询产品满减表(只针对同商品)列表');

        -- 配置运费模版权限
        INSERT INTO sys_menu (id, menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES (nextval('sys_menu_id_seq'), '运费模版管理', 2, 1, 6, '/pms/feightTemplate', '', '', '运费模版管理');

        SELECT currval('sys_menu_id_seq') INTO parentId;

        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('添加运费模版', 3, 1, parentId, '', '/api/pms/feightTemplate/addFeightTemplate', '', '添加运费模版');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('删除运费模版', 3, 2, parentId, '', '/api/pms/feightTemplate/deleteFeightTemplate', '', '删除运费模版');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新运费模版', 3, 3, parentId, '', '/api/pms/feightTemplate/updateFeightTemplate', '', '更新运费模版');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新运费模版状态', 3, 4, parentId, '', '/api/pms/feightTemplate/updateFeightTemplateStatus', '', '更新运费模版状态');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询运费模版详情', 3, 5, parentId, '', '/api/pms/feightTemplate/queryFeightTemplateDetail', '', '查询运费模版详情');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询运费模版列表', 3, 6, parentId, '', '/api/pms/feightTemplate/queryFeightTemplateList', '', '查询运费模版列表');

        -- 配置产品阶梯价格表(只针对同商品)权限
        INSERT INTO sys_menu (id, menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES (nextval('sys_menu_id_seq'), '产品阶梯价格表(只针对同商品)管理', 2, 1, 6, '/pms/productLadder', '', '', '产品阶梯价格表(只针对同商品)管理');

        SELECT currval('sys_menu_id_seq') INTO parentId;

        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('添加产品阶梯价格表(只针对同商品)', 3, 1, parentId, '', '/api/pms/productLadder/addProductLadder', '', '添加产品阶梯价格表(只针对同商品)');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('删除产品阶梯价格表(只针对同商品)', 3, 2, parentId, '', '/api/pms/productLadder/deleteProductLadder', '', '删除产品阶梯价格表(只针对同商品)');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新产品阶梯价格表(只针对同商品)', 3, 3, parentId, '', '/api/pms/productLadder/updateProductLadder', '', '更新产品阶梯价格表(只针对同商品)');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新产品阶梯价格表(只针对同商品)状态', 3, 4, parentId, '', '/api/pms/productLadder/updateProductLadderStatus', '', '更新产品阶梯价格表(只针对同商品)状态');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询产品阶梯价格表(只针对同商品)详情', 3, 5, parentId, '', '/api/pms/productLadder/queryProductLadderDetail', '', '查询产品阶梯价格表(只针对同商品)详情');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询产品阶梯价格表(只针对同商品)列表', 3, 6, parentId, '', '/api/pms/productLadder/queryProductLadderList', '', '查询产品阶梯价格表(只针对同商品)列表');

        -- 配置商品sku表权限
        INSERT INTO sys_menu (id, menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES (nextval('sys_menu_id_seq'), '商品sku表管理', 2, 1, 6, '/pms/productSku', '', '', '商品sku表管理');

        SELECT currval('sys_menu_id_seq') INTO parentId;

        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('添加商品sku表', 3, 1, parentId, '', '/api/pms/productSku/addProductSku', '', '添加商品sku表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('删除商品sku表', 3, 2, parentId, '', '/api/pms/productSku/deleteProductSku', '', '删除商品sku表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新商品sku表', 3, 3, parentId, '', '/api/pms/productSku/updateProductSku', '', '更新商品sku表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新商品sku表状态', 3, 4, parentId, '', '/api/pms/productSku/updateProductSkuStatus', '', '更新商品sku表状态');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询商品sku表详情', 3, 5, parentId, '', '/api/pms/productSku/queryProductSkuDetail', '', '查询商品sku表详情');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询商品sku表列表', 3, 6, parentId, '', '/api/pms/productSku/queryProductSkuList', '', '查询商品sku表列表');

        -- 配置商品规格表权限
        INSERT INTO sys_menu (id, menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES (nextval('sys_menu_id_seq'), '商品规格表管理', 2, 1, 6, '/pms/productSpec', '', '', '商品规格表管理');

        SELECT currval('sys_menu_id_seq') INTO parentId;

        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('添加商品规格表', 3, 1, parentId, '', '/api/pms/productSpec/addProductSpec', '', '添加商品规格表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('删除商品规格表', 3, 2, parentId, '', '/api/pms/productSpec/deleteProductSpec', '', '删除商品规格表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新商品规格表', 3, 3, parentId, '', '/api/pms/productSpec/updateProductSpec', '', '更新商品规格表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新商品规格表状态', 3, 4, parentId, '', '/api/pms/productSpec/updateProductSpecStatus', '', '更新商品规格表状态');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询商品规格表详情', 3, 5, parentId, '', '/api/pms/productSpec/queryProductSpecDetail', '', '查询商品规格表详情');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询商品规格表列表', 3, 6, parentId, '', '/api/pms/productSpec/queryProductSpecList', '', '查询商品规格表列表');

        -- 配置商品规格值表权限
        INSERT INTO sys_menu (id, menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES (nextval('sys_menu_id_seq'), '商品规格值表管理', 2, 1, 6, '/pms/productSpecValue', '', '', '商品规格值表管理');

        SELECT currval('sys_menu_id_seq') INTO parentId;

        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('添加商品规格值表', 3, 1, parentId, '', '/api/pms/productSpecValue/addProductSpecValue', '', '添加商品规格值表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('删除商品规格值表', 3, 2, parentId, '', '/api/pms/productSpecValue/deleteProductSpecValue', '', '删除商品规格值表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新商品规格值表', 3, 3, parentId, '', '/api/pms/productSpecValue/updateProductSpecValue', '', '更新商品规格值表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新商品规格值表状态', 3, 4, parentId, '', '/api/pms/productSpecValue/updateProductSpecValueStatus', '', '更新商品规格值表状态');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询商品规格值表详情', 3, 5, parentId, '', '/api/pms/productSpecValue/queryProductSpecValueDetail', '', '查询商品规格值表详情');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询商品规格值表列表', 3, 6, parentId, '', '/api/pms/productSpecValue/queryProductSpecValueList', '', '查询商品规格值表列表');

        -- 配置商品spu表权限
        INSERT INTO sys_menu (id, menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES (nextval('sys_menu_id_seq'), '商品spu表管理', 2, 1, 6, '/pms/productSpu', '', '', '商品spu表管理');

        SELECT currval('sys_menu_id_seq') INTO parentId;

        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('添加商品spu表', 3, 1, parentId, '', '/api/pms/productSpu/addProductSpu', '', '添加商品spu表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('删除商品spu表', 3, 2, parentId, '', '/api/pms/productSpu/deleteProductSpu', '', '删除商品spu表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新商品spu表', 3, 3, parentId, '', '/api/pms/productSpu/updateProductSpu', '', '更新商品spu表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新商品spu表状态', 3, 4, parentId, '', '/api/pms/productSpu/updateProductSpuStatus', '', '更新商品spu表状态');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询商品spu表详情', 3, 5, parentId, '', '/api/pms/productSpu/queryProductSpuDetail', '', '查询商品spu表详情');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询商品spu表列表', 3, 6, parentId, '', '/api/pms/productSpu/queryProductSpuList', '', '查询商品spu表列表');
    END $$;


DO $$
    DECLARE
        parentId bigint;
    BEGIN


        -- 配置购物车表权限
        INSERT INTO sys_menu (id, menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES (nextval('sys_menu_id_seq'), '购物车表管理', 2, 1, 7, '/oms/cartItem', '', '', '购物车表管理');

        SELECT currval('sys_menu_id_seq') INTO parentId;

        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('添加购物车表', 3, 1, parentId, '', '/api/oms/cartItem/addCartItem', '', '添加购物车表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('删除购物车表', 3, 2, parentId, '', '/api/oms/cartItem/deleteCartItem', '', '删除购物车表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新购物车表', 3, 3, parentId, '', '/api/oms/cartItem/updateCartItem', '', '更新购物车表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新购物车表状态', 3, 4, parentId, '', '/api/oms/cartItem/updateCartItemStatus', '', '更新购物车表状态');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询购物车表详情', 3, 5, parentId, '', '/api/oms/cartItem/queryCartItemDetail', '', '查询购物车表详情');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询购物车表列表', 3, 6, parentId, '', '/api/oms/cartItem/queryCartItemList', '', '查询购物车表列表');

        -- 配置公司收发货地址表权限
        INSERT INTO sys_menu (id, menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES (nextval('sys_menu_id_seq'), '公司收发货地址表管理', 2, 1, 7, '/oms/companyAddress', '', '', '公司收发货地址表管理');

        SELECT currval('sys_menu_id_seq') INTO parentId;

        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('添加公司收发货地址表', 3, 1, parentId, '', '/api/oms/companyAddress/addCompanyAddress', '', '添加公司收发货地址表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('删除公司收发货地址表', 3, 2, parentId, '', '/api/oms/companyAddress/deleteCompanyAddress', '', '删除公司收发货地址表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新公司收发货地址表', 3, 3, parentId, '', '/api/oms/companyAddress/updateCompanyAddress', '', '更新公司收发货地址表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新公司收发货地址表状态', 3, 4, parentId, '', '/api/oms/companyAddress/updateCompanyAddressStatus', '', '更新公司收发货地址表状态');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询公司收发货地址表详情', 3, 5, parentId, '', '/api/oms/companyAddress/queryCompanyAddressDetail', '', '查询公司收发货地址表详情');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询公司收发货地址表列表', 3, 6, parentId, '', '/api/oms/companyAddress/queryCompanyAddressList', '', '查询公司收发货地址表列表');

        -- 配置订单表权限
        INSERT INTO sys_menu (id, menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES (nextval('sys_menu_id_seq'), '订单表管理', 2, 1, 7, '/oms/order', '', '', '订单表管理');

        SELECT currval('sys_menu_id_seq') INTO parentId;

        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('添加订单表', 3, 1, parentId, '', '/api/oms/order/addOrder', '', '添加订单表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('删除订单表', 3, 2, parentId, '', '/api/oms/order/deleteOrder', '', '删除订单表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新订单表', 3, 3, parentId, '', '/api/oms/order/updateOrder', '', '更新订单表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新订单表状态', 3, 4, parentId, '', '/api/oms/order/updateOrderStatus', '', '更新订单表状态');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询订单表详情', 3, 5, parentId, '', '/api/oms/order/queryOrderDetail', '', '查询订单表详情');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询订单表列表', 3, 6, parentId, '', '/api/oms/order/queryOrderList', '', '查询订单表列表');

        -- 配置订单收货地址表权限
        INSERT INTO sys_menu (id, menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES (nextval('sys_menu_id_seq'), '订单收货地址表管理', 2, 1, 7, '/oms/orderDelivery', '', '', '订单收货地址表管理');

        SELECT currval('sys_menu_id_seq') INTO parentId;

        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('添加订单收货地址表', 3, 1, parentId, '', '/api/oms/orderDelivery/addOrderDelivery', '', '添加订单收货地址表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('删除订单收货地址表', 3, 2, parentId, '', '/api/oms/orderDelivery/deleteOrderDelivery', '', '删除订单收货地址表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新订单收货地址表', 3, 3, parentId, '', '/api/oms/orderDelivery/updateOrderDelivery', '', '更新订单收货地址表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新订单收货地址表状态', 3, 4, parentId, '', '/api/oms/orderDelivery/updateOrderDeliveryStatus', '', '更新订单收货地址表状态');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询订单收货地址表详情', 3, 5, parentId, '', '/api/oms/orderDelivery/queryOrderDeliveryDetail', '', '查询订单收货地址表详情');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询订单收货地址表列表', 3, 6, parentId, '', '/api/oms/orderDelivery/queryOrderDeliveryList', '', '查询订单收货地址表列表');

        -- 配置订单商品表权限
        INSERT INTO sys_menu (id, menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES (nextval('sys_menu_id_seq'), '订单商品表管理', 2, 1, 7, '/oms/orderItem', '', '', '订单商品表管理');

        SELECT currval('sys_menu_id_seq') INTO parentId;

        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('添加订单商品表', 3, 1, parentId, '', '/api/oms/orderItem/addOrderItem', '', '添加订单商品表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('删除订单商品表', 3, 2, parentId, '', '/api/oms/orderItem/deleteOrderItem', '', '删除订单商品表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新订单商品表', 3, 3, parentId, '', '/api/oms/orderItem/updateOrderItem', '', '更新订单商品表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新订单商品表状态', 3, 4, parentId, '', '/api/oms/orderItem/updateOrderItemStatus', '', '更新订单商品表状态');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询订单商品表详情', 3, 5, parentId, '', '/api/oms/orderItem/queryOrderItemDetail', '', '查询订单商品表详情');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询订单商品表列表', 3, 6, parentId, '', '/api/oms/orderItem/queryOrderItemList', '', '查询订单商品表列表');

        -- 配置订单主表权限
        INSERT INTO sys_menu (id, menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES (nextval('sys_menu_id_seq'), '订单主表管理', 2, 1, 7, '/oms/orderMain', '', '', '订单主表管理');

        SELECT currval('sys_menu_id_seq') INTO parentId;

        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('添加订单主表', 3, 1, parentId, '', '/api/oms/orderMain/addOrderMain', '', '添加订单主表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('删除订单主表', 3, 2, parentId, '', '/api/oms/orderMain/deleteOrderMain', '', '删除订单主表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新订单主表', 3, 3, parentId, '', '/api/oms/orderMain/updateOrderMain', '', '更新订单主表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新订单主表状态', 3, 4, parentId, '', '/api/oms/orderMain/updateOrderMainStatus', '', '更新订单主表状态');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询订单主表详情', 3, 5, parentId, '', '/api/oms/orderMain/queryOrderMainDetail', '', '查询订单主表详情');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询订单主表列表', 3, 6, parentId, '', '/api/oms/orderMain/queryOrderMainList', '', '查询订单主表列表');

        -- 配置订单操作记录表权限
        INSERT INTO sys_menu (id, menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES (nextval('sys_menu_id_seq'), '订单操作记录表管理', 2, 1, 7, '/oms/orderOperationLog', '', '', '订单操作记录表管理');

        SELECT currval('sys_menu_id_seq') INTO parentId;

        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('添加订单操作记录表', 3, 1, parentId, '', '/api/oms/orderOperationLog/addOrderOperationLog', '', '添加订单操作记录表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('删除订单操作记录表', 3, 2, parentId, '', '/api/oms/orderOperationLog/deleteOrderOperationLog', '', '删除订单操作记录表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新订单操作记录表', 3, 3, parentId, '', '/api/oms/orderOperationLog/updateOrderOperationLog', '', '更新订单操作记录表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新订单操作记录表状态', 3, 4, parentId, '', '/api/oms/orderOperationLog/updateOrderOperationLogStatus', '', '更新订单操作记录表状态');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询订单操作记录表详情', 3, 5, parentId, '', '/api/oms/orderOperationLog/queryOrderOperationLogDetail', '', '查询订单操作记录表详情');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询订单操作记录表列表', 3, 6, parentId, '', '/api/oms/orderOperationLog/queryOrderOperationLogList', '', '查询订单操作记录表列表');

        -- 配置订单支付表权限
        INSERT INTO sys_menu (id, menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES (nextval('sys_menu_id_seq'), '订单支付表管理', 2, 1, 7, '/oms/orderPayment', '', '', '订单支付表管理');

        SELECT currval('sys_menu_id_seq') INTO parentId;

        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('添加订单支付表', 3, 1, parentId, '', '/api/oms/orderPayment/addOrderPayment', '', '添加订单支付表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('删除订单支付表', 3, 2, parentId, '', '/api/oms/orderPayment/deleteOrderPayment', '', '删除订单支付表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新订单支付表', 3, 3, parentId, '', '/api/oms/orderPayment/updateOrderPayment', '', '更新订单支付表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新订单支付表状态', 3, 4, parentId, '', '/api/oms/orderPayment/updateOrderPaymentStatus', '', '更新订单支付表状态');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询订单支付表详情', 3, 5, parentId, '', '/api/oms/orderPayment/queryOrderPaymentDetail', '', '查询订单支付表详情');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询订单支付表列表', 3, 6, parentId, '', '/api/oms/orderPayment/queryOrderPaymentList', '', '查询订单支付表列表');

        -- 配置订单优惠信息表权限
        INSERT INTO sys_menu (id, menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES (nextval('sys_menu_id_seq'), '订单优惠信息表管理', 2, 1, 7, '/oms/orderPromotion', '', '', '订单优惠信息表管理');

        SELECT currval('sys_menu_id_seq') INTO parentId;

        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('添加订单优惠信息表', 3, 1, parentId, '', '/api/oms/orderPromotion/addOrderPromotion', '', '添加订单优惠信息表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('删除订单优惠信息表', 3, 2, parentId, '', '/api/oms/orderPromotion/deleteOrderPromotion', '', '删除订单优惠信息表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新订单优惠信息表', 3, 3, parentId, '', '/api/oms/orderPromotion/updateOrderPromotion', '', '更新订单优惠信息表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新订单优惠信息表状态', 3, 4, parentId, '', '/api/oms/orderPromotion/updateOrderPromotionStatus', '', '更新订单优惠信息表状态');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询订单优惠信息表详情', 3, 5, parentId, '', '/api/oms/orderPromotion/queryOrderPromotionDetail', '', '查询订单优惠信息表详情');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询订单优惠信息表列表', 3, 6, parentId, '', '/api/oms/orderPromotion/queryOrderPromotionList', '', '查询订单优惠信息表列表');

        -- 配置退货/售后主表权限
        INSERT INTO sys_menu (id, menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES (nextval('sys_menu_id_seq'), '退货/售后主表管理', 2, 1, 7, '/oms/orderReturn', '', '', '退货/售后主表管理');

        SELECT currval('sys_menu_id_seq') INTO parentId;

        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('添加退货/售后主表', 3, 1, parentId, '', '/api/oms/orderReturn/addOrderReturn', '', '添加退货/售后主表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('删除退货/售后主表', 3, 2, parentId, '', '/api/oms/orderReturn/deleteOrderReturn', '', '删除退货/售后主表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新退货/售后主表', 3, 3, parentId, '', '/api/oms/orderReturn/updateOrderReturn', '', '更新退货/售后主表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新退货/售后主表状态', 3, 4, parentId, '', '/api/oms/orderReturn/updateOrderReturnStatus', '', '更新退货/售后主表状态');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询退货/售后主表详情', 3, 5, parentId, '', '/api/oms/orderReturn/queryOrderReturnDetail', '', '查询退货/售后主表详情');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询退货/售后主表列表', 3, 6, parentId, '', '/api/oms/orderReturn/queryOrderReturnList', '', '查询退货/售后主表列表');

        -- 配置退货/售后明细表权限
        INSERT INTO sys_menu (id, menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES (nextval('sys_menu_id_seq'), '退货/售后明细表管理', 2, 1, 7, '/oms/orderReturnItem', '', '', '退货/售后明细表管理');

        SELECT currval('sys_menu_id_seq') INTO parentId;

        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('添加退货/售后明细表', 3, 1, parentId, '', '/api/oms/orderReturnItem/addOrderReturnItem', '', '添加退货/售后明细表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('删除退货/售后明细表', 3, 2, parentId, '', '/api/oms/orderReturnItem/deleteOrderReturnItem', '', '删除退货/售后明细表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新退货/售后明细表', 3, 3, parentId, '', '/api/oms/orderReturnItem/updateOrderReturnItem', '', '更新退货/售后明细表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新退货/售后明细表状态', 3, 4, parentId, '', '/api/oms/orderReturnItem/updateOrderReturnItemStatus', '', '更新退货/售后明细表状态');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询退货/售后明细表详情', 3, 5, parentId, '', '/api/oms/orderReturnItem/queryOrderReturnItemDetail', '', '查询退货/售后明细表详情');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询退货/售后明细表列表', 3, 6, parentId, '', '/api/oms/orderReturnItem/queryOrderReturnItemList', '', '查询退货/售后明细表列表');

        -- 配置退货原因表权限
        INSERT INTO sys_menu (id, menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES (nextval('sys_menu_id_seq'), '退货原因表管理', 2, 1, 7, '/oms/orderReturnReason', '', '', '退货原因表管理');

        SELECT currval('sys_menu_id_seq') INTO parentId;

        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('添加退货原因表', 3, 1, parentId, '', '/api/oms/orderReturnReason/addOrderReturnReason', '', '添加退货原因表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('删除退货原因表', 3, 2, parentId, '', '/api/oms/orderReturnReason/deleteOrderReturnReason', '', '删除退货原因表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新退货原因表', 3, 3, parentId, '', '/api/oms/orderReturnReason/updateOrderReturnReason', '', '更新退货原因表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新退货原因表状态', 3, 4, parentId, '', '/api/oms/orderReturnReason/updateOrderReturnReasonStatus', '', '更新退货原因表状态');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询退货原因表详情', 3, 5, parentId, '', '/api/oms/orderReturnReason/queryOrderReturnReasonDetail', '', '查询退货原因表详情');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询退货原因表列表', 3, 6, parentId, '', '/api/oms/orderReturnReason/queryOrderReturnReasonList', '', '查询退货原因表列表');

        -- 配置订单设置表权限
        INSERT INTO sys_menu (id, menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES (nextval('sys_menu_id_seq'), '订单设置表管理', 2, 1, 7, '/oms/orderSetting', '', '', '订单设置表管理');

        SELECT currval('sys_menu_id_seq') INTO parentId;

        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('添加订单设置表', 3, 1, parentId, '', '/api/oms/orderSetting/addOrderSetting', '', '添加订单设置表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('删除订单设置表', 3, 2, parentId, '', '/api/oms/orderSetting/deleteOrderSetting', '', '删除订单设置表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新订单设置表', 3, 3, parentId, '', '/api/oms/orderSetting/updateOrderSetting', '', '更新订单设置表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新订单设置表状态', 3, 4, parentId, '', '/api/oms/orderSetting/updateOrderSettingStatus', '', '更新订单设置表状态');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询订单设置表详情', 3, 5, parentId, '', '/api/oms/orderSetting/queryOrderSettingDetail', '', '查询订单设置表详情');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询订单设置表列表', 3, 6, parentId, '', '/api/oms/orderSetting/queryOrderSettingList', '', '查询订单设置表列表');
    END $$;

DO $$
    DECLARE
        parentId bigint;
    BEGIN


        -- 配置优惠券表权限
        INSERT INTO sys_menu (id, menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES (nextval('sys_menu_id_seq'), '优惠券表管理', 2, 1, 8, '/sms/coupon', '', '', '优惠券表管理');

        SELECT currval('sys_menu_id_seq') INTO parentId;

        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('添加优惠券表', 3, 1, parentId, '', '/api/sms/coupon/addCoupon', '', '添加优惠券表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('删除优惠券表', 3, 2, parentId, '', '/api/sms/coupon/deleteCoupon', '', '删除优惠券表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新优惠券表', 3, 3, parentId, '', '/api/sms/coupon/updateCoupon', '', '更新优惠券表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新优惠券表状态', 3, 4, parentId, '', '/api/sms/coupon/updateCouponStatus', '', '更新优惠券表状态');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询优惠券表详情', 3, 5, parentId, '', '/api/sms/coupon/queryCouponDetail', '', '查询优惠券表详情');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询优惠券表列表', 3, 6, parentId, '', '/api/sms/coupon/queryCouponList', '', '查询优惠券表列表');

        -- 配置优惠券领取记录表权限
        INSERT INTO sys_menu (id, menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES (nextval('sys_menu_id_seq'), '优惠券领取记录表管理', 2, 1, 8, '/sms/couponRecord', '', '', '优惠券领取记录表管理');

        SELECT currval('sys_menu_id_seq') INTO parentId;

        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('添加优惠券领取记录表', 3, 1, parentId, '', '/api/sms/couponRecord/addCouponRecord', '', '添加优惠券领取记录表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('删除优惠券领取记录表', 3, 2, parentId, '', '/api/sms/couponRecord/deleteCouponRecord', '', '删除优惠券领取记录表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新优惠券领取记录表', 3, 3, parentId, '', '/api/sms/couponRecord/updateCouponRecord', '', '更新优惠券领取记录表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新优惠券领取记录表状态', 3, 4, parentId, '', '/api/sms/couponRecord/updateCouponRecordStatus', '', '更新优惠券领取记录表状态');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询优惠券领取记录表详情', 3, 5, parentId, '', '/api/sms/couponRecord/queryCouponRecordDetail', '', '查询优惠券领取记录表详情');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询优惠券领取记录表列表', 3, 6, parentId, '', '/api/sms/couponRecord/queryCouponRecordList', '', '查询优惠券领取记录表列表');

        -- 配置优惠券使用范围表权限
        INSERT INTO sys_menu (id, menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES (nextval('sys_menu_id_seq'), '优惠券使用范围表管理', 2, 1, 8, '/sms/couponScope', '', '', '优惠券使用范围表管理');

        SELECT currval('sys_menu_id_seq') INTO parentId;

        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('添加优惠券使用范围表', 3, 1, parentId, '', '/api/sms/couponScope/addCouponScope', '', '添加优惠券使用范围表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('删除优惠券使用范围表', 3, 2, parentId, '', '/api/sms/couponScope/deleteCouponScope', '', '删除优惠券使用范围表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新优惠券使用范围表', 3, 3, parentId, '', '/api/sms/couponScope/updateCouponScope', '', '更新优惠券使用范围表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新优惠券使用范围表状态', 3, 4, parentId, '', '/api/sms/couponScope/updateCouponScopeStatus', '', '更新优惠券使用范围表状态');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询优惠券使用范围表详情', 3, 5, parentId, '', '/api/sms/couponScope/queryCouponScopeDetail', '', '查询优惠券使用范围表详情');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询优惠券使用范围表列表', 3, 6, parentId, '', '/api/sms/couponScope/queryCouponScopeList', '', '查询优惠券使用范围表列表');

        -- 配置优惠券类型表权限
        INSERT INTO sys_menu (id, menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES (nextval('sys_menu_id_seq'), '优惠券类型表管理', 2, 1, 8, '/sms/couponType', '', '', '优惠券类型表管理');

        SELECT currval('sys_menu_id_seq') INTO parentId;

        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('添加优惠券类型表', 3, 1, parentId, '', '/api/sms/couponType/addCouponType', '', '添加优惠券类型表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('删除优惠券类型表', 3, 2, parentId, '', '/api/sms/couponType/deleteCouponType', '', '删除优惠券类型表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新优惠券类型表', 3, 3, parentId, '', '/api/sms/couponType/updateCouponType', '', '更新优惠券类型表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新优惠券类型表状态', 3, 4, parentId, '', '/api/sms/couponType/updateCouponTypeStatus', '', '更新优惠券类型表状态');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询优惠券类型表详情', 3, 5, parentId, '', '/api/sms/couponType/queryCouponTypeDetail', '', '查询优惠券类型表详情');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询优惠券类型表列表', 3, 6, parentId, '', '/api/sms/couponType/queryCouponTypeList', '', '查询优惠券类型表列表');

        -- 配置首页轮播广告表权限
        INSERT INTO sys_menu (id, menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES (nextval('sys_menu_id_seq'), '首页轮播广告表管理', 2, 1, 8, '/sms/homeAdvertise', '', '', '首页轮播广告表管理');

        SELECT currval('sys_menu_id_seq') INTO parentId;

        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('添加首页轮播广告表', 3, 1, parentId, '', '/api/sms/homeAdvertise/addHomeAdvertise', '', '添加首页轮播广告表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('删除首页轮播广告表', 3, 2, parentId, '', '/api/sms/homeAdvertise/deleteHomeAdvertise', '', '删除首页轮播广告表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新首页轮播广告表', 3, 3, parentId, '', '/api/sms/homeAdvertise/updateHomeAdvertise', '', '更新首页轮播广告表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新首页轮播广告表状态', 3, 4, parentId, '', '/api/sms/homeAdvertise/updateHomeAdvertiseStatus', '', '更新首页轮播广告表状态');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询首页轮播广告表详情', 3, 5, parentId, '', '/api/sms/homeAdvertise/queryHomeAdvertiseDetail', '', '查询首页轮播广告表详情');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询首页轮播广告表列表', 3, 6, parentId, '', '/api/sms/homeAdvertise/queryHomeAdvertiseList', '', '查询首页轮播广告表列表');

        -- 配置秒杀活动表权限
        INSERT INTO sys_menu (id, menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES (nextval('sys_menu_id_seq'), '秒杀活动表管理', 2, 1, 8, '/sms/seckillActivity', '', '', '秒杀活动表管理');

        SELECT currval('sys_menu_id_seq') INTO parentId;

        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('添加秒杀活动表', 3, 1, parentId, '', '/api/sms/seckillActivity/addSeckillActivity', '', '添加秒杀活动表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('删除秒杀活动表', 3, 2, parentId, '', '/api/sms/seckillActivity/deleteSeckillActivity', '', '删除秒杀活动表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新秒杀活动表', 3, 3, parentId, '', '/api/sms/seckillActivity/updateSeckillActivity', '', '更新秒杀活动表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新秒杀活动表状态', 3, 4, parentId, '', '/api/sms/seckillActivity/updateSeckillActivityStatus', '', '更新秒杀活动表状态');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询秒杀活动表详情', 3, 5, parentId, '', '/api/sms/seckillActivity/querySeckillActivityDetail', '', '查询秒杀活动表详情');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询秒杀活动表列表', 3, 6, parentId, '', '/api/sms/seckillActivity/querySeckillActivityList', '', '查询秒杀活动表列表');

        -- 配置秒杀商品表权限
        INSERT INTO sys_menu (id, menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES (nextval('sys_menu_id_seq'), '秒杀商品表管理', 2, 1, 8, '/sms/seckillProduct', '', '', '秒杀商品表管理');

        SELECT currval('sys_menu_id_seq') INTO parentId;

        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('添加秒杀商品表', 3, 1, parentId, '', '/api/sms/seckillProduct/addSeckillProduct', '', '添加秒杀商品表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('删除秒杀商品表', 3, 2, parentId, '', '/api/sms/seckillProduct/deleteSeckillProduct', '', '删除秒杀商品表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新秒杀商品表', 3, 3, parentId, '', '/api/sms/seckillProduct/updateSeckillProduct', '', '更新秒杀商品表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新秒杀商品表状态', 3, 4, parentId, '', '/api/sms/seckillProduct/updateSeckillProductStatus', '', '更新秒杀商品表状态');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询秒杀商品表详情', 3, 5, parentId, '', '/api/sms/seckillProduct/querySeckillProductDetail', '', '查询秒杀商品表详情');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询秒杀商品表列表', 3, 6, parentId, '', '/api/sms/seckillProduct/querySeckillProductList', '', '查询秒杀商品表列表');

        -- 配置秒杀预约表权限
        INSERT INTO sys_menu (id, menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES (nextval('sys_menu_id_seq'), '秒杀预约表管理', 2, 1, 8, '/sms/seckillReservation', '', '', '秒杀预约表管理');

        SELECT currval('sys_menu_id_seq') INTO parentId;

        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('添加秒杀预约表', 3, 1, parentId, '', '/api/sms/seckillReservation/addSeckillReservation', '', '添加秒杀预约表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('删除秒杀预约表', 3, 2, parentId, '', '/api/sms/seckillReservation/deleteSeckillReservation', '', '删除秒杀预约表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新秒杀预约表', 3, 3, parentId, '', '/api/sms/seckillReservation/updateSeckillReservation', '', '更新秒杀预约表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新秒杀预约表状态', 3, 4, parentId, '', '/api/sms/seckillReservation/updateSeckillReservationStatus', '', '更新秒杀预约表状态');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询秒杀预约表详情', 3, 5, parentId, '', '/api/sms/seckillReservation/querySeckillReservationDetail', '', '查询秒杀预约表详情');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询秒杀预约表列表', 3, 6, parentId, '', '/api/sms/seckillReservation/querySeckillReservationList', '', '查询秒杀预约表列表');

        -- 配置秒杀场次表权限
        INSERT INTO sys_menu (id, menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES (nextval('sys_menu_id_seq'), '秒杀场次表管理', 2, 1, 8, '/sms/seckillSession', '', '', '秒杀场次表管理');

        SELECT currval('sys_menu_id_seq') INTO parentId;

        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('添加秒杀场次表', 3, 1, parentId, '', '/api/sms/seckillSession/addSeckillSession', '', '添加秒杀场次表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('删除秒杀场次表', 3, 2, parentId, '', '/api/sms/seckillSession/deleteSeckillSession', '', '删除秒杀场次表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新秒杀场次表', 3, 3, parentId, '', '/api/sms/seckillSession/updateSeckillSession', '', '更新秒杀场次表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新秒杀场次表状态', 3, 4, parentId, '', '/api/sms/seckillSession/updateSeckillSessionStatus', '', '更新秒杀场次表状态');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询秒杀场次表详情', 3, 5, parentId, '', '/api/sms/seckillSession/querySeckillSessionDetail', '', '查询秒杀场次表详情');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询秒杀场次表列表', 3, 6, parentId, '', '/api/sms/seckillSession/querySeckillSessionList', '', '查询秒杀场次表列表');
    END $$;

DO $$
    DECLARE
        parentId bigint;
    BEGIN


        -- 配置帮助表权限
        INSERT INTO sys_menu (id, menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES (nextval('sys_menu_id_seq'), '帮助表管理', 2, 1, 9, '/cms/help', '', '', '帮助表管理');

        SELECT currval('sys_menu_id_seq') INTO parentId;

        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('添加帮助表', 3, 1, parentId, '', '/api/cms/help/addHelp', '', '添加帮助表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('删除帮助表', 3, 2, parentId, '', '/api/cms/help/deleteHelp', '', '删除帮助表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新帮助表', 3, 3, parentId, '', '/api/cms/help/updateHelp', '', '更新帮助表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新帮助表状态', 3, 4, parentId, '', '/api/cms/help/updateHelpStatus', '', '更新帮助表状态');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询帮助表详情', 3, 5, parentId, '', '/api/cms/help/queryHelpDetail', '', '查询帮助表详情');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询帮助表列表', 3, 6, parentId, '', '/api/cms/help/queryHelpList', '', '查询帮助表列表');

        -- 配置帮助分类表权限
        INSERT INTO sys_menu (id, menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES (nextval('sys_menu_id_seq'), '帮助分类表管理', 2, 1, 9, '/cms/helpCategory', '', '', '帮助分类表管理');

        SELECT currval('sys_menu_id_seq') INTO parentId;

        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('添加帮助分类表', 3, 1, parentId, '', '/api/cms/helpCategory/addHelpCategory', '', '添加帮助分类表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('删除帮助分类表', 3, 2, parentId, '', '/api/cms/helpCategory/deleteHelpCategory', '', '删除帮助分类表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新帮助分类表', 3, 3, parentId, '', '/api/cms/helpCategory/updateHelpCategory', '', '更新帮助分类表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新帮助分类表状态', 3, 4, parentId, '', '/api/cms/helpCategory/updateHelpCategoryStatus', '', '更新帮助分类表状态');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询帮助分类表详情', 3, 5, parentId, '', '/api/cms/helpCategory/queryHelpCategoryDetail', '', '查询帮助分类表详情');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询帮助分类表列表', 3, 6, parentId, '', '/api/cms/helpCategory/queryHelpCategoryList', '', '查询帮助分类表列表');

        -- 配置用户举报表权限
        INSERT INTO sys_menu (id, menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES (nextval('sys_menu_id_seq'), '用户举报表管理', 2, 1, 9, '/cms/memberReport', '', '', '用户举报表管理');

        SELECT currval('sys_menu_id_seq') INTO parentId;

        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('添加用户举报表', 3, 1, parentId, '', '/api/cms/memberReport/addMemberReport', '', '添加用户举报表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('删除用户举报表', 3, 2, parentId, '', '/api/cms/memberReport/deleteMemberReport', '', '删除用户举报表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新用户举报表', 3, 3, parentId, '', '/api/cms/memberReport/updateMemberReport', '', '更新用户举报表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新用户举报表状态', 3, 4, parentId, '', '/api/cms/memberReport/updateMemberReportStatus', '', '更新用户举报表状态');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询用户举报表详情', 3, 5, parentId, '', '/api/cms/memberReport/queryMemberReportDetail', '', '查询用户举报表详情');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询用户举报表列表', 3, 6, parentId, '', '/api/cms/memberReport/queryMemberReportList', '', '查询用户举报表列表');

        -- 配置优选专区权限
        INSERT INTO sys_menu (id, menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES (nextval('sys_menu_id_seq'), '优选专区管理', 2, 1, 9, '/cms/preferredArea', '', '', '优选专区管理');

        SELECT currval('sys_menu_id_seq') INTO parentId;

        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('添加优选专区', 3, 1, parentId, '', '/api/cms/preferredArea/addPreferredArea', '', '添加优选专区');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('删除优选专区', 3, 2, parentId, '', '/api/cms/preferredArea/deletePreferredArea', '', '删除优选专区');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新优选专区', 3, 3, parentId, '', '/api/cms/preferredArea/updatePreferredArea', '', '更新优选专区');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新优选专区状态', 3, 4, parentId, '', '/api/cms/preferredArea/updatePreferredAreaStatus', '', '更新优选专区状态');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询优选专区详情', 3, 5, parentId, '', '/api/cms/preferredArea/queryPreferredAreaDetail', '', '查询优选专区详情');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询优选专区列表', 3, 6, parentId, '', '/api/cms/preferredArea/queryPreferredAreaList', '', '查询优选专区列表');

        -- 配置优选专区和产品关系表权限
        INSERT INTO sys_menu (id, menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES (nextval('sys_menu_id_seq'), '优选专区和产品关系表管理', 2, 1, 9, '/cms/preferredAreaProductRelation', '', '', '优选专区和产品关系表管理');

        SELECT currval('sys_menu_id_seq') INTO parentId;

        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('添加优选专区和产品关系表', 3, 1, parentId, '', '/api/cms/preferredAreaProductRelation/addPreferredAreaProductRelation', '', '添加优选专区和产品关系表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('删除优选专区和产品关系表', 3, 2, parentId, '', '/api/cms/preferredAreaProductRelation/deletePreferredAreaProductRelation', '', '删除优选专区和产品关系表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新优选专区和产品关系表', 3, 3, parentId, '', '/api/cms/preferredAreaProductRelation/updatePreferredAreaProductRelation', '', '更新优选专区和产品关系表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新优选专区和产品关系表状态', 3, 4, parentId, '', '/api/cms/preferredAreaProductRelation/updatePreferredAreaProductRelationStatus', '', '更新优选专区和产品关系表状态');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询优选专区和产品关系表详情', 3, 5, parentId, '', '/api/cms/preferredAreaProductRelation/queryPreferredAreaProductRelationDetail', '', '查询优选专区和产品关系表详情');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询优选专区和产品关系表列表', 3, 6, parentId, '', '/api/cms/preferredAreaProductRelation/queryPreferredAreaProductRelationList', '', '查询优选专区和产品关系表列表');

        -- 配置专题表权限
        INSERT INTO sys_menu (id, menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES (nextval('sys_menu_id_seq'), '专题表管理', 2, 1, 9, '/cms/subject', '', '', '专题表管理');

        SELECT currval('sys_menu_id_seq') INTO parentId;

        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('添加专题表', 3, 1, parentId, '', '/api/cms/subject/addSubject', '', '添加专题表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('删除专题表', 3, 2, parentId, '', '/api/cms/subject/deleteSubject', '', '删除专题表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新专题表', 3, 3, parentId, '', '/api/cms/subject/updateSubject', '', '更新专题表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新专题表状态', 3, 4, parentId, '', '/api/cms/subject/updateSubjectStatus', '', '更新专题表状态');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询专题表详情', 3, 5, parentId, '', '/api/cms/subject/querySubjectDetail', '', '查询专题表详情');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询专题表列表', 3, 6, parentId, '', '/api/cms/subject/querySubjectList', '', '查询专题表列表');

        -- 配置专题分类表权限
        INSERT INTO sys_menu (id, menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES (nextval('sys_menu_id_seq'), '专题分类表管理', 2, 1, 9, '/cms/subjectCategory', '', '', '专题分类表管理');

        SELECT currval('sys_menu_id_seq') INTO parentId;

        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('添加专题分类表', 3, 1, parentId, '', '/api/cms/subjectCategory/addSubjectCategory', '', '添加专题分类表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('删除专题分类表', 3, 2, parentId, '', '/api/cms/subjectCategory/deleteSubjectCategory', '', '删除专题分类表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新专题分类表', 3, 3, parentId, '', '/api/cms/subjectCategory/updateSubjectCategory', '', '更新专题分类表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新专题分类表状态', 3, 4, parentId, '', '/api/cms/subjectCategory/updateSubjectCategoryStatus', '', '更新专题分类表状态');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询专题分类表详情', 3, 5, parentId, '', '/api/cms/subjectCategory/querySubjectCategoryDetail', '', '查询专题分类表详情');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询专题分类表列表', 3, 6, parentId, '', '/api/cms/subjectCategory/querySubjectCategoryList', '', '查询专题分类表列表');

        -- 配置专题评论表权限
        INSERT INTO sys_menu (id, menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES (nextval('sys_menu_id_seq'), '专题评论表管理', 2, 1, 9, '/cms/subjectComment', '', '', '专题评论表管理');

        SELECT currval('sys_menu_id_seq') INTO parentId;

        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('添加专题评论表', 3, 1, parentId, '', '/api/cms/subjectComment/addSubjectComment', '', '添加专题评论表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('删除专题评论表', 3, 2, parentId, '', '/api/cms/subjectComment/deleteSubjectComment', '', '删除专题评论表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新专题评论表', 3, 3, parentId, '', '/api/cms/subjectComment/updateSubjectComment', '', '更新专题评论表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新专题评论表状态', 3, 4, parentId, '', '/api/cms/subjectComment/updateSubjectCommentStatus', '', '更新专题评论表状态');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询专题评论表详情', 3, 5, parentId, '', '/api/cms/subjectComment/querySubjectCommentDetail', '', '查询专题评论表详情');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询专题评论表列表', 3, 6, parentId, '', '/api/cms/subjectComment/querySubjectCommentList', '', '查询专题评论表列表');

        -- 配置专题商品关系表权限
        INSERT INTO sys_menu (id, menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES (nextval('sys_menu_id_seq'), '专题商品关系表管理', 2, 1, 9, '/cms/subjectProductRelation', '', '', '专题商品关系表管理');

        SELECT currval('sys_menu_id_seq') INTO parentId;

        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('添加专题商品关系表', 3, 1, parentId, '', '/api/cms/subjectProductRelation/addSubjectProductRelation', '', '添加专题商品关系表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('删除专题商品关系表', 3, 2, parentId, '', '/api/cms/subjectProductRelation/deleteSubjectProductRelation', '', '删除专题商品关系表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新专题商品关系表', 3, 3, parentId, '', '/api/cms/subjectProductRelation/updateSubjectProductRelation', '', '更新专题商品关系表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新专题商品关系表状态', 3, 4, parentId, '', '/api/cms/subjectProductRelation/updateSubjectProductRelationStatus', '', '更新专题商品关系表状态');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询专题商品关系表详情', 3, 5, parentId, '', '/api/cms/subjectProductRelation/querySubjectProductRelationDetail', '', '查询专题商品关系表详情');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询专题商品关系表列表', 3, 6, parentId, '', '/api/cms/subjectProductRelation/querySubjectProductRelationList', '', '查询专题商品关系表列表');

        -- 配置话题表权限
        INSERT INTO sys_menu (id, menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES (nextval('sys_menu_id_seq'), '话题表管理', 2, 1, 9, '/cms/topic', '', '', '话题表管理');

        SELECT currval('sys_menu_id_seq') INTO parentId;

        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('添加话题表', 3, 1, parentId, '', '/api/cms/topic/addTopic', '', '添加话题表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('删除话题表', 3, 2, parentId, '', '/api/cms/topic/deleteTopic', '', '删除话题表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新话题表', 3, 3, parentId, '', '/api/cms/topic/updateTopic', '', '更新话题表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新话题表状态', 3, 4, parentId, '', '/api/cms/topic/updateTopicStatus', '', '更新话题表状态');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询话题表详情', 3, 5, parentId, '', '/api/cms/topic/queryTopicDetail', '', '查询话题表详情');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询话题表列表', 3, 6, parentId, '', '/api/cms/topic/queryTopicList', '', '查询话题表列表');

        -- 配置话题分类表权限
        INSERT INTO sys_menu (id, menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES (nextval('sys_menu_id_seq'), '话题分类表管理', 2, 1, 9, '/cms/topicCategory', '', '', '话题分类表管理');

        SELECT currval('sys_menu_id_seq') INTO parentId;

        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('添加话题分类表', 3, 1, parentId, '', '/api/cms/topicCategory/addTopicCategory', '', '添加话题分类表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('删除话题分类表', 3, 2, parentId, '', '/api/cms/topicCategory/deleteTopicCategory', '', '删除话题分类表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新话题分类表', 3, 3, parentId, '', '/api/cms/topicCategory/updateTopicCategory', '', '更新话题分类表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新话题分类表状态', 3, 4, parentId, '', '/api/cms/topicCategory/updateTopicCategoryStatus', '', '更新话题分类表状态');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询话题分类表详情', 3, 5, parentId, '', '/api/cms/topicCategory/queryTopicCategoryDetail', '', '查询话题分类表详情');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询话题分类表列表', 3, 6, parentId, '', '/api/cms/topicCategory/queryTopicCategoryList', '', '查询话题分类表列表');

        -- 配置话题评论表权限
        INSERT INTO sys_menu (id, menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES (nextval('sys_menu_id_seq'), '话题评论表管理', 2, 1, 9, '/cms/topicComment', '', '', '话题评论表管理');

        SELECT currval('sys_menu_id_seq') INTO parentId;

        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('添加话题评论表', 3, 1, parentId, '', '/api/cms/topicComment/addTopicComment', '', '添加话题评论表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('删除话题评论表', 3, 2, parentId, '', '/api/cms/topicComment/deleteTopicComment', '', '删除话题评论表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新话题评论表', 3, 3, parentId, '', '/api/cms/topicComment/updateTopicComment', '', '更新话题评论表');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('更新话题评论表状态', 3, 4, parentId, '', '/api/cms/topicComment/updateTopicCommentStatus', '', '更新话题评论表状态');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询话题评论表详情', 3, 5, parentId, '', '/api/cms/topicComment/queryTopicCommentDetail', '', '查询话题评论表详情');
        INSERT INTO sys_menu (menu_name, menu_type, menu_sort, parent_id, menu_url, api_url, menu_icon, remark) VALUES ('查询话题评论表列表', 3, 6, parentId, '', '/api/cms/topicComment/queryTopicCommentList', '', '查询话题评论表列表');
    END $$;



