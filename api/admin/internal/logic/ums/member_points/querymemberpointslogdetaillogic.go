package member_points

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

// QueryMemberPointsLogDetailLogic 查询会员积分记录详情
/*
Author: LiuFeiHua
Date: 2025/05/22 15:25:56
*/
type QueryMemberPointsLogDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryMemberPointsLogDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryMemberPointsLogDetailLogic {
	return &QueryMemberPointsLogDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// QueryMemberPointsLogDetail 查询会员积分记录详情
func (l *QueryMemberPointsLogDetailLogic) QueryMemberPointsLogDetail(req *types.QueryMemberPointsLogDetailReq) (resp *types.QueryMemberPointsLogDetailResp, err error) {

	detail, err := l.svcCtx.MemberPointsLogService.QueryMemberPointsLogDetail(l.ctx, &umsclient.QueryMemberPointsLogDetailReq{
		Id: req.Id,
	})

	if err != nil {
		logc.Errorf(l.ctx, "查询会员积分记录详情失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	data := types.QueryMemberPointsLogDetailData{
		Id:           detail.Id,           //
		MemberId:     detail.MemberId,     // 会员ID
		ChangeType:   detail.ChangeType,   // 变更类型：1-添加积分，2-减少积分
		ChangePoints: detail.ChangePoints, // 变更积分
		SourceType:   detail.SourceType,   // 来源类型：0-其他，1-订单，2-活动，3-签到，4-管理员修改
		Description:  detail.Description,  // 描述
		OperateMan:   detail.OperateMan,   // 操作人员
		OperateNote:  detail.OperateNote,  // 操作备注
		CreateTime:   detail.CreateTime,   // 创建时间

	}
	return &types.QueryMemberPointsLogDetailResp{
		Code:    "000000",
		Message: "查询会员积分记录成功",
		Data:    data,
	}, nil
}
