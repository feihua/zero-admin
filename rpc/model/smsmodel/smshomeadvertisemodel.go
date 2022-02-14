package smsmodel

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
	smsHomeAdvertiseFieldNames          = builderx.FieldNames(&SmsHomeAdvertise{})
	smsHomeAdvertiseRows                = strings.Join(smsHomeAdvertiseFieldNames, ",")
	smsHomeAdvertiseRowsExpectAutoSet   = strings.Join(stringx.Remove(smsHomeAdvertiseFieldNames, "id", "create_time", "update_time"), ",")
	smsHomeAdvertiseRowsWithPlaceHolder = strings.Join(stringx.Remove(smsHomeAdvertiseFieldNames, "id", "create_time", "update_time"), "=?,") + "=?"
)

type (
	SmsHomeAdvertiseModel struct {
		conn  sqlx.SqlConn
		table string
	}

	SmsHomeAdvertise struct {
		Id         int64     `db:"id"`
		Name       string    `db:"name"`
		Type       int64     `db:"type"` // 轮播位置：0->PC首页轮播；1->app首页轮播
		Pic        string    `db:"pic"`
		StartTime  time.Time `db:"start_time"`
		EndTime    time.Time `db:"end_time"`
		Status     int64     `db:"status"`      // 上下线状态：0->下线；1->上线
		ClickCount int64     `db:"click_count"` // 点击数
		OrderCount int64     `db:"order_count"` // 下单数
		Url        string    `db:"url"`         // 链接地址
		Note       string    `db:"note"`        // 备注
		Sort       int64     `db:"sort"`        // 排序
	}
)

func NewSmsHomeAdvertiseModel(conn sqlx.SqlConn) *SmsHomeAdvertiseModel {
	return &SmsHomeAdvertiseModel{
		conn:  conn,
		table: "sms_home_advertise",
	}
}

func (m *SmsHomeAdvertiseModel) Insert(data SmsHomeAdvertise) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, smsHomeAdvertiseRowsExpectAutoSet)
	ret, err := m.conn.Exec(query, data.Name, data.Type, data.Pic, data.StartTime, data.EndTime, data.Status, data.ClickCount, data.OrderCount, data.Url, data.Note, data.Sort)
	return ret, err
}

func (m *SmsHomeAdvertiseModel) FindOne(id int64) (*SmsHomeAdvertise, error) {
	query := fmt.Sprintf("select %s from %s where id = ? limit 1", smsHomeAdvertiseRows, m.table)
	var resp SmsHomeAdvertise
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

func (m *SmsHomeAdvertiseModel) FindAll(Current int64, PageSize int64) (*[]SmsHomeAdvertise, error) {

	query := fmt.Sprintf("select %s from %s limit ?,?", smsHomeAdvertiseRows, m.table)
	var resp []SmsHomeAdvertise
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

func (m *SmsHomeAdvertiseModel) Count() (int64, error) {
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

func (m *SmsHomeAdvertiseModel) Update(data SmsHomeAdvertise) error {
	query := fmt.Sprintf("update %s set %s where id = ?", m.table, smsHomeAdvertiseRowsWithPlaceHolder)
	_, err := m.conn.Exec(query, data.Name, data.Type, data.Pic, data.StartTime, data.EndTime, data.Status, data.ClickCount, data.OrderCount, data.Url, data.Note, data.Sort, data.Id)
	return err
}

func (m *SmsHomeAdvertiseModel) Delete(id int64) error {
	query := fmt.Sprintf("delete from %s where id = ?", m.table)
	_, err := m.conn.Exec(query, id)
	return err
}
