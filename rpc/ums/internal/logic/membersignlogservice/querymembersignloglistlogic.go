package membersignlogservicelogic

import (
	"context"
	"errors"
	"time"

	"github.com/feihua/zero-admin/pkg/time_util"
	"github.com/feihua/zero-admin/rpc/ums/gen/query"
	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
)

// QueryMemberSignLogListLogic 查询会员签到记录列表
/*
Author: LiuFeiHua
Date: 2025/05/22 14:12:14
*/
type QueryMemberSignLogListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryMemberSignLogListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryMemberSignLogListLogic {
	return &QueryMemberSignLogListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryMemberSignLogList 查询会员签到记录列表
func (l *QueryMemberSignLogListLogic) QueryMemberSignLogList(in *umsclient.QueryMemberSignLogListReq) (*umsclient.QueryMemberSignLogListResp, error) {
	memberSignLog := query.UmsMemberSignLog
	q := memberSignLog.WithContext(l.ctx)
	if in.MemberId != 0 {
		q = q.Where(memberSignLog.MemberID.Eq(in.MemberId))
	}
	if len(in.SignDate) > 0 {
		signDate, err := time.Parse("2006-01-02", in.SignDate)
		if err != nil {
			return nil, errors.New("日期格式错误")
		}

		q = q.Where(memberSignLog.SignDate.Eq(signDate))
	}

	result, count, err := q.FindByPage(int((in.PageNum-1)*in.PageSize), int(in.PageSize))

	if err != nil {
		logc.Errorf(l.ctx, "查询会员签到记录列表失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("查询会员签到记录列表失败")
	}

	var list []*umsclient.MemberSignLogListData

	for _, item := range result {
		list = append(list, &umsclient.MemberSignLogListData{
			Id:           item.ID,                              //
			MemberId:     item.MemberID,                        // 会员ID
			SignDate:     time_util.TimeToStr1(item.SignDate),  // 签到日期
			ContinueDays: item.ContinueDays,                    // 连续签到天数
			Points:       item.Points,                          // 获得积分
			CreateTime:   time_util.TimeToStr(item.CreateTime), //

		})
	}

	return &umsclient.QueryMemberSignLogListResp{
		Total: count,
		List:  list,
	}, nil
}
