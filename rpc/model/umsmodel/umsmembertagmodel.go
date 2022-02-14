package umsmodel

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
	"github.com/zeromicro/go-zero/tools/goctl/model/sql/builderx"
)

var (
	umsMemberTagFieldNames          = builderx.FieldNames(&UmsMemberTag{})
	umsMemberTagRows                = strings.Join(umsMemberTagFieldNames, ",")
	umsMemberTagRowsExpectAutoSet   = strings.Join(stringx.Remove(umsMemberTagFieldNames, "id", "create_time", "update_time"), ",")
	umsMemberTagRowsWithPlaceHolder = strings.Join(stringx.Remove(umsMemberTagFieldNames, "id", "create_time", "update_time"), "=?,") + "=?"
)

type (
	UmsMemberTagModel struct {
		conn  sqlx.SqlConn
		table string
	}

	UmsMemberTag struct {
		Id                int64   `db:"id"`
		Name              string  `db:"name"`
		FinishOrderCount  int64   `db:"finish_order_count"`  // 自动打标签完成订单数量
		FinishOrderAmount float64 `db:"finish_order_amount"` // 自动打标签完成订单金额
	}
)

func NewUmsMemberTagModel(conn sqlx.SqlConn) *UmsMemberTagModel {
	return &UmsMemberTagModel{
		conn:  conn,
		table: "ums_member_tag",
	}
}

func (m *UmsMemberTagModel) Insert(data UmsMemberTag) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?)", m.table, umsMemberTagRowsExpectAutoSet)
	ret, err := m.conn.Exec(query, data.Name, data.FinishOrderCount, data.FinishOrderAmount)
	return ret, err
}

func (m *UmsMemberTagModel) FindOne(id int64) (*UmsMemberTag, error) {
	query := fmt.Sprintf("select %s from %s where id = ? limit 1", umsMemberTagRows, m.table)
	var resp UmsMemberTag
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

func (m *UmsMemberTagModel) FindAll(Current int64, PageSize int64) (*[]UmsMemberTag, error) {

	query := fmt.Sprintf("select %s from %s limit ?,?", umsMemberTagRows, m.table)
	var resp []UmsMemberTag
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

func (m *UmsMemberTagModel) Count() (int64, error) {
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

func (m *UmsMemberTagModel) Update(data UmsMemberTag) error {
	query := fmt.Sprintf("update %s set %s where id = ?", m.table, umsMemberTagRowsWithPlaceHolder)
	_, err := m.conn.Exec(query, data.Name, data.FinishOrderCount, data.FinishOrderAmount, data.Id)
	return err
}

func (m *UmsMemberTagModel) Delete(id int64) error {
	query := fmt.Sprintf("delete from %s where id = ?", m.table)
	_, err := m.conn.Exec(query, id)
	return err
}
