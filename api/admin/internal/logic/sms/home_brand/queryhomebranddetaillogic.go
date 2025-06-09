package home_brand

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryHomeBrandDetailLogic 查询首页推荐品牌表详情
/*
Author: 刘飞华
Date: 2025/02/07 10:11:43
*/
type QueryHomeBrandDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryHomeBrandDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryHomeBrandDetailLogic {
	return &QueryHomeBrandDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// QueryHomeBrandDetail 查询首页推荐品牌表详情
func (l *QueryHomeBrandDetailLogic) QueryHomeBrandDetail(req *types.QueryHomeBrandDetailReq) (resp *types.QueryHomeBrandDetailResp, err error) {

	detail, err := l.svcCtx.ProductBrandService.QueryProductBrandDetail(l.ctx, &pmsclient.QueryProductBrandDetailReq{
		Id: req.Id,
	})

	if err != nil {
		logc.Errorf(l.ctx, "查询商品品牌详情失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	data := types.QueryHomeBrandDetailData{
		Id:                  detail.Id,                  //
		Name:                detail.Name,                // 品牌名称
		Logo:                detail.Logo,                // 品牌logo
		BigPic:              detail.BigPic,              // 专区大图
		Description:         detail.Description,         // 描述
		FirstLetter:         detail.FirstLetter,         // 首字母
		Sort:                detail.Sort,                // 排序
		RecommendStatus:     detail.RecommendStatus,     // 推荐状态
		ProductCount:        detail.ProductCount,        // 产品数量
		ProductCommentCount: detail.ProductCommentCount, // 产品评论数量
		IsEnabled:           detail.IsEnabled,           // 是否启用
		CreateBy:            detail.CreateBy,            // 创建人ID
		CreateTime:          detail.CreateTime,          // 创建时间
		UpdateBy:            detail.UpdateBy,            // 更新人ID
		UpdateTime:          detail.UpdateTime,          // 更新时间

	}

	return &types.QueryHomeBrandDetailResp{
		Code:    "000000",
		Message: "查询首页推荐品牌表成功",
		Data:    data,
	}, nil
}
