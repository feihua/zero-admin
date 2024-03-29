// Code generated by goctl. DO NOT EDIT.

package pmsmodel

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	pmsProductVertifyRecordFieldNames          = builder.RawFieldNames(&PmsProductVertifyRecord{})
	pmsProductVertifyRecordRows                = strings.Join(pmsProductVertifyRecordFieldNames, ",")
	pmsProductVertifyRecordRowsExpectAutoSet   = strings.Join(stringx.Remove(pmsProductVertifyRecordFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	pmsProductVertifyRecordRowsWithPlaceHolder = strings.Join(stringx.Remove(pmsProductVertifyRecordFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	pmsProductVertifyRecordModel interface {
		Insert(ctx context.Context, data *PmsProductVertifyRecord) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*PmsProductVertifyRecord, error)
		Update(ctx context.Context, data *PmsProductVertifyRecord) error
		Delete(ctx context.Context, id int64) error
	}

	defaultPmsProductVertifyRecordModel struct {
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

func newPmsProductVertifyRecordModel(conn sqlx.SqlConn) *defaultPmsProductVertifyRecordModel {
	return &defaultPmsProductVertifyRecordModel{
		conn:  conn,
		table: "`pms_product_vertify_record`",
	}
}

func (m *defaultPmsProductVertifyRecordModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultPmsProductVertifyRecordModel) FindOne(ctx context.Context, id int64) (*PmsProductVertifyRecord, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", pmsProductVertifyRecordRows, m.table)
	var resp PmsProductVertifyRecord
	err := m.conn.QueryRowCtx(ctx, &resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultPmsProductVertifyRecordModel) Insert(ctx context.Context, data *PmsProductVertifyRecord) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?)", m.table, pmsProductVertifyRecordRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.ProductId, data.VertifyMan, data.Status, data.Detail)
	return ret, err
}

func (m *defaultPmsProductVertifyRecordModel) Update(ctx context.Context, data *PmsProductVertifyRecord) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, pmsProductVertifyRecordRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.ProductId, data.VertifyMan, data.Status, data.Detail, data.Id)
	return err
}

func (m *defaultPmsProductVertifyRecordModel) tableName() string {
	return m.table
}
