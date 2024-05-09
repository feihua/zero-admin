package brandservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/pms/gen/model"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"

	"github.com/feihua/zero-admin/rpc/pms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

// BrandUpdateLogic 商品品牌
/*
Author: LiuFeiHua
Date: 2024/5/8 10:38
*/
type BrandUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBrandUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BrandUpdateLogic {
	return &BrandUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// BrandUpdate 更新商品品牌
func (l *BrandUpdateLogic) BrandUpdate(in *pmsclient.BrandUpdateReq) (*pmsclient.BrandUpdateResp, error) {
	q := query.PmsBrand
	_, err := q.WithContext(l.ctx).Updates(&model.PmsBrand{
		ID:                  in.Id,
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

	return &pmsclient.BrandUpdateResp{}, nil
}
