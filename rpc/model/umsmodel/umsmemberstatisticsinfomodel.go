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
	umsMemberStatisticsInfoFieldNames          = builderx.FieldNames(&UmsMemberStatisticsInfo{})
	umsMemberStatisticsInfoRows                = strings.Join(umsMemberStatisticsInfoFieldNames, ",")
	umsMemberStatisticsInfoRowsExpectAutoSet   = strings.Join(stringx.Remove(umsMemberStatisticsInfoFieldNames, "id", "create_time", "update_time"), ",")
	umsMemberStatisticsInfoRowsWithPlaceHolder = strings.Join(stringx.Remove(umsMemberStatisticsInfoFieldNames, "id", "create_time", "update_time"), "=?,") + "=?"
)

type (
	UmsMemberStatisticsInfoModel struct {
		conn  sqlx.SqlConn
		table string
	}

	UmsMemberStatisticsInfo struct {
		Id                  int64     `db:"id"`
		MemberId            int64     `db:"member_id"`
		ConsumeAmount       float64   `db:"consume_amount"`     // 累计消费金额
		OrderCount          int64     `db:"order_count"`        // 订单数量
		CouponCount         int64     `db:"coupon_count"`       // 优惠券数量
		CommentCount        int64     `db:"comment_count"`      // 评价数
		ReturnOrderCount    int64     `db:"return_order_count"` // 退货数量
		LoginCount          int64     `db:"login_count"`        // 登录次数
		AttendCount         int64     `db:"attend_count"`       // 关注数量
		FansCount           int64     `db:"fans_count"`         // 粉丝数量
		CollectProductCount int64     `db:"collect_product_count"`
		CollectSubjectCount int64     `db:"collect_subject_count"`
		CollectTopicCount   int64     `db:"collect_topic_count"`
		CollectCommentCount int64     `db:"collect_comment_count"`
		InviteFriendCount   int64     `db:"invite_friend_count"`
		RecentOrderTime     time.Time `db:"recent_order_time"` // 最后一次下订单时间
	}
)

func NewUmsMemberStatisticsInfoModel(conn sqlx.SqlConn) *UmsMemberStatisticsInfoModel {
	return &UmsMemberStatisticsInfoModel{
		conn:  conn,
		table: "ums_member_statistics_info",
	}
}

func (m *UmsMemberStatisticsInfoModel) Insert(data UmsMemberStatisticsInfo) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, umsMemberStatisticsInfoRowsExpectAutoSet)
	ret, err := m.conn.Exec(query, data.MemberId, data.ConsumeAmount, data.OrderCount, data.CouponCount, data.CommentCount, data.ReturnOrderCount, data.LoginCount, data.AttendCount, data.FansCount, data.CollectProductCount, data.CollectSubjectCount, data.CollectTopicCount, data.CollectCommentCount, data.InviteFriendCount, data.RecentOrderTime)
	return ret, err
}

func (m *UmsMemberStatisticsInfoModel) FindOne(id int64) (*UmsMemberStatisticsInfo, error) {
	query := fmt.Sprintf("select %s from %s where id = ? limit 1", umsMemberStatisticsInfoRows, m.table)
	var resp UmsMemberStatisticsInfo
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

func (m *UmsMemberStatisticsInfoModel) FindAll(Current int64, PageSize int64) (*[]UmsMemberStatisticsInfo, error) {

	query := fmt.Sprintf("select %s from %s limit ?,?", umsMemberStatisticsInfoRows, m.table)
	var resp []UmsMemberStatisticsInfo
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

func (m *UmsMemberStatisticsInfoModel) Count() (int64, error) {
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

func (m *UmsMemberStatisticsInfoModel) Update(data UmsMemberStatisticsInfo) error {
	query := fmt.Sprintf("update %s set %s where id = ?", m.table, umsMemberStatisticsInfoRowsWithPlaceHolder)
	_, err := m.conn.Exec(query, data.MemberId, data.ConsumeAmount, data.OrderCount, data.CouponCount, data.CommentCount, data.ReturnOrderCount, data.LoginCount, data.AttendCount, data.FansCount, data.CollectProductCount, data.CollectSubjectCount, data.CollectTopicCount, data.CollectCommentCount, data.InviteFriendCount, data.RecentOrderTime, data.Id)
	return err
}

func (m *UmsMemberStatisticsInfoModel) Delete(id int64) error {
	query := fmt.Sprintf("delete from %s where id = ?", m.table)
	_, err := m.conn.Exec(query, id)
	return err
}
