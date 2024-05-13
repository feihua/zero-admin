package homebrand

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// HomeBrandAddLogic 首页品牌信息
/*
Author: LiuFeiHua
Date: 2024/5/13 15:53
*/
type HomeBrandAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHomeBrandAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) HomeBrandAddLogic {
	return HomeBrandAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// HomeBrandAdd 添加首页品牌信息
func (l *HomeBrandAddLogic) HomeBrandAdd(req types.AddHomeBrandReq) (*types.AddHomeBrandResp, error) {
	brandListResp, _ := l.svcCtx.BrandService.BrandListByIds(l.ctx, &pmsclient.BrandListByIdsReq{Ids: req.BrandIds})

	var list []*smsclient.HomeBrandAddData

	for _, item := range brandListResp.List {
		list = append(list, &smsclient.HomeBrandAddData{
			BrandId:         item.Id,
			BrandName:       item.Name,
			RecommendStatus: item.ShowStatus,
			Sort:            item.Sort,
		})
	}

	_, err := l.svcCtx.HomeBrandService.HomeBrandAdd(l.ctx, &smsclient.HomeBrandAddReq{
		BrandAddData: list,
	})

	if err != nil {
		logc.Errorf(l.ctx, "添加首页品牌信息失败,参数：%+v,响应：%s", req, err.Error())
		return nil, errorx.NewDefaultError("添加首页品牌失败")
	}

	return &types.AddHomeBrandResp{
		Code:    "000000",
		Message: "添加首页品牌成功",
	}, nil
}
