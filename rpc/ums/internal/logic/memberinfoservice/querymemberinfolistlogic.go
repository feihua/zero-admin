package memberinfoservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/pkg/time_util"
	"github.com/feihua/zero-admin/rpc/ums/gen/query"
	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
)

// QueryMemberInfoListLogic 查询会员信息列表
/*
Author: LiuFeiHua
Date: 2025/05/21 14:18:26
*/
type QueryMemberInfoListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryMemberInfoListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryMemberInfoListLogic {
	return &QueryMemberInfoListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryMemberInfoList 查询会员信息列表
func (l *QueryMemberInfoListLogic) QueryMemberInfoList(in *umsclient.QueryMemberInfoListReq) (*umsclient.QueryMemberInfoListResp, error) {
	memberInfo := query.UmsMemberInfo
	q := memberInfo.WithContext(l.ctx)
	if in.MemberId != 0 {
		q = q.Where(memberInfo.MemberID.Eq(in.MemberId))
	}
	if len(in.Nickname) > 0 {
		q = q.Where(memberInfo.Nickname.Like("%" + in.Nickname + "%"))
	}
	if len(in.Mobile) > 0 {
		q = q.Where(memberInfo.Mobile.Like("%" + in.Mobile + "%"))
	}
	if in.Source != 3 {
		q = q.Where(memberInfo.Source.Eq(in.Source))
	}

	if in.IsEnabled != 2 {
		q = q.Where(memberInfo.IsEnabled.Eq(in.IsEnabled))
	}

	result, count, err := q.FindByPage(int((in.PageNum-1)*in.PageSize), int(in.PageSize))

	if err != nil {
		logc.Errorf(l.ctx, "查询会员信息列表失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("查询会员信息列表失败")
	}

	var list []*umsclient.MemberInfoListData

	for _, item := range result {
		list = append(list, &umsclient.MemberInfoListData{
			Id:           item.ID,                                 // 主键ID
			MemberId:     item.MemberID,                           // 会员ID
			LevelId:      item.LevelID,                            // 等级ID
			Nickname:     item.Nickname,                           // 昵称
			Mobile:       item.Mobile,                             // 手机号码
			Source:       item.Source,                             // 注册来源：0-PC，1-APP，2-小程序
			Avatar:       item.Avatar,                             // 头像
			Signature:    item.Signature,                          // 个性签名
			Gender:       item.Gender,                             // 性别：0-未知，1-男，2-女
			Birthday:     time_util.TimeToString(item.Birthday),   // 生日
			GrowthPoint:  item.GrowthPoint,                        // 成长值
			Points:       item.Points,                             // 积分
			TotalPoints:  item.TotalPoints,                        // 累计获得积分
			SpendAmount:  float32(item.SpendAmount),               // 累计消费金额
			OrderCount:   item.OrderCount,                         // 订单数
			CouponCount:  item.CouponCount,                        // 优惠券数量
			CommentCount: item.CommentCount,                       // 评价数
			ReturnCount:  item.ReturnCount,                        // 退货数
			LotteryTimes: item.LotteryTimes,                       // 剩余抽奖次数
			LastLogin:    time_util.TimeToString(item.LastLogin),  // 最后登录
			IsEnabled:    item.IsEnabled,                          // 是否启用：0-禁用，1-启用
			CreateTime:   time_util.TimeToStr(item.CreateTime),    // 创建时间
			UpdateTime:   time_util.TimeToString(item.UpdateTime), // 更新时间

		})
	}

	return &umsclient.QueryMemberInfoListResp{
		Total: count,
		List:  list,
	}, nil
}
