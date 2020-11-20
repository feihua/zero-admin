package logic

import (
	"context"
	"go-zero-admin/api/internal/svc"
	"go-zero-admin/api/internal/types"
	"go-zero-admin/rpc/sys/sysclient"

	"github.com/tal-tech/go-zero/core/logx"
)

type DictListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDictListLogic(ctx context.Context, svcCtx *svc.ServiceContext) DictListLogic {
	return DictListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DictListLogic) DictList(req types.ListDictReq) (*types.ListDictResp, error) {
	resp, err := l.svcCtx.Sys.DictList(l.ctx, &sysclient.DictListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
	})

	if err != nil {
		return nil, err
	}

	var list []*types.ListDictData

	for _, dict := range resp.List {
		list = append(list, &types.ListDictData{
			dict.Id,
			dict.Value,
			dict.Label,
			dict.Type,
			dict.Description,
			0,
			dict.CreateBy,
			dict.CreateTime,
			dict.LastUpdateBy,
			dict.LastUpdateTime,
			dict.Remarks,
			dict.DelFlag,
		})
	}

	return &types.ListDictResp{
		Current:  req.Current,
		Data:     list,
		PageSize: req.PageSize,
		Success:  true,
		Total:    resp.Total,
	}, nil
}
