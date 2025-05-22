package member_tag

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

// QueryMemberTagDetailLogic 查询用户标签详情
/*
Author: LiuFeiHua
Date: 2025/05/22 10:44:59
*/
type QueryMemberTagDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryMemberTagDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryMemberTagDetailLogic {
	return &QueryMemberTagDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// QueryMemberTagDetail 查询用户标签详情
func (l *QueryMemberTagDetailLogic) QueryMemberTagDetail(req *types.QueryMemberTagDetailReq) (resp *types.QueryMemberTagDetailResp, err error) {

	detail, err := l.svcCtx.MemberTagService.QueryMemberTagDetail(l.ctx, &umsclient.QueryMemberTagDetailReq{
		Id: req.Id,
	})

	if err != nil {
		logc.Errorf(l.ctx, "查询用户标签详情失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	data := types.QueryMemberTagDetailData{
		Id:                detail.Id,                         // 主键ID
		TagName:           detail.TagName,                    // 标签名称
		Description:       detail.Description,                // 标签描述
		FinishOrderCount:  detail.FinishOrderCount,           // 自动打标签完成订单数量
		FinishOrderAmount: float64(detail.FinishOrderAmount), // 自动打标签完成订单金额
		Status:            detail.Status,                     // 状态：0-禁用，1-启用
		CreateBy:          detail.CreateBy,                   // 创建人ID
		CreateTime:        detail.CreateTime,                 // 创建时间
		UpdateBy:          detail.UpdateBy,                   // 更新人ID
		UpdateTime:        detail.UpdateTime,                 // 更新时间

	}
	return &types.QueryMemberTagDetailResp{
		Code:    "000000",
		Message: "查询用户标签成功",
		Data:    data,
	}, nil
}
