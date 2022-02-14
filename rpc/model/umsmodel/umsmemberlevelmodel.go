package umsmodel

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
	umsMemberLevelFieldNames          = builderx.FieldNames(&UmsMemberLevel{})
	umsMemberLevelRows                = strings.Join(umsMemberLevelFieldNames, ",")
	umsMemberLevelRowsExpectAutoSet   = strings.Join(stringx.Remove(umsMemberLevelFieldNames, "id", "create_time", "update_time"), ",")
	umsMemberLevelRowsWithPlaceHolder = strings.Join(stringx.Remove(umsMemberLevelFieldNames, "id", "create_time", "update_time"), "=?,") + "=?"
)

type (
	UmsMemberLevelModel struct {
		conn  sqlx.SqlConn
		table string
	}

	UmsMemberLevel struct {
		Id                    int64   `db:"id"`
		Name                  string  `db:"name"`
		GrowthPoint           int64   `db:"growth_point"`
		DefaultStatus         int64   `db:"default_status"`          // 是否为默认等级：0->不是；1->是
		FreeFreightPoint      float64 `db:"free_freight_point"`      // 免运费标准
		CommentGrowthPoint    int64   `db:"comment_growth_point"`    // 每次评价获取的成长值
		PriviledgeFreeFreight int64   `db:"priviledge_free_freight"` // 是否有免邮特权
		PriviledgeSignIn      int64   `db:"priviledge_sign_in"`      // 是否有签到特权
		PriviledgeComment     int64   `db:"priviledge_comment"`      // 是否有评论获奖励特权
		PriviledgePromotion   int64   `db:"priviledge_promotion"`    // 是否有专享活动特权
		PriviledgeMemberPrice int64   `db:"priviledge_member_price"` // 是否有会员价格特权
		PriviledgeBirthday    int64   `db:"priviledge_birthday"`     // 是否有生日特权
		Note                  string  `db:"note"`
	}
)

func NewUmsMemberLevelModel(conn sqlx.SqlConn) *UmsMemberLevelModel {
	return &UmsMemberLevelModel{
		conn:  conn,
		table: "ums_member_level",
	}
}

func (m *UmsMemberLevelModel) Insert(data UmsMemberLevel) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, umsMemberLevelRowsExpectAutoSet)
	ret, err := m.conn.Exec(query, data.Name, data.GrowthPoint, data.DefaultStatus, data.FreeFreightPoint, data.CommentGrowthPoint, data.PriviledgeFreeFreight, data.PriviledgeSignIn, data.PriviledgeComment, data.PriviledgePromotion, data.PriviledgeMemberPrice, data.PriviledgeBirthday, data.Note)
	return ret, err
}

func (m *UmsMemberLevelModel) FindOne(id int64) (*UmsMemberLevel, error) {
	query := fmt.Sprintf("select %s from %s where id = ? limit 1", umsMemberLevelRows, m.table)
	var resp UmsMemberLevel
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

func (m *UmsMemberLevelModel) FindAll(Current int64, PageSize int64) (*[]UmsMemberLevel, error) {

	query := fmt.Sprintf("select %s from %s limit ?,?", umsMemberLevelRows, m.table)
	var resp []UmsMemberLevel
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

func (m *UmsMemberLevelModel) Count() (int64, error) {
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

func (m *UmsMemberLevelModel) Update(data UmsMemberLevel) error {
	query := fmt.Sprintf("update %s set %s where id = ?", m.table, umsMemberLevelRowsWithPlaceHolder)
	_, err := m.conn.Exec(query, data.Name, data.GrowthPoint, data.DefaultStatus, data.FreeFreightPoint, data.CommentGrowthPoint, data.PriviledgeFreeFreight, data.PriviledgeSignIn, data.PriviledgeComment, data.PriviledgePromotion, data.PriviledgeMemberPrice, data.PriviledgeBirthday, data.Note, data.Id)
	return err
}

func (m *UmsMemberLevelModel) Delete(id int64) error {
	query := fmt.Sprintf("delete from %s where id = ?", m.table)
	_, err := m.conn.Exec(query, id)
	return err
}
