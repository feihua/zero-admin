package productspuservicelogic

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

// QueryProductSpuListLogic 查询商品SPU列表
/*
Author: LiuFeiHua
Date: 2025/06/16 14:37:38
*/
type QueryProductSpuListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryProductSpuListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryProductSpuListLogic {
	return &QueryProductSpuListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryProductSpuList 查询商品SPU列表
func (l *QueryProductSpuListLogic) QueryProductSpuList(in *pmsclient.QueryProductSpuListReq) (*pmsclient.QueryProductSpuListResp, error) {
	productSpu := query.PmsProductSpu
	q := productSpu.WithContext(l.ctx)
	if len(in.Name) > 0 {
		q = q.Where(productSpu.Name.Like("%" + in.Name + "%"))
	}
	if in.CategoryId != 0 {
		q = q.Where(productSpu.CategoryID.Eq(in.CategoryId))
	}
	if len(in.CategoryIds) > 0 {
		q = q.Where(productSpu.CategoryIds.Like("%" + in.CategoryIds + "%"))
	}
	if len(in.CategoryName) > 0 {
		q = q.Where(productSpu.CategoryName.Like("%" + in.CategoryName + "%"))
	}
	if in.BrandId != 0 {
		q = q.Where(productSpu.BrandID.Eq(in.BrandId))
	}
	if len(in.BrandName) > 0 {
		q = q.Where(productSpu.BrandName.Like("%" + in.BrandName + "%"))
	}

	if len(in.Keywords) > 0 {
		q = q.Where(productSpu.Keywords.Like("%" + in.Keywords + "%"))
	}

	if in.PublishStatus != 2 {
		q = q.Where(productSpu.PublishStatus.Eq(in.PublishStatus))
	}
	if in.NewStatus != 2 {
		q = q.Where(productSpu.NewStatus.Eq(in.NewStatus))
	}
	if in.RecommendStatus != 2 {
		q = q.Where(productSpu.RecommendStatus.Eq(in.RecommendStatus))
	}
	if in.VerifyStatus != 2 {
		q = q.Where(productSpu.VerifyStatus.Eq(in.VerifyStatus))
	}
	if in.PreviewStatus != 2 {
		q = q.Where(productSpu.PreviewStatus.Eq(in.PreviewStatus))
	}

	if in.PromotionType != 6 {
		q = q.Where(productSpu.PromotionType.Eq(in.PromotionType))
	}

	result, count, err := q.FindByPage(int((in.PageNum-1)*in.PageSize), int(in.PageSize))

	if err != nil {
		logc.Errorf(l.ctx, "查询商品SPU列表失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("查询商品SPU列表失败")
	}

	var list []*pmsclient.ProductSpuListData

	for _, item := range result {
		list = append(list, &pmsclient.ProductSpuListData{
			Id:                  item.ID,                                          // 商品SpuId
			ProductSn:           item.ProductSn,                                   // 商品货号
			Name:                item.Name,                                        // 商品名称
			CategoryId:          item.CategoryID,                                  // 商品分类ID
			CategoryIds:         item.CategoryIds,                                 // 商品分类ID集合
			CategoryName:        item.CategoryName,                                // 商品分类名称
			BrandId:             item.BrandID,                                     // 品牌ID
			BrandName:           item.BrandName,                                   // 品牌名称
			Unit:                item.Unit,                                        // 单位
			Weight:              float32(item.Weight),                             // 重量(kg)
			Keywords:            item.Keywords,                                    // 关键词
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
			SubTitle:            item.SubTitle,                                    // 副标题
			DetailHtml:          item.DetailHTML,                                  // 产品详情网页内容
			DetailMobileHtml:    item.DetailMobileHTML,                            // 移动端网页详情
			CreateBy:            item.CreateBy,                                    // 创建人ID
			CreateTime:          time_util.TimeToStr(item.CreateTime),             // 创建时间
			UpdateBy:            pointerprocess.DefaltData(item.UpdateBy).(int64), // 更新人ID
			UpdateTime:          time_util.TimeToString(item.UpdateTime),          // 更新时间

		})
	}

	return &pmsclient.QueryProductSpuListResp{
		Total: count,
		List:  list,
	}, nil
}
