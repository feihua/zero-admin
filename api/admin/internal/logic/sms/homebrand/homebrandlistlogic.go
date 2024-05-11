package logic

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

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

func (l *HomeBrandListLogic) HomeBrandList(req types.ListHomeBrandReq) (*types.ListHomeBrandResp, error) {
	resp, err := l.svcCtx.HomeBrandService.HomeBrandList(l.ctx, &smsclient.HomeBrandListReq{
		BrandName:       req.BrandName,
		RecommendStatus: req.RecommendStatus,
		Current:         req.Current,
		PageSize:        req.PageSize,
	})

	if err != nil {
		logc.Errorf(l.ctx, "参数: %+v,查询首页品牌列表异常:%s", req, err.Error())
		return nil, errorx.NewDefaultError("查询首页品牌失败")
	}

	var list []*types.ListtHomeBrandData

	for _, item := range resp.List {
		list = append(list, &types.ListtHomeBrandData{
			Id:              item.Id,
			BrandId:         item.BrandId,
			BrandName:       item.BrandName,
			RecommendStatus: item.RecommendStatus,
			Sort:            item.Sort,
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
