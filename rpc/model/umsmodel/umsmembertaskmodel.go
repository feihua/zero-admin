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
	umsMemberTaskFieldNames          = builderx.FieldNames(&UmsMemberTask{})
	umsMemberTaskRows                = strings.Join(umsMemberTaskFieldNames, ",")
	umsMemberTaskRowsExpectAutoSet   = strings.Join(stringx.Remove(umsMemberTaskFieldNames, "id", "create_time", "update_time"), ",")
	umsMemberTaskRowsWithPlaceHolder = strings.Join(stringx.Remove(umsMemberTaskFieldNames, "id", "create_time", "update_time"), "=?,") + "=?"
)

type (
	UmsMemberTaskModel struct {
		conn  sqlx.SqlConn
		table string
	}

	UmsMemberTask struct {
		Id           int64  `db:"id"`
		Name         string `db:"name"`
		Growth       int64  `db:"growth"`       // 赠送成长值
		Intergration int64  `db:"intergration"` // 赠送积分
		Type         int64  `db:"type"`         // 任务类型：0->新手任务；1->日常任务
	}
)

func NewUmsMemberTaskModel(conn sqlx.SqlConn) *UmsMemberTaskModel {
	return &UmsMemberTaskModel{
		conn:  conn,
		table: "ums_member_task",
	}
}

func (m *UmsMemberTaskModel) Insert(data UmsMemberTask) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?)", m.table, umsMemberTaskRowsExpectAutoSet)
	ret, err := m.conn.Exec(query, data.Name, data.Growth, data.Intergration, data.Type)
	return ret, err
}

func (m *UmsMemberTaskModel) FindOne(id int64) (*UmsMemberTask, error) {
	query := fmt.Sprintf("select %s from %s where id = ? limit 1", umsMemberTaskRows, m.table)
	var resp UmsMemberTask
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

func (m *UmsMemberTaskModel) FindAll(Current int64, PageSize int64) (*[]UmsMemberTask, error) {

	query := fmt.Sprintf("select %s from %s limit ?,?", umsMemberTaskRows, m.table)
	var resp []UmsMemberTask
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

func (m *UmsMemberTaskModel) Count() (int64, error) {
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

func (m *UmsMemberTaskModel) Update(data UmsMemberTask) error {
	query := fmt.Sprintf("update %s set %s where id = ?", m.table, umsMemberTaskRowsWithPlaceHolder)
	_, err := m.conn.Exec(query, data.Name, data.Growth, data.Intergration, data.Type, data.Id)
	return err
}

func (m *UmsMemberTaskModel) Delete(id int64) error {
	query := fmt.Sprintf("delete from %s where id = ?", m.table)
	_, err := m.conn.Exec(query, id)
	return err
}
