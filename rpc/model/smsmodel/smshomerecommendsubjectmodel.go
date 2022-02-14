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
	smsHomeRecommendSubjectFieldNames          = builderx.FieldNames(&SmsHomeRecommendSubject{})
	smsHomeRecommendSubjectRows                = strings.Join(smsHomeRecommendSubjectFieldNames, ",")
	smsHomeRecommendSubjectRowsExpectAutoSet   = strings.Join(stringx.Remove(smsHomeRecommendSubjectFieldNames, "id", "create_time", "update_time"), ",")
	smsHomeRecommendSubjectRowsWithPlaceHolder = strings.Join(stringx.Remove(smsHomeRecommendSubjectFieldNames, "id", "create_time", "update_time"), "=?,") + "=?"
)

type (
	SmsHomeRecommendSubjectModel struct {
		conn  sqlx.SqlConn
		table string
	}

	SmsHomeRecommendSubject struct {
		Id              int64  `db:"id"`
		SubjectId       int64  `db:"subject_id"`
		SubjectName     string `db:"subject_name"`
		RecommendStatus int64  `db:"recommend_status"`
		Sort            int64  `db:"sort"`
	}
)

func NewSmsHomeRecommendSubjectModel(conn sqlx.SqlConn) *SmsHomeRecommendSubjectModel {
	return &SmsHomeRecommendSubjectModel{
		conn:  conn,
		table: "sms_home_recommend_subject",
	}
}

func (m *SmsHomeRecommendSubjectModel) Insert(data SmsHomeRecommendSubject) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?)", m.table, smsHomeRecommendSubjectRowsExpectAutoSet)
	ret, err := m.conn.Exec(query, data.SubjectId, data.SubjectName, data.RecommendStatus, data.Sort)
	return ret, err
}

func (m *SmsHomeRecommendSubjectModel) FindOne(id int64) (*SmsHomeRecommendSubject, error) {
	query := fmt.Sprintf("select %s from %s where id = ? limit 1", smsHomeRecommendSubjectRows, m.table)
	var resp SmsHomeRecommendSubject
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

func (m *SmsHomeRecommendSubjectModel) FindAll(Current int64, PageSize int64) (*[]SmsHomeRecommendSubject, error) {

	query := fmt.Sprintf("select %s from %s limit ?,?", smsHomeRecommendSubjectRows, m.table)
	var resp []SmsHomeRecommendSubject
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

func (m *SmsHomeRecommendSubjectModel) Count() (int64, error) {
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

func (m *SmsHomeRecommendSubjectModel) Update(data SmsHomeRecommendSubject) error {
	query := fmt.Sprintf("update %s set %s where id = ?", m.table, smsHomeRecommendSubjectRowsWithPlaceHolder)
	_, err := m.conn.Exec(query, data.SubjectId, data.SubjectName, data.RecommendStatus, data.Sort, data.Id)
	return err
}

func (m *SmsHomeRecommendSubjectModel) Delete(id int64) error {
	query := fmt.Sprintf("delete from %s where id = ?", m.table)
	_, err := m.conn.Exec(query, id)
	return err
}
