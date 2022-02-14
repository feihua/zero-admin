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
	pmsProductFullReductionFieldNames          = builderx.FieldNames(&PmsProductFullReduction{})
	pmsProductFullReductionRows                = strings.Join(pmsProductFullReductionFieldNames, ",")
	pmsProductFullReductionRowsExpectAutoSet   = strings.Join(stringx.Remove(pmsProductFullReductionFieldNames, "id", "create_time", "update_time"), ",")
	pmsProductFullReductionRowsWithPlaceHolder = strings.Join(stringx.Remove(pmsProductFullReductionFieldNames, "id", "create_time", "update_time"), "=?,") + "=?"
)

type (
	PmsProductFullReductionModel struct {
		conn  sqlx.SqlConn
		table string
	}

	PmsProductFullReduction struct {
		Id          int64   `db:"id"`
		ProductId   int64   `db:"product_id"`
		FullPrice   float64 `db:"full_price"`
		ReducePrice float64 `db:"reduce_price"`
	}
)

func NewPmsProductFullReductionModel(conn sqlx.SqlConn) *PmsProductFullReductionModel {
	return &PmsProductFullReductionModel{
		conn:  conn,
		table: "pms_product_full_reduction",
	}
}

func (m *PmsProductFullReductionModel) Insert(data PmsProductFullReduction) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?)", m.table, pmsProductFullReductionRowsExpectAutoSet)
	ret, err := m.conn.Exec(query, data.ProductId, data.FullPrice, data.ReducePrice)
	return ret, err
}

func (m *PmsProductFullReductionModel) FindOne(id int64) (*PmsProductFullReduction, error) {
	query := fmt.Sprintf("select %s from %s where id = ? limit 1", pmsProductFullReductionRows, m.table)
	var resp PmsProductFullReduction
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

func (m *PmsProductFullReductionModel) FindAll(Current int64, PageSize int64) (*[]PmsProductFullReduction, error) {

	query := fmt.Sprintf("select %s from %s limit ?,?", pmsProductFullReductionRows, m.table)
	var resp []PmsProductFullReduction
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

func (m *PmsProductFullReductionModel) Count() (int64, error) {
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

func (m *PmsProductFullReductionModel) Update(data PmsProductFullReduction) error {
	query := fmt.Sprintf("update %s set %s where id = ?", m.table, pmsProductFullReductionRowsWithPlaceHolder)
	_, err := m.conn.Exec(query, data.ProductId, data.FullPrice, data.ReducePrice, data.Id)
	return err
}

func (m *PmsProductFullReductionModel) Delete(id int64) error {
	query := fmt.Sprintf("delete from %s where id = ?", m.table)
	_, err := m.conn.Exec(query, id)
	return err
}
