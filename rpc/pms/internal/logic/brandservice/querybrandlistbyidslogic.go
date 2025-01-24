package brandservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/pkg/time_util"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/pms/internal/svc"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryBrandListByIdsLogic 根据ids查询品牌信息
/*
Author: LiuFeiHua
Date: 2024/6/12 16:29
*/
type QueryBrandListByIdsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryBrandListByIdsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryBrandListByIdsLogic {
	return &QueryBrandListByIdsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryBrandListByIds 根据ids查询品牌信息
func (l *QueryBrandListByIdsLogic) QueryBrandListByIds(in *pmsclient.QueryBrandListByIdsReq) (*pmsclient.QueryBrandListResp, error) {
	result, err := query.PmsBrand.WithContext(l.ctx).Where(query.PmsBrand.ID.In(in.Ids...)).Find()

	if err != nil {
		logc.Errorf(l.ctx, "查询商品品牌列表信息失败,参数：%+v,异常:%s", in, err.Error())
		return nil, err
	}

	var list []*pmsclient.BrandListData
	for _, item := range result {

		list = append(list, &pmsclient.BrandListData{
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
		})
	}

	return &pmsclient.QueryBrandListResp{
		Total: 0,
		List:  list,
	}, nil
}
