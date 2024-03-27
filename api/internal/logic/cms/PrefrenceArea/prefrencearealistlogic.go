package PrefrenceArea

import (
	"context"
	"encoding/json"
	"github.com/feihua/zero-admin/api/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/cms/cmsclient"

	"github.com/feihua/zero-admin/api/internal/svc"
	"github.com/feihua/zero-admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

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

func (l *PrefrenceAreaListLogic) PrefrenceAreaList(req *types.ListPrefrenceAreaReq) (resp *types.ListPrefrenceAreaResp, err error) {
	prefrenceAreaList, err := l.svcCtx.PrefrenceAreaService.PrefrenceAreaList(l.ctx, &cmsclient.PrefrenceAreaListReq{
		Current:    req.Current,
		PageSize:   req.PageSize,
		Name:       req.Name,
		SubTitle:   req.SubTitle,
		ShowStatus: req.ShowStatus,
	})

	if err != nil {
		data, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("参数: %s,查询商品优选列表异常:%s", string(data), err.Error())
		return nil, errorx.NewDefaultError("查询商品优选失败")
	}

	var list []*types.ListtPrefrenceAreaData

	for _, item := range prefrenceAreaList.List {
		list = append(list, &types.ListtPrefrenceAreaData{
			Id:         item.Id,
			Name:       item.Name,
			SubTitle:   item.SubTitle,
			Pic:        item.Pic,
			Sort:       item.Sort,
			ShowStatus: item.ShowStatus,
		})
	}

	return &types.ListPrefrenceAreaResp{
		Current:  req.Current,
		Data:     list,
		PageSize: req.PageSize,
		Success:  true,
		Total:    prefrenceAreaList.Total,
		Code:     "000000",
		Message:  "查询商品优选成功",
	}, nil
}
