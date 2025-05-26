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
	"gorm.io/gorm"
)

// QueryProductBrandDetailLogic 查询商品品牌详情
/*
Author: LiuFeiHua
Date: 2025/05/26 10:33:54
*/
type QueryProductBrandDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryProductBrandDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryProductBrandDetailLogic {
	return &QueryProductBrandDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryProductBrandDetail 查询商品品牌详情
func (l *QueryProductBrandDetailLogic) QueryProductBrandDetail(in *pmsclient.QueryProductBrandDetailReq) (*pmsclient.QueryProductBrandDetailResp, error) {
	item, err := query.PmsProductBrand.WithContext(l.ctx).Where(query.PmsProductBrand.ID.Eq(in.Id)).First()

	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		logc.Errorf(l.ctx, "商品品牌不存在, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errors.New("商品品牌不存在")
	case err != nil:
		logc.Errorf(l.ctx, "查询商品品牌异常, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errors.New("查询商品品牌异常")
	}

	data := &pmsclient.QueryProductBrandDetailResp{
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
	}

	return data, nil
}
