package productbrandservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/pkg/pointerprocess"
	"github.com/feihua/zero-admin/pkg/time_util"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"
	"github.com/feihua/zero-admin/rpc/pms/internal/svc"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
)

// QueryProductBrandListLogic 查询商品品牌列表
/*
Author: LiuFeiHua
Date: 2025/05/26 10:33:54
*/
type QueryProductBrandListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryProductBrandListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryProductBrandListLogic {
	return &QueryProductBrandListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryProductBrandList 查询商品品牌列表
func (l *QueryProductBrandListLogic) QueryProductBrandList(in *pmsclient.QueryProductBrandListReq) (*pmsclient.QueryProductBrandListResp, error) {
	productBrand := query.PmsProductBrand
	q := productBrand.WithContext(l.ctx)
	if len(in.Name) > 0 {
		q = q.Where(productBrand.Name.Like("%" + in.Name + "%"))
	}

	if in.RecommendStatus != 2 {
		q = q.Where(productBrand.RecommendStatus.Eq(in.RecommendStatus))
	}
	if in.IsEnabled != 2 {
		q = q.Where(productBrand.IsEnabled.Eq(in.IsEnabled))
	}

	result, count, err := q.FindByPage(int((in.PageNum-1)*in.PageSize), int(in.PageSize))

	if err != nil {
		logc.Errorf(l.ctx, "查询商品品牌列表失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("查询商品品牌列表失败")
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
		Total: count,
		List:  list,
	}, nil
}
