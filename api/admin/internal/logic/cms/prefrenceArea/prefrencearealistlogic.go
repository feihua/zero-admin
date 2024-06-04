package prefrenceArea

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/cms/cmsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"strings"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// PrefrenceAreaListLogic 商品优选
/*
Author: LiuFeiHua
Date: 2024/5/11 17:18
*/
type PrefrenceAreaListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPrefrenceAreaListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PrefrenceAreaListLogic {
	return &PrefrenceAreaListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// PrefrenceAreaList 查询商品优选
func (l *PrefrenceAreaListLogic) PrefrenceAreaList(req *types.ListPrefrenceAreaReq) (*types.ListPrefrenceAreaResp, error) {
	resp, err := l.svcCtx.PreferredAreaService.PreferredAreaList(l.ctx, &cmsclient.PreferredAreaListReq{
		Current:    req.Current,
		PageSize:   req.PageSize,
		Name:       strings.TrimSpace(req.Name),
		SubTitle:   strings.TrimSpace(req.SubTitle),
		ShowStatus: req.ShowStatus,
	})

	if err != nil {
		logc.Errorf(l.ctx, "参数: %+v,查询商品优选列表异常:%s", req, err.Error())
		return nil, errorx.NewDefaultError("查询商品优选失败")
	}

	var list []*types.ListPrefrenceAreaData

	for _, item := range resp.List {
		list = append(list, &types.ListPrefrenceAreaData{
			Id:         item.Id,
			Name:       item.Name,
			SubTitle:   item.SubTitle,
			Pic:        item.Pic,
			Sort:       item.Sort,
			ShowStatus: item.ShowStatus,
			CreateBy:   item.CreateBy,
			CreateTime: item.CreateTime,
			UpdateBy:   item.UpdateBy,
			UpdateTime: item.UpdateTime,
		})
	}

	return &types.ListPrefrenceAreaResp{
		Current:  req.Current,
		Data:     list,
		PageSize: req.PageSize,
		Success:  true,
		Total:    resp.Total,
		Code:     "000000",
		Message:  "查询商品优选成功",
	}, nil
}
