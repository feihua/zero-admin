package brandservicelogic

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

// UpdateBrandLogic 更新品牌
/*
Author: LiuFeiHua
Date: 2025/01/24 09:08:05
*/
type UpdateBrandLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateBrandLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateBrandLogic {
	return &UpdateBrandLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateBrand 更新品牌
func (l *UpdateBrandLogic) UpdateBrand(in *pmsclient.UpdateBrandReq) (*pmsclient.UpdateBrandResp, error) {
	q := query.PmsBrand

	brand, err := q.WithContext(l.ctx).Where(q.ID.Eq(in.Id)).First()

	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		return nil, errors.New("品牌不存在")
	case err != nil:
		logc.Errorf(l.ctx, "查询品牌异常, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errors.New("查询品牌异常")
	}

	name := in.Name
	count, err := q.WithContext(l.ctx).Where(q.Name.Eq(name), q.ID.Neq(in.Id)).Count()

	if err != nil {
		logc.Errorf(l.ctx, "根据品牌名称：%s,查询品牌失败,异常:%s", name, err.Error())
		return nil, errors.New(fmt.Sprintf("更新品牌失败"))
	}

	if count > 0 {
		return nil, errors.New(fmt.Sprintf("品牌名称：%s,已存在", name))
	}

	now := time.Now()
	item := &model.PmsBrand{
		ID:                  in.Id,                     //
		Name:                in.Name,                   // 品牌名称
		FirstLetter:         in.FirstLetter,            // 首字母
		Sort:                in.Sort,                   // 排序
		FactoryStatus:       in.FactoryStatus,          // 是否为品牌制造商：0->不是；1->是
		ShowStatus:          in.ShowStatus,             // 品牌显示状态
		RecommendStatus:     in.RecommendStatus,        // 推荐状态
		ProductCount:        brand.ProductCount,        // 产品数量
		ProductCommentCount: brand.ProductCommentCount, // 产品评论数量
		Logo:                in.Logo,                   // 品牌logo
		BigPic:              in.BigPic,                 // 专区大图
		BrandStory:          in.BrandStory,             // 品牌故事
		CreateBy:            brand.CreateBy,            // 创建者
		CreateTime:          brand.CreateTime,          // 创建时间
		UpdateBy:            in.UpdateBy,               // 更新者
		UpdateTime:          &now,                      // 更新时间
	}

	// 2.品牌存在时,则直接更新品牌
	err = l.svcCtx.DB.Model(&model.PmsBrand{}).WithContext(l.ctx).Where(query.PmsBrand.ID.Eq(in.Id)).Save(item).Error

	if err != nil {
		logc.Errorf(l.ctx, "更新品牌失败,参数:%+v,异常:%s", item, err.Error())
		return nil, errors.New("更新品牌失败")
	}

	return &pmsclient.UpdateBrandResp{}, nil
}
