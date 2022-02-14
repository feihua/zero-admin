package pmsmodel

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
	pmsCommentFieldNames          = builderx.FieldNames(&PmsComment{})
	pmsCommentRows                = strings.Join(pmsCommentFieldNames, ",")
	pmsCommentRowsExpectAutoSet   = strings.Join(stringx.Remove(pmsCommentFieldNames, "id", "create_time", "update_time"), ",")
	pmsCommentRowsWithPlaceHolder = strings.Join(stringx.Remove(pmsCommentFieldNames, "id", "create_time", "update_time"), "=?,") + "=?"
)

type (
	PmsCommentModel struct {
		conn  sqlx.SqlConn
		table string
	}

	PmsComment struct {
		Id               int64     `db:"id"`
		ProductId        int64     `db:"product_id"`
		MemberNickName   string    `db:"member_nick_name"`
		ProductName      string    `db:"product_name"`
		Star             int64     `db:"star"`      // 评价星数：0->5
		MemberIp         string    `db:"member_ip"` // 评价的ip
		CreateTime       time.Time `db:"create_time"`
		ShowStatus       int64     `db:"show_status"`
		ProductAttribute string    `db:"product_attribute"` // 购买时的商品属性
		CollectCouont    int64     `db:"collect_couont"`
		ReadCount        int64     `db:"read_count"`
		Content          string    `db:"content"`
		Pics             string    `db:"pics"`        // 上传图片地址，以逗号隔开
		MemberIcon       string    `db:"member_icon"` // 评论用户头像
		ReplayCount      int64     `db:"replay_count"`
	}
)

func NewPmsCommentModel(conn sqlx.SqlConn) *PmsCommentModel {
	return &PmsCommentModel{
		conn:  conn,
		table: "pms_comment",
	}
}

func (m *PmsCommentModel) Insert(data PmsComment) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, pmsCommentRowsExpectAutoSet)
	ret, err := m.conn.Exec(query, data.ProductId, data.MemberNickName, data.ProductName, data.Star, data.MemberIp, data.ShowStatus, data.ProductAttribute, data.CollectCouont, data.ReadCount, data.Content, data.Pics, data.MemberIcon, data.ReplayCount)
	return ret, err
}

func (m *PmsCommentModel) FindOne(id int64) (*PmsComment, error) {
	query := fmt.Sprintf("select %s from %s where id = ? limit 1", pmsCommentRows, m.table)
	var resp PmsComment
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

func (m *PmsCommentModel) FindAll(Current int64, PageSize int64) (*[]PmsComment, error) {

	query := fmt.Sprintf("select %s from %s limit ?,?", pmsCommentRows, m.table)
	var resp []PmsComment
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

func (m *PmsCommentModel) Count() (int64, error) {
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

func (m *PmsCommentModel) Update(data PmsComment) error {
	query := fmt.Sprintf("update %s set %s where id = ?", m.table, pmsCommentRowsWithPlaceHolder)
	_, err := m.conn.Exec(query, data.ProductId, data.MemberNickName, data.ProductName, data.Star, data.MemberIp, data.ShowStatus, data.ProductAttribute, data.CollectCouont, data.ReadCount, data.Content, data.Pics, data.MemberIcon, data.ReplayCount, data.Id)
	return err
}

func (m *PmsCommentModel) Delete(id int64) error {
	query := fmt.Sprintf("delete from %s where id = ?", m.table)
	_, err := m.conn.Exec(query, id)
	return err
}
