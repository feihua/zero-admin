package sysmodel

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
	"github.com/zeromicro/go-zero/tools/goctl/model/sql/builderx"
)

var (
	sysMenuFieldNames          = builderx.FieldNames(&SysMenu{})
	sysMenuRows                = strings.Join(sysMenuFieldNames, ",")
	sysMenuRowsExpectAutoSet   = strings.Join(stringx.Remove(sysMenuFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	sysMenuRowsWithPlaceHolder = strings.Join(stringx.Remove(sysMenuFieldNames, "`id`", "`create_by`", "`create_time`", "`update_time`"), "=?,") + "=?"
)

type (
	SysMenuModel struct {
		conn  sqlx.SqlConn
		table string
	}

	SysMenu struct {
		Id             int64     `db:"id"`               // 编号
		Name           string    `db:"name"`             // 菜单名称
		ParentId       int64     `db:"parent_id"`        // 父菜单ID，一级菜单为0
		Url            string    `db:"url"`              // 菜单URL,类型：1.普通页面（如用户管理， /sysmodel/user） 2.嵌套完整外部页面，以http(s)开头的链接 3.嵌套服务器页面，使用iframe:前缀+目标URL(如SQL监控， iframe:/druid/login.html, iframe:前缀会替换成服务器地址)
		Perms          string    `db:"perms"`            // 授权(多个用逗号分隔，如：sysmodel:user:add,sysmodel:user:edit)
		Type           int64     `db:"type"`             // 类型   0：目录   1：菜单   2：按钮
		Icon           string    `db:"icon"`             // 菜单图标
		OrderNum       int64     `db:"order_num"`        // 排序
		CreateBy       string    `db:"create_by"`        // 创建人
		CreateTime     time.Time `db:"create_time"`      // 创建时间
		LastUpdateBy   string    `db:"last_update_by"`   // 更新人
		LastUpdateTime time.Time `db:"last_update_time"` // 更新时间
		DelFlag        int64     `db:"del_flag"`         // 是否删除  -1：已删除  0：正常
		VuePath        string    `db:"vue_path"`         // vue系统的path
		VueComponent   string    `db:"vue_component"`    // vue的页面
		VueIcon        string    `db:"vue_icon"`         // vue的图标
		VueRedirect    string    `db:"vue_redirect"`     // vue的路由重定向
		BackgroundUrl  string    `db:"background_url"`   // 后台地址
	}
)

func NewSysMenuModel(conn sqlx.SqlConn) *SysMenuModel {
	return &SysMenuModel{
		conn:  conn,
		table: "sys_menu",
	}
}

func (m *SysMenuModel) Insert(data SysMenu) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, sysMenuRowsExpectAutoSet)
	ret, err := m.conn.Exec(query, data.Name, data.ParentId, data.Url, data.Perms, data.Type, data.Icon, data.OrderNum, data.CreateBy, data.LastUpdateBy, data.LastUpdateTime, data.DelFlag, data.VuePath, data.VueComponent, data.VueIcon, data.VueRedirect)
	return ret, err
}

func (m *SysMenuModel) FindOne(id int64) (*SysMenu, error) {
	query := fmt.Sprintf("select %s from %s where id = ? limit 1", sysMenuRows, m.table)
	var resp SysMenu
	err := m.conn.QueryRow(&resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *SysMenuModel) FindAll(Current int64, PageSize int64) (*[]SysMenu, error) {

	query := fmt.Sprintf("select %s from %s order by id limit ?,?", sysMenuRows, m.table)
	var resp []SysMenu
	err := m.conn.QueryRows(&resp, query, (Current-1)*PageSize, PageSize)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *SysMenuModel) FindAllByUserId(userId int64) (*[]SysMenu, error) {

	query := "select sm.* from sys_user_role sur left join sys_role sr on sur.role_id = sr.id left join sys_role_menu srm on sr.id = srm.role_id left join sys_menu sm on srm.menu_id = sm.id where sur.user_id=? order by sm.id"
	var resp []SysMenu
	err := m.conn.QueryRows(&resp, query, userId)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *SysMenuModel) Count() (int64, error) {
	query := fmt.Sprintf("select count(*) as count from %s", m.table)

	var count int64
	err := m.conn.QueryRow(&count, query)

	switch err {
	case nil:
		return count, nil
	case sqlc.ErrNotFound:
		return 0, ErrNotFound
	default:
		return 0, err
	}
}
func (m *SysMenuModel) Update(data SysMenu) error {
	query := fmt.Sprintf("update %s set %s where id = ?", m.table, sysMenuRowsWithPlaceHolder)
	_, err := m.conn.Exec(query, data.Name, data.ParentId, data.Url, data.Perms, data.Type, data.Icon, data.OrderNum, data.LastUpdateBy, data.LastUpdateTime, data.DelFlag, data.VuePath, data.VueComponent, data.VueIcon, data.VueRedirect, data.Id)
	return err
}

func (m *SysMenuModel) Delete(id int64) error {
	query := fmt.Sprintf("delete from %s where id = ?", m.table)
	_, err := m.conn.Exec(query, id)
	return err
}
