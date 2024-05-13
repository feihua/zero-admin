package logic

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MemberPriceListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMemberPriceListLogic(ctx context.Context, svcCtx *svc.ServiceContext) MemberPriceListLogic {
	return MemberPriceListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MemberPriceListLogic) MemberPriceList(req types.ListMemberPriceReq) (*types.ListMemberPriceResp, error) {
	resp, err := l.svcCtx.MemberPriceService.MemberPriceList(l.ctx, &pmsclient.MemberPriceListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
	})

	if err != nil {
		logc.Errorf(l.ctx, "参数: %+v,查询会员价格列表异常:%s", req, err.Error())
		return nil, errorx.NewDefaultError("查询会员价格失败")
	}

	var list []*types.ListMemberPriceData

	for _, item := range resp.List {
		list = append(list, &types.ListMemberPriceData{
			Id:              item.Id,
			ProductId:       item.ProductId,
			MemberLevelId:   item.MemberLevelId,
			MemberPrice:     item.MemberPrice,
			MemberLevelName: item.MemberLevelName,
		})
	}

	return &types.ListMemberPriceResp{
		Current:  req.Current,
		Data:     list,
		PageSize: req.PageSize,
		Success:  true,
		Total:    resp.Total,
		Code:     "000000",
		Message:  "查询会员价格成功",
	}, nil
}
