package productservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/pkg/time_util"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/pms/internal/svc"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryProductListByIdsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryProductListByIdsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryProductListByIdsLogic {
	return &QueryProductListByIdsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *QueryProductListByIdsLogic) QueryProductListByIds(in *pmsclient.QueryProductByIdsReq) (*pmsclient.QueryProductListResp, error) {
	q := query.PmsProduct
	result, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Find()

	if err != nil {
		logc.Errorf(l.ctx, "查询商品列表信息失败,参数：%+v,异常:%s", in, err.Error())
		return nil, err
	}

	var list []*pmsclient.ProductListData
	for _, item := range result {

		list = append(list, &pmsclient.ProductListData{
			Id:                         item.ID,                                      //
			BrandId:                    item.BrandID,                                 // 品牌id
			ProductCategoryId:          item.ProductCategoryID,                       // 商品分类id
			FeightTemplateId:           item.FeightTemplateID,                        // 商品运费模板id
			ProductAttributeCategoryId: item.ProductAttributeCategoryID,              // 商品属性分类id
			Name:                       item.Name,                                    // 商品名称
			Pic:                        item.Pic,                                     // 商品图片
			ProductSn:                  item.ProductSn,                               // 货号
			DeleteStatus:               item.DeleteStatus,                            // 删除状态：0->未删除；1->已删除
			PublishStatus:              item.PublishStatus,                           // 上架状态：0->下架；1->上架
			NewStatus:                  item.NewStatus,                               // 新品状态:0->不是新品；1->新品
			RecommandStatus:            item.RecommandStatus,                         // 推荐状态；0->不推荐；1->推荐
			VerifyStatus:               item.VerifyStatus,                            // 审核状态：0->未审核；1->审核通过
			Sort:                       item.Sort,                                    // 排序
			Sale:                       item.Sale,                                    // 销量
			Price:                      item.Price,                                   // 商品价格
			PromotionPrice:             item.PromotionPrice,                          // 促销价格
			GiftGrowth:                 item.GiftGrowth,                              // 赠送的成长值
			GiftPoint:                  item.GiftPoint,                               // 赠送的积分
			UsePointLimit:              item.UsePointLimit,                           // 限制使用的积分数
			SubTitle:                   item.SubTitle,                                // 副标题
			Description:                item.Description,                             // 商品描述
			OriginalPrice:              item.OriginalPrice,                           // 市场价
			Stock:                      item.Stock,                                   // 库存
			LowStock:                   item.LowStock,                                // 库存预警值
			Unit:                       item.Unit,                                    // 单位
			Weight:                     item.Weight,                                  // 商品重量，默认为克
			PreviewStatus:              item.PreviewStatus,                           // 是否为预告商品：0->不是；1->是
			ServiceIds:                 item.ServiceIds,                              // 以逗号分割的产品服务：1->无忧退货；2->快速退款；3->免费包邮
			Keywords:                   item.Keywords,                                // 搜索关键字
			Note:                       item.Note,                                    // 备注
			AlbumPics:                  item.AlbumPics,                               // 画册图片，连产品图片限制为5张，以逗号分割
			DetailTitle:                item.DetailTitle,                             // 详情标题
			DetailDesc:                 item.DetailDesc,                              // 详情描述
			DetailHtml:                 item.DetailHTML,                              // 产品详情网页内容
			DetailMobileHtml:           item.DetailMobileHTML,                        // 移动端网页详情
			PromotionStartTime:         time_util.TimeToStr(item.PromotionStartTime), // 促销开始时间
			PromotionEndTime:           time_util.TimeToStr(item.PromotionEndTime),   // 促销结束时间
			PromotionPerLimit:          item.PromotionPerLimit,                       // 活动限购数量
			PromotionType:              item.PromotionType,                           // 促销类型：0->没有促销使用原价;1->使用促销价；2->使用会员价；3->使用阶梯价格；4->使用满减价格；5->限时购
			BrandName:                  item.BrandName,                               // 品牌名称
			ProductCategoryName:        item.ProductCategoryName,                     // 商品分类名称
			ProductCategoryIdArray:     item.ProductCategoryIDArray,                  // 商品分类id字符串
		})
	}

	logc.Infof(l.ctx, "查询商品列表信息,参数：%+v,响应：%+v", in, list)
	return &pmsclient.QueryProductListResp{
		Total: 0,
		List:  list,
	}, nil

}
