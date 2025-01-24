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

type QueryBrandListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryBrandListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryBrandListLogic {
	return &QueryBrandListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询品牌表列表
func (l *QueryBrandListLogic) QueryBrandList(in *pmsclient.QueryBrandListReq) (*pmsclient.QueryBrandListResp, error) {
	q := query.PmsBrand.WithContext(l.ctx)
	if len(in.Name) > 0 {
		q = q.Where(query.PmsBrand.Name.Like("%" + in.Name + "%"))
	}
	if in.FactoryStatus != 2 {
		q = q.Where(query.PmsBrand.FactoryStatus.Eq(in.FactoryStatus))
	}
	if in.ShowStatus != 2 {
		q = q.Where(query.PmsBrand.ShowStatus.Eq(in.ShowStatus))
	}
	if in.RecommendStatus != 2 {
		q = q.Where(query.PmsBrand.RecommendStatus.Eq(in.RecommendStatus))
	}

	result, count, err := q.FindByPage(int((in.PageNum-1)*in.PageSize), int(in.PageSize))

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
		Total: count,
		List:  list,
	}, nil

}
