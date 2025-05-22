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

// QueryMemberSignLogListLogic 查询会员签到记录列表
/*
Author: LiuFeiHua
Date: 2025/05/22 14:12:14
*/
type QueryMemberSignLogListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryMemberSignLogListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryMemberSignLogListLogic {
	return &QueryMemberSignLogListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// QueryMemberSignLogList 查询会员签到记录列表
func (l *QueryMemberSignLogListLogic) QueryMemberSignLogList(req *types.QueryMemberSignLogListReq) (resp *types.QueryMemberSignLogListResp, err error) {
	result, err := l.svcCtx.MemberSignLogService.QueryMemberSignLogList(l.ctx, &umsclient.QueryMemberSignLogListReq{
		PageNum:  req.Current,
		PageSize: req.PageSize,
		MemberId: req.MemberId, // 会员ID
		SignDate: req.SignDate, // 签到日期
	})

	if err != nil {
		logc.Errorf(l.ctx, "查询字会员签到记录列表失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	var list []*types.QueryMemberSignLogListData

	for _, detail := range result.List {
		list = append(list, &types.QueryMemberSignLogListData{
			Id:           detail.Id,           //
			MemberId:     detail.MemberId,     // 会员ID
			SignDate:     detail.SignDate,     // 签到日期
			ContinueDays: detail.ContinueDays, // 连续签到天数
			Points:       detail.Points,       // 获得积分
			CreateTime:   detail.CreateTime,   //

		})
	}

	return &types.QueryMemberSignLogListResp{
		Code:     "000000",
		Message:  "查询会员签到记录列表成功",
		Current:  req.Current,
		Data:     list,
		PageSize: req.PageSize,
		Success:  true,
		Total:    result.Total,
	}, nil
}
