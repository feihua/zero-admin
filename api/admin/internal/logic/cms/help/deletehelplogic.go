package help

import (
	"context"

	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"
	"github.com/feihua/zero-admin/rpc/cms/cmsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"

	"github.com/zeromicro/go-zero/core/logx"
)

// DeleteHelpLogic 删除帮助
/*
Author: 刘飞华
Date: 2026/08/27 10:50:29
*/
type DeleteHelpLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteHelpLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteHelpLogic {
	return &DeleteHelpLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// DeleteHelp 删除帮助
func (l *DeleteHelpLogic) DeleteHelp(req *types.DeleteHelpReq) (resp *types.DeleteHelpResp, err error) {
	_, err = l.svcCtx.HelpService.DeleteHelp(l.ctx, &cmsclient.DeleteHelpReq{
		Ids: req.Ids,
	})

	if err != nil {
		logc.Errorf(l.ctx, "删除帮助失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return &types.DeleteHelpResp{
		Code:    "000000",
		Message: "删除帮助成功",
	}, nil
}
