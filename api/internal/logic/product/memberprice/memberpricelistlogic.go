package logic

import (
	"context"
	"fmt"
	"go-zero-admin/rpc/pms/pmsclient"

	"go-zero-admin/api/internal/svc"
	"go-zero-admin/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
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
	resp, err := l.svcCtx.Pms.MemberPriceList(l.ctx, &pmsclient.MemberPriceListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
	})

	if err != nil {
		return nil, err
	}

	for _, data := range resp.List {
		fmt.Println(data)
	}
	var list []*types.ListtMemberPriceData

	for _, item := range resp.List {
		list = append(list, &types.ListtMemberPriceData{
			Id:              item.Id,
			ProductId:       item.ProductId,
			MemberLevelId:   item.MemberLevelId,
			MemberPrice:     float64(item.MemberPrice),
			MemberLevelName: item.MemberLevelName,
		})
	}

	return &types.ListMemberPriceResp{
		Current:  req.Current,
		Data:     nil,
		PageSize: req.PageSize,
		Success:  true,
		Total:    resp.Total,
	}, nil
}
