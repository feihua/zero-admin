package logic

import (
	"context"
	"encoding/json"
	"github.com/feihua/zero-admin/api/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"

	"github.com/feihua/zero-admin/api/internal/svc"
	"github.com/feihua/zero-admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

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
		reqStr, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("添加首页品牌信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, errorx.NewDefaultError("添加首页品牌失败")
	}

	return &types.AddHomeBrandResp{
		Code:    "000000",
		Message: "添加首页品牌成功",
	}, nil
}
