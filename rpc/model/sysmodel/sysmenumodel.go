package sysmodel

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"strings"
)

var _ SysMenuModel = (*customSysMenuModel)(nil)

type (
	// SysMenuModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSysMenuModel.
	SysMenuModel interface {
		sysMenuModel
		Count(ctx context.Context) (int64, error)
		FindAll(ctx context.Context, Current int64, PageSize int64) (*[]SysMenu, error)
		DeleteByIds(ctx context.Context, ids []int64) error
		FindAllByUserId(ctx context.Context, userId int64) (*[]SysMenu, error)
	}

	customSysMenuModel struct {
		*defaultSysMenuModel
	}
)

// NewSysMenuModel returns a model for the database table.
func NewSysMenuModel(conn sqlx.SqlConn) SysMenuModel {
	return &customSysMenuModel{
		defaultSysMenuModel: newSysMenuModel(conn),
	}
}

func (m *customSysMenuModel) FindAll(ctx context.Context, Current int64, PageSize int64) (*[]SysMenu, error) {

	query := fmt.Sprintf("select %s from %s limit ?,?", sysMenuRows, m.table)
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

func (m *customSysMenuModel) Count(ctx context.Context) (int64, error) {
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

func (m *customSysMenuModel) DeleteByIds(ctx context.Context, ids []int64) error {
	query := fmt.Sprintf("delete from %s where `id` in (?)", m.table)
	_, err := m.conn.ExecCtx(ctx, query, strings.Replace(strings.Trim(fmt.Sprint(ids), "[]"), " ", ",", -1))
	return err
}
func (m *customSysMenuModel) FindAllByUserId(ctx context.Context, userId int64) (*[]SysMenu, error) {

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
