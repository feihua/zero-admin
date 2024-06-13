package brandservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/pms/gen/model"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"

	"github.com/feihua/zero-admin/rpc/pms/internal/svc"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// AddBrandLogic 添加品牌表
/*
Author: LiuFeiHua
Date: 2024/6/12 16:25
*/
type AddBrandLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddBrandLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddBrandLogic {
	return &AddBrandLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// AddBrand 添加品牌表
func (l *AddBrandLogic) AddBrand(in *pmsclient.AddBrandReq) (*pmsclient.AddBrandResp, error) {
	err := query.PmsBrand.WithContext(l.ctx).Create(&model.PmsBrand{
		Name:                in.Name,
		FirstLetter:         in.FirstLetter,
		Sort:                in.Sort,
		FactoryStatus:       in.FactoryStatus,
		ShowStatus:          in.ShowStatus,
		ProductCount:        in.ProductCount,
		ProductCommentCount: in.ProductCommentCount,
		Logo:                in.Logo,
		BigPic:              in.BigPic,
		BrandStory:          in.BrandStory,
	})

	if err != nil {
		return nil, err
	}

	return &pmsclient.AddBrandResp{}, nil
}
