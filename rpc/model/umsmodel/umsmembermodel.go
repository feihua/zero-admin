package umsmodel

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
	umsMemberFieldNames          = builderx.FieldNames(&UmsMember{})
	umsMemberRows                = strings.Join(umsMemberFieldNames, ",")
	umsMemberRowsExpectAutoSet   = strings.Join(stringx.Remove(umsMemberFieldNames, "id", "create_time", "update_time"), ",")
	umsMemberRowsWithPlaceHolder = strings.Join(stringx.Remove(umsMemberFieldNames, "id", "create_time", "update_time"), "=?,") + "=?"
)

type (
	UmsMemberModel struct {
		conn  sqlx.SqlConn
		table string
	}

	UmsMember struct {
		Id                    int64     `db:"id"`
		MemberLevelId         int64     `db:"member_level_id"`
		Username              string    `db:"username"`               // 用户名
		Password              string    `db:"password"`               // 密码
		Nickname              string    `db:"nickname"`               // 昵称
		Phone                 string    `db:"phone"`                  // 手机号码
		Status                int64     `db:"status"`                 // 帐号启用状态:0->禁用；1->启用
		CreateTime            time.Time `db:"create_time"`            // 注册时间
		Icon                  string    `db:"icon"`                   // 头像
		Gender                int64     `db:"gender"`                 // 性别：0->未知；1->男；2->女
		Birthday              time.Time `db:"birthday"`               // 生日
		City                  string    `db:"city"`                   // 所做城市
		Job                   string    `db:"job"`                    // 职业
		PersonalizedSignature string    `db:"personalized_signature"` // 个性签名
		SourceType            int64     `db:"source_type"`            // 用户来源
		Integration           int64     `db:"integration"`            // 积分
		Growth                int64     `db:"growth"`                 // 成长值
		LuckeyCount           int64     `db:"luckey_count"`           // 剩余抽奖次数
		HistoryIntegration    int64     `db:"history_integration"`    // 历史积分数量
	}
)

func NewUmsMemberModel(conn sqlx.SqlConn) *UmsMemberModel {
	return &UmsMemberModel{
		conn:  conn,
		table: "ums_member",
	}
}

func (m *UmsMemberModel) Insert(data UmsMember) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, umsMemberRowsExpectAutoSet)
	ret, err := m.conn.Exec(query, data.MemberLevelId, data.Username, data.Password, data.Nickname, data.Phone, data.Status, data.Icon, data.Gender, data.Birthday, data.City, data.Job, data.PersonalizedSignature, data.SourceType, data.Integration, data.Growth, data.LuckeyCount, data.HistoryIntegration)
	return ret, err
}

func (m *UmsMemberModel) FindOne(id int64) (*UmsMember, error) {
	query := fmt.Sprintf("select %s from %s where id = ? limit 1", umsMemberRows, m.table)
	var resp UmsMember
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

func (m *UmsMemberModel) FindAll(Current int64, PageSize int64) (*[]UmsMember, error) {

	query := fmt.Sprintf("select %s from %s limit ?,?", umsMemberRows, m.table)
	var resp []UmsMember
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

func (m *UmsMemberModel) Count() (int64, error) {
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

func (m *UmsMemberModel) FindOneByUsername(username string) (*UmsMember, error) {
	var resp UmsMember
	query := fmt.Sprintf("select %s from %s where username = ? limit 1", umsMemberRows, m.table)
	err := m.conn.QueryRow(&resp, query, username)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *UmsMemberModel) FindOneByPhone(phone string) (*UmsMember, error) {
	var resp UmsMember
	query := fmt.Sprintf("select %s from %s where phone = ? limit 1", umsMemberRows, m.table)
	err := m.conn.QueryRow(&resp, query, phone)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *UmsMemberModel) Update(data UmsMember) error {
	query := fmt.Sprintf("update %s set %s where id = ?", m.table, umsMemberRowsWithPlaceHolder)
	_, err := m.conn.Exec(query, data.MemberLevelId, data.Username, data.Password, data.Nickname, data.Phone, data.Status, data.Icon, data.Gender, data.Birthday, data.City, data.Job, data.PersonalizedSignature, data.SourceType, data.Integration, data.Growth, data.LuckeyCount, data.HistoryIntegration, data.Id)
	return err
}

func (m *UmsMemberModel) Delete(id int64) error {
	query := fmt.Sprintf("delete from %s where id = ?", m.table)
	_, err := m.conn.Exec(query, id)
	return err
}
