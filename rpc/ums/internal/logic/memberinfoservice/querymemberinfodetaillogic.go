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
	"gorm.io/gorm"
)

// QueryMemberInfoDetailLogic 查询会员信息详情
/*
Author: LiuFeiHua
Date: 2025/05/21 14:18:26
*/
type QueryMemberInfoDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryMemberInfoDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryMemberInfoDetailLogic {
	return &QueryMemberInfoDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryMemberInfoDetail 查询会员信息详情
func (l *QueryMemberInfoDetailLogic) QueryMemberInfoDetail(in *umsclient.QueryMemberInfoDetailReq) (*umsclient.QueryMemberInfoDetailResp, error) {
	item, err := query.UmsMemberInfo.WithContext(l.ctx).Where(query.UmsMemberInfo.MemberID.Eq(in.MemberId)).First()

	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		logc.Errorf(l.ctx, "会员不存在, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errors.New("会员不存在")
	case err != nil:
		logc.Errorf(l.ctx, "查询会员异常, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errors.New("查询会员异常")
	}

	data := &umsclient.QueryMemberInfoDetailResp{
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
	}

	return data, nil
}
