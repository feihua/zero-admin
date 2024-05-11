package brand

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/api/web/internal/svc"
	"github.com/feihua/zero-admin/api/web/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// BrandListLogic
/*
Author: LiuFeiHua
Date: 2024/4/7 17:29
*/
type BrandListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewBrandListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BrandListLogic {
	return &BrandListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// BrandList 品牌列表
func (l *BrandListLogic) BrandList() (*types.BrandListResp, error) {
	req := &pmsclient.BrandListReq{
		Current:       1,
		PageSize:      100,
		Name:          "",
		FactoryStatus: 2,
		ShowStatus:    1,
	}
	resp, err := l.svcCtx.BrandService.BrandList(l.ctx, req)

	if err != nil {
		logc.Errorf(l.ctx, "参数: %+v,查询商品品牌列表异常:%s", req, err.Error())
		return nil, errors.New("查询商品品牌失败")
	}

	var list []types.BrandListData

	for _, item := range resp.List {
		list = append(list, types.BrandListData{
			Id:   item.Id,
			Name: item.Name,
		})
	}

	return &types.BrandListResp{
		Code:    0,
		Message: "查询商品品牌列表",
		Data:    list,
	}, nil
}
