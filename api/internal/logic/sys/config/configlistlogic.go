package logic

import (
	"context"
	"zero-admin/api/internal/common/errorx"
	"zero-admin/api/internal/svc"
	"zero-admin/api/internal/types"
	"zero-admin/rpc/sys/sysclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type ConfigListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewConfigListLogic(ctx context.Context, svcCtx *svc.ServiceContext) ConfigListLogic {
	return ConfigListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ConfigListLogic) ConfigList(req types.ListConfigReq) (*types.ListConfigResp, error) {
	resp, err := l.svcCtx.Sys.ConfigList(l.ctx, &sysclient.ConfigListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
	})

	if err != nil {
		return nil, errorx.NewDefaultError("查询参数配置失败")
	}

	var list []*types.ListConfigData

	for _, config := range resp.List {
		list = append(list, &types.ListConfigData{
			Id:             config.Id,
			Value:          config.Value,
			Label:          config.Label,
			Type:           config.Type,
			Description:    config.Description,
			Sort:           config.Sort,
			CreateBy:       config.CreateBy,
			CreateTime:     config.CreateTime,
			LastUpdateBy:   config.LastUpdateBy,
			LastUpdateTime: config.LastUpdateTime,
			Remarks:        config.Remarks,
			DelFlag:        config.DelFlag,
		})
	}

	return &types.ListConfigResp{
		Current:  req.Current,
		Data:     list,
		PageSize: req.PageSize,
		Success:  true,
		Total:    resp.Total,
	}, nil
}
