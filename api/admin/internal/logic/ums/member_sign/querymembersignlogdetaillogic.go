package member_sign

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryMemberSignLogDetailLogic 查询会员签到记录详情
/*
Author: LiuFeiHua
Date: 2025/05/22 14:12:14
*/
type QueryMemberSignLogDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryMemberSignLogDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryMemberSignLogDetailLogic {
	return &QueryMemberSignLogDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// QueryMemberSignLogDetail 查询会员签到记录详情
func (l *QueryMemberSignLogDetailLogic) QueryMemberSignLogDetail(req *types.QueryMemberSignLogDetailReq) (resp *types.QueryMemberSignLogDetailResp, err error) {

	detail, err := l.svcCtx.MemberSignLogService.QueryMemberSignLogDetail(l.ctx, &umsclient.QueryMemberSignLogDetailReq{
		Id: req.Id,
	})

	if err != nil {
		logc.Errorf(l.ctx, "查询会员签到记录详情失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	data := types.QueryMemberSignLogDetailData{
		Id:           detail.Id,           //
		MemberId:     detail.MemberId,     // 会员ID
		SignDate:     detail.SignDate,     // 签到日期
		ContinueDays: detail.ContinueDays, // 连续签到天数
		Points:       detail.Points,       // 获得积分
		CreateTime:   detail.CreateTime,   //

	}
	return &types.QueryMemberSignLogDetailResp{
		Code:    "000000",
		Message: "查询会员签到记录成功",
		Data:    data,
	}, nil
}
