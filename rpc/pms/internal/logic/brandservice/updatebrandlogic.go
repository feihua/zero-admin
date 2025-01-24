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
	q := query.PmsBrand.WithContext(l.ctx)

	// 1.根据品牌id查询品牌是否已存在
	_, err := q.Where(query.PmsBrand.ID.Eq(in.Id)).First()

	if err != nil {
		logc.Errorf(l.ctx, "根据品牌id：%d,查询品牌失败,异常:%s", in.Id, err.Error())
		return nil, errors.New(fmt.Sprintf("查询品牌失败"))
	}

	item := &model.PmsBrand{
		ID:                  in.Id,                  //
		Name:                in.Name,                // 品牌名称
		FirstLetter:         in.FirstLetter,         // 首字母
		Sort:                in.Sort,                // 排序
		FactoryStatus:       in.FactoryStatus,       // 是否为品牌制造商：0->不是；1->是
		ShowStatus:          in.ShowStatus,          // 品牌显示状态
		RecommendStatus:     in.RecommendStatus,     // 推荐状态
		ProductCount:        in.ProductCount,        // 产品数量
		ProductCommentCount: in.ProductCommentCount, // 产品评论数量
		Logo:                in.Logo,                // 品牌logo
		BigPic:              in.BigPic,              // 专区大图
		BrandStory:          in.BrandStory,          // 品牌故事
		UpdateBy:            in.UpdateBy,            // 更新者
	}

	// 2.品牌存在时,则直接更新品牌
	_, err = q.Updates(item)

	if err != nil {
		logc.Errorf(l.ctx, "更新品牌失败,参数:%+v,异常:%s", item, err.Error())
		return nil, errors.New("更新品牌失败")
	}

	logc.Infof(l.ctx, "更新品牌成功,参数：%+v", in)
	return &pmsclient.UpdateBrandResp{}, nil
}
