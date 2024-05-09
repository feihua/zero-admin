package brandservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/pms/gen/model"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"

	"github.com/feihua/zero-admin/rpc/pms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

// BrandAddLogic 商品品牌
/*
Author: LiuFeiHua
Date: 2024/5/8 10:38
*/
type BrandAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBrandAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BrandAddLogic {
	return &BrandAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// BrandAdd 添加商品品牌
func (l *BrandAddLogic) BrandAdd(in *pmsclient.BrandAddReq) (*pmsclient.BrandAddResp, error) {
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

	return &pmsclient.BrandAddResp{}, nil
}
