package logic

import (
	"context"
	"zero-admin/api/internal/common/errorx"
	"zero-admin/rpc/sys/sysclient"

	"zero-admin/api/internal/svc"
	"zero-admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// DictDeleteLogic
/*
Author: LiuFeiHua
Date: 2023/12/18 17:17
*/
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

// DictDelete 删除字典信息
func (l *DictDeleteLogic) DictDelete(req types.DeleteDictReq) (*types.DeleteDictResp, error) {
	_, err := l.svcCtx.DictService.DictDelete(l.ctx, &sysclient.DictDeleteReq{
		Ids: req.Ids,
	})

	if err != nil {
		logx.WithContext(l.ctx).Errorf("根据dictId: %d,删除字典异常:%s", req.Ids, err.Error())
		return nil, errorx.NewDefaultError("删除字典失败")
	}

	return &types.DeleteDictResp{
		Code:    "000000",
		Message: "删除字典成功",
	}, nil
}
