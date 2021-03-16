package logic

import (
	"context"
	"fmt"
	"go-zero-admin/rpc/sms/smsclient"

	"go-zero-admin/api/internal/svc"
	"go-zero-admin/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type HomeBrandListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHomeBrandListLogic(ctx context.Context, svcCtx *svc.ServiceContext) HomeBrandListLogic {
	return HomeBrandListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HomeBrandListLogic) HomeBrandList(req types.ListHomeBrandReq) (*types.ListHomeBrandResp, error) {
	resp, err := l.svcCtx.Sms.HomeBrandList(l.ctx, &smsclient.HomeBrandListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
	})

	if err != nil {
		return nil, err
	}

	for _, data := range resp.List {

		fmt.Println(data)
	}
	var list []*types.ListtHomeBrandData

	for _, item := range resp.List {
		list = append(list, &types.ListtHomeBrandData{
			Id:              item.Id,
			BrandId:         item.BrandId,
			BrandName:       item.BrandName,
			RecommendStatus: item.RecommendStatus,
			Sort:            item.Sort,
		})
	}

	return &types.ListHomeBrandResp{
		Current:  req.Current,
		Data:     list,
		PageSize: req.PageSize,
		Success:  true,
		Total:    resp.Total,
		Code:     "000000",
		Message:  "",
	}, nil
}
