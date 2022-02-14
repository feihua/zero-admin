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
	smsFlashPromotionProductRelationFieldNames          = builderx.FieldNames(&SmsFlashPromotionProductRelation{})
	smsFlashPromotionProductRelationRows                = strings.Join(smsFlashPromotionProductRelationFieldNames, ",")
	smsFlashPromotionProductRelationRowsExpectAutoSet   = strings.Join(stringx.Remove(smsFlashPromotionProductRelationFieldNames, "id", "create_time", "update_time"), ",")
	smsFlashPromotionProductRelationRowsWithPlaceHolder = strings.Join(stringx.Remove(smsFlashPromotionProductRelationFieldNames, "id", "create_time", "update_time"), "=?,") + "=?"
)

type (
	SmsFlashPromotionProductRelationModel struct {
		conn  sqlx.SqlConn
		table string
	}

	SmsFlashPromotionProductRelation struct {
		Id                      int64   `db:"id"` // 编号
		FlashPromotionId        int64   `db:"flash_promotion_id"`
		FlashPromotionSessionId int64   `db:"flash_promotion_session_id"` // 编号
		ProductId               int64   `db:"product_id"`
		FlashPromotionPrice     float64 `db:"flash_promotion_price"` // 限时购价格
		FlashPromotionCount     int64   `db:"flash_promotion_count"` // 限时购数量
		FlashPromotionLimit     int64   `db:"flash_promotion_limit"` // 每人限购数量
		Sort                    int64   `db:"sort"`                  // 排序
	}
)

func NewSmsFlashPromotionProductRelationModel(conn sqlx.SqlConn) *SmsFlashPromotionProductRelationModel {
	return &SmsFlashPromotionProductRelationModel{
		conn:  conn,
		table: "sms_flash_promotion_product_relation",
	}
}

func (m *SmsFlashPromotionProductRelationModel) Insert(data SmsFlashPromotionProductRelation) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?)", m.table, smsFlashPromotionProductRelationRowsExpectAutoSet)
	ret, err := m.conn.Exec(query, data.FlashPromotionId, data.FlashPromotionSessionId, data.ProductId, data.FlashPromotionPrice, data.FlashPromotionCount, data.FlashPromotionLimit, data.Sort)
	return ret, err
}

func (m *SmsFlashPromotionProductRelationModel) FindOne(id int64) (*SmsFlashPromotionProductRelation, error) {
	query := fmt.Sprintf("select %s from %s where id = ? limit 1", smsFlashPromotionProductRelationRows, m.table)
	var resp SmsFlashPromotionProductRelation
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

func (m *SmsFlashPromotionProductRelationModel) FindAll(Current int64, PageSize int64) (*[]SmsFlashPromotionProductRelation, error) {

	query := fmt.Sprintf("select %s from %s limit ?,?", smsFlashPromotionProductRelationRows, m.table)
	var resp []SmsFlashPromotionProductRelation
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

func (m *SmsFlashPromotionProductRelationModel) Count() (int64, error) {
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

func (m *SmsFlashPromotionProductRelationModel) Update(data SmsFlashPromotionProductRelation) error {
	query := fmt.Sprintf("update %s set %s where id = ?", m.table, smsFlashPromotionProductRelationRowsWithPlaceHolder)
	_, err := m.conn.Exec(query, data.FlashPromotionId, data.FlashPromotionSessionId, data.ProductId, data.FlashPromotionPrice, data.FlashPromotionCount, data.FlashPromotionLimit, data.Sort, data.Id)
	return err
}

func (m *SmsFlashPromotionProductRelationModel) Delete(id int64) error {
	query := fmt.Sprintf("delete from %s where id = ?", m.table)
	_, err := m.conn.Exec(query, id)
	return err
}
