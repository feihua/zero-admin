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

type QueryPreferredAreaListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryPreferredAreaListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryPreferredAreaListLogic {
	return &QueryPreferredAreaListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryPreferredAreaListLogic) QueryPreferredAreaList(req *types.QueryPreferredAreaListReq) (resp *types.QueryPreferredAreaListResp, err error) {
	result, err := l.svcCtx.PreferredAreaService.QueryPreferredAreaList(l.ctx, &cmsclient.QueryPreferredAreaListReq{
		PageNum:    req.Current,
		PageSize:   req.PageSize,
		Name:       strings.TrimSpace(req.Name),
		SubTitle:   strings.TrimSpace(req.SubTitle),
		ShowStatus: req.ShowStatus,
	})

	if err != nil {
		logc.Errorf(l.ctx, "参数: %+v,查询商品优选列表异常:%s", req, err.Error())
		return nil, errorx.NewDefaultError("查询商品优选失败")
	}

	var list []*types.QueryPreferredAreaListData

	for _, item := range result.List {
		list = append(list, &types.QueryPreferredAreaListData{
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

	return &types.QueryPreferredAreaListResp{
		Current:  req.Current,
		Data:     list,
		PageSize: req.PageSize,
		Success:  true,
		Total:    result.Total,
		Code:     "000000",
		Message:  "查询商品优选成功",
	}, nil
}
