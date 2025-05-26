package productbrandservicelogic

import (
	"context"
	"errors"
	"fmt"
	"github.com/feihua/zero-admin/rpc/pms/gen/model"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"
	"github.com/feihua/zero-admin/rpc/pms/internal/svc"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
)

// AddProductBrandLogic 添加商品品牌
/*
Author: LiuFeiHua
Date: 2025/05/26 10:33:54
*/
type AddProductBrandLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddProductBrandLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddProductBrandLogic {
	return &AddProductBrandLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// AddProductBrand 添加商品品牌
func (l *AddProductBrandLogic) AddProductBrand(in *pmsclient.AddProductBrandReq) (*pmsclient.AddProductBrandResp, error) {
	q := query.PmsProductBrand

	count, _ := q.WithContext(l.ctx).Where(q.Name.Eq(in.Name)).Count()

	if count > 0 {
		return nil, errors.New(fmt.Sprintf("品牌名称：%s,已存在", in.Name))
	}
	item := &model.PmsProductBrand{
		Name:                in.Name,                // 品牌名称
		Logo:                in.Logo,                // 品牌logo
		BigPic:              in.BigPic,              // 专区大图
		Description:         in.Description,         // 描述
		FirstLetter:         in.FirstLetter,         // 首字母
		Sort:                in.Sort,                // 排序
		RecommendStatus:     in.RecommendStatus,     // 推荐状态
		ProductCount:        in.ProductCount,        // 产品数量
		ProductCommentCount: in.ProductCommentCount, // 产品评论数量
		IsEnabled:           in.IsEnabled,           // 是否启用
		CreateBy:            in.CreateBy,            // 创建人ID
	}

	err := q.WithContext(l.ctx).Create(item)
	if err != nil {
		logc.Errorf(l.ctx, "添加商品品牌失败,参数:%+v,异常:%s", item, err.Error())
		return nil, errors.New("添加商品品牌失败")
	}

	return &pmsclient.AddProductBrandResp{
		BrandId: item.ID,
	}, nil
}
