package brandservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/pkg/time_util"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"
	"github.com/feihua/zero-admin/rpc/pms/internal/svc"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
)

// QueryBrandDetailLogic 查询品牌详情
/*
Author: LiuFeiHua
Date: 2025/01/24 09:08:05
*/
type QueryBrandDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryBrandDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryBrandDetailLogic {
	return &QueryBrandDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryBrandDetail 查询品牌详情
func (l *QueryBrandDetailLogic) QueryBrandDetail(in *pmsclient.QueryBrandDetailReq) (*pmsclient.QueryBrandDetailResp, error) {
	item, err := query.PmsBrand.WithContext(l.ctx).Where(query.PmsBrand.ID.Eq(in.Id)).First()

	if err != nil {
		logc.Errorf(l.ctx, "查询品牌详情失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("查询品牌详情失败")
	}

	data := &pmsclient.QueryBrandDetailResp{
		Id:                  item.ID,                                 //
		Name:                item.Name,                               // 品牌名称
		FirstLetter:         item.FirstLetter,                        // 首字母
		Sort:                item.Sort,                               // 排序
		FactoryStatus:       item.FactoryStatus,                      // 是否为品牌制造商：0->不是；1->是
		ShowStatus:          item.ShowStatus,                         // 品牌显示状态
		RecommendStatus:     item.RecommendStatus,                    // 推荐状态
		ProductCount:        item.ProductCount,                       // 产品数量
		ProductCommentCount: item.ProductCommentCount,                // 产品评论数量
		Logo:                item.Logo,                               // 品牌logo
		BigPic:              item.BigPic,                             // 专区大图
		BrandStory:          item.BrandStory,                         // 品牌故事
		CreateBy:            item.CreateBy,                           // 创建者
		CreateTime:          time_util.TimeToStr(item.CreateTime),    // 创建时间
		UpdateBy:            item.UpdateBy,                           // 更新者
		UpdateTime:          time_util.TimeToString(item.UpdateTime), // 更新时间
	}

	return data, nil
}
