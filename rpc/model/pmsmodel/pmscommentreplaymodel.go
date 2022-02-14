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
	pmsCommentReplayFieldNames          = builderx.FieldNames(&PmsCommentReplay{})
	pmsCommentReplayRows                = strings.Join(pmsCommentReplayFieldNames, ",")
	pmsCommentReplayRowsExpectAutoSet   = strings.Join(stringx.Remove(pmsCommentReplayFieldNames, "id", "create_time", "update_time"), ",")
	pmsCommentReplayRowsWithPlaceHolder = strings.Join(stringx.Remove(pmsCommentReplayFieldNames, "id", "create_time", "update_time"), "=?,") + "=?"
)

type (
	PmsCommentReplayModel struct {
		conn  sqlx.SqlConn
		table string
	}

	PmsCommentReplay struct {
		Id             int64     `db:"id"`
		CommentId      int64     `db:"comment_id"`
		MemberNickName string    `db:"member_nick_name"`
		MemberIcon     string    `db:"member_icon"`
		Content        string    `db:"content"`
		CreateTime     time.Time `db:"create_time"`
		Type           int64     `db:"type"` // 评论人员类型；0->会员；1->管理员
	}
)

func NewPmsCommentReplayModel(conn sqlx.SqlConn) *PmsCommentReplayModel {
	return &PmsCommentReplayModel{
		conn:  conn,
		table: "pms_comment_replay",
	}
}

func (m *PmsCommentReplayModel) Insert(data PmsCommentReplay) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?)", m.table, pmsCommentReplayRowsExpectAutoSet)
	ret, err := m.conn.Exec(query, data.CommentId, data.MemberNickName, data.MemberIcon, data.Content, data.Type)
	return ret, err
}

func (m *PmsCommentReplayModel) FindOne(id int64) (*PmsCommentReplay, error) {
	query := fmt.Sprintf("select %s from %s where id = ? limit 1", pmsCommentReplayRows, m.table)
	var resp PmsCommentReplay
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

func (m *PmsCommentReplayModel) FindAll(Current int64, PageSize int64) (*[]PmsCommentReplay, error) {

	query := fmt.Sprintf("select %s from %s limit ?,?", pmsCommentReplayRows, m.table)
	var resp []PmsCommentReplay
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

func (m *PmsCommentReplayModel) Count() (int64, error) {
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

func (m *PmsCommentReplayModel) Update(data PmsCommentReplay) error {
	query := fmt.Sprintf("update %s set %s where id = ?", m.table, pmsCommentReplayRowsWithPlaceHolder)
	_, err := m.conn.Exec(query, data.CommentId, data.MemberNickName, data.MemberIcon, data.Content, data.Type, data.Id)
	return err
}

func (m *PmsCommentReplayModel) Delete(id int64) error {
	query := fmt.Sprintf("delete from %s where id = ?", m.table)
	_, err := m.conn.Exec(query, id)
	return err
}
