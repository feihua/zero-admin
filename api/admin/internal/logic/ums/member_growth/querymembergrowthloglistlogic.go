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

// QueryMemberGrowthLogListLogic 查询会员成长值记录列表
/*
Author: LiuFeiHua
Date: 2025/05/22 15:25:56
*/
type QueryMemberGrowthLogListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryMemberGrowthLogListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryMemberGrowthLogListLogic {
	return &QueryMemberGrowthLogListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// QueryMemberGrowthLogList 查询会员成长值记录列表
func (l *QueryMemberGrowthLogListLogic) QueryMemberGrowthLogList(req *types.QueryMemberGrowthLogListReq) (resp *types.QueryMemberGrowthLogListResp, err error) {
	result, err := l.svcCtx.MemberGrowthLogService.QueryMemberGrowthLogList(l.ctx, &umsclient.QueryMemberGrowthLogListReq{
		PageNum:    req.Current,
		PageSize:   req.PageSize,
		MemberId:   req.MemberId,   // 会员ID
		ChangeType: req.ChangeType, // 变更类型：1-添加成长值，2-减少成长值
		SourceType: req.SourceType, // 来源类型：0-其他，1-订单，2-活动，3-签到，4-管理员修改
	})

	if err != nil {
		logc.Errorf(l.ctx, "查询字会员成长值记录列表失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	var list []*types.QueryMemberGrowthLogListData

	for _, detail := range result.List {
		list = append(list, &types.QueryMemberGrowthLogListData{
			Id:           detail.Id,           //
			MemberId:     detail.MemberId,     // 会员ID
			ChangeType:   detail.ChangeType,   // 变更类型：1-添加成长值，2-减少成长值
			ChangeGrowth: detail.ChangeGrowth, // 变更成长值
			SourceType:   detail.SourceType,   // 来源类型：0-其他，1-订单，2-活动，3-签到，4-管理员修改
			Description:  detail.Description,  // 描述
			OperateMan:   detail.OperateMan,   // 操作人员
			OperateNote:  detail.OperateNote,  // 操作备注
			CreateTime:   detail.CreateTime,   // 创建时间

		})
	}

	return &types.QueryMemberGrowthLogListResp{
		Code:     "000000",
		Message:  "查询会员成长值记录列表成功",
		Current:  req.Current,
		Data:     list,
		PageSize: req.PageSize,
		Success:  true,
		Total:    result.Total,
	}, nil
}
