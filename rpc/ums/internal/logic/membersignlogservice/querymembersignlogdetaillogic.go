package membersignlogservicelogic

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

// QueryMemberSignLogDetailLogic 查询会员签到记录详情
/*
Author: LiuFeiHua
Date: 2025/05/22 14:12:14
*/
type QueryMemberSignLogDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryMemberSignLogDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryMemberSignLogDetailLogic {
	return &QueryMemberSignLogDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryMemberSignLogDetail 查询会员签到记录详情
func (l *QueryMemberSignLogDetailLogic) QueryMemberSignLogDetail(in *umsclient.QueryMemberSignLogDetailReq) (*umsclient.QueryMemberSignLogDetailResp, error) {
	item, err := query.UmsMemberSignLog.WithContext(l.ctx).Where(query.UmsMemberSignLog.ID.Eq(in.Id)).First()

	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		logc.Errorf(l.ctx, "会员签到记录不存在, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errors.New("会员签到记录不存在")
	case err != nil:
		logc.Errorf(l.ctx, "查询会员签到记录异常, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errors.New("查询会员签到记录异常")
	}

	data := &umsclient.QueryMemberSignLogDetailResp{
		Id:           item.ID,                              //
		MemberId:     item.MemberID,                        // 会员ID
		SignDate:     time_util.TimeToStr1(item.SignDate),  // 签到日期
		ContinueDays: item.ContinueDays,                    // 连续签到天数
		Points:       item.Points,                          // 获得积分
		CreateTime:   time_util.TimeToStr(item.CreateTime), //
	}

	return data, nil
}
