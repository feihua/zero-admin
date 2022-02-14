package pmsmodel

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
	pmsProductVertifyRecordFieldNames          = builderx.FieldNames(&PmsProductVertifyRecord{})
	pmsProductVertifyRecordRows                = strings.Join(pmsProductVertifyRecordFieldNames, ",")
	pmsProductVertifyRecordRowsExpectAutoSet   = strings.Join(stringx.Remove(pmsProductVertifyRecordFieldNames, "id", "create_time", "update_time"), ",")
	pmsProductVertifyRecordRowsWithPlaceHolder = strings.Join(stringx.Remove(pmsProductVertifyRecordFieldNames, "id", "create_time", "update_time"), "=?,") + "=?"
)

type (
	PmsProductVertifyRecordModel struct {
		conn  sqlx.SqlConn
		table string
	}

	PmsProductVertifyRecord struct {
		Id         int64     `db:"id"`
		ProductId  int64     `db:"product_id"`
		CreateTime time.Time `db:"create_time"`
		VertifyMan string    `db:"vertify_man"` // 审核人
		Status     int64     `db:"status"`
		Detail     string    `db:"detail"` // 反馈详情
	}
)

func NewPmsProductVertifyRecordModel(conn sqlx.SqlConn) *PmsProductVertifyRecordModel {
	return &PmsProductVertifyRecordModel{
		conn:  conn,
		table: "pms_product_vertify_record",
	}
}

func (m *PmsProductVertifyRecordModel) Insert(data PmsProductVertifyRecord) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?)", m.table, pmsProductVertifyRecordRowsExpectAutoSet)
	ret, err := m.conn.Exec(query, data.ProductId, data.VertifyMan, data.Status, data.Detail)
	return ret, err
}

func (m *PmsProductVertifyRecordModel) FindOne(id int64) (*PmsProductVertifyRecord, error) {
	query := fmt.Sprintf("select %s from %s where id = ? limit 1", pmsProductVertifyRecordRows, m.table)
	var resp PmsProductVertifyRecord
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

func (m *PmsProductVertifyRecordModel) FindAll(Current int64, PageSize int64) (*[]PmsProductVertifyRecord, error) {

	query := fmt.Sprintf("select %s from %s limit ?,?", pmsProductVertifyRecordRows, m.table)
	var resp []PmsProductVertifyRecord
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

func (m *PmsProductVertifyRecordModel) Count() (int64, error) {
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

func (m *PmsProductVertifyRecordModel) Update(data PmsProductVertifyRecord) error {
	query := fmt.Sprintf("update %s set %s where id = ?", m.table, pmsProductVertifyRecordRowsWithPlaceHolder)
	_, err := m.conn.Exec(query, data.ProductId, data.VertifyMan, data.Status, data.Detail, data.Id)
	return err
}

func (m *PmsProductVertifyRecordModel) Delete(id int64) error {
	query := fmt.Sprintf("delete from %s where id = ?", m.table)
	_, err := m.conn.Exec(query, id)
	return err
}
