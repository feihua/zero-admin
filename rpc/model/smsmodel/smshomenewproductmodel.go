package smsmodel

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
	smsHomeNewProductFieldNames          = builderx.FieldNames(&SmsHomeNewProduct{})
	smsHomeNewProductRows                = strings.Join(smsHomeNewProductFieldNames, ",")
	smsHomeNewProductRowsExpectAutoSet   = strings.Join(stringx.Remove(smsHomeNewProductFieldNames, "id", "create_time", "update_time"), ",")
	smsHomeNewProductRowsWithPlaceHolder = strings.Join(stringx.Remove(smsHomeNewProductFieldNames, "id", "create_time", "update_time"), "=?,") + "=?"
)

type (
	SmsHomeNewProductModel struct {
		conn  sqlx.SqlConn
		table string
	}

	SmsHomeNewProduct struct {
		Id              int64  `db:"id"`
		ProductId       int64  `db:"product_id"`
		ProductName     string `db:"product_name"`
		RecommendStatus int64  `db:"recommend_status"`
		Sort            int64  `db:"sort"`
	}
)

func NewSmsHomeNewProductModel(conn sqlx.SqlConn) *SmsHomeNewProductModel {
	return &SmsHomeNewProductModel{
		conn:  conn,
		table: "sms_home_new_product",
	}
}

func (m *SmsHomeNewProductModel) Insert(data SmsHomeNewProduct) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?)", m.table, smsHomeNewProductRowsExpectAutoSet)
	ret, err := m.conn.Exec(query, data.ProductId, data.ProductName, data.RecommendStatus, data.Sort)
	return ret, err
}

func (m *SmsHomeNewProductModel) FindOne(id int64) (*SmsHomeNewProduct, error) {
	query := fmt.Sprintf("select %s from %s where id = ? limit 1", smsHomeNewProductRows, m.table)
	var resp SmsHomeNewProduct
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

func (m *SmsHomeNewProductModel) FindAll(Current int64, PageSize int64) (*[]SmsHomeNewProduct, error) {

	query := fmt.Sprintf("select %s from %s limit ?,?", smsHomeNewProductRows, m.table)
	var resp []SmsHomeNewProduct
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

func (m *SmsHomeNewProductModel) Count() (int64, error) {
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

func (m *SmsHomeNewProductModel) Update(data SmsHomeNewProduct) error {
	query := fmt.Sprintf("update %s set %s where id = ?", m.table, smsHomeNewProductRowsWithPlaceHolder)
	_, err := m.conn.Exec(query, data.ProductId, data.ProductName, data.RecommendStatus, data.Sort, data.Id)
	return err
}

func (m *SmsHomeNewProductModel) Delete(id int64) error {
	query := fmt.Sprintf("delete from %s where id = ?", m.table)
	_, err := m.conn.Exec(query, id)
	return err
}
