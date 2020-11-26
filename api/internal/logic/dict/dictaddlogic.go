package logic

import (
	"context"
	"go-zero-admin/rpc/sys/sysclient"

	"go-zero-admin/api/internal/svc"
	"go-zero-admin/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type DictAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDictAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) DictAddLogic {
	return DictAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DictAddLogic) DictAdd(req types.AddDictReq) (*types.AddDictResp, error) {
	_, err := l.svcCtx.Sys.DictAdd(l.ctx, &sysclient.DictAddReq{
		Value:       req.Value,
		Label:       req.Label,
		Type:        req.Type,
		Description: req.Description,
		Sort:        int64(req.Sort),
		Remarks:     req.Remarks,
		//todo 从token里面拿
		CreateBy: "admin",
	})

	if err != nil {
		return nil, err
	}

	return &types.AddDictResp{}, nil
}
