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
	smsHomeRecommendProductFieldNames          = builderx.FieldNames(&SmsHomeRecommendProduct{})
	smsHomeRecommendProductRows                = strings.Join(smsHomeRecommendProductFieldNames, ",")
	smsHomeRecommendProductRowsExpectAutoSet   = strings.Join(stringx.Remove(smsHomeRecommendProductFieldNames, "id", "create_time", "update_time"), ",")
	smsHomeRecommendProductRowsWithPlaceHolder = strings.Join(stringx.Remove(smsHomeRecommendProductFieldNames, "id", "create_time", "update_time"), "=?,") + "=?"
)

type (
	SmsHomeRecommendProductModel struct {
		conn  sqlx.SqlConn
		table string
	}

	SmsHomeRecommendProduct struct {
		Id              int64  `db:"id"`
		ProductId       int64  `db:"product_id"`
		ProductName     string `db:"product_name"`
		RecommendStatus int64  `db:"recommend_status"`
		Sort            int64  `db:"sort"`
	}
)

func NewSmsHomeRecommendProductModel(conn sqlx.SqlConn) *SmsHomeRecommendProductModel {
	return &SmsHomeRecommendProductModel{
		conn:  conn,
		table: "sms_home_recommend_product",
	}
}

func (m *SmsHomeRecommendProductModel) Insert(data SmsHomeRecommendProduct) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?)", m.table, smsHomeRecommendProductRowsExpectAutoSet)
	ret, err := m.conn.Exec(query, data.ProductId, data.ProductName, data.RecommendStatus, data.Sort)
	return ret, err
}

func (m *SmsHomeRecommendProductModel) FindOne(id int64) (*SmsHomeRecommendProduct, error) {
	query := fmt.Sprintf("select %s from %s where id = ? limit 1", smsHomeRecommendProductRows, m.table)
	var resp SmsHomeRecommendProduct
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

func (m *SmsHomeRecommendProductModel) FindAll(Current int64, PageSize int64) (*[]SmsHomeRecommendProduct, error) {

	query := fmt.Sprintf("select %s from %s limit ?,?", smsHomeRecommendProductRows, m.table)
	var resp []SmsHomeRecommendProduct
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

func (m *SmsHomeRecommendProductModel) Count() (int64, error) {
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

func (m *SmsHomeRecommendProductModel) Update(data SmsHomeRecommendProduct) error {
	query := fmt.Sprintf("update %s set %s where id = ?", m.table, smsHomeRecommendProductRowsWithPlaceHolder)
	_, err := m.conn.Exec(query, data.ProductId, data.ProductName, data.RecommendStatus, data.Sort, data.Id)
	return err
}

func (m *SmsHomeRecommendProductModel) Delete(id int64) error {
	query := fmt.Sprintf("delete from %s where id = ?", m.table)
	_, err := m.conn.Exec(query, id)
	return err
}
