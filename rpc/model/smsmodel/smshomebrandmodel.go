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
	smsHomeBrandFieldNames          = builderx.FieldNames(&SmsHomeBrand{})
	smsHomeBrandRows                = strings.Join(smsHomeBrandFieldNames, ",")
	smsHomeBrandRowsExpectAutoSet   = strings.Join(stringx.Remove(smsHomeBrandFieldNames, "id", "create_time", "update_time"), ",")
	smsHomeBrandRowsWithPlaceHolder = strings.Join(stringx.Remove(smsHomeBrandFieldNames, "id", "create_time", "update_time"), "=?,") + "=?"
)

type (
	SmsHomeBrandModel struct {
		conn  sqlx.SqlConn
		table string
	}

	SmsHomeBrand struct {
		Id              int64  `db:"id"`
		BrandId         int64  `db:"brand_id"`
		BrandName       string `db:"brand_name"`
		RecommendStatus int64  `db:"recommend_status"`
		Sort            int64  `db:"sort"`
	}
)

func NewSmsHomeBrandModel(conn sqlx.SqlConn) *SmsHomeBrandModel {
	return &SmsHomeBrandModel{
		conn:  conn,
		table: "sms_home_brand",
	}
}

func (m *SmsHomeBrandModel) Insert(data SmsHomeBrand) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?)", m.table, smsHomeBrandRowsExpectAutoSet)
	ret, err := m.conn.Exec(query, data.BrandId, data.BrandName, data.RecommendStatus, data.Sort)
	return ret, err
}

func (m *SmsHomeBrandModel) FindOne(id int64) (*SmsHomeBrand, error) {
	query := fmt.Sprintf("select %s from %s where id = ? limit 1", smsHomeBrandRows, m.table)
	var resp SmsHomeBrand
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

func (m *SmsHomeBrandModel) FindAll(Current int64, PageSize int64) (*[]SmsHomeBrand, error) {

	query := fmt.Sprintf("select %s from %s limit ?,?", smsHomeBrandRows, m.table)
	var resp []SmsHomeBrand
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

func (m *SmsHomeBrandModel) Count() (int64, error) {
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

func (m *SmsHomeBrandModel) Update(data SmsHomeBrand) error {
	query := fmt.Sprintf("update %s set %s where id = ?", m.table, smsHomeBrandRowsWithPlaceHolder)
	_, err := m.conn.Exec(query, data.BrandId, data.BrandName, data.RecommendStatus, data.Sort, data.Id)
	return err
}

func (m *SmsHomeBrandModel) Delete(id int64) error {
	query := fmt.Sprintf("delete from %s where id = ?", m.table)
	_, err := m.conn.Exec(query, id)
	return err
}
