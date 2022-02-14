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
	pmsMemberPriceFieldNames          = builderx.FieldNames(&PmsMemberPrice{})
	pmsMemberPriceRows                = strings.Join(pmsMemberPriceFieldNames, ",")
	pmsMemberPriceRowsExpectAutoSet   = strings.Join(stringx.Remove(pmsMemberPriceFieldNames, "id", "create_time", "update_time"), ",")
	pmsMemberPriceRowsWithPlaceHolder = strings.Join(stringx.Remove(pmsMemberPriceFieldNames, "id", "create_time", "update_time"), "=?,") + "=?"
)

type (
	PmsMemberPriceModel struct {
		conn  sqlx.SqlConn
		table string
	}

	PmsMemberPrice struct {
		Id              int64   `db:"id"`
		ProductId       int64   `db:"product_id"`
		MemberLevelId   int64   `db:"member_level_id"`
		MemberPrice     float64 `db:"member_price"` // 会员价格
		MemberLevelName string  `db:"member_level_name"`
	}
)

func NewPmsMemberPriceModel(conn sqlx.SqlConn) *PmsMemberPriceModel {
	return &PmsMemberPriceModel{
		conn:  conn,
		table: "pms_member_price",
	}
}

func (m *PmsMemberPriceModel) Insert(data PmsMemberPrice) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?)", m.table, pmsMemberPriceRowsExpectAutoSet)
	ret, err := m.conn.Exec(query, data.ProductId, data.MemberLevelId, data.MemberPrice, data.MemberLevelName)
	return ret, err
}

func (m *PmsMemberPriceModel) FindOne(id int64) (*PmsMemberPrice, error) {
	query := fmt.Sprintf("select %s from %s where id = ? limit 1", pmsMemberPriceRows, m.table)
	var resp PmsMemberPrice
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

func (m *PmsMemberPriceModel) FindAll(Current int64, PageSize int64) (*[]PmsMemberPrice, error) {

	query := fmt.Sprintf("select %s from %s limit ?,?", pmsMemberPriceRows, m.table)
	var resp []PmsMemberPrice
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

func (m *PmsMemberPriceModel) Count() (int64, error) {
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

func (m *PmsMemberPriceModel) Update(data PmsMemberPrice) error {
	query := fmt.Sprintf("update %s set %s where id = ?", m.table, pmsMemberPriceRowsWithPlaceHolder)
	_, err := m.conn.Exec(query, data.ProductId, data.MemberLevelId, data.MemberPrice, data.MemberLevelName, data.Id)
	return err
}

func (m *PmsMemberPriceModel) Delete(id int64) error {
	query := fmt.Sprintf("delete from %s where id = ?", m.table)
	_, err := m.conn.Exec(query, id)
	return err
}
