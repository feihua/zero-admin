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
	"gorm.io/gorm"
	"time"
)

// UpdateProductBrandLogic 更新商品品牌
/*
Author: LiuFeiHua
Date: 2025/05/26 10:33:54
*/
type UpdateProductBrandLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateProductBrandLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateProductBrandLogic {
	return &UpdateProductBrandLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateProductBrand 更新商品品牌
func (l *UpdateProductBrandLogic) UpdateProductBrand(in *pmsclient.UpdateProductBrandReq) (*pmsclient.UpdateProductBrandResp, error) {
	brand := query.PmsProductBrand
	q := brand.WithContext(l.ctx)

	// 1.根据商品品牌id查询商品品牌是否已存在
	detail, err := q.Where(brand.ID.Eq(in.Id)).First()

	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		logc.Errorf(l.ctx, "商品品牌不存在, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errors.New("商品品牌不存在")
	case err != nil:
		logc.Errorf(l.ctx, "查询商品品牌异常, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errors.New("查询商品品牌异常")
	}

	count, _ := q.WithContext(l.ctx).Where(brand.ID.Neq(in.Id), brand.Name.Eq(in.Name)).Count()

	if count > 0 {
		return nil, errors.New(fmt.Sprintf("品牌名称：%s,已存在", in.Name))
	}
	now := time.Now()
	item := &model.PmsProductBrand{
		ID:                  in.Id,                  //
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
		CreateBy:            detail.CreateBy,        // 创建者
		CreateTime:          detail.CreateTime,      // 创建时间
		UpdateBy:            &in.UpdateBy,           // 更新者
		UpdateTime:          &now,                   // 更新时间
	}

	// 2.商品品牌存在时,则直接更新商品品牌
	err = l.svcCtx.DB.Model(&model.PmsProductBrand{}).WithContext(l.ctx).Where(brand.ID.Eq(in.Id)).Save(item).Error

	if err != nil {
		logc.Errorf(l.ctx, "更新商品品牌失败,参数:%+v,异常:%s", item, err.Error())
		return nil, errors.New("更新商品品牌失败")
	}

	return &pmsclient.UpdateProductBrandResp{}, nil
}
