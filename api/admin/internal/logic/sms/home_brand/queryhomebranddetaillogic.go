package home_brand

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"
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

	detail, err := l.svcCtx.HomeBrandService.QueryHomeBrandDetail(l.ctx, &smsclient.QueryHomeBrandDetailReq{
		Id: req.Id,
	})

	if err != nil {
		logc.Errorf(l.ctx, "查询首页推荐品牌表详情失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	data := types.QueryHomeBrandDetailData{
		Id:              detail.Id,              //
		BrandId:         detail.BrandId,         // 商品品牌id
		BrandName:       detail.BrandName,       // 商品品牌名称
		RecommendStatus: detail.RecommendStatus, // 推荐状态：0->不推荐;1->推荐
		Sort:            detail.Sort,            // 排序
	}
	return &types.QueryHomeBrandDetailResp{
		Code:    "000000",
		Message: "查询首页推荐品牌表成功",
		Data:    data,
	}, nil
}
