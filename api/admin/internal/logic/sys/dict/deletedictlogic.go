package dict

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// DeleteDictLogic
/*
Author: LiuFeiHua
Date: 2023/12/18 17:17
*/
type DeleteDictLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteDictLogic(ctx context.Context, svcCtx *svc.ServiceContext) DeleteDictLogic {
	return DeleteDictLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// DeleteDict 删除字典信息
func (l *DeleteDictLogic) DeleteDict(req *types.DeleteDictReq) (*types.DeleteDictResp, error) {
	_, err := l.svcCtx.DictService.DictDelete(l.ctx, &sysclient.DictDeleteReq{
		Ids: req.Ids,
	})

	if err != nil {
		logc.Errorf(l.ctx, "根据dictId: %+v,删除字典异常:%s", req, err.Error())
		return nil, errorx.NewDefaultError("删除字典失败")
	}

	return &types.DeleteDictResp{
		Code:    "000000",
		Message: "删除字典成功",
	}, nil
}
