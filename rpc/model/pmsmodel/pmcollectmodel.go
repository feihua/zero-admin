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
	pmsCollectFieldNames          = builderx.FieldNames(&PmsCollect{})
	pmsCollectRows                = strings.Join(pmsCollectFieldNames, ",")
	pmsCollectRowsExpectAutoSet   = strings.Join(stringx.Remove(pmsCollectFieldNames, "id", "create_time", "update_time"), ",")
	pmsCollectRowsWithPlaceHolder = strings.Join(stringx.Remove(pmsCollectFieldNames, "id", "create_time", "update_time"), "=?,") + "=?"
)

type (
	PmsCollectModel struct {
		conn  sqlx.SqlConn
		table string
	}

	PmsCollect struct {
		Id          int32     `db:"id"`
		UserId      int32     `db:"user_id"`      // 用户表的用户ID
		ValueId     int32     `db:"value_id"`     // 如果collect_type=0，则是商品ID；如果collect_type=1，则是专题ID
		CollectType int32     `db:"collect_type"` // 收藏类型，如果collect_type=0，则是商品ID；如果collect_typee=1，则是专题ID
		AddTime     time.Time `db:"add_time"`     // 创建时间
		UpdateTime  time.Time `db:"update_time"`  // 更新时间
		Deleted     int8      `db:"deleted"`      // 逻辑删除
	}
)

func NewPmsCollectModel(conn sqlx.SqlConn) *PmsCollectModel {
	return &PmsCollectModel{
		conn:  conn,
		table: "pms_product_collect",
	}
}

func (m *PmsCollectModel) Insert(data PmsCollect) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?)", m.table, pmsCollectRowsExpectAutoSet)
	ret, err := m.conn.Exec(query, data.UserId, data.ValueId, data.CollectType, data.AddTime)
	return ret, err
}

func (m *PmsCollectModel) FindOne(MemberId int32, ValueId int32) (*PmsCollect, error) {
	query := fmt.Sprintf("select %s from %s where user_id = ? and value_id = ? limit 1", pmsCollectRows, m.table)
	var resp PmsCollect
	err := m.conn.QueryRow(&resp, query, MemberId, ValueId)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *PmsCollectModel) FindAll(MemberId int64) (*[]PmsCollect, error) {

	query := fmt.Sprintf("select %s from %s where user_id = ?", pmsCollectRows, m.table)
	var resp []PmsCollect
	err := m.conn.QueryRows(&resp, query, MemberId)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *PmsCollectModel) Count(MemberId int64) (int64, error) {
	query := fmt.Sprintf("select count(*) as count from %s where user_id = ? ", m.table)

	var count int64
	err := m.conn.QueryRow(&count, query, MemberId)

	switch err {
	case nil:
		return count, nil
	case sqlc.ErrNotFound:
		return 0, ErrNotFound
	default:
		return 0, err
	}
}

func (m *PmsCollectModel) Delete(id int64) error {
	query := fmt.Sprintf("delete from %s where id = ?", m.table)
	_, err := m.conn.Exec(query, id)
	return err
}
