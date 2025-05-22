package member_growth

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

// QueryMemberGrowthLogDetailLogic 查询会员成长值记录详情
/*
Author: LiuFeiHua
Date: 2025/05/22 15:25:56
*/
type QueryMemberGrowthLogDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryMemberGrowthLogDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryMemberGrowthLogDetailLogic {
	return &QueryMemberGrowthLogDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// QueryMemberGrowthLogDetail 查询会员成长值记录详情
func (l *QueryMemberGrowthLogDetailLogic) QueryMemberGrowthLogDetail(req *types.QueryMemberGrowthLogDetailReq) (resp *types.QueryMemberGrowthLogDetailResp, err error) {

	detail, err := l.svcCtx.MemberGrowthLogService.QueryMemberGrowthLogDetail(l.ctx, &umsclient.QueryMemberGrowthLogDetailReq{
		Id: req.Id,
	})

	if err != nil {
		logc.Errorf(l.ctx, "查询会员成长值记录详情失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	data := types.QueryMemberGrowthLogDetailData{
		Id:           detail.Id,           //
		MemberId:     detail.MemberId,     // 会员ID
		ChangeType:   detail.ChangeType,   // 变更类型：1-添加成长值，2-减少成长值
		ChangeGrowth: detail.ChangeGrowth, // 变更成长值
		SourceType:   detail.SourceType,   // 来源类型：0-其他，1-订单，2-活动，3-签到，4-管理员修改
		Description:  detail.Description,  // 描述
		OperateMan:   detail.OperateMan,   // 操作人员
		OperateNote:  detail.OperateNote,  // 操作备注
		CreateTime:   detail.CreateTime,   // 创建时间

	}
	return &types.QueryMemberGrowthLogDetailResp{
		Code:    "000000",
		Message: "查询会员成长值记录成功",
		Data:    data,
	}, nil
}
