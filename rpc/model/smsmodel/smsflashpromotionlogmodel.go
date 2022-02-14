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
	smsFlashPromotionLogFieldNames          = builderx.FieldNames(&SmsFlashPromotionLog{})
	smsFlashPromotionLogRows                = strings.Join(smsFlashPromotionLogFieldNames, ",")
	smsFlashPromotionLogRowsExpectAutoSet   = strings.Join(stringx.Remove(smsFlashPromotionLogFieldNames, "id", "create_time", "update_time"), ",")
	smsFlashPromotionLogRowsWithPlaceHolder = strings.Join(stringx.Remove(smsFlashPromotionLogFieldNames, "id", "create_time", "update_time"), "=?,") + "=?"
)

type (
	SmsFlashPromotionLogModel struct {
		conn  sqlx.SqlConn
		table string
	}

	SmsFlashPromotionLog struct {
		Id            int64     `db:"id"`
		MemberId      int64     `db:"member_id"`
		ProductId     int64     `db:"product_id"`
		MemberPhone   string    `db:"member_phone"`
		ProductName   string    `db:"product_name"`
		SubscribeTime time.Time `db:"subscribe_time"` // 会员订阅时间
		SendTime      time.Time `db:"send_time"`
	}
)

func NewSmsFlashPromotionLogModel(conn sqlx.SqlConn) *SmsFlashPromotionLogModel {
	return &SmsFlashPromotionLogModel{
		conn:  conn,
		table: "sms_flash_promotion_log",
	}
}

func (m *SmsFlashPromotionLogModel) Insert(data SmsFlashPromotionLog) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?)", m.table, smsFlashPromotionLogRowsExpectAutoSet)
	ret, err := m.conn.Exec(query, data.MemberId, data.ProductId, data.MemberPhone, data.ProductName, data.SubscribeTime, data.SendTime)
	return ret, err
}

func (m *SmsFlashPromotionLogModel) FindOne(id int64) (*SmsFlashPromotionLog, error) {
	query := fmt.Sprintf("select %s from %s where id = ? limit 1", smsFlashPromotionLogRows, m.table)
	var resp SmsFlashPromotionLog
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

func (m *SmsFlashPromotionLogModel) FindAll(Current int64, PageSize int64) (*[]SmsFlashPromotionLog, error) {

	query := fmt.Sprintf("select %s from %s limit ?,?", smsFlashPromotionLogRows, m.table)
	var resp []SmsFlashPromotionLog
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

func (m *SmsFlashPromotionLogModel) Count() (int64, error) {
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

func (m *SmsFlashPromotionLogModel) Update(data SmsFlashPromotionLog) error {
	query := fmt.Sprintf("update %s set %s where id = ?", m.table, smsFlashPromotionLogRowsWithPlaceHolder)
	_, err := m.conn.Exec(query, data.MemberId, data.ProductId, data.MemberPhone, data.ProductName, data.SubscribeTime, data.SendTime, data.Id)
	return err
}

func (m *SmsFlashPromotionLogModel) Delete(id int64) error {
	query := fmt.Sprintf("delete from %s where id = ?", m.table)
	_, err := m.conn.Exec(query, id)
	return err
}
