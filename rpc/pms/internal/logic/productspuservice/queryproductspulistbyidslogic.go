package productspuservicelogic

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

type QueryProductSpuListByIdsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryProductSpuListByIdsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryProductSpuListByIdsLogic {
	return &QueryProductSpuListByIdsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryProductSpuListByIds 根据id集合查询商品信息
func (l *QueryProductSpuListByIdsLogic) QueryProductSpuListByIds(in *pmsclient.QueryProductSpuByIdsReq) (*pmsclient.QueryProductSpuListResp, error) {
	q := query.PmsProductSpu
	result, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Find()

	if err != nil {
		logc.Errorf(l.ctx, "查询商品列表失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("查询商品列表失败")
	}

	var list []*pmsclient.ProductSpuListData

	for _, item := range result {
		list = append(list, &pmsclient.ProductSpuListData{
			Id:                  item.ID,                                          // 商品SpuId
			Name:                item.Name,                                        // 商品名称
			CategoryId:          item.CategoryID,                                  // 商品分类ID
			CategoryIds:         item.CategoryIds,                                 // 商品分类ID集合
			CategoryName:        item.CategoryName,                                // 商品分类名称
			BrandId:             item.BrandID,                                     // 品牌ID
			BrandName:           item.BrandName,                                   // 品牌名称
			Unit:                item.Unit,                                        // 单位
			Weight:              float32(item.Weight),                             // 重量(kg)
			Keywords:            item.Keywords,                                    // 关键词
			Brief:               item.Brief,                                       // 简介
			Description:         item.Description,                                 // 详细描述
			AlbumPics:           item.AlbumPics,                                   // 画册图片，最多8张，以逗号分割
			MainPic:             item.MainPic,                                     // 主图
			PriceRange:          item.PriceRange,                                  // 价格区间
			PublishStatus:       item.PublishStatus,                               // 上架状态：0-下架，1-上架
			NewStatus:           item.NewStatus,                                   // 新品状态:0->不是新品；1->新品
			RecommendStatus:     item.RecommendStatus,                             // 推荐状态；0->不推荐；1->推荐
			VerifyStatus:        item.VerifyStatus,                                // 审核状态：0->未审核；1->审核通过
			PreviewStatus:       item.PreviewStatus,                               // 是否为预告商品：0->不是；1->是
			Sort:                item.Sort,                                        // 排序
			NewStatusSort:       item.NewStatusSort,                               // 新品排序
			RecommendStatusSort: item.RecommendStatusSort,                         // 推荐排序
			Sales:               item.Sales,                                       // 销量
			Stock:               item.Stock,                                       // 库存
			LowStock:            item.LowStock,                                    // 预警库存
			PromotionType:       item.PromotionType,                               // 促销类型：0->没有促销使用原价;1->使用促销价；2->使用会员价；3->使用阶梯价格；4->使用满减价格；5->秒杀
			DetailTitle:         item.DetailTitle,                                 // 详情标题
			DetailDesc:          item.DetailDesc,                                  // 详情描述
			DetailHtml:          item.DetailHTML,                                  // 产品详情网页内容
			DetailMobileHtml:    item.DetailMobileHTML,                            // 移动端网页详情
			CreateBy:            item.CreateBy,                                    // 创建人ID
			CreateTime:          time_util.TimeToStr(item.CreateTime),             // 创建时间
			UpdateBy:            pointerprocess.DefaltData(item.UpdateBy).(int64), // 更新人ID
			UpdateTime:          time_util.TimeToString(item.UpdateTime),          // 更新时间

		})
	}

	return &pmsclient.QueryProductSpuListResp{
		Total: 0,
		List:  list,
	}, nil

}
