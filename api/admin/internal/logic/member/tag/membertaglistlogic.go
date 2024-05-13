package logic

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MemberTagListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMemberTagListLogic(ctx context.Context, svcCtx *svc.ServiceContext) MemberTagListLogic {
	return MemberTagListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MemberTagListLogic) MemberTagList(req types.ListMemberTagReq) (*types.ListMemberTagResp, error) {
	resp, err := l.svcCtx.MemberTagService.MemberTagList(l.ctx, &umsclient.MemberTagListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
	})

	if err != nil {
		logc.Errorf(l.ctx, "参数: %+v,查询会员标签列表异常:%s", req, err.Error())
		return nil, errorx.NewDefaultError("查询会员标签失败")
	}

	var list []*types.ListMemberTagData

	for _, item := range resp.List {
		list = append(list, &types.ListMemberTagData{
			Id:                item.Id,
			Name:              item.Name,
			FinishOrderCount:  item.FinishOrderCount,
			FinishOrderAmount: item.FinishOrderAmount,
		})
	}

	return &types.ListMemberTagResp{
		Current:  req.Current,
		Data:     list,
		PageSize: req.PageSize,
		Success:  true,
		Total:    resp.Total,
		Code:     "000000",
		Message:  "查询会员标签成功",
	}, nil
}
