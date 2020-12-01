package model

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/tal-tech/go-zero/core/stores/sqlc"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"github.com/tal-tech/go-zero/core/stringx"
	"github.com/tal-tech/go-zero/tools/goctl/model/sql/builderx"
)

var (
	smsFlashPromotionSessionFieldNames          = builderx.FieldNames(&SmsFlashPromotionSession{})
	smsFlashPromotionSessionRows                = strings.Join(smsFlashPromotionSessionFieldNames, ",")
	smsFlashPromotionSessionRowsExpectAutoSet   = strings.Join(stringx.Remove(smsFlashPromotionSessionFieldNames, "id", "create_time", "update_time"), ",")
	smsFlashPromotionSessionRowsWithPlaceHolder = strings.Join(stringx.Remove(smsFlashPromotionSessionFieldNames, "id", "create_time", "update_time"), "=?,") + "=?"
)

type (
	SmsFlashPromotionSessionModel struct {
		conn  sqlx.SqlConn
		table string
	}

	SmsFlashPromotionSession struct {
		Id         int64     `db:"id"`          // 编号
		Name       string    `db:"name"`        // 场次名称
		StartTime  string    `db:"start_time"`  // 每日开始时间
		EndTime    string    `db:"end_time"`    // 每日结束时间
		Status     int64     `db:"status"`      // 启用状态：0->不启用；1->启用
		CreateTime time.Time `db:"create_time"` // 创建时间
	}
)

func NewSmsFlashPromotionSessionModel(conn sqlx.SqlConn) *SmsFlashPromotionSessionModel {
	return &SmsFlashPromotionSessionModel{
		conn:  conn,
		table: "sms_flash_promotion_session",
	}
}

func (m *SmsFlashPromotionSessionModel) Insert(data SmsFlashPromotionSession) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?)", m.table, smsFlashPromotionSessionRowsExpectAutoSet)
	ret, err := m.conn.Exec(query, data.Name, data.StartTime, data.EndTime, data.Status)
	return ret, err
}

func (m *SmsFlashPromotionSessionModel) FindOne(id int64) (*SmsFlashPromotionSession, error) {
	query := fmt.Sprintf("select %s from %s where id = ? limit 1", smsFlashPromotionSessionRows, m.table)
	var resp SmsFlashPromotionSession
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

func (m *SmsFlashPromotionSessionModel) Update(data SmsFlashPromotionSession) error {
	query := fmt.Sprintf("update %s set %s where id = ?", m.table, smsFlashPromotionSessionRowsWithPlaceHolder)
	_, err := m.conn.Exec(query, data.Name, data.StartTime, data.EndTime, data.Status, data.Id)
	return err
}

func (m *SmsFlashPromotionSessionModel) Delete(id int64) error {
	query := fmt.Sprintf("delete from %s where id = ?", m.table)
	_, err := m.conn.Exec(query, id)
	return err
}
