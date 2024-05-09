package logic

import (
	"context"
	"encoding/json"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"

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
		data, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("参数: %s,查询会员价格列表异常:%s", string(data), err.Error())
		return nil, errorx.NewDefaultError("查询会员价格失败")
	}

	var list []*types.ListtMemberPriceData

	for _, item := range resp.List {
		list = append(list, &types.ListtMemberPriceData{
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
