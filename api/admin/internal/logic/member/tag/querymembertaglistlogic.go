package tag

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryMemberTagListLogic 查询会员标签列表
/*
Author: LiuFeiHua
Date: 2024/5/23 9:21
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

// QueryMemberTagList 查询会员标签列表
func (l *QueryMemberTagListLogic) QueryMemberTagList(req *types.QueryMemberTagListReq) (resp *types.QueryMemberTagListResp, err error) {
	result, err := l.svcCtx.MemberTagService.QueryMemberTagList(l.ctx, &umsclient.QueryMemberTagListReq{
		PageNum:  req.Current,
		PageSize: req.PageSize,
		Status:   req.Status,
		TagName:  req.TagName,
	})

	if err != nil {
		logc.Errorf(l.ctx, "参数: %+v,查询会员标签列表异常:%s", req, err.Error())
		return nil, errorx.NewDefaultError("查询会员标签失败")
	}

	var list []*types.QueryMemberTagListData

	for _, item := range result.List {
		list = append(list, &types.QueryMemberTagListData{
			Id:                item.Id,
			TagName:           item.TagName,
			FinishOrderCount:  item.FinishOrderCount,
			FinishOrderAmount: item.FinishOrderAmount,
			Status:            item.Status,
		})
	}

	return &types.QueryMemberTagListResp{
		Current:  req.Current,
		Data:     list,
		PageSize: req.PageSize,
		Success:  true,
		Total:    result.Total,
		Code:     "000000",
		Message:  "查询会员标签成功",
	}, nil
}
