package logic

import (
	"context"
	"go-zero-admin/rpc/model/pmsmodel"

	"go-zero-admin/rpc/pms/internal/svc"
	"go-zero-admin/rpc/pms/pms"

	"github.com/tal-tech/go-zero/core/logx"
)

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

func (l *BrandAddLogic) BrandAdd(in *pms.BrandAddReq) (*pms.BrandAddResp, error) {
	_, err := l.svcCtx.PmsBrandModel.Insert(pmsmodel.PmsBrand{
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

	return &pms.BrandAddResp{}, nil
}
