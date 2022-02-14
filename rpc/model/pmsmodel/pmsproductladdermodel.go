package pmsmodel

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
	pmsProductLadderFieldNames          = builderx.FieldNames(&PmsProductLadder{})
	pmsProductLadderRows                = strings.Join(pmsProductLadderFieldNames, ",")
	pmsProductLadderRowsExpectAutoSet   = strings.Join(stringx.Remove(pmsProductLadderFieldNames, "id", "create_time", "update_time"), ",")
	pmsProductLadderRowsWithPlaceHolder = strings.Join(stringx.Remove(pmsProductLadderFieldNames, "id", "create_time", "update_time"), "=?,") + "=?"
)

type (
	PmsProductLadderModel struct {
		conn  sqlx.SqlConn
		table string
	}

	PmsProductLadder struct {
		Id        int64   `db:"id"`
		ProductId int64   `db:"product_id"`
		Count     int64   `db:"count"`    // 满足的商品数量
		Discount  float64 `db:"discount"` // 折扣
		Price     float64 `db:"price"`    // 折后价格
	}
)

func NewPmsProductLadderModel(conn sqlx.SqlConn) *PmsProductLadderModel {
	return &PmsProductLadderModel{
		conn:  conn,
		table: "pms_product_ladder",
	}
}

func (m *PmsProductLadderModel) Insert(data PmsProductLadder) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?)", m.table, pmsProductLadderRowsExpectAutoSet)
	ret, err := m.conn.Exec(query, data.ProductId, data.Count, data.Discount, data.Price)
	return ret, err
}

func (m *PmsProductLadderModel) FindOne(id int64) (*PmsProductLadder, error) {
	query := fmt.Sprintf("select %s from %s where id = ? limit 1", pmsProductLadderRows, m.table)
	var resp PmsProductLadder
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

func (m *PmsProductLadderModel) FindAll(Current int64, PageSize int64) (*[]PmsProductLadder, error) {

	query := fmt.Sprintf("select %s from %s limit ?,?", pmsProductLadderRows, m.table)
	var resp []PmsProductLadder
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

func (m *PmsProductLadderModel) Count() (int64, error) {
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

func (m *PmsProductLadderModel) Update(data PmsProductLadder) error {
	query := fmt.Sprintf("update %s set %s where id = ?", m.table, pmsProductLadderRowsWithPlaceHolder)
	_, err := m.conn.Exec(query, data.ProductId, data.Count, data.Discount, data.Price, data.Id)
	return err
}

func (m *PmsProductLadderModel) Delete(id int64) error {
	query := fmt.Sprintf("delete from %s where id = ?", m.table)
	_, err := m.conn.Exec(query, id)
	return err
}
