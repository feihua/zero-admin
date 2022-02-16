package logic

import (
	"context"
	"zero-admin/api/internal/common/errorx"
	"zero-admin/rpc/sys/sysclient"

	"zero-admin/api/internal/svc"
	"zero-admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DictDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDictDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) DictDeleteLogic {
	return DictDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DictDeleteLogic) DictDelete(req types.DeleteDictReq) (*types.DeleteDictResp, error) {
	_, err := l.svcCtx.Sys.DictDelete(l.ctx, &sysclient.DictDeleteReq{
		Id: req.Id,
	})

	if err != nil {
		logx.WithContext(l.ctx).Errorf("根据dictId: %d,删除字典异常:%s", req.Id, err.Error())
		return nil, errorx.NewDefaultError("删除字典失败")
	}

	return &types.DeleteDictResp{
		Code:    "000000",
		Message: "删除字典成功",
	}, nil
}
