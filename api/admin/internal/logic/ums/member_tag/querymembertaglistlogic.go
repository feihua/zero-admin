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

// QueryMemberTagListLogic 查询用户标签列表
/*
Author: LiuFeiHua
Date: 2025/05/22 10:44:59
*/
type QueryMemberTagListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryMemberTagListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryMemberTagListLogic {
	return &QueryMemberTagListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// QueryMemberTagList 查询用户标签列表
func (l *QueryMemberTagListLogic) QueryMemberTagList(req *types.QueryMemberTagListReq) (resp *types.QueryMemberTagListResp, err error) {
	result, err := l.svcCtx.MemberTagService.QueryMemberTagList(l.ctx, &umsclient.QueryMemberTagListReq{
		PageNum:  req.Current,
		PageSize: req.PageSize,
		TagName:  req.TagName, // 标签名称
		Status:   req.Status,  // 状态：0-禁用，1-启用
	})

	if err != nil {
		logc.Errorf(l.ctx, "查询字用户标签列表失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	var list []*types.QueryMemberTagListData

	for _, detail := range result.List {
		list = append(list, &types.QueryMemberTagListData{
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

		})
	}

	return &types.QueryMemberTagListResp{
		Code:     "000000",
		Message:  "查询用户标签列表成功",
		Current:  req.Current,
		Data:     list,
		PageSize: req.PageSize,
		Success:  true,
		Total:    result.Total,
	}, nil
}
