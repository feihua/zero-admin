package productbrandservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/pkg/pointerprocess"
	"github.com/feihua/zero-admin/pkg/time_util"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/pms/internal/svc"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

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

func (l *QueryBrandListByIdsLogic) QueryBrandListByIds(in *pmsclient.QueryBrandListByIdsReq) (*pmsclient.QueryProductBrandListResp, error) {
	result, err := query.PmsProductBrand.WithContext(l.ctx).Where(query.PmsProductBrand.ID.In(in.Ids...)).Find()

	if err != nil {
		logc.Errorf(l.ctx, "根据ids查询品牌信息失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("根据ids查询品牌信息失败")
	}

	var list []*pmsclient.ProductBrandListData
	for _, item := range result {

		list = append(list, &pmsclient.ProductBrandListData{
			Id:                  item.ID,                                          //
			Name:                item.Name,                                        // 品牌名称
			Logo:                item.Logo,                                        // 品牌logo
			BigPic:              item.BigPic,                                      // 专区大图
			Description:         item.Description,                                 // 描述
			FirstLetter:         item.FirstLetter,                                 // 首字母
			Sort:                item.Sort,                                        // 排序
			RecommendStatus:     item.RecommendStatus,                             // 推荐状态
			ProductCount:        item.ProductCount,                                // 产品数量
			ProductCommentCount: item.ProductCommentCount,                         // 产品评论数量
			IsEnabled:           item.IsEnabled,                                   // 是否启用
			CreateBy:            item.CreateBy,                                    // 创建人ID
			CreateTime:          time_util.TimeToStr(item.CreateTime),             // 创建时间
			UpdateBy:            pointerprocess.DefaltData(item.UpdateBy).(int64), // 更新人ID
			UpdateTime:          time_util.TimeToString(item.UpdateTime),          // 更新时间
		})
	}

	return &pmsclient.QueryProductBrandListResp{
		Total: 0,
		List:  list,
	}, nil

}
