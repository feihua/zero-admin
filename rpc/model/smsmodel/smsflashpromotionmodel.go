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
	smsFlashPromotionFieldNames          = builderx.FieldNames(&SmsFlashPromotion{})
	smsFlashPromotionRows                = strings.Join(smsFlashPromotionFieldNames, ",")
	smsFlashPromotionRowsExpectAutoSet   = strings.Join(stringx.Remove(smsFlashPromotionFieldNames, "id", "create_time", "update_time"), ",")
	smsFlashPromotionRowsWithPlaceHolder = strings.Join(stringx.Remove(smsFlashPromotionFieldNames, "id", "create_time", "update_time"), "=?,") + "=?"
)

type (
	SmsFlashPromotionModel struct {
		conn  sqlx.SqlConn
		table string
	}

	SmsFlashPromotion struct {
		Id         int64     `db:"id"`
		Title      string    `db:"title"`
		StartDate  time.Time `db:"start_date"`  // 开始日期
		EndDate    time.Time `db:"end_date"`    // 结束日期
		Status     int64     `db:"status"`      // 上下线状态
		CreateTime time.Time `db:"create_time"` // 秒杀时间段名称
	}
)

func NewSmsFlashPromotionModel(conn sqlx.SqlConn) *SmsFlashPromotionModel {
	return &SmsFlashPromotionModel{
		conn:  conn,
		table: "sms_flash_promotion",
	}
}

func (m *SmsFlashPromotionModel) Insert(data SmsFlashPromotion) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?)", m.table, smsFlashPromotionRowsExpectAutoSet)
	ret, err := m.conn.Exec(query, data.Title, data.StartDate, data.EndDate, data.Status)
	return ret, err
}

func (m *SmsFlashPromotionModel) FindOne(id int64) (*SmsFlashPromotion, error) {
	query := fmt.Sprintf("select %s from %s where id = ? limit 1", smsFlashPromotionRows, m.table)
	var resp SmsFlashPromotion
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

func (m *SmsFlashPromotionModel) FindAll(Current int64, PageSize int64) (*[]SmsFlashPromotion, error) {

	query := fmt.Sprintf("select %s from %s limit ?,?", smsFlashPromotionRows, m.table)
	var resp []SmsFlashPromotion
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

func (m *SmsFlashPromotionModel) Count() (int64, error) {
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

func (m *SmsFlashPromotionModel) Update(data SmsFlashPromotion) error {
	query := fmt.Sprintf("update %s set %s where id = ?", m.table, smsFlashPromotionRowsWithPlaceHolder)
	_, err := m.conn.Exec(query, data.Title, data.StartDate, data.EndDate, data.Status, data.Id)
	return err
}

func (m *SmsFlashPromotionModel) Delete(id int64) error {
	query := fmt.Sprintf("delete from %s where id = ?", m.table)
	_, err := m.conn.Exec(query, id)
	return err
}
