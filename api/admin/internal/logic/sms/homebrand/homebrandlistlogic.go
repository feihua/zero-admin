package homebrand

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"strings"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// HomeBrandListLogic 首页品牌信息
/*
Author: LiuFeiHua
Date: 2024/5/13 15:53
*/
type HomeBrandListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHomeBrandListLogic(ctx context.Context, svcCtx *svc.ServiceContext) HomeBrandListLogic {
	return HomeBrandListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// HomeBrandList 查询首页品牌信息
func (l *HomeBrandListLogic) HomeBrandList(req types.ListHomeBrandReq) (*types.ListHomeBrandResp, error) {
	resp, err := l.svcCtx.HomeBrandService.QueryHomeBrandList(l.ctx, &smsclient.QueryHomeBrandListReq{
		BrandName:       strings.TrimSpace(req.BrandName),
		RecommendStatus: req.RecommendStatus,
		PageNum:         req.Current,
		PageSize:        req.PageSize,
	})

	if err != nil {
		logc.Errorf(l.ctx, "参数: %+v,查询首页品牌列表异常:%s", req, err.Error())
		return nil, errorx.NewDefaultError("查询首页品牌失败")
	}

	var list []*types.ListHomeBrandData

	for _, item := range resp.List {
		list = append(list, &types.ListHomeBrandData{
			Id:              item.Id,              //
			BrandId:         item.BrandId,         // 商品品牌id
			BrandName:       item.BrandName,       // 商品品牌名称
			RecommendStatus: item.RecommendStatus, // 推荐状态：0->不推荐;1->推荐
			Sort:            item.Sort,            // 排序
		})
	}

	return &types.ListHomeBrandResp{
		Current:  req.Current,
		Data:     list,
		PageSize: req.PageSize,
		Success:  true,
		Total:    resp.Total,
		Code:     "000000",
		Message:  "查询首页品牌成功",
	}, nil
}
