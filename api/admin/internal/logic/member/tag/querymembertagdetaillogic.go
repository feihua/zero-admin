package tag

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

// QueryMemberTagDetailLogic 查询用户标签表详情
/*
Author: 刘飞华
Date: 2025/02/05 10:34:53
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

// QueryMemberTagDetail 查询用户标签表详情
func (l *QueryMemberTagDetailLogic) QueryMemberTagDetail(req *types.QueryMemberTagDetailReq) (resp *types.QueryMemberTagDetailResp, err error) {

	detail, err := l.svcCtx.MemberTagService.QueryMemberTagDetail(l.ctx, &umsclient.QueryMemberTagDetailReq{
		Id: req.Id,
	})

	if err != nil {
		logc.Errorf(l.ctx, "查询用户标签表详情失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	data := types.QueryMemberTagDetailData{
		Id:                detail.Id,                //
		TagName:           detail.TagName,           // 标签名称
		FinishOrderCount:  detail.FinishOrderCount,  // 自动打标签完成订单数量
		Status:            detail.Status,            // 状态：0->禁用；1->启用
		FinishOrderAmount: detail.FinishOrderAmount, // 自动打标签完成订单金额
	}
	return &types.QueryMemberTagDetailResp{
		Code:    "000000",
		Message: "查询用户标签表成功",
		Data:    data,
	}, nil
}
