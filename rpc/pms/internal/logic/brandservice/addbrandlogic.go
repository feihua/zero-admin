package brandservicelogic

import (
	"context"
	"errors"
	"fmt"
	"github.com/feihua/zero-admin/rpc/pms/gen/model"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/pms/internal/svc"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// AddBrandLogic 添加品牌
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

// AddBrand 添加品牌
func (l *AddBrandLogic) AddBrand(in *pmsclient.AddBrandReq) (*pmsclient.AddBrandResp, error) {
	q := query.PmsBrand
	name := in.Name
	count, err := q.WithContext(l.ctx).Where(q.Name.Eq(name)).Count()

	if err != nil {
		logc.Errorf(l.ctx, "根据品牌名称：%s,查询品牌失败,异常:%s", name, err.Error())
		return nil, errors.New(fmt.Sprintf("添加品牌失败"))
	}

	if count > 0 {
		return nil, errors.New(fmt.Sprintf("品牌名称：%s,已存在", name))
	}

	brand := model.PmsBrand{
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
		CreateBy:            in.CreateBy,            // 创建者
	}
	err = q.WithContext(l.ctx).Create(&brand)

	if err != nil {
		logc.Errorf(l.ctx, "添加品牌失败失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("添加品牌失败失败")
	}

	return &pmsclient.AddBrandResp{
		BrandId: brand.ID,
	}, nil
}
